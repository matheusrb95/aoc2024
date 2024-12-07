package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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

	var xmasMatrix [][]byte
	var result int

	input := bufio.NewScanner(f)
	for input.Scan() {
		xmasMatrix = append(xmasMatrix, []byte(input.Text()))
	}

	for i := 0; i < len(xmasMatrix); i++ {
		for j := 0; j < len(xmasMatrix[i]); j++ {
			if xmasMatrix[i][j] == 'X' {
				// horizontal
				if j < len(xmasMatrix[i])-3 {
					if xmasMatrix[i][j+1] == 'M' {
						if xmasMatrix[i][j+2] == 'A' {
							if xmasMatrix[i][j+3] == 'S' {
								result++
							}
						}
					}
				}

				// horizontal reverted
				if j >= 3 {
					if xmasMatrix[i][j-1] == 'M' {
						if xmasMatrix[i][j-2] == 'A' {
							if xmasMatrix[i][j-3] == 'S' {
								result++
							}
						}
					}
				}

				// vertical
				if i < len(xmasMatrix)-3 {
					if xmasMatrix[i+1][j] == 'M' {
						if xmasMatrix[i+2][j] == 'A' {
							if xmasMatrix[i+3][j] == 'S' {
								result++
							}
						}
					}
				}

				// vertical reverted
				if i >= 3 {
					if xmasMatrix[i-1][j] == 'M' {
						if xmasMatrix[i-2][j] == 'A' {
							if xmasMatrix[i-3][j] == 'S' {
								result++
							}
						}
					}
				}

				// diagonal right up
				if i >= 3 && j < len(xmasMatrix[i])-3 {
					if xmasMatrix[i-1][j+1] == 'M' {
						if xmasMatrix[i-2][j+2] == 'A' {
							if xmasMatrix[i-3][j+3] == 'S' {
								result++
							}
						}
					}
				}

				// diagonal left up
				if i >= 3 && j >= 3 {
					if xmasMatrix[i-1][j-1] == 'M' {
						if xmasMatrix[i-2][j-2] == 'A' {
							if xmasMatrix[i-3][j-3] == 'S' {
								result++
							}
						}
					}
				}

				// diagonal right down
				if i < len(xmasMatrix)-3 && j >= 3 {
					if xmasMatrix[i+1][j-1] == 'M' {
						if xmasMatrix[i+2][j-2] == 'A' {
							if xmasMatrix[i+3][j-3] == 'S' {
								result++
							}
						}
					}
				}

				// diagonal left down
				if i < len(xmasMatrix)-3 && j < len(xmasMatrix[i])-3 {
					if xmasMatrix[i+1][j+1] == 'M' {
						if xmasMatrix[i+2][j+2] == 'A' {
							if xmasMatrix[i+3][j+3] == 'S' {
								result++
							}
						}
					}
				}
			}
		}
	}

	fmt.Println(result)

	return 0
}
