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

func classCancelled(r *bufio.Reader) bool {
	inParams, _ := getNextRow(r)
	studentsNeeded := inParams[1]
	times, _ := getNextRow(r)

	var onTime int
	for _, v := range times {
		if v <= 0 {
			onTime++
		}
	}

	if onTime < studentsNeeded {
		return true
	}
	return false
}

func main() {
	var numTests int
	fmt.Scanln(&numTests)

	in := bufio.NewReader(os.Stdin)

	for i := 0; i < numTests; i++ {
		if classCancelled(in) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
