package executor

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
	"unicode"

	"github.com/Gx2-Studio/ssed/pkg/ast"
)

func Execute(cmd ast.Command, input io.Reader, output io.Writer) error {
	switch command := cmd.(type) {
	case *ast.ReplaceCommand:
		return executeReplace(command, input, output)
	case *ast.DeleteCommand:
		return executeDelete(command, input, output)
	case *ast.ShowCommand:
		return executeShow(command, input, output)
	case *ast.InsertCommand:
		return executeInsert(command, input, output)
	case *ast.TransformCommand:
		return executeTransform(command, input, output)
	case *ast.CountCommand:
		return executeCount(command, input, output)
	case *ast.CompoundCommand:
		return executeCompound(command, input, output)
	default:
		return nil
	}
}

func executeCompound(cmd *ast.CompoundCommand, input io.Reader, output io.Writer) error {
	if len(cmd.Commands) == 0 {
		return nil
	}

	if len(cmd.Commands) == 1 {
		return Execute(cmd.Commands[0], input, output)
	}

	numPipes := len(cmd.Commands) - 1
	pipes := make([]*io.PipeWriter, numPipes)
	readers := make([]*io.PipeReader, numPipes)

	for i := 0; i < numPipes; i++ {
		readers[i], pipes[i] = io.Pipe()
	}

	errChan := make(chan error, len(cmd.Commands))

	for i := 0; i < numPipes; i++ {
		go func(idx int) {
			var cmdInput io.Reader
			if idx == 0 {
				cmdInput = input
			} else {
				cmdInput = readers[idx-1]
			}

			err := Execute(cmd.Commands[idx], cmdInput, pipes[idx])
			pipes[idx].CloseWithError(err)
			errChan <- err
		}(i)
	}

	lastInput := readers[numPipes-1]
	lastErr := Execute(cmd.Commands[numPipes], lastInput, output)

	for i := 0; i < numPipes; i++ {
		if err := <-errChan; err != nil {
			return err
		}
	}

	return lastErr
}

func executeReplace(cmd *ast.ReplaceCommand, input io.Reader, output io.Writer) error {
	scanner := bufio.NewScanner(input)

	var re *regexp.Regexp

	if cmd.IsRegex {
		var err error

		re, err = regexp.Compile(cmd.Source)
		if err != nil {
			return err
		}
	}

	for scanner.Scan() {
		line := scanner.Text()

		if cmd.IsRegex {
			line = re.ReplaceAllString(line, cmd.Replacement)
		} else {
			line = strings.ReplaceAll(line, cmd.Source, cmd.Replacement)
		}

		_, err := io.WriteString(output, line+"\n")
		if err != nil {
			return err
		}
	}

	return scanner.Err()
}

func executeDelete(cmd *ast.DeleteCommand, input io.Reader, output io.Writer) error {
	scanner := bufio.NewScanner(input)

	if cmd.LastN > 0 {
		var lines []string

		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			return err
		}

		cutoff := len(lines) - cmd.LastN
		if cutoff < 0 {
			cutoff = 0
		}

		for i := 0; i < cutoff; i++ {
			if _, err := io.WriteString(output, lines[i]+"\n"); err != nil {
				return err
			}
		}

		return nil
	}

	lineNum := 0

	var re *regexp.Regexp

	if cmd.IsRegex {
		var err error

		re, err = regexp.Compile(cmd.Target)
		if err != nil {
			return err
		}
	}

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()

		if cmd.FirstN > 0 && lineNum <= cmd.FirstN {
			continue
		}

		if cmd.LineRange != nil {
			if cmd.LineRange.HasRange() {
				if lineNum >= cmd.LineRange.Start && lineNum <= cmd.LineRange.End {
					continue
				}
			} else {
				if lineNum == cmd.LineRange.Start {
					continue
				}
			}
		} else if cmd.Target != "" {
			match := matchPattern(line, cmd.Target, cmd.IsRegex, cmd.PatternType, cmd.WholeWord, re)
			if cmd.Negated {
				match = !match
			}

			if match {
				continue
			}
		}

		_, err := io.WriteString(output, line+"\n")
		if err != nil {
			return err
		}
	}

	return scanner.Err()
}

