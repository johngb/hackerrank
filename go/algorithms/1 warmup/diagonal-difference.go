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

func abs(n int) int {
	if n <= 0 {
		return -n
	}
	return n
}

func main() {
	var size int
	fmt.Scanln(&size)

	in := bufio.NewReader(os.Stdin)
	var sum int
	for rn := 0; rn < size; rn++ {
		row, err := getNextRow(in)
		if err != nil {
			fmt.Println("Error: ", err.Error())
			os.Exit(1)
		}
		sum += row[rn] - row[size-rn-1]
	}

	fmt.Println(abs(sum))
}
