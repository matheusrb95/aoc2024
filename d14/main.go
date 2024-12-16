package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	width  int = 101
	height int = 103
	times  int = 10_000
)

type Position struct {
	X int
	Y int
}

type Velocity struct {
	X int
	Y int
}

type Robot struct {
	Pos Position
	Vel Velocity
}

type Quad struct {
	Positions []Position
}

func main() {
	os.Exit(run())
}

func (r *Robot) move() {
	r.Pos.X = r.Pos.X + r.Vel.X
	if r.Pos.X > width-1 {
		r.Pos.X = r.Pos.X - width
	} else if r.Pos.X < 0 {
		r.Pos.X = width + r.Pos.X
	}

	r.Pos.Y = r.Pos.Y + r.Vel.Y
	if r.Pos.Y > height-1 {
		r.Pos.Y = r.Pos.Y - height
	} else if r.Pos.Y < 0 {
		r.Pos.Y = height + r.Pos.Y
	}
}

func run() int {
	input, err := readInput("input.txt")
	if err != nil {
		log.Println(err)
		return 0
	}

	for i := 1; i <= times; i++ {
		bathroom := make(map[Position]int)
		for _, robot := range input {
			robot.move()
			bathroom[robot.Pos]++
		}

		if !overlaps(bathroom) {
			Print(i, input)
		}
	}

	// quads := NewQuads()

	// var robots int
	// result := 1

	// for _, quad := range quads {
	// 	for _, pos := range quad.Positions {
	// 		robots += bathroom[pos]
	// 	}

	// 	result *= robots
	// 	robots = 0
	// }

	// fmt.Println(result)

	return 0
}

func overlaps(bathroom map[Position]int) bool {
	for _, v := range bathroom {
		if v > 1 {
			return true
		}
	}

	return false
}

func Print(n int, robots []*Robot) {
	matrix := make([][]rune, height)

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			matrix[i] = append(matrix[i], '.')
		}
	}

	for _, robot := range robots {
		matrix[robot.Pos.X][robot.Pos.Y] = '#'
	}

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			fmt.Print(string(matrix[i][j]))
		}
		fmt.Println()
	}
	fmt.Printf("%d: ======================================================================================================\n", n)
}

func NewQuads() []Quad {
	x := width / 2
	y := height / 2

	var result []Quad

	positions := make([]Position, 0)
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			positions = append(positions, Position{X: i, Y: j})
		}
	}
	result = append(result, Quad{Positions: positions})

	positions = make([]Position, 0)
	for i := x + 1; i < width; i++ {
		for j := 0; j < y; j++ {
			positions = append(positions, Position{X: i, Y: j})
		}
	}
	result = append(result, Quad{Positions: positions})

	positions = make([]Position, 0)
	for i := 0; i < x; i++ {
		for j := y + 1; j < height; j++ {
			positions = append(positions, Position{X: i, Y: j})
		}
	}
	result = append(result, Quad{Positions: positions})

	positions = make([]Position, 0)
	for i := x + 1; i < width; i++ {
		for j := y + 1; j < height; j++ {
			positions = append(positions, Position{X: i, Y: j})
		}
	}
	result = append(result, Quad{Positions: positions})

	return result
}

func readInput(filepath string) ([]*Robot, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	result := make([]*Robot, 0)
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		values := strings.Split(line, " ")

		var robot Robot

		p := strings.Split(values[0], "=")
		pos := strings.Split(p[1], ",")
		robot.Pos.X, err = strconv.Atoi(pos[0])
		if err != nil {
			return nil, err
		}
		robot.Pos.Y, err = strconv.Atoi(pos[1])
		if err != nil {
			return nil, err
		}

		v := strings.Split(values[1], "=")
		vel := strings.Split(v[1], ",")
		robot.Vel.X, err = strconv.Atoi(vel[0])
		if err != nil {
			return nil, err
		}
		robot.Vel.Y, err = strconv.Atoi(vel[1])
		if err != nil {
			return nil, err
		}

		result = append(result, &robot)
	}

	return result, nil
}
