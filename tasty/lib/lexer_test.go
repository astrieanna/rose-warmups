package tasty

import (
	"io"
	"strings"
	"testing"
)

func TestLexingString(t *testing.T) {
	var r io.Reader = strings.NewReader("some io.Reader stream to be read\n")
	var correct_parsed []TokenInstance
	correct_parsed = []TokenInstance{TokenInstance{WORD, "some"},
		TokenInstance{WS, " "}, TokenInstance{WORD, "io"},
		TokenInstance{PERIOD, "."}, TokenInstance{WORD, "Reader"},
		TokenInstance{WS, " "}, TokenInstance{WORD, "stream"},
		TokenInstance{WS, " "}, TokenInstance{WORD, "to"},
		TokenInstance{WS, " "}, TokenInstance{WORD, "be"},
		TokenInstance{WS, " "}, TokenInstance{WORD, "read"},
		TokenInstance{NEWLINE, "\n"}}

	s := NewScanner(r)
	for _, ti := range correct_parsed {
		token, literal := s.Scan()
		if token != ti.Type {
			t.Error("Expected ", ti.Type, " but got ", token, " with value ", literal)
		}
		if literal != ti.Literal {
			t.Error("Expected ", ti.Literal, " but got ", literal, " with type ", token)
		}
	}
	token, _ := s.Scan()
	if token != EOF {
		t.Error("Expected EOF, but got ", token)
	}

}

func TestLexingIngredients(t *testing.T) {
	var r io.Reader = strings.NewReader("2c mayo\n  8 T brown sugar\n 2t cayenne\n")
	var correct_parsed []TokenInstance
	correct_parsed = []TokenInstance{TokenInstance{INTEGER, "2"},
		TokenInstance{KW_CUP, "c"}, TokenInstance{WS, " "},
		TokenInstance{WORD, "mayo"}, TokenInstance{NEWLINE, "\n"},
		TokenInstance{WS, "  "}, TokenInstance{INTEGER, "8"},
		TokenInstance{WS, " "}, TokenInstance{KW_TABLESPOON, "T"},
		TokenInstance{WS, " "}, TokenInstance{WORD, "brown"},
		TokenInstance{WS, " "}, TokenInstance{WORD, "sugar"},
		TokenInstance{NEWLINE, "\n"}, TokenInstance{WS, " "},
		TokenInstance{INTEGER, "2"}, TokenInstance{KW_TEASPOON, "t"},
		TokenInstance{WS, " "}, TokenInstance{WORD, "cayenne"},
		TokenInstance{NEWLINE, "\n"}}

	s := NewScanner(r)
	for _, ti := range correct_parsed {
		token, literal := s.Scan()
		if token != ti.Type {
			t.Error("Expected ", ti.Type, " but got ", token, " with value ", literal)
		}
		if literal != ti.Literal {
			t.Error("Expected ", ti.Literal, " but got ", literal, " with type ", token)
		}
	}
	token, _ := s.Scan()
	if token != EOF {
		t.Error("Expected EOF, but got ", token)
	}

}

func TestLexingIngredientsDoubled(t *testing.T) {
	var r io.Reader = strings.NewReader("1q mayo\n  1 c brown sugar\n 1.3333333333333333T cayenne\n")
	var correct_parsed []TokenInstance
	correct_parsed = []TokenInstance{TokenInstance{INTEGER, "1"},
		TokenInstance{KW_QUART, "q"}, TokenInstance{WS, " "},
		TokenInstance{WORD, "mayo"}, TokenInstance{NEWLINE, "\n"},
		TokenInstance{WS, "  "}, TokenInstance{INTEGER, "1"},
		TokenInstance{WS, " "}, TokenInstance{KW_CUP, "c"},
		TokenInstance{WS, " "}, TokenInstance{WORD, "brown"},
		TokenInstance{WS, " "}, TokenInstance{WORD, "sugar"},
		TokenInstance{NEWLINE, "\n"}, TokenInstance{WS, " "},
		TokenInstance{INTEGER, "1"}, TokenInstance{PERIOD, "."},
		TokenInstance{INTEGER, "3333333333333333"}, TokenInstance{KW_TABLESPOON, "T"},
		TokenInstance{WS, " "}, TokenInstance{WORD, "cayenne"},
		TokenInstance{NEWLINE, "\n"}}

	s := NewScanner(r)
	for _, ti := range correct_parsed {
		token, literal := s.Scan()
		if token != ti.Type {
			t.Error("Expected ", ti.Type, " but got ", token, " with value ", literal)
		}
		if literal != ti.Literal {
			t.Error("Expected ", ti.Literal, " but got ", literal, " with type ", token)
		}
	}
	token, _ := s.Scan()
	if token != EOF {
		t.Error("Expected EOF, but got ", token)
	}

}
