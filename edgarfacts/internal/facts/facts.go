package facts

import (
	"errors"
	"fmt"
	"net/http"
	"os"	
)

// Load Company Facts
func LoadFacts(cik, name, organization, email string) ([]byte, error) {

	//Define URL for API
	url := fmt.Sprintf("https://data.sec.gov/api/xbrl/companyfacts/CIK%s.json", cik)

	//Create client
	client := *http.Client{}

	//Prepare request
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	// Set custom user-agent
	userAgent := fmt.Sprintf("%s %s %s", organization, name, email)
	req.Header.Set("User-Agent", userAgent)

	// Make request
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	//Check Response Code
	if response.StatusCode != http.StatusOK {
		errorStatus := errors.New(
			fmt.Sprintf("Status Code != OK, %v", response.StatusCode)
		)
		return nil, StatusCode
	}

	// Read response body
	body, err := os.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
