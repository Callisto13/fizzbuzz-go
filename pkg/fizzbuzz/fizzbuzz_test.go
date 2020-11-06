package fizzbuzz_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/callisto13/fizzbuzz-go/pkg/fizzbuzz"
)

var _ = Describe("Fizzbuzz", func() {
	Context("knows when a number", func() {
		It("is divisible by 3", func() {
			Expect(fizzbuzz.IsDivisibleBy(3, 3)).To(BeTrue())
		})

		It("is NOT divisible by 3", func() {
			Expect(fizzbuzz.IsDivisibleBy(1, 3)).To(BeFalse())
		})

		It("is divisible by 5", func() {
			Expect(fizzbuzz.IsDivisibleBy(5, 5)).To(BeTrue())
		})

		It("is NOT divisible by 5", func() {
			Expect(fizzbuzz.IsDivisibleBy(1, 5)).To(BeFalse())
		})

		It("is divisible by 3 and  5", func() {
			Expect(fizzbuzz.IsDivisibleBy(15, 15)).To(BeTrue())
		})

		It("is NOT divisible by 3 or 5", func() {
			Expect(fizzbuzz.IsDivisibleBy(1, 15)).To(BeFalse())
		})
	})

	Context("when playing the game, Fizzbuzz says...", func() {
		It("`fizz` when a number is divisible by 3", func() {
			Expect(fizzbuzz.Says(3)).To(Equal("fizz"))
		})

		It("`buzz` when a number is divisible by 5", func() {
			Expect(fizzbuzz.Says(5)).To(Equal("buzz"))
		})

		It("`buzz` when a number is divisible by 3 and 5", func() {
			Expect(fizzbuzz.Says(15)).To(Equal("fizzbuzz"))
		})

		It("the number when it is NOT divisible by 3 or 5", func() {
			Expect(fizzbuzz.Says(7)).To(Equal("7"))
		})
	})
})
