package tasty

type Token int // a lexical token
const (
	ILLEGAL Token = iota
	EOF
	WS
	NEWLINE
	PERIOD // inside floating point number

	KW_TEASPOON
	KW_TABLESPOON
	KW_CUP
	KW_QUART

	WORD
	INTEGER
)

type TokenInstance struct {
	Type    Token
	Literal string
}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}

var eof = rune(0)
