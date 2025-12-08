package parser

import (
	"strings"
	"testing"

	"github.com/Gx2-Studio/ssed/pkg/ast"
	"github.com/Gx2-Studio/ssed/pkg/lexer"
)

func TestParseReplace(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		source      string
		replacement string
	}{
		{"simple replace", "replace foo with bar", "foo", "bar"},
		{"replace with strings", "replace 'hello' with 'world'", "hello", "world"},
		{"replace with numbers", "replace 123 with 456", "123", "456"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lex := lexer.New(tt.input)
			p := New(lex)
			cmd := p.Parse()

			replaceCmd, ok := cmd.(*ast.ReplaceCommand)
			if !ok {
				t.Fatalf("expected ReplaceCommand, got %T", cmd)
			}

			if replaceCmd.Source != tt.source {
				t.Errorf("expected source %q, got %q", tt.source, replaceCmd.Source)
			}

			if replaceCmd.Replacement != tt.replacement {
				t.Errorf("expected replacement %q, got %q", tt.replacement, replaceCmd.Replacement)
			}
		})
	}
}

func TestParseDelete(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		target string
	}{
		{"simple delete", "delete foo", "foo"},
		{"delete string", "delete 'error message'", "error message"},
		{"delete number", "delete 404", "404"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lex := lexer.New(tt.input)
			p := New(lex)
			cmd := p.Parse()

			deleteCmd, ok := cmd.(*ast.DeleteCommand)
			if !ok {
				t.Fatalf("expected DeleteCommand, got %T", cmd)
			}

			if deleteCmd.Target != tt.target {
				t.Errorf("expected target %q, got %q", tt.target, deleteCmd.Target)
			}
		})
	}
}

func TestParseShow(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		target string
	}{
		{"simple show", "show foo", "foo"},
		{"show string", "show 'warning'", "warning"},
		{"show number", "show 500", "500"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lex := lexer.New(tt.input)
			p := New(lex)
			cmd := p.Parse()

			showCmd, ok := cmd.(*ast.ShowCommand)
			if !ok {
				t.Fatalf("expected ShowCommand, got %T", cmd)
			}

			if showCmd.Target != tt.target {
				t.Errorf("expected target %q, got %q", tt.target, showCmd.Target)
			}
		})
	}
}

func TestParseInsert(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		text      string
		position  ast.InsertPosition
		reference string
	}{
		{"insert before", "insert header before title", "header", ast.InsertBefore, "title"},
		{"insert after", "insert footer after content", "footer", ast.InsertAfter, "content"},
		{"insert first", "insert 'preamble' first", "preamble", ast.InsertPrepend, ""},
		{"insert last", "insert 'end' last", "end", ast.InsertAppend, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lex := lexer.New(tt.input)
			p := New(lex)
			cmd := p.Parse()

			insertCmd, ok := cmd.(*ast.InsertCommand)
			if !ok {
				t.Fatalf("expected InsertCommand, got %T", cmd)
			}

			if insertCmd.Text != tt.text {
				t.Errorf("expected text %q, got %q", tt.text, insertCmd.Text)
			}

			if insertCmd.Position != tt.position {
				t.Errorf("expected position %v, got %v", tt.position, insertCmd.Position)
			}

			if insertCmd.Reference != tt.reference {
				t.Errorf("expected reference %q, got %q", tt.reference, insertCmd.Reference)
			}
		})
	}
}

func TestParseDeleteLine(t *testing.T) {
	tests := []struct {
		name  string
		input string
		start int
		end   int
	}{
		{"delete single line", "delete line 5", 5, 0},
		{"delete line range", "delete lines 5 to 10", 5, 10},
		{"delete line 1", "delete line 1", 1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lex := lexer.New(tt.input)
			p := New(lex)
			cmd := p.Parse()

			deleteCmd, ok := cmd.(*ast.DeleteCommand)
			if !ok {
				t.Fatalf("expected DeleteCommand, got %T", cmd)
			}

			if deleteCmd.LineRange == nil {
				t.Fatal("expected LineRange, got nil")
			}

			if deleteCmd.LineRange.Start != tt.start {
				t.Errorf("expected start %d, got %d", tt.start, deleteCmd.LineRange.Start)
			}

			if deleteCmd.LineRange.End != tt.end {
				t.Errorf("expected end %d, got %d", tt.end, deleteCmd.LineRange.End)
			}
		})
	}
}

