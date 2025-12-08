package executor

import (
	"bytes"
	"strings"
	"testing"

	"github.com/Gx2-Studio/ssed/pkg/ast"
)

func TestExecuteReplace(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		source      string
		replacement string
		expected    string
	}{
		{
			"simple replace",
			"hello world\n",
			"world",
			"go",
			"hello go\n",
		},
		{
			"replace multiple occurrences on same line",
			"foo bar foo\n",
			"foo",
			"baz",
			"baz bar baz\n",
		},
		{
			"replace across multiple lines",
			"hello world\nworld hello\n",
			"world",
			"earth",
			"hello earth\nearth hello\n",
		},
		{
			"no match does nothing",
			"hello world\n",
			"xyz",
			"abc",
			"hello world\n",
		},
		{
			"empty input",
			"",
			"foo",
			"bar",
			"",
		},
		{
			"replace with empty string (delete)",
			"hello world\n",
			" world",
			"",
			"hello\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &ast.ReplaceCommand{
				Source:      tt.source,
				Replacement: tt.replacement,
			}
			input := strings.NewReader(tt.input)
			var output bytes.Buffer

			err := Execute(cmd, input, &output)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if output.String() != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, output.String())
			}
		})
	}
}

func TestExecuteDelete(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		target   string
		expected string
	}{
		{
			"delete lines containing target",
			"hello world\nerror here\ngoodbye\n",
			"error",
			"hello world\ngoodbye\n",
		},
		{
			"delete all matching lines",
			"keep this\ndelete this\nkeep this too\ndelete this\n",
			"delete",
			"keep this\nkeep this too\n",
		},
		{
			"no match keeps all lines",
			"line one\nline two\n",
			"xyz",
			"line one\nline two\n",
		},
		{
			"empty input",
			"",
			"foo",
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &ast.DeleteCommand{Target: tt.target}
			input := strings.NewReader(tt.input)
			var output bytes.Buffer

			err := Execute(cmd, input, &output)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if output.String() != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, output.String())
			}
		})
	}
}

func TestExecuteShow(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		target   string
		expected string
	}{
		{
			"show lines containing target",
			"hello world\nerror here\ngoodbye\n",
			"error",
			"error here\n",
		},
		{
			"show multiple matching lines",
			"match this\nno dice\nmatch this too\n",
			"match",
			"match this\nmatch this too\n",
		},
		{
			"no match shows nothing",
			"line one\nline two\n",
			"xyz",
			"",
		},
		{
			"empty input",
			"",
			"foo",
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &ast.ShowCommand{Target: tt.target}
			input := strings.NewReader(tt.input)
			var output bytes.Buffer

			err := Execute(cmd, input, &output)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if output.String() != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, output.String())
			}
		})
	}
}

func TestExecuteDeleteLine(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		start    int
		end      int
		expected string
	}{
		{
			"delete single line",
			"line one\nline two\nline three\n",
			2,
			0,
			"line one\nline three\n",
		},
		{
			"delete line range",
			"line one\nline two\nline three\nline four\nline five\n",
			2,
			4,
			"line one\nline five\n",
		},
		{
			"delete first line",
			"line one\nline two\nline three\n",
			1,
			0,
			"line two\nline three\n",
		},
		{
			"delete last line",
			"line one\nline two\nline three\n",
			3,
			0,
			"line one\nline two\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &ast.DeleteCommand{
				LineRange: &ast.LineRange{Start: tt.start, End: tt.end},
			}
			input := strings.NewReader(tt.input)
			var output bytes.Buffer

			err := Execute(cmd, input, &output)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if output.String() != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, output.String())
			}
		})
	}
}

func TestExecuteShowLineNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			"show all lines with numbers",
			"hello\nworld\ntest\n",
			"     1\thello\n     2\tworld\n     3\ttest\n",
		},
		{
			"empty input",
			"",
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &ast.ShowCommand{ShowLineNumbers: true}
			input := strings.NewReader(tt.input)
			var output bytes.Buffer

			err := Execute(cmd, input, &output)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if output.String() != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, output.String())
			}
		})
	}
}

