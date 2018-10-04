package hello_test

import (
	"io/ioutil"
	"net"
	"net/http"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Hello", func() {
	var (
		session *gexec.Session
	)

	BeforeEach(func() {
		command := exec.Command(pathToCmd)
		session, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(func() bool {
			_, err := net.Dial("tcp", ":8080")
			return err == nil
		}).Should(BeTrue())
	})

	AfterEach(func() {
		session.Terminate()
	})

	It("greets the world", func() {
		res, err := http.Get("http://localhost:8080/hello")
		Expect(err).NotTo(HaveOccurred())

		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		Expect(err).NotTo(HaveOccurred())
		Expect(string(body)).To(Equal("Hello World!"))
	})

})
