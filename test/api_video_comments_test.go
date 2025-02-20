/*
PeerTube

Testing VideoCommentsAPIService

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

func Test_peertube_api_sdk_go_VideoCommentsAPIService(t *testing.T) {

	configuration := api.NewConfiguration()
	apiClient := api.NewAPIClient(configuration)

	t.Run("Test VideoCommentsAPIService ApiV1UsersMeVideosCommentsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.VideoCommentsAPI.ApiV1UsersMeVideosCommentsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VideoCommentsAPIService ApiV1VideosCommentsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.VideoCommentsAPI.ApiV1VideosCommentsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VideoCommentsAPIService ApiV1VideosIdCommentThreadsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id ApiV1VideosOwnershipIdAcceptPostIdParameter

		resp, httpRes, err := apiClient.VideoCommentsAPI.ApiV1VideosIdCommentThreadsGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VideoCommentsAPIService ApiV1VideosIdCommentThreadsPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id ApiV1VideosOwnershipIdAcceptPostIdParameter

		resp, httpRes, err := apiClient.VideoCommentsAPI.ApiV1VideosIdCommentThreadsPost(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VideoCommentsAPIService ApiV1VideosIdCommentThreadsThreadIdGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id ApiV1VideosOwnershipIdAcceptPostIdParameter
		var threadId int32

		resp, httpRes, err := apiClient.VideoCommentsAPI.ApiV1VideosIdCommentThreadsThreadIdGet(context.Background(), id, threadId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VideoCommentsAPIService ApiV1VideosIdCommentsCommentIdApprovePost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id ApiV1VideosOwnershipIdAcceptPostIdParameter
		var commentId int32

		httpRes, err := apiClient.VideoCommentsAPI.ApiV1VideosIdCommentsCommentIdApprovePost(context.Background(), id, commentId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VideoCommentsAPIService ApiV1VideosIdCommentsCommentIdDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id ApiV1VideosOwnershipIdAcceptPostIdParameter
		var commentId int32

		httpRes, err := apiClient.VideoCommentsAPI.ApiV1VideosIdCommentsCommentIdDelete(context.Background(), id, commentId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VideoCommentsAPIService ApiV1VideosIdCommentsCommentIdPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id ApiV1VideosOwnershipIdAcceptPostIdParameter
		var commentId int32

		resp, httpRes, err := apiClient.VideoCommentsAPI.ApiV1VideosIdCommentsCommentIdPost(context.Background(), id, commentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
