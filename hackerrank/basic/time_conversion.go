/*
https://www.hackerrank.com/challenges/one-week-preparation-kit-time-conversion/problem

Convert 12 hour time into 24 hour time.
	- 12h: 12:00:00 AM = 24: 00:00:00

Return a new string representing the input time in 24 hour format.
--------------------------
Break down the input string into hh, mm, ss and AM/PM

If AM - 24h hh space is 00
If PM - 24h hh space is 12 + value (except if its already 12 - then its 12)
mm:ss should carry over from 12h time
Format an output string with zero prefix if num is less than 10.

Sample:

07:05:45PM

*/

package basic

import (
	"fmt"
	"strconv"
	"strings"
)

func TimeConversion(s string) string {
	miltime := ""

	clone := strings.Clone(s)

	// Get rid of AM/PM
	clone = strings.TrimRight(clone, "APM")

	// Break apart the values
	parts := strings.Split(clone, ":")

	// Convert from string to int
	hh, _ := strconv.ParseInt(parts[0], 10, 0)
	mm, _ := strconv.ParseInt(parts[1], 10, 0)
	ss, _ := strconv.ParseInt(parts[2], 10, 0)

	// AM or PM
	if strings.Contains(s, "AM") {
		// 12h on 24h is 00
		if hh == 12 {
			hh = 0
		}
		miltime = fmt.Sprintf("%02d:%02d:%02d", hh, mm, ss)
		fmt.Println(miltime)
	} else {
		// Add 12h to get 24h time
		if hh != 12 {
			hh += 12
		}
		miltime = fmt.Sprintf("%02d:%02d:%02d", hh, mm, ss)
		fmt.Println(miltime)
	}

	return miltime
}
