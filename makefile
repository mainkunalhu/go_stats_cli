.PHONY: run build

# Run the application with a default folder
run:
	go run internal/app/main.go -add "/home/kunal/Desktop/go-projects/git_stats_cli"

# Run the application with a custom email
stats:
	go run internal/app/main.go -email "natwarlaluzumaki@gmail.com"

# Build the binary
build:
	go build -o gocli internal/app/main.go