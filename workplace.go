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
	"github.com/TeamWarp/warp-go-sdk/internal/param"
	"github.com/TeamWarp/warp-go-sdk/internal/requestconfig"
	"github.com/TeamWarp/warp-go-sdk/option"
	"github.com/TeamWarp/warp-go-sdk/packages/pagination"
)

// Endpoints for workplace management. Create, list, and update workplaces within
// your company.
//
// WorkplaceService contains methods and other services that help with interacting
// with the warp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkplaceService] method instead.
type WorkplaceService struct {
	Options []option.RequestOption
}

// NewWorkplaceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkplaceService(opts ...option.RequestOption) (r *WorkplaceService) {
	r = &WorkplaceService{}
	r.Options = opts
	return
}

// Create a new workplace.
func (r *WorkplaceService) New(ctx context.Context, body WorkplaceNewParams, opts ...option.RequestOption) (res *WorkplaceNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/workplaces"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update an existing workplace.
func (r *WorkplaceService) Update(ctx context.Context, id string, body WorkplaceUpdateParams, opts ...option.RequestOption) (res *WorkplaceUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workplaces/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List all workplaces for your company.
func (r *WorkplaceService) List(ctx context.Context, query WorkplaceListParams, opts ...option.RequestOption) (res *pagination.CursorPage[WorkplaceListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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
	CreatedAt string                     `json:"createdAt" api:"required"`
	Name      string                     `json:"name" api:"required"`
	Status    WorkplaceNewResponseStatus `json:"status" api:"required"`
	Type      WorkplaceNewResponseType   `json:"type" api:"required"`
	JSON      workplaceNewResponseJSON   `json:"-"`
}

// workplaceNewResponseJSON contains the JSON metadata for the struct
// [WorkplaceNewResponse]
type workplaceNewResponseJSON struct {
	ID          apijson.Field
	Address     apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkplaceNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workplaceNewResponseJSON) RawJSON() string {
	return r.raw
}

// A valid US address
type WorkplaceNewResponseAddress struct {
	City    string                             `json:"city" api:"required"`
	Country WorkplaceNewResponseAddressCountry `json:"country" api:"required"`
	// a non empty string
	Line1      string                           `json:"line1" api:"required"`
	PostalCode string                           `json:"postalCode" api:"required"`
	State      WorkplaceNewResponseAddressState `json:"state" api:"required"`
	Line2      string                           `json:"line2" api:"nullable"`
	JSON       workplaceNewResponseAddressJSON  `json:"-"`
}

// workplaceNewResponseAddressJSON contains the JSON metadata for the struct
// [WorkplaceNewResponseAddress]
type workplaceNewResponseAddressJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Line1       apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	Line2       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkplaceNewResponseAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workplaceNewResponseAddressJSON) RawJSON() string {
	return r.raw
}

type WorkplaceNewResponseAddressCountry string

const (
	WorkplaceNewResponseAddressCountryUs WorkplaceNewResponseAddressCountry = "US"
)

func (r WorkplaceNewResponseAddressCountry) IsKnown() bool {
	switch r {
	case WorkplaceNewResponseAddressCountryUs:
		return true
	}
	return false
}

type WorkplaceNewResponseAddressState string

const (
	WorkplaceNewResponseAddressStateAl WorkplaceNewResponseAddressState = "AL"
	WorkplaceNewResponseAddressStateAk WorkplaceNewResponseAddressState = "AK"
	WorkplaceNewResponseAddressStateAz WorkplaceNewResponseAddressState = "AZ"
	WorkplaceNewResponseAddressStateAr WorkplaceNewResponseAddressState = "AR"
	WorkplaceNewResponseAddressStateCa WorkplaceNewResponseAddressState = "CA"
	WorkplaceNewResponseAddressStateCo WorkplaceNewResponseAddressState = "CO"
	WorkplaceNewResponseAddressStateCt WorkplaceNewResponseAddressState = "CT"
	WorkplaceNewResponseAddressStateDc WorkplaceNewResponseAddressState = "DC"
	WorkplaceNewResponseAddressStateDe WorkplaceNewResponseAddressState = "DE"
	WorkplaceNewResponseAddressStateFl WorkplaceNewResponseAddressState = "FL"
	WorkplaceNewResponseAddressStateGa WorkplaceNewResponseAddressState = "GA"
	WorkplaceNewResponseAddressStateHi WorkplaceNewResponseAddressState = "HI"
	WorkplaceNewResponseAddressStateID WorkplaceNewResponseAddressState = "ID"
	WorkplaceNewResponseAddressStateIl WorkplaceNewResponseAddressState = "IL"
	WorkplaceNewResponseAddressStateIn WorkplaceNewResponseAddressState = "IN"
	WorkplaceNewResponseAddressStateIa WorkplaceNewResponseAddressState = "IA"
	WorkplaceNewResponseAddressStateKs WorkplaceNewResponseAddressState = "KS"
	WorkplaceNewResponseAddressStateKy WorkplaceNewResponseAddressState = "KY"
	WorkplaceNewResponseAddressStateLa WorkplaceNewResponseAddressState = "LA"
	WorkplaceNewResponseAddressStateMe WorkplaceNewResponseAddressState = "ME"
	WorkplaceNewResponseAddressStateMd WorkplaceNewResponseAddressState = "MD"
	WorkplaceNewResponseAddressStateMa WorkplaceNewResponseAddressState = "MA"
	WorkplaceNewResponseAddressStateMi WorkplaceNewResponseAddressState = "MI"
	WorkplaceNewResponseAddressStateMn WorkplaceNewResponseAddressState = "MN"
	WorkplaceNewResponseAddressStateMs WorkplaceNewResponseAddressState = "MS"
	WorkplaceNewResponseAddressStateMo WorkplaceNewResponseAddressState = "MO"
	WorkplaceNewResponseAddressStateMt WorkplaceNewResponseAddressState = "MT"
	WorkplaceNewResponseAddressStateNe WorkplaceNewResponseAddressState = "NE"
	WorkplaceNewResponseAddressStateNv WorkplaceNewResponseAddressState = "NV"
	WorkplaceNewResponseAddressStateNh WorkplaceNewResponseAddressState = "NH"
	WorkplaceNewResponseAddressStateNj WorkplaceNewResponseAddressState = "NJ"
	WorkplaceNewResponseAddressStateNm WorkplaceNewResponseAddressState = "NM"
	WorkplaceNewResponseAddressStateNy WorkplaceNewResponseAddressState = "NY"
	WorkplaceNewResponseAddressStateNc WorkplaceNewResponseAddressState = "NC"
	WorkplaceNewResponseAddressStateNd WorkplaceNewResponseAddressState = "ND"
	WorkplaceNewResponseAddressStateOh WorkplaceNewResponseAddressState = "OH"
	WorkplaceNewResponseAddressStateOk WorkplaceNewResponseAddressState = "OK"
	WorkplaceNewResponseAddressStateOr WorkplaceNewResponseAddressState = "OR"
	WorkplaceNewResponseAddressStatePa WorkplaceNewResponseAddressState = "PA"
	WorkplaceNewResponseAddressStateRi WorkplaceNewResponseAddressState = "RI"
	WorkplaceNewResponseAddressStateSc WorkplaceNewResponseAddressState = "SC"
	WorkplaceNewResponseAddressStateSd WorkplaceNewResponseAddressState = "SD"
	WorkplaceNewResponseAddressStateTn WorkplaceNewResponseAddressState = "TN"
	WorkplaceNewResponseAddressStateTx WorkplaceNewResponseAddressState = "TX"
	WorkplaceNewResponseAddressStateUt WorkplaceNewResponseAddressState = "UT"
	WorkplaceNewResponseAddressStateVt WorkplaceNewResponseAddressState = "VT"
	WorkplaceNewResponseAddressStateVa WorkplaceNewResponseAddressState = "VA"
	WorkplaceNewResponseAddressStateWa WorkplaceNewResponseAddressState = "WA"
	WorkplaceNewResponseAddressStateWv WorkplaceNewResponseAddressState = "WV"
	WorkplaceNewResponseAddressStateWi WorkplaceNewResponseAddressState = "WI"
	WorkplaceNewResponseAddressStateWy WorkplaceNewResponseAddressState = "WY"
)

func (r WorkplaceNewResponseAddressState) IsKnown() bool {
	switch r {
	case WorkplaceNewResponseAddressStateAl, WorkplaceNewResponseAddressStateAk, WorkplaceNewResponseAddressStateAz, WorkplaceNewResponseAddressStateAr, WorkplaceNewResponseAddressStateCa, WorkplaceNewResponseAddressStateCo, WorkplaceNewResponseAddressStateCt, WorkplaceNewResponseAddressStateDc, WorkplaceNewResponseAddressStateDe, WorkplaceNewResponseAddressStateFl, WorkplaceNewResponseAddressStateGa, WorkplaceNewResponseAddressStateHi, WorkplaceNewResponseAddressStateID, WorkplaceNewResponseAddressStateIl, WorkplaceNewResponseAddressStateIn, WorkplaceNewResponseAddressStateIa, WorkplaceNewResponseAddressStateKs, WorkplaceNewResponseAddressStateKy, WorkplaceNewResponseAddressStateLa, WorkplaceNewResponseAddressStateMe, WorkplaceNewResponseAddressStateMd, WorkplaceNewResponseAddressStateMa, WorkplaceNewResponseAddressStateMi, WorkplaceNewResponseAddressStateMn, WorkplaceNewResponseAddressStateMs, WorkplaceNewResponseAddressStateMo, WorkplaceNewResponseAddressStateMt, WorkplaceNewResponseAddressStateNe, WorkplaceNewResponseAddressStateNv, WorkplaceNewResponseAddressStateNh, WorkplaceNewResponseAddressStateNj, WorkplaceNewResponseAddressStateNm, WorkplaceNewResponseAddressStateNy, WorkplaceNewResponseAddressStateNc, WorkplaceNewResponseAddressStateNd, WorkplaceNewResponseAddressStateOh, WorkplaceNewResponseAddressStateOk, WorkplaceNewResponseAddressStateOr, WorkplaceNewResponseAddressStatePa, WorkplaceNewResponseAddressStateRi, WorkplaceNewResponseAddressStateSc, WorkplaceNewResponseAddressStateSd, WorkplaceNewResponseAddressStateTn, WorkplaceNewResponseAddressStateTx, WorkplaceNewResponseAddressStateUt, WorkplaceNewResponseAddressStateVt, WorkplaceNewResponseAddressStateVa, WorkplaceNewResponseAddressStateWa, WorkplaceNewResponseAddressStateWv, WorkplaceNewResponseAddressStateWi, WorkplaceNewResponseAddressStateWy:
		return true
	}
	return false
}

type WorkplaceNewResponseStatus string

const (
	WorkplaceNewResponseStatusActive   WorkplaceNewResponseStatus = "active"
	WorkplaceNewResponseStatusArchived WorkplaceNewResponseStatus = "archived"
)

func (r WorkplaceNewResponseStatus) IsKnown() bool {
	switch r {
	case WorkplaceNewResponseStatusActive, WorkplaceNewResponseStatusArchived:
		return true
	}
	return false
}

type WorkplaceNewResponseType string

const (
	WorkplaceNewResponseTypeRemote WorkplaceNewResponseType = "remote"
	WorkplaceNewResponseTypeOffice WorkplaceNewResponseType = "office"
)

func (r WorkplaceNewResponseType) IsKnown() bool {
	switch r {
	case WorkplaceNewResponseTypeRemote, WorkplaceNewResponseTypeOffice:
		return true
	}
	return false
}

type WorkplaceUpdateResponse struct {
	// Public workplace identifier
	ID string `json:"id" api:"required"`
	// A valid US address
	Address WorkplaceUpdateResponseAddress `json:"address" api:"required"`
	// a string to be decoded into a Date
	CreatedAt string                        `json:"createdAt" api:"required"`
	Name      string                        `json:"name" api:"required"`
	Status    WorkplaceUpdateResponseStatus `json:"status" api:"required"`
	Type      WorkplaceUpdateResponseType   `json:"type" api:"required"`
	JSON      workplaceUpdateResponseJSON   `json:"-"`
}

// workplaceUpdateResponseJSON contains the JSON metadata for the struct
// [WorkplaceUpdateResponse]
type workplaceUpdateResponseJSON struct {
	ID          apijson.Field
	Address     apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkplaceUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workplaceUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// A valid US address
type WorkplaceUpdateResponseAddress struct {
	City    string                                `json:"city" api:"required"`
	Country WorkplaceUpdateResponseAddressCountry `json:"country" api:"required"`
	// a non empty string
	Line1      string                              `json:"line1" api:"required"`
	PostalCode string                              `json:"postalCode" api:"required"`
	State      WorkplaceUpdateResponseAddressState `json:"state" api:"required"`
	Line2      string                              `json:"line2" api:"nullable"`
	JSON       workplaceUpdateResponseAddressJSON  `json:"-"`
}

// workplaceUpdateResponseAddressJSON contains the JSON metadata for the struct
// [WorkplaceUpdateResponseAddress]
type workplaceUpdateResponseAddressJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Line1       apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	Line2       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkplaceUpdateResponseAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workplaceUpdateResponseAddressJSON) RawJSON() string {
	return r.raw
}

type WorkplaceUpdateResponseAddressCountry string

const (
	WorkplaceUpdateResponseAddressCountryUs WorkplaceUpdateResponseAddressCountry = "US"
)

func (r WorkplaceUpdateResponseAddressCountry) IsKnown() bool {
	switch r {
	case WorkplaceUpdateResponseAddressCountryUs:
		return true
	}
	return false
}

type WorkplaceUpdateResponseAddressState string

const (
	WorkplaceUpdateResponseAddressStateAl WorkplaceUpdateResponseAddressState = "AL"
	WorkplaceUpdateResponseAddressStateAk WorkplaceUpdateResponseAddressState = "AK"
	WorkplaceUpdateResponseAddressStateAz WorkplaceUpdateResponseAddressState = "AZ"
	WorkplaceUpdateResponseAddressStateAr WorkplaceUpdateResponseAddressState = "AR"
	WorkplaceUpdateResponseAddressStateCa WorkplaceUpdateResponseAddressState = "CA"
	WorkplaceUpdateResponseAddressStateCo WorkplaceUpdateResponseAddressState = "CO"
	WorkplaceUpdateResponseAddressStateCt WorkplaceUpdateResponseAddressState = "CT"
	WorkplaceUpdateResponseAddressStateDc WorkplaceUpdateResponseAddressState = "DC"
	WorkplaceUpdateResponseAddressStateDe WorkplaceUpdateResponseAddressState = "DE"
	WorkplaceUpdateResponseAddressStateFl WorkplaceUpdateResponseAddressState = "FL"
	WorkplaceUpdateResponseAddressStateGa WorkplaceUpdateResponseAddressState = "GA"
	WorkplaceUpdateResponseAddressStateHi WorkplaceUpdateResponseAddressState = "HI"
	WorkplaceUpdateResponseAddressStateID WorkplaceUpdateResponseAddressState = "ID"
	WorkplaceUpdateResponseAddressStateIl WorkplaceUpdateResponseAddressState = "IL"
	WorkplaceUpdateResponseAddressStateIn WorkplaceUpdateResponseAddressState = "IN"
	WorkplaceUpdateResponseAddressStateIa WorkplaceUpdateResponseAddressState = "IA"
	WorkplaceUpdateResponseAddressStateKs WorkplaceUpdateResponseAddressState = "KS"
	WorkplaceUpdateResponseAddressStateKy WorkplaceUpdateResponseAddressState = "KY"
	WorkplaceUpdateResponseAddressStateLa WorkplaceUpdateResponseAddressState = "LA"
	WorkplaceUpdateResponseAddressStateMe WorkplaceUpdateResponseAddressState = "ME"
	WorkplaceUpdateResponseAddressStateMd WorkplaceUpdateResponseAddressState = "MD"
	WorkplaceUpdateResponseAddressStateMa WorkplaceUpdateResponseAddressState = "MA"
	WorkplaceUpdateResponseAddressStateMi WorkplaceUpdateResponseAddressState = "MI"
	WorkplaceUpdateResponseAddressStateMn WorkplaceUpdateResponseAddressState = "MN"
	WorkplaceUpdateResponseAddressStateMs WorkplaceUpdateResponseAddressState = "MS"
	WorkplaceUpdateResponseAddressStateMo WorkplaceUpdateResponseAddressState = "MO"
	WorkplaceUpdateResponseAddressStateMt WorkplaceUpdateResponseAddressState = "MT"
	WorkplaceUpdateResponseAddressStateNe WorkplaceUpdateResponseAddressState = "NE"
	WorkplaceUpdateResponseAddressStateNv WorkplaceUpdateResponseAddressState = "NV"
	WorkplaceUpdateResponseAddressStateNh WorkplaceUpdateResponseAddressState = "NH"
	WorkplaceUpdateResponseAddressStateNj WorkplaceUpdateResponseAddressState = "NJ"
	WorkplaceUpdateResponseAddressStateNm WorkplaceUpdateResponseAddressState = "NM"
	WorkplaceUpdateResponseAddressStateNy WorkplaceUpdateResponseAddressState = "NY"
	WorkplaceUpdateResponseAddressStateNc WorkplaceUpdateResponseAddressState = "NC"
	WorkplaceUpdateResponseAddressStateNd WorkplaceUpdateResponseAddressState = "ND"
	WorkplaceUpdateResponseAddressStateOh WorkplaceUpdateResponseAddressState = "OH"
	WorkplaceUpdateResponseAddressStateOk WorkplaceUpdateResponseAddressState = "OK"
	WorkplaceUpdateResponseAddressStateOr WorkplaceUpdateResponseAddressState = "OR"
	WorkplaceUpdateResponseAddressStatePa WorkplaceUpdateResponseAddressState = "PA"
	WorkplaceUpdateResponseAddressStateRi WorkplaceUpdateResponseAddressState = "RI"
	WorkplaceUpdateResponseAddressStateSc WorkplaceUpdateResponseAddressState = "SC"
	WorkplaceUpdateResponseAddressStateSd WorkplaceUpdateResponseAddressState = "SD"
	WorkplaceUpdateResponseAddressStateTn WorkplaceUpdateResponseAddressState = "TN"
	WorkplaceUpdateResponseAddressStateTx WorkplaceUpdateResponseAddressState = "TX"
	WorkplaceUpdateResponseAddressStateUt WorkplaceUpdateResponseAddressState = "UT"
	WorkplaceUpdateResponseAddressStateVt WorkplaceUpdateResponseAddressState = "VT"
	WorkplaceUpdateResponseAddressStateVa WorkplaceUpdateResponseAddressState = "VA"
	WorkplaceUpdateResponseAddressStateWa WorkplaceUpdateResponseAddressState = "WA"
	WorkplaceUpdateResponseAddressStateWv WorkplaceUpdateResponseAddressState = "WV"
	WorkplaceUpdateResponseAddressStateWi WorkplaceUpdateResponseAddressState = "WI"
	WorkplaceUpdateResponseAddressStateWy WorkplaceUpdateResponseAddressState = "WY"
)

func (r WorkplaceUpdateResponseAddressState) IsKnown() bool {
	switch r {
	case WorkplaceUpdateResponseAddressStateAl, WorkplaceUpdateResponseAddressStateAk, WorkplaceUpdateResponseAddressStateAz, WorkplaceUpdateResponseAddressStateAr, WorkplaceUpdateResponseAddressStateCa, WorkplaceUpdateResponseAddressStateCo, WorkplaceUpdateResponseAddressStateCt, WorkplaceUpdateResponseAddressStateDc, WorkplaceUpdateResponseAddressStateDe, WorkplaceUpdateResponseAddressStateFl, WorkplaceUpdateResponseAddressStateGa, WorkplaceUpdateResponseAddressStateHi, WorkplaceUpdateResponseAddressStateID, WorkplaceUpdateResponseAddressStateIl, WorkplaceUpdateResponseAddressStateIn, WorkplaceUpdateResponseAddressStateIa, WorkplaceUpdateResponseAddressStateKs, WorkplaceUpdateResponseAddressStateKy, WorkplaceUpdateResponseAddressStateLa, WorkplaceUpdateResponseAddressStateMe, WorkplaceUpdateResponseAddressStateMd, WorkplaceUpdateResponseAddressStateMa, WorkplaceUpdateResponseAddressStateMi, WorkplaceUpdateResponseAddressStateMn, WorkplaceUpdateResponseAddressStateMs, WorkplaceUpdateResponseAddressStateMo, WorkplaceUpdateResponseAddressStateMt, WorkplaceUpdateResponseAddressStateNe, WorkplaceUpdateResponseAddressStateNv, WorkplaceUpdateResponseAddressStateNh, WorkplaceUpdateResponseAddressStateNj, WorkplaceUpdateResponseAddressStateNm, WorkplaceUpdateResponseAddressStateNy, WorkplaceUpdateResponseAddressStateNc, WorkplaceUpdateResponseAddressStateNd, WorkplaceUpdateResponseAddressStateOh, WorkplaceUpdateResponseAddressStateOk, WorkplaceUpdateResponseAddressStateOr, WorkplaceUpdateResponseAddressStatePa, WorkplaceUpdateResponseAddressStateRi, WorkplaceUpdateResponseAddressStateSc, WorkplaceUpdateResponseAddressStateSd, WorkplaceUpdateResponseAddressStateTn, WorkplaceUpdateResponseAddressStateTx, WorkplaceUpdateResponseAddressStateUt, WorkplaceUpdateResponseAddressStateVt, WorkplaceUpdateResponseAddressStateVa, WorkplaceUpdateResponseAddressStateWa, WorkplaceUpdateResponseAddressStateWv, WorkplaceUpdateResponseAddressStateWi, WorkplaceUpdateResponseAddressStateWy:
		return true
	}
	return false
}

type WorkplaceUpdateResponseStatus string

const (
	WorkplaceUpdateResponseStatusActive   WorkplaceUpdateResponseStatus = "active"
	WorkplaceUpdateResponseStatusArchived WorkplaceUpdateResponseStatus = "archived"
)

func (r WorkplaceUpdateResponseStatus) IsKnown() bool {
	switch r {
	case WorkplaceUpdateResponseStatusActive, WorkplaceUpdateResponseStatusArchived:
		return true
	}
	return false
}

type WorkplaceUpdateResponseType string

const (
	WorkplaceUpdateResponseTypeRemote WorkplaceUpdateResponseType = "remote"
	WorkplaceUpdateResponseTypeOffice WorkplaceUpdateResponseType = "office"
)

func (r WorkplaceUpdateResponseType) IsKnown() bool {
	switch r {
	case WorkplaceUpdateResponseTypeRemote, WorkplaceUpdateResponseTypeOffice:
		return true
	}
	return false
}

type WorkplaceListResponse struct {
	// Public workplace identifier
	ID string `json:"id" api:"required"`
	// A valid US address
	Address WorkplaceListResponseAddress `json:"address" api:"required"`
	// a string to be decoded into a Date
	CreatedAt string                      `json:"createdAt" api:"required"`
	Name      string                      `json:"name" api:"required"`
	Status    WorkplaceListResponseStatus `json:"status" api:"required"`
	Type      WorkplaceListResponseType   `json:"type" api:"required"`
	JSON      workplaceListResponseJSON   `json:"-"`
}

// workplaceListResponseJSON contains the JSON metadata for the struct
// [WorkplaceListResponse]
type workplaceListResponseJSON struct {
	ID          apijson.Field
	Address     apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkplaceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workplaceListResponseJSON) RawJSON() string {
	return r.raw
}

// A valid US address
type WorkplaceListResponseAddress struct {
	City    string                              `json:"city" api:"required"`
	Country WorkplaceListResponseAddressCountry `json:"country" api:"required"`
	// a non empty string
	Line1      string                            `json:"line1" api:"required"`
	PostalCode string                            `json:"postalCode" api:"required"`
	State      WorkplaceListResponseAddressState `json:"state" api:"required"`
	Line2      string                            `json:"line2" api:"nullable"`
	JSON       workplaceListResponseAddressJSON  `json:"-"`
}

// workplaceListResponseAddressJSON contains the JSON metadata for the struct
// [WorkplaceListResponseAddress]
type workplaceListResponseAddressJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Line1       apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	Line2       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkplaceListResponseAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workplaceListResponseAddressJSON) RawJSON() string {
	return r.raw
}

type WorkplaceListResponseAddressCountry string

const (
	WorkplaceListResponseAddressCountryUs WorkplaceListResponseAddressCountry = "US"
)

func (r WorkplaceListResponseAddressCountry) IsKnown() bool {
	switch r {
	case WorkplaceListResponseAddressCountryUs:
		return true
	}
	return false
}

type WorkplaceListResponseAddressState string

const (
	WorkplaceListResponseAddressStateAl WorkplaceListResponseAddressState = "AL"
	WorkplaceListResponseAddressStateAk WorkplaceListResponseAddressState = "AK"
	WorkplaceListResponseAddressStateAz WorkplaceListResponseAddressState = "AZ"
	WorkplaceListResponseAddressStateAr WorkplaceListResponseAddressState = "AR"
	WorkplaceListResponseAddressStateCa WorkplaceListResponseAddressState = "CA"
	WorkplaceListResponseAddressStateCo WorkplaceListResponseAddressState = "CO"
	WorkplaceListResponseAddressStateCt WorkplaceListResponseAddressState = "CT"
	WorkplaceListResponseAddressStateDc WorkplaceListResponseAddressState = "DC"
	WorkplaceListResponseAddressStateDe WorkplaceListResponseAddressState = "DE"
	WorkplaceListResponseAddressStateFl WorkplaceListResponseAddressState = "FL"
	WorkplaceListResponseAddressStateGa WorkplaceListResponseAddressState = "GA"
	WorkplaceListResponseAddressStateHi WorkplaceListResponseAddressState = "HI"
	WorkplaceListResponseAddressStateID WorkplaceListResponseAddressState = "ID"
	WorkplaceListResponseAddressStateIl WorkplaceListResponseAddressState = "IL"
	WorkplaceListResponseAddressStateIn WorkplaceListResponseAddressState = "IN"
	WorkplaceListResponseAddressStateIa WorkplaceListResponseAddressState = "IA"
	WorkplaceListResponseAddressStateKs WorkplaceListResponseAddressState = "KS"
	WorkplaceListResponseAddressStateKy WorkplaceListResponseAddressState = "KY"
	WorkplaceListResponseAddressStateLa WorkplaceListResponseAddressState = "LA"
	WorkplaceListResponseAddressStateMe WorkplaceListResponseAddressState = "ME"
	WorkplaceListResponseAddressStateMd WorkplaceListResponseAddressState = "MD"
	WorkplaceListResponseAddressStateMa WorkplaceListResponseAddressState = "MA"
	WorkplaceListResponseAddressStateMi WorkplaceListResponseAddressState = "MI"
	WorkplaceListResponseAddressStateMn WorkplaceListResponseAddressState = "MN"
	WorkplaceListResponseAddressStateMs WorkplaceListResponseAddressState = "MS"
	WorkplaceListResponseAddressStateMo WorkplaceListResponseAddressState = "MO"
	WorkplaceListResponseAddressStateMt WorkplaceListResponseAddressState = "MT"
	WorkplaceListResponseAddressStateNe WorkplaceListResponseAddressState = "NE"
	WorkplaceListResponseAddressStateNv WorkplaceListResponseAddressState = "NV"
	WorkplaceListResponseAddressStateNh WorkplaceListResponseAddressState = "NH"
	WorkplaceListResponseAddressStateNj WorkplaceListResponseAddressState = "NJ"
	WorkplaceListResponseAddressStateNm WorkplaceListResponseAddressState = "NM"
	WorkplaceListResponseAddressStateNy WorkplaceListResponseAddressState = "NY"
	WorkplaceListResponseAddressStateNc WorkplaceListResponseAddressState = "NC"
	WorkplaceListResponseAddressStateNd WorkplaceListResponseAddressState = "ND"
	WorkplaceListResponseAddressStateOh WorkplaceListResponseAddressState = "OH"
	WorkplaceListResponseAddressStateOk WorkplaceListResponseAddressState = "OK"
	WorkplaceListResponseAddressStateOr WorkplaceListResponseAddressState = "OR"
	WorkplaceListResponseAddressStatePa WorkplaceListResponseAddressState = "PA"
	WorkplaceListResponseAddressStateRi WorkplaceListResponseAddressState = "RI"
	WorkplaceListResponseAddressStateSc WorkplaceListResponseAddressState = "SC"
	WorkplaceListResponseAddressStateSd WorkplaceListResponseAddressState = "SD"
	WorkplaceListResponseAddressStateTn WorkplaceListResponseAddressState = "TN"
	WorkplaceListResponseAddressStateTx WorkplaceListResponseAddressState = "TX"
	WorkplaceListResponseAddressStateUt WorkplaceListResponseAddressState = "UT"
	WorkplaceListResponseAddressStateVt WorkplaceListResponseAddressState = "VT"
	WorkplaceListResponseAddressStateVa WorkplaceListResponseAddressState = "VA"
	WorkplaceListResponseAddressStateWa WorkplaceListResponseAddressState = "WA"
	WorkplaceListResponseAddressStateWv WorkplaceListResponseAddressState = "WV"
	WorkplaceListResponseAddressStateWi WorkplaceListResponseAddressState = "WI"
	WorkplaceListResponseAddressStateWy WorkplaceListResponseAddressState = "WY"
)

func (r WorkplaceListResponseAddressState) IsKnown() bool {
	switch r {
	case WorkplaceListResponseAddressStateAl, WorkplaceListResponseAddressStateAk, WorkplaceListResponseAddressStateAz, WorkplaceListResponseAddressStateAr, WorkplaceListResponseAddressStateCa, WorkplaceListResponseAddressStateCo, WorkplaceListResponseAddressStateCt, WorkplaceListResponseAddressStateDc, WorkplaceListResponseAddressStateDe, WorkplaceListResponseAddressStateFl, WorkplaceListResponseAddressStateGa, WorkplaceListResponseAddressStateHi, WorkplaceListResponseAddressStateID, WorkplaceListResponseAddressStateIl, WorkplaceListResponseAddressStateIn, WorkplaceListResponseAddressStateIa, WorkplaceListResponseAddressStateKs, WorkplaceListResponseAddressStateKy, WorkplaceListResponseAddressStateLa, WorkplaceListResponseAddressStateMe, WorkplaceListResponseAddressStateMd, WorkplaceListResponseAddressStateMa, WorkplaceListResponseAddressStateMi, WorkplaceListResponseAddressStateMn, WorkplaceListResponseAddressStateMs, WorkplaceListResponseAddressStateMo, WorkplaceListResponseAddressStateMt, WorkplaceListResponseAddressStateNe, WorkplaceListResponseAddressStateNv, WorkplaceListResponseAddressStateNh, WorkplaceListResponseAddressStateNj, WorkplaceListResponseAddressStateNm, WorkplaceListResponseAddressStateNy, WorkplaceListResponseAddressStateNc, WorkplaceListResponseAddressStateNd, WorkplaceListResponseAddressStateOh, WorkplaceListResponseAddressStateOk, WorkplaceListResponseAddressStateOr, WorkplaceListResponseAddressStatePa, WorkplaceListResponseAddressStateRi, WorkplaceListResponseAddressStateSc, WorkplaceListResponseAddressStateSd, WorkplaceListResponseAddressStateTn, WorkplaceListResponseAddressStateTx, WorkplaceListResponseAddressStateUt, WorkplaceListResponseAddressStateVt, WorkplaceListResponseAddressStateVa, WorkplaceListResponseAddressStateWa, WorkplaceListResponseAddressStateWv, WorkplaceListResponseAddressStateWi, WorkplaceListResponseAddressStateWy:
		return true
	}
	return false
}

type WorkplaceListResponseStatus string

const (
	WorkplaceListResponseStatusActive   WorkplaceListResponseStatus = "active"
	WorkplaceListResponseStatusArchived WorkplaceListResponseStatus = "archived"
)

func (r WorkplaceListResponseStatus) IsKnown() bool {
	switch r {
	case WorkplaceListResponseStatusActive, WorkplaceListResponseStatusArchived:
		return true
	}
	return false
}

type WorkplaceListResponseType string

const (
	WorkplaceListResponseTypeRemote WorkplaceListResponseType = "remote"
	WorkplaceListResponseTypeOffice WorkplaceListResponseType = "office"
)

func (r WorkplaceListResponseType) IsKnown() bool {
	switch r {
	case WorkplaceListResponseTypeRemote, WorkplaceListResponseTypeOffice:
		return true
	}
	return false
}

type WorkplaceNewParams struct {
	// A valid US address
	Address param.Field[WorkplaceNewParamsAddress] `json:"address" api:"required"`
	// a non empty string
	Name param.Field[string]                 `json:"name" api:"required"`
	Type param.Field[WorkplaceNewParamsType] `json:"type" api:"required"`
}

func (r WorkplaceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A valid US address
type WorkplaceNewParamsAddress struct {
	City    param.Field[string]                           `json:"city" api:"required"`
	Country param.Field[WorkplaceNewParamsAddressCountry] `json:"country" api:"required"`
	// a non empty string
	Line1      param.Field[string]                         `json:"line1" api:"required"`
	PostalCode param.Field[string]                         `json:"postalCode" api:"required"`
	State      param.Field[WorkplaceNewParamsAddressState] `json:"state" api:"required"`
	Line2      param.Field[string]                         `json:"line2"`
}

func (r WorkplaceNewParamsAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkplaceNewParamsAddressCountry string

const (
	WorkplaceNewParamsAddressCountryUs WorkplaceNewParamsAddressCountry = "US"
)

func (r WorkplaceNewParamsAddressCountry) IsKnown() bool {
	switch r {
	case WorkplaceNewParamsAddressCountryUs:
		return true
	}
	return false
}

type WorkplaceNewParamsAddressState string

const (
	WorkplaceNewParamsAddressStateAl WorkplaceNewParamsAddressState = "AL"
	WorkplaceNewParamsAddressStateAk WorkplaceNewParamsAddressState = "AK"
	WorkplaceNewParamsAddressStateAz WorkplaceNewParamsAddressState = "AZ"
	WorkplaceNewParamsAddressStateAr WorkplaceNewParamsAddressState = "AR"
	WorkplaceNewParamsAddressStateCa WorkplaceNewParamsAddressState = "CA"
	WorkplaceNewParamsAddressStateCo WorkplaceNewParamsAddressState = "CO"
	WorkplaceNewParamsAddressStateCt WorkplaceNewParamsAddressState = "CT"
	WorkplaceNewParamsAddressStateDc WorkplaceNewParamsAddressState = "DC"
	WorkplaceNewParamsAddressStateDe WorkplaceNewParamsAddressState = "DE"
	WorkplaceNewParamsAddressStateFl WorkplaceNewParamsAddressState = "FL"
	WorkplaceNewParamsAddressStateGa WorkplaceNewParamsAddressState = "GA"
	WorkplaceNewParamsAddressStateHi WorkplaceNewParamsAddressState = "HI"
	WorkplaceNewParamsAddressStateID WorkplaceNewParamsAddressState = "ID"
	WorkplaceNewParamsAddressStateIl WorkplaceNewParamsAddressState = "IL"
	WorkplaceNewParamsAddressStateIn WorkplaceNewParamsAddressState = "IN"
	WorkplaceNewParamsAddressStateIa WorkplaceNewParamsAddressState = "IA"
	WorkplaceNewParamsAddressStateKs WorkplaceNewParamsAddressState = "KS"
	WorkplaceNewParamsAddressStateKy WorkplaceNewParamsAddressState = "KY"
	WorkplaceNewParamsAddressStateLa WorkplaceNewParamsAddressState = "LA"
	WorkplaceNewParamsAddressStateMe WorkplaceNewParamsAddressState = "ME"
	WorkplaceNewParamsAddressStateMd WorkplaceNewParamsAddressState = "MD"
	WorkplaceNewParamsAddressStateMa WorkplaceNewParamsAddressState = "MA"
	WorkplaceNewParamsAddressStateMi WorkplaceNewParamsAddressState = "MI"
	WorkplaceNewParamsAddressStateMn WorkplaceNewParamsAddressState = "MN"
	WorkplaceNewParamsAddressStateMs WorkplaceNewParamsAddressState = "MS"
	WorkplaceNewParamsAddressStateMo WorkplaceNewParamsAddressState = "MO"
	WorkplaceNewParamsAddressStateMt WorkplaceNewParamsAddressState = "MT"
	WorkplaceNewParamsAddressStateNe WorkplaceNewParamsAddressState = "NE"
	WorkplaceNewParamsAddressStateNv WorkplaceNewParamsAddressState = "NV"
	WorkplaceNewParamsAddressStateNh WorkplaceNewParamsAddressState = "NH"
	WorkplaceNewParamsAddressStateNj WorkplaceNewParamsAddressState = "NJ"
	WorkplaceNewParamsAddressStateNm WorkplaceNewParamsAddressState = "NM"
	WorkplaceNewParamsAddressStateNy WorkplaceNewParamsAddressState = "NY"
	WorkplaceNewParamsAddressStateNc WorkplaceNewParamsAddressState = "NC"
	WorkplaceNewParamsAddressStateNd WorkplaceNewParamsAddressState = "ND"
	WorkplaceNewParamsAddressStateOh WorkplaceNewParamsAddressState = "OH"
	WorkplaceNewParamsAddressStateOk WorkplaceNewParamsAddressState = "OK"
	WorkplaceNewParamsAddressStateOr WorkplaceNewParamsAddressState = "OR"
	WorkplaceNewParamsAddressStatePa WorkplaceNewParamsAddressState = "PA"
	WorkplaceNewParamsAddressStateRi WorkplaceNewParamsAddressState = "RI"
	WorkplaceNewParamsAddressStateSc WorkplaceNewParamsAddressState = "SC"
	WorkplaceNewParamsAddressStateSd WorkplaceNewParamsAddressState = "SD"
	WorkplaceNewParamsAddressStateTn WorkplaceNewParamsAddressState = "TN"
	WorkplaceNewParamsAddressStateTx WorkplaceNewParamsAddressState = "TX"
	WorkplaceNewParamsAddressStateUt WorkplaceNewParamsAddressState = "UT"
	WorkplaceNewParamsAddressStateVt WorkplaceNewParamsAddressState = "VT"
	WorkplaceNewParamsAddressStateVa WorkplaceNewParamsAddressState = "VA"
	WorkplaceNewParamsAddressStateWa WorkplaceNewParamsAddressState = "WA"
	WorkplaceNewParamsAddressStateWv WorkplaceNewParamsAddressState = "WV"
	WorkplaceNewParamsAddressStateWi WorkplaceNewParamsAddressState = "WI"
	WorkplaceNewParamsAddressStateWy WorkplaceNewParamsAddressState = "WY"
)

func (r WorkplaceNewParamsAddressState) IsKnown() bool {
	switch r {
	case WorkplaceNewParamsAddressStateAl, WorkplaceNewParamsAddressStateAk, WorkplaceNewParamsAddressStateAz, WorkplaceNewParamsAddressStateAr, WorkplaceNewParamsAddressStateCa, WorkplaceNewParamsAddressStateCo, WorkplaceNewParamsAddressStateCt, WorkplaceNewParamsAddressStateDc, WorkplaceNewParamsAddressStateDe, WorkplaceNewParamsAddressStateFl, WorkplaceNewParamsAddressStateGa, WorkplaceNewParamsAddressStateHi, WorkplaceNewParamsAddressStateID, WorkplaceNewParamsAddressStateIl, WorkplaceNewParamsAddressStateIn, WorkplaceNewParamsAddressStateIa, WorkplaceNewParamsAddressStateKs, WorkplaceNewParamsAddressStateKy, WorkplaceNewParamsAddressStateLa, WorkplaceNewParamsAddressStateMe, WorkplaceNewParamsAddressStateMd, WorkplaceNewParamsAddressStateMa, WorkplaceNewParamsAddressStateMi, WorkplaceNewParamsAddressStateMn, WorkplaceNewParamsAddressStateMs, WorkplaceNewParamsAddressStateMo, WorkplaceNewParamsAddressStateMt, WorkplaceNewParamsAddressStateNe, WorkplaceNewParamsAddressStateNv, WorkplaceNewParamsAddressStateNh, WorkplaceNewParamsAddressStateNj, WorkplaceNewParamsAddressStateNm, WorkplaceNewParamsAddressStateNy, WorkplaceNewParamsAddressStateNc, WorkplaceNewParamsAddressStateNd, WorkplaceNewParamsAddressStateOh, WorkplaceNewParamsAddressStateOk, WorkplaceNewParamsAddressStateOr, WorkplaceNewParamsAddressStatePa, WorkplaceNewParamsAddressStateRi, WorkplaceNewParamsAddressStateSc, WorkplaceNewParamsAddressStateSd, WorkplaceNewParamsAddressStateTn, WorkplaceNewParamsAddressStateTx, WorkplaceNewParamsAddressStateUt, WorkplaceNewParamsAddressStateVt, WorkplaceNewParamsAddressStateVa, WorkplaceNewParamsAddressStateWa, WorkplaceNewParamsAddressStateWv, WorkplaceNewParamsAddressStateWi, WorkplaceNewParamsAddressStateWy:
		return true
	}
	return false
}

type WorkplaceNewParamsType string

const (
	WorkplaceNewParamsTypeRemote WorkplaceNewParamsType = "remote"
	WorkplaceNewParamsTypeOffice WorkplaceNewParamsType = "office"
)

func (r WorkplaceNewParamsType) IsKnown() bool {
	switch r {
	case WorkplaceNewParamsTypeRemote, WorkplaceNewParamsTypeOffice:
		return true
	}
	return false
}

type WorkplaceUpdateParams struct {
	Name param.Field[string] `json:"name"`
}

func (r WorkplaceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkplaceListParams struct {
	// Public workplace identifier
	AfterID param.Field[string] `query:"afterId"`
	// Public workplace identifier
	BeforeID param.Field[string] `query:"beforeId"`
	// a number less than or equal to 100
	Limit param.Field[string] `query:"limit"`
}

// URLQuery serializes [WorkplaceListParams]'s query parameters as `url.Values`.
func (r WorkplaceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
