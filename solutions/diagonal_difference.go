/*
Given a square matrix of size , calculate the absolute difference between the sums of its diagonals.

Input Format

The first line contains a single integer, . The next  lines denote the matrix's rows, with each line containing space-separated integers describing the columns.

Constraints

Output Format

Print the absolute difference between the two sums of the matrix's diagonals as a single integer.

Sample Input

3
11 2 4
4 5 6
10 8 -12
Sample Output

15
Explanation

The primary diagonal is:

11
   5
     -12
Sum across the primary diagonal: 11 + 5 - 12 = 4

The secondary diagonal is:

     4
   5
10
Sum across the secondary diagonal: 4 + 5 + 10 = 19
Difference: |4 - 19| = 15

Note: |x| is absolute value function
*/

package solutions

func abs(number int) int {
	if number < 0 {
		return number * -1
	}
	return number

}

// DiagonalDifference calculates the difference between the diagonals in the matrix
func DiagonalDifference(input [][]int) int {
	var leftDiagonal int
	var rightDiagonal int
	arrSize := len(input) - 1

	for i := range input {
		leftDiagonal = leftDiagonal + input[i][i]
		rightDiagonal = rightDiagonal + input[i][arrSize-i]
	}

	sum := leftDiagonal - rightDiagonal

	absoluteValue := abs(sum)
	return absoluteValue
}
