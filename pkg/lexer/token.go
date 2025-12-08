package lexer

type TokenType string

const (
	REPLACE    TokenType = "REPLACE"
	DELETE     TokenType = "DELETE"
	INSERT     TokenType = "INSERT"
	SHOW       TokenType = "SHOW"
	WITH       TokenType = "WITH"
	IN         TokenType = "IN"
	FROM       TokenType = "FROM"
	ALL        TokenType = "ALL"
	FIRST      TokenType = "FIRST"
	LAST       TokenType = "LAST"
	BEFORE     TokenType = "BEFORE"
	AFTER      TokenType = "AFTER"
	LINE       TokenType = "LINE"
	LINES      TokenType = "LINES"
	TO         TokenType = "TO"
	TRANSFORM  TokenType = "TRANSFORM"
	CONVERT    TokenType = "CONVERT"
	UPPERCASE  TokenType = "UPPERCASE"
	LOWERCASE  TokenType = "LOWERCASE"
	TITLECASE  TokenType = "TITLECASE"
	TRIM       TokenType = "TRIM"
	WHITESPACE TokenType = "WHITESPACE"
	TRAILING   TokenType = "TRAILING"
	LEADING    TokenType = "LEADING"
	SPACES     TokenType = "SPACES"
	REMOVE     TokenType = "REMOVE"
	COUNT      TokenType = "COUNT"
	CONTAINING TokenType = "CONTAINING"
	STARTING   TokenType = "STARTING"
	ENDING     TokenType = "ENDING"
	BETWEEN    TokenType = "BETWEEN"
	AND        TokenType = "AND"
	NOT        TokenType = "NOT"
	WHOLE      TokenType = "WHOLE"
	WORD       TokenType = "WORD"
	STOP       TokenType = "STOP"
	AT         TokenType = "AT"
	NUMBERS    TokenType = "NUMBERS"
	THEN       TokenType = "THEN"

	IDENTIFIER TokenType = "IDENTIFIER"
	STRING     TokenType = "STRING"
	NUMBER     TokenType = "NUMBER"
	REGEX      TokenType = "REGEX"

	EOF     TokenType = "EOF"
	ILLEGAL TokenType = "ILLEGAL"
	NEWLINE TokenType = "NEWLINE"
)

var keywords = map[string]TokenType{
	"replace":    REPLACE,
	"delete":     DELETE,
	"insert":     INSERT,
	"show":       SHOW,
	"with":       WITH,
	"in":         IN,
	"from":       FROM,
	"all":        ALL,
	"first":      FIRST,
	"last":       LAST,
	"before":     BEFORE,
	"after":      AFTER,
	"line":       LINE,
	"lines":      LINES,
	"to":         TO,
	"transform":  TRANSFORM,
	"convert":    CONVERT,
	"uppercase":  UPPERCASE,
	"lowercase":  LOWERCASE,
	"titlecase":  TITLECASE,
	"trim":       TRIM,
	"whitespace": WHITESPACE,
	"trailing":   TRAILING,
	"leading":    LEADING,
	"spaces":     SPACES,
	"remove":     REMOVE,
	"count":      COUNT,
	"containing": CONTAINING,
	"starting":   STARTING,
	"ending":     ENDING,
	"between":    BETWEEN,
	"and":        AND,
	"not":        NOT,
	"whole":      WHOLE,
	"word":       WORD,
	"stop":       STOP,
	"at":         AT,
	"numbers":    NUMBERS,
	"then":       THEN,
}

type Position struct {
	Line   int
	Column int
}

type Token struct {
	Type    TokenType
	Literal string
	Pos     Position
}

func LookupIdent(ident string) TokenType {
	if ttype, ok := keywords[ident]; ok {
		return ttype
	}

	return IDENTIFIER
}
