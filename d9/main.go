package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	disk := make(map[int]int)
	free := make(map[int]int)

	var id int
	for i := 0; i < len(input); i++ {
		if even(i) {
			disk[id] = input[i]
			continue
		}

		free[id] = input[i]
		id++
	}

	block := make([]int, 0)

	for i := 0; i <= id; i++ {
		for j := 0; j < disk[i]; j++ {
			block = append(block, i)
		}

		for j := 0; j < free[i]; j++ {
			block = append(block, -1)
		}
	}

	for i := 0; i < len(block); i++ {
		if block[i] != -1 {
			continue
		}

		for j := len(block) - 1; j > 0; j-- {
			if block[j] == -1 {
				continue
			}

			block = swapElement(block, i, j)
			block = removeEndDots(block)
			break
		}
	}

	var result int
	for i := 0; i < len(block); i++ {
		result += i * block[i]
	}

	fmt.Println(result)

	return 0
}

func even(num int) bool {
	if num%2 != 0 {
		return false
	}

	return true
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
		result, err = stringToIntegers(input.Text())
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func stringToIntegers(line string) ([]int, error) {
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

func swapElement(slice []int, i int, j int) []int {
	r := slice[i]
	slice[i] = slice[j]
	slice[j] = r

	return slice
}

func removeEndDots(slice []int) []int {
	for j := len(slice) - 1; j > 0; j-- {
		if slice[j] != -1 {
			break
		}

		slice = slice[:j]

	}

	return slice
}
