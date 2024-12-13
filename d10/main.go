package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var lenRow, lenCol int

type Pos struct {
	row int
	col int
}

func main() {
	os.Exit(run())
}

func run() int {
	topographicMap, err := readInput("input.txt")
	if err != nil {
		log.Println(err)
		return 1
	}

	lenRow, lenCol = len(topographicMap[0]), len(topographicMap)

	trailhead := make(map[Pos]int)

	for row := 0; row < len(topographicMap); row++ {
		for col := 0; col < len(topographicMap[row]); col++ {
			x := topographicMap[row][col]
			if x == 0 {
				pos := Pos{row: row, col: col}
				traceRoute(topographicMap, pos, 0, trailhead)
			}

		}
	}

	var result int
	for _, v := range trailhead {
		result += v
	}

	fmt.Println(result)

	return 0
}

func traceRoute(topographicMap [][]int, pos Pos, nextStep int, trailhead map[Pos]int) {
	if pos.col >= lenCol || pos.col < 0 || pos.row >= lenRow || pos.row < 0 {
		return
	}

	step := topographicMap[pos.row][pos.col]

	if step != nextStep {
		return
	}

	if nextStep == 9 {
		trailhead[pos]++
		return
	}

	// for i := -1; i <= 1; i++ {
	// 	for j := -1; j <= 1; j++ {
	// 		if i == 0 && j == 0 {
	// 			continue
	// 		}

	// 		newPos := Pos{row: pos.row + i, col: pos.col + j}
	// 		traceRoute(topographicMap, newPos, nextStep+1, trailhead)
	// 	}
	// }
	for _, dir := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		newPos := Pos{row: pos.row + dir[0], col: pos.col + dir[1]}
		traceRoute(topographicMap, newPos, nextStep+1, trailhead)
	}
}

func readInput(filepath string) ([][]int, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	result := make([][]int, 0)
	input := bufio.NewScanner(f)
	for input.Scan() {
		integers, err := stringToIntegers(input.Text())
		if err != nil {
			return nil, err
		}

		result = append(result, integers)
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
