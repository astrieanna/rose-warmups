package tasty

// Multiplies each amount in the Recipe by multiple, then normalizes the units.
func MultiplyRecipe(multiple float64, original Recipe) Recipe {
	var scaled = make(Recipe)
	for key, value := range original {
		var n Amount = value
		n.Quantity = multiple * value.Quantity
		n = normalizeAmount(n)
		scaled[key] = n
	}
	return scaled
}

// The unit will be upgraded if the amount is at least one whole of a higher unit.
// ## Conversions
// 3 t = 1 T
// 16 T = 1 c
// 4 c = 1 q
func normalizeAmount(a Amount) Amount {
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
