/*

Alice and Bob each created one problem for HackerRank. A reviewer rates the two challenges, awarding points on a scale from  to  for three categories: problem clarity, originality, and difficulty.

We define the rating for Alice's challenge to be the triplet , and the rating for Bob's challenge to be the triplet .

Your task is to find their comparison points by comparing  with ,  with , and  with .

If , then Alice is awarded  point.
If , then Bob is awarded  point.
If , then neither person receives a point.
Comparison points is the total points a person earned.

Given  and , can you compare the two challenges and print their respective comparison points?

Input Format

The first line contains  space-separated integers, , , and , describing the respective values in triplet .
The second line contains  space-separated integers, , , and , describing the respective values in triplet .

Constraints

Output Format

Print two space-separated integers denoting the respective comparison points earned by Alice and Bob.

Sample Input

5 6 7
3 6 10
Sample Output

1 1
Explanation

In this example:

Now, let's compare each individual score:

, so Alice receives  point.
, so nobody receives a point.
, so Bob receives  point.
Alice's comparison score is , and Bob's comparison score is . Thus, we print 1 1 (Alice's comparison score followed by Bob's comparison score) on a single line.

*/

package solutions

import (
	"testing"

	. "github.com/franela/goblin"
)

// TestSimpleArraySum is a component wrapper over the Simple Array Sum solution
func TestCompareTheTriplets(t *testing.T) {
	g := Goblin(t)
	g.Describe("Compare The Triplets", func() {

		g.It("Input: [5 6 7] [3 6 10] => [1 1]", func() {
			aliceInput := []int{5, 6, 7}
			bobInput := []int{3, 6, 10}
			expected := []int{1, 1}

			actual := CompareTheTriplets(aliceInput, bobInput)
			g.Assert(actual).Equal(expected)
		})

		g.It("Input: [1 1 1] [1 1 1] => [0 0]", func() {
			aliceInput := []int{1, 1, 1}
			bobInput := []int{1, 1, 1}
			expected := []int{0, 0}

			actual := CompareTheTriplets(aliceInput, bobInput)
			g.Assert(actual).Equal(expected)
		})

		g.It("Input: [2 2 2] [1 1 1] => [3 0]", func() {
			aliceInput := []int{2, 2, 2}
			bobInput := []int{1, 1, 1}
			expected := []int{3, 0}

			actual := CompareTheTriplets(aliceInput, bobInput)
			g.Assert(actual).Equal(expected)
		})

	})

}
