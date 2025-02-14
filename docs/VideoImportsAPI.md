# \VideoImportsAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1VideosImportsIdCancelPost**](VideoImportsAPI.md#ApiV1VideosImportsIdCancelPost) | **Post** /api/v1/videos/imports/{id}/cancel | Cancel video import
[**ApiV1VideosImportsIdDelete**](VideoImportsAPI.md#ApiV1VideosImportsIdDelete) | **Delete** /api/v1/videos/imports/{id} | Delete video import
[**ImportVideo**](VideoImportsAPI.md#ImportVideo) | **Post** /api/v1/videos/imports | Import a video



## ApiV1VideosImportsIdCancelPost

> ApiV1VideosImportsIdCancelPost(ctx, id).Execute()

Cancel video import



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RustLangLatam/peertube_api_sdk"
)

func main() {
	id := int32(56) // int32 | Entity id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoImportsAPI.ApiV1VideosImportsIdCancelPost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoImportsAPI.ApiV1VideosImportsIdCancelPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Entity id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1VideosImportsIdCancelPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1VideosImportsIdDelete

> ApiV1VideosImportsIdDelete(ctx, id).Execute()

Delete video import



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RustLangLatam/peertube_api_sdk"
)

func main() {
	id := int32(56) // int32 | Entity id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoImportsAPI.ApiV1VideosImportsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoImportsAPI.ApiV1VideosImportsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Entity id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1VideosImportsIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportVideo

> VideoUploadResponse ImportVideo(ctx).Name(name).ChannelId(channelId).TargetUrl(targetUrl).MagnetUri(magnetUri).Torrentfile(torrentfile).Privacy(privacy).Category(category).Licence(licence).Language(language).Description(description).WaitTranscoding(waitTranscoding).GenerateTranscription(generateTranscription).Support(support).Nsfw(nsfw).Tags(tags).CommentsEnabled(commentsEnabled).CommentsPolicy(commentsPolicy).DownloadEnabled(downloadEnabled).OriginallyPublishedAt(originallyPublishedAt).ScheduleUpdate(scheduleUpdate).Thumbnailfile(thumbnailfile).Previewfile(previewfile).VideoPasswords(videoPasswords).Execute()

Import a video



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/RustLangLatam/peertube_api_sdk"
)

func main() {
	name := "name_example" // string | Video name
	channelId := int32(56) // int32 | Channel id that will contain this video
	targetUrl := TODO // string | remote URL where to find the import's source video (optional)
	magnetUri := TODO // string | magnet URI allowing to resolve the import's source video (optional)
	torrentfile := TODO // *os.File | Torrent file containing only the video file (optional)
	privacy := openapiclient.VideoPrivacySet(1) // VideoPrivacySet |  (optional)
	category := int32(56) // int32 | category id of the video (see [/videos/categories](#operation/getCategories)) (optional)
	licence := int32(56) // int32 | licence id of the video (see [/videos/licences](#operation/getLicences)) (optional)
	language := "language_example" // string | language id of the video (see [/videos/languages](#operation/getLanguages)) (optional)
	description := "description_example" // string | Video description (optional)
	waitTranscoding := true // bool | Whether or not we wait transcoding before publish the video (optional)
	generateTranscription := true // bool | **PeerTube >= 6.2** If enabled by the admin, automatically generate a subtitle of the video (optional)
	support := "support_example" // string | A text tell the audience how to support the video creator (optional)
	nsfw := true // bool | Whether or not this video contains sensitive content (optional)
	tags := []string{"Inner_example"} // []string | Video tags (maximum 5 tags each between 2 and 30 characters) (optional)
	commentsEnabled := true // bool | Deprecated in 6.2, use commentsPolicy instead (optional)
	commentsPolicy := openapiclient.VideoCommentsPolicySet(1) // VideoCommentsPolicySet |  (optional)
	downloadEnabled := true // bool | Enable or disable downloading for this video (optional)
	originallyPublishedAt := time.Now() // time.Time | Date when the content was originally published (optional)
	scheduleUpdate := *openapiclient.NewVideoScheduledUpdate(time.Now()) // VideoScheduledUpdate |  (optional)
	thumbnailfile := os.NewFile(1234, "some_file") // *os.File | Video thumbnail file (optional)
	previewfile := os.NewFile(1234, "some_file") // *os.File | Video preview file (optional)
	videoPasswords := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoImportsAPI.ImportVideo(context.Background()).Name(name).ChannelId(channelId).TargetUrl(targetUrl).MagnetUri(magnetUri).Torrentfile(torrentfile).Privacy(privacy).Category(category).Licence(licence).Language(language).Description(description).WaitTranscoding(waitTranscoding).GenerateTranscription(generateTranscription).Support(support).Nsfw(nsfw).Tags(tags).CommentsEnabled(commentsEnabled).CommentsPolicy(commentsPolicy).DownloadEnabled(downloadEnabled).OriginallyPublishedAt(originallyPublishedAt).ScheduleUpdate(scheduleUpdate).Thumbnailfile(thumbnailfile).Previewfile(previewfile).VideoPasswords(videoPasswords).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoImportsAPI.ImportVideo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportVideo`: VideoUploadResponse
	fmt.Fprintf(os.Stdout, "Response from `VideoImportsAPI.ImportVideo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportVideoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Video name | 
 **channelId** | **int32** | Channel id that will contain this video | 
 **targetUrl** | [**string**](string.md) | remote URL where to find the import&#39;s source video | 
 **magnetUri** | [**string**](string.md) | magnet URI allowing to resolve the import&#39;s source video | 
 **torrentfile** | [***os.File**](*os.File.md) | Torrent file containing only the video file | 
 **privacy** | [**VideoPrivacySet**](VideoPrivacySet.md) |  | 
 **category** | **int32** | category id of the video (see [/videos/categories](#operation/getCategories)) | 
 **licence** | **int32** | licence id of the video (see [/videos/licences](#operation/getLicences)) | 
 **language** | **string** | language id of the video (see [/videos/languages](#operation/getLanguages)) | 
 **description** | **string** | Video description | 
 **waitTranscoding** | **bool** | Whether or not we wait transcoding before publish the video | 
 **generateTranscription** | **bool** | **PeerTube &gt;&#x3D; 6.2** If enabled by the admin, automatically generate a subtitle of the video | 
 **support** | **string** | A text tell the audience how to support the video creator | 
 **nsfw** | **bool** | Whether or not this video contains sensitive content | 
 **tags** | **[]string** | Video tags (maximum 5 tags each between 2 and 30 characters) | 
 **commentsEnabled** | **bool** | Deprecated in 6.2, use commentsPolicy instead | 
 **commentsPolicy** | [**VideoCommentsPolicySet**](VideoCommentsPolicySet.md) |  | 
 **downloadEnabled** | **bool** | Enable or disable downloading for this video | 
 **originallyPublishedAt** | **time.Time** | Date when the content was originally published | 
 **scheduleUpdate** | [**VideoScheduledUpdate**](VideoScheduledUpdate.md) |  | 
 **thumbnailfile** | ***os.File** | Video thumbnail file | 
 **previewfile** | ***os.File** | Video preview file | 
 **videoPasswords** | **[]string** |  | 

### Return type

[**VideoUploadResponse**](VideoUploadResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