func executeShow(cmd *ast.ShowCommand, input io.Reader, output io.Writer) error {
	scanner := bufio.NewScanner(input)

	if cmd.LastN > 0 {
		var lines []string

		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			return err
		}

		start := len(lines) - cmd.LastN
		if start < 0 {
			start = 0
		}

		for i := start; i < len(lines); i++ {
			if _, err := io.WriteString(output, lines[i]+"\n"); err != nil {
				return err
			}
		}

		return nil
	}

	lineNum := 0

	var re *regexp.Regexp

	if cmd.IsRegex {
		var err error

		re, err = regexp.Compile(cmd.Target)
		if err != nil {
			return err
		}
	}

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()

		if cmd.ShowLineNumbers {
			_, err := fmt.Fprintf(output, "%6d\t%s\n", lineNum, line)
			if err != nil {
				return err
			}

			continue
		}

		if cmd.FirstN > 0 {
			if lineNum <= cmd.FirstN {
				if _, err := io.WriteString(output, line+"\n"); err != nil {
					return err
				}
			}

			continue
		}

		if cmd.LineRange != nil {
			if cmd.LineRange.HasRange() {
				if lineNum < cmd.LineRange.Start || lineNum > cmd.LineRange.End {
					continue
				}
			} else {
				if lineNum != cmd.LineRange.Start {
					continue
				}
			}
		} else if cmd.Target != "" {
			match := matchPattern(line, cmd.Target, cmd.IsRegex, cmd.PatternType, cmd.WholeWord, re)
			if cmd.Negated {
				match = !match
			}

			if !match {
				continue
			}
		}

		_, err := io.WriteString(output, line+"\n")
		if err != nil {
			return err
		}
	}

	return scanner.Err()
}

func matchPattern(
	line, target string,
	isRegex bool,
	patternType ast.PatternType,
	wholeWord bool,
	re *regexp.Regexp,
) bool {
	if isRegex {
		return re.MatchString(line)
	}

	if wholeWord {
		pattern := `\b` + regexp.QuoteMeta(target) + `\b`
		matched, _ := regexp.MatchString(pattern, line)

		return matched
	}

	switch patternType {
	case ast.PatternStartsWith:
		return strings.HasPrefix(line, target)
	case ast.PatternEndsWith:
		return strings.HasSuffix(line, target)
	default:
		return strings.Contains(line, target)
	}
}

func executeInsert(cmd *ast.InsertCommand, input io.Reader, output io.Writer) error {
	scanner := bufio.NewScanner(input)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	if cmd.Position == ast.InsertPrepend {
		if _, err := io.WriteString(output, cmd.Text+"\n"); err != nil {
			return err
		}
	}

	for _, line := range lines {
		if cmd.Position == ast.InsertBefore && strings.Contains(line, cmd.Reference) {
			if _, err := io.WriteString(output, cmd.Text+"\n"); err != nil {
				return err
			}
		}

		if _, err := io.WriteString(output, line+"\n"); err != nil {
			return err
		}

		if cmd.Position == ast.InsertAfter && strings.Contains(line, cmd.Reference) {
			if _, err := io.WriteString(output, cmd.Text+"\n"); err != nil {
				return err
			}
		}
	}

	if cmd.Position == ast.InsertAppend {
		if _, err := io.WriteString(output, cmd.Text+"\n"); err != nil {
			return err
		}
	}

	return nil
}

func executeTransform(cmd *ast.TransformCommand, input io.Reader, output io.Writer) error {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()

		switch cmd.Type {
		case ast.TransformUppercase:
			line = strings.ToUpper(line)
		case ast.TransformLowercase:
			line = strings.ToLower(line)
		case ast.TransformTitlecase:
			line = toTitleCase(line)
		case ast.TransformTrim:
			line = strings.TrimSpace(line)
		case ast.TransformTrimLeading:
			line = strings.TrimLeftFunc(line, unicode.IsSpace)
		case ast.TransformTrimTrailing:
			line = strings.TrimRightFunc(line, unicode.IsSpace)
		}

		if _, err := io.WriteString(output, line+"\n"); err != nil {
			return err
		}
	}

	return scanner.Err()
}

func toTitleCase(s string) string {
	words := strings.Fields(s)

	for i, word := range words {
		if len(word) > 0 {
			runes := []rune(word)
			runes[0] = unicode.ToUpper(runes[0])

			for j := 1; j < len(runes); j++ {
				runes[j] = unicode.ToLower(runes[j])
			}

			words[i] = string(runes)
		}
	}

	return strings.Join(words, " ")
}

func executeCount(cmd *ast.CountCommand, input io.Reader, output io.Writer) error {
	scanner := bufio.NewScanner(input)
	count := 0

	var re *regexp.Regexp

	if cmd.IsRegex {
		var err error

		re, err = regexp.Compile(cmd.Target)
		if err != nil {
			return err
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		var match bool

		if cmd.IsRegex {
			match = re.MatchString(line)
		} else {
			match = strings.Contains(line, cmd.Target)
		}

		if match {
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	_, err := fmt.Fprintf(output, "%d\n", count)

	return err
}
