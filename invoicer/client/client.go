package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sushant102004/Traffic-Toll-Microservice/types"
)

// It will be used to do POST requests to Aggregator API.

type AggClient struct {
	Endpoint string
}

func NewAggClient(endpoint string) *AggClient {
	return &AggClient{
		Endpoint: endpoint,
	}
}

func (c *AggClient) PostDataToAPI(cDistance types.CalculatedDistance) error {
	b, err := json.Marshal(cDistance)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.Endpoint, bytes.NewReader(b))
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("The service response was %d", resp.StatusCode)
	}
	return nil

}
