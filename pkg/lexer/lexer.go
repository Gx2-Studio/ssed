package lexer

import (
	"strings"
	"unicode"
)

type Lexer struct {
	input     string
	character byte
	position  int
	line      int
	column    int
}

func New(input string) *Lexer {
	var lexer *Lexer = &Lexer{
		input:    input,
		position: 0,
		line:     1,
		column:   0,
	}

	lexer.readChar()

	return lexer
}

func (lexer *Lexer) NextToken() Token {
	lexer.skipWhitespace()

	pos := Position{Line: lexer.line, Column: lexer.column}
	var t Token = Token{Pos: pos}

	switch {
	case unicode.IsDigit(rune(lexer.character)):
		t.Literal = lexer.readNumber()
		t.Type = NUMBER

	case unicode.IsLetter(rune(lexer.character)):
		t.Literal = lexer.readIdentifier()
		t.Type = LookupIdent(t.Literal)

	case lexer.character == '\'' || lexer.character == '"':
		t.Literal = lexer.readString()
		t.Type = STRING

	case lexer.character == '/':
		t.Literal = lexer.readRegex()
		t.Type = REGEX

	case lexer.character == byte(0):
		t.Type = EOF

	default:
		t.Literal = string(lexer.character)
		t.Type = ILLEGAL

		lexer.readChar()
	}

	return t
}

func (lexer *Lexer) readChar() {
	if lexer.position >= len(lexer.input) {
		lexer.character = byte(0)
	} else {
		if lexer.character == '\n' {
			lexer.line++
			lexer.column = 0
		}

		lexer.character = lexer.input[lexer.position]
		lexer.position = lexer.position + 1
		lexer.column++
	}
}

func (lexer *Lexer) peekChar() byte {
	if lexer.position < len(lexer.input) {
		return lexer.input[lexer.position]
	}

	return byte(0)
}

func (lexer *Lexer) skipWhitespace() {
	for unicode.IsSpace(rune(lexer.character)) {
		lexer.readChar()
	}
}

func (lexer *Lexer) readIdentifier() string {
	var b strings.Builder

	for unicode.IsLetter(rune(lexer.character)) {
		b.WriteByte(lexer.character)
		lexer.readChar()
	}

	return b.String()
}

func (lexer *Lexer) readString() string {
	var b strings.Builder

	openingChar := lexer.character

	lexer.readChar()

	for lexer.character != openingChar && lexer.character != byte(0) {
		if lexer.character == '\\' {
			next := lexer.peekChar()
			if next == openingChar || next == '\\' {
				lexer.readChar()
			}
		}

		b.WriteByte(lexer.character)
		lexer.readChar()
	}

	lexer.readChar()

	return b.String()
}

func (lexer *Lexer) readNumber() string {
	var b strings.Builder

	for unicode.IsDigit(rune(lexer.character)) && lexer.character != byte(0) {
		b.WriteByte(lexer.character)
		lexer.readChar()
	}

	if lexer.character == '.' && unicode.IsDigit(rune(lexer.peekChar())) {
		b.WriteByte('.')

		lexer.readChar()

		for unicode.IsDigit(rune(lexer.character)) && lexer.character != byte(0) {
			b.WriteByte(lexer.character)
			lexer.readChar()
		}
	}

	return b.String()
}

func (lexer *Lexer) readRegex() string {
	var b strings.Builder

	lexer.readChar()

	for lexer.character != '/' && lexer.character != byte(0) {
		if lexer.character == '\\' && lexer.peekChar() == '/' {
			b.WriteByte(lexer.character)
			lexer.readChar()
		}

		b.WriteByte(lexer.character)
		lexer.readChar()
	}

	lexer.readChar()

	return b.String()
}
