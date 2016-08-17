package tasty

// Represents a lexical token
type Token int

// Represent the kinds of tokens we can have in a recipe
const (
	ILLEGAL Token = iota
	EOF
	WS
	NEWLINE
	PERIOD

	KW_TEASPOON
	KW_TABLESPOON
	KW_CUP
	KW_QUART

	WORD
	INTEGER
)

// Represents a token type and the literal characters it was parsed from.
type TokenInstance struct {
	Type    Token
	Literal string
}

// Returns true if ch is a space or a tab.
func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t'
}

// Returns true if ch is a upper or lowercase letter.
func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

// Returns true if ch is a numeric digit
func isDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}

var eof = rune(0)
