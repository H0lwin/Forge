package main

import (
	"context"
	"fmt"
	"os"

	"forge/cmd"
)

func main() {
	if err := cmd.Execute(context.Background(), os.Stdout, os.Stderr); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "forge error: %v\n", err)
		os.Exit(1)
	}
}
