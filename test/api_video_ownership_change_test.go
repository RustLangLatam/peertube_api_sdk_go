/*
PeerTube

Testing VideoOwnershipChangeAPIService

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

func Test_peertube_api_sdk_VideoOwnershipChangeAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test VideoOwnershipChangeAPIService ApiV1VideosIdGiveOwnershipPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id openapiclient.ApiV1VideosOwnershipIdAcceptPostIdParameter

		httpRes, err := apiClient.VideoOwnershipChangeAPI.ApiV1VideosIdGiveOwnershipPost(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VideoOwnershipChangeAPIService ApiV1VideosOwnershipGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.VideoOwnershipChangeAPI.ApiV1VideosOwnershipGet(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VideoOwnershipChangeAPIService ApiV1VideosOwnershipIdAcceptPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id openapiclient.ApiV1VideosOwnershipIdAcceptPostIdParameter

		httpRes, err := apiClient.VideoOwnershipChangeAPI.ApiV1VideosOwnershipIdAcceptPost(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VideoOwnershipChangeAPIService ApiV1VideosOwnershipIdRefusePost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id openapiclient.ApiV1VideosOwnershipIdAcceptPostIdParameter

		httpRes, err := apiClient.VideoOwnershipChangeAPI.ApiV1VideosOwnershipIdRefusePost(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
