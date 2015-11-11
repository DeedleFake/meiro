package shuffle

import (
	"math/rand"
	"sort"
)

func Shuffle(data sort.Interface) {
	length := data.Len()

	for i := 0; i < length; i++ {
		i2 := rand.Intn(i + 1)
		data.Swap(i, i2)
	}
}
