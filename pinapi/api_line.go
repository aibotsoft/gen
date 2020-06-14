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
)

// Linger please
var (
	_ _context.Context
)

// LineApiService LineApi service
type LineApiService service

type apiLineSpecialV1GetRequest struct {
	ctx _context.Context
	apiService *LineApiService
	oddsFormat *string
	specialId *int64
	contestantId *int64
	handicap *int
}


func (r apiLineSpecialV1GetRequest) OddsFormat(oddsFormat string) apiLineSpecialV1GetRequest {
	r.oddsFormat = &oddsFormat
	return r
}

func (r apiLineSpecialV1GetRequest) SpecialId(specialId int64) apiLineSpecialV1GetRequest {
	r.specialId = &specialId
	return r
}

func (r apiLineSpecialV1GetRequest) ContestantId(contestantId int64) apiLineSpecialV1GetRequest {
	r.contestantId = &contestantId
	return r
}

func (r apiLineSpecialV1GetRequest) Handicap(handicap int) apiLineSpecialV1GetRequest {
	r.handicap = &handicap
	return r
}

/*
LineSpecialV1Get Get Special Line - v1
Returns special lines and calculate odds.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return apiLineSpecialV1GetRequest
*/
func (a *LineApiService) LineSpecialV1Get(ctx _context.Context) apiLineSpecialV1GetRequest {
	return apiLineSpecialV1GetRequest{
		apiService: a,
		ctx: ctx,
	}
}

/*
Execute executes the request
 @return SpecialLineResponse
*/
func (r apiLineSpecialV1GetRequest) Execute() (SpecialLineResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SpecialLineResponse
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "LineApiService.LineSpecialV1Get")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/line/special"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	
	if r.oddsFormat == nil {
		return localVarReturnValue, nil, reportError("oddsFormat is required and must be specified")
	}
	
	if r.specialId == nil {
		return localVarReturnValue, nil, reportError("specialId is required and must be specified")
	}
	
	if r.contestantId == nil {
		return localVarReturnValue, nil, reportError("contestantId is required and must be specified")
	}
	
	localVarQueryParams.Add("oddsFormat", parameterToString(*r.oddsFormat, ""))
	localVarQueryParams.Add("specialId", parameterToString(*r.specialId, ""))
	localVarQueryParams.Add("contestantId", parameterToString(*r.contestantId, ""))
	if r.handicap != nil {
		localVarQueryParams.Add("handicap", parameterToString(*r.handicap, ""))
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
type apiLineStraightV1GetRequest struct {
	ctx _context.Context
	apiService *LineApiService
	leagueId *int
	oddsFormat *string
	sportId *int
	eventId *int64
	periodNumber *int
	betType *string
	handicap *float64
	team *string
	side *string
}


func (r apiLineStraightV1GetRequest) LeagueId(leagueId int) apiLineStraightV1GetRequest {
	r.leagueId = &leagueId
	return r
}

func (r apiLineStraightV1GetRequest) OddsFormat(oddsFormat string) apiLineStraightV1GetRequest {
	r.oddsFormat = &oddsFormat
	return r
}

func (r apiLineStraightV1GetRequest) SportId(sportId int) apiLineStraightV1GetRequest {
	r.sportId = &sportId
	return r
}

func (r apiLineStraightV1GetRequest) EventId(eventId int64) apiLineStraightV1GetRequest {
	r.eventId = &eventId
	return r
}

func (r apiLineStraightV1GetRequest) PeriodNumber(periodNumber int) apiLineStraightV1GetRequest {
	r.periodNumber = &periodNumber
	return r
}

func (r apiLineStraightV1GetRequest) BetType(betType string) apiLineStraightV1GetRequest {
	r.betType = &betType
	return r
}

func (r apiLineStraightV1GetRequest) Handicap(handicap float64) apiLineStraightV1GetRequest {
	r.handicap = &handicap
	return r
}

func (r apiLineStraightV1GetRequest) Team(team string) apiLineStraightV1GetRequest {
	r.team = &team
	return r
}

func (r apiLineStraightV1GetRequest) Side(side string) apiLineStraightV1GetRequest {
	r.side = &side
	return r
}

/*
LineStraightV1Get Get Straight Line - v1
Returns latest line.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return apiLineStraightV1GetRequest
*/
func (a *LineApiService) LineStraightV1Get(ctx _context.Context) apiLineStraightV1GetRequest {
	return apiLineStraightV1GetRequest{
		apiService: a,
		ctx: ctx,
	}
}

/*
Execute executes the request
 @return LineResponse
*/
func (r apiLineStraightV1GetRequest) Execute() (LineResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  LineResponse
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "LineApiService.LineStraightV1Get")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/line"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	
	if r.leagueId == nil {
		return localVarReturnValue, nil, reportError("leagueId is required and must be specified")
	}
	
	if r.oddsFormat == nil {
		return localVarReturnValue, nil, reportError("oddsFormat is required and must be specified")
	}
	
	if r.sportId == nil {
		return localVarReturnValue, nil, reportError("sportId is required and must be specified")
	}
	
	if r.eventId == nil {
		return localVarReturnValue, nil, reportError("eventId is required and must be specified")
	}
	
	if r.periodNumber == nil {
		return localVarReturnValue, nil, reportError("periodNumber is required and must be specified")
	}
	
	if r.betType == nil {
		return localVarReturnValue, nil, reportError("betType is required and must be specified")
	}
			
	localVarQueryParams.Add("leagueId", parameterToString(*r.leagueId, ""))
	if r.handicap != nil {
		localVarQueryParams.Add("handicap", parameterToString(*r.handicap, ""))
	}
	localVarQueryParams.Add("oddsFormat", parameterToString(*r.oddsFormat, ""))
	localVarQueryParams.Add("sportId", parameterToString(*r.sportId, ""))
	localVarQueryParams.Add("eventId", parameterToString(*r.eventId, ""))
	localVarQueryParams.Add("periodNumber", parameterToString(*r.periodNumber, ""))
	localVarQueryParams.Add("betType", parameterToString(*r.betType, ""))
	if r.team != nil {
		localVarQueryParams.Add("team", parameterToString(*r.team, ""))
	}
	if r.side != nil {
		localVarQueryParams.Add("side", parameterToString(*r.side, ""))
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