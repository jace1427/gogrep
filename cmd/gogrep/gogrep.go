package main

import (
	"fmt"
	"os"

	"github.com/jace1427/gogrep/internal/gogrep"
	"github.com/jace1427/gogrep/internal/searcher"
)

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	opts, err := gogrep.NewCommand()
	CheckError(err)

	err = searcher.Search(*opts)
	CheckError(err)
}