func TestParseShowLineNumbers(t *testing.T) {
	lex := lexer.New("show line numbers")
	p := New(lex)
	cmd := p.Parse()

	showCmd, ok := cmd.(*ast.ShowCommand)
	if !ok {
		t.Fatalf("expected ShowCommand, got %T", cmd)
	}

	if !showCmd.ShowLineNumbers {
		t.Error("expected ShowLineNumbers to be true")
	}
}

func TestParseShowLine(t *testing.T) {
	tests := []struct {
		name  string
		input string
		start int
		end   int
	}{
		{"show single line", "show line 3", 3, 0},
		{"show line range", "show lines 1 to 5", 1, 5},
		{"show last lines", "show lines 10 to 20", 10, 20},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lex := lexer.New(tt.input)
			p := New(lex)
			cmd := p.Parse()

			showCmd, ok := cmd.(*ast.ShowCommand)
			if !ok {
				t.Fatalf("expected ShowCommand, got %T", cmd)
			}

			if showCmd.LineRange == nil {
				t.Fatal("expected LineRange, got nil")
			}

			if showCmd.LineRange.Start != tt.start {
				t.Errorf("expected start %d, got %d", tt.start, showCmd.LineRange.Start)
			}

			if showCmd.LineRange.End != tt.end {
				t.Errorf("expected end %d, got %d", tt.end, showCmd.LineRange.End)
			}
		})
	}
}

func TestParseReplaceRegex(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		source      string
		isRegex     bool
		replacement string
	}{
		{"regex replace", "replace /[0-9]+/ with NUM", "[0-9]+", true, "NUM"},
		{"regex with anchors", "replace /^foo/ with bar", "^foo", true, "bar"},
		{"literal replace", "replace foo with bar", "foo", false, "bar"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lex := lexer.New(tt.input)
			p := New(lex)
			cmd := p.Parse()

			replaceCmd, ok := cmd.(*ast.ReplaceCommand)
			if !ok {
				t.Fatalf("expected ReplaceCommand, got %T", cmd)
			}

			if replaceCmd.Source != tt.source {
				t.Errorf("expected source %q, got %q", tt.source, replaceCmd.Source)
			}

			if replaceCmd.IsRegex != tt.isRegex {
				t.Errorf("expected IsRegex %v, got %v", tt.isRegex, replaceCmd.IsRegex)
			}

			if replaceCmd.Replacement != tt.replacement {
				t.Errorf("expected replacement %q, got %q", tt.replacement, replaceCmd.Replacement)
			}
		})
	}
}

func TestParseDeleteRegex(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		target  string
		isRegex bool
	}{
		{"regex delete", "delete /error.*/", "error.*", true},
		{"literal delete", "delete error", "error", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lex := lexer.New(tt.input)
			p := New(lex)
			cmd := p.Parse()

			deleteCmd, ok := cmd.(*ast.DeleteCommand)
			if !ok {
				t.Fatalf("expected DeleteCommand, got %T", cmd)
			}

			if deleteCmd.Target != tt.target {
				t.Errorf("expected target %q, got %q", tt.target, deleteCmd.Target)
			}

			if deleteCmd.IsRegex != tt.isRegex {
				t.Errorf("expected IsRegex %v, got %v", tt.isRegex, deleteCmd.IsRegex)
			}
		})
	}
}

func TestParseShowRegex(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		target  string
		isRegex bool
	}{
		{"regex show", "show /^ERROR/", "^ERROR", true},
		{"literal show", "show error", "error", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lex := lexer.New(tt.input)
			p := New(lex)
			cmd := p.Parse()

			showCmd, ok := cmd.(*ast.ShowCommand)
			if !ok {
				t.Fatalf("expected ShowCommand, got %T", cmd)
			}

			if showCmd.Target != tt.target {
				t.Errorf("expected target %q, got %q", tt.target, showCmd.Target)
			}

			if showCmd.IsRegex != tt.isRegex {
				t.Errorf("expected IsRegex %v, got %v", tt.isRegex, showCmd.IsRegex)
			}
		})
	}
}

