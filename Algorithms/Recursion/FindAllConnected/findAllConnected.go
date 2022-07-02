package main

import (
	"fmt"
	"math"
)

func findConnects(matrix [][]int, M, N, r, c int) int {
	answer := 0
	if r < 0 || c < 0 || r >= M || c >= N {
		answer = 0
	} else if matrix[r][c] == 1 {
		matrix[r][c] = 0
		answer = 1 +
			findConnects(matrix, M, N, r-1, c) +
			findConnects(matrix, M, N, r+1, c) +
			findConnects(matrix, M, N, r, c-1) +
			findConnects(matrix, M, N, r, c+1) +
			findConnects(matrix, M, N, r-1, c-1) +
			findConnects(matrix, M, N, r-1, c+1) +
			findConnects(matrix, M, N, r+1, c+1) +
			findConnects(matrix, M, N, r+1, c-1)
	}
	return answer
}

func findMaxConnects(matrix [][]int, M, N int) int {
	maxConnects := 0
	for r := 0; r < M; r++ {
		for c := 0; c < N; c++ {
			maxConnects = int(math.Max(float64(maxConnects), float64(findConnects(matrix, M, N, r, c))))
		}
	}
	return maxConnects
}

func main() {
	matrix := [][]int{{1, 1, 0, 0, 0}, {0, 1, 1, 0, 0}, {0, 0, 1, 0, 1}, {1, 0, 0, 0, 1}, {0, 1, 0, 1, 1}}
	count := findMaxConnects(matrix, 5, 5)
	fmt.Println(count)
}
