package ngram // Decodes Google ngram v3 format (latest known, 20200217)

import (
	"bufio"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"

	gipkg "github.com/jtagcat/git-id/pkg"
	simple "github.com/jtagcat/simple/v2/pkg"
	"github.com/rk-ee/virkvirivirvavirn/pkg"
	"golang.org/x/sync/errgroup"
)

var ErrFormat = errors.New("invalid format")

type Ngram struct {
	Value string
	Type  string // enum: "", one of tags or standaloneTags
	Years map[int]NYear
}
type NYear struct {
	Occurrences int // overall occurrences in books of the year
	Unique      int // in how many books occurred in
}

// https://books.google.com/ngrams/info
var tags = []string{
	"NOUN",
	"VERB",
	"ADJ",  // adjective
	"ADV",  // adverb
	"PRON", // pronoun
	"DET",  // determiner or article
	"ADP",  // an adposition: either a preposition or a postposition
	"NUM",  // numeral
	"CONJ", // conjunction
	"PRT",  // particle
}

var standaloneTags = []string{
	"ROOT",  // root of the parse tree
	"START", // start of a sentence
	"END",   // end of a sentence
}

// Loads up ngrams from a dirpath/*.gz files
// Beware this is not the smartest way of doing it, with 202002217 1gram dataset, output is estimated to be 448GiB.
func NgramFromDir(dirpath string) (output []Ngram, _ error) {
	return simple.Parallel(func(g *errgroup.Group, returnc chan Ngram) error {
		return filepath.WalkDir(dirpath, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return fmt.Errorf("%q: %e", path, err)
			}

			if !d.IsDir() && strings.HasSuffix(path, ".gz") {
				path := path
				g.Go(func() error {
					f, err := gipkg.OpenFileExisting(path, os.O_RDONLY)
					if err != nil {
						return fmt.Errorf("%s: %w", path, err)
					}
					defer f.Close()
					gz, err := gzip.NewReader(f)
					if err != nil {
						return fmt.Errorf("%s: %w", path, err)
					}
					defer gz.Close()

					decodedFile, err := DecodeNgram(gz)
					if err != nil {
						return err
					}
					for _, o := range decodedFile {
						returnc <- o
					}
					return nil
				})
			}
			return nil
		})
	})
}

func NgramFromFile(path string) (output []Ngram, _ error) {
	f, err := gipkg.OpenFileExisting(path, os.O_RDONLY)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	gz, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	defer gz.Close()

	return DecodeNgram(gz)
}

// Empty lines will be silently skipped. output is not sorted.
func DecodeNgram(input io.Reader) (output []Ngram, _ error) {
	return simple.Parallel(func(g *errgroup.Group, returnc chan Ngram) error {
		scanner := bufio.NewScanner(input)
		for linenum := 1; scanner.Scan(); linenum++ {
			line, linenum := scanner.Text(), linenum // https://golang.org/doc/faq#closures_and_goroutines
			g.Go(func() error {
				return decodeNgramLine(line, linenum, returnc)
			})
		}
		return scanner.Err()
	})
}

func decodeNgramLine(line string, linenum int, result chan Ngram) error {
	if head, body, ok := strings.Cut(strings.TrimSpace(line), "\t"); !ok && line != "" {
		return fmt.Errorf("line %d does not contain any \\t: %w", linenum, ErrFormat)
	} else if ok {
		current := Ngram{Years: make(map[int]NYear)}

		parseHead := func(wg *sync.WaitGroup, head string, current *Ngram) {
			defer wg.Done()

			if pkg.StringsHasBothix(head, "_") { // _NOUN_
				inside := pkg.StringsTrimBothix(head, "_")
				for _, tag := range append(tags, standaloneTags...) {
					if inside == tag { // case sensitive
						current.Type = tag
						return
					}
				}
			}

			for _, tag := range tags { // something_NOUN
				if strings.HasSuffix(head, "_"+tag) {
					current.Value = strings.TrimSuffix(head, "_"+tag)
					current.Type = tag
					return
				}
			}

			current.Value = head // something
		}
		var wg sync.WaitGroup
		wg.Add(1)
		go parseHead(&wg, head, &current)

		// 1969,567,89\t...
		yearsCSV := strings.Split(body, "\t")
		for _, csvRaw := range yearsCSV { // avoiding extra goroutines for readability
			// and insignificant or potentially negative perf improvements

			// 1969,567,89
			csv := strings.Split(csvRaw, ",")
			if len(csv) != 3 {
				return fmt.Errorf("line %d year data %q must contain 3 csv values: %w", linenum, csvRaw, ErrFormat)
			}

			cint, err := pkg.StrconvSliceAtoi(csv)
			if err != nil {
				return fmt.Errorf("line %d year data %q must contain 3 csv integer values: %w", linenum, csvRaw, ErrFormat)
			}

			// assuming no duplicate years
			current.Years[cint[0]] = NYear{Occurrences: cint[1], Unique: cint[2]}
		}

		wg.Wait()
		result <- current
	}
	return nil
}
