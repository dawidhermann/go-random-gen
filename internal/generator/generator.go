// Selects random entry
package generator

import (
	"errors"
	"math"
	"math/rand"
	"slices"
	"time"

	"github.com/dawidhermann/go-random-gen/internal/parser"
)

var ErrNoEntries = errors.New("entries slice is empty")

// Returns random entry from passed slice
func GetRandomEntry(entries []parser.Entry) (parser.Entry, error) {
	entriesSize := len(entries)
	if entriesSize == 0 {
		return parser.Entry{}, ErrNoEntries
	}
	sortEntries(entries)
	size := math.Ceil(float64(entriesSize) / 3)
	availableEntries := entries[0:int(size)]
	randomEntryIndex := getRandomNumber(len(availableEntries))
	return entries[randomEntryIndex], nil
}

// Sorts entries by time
func sortEntries(entries []parser.Entry) {
	slices.SortFunc(entries, func(a, b parser.Entry) int {
		return a.Time.Compare(b.Time)
	})
}

// generates random number between 0 and slice length
func getRandomNumber(max int) int {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	return r.Intn(max)
}
