package boj10815

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	r         = bufio.NewReader(os.Stdin)
	w         = bufio.NewWriter(os.Stdout)
	N         int
	M         int
	myCard    []int
	totalCard []int
	result    []int
)

func main() {
	defer w.Flush()

	fmt.Fscan(r, &N)

	myCard = make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(r, &myCard[i])
	}

	fmt.Fscan(r, &M)
	totalCard = make([]int, M)
	for i := 0; i < M; i++ {
		fmt.Fscan(r, &totalCard[i])
	}

	result = make([]int, M)

	sort.Ints(myCard)

	for i := 0; i < M; i++ {
		start := 0
		end := N - 1
		for start <= end {
			mid := (start + end) / 2
			if myCard[mid] == totalCard[i] {
				result[i] = 1
				break
			} else if myCard[mid] > totalCard[i] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}

	// print result with spaces
	for i := 0; i < M; i++ {
		fmt.Fprint(w, result[i], " ")
	}

}
