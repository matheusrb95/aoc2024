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

	lenRow := len(xmasMatrix[0])
	lenColumn := len(xmasMatrix)

	for i := 0; i < lenColumn; i++ {
		for j := 0; j < lenRow; j++ {
			if j >= 1 && j < lenRow-1 && i >= 1 && i < lenColumn-1 {
				if xmasMatrix[i][j] == 'A' {
					if (leftUp(xmasMatrix, i, j) == 'M' && rightDown(xmasMatrix, i, j) == 'S') || (leftUp(xmasMatrix, i, j) == 'S' && rightDown(xmasMatrix, i, j) == 'M') {
						if (rightUp(xmasMatrix, i, j) == 'M' && leftDown(xmasMatrix, i, j) == 'S') || (rightUp(xmasMatrix, i, j) == 'S' && leftDown(xmasMatrix, i, j) == 'M') {
							result++
						}
					}
				}
			}
		}
	}

	fmt.Println(result)

	return 0
}

func rightUp(matrix [][]byte, i, j int) byte {
	return matrix[i-1][j+1]
}

func rightDown(matrix [][]byte, i, j int) byte {
	return matrix[i+1][j+1]
}

func leftUp(matrix [][]byte, i, j int) byte {
	return matrix[i-1][j-1]
}

func leftDown(matrix [][]byte, i, j int) byte {
	return matrix[i+1][j-1]
}
