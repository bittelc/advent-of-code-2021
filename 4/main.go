package main

import (
	"fmt"
)

type cards [][]int

func main() {
	cards := getInputs()
	fmt.Println(cards)
	for card := range cards {

	}
	for i := 0; ; i++ {
		drawnNum := draw(i)
		fmt.Println("drew:", drawnNum)
		// for val := range cards {
		// 	if card == val {
		// 		card = 0
		// 	}
		// }
	}
}

func draw(i int) int {
	inputs := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}
	return inputs[i]

}

func getInputs() cards {
	input := cards{
		[]int{
			22, 13, 17, 11, 0,
		},
		{
			8, 2, 23, 4, 24,
		},
		{
			21, 9, 14, 16, 7,
		},
		{
			6, 10, 3, 18, 5,
		},
		{
			1, 12, 20, 15, 19,
		},
		{
			1, 12, 20, 15, 19,
		},
	}
	return input
}
