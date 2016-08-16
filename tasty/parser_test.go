package tasty

import (
	"io"
	"strings"
	"testing"
)

func TestParsingIngredients(t *testing.T) {
	var r io.Reader = strings.NewReader("2c mayo\n  8 T brown sugar\n 2t cayenne\n")
	correct_recipe := Recipe{
		"mayo":        Amount{2, CUP},
		"brown sugar": Amount{8, TABLESPOON},
		"cayenne":     Amount{2, TEASPOON},
	}

	var p *Parser = NewParser(r)
	pr, err := p.Parse()
	if err != nil {
		t.Error("Got parsing error: ", err)
		return
	}
	parsed_recipe := *pr
	for key, value := range correct_recipe {
		if parsed_recipe[key] != value {
			t.Error("Expected ", value, ", but got ", parsed_recipe[key])
		}
	}
	for key, value := range parsed_recipe {
		if correct_recipe[key] != value {
			t.Error("Expected ", correct_recipe[key], ", but got ", value)
		}
	}
}

func TestParsingIngredientsDoubled(t *testing.T) {
	var r io.Reader = strings.NewReader("1q mayo\n  1 c brown sugar\n 1.3333333333333333T cayenne\n")
	correct_recipe := Recipe{
		"mayo":        Amount{1, QUART},
		"brown sugar": Amount{1, CUP},
		"cayenne":     Amount{1.3333333333333333, TABLESPOON},
	}

	var p *Parser = NewParser(r)
	pr, err := p.Parse()
    if err != nil {
		t.Error("Got parsing error: ", err)
		return
	}
	parsed_recipe := *pr
	for key, value := range correct_recipe {
		if parsed_recipe[key] != value {
			t.Error("Expected ", value, ", but got ", parsed_recipe[key])
		}
	}
	for key, value := range parsed_recipe {
		if correct_recipe[key] != value {
			t.Error("Expected ", correct_recipe[key], ", but got ", value)
		}
	}

}