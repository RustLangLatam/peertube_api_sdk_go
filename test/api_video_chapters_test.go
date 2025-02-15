/*
PeerTube

Testing VideoChaptersAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package peertube_api_sdk

import (
	"context"
	openapiclient "github.com/RustLangLatam/peertube_api_sdk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_peertube_api_sdk_VideoChaptersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test VideoChaptersAPIService GetVideoChapters", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id openapiclient.ApiV1VideosOwnershipIdAcceptPostIdParameter

		resp, httpRes, err := apiClient.VideoChaptersAPI.GetVideoChapters(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VideoChaptersAPIService ReplaceVideoChapters", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id openapiclient.ApiV1VideosOwnershipIdAcceptPostIdParameter

		httpRes, err := apiClient.VideoChaptersAPI.ReplaceVideoChapters(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