func TestExecuteShowLine(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		start    int
		end      int
		expected string
	}{
		{
			"show single line",
			"line one\nline two\nline three\n",
			2,
			0,
			"line two\n",
		},
		{
			"show line range",
			"line one\nline two\nline three\nline four\nline five\n",
			2,
			4,
			"line two\nline three\nline four\n",
		},
		{
			"show first line",
			"line one\nline two\nline three\n",
			1,
			0,
			"line one\n",
		},
		{
			"show last line",
			"line one\nline two\nline three\n",
			3,
			0,
			"line three\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &ast.ShowCommand{
				LineRange: &ast.LineRange{Start: tt.start, End: tt.end},
			}
			input := strings.NewReader(tt.input)
			var output bytes.Buffer

			err := Execute(cmd, input, &output)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if output.String() != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, output.String())
			}
		})
	}
}

func TestExecuteReplaceRegex(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		source      string
		replacement string
		expected    string
	}{
		{
			"regex replace digits",
			"price: 100 dollars\nprice: 200 dollars\n",
			"[0-9]+",
			"NUM",
			"price: NUM dollars\nprice: NUM dollars\n",
		},
		{
			"regex replace with anchors",
			"hello world\nworld hello\n",
			"^hello",
			"hi",
			"hi world\nworld hello\n",
		},
		{
			"regex replace word boundary",
			"cat catalog cats\n",
			"\\bcat\\b",
			"dog",
			"dog catalog cats\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &ast.ReplaceCommand{
				Source:      tt.source,
				IsRegex:     true,
				Replacement: tt.replacement,
			}
			input := strings.NewReader(tt.input)
			var output bytes.Buffer

			err := Execute(cmd, input, &output)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if output.String() != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, output.String())
			}
		})
	}
}

func TestExecuteDeleteRegex(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		target   string
		expected string
	}{
		{
			"delete lines matching regex",
			"INFO: ok\nERROR: fail\nINFO: good\n",
			"^ERROR",
			"INFO: ok\nINFO: good\n",
		},
		{
			"delete lines with numbers",
			"line one\nline 123\nline two\n",
			"[0-9]+",
			"line one\nline two\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &ast.DeleteCommand{Target: tt.target, IsRegex: true}
			input := strings.NewReader(tt.input)
			var output bytes.Buffer

			err := Execute(cmd, input, &output)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if output.String() != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, output.String())
			}
		})
	}
}

func TestExecuteShowRegex(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		target   string
		expected string
	}{
		{
			"show lines matching regex",
			"INFO: ok\nERROR: fail\nWARN: maybe\n",
			"^(ERROR|WARN)",
			"ERROR: fail\nWARN: maybe\n",
		},
		{
			"show lines ending with pattern",
			"hello world\ngoodbye\nhello there\n",
			"world$",
			"hello world\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &ast.ShowCommand{Target: tt.target, IsRegex: true}
			input := strings.NewReader(tt.input)
			var output bytes.Buffer

			err := Execute(cmd, input, &output)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if output.String() != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, output.String())
			}
		})
	}
}

func TestExecuteInsert(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		text      string
		position  ast.InsertPosition
		reference string
		expected  string
	}{
		{
			"insert before matching line",
			"line one\ntarget line\nline three\n",
			"INSERTED",
			ast.InsertBefore,
			"target",
			"line one\nINSERTED\ntarget line\nline three\n",
		},
		{
			"insert after matching line",
			"line one\ntarget line\nline three\n",
			"INSERTED",
			ast.InsertAfter,
			"target",
			"line one\ntarget line\nINSERTED\nline three\n",
		},
		{
			"prepend to start",
			"line one\nline two\n",
			"HEADER",
			ast.InsertPrepend,
			"",
			"HEADER\nline one\nline two\n",
		},
		{
			"append to end",
			"line one\nline two\n",
			"FOOTER",
			ast.InsertAppend,
			"",
			"line one\nline two\nFOOTER\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &ast.InsertCommand{
				Text:      tt.text,
				Position:  tt.position,
				Reference: tt.reference,
			}
			input := strings.NewReader(tt.input)
			var output bytes.Buffer

			err := Execute(cmd, input, &output)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if output.String() != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, output.String())
			}
		})
	}
}

func TestExecuteShowFirstLastN(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		firstN   int
		lastN    int
		expected string
	}{
		{
			"show first 3 lines",
			"line 1\nline 2\nline 3\nline 4\nline 5\n",
			3,
			0,
			"line 1\nline 2\nline 3\n",
		},
		{
			"show last 2 lines",
			"line 1\nline 2\nline 3\nline 4\nline 5\n",
			0,
			2,
			"line 4\nline 5\n",
		},
		{
			"show first more than available",
			"line 1\nline 2\n",
			5,
			0,
			"line 1\nline 2\n",
		},
		{
			"show last more than available",
			"line 1\nline 2\n",
			0,
			5,
			"line 1\nline 2\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &ast.ShowCommand{FirstN: tt.firstN, LastN: tt.lastN}
			input := strings.NewReader(tt.input)
			var output bytes.Buffer

			err := Execute(cmd, input, &output)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if output.String() != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, output.String())
			}
		})
	}
}

