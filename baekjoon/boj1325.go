package boj1325

import (
	"bufio"
	"fmt"
	"os"
)

var (
	r       = bufio.NewReader(os.Stdin)
	w       = bufio.NewWriter(os.Stdout)
	N       int
	M       int
	graph   [][]int
	visited []bool
	answer  []int
	cnt     int
	max     int
)

func main() {
	defer w.Flush()

	var N int
	fmt.Fscanln(r, &N, &M)

	graph = make([][]int, N+1)

	for i := 0; i < M; i++ {
		var a, b int
		fmt.Fscanln(r, &a, &b)
		graph[b] = append(graph[b], a)
	}

	for i := 1; i <= N; i++ {
		visited = make([]bool, len(graph))
		cnt = 0

		dfs(i)
		if cnt > max {
			max = cnt
			answer = nil
			answer = append(answer, i)
		}
		if cnt == max {
			answer = append(answer, i)
		}
	}

	answer = makeSliceUnique(answer)

	for i := 0; i < len(answer); i++ {
		fmt.Fprintln(w, answer[i])
	}

}

func dfs(v int) {
	visited[v] = true

	for i := 0; i < len(graph[v]); i++ {
		next := graph[v][i]
		if !visited[next] {
			cnt++
			dfs(next)
		}
	}
}

// makeSliceUnique removes duplicate elements from a slice
func makeSliceUnique(slice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
