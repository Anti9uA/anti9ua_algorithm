package boj2110

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	r   = bufio.NewReader(os.Stdin)
	w   = bufio.NewWriter(os.Stdout)
	N   int
	C   int
	arr []int
)

func main() {
	defer w.Flush()

	fmt.Fscan(r, &N, &C)
	for i := 0; i < N; i++ {
		var a int
		fmt.Fscan(r, &a)
		arr = append(arr, a)
	}

	sort.Ints(arr)

	start := 1
	end := arr[N-1] - arr[0]

	answer := 0

	for start <= end {
		mid := (start + end) / 2
		value := arr[0]
		count := 1

		for i := 1; i < N; i++ {
			if arr[i] >= value+mid {
				value = arr[i]
				count++
			}
		}

		if count >= C {
			start = mid + 1
			answer = mid
		} else {
			end = mid - 1
		}
	}

	fmt.Fprintln(w, answer)
}
