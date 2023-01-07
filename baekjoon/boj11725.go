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
	tree    [][]int
	visited []bool
)

func main() {
	defer w.Flush()

	var N int
	fmt.Fscanln(r, &N)

	tree = make([][]int, N+1)
	visited = make([]bool, N+1)

	for i := 0; i < N-1; i++ {
		var u, v int
		fmt.Fscanln(r, &u, &v)
		tree[u] = append(tree[u], v)
		tree[v] = append(tree[v], u)
	}

	parents := bfs()
	for i := 2; i <= N; i++ {
		fmt.Fprintln(w, parents[i])
	}

}

func bfs() (parents []int) {
	visited[1] = true
	queue := []int{1}
	parents = make([]int, len(tree))

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, child := range tree[node] {
			if !visited[child] {
				visited[child] = true
				parents[child] = node
				queue = append(queue, child)
			}
		}
	}
	return
}
