package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	simple "github.com/jtagcat/simple/v2/pkg"
	"github.com/rk-ee/virkvirivirvavirn/pkg/ngram"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

var pastakasAnanassCmd = &cobra.Command{
	Use:   "ppap",
	Short: "Find matches between eki and ngram0 dataset",
	Run: func(_ *cobra.Command, _ []string) {
		ngrams, err := ngram.NgramFromFile(pastakasAnanass_opts.ngramFile)
		if err != nil {
			log.Fatal(err)
		}

		var ngramWords []commWord
		for _, n := range ngrams {
			var total int
			for _, p := range n.Years {
				total += p.Unique
			}
			ngramWords = append(ngramWords, commWord{Value: n.Value, Popularity: total, MorphCode: n.Type})
		}

		ji, err := os.ReadFile(pastakasAnanass_opts.ekiPath)
		if err != nil {
			log.Fatal(err)
		}
		var ekiWords []commWord
		err = json.Unmarshal(ji, &ekiWords)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Matching %d 1grams to %d ekiWords", len(ngramWords), len(ekiWords))

		// Ngram.Value x ekiWords

		matches := matchCommWords(ekiWords, "et", ngramWords, "en")

		jo, err := json.MarshalIndent(matches, "", "  ")
		if err != nil {
			log.Fatalf("Failed to marshal EKI commWords: %e", err)
		}
		err = ioutil.WriteFile(pastakasAnanass_opts.matchOutputPath, jo, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed writing file: %e", err)
		}
	},
}

func matchCommWords(a []commWord, aID string, b []commWord, bID string) (matches *[]map[string]commWord) {
	res, _ := simple.Parallel(func(g *errgroup.Group, returnc chan map[string]commWord) error {
		for _, ax := range a {
			ax := ax
			g.Go(func() error {
				for _, bx := range b {
					if strings.EqualFold(ax.Value, bx.Value) {
						m := make(map[string]commWord)
						m[aID] = ax
						m[bID] = bx
						returnc <- m
					}
				}
				return nil
			})
		}
		return nil
	})
	return &res
}

var pastakasAnanass_opts struct {
	ekiPath         string
	ngramFile       string
	matchOutputPath string
}

func init() {
	rootCmd.AddCommand(pastakasAnanassCmd)
	pastakasAnanassCmd.PersistentFlags().StringVarP(&pastakasAnanass_opts.ekiPath, "eki", "e", "", "Processed eki input to act on")
	pastakasAnanassCmd.PersistentFlags().StringVarP(&pastakasAnanass_opts.ngramFile, "ngramf", "n", "", "Path to ngram .gz file.")
	pastakasAnanassCmd.PersistentFlags().StringVarP(&pastakasAnanass_opts.matchOutputPath, "out", "o", "", "File to output JSON pulls to")
}
