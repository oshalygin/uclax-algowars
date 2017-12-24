/*

Given five positive integers, find the minimum and maximum values that can be calculated by summing exactly four of the five integers. Then print the respective minimum and maximum values as a single line of two space-separated long integers.

Input Format

A single line of five space-separated integers.

Constraints

Each integer is in the inclusive range .
Output Format

Print two space-separated long integers denoting the respective minimum and maximum values that can be calculated by summing exactly four of the five integers. (The output can be greater than 32 bit integer.)

Sample Input

1 2 3 4 5
Sample Output

10 14
Explanation

Our initial numbers are , , , , and . We can calculate the following sums using four of the five integers:

If we sum everything except , our sum is .
If we sum everything except , our sum is .
If we sum everything except , our sum is .
If we sum everything except , our sum is .
If we sum everything except , our sum is .
As you can see, the minimal sum is  and the maximal sum is . Thus, we print these minimal and maximal sums as two space-separated integers on a new line.

Hints: Beware of integer overflow! Use 64-bit Integer.

*/

package solutions

type MiniMax struct {
	Minimum int64
	Maximum int64
}

func MiniMaxSum(input []int64) MiniMax {
	miniMax := MiniMax{Minimum: 0, Maximum: 0}

	var minimumInt int64 = 1
	var maximumInt int64 = 1
	var minimumIntIndex int
	var maximumIntIndex int

	for i := 0; i < len(input); i++ {
		if input[i] > maximumInt {
			maximumInt = input[i]
			maximumIntIndex = i
		}

		if input[i] <= minimumInt {
			minimumInt = input[i]
			minimumIntIndex = i
		}
	}

	for i, number := range input {
		if minimumIntIndex != i {
			miniMax.Maximum += number
		}
		if maximumIntIndex != i {
			miniMax.Minimum += number
		}
	}

	return miniMax
}
