package zipf

import (
	"errors"
	"regexp"
)

var (
	rexWords   = regexp.MustCompile(`\W+`)
	rexSymbols = regexp.MustCompile(`[\w\s]+`)
)

// SplitWord return any valid word inside s.
func SplitWord(s string) ([]string, error) {
	ss := rexWords.Split(s, -1)
	var res []string
	for _, x := range ss {
		if x == "" {
			continue
		}
		res = append(res, x)
	}
	if len(res) == 0 {
		return nil, errors.New("not words found")
	}
	return res, nil
}

// SplitSymbol return any non-valid char inside s.
func SplitSymbol(s string) ([]string, error) {
	ss := rexSymbols.Split(s, -1)
	var res []string
	for _, x := range ss {
		if x == "" {
			continue
		}
		res = append(res, x)
	}
	if len(res) == 0 {
		return nil, errors.New("not symbols found")
	}
	return res, nil
}
