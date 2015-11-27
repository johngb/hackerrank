package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var size int
	fmt.Scanln(&size)

	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	if err != nil && err.Error() != "EOF" {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	sArr := strings.Fields(line)
	var res int64
	for _, v := range sArr {
		n, err := strconv.ParseInt(string(v), 10, 64)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(2)
		}
		res += n
	}

	fmt.Println(res)
}
