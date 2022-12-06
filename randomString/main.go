package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	var N int // 処理番号

	fmt.Scan(&N)

	answer := Solution(N)
	fmt.Println(answer)

}

func Solution(number int) (ans string) {
	rand.Seed(time.Now().UnixNano())

	alphabet := "abcdefghijklmnopqrstuvwxyz"
	check := make(map[int]int, 26)
	var length []string
	for i := 0; i < number; i++ {
		n := rand.Intn(25)
		count := rand.Intn(number/2-1) + 1
		if check[n]%2 == 0 {
			if count%2 != 0 {
				count++
				if i+count < number {
					continue
				}
				for j := 0; j < count; j++ {
					length = append(length, alphabet[n:n+1])
					i++
					check[n]++
				}
			}
		} else {
			if count%2 == 0 {
				count++
				if i+count < number {
					continue
				}
				for j := 0; j < count; j++ {
					length = append(length, alphabet[n:n+1])
					i++
					check[n]++
				}
			}

		}
		length = append(length, alphabet[n:n+1])
		check[n]++
	}

	rand.Shuffle(len(length), func(i, j int) { length[i], length[j] = length[j], length[i] })
	ans = strings.Join(length, "")

	return ans

}