func TestExecuteDeleteFirstLastN(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		firstN   int
		lastN    int
		expected string
	}{
		{
			"delete first 2 lines",
			"line 1\nline 2\nline 3\nline 4\nline 5\n",
			2,
			0,
			"line 3\nline 4\nline 5\n",
		},
		{
			"delete last 2 lines",
			"line 1\nline 2\nline 3\nline 4\nline 5\n",
			0,
			2,
			"line 1\nline 2\nline 3\n",
		},
		{
			"delete first more than available",
			"line 1\nline 2\n",
			5,
			0,
			"",
		},
		{
			"delete last more than available",
			"line 1\nline 2\n",
			0,
			5,
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &ast.DeleteCommand{FirstN: tt.firstN, LastN: tt.lastN}
			input := strings.NewReader(tt.input)
			var output bytes.Buffer

			err := Execute(cmd, input, &output)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if output.String() != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, output.String())
			}
		})
	}
}

func TestExecuteShowPatternTypes(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		target      string
		patternType ast.PatternType
		negated     bool
		expected    string
	}{
		{
			"show lines starting with",
			"# comment\nnormal line\n# another comment\n",
			"#",
			ast.PatternStartsWith,
			false,
			"# comment\n# another comment\n",
		},
		{
			"show lines ending with",
			"line one;\nline two\nline three;\n",
			";",
			ast.PatternEndsWith,
			false,
			"line one;\nline three;\n",
		},
		{
			"show lines containing",
			"hello world\ngoodbye world\nhello there\n",
			"world",
			ast.PatternContains,
			false,
			"hello world\ngoodbye world\n",
		},
		{
			"show lines NOT starting with",
			"# comment\nnormal line\n# another comment\n",
			"#",
			ast.PatternStartsWith,
			true,
			"normal line\n",
		},
		{
			"show lines NOT containing",
			"hello world\ngoodbye world\nhello there\n",
			"world",
			ast.PatternContains,
			true,
			"hello there\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &ast.ShowCommand{Target: tt.target, PatternType: tt.patternType, Negated: tt.negated}
			input := strings.NewReader(tt.input)
			var output bytes.Buffer

			err := Execute(cmd, input, &output)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if output.String() != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, output.String())
			}
		})
	}
}

func TestExecuteWholeWord(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		target   string
		expected string
	}{
		{
			"show lines with whole word cat",
			"cat is here\ncatalog of things\nthe cat meows\n",
			"cat",
			"cat is here\nthe cat meows\n",
		},
		{
			"show lines with whole word the",
			"the quick fox\nthen and there\nthe end\n",
			"the",
			"the quick fox\nthe end\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &ast.ShowCommand{Target: tt.target, WholeWord: true}
			input := strings.NewReader(tt.input)
			var output bytes.Buffer

			err := Execute(cmd, input, &output)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if output.String() != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, output.String())
			}
		})
	}
}

func TestExecuteDeletePatternTypes(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		target      string
		patternType ast.PatternType
		expected    string
	}{
		{
			"delete lines starting with",
			"# comment\nnormal line\n# another comment\n",
			"#",
			ast.PatternStartsWith,
			"normal line\n",
		},
		{
			"delete lines ending with",
			"line one;\nline two\nline three;\n",
			";",
			ast.PatternEndsWith,
			"line two\n",
		},
		{
			"delete lines containing",
			"hello world\ngoodbye world\nhello there\n",
			"world",
			ast.PatternContains,
			"hello there\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &ast.DeleteCommand{Target: tt.target, PatternType: tt.patternType}
			input := strings.NewReader(tt.input)
			var output bytes.Buffer

			err := Execute(cmd, input, &output)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if output.String() != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, output.String())
			}
		})
	}
}

