package main

import (
	"fmt"
)

type cards [][]int

func main() {
	cards := getInputs()
	fmt.Println(cards)
	for i := 0; ; i++ {
		drawnNum := draw(i)
		fmt.Println("drew:", drawnNum)
		for numCard, card := range cards {
			fmt.Println("card:", card)
			for j := 0; j <= 4; j++ {
				if drawnNum == card[j] {
					fmt.Printf("\told card: %v, drawNum: %v, j: %v\n", card, drawnNum, j)
					card[j] = -1
					fmt.Println("\tnew card:", card)
					sum := 0
					for k := 0; k <= 4; k++ {
						sum = sum + card[k]
					}
					if sum == -5 {
						fmt.Printf("\tbingo! card: %v, drawNum: %v, j: %v\n", card, drawnNum, j)
						fmt.Println("\tcard number:", numCard)
						originalCard := getInputs()[numCard]
						fmt.Println("original card:", originalCard)
						return
					}
				}
			}
		}
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
			21, 9, 14, 16, 7, // this is first one to mark bingo
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