func TestParseIllegal(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"unknown command", "unknown foo"},
		{"replace missing with", "replace foo bar"},
		{"empty input", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lex := lexer.New(tt.input)
			p := New(lex)
			cmd := p.Parse()

			_, ok := cmd.(*ast.Illegal)
			if !ok {
				t.Fatalf("expected Illegal, got %T", cmd)
			}
		})
	}
}

func TestParseErrorMessages(t *testing.T) {
	tests := []struct {
		name            string
		input           string
		expectedContain string
	}{
		{
			"unknown command error",
			"unknown foo",
			"unknown command",
		},
		{
			"replace missing with",
			"replace foo bar",
			"expected 'with'",
		},
		{
			"empty input error",
			"",
			"empty input",
		},
		{
			"delete line missing number",
			"delete line foo",
			"expected line number",
		},
		{
			"insert missing position",
			"insert header foo",
			"expected 'before'",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lex := lexer.New(tt.input)
			p := New(lex)
			cmd := p.Parse()

			illegal, ok := cmd.(*ast.Illegal)
			if !ok {
				t.Fatalf("expected Illegal, got %T", cmd)
			}

			if illegal.Message == "" {
				t.Error("expected error message to be set")
			}

			if !strings.Contains(illegal.Message, tt.expectedContain) {
				t.Errorf(
					"expected message to contain %q, got %q",
					tt.expectedContain,
					illegal.Message,
				)
			}
		})
	}
}

func TestParseTransform(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		transformType ast.TransformType
	}{
		{"convert to uppercase", "convert to uppercase", ast.TransformUppercase},
		{"convert to lowercase", "convert to lowercase", ast.TransformLowercase},
		{"convert to titlecase", "convert to titlecase", ast.TransformTitlecase},
		{"trim whitespace", "trim whitespace", ast.TransformTrim},
		{"trim only", "trim", ast.TransformTrim},
		{"remove trailing spaces", "remove trailing spaces", ast.TransformTrimTrailing},
		{"remove trailing whitespace", "remove trailing whitespace", ast.TransformTrimTrailing},
		{"remove leading spaces", "remove leading spaces", ast.TransformTrimLeading},
		{"remove leading whitespace", "remove leading whitespace", ast.TransformTrimLeading},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lex := lexer.New(tt.input)
			p := New(lex)
			cmd := p.Parse()

			transformCmd, ok := cmd.(*ast.TransformCommand)
			if !ok {
				t.Fatalf("expected TransformCommand, got %T", cmd)
			}

			if transformCmd.Type != tt.transformType {
				t.Errorf("expected type %v, got %v", tt.transformType, transformCmd.Type)
			}
		})
	}
}

func TestParseCount(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		target  string
		isRegex bool
	}{
		{"count literal", "count error", "error", false},
		{"count lines containing", "count lines containing error", "error", false},
		{"count containing", "count containing warning", "warning", false},
		{"count regex", "count /^ERROR/", "^ERROR", true},
		{"count lines containing regex", "count lines containing /[0-9]+/", "[0-9]+", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lex := lexer.New(tt.input)
			p := New(lex)
			cmd := p.Parse()

			countCmd, ok := cmd.(*ast.CountCommand)
			if !ok {
				t.Fatalf("expected CountCommand, got %T", cmd)
			}

			if countCmd.Target != tt.target {
				t.Errorf("expected target %q, got %q", tt.target, countCmd.Target)
			}

			if countCmd.IsRegex != tt.isRegex {
				t.Errorf("expected IsRegex %v, got %v", tt.isRegex, countCmd.IsRegex)
			}
		})
	}
}

