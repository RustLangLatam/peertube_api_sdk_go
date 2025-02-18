// Package main provides an example of how to use the PeerTube API SDK to upload a video.
package main

import (
	"context"
	"fmt"
	"github.com/RustLangLatam/peertube_api_sdk_go/api"
	"log"
	"os"
	"time"
)

// main is the entry point of the program.
func main() {
	// Define the base URL of the PeerTube instance.
	const baseUrl = "https://peertube.orderi.co"

	// Create a new context for the API requests.
	ctx := context.Background()

	// Initialize API client
	config := api.NewConfigurationFromBaseURL(baseUrl)

	// Create a new API client.
	apiClient := api.NewAPIClient(config)

	// Get the OAuth client from the API.
	oauthClient, _, err := apiClient.SessionAPI.GetOAuthClient(ctx).Execute()
	if err != nil {
		// Log the error and exit if it occurs.
		log.Fatal(err)
	}

	// Create a new OAuth token request.
	oauthToken, _, err := apiClient.SessionAPI.GetOAuthToken(ctx).
		// Set the client ID from the OAuth client.
		ClientId(oauthClient.GetClientId()).
		// Set the client secret from the OAuth client.
		ClientSecret(oauthClient.GetClientSecret()).
		// Set the password for the OAuth token request.
		Password("1ncC1hrpB9KD@peer").
		// Set the username for the OAuth token request.
		Username("root").
		// Set the grant type for the OAuth token request.
		GrantType("password").
		Execute()

	if err != nil {
		// Log the error and exit if it occurs.
		log.Fatal(err)
	}

	// Add the OAuth token to the API client configuration.
	config.AddDefaultHeader("Authorization", "Bearer "+oauthToken.GetAccessToken())

	// Define the file path of the uploadResponse to upload.
	filePath := "resources/video_test.mp4"

	// Open the uploadResponse file.
	videoFile, err := os.Open(filePath)
	if err != nil {
		// Log the error and exit if it occurs.
		fmt.Println("Error opening file :", err)
		return
	}

	// Defer the closing of the uploadResponse file.
	defer func(videoFile *os.File) {
		// Close the uploadResponse file.
		err := videoFile.Close()
		if err != nil {
			// Log the error if it occurs.
			log.Println(err)
		}
	}(videoFile)

	// Upload the uploadResponse using the API client.
	uploadResponse, _, err := apiClient.VideoAPI.UploadLegacy(ctx).
		// Set the name of the uploadResponse.
		Name("uploadResponse test sdk golang").
		// Set the channel ID of the uploadResponse.
		ChannelId(1).
		// Set the privacy of the uploadResponse.
		Privacy(1).
		// Set the category of the uploadResponse.
		Category(1).
		// Set the description of the uploadResponse.
		Description("test description golang sdk").
		// Enable comments on the uploadResponse.
		CommentsEnabled(true).
		// Enable downloads of the uploadResponse.
		DownloadEnabled(true).
		// Set the original publication date of the uploadResponse.
		OriginallyPublishedAt(time.Now()).
		// Set the uploadResponse file to upload.
		Videofile(videoFile).
		// Execute the uploadResponse upload request.
		Execute()

	if err != nil {
		// Log the error and exit if it occurs.
		log.Fatal(err)
	}

	// Marshal the uploadResponse response to JSON.
	video := uploadResponse.GetVideo()

	// Print the uploadResponse response.
	fmt.Printf("Id: %v\n", video.GetId())
	fmt.Printf("ShortUUID: %v\n", video.GetShortUUID())
	fmt.Printf("Uuid: %v\n", video.GetUuid())
}
