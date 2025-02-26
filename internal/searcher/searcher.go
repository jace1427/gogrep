package searcher

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jace1427/gogrep/internal/gogrep"
)

func Search(opts gogrep.Options) error {
	const (
		blue  = "\033[34m"
		reset = "\033[0m"
	)

	useColor := opts.Color && isTerminal()

	for _, path := range opts.Files {
		file, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("Failed to open %s: %s", path, err)
		}

		scanner := bufio.NewScanner(file)
		lineNum := 0

		for scanner.Scan() {
			lineNum++
			line := scanner.Text()

			matched := opts.Pattern.MatchString(line)

			// Print the line if it matches the pattern (or doesn't, if -v is set)
			if matched != opts.InvertMatch {
				if path != "" {
					fmt.Printf("%s:", path)
				}

				if opts.ShowLineNumber {
					fmt.Printf("%d:", lineNum)
				}

				if matched && useColor {
					// Find all matches in the line and colorize them
					coloredLine := opts.Pattern.ReplaceAllStringFunc(line, func(match string) string {
						return blue + match + reset
					})
					fmt.Println(coloredLine)
				} else {
					// Print without coloring
					fmt.Println(line)
				}
			}

		}

		file.Close()
		if err := scanner.Err(); err != nil {
			return fmt.Errorf("Error reading: %s", err)
		}
	}

	return nil
}

func isTerminal() bool {
	fileInfo, err := os.Stdout.Stat()
	if err != nil {
		return false
	}
	return (fileInfo.Mode() & os.ModeCharDevice) != 0
}
