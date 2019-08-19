package resources

import (
	"bufio"
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var _ = Describe("httpDataplaneInspectionClient", func() {
	Describe("List()", func() {
		It("should create url with tags and parse response", func() {
			meshName := "default"
			tags := map[string]string{
				"service": "mobile",
				"version": "v1",
			}

			client := httpDataplaneInspectionClient{
				Client: &http.Client{
					Transport: RoundTripperFunc(func(req *http.Request) (*http.Response, error) {
						Expect(req.URL.String()).To(Equal("/meshes/default/dataplane-inspections?tag=service%3Amobile&tag=version%3Av1"))

						file, err := os.Open(filepath.Join("testdata", "list-dataplane-inspections.json"))
						if err != nil {
							return nil, err
						}
						return &http.Response{
							StatusCode: http.StatusOK,
							Body:       ioutil.NopCloser(bufio.NewReader(file)),
						}, nil
					}),
				},
			}

			// when
			list, err := client.List(context.Background(), meshName, tags)
			// then
			Expect(err).ToNot(HaveOccurred())
			// and
			Expect(list.Items).To(HaveLen(1))
			Expect(list.Items[0].Meta.GetName()).To(Equal("one"))
			Expect(list.Items[0].Spec.Dataplane.Networking.Inbound[0].Tags).To(HaveKeyWithValue("service", "mobile"))
			Expect(list.Items[0].Spec.Dataplane.Networking.Inbound[0].Tags).To(HaveKeyWithValue("version", "v1"))

			Expect(list.Items[0].Spec.DataplaneInsight.Subscriptions).To(HaveLen(2))
		})
	})
})
