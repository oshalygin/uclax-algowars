package solutions

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestDiagonalDifference(t *testing.T) {
	g := Goblin(t)
	g.Describe("Diagonal Difference", func() {

		g.It("Input: [[11, 2, 4], [4, 5, 6], [10, 8, -12]] => 15", func() {

			input := [][]int{
				[]int{11, 2, 4}, []int{4, 5, 6}, []int{10, 8, -12},
			}

			expected := 15

			actual := DiagonalDifference(input)
			g.Assert(actual).Equal(expected)
		})

	})

}
