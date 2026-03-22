package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/tedlaraghu/golang/edgarfacts/internal/facts"
	"github.com/tedlaraghu/golang/edgarfacts/internal/storage"
)

func main() {
	//Parse command line arguments
	var cik string
	var organization string
	var name string
	var email string

	flag.StringVar(&cik, "cik", "", "CIK of the company")
	flag.StringVar(&organization, "organization", "", "Organization name")
	flag.StringVar(&name, "name", "", "Name of the user")
	flag.StringVar(&email, "email", "", "Email")

	flag.Parse()

	//validate command line arguments
	if len(cik) != 10 {
		panic("CIK must be of length 10")
	}

	if organization == "" {
		panic("Organization name is required")
	}

	if name == "" {
		panic("Name is required")
	}

	if email == "" {
		panic("Email is required")
	}

	//Load Environment Variables
	bucketName := os.Getenv("BUCKET")
	folderPath := os.Getenv("STAGE")

	if bucketName == "" || folderPath == "" {
		panic("Error reading ENV")
	}

	//configure logger
	logger := log.New(os.Stdout, "", log.LstdFlags)

	// Load Data
	logger.Printf("Loading Facts for CIK :%s\n", cik)

	factsData, err := facts.LoadFacts(cik, organization, name, email)

	if err != nil {
		panic(err)
	}

	//Upload to Google storage
	fileName := fmt.Sprintf("%s.json", cik)
	filePath := filepath.Join(folderPath, fileName)

	logger.Printf("Uploading Facts to %s on bucket %s\n", fileName, bucketName)

	err = storage.UploadBytes(factsData, bucketName, filePath)
	if err != nil {
		panic(err)
	}

}

// func main() {
// 	cik := "0000886982"
// 	organization := "My Organization"
// 	name := "RT"
// 	email := "rt@example.com"

// 	factData, err := facts.LoadFacts(cik, name, organization, email)

// 	if err != nil {
// 		panic(err)
// 	}

// 	// fmt.Println(string(factData))

// 	bucketName := "gostoragert"
// 	filePath := fmt.Sprintf("sec/edgar/facts/stage/%s.json", cik)

// 	err = storage.UploadBytes(factData, bucketName, filePath)
// 	if err != nil {
// 		panic(err)
// 	}

// 	//Log Upload
// 	fmt.Printf("Uploaded %s\n", cik)

// }
