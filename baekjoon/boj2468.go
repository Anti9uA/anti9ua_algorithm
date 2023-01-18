package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	r       = bufio.NewReader(os.Stdin)
	w       = bufio.NewWriter(os.Stdout)
	N       int
	rain    [][]int
	visited [][]bool
	// 상 하 좌 우
	dx = []int{-1, 1, 0, 0}
	dy = []int{0, 0, -1, 1}
)

func main() {
	defer w.Flush()

	fmt.Fscanln(r, &N)

	rain = make([][]int, N)
	visited = make([][]bool, N)

	answer := 1

	for i := 0; i < N; i++ {
		rain[i] = make([]int, N)
		visited[i] = make([]bool, N)
		for j := 0; j < N; j++ {
			fmt.Fscanf(r, "%d ", &rain[i][j])
		}
	}

	for i := 1; i < 101; i++ {
		height := 0
		reset()
		for y := 0; y < N; y++ {
			for x := 0; x < N; x++ {
				if rain[y][x] > i && !visited[y][x] {
					visited[y][x] = true
					dfs(x, y, i)
					height++
				}
			}
		}

		if height > answer {
			answer = height
		}

		if height == 0 {
			continue
		}
	}
	fmt.Fprint(w, answer)
}

// DFS를 통한 탐색
func dfs(x, y, h int) {
	for i := 0; i < 4; i++ {
		nx := x + dx[i]
		ny := y + dy[i]

		if nx < 0 || nx >= N || ny < 0 || ny >= N {
			continue
		}

		if visited[ny][nx] {
			continue
		}

		if rain[ny][nx] <= h {
			continue
		}

		visited[ny][nx] = true
		dfs(nx, ny, h)
	}
}

// reset visited
func reset() {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			visited[i][j] = false
		}
	}
}
