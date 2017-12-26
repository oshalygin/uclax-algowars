package solutions_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/oshalygin/uclax-algowars/solutions"
)

var _ = Describe("BirthdayCake", func() {
	It("3 2 1 3 => 2", func() {
		input := []int{3, 2, 1, 3}
		expected := 2

		actual := BirthdayCake(input)
		Expect(actual).To(Equal(expected))
	})

	It("1 1 1 1 => 4", func() {
		input := []int{1, 1, 1, 1}
		expected := 4

		actual := BirthdayCake(input)
		Expect(actual).To(Equal(expected))
	})

	It("1 1 1 1 7 4 4 => 7", func() {
		input := []int{1, 1, 1, 1, 7, 4, 4}
		expected := 1

		actual := BirthdayCake(input)
		Expect(actual).To(Equal(expected))
	})
})
