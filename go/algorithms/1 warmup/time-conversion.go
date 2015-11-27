package main

import (
	"fmt"
	"strconv"
)

func parseTime(ts string) string {
	hour, _ := strconv.Atoi(ts[0:2])

	var newTime string
	var isPM bool

	// If it's a PM time
	if ts[len(ts)-2:len(ts)] == "PM" {
		isPM = true
	}

	switch {

	// If it's AM and the time shows as 12
	case hour == 12 && !isPM:
		newTime = new24Hour(0, ts)

	case hour == 12 && isPM:
		newTime = new24Hour(12, ts)

	case !isPM:
		newTime = new24Hour(hour, ts)

	case isPM:
		newTime = new24Hour(hour+12, ts)
	}
	return newTime
}

func new24Hour(hour int, ts string) string {
	var hs string
	if hour < 10 {
		hs = "0" + strconv.Itoa(hour)
	} else {
		hs = strconv.Itoa(hour)
	}
	return hs + ts[2:len(ts)-2]
}

func main() {
	var ts string
	fmt.Scanln(&ts)
	fmt.Println(parseTime(ts))
}
