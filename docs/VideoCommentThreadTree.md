# VideoCommentThreadTree

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to [**VideoComment**](VideoComment.md) |  | [optional] 
**Children** | Pointer to [**[]VideoCommentThreadTree**](VideoCommentThreadTree.md) |  | [optional] 

## Methods

### NewVideoCommentThreadTree

`func NewVideoCommentThreadTree() *VideoCommentThreadTree`

NewVideoCommentThreadTree instantiates a new VideoCommentThreadTree object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoCommentThreadTreeWithDefaults

`func NewVideoCommentThreadTreeWithDefaults() *VideoCommentThreadTree`

NewVideoCommentThreadTreeWithDefaults instantiates a new VideoCommentThreadTree object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *VideoCommentThreadTree) GetComment() VideoComment`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *VideoCommentThreadTree) GetCommentOk() (*VideoComment, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *VideoCommentThreadTree) SetComment(v VideoComment)`

SetComment sets Comment field to given value.

### HasComment

`func (o *VideoCommentThreadTree) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetChildren

`func (o *VideoCommentThreadTree) GetChildren() []VideoCommentThreadTree`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *VideoCommentThreadTree) GetChildrenOk() (*[]VideoCommentThreadTree, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *VideoCommentThreadTree) SetChildren(v []VideoCommentThreadTree)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *VideoCommentThreadTree) HasChildren() bool`

HasChildren returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


