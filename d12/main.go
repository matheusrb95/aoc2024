package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Coordinate struct {
	X int
	Y int
}

type Garden struct {
	Field   []string
	Width   int
	Height  int
	Regions []map[Coordinate]bool
}

func main() {
	os.Exit(run())
}

func run() int {
	g, err := readInput("input.txt")
	if err != nil {
		log.Println(err)
		return 1
	}

	g.Regions = g.mapRegions()

	var result int
	for _, region := range g.Regions {
		result += len(region) * calculateSides(region)
	}

	fmt.Println(result)

	return 0
}

func calculateSides(region map[Coordinate]bool) int {
	sides := make([]map[Coordinate]bool, 0)

	for _, dir := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		side := make(map[Coordinate]bool)

		for coord := range region {
			newCoord := Coordinate{X: coord.X + dir[0], Y: coord.Y + dir[1]}
			if _, ok := region[newCoord]; ok {
				continue
			}

			side[coord] = true
		}
		sides = append(sides, side)
	}

	var result int

	for i, side := range sides {
		for coord := range side {
			if i < 2 {
				newCoord := Coordinate{X: coord.X, Y: coord.Y + 1}
				if _, ok := side[newCoord]; ok {
					continue
				}

				result++
			} else {
				newCoord := Coordinate{X: coord.X + 1, Y: coord.Y}
				if _, ok := side[newCoord]; ok {
					continue
				}

				result++
			}
		}
	}

	return result
}

func calculatePerimeter(region map[Coordinate]bool) int {
	boundaries := make(map[Coordinate]int)

	for coord := range region {
		for _, dir := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			newCoord := Coordinate{X: coord.X + dir[0], Y: coord.Y + dir[1]}
			if _, ok := region[newCoord]; !ok {
				continue
			}

			boundaries[coord]++
		}
	}

	var result int
	for coord := range region {
		switch boundaries[coord] {
		case 1:
			result += 3
		case 2:
			result += 2
		case 3:
			result += 1
		case 4:
			result += 0
		default:
			result += 4
		}

	}

	return result
}

func (g *Garden) mapRegions() []map[Coordinate]bool {
	result := make([]map[Coordinate]bool, 0)

	for i := 0; i < g.Height; i++ {

	FieldY:
		for j := 0; j < g.Width; j++ {
			for _, region := range result {
				if _, ok := region[Coordinate{X: i, Y: j}]; ok {
					continue FieldY
				}
			}

			plant := rune(g.Field[i][j])
			region := make(map[Coordinate]bool)
			g.resolveRegion(Coordinate{X: i, Y: j}, plant, region)

			result = append(result, region)
		}
	}

	return result
}

func (g *Garden) resolveRegion(c Coordinate, plant rune, region map[Coordinate]bool) {
	if c.X >= g.Height || c.X < 0 || c.Y >= g.Width || c.Y < 0 {
		return
	}

	if _, ok := region[c]; ok {
		return
	}

	if g.Field[c.X][c.Y] != byte(plant) {
		return
	}

	region[c] = true

	for _, dir := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		nextCoord := Coordinate{X: c.X + dir[0], Y: c.Y + dir[1]}
		g.resolveRegion(nextCoord, plant, region)
	}
}

func readInput(filepath string) (*Garden, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	field := make([]string, 0)
	input := bufio.NewScanner(f)
	for input.Scan() {
		field = append(field, input.Text())
	}

	result := Garden{
		Field:  field,
		Width:  len(field[0]),
		Height: len(field),
	}

	return &result, nil
}
