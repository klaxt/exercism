package lasagnamaster

// site requires lasagna name
// conflicts with existing lasagna test also with package
//package lasagna

func PreparationTime(layers []string, prepTime int) (estimate int) {
	time := 2
	if prepTime > 0 {
		time = prepTime
	}
	return len(layers) * time
}

const (
	NOODLES_PER_LAYER = 50
	SAUCE_PER_LAYER   = 0.2
)

func Quantities(layers []string) (noodles int, sauce float64) {
	noodleLayers, sauceLayers := 0, 0
	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodleLayers++
		}
		if layer == "sauce" {
			sauceLayers++
		}
	}
	noodles = noodleLayers * NOODLES_PER_LAYER
	sauce = float64(sauceLayers) * SAUCE_PER_LAYER
	return
}

func AddSecretIngredient(secretIngredients []string, ingredients []string) {
	ingredients[len(ingredients)-1] = secretIngredients[len(secretIngredients)-1]
}

const PORTIONS = 2

func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaledQuantities := []float64{}
	for _, q := range quantities {
		scaledQuantities = append(scaledQuantities, q*float64(portions)/float64(PORTIONS))
	}
	return scaledQuantities
}
