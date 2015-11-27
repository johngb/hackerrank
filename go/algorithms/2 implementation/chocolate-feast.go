package main

import "fmt"

func calcNumChoc(money, priceChoc, numWrappersForChoc int) int {
	numChoc := money / priceChoc
	wrappers := numChoc
	for wrappers >= numWrappersForChoc {
		numNewChocs := wrappers / numWrappersForChoc
		numChoc += numNewChocs
		wrappers = wrappers%numWrappersForChoc + numNewChocs
	}

	return numChoc
}

func main() {

	var numTests, money, priceChoc, numWrappersForChoc int

	fmt.Scanln(&numTests)
	res := make([]int, numTests)

	for i := 0; i < numTests; i++ {
		fmt.Scanln(&money, &priceChoc, &numWrappersForChoc)

		res[i] = calcNumChoc(money, priceChoc, numWrappersForChoc)
	}

	for _, v := range res {
		fmt.Println(v)
	}
}
