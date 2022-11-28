package main

import (
	"fmt"
)

// dishInformations stores dish information
type dishInformations struct {
	D int // dish number
	S int // init stock
	P int // price
	N int // number of dish
}

type dishes struct {
	dishInformations map[int]*dishInformations
}

type kitchens struct {
	K                []bool
	dishInformations map[int]*dishInformations
	unCookingOrder   map[int]*dishInformations
}

// dishStocks stores the number of dishes in stock
type dishStocks struct {
	N int // number of dish
}

type tableInformations struct {
	N map[int]int // number of dish
}

type tables struct {
	tableInformations map[int]*tableInformations
}

func main() {
	var N int // 処理番号

	fmt.Scan(&N)
	dishes := dishes{}
	tables := tables{}
	kitchens := kitchens{}
	dishes.dishInformations = make(map[int]*dishInformations, 1000000)
	tables.tableInformations = make(map[int]*tableInformations, 100000)
	kitchens.dishInformations = make(map[int]*dishInformations, 1000000)

	switch N {
	case 1:
		dishes.initDish()
		tables.orderProcess(&dishes)
	case 2:
		kitchens.initKitchen()

	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")

	}
}

func (dishes *dishes) initDish() {
	var M int //メニュー数
	fmt.Scan(&M)
	for i := 0; i < M; i++ {
		var di dishInformations
		fmt.Scan(&di.D, &di.S, &di.P)
		di.N = di.S
		dishes.dishInformations[di.D] = &di
	}
}

// orderProcess taking a customer's order
// Step 1
func (tables *tables) orderProcess(dish *dishes) {
	var order string
	var T, D, N int
	for {
		fmt.Scan(&order)
		switch order {
		case "order":
			fmt.Scan(&T, &D, &N)
			if dish.dishInformations[D].N < N {
				fmt.Println("sould out", T)
			} else {
				for i := 0; i < N; i++ {
					dish.dishInformations[D].N--
					fmt.Println("received order", T, D)
				}
			}
		default:
			return

		}
	}

}

func (kitchens *kitchens) initKitchen() {
	var M, K int
	fmt.Scan(&M, &K)
	kitchens.K = make([]bool, K)
	for i := 0; i < M; i++ {
		var di dishInformations
		fmt.Scan(&di.D, &di.S, &di.P)
		di.N = di.S
		kitchens.dishInformations[di.D] = &di
	}
}

// cookingProcess is the process of increasing the food stock
// Step 2
func (kitchens *kitchens) cookingProcess(T, D, N int) {
	var received, order string
	var R int // order acceptance information
	var C int // completion information

	for {
		fmt.Scan(&received, &order, &R, &C)
		switch received {
		case "received":
			switch order {
			case "order":
				if kitchens.dishInformations[D].N < N {
					fmt.Println("sould out", T)
				} else {
					for i := 0; i < N; i++ {
						kitchens.dishInformations[D].N--
						fmt.Println("received order", T, D)
					}
				}
			default:
				return

			}
		default:
			return

		}
	}

}
