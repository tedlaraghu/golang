package facts

import (
	"fmt"
	"internal/facts"
)

func main() {
	cik := "0000886982"
	organization := "My Organization"
	name := "RT"
	email := "rt@example.com"

	facts, err := facts.LoadFacts(cik, name, organization, email)

	if err != nil {
		panic(err)
	}

	fmt.Printf(string(facts))
}
