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
	antinodes := make([]Pos, 0)

	for _, v := range antennas {
		combinations := combinations(v)

		for _, combination := range combinations {
			antinodes = append(antinodes, combination[0])
			antinodes = append(antinodes, combination[1])
			ressonantHarmonic(combination[0], combination[1], &antinodes)
			ressonantHarmonic(combination[1], combination[0], &antinodes)
		}
	}

	for _, pos := range antinodes {
		uniqueAntinodes[pos] = true
	}

	fmt.Println(len(uniqueAntinodes))

	return 0
}

func ressonantHarmonic(a1, a2 Pos, antinodes *[]Pos) {
	antinode := Pos{row: a1.row + (a1.row - a2.row), col: a1.col + (a1.col - a2.col)}

	if antinode.row < 0 || antinode.col < 0 || antinode.row >= lenRow || antinode.col >= lenCol {
		return
	}

	*antinodes = append(*antinodes, antinode)
	ressonantHarmonic(antinode, a1, antinodes)
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
