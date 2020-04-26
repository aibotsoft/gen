# ClientBalanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableBalance** | Pointer to **float64** | Amount available for betting. | 
**OutstandingTransactions** | Pointer to **float64** | Sum of not yet settled bet amounts. | 
**GivenCredit** | Pointer to **float64** | Client’s credit. | 
**Currency** | Pointer to **string** | Client’s currency code. | 

## Methods

### GetAvailableBalance

`func (o *ClientBalanceResponse) GetAvailableBalance() float64`

GetAvailableBalance returns the AvailableBalance field if non-nil, zero value otherwise.

### GetAvailableBalanceOk

`func (o *ClientBalanceResponse) GetAvailableBalanceOk() (float64, bool)`

GetAvailableBalanceOk returns a tuple with the AvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAvailableBalance

`func (o *ClientBalanceResponse) HasAvailableBalance() bool`

HasAvailableBalance returns a boolean if a field has been set.

### SetAvailableBalance

`func (o *ClientBalanceResponse) SetAvailableBalance(v float64)`

SetAvailableBalance gets a reference to the given float64 and assigns it to the AvailableBalance field.

### GetOutstandingTransactions

`func (o *ClientBalanceResponse) GetOutstandingTransactions() float64`

GetOutstandingTransactions returns the OutstandingTransactions field if non-nil, zero value otherwise.

### GetOutstandingTransactionsOk

`func (o *ClientBalanceResponse) GetOutstandingTransactionsOk() (float64, bool)`

GetOutstandingTransactionsOk returns a tuple with the OutstandingTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOutstandingTransactions

`func (o *ClientBalanceResponse) HasOutstandingTransactions() bool`

HasOutstandingTransactions returns a boolean if a field has been set.

### SetOutstandingTransactions

`func (o *ClientBalanceResponse) SetOutstandingTransactions(v float64)`

SetOutstandingTransactions gets a reference to the given float64 and assigns it to the OutstandingTransactions field.

### GetGivenCredit

`func (o *ClientBalanceResponse) GetGivenCredit() float64`

GetGivenCredit returns the GivenCredit field if non-nil, zero value otherwise.

### GetGivenCreditOk

`func (o *ClientBalanceResponse) GetGivenCreditOk() (float64, bool)`

GetGivenCreditOk returns a tuple with the GivenCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGivenCredit

`func (o *ClientBalanceResponse) HasGivenCredit() bool`

HasGivenCredit returns a boolean if a field has been set.

### SetGivenCredit

`func (o *ClientBalanceResponse) SetGivenCredit(v float64)`

SetGivenCredit gets a reference to the given float64 and assigns it to the GivenCredit field.

### GetCurrency

`func (o *ClientBalanceResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ClientBalanceResponse) GetCurrencyOk() (string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCurrency

`func (o *ClientBalanceResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrency

`func (o *ClientBalanceResponse) SetCurrency(v string)`

SetCurrency gets a reference to the given string and assigns it to the Currency field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


