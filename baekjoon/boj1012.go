package boj1012

import (
	"bufio"
	"fmt"
	"os"
)

var (
	r       = bufio.NewReader(os.Stdin)
	w       = bufio.NewWriter(os.Stdout)
	T       int
	M       int
	N       int
	K       int
	farm    [][]int
	visited [][]bool
	// 상 하 좌 우
	dx = []int{-1, 1, 0, 0}
	dy = []int{0, 0, -1, 1}
)

func main() {
	defer w.Flush()
	farm = make([][]int, 51)
	visited = make([][]bool, 51)
	for i := 0; i < 51; i++ {
		farm[i] = make([]int, 51)
		visited[i] = make([]bool, 51)
	}

	fmt.Fscanln(r, &T)
	for i := 0; i < T; i++ {
		fmt.Fscanln(r, &M, &N, &K)
		for i := 0; i < K; i++ {
			var x, y int
			fmt.Fscanln(r, &x, &y)
			farm[y][x] = 1
		}

		bug := 0
		for y := 0; y < N; y++ {
			for x := 0; x < M; x++ {
				if farm[y][x] == 1 && !visited[y][x] {
					visited[y][x] = true
					dfs(x, y)
					bug++
				}
			}
		}
		fmt.Fprintln(w, bug)
		reset()
	}
}

// DFS를 통한 탐색
func dfs(x, y int) {
	for i := 0; i < 4; i++ {
		nx := x + dx[i]
		ny := y + dy[i]

		if nx < 0 || nx >= M || ny < 0 || ny >= N {
			continue
		}

		if visited[ny][nx] {
			continue
		}

		if farm[ny][nx] == 0 {
			continue
		}

		visited[ny][nx] = true
		dfs(nx, ny)
	}
}

// reset visited
func reset() {
	for i := 0; i < 51; i++ {
		for j := 0; j < 51; j++ {
			visited[i][j] = false
			farm[i][j] = 0
		}
	}
}
