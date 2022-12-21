package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	str := sc.Text()
	fmt.Println(IsPalindrome(str))
}

func IsPalindrome(s string) bool {
	str := []rune(s)
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-1-i] {
			return false
		}
	}
	return true
}
