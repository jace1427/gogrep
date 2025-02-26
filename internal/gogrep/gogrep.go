package gogrep

import (
	"flag"
	"fmt"
	"os"
	"regexp"
)

type Options struct {
	Pattern        *regexp.Regexp
	Files          []string
	IgnoreCase     bool
	ShowLineNumber bool
	InvertMatch    bool
	Color          bool
	Recursive      bool
}

func NewCommand() (*Options, error) {
	// Define flag set
	fs := flag.NewFlagSet("gogrep", flag.ExitOnError)
	ignoreCase := fs.Bool("i", false, "Ignore case distinctions")
	showLineNumber := fs.Bool("n", false, "Print line number with output lines")
	invertMatch := fs.Bool("v", false, "Select non-matching lines")
	color := fs.Bool("c", true, "Print in color")
	recursive := fs.Bool("r", false, "Search all subdirectories recursively")

	// Parse flags
	err := fs.Parse(os.Args[1:])
	if err != nil {
		return nil, err
	}

	// Get positional arguments
	args := fs.Args()
	if len(args) < 1 {
		fmt.Println("Usage: gogrep [OPTIONS] PATTERN PATH...")
		fs.PrintDefaults()
		return nil, fmt.Errorf("insufficient arguments: pattern is required.")
	}

	// Add files if provided
	if len(args) < 2 {
		fmt.Println("Usage: gogrep [OPTIONS] PATTERN PATH...")
		fs.PrintDefaults()
		return nil, fmt.Errorf("insufficient arguments: path is required.")
	}

	pattern := args[0]
	// Apply case insensitivity if requested
	if *ignoreCase {
		pattern = "(?i)" + pattern
	}

	// Compile the regular expression
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("Invalid pattern: %s", err)
	}

	// Create options struct
	opts := &Options{
		Pattern:        regex,
		Files:          args[1:],
		IgnoreCase:     *ignoreCase,
		ShowLineNumber: *showLineNumber,
		InvertMatch:    *invertMatch,
		Color:          *color,
		Recursive:      *recursive,
	}

	return opts, nil
}
