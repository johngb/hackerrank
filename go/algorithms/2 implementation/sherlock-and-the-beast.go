package main

import "fmt"

func getDecentNumber(n int) string {
	if n <= 0 {
		return "-1"
	}

	for fives := n; fives >= 0; fives-- {
		if fives%3 == 0 {
			threes := n - fives
			if checkValidThrees(threes, fives) {
				return makeDecentNumber(threes, fives)
			}
		}
	}
	return "-1"
}

func checkValidThrees(threes, fives int) bool {
	if threes%5 == 0 {
		return true
	}
	return false
}

func makeDecentNumber(threes, fives int) string {
	res := []byte{}
	for i := 0; i < fives; i++ {
		res = append(res, '5')
	}
	for i := 0; i < threes; i++ {
		res = append(res, '3')
	}
	return string(res)
}

func main() {

	var numTests int
	fmt.Scanln(&numTests)

	var digits int
	for i := 0; i < numTests; i++ {
		fmt.Scanln(&digits)
		fmt.Println(getDecentNumber(digits))
	}
}
