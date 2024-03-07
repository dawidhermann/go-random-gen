// Manages file content
package persistence

import (
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
)

// Reads file
func ReadFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Adds new item to the file
func AddNewEntry(filename string, content string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	id := uuid.New()
	t := time.Now()
	marshaledTime, err := t.MarshalText()
	if err != nil {
		return err
	}
	record := fmt.Sprintf("%s,%s,%s", id, content, marshaledTime)
	if _, err := file.WriteString(record); err != nil {
		return err
	}
	return nil
}
