MAIN_PATH = .
BINARY_NAME = desktop

all: build

templ:
	@echo "🎨 Generating Templ files..."
	@templ generate

build: templ
	@echo "🏗️  Building project..."
	@go build -tags wails -o bin/$(BINARY_NAME) $(MAIN_PATH)

run: templ
	@echo "🚀 Starting server..."
	@go run -tags wails $(MAIN_PATH)

tidy:
	@echo "🧹 Tidying modules..."
	@go mod tidy

clean:
	@echo "🗑️  Cleaning up..."
	@rm -rf bin/
	@find . -name "*_templ.go" -delete

watch:
	@echo "👀 Starting watch mode (Air)..."
	@air