// Package main contains zipf's law analyzer for any text file.
package main

import (
	"flag"
	"log"
	"os"

	"github.com/jimmy-go/zipf"
)

var (
	dir     = flag.String("path", "", "Directory.")
	limit   = flag.Int("limit", 100, "Number of words to display.")
	symbols = flag.Bool("symbols-enabled", false, "Count symbols enabled.")
)

func main() {
	flag.Parse()
	log.SetFlags(log.Lshortfile)

	z, err := zipf.New(*dir, *limit, *symbols, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
	if err := z.Run(); err != nil {
		log.Fatal(err)
	}
}
