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
	f, err := os.Open("test.txt")
	if err != nil {
		log.Println(err)
		return 1
	}
	defer f.Close()

	var sequence bool
	var sequences [][]int
	rules := make(map[int][]int)

	input := bufio.NewScanner(f)
	for input.Scan() {
		if input.Text() == "" {
			sequence = true
			continue
		}

		if sequence {
			sequence := strings.Split(input.Text(), ",")
			sequenceInt, err := stringsToIntegers(sequence)
			if err != nil {
				log.Println(err)
				return 1
			}

			sequences = append(sequences, sequenceInt)
			continue
		}

		rule := strings.Split(input.Text(), "|")
		ruleInt, err := stringsToIntegers(rule)
		if err != nil {
			log.Println(err)
			return 1
		}

		rules[ruleInt[0]] = append(rules[ruleInt[0]], ruleInt[1])
	}

	var validSequences [][]int

	for _, sequence := range sequences {
		if valid(rules, sequence) {
			validSequences = append(validSequences, sequence)
		}
	}

	var result int
	for _, order := range validSequences {
		result += order[len(order)/2]
	}
	fmt.Println(result)

	return 0
}

func valid(rules map[int][]int, sequence []int) bool {
	for j := 1; j < len(sequence); j++ {
		if !contains(rules[sequence[j-1]], sequence[j]) {
			return false
		}
	}
	return true
}

func stringsToIntegers(lines []string) ([]int, error) {
	integers := make([]int, 0, len(lines))
	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		integers = append(integers, n)
	}
	return integers, nil
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
