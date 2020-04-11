package integration_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Integration", func() {
	var (
		fizzbuzzCommand *exec.Cmd
		fizzbuzzArgs    []string
		session         *gexec.Session
	)

	JustBeforeEach(func() {
		fizzbuzzCommand = exec.Command(fizzbuzzBinary, fizzbuzzArgs...)

		var err error
		session, err = gexec.Start(fizzbuzzCommand, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
	})

	Context("when a command line argument is received", func() {
		BeforeEach(func() {
			fizzbuzzArgs = []string{"3"}
		})

		It("prints the results", func() {
			Eventually(session.Out).Should(gbytes.Say("fizz\n"))
		})
	})

	Context("when several command line argument are received", func() {
		BeforeEach(func() {
			fizzbuzzArgs = []string{"1", "2", "3", "4", "5"}
		})

		It("it processes them all", func() {
			Eventually(session.Out).Should(gbytes.Say("1\n2\nfizz\n4\nbuzz\n"))
		})
	})

	Context("when the command line argument is not a number", func() {
		BeforeEach(func() {
			fizzbuzzArgs = []string{"!"}
		})

		It("it prints an error", func() {
			Eventually(session.Out).Should(gbytes.Say("input is not a number, skipping\n"))
		})
	})

	Context("when no arguments are provided", func() {
		BeforeEach(func() {
			fizzbuzzArgs = []string{}
		})

		It("it prints usage info", func() {
			Eventually(session.Out).Should(gbytes.Say("usage: fizzbuzz <number> \\[<number> <number> ...\\]\n"))
		})
	})
})
