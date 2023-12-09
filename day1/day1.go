package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func part1() {
	file, err := os.Open("inputs/day1_1.txt")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var nums []string

	for fileScanner.Scan() {

		alreadyFirst := false
		var first string
		var last string

		for _, char := range fileScanner.Text() {

			_, err := strconv.ParseInt(string(char), 10, 64)
			if err != nil {
				continue
			}

			if !alreadyFirst {
				first = string(char)
				alreadyFirst = true
				continue
			}

			last = string(char)

		}

		if last == "" {
			last = first
		}

		nums = append(nums, first+last)

	}

	var sum int64

	for _, num := range nums {
		n, err := strconv.ParseInt(num, 10, 64)
		if err != nil {
			continue
		}

		sum += n

	}

	fmt.Println(sum)
}

const (
	One   = "one"
	Two   = "two"
	Three = "three"
	Four  = "four"
	Five  = "five"
	Six   = "six"
	Seven = "seven"
	Eight = "eight"
	Nine  = "nine"
)

func matchTextDigit(i int, line string) (string, error) {

	var num string

	if i+3 <= len(line) && line[i:i+3] == One {
		num = "1"
	} else if i+3 <= len(line) && line[i:i+3] == Two {
		num = "2"
	} else if i+5 <= len(line) && line[i:i+5] == Three {
		num = "3"
	} else if i+4 <= len(line) && line[i:i+4] == Four {
		num = "4"
	} else if i+4 <= len(line) && line[i:i+4] == Five {
		num = "5"
	} else if i+3 <= len(line) && line[i:i+3] == Six {
		num = "6"
	} else if i+5 <= len(line) && line[i:i+5] == Seven {
		num = "7"
	} else if i+5 <= len(line) && line[i:i+5] == Eight {
		num = "8"
	} else if i+4 <= len(line) && line[i:i+4] == Nine {
		num = "9"
	} else {
		return "", fmt.Errorf("not found")
	}

	return num, nil
}

func part2() {
	file, err := os.Open("inputs/day1_2.txt")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var nums []string

	for fileScanner.Scan() {

		alreadyFirst := false
		var first string
		var last string

		for i, char := range fileScanner.Text() {

			c := string(char)

			line := fileScanner.Text()

			_, err := strconv.ParseInt(c, 10, 64)
			if err != nil {
				c, err = matchTextDigit(i, line)
				if err != nil {
					continue
				}

			}

			if !alreadyFirst {
				first = c
				alreadyFirst = true
			} else {
				last = c
			}
		}

		if last == "" {
			last = first
		}

		nums = append(nums, first+last)

	}

	var sum int64

	for _, num := range nums {
		n, err := strconv.ParseInt(num, 10, 64)
		if err != nil {
			continue
		}

		sum += n

	}

	fmt.Println(sum)
}

func main() {

	//part1()

	part2()

}