func TestParseFirstLastLines(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		firstN  int
		lastN   int
		cmdType string // "show" or "delete"
	}{
		{"show first 5 lines", "show first 5 lines", 5, 0, "show"},
		{"show first 3", "show first 3", 3, 0, "show"},
		{"show last 10 lines", "show last 10 lines", 0, 10, "show"},
		{"show last 2", "show last 2", 0, 2, "show"},
		{"delete first 5 lines", "delete first 5 lines", 5, 0, "delete"},
		{"delete last 3 lines", "delete last 3 lines", 0, 3, "delete"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lex := lexer.New(tt.input)
			p := New(lex)
			cmd := p.Parse()

			if tt.cmdType == "show" {
				showCmd, ok := cmd.(*ast.ShowCommand)
				if !ok {
					t.Fatalf("expected ShowCommand, got %T", cmd)
				}

				if showCmd.FirstN != tt.firstN {
					t.Errorf("expected FirstN %d, got %d", tt.firstN, showCmd.FirstN)
				}

				if showCmd.LastN != tt.lastN {
					t.Errorf("expected LastN %d, got %d", tt.lastN, showCmd.LastN)
				}
			} else {
				deleteCmd, ok := cmd.(*ast.DeleteCommand)
				if !ok {
					t.Fatalf("expected DeleteCommand, got %T", cmd)
				}

				if deleteCmd.FirstN != tt.firstN {
					t.Errorf("expected FirstN %d, got %d", tt.firstN, deleteCmd.FirstN)
				}

				if deleteCmd.LastN != tt.lastN {
					t.Errorf("expected LastN %d, got %d", tt.lastN, deleteCmd.LastN)
				}
			}
		})
	}
}

func TestParseNaturalPatterns(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		target      string
		patternType ast.PatternType
		negated     bool
		wholeWord   bool
		cmdType     string // "show" or "delete"
	}{
		{
			"show lines starting with",
			"show lines starting with #",
			"#",
			ast.PatternStartsWith,
			false,
			false,
			"show",
		},
		{
			"show lines starting with (with)",
			"show lines starting with foo",
			"foo",
			ast.PatternStartsWith,
			false,
			false,
			"show",
		},
		{
			"show lines ending with",
			"show lines ending with ;",
			";",
			ast.PatternEndsWith,
			false,
			false,
			"show",
		},
		{
			"show lines containing",
			"show lines containing error",
			"error",
			ast.PatternContains,
			false,
			false,
			"show",
		},
		{
			"delete lines starting with",
			"delete lines starting with #",
			"#",
			ast.PatternStartsWith,
			false,
			false,
			"delete",
		},
		{
			"delete lines ending with",
			"delete lines ending with ;",
			";",
			ast.PatternEndsWith,
			false,
			false,
			"delete",
		},
		{
			"delete lines containing",
			"delete lines containing debug",
			"debug",
			ast.PatternContains,
			false,
			false,
			"delete",
		},
		// Negated patterns
		{
			"show lines not starting with",
			"show lines not starting with #",
			"#",
			ast.PatternStartsWith,
			true,
			false,
			"show",
		},
		{
			"show lines not ending with",
			"show lines not ending with ;",
			";",
			ast.PatternEndsWith,
			true,
			false,
			"show",
		},
		{
			"show lines not containing",
			"show lines not containing error",
			"error",
			ast.PatternContains,
			true,
			false,
			"show",
		},
		{
			"delete lines not containing",
			"delete lines not containing debug",
			"debug",
			ast.PatternContains,
			true,
			false,
			"delete",
		},
		// Whole word patterns
		{
			"show lines containing whole word",
			"show lines containing whole word cat",
			"cat",
			ast.PatternContains,
			false,
			true,
			"show",
		},
		{
			"delete lines containing whole word",
			"delete lines containing whole word dog",
			"dog",
			ast.PatternContains,
			false,
			true,
			"delete",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lex := lexer.New(tt.input)
			p := New(lex)
			cmd := p.Parse()

			if tt.cmdType == "show" {
				showCmd, ok := cmd.(*ast.ShowCommand)
				if !ok {
					t.Fatalf("expected ShowCommand, got %T", cmd)
				}

				if showCmd.Target != tt.target {
					t.Errorf("expected target %q, got %q", tt.target, showCmd.Target)
				}

				if showCmd.PatternType != tt.patternType {
					t.Errorf("expected patternType %v, got %v", tt.patternType, showCmd.PatternType)
				}

				if showCmd.Negated != tt.negated {
					t.Errorf("expected negated %v, got %v", tt.negated, showCmd.Negated)
				}

				if showCmd.WholeWord != tt.wholeWord {
					t.Errorf("expected wholeWord %v, got %v", tt.wholeWord, showCmd.WholeWord)
				}
			} else {
				deleteCmd, ok := cmd.(*ast.DeleteCommand)
				if !ok {
					t.Fatalf("expected DeleteCommand, got %T", cmd)
				}

				if deleteCmd.Target != tt.target {
					t.Errorf("expected target %q, got %q", tt.target, deleteCmd.Target)
				}

				if deleteCmd.PatternType != tt.patternType {
					t.Errorf("expected patternType %v, got %v", tt.patternType, deleteCmd.PatternType)
				}

				if deleteCmd.Negated != tt.negated {
					t.Errorf("expected negated %v, got %v", tt.negated, deleteCmd.Negated)
				}

				if deleteCmd.WholeWord != tt.wholeWord {
					t.Errorf("expected wholeWord %v, got %v", tt.wholeWord, deleteCmd.WholeWord)
				}
			}
		})
	}
}

