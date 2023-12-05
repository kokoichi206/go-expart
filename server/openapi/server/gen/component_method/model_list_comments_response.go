/*
Pet store schema

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package component_method

import (
	"encoding/json"
)

// checks if the ListCommentsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCommentsResponse{}

// ListCommentsResponse list comments
type ListCommentsResponse struct {
	Result []Comment `json:"result"`
	Total int32 `json:"total"`
}

// NewListCommentsResponse instantiates a new ListCommentsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCommentsResponse(result []Comment, total int32) *ListCommentsResponse {
	this := ListCommentsResponse{}
	this.Result = result
	this.Total = total
	return &this
}

// NewListCommentsResponseWithDefaults instantiates a new ListCommentsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCommentsResponseWithDefaults() *ListCommentsResponse {
	this := ListCommentsResponse{}
	return &this
}

// GetResult returns the Result field value
func (o *ListCommentsResponse) GetResult() []Comment {
	if o == nil {
		var ret []Comment
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *ListCommentsResponse) GetResultOk() ([]Comment, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result, true
}

// SetResult sets field value
func (o *ListCommentsResponse) SetResult(v []Comment) {
	o.Result = v
}

// GetTotal returns the Total field value
func (o *ListCommentsResponse) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ListCommentsResponse) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *ListCommentsResponse) SetTotal(v int32) {
	o.Total = v
}

func (o ListCommentsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCommentsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	toSerialize["total"] = o.Total
	return toSerialize, nil
}

type NullableListCommentsResponse struct {
	value *ListCommentsResponse
	isSet bool
}

func (v NullableListCommentsResponse) Get() *ListCommentsResponse {
	return v.value
}

func (v *NullableListCommentsResponse) Set(val *ListCommentsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListCommentsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListCommentsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCommentsResponse(val *ListCommentsResponse) *NullableListCommentsResponse {
	return &NullableListCommentsResponse{value: val, isSet: true}
}

func (v NullableListCommentsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCommentsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

