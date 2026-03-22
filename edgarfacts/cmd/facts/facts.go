package main

import (
	"fmt"
	// "internal/facts"
	"github.com/tedlaraghu/golang/edgarfacts/internal/facts"
)

func main() {
	cik := "X0000886982"
	organization := "My Organization"
	name := "RT"
	email := "rt@example.com"

	factData, err := facts.LoadFacts(cik, name, organization, email)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(factData))
}
