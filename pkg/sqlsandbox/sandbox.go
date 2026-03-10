package sqlsandbox

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

// SchemaInfo represents the JSONB schema_info from a challenge.
type SchemaInfo struct {
	Tables []TableDef `json:"tables"`
}

// TableDef defines a table to create in the sandbox.
type TableDef struct {
	Name    string                   `json:"name"`
	Columns []ColumnDef              `json:"columns"`
	Rows    []map[string]interface{} `json:"rows"`
}

// ColumnDef defines a column in a table.
type ColumnDef struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// ExecuteResult holds the result of a SQL sandbox execution.
type ExecuteResult struct {
	Columns []string                 `json:"columns"`
	Rows    []map[string]interface{} `json:"rows"`
	Error   string                   `json:"error,omitempty"`
}

// Sandbox manages temporary PostgreSQL schema execution.
type Sandbox struct {
	db *sql.DB
}

// New creates a new Sandbox using an existing database connection.
func New(db *sql.DB) *Sandbox {
	return &Sandbox{db: db}
}

// Execute runs a user's SQL query inside a temporary schema with seed data.
// 1. Creates a temporary schema
// 2. Creates tables and inserts seed data from schema_info
// 3. Runs the user's query
// 4. Returns result rows
// 5. Drops the schema (cleanup)
func (s *Sandbox) Execute(schemaInfoJSON json.RawMessage, userQuery string) (*ExecuteResult, error) {
	// Generate unique schema name
	schemaName := "sandbox_" + strings.ReplaceAll(uuid.New().String()[:8], "-", "")

	// Create schema
	if _, err := s.db.Exec(fmt.Sprintf("CREATE SCHEMA %s", schemaName)); err != nil {
		return nil, fmt.Errorf("failed to create sandbox schema: %w", err)
	}

	// Ensure cleanup
	defer func() {
		s.db.Exec(fmt.Sprintf("DROP SCHEMA %s CASCADE", schemaName))
	}()

	// Set search_path to the sandbox schema
	if _, err := s.db.Exec(fmt.Sprintf("SET search_path TO %s", schemaName)); err != nil {
		return nil, fmt.Errorf("failed to set search_path: %w", err)
	}
	// Reset search_path on return
	defer func() {
		s.db.Exec("SET search_path TO public")
	}()

	// Parse schema_info
	if schemaInfoJSON != nil && len(schemaInfoJSON) > 0 && string(schemaInfoJSON) != "null" {
		var schema SchemaInfo
		if err := json.Unmarshal(schemaInfoJSON, &schema); err != nil {
			return &ExecuteResult{Error: "invalid schema_info: " + err.Error()}, nil
		}

		// Create tables and insert data
		for _, table := range schema.Tables {
			if err := s.createTable(schemaName, table); err != nil {
				return &ExecuteResult{Error: "failed to create table: " + err.Error()}, nil
			}
			if err := s.insertRows(schemaName, table); err != nil {
				return &ExecuteResult{Error: "failed to insert seed data: " + err.Error()}, nil
			}
		}
	}

	// Execute user query
	rows, err := s.db.Query(userQuery)
	if err != nil {
		return &ExecuteResult{Error: err.Error()}, nil
	}
	defer rows.Close()

	// Read columns
	columns, err := rows.Columns()
	if err != nil {
		return &ExecuteResult{Error: "failed to read columns: " + err.Error()}, nil
	}

	// Read rows
	var resultRows []map[string]interface{}
	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range values {
			valuePtrs[i] = &values[i]
		}
		if err := rows.Scan(valuePtrs...); err != nil {
			return &ExecuteResult{Error: "failed to scan row: " + err.Error()}, nil
		}

		// Convert []byte to string for JSON serialization
		row := make(map[string]interface{})
		for i, v := range values {
			colName := columns[i]
			if b, ok := v.([]byte); ok {
				row[colName] = string(b)
			} else {
				row[colName] = v
			}
		}
		resultRows = append(resultRows, row)
	}

	return &ExecuteResult{
		Columns: columns,
		Rows:    resultRows,
	}, nil
}

// createTable generates and executes a CREATE TABLE statement.
func (s *Sandbox) createTable(schema string, table TableDef) error {
	var colDefs []string
	for _, col := range table.Columns {
		colDefs = append(colDefs, fmt.Sprintf("%s %s", col.Name, col.Type))
	}
	query := fmt.Sprintf("CREATE TABLE %s.%s (%s)", schema, table.Name, strings.Join(colDefs, ", "))
	_, err := s.db.Exec(query)
	return err
}

// insertRows generates and executes INSERT statements for seed data.
func (s *Sandbox) insertRows(schema string, table TableDef) error {
	if len(table.Rows) == 0 || len(table.Columns) == 0 {
		return nil
	}

	// Build column names
	var colNames []string
	for _, col := range table.Columns {
		colNames = append(colNames, col.Name)
	}

	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	for _, row := range table.Rows {
		var placeholders []string
		var values []interface{}

		for i, col := range table.Columns {
			placeholders = append(placeholders, fmt.Sprintf("$%d", i+1))
			values = append(values, row[col.Name])
		}

		query := fmt.Sprintf(
			"INSERT INTO %s.%s (%s) VALUES (%s)",
			schema, table.Name,
			strings.Join(colNames, ", "),
			strings.Join(placeholders, ", "),
		)
		if _, err := tx.Exec(query, values...); err != nil {
			return fmt.Errorf("insert into %s failed: %w", table.Name, err)
		}
	}

	return tx.Commit()
}
