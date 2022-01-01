package owlbot_client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type OwlClient struct {
	Hclient   *http.Client
	Apikey    string
	Host      string
	Useragent string
}

// could be made generic, but would need prefix, and authorization vars-validation
// is there a good lib available?
func (c *OwlClient) owlGet(path string, output_struct interface{}) error {
	if c.Apikey == "" {
		return fmt.Errorf("API key not set")
	}
	if c.Host == "" {
		return fmt.Errorf("host is not set")
	}
	Host := strings.TrimSuffix(c.Host, "/")
	if !strings.HasPrefix(path, "/api/v4/") {
		return fmt.Errorf("path should start with /api/v4/")
	}

	req, err := http.NewRequest(http.MethodGet, Host+path, nil)
	if err != nil {
		return fmt.Errorf("error crafting request: %w", err)
	}
	req.Header.Set("Authorization", "Token "+c.Apikey)
	req.Header.Set("User-Agent", c.Useragent)

	r, err := c.Hclient.Do(req)
	if err != nil {
		return fmt.Errorf("error performing request: %w", err)
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("error reading body: %w", err)
	}

	err = json.Unmarshal(body, &output_struct)
	if err != nil {
		return fmt.Errorf("error unmarshaling: %w", err)
	}
	return nil
}

func GetWord(c *OwlClient, word string) (out Result, err error) {
	return out, c.owlGet("/api/v4/dictionary/"+word, &out)
}
