package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	r    = bufio.NewReader(os.Stdin)
	w    = bufio.NewWriter(os.Stdout)
	N    int
	arr  []int
	temp []int
)

func main() {
	defer w.Flush()
	fmt.Fscan(r, &N)
	for i := 0; i < N; i++ {
		var a int
		fmt.Fscan(r, &a)
		arr = append(arr, a)
	}

	sort.Ints(arr)

	if N%2 == 1 {
		temp = append(temp, arr[len(arr)-1])
		arr = arr[:len(arr)-1]
	}

	for i := 0; i < N/2; i++ {
		temp = append(temp, arr[i]+arr[len(arr)-1-i])
	}

	sort.Ints(temp)
	fmt.Fprint(w, temp[len(temp)-1])
}
