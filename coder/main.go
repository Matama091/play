package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	var S string
	fmt.Scan(&S)
	var a, A int

	if !check_regexp(`[A-Z]`, S) || !check_regexp(`[a-b]`, S) {
		fmt.Println("No")
		return
	}

	for _, c := range S {
		a = strings.Index(S, string(c))
		A = strings.LastIndex(S, string(c))
		fmt.Println(a, A, string(c))
		if !(a == A) {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")

}

func check_regexp(reg, str string) bool {
	return regexp.MustCompile(reg).Match([]byte(str))
}
