package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/RustLangLatam/peertube_api_sdk_go/api"
)

func main() {
	// Define the base URL of the PeerTube instance and the video ID to fetch.
	const (
		// The base URL of the PeerTube instance.
		baseUrl = "https://peertube.orderi.co"
	)

	// Initialize the API client configuration with the base URL and disable debugging.
	config := api.NewConfigurationFromBaseURL(baseUrl, false)

	// Create a new API client with the configuration.
	apiClient := api.NewAPIClient(config)

	// Fetch the Licences list using the API client.
	// The context.Background() is used to create a background context for the request.
	video, response, err := apiClient.VideoAPI.GetLicences(context.Background()).Execute()

	// Check if the request failed or the response status code is not 200.
	if err != nil || response.StatusCode != 200 {
		log.Fatalf("Failed to fetch licences: %v (status code: %d)", err, response.StatusCode)
	}

	// Marshal the licences object into indented JSON.
	videoJSON, err := json.MarshalIndent(video, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal licences to JSON: %v", err)
	}

	// Print the licences details in indented JSON format.
	fmt.Printf("Licences:\n%s\n", videoJSON)
}
