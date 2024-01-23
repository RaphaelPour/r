package main_test

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/bsm/gomega/gbytes"
	"github.com/bsm/gomega/gexec"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("R", func() {
	When("binary is used correctly", Ordered, func() {
		var tempFile *os.File
		var tempDir string

		BeforeAll(func() {
			var err error
			tempDir, err = os.MkdirTemp(os.TempDir(), "")
			Expect(err).Should(BeNil())
		})

		AfterAll(func() {
			os.RemoveAll(tempDir)
		})

		BeforeEach(func(ctx SpecContext) {
			var err error
			tempFile, err = os.CreateTemp(tempDir, "")
			Expect(err).ShouldNot(HaveOccurred())
		})

		helper.WithRunningDomain(
			It("renames file", func() {
				renamed := filepath.Base(tempFile.Name()) + "2"
				session, err := gexec.Start(
					exec.Command("ci-build/r", tempFile.Name(), renamed),
					GinkgoWriter,
					GinkgoWriter,
				)
				Expect(err).ShouldNot(HaveOccurred())
				Eventually(session).Should(gexec.Exit())

				_, err = os.Stat(tempFile.Name())
				Expect(err).To(MatchError(os.IsNotExist, "It is not existing"))

				_, err = os.Stat(filepath.Join(tempDir, renamed))
				Expect(err).Should(BeNil())
			})
		)
	})

	When("binary is used incorrectly", func() {
		It("fails without any argument", func() {
			command := exec.Command("ci-build/r")
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).ShouldNot(HaveOccurred())
			Eventually(session).Should(gexec.Exit())
			Expect(session.Out).Should(gbytes.Say("usage: r <path> <new_filename>"))
			Expect(session.Err).Should(gbytes.Say(""))
		})
	})
})
