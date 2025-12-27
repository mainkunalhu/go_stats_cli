package main

import (
	"flag"

	"mainkunalhu.com/go_stats_cli/cmd/scan"
	"mainkunalhu.com/go_stats_cli/cmd/stats"
)

func main() {
	var folder string
	var email string
	flag.StringVar(&folder, "add", "", "adadd a new folder to scan for Git repositoriesd")
	flag.StringVar(&email, "email", "your@email.com", "the email to scan")
	flag.Parse()

	if folder != "" {
		scan.Scan(folder)
		return
	}

	stats.Stats(email)
}
