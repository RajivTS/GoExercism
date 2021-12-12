package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgTime int) int {
	if avgTime == 0 {
		avgTime = 2
	}
	return len(layers) * avgTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var noodlesQuantity int
	var sauceQuantity float64
	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodlesQuantity += 50
		case "sauce":
			sauceQuantity += 0.2
		}
	}
	return noodlesQuantity, sauceQuantity
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) []string {
	finalList := make([]string, len(myList))
	copy(finalList, myList)
	finalList = append(finalList, friendsList[len(friendsList)-1])
	return finalList
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, numPortions int) []float64 {
	scaledQuantities := make([]float64, len(quantities))
	copy(scaledQuantities, quantities)
	for i := 0; i < len(scaledQuantities); i++ {
		scaledQuantities[i] /= 2
		scaledQuantities[i] *= float64(numPortions)
	}
	return scaledQuantities
}
