/*
PeerTube

Testing ChannelsSyncAPIService

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

func Test_peertube_api_sdk_ChannelsSyncAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ChannelsSyncAPIService AddVideoChannelSync", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ChannelsSyncAPI.AddVideoChannelSync(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChannelsSyncAPIService DelVideoChannelSync", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var channelSyncId int32

		httpRes, err := apiClient.ChannelsSyncAPI.DelVideoChannelSync(context.Background(), channelSyncId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChannelsSyncAPIService TriggerVideoChannelSync", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var channelSyncId int32

		httpRes, err := apiClient.ChannelsSyncAPI.TriggerVideoChannelSync(context.Background(), channelSyncId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
