/*
PeerTube

Testing AutomaticTagsAPIService

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

func Test_peertube_api_sdk_go_AutomaticTagsAPIService(t *testing.T) {

	configuration := api.NewConfiguration()
	apiClient := api.NewAPIClient(configuration)

	t.Run("Test AutomaticTagsAPIService ApiV1AutomaticTagsAccountsAccountNameAvailableGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accountName string

		resp, httpRes, err := apiClient.AutomaticTagsAPI.ApiV1AutomaticTagsAccountsAccountNameAvailableGet(context.Background(), accountName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AutomaticTagsAPIService ApiV1AutomaticTagsPoliciesAccountsAccountNameCommentsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accountName string

		resp, httpRes, err := apiClient.AutomaticTagsAPI.ApiV1AutomaticTagsPoliciesAccountsAccountNameCommentsGet(context.Background(), accountName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AutomaticTagsAPIService ApiV1AutomaticTagsPoliciesAccountsAccountNameCommentsPut", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accountName string

		httpRes, err := apiClient.AutomaticTagsAPI.ApiV1AutomaticTagsPoliciesAccountsAccountNameCommentsPut(context.Background(), accountName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AutomaticTagsAPIService ApiV1AutomaticTagsServerAvailableGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AutomaticTagsAPI.ApiV1AutomaticTagsServerAvailableGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
