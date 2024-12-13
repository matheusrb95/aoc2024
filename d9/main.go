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

	a := make(map[int][]int)
	b := make(map[int][]int)

	for i := 0; i < len(disk); i++ {
		for j := 0; j < disk[i]; j++ {
			a[i] = append(a[i], i)
		}

		for j := 0; j < free[i]; j++ {
			b[i] = append(b[i], -1)
		}
	}

	for i := 0; i < len(free); i++ {
		for j := len(disk) - 1; j > 0; j-- {
			if _, ok := disk[j]; !ok {
				continue
			}

			if disk[j] > free[i] {
				continue
			}

			free[i] -= disk[j]

			var temp1, temp2 []int
			for k := 0; k < disk[j]; k++ {
				temp1 = append(temp1, -1)
			}

			for k := 0; k < free[i]; k++ {
				temp2 = append(temp2, -1)
			}

			a[i] = append(a[i], a[j]...)
			a[j] = temp1
			b[i] = temp2

			delete(disk, j)
		}
	}

	var format []int

	for x := 0; x < len(a); x++ {
		format = append(format, a[x]...)
		format = append(format, b[x]...)
	}

	var result int
	for i, num := range format {
		if num == -1 {
			continue
		}

		result += i * num
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
