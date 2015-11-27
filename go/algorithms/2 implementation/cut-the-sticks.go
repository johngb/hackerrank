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

func cutSticks(sticks []int) []int {
	min := getMin(sticks)

	var res []int
	for _, v := range sticks {
		newLen := v - min
		if newLen != 0 {
			res = append(res, newLen)
		}
	}

	return res
}

func getMin(arr []int) int {
	min := 10000
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min
}

func main() {

	in := bufio.NewReader(os.Stdin)
	_, _ = getNextRow(in)

	sticks, _ := getNextRow(in)
	for len(sticks) != 0 {
		fmt.Println(len(sticks))
		sticks = cutSticks(sticks)
	}

}
