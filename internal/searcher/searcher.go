package searcher

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jace1427/gogrep/internal/gogrep"
)

// ANSI color codes
const blue = "\033[34m"
const reset = "\033[0m"

func Search(opts gogrep.Options) error {
	// Process each path (file or directory)
	for _, path := range opts.Files {
		// Check if path is a directory
		fileInfo, err := os.Stat(path)
		if err != nil {
			return fmt.Errorf("Failed to access %s: %s", path, err)
		}

		if fileInfo.IsDir() {
			// Path is a directory, search all files in it
			if searchDirectory(path, opts) != nil {
				return err
			}
		} else {
			// Path is a file, search it directly
			if searchFile(path, opts) != nil {
				return err
			}
		}
	}

	return nil
}

// searchDirectory searches all files in the given directory
func searchDirectory(dirPath string, opts gogrep.Options) error {
	// Read directory contents
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return fmt.Errorf("Failed to read directory %s: %s", dirPath, err)
	}

	// Process each entry
	for _, entry := range entries {
		if entry.IsDir() {
			if opts.Recursive {
				subDirPath := filepath.Join(dirPath, entry.Name())
				searchDirectory(subDirPath, opts)
			} else {
				continue
			}
		} else {
			// Get full path
			filePath := filepath.Join(dirPath, entry.Name())

			// Search the file
			err := searchFile(filePath, opts)
			if err != nil {
				// Print error but continue with other files
				fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			}
		}
	}

	return nil
}

// searchFile searches a single file for the pattern
func searchFile(filePath string, opts gogrep.Options) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("Failed to open %s: %s", filePath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()

		matched := opts.Pattern.MatchString(line)

		// Print the line if it matches the pattern (or doesn't, if -v is set)
		if matched != opts.InvertMatch {
			// Print file path
			fmt.Printf("%s:", filePath)

			// Print line number if requested
			if opts.ShowLineNumber {
				fmt.Printf("%d:", lineNum)
			}

			if matched {
				// Find all matches in the line and colorize them
				coloredLine := opts.Pattern.ReplaceAllStringFunc(line, func(match string) string {
					return blue + match + reset
				})
				fmt.Println(coloredLine)
			} else {
				// For -v option, just print the line without coloring
				fmt.Println(line)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("Error reading %s: %s", filePath, err)
	}

	return nil
}
