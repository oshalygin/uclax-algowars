package solutions_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/oshalygin/uclax-algowars/solutions"
)

var _ = Describe("AVeryBigSum", func() {

	It("Input: 1000000001 1000000002 1000000003 1000000004 1000000005 => 5000000015", func() {
		input := []int64{
			1000000001, 1000000002, 1000000003, 1000000004, 1000000005,
		}

		var expected int64 = 5000000015

		actual := AVeryBigSum(input)
		Expect(actual).To(Equal(expected))
	})
})
