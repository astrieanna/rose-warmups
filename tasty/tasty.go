package main

import (
	"bufio"
	"fmt"
	"github.com/astrieanna/rose-warmups/tasty/lib"
	"os"
)

// Tasty: Scaling a Recipe
func main() {
	// Open File
	filename := os.Args[1] // First arg
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)

	// Parse File
	parser := tasty.NewParser(reader)
	pr, err := parser.Parse()
	if err != nil {
		panic(err)
	}
	parsed_recipe := *pr

	// Double Recipe
	scaled_recipe := tasty.MultiplyRecipe(2, parsed_recipe)

	// Print Results
	fmt.Println("Original Recipe")
	tasty.PrintRecipe(parsed_recipe)
	fmt.Println()
	fmt.Println("Doubled Recipe")
	tasty.PrintRecipe(scaled_recipe)
}
