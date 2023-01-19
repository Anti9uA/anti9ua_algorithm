package boj10026

import (
	"bufio"
	"fmt"
	"os"
)

var (
	r       = bufio.NewReader(os.Stdin)
	w       = bufio.NewWriter(os.Stdout)
	N       int
	photo   [][]string
	visited [][]bool
	// 상 하 좌 우
	dx = []int{-1, 1, 0, 0}
	dy = []int{0, 0, -1, 1}
)

func main() {
	defer w.Flush()

	fmt.Fscanln(r, &N)

	photo = make([][]string, N)
	visited = make([][]bool, N)
	for i := 0; i < N; i++ {
		photo[i] = make([]string, N)
		visited[i] = make([]bool, N)
		var temp string
		fmt.Fscanln(r, &temp)

		for j := 0; j < N; j++ {
			photo[i][j] = string(temp[j])
		}
	}

	normal := 0
	weakeness := 0
	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			if !visited[y][x] {
				visited[y][x] = true
				dfsNormal(x, y)
				normal++
			}
		}
	}
	reset()
	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			if !visited[y][x] && photo[y][x] != "B" {
				visited[y][x] = true
				dfsWeakeness(x, y, true)
				weakeness++
			}
			if !visited[y][x] && photo[y][x] == "B" {
				visited[y][x] = true
				dfsWeakeness(x, y, false)
				weakeness++
			}
		}
	}

	fmt.Fprintln(w, normal, weakeness)
}

// DFS for red-green weakness
// connected R and G without B
func dfsWeakeness(x, y int, flag bool) {
	for i := 0; i < 4; i++ {
		nx := x + dx[i]
		ny := y + dy[i]

		if nx < 0 || nx >= N || ny < 0 || ny >= N {
			continue
		}

		if visited[ny][nx] {
			continue
		}

		if flag {
			if photo[ny][nx] == "B" {
				continue
			}
		} else {
			if photo[ny][nx] != "B" {
				continue
			}
		}

		visited[ny][nx] = true
		dfsWeakeness(nx, ny, flag)
	}
}

// DFS for normal
// conneced only sample color
func dfsNormal(x, y int) {
	for i := 0; i < 4; i++ {
		nx := x + dx[i]
		ny := y + dy[i]

		if nx < 0 || nx >= N || ny < 0 || ny >= N {
			continue
		}

		if visited[ny][nx] {
			continue
		}

		if photo[y][x] != photo[ny][nx] {
			continue
		}

		visited[ny][nx] = true
		dfsNormal(nx, ny)
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
