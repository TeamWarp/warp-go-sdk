// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package warphr

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/TeamWarp/warp-go-sdk/internal/apijson"
	"github.com/TeamWarp/warp-go-sdk/internal/apiquery"
	"github.com/TeamWarp/warp-go-sdk/internal/requestconfig"
	"github.com/TeamWarp/warp-go-sdk/option"
	"github.com/TeamWarp/warp-go-sdk/packages/pagination"
	"github.com/TeamWarp/warp-go-sdk/packages/param"
	"github.com/TeamWarp/warp-go-sdk/packages/respjson"
)

// Endpoints for workplace management. Create, list, and update workplaces within
// your company.
//
// WorkplaceService contains methods and other services that help with interacting
// with the warp-hr API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkplaceService] method instead.
type WorkplaceService struct {
	options []option.RequestOption
}

// NewWorkplaceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkplaceService(opts ...option.RequestOption) (r WorkplaceService) {
	r = WorkplaceService{}
	r.options = opts
	return
}

// Create a new workplace.
func (r *WorkplaceService) New(ctx context.Context, body WorkplaceNewParams, opts ...option.RequestOption) (res *WorkplaceNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/workplaces"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update an existing workplace.
func (r *WorkplaceService) Update(ctx context.Context, id string, body WorkplaceUpdateParams, opts ...option.RequestOption) (res *WorkplaceUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workplaces/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List all workplaces for your company.
func (r *WorkplaceService) List(ctx context.Context, query WorkplaceListParams, opts ...option.RequestOption) (res *pagination.CursorPage[WorkplaceListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/workplaces"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List all workplaces for your company.
func (r *WorkplaceService) ListAutoPaging(ctx context.Context, query WorkplaceListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[WorkplaceListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

type WorkplaceNewResponse struct {
	// Public workplace identifier
	ID string `json:"id" api:"required"`
	// A valid US address
	Address WorkplaceNewResponseAddress `json:"address" api:"required"`
	// a string to be decoded into a Date
	CreatedAt string `json:"createdAt" api:"required"`
	Name      string `json:"name" api:"required"`
	// Any of "active", "archived".
	Status WorkplaceNewResponseStatus `json:"status" api:"required"`
	// Any of "remote", "office".
	Type WorkplaceNewResponseType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Address     respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkplaceNewResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkplaceNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A valid US address
type WorkplaceNewResponseAddress struct {
	City string `json:"city" api:"required"`
	// Any of "US".
	Country string `json:"country" api:"required"`
	// a non empty string
	Line1      string `json:"line1" api:"required"`
	PostalCode string `json:"postalCode" api:"required"`
	// Any of "AL", "AK", "AZ", "AR", "CA", "CO", "CT", "DC", "DE", "FL", "GA", "HI",
	// "ID", "IL", "IN", "IA", "KS", "KY", "LA", "ME", "MD", "MA", "MI", "MN", "MS",
	// "MO", "MT", "NE", "NV", "NH", "NJ", "NM", "NY", "NC", "ND", "OH", "OK", "OR",
	// "PA", "RI", "SC", "SD", "TN", "TX", "UT", "VT", "VA", "WA", "WV", "WI", "WY".
	State string `json:"state" api:"required"`
	Line2 string `json:"line2" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Country     respjson.Field
		Line1       respjson.Field
		PostalCode  respjson.Field
		State       respjson.Field
		Line2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkplaceNewResponseAddress) RawJSON() string { return r.JSON.raw }
func (r *WorkplaceNewResponseAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkplaceNewResponseStatus string

const (
	WorkplaceNewResponseStatusActive   WorkplaceNewResponseStatus = "active"
	WorkplaceNewResponseStatusArchived WorkplaceNewResponseStatus = "archived"
)

type WorkplaceNewResponseType string

const (
	WorkplaceNewResponseTypeRemote WorkplaceNewResponseType = "remote"
	WorkplaceNewResponseTypeOffice WorkplaceNewResponseType = "office"
)

type WorkplaceUpdateResponse struct {
	// Public workplace identifier
	ID string `json:"id" api:"required"`
	// A valid US address
	Address WorkplaceUpdateResponseAddress `json:"address" api:"required"`
	// a string to be decoded into a Date
	CreatedAt string `json:"createdAt" api:"required"`
	Name      string `json:"name" api:"required"`
	// Any of "active", "archived".
	Status WorkplaceUpdateResponseStatus `json:"status" api:"required"`
	// Any of "remote", "office".
	Type WorkplaceUpdateResponseType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Address     respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkplaceUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkplaceUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A valid US address
type WorkplaceUpdateResponseAddress struct {
	City string `json:"city" api:"required"`
	// Any of "US".
	Country string `json:"country" api:"required"`
	// a non empty string
	Line1      string `json:"line1" api:"required"`
	PostalCode string `json:"postalCode" api:"required"`
	// Any of "AL", "AK", "AZ", "AR", "CA", "CO", "CT", "DC", "DE", "FL", "GA", "HI",
	// "ID", "IL", "IN", "IA", "KS", "KY", "LA", "ME", "MD", "MA", "MI", "MN", "MS",
	// "MO", "MT", "NE", "NV", "NH", "NJ", "NM", "NY", "NC", "ND", "OH", "OK", "OR",
	// "PA", "RI", "SC", "SD", "TN", "TX", "UT", "VT", "VA", "WA", "WV", "WI", "WY".
	State string `json:"state" api:"required"`
	Line2 string `json:"line2" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Country     respjson.Field
		Line1       respjson.Field
		PostalCode  respjson.Field
		State       respjson.Field
		Line2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkplaceUpdateResponseAddress) RawJSON() string { return r.JSON.raw }
func (r *WorkplaceUpdateResponseAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkplaceUpdateResponseStatus string

const (
	WorkplaceUpdateResponseStatusActive   WorkplaceUpdateResponseStatus = "active"
	WorkplaceUpdateResponseStatusArchived WorkplaceUpdateResponseStatus = "archived"
)

type WorkplaceUpdateResponseType string

const (
	WorkplaceUpdateResponseTypeRemote WorkplaceUpdateResponseType = "remote"
	WorkplaceUpdateResponseTypeOffice WorkplaceUpdateResponseType = "office"
)

type WorkplaceListResponse struct {
	// Public workplace identifier
	ID string `json:"id" api:"required"`
	// A valid US address
	Address WorkplaceListResponseAddress `json:"address" api:"required"`
	// a string to be decoded into a Date
	CreatedAt string `json:"createdAt" api:"required"`
	Name      string `json:"name" api:"required"`
	// Any of "active", "archived".
	Status WorkplaceListResponseStatus `json:"status" api:"required"`
	// Any of "remote", "office".
	Type WorkplaceListResponseType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Address     respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkplaceListResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkplaceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A valid US address
type WorkplaceListResponseAddress struct {
	City string `json:"city" api:"required"`
	// Any of "US".
	Country string `json:"country" api:"required"`
	// a non empty string
	Line1      string `json:"line1" api:"required"`
	PostalCode string `json:"postalCode" api:"required"`
	// Any of "AL", "AK", "AZ", "AR", "CA", "CO", "CT", "DC", "DE", "FL", "GA", "HI",
	// "ID", "IL", "IN", "IA", "KS", "KY", "LA", "ME", "MD", "MA", "MI", "MN", "MS",
	// "MO", "MT", "NE", "NV", "NH", "NJ", "NM", "NY", "NC", "ND", "OH", "OK", "OR",
	// "PA", "RI", "SC", "SD", "TN", "TX", "UT", "VT", "VA", "WA", "WV", "WI", "WY".
	State string `json:"state" api:"required"`
	Line2 string `json:"line2" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Country     respjson.Field
		Line1       respjson.Field
		PostalCode  respjson.Field
		State       respjson.Field
		Line2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkplaceListResponseAddress) RawJSON() string { return r.JSON.raw }
func (r *WorkplaceListResponseAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkplaceListResponseStatus string

const (
	WorkplaceListResponseStatusActive   WorkplaceListResponseStatus = "active"
	WorkplaceListResponseStatusArchived WorkplaceListResponseStatus = "archived"
)

type WorkplaceListResponseType string

const (
	WorkplaceListResponseTypeRemote WorkplaceListResponseType = "remote"
	WorkplaceListResponseTypeOffice WorkplaceListResponseType = "office"
)

type WorkplaceNewParams struct {
	// A valid US address
	Address WorkplaceNewParamsAddress `json:"address,omitzero" api:"required"`
	// a non empty string
	Name string `json:"name" api:"required"`
	// Any of "remote", "office".
	Type WorkplaceNewParamsType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r WorkplaceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkplaceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkplaceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A valid US address
//
// The properties City, Country, Line1, PostalCode, State are required.
type WorkplaceNewParamsAddress struct {
	City string `json:"city" api:"required"`
	// Any of "US".
	Country string `json:"country,omitzero" api:"required"`
	// a non empty string
	Line1      string `json:"line1" api:"required"`
	PostalCode string `json:"postalCode" api:"required"`
	// Any of "AL", "AK", "AZ", "AR", "CA", "CO", "CT", "DC", "DE", "FL", "GA", "HI",
	// "ID", "IL", "IN", "IA", "KS", "KY", "LA", "ME", "MD", "MA", "MI", "MN", "MS",
	// "MO", "MT", "NE", "NV", "NH", "NJ", "NM", "NY", "NC", "ND", "OH", "OK", "OR",
	// "PA", "RI", "SC", "SD", "TN", "TX", "UT", "VT", "VA", "WA", "WV", "WI", "WY".
	State string            `json:"state,omitzero" api:"required"`
	Line2 param.Opt[string] `json:"line2,omitzero"`
	paramObj
}

func (r WorkplaceNewParamsAddress) MarshalJSON() (data []byte, err error) {
	type shadow WorkplaceNewParamsAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkplaceNewParamsAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkplaceNewParamsAddress](
		"country", "US",
	)
	apijson.RegisterFieldValidator[WorkplaceNewParamsAddress](
		"state", "AL", "AK", "AZ", "AR", "CA", "CO", "CT", "DC", "DE", "FL", "GA", "HI", "ID", "IL", "IN", "IA", "KS", "KY", "LA", "ME", "MD", "MA", "MI", "MN", "MS", "MO", "MT", "NE", "NV", "NH", "NJ", "NM", "NY", "NC", "ND", "OH", "OK", "OR", "PA", "RI", "SC", "SD", "TN", "TX", "UT", "VT", "VA", "WA", "WV", "WI", "WY",
	)
}

type WorkplaceNewParamsType string

const (
	WorkplaceNewParamsTypeRemote WorkplaceNewParamsType = "remote"
	WorkplaceNewParamsTypeOffice WorkplaceNewParamsType = "office"
)

type WorkplaceUpdateParams struct {
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r WorkplaceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkplaceUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkplaceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkplaceListParams struct {
	// Public workplace identifier
	AfterID param.Opt[string] `query:"afterId,omitzero" json:"-"`
	// Public workplace identifier
	BeforeID param.Opt[string] `query:"beforeId,omitzero" json:"-"`
	// a number less than or equal to 100
	Limit param.Opt[string] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkplaceListParams]'s query parameters as `url.Values`.
func (r WorkplaceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
