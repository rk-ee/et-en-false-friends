package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	eki "github.com/rk-ee/virkvirivirvavirn/pkg/ekilex_client"
	"github.com/spf13/cobra"
)

// ekipullCmd represents the ekipull command
var ekipullCmd = &cobra.Command{
	Use:   "ekipull",
	Short: "Pull words from ekilex",
	RunE: func(_ *cobra.Command, _ []string) error {
		if ekipull_opts.ekiAPIkey == "" {
			return errors.New("API key empty")
		}
		if ekipull_opts.ekiAPIhost == "" {
			return errors.New("API host empty")
		}
		if ekipull_opts.jsonout == "" {
			return errors.New("no output file specified")
		}
		if err := ioutil.WriteFile(ekipull_opts.jsonout, []byte("delete me, testing writability"), os.ModePerm); err != nil {
			return fmt.Errorf("error writing permtest file to %s: %v", ekipull_opts.jsonout, err)
		}
		if err := os.Remove(ekipull_opts.jsonout); err != nil {
			return fmt.Errorf("error removing permtest file from %s: %v", ekipull_opts.jsonout, err)
		}

		client := &eki.EkiClient{
			Hclient:   &http.Client{},
			Apikey:    ekipull_opts.ekiAPIkey,
			Host:      ekipull_opts.ekiAPIhost,
			Useragent: ekipull_opts.ekiUserAgent,
		}

		log.Println("getting wordlist")
		words, err := eki.GetWordList(client, ekipull_opts.ekiDatasets, ekipull_opts.searchterm)
		if err != nil {
			log.Printf("err while getting list: %q", err)
		}

		var ekifullwords []eki.WordDetails
		// c := goccm.New(5) // max concurrent routines
		for _, w := range words.Words {
			log.Printf("details for %v", w.Value)
			d, err := eki.GetWordDetails(client, w.ID)
			if err != nil {
				log.Printf("err while getting details: %q", err)
			}

			ekifullwords = append(ekifullwords, d)
		}

		j, err := json.MarshalIndent(ekifullwords, "", "  ") // "*"
		if err != nil {
			log.Printf("Failed to marshal pulled EKI content: %s", err)
		}
		return ioutil.WriteFile(ekipull_opts.jsonout, j, os.ModePerm)
	},
}

var ekipull_opts struct {
	ekiAPIkey    string
	ekiAPIhost   string
	ekiUserAgent string
	ekiDatasets  []string
	searchterm   string
	jsonout      string
}

func init() {
	rootCmd.AddCommand(ekipullCmd)

	ekipullCmd.PersistentFlags().StringVarP(&ekipull_opts.ekiAPIkey, "apikey", "k", "", "Ekilex API key")
	ekipullCmd.PersistentFlags().StringVarP(&ekipull_opts.ekiAPIhost, "host", "l", "https://ekilex.ee", "Ekilex API host")
	ekipullCmd.PersistentFlags().StringVarP(&ekipull_opts.ekiUserAgent, "user-agent", "a", "vvvv v0", "User Agent to use with Ekilex API")
	ekipullCmd.PersistentFlags().StringArrayVarP(&ekipull_opts.ekiDatasets, "datasets", "d", []string{"eki"}, "Remote datasets codes to operate with")
	ekipullCmd.PersistentFlags().StringVarP(&ekipull_opts.searchterm, "search", "s", "*", "Search term to list words")
	ekipullCmd.PersistentFlags().StringVarP(&ekipull_opts.jsonout, "out", "o", "", "File to output JSON pulls to")
}
