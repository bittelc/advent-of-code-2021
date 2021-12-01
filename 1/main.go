package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	prior := 100000
	answer := 0
	for scanner.Scan() {
		fmt.Println("lines:", scanner.Text())
		next, err := strconv.Atoi(scanner.Text())
		check(err)
		if next > prior {
			answer++
		}
		prior = next
	}
	fmt.Println("answer:", answer)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
