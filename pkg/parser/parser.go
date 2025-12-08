package parser

import (
	"fmt"
	"strconv"

	"github.com/Gx2-Studio/ssed/pkg/ast"
	"github.com/Gx2-Studio/ssed/pkg/lexer"
)

type Parser struct {
	lex       *lexer.Lexer
	curToken  lexer.Token
	peekToken lexer.Token
}

func (p *Parser) makeError(format string, args ...interface{}) *ast.Illegal {
	return &ast.Illegal{
		Identifier: p.curToken.Literal,
		Message: fmt.Sprintf(
			"line %d, column %d: "+format,
			append([]interface{}{p.curToken.Pos.Line, p.curToken.Pos.Column}, args...)...),
		Line:   p.curToken.Pos.Line,
		Column: p.curToken.Pos.Column,
	}
}

func New(lex *lexer.Lexer) *Parser {
	p := &Parser{lex: lex}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lex.NextToken()
}

func (p *Parser) Parse() ast.Command {
	cmd := p.parseSingleCommand()
	if cmd == nil {
		return nil
	}

	if _, isIllegal := cmd.(*ast.Illegal); isIllegal {
		return cmd
	}

	var commands []ast.Command

	commands = append(commands, cmd)

	for p.curToken.Type == lexer.THEN || p.peekToken.Type == lexer.THEN {
		if p.curToken.Type != lexer.THEN {
			p.nextToken()
		}

		p.nextToken()

		nextCmd := p.parseSingleCommand()
		if nextCmd == nil {
			return p.makeError("expected command after 'then'")
		}

		if _, isIllegal := nextCmd.(*ast.Illegal); isIllegal {
			return nextCmd
		}

		commands = append(commands, nextCmd)
	}

	if len(commands) == 1 {
		return commands[0]
	}

	return &ast.CompoundCommand{Commands: commands}
}

func (p *Parser) parseSingleCommand() ast.Command {
	switch p.curToken.Type {
	case lexer.REPLACE:
		return p.parseReplace()
	case lexer.DELETE:
		return p.parseDelete()
	case lexer.SHOW:
		return p.parseShow()
	case lexer.INSERT:
		return p.parseInsert()
	case lexer.CONVERT, lexer.TRIM, lexer.REMOVE:
		return p.parseTransform()
	case lexer.COUNT:
		return p.parseCount()
	case lexer.EOF:
		return p.makeError(
			"empty input, expected a command (replace, delete, show, insert, convert, count)",
		)
	default:
		return p.makeError(
			"unknown command %q, expected replace, delete, show, insert, convert, or count",
			p.curToken.Literal,
		)
	}
}

func (p *Parser) parseReplace() ast.Command {
	p.nextToken()

	if p.curToken.Type == lexer.EOF {
		return p.makeError("expected pattern to replace, got end of input")
	}

	source := p.curToken.Literal
	isRegex := p.curToken.Type == lexer.REGEX

	p.nextToken()

	if p.curToken.Type != lexer.WITH {
		return p.makeError("expected 'with' after %q in replace command", source)
	}

	p.nextToken()

	var replacement string

	if p.curToken.Type != lexer.EOF {
		replacement = p.curToken.Literal
	}

	return &ast.ReplaceCommand{
		Source:      source,
		IsRegex:     isRegex,
		Replacement: replacement,
	}
}

func (p *Parser) parseDelete() ast.Command {
	p.nextToken()

	if p.curToken.Type == lexer.FIRST || p.curToken.Type == lexer.LAST {
		isFirst := p.curToken.Type == lexer.FIRST

		p.nextToken()

		if p.curToken.Type != lexer.NUMBER {
			return p.makeError("expected number after 'first' or 'last'")
		}

		n, err := strconv.Atoi(p.curToken.Literal)
		if err != nil {
			return p.makeError("invalid number %q", p.curToken.Literal)
		}

		p.nextToken()

		if p.curToken.Type == lexer.LINES || p.curToken.Type == lexer.LINE {
			p.nextToken()
		}

		if isFirst {
			return &ast.DeleteCommand{FirstN: n}
		}

		return &ast.DeleteCommand{LastN: n}
	}

	if p.curToken.Type == lexer.LINE || p.curToken.Type == lexer.LINES {
		if p.curToken.Type == lexer.LINES {
			patternType, target, isRegex, negated, wholeWord, ok := p.parseNaturalPattern()
			if ok {
				return &ast.DeleteCommand{
					Target:      target,
					IsRegex:     isRegex,
					PatternType: patternType,
					Negated:     negated,
					WholeWord:   wholeWord,
				}
			}
		}

		return p.parseLineRange(func(lr *ast.LineRange) ast.Command {
			return &ast.DeleteCommand{LineRange: lr}
		})
	}

	target := p.curToken.Literal
	isRegex := p.curToken.Type == lexer.REGEX

	return &ast.DeleteCommand{Target: target, IsRegex: isRegex}
}

