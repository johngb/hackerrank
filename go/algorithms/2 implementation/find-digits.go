package main

import (
	"fmt"
	"strconv"
)

func main() {
	var numTests, num int
	fmt.Scanln(&numTests)

	results := make([]int, numTests)

	for i := 0; i < numTests; i++ {
		fmt.Scanln(&num)
		results[i] = findDigits(num)
	}

	for _, v := range results {
		fmt.Println(v)
	}
}

func findDigits(n int) int {
	nStr := strconv.Itoa(n)
	var res int

	lenN := len(nStr)
	for i := 0; i < lenN; i++ {
		digit, _ := strconv.Atoi(string(nStr[i]))
		if digit == 0 {
			continue
		}
		if n%digit == 0 {
			res++
		}
	}
	return res
}
