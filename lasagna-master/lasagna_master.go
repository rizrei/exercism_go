package lasagnamaster

const TIME_PER_LAYER = 2

func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		timePerLayer = TIME_PER_LAYER
	}

	return len(layers) * timePerLayer
}

const GRAMS_PER_NOODLE = 50
const LITTERS_PER_SAUCE = 0.2

func Quantities(layers []string) (noodlesQuantity int, sauceQuantity float64) {
	noodlesQuantity = 0
	sauceQuantity = 0.0

	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodlesQuantity += GRAMS_PER_NOODLE
		case "sauce":
			sauceQuantity += LITTERS_PER_SAUCE
		}
	}

	return
}

func AddSecretIngredient(friendIngedients, myIngedients []string) {
	myIngedients = append(myIngedients[0:len(myIngedients)-1], friendIngedients[len(friendIngedients)-1])
}

func ScaleRecipe(quantities []float64, numberOfPortions int) []float64 {
	var scaledQuantities []float64

	for _, q := range quantities {
		scaledQuantities = append(scaledQuantities, q*float64(numberOfPortions)/2.0)
	}

	return scaledQuantities
}
