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
	totals, values, err := readInput("input.txt")
	if err != nil {
		log.Println(err)
		return 1
	}

	var sum int
	for i := 0; i < len(totals); i++ {
		if valid(totals[i], values[i], 0) {
			sum += totals[i]
		}
	}

	fmt.Println(sum)

	return 0
}

func valid(total int, values []int, mult int) bool {
	if len(values) == mult {
		return false
	}

	combinations := generateOperatorCombinations(len(values)-1, mult)
	for _, operators := range combinations {
		if total == evaluate(values, operators) {
			return true
		}
	}

	return valid(total, values, mult+1)
}

func evaluate(values []int, operators string) int {
	result := values[0]
	for i := 1; i < len(values); i++ {
		if operators[i-1] == '+' {
			result += values[i]
		} else if operators[i-1] == '*' {
			result *= values[i]
		}
	}

	return result
}

func generateOperatorCombinations(len int, mult int) []string {
	var combinations []string

	base := make([]rune, len)

	for i := 0; i < len; i++ {
		base[i] = '+'
	}

	generateCombinations(&combinations, base, mult, 0)

	return combinations
}

func generateCombinations(combinations *[]string, base []rune, mult, start int) {
	if mult == 0 {
		*combinations = append(*combinations, string(base))
		return
	}

	for i := start; i < len(base); i++ {
		base[i] = '*'
		generateCombinations(combinations, base, mult-1, i+1)
		base[i] = '+'
	}
}

func readInput(filepath string) ([]int, [][]int, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	var totals []int
	var values [][]int

	input := bufio.NewScanner(f)
	for input.Scan() {
		text := strings.Split(input.Text(), ":")
		total, err := strconv.Atoi(text[0])
		if err != nil {
			return nil, nil, err
		}

		totals = append(totals, total)

		value, err := stringsToIntegers(strings.Split(text[1], " "))
		if err != nil {
			return nil, nil, err
		}

		values = append(values, value)
	}

	return totals, values, nil
}

func stringsToIntegers(lines []string) ([]int, error) {
	integers := make([]int, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			continue
		}

		n, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		integers = append(integers, n)
	}
	return integers, nil
}
