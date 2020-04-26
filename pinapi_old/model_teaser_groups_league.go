/*
 * Pinnacle - Customer API Reference
 *
 *  # Authentication   API uses HTTP Basic access authentication. You need to send Authorization HTTP Request header:    `Authorization: Basic <Base64 value of UTF-8 encoded \"username:password\">`  Example:  `Authorization: Basic U03MyOT23YbzMDc6d3c3O1DQ1`
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pinapi_old

import (
	"bytes"
	"encoding/json"
)

// TeaserGroupsLeague struct for TeaserGroupsLeague
type TeaserGroupsLeague struct {
	// Unique identifier. League details can be retrieved from a call to v2/leagues endpoint.
	Id     *int32               `json:"id,omitempty" xml:"id"`
	Spread *TeaserGroupsBetType `json:"spread,omitempty" xml:"spread"`
	Total  *TeaserGroupsBetType `json:"total,omitempty" xml:"total"`
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TeaserGroupsLeague) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TeaserGroupsLeague) GetIdOk() (int32, bool) {
	if o == nil || o.Id == nil {
		var ret int32
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TeaserGroupsLeague) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *TeaserGroupsLeague) SetId(v int32) {
	o.Id = &v
}

// GetSpread returns the Spread field value if set, zero value otherwise.
func (o *TeaserGroupsLeague) GetSpread() TeaserGroupsBetType {
	if o == nil || o.Spread == nil {
		var ret TeaserGroupsBetType
		return ret
	}
	return *o.Spread
}

// GetSpreadOk returns a tuple with the Spread field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TeaserGroupsLeague) GetSpreadOk() (TeaserGroupsBetType, bool) {
	if o == nil || o.Spread == nil {
		var ret TeaserGroupsBetType
		return ret, false
	}
	return *o.Spread, true
}

// HasSpread returns a boolean if a field has been set.
func (o *TeaserGroupsLeague) HasSpread() bool {
	if o != nil && o.Spread != nil {
		return true
	}

	return false
}

// SetSpread gets a reference to the given TeaserGroupsBetType and assigns it to the Spread field.
func (o *TeaserGroupsLeague) SetSpread(v TeaserGroupsBetType) {
	o.Spread = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *TeaserGroupsLeague) GetTotal() TeaserGroupsBetType {
	if o == nil || o.Total == nil {
		var ret TeaserGroupsBetType
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TeaserGroupsLeague) GetTotalOk() (TeaserGroupsBetType, bool) {
	if o == nil || o.Total == nil {
		var ret TeaserGroupsBetType
		return ret, false
	}
	return *o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *TeaserGroupsLeague) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given TeaserGroupsBetType and assigns it to the Total field.
func (o *TeaserGroupsLeague) SetTotal(v TeaserGroupsBetType) {
	o.Total = &v
}

type NullableTeaserGroupsLeague struct {
	Value        TeaserGroupsLeague
	ExplicitNull bool
}

func (v NullableTeaserGroupsLeague) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableTeaserGroupsLeague) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
