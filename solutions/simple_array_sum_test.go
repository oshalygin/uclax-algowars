/*

Given an array of  integers, can you find the sum of its elements?

Input Format

The first line contains an integer, , denoting the size of the array.
The second line contains  space-separated integers representing the array's elements.

Output Format

Print the sum of the array's elements as a single integer.

Sample Input

6
1 2 3 4 10 11

Sample Output

31

*/

package tests

import (
	"strconv"
	"strings"
	"testing"

	. "github.com/franela/goblin"
)

func getArraySum(input string) int {
	array := strings.Split(input, " ")

	sum := 0
	for _, value := range array {
		integerValue, _ := strconv.Atoi(value)
		sum = sum + integerValue
	}

	return sum
}

// TestSimpleArraySum is a component wrapper over the Simple Array Sum solution
func TestSimpleArraySum(t *testing.T) {
	g := Goblin(t)
	g.Describe("Simple Array Sum", func() {

		g.It("Input: 1 2 3 4 10 11 => 31", func() {
			userInput := "1 2 3 4 10 11"
			expected := 31

			actual := getArraySum(userInput)
			g.Assert(actual).Equal(expected)
		})

		g.It("Input: 54 5 3 2 780 85 => 31", func() {
			userInput := "54 5 3 2 780 85"
			expected := 929

			actual := getArraySum(userInput)
			g.Assert(actual).Equal(expected)
		})
	})

}
