package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const exp = `mul\(([0-9]{1,3}),([0-9]{1,3})\)|do\(\)|don\'t\(\)`

func main() {
	os.Exit(run())
}

func run() int {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Println(err)
		return 1
	}
	defer f.Close()

	var sum int

	re := regexp.MustCompile(exp)

	input := bufio.NewScanner(f)
	var memory []string
	for input.Scan() {
		matches := re.FindAll(input.Bytes(), -1)
		for _, match := range matches {
			memory = append(memory, string(match))
		}
	}

	fmt.Println("%q\n", memory)

	do := true
	for _, value := range memory {
		if value == "don't()" {
			do = false
			continue
		}

		if value == "do()" {
			do = true
			continue
		}

		if !do {
			continue
		}

		records := strings.Split(value[4:len(value)-1], ",")
		recordInt, err := stringsToIntegers(records)
		if err != nil {
			log.Println(err)
			return 1
		}

		fmt.Printf("%q\n", value)
		sum += recordInt[0] * recordInt[1]
	}

	fmt.Println(sum)

	return 0
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
