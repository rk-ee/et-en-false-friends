package ekilex_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	simple "github.com/jtagcat/simple/v2/pkg"
	retrywait "k8s.io/apimachinery/pkg/util/wait"
)

type EkiClient struct {
	Hclient   *http.Client
	Apikey    string
	Host      string
	Useragent string
}

func (c *EkiClient) Get(path string, output_struct interface{}) error {
	if !strings.HasPrefix(path, "/api/") {
		return fmt.Errorf("path should start with '/api/', %v does not", path)
	}
	if c.Apikey == "" {
		return errors.New("API key empty")
	}
	if c.Host == "" {
		return errors.New("host is empty")
	}
	Host := strings.TrimSuffix(c.Host, "/")

	return simple.RetryOnError(retrywait.Backoff{
		Duration: 10 * time.Second,
		Steps:    4,
		Factor:   2,
		Jitter:   1,
	}, func() (bool, error) {
		req, err := http.NewRequest(http.MethodGet, Host+path, nil)
		if err != nil {
			return false, err
		}
		req.Header.Set("ekilex-api-key", c.Apikey)
		req.Header.Set("User-Agent", c.Useragent)

		r, err := c.Hclient.Do(req)
		if err != nil {
			return true, err
		}
		defer r.Body.Close()

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return true, err
		}

		err = json.Unmarshal(body, &output_struct)
		if err != nil {
			return false, err
		}
		return false, nil
	})
}
