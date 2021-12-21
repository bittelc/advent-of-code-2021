package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type row []int
type card []row
type cards []card

func main() {
	crds := getInputs()
	fmt.Println("crds:", crds)
	for i := 0; ; i++ {
		drewNum := draw(i)
		fmt.Println("drewNum:", drewNum)
		// markCards(&crds, drewNum)
		// winner, crd := checkForRow(crds)
		// if winner {
		// 	fmt.Println("there's a row winner from card %v containing %v", crd, i)
		// }
		// winner, crd = checkForColumn(crds)
		// if winner {
		// 	fmt.Println("there's a column winner from card %v containing %v", crd, i)
		// }
	}
}

func markCards(crds *cards, drewNum int) {
	for _, crd := range *crds {
		for _, rw := range crd {
			for _, val := range rw {
				fmt.Println("val:", val)
			}
		}
	}
}

func draw(i int) int {
	inputs := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}
	return inputs[i]
}

func getInputs() cards {
	crds := cards{}
	file, err := os.Open("test_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	crd := card{}
	rw := row{}
	for scanner.Scan() {
		line := scanner.Text()
		rw = nil
		rw = parseRow(line)
		fmt.Println("rw", rw)
		if len(rw) == 0 {
			crds = append(crds, crd)
			crd = nil
			continue
		}
		crd = append(crd, rw)
	}
	crds = append(crds, crd)
	return crds
}

func parseRow(s string) row {
	rw := row{}
	parsedS := strings.Fields(s)
	for _, val := range parsedS {
		i, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalf("error converting string value in row to int: %v", err)
		}
		rw = append(rw, i)
	}
	return rw
}
