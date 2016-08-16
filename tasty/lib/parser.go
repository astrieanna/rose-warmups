package tasty

import (
	"fmt"
	"io"
	"strconv"
)

// Parser represents a parser.
type Parser struct {
	s   *Scanner
	buf struct {
		tok         Token  // last read token
		lit         string // last read literal
		isUnscanned bool   // true if you should read buf first
	}
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scan() (tok Token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.isUnscanned {
		p.buf.isUnscanned = false
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() { p.buf.isUnscanned = true }

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) scanIgnoreWhitespace() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}
	return
}

func (p *Parser) parseNumber() (*float64, error) {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != INTEGER {
		return nil, fmt.Errorf("found %q, expected number", lit)
	}
	var first_half string = lit
	tok, _ = p.scan()
	if tok == PERIOD {
		tok, lit = p.scan()
		if tok != INTEGER {
			return nil, fmt.Errorf("found %q, expected second half of float", lit)
		}
		f, err := strconv.ParseFloat(first_half+"."+lit, 64)
		if err != nil {
			return nil, err
		} else {
			return &f, nil
		}
	} else {
		p.unscan()
		f, err := strconv.ParseFloat(first_half, 64)
		if err != nil {
			return nil, err
		} else {
			return &f, nil
		}
	}
}

func (p *Parser) parseIngredientName() (*string, error) {
	var name *string
	tok, lit := p.scanIgnoreWhitespace()
	if tok == WORD {
		n := lit
		name = &n
	} else {
		return nil, fmt.Errorf("found %q, expected word", lit)
	}

	for {
		tok, lit = p.scanIgnoreWhitespace()
		if tok == NEWLINE {
			return name, nil
		} else if tok == EOF {
			p.unscan()
			return name, nil
		} else if tok != WORD {
			return nil, fmt.Errorf("found %q, expected word", lit)
		} else {
			n := *name + " " + lit
			name = &n
		}
	}
}

func (p *Parser) Parse() (*Recipe, error) {
	var r Recipe = make(Recipe) // we're trying to parse an ingredients list
	for {
		value, err := p.parseNumber()
		if err != nil {
			return nil, err
		}

		tok, lit := p.scanIgnoreWhitespace()
		var unit Unit
		switch tok {
		case KW_TEASPOON:
			unit = TEASPOON
		case KW_TABLESPOON:
			unit = TABLESPOON
		case KW_CUP:
			unit = CUP
		case KW_QUART:
			unit = QUART
		default:
			return nil, fmt.Errorf("found %q, expected a unit", lit)
		}

		ingredient, err := p.parseIngredientName()
		if err != nil {
			return nil, err
		}

		r[*ingredient] = Amount{*value, unit}

		tok, lit = p.scanIgnoreWhitespace()
		if tok == EOF {
			return &r, nil // we got it! :)
		} else {
			p.unscan()
		}
	}
}
