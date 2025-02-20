/*
PeerTube

Testing JobAPIService

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

func Test_peertube_api_sdk_go_JobAPIService(t *testing.T) {

	configuration := api.NewConfiguration()
	apiClient := api.NewAPIClient(configuration)

	t.Run("Test JobAPIService ApiV1JobsPausePost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.JobAPI.ApiV1JobsPausePost(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test JobAPIService ApiV1JobsResumePost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.JobAPI.ApiV1JobsResumePost(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test JobAPIService GetJobs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var state string

		resp, httpRes, err := apiClient.JobAPI.GetJobs(context.Background(), state).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
