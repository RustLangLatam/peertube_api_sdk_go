package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/RustLangLatam/peertube_api_sdk_go/api"
	"github.com/RustLangLatam/peertube_api_sdk_go/models"
	"github.com/RustLangLatam/peertube_api_sdk_go/utils"
)

func main() {
	// Define the base URL of the PeerTube instance and the video ID to fetch.
	const (
		// The base URL of the PeerTube instance.
		baseUrl = "https://peertube.orderi.co"
		// The ID of the video to fetch.
		videoID = 29
	)

	// Initialize the API client configuration with the base URL and disable debugging.
	config := api.NewConfigurationFromBaseURL(baseUrl, false)

	// Create a new API client with the configuration.
	apiClient := api.NewAPIClient(config)

	// Convert the video ID to the required API parameter type.
	videoIDParam := models.Int32AsApiV1VideosOwnershipIdAcceptPostIdParameter(utils.PtrInt32(videoID))

	// Fetch the video details using the API client.
	// The context.Background() is used to create a background context for the request.
	video, response, err := apiClient.VideoAPI.GetVideo(context.Background(), videoIDParam).Execute()

	// Check if the request failed or the response status code is not 200.
	if err != nil || response.StatusCode != 200 {
		log.Fatalf("Failed to fetch video: %v (status code: %d)", err, response.StatusCode)
	}

	// Marshal the video object into indented JSON.
	videoJSON, err := json.MarshalIndent(video, "", "  ")
	if err != nil {
		// Log a fatal error if the marshaling fails.
		log.Fatalf("Failed to marshal video to JSON: %v", err)
	}

	// Print the video details in indented JSON format.
	fmt.Printf("Video Details:\n%s\n", videoJSON)
}
