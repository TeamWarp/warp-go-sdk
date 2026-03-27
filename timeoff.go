// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package warphr

import (
	"context"
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

// Endpoints for worker time off management. See time off requests, which workers
// are assigned to which policies, or worker remaining balances.
//
// TimeOffService contains methods and other services that help with interacting
// with the warp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTimeOffService] method instead.
type TimeOffService struct {
	Options []option.RequestOption
	// Endpoints for worker time off management. See time off requests, which workers
	// are assigned to which policies, or worker remaining balances.
	Policies *TimeOffPolicyService
}

// NewTimeOffService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTimeOffService(opts ...option.RequestOption) (r *TimeOffService) {
	r = &TimeOffService{}
	r.Options = opts
	r.Policies = NewTimeOffPolicyService(opts...)
	return
}

// Time off assignments are mappings between workers and time off policies. Useful
// for finding out which policies a worker is assigned to, or which workers are
// assigned to a given policy.
func (r *TimeOffService) ListAssignments(ctx context.Context, query TimeOffListAssignmentsParams, opts ...option.RequestOption) (res *pagination.CursorPage[TimeOffListAssignmentsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/time_off/assignments"
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

// Time off assignments are mappings between workers and time off policies. Useful
// for finding out which policies a worker is assigned to, or which workers are
// assigned to a given policy.
func (r *TimeOffService) ListAssignmentsAutoPaging(ctx context.Context, query TimeOffListAssignmentsParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[TimeOffListAssignmentsResponse] {
	return pagination.NewCursorPageAutoPager(r.ListAssignments(ctx, query, opts...))
}

// Get worker remaining time-off balances.
func (r *TimeOffService) ListBalances(ctx context.Context, query TimeOffListBalancesParams, opts ...option.RequestOption) (res *pagination.CursorPage[TimeOffListBalancesResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/time_off/balances"
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

// Get worker remaining time-off balances.
func (r *TimeOffService) ListBalancesAutoPaging(ctx context.Context, query TimeOffListBalancesParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[TimeOffListBalancesResponse] {
	return pagination.NewCursorPageAutoPager(r.ListBalances(ctx, query, opts...))
}

// Get the time off requests that workers in your company have made.
func (r *TimeOffService) ListRequests(ctx context.Context, query TimeOffListRequestsParams, opts ...option.RequestOption) (res *pagination.CursorPage[TimeOffListRequestsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/time_off/requests"
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

// Get the time off requests that workers in your company have made.
func (r *TimeOffService) ListRequestsAutoPaging(ctx context.Context, query TimeOffListRequestsParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[TimeOffListRequestsResponse] {
	return pagination.NewCursorPageAutoPager(r.ListRequests(ctx, query, opts...))
}

type TimeOffListAssignmentsResponse struct {
	ID string `json:"id" api:"required"`
	// a string to be decoded into a Date
	AssignedAt string `json:"assignedAt" api:"required"`
	// a string starting with "top\_"
	PolicyID string `json:"policyId" api:"required"`
	// The id of the worker.
	WorkerID string                             `json:"workerId" api:"required"`
	JSON     timeOffListAssignmentsResponseJSON `json:"-"`
}

// timeOffListAssignmentsResponseJSON contains the JSON metadata for the struct
// [TimeOffListAssignmentsResponse]
type timeOffListAssignmentsResponseJSON struct {
	ID          apijson.Field
	AssignedAt  apijson.Field
	PolicyID    apijson.Field
	WorkerID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimeOffListAssignmentsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timeOffListAssignmentsResponseJSON) RawJSON() string {
	return r.raw
}

type TimeOffListBalancesResponse struct {
	ID              string  `json:"id" api:"required"`
	AccruedLocked   float64 `json:"accruedLocked" api:"required"`
	AccruedUnlocked float64 `json:"accruedUnlocked" api:"required"`
	Available       float64 `json:"available" api:"required"`
	Holds           float64 `json:"holds" api:"required"`
	LegacyWorkerID  string  `json:"legacyWorkerId" api:"required"`
	// a string starting with "top\_"
	PolicyID string                          `json:"policyId" api:"required"`
	Used     float64                         `json:"used" api:"required"`
	JSON     timeOffListBalancesResponseJSON `json:"-"`
}

// timeOffListBalancesResponseJSON contains the JSON metadata for the struct
// [TimeOffListBalancesResponse]
type timeOffListBalancesResponseJSON struct {
	ID              apijson.Field
	AccruedLocked   apijson.Field
	AccruedUnlocked apijson.Field
	Available       apijson.Field
	Holds           apijson.Field
	LegacyWorkerID  apijson.Field
	PolicyID        apijson.Field
	Used            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TimeOffListBalancesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timeOffListBalancesResponseJSON) RawJSON() string {
	return r.raw
}

type TimeOffListRequestsResponse struct {
	ID string `json:"id" api:"required"`
	// a string to be decoded into a Date
	CreatedAt string `json:"createdAt" api:"required"`
	// a string to be decoded into a Date
	EndAt            string  `json:"endAt" api:"required"`
	Reason           string  `json:"reason" api:"required,nullable"`
	RequestedMinutes float64 `json:"requestedMinutes" api:"required"`
	// a string to be decoded into a Date
	StartAt string                            `json:"startAt" api:"required"`
	Status  TimeOffListRequestsResponseStatus `json:"status" api:"required"`
	// a string starting with "top\_"
	TimeOffPolicyID string `json:"timeOffPolicyId" api:"required"`
	// The time zone that the worker is requesting time off in.
	TimeZone string `json:"timeZone" api:"required,nullable"`
	// The id of the worker.
	WorkerID string                          `json:"workerId" api:"required"`
	JSON     timeOffListRequestsResponseJSON `json:"-"`
}

// timeOffListRequestsResponseJSON contains the JSON metadata for the struct
// [TimeOffListRequestsResponse]
type timeOffListRequestsResponseJSON struct {
	ID               apijson.Field
	CreatedAt        apijson.Field
	EndAt            apijson.Field
	Reason           apijson.Field
	RequestedMinutes apijson.Field
	StartAt          apijson.Field
	Status           apijson.Field
	TimeOffPolicyID  apijson.Field
	TimeZone         apijson.Field
	WorkerID         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TimeOffListRequestsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timeOffListRequestsResponseJSON) RawJSON() string {
	return r.raw
}

type TimeOffListRequestsResponseStatus string

const (
	TimeOffListRequestsResponseStatusPending  TimeOffListRequestsResponseStatus = "pending"
	TimeOffListRequestsResponseStatusApproved TimeOffListRequestsResponseStatus = "approved"
	TimeOffListRequestsResponseStatusDenied   TimeOffListRequestsResponseStatus = "denied"
)

func (r TimeOffListRequestsResponseStatus) IsKnown() bool {
	switch r {
	case TimeOffListRequestsResponseStatusPending, TimeOffListRequestsResponseStatusApproved, TimeOffListRequestsResponseStatusDenied:
		return true
	}
	return false
}

type TimeOffListAssignmentsParams struct {
	AfterID  param.Field[string] `query:"afterId"`
	BeforeID param.Field[string] `query:"beforeId"`
	// a number less than or equal to 100
	Limit     param.Field[string]   `query:"limit"`
	PolicyIDs param.Field[[]string] `query:"policyIds"`
	WorkerIDs param.Field[[]string] `query:"workerIds"`
}

// URLQuery serializes [TimeOffListAssignmentsParams]'s query parameters as
// `url.Values`.
func (r TimeOffListAssignmentsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TimeOffListBalancesParams struct {
	AfterID  param.Field[string] `query:"afterId"`
	BeforeID param.Field[string] `query:"beforeId"`
	// a string to be decoded into a Date
	EndDate param.Field[string] `query:"endDate"`
	// a number less than or equal to 100
	Limit     param.Field[string]   `query:"limit"`
	PolicyIDs param.Field[[]string] `query:"policyIds"`
	// a string to be decoded into a Date
	StartDate param.Field[string]   `query:"startDate"`
	WorkerIDs param.Field[[]string] `query:"workerIds"`
}

// URLQuery serializes [TimeOffListBalancesParams]'s query parameters as
// `url.Values`.
func (r TimeOffListBalancesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TimeOffListRequestsParams struct {
	AfterID  param.Field[string] `query:"afterId"`
	BeforeID param.Field[string] `query:"beforeId"`
	// a string to be decoded into a Date
	EndsBefore param.Field[string] `query:"endsBefore"`
	// a string to be decoded into a Date
	EndsOnOrAfter param.Field[string] `query:"endsOnOrAfter"`
	// a number less than or equal to 100
	Limit     param.Field[string]   `query:"limit"`
	PolicyIDs param.Field[[]string] `query:"policyIds"`
	// a string to be decoded into a Date
	StartsBefore param.Field[string] `query:"startsBefore"`
	// a string to be decoded into a Date
	StartsOnOrAfter param.Field[string]                            `query:"startsOnOrAfter"`
	Statuses        param.Field[[]TimeOffListRequestsParamsStatus] `query:"statuses"`
	WorkerIDs       param.Field[[]string]                          `query:"workerIds"`
}

// URLQuery serializes [TimeOffListRequestsParams]'s query parameters as
// `url.Values`.
func (r TimeOffListRequestsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TimeOffListRequestsParamsStatus string

const (
	TimeOffListRequestsParamsStatusPending  TimeOffListRequestsParamsStatus = "pending"
	TimeOffListRequestsParamsStatusApproved TimeOffListRequestsParamsStatus = "approved"
	TimeOffListRequestsParamsStatusDenied   TimeOffListRequestsParamsStatus = "denied"
)

func (r TimeOffListRequestsParamsStatus) IsKnown() bool {
	switch r {
	case TimeOffListRequestsParamsStatusPending, TimeOffListRequestsParamsStatusApproved, TimeOffListRequestsParamsStatusDenied:
		return true
	}
	return false
}
