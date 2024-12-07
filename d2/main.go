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
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var result [][]string

	input := bufio.NewScanner(f)
	for input.Scan() {
		report := strings.Split(input.Text(), " ")

		reportInt, err := stringsToIntegers(report)
		if err != nil {
			log.Println(err)
			return 1
		}

		var element int
		len := len(reportInt)
		for i := 0; i <= len; i++ {
			if safe(reportInt) {
				result = append(result, report)
				break
			}

			if i > 0 {
				reportInt = restoreElement(reportInt, i-1, element)
			}

			if i < len {
				reportInt, element = removeElement(reportInt, i)
			} else {
				reportInt, element = removeElement(reportInt, i-1)
			}
		}
	}

	fmt.Println(result)

	return 0
}

func safe(slice []int) bool {
	var decreasing, increasing bool
	for i := 0; i < len(slice)-1; i++ {
		dif := slice[i] - slice[i+1]

		if dif < -3 || dif > 3 || dif == 0 {
			return false
		} else if dif > 0 {
			increasing = true
		} else {
			decreasing = true
		}

		if increasing && decreasing {
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

func removeElement(slice []int, index int) ([]int, int) {
	v := slice[index]
	return append(slice[:index], slice[index+1:]...), v
}

func restoreElement(slice []int, index, value int) []int {
	if len(slice) == index {
		return append(slice, value)
	}
	slice = append(slice[:index+1], slice[index:]...)
	slice[index] = value

	return slice
}
