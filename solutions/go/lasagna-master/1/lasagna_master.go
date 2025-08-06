package lasagna

func PreparationTime(layers []string, avgTime int) int {
	if avgTime != 0 {
		return len(layers) * avgTime
	} else {
		return len(layers) * 2
	}
}

func Quantities(layers []string) (int, float64) {
	var (
		layer   = ""
		sauce   = 0
		noodles = 0
	)

	for _, layer = range layers {
		if layer == "noodles" {
			noodles++
		} else if layer == "sauce" {
			sauce++
		}
	}

	return 50 * noodles, 0.2 * float64(sauce)
}

func AddSecretIngredient(friendsList []string, myRecipe []string) {
	myRecipe[len(myRecipe)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(twoPortionAmounts []float64, portions int) []float64 {
	var (
		amounts = make([]float64, len(twoPortionAmounts))
		factor  = float64(portions) / 2
		amount  = float64(0)
		i       = 0
	)

	for i, amount = range twoPortionAmounts {
		amounts[i] = amount * factor
	}

	return amounts
}
