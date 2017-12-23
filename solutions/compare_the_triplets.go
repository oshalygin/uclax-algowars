package solutions

// CompareTheTriplets solution
func CompareTheTriplets(aliceInput []int, bobInput []int) []int {
	alice := 0
	bob := 0

	for i := range aliceInput {
		if aliceInput[i] > bobInput[i] {
			alice++
		} else if aliceInput[i] < bobInput[i] {
			bob++
		}
	}

	return []int{alice, bob}
}
