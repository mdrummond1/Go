package main

import (
	"fmt"
	"strconv"
)

func add(x int, y int) int {
	var z int
	z = x + y

	return z
}

func main() {

	for i := 0; i <= 100; i += 2 {

		//fmt.Println("even!")
		fmt.Println("i: " + strconv.Itoa(i))

	}

}
