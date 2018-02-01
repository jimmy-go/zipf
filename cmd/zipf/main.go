// Package main contains zipf's law analyzer for any text file.
package main

import (
	"flag"
	"log"

	"github.com/jimmy-go/zipf"
)

var (
	dir   = flag.String("path", "", "Directory.")
	limit = flag.Int("limit", 500, "Number of words to display.")
)

func main() {
	flag.Parse()

	z := zipf.New(*dir, *limit)
	if err := z.Run(); err != nil {
		log.Fatal(err)
	}
}
