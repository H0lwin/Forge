package main

import (
	"context"
	"os"

	"forge/cmd"
)

func main() {
	if err := cmd.Execute(context.Background(), os.Stdout, os.Stderr); err != nil {
		os.Exit(1)
	}
}
