package solutions_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/oshalygin/uclax-algowars/solutions"
)

var _ = Describe("SimpleArraySum", func() {

	It("Input: 1 2 3 4 10 11 => 31", func() {
		userInput := []int{1, 2, 3, 4, 10, 11}
		expected := 31

		actual := GetArraySum(userInput)

		Expect(actual).To(Equal(expected))
	})

	It("Input: 54 5 3 2 780 85 => 31", func() {
		userInput := []int{54, 5, 3, 2, 780, 85}
		expected := 929

		actual := GetArraySum(userInput)
		Expect(actual).To(Equal(expected))
	})

})