func TestExecuteTransform(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		transformType ast.TransformType
		expected      string
	}{
		{
			"convert to uppercase",
			"hello world\ngoodbye\n",
			ast.TransformUppercase,
			"HELLO WORLD\nGOODBYE\n",
		},
		{
			"convert to lowercase",
			"HELLO WORLD\nGOODBYE\n",
			ast.TransformLowercase,
			"hello world\ngoodbye\n",
		},
		{
			"convert to titlecase",
			"hello world\ngoodbye friend\n",
			ast.TransformTitlecase,
			"Hello World\nGoodbye Friend\n",
		},
		{
			"trim whitespace",
			"  hello world  \n\tgoodbye\t\n",
			ast.TransformTrim,
			"hello world\ngoodbye\n",
		},
		{
			"trim leading whitespace",
			"  hello world  \n\tgoodbye\n",
			ast.TransformTrimLeading,
			"hello world  \ngoodbye\n",
		},
		{
			"trim trailing whitespace",
			"  hello world  \n\tgoodbye  \n",
			ast.TransformTrimTrailing,
			"  hello world\n\tgoodbye\n",
		},
		{
			"empty input",
			"",
			ast.TransformUppercase,
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &ast.TransformCommand{Type: tt.transformType}
			input := strings.NewReader(tt.input)
			var output bytes.Buffer

			err := Execute(cmd, input, &output)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if output.String() != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, output.String())
			}
		})
	}
}

func TestExecuteCount(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		target   string
		isRegex  bool
		expected string
	}{
		{
			"count lines containing literal",
			"hello world\nerror here\nerror again\ngoodbye\n",
			"error",
			false,
			"2\n",
		},
		{
			"count lines with no match",
			"hello world\ngoodbye\n",
			"error",
			false,
			"0\n",
		},
		{
			"count with regex pattern",
			"INFO: ok\nERROR: fail\nWARN: maybe\nERROR: again\n",
			"^ERROR",
			true,
			"2\n",
		},
		{
			"count all lines matching regex",
			"line 1\nline 2\nno number\n",
			"[0-9]+",
			true,
			"2\n",
		},
		{
			"empty input",
			"",
			"foo",
			false,
			"0\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &ast.CountCommand{Target: tt.target, IsRegex: tt.isRegex}
			input := strings.NewReader(tt.input)
			var output bytes.Buffer

			err := Execute(cmd, input, &output)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if output.String() != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, output.String())
			}
		})
	}
}

func TestExecuteCompound(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		commands []ast.Command
		expected string
	}{
		{
			"delete then replace",
			"# comment\nhello world\n# another comment\ngoodbye world\n",
			[]ast.Command{
				&ast.DeleteCommand{Target: "#", PatternType: ast.PatternStartsWith},
				&ast.ReplaceCommand{Source: "world", Replacement: "universe"},
			},
			"hello universe\ngoodbye universe\n",
		},
		{
			"replace then uppercase",
			"hello world\ngoodbye world\n",
			[]ast.Command{
				&ast.ReplaceCommand{Source: "world", Replacement: "earth"},
				&ast.TransformCommand{Type: ast.TransformUppercase},
			},
			"HELLO EARTH\nGOODBYE EARTH\n",
		},
		{
			"trim then uppercase",
			"  hello  \n  world  \n",
			[]ast.Command{
				&ast.TransformCommand{Type: ast.TransformTrim},
				&ast.TransformCommand{Type: ast.TransformUppercase},
			},
			"HELLO\nWORLD\n",
		},
		{
			"three commands chained",
			"  # comment  \n  hello WORLD  \n  # another  \n  goodbye EARTH  \n",
			[]ast.Command{
				&ast.TransformCommand{Type: ast.TransformTrim},
				&ast.DeleteCommand{Target: "#", PatternType: ast.PatternStartsWith},
				&ast.TransformCommand{Type: ast.TransformLowercase},
			},
			"hello world\ngoodbye earth\n",
		},
		{
			"delete lines then show first",
			"line 1\nkeep this\nline 3\nkeep that\nline 5\n",
			[]ast.Command{
				&ast.DeleteCommand{Target: "line"},
				&ast.ShowCommand{FirstN: 1},
			},
			"keep this\n",
		},
		{
			"single command in compound",
			"hello world\n",
			[]ast.Command{
				&ast.ReplaceCommand{Source: "world", Replacement: "earth"},
			},
			"hello earth\n",
		},
		{
			"empty compound",
			"hello world\n",
			[]ast.Command{},
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &ast.CompoundCommand{Commands: tt.commands}
			input := strings.NewReader(tt.input)
			var output bytes.Buffer

			err := Execute(cmd, input, &output)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if output.String() != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, output.String())
			}
		})
	}
}
