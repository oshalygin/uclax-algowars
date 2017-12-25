package solutions_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/oshalygin/uclax-algowars/solutions"
)

var _ = Describe("Camelcase", func() {
	It("camelCase => 2", func() {
		input := "camelCase"
		expected := 2

		actual := CamelCase(input)
		Expect(actual).To(Equal(expected))
	})

	It("helloMyNameIsOlegAndIAmWritingASentenceHere => 12", func() {
		input := "helloMyNameIsOlegAndIAmWritingASentenceHere"
		expected := 12

		actual := CamelCase(input)
		Expect(actual).To(Equal(expected))
	})
})
