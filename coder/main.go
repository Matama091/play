package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func read() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	var N, K int
	fmt.Scan(&N, &K)

	max := 0
	A := make([]int, N)
	B := make([]int, K)
	for i := 0; i < N; i++ {
		A[i], _ = strconv.Atoi(read())

		if max < A[i] {
			max = A[i]
		}

	}

	for i := 0; i < K; i++ {
		B[i], _ = strconv.Atoi(read())
	}

	count := 0
	for i := 0; i < N; i++ {
		if max == A[i] {
			count++
		}
	}

	number := make([]int, count)
	c := 0
	for i := 0; i < N; i++ {
		if A[i] == max {
			number[c] = i + 1
			c++
		}
	}

	var yes bool
	for i := 0; i < count; i++ {
		for k := 0; k < K; k++ {
			if number[i] == B[k] {
				yes = true
				break
			}
		}
	}

	if yes {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
