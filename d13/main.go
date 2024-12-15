package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

type Button struct {
	X int
	Y int
}

type Arcade struct {
	A     Button
	B     Button
	Prize Coordinate
}

func main() {
	os.Exit(run())
}

func run() int {
	input, err := readInput("input.txt")
	if err != nil {
		log.Println(err)
		return 1
	}

	var result int
	for _, a := range input {
		solution, err := a.SolveDirectSubstitution()
		if err != nil {
			continue
		}

		result += solution[0]*3 + solution[1]*1
	}

	fmt.Println(result)

	return 0
}

func (a *Arcade) SolveDirectSubstitution() ([]int, error) {
	x0, y0 := a.A.X, a.A.Y
	x1, y1 := a.B.X, a.B.Y
	cx, cy := a.Prize.X, a.Prize.Y

	Bdividend := (cy*x0 - cx*y0)
	Bdivisor := (y1*x0 - y0*x1)
	if Bdivisor == 0 {
		return nil, errors.New("Zero division for B. No solution!")
	} else if Bdividend%Bdivisor != 0 {
		return nil, errors.New("Non integer solution!")
	}

	B := Bdividend / Bdivisor
	Adividend := (cx - B*x1)
	Adivisor := x0
	if Adivisor == 0 {
		return nil, errors.New("Zero division for A. No solution!")
	} else if Adividend%Adivisor != 0 {
		return nil, errors.New("Non integer solution!")
	}
	A := Adividend / Adivisor

	return []int{A, B}, nil
}

func readInput(filepath string) ([]Arcade, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	result := make([]Arcade, 0)

	input := bufio.NewScanner(f)
	var arcade Arcade

	row := 1
	for input.Scan() {
		line := input.Text()

		if line == "" {
			arcade = Arcade{}
			continue
		}

		switch row {
		case 1:
			arcade.A.X, arcade.A.Y, err = readLine(line, "+")
			if err != nil {
				return nil, err
			}
		case 2:
			arcade.B.X, arcade.B.Y, err = readLine(line, "+")
			if err != nil {
				return nil, err
			}
		case 3:
			arcade.Prize.X, arcade.Prize.Y, err = readLine(line, "=")
			if err != nil {
				return nil, err
			}

			result = append(result, arcade)
			row = 0
		}
		row++
	}

	return result, nil
}

func readLine(line, separator string) (int, int, error) {
	line = strings.Split(line, ":")[1]
	values := strings.Split(line, ",")
	x, err := strconv.Atoi(strings.Split(values[0], separator)[1])
	if err != nil {
		return 0, 0, err
	}

	y, err := strconv.Atoi(strings.Split(values[1], separator)[1])
	if err != nil {
		return 0, 0, err
	}

	return x, y, nil
}
