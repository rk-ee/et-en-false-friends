package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	simple "github.com/jtagcat/simple/v2/pkg"
	"github.com/spf13/cobra"

	eki "github.com/rk-ee/virkvirivirvavirn/pkg/ekilex_client"
	retrywait "k8s.io/apimachinery/pkg/util/wait"
)

// ekiprocessCmd represents the ekiprocess command
var ekiprocessCmd = &cobra.Command{
	Use:   "ekiprocess",
	Short: "Compose commWords from pulled ekilex-eki data",
	Run: func(_ *cobra.Command, _ []string) {
		lemma, err := getLemma(ekiprocess_opts.lemmaloc)
		if err != nil {
			log.Fatal(err)
		}

		ji, err := os.ReadFile(ekiprocess_opts.jsonin)
		if err != nil {
			log.Fatal(err)
		}
		var ekiContent []eki.WordDetails
		err = json.Unmarshal(ji, &ekiContent)
		if err != nil {
			log.Fatal(err)
		}

		combined := combineLemma(ekiContent, lemma)

		jo, err := json.MarshalIndent(combined, "", "  ")
		if err != nil {
			log.Printf("Failed to marshal EKI commWords: %s", err)
		}
		err = ioutil.WriteFile(ekiprocess_opts.jsonout, jo, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed writing file: %e", err)
		}
	},
}

type lemmaType map[string]int

func getLemma(url string) (lemmaType, error) {
	out := make(lemmaType)

	var c http.Client
	err := simple.RetryOnError(retrywait.Backoff{
		Duration: 2 * time.Second,
		Steps:    4,
		Factor:   1.5,
		Jitter:   1,
	}, func() (bool, error) {
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return false, err
		}

		r, err := c.Do(req)
		if err != nil {
			return true, err
		}
		defer r.Body.Close()

		lscanner := bufio.NewScanner(r.Body)
		for i := 1; lscanner.Scan(); i++ {
			line := strings.TrimSpace(
				strings.TrimPrefix(lscanner.Text(), "\ufeff")) // uft-8-sig
			split := strings.SplitN(line, " ", 2)
			popularity, err := strconv.Atoi(split[0])
			if err != nil {
				return true, fmt.Errorf("column 0 is not int‽ line %d: %v", i, line)
			}
			out[split[1]] = popularity
		}

		return false, nil
	})

	return out, err
}

func combineLemma(ekiContent []eki.WordDetails, lemma lemmaType) (outWords []commWord) {
	for _, o := range ekiContent {
		var w commWord

		// definitions: .[].lexemes[].meaning.definitions[0].value' (def0: first, same thing can have multiple meanings)
		for _, l := range o.Lexemes {
			for _, d := range l.Meaning.Definitions {
				w.Definitions = append(w.Definitions, d.Value)
			}
		}

		// oppurtunistic popularity
		if v, ok := lemma[o.Word.Value]; ok {
			w.Popularity = v
		}

		w.Value = o.Word.Value
		for _, p := range o.Paradigms { // This is the fault in process
			for _, f := range p.Forms {
				outWords = append(outWords, commWord{
					Definitions: w.Definitions, Popularity: w.Popularity,
					MorphCode: f.MorphCode, Value: f.Value,
				})
			}
		}

	}
	return outWords
}

var ekiprocess_opts struct {
	lemmaloc string
	jsonin   string
	jsonout  string
}

func init() {
	rootCmd.AddCommand(ekiprocessCmd)

	ekiprocessCmd.PersistentFlags().StringVarP(&ekiprocess_opts.lemmaloc, "lemma-location", "l", "https://www.cl.ut.ee/ressursid/sagedused1/failid/lemma_kahanevas.txt", "where to fetch lemma dataset from")
	ekiprocessCmd.PersistentFlags().StringVarP(&ekiprocess_opts.jsonin, "in", "i", "", "JSON input to act on")
	ekiprocessCmd.PersistentFlags().StringVarP(&ekiprocess_opts.jsonout, "out", "o", "", "Path to output JSON to")
}
