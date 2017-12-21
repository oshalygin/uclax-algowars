package solutions

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestPlusMinus(t *testing.T) {
	g := Goblin(t)
	g.Describe("Plus Minus", func() {

		g.It("Input: [[-4, 3, -9, 0, 4, 1]] => 0.500000 0.333333 0.166667", func() {

			input := []int{-4, 3, -9, 0, 4, 1}

			expected := PlusMinusFractions{positive: 0.500000, negative: 0.3333333333333333, zero: 0.16666666666666666}

			actual := PlusMinus(input)
			g.Assert(actual).Equal(expected)
		})

		g.It("Input: [[4, 3, 9, 2, 4, 1]] => 1 0 0", func() {

			input := []int{4, 3, 9, 2, 4, 1}

			expected := PlusMinusFractions{positive: 1, negative: 0, zero: 0}

			actual := PlusMinus(input)
			g.Assert(actual).Equal(expected)
		})

	})

}
