/*

Given a time in -hour AM/PM format, convert it to military (-hour) time.

Note: Midnight is  on a -hour clock, and  on a -hour clock. Noon is  on a -hour clock, and  on a -hour clock.

Input Format

A single string containing a time in -hour clock format (i.e.:  or ), where  and .

Output Format

Convert and print the given time in -hour format, where .

Sample Input

07:05:45PM
Sample Output

19:05:45

*/

package solutions

import (
	"strconv"
	"strings"
)

func TimeConversion(date string) string {
	isPM := strings.Contains(date, "PM")

	replacer := strings.NewReplacer("PM", "", "AM", "")
	updated := replacer.Replace(date)

	if !isPM {
		return updated
	}

	splitTime := strings.Split(updated, ":")
	hours, _ := strconv.Atoi(splitTime[0])
	hours = hours + 12

	splitTime[0] = strconv.Itoa(hours)

	updated = strings.Join(splitTime, ":")
	return updated
}
