/*
PeerTube

Testing HomepageAPIService

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

func Test_peertube_api_sdk_HomepageAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test HomepageAPIService ApiV1CustomPagesHomepageInstanceGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.HomepageAPI.ApiV1CustomPagesHomepageInstanceGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HomepageAPIService ApiV1CustomPagesHomepageInstancePut", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.HomepageAPI.ApiV1CustomPagesHomepageInstancePut(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
