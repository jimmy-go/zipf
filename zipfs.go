// Package zipfs contains
//
// The MIT License (MIT)
//
// Copyright (c) 2016 Angel Del Castillo
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
package zipfs

// TODO; make concurrent walk path.

import (
	"bufio"
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

var (
	errEmpty = errors.New("zipfs: term is empty")
)

// Zipfs struct contains all data from analysis.
type Zipfs struct {
	words      map[string]int32
	counts     map[int32]string
	collection []Term
	sync.RWMutex
}

// New returns a Zipfs analyzer.
func New() *Zipfs {
	z := &Zipfs{
		words:  make(map[string]int32),
		counts: make(map[int32]string),
	}
	return z
}

// Add queue words to the map of words and sums 1 to existent words.
func (z *Zipfs) Add(s string) error {
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
func (z *Zipfs) Walk(dir string) error {
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
				z.Add(w)
			}
		}
		return nil
	})
	return err
}

// Report report words count without order.
func (z *Zipfs) Report(format string, limit int) {
	z.RLock()
	defer z.RUnlock()

	var i int
	for k, c := range z.words {
		i++
		if i > limit {
			continue
		}
		z.collection = append(z.collection, Term{Word: k, Count: c})
	}

	sort.Sort(ByCount(z.collection))

	switch format {
	default:
		for i := range z.collection {
			log.Printf("[%s] [%v]", z.collection[i].Word, z.collection[i].Count)
		}
	case "json":
		json.NewEncoder(os.Stdout).Encode(z.collection)
	}
}

func readLines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		line := scan.Text()
		if len(line) < 1 {
			continue
		}
		lines = append(lines, line)
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

// ByCount implement sort interface
type ByCount []Term

func (a ByCount) Len() int           { return len(a) }
func (a ByCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCount) Less(i, j int) bool { return a[i].Count > a[j].Count }
