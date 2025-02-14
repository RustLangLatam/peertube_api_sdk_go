# CommentThreadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | Total threads (included deleted ones) on this video | [optional] 
**TotalNotDeletedComments** | Pointer to **int32** | Total not-deleted threads (included deleted ones) on this video | [optional] 
**Data** | Pointer to [**[]VideoComment**](VideoComment.md) |  | [optional] 

## Methods

### NewCommentThreadResponse

`func NewCommentThreadResponse() *CommentThreadResponse`

NewCommentThreadResponse instantiates a new CommentThreadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentThreadResponseWithDefaults

`func NewCommentThreadResponseWithDefaults() *CommentThreadResponse`

NewCommentThreadResponseWithDefaults instantiates a new CommentThreadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *CommentThreadResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CommentThreadResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CommentThreadResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CommentThreadResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTotalNotDeletedComments

`func (o *CommentThreadResponse) GetTotalNotDeletedComments() int32`

GetTotalNotDeletedComments returns the TotalNotDeletedComments field if non-nil, zero value otherwise.

### GetTotalNotDeletedCommentsOk

`func (o *CommentThreadResponse) GetTotalNotDeletedCommentsOk() (*int32, bool)`

GetTotalNotDeletedCommentsOk returns a tuple with the TotalNotDeletedComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNotDeletedComments

`func (o *CommentThreadResponse) SetTotalNotDeletedComments(v int32)`

SetTotalNotDeletedComments sets TotalNotDeletedComments field to given value.

### HasTotalNotDeletedComments

`func (o *CommentThreadResponse) HasTotalNotDeletedComments() bool`

HasTotalNotDeletedComments returns a boolean if a field has been set.

### GetData

`func (o *CommentThreadResponse) GetData() []VideoComment`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CommentThreadResponse) GetDataOk() (*[]VideoComment, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CommentThreadResponse) SetData(v []VideoComment)`

SetData sets Data field to given value.

### HasData

`func (o *CommentThreadResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


