package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Pos struct {
	row int
	col int
}

var (
	lenRow, lenCol int
)

func main() {
	os.Exit(run())
}

func run() int {
	antennas, err := readInput("input.txt")
	if err != nil {
		log.Println(err)
		return 1
	}

	uniqueAntinodes := make(map[Pos]bool)

	for _, v := range antennas {
		combinations := combinations(v)

		for _, combination := range combinations {
			antinodes := antinode(combination[0], combination[1])
			for _, pos := range antinodes {
				uniqueAntinodes[pos] = true
			}
		}
	}

	fmt.Println(uniqueAntinodes)

	var result int
	for pos := range uniqueAntinodes {
		if !insideBoundary(pos) {
			continue
		}

		result++
	}

	fmt.Println(result)

	return 0
}

func insideBoundary(pos Pos) bool {
	if pos.row < 0 || pos.col < 0 {
		return false
	}

	if pos.row >= lenRow || pos.col >= lenCol {
		return false
	}

	return true
}

func antinode(a1, a2 Pos) []Pos {
	result := make([]Pos, 0)
	diff := diffPoints(a1, a2)

	if (a1.row > a2.row && a1.col > a2.col) || (a2.row > a1.row && a2.col > a1.col) {
		result = append(result, Pos{row: a1.row - diff.row, col: a1.col - diff.col})
		result = append(result, Pos{row: a2.row + diff.row, col: a2.col + diff.col})

		return result
	}

	result = append(result, Pos{row: a1.row - diff.row, col: a1.col + diff.col})
	result = append(result, Pos{row: a2.row + diff.row, col: a2.col - diff.col})

	return result
}

func diffPoints(a, b Pos) Pos {
	return Pos{row: abs(a.row - b.row), col: abs(a.col - b.col)}
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}

func combinations(antennas []Pos) [][]Pos {
	result := make([][]Pos, 0)

	for i := 0; i < len(antennas)-1; i++ {
		for j := i + 1; j < len(antennas); j++ {
			combination := []Pos{antennas[i], antennas[j]}
			result = append(result, combination)
		}
	}

	return result
}

func readInput(filepath string) (map[rune][]Pos, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	result := make(map[rune][]Pos)

	input := bufio.NewScanner(f)
	var row int
	for input.Scan() {
		for col, c := range input.Text() {
			if c == '.' {
				continue
			}

			result[c] = append(result[c], Pos{row: row, col: col})
		}

		row++
	}

	lenRow, lenCol = row, row

	return result, nil
}
