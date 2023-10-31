/*
Pet store schema

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package component

// checks if the Comment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Comment{}

// Comment comment
type Comment struct {
	Comment1 string `json:"comment1"`
	Date string `json:"date"`
	Kijicode string `json:"kijicode"`
	Body string `json:"body"`
	Highlight []string `json:"highlight"`
}

// NewComment instantiates a new Comment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComment(comment1 string, date string, kijicode string, body string, highlight []string) *Comment {
	this := Comment{}
	this.Comment1 = comment1
	this.Date = date
	this.Kijicode = kijicode
	this.Body = body
	this.Highlight = highlight
	return &this
}

// NewCommentWithDefaults instantiates a new Comment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommentWithDefaults() *Comment {
	this := Comment{}
	return &this
}

// GetComment1 returns the Comment1 field value
func (o *Comment) GetComment1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comment1
}

// GetComment1Ok returns a tuple with the Comment1 field value
// and a boolean to check if the value has been set.
func (o *Comment) GetComment1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comment1, true
}

// SetComment1 sets field value
func (o *Comment) SetComment1(v string) {
	o.Comment1 = v
}

// GetDate returns the Date field value
func (o *Comment) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *Comment) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *Comment) SetDate(v string) {
	o.Date = v
}

// GetKijicode returns the Kijicode field value
func (o *Comment) GetKijicode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kijicode
}

// GetKijicodeOk returns a tuple with the Kijicode field value
// and a boolean to check if the value has been set.
func (o *Comment) GetKijicodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kijicode, true
}

// SetKijicode sets field value
func (o *Comment) SetKijicode(v string) {
	o.Kijicode = v
}

// GetBody returns the Body field value
func (o *Comment) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *Comment) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *Comment) SetBody(v string) {
	o.Body = v
}

// GetHighlight returns the Highlight field value
func (o *Comment) GetHighlight() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Highlight
}

// GetHighlightOk returns a tuple with the Highlight field value
// and a boolean to check if the value has been set.
func (o *Comment) GetHighlightOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Highlight, true
}

// SetHighlight sets field value
func (o *Comment) SetHighlight(v []string) {
	o.Highlight = v
}

func (o Comment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["comment1"] = o.Comment1
	toSerialize["date"] = o.Date
	toSerialize["kijicode"] = o.Kijicode
	toSerialize["body"] = o.Body
	toSerialize["highlight"] = o.Highlight
	return toSerialize, nil
}

type NullableComment struct {
	value *Comment
	isSet bool
}

func (v NullableComment) Get() *Comment {
	return v.value
}

func (v *NullableComment) Set(val *Comment) {
	v.value = val
	v.isSet = true
}

func (v NullableComment) IsSet() bool {
	return v.isSet
}

func (v *NullableComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComment(val *Comment) *NullableComment {
	return &NullableComment{value: val, isSet: true}
}
