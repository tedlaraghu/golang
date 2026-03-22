package main

import (
	"fmt"

	"github.com/tedlaraghu/golang/edgarfacts/internal/facts"
	"github.com/tedlaraghu/golang/edgarfacts/internal/storage"
)

func main() {
	cik := "0000886982"
	organization := "My Organization"
	name := "RT"
	email := "rt@example.com"

	factData, err := facts.LoadFacts(cik, name, organization, email)

	if err != nil {
		panic(err)
	}

	// fmt.Println(string(factData))

	bucketName := "gostoragert"
	filePath := fmt.Sprintf("sec/edgar/facts/stage/%s.json", cik)

	err = storage.UploadBytes(factData, bucketName, filePath)
	if err != nil {
		panic(err)
	}

	//Log Upload
	fmt.Printf("Uploaded %s\n", cik)

}
