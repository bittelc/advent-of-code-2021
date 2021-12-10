package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type counter struct {
	zero int
	one  int
}

func main() {
	gammaBinary := calculateGammeRateBinary()
	gammaDec := calculateDecFromBin(gammaBinary)
	epsilonBinary := calculateEpsilonRateBinary()
	epsilonDec := calculateDecFromBin(epsilonBinary)
	fmt.Println("decimal result:", gammaDec*epsilonDec)
}

func calculateGammeRateBinary() string {
	gammaRateBin := ""
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// optionally, resize scanner's capacity for lines over 64K, see next example
	count := counter{}
	// for i := 0; i <= 11; i++ {
	for i := 0; i <= 11; i++ {
		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			log.Fatal(err)
		}
		count.zero = 0
		count.one = 0
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			char := line[i : i+1]
			if char == "0" {
				count.zero++
			} else if char == "1" {
				count.one++
			} else {
				log.Fatalln("Yo none of it matched, problem")
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		if count.zero > count.one {
			gammaRateBin = gammaRateBin + "0"
		} else if count.zero < count.one {
			gammaRateBin = gammaRateBin + "1"
		} else {
			log.Fatalln("uh oh, equal number of 1s and 0s in the column, count:", count)
		}
		fmt.Println("gammaRateBin:", gammaRateBin)
	}

	return gammaRateBin
}

func calculateEpsilonRateBinary() string {
	epsilonRateBin := ""
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// optionally, resize scanner's capacity for lines over 64K, see next example
	count := counter{}
	// for i := 0; i <= 11; i++ {
	for i := 0; i <= 11; i++ {
		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			log.Fatal(err)
		}
		count.zero = 0
		count.one = 0
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			char := line[i : i+1]
			if char == "1" {
				count.zero++
			} else if char == "0" {
				count.one++
			} else {
				log.Fatalln("Yo none of it matched, problem")
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		if count.zero > count.one {
			epsilonRateBin = epsilonRateBin + "0"
		} else if count.zero < count.one {
			epsilonRateBin = epsilonRateBin + "1"
		} else {
			log.Fatalln("uh oh, equal number of 1s and 0s in the column, count:", count)
		}
		fmt.Println("epsilonRateBin:", epsilonRateBin)
	}

	return epsilonRateBin
}

func calculateDecFromBin(num string) int64 {
	dec, err := strconv.ParseInt(num, 2, 64)
	if err != nil {
		log.Fatalln(err)
	}
	return dec
}
