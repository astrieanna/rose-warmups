package tasty

// Common types and consts for representing a recipe
// Also, printing a recipe

import (
	"fmt"
	"sort"
)

type Unit int

const (
	TEASPOON Unit = iota
	TABLESPOON
	CUP
	QUART
)

type Amount struct {
	Quantity float64
	Unit     Unit
}
type Recipe map[string]Amount

func UnitToString(u Unit) string {
	switch u {
	case TEASPOON:
		return "t"
	case TABLESPOON:
		return "T"
	case CUP:
		return "c"
	case QUART:
		return "q"
	default:
		panic("unrecognized unit")
	}
}

func PrintRecipe(recipe Recipe) {
	// Arrange ingredients in sorted order
	var ingredients []string
	for k := range recipe {
		ingredients = append(ingredients, k)
	}
	sort.Strings(ingredients)

	// Print out amount for each ingredient
	for _, key := range ingredients {
		fmt.Println(recipe[key].Quantity, UnitToString(recipe[key].Unit), " ", key)
	}
}
