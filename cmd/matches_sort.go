package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sort"

	"github.com/spf13/cobra"
)

var filterSortCmd = &cobra.Command{
	Use:   "fs",
	Short: "Sorts the output of matches.",
	Run: func(_ *cobra.Command, _ []string) {
		ji, err := os.ReadFile(filterSort_opts.input)
		if err != nil {
			log.Fatal(err)
		}
		var words []map[string]commWord
		err = json.Unmarshal(ji, &words)
		if err != nil {
			log.Fatal(err)
		}

		maxPopularity := make(map[string]int)
		for _, o := range words {
			for l, v := range o {
				if v.Popularity > maxPopularity[l] {
					maxPopularity[l] = v.Popularity
				}
			}
		}

		sort.Slice(words, func(i, j int) bool {
			sumPopularity := make(map[int]int)
			for _, n := range []int{i, j} {
				for l, v := range words[n] {
					sumPopularity[n] *= v.Popularity / maxPopularity[l]
				}
			}
			return sumPopularity[i] < sumPopularity[j]
		})

		//
		jo, err := json.MarshalIndent(words, "", "  ")
		if err != nil {
			log.Fatalf("Failed to marshal output: %e", err)
		}
		err = ioutil.WriteFile(filterSort_opts.output, jo, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed writing file: %e", err)
		}
	},
}

var filterSort_opts struct {
	input  string
	output string
}

func init() {
	rootCmd.AddCommand(filterSortCmd)
	filterSortCmd.PersistentFlags().StringVarP(&filterSort_opts.input, "in", "i", "", "A JSON file of matches.")
	filterSortCmd.PersistentFlags().StringVarP(&filterSort_opts.output, "out", "o", "", "File to output JSON to")
}
