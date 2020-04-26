# OddsPeriod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineId** | Pointer to **int64** | Line Id. | [optional] 
**Number** | Pointer to **int32** | This represents the period of the match. For example, for soccer we have  0 (Game), 1 (1st Half) &amp; 2 (2nd Half) | [optional] 
**Cutoff** | Pointer to [**time.Time**](time.Time.md) | Period’s wagering cut-off date in UTC. | [optional] 
**Status** | Pointer to **int32** | 1 - online, period is open for betting  2 - offline, period is not open for betting  | [optional] 
**MaxSpread** | Pointer to **float64** | Maximum spread bet volume. See [How to calculate max risk from the max volume](https://github.com/pinnacleapi/pinnacleapi-documentation/blob/master/FAQ.md#how-to-calculate-max-risk-from-the-max-volume-limits-in-odds) | [optional] 
**MaxMoneyline** | Pointer to **float64** | Maximum moneyline bet volume. See [How to calculate max risk from the max volume](https://github.com/pinnacleapi/pinnacleapi-documentation/blob/master/FAQ.md#how-to-calculate-max-risk-from-the-max-volume-limits-in-odds) | [optional] 
**MaxTotal** | Pointer to **float64** | Maximum total points bet volume. See [How to calculate max risk from the max volume](https://github.com/pinnacleapi/pinnacleapi-documentation/blob/master/FAQ.md#how-to-calculate-max-risk-from-the-max-volume-limits-in-odds) | [optional] 
**MaxTeamTotal** | Pointer to **float64** | Maximum team total points bet volume. See [How to calculate max risk from the max volume](https://github.com/pinnacleapi/pinnacleapi-documentation/blob/master/FAQ.md#how-to-calculate-max-risk-from-the-max-volume-limits-in-odds) | [optional] 
**Spreads** | Pointer to [**[]OddsSpread**](OddsSpread.md) | Container for spread odds. | [optional] 
**Moneyline** | Pointer to [**OddsMoneyline**](OddsMoneyline.md) |  | [optional] 
**Totals** | Pointer to [**[]OddsTotal**](OddsTotal.md) | Container for team total points. | [optional] 
**TeamTotal** | Pointer to [**OddsTeamTotals**](OddsTeamTotals.md) |  | [optional] 

## Methods

### GetLineId

`func (o *OddsPeriod) GetLineId() int64`

GetLineId returns the LineId field if non-nil, zero value otherwise.

### GetLineIdOk

`func (o *OddsPeriod) GetLineIdOk() (int64, bool)`

GetLineIdOk returns a tuple with the LineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLineId

`func (o *OddsPeriod) HasLineId() bool`

HasLineId returns a boolean if a field has been set.

### SetLineId

`func (o *OddsPeriod) SetLineId(v int64)`

SetLineId gets a reference to the given int64 and assigns it to the LineId field.

### GetNumber

`func (o *OddsPeriod) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *OddsPeriod) GetNumberOk() (int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNumber

`func (o *OddsPeriod) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumber

`func (o *OddsPeriod) SetNumber(v int32)`

SetNumber gets a reference to the given int32 and assigns it to the Number field.

### GetCutoff

`func (o *OddsPeriod) GetCutoff() time.Time`

GetCutoff returns the Cutoff field if non-nil, zero value otherwise.

### GetCutoffOk

`func (o *OddsPeriod) GetCutoffOk() (time.Time, bool)`

GetCutoffOk returns a tuple with the Cutoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCutoff

`func (o *OddsPeriod) HasCutoff() bool`

HasCutoff returns a boolean if a field has been set.

### SetCutoff

`func (o *OddsPeriod) SetCutoff(v time.Time)`

SetCutoff gets a reference to the given time.Time and assigns it to the Cutoff field.

### GetStatus

`func (o *OddsPeriod) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OddsPeriod) GetStatusOk() (int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *OddsPeriod) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *OddsPeriod) SetStatus(v int32)`

SetStatus gets a reference to the given int32 and assigns it to the Status field.

### GetMaxSpread

`func (o *OddsPeriod) GetMaxSpread() float64`

GetMaxSpread returns the MaxSpread field if non-nil, zero value otherwise.

### GetMaxSpreadOk

`func (o *OddsPeriod) GetMaxSpreadOk() (float64, bool)`

GetMaxSpreadOk returns a tuple with the MaxSpread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaxSpread

`func (o *OddsPeriod) HasMaxSpread() bool`

HasMaxSpread returns a boolean if a field has been set.

### SetMaxSpread

`func (o *OddsPeriod) SetMaxSpread(v float64)`

SetMaxSpread gets a reference to the given float64 and assigns it to the MaxSpread field.

### GetMaxMoneyline

`func (o *OddsPeriod) GetMaxMoneyline() float64`

GetMaxMoneyline returns the MaxMoneyline field if non-nil, zero value otherwise.

### GetMaxMoneylineOk

`func (o *OddsPeriod) GetMaxMoneylineOk() (float64, bool)`

GetMaxMoneylineOk returns a tuple with the MaxMoneyline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaxMoneyline

`func (o *OddsPeriod) HasMaxMoneyline() bool`

HasMaxMoneyline returns a boolean if a field has been set.

### SetMaxMoneyline

`func (o *OddsPeriod) SetMaxMoneyline(v float64)`

SetMaxMoneyline gets a reference to the given float64 and assigns it to the MaxMoneyline field.

### GetMaxTotal

`func (o *OddsPeriod) GetMaxTotal() float64`

GetMaxTotal returns the MaxTotal field if non-nil, zero value otherwise.

### GetMaxTotalOk

`func (o *OddsPeriod) GetMaxTotalOk() (float64, bool)`

GetMaxTotalOk returns a tuple with the MaxTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaxTotal

`func (o *OddsPeriod) HasMaxTotal() bool`

HasMaxTotal returns a boolean if a field has been set.

### SetMaxTotal

`func (o *OddsPeriod) SetMaxTotal(v float64)`

SetMaxTotal gets a reference to the given float64 and assigns it to the MaxTotal field.

### GetMaxTeamTotal

`func (o *OddsPeriod) GetMaxTeamTotal() float64`

GetMaxTeamTotal returns the MaxTeamTotal field if non-nil, zero value otherwise.

### GetMaxTeamTotalOk

`func (o *OddsPeriod) GetMaxTeamTotalOk() (float64, bool)`

GetMaxTeamTotalOk returns a tuple with the MaxTeamTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaxTeamTotal

`func (o *OddsPeriod) HasMaxTeamTotal() bool`

HasMaxTeamTotal returns a boolean if a field has been set.

### SetMaxTeamTotal

`func (o *OddsPeriod) SetMaxTeamTotal(v float64)`

SetMaxTeamTotal gets a reference to the given float64 and assigns it to the MaxTeamTotal field.

### GetSpreads

`func (o *OddsPeriod) GetSpreads() []OddsSpread`

GetSpreads returns the Spreads field if non-nil, zero value otherwise.

### GetSpreadsOk

`func (o *OddsPeriod) GetSpreadsOk() ([]OddsSpread, bool)`

GetSpreadsOk returns a tuple with the Spreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSpreads

`func (o *OddsPeriod) HasSpreads() bool`

HasSpreads returns a boolean if a field has been set.

### SetSpreads

`func (o *OddsPeriod) SetSpreads(v []OddsSpread)`

SetSpreads gets a reference to the given []OddsSpread and assigns it to the Spreads field.

### GetMoneyline

`func (o *OddsPeriod) GetMoneyline() OddsMoneyline`

GetMoneyline returns the Moneyline field if non-nil, zero value otherwise.

### GetMoneylineOk

`func (o *OddsPeriod) GetMoneylineOk() (OddsMoneyline, bool)`

GetMoneylineOk returns a tuple with the Moneyline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMoneyline

`func (o *OddsPeriod) HasMoneyline() bool`

HasMoneyline returns a boolean if a field has been set.

### SetMoneyline

`func (o *OddsPeriod) SetMoneyline(v OddsMoneyline)`

SetMoneyline gets a reference to the given OddsMoneyline and assigns it to the Moneyline field.

### GetTotals

`func (o *OddsPeriod) GetTotals() []OddsTotal`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *OddsPeriod) GetTotalsOk() ([]OddsTotal, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotals

`func (o *OddsPeriod) HasTotals() bool`

HasTotals returns a boolean if a field has been set.

### SetTotals

`func (o *OddsPeriod) SetTotals(v []OddsTotal)`

SetTotals gets a reference to the given []OddsTotal and assigns it to the Totals field.

### GetTeamTotal

`func (o *OddsPeriod) GetTeamTotal() OddsTeamTotals`

GetTeamTotal returns the TeamTotal field if non-nil, zero value otherwise.

### GetTeamTotalOk

`func (o *OddsPeriod) GetTeamTotalOk() (OddsTeamTotals, bool)`

GetTeamTotalOk returns a tuple with the TeamTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTeamTotal

`func (o *OddsPeriod) HasTeamTotal() bool`

HasTeamTotal returns a boolean if a field has been set.

### SetTeamTotal

`func (o *OddsPeriod) SetTeamTotal(v OddsTeamTotals)`

SetTeamTotal gets a reference to the given OddsTeamTotals and assigns it to the TeamTotal field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


