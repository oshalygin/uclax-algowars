package solutions_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/oshalygin/uclax-algowars/solutions"
)

var _ = Describe("GradingStudents", func() {
	It("73 67 38 33 => 75 67 40 33", func() {
		grades := []int{73, 67, 38, 33}
		expected := []int{75, 67, 40, 33}

		actual := GradingStudents(grades)
		Expect(actual).To(Equal(expected))
	})
})
