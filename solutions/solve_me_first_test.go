/*
Welcome to HackerRank! The purpose of this challenge is to familiarize you with reading input from stdin (the standard input stream) and writing output to stdout (the standard output stream) using our environment.

Review the code provided in the editor below, then complete the solveMeFirst function so that it returns the sum of two integers read from stdin. Take some time to understand this code so you're prepared to write it yourself in future challenges.

Select a language below, and start coding!

Input Format

Code that reads input from stdin is provided for you in the editor. There are  lines of input, and each line contains a single integer.

Output Format

Code that prints the sum calculated and returned by solveMeFirst is provided for you in the editor.
*/

package tests

import (
	"testing"

	. "github.com/franela/goblin"
)

func SolveMeFirst(firstValue int, secondValue int) int {
	return firstValue + secondValue
}

func Test_SolveMeFirst(t *testing.T) {
	g := Goblin(t)
	g.Describe("Solve Me First", func() {
		g.It("2 + 3 = 5", func() {
			a := 2
			b := 3

			expected := 4

			actual := SolveMeFirst(a, b)

			g.Assert(actual).Equal(expected)

		})
	})
}