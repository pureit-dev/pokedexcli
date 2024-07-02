package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Define a struct that matches the structure of the JSON response
type LocationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetPokeLocations() error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// Create an instance of the LocationAreaResponse struct
	var locationArea LocationAreaResponse

	// Unmarshal the JSON data into the struct
	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		return err
	}

	// Now you can access the decoded data
	fmt.Printf("Total locations: %d\n", locationArea.Count)
	fmt.Printf("Next page URL: %s\n", locationArea.Next)
	
	// Print the names of the first few locations
	for i, result := range locationArea.Results {
		if i >= 5 {
			break
		}
		fmt.Printf("Location %d: %s\n", i+1, result.Name)
	}

	return nil
}
