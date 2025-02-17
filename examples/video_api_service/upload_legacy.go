// Package main provides an example of how to use the PeerTube API SDK to upload a video.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	// Import the PeerTube API SDK.
	openapiclient "github.com/RustLangLatam/peertube_api_sdk"
)

// main is the entry point of the program.
func main() {
	// Define the base URL of the PeerTube instance.
	const baseUrl = "https://peertube.orderi.co"

	// Create a new context for the API requests.
	ctx := context.Background()

	// Initialize a new API client configuration.
	config := openapiclient.NewConfigurationFromBaseURL(baseUrl)

	// Create a new API client instance.
	apiClient := openapiclient.NewAPIClient(config)

	// Get the OAuth client from the API.
	oauthClient, _, err := apiClient.SessionAPI.GetOAuthClient(ctx).Execute()
	if err != nil {
		// Log the error and exit if it occurs.
		log.Fatal(err)
	}

	// Create a new OAuth token request.
	oauthToken := apiClient.SessionAPI.GetOAuthToken(ctx).
		// Set the client ID from the OAuth client.
		ClientId(oauthClient.GetClientId()).
		// Set the client secret from the OAuth client.
		ClientSecret(oauthClient.GetClientSecret()).
		// Set the password for the OAuth token request.
		Password("1ncC1hrpB9KD@peer").
		// Set the username for the OAuth token request.
		Username("root").
		// Set the grant type for the OAuth token request.
		GrantType("password")

	// Execute the OAuth token request.
	response, _, err := oauthToken.Execute()
	if err != nil {
		// Log the error and exit if it occurs.
		log.Fatal(err)
	}

	// Add the OAuth token to the API client configuration.
	config.AddDefaultHeader("Authorization", "Bearer "+response.GetAccessToken())

	// Define the file path of the video to upload.
	filePath := "resources/video_test.mp4"

	// Open the video file.
	videoFile, err := os.Open(filePath)
	if err != nil {
		// Log the error and exit if it occurs.
		fmt.Println("Error opening file :", err)
		return
	}

	// Defer the closing of the video file.
	defer func(videoFile *os.File) {
		// Close the video file.
		err := videoFile.Close()
		if err != nil {
			// Log the error if it occurs.
			log.Println(err)
		}
	}(videoFile)

	// Upload the video using the API client.
	video, _, err := apiClient.VideoAPI.UploadLegacy(context.Background()).
		// Set the name of the video.
		Name("test sdk golang").
		// Set the channel ID of the video.
		ChannelId(1).
		// Set the privacy of the video.
		Privacy(1).
		// Set the category of the video.
		Category(1).
		// Set the description of the video.
		Description("test description sdk").
		// Enable comments on the video.
		CommentsEnabled(true).
		// Enable downloads of the video.
		DownloadEnabled(true).
		// Set the original publication date of the video.
		OriginallyPublishedAt(time.Now()).
		// Set the video file to upload.
		Videofile(videoFile).
		// Execute the video upload request.
		Execute()

	if err != nil {
		// Log the error and exit if it occurs.
		log.Fatal(err)
	}

	// Marshal the video response to JSON.
	videoMap, _ := video.MarshalJSON()

	// Print the video response.
	fmt.Printf("Video: %v\n", string(videoMap))
}
