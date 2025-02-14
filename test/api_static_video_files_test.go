/*
PeerTube

Testing StaticVideoFilesAPIService

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

func Test_peertube_api_sdk_StaticVideoFilesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test StaticVideoFilesAPIService StaticStreamingPlaylistsHlsFilenameGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var filename string

		httpRes, err := apiClient.StaticVideoFilesAPI.StaticStreamingPlaylistsHlsFilenameGet(context.Background(), filename).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StaticVideoFilesAPIService StaticStreamingPlaylistsHlsPrivateFilenameGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var filename string

		httpRes, err := apiClient.StaticVideoFilesAPI.StaticStreamingPlaylistsHlsPrivateFilenameGet(context.Background(), filename).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StaticVideoFilesAPIService StaticWebVideosFilenameGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var filename string

		httpRes, err := apiClient.StaticVideoFilesAPI.StaticWebVideosFilenameGet(context.Background(), filename).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StaticVideoFilesAPIService StaticWebVideosPrivateFilenameGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var filename string

		httpRes, err := apiClient.StaticVideoFilesAPI.StaticWebVideosPrivateFilenameGet(context.Background(), filename).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
