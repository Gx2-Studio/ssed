package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/Gx2-Studio/ssed/pkg/executor"
	"github.com/Gx2-Studio/ssed/pkg/lexer"
	"github.com/Gx2-Studio/ssed/pkg/parser"
)

var version = "0.1.0"

func main() {
	if err := Run(os.Args[1:], os.Stdin, os.Stdout, os.Stderr); err != nil {
		os.Exit(1)
	}
}

func Run(args []string, stdin io.Reader, stdout, stderr io.Writer) error {
	var preview, inPlace, quiet bool
	var backup string

	rootCmd := &cobra.Command{
		Use:   "ssed <query> [file...]",
		Short: "Simple sed - text transformation in plain English",
		Long: `ssed is a natural language interface for text transformation.

Instead of learning sed syntax, just describe what you want in plain English.

Examples:
  ssed "replace foo with bar" file.txt
  ssed "delete error" < input.txt
  ssed "show warning" app.log
  cat data.txt | ssed "replace hello with hi"`,
		Args:          cobra.MinimumNArgs(1),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runQuery(args, stdin, stdout, stderr, preview, inPlace, backup, quiet)
		},
	}

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(stdout, "ssed version %s\n", version)
		},
	}

	examplesCmd := &cobra.Command{
		Use:   "examples",
		Short: "Show usage examples",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(stdout, `ssed Usage Examples
==================

Replace text:
  ssed "replace hello with hi" greeting.txt
  ssed "replace foo with bar" < input.txt
  echo "hello world" | ssed "replace world with go"

Delete lines:
  ssed "delete error" app.log
  ssed "delete warning" < input.txt

Show/filter lines:
  ssed "show error" app.log
  ssed "show TODO" *.go

Insert text:
  ssed "insert header before title" doc.txt
  ssed "insert footer last" README.md

Options:
  ssed "replace foo with bar" file.txt --preview    # Preview changes
  ssed "replace foo with bar" file.txt -i           # Edit in-place
  ssed "replace foo with bar" file.txt -i --backup .bak  # With backup`)
		},
	}

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(examplesCmd)

	rootCmd.Flags().BoolVarP(&preview, "preview", "p", false, "Preview changes without applying")
	rootCmd.Flags().BoolVarP(&inPlace, "in-place", "i", false, "Edit files in-place")
	rootCmd.Flags().StringVarP(&backup, "backup", "b", "", "Backup suffix for in-place editing (e.g., .bak)")
	rootCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Suppress output (only show errors)")

	rootCmd.SetArgs(args)
	rootCmd.SetIn(stdin)
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	return rootCmd.Execute()
}

func runQuery(args []string, stdin io.Reader, stdout, stderr io.Writer, preview, inPlace bool, backup string, quiet bool) error {
	query := args[0]

	lex := lexer.New(query)
	p := parser.New(lex)
	ast := p.Parse()

	if ast == nil {
		return fmt.Errorf("failed to parse query: %s", query)
	}

	if ast.TokenLiteral() == "ILLEGAL" {
		return fmt.Errorf("unknown command in query: %s", query)
	}

	var inputs []io.Reader
	var filenames []string

	if len(args) > 1 {
		for _, filename := range args[1:] {
			file, err := os.Open(filename)
			if err != nil {
				return fmt.Errorf("error opening file %s: %w", filename, err)
			}

			defer file.Close()

			inputs = append(inputs, file)
			filenames = append(filenames, filename)
		}
	} else {
		inputs = append(inputs, stdin)
		filenames = append(filenames, "stdin")
	}

	for idx, input := range inputs {
		var output io.Writer
		var outputBuf *strings.Builder

		if preview {
			outputBuf = &strings.Builder{}
			output = outputBuf
		} else if inPlace && filenames[idx] != "stdin" {
			outputBuf = &strings.Builder{}
			output = outputBuf
		} else {
			output = stdout
		}

		var inputReader io.Reader = input

		if inPlace && filenames[idx] != "stdin" {
			content, err := io.ReadAll(input)
			if err != nil {
				return fmt.Errorf("error reading file %s: %w", filenames[idx], err)
			}

			inputReader = strings.NewReader(string(content))
		}

		err := executor.Execute(ast, inputReader, output)
		if err != nil {
			return fmt.Errorf("execution error: %w", err)
		}

		if preview && outputBuf != nil {
			fmt.Fprintf(stdout, "=== Preview for %s ===\n", filenames[idx])
			fmt.Fprintln(stdout, outputBuf.String())
			fmt.Fprintln(stdout, "=== End preview (no changes made) ===")
		}

		if inPlace && filenames[idx] != "stdin" && outputBuf != nil {
			if backup != "" {
				backupName := filenames[idx] + backup
				if err := copyFile(filenames[idx], backupName); err != nil {
					return fmt.Errorf("error creating backup: %w", err)
				}

				if !quiet {
					fmt.Fprintf(stderr, "Backup created: %s\n", backupName)
				}
			}

			if err := atomicWriteFile(filenames[idx], []byte(outputBuf.String())); err != nil {
				return fmt.Errorf("error writing file %s: %w", filenames[idx], err)
			}

			if !quiet {
				fmt.Fprintf(stderr, "Modified: %s\n", filenames[idx])
			}
		}
	}

	return nil
}

func copyFile(src, dst string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}

	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}

	defer destination.Close()

	buf := bufio.NewReader(source)
	_, err = io.Copy(destination, buf)

	return err
}

func atomicWriteFile(filename string, data []byte) error {
	info, err := os.Stat(filename)
	if err != nil {
		return err
	}

	dir := filepath.Dir(filename)

	tempFile, err := os.CreateTemp(dir, ".ssed-*")
	if err != nil {
		return err
	}

	tempName := tempFile.Name()

	defer func() {
		if tempFile != nil {
			tempFile.Close()
			os.Remove(tempName)
		}
	}()

	if _, err := tempFile.Write(data); err != nil {
		return err
	}

	if err := tempFile.Chmod(info.Mode()); err != nil {
		return err
	}

	if err := tempFile.Close(); err != nil {
		return err
	}

	tempFile = nil

	return os.Rename(tempName, filename)
}