func (p *Parser) parseShow() ast.Command {
	p.nextToken()

	if p.curToken.Type == lexer.FIRST || p.curToken.Type == lexer.LAST {
		isFirst := p.curToken.Type == lexer.FIRST

		p.nextToken()

		if p.curToken.Type != lexer.NUMBER {
			return p.makeError("expected number after 'first' or 'last'")
		}

		n, err := strconv.Atoi(p.curToken.Literal)
		if err != nil {
			return p.makeError("invalid number %q", p.curToken.Literal)
		}

		p.nextToken()

		if p.curToken.Type == lexer.LINES || p.curToken.Type == lexer.LINE {
			p.nextToken()
		}

		if isFirst {
			return &ast.ShowCommand{FirstN: n}
		}

		return &ast.ShowCommand{LastN: n}
	}

	if p.curToken.Type == lexer.LINE || p.curToken.Type == lexer.LINES {
		if p.curToken.Type == lexer.LINE && p.peekToken.Type == lexer.NUMBERS {
			p.nextToken()
			p.nextToken()

			return &ast.ShowCommand{ShowLineNumbers: true}
		}

		if p.curToken.Type == lexer.LINES {
			patternType, target, isRegex, negated, wholeWord, ok := p.parseNaturalPattern()
			if ok {
				return &ast.ShowCommand{
					Target:      target,
					IsRegex:     isRegex,
					PatternType: patternType,
					Negated:     negated,
					WholeWord:   wholeWord,
				}
			}
		}

		return p.parseLineRange(func(lr *ast.LineRange) ast.Command {
			return &ast.ShowCommand{LineRange: lr}
		})
	}

	target := p.curToken.Literal
	isRegex := p.curToken.Type == lexer.REGEX

	return &ast.ShowCommand{Target: target, IsRegex: isRegex}
}

func (p *Parser) parseNaturalPattern() (ast.PatternType, string, bool, bool, bool, bool) {
	negated := false

	if p.peekToken.Type == lexer.NOT {
		negated = true

		p.nextToken()
	}

	if p.peekToken.Type == lexer.STARTING {
		p.nextToken()
		p.nextToken()

		if p.curToken.Type == lexer.WITH {
			p.nextToken()
		}

		target := p.curToken.Literal
		isRegex := p.curToken.Type == lexer.REGEX

		return ast.PatternStartsWith, target, isRegex, negated, false, true
	}

	if p.peekToken.Type == lexer.ENDING {
		p.nextToken()
		p.nextToken()

		if p.curToken.Type == lexer.WITH {
			p.nextToken()
		}

		target := p.curToken.Literal
		isRegex := p.curToken.Type == lexer.REGEX

		return ast.PatternEndsWith, target, isRegex, negated, false, true
	}

	if p.peekToken.Type == lexer.CONTAINING {
		p.nextToken()
		p.nextToken()

		wholeWord := false

		if p.curToken.Type == lexer.WHOLE && p.peekToken.Type == lexer.WORD {
			wholeWord = true

			p.nextToken()
			p.nextToken()
		}

		target := p.curToken.Literal
		isRegex := p.curToken.Type == lexer.REGEX

		return ast.PatternContains, target, isRegex, negated, wholeWord, true
	}

	return ast.PatternContains, "", false, false, false, false
}

func (p *Parser) parseLineRange(makeCmd func(*ast.LineRange) ast.Command) ast.Command {
	p.nextToken()

	if p.curToken.Type != lexer.NUMBER {
		return p.makeError("expected line number, got %q", p.curToken.Literal)
	}

	start, err := strconv.Atoi(p.curToken.Literal)
	if err != nil {
		return p.makeError("invalid line number %q", p.curToken.Literal)
	}

	lr := &ast.LineRange{Start: start}

	if p.peekToken.Type == lexer.TO {
		p.nextToken()
		p.nextToken()

		if p.curToken.Type != lexer.NUMBER {
			return p.makeError("expected end line number after 'to', got %q", p.curToken.Literal)
		}

		end, err := strconv.Atoi(p.curToken.Literal)
		if err != nil {
			return p.makeError("invalid end line number %q", p.curToken.Literal)
		}

		lr.End = end
	}

	return makeCmd(lr)
}

