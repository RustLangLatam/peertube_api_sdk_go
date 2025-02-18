package main

import (
	"context"
	"fmt"
	"github.com/RustLangLatam/peertube_api_sdk_go/api"
	"log"
)

func main() {
	baseUrl := "https://peertube.orderi.co"

	// Initialize API client
	config := api.NewConfigurationFromBaseURL(baseUrl)

	apiClient := api.NewAPIClient(config)

	ctx := context.Background()
	oauthClient, _, err := apiClient.SessionAPI.GetOAuthClient(ctx).Execute()
	if err != nil {
		log.Fatal(err)
	}

	oauthToken := apiClient.SessionAPI.GetOAuthToken(ctx).
		ClientId(oauthClient.GetClientId()).
		ClientSecret(oauthClient.GetClientSecret()).
		Password("1ncC1hrpB9KD@peer").
		Username("root").
		GrantType("password")

	if response, _, err := oauthToken.Execute(); err != nil {
		log.Fatal(err)
	} else {
		videoMap, _ := response.MarshalJSON()
		fmt.Printf("Video: %v\n", string(videoMap))
	}
}
