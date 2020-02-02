package util

import (
	"io/ioutil"
	"log"
)

/*
WriteFile takes a string and a path and writes to that file
*/
func WriteFile( path string, content string) (error) {
    message := []byte(content)
	err := ioutil.WriteFile(path, message, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return err
}
