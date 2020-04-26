/*
 * Pinnacle - Customer API Reference
 *
 *  # Authentication   API uses HTTP Basic access authentication. You need to send Authorization HTTP Request header:    `Authorization: Basic <Base64 value of UTF-8 encoded \"username:password\">`  Example:  `Authorization: Basic U03MyOT23YbzMDc6d3c3O1DQ1` 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package pinapi

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"reflect"
)

// Linger please
var (
	_ _context.Context
)

// FixturesApiService FixturesApi service
type FixturesApiService service

type apiFixturesSettledV1GetRequest struct {
	ctx _context.Context
	apiService *FixturesApiService
	sportId *int
	leagueIds *[]int
	since *int
}


func (r apiFixturesSettledV1GetRequest) SportId(sportId int) apiFixturesSettledV1GetRequest {
	r.sportId = &sportId
	return r
}

func (r apiFixturesSettledV1GetRequest) LeagueIds(leagueIds []int) apiFixturesSettledV1GetRequest {
	r.leagueIds = &leagueIds
	return r
}

func (r apiFixturesSettledV1GetRequest) Since(since int) apiFixturesSettledV1GetRequest {
	r.since = &since
	return r
}

/*
FixturesSettledV1Get Get Settled Fixtures - v1
Returns fixtures settled in the last 24 hours for the given sport.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return apiFixturesSettledV1GetRequest
*/
func (a *FixturesApiService) FixturesSettledV1Get(ctx _context.Context) apiFixturesSettledV1GetRequest {
	return apiFixturesSettledV1GetRequest{
		apiService: a,
		ctx: ctx,
	}
}

