package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// var drawNumList []int = []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}

var drawNumList []int = []int{4, 77, 78, 12, 91, 82, 48, 59, 28, 26, 34, 10, 71, 89, 54, 63, 66, 75, 15, 22, 39, 55, 83, 47, 81, 74, 2, 46, 25, 98, 29, 21, 85, 96, 3, 16, 60, 31, 99, 86, 52, 17, 69, 27, 73, 49, 95, 35, 9, 53, 64, 88, 37, 72, 92, 70, 5, 65, 79, 61, 38, 14, 7, 44, 43, 8, 42, 45, 23, 41, 57, 80, 51, 90, 84, 11, 93, 40, 50, 33, 56, 67, 68, 32, 6, 94, 97, 13, 87, 30, 18, 76, 36, 24, 19, 20, 1, 0, 58, 62}

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
		for j := 0; j <= len(crds); j++ {
			winnerEle, _ := checkForRow(&crds)
			if winnerEle == -1 {
				continue
			}
			removeCrdFromCrds(winnerEle, &crds)
		}
		for j := 0; j <= len(crds); j++ {
			winnerEle, _ := checkForColumn(&crds)
			if winnerEle == -1 {
				continue
			}
			removeCrdFromCrds(winnerEle, &crds)
		}
		fmt.Println("crds length", len(crds))
		if len(crds) == 2 {
			fmt.Println("final crd:", drewNum*sumUnmarked(&crds[0]))
		}
	}
}

func removeCrdFromCrds(ele int, crds *cards) {
	*crds = append((*crds)[:ele], (*crds)[ele+1:]...)
}

func sumUnmarked(crd *card) int {
	sum := 0
	for _, rw := range *crd {
		for _, keyval := range rw {
			for key, val := range keyval {
				if val == false {
					sum = sum + key
				}
			}
		}
	}
	return sum
}

func checkForColumn(crds *cards) (int, card) {
	for j, crd := range *crds {
		for i := 0; i <= 4; i++ {
			consecutive := 0
		col:
			for _, rw := range crd {
				el := rw[i]
				for _, marked := range el {
					if marked == false {
						break col
					}
					consecutive++
					if consecutive == 5 {
						return j, crd
					}
				}
			}
		}
	}
	return -1, nil
}

func checkForRow(crds *cards) (int, card) {
	for i, crd := range *crds {
	rww:
		for _, rw := range crd {
			for _, keyval := range rw {
				for _, val := range keyval {
					if val == false {
						break rww
					}
				}
			}
			return i, crd
		}
	}
	return -1, nil
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
	file, err := os.Open("input.txt")
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
