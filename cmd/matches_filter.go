package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var filterMatchesCmd = &cobra.Command{
	Use:   "fm",
	Short: "Filters and processes the output of matches.",
	Run: func(_ *cobra.Command, _ []string) {
		ji, err := os.ReadFile(filterMatches_opts.input)
		if err != nil {
			log.Fatal(err)
		}
		var words []map[string]commWord
		err = json.Unmarshal(ji, &words)
		if err != nil {
			log.Fatal(err)
		}

		for i, w := range words {
			for l, v := range w {
				words[i][l] = commWord{
					Value:       strings.ToLower(v.Value),
					Definitions: v.Definitions, Popularity: v.Popularity, MorphCode: v.MorphCode,
				}
			}
		}

		//
		jo, err := json.MarshalIndent(words, "", "  ")
		if err != nil {
			log.Fatalf("Failed to marshal output: %e", err)
		}
		err = ioutil.WriteFile(filterMatches_opts.output, jo, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed writing file: %e", err)
		}
	},
}

var filterMatches_opts struct {
	input  string
	output string
}

func init() {
	rootCmd.AddCommand(filterMatchesCmd)
	filterMatchesCmd.PersistentFlags().StringVarP(&filterMatches_opts.input, "in", "i", "", "A JSON file of matches.")
	filterMatchesCmd.PersistentFlags().StringVarP(&filterMatches_opts.output, "out", "o", "", "File to output JSON to")
}
