# CommentThreadPostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to [**VideoComment**](VideoComment.md) |  | [optional] 

## Methods

### NewCommentThreadPostResponse

`func NewCommentThreadPostResponse() *CommentThreadPostResponse`

NewCommentThreadPostResponse instantiates a new CommentThreadPostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentThreadPostResponseWithDefaults

`func NewCommentThreadPostResponseWithDefaults() *CommentThreadPostResponse`

NewCommentThreadPostResponseWithDefaults instantiates a new CommentThreadPostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *CommentThreadPostResponse) GetComment() VideoComment`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CommentThreadPostResponse) GetCommentOk() (*VideoComment, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CommentThreadPostResponse) SetComment(v VideoComment)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CommentThreadPostResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