func (p *Parser) parseInsert() ast.Command {
	p.nextToken()

	if p.curToken.Type == lexer.EOF {
		return p.makeError("expected text to insert, got end of input")
	}

	text := p.curToken.Literal

	p.nextToken()

	var position ast.InsertPosition

	switch p.curToken.Type {
	case lexer.BEFORE:
		position = ast.InsertBefore
	case lexer.AFTER:
		position = ast.InsertAfter
	case lexer.FIRST:
		position = ast.InsertPrepend

		return &ast.InsertCommand{Text: text, Position: position, Reference: ""}
	case lexer.LAST:
		position = ast.InsertAppend

		return &ast.InsertCommand{Text: text, Position: position, Reference: ""}
	default:
		return p.makeError(
			"expected 'before', 'after', 'first', or 'last' in insert command, got %q",
			p.curToken.Literal,
		)
	}

	p.nextToken()

	if p.curToken.Type == lexer.EOF {
		return p.makeError(
			"expected reference pattern after '%s'",
			map[ast.InsertPosition]string{ast.InsertBefore: "before", ast.InsertAfter: "after"}[position],
		)
	}

	reference := p.curToken.Literal

	return &ast.InsertCommand{
		Text:      text,
		Position:  position,
		Reference: reference,
	}
}

func (p *Parser) parseTransform() ast.Command {
	startToken := p.curToken.Type

	switch startToken {
	case lexer.CONVERT:
		p.nextToken()

		if p.curToken.Type != lexer.TO {
			return p.makeError("expected 'to' after 'convert'")
		}

		p.nextToken()

		switch p.curToken.Type {
		case lexer.UPPERCASE:
			return &ast.TransformCommand{Type: ast.TransformUppercase}
		case lexer.LOWERCASE:
			return &ast.TransformCommand{Type: ast.TransformLowercase}
		case lexer.TITLECASE:
			return &ast.TransformCommand{Type: ast.TransformTitlecase}
		default:
			return p.makeError(
				"expected 'uppercase', 'lowercase', or 'titlecase' after 'convert to'",
			)
		}

	case lexer.TRIM:
		p.nextToken()

		if p.curToken.Type == lexer.WHITESPACE || p.curToken.Type == lexer.EOF ||
			p.curToken.Type == lexer.THEN {
			return &ast.TransformCommand{Type: ast.TransformTrim}
		}

		return p.makeError("expected 'whitespace' or end of input after 'trim'")

	case lexer.REMOVE:
		p.nextToken()

		switch p.curToken.Type {
		case lexer.TRAILING:
			p.nextToken()

			if p.curToken.Type == lexer.SPACES || p.curToken.Type == lexer.WHITESPACE ||
				p.curToken.Type == lexer.EOF {
				return &ast.TransformCommand{Type: ast.TransformTrimTrailing}
			}

			return p.makeError("expected 'spaces' or 'whitespace' after 'remove trailing'")
		case lexer.LEADING:
			p.nextToken()

			if p.curToken.Type == lexer.SPACES || p.curToken.Type == lexer.WHITESPACE ||
				p.curToken.Type == lexer.EOF {
				return &ast.TransformCommand{Type: ast.TransformTrimLeading}
			}

			return p.makeError("expected 'spaces' or 'whitespace' after 'remove leading'")
		default:
			return p.makeError("expected 'trailing' or 'leading' after 'remove'")
		}
	}

	return p.makeError("unexpected token in transform command")
}

func (p *Parser) parseCount() ast.Command {
	p.nextToken()

	if p.curToken.Type == lexer.LINES {
		p.nextToken()
	}

	if p.curToken.Type == lexer.CONTAINING {
		p.nextToken()
	}

	if p.curToken.Type == lexer.EOF {
		return p.makeError("expected pattern after 'count'")
	}

	target := p.curToken.Literal
	isRegex := p.curToken.Type == lexer.REGEX

	return &ast.CountCommand{Target: target, IsRegex: isRegex}
}
