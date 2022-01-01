package ekilex_client

import (
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
)

type ConfEnv struct {
	Apikey string
}

func tomlSecrets() ConfEnv {
	var testenv ConfEnv
	if _, err := toml.DecodeFile("testenv.toml", &testenv); err != nil {
		log.Fatalf("Error reading sercrets file: %q", err)
	}
	return testenv
}

func testingClient() *EkiClient {
	testenv := tomlSecrets()
	return &EkiClient{
		Hclient:   &http.Client{},
		Apikey:    testenv.Apikey,
		Host:      "https://ekilex.ee",
		Useragent: "gotest",
	}
}
