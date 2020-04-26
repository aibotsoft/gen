# OddsMoneyline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Home** | Pointer to **float64** | Away team price | [optional] 
**Away** | Pointer to **float64** | Away team price. | [optional] 
**Draw** | Pointer to **float64** | Draw price. This is present only for events we offer price for draw. | [optional] 

## Methods

### GetHome

`func (o *OddsMoneyline) GetHome() float64`

GetHome returns the Home field if non-nil, zero value otherwise.

### GetHomeOk

`func (o *OddsMoneyline) GetHomeOk() (float64, bool)`

GetHomeOk returns a tuple with the Home field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHome

`func (o *OddsMoneyline) HasHome() bool`

HasHome returns a boolean if a field has been set.

### SetHome

`func (o *OddsMoneyline) SetHome(v float64)`

SetHome gets a reference to the given float64 and assigns it to the Home field.

### GetAway

`func (o *OddsMoneyline) GetAway() float64`

GetAway returns the Away field if non-nil, zero value otherwise.

### GetAwayOk

`func (o *OddsMoneyline) GetAwayOk() (float64, bool)`

GetAwayOk returns a tuple with the Away field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAway

`func (o *OddsMoneyline) HasAway() bool`

HasAway returns a boolean if a field has been set.

### SetAway

`func (o *OddsMoneyline) SetAway(v float64)`

SetAway gets a reference to the given float64 and assigns it to the Away field.

### GetDraw

`func (o *OddsMoneyline) GetDraw() float64`

GetDraw returns the Draw field if non-nil, zero value otherwise.

### GetDrawOk

`func (o *OddsMoneyline) GetDrawOk() (float64, bool)`

GetDrawOk returns a tuple with the Draw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDraw

`func (o *OddsMoneyline) HasDraw() bool`

HasDraw returns a boolean if a field has been set.

### SetDraw

`func (o *OddsMoneyline) SetDraw(v float64)`

SetDraw gets a reference to the given float64 and assigns it to the Draw field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


