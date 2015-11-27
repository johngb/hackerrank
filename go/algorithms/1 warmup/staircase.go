package main

import "fmt"

func spaces(n int) string {
	var r string
	for i := 0; i < n; i++ {
		r += " "
	}
	return r
}

func hashes(n int) string {
	var r string
	for i := 0; i < n; i++ {
		r += "#"
	}
	return r
}

func main() {
	var size int
	fmt.Scanln(&size)

	for i := 0; i < size; i++ {
		line := spaces(size-i-1) + hashes(i+1)
		fmt.Println(line)
	}
}
