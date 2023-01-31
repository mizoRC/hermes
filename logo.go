package main

import (
	"fmt"
	"io/ioutil"
)

import (
	"log"
	"os"
)

func DisplayLogo() {
	currentPath, err := os.Getwd()
	logoPath := fmt.Sprintf("%s/logo.txt", currentPath)
	file, err := os.Open(logoPath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	b, err := ioutil.ReadFile(logoPath)

	fmt.Println(string(b))
}
