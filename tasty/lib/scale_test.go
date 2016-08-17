package tasty

// Runs the given example for the 'tasty' problem

import (
	"testing"
)

func ExampleMultiplyRecipe() {
	original := Recipe{
		"mayo":        Amount{2, CUP},
		"brown sugar": Amount{8, TABLESPOON},
		"cayenne":     Amount{2, TEASPOON},
	}
	doubled := MultiplyRecipe(2, original)
    PrintRecipe(doubled)
    // Output:
    //1 c   brown sugar
    //1.3333333333333333 T   cayenne
    //1 q   mayo

}

func TestDoubling(t *testing.T) {
	// Sample Input
	var original = Recipe{
		"mayo":        Amount{2, CUP},
		"brown sugar": Amount{8, TABLESPOON},
		"cayenne":     Amount{2, TEASPOON},
	}

	var correct_doubled = Recipe{
		"mayo":        Amount{1, QUART},
		"brown sugar": Amount{1, CUP},
		"cayenne":     Amount{1.3333333333333333, TABLESPOON},
	}

	// Double the amounts
	var doubled = MultiplyRecipe(2, original)

	// Check all keys in doubled are correct
	for key, value := range doubled {
		if value != correct_doubled[key] {
			t.Error("For ", key, ", expected ", correct_doubled[key], " got ",
				value)
		}
	}

	// Check all keys in correct_doubled are correct
	// catches i.e. missing keys in doubled
	for key, value := range correct_doubled {
		if value != correct_doubled[key] {
			t.Error("For ", key, ", expected ", correct_doubled[key], " got ",
				value)
		}
	}
}
