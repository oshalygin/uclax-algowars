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

	It("07:05:45AM => 07:05:45", func() {
		input := "07:05:45AM"
		expected := "07:05:45"

		actual := TimeConversion(input)

		Expect(actual).To(Equal(expected))
	})

	It("00:00:00AM => 00:00:00", func() {
		input := "00:00:00AM"
		expected := "00:00:00"

		actual := TimeConversion(input)

		Expect(actual).To(Equal(expected))
	})

	It("12:00:00PM => 12:00:00", func() {
		input := "12:00:00PM"
		expected := "12:00:00"

		actual := TimeConversion(input)

		Expect(actual).To(Equal(expected))
	})

	It("12:00:00AM => 12:00:00", func() {
		input := "12:00:00AM"
		expected := "00:00:00"

		actual := TimeConversion(input)

		Expect(actual).To(Equal(expected))
	})
})
