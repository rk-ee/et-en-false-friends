package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLemma(t *testing.T) { // + test if resource has changed
	out, err := getLemma("https://www.cl.ut.ee/ressursid/sagedused1/failid/lemma_kahanevas.txt")
	assert.Nil(t, err)
	assert.Equal(t, 403162, out["ja"])
}

// This is the fault in process
// func TestCombineLemma(t *testing.T) {
// 	_, err := getLemma("https://www.cl.ut.ee/ressursid/sagedused1/failid/lemma_kahanevas.txt")
// 	assert.Nil(t, err)
// 	//	out := combineLemma([]eki.WordDetails{}, lemma)
// }
