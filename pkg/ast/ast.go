package ast

type Node interface {
	TokenLiteral() string
}

type Command interface {
	Node
	commandNode()
}

type LineRange struct {
	Start int
	End   int
}

type PatternType int

const (
	PatternContains PatternType = iota
	PatternStartsWith
	PatternEndsWith
)

func (lr LineRange) HasRange() bool {
	return lr.End > 0
}

type InsertPosition int

const (
	InsertBefore InsertPosition = iota
	InsertAfter
	InsertPrepend
	InsertAppend
)

type ReplaceCommand struct {
	Source      string
	IsRegex     bool
	Replacement string
}

func (r *ReplaceCommand) commandNode() {
}

func (r *ReplaceCommand) TokenLiteral() string {
	return "REPLACE"
}

type DeleteCommand struct {
	Target      string
	IsRegex     bool
	PatternType PatternType
	Negated     bool
	WholeWord   bool
	LineRange   *LineRange
	FirstN      int
	LastN       int
}

func (d *DeleteCommand) commandNode() {
}

func (d *DeleteCommand) TokenLiteral() string {
	return "DELETE"
}

type ShowCommand struct {
	Target          string
	IsRegex         bool
	PatternType     PatternType
	Negated         bool
	WholeWord       bool
	LineRange       *LineRange
	ShowLineNumbers bool
	FirstN          int
	LastN           int
}

func (s *ShowCommand) commandNode() {
}

func (s *ShowCommand) TokenLiteral() string {
	return "SHOW"
}

type InsertCommand struct {
	Text      string
	Position  InsertPosition
	Reference string
}

func (i *InsertCommand) commandNode() {
}

func (i *InsertCommand) TokenLiteral() string {
	return "INSERT"
}

type TransformType int

const (
	TransformUppercase TransformType = iota
	TransformLowercase
	TransformTitlecase
	TransformTrim
	TransformTrimLeading
	TransformTrimTrailing
)

type TransformCommand struct {
	Type TransformType
}

func (t *TransformCommand) commandNode() {
}

func (t *TransformCommand) TokenLiteral() string {
	return "TRANSFORM"
}

type CountCommand struct {
	Target  string
	IsRegex bool
}

func (c *CountCommand) commandNode() {
}

func (c *CountCommand) TokenLiteral() string {
	return "COUNT"
}

type CompoundCommand struct {
	Commands []Command
}

func (c *CompoundCommand) commandNode() {
}

func (c *CompoundCommand) TokenLiteral() string {
	return "COMPOUND"
}

type Illegal struct {
	Identifier string
	Message    string
	Line       int
	Column     int
}

func (i *Illegal) commandNode() {
}

func (i *Illegal) TokenLiteral() string {
	return "ILLEGAL"
}

func (i *Illegal) Error() string {
	if i.Message != "" {
		return i.Message
	}

	return "unexpected token: " + i.Identifier
}
