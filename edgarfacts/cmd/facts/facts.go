package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

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
	flag.StringVar(&organization, "organization", "", "organization")
	flag.StringVar(&name, "name", "", "name")
	flag.StringVar(&email, "email", "", "email")

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

	fmt.Printf("Folder Path: %s\n", folderPath)

	//Upload to Google storage
	// fileName := fmt.Sprintf("%s.json", cik)
	// filePath := filepath.Join(folderPath, fileName)

	folderPath = strings.ReplaceAll(folderPath, "\\", "/")
	fileName := fmt.Sprintf("%s.json", cik)
	filePath := path.Join(folderPath, fileName)

	logger.Printf("Uploading Facts to %s on bucket %s\n", fileName, bucketName)

	err = storage.UploadBytes(factsData, bucketName, filePath)
	if err != nil {
		panic(err)
	}

}
