package parser

import (
	"testing"

	"github.com/Gx2-Studio/ssed/pkg/lexer"
)

func FuzzParser(f *testing.F) {
	seeds := []string{
		// Basic commands
		"replace foo with bar",
		"delete error",
		"show warning",
		"insert header first",
		"insert footer last",
		"insert text before marker",
		"insert text after marker",
		"convert to uppercase",
		"convert to lowercase",
		"convert to titlecase",
		"trim whitespace",
		"trim",
		"remove leading whitespace",
		"remove trailing whitespace",
		"count error",
		"count lines containing error",

		// Line operations
		"delete line 5",
		"delete lines 1-10",
		"show line 1",
		"show lines 5-10",
		"show first 5 lines",
		"show last 10 lines",
		"delete first 3 lines",
		"delete last 7 lines",

		// Pattern matching
		"delete lines starting with #",
		"delete lines ending with ;",
		"show lines containing error",
		"show lines not containing debug",
		"delete lines not starting with //",
		"show lines containing whole word cat",

		// Regex
		"replace /[0-9]+/ with NUM",
		"delete /^#/",
		"show /error|warning/",
		"count /[A-Z]+/",

		// Compound commands
		"delete then replace foo with bar",
		"trim then convert to uppercase",
		"delete error then delete warning then delete info",
		"show first 10 lines then delete empty then convert to lowercase",
		"replace foo with bar then replace bar with baz then replace baz with qux",

		// Quoted strings
		`replace "hello world" with "goodbye world"`,
		`delete "special chars !@#$%"`,
		`show 'single quoted'`,
		`insert "new line" before "marker"`,

		// Empty and whitespace
		"",
		"   ",
		"\t\n\r",
		"\n\n\n",

		// Invalid commands
		"invalid command here",
		"replace",
		"replace foo",
		"replace foo with",
		"delete",
		"show",
		"insert",
		"insert foo",
		"convert",
		"convert to",
		"convert to invalid",

		// Null bytes and control characters
		"\x00",
		"\x00\x01\x02\x03",
		"replace \x00 with bar",

		// Unicode
		"replace 你好 with 世界",
		"delete línea",
		"show 日本語",
		"replace émoji with 🎉",

		// Numbers edge cases
		"delete line 0",
		"delete line -1",
		"delete line 999999999999999999999999999999",
		"show first 0 lines",
		"delete lines 10-5",
		"show line 1.5",

		// Repeated keywords
		"replace replace with replace",
		"delete delete",
		"show show",
		"with with with",
		"then then then",
		"lines lines lines",

		// Malformed compound
		"then",
		"then delete foo",
		"delete foo then",
		"delete foo then then delete bar",

		// Very long input
		"replace " + string(make([]byte, 1000)) + " with bar",

		// Regex edge cases
		"replace /^$/ with empty",
		"delete /.*+?/",
		"show /[[[/",
		"replace /(((/  with foo",

		// Mixed special characters
		`replace "!@#$%^&*()" with '<>?:"{}'`,

		// Deep compound nesting
		"delete a then delete b then delete c then delete d then delete e then delete f",

		// Keywords as values
		"replace with with foo",
		"replace then with bar",
		"delete lines containing with",
		"show lines starting with then",
	}

	for _, seed := range seeds {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, input string) {
		lex := lexer.New(input)
		p := New(lex)
		_ = p.Parse()
	})
}
