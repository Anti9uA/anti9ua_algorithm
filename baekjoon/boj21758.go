package boj21758

import (
	"bufio"
	"fmt"
	"os"
)

var (
	r   = bufio.NewReader(os.Stdin)
	w   = bufio.NewWriter(os.Stdout)
	N   int
	arr []int
	sum []int
)

func main() {
	defer w.Flush()

	fmt.Fscan(r, &N)

	arr = make([]int, N+1)
	for i := 1; i <= N; i++ {
		fmt.Fscan(r, &arr[i])
	}

	sum = make([]int, N+1)
	for i := 1; i <= N; i++ {
		sum[i] = sum[i-1] + arr[i]
	}

	var current int
	var result int

	for i := 2; i < N; i++ {
		current = (sum[N] - sum[1] - arr[i]) + (sum[N] - sum[i])
		result = getMax(result, current)
	}

	for i := 2; i < N; i++ {
		current = sum[N-1] - arr[i] + sum[i-1]
		result = getMax(result, current)
	}

	for i := 2; i < N; i++ {
		current = (sum[i] - sum[1]) + (sum[N-1] - sum[i-1])
		result = getMax(result, current)
	}

	fmt.Fprintln(w, result)

}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
