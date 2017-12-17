package solutions

import (
	"testing"

	. "github.com/franela/goblin"
)

// TestSimpleArraySum is a component wrapper over the Simple Array Sum solution
func TestAVeryBigSum(t *testing.T) {
	g := Goblin(t)
	g.Describe("A Very Big Sum", func() {

		g.It("Input: 1000000001 1000000002 1000000003 1000000004 1000000005 => 5000000015", func() {
			input := []int64{
				1000000001, 1000000002, 1000000003, 1000000004, 1000000005,
			}

			expected := 5000000015

			actual := AVeryBigSum(input)
			g.Assert(actual).Equal(expected)
		})

	})

}
