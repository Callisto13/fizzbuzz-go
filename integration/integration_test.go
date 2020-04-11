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
		fizzbuzzBinary  string
		fizzbuzzCommand *exec.Cmd
	)

	BeforeEach(func() {
		var err error
		fizzbuzzBinary, err = gexec.Build("github.com/callisto13/fizzbuzz", "-mod=vendor")
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		gexec.CleanupBuildArtifacts()
	})

	It("when the command line argument is 3, it prints 'fizz'", func() {
		fizzbuzzCommand = exec.Command(fizzbuzzBinary, "3")

		session, err := gexec.Start(fizzbuzzCommand, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
		Eventually(session.Out).Should(gbytes.Say("fizz\n"))
	})

	It("when the command line argument is 5, it prints 'buzz'", func() {
		fizzbuzzCommand = exec.Command(fizzbuzzBinary, "5")

		session, err := gexec.Start(fizzbuzzCommand, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
		Eventually(session.Out).Should(gbytes.Say("buzz\n"))
	})

	It("when the command line argument is not a number, it prints an error", func() {
		fizzbuzzCommand = exec.Command(fizzbuzzBinary, "!")

		session, err := gexec.Start(fizzbuzzCommand, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
		Eventually(session.Out).Should(gbytes.Say("arguments must be numbers\n"))
	})
})
