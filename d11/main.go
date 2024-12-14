package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var blinkTimes int = 75

type NumAndBlinks struct {
	Num    int
	Blinks int
}

func main() {
	os.Exit(run())
}

func run() int {
	input, err := readInput("input.txt")
	if err != nil {
		log.Println(err)
		return 1
	}

	count := make(map[NumAndBlinks]int)

	var result int
	for _, num := range input {
		result += calc(num, blinkTimes, count)
	}

	fmt.Println(result)

	return 0
}

func calc(num, blinks int, count map[NumAndBlinks]int) int {
	if blinks == 0 {
		return 1
	}

	if total := count[NumAndBlinks{Num: num, Blinks: blinks}]; total > 0 {
		return total
	}

	var result int
	switch rule(num) {
	case 1:
		result = calc(1, blinks-1, count)
	case 2:
		first, second := split(num)
		result = calc(first, blinks-1, count) + calc(second, blinks-1, count)
	default:
		result = calc(num*2024, blinks-1, count)
	}

	count[NumAndBlinks{Num: num, Blinks: blinks}] = result
	return result
}

func rule(num int) int {
	if num == 0 {
		return 1
	}

	numStr := strconv.Itoa(num)
	if len(numStr)%2 == 0 {
		return 2
	}

	return 0
}

func split(num int) (int, int) {
	strNum := strconv.Itoa(num)
	mid := len(strNum) / 2
	firstHalf, secondHalf := strNum[:mid], strNum[mid:]

	firstInt, _ := strconv.Atoi(firstHalf)
	secondInt, _ := strconv.Atoi(secondHalf)

	return firstInt, secondInt
}

func readInput(filepath string) ([]int, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	result := make([]int, 0)
	input := bufio.NewScanner(f)
	for input.Scan() {
		result, err = stringToIntegers(strings.Split(input.Text(), " "))
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func stringToIntegers(line []string) ([]int, error) {
	result := make([]int, 0, len(line))
	for _, num := range line {
		n, err := strconv.Atoi(string(num))
		if err != nil {
			return nil, err
		}
		result = append(result, n)
	}
	return result, nil
}
