package main

import (
  "fmt"
  "sort"
  "recipeparser/lexer"
)
// Tasty: Scaling a Recipe
// input: tasty(2, {'mayo', (2, 'c'),
//                  'brown sugar', (8, 'T'),
//                  'cayenne', (2, 't')})
// output: {'mayo': (1,'q'),
//          'brown sugar', (1,'c'),
//          'cayenne', (1.333..., 'T') }


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

func print_recipe(recipe Recipe, ingredients []string) {
	for _, key := range ingredients {
		fmt.Println(recipe[key].Quantity, recipe[key].Unit, " ", key)
	}
}

func multiply_recipe(multiple float64, original Recipe) Recipe {
	var scaled = make(Recipe)
	for key, value := range original {
		var n Amount = value
		n.Quantity = multiple * value.Quantity
		n = normalize_amount(n)
		scaled[key] = n
	}
	return scaled
}

// Conversions
// 3 t = 1 T
// 16 T = 1 c
// 4 c = 1 q
func normalize_amount(a Amount) Amount {
	if a.Unit == TEASPOON && a.Quantity >= 3 {
		a.Unit = TABLESPOON
		a.Quantity = a.Quantity / 3
	}

	if a.Unit == TABLESPOON && a.Quantity >= 16 {
		a.Unit = CUP
		a.Quantity = a.Quantity / 16
	}

	if a.Unit == CUP && a.Quantity >= 4 {
		a.Unit = QUART
		a.Quantity = a.Quantity / 4
	}
	return a
}

func main() {
	// Sample Input
	var original Recipe = make(Recipe)
	original["mayo"] = Amount{2, "c"}
	original["brown sugar"] = Amount{8, "T"}
	original["cayenne"] = Amount{2, "t"}

	// Double the amounts
	var doubled = multiply_recipe(2, original)

	// Arrange ingredients in sorted order
	var keys []string
	for k := range original {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Print out results
	fmt.Println("Rose's Recipe")
	print_recipe(original, keys)
	fmt.Println()
	fmt.Println("The Doubled Recipe")
	print_recipe(doubled, keys)
}
