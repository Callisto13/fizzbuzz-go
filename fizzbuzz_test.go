package fizzbuzz_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/callisto13/fizzbuzz"
)

var _ = Describe("Fizzbuzz", func() {
	Context("knows when a number is", func() {
		It("divisible by 3", func() {
			Expect(fizzbuzz.IsDivisibleByThree(3)).To(BeTrue())
		})
	})
})
