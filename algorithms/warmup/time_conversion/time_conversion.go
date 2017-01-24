package main

import "strings"
import "strconv"

// 12:00:00AM ===== 00:00:00
// 12:00:00PM ===== 12:00:00

func time_convertion(tm string) string {

	if tm[8:] == "AM" {
		if tm[0:2] == "12" {
			return strings.Join([]string{"00", tm[2:8]}, "")
		}
		return tm[:8]
	}
	if tm[0:2] == "12" {
		return tm[0:8]
	}
	hour, _ := strconv.Atoi(tm[0:2])
	hour += 12
	return strings.Join([]string{strconv.Itoa(hour), tm[2:8]}, "")
}
