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
	K    int
	arr  []int
	temp []int
)

func main() {
	defer w.Flush()

	fmt.Fscan(r, &N, &K)
	for i := 0; i < N; i++ {
		var a int
		fmt.Fscan(r, &a)
		arr = append(arr, a)
	}

	ans := 0

	for i := 0; i < N-1; i++ {
		temp = append(temp, arr[i+1]-arr[i])
	}

	sort.Ints(temp)

	for i := 0; i < N-K; i++ {
		ans += temp[i]
	}

	fmt.Fprintln(w, ans)
}

// 그리디 사용했음
// 조별로 인원수가 같을 필요가 없다는게 핵심
// 옆에 있는 사람과의 차이를 구해서 정렬한 후
// 가장 작은 차이를 가진 사람들을 빼주면 된다.
