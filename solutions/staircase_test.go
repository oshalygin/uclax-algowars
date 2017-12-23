package solutions_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/oshalygin/uclax-algowars/solutions"
)

var _ = Describe("Staircase", func() {
	It("6", func() {
		height := 6
		expected := []string{"#", "##", "###", "####", "#####", "######"}

		actual := StairCase(height)
		Expect(actual).To(Equal(expected))

	})

	It("1", func() {
		height := 1
		expected := []string{"#"}

		actual := StairCase(height)
		Expect(actual).To(Equal(expected))

	})

	It("2", func() {
		height := 2
		expected := []string{"#", "##"}

		actual := StairCase(height)
		Expect(actual).To(Equal(expected))

	})
})
