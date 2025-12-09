package lexer

import "testing"

func FuzzLexer(f *testing.F) {
	seeds := []string{
		"replace foo with bar",
		"delete error",
		"show warning",
		`replace "hello world" with 'goodbye'`,
		"replace /[0-9]+/ with NUM",
		"delete lines starting with #",
		"show first 10 lines",
		"delete last 5 lines",
		"convert to uppercase",
		"trim whitespace",
		"count error",
		"insert header first",
		"delete then replace foo with bar then convert to uppercase",

		// Empty and whitespace
		"",
		"   ",
		"\t\n\r",
		"\n\n\n",

		// Null bytes and control characters
		"\x00",
		"\x00\x01\x02\x03",
		"replace \x00 with bar",
		"delete \x1f\x7f",

		// Unterminated strings and regex
		`"unterminated`,
		`'unterminated`,
		"/unterminated regex",
		`"`,
		`'`,
		"/",

		// Escaped characters
		`"escaped \"quote\""`,
		`'escaped \'quote\''`,
		`"double \\ backslash"`,
		`"\\\"\\\""`,

		// Unicode and multi-byte
		"replace 你好 with 世界",
		"delete línea",
		"show 日本語",
		"replace émoji with 🎉",
		"delete \u0000\u0001",
		"replace \xc0\xc1 with foo",
		"\xff\xfe",
		"replace \xed\xa0\x80 with bar",

		// Very long tokens
		"replace " + string(make([]byte, 1000)) + " with bar",
		`"` + string(make([]byte, 10000)) + `"`,

		// Repeated keywords
		"replace replace replace",
		"with with with with",
		"delete delete delete",
		"then then then then",
		"show show show show",

		// Numbers edge cases
		"delete line 0",
		"delete line -1",
		"delete line 999999999999999999999999999999",
		"show first 0 lines",
		"delete lines 10-5",
		"show line 1.5",
		"delete line 1e10",

		// Nested quotes
		`"outer 'inner' outer"`,
		`'outer "inner" outer'`,
		`"a\"b\"c\"d\"e"`,

		// Regex edge cases
		"/^$/",
		"/.*+?/",
		"/[[[/",
		"/(((/",
		`/\\/`,
		"/a{999999}/",
		"/(a|b|c|d|e|f|g)+/",

		// Mixed special characters
		`replace "!@#$%^&*()" with '<>?:"{}'`,
		"delete |\\[]{}",
		"show ~`",

		// Only special characters
		"!@#$%^&*()",
		"<>?/\\|[]{}",
		"+-=_",

		// Very deep nesting simulation
		"delete then delete then delete then delete then delete then delete foo",
	}

	for _, seed := range seeds {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, input string) {
		lex := New(input)
		for range 100000 {
			tok := lex.NextToken()
			if tok.Type == EOF {
				break
			}
		}
	})
}
