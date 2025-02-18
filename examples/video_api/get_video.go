package main

import (
	"context"
	"fmt"
	"github.com/RustLangLatam/peertube_api_sdk_go/api"
	"github.com/RustLangLatam/peertube_api_sdk_go/models"
	"github.com/RustLangLatam/peertube_api_sdk_go/utils"
	"log"
)

func main() {
	const baseUrl = "https://peertube.orderi.co"

	// Initialize API client
	config := api.NewConfigurationFromBaseURL(baseUrl, true)

	apiClient := api.NewAPIClient(config)

	id := utils.PtrInt32(29)

	// Fetch a video by ID
	videoID := models.Int32AsApiV1VideosOwnershipIdAcceptPostIdParameter(id)

	video, _, err := apiClient.VideoAPI.GetVideo(context.Background(), videoID).Execute()
	if err != nil {
		log.Fatal(err)
	}

	videoMap, _ := video.MarshalJSON()

	fmt.Printf("Video Title: %v\n", string(videoMap))
}
