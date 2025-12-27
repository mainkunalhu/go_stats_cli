# Git Stats CLI

A tool to track your local Git commits and show a contribution heatmap in your terminal.

## How to Use

### 1. Scan for Git repos

Run this to save a folder to your tracking list:

```bash
go run internal/app/main.go -add "/path/to/your/code"

```

### 2. See your stats

Run this to see your 6-month heatmap:

```bash
go run internal/app/main.go -email "your-email@example.com"

```

## Makefile Commands

- `make run` - Scans the current folder.
- `make stats` - Shows your commit graph.
- `make build` - Creates an executable file.
