// Package main provides an example of how to use the PeerTube API SDK to upload a video.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/RustLangLatam/peertube_api_sdk_go/api"
)

func main() {
	// Define the base URL of the PeerTube instance.
	// This URL should be replaced with the actual URL of the PeerTube instance you want to interact with.
	baseUrl := "https://peertube.orderi.co"

	// Initialize API client
	// Create a new configuration for the API client using the base URL.
	config := api.NewConfigurationFromBaseURL(baseUrl)

	// Create a new API client.
	// This client will be used to make API requests to the PeerTube instance.
	apiClient := api.NewAPIClient(config)

	// Get the OAuth client from the API.
	// This is the first step in the OAuth authentication flow.
	oauthClient, _, err := apiClient.SessionAPI.GetOAuthClient(context.Background()).Execute()
	if err != nil {
		// If an error occurs, log the error and exit the program.
		log.Fatal(err)
	}

	// Create a new OAuth token request.
	// This request will be used to obtain an access token for the API client.
	oauthToken, _, err := apiClient.SessionAPI.GetOAuthToken(context.Background()).
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
		// Execute the OAuth token request.
		Execute()

	if err != nil {
		// If an error occurs, log the error and exit the program.
		log.Fatal(err)
	}

	// Add the OAuth token to the API client configuration.
	// This will allow the API client to make authenticated requests to the PeerTube instance.
	config.AddDefaultHeader("Authorization", "Bearer "+oauthToken.GetAccessToken())

	// Open the video file.
	// This file will be uploaded to the PeerTube instance.
	videoFile, err := os.Open("resources/video_test.mp4")
	if err != nil {
		// If an error occurs, print the error and return from the function.
		fmt.Println("Error opening file :", err)
		return
	}
	// Defer closing the video file until the function returns.
	defer videoFile.Close()

	// Open the thumbnail file.
	// This file will be uploaded to the PeerTube instance as the thumbnail for the video.
	thumbnail, err := os.Open("resources/thumbnail_test.png")
	if err != nil {
		// If an error occurs, print the error and return from the function.
		fmt.Println("Error opening file :", err)
		return
	}
	// Defer closing the thumbnail file until the function returns.
	defer thumbnail.Close()

	// Upload the video using the API client.
	// This will create a new video on the PeerTube instance.
	uploadResponse, _, err := apiClient.VideoAPI.UploadLegacy(context.Background()).
		// Set the name of the video.
		Name("uploadResponse test sdk golang").
		// Set the channel ID for the video.
		ChannelId(1).
		// Set the privacy level for the video.
		Privacy(1).
		// Set the category for the video.
		Category(1).
		// Set the description for the video.
		Description("test description golang sdk").
		// Enable comments for the video.
		CommentsEnabled(true).
		// Enable downloads for the video.
		DownloadEnabled(true).
		// Set the original publication date for the video.
		OriginallyPublishedAt(time.Now()).
		// Set the video file for the upload.
		Videofile(videoFile).
		// Execute the upload request.
		Execute()

	if err != nil {
		// If an error occurs, log the error and exit the program.
		log.Fatal(err)
	}

	// Marshal the uploadResponse response to JSON.
	// This will convert the response to a JSON object.
	video := uploadResponse.GetVideo()

	// Print the uploadResponse response.
	// This will print the ID, short UUID, and UUID of the newly created video.
	fmt.Printf("Id: %v\n", video.GetId())
	fmt.Printf("ShortUUID: %v\n", video.GetShortUUID())
	fmt.Printf("Uuid: %v\n", video.GetUuid())
}
