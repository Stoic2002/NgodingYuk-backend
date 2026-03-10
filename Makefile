# Makefile for NgodingYuk Backend

.PHONY: all run build seed test tidy clean

APP_NAME=ngodingyuk-server
MAIN_PATH=cmd/server/main.go
SEED_PATH=./cmd/seed/
BUILD_DIR=tmp

all: tidy build

run:
	@echo "🚀 Running backend server..."
	@go run $(MAIN_PATH)

migrate:
	@echo "🚀 Running database migrations and starting server..."
	@DB_MIGRATE=true go run $(MAIN_PATH)

build:
	@echo "🔨 Building backend server..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PATH)
	@echo "✅ Build complete: $(BUILD_DIR)/$(APP_NAME)"

seed:
	@echo "🌱 Running database seeder..."
	@go run $(SEED_PATH)

seed-v2:
	@echo "🌱 Running database seeder v2..."
	@go run ./cmd/seed_v2/

test:
	@echo "🧪 Running tests..."
	@go test -v ./...

tidy:
	@echo "🧹 Tidying module dependencies..."
	@go mod tidy

clean:
	@echo "🗑️ Cleaning up..."
	@rm -rf $(BUILD_DIR)
	@echo "✅ Clean complete."
