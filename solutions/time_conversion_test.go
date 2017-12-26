package solutions_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/oshalygin/uclax-algowars/solutions"
)

var _ = Describe("TimeConversion", func() {
	It("07:05:45PM => 19:05:45", func() {
		input := "07:05:45PM"
		expected := "19:05:45"

		actual := TimeConversion(input)

		Expect(actual).To(Equal(expected))
	})
})
