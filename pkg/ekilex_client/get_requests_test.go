package ekilex_client

import (
	"reflect"
	"testing"
	"time"
)

func TestGetWordList(t *testing.T) {
	want := WordList{
		1, []Word{
			{
				232187, "seitse", "seitse",
				1, false, "est", false, false, false,
				[]string{"eki", "ety", "les"},
				0, nil,
				time.Date(2021, 8, 26, 10, 42, 0o6, 191000000, time.FixedZone("", 0)),
			},
		},
	}

	got, err := GetWordList(testingClient(), []string{"eki"}, "seitse")
	if err != nil {
		t.Errorf("testable returned error: %v", err)
	}
	if reflect.DeepEqual(got, want) == false {
		t.Errorf("got %v\nwant %v", got, want)
	}
}

func TestGetWord(t *testing.T) {
	got, err := GetWordDetails(testingClient(), 232187)
	if err != nil {
		t.Errorf("testable returned error: %v", err)
	}

	if got.Word.Value != "seitse" {
		t.Errorf("Word value: got %v, want %v", got.Word.Value, "seitse")
	}

	if got.Paradigms[0].Forms[1].Value != "seitsme" {
		t.Errorf("Form value: got %v, want %v", got.Paradigms[0].Forms[1].Value, "seitsme")
	}

	if got.Lexemes[0].Meaning.Definitions[0].Value != "põhiarv 7" {
		t.Errorf("Definition value: got %v, want %v", got.Lexemes[0].Meaning.Definitions[0].Value, "põhiarv 7")
	}
}
