package main

import (
	"fmt"
	"github.com/astrieanna/rose-warmups/tasty"
)

// Tasty: Scaling a Recipe
// input: tasty(2, {'mayo', (2, 'c'),
//                  'brown sugar', (8, 'T'),
//                  'cayenne', (2, 't')})
// output: {'mayo': (1,'q'),
//          'brown sugar', (1,'c'),
//          'cayenne', (1.333..., 'T') }

func main() {
	// Sample Input
	var original tasty.Recipe = make(tasty.Recipe)
	original["mayo"] = tasty.Amount{2, tasty.CUP}
	original["brown sugar"] = tasty.Amount{8, tasty.TABLESPOON}
	original["cayenne"] = tasty.Amount{2, tasty.TEASPOON}

	// Double the amounts
	var doubled = tasty.MultiplyRecipe(2, original)

	// Print out results
	fmt.Println("Rose's Recipe")
	tasty.PrintRecipe(original)
	fmt.Println()
	fmt.Println("The Doubled Recipe")
	tasty.PrintRecipe(doubled)
}
