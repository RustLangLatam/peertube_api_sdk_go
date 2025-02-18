/*
PeerTube

Testing VideoPasswordsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package peertube_api_sdk_go

import (
	"context"
	"github.com/RustLangLatam/peertube_api_sdk_go/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_peertube_api_sdk_go_VideoPasswordsAPIService(t *testing.T) {

	configuration := api.NewConfiguration()
	apiClient := api.NewAPIClient(configuration)

	t.Run("Test VideoPasswordsAPIService ApiV1VideosIdPasswordsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id ApiV1VideosOwnershipIdAcceptPostIdParameter

		resp, httpRes, err := apiClient.VideoPasswordsAPI.ApiV1VideosIdPasswordsGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VideoPasswordsAPIService ApiV1VideosIdPasswordsPut", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id ApiV1VideosOwnershipIdAcceptPostIdParameter

		httpRes, err := apiClient.VideoPasswordsAPI.ApiV1VideosIdPasswordsPut(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VideoPasswordsAPIService ApiV1VideosIdPasswordsVideoPasswordIdDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id ApiV1VideosOwnershipIdAcceptPostIdParameter
		var videoPasswordId int32

		httpRes, err := apiClient.VideoPasswordsAPI.ApiV1VideosIdPasswordsVideoPasswordIdDelete(context.Background(), id, videoPasswordId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
