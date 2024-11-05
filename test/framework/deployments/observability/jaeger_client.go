package observability

import (
	"encoding/json"

	"net/http"
	"sort"
)

type jaegerServicesOutput struct {
	Data []string `json:"data"`
}

func tracedServices(url string) ([]string, error) {
	resp, err := http.Get(url + "/api/services")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	output := &jaegerServicesOutput{}
	if err := json.NewDecoder(resp.Body).Decode(output); err != nil {
		return nil, err
	}
	sort.Strings(output.Data)
	return output.Data, nil
}
