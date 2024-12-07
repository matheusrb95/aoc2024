package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("d1/list.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var l1, l2 []int

	input := bufio.NewScanner(f)
	for input.Scan() {
		s := strings.Split(input.Text(), "   ")

		x, err := strconv.Atoi(s[0])
		if err != nil {
			fmt.Print(err)
		}
		l1 = append(l1, x)

		x, err = strconv.Atoi(s[1])
		if err != nil {
			fmt.Print(err)
		}
		l2 = append(l2, x)
	}

	quickSortStart(l1)
	quickSortStart(l2)

	fmt.Printf("Part1: %d\n", sumDifBetweenSlices(l1, l2))
	fmt.Printf("Part2: %d\n", similarityScore(l1, l2))
}

func similarityScore(l1, l2 []int) int {
	m1 := make(map[int]int)
	m2 := make(map[int]int)

	for _, v1 := range l1 {
		m1[v1]++
	}

	for _, v2 := range l2 {
		if m1[v2] != 0 {
			m2[v2]++
		}
	}

	var sum int
	for k, v := range m1 {
		sum += k * v * m2[k]
	}

	return sum
}

func sumDifBetweenSlices(l1, l2 []int) (sum int) {
	for i, v := range l1 {
		sum += abs(v - l2[i])
	}

	return
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}

func quickSortStart(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]

	i := low - 1

	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			swap(&arr[i], &arr[j])
		}
	}

	swap(&arr[i+1], &arr[high])

	return i + 1
}

func swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}
