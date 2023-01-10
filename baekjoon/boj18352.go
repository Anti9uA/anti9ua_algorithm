package boj18352

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var n, m, k, x int
	fmt.Fscan(r, &n, &m, &k, &x)

	adjList := make([][]int, n+1)

	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(r, &a, &b)
		adjList[a] = append(adjList[a], b)
	}

	var q []int
	distance := make([]int, n+1)
	for i := range distance {
		distance[i] = -1
	}

	q = append(q, x)
	distance[x] = 0
	var ans []int
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		if distance[cur] >= k {
			break
		}

		for _, next := range adjList[cur] {
			if distance[next] == -1 {
				q = append(q, next)
				distance[next] = distance[cur] + 1
				if distance[next] == k {
					ans = append(ans, next)
				}
			}
		}
	}

	sort.Slice(ans, func(i, j int) bool {
		return ans[i] < ans[j]
	})

	if len(ans) == 0 {
		fmt.Fprintln(w, -1)
	} else {
		for _, x := range ans {
			fmt.Fprintln(w, x)
		}
	}

	w.Flush()
}
