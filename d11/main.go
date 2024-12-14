package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	os.Exit(run())
}

func run() int {
	input, err := readInput("input.txt")
	if err != nil {
		log.Println(err)
		return 1
	}

	blinkTimes := 25
	for i := 0; i < blinkTimes; i++ {
		result := make([]int, 0)
		lenInput := len(input)
		for j := 0; j < lenInput; j++ {
			switch rule(input[j]) {
			case 1:
				result = append(result, 1)
			case 2:
				slice, err := split(input, j)
				if err != nil {
					log.Println(err)
					return 1
				}
				result = append(result, slice...)
			default:
				result = append(result, input[j]*2024)
			}
		}

		input = result
	}

	fmt.Print(len(input))

	return 0
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

func split(slice []int, index int) ([]int, error) {
	strNum := strconv.Itoa(slice[index])
	mid := len(strNum) / 2
	firstHalf, secondHalf := strNum[:mid], strNum[mid:]

	firstInt, err := strconv.Atoi(firstHalf)
	if err != nil {
		return nil, err
	}
	secondInt, err := strconv.Atoi(secondHalf)
	if err != nil {
		return nil, err
	}

	return []int{firstInt, secondInt}, nil
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
