package shuffle

import (
	"math/rand"
	"sort"
)

// Shuffle sorts data in a randomized order. It uses the default
// Source in math/rand, so if clients want to manipulate the outcome,
// they should call the appropriate functions in math/rand.
//
// TODO: Add ability to use custom Sources.
func Shuffle(data sort.Interface) {
	length := data.Len()

	for i := 0; i < length; i++ {
		i2 := rand.Intn(i + 1)
		data.Swap(i, i2)
	}
}
