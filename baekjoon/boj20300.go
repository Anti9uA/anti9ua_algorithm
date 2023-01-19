package boj20300

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

// 그리디로 풀었음
// 가장 작은 수와 가장 큰 수를 더한 값이 최소 근손실이 되는 것을 이용
// 1. 배열을 정렬한다.
// 2. 배열의 길이가 홀수라면, 가장 큰 수를 temp에 넣는다.
// 3. temp에는 배열의 가장 작은 수와 가장 큰 수를 더한 값이 들어간다.
// 4. temp의 가장 큰 수를 출력한다.
