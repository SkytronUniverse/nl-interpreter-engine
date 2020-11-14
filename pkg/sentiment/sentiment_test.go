package sentiment_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"

	. "github.com/nl-interpreter-engine/pkg/sentiment"
	"github.com/nl-interpreter-engine/web/webfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sentiment", func() {

	var (
		fakeHTTPClient  *webfakes.FakeHTTPClient
		sentimentClient *Client
	)

	BeforeEach(func() {
		fakeHTTPClient = &webfakes.FakeHTTPClient{}
		sentimentClient = New(fakeHTTPClient)
		Expect(sentimentClient.HTTPClient).NotTo(BeNil())
	})

	Describe("InterpretSentiment", func() {
		BeforeEach(func() {
			fakeHTTPClient.DoStub = func(req *http.Request) (*http.Response, error) {
				resp := &http.Response{
					Body: ioutil.NopCloser(bytes.NewBufferString(`{"sentiment": "something"}`)),
				}
				return resp, nil
			}

		})
		It("Does not error", func() {
			bytes, err := sentimentClient.InterpretSentiment("fakefiles/fakefile.txt")
			Expect(bytes).NotTo(BeNil())
			Expect(err).NotTo(HaveOccurred())
		})

		It("Errors because file DNE", func() {
			_, err := sentimentClient.InterpretSentiment("file.dne")
			Expect(err).To(HaveOccurred())
		})

		Describe("HTTP request is not successful", func() {
			BeforeEach(func() {
				fakeHTTPClient.DoStub = func(req *http.Request) (*http.Response, error) {
					return nil, errors.New("It can only be attributable to human error”")
				}
			})

			It("Returns an error", func() {
				_, err := sentimentClient.InterpretSentiment("fakefiles/fakefile.txt")
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errors.New("It can only be attributable to human error”")))
			})
		})
	})

})
