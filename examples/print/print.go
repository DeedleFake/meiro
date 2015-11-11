package main

import (
	"fmt"
	"github.com/DeedleFake/meiro"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	m := meiro.Random(5, 5)

	fmt.Printf("|")
	for x := 0; x < m.Width(); x++ {
		fmt.Printf("--|")
	}
	fmt.Println()

	for y := 1; y < m.Height(); y++ {
		fmt.Printf("|")
		for x := 0; x < m.Width(); x++ {
			fmt.Printf("  ")

			if m.At(x, y).Right() == nil {
				fmt.Printf("|")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()

		fmt.Printf("|")
		for x := 0; x < m.Width(); x++ {
			if m.At(x, y).Down() == nil {
				fmt.Printf("--|")
			} else {
				fmt.Printf("  |")
			}
		}
		fmt.Println()
	}
}
