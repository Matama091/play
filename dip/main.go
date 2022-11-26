package main

import (
	"fmt"
)

func main() {
	var N int // 人数
	fmt.Scan(&N)

	Number := make(map[int]bool, N)

	A := make([]int, N) // 数字
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])

	}

	for i := 0; i < N; i++ {
		if Number[A[i]] {
			fmt.Println(i + 1)
			return
		}

		Number[A[i]] = true
	}

	fmt.Println("-1")
}
