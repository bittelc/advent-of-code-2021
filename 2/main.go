package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const f = "forward"
const d = "down"
const u = "up"

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()
	depth := 0
	horizontal := 0
	aim := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		switch {
		case strings.Contains(scanner.Text(), f):
			i := parseInt(scanner.Text(), f)
			horizontal = horizontal + i
			depth = depth + aim*i
		case strings.Contains(scanner.Text(), d):
			// depth = depth + parseInt(scanner.Text(), d)
			aim = aim + parseInt(scanner.Text(), d)
		case strings.Contains(scanner.Text(), u):
			// depth = depth - parseInt(scanner.Text(), u)
			aim = aim - parseInt(scanner.Text(), u)
		}
	}
	fmt.Println("horizontal", horizontal)
	fmt.Println("depth", depth)
	fmt.Println("multiplied:", horizontal*depth)
}

func parseInt(fullText, dropText string) int {
	s := strings.TrimPrefix(fullText, fmt.Sprintf("%v ", dropText))
	i, err := strconv.Atoi(s)
	check(err)
	return i
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
