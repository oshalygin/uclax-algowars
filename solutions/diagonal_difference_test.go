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

		g.It("Input: [[11, 2, 4, 3], [4, 5, 6, 3], [10, 8, -12, 3], [1, 1, 1, 1]] => 13", func() {

			input := [][]int{
				[]int{11, 2, 4, 3}, []int{4, 5, 6, 3}, []int{10, 8, -12, 3}, []int{1, 1, 1, 1},
			}

			expected := 13

			actual := DiagonalDifference(input)
			g.Assert(actual).Equal(expected)
		})

	})

}
