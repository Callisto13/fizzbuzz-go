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

		It("knows when a number is NOT divisible by 3", func() {
			Expect(fizzbuzz.IsDivisibleByThree(1)).To(BeFalse())
		})

		It("knows when a number is divisible by 5", func() {
			Expect(fizzbuzz.IsDivisibleByFive(5)).To(BeTrue())
		})

		It("knows when a number is NOT divisible by 5", func() {
			Expect(fizzbuzz.IsDivisibleByFive(1)).To(BeFalse())
		})

		It("knows when a number is divisible by 3 and  5", func() {
			Expect(fizzbuzz.IsDivisibleByThreeAndFive(15)).To(BeTrue())
		})

		It("knows when a number is NOT divisible by 3 or 5", func() {
			Expect(fizzbuzz.IsDivisibleByThreeAndFive(1)).To(BeFalse())
		})
	})
})
