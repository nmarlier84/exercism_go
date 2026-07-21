package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avg int) int {
	if avg == 0 {
		avg = 2
	}

	return len(layers) * avg
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, v := range layers {
		switch v {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	newQuantities := append([]float64(nil), quantities...)
	for i, v := range newQuantities {
		newQuantities[i] = v * float64(portions) / 2.0
	}
	return newQuantities
}
