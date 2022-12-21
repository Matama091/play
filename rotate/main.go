package main

import (
	"errors"
	"fmt"
)

func main() {
	// sc := bufio.NewScanner(os.Stdin)
	// sc.Scan()
	// if err := sc.Err(); err != nil {
	// 	log.Fatal(err)
	// }
	// str := sc.Text()
	// number, err := strconv.Atoi(str)
	// if err != nil {
	// 	log.Fatal()
	// }
	matrix := [][]int{{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(matrix)

	fmt.Println(Rotate(matrix))
}

func Rotate(matrix [][]int) ([][]int, error) {
	n := len(matrix)

	for i := 0; i < n; i++ {
		if n != len(matrix[i]) {
			err := errors.New("not square array")
			return nil, err
		}
	}

	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}

	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			ans[n-1-col][row] = matrix[row][col]
		}
	}

	return ans, nil
}
