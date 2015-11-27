package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getNextRow(r *bufio.Reader) ([]int, error) {
	line, err := r.ReadString('\n')
	if err != nil && err.Error() != "EOF" {
		return []int{}, err
	}

	sArr := strings.Fields(line)
	res := make([]int, len(sArr))
	for k, v := range sArr {
		n, err := strconv.Atoi(string(v))
		if err != nil {
			return []int{}, err
		}
		res[k] = n
	}
	return res, nil
}

func main() {
	var size float64
	fmt.Scanln(&size)

	in := bufio.NewReader(os.Stdin)
	nums, err := getNextRow(in)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}

	var numPos, numNeg, numZero int
	for _, v := range nums {
		switch {
		case v < 0:
			numNeg++
		case v == 0:
			numZero++
		case v > 0:
			numPos++
		}
	}

	fmt.Printf("%.3f\n", float64(numPos)/size)
	fmt.Printf("%.3f\n", float64(numNeg)/size)
	fmt.Printf("%.3f", float64(numZero)/size)
}
