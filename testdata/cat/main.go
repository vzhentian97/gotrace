package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	r, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "read error: %v", err)
		os.Exit(1)
		return
	}
	fmt.Fprintf(os.Stdout, "%s", string(r))
}
