package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/arulkarim/ngodingyuk-server/pkg/gorunner"
	"github.com/arulkarim/ngodingyuk-server/pkg/sqlsandbox"
)

type ExecuteService struct {
	sqlDB *sql.DB // raw DB connection for SQL sandbox
}

func NewExecuteService(sqlDB *sql.DB) *ExecuteService {
	return &ExecuteService{sqlDB: sqlDB}
}

// RunResult represents the result of code execution.
type RunResult struct {
	Passed   bool   `json:"passed"`
	Input    string `json:"input,omitempty"`
	Expected string `json:"expected,omitempty"`
	Actual   string `json:"actual"`
	Error    string `json:"error,omitempty"`
}

// ExecuteResponse is the full response for a code run/submit.
type ExecuteResponse struct {
	Results   []RunResult `json:"results"`
	AllPassed bool        `json:"all_passed"`
	Error     string      `json:"error,omitempty"`
}

// RunGo executes Go code against test cases.
func (s *ExecuteService) RunGo(code string, testCases json.RawMessage) (*ExecuteResponse, error) {
	var cases []struct {
		Input    string `json:"input"`
		Expected string `json:"expected"`
	}
	if testCases != nil && string(testCases) != "null" {
		if err := json.Unmarshal(testCases, &cases); err != nil {
			return nil, fmt.Errorf("invalid test_cases: %w", err)
		}
	}

	if len(cases) == 0 {
		// Run without test cases — just execute and return output
		result, err := gorunner.Run(code, "")
		if err != nil {
			return &ExecuteResponse{Error: err.Error()}, nil
		}
		output := result.Stdout
		if result.Stderr != "" {
			output = result.Stderr
		}
		return &ExecuteResponse{
			Results: []RunResult{{
				Passed: result.ExitCode == 0,
				Actual: output,
			}},
			AllPassed: result.ExitCode == 0,
		}, nil
	}

	var results []RunResult
	allPassed := true

	for _, tc := range cases {
		result, err := gorunner.Run(code, tc.Input)
		if err != nil {
			results = append(results, RunResult{
				Passed:   false,
				Input:    tc.Input,
				Expected: tc.Expected,
				Actual:   "Error: " + err.Error(),
			})
			allPassed = false
			continue
		}

		actual := result.Stdout
		if result.ExitCode != 0 {
			actual = result.Stderr
		}
		expected := strings.TrimSpace(tc.Expected)
		passed := strings.TrimSpace(actual) == expected

		if !passed {
			allPassed = false
		}

		results = append(results, RunResult{
			Passed:   passed,
			Input:    tc.Input,
			Expected: tc.Expected,
			Actual:   strings.TrimSpace(actual),
		})
	}

	return &ExecuteResponse{
		Results:   results,
		AllPassed: allPassed,
	}, nil
}

// RunSQL executes SQL code in a PostgreSQL sandbox.
func (s *ExecuteService) RunSQL(code string, schemaInfo json.RawMessage, testCases json.RawMessage) (*ExecuteResponse, error) {
	sandbox := sqlsandbox.New(s.sqlDB)

	result, err := sandbox.Execute(schemaInfo, code)
	if err != nil {
		return &ExecuteResponse{Error: err.Error()}, nil
	}

	if result.Error != "" {
		return &ExecuteResponse{
			Results: []RunResult{{
				Passed: false,
				Actual: result.Error,
				Error:  result.Error,
			}},
			AllPassed: false,
		}, nil
	}

	// Compare with expected output test cases
	if testCases == nil || string(testCases) == "null" {
		// No expected output — just return the result
		actualJSON, _ := json.Marshal(map[string]interface{}{
			"columns": result.Columns,
			"rows":    result.Rows,
		})
		return &ExecuteResponse{
			Results: []RunResult{{
				Passed: true,
				Actual: string(actualJSON),
			}},
			AllPassed: true,
		}, nil
	}

	var parsedCases []struct {
		Description string      `json:"description"`
		Type        string      `json:"type"`
		Expected    interface{} `json:"expected"`
		Row         *int        `json:"row"`
		Col         *string     `json:"col"`
	}
	if err := json.Unmarshal(testCases, &parsedCases); err != nil {
		return &ExecuteResponse{Error: "invalid test_cases format: " + err.Error()}, nil
	}

	var results []RunResult
	allPassed := true

	for i, tc := range parsedCases {
		passed := false
		var actualVal interface{}

		switch tc.Type {
		case "row_count":
			actualVal = len(result.Rows)
			expectedFloat, ok := tc.Expected.(float64)
			if ok && float64(len(result.Rows)) == expectedFloat {
				passed = true
			}
		case "has_column":
			expectedStr, ok := tc.Expected.(string)
			actualCols := strings.Join(result.Columns, ", ")
			actualVal = actualCols
			if ok {
				for _, col := range result.Columns {
					if col == expectedStr {
						passed = true
						break
					}
				}
			}
		case "cell_value":
			if tc.Row != nil && tc.Col != nil && *tc.Row < len(result.Rows) {
				colIdx := -1
				for i, c := range result.Columns {
					if c == *tc.Col {
						colIdx = i
						break
					}
				}
				if colIdx != -1 && colIdx < len(result.Rows[*tc.Row]) {
					actualVal = result.Rows[*tc.Row][colIdx]
					actualStr := fmt.Sprintf("%v", actualVal)
					expectedStr := fmt.Sprintf("%v", tc.Expected)
					if actualStr == expectedStr {
						passed = true
					}
				} else {
					actualVal = "column not found or row incomplete"
				}
			} else {
				actualVal = "row out of bounds"
			}
		default:
			actualVal = "unknown test type"
		}

		if !passed {
			allPassed = false
		}

		desc := tc.Description
		if desc == "" {
			desc = fmt.Sprintf("Test Case %d", i+1)
		}

		expectedStr := fmt.Sprintf("%v", tc.Expected)
		actualStr := fmt.Sprintf("%v", actualVal)

		if tc.Type == "row_count" {
			expectedStr = fmt.Sprintf("%v rows", tc.Expected)
			actualStr = fmt.Sprintf("%v rows", actualVal)
		} else if tc.Type == "has_column" {
			expectedStr = fmt.Sprintf("column: %v", tc.Expected)
		} else if tc.Type == "cell_value" {
			expectedStr = fmt.Sprintf("cell[%v]: %v", *tc.Col, tc.Expected)
		}

		results = append(results, RunResult{
			Passed:   passed,
			Input:    tc.Type + " -> " + desc,
			Expected: expectedStr,
			Actual:   actualStr,
		})
	}

	return &ExecuteResponse{
		Results:   results,
		AllPassed: allPassed,
	}, nil
}
