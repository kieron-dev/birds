package hello_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"testing"
)

var (
	pathToCmd string
	err       error
)

var _ = BeforeSuite(func() {
	pathToCmd, err = gexec.Build("github.com/kieron-pivotal/birdpedia/cmd")
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

func TestHello(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Hello Suite")
}
