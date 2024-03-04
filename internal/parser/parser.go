// Parses data from and to csv file
package parser

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	ErrEntryNotValid = errors.New("entry is not valid")
)

// Generator entry
type Entry struct {
	Id     string
	Person string
	Time   time.Time
}

// Creates new entry
func NewEntry(data string) (Entry, error) {
	entries := strings.Split(data, ",")
	if len(entries) != 3 {
		return Entry{}, ErrEntryNotValid
	}
	time, err := time.Parse(time.RFC3339, entries[3])
	if err != nil {
		return Entry{}, err
	}
	entry := Entry{
		Id:     entries[0],
		Person: entries[1],
		Time:   time,
	}
	return entry, nil
}

// Converts entry to CSV string
func (entry Entry) ToCsvString() (string, error) {
	timeStr, err := entry.Time.MarshalText()
	if err != nil {
		return "", err
	}
	csvStr := fmt.Sprintf("%s,%s,%s", entry.Id, entry.Person, timeStr)
	return csvStr, nil
}
