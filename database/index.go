package database

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type fileDatabase struct {
	file string
}

func NewFileDatabase (path string) fileDatabase{
	log.Println("SALVER!")
	return fileDatabase{path}
}

// Set receive a key and value and write in txt file
func (file fileDatabase) Set (key, value string) error {
	osFile, err := os.OpenFile(file.file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		return err
	}
	defer osFile.Close()
	if _, err = osFile.WriteString(fmt.Sprintf("%s:%s", key, value)); err != nil {
		return err
	}
	return nil
}

func (file  fileDatabase) Get(key string) (string, error) {
	osFile, err := os.OpenFile(file.file, os.O_RDONLY, 0600)
	if err != nil {
		return "", err
	}
	defer osFile.Close()
	scanner := bufio.NewScanner(osFile)
	var last string
	for scanner.Scan() {
		row := scanner.Text()
		parts := strings.Split(row, ":")
		if len(parts) < 2 {
			return "", errors.New("invalid")
		}
		if parts[0] == key {
			last = parts[1]
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	if last != "" {
		return last, nil
	}

	return "", errors.New("not found")

}