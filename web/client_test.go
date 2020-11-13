package web_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/nl-interpreter-engine/web"
	. "github.com/nl-interpreter-engine/web"
)

var _ = Describe("Client", func() {

	Describe("NewWebClient", func() {
		It("Returns a new web client", func() {
			client := NewWebClient(nil)
			Expect(client).NotTo(BeNil())
			_, ok := client.HTTPClient.(web.HTTPClient)
			Expect(ok).To(BeTrue())
		})
	})

})
