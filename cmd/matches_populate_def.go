package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	simple "github.com/jtagcat/simple/v2/pkg"
	"github.com/spf13/cobra"
	retrywait "k8s.io/apimachinery/pkg/util/wait"
)

var populateDefCmd = &cobra.Command{
	Use:   "pd",
	Short: "Populates definition for english words (unused)",
	Run: func(_ *cobra.Command, _ []string) {
		ji, err := os.ReadFile(populateDef_opts.input)
		if err != nil {
			log.Fatal(err)
		}
		var words []map[string]commWord
		err = json.Unmarshal(ji, &words)
		if err != nil {
			log.Fatal(err)
		}

		for _, w := range words {
			word := w["en"].Value
			defs, err := getDef(word)
			if err != nil {
				log.Println(fmt.Errorf("Failed getting defs for %s, %w", word, err))
				continue
			}
			new := w["en"]
			new.Definitions = defs
			w["en"] = new
		}

		//
		jo, err := json.MarshalIndent(words, "", "  ")
		if err != nil {
			log.Fatalf("Failed to marshal output: %e", err)
		}
		err = ioutil.WriteFile(populateDef_opts.output, jo, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed writing file: %e", err)
		}
	},
}

type DefType struct {
	Word      string
	Phonetic  string
	Phonetics []DefPhonetics
	Origin    string
	Meanings  []DefMeaning
}
type DefPhonetics struct {
	Text     string
	AudioUrl string
}
type DefMeaning struct {
	PartOfSpeech string
	Definitions  []DefDef
}
type DefDef struct {
	Definition string
	Example    string
	Synonyms   []string
	Antonyms   []string
}

func getDef(word string) (out []string, _ error) {
	c := &http.Client{Timeout: 5 * time.Second}
	err := simple.RetryOnError(retrywait.Backoff{
		Duration: 2 * time.Second,
		Steps:    4,
		Factor:   1.5,
		Jitter:   1,
	}, func() (bool, error) {
		r, err := c.Get("https://api.dictionaryapi.dev/api/v2/entries/en/" + word)
		if err != nil {
			return true, err
		}
		defer r.Body.Close()

		var target []DefType
		err = json.NewDecoder(r.Body).Decode(&target)
		if err != nil {
			return true, err
		}

		for _, t := range target {
			for _, m := range t.Meanings {
				for _, d := range m.Definitions {
					out = append(out, d.Definition)
				}
			}
		}

		return false, nil
	})
	return out, err
}

var populateDef_opts struct {
	input  string
	output string
}

func init() {
	rootCmd.AddCommand(populateDefCmd)
	populateDefCmd.PersistentFlags().StringVarP(&populateDef_opts.input, "in", "i", "", "A JSON file of matches.")
	populateDefCmd.PersistentFlags().StringVarP(&populateDef_opts.output, "out", "o", "", "File to output JSON to")
}
