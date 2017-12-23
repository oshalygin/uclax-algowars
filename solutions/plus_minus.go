/*
Given an array of integers, calculate which fraction of its elements are positive, which fraction of its elements are negative, and which fraction of its elements are zeroes, respectively. Print the decimal value of each fraction on a new line.

Note: This challenge introduces precision problems. The test cases are scaled to six decimal places, though answers with absolute error of up to  are acceptable.

Input Format

The first line contains an integer, , denoting the size of the array.
The second line contains  space-separated integers describing an array of numbers .

Output Format

You must print the following  lines:

A decimal representing of the fraction of positive numbers in the array compared to its size.
A decimal representing of the fraction of negative numbers in the array compared to its size.
A decimal representing of the fraction of zeroes in the array compared to its size.
*/

package solutions

// PlusMinusFractions defines the data object that stores the positive, negative and zero values of an array
type PlusMinusFractions struct {
	Positive float64
	Negative float64
	Zero     float64
}

// PlusMinus is a solution to the Hackerrank problem
func PlusMinus(input []int) PlusMinusFractions {

	result := PlusMinusFractions{}

	n := len(input)

	positives := 0.00
	negatives := 0.00
	zeros := 0.00

	for i := range input {
		if input[i] == 0 {
			zeros++
		}
		if input[i] > 0 {
			positives++
		}
		if input[i] < 0 {
			negatives++
		}
	}

	result.Positive = positives / float64(n)
	result.Negative = negatives / float64(n)
	result.Zero = zeros / float64(n)

	return result
}
