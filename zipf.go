package zipf

import (
	"bufio"
	"errors"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

var (
	errEmpty = errors.New("zipf: term is empty")
)

// Zipf type.
type Zipf struct {
	path       string
	limit      int
	words      map[string]int32
	counts     map[int32]string
	collection []Term
	sync.RWMutex
}

// New returns a Zipf analiser.
func New(dir string, limit int) *Zipf {
	z := &Zipf{
		path:   dir,
		limit:  limit,
		words:  make(map[string]int32),
		counts: make(map[int32]string),
	}
	return z
}

// Add queue words to the map of words and sums 1 to existent words.
func (z *Zipf) Add(s string) error {
	z.RLock()
	defer z.RUnlock()

	if len(s) < 1 {
		return errEmpty
	}
	count, ok := z.words[s]
	if !ok {
		z.words[s] = 1
	}
	z.words[s] = count + 1
	return nil
}

// Walk read all files in dir and populate the word's count.
func (z *Zipf) Walk(dir string) error {
	err := filepath.Walk(dir, func(name string, info os.FileInfo, err error) error {
		// skip directories
		if info != nil && info.IsDir() {
			return nil
		}

		// read file
		lines, err := readLines(name)
		if err != nil {
			return err
		}

		for i := range lines {
			line := lines[i]
			// skip empty lines
			if len(line) < 1 {
				continue
			}

			words := processLine(line)
			if err != nil {
				// we don't return error here because .DS_Store file is created automatically
				//
				// if buggy we need a rule to skip files later.
				log.Printf("Walk : err [%s]", err)
				continue
			}
			for _, w := range words {
				if err := z.Add(w); err != nil {
					return err
				}
			}
		}
		return nil
	})
	return err
}

// Report report words count without order.
func (z *Zipf) Report() error {
	z.RLock()
	defer z.RUnlock()

	var i int
	for k, c := range z.words {
		i++
		if i > z.limit {
			continue
		}
		z.collection = append(z.collection, Term{Word: k, Count: c})
	}

	sort.Sort(ByCountAsc(z.collection))

	for i := range z.collection {
		log.Printf("[%s] [%v]", z.collection[i].Word, z.collection[i].Count)
	}
	return nil
}

func readLines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var lines []string
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		line := scan.Text()
		if len(line) < 1 {
			continue
		}
		lines = append(lines, line)
	}
	if err := f.Close(); err != nil {
		return nil, err
	}
	return lines, scan.Err()
}

func processLine(line string) []string {
	return strings.Split(line, " ")
}

// Term struct contain final struct for terms/words
type Term struct {
	Word  string `json:"word"`
	Count int32  `json:"count"`
}
