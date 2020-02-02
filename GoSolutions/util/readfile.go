package util

import (
	"io/ioutil"
	"log"
)

/*
ReadFile reads a file and outputs its string representation
*/
func ReadFile(path string) (string, error) {
    content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return string(content), err
}
