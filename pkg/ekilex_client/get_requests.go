package ekilex_client

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func GetWordList(c *EkiClient, datasets []string, searchterm string) (output WordList, err error) {
	if datasets == nil {
		return output, errors.New("no datasets specified")
	}
	if searchterm == "" {
		return output, errors.New("no search term specified")
	}

	err = c.Get(fmt.Sprintf("/api/word/search/%s/%s", searchterm, strings.Join(datasets, ",")), &output)
	return output, err
}

func GetWordDetails(c *EkiClient, wordID int) (output WordDetails, err error) {
	err = c.Get("/api/word/details/"+strconv.Itoa(wordID), &output)
	return output, err
}
