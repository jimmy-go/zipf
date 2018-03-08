package zipf

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitWord(t *testing.T) {
	table := []struct {
		Purpose string
		In      string
		Exp     []string
		Err     error
	}{
		{"1. OK", "func Do() error", []string{"func", "Do", "error"}, nil},
		{"2. OK", "func Do()", []string{"func", "Do"}, nil},
		{"3. Fail", ")//)", nil, errors.New("not words found")},
	}
	for _, x := range table {
		ss, err := SplitWord(x.In)
		assert.EqualValues(t, x.Err, err, x.Purpose)
		assert.EqualValues(t, x.Exp, ss, x.Purpose)
	}
}

func TestSplitSymbol(t *testing.T) {
	table := []struct {
		Purpose string
		In      string
		Exp     []string
		Err     error
	}{
		{"1. OK", "$var1 := something", []string{"$", ":="}, nil},
		{"2. Fail", "abc", nil, errors.New("not symbols found")},
	}
	for _, x := range table {
		ss, err := SplitSymbol(x.In)
		assert.EqualValues(t, x.Err, err, x.Purpose)
		assert.EqualValues(t, x.Exp, ss, x.Purpose)
	}
}
