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

func maxVehicleSize(segs []int, entry, exit int) int {
	if entry == exit {
		return segs[entry]
	}

	// fmt.Printf("segs: %v, entry: %d, exit: %d", segs, entry, exit)
	min := 3
	for i := entry; i <= exit; i++ {
		if segs[i] < min {
			min = segs[i]
		}
	}

	return min
}

func main() {
	var numTests int

	in := bufio.NewReader(os.Stdin)
	row, _ := getNextRow(in)
	numTests = row[1]

	freewaySegs, _ := getNextRow(in)

	var entry, exit int
	res := make([]int, numTests)
	for i := 0; i < numTests; i++ {
		row, _ := getNextRow(in)
		entry, exit = row[0], row[1]

		maxsize := maxVehicleSize(freewaySegs, entry, exit)
		res[i] = maxsize
	}

	for _, v := range res {
		fmt.Println(v)
	}
}
