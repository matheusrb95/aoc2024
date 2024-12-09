package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var lenInputX, lenInputY int

type Vector struct {
	y int
	x int
}

type Guard struct {
	Pos Vector
	Dir Vector
}

func main() {
	os.Exit(run())
}

func run() int {
	input, err := readInput("test.txt")
	if err != nil {
		log.Println(err)
		return 1
	}

	lenInputX = len(input[0])
	lenInputY = len(input)

	obstacles := obstacles(input)
	g := NewGuard(input)

	path := make(map[Vector]bool)
	g.calculateRoute(obstacles, path)
	fmt.Println(len(path))

	return 0
}

func (g *Guard) tracePath(obstacle Vector, path map[Vector]bool) {
	diffX := obstacle.x - g.Pos.x
	diffY := obstacle.y - g.Pos.y

	steps := max(abs(diffX), abs(diffY)) - 1

	for i := 1; i <= steps; i++ {
		newX := g.Pos.x + (i * g.Dir.x)
		newY := g.Pos.y + (i * g.Dir.y)

		path[Vector{x: newX, y: newY}] = true
		fmt.Println(newX, newY)
	}
}

func (g *Guard) calculateRoute(obstacles []Vector, path map[Vector]bool) {
	closestObstacle := g.closestObstacle(obstacles)
	if closestObstacle == (Vector{}) {
		var newX int
		var newY int

		if g.Dir.x == -1 {
			newX = 0
			newY = g.Pos.y
		} else if g.Dir.x == 1 {
			newX = lenInputX
			newY = g.Pos.y
		} else if g.Dir.y == -1 {
			newX = g.Pos.x
			newY = 0
		} else if g.Dir.y == 1 {
			newX = g.Pos.x
			newY = lenInputY
		}

		obstacle := Vector{x: newX, y: newY}
		g.tracePath(obstacle, path)
		return
	}

	g.tracePath(closestObstacle, path)
	g.Pos.x = closestObstacle.x - g.Dir.x
	g.Pos.y = closestObstacle.y - g.Dir.y
	g.rotateDir()
	g.calculateRoute(obstacles, path)

	return
}

func (g *Guard) closestObstacle(obstacles []Vector) Vector {
	result := Vector{}

	for _, obstacle := range obstacles {
		if g.Dir.y == 1 {
			if g.Pos.x == obstacle.x {
				// obstacle in oposite direction
				if obstacle.y < g.Pos.y {
					continue
				}

				if obstacle.y > result.y && result != (Vector{}) {
					continue
				}

				result = obstacle
			}
		} else if g.Dir.y == -1 {
			if g.Pos.x == obstacle.x {
				// obstacle in oposite direction
				if obstacle.y > g.Pos.y {
					continue
				}

				if obstacle.y < result.y && result != (Vector{}) {
					continue
				}

				result = obstacle
			}
		} else if g.Dir.x == -1 {
			if g.Pos.y == obstacle.y {
				// obstacle in oposite direction
				if obstacle.x > g.Pos.x {
					continue
				}

				if obstacle.x < result.x && result != (Vector{}) {
					continue
				}

				result = obstacle
			}
		} else if g.Dir.x == 1 {
			if g.Pos.y == obstacle.y {
				// obstacle in oposite direction
				if obstacle.x < g.Pos.x {
					continue
				}

				if obstacle.x > result.x && result != (Vector{}) {
					continue
				}

				result = obstacle
			}
		}
	}

	return result
}

func (g *Guard) rotateDir() {
	y := g.Dir.y
	g.Dir.y = g.Dir.x
	g.Dir.x = -y
}

func NewGuard(input [][]byte) *Guard {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == '^' {
				return &Guard{
					Pos: Vector{y: i, x: j},
					Dir: Vector{y: -1, x: 0},
				}
			} else if input[i][j] == '>' {
				return &Guard{
					Pos: Vector{y: i, x: j},
					Dir: Vector{y: 0, x: 1},
				}
			} else if input[i][j] == '<' {
				return &Guard{
					Pos: Vector{y: i, x: j},
					Dir: Vector{y: 0, x: -1},
				}
			} else if input[i][j] == 'v' {
				return &Guard{
					Pos: Vector{y: i, x: j},
					Dir: Vector{y: 1, x: 0},
				}
			}
		}
	}

	return nil
}

func obstacles(input [][]byte) []Vector {
	var result []Vector

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == '#' {
				result = append(result, Vector{y: i, x: j})
			}
		}
	}

	return result
}

func readInput(filepath string) ([][]byte, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var result [][]byte

	input := bufio.NewScanner(f)
	for input.Scan() {
		result = append(result, []byte(input.Text()))
	}

	return result, nil
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}
