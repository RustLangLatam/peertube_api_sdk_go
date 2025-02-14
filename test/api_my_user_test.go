/*
PeerTube

Testing MyUserAPIService

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

func Test_peertube_api_sdk_MyUserAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MyUserAPIService ApiV1UsersMeAvatarDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.MyUserAPI.ApiV1UsersMeAvatarDelete(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MyUserAPIService ApiV1UsersMeAvatarPickPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.MyUserAPI.ApiV1UsersMeAvatarPickPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MyUserAPIService ApiV1UsersMeVideoQuotaUsedGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.MyUserAPI.ApiV1UsersMeVideoQuotaUsedGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MyUserAPIService ApiV1UsersMeVideosGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.MyUserAPI.ApiV1UsersMeVideosGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MyUserAPIService ApiV1UsersMeVideosImportsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.MyUserAPI.ApiV1UsersMeVideosImportsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MyUserAPIService ApiV1UsersMeVideosVideoIdRatingGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var videoId int32

		resp, httpRes, err := apiClient.MyUserAPI.ApiV1UsersMeVideosVideoIdRatingGet(context.Background(), videoId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MyUserAPIService GetUserInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.MyUserAPI.GetUserInfo(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MyUserAPIService PutUserInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.MyUserAPI.PutUserInfo(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
