package main

import (
	"context"
	"fmt"
	"log"

	openapiclient "github.com/RustLangLatam/peertube_api_sdk"
)

func main() {
	const baseUrl = "https://peertube.orderi.co"

	// Initialize API client
	config := openapiclient.NewConfigurationFromBaseURL(baseUrl)

	apiClient := openapiclient.NewAPIClient(config)

	id := openapiclient.PtrInt32(29)

	// Fetch a video by ID
	videoID := openapiclient.Int32AsApiV1VideosOwnershipIdAcceptPostIdParameter(id)

	video, _, err := apiClient.VideoAPI.GetVideo(context.Background(), videoID).Execute()
	if err != nil {
		log.Fatal(err)
	}

	videoMap, _ := video.MarshalJSON()

	fmt.Printf("Video Title: %v\n", string(videoMap))
}
