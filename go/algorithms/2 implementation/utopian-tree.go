package main

import (
	"fmt"
	"strconv"
)

func main() {
	var numTests, growthCycles int
	fmt.Scanln(&numTests)

	for i := 0; i < numTests; i++ {
		fmt.Scanln(&growthCycles)
		fmt.Println(strconv.Itoa(calcHeight(growthCycles)))
	}
}

func calcHeight(c int) int {
	height := 1
	for i := 1; i <= c; i++ {
		if i%2 == 1 {
			height *= 2
		} else {
			height += 1
		}
	}
	return height
}
