package main

import (
	"fmt"
)

func main() {
outer:
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j)
			if i == j {
				break outer
			}
		}

	}

	//Label
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("<%d-%d>", i, j)
			if i == j {
				goto aaa
			}
		}
	}

aaa:
	fmt.Println("Exit")
}

//https://go.dev/play/p/s1MvdsRyShu
