package lexer

import "testing"

func TestNextToken(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []Token
	}{
		{
			"single keyword", "replace", []Token{
				{Type: REPLACE, Literal: "replace"},
				{Type: EOF, Literal: ""},
			},
		},
		{
			"full replace command", "replace foo with bar", []Token{
				{Type: REPLACE, Literal: "replace"},
				{Type: IDENTIFIER, Literal: "foo"},
				{Type: WITH, Literal: "with"},
				{Type: IDENTIFIER, Literal: "bar"},
				{Type: EOF, Literal: ""},
			},
		},
		{
			"uppercase keyword becomes identifier", "REPLACE", []Token{
				{Type: IDENTIFIER, Literal: "REPLACE"},
				{Type: EOF, Literal: ""},
			},
		},
		{
			"single quoted string", "'hello world'", []Token{
				{Type: STRING, Literal: "hello world"},
				{Type: EOF, Literal: ""},
			},
		},
		{
			"double quoted string", `"hello world"`, []Token{
				{Type: STRING, Literal: "hello world"},
				{Type: EOF, Literal: ""},
			},
		},
		{
			"integer number", "42", []Token{
				{Type: NUMBER, Literal: "42"},
				{Type: EOF, Literal: ""},
			},
		},
		{
			"decimal number", "3.14", []Token{
				{Type: NUMBER, Literal: "3.14"},
				{Type: EOF, Literal: ""},
			},
		},
		{
			"empty input", "", []Token{
				{Type: EOF, Literal: ""},
			},
		},
		{
			"simple regex", "/foo/", []Token{
				{Type: REGEX, Literal: "foo"},
				{Type: EOF, Literal: ""},
			},
		},
		{
			"regex with special chars", "/^hello.*world$/", []Token{
				{Type: REGEX, Literal: "^hello.*world$"},
				{Type: EOF, Literal: ""},
			},
		},
		{
			"regex in replace command", "replace /[0-9]+/ with NUM", []Token{
				{Type: REPLACE, Literal: "replace"},
				{Type: REGEX, Literal: "[0-9]+"},
				{Type: WITH, Literal: "with"},
				{Type: IDENTIFIER, Literal: "NUM"},
				{Type: EOF, Literal: ""},
			},
		},
		{
			"then keyword", "delete foo then replace bar with baz", []Token{
				{Type: DELETE, Literal: "delete"},
				{Type: IDENTIFIER, Literal: "foo"},
				{Type: THEN, Literal: "then"},
				{Type: REPLACE, Literal: "replace"},
				{Type: IDENTIFIER, Literal: "bar"},
				{Type: WITH, Literal: "with"},
				{Type: IDENTIFIER, Literal: "baz"},
				{Type: EOF, Literal: ""},
			},
		},
		{
			"escaped double quote in string", `"foo \"bar\" baz"`, []Token{
				{Type: STRING, Literal: `foo "bar" baz`},
				{Type: EOF, Literal: ""},
			},
		},
		{
			"escaped single quote in string", `'foo \'bar\' baz'`, []Token{
				{Type: STRING, Literal: `foo 'bar' baz`},
				{Type: EOF, Literal: ""},
			},
		},
		{
			"escaped backslash in string", `"foo \\ bar"`, []Token{
				{Type: STRING, Literal: `foo \ bar`},
				{Type: EOF, Literal: ""},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			lexer := New(test.input)
			for _, expected := range test.expected {
				token := lexer.NextToken()

				if token.Type != expected.Type {
					t.Errorf("expected token type %s, got %s", expected.Type, token.Type)
				}

				if token.Literal != expected.Literal {
					t.Errorf("expected token literal %s, got %s", expected.Literal, token.Literal)
				}

			}
		})
	}
}
