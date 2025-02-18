/*
PeerTube

Testing AccountBlocksAPIService

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

func Test_peertube_api_sdk_go_AccountBlocksAPIService(t *testing.T) {

	configuration := api.NewConfiguration()
	apiClient := api.NewAPIClient(configuration)

	t.Run("Test AccountBlocksAPIService ApiV1BlocklistStatusGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AccountBlocksAPI.ApiV1BlocklistStatusGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountBlocksAPIService ApiV1ServerBlocklistAccountsAccountNameDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accountName string

		httpRes, err := apiClient.AccountBlocksAPI.ApiV1ServerBlocklistAccountsAccountNameDelete(context.Background(), accountName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountBlocksAPIService ApiV1ServerBlocklistAccountsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.AccountBlocksAPI.ApiV1ServerBlocklistAccountsGet(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountBlocksAPIService ApiV1ServerBlocklistAccountsPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.AccountBlocksAPI.ApiV1ServerBlocklistAccountsPost(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
