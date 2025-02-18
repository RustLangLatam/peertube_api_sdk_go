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
	config := api.NewConfigurationFromBaseURL(baseUrl, true)

	apiClient := api.NewAPIClient(config)

	ctx := context.Background()
	oauthClient, _, err := apiClient.SessionAPI.GetOAuthClient(ctx).Execute()
	if err != nil {
		log.Fatal(err)
	}

	oauthToken, _, err := apiClient.SessionAPI.GetOAuthToken(ctx).
		ClientId(oauthClient.GetClientId()).
		ClientSecret(oauthClient.GetClientSecret()).
		Password("1ncC1hrpB9KD@peer").
		Username("root").
		GrantType("password").
		Execute()

	if err != nil {
		log.Fatal(err)
	}

	oauthTokenMap, _ := oauthToken.MarshalJSON()
	fmt.Printf("OauthToken: %v\n", string(oauthTokenMap))
}
