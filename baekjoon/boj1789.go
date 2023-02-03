package boj1789

import (
	"bufio"
	"fmt"
	"os"
)

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
	N int
)

func main() {
	defer w.Flush()

	fmt.Fscan(r, &N)

	var sum int64 = 0
	var num int64 = 1
	var result int64 = 0

	for {
		sum += num
		result += 1
		if sum > int64(N) {
			result -= 1
			break
		}
		num += 1
	}

	fmt.Fprintln(w, result)
}
