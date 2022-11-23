package main

import (
	"bufio"
	"fmt"
	"os"
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

	B := make([][]int, K)
	for i := 0; i < K; i++ {
		B[i] = make([]int, i/K)
	}

	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
		B[i%K] = append(B[i%K], A[i])
	}

	if K == 1 {
		fmt.Println("Yes")
		return
	}

	for i := 0; i < K; i++ {
		Quicksort(B[i])
	}

	var C []int
	for i := 0; i < N; i++ {
		C = append(C, B[i%K][i/K])
	}

	if Comparison(C) {
		fmt.Println("Yes")
		return
	}

	fmt.Println("No")

}

func Comparison(A []int) bool {
	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			return false
		}
	}
	return true
}

func Quicksort(s []int) []int {
	if len(s) == 1 || len(s) == 0 { //list size 0 or 1 is no sort
		return s
	} else {
		pivot := s[0] //make a pivot first in the list
		place := 0

		for j := 0; j < len(s)-1; j++ {
			if s[j+1] < pivot { // if it is smaller than the pivot
				s[j+1], s[place+1] = s[place+1], s[j+1]
				place++
			}
		}
		s[0], s[place] = s[place], s[0]

		first := Quicksort(s[:place])
		second := Quicksort(s[place+1:])
		first = append(first, s[place])

		first = append(first, second...)
		return first
	}

}
