package utils

import (
	"io/ioutil"
	"log"
)

func ReadInputFile(path string) ([]byte, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return b, nil
}