/*
Execute executes the request
 @return SettledFixturesSport
*/
func (r apiFixturesSettledV1GetRequest) Execute() (SettledFixturesSport, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SettledFixturesSport
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "FixturesApiService.FixturesSettledV1Get")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/fixtures/settled"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	
	if r.sportId == nil {
		return localVarReturnValue, nil, reportError("sportId is required and must be specified")
	}
		
	localVarQueryParams.Add("sportId", parameterToString(*r.sportId, ""))
	if r.leagueIds != nil {
		t := *r.leagueIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("leagueIds", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("leagueIds", parameterToString(t, "multi"))
		}
	}
	if r.since != nil {
		localVarQueryParams.Add("since", parameterToString(*r.since, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v SettledFixturesSport
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ExtendedErrorResponse
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
type apiFixturesSpecialV1GetRequest struct {
	ctx _context.Context
	apiService *FixturesApiService
	sportId *int
	leagueIds *[]int
	since *int64
	category *string
	eventId *int64
	specialId *int64
}


func (r apiFixturesSpecialV1GetRequest) SportId(sportId int) apiFixturesSpecialV1GetRequest {
	r.sportId = &sportId
	return r
}

func (r apiFixturesSpecialV1GetRequest) LeagueIds(leagueIds []int) apiFixturesSpecialV1GetRequest {
	r.leagueIds = &leagueIds
	return r
}

func (r apiFixturesSpecialV1GetRequest) Since(since int64) apiFixturesSpecialV1GetRequest {
	r.since = &since
	return r
}

func (r apiFixturesSpecialV1GetRequest) Category(category string) apiFixturesSpecialV1GetRequest {
	r.category = &category
	return r
}

func (r apiFixturesSpecialV1GetRequest) EventId(eventId int64) apiFixturesSpecialV1GetRequest {
	r.eventId = &eventId
	return r
}

func (r apiFixturesSpecialV1GetRequest) SpecialId(specialId int64) apiFixturesSpecialV1GetRequest {
	r.specialId = &specialId
	return r
}

/*
FixturesSpecialV1Get Get Special Fixtures - v1
Returns all **non-settled** specials for the given sport.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return apiFixturesSpecialV1GetRequest
*/
func (a *FixturesApiService) FixturesSpecialV1Get(ctx _context.Context) apiFixturesSpecialV1GetRequest {
	return apiFixturesSpecialV1GetRequest{
		apiService: a,
		ctx: ctx,
	}
}

/*
Execute executes the request
 @return SpecialsFixturesResponse
*/
func (r apiFixturesSpecialV1GetRequest) Execute() (SpecialsFixturesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SpecialsFixturesResponse
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "FixturesApiService.FixturesSpecialV1Get")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/fixtures/special"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	
	if r.sportId == nil {
		return localVarReturnValue, nil, reportError("sportId is required and must be specified")
	}
					
	localVarQueryParams.Add("sportId", parameterToString(*r.sportId, ""))
	if r.leagueIds != nil {
		t := *r.leagueIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("leagueIds", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("leagueIds", parameterToString(t, "multi"))
		}
	}
	if r.since != nil {
		localVarQueryParams.Add("since", parameterToString(*r.since, ""))
	}
	if r.category != nil {
		localVarQueryParams.Add("category", parameterToString(*r.category, ""))
	}
	if r.eventId != nil {
		localVarQueryParams.Add("eventId", parameterToString(*r.eventId, ""))
	}
	if r.specialId != nil {
		localVarQueryParams.Add("specialId", parameterToString(*r.specialId, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v SpecialsFixturesResponse
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ExtendedErrorResponse
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
type apiFixturesSpecialsSettledV1GetRequest struct {
	ctx _context.Context
	apiService *FixturesApiService
	sportId *int
	leagueIds *[]int
	since *int64
}


func (r apiFixturesSpecialsSettledV1GetRequest) SportId(sportId int) apiFixturesSpecialsSettledV1GetRequest {
	r.sportId = &sportId
	return r
}

func (r apiFixturesSpecialsSettledV1GetRequest) LeagueIds(leagueIds []int) apiFixturesSpecialsSettledV1GetRequest {
	r.leagueIds = &leagueIds
	return r
}

func (r apiFixturesSpecialsSettledV1GetRequest) Since(since int64) apiFixturesSpecialsSettledV1GetRequest {
	r.since = &since
	return r
}

/*
FixturesSpecialsSettledV1Get Get Settled Special Fixtures - v1
Returns all specials which are settled in the last 24 hours for the given Sport.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return apiFixturesSpecialsSettledV1GetRequest
*/
func (a *FixturesApiService) FixturesSpecialsSettledV1Get(ctx _context.Context) apiFixturesSpecialsSettledV1GetRequest {
	return apiFixturesSpecialsSettledV1GetRequest{
		apiService: a,
		ctx: ctx,
	}
}

/*
Execute executes the request
 @return SettledSpecialsResponse
*/
func (r apiFixturesSpecialsSettledV1GetRequest) Execute() (SettledSpecialsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SettledSpecialsResponse
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "FixturesApiService.FixturesSpecialsSettledV1Get")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/fixtures/special/settled"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	
	if r.sportId == nil {
		return localVarReturnValue, nil, reportError("sportId is required and must be specified")
	}
		
	localVarQueryParams.Add("sportId", parameterToString(*r.sportId, ""))
	if r.leagueIds != nil {
		t := *r.leagueIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("leagueIds", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("leagueIds", parameterToString(t, "multi"))
		}
	}
	if r.since != nil {
		localVarQueryParams.Add("since", parameterToString(*r.since, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v SettledSpecialsResponse
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ExtendedErrorResponse
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
type apiFixturesV1GetRequest struct {
	ctx _context.Context
	apiService *FixturesApiService
	sportId *int
	leagueIds *[]int
	isLive *bool
	since *int64
	eventIds *[]int
}


func (r apiFixturesV1GetRequest) SportId(sportId int) apiFixturesV1GetRequest {
	r.sportId = &sportId
	return r
}

func (r apiFixturesV1GetRequest) LeagueIds(leagueIds []int) apiFixturesV1GetRequest {
	r.leagueIds = &leagueIds
	return r
}

func (r apiFixturesV1GetRequest) IsLive(isLive bool) apiFixturesV1GetRequest {
	r.isLive = &isLive
	return r
}

func (r apiFixturesV1GetRequest) Since(since int64) apiFixturesV1GetRequest {
	r.since = &since
	return r
}

func (r apiFixturesV1GetRequest) EventIds(eventIds []int) apiFixturesV1GetRequest {
	r.eventIds = &eventIds
	return r
}

/*
FixturesV1Get Get Fixtures - v1
Returns all **non-settled** events for the given sport. Please note that it is possible that the event is in Get Fixtures response but not in Get Odds. This happens when the odds are not currently available for wagering. Please note that it is possible to receive the same exact response when using **since** parameter. This is rare and can be caused by internal updates of event properties.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return apiFixturesV1GetRequest
*/
func (a *FixturesApiService) FixturesV1Get(ctx _context.Context) apiFixturesV1GetRequest {
	return apiFixturesV1GetRequest{
		apiService: a,
		ctx: ctx,
	}
}

/*
Execute executes the request
 @return FixturesResponse
*/
func (r apiFixturesV1GetRequest) Execute() (FixturesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  FixturesResponse
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "FixturesApiService.FixturesV1Get")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/fixtures"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	
	if r.sportId == nil {
		return localVarReturnValue, nil, reportError("sportId is required and must be specified")
	}
				
	localVarQueryParams.Add("sportId", parameterToString(*r.sportId, ""))
	if r.leagueIds != nil {
		t := *r.leagueIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("leagueIds", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("leagueIds", parameterToString(t, "multi"))
		}
	}
	if r.isLive != nil {
		localVarQueryParams.Add("isLive", parameterToString(*r.isLive, ""))
	}
	if r.since != nil {
		localVarQueryParams.Add("since", parameterToString(*r.since, ""))
	}
	if r.eventIds != nil {
		t := *r.eventIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("eventIds", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("eventIds", parameterToString(t, "multi"))
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v FixturesResponse
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ExtendedErrorResponse
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
