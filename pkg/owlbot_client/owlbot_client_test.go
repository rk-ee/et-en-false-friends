package owlbot_client

import (
	"net/http"
	"sync"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

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

func testingClient() *OwlClient {
	testenv := tomlSecrets()
	return &OwlClient{
		Hclient:   &http.Client{},
		Apikey:    testenv.Apikey,
		Host:      "https://owlbot.info",
		Useragent: "gotest",
	}
}

func TestGetWord(t *testing.T) {
	for ip := 1; ip < 100; ip++ {
		var wg sync.WaitGroup
		defer wg.Wait()
		log.Printf("starting %d", ip)
		for i := 1; i < 100; i++ {
			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				got, err := GetWord(testingClient(), "owl")
				assert.Nil(t, err)
				assert.Equal(t, Result{"owl", []ResultDefinition{{"noun",
					"a nocturnal bird of prey with large eyes, a facial disc, a hooked beak, and typically a loud hooting call.",
					"I love reaching out into that absolute silence, when you can hear the owl or the wind.",
					"https://media.owlbot.info/dictionary/images/hhhhhhhhhhhhhhhhhhhu.jpg.400x400_q85_box-15,0,209,194_crop_detail.jpg",
					"🦉"}}, "oul"}, got)
			}(&wg)
		}
	}
}

// ownbot is 30req/1m