func TestParseTransformErrors(t *testing.T) {
	tests := []struct {
		name            string
		input           string
		expectedContain string
	}{
		{
			"convert missing to",
			"convert uppercase",
			"expected 'to'",
		},
		{
			"convert invalid target",
			"convert to invalid",
			"expected 'uppercase'",
		},
		{
			"trim invalid",
			"trim invalid",
			"expected 'whitespace'",
		},
		{
			"remove invalid",
			"remove invalid",
			"expected 'trailing' or 'leading'",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lex := lexer.New(tt.input)
			p := New(lex)
			cmd := p.Parse()

			illegal, ok := cmd.(*ast.Illegal)
			if !ok {
				t.Fatalf("expected Illegal, got %T", cmd)
			}

			if !strings.Contains(illegal.Message, tt.expectedContain) {
				t.Errorf(
					"expected message to contain %q, got %q",
					tt.expectedContain,
					illegal.Message,
				)
			}
		})
	}
}

func TestParseCompound(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		cmdCount int
		cmdTypes []string
	}{
		{
			"two commands",
			"delete foo then replace bar with baz",
			2,
			[]string{"DELETE", "REPLACE"},
		},
		{
			"three commands",
			"delete foo then replace bar with baz then convert to uppercase",
			3,
			[]string{"DELETE", "REPLACE", "TRANSFORM"},
		},
		{
			"single command no then",
			"delete foo",
			1,
			[]string{"DELETE"},
		},
		{
			"show then delete",
			"show error then delete warning",
			2,
			[]string{"SHOW", "DELETE"},
		},
		{
			"transform chain",
			"trim then convert to uppercase",
			2,
			[]string{"TRANSFORM", "TRANSFORM"},
		},
		{
			"delete lines then replace",
			"delete lines starting with '#' then replace TODO with DONE",
			2,
			[]string{"DELETE", "REPLACE"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lex := lexer.New(tt.input)
			p := New(lex)
			cmd := p.Parse()

			if tt.cmdCount == 1 {
				if cmd.TokenLiteral() != tt.cmdTypes[0] {
					t.Errorf("expected %s, got %s", tt.cmdTypes[0], cmd.TokenLiteral())
				}

				return
			}

			compoundCmd, ok := cmd.(*ast.CompoundCommand)
			if !ok {
				t.Fatalf("expected CompoundCommand, got %T", cmd)
			}

			if len(compoundCmd.Commands) != tt.cmdCount {
				t.Errorf("expected %d commands, got %d", tt.cmdCount, len(compoundCmd.Commands))
			}

			for i, expectedType := range tt.cmdTypes {
				if i >= len(compoundCmd.Commands) {
					break
				}

				if compoundCmd.Commands[i].TokenLiteral() != expectedType {
					t.Errorf(
						"command %d: expected %s, got %s",
						i,
						expectedType,
						compoundCmd.Commands[i].TokenLiteral(),
					)
				}
			}
		})
	}
}

func TestParseCompoundErrors(t *testing.T) {
	tests := []struct {
		name            string
		input           string
		expectedContain string
	}{
		{
			"then without second command",
			"delete foo then",
			"empty input",
		},
		{
			"then with invalid command",
			"delete foo then unknown bar",
			"unknown command",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lex := lexer.New(tt.input)
			p := New(lex)
			cmd := p.Parse()

			illegal, ok := cmd.(*ast.Illegal)
			if !ok {
				t.Fatalf("expected Illegal, got %T", cmd)
			}

			if !strings.Contains(illegal.Message, tt.expectedContain) {
				t.Errorf(
					"expected message to contain %q, got %q",
					tt.expectedContain,
					illegal.Message,
				)
			}
		})
	}
}
