package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var drawNumList []int = []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}

type space map[int]bool
type row []space
type card []row
type cards []card

func main() {
	crds := getInputs()
	for i := 0; i < len(drawNumList); i++ {
		drewNum := draw(i)
		fmt.Println("drew number:", drewNum)
		markCards(&crds, drewNum)
		winner, crd := checkForRow(&crds)
		if winner {
			fmt.Printf("there's a row winner from card %+v containing %v\n", crd, drewNum)
			return
		}
		// winner, crd = checkForColumn(crds)
		// if winner {
		// 	fmt.Println("there's a column winner from card %v containing %v", crd, i)
		// }
	}
}

func checkForColumn(crds *cards) (bool, card) {
	for _, crd := range *crds {
		for i := 0; i <= 4; i++ {
			keyval := 
		}
	}
	return false, nil
}

func checkForRow(crds *cards) (bool, card) {
	for _, crd := range *crds {
	rww:
		for _, rw := range crd {
			for _, keyval := range rw {
				for _, val := range keyval {
					if val == false {
						break rww
					}
				}
			}
			return true, crd
		}
	}
	return false, nil
}

func markCards(crds *cards, drewNum int) {
	for _, crd := range *crds {
		for _, rw := range crd {
			for _, val := range rw {
				for key := range val {
					if drewNum == key {
						val[drewNum] = true
					}
				}
			}
		}
	}
}

func draw(i int) int {
	return drawNumList[i]
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
	spce := space{}
	parsedS := strings.Fields(s)
	for _, val := range parsedS {
		spce = space{}
		i, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalf("error converting string value in row to int: %v", err)
		}
		spce[i] = false
		rw = append(rw, spce)
	}
	return rw
}
