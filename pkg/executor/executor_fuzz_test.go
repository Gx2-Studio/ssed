package executor

import (
	"bytes"
	"strings"
	"testing"

	"github.com/Gx2-Studio/ssed/pkg/lexer"
	"github.com/Gx2-Studio/ssed/pkg/parser"
)

func FuzzExecutor(f *testing.F) {
	type seed struct {
		query string
		input string
	}

	seeds := []seed{
		// Basic operations
		{"replace foo with bar", "foo bar foo\nbaz foo\n"},
		{"delete error", "info line\nerror line\nwarning line\n"},
		{"show warning", "info\nwarning here\nerror\n"},
		{"convert to uppercase", "hello world\n"},
		{"convert to lowercase", "HELLO WORLD\n"},
		{"trim whitespace", "  hello  \n  world  \n"},
		{"count error", "error\nerror\ninfo\n"},

		// Line operations
		{"delete line 2", "line1\nline2\nline3\n"},
		{"show lines 1-3", "a\nb\nc\nd\ne\n"},
		{"show first 2 lines", "1\n2\n3\n4\n5\n"},
		{"show last 2 lines", "1\n2\n3\n4\n5\n"},
		{"delete first 2 lines", "1\n2\n3\n4\n5\n"},
		{"delete last 2 lines", "1\n2\n3\n4\n5\n"},

		// Pattern matching
		{"delete lines starting with #", "# comment\ncode\n# another\n"},
		{"show lines ending with ;", "int x;\nno semi\nint y;\n"},
		{"show lines containing error", "info\nerror here\nwarning\n"},

		// Insert operations
		{"insert HEADER first", "line1\nline2\n"},
		{"insert FOOTER last", "line1\nline2\n"},
		{"insert NEW before target", "before\ntarget line\nafter\n"},
		{"insert NEW after target", "before\ntarget line\nafter\n"},

		// Compound commands
		{"delete # then convert to uppercase", "# comment\nhello\n# another\nworld\n"},
		{"trim then replace foo with bar", "  foo  \n  bar  \n"},
		{"show error then count error", "error1\ninfo\nerror2\n"},

		// Regex
		{"replace /[0-9]+/ with NUM", "test123\n456test\n"},
		{"delete /^#/", "# comment\ncode\n"},
		{"show /error|warn/", "info\nerror here\nwarning\n"},

		// Empty input
		{"replace foo with bar", ""},
		{"delete error", ""},
		{"show warning", ""},

		// Large input
		{"replace a with b", strings.Repeat("a\n", 10000)},
		{"delete a", strings.Repeat("a\nb\n", 5000)},
		{"show a", strings.Repeat("a\nb\n", 5000)},

		// Unicode
		{"replace 你好 with 世界", "你好\nworld\n"},
		{"delete 日本", "日本語\nenglish\n"},

		// Very long lines
		{"replace x with y", strings.Repeat("x", 100000) + "\n"},
		{"show x", strings.Repeat("x", 100000) + "\n"},

		// Many lines
		{"show first 5 lines", strings.Repeat("line\n", 100000)},
		{"show last 5 lines", strings.Repeat("line\n", 100000)},
		{"delete first 5 lines", strings.Repeat("line\n", 100000)},
		{"delete last 5 lines", strings.Repeat("line\n", 100000)},

		// Special characters in input
		{"replace x with y", "!@#$%^&*()\n<>?/\\|[]{}\n"},
		{"show x", "\x00\x01\x02\x03\n"},

		// Binary-like input
		{"replace a with b", string([]byte{0, 1, 2, 255, 254, 253})},

		// Regex that might be slow
		{"replace /a+/ with x", strings.Repeat("a", 1000) + "\n"},
		{"delete /.*error.*/", "this is an error line\n"},

		// Edge case: no newline at end
		{"replace foo with bar", "foo"},
		{"delete foo", "foo"},
		{"show foo", "foo"},

		// Edge case: only newlines
		{"delete empty", "\n\n\n\n\n"},
		{"show lines", "\n\n\n"},
	}

	for _, s := range seeds {
		f.Add(s.query, s.input)
	}

	f.Fuzz(func(t *testing.T, query, input string) {
		lex := lexer.New(query)
		p := parser.New(lex)
		cmd := p.Parse()

		if cmd == nil {
			return
		}

		if cmd.TokenLiteral() == "ILLEGAL" {
			return
		}

		var output bytes.Buffer
		_ = Execute(cmd, strings.NewReader(input), &output)
	})
}
