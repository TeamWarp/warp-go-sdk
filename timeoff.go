// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package warphr

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/warp-hr-go/internal/apijson"
	"github.com/stainless-sdks/warp-hr-go/internal/apiquery"
	"github.com/stainless-sdks/warp-hr-go/internal/requestconfig"
	"github.com/stainless-sdks/warp-hr-go/option"
	"github.com/stainless-sdks/warp-hr-go/packages/pagination"
	"github.com/stainless-sdks/warp-hr-go/packages/param"
	"github.com/stainless-sdks/warp-hr-go/packages/respjson"
)

// Endpoints for worker time off management. See time off requests, which workers
// are assigned to which policies, or worker remaining balances.
//
// TimeOffService contains methods and other services that help with interacting
// with the warp-hr API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTimeOffService] method instead.
type TimeOffService struct {
	options []option.RequestOption
	// Endpoints for worker time off management. See time off requests, which workers
	// are assigned to which policies, or worker remaining balances.
	Policies TimeOffPolicyService
}

// NewTimeOffService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTimeOffService(opts ...option.RequestOption) (r TimeOffService) {
	r = TimeOffService{}
	r.options = opts
	r.Policies = NewTimeOffPolicyService(opts...)
	return
}

// Time off assignments are mappings between workers and time off policies. Useful
// for finding out which policies a worker is assigned to, or which workers are
// assigned to a given policy.
func (r *TimeOffService) ListAssignments(ctx context.Context, query TimeOffListAssignmentsParams, opts ...option.RequestOption) (res *pagination.CursorPage[TimeOffListAssignmentsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
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
	opts = slices.Concat(r.options, opts)
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
	opts = slices.Concat(r.options, opts)
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
	WorkerID string `json:"workerId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AssignedAt  respjson.Field
		PolicyID    respjson.Field
		WorkerID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TimeOffListAssignmentsResponse) RawJSON() string { return r.JSON.raw }
func (r *TimeOffListAssignmentsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TimeOffListBalancesResponse struct {
	ID              string  `json:"id" api:"required"`
	AccruedLocked   float64 `json:"accruedLocked" api:"required"`
	AccruedUnlocked float64 `json:"accruedUnlocked" api:"required"`
	Available       float64 `json:"available" api:"required"`
	Holds           float64 `json:"holds" api:"required"`
	LegacyWorkerID  string  `json:"legacyWorkerId" api:"required"`
	// a string starting with "top\_"
	PolicyID string  `json:"policyId" api:"required"`
	Used     float64 `json:"used" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AccruedLocked   respjson.Field
		AccruedUnlocked respjson.Field
		Available       respjson.Field
		Holds           respjson.Field
		LegacyWorkerID  respjson.Field
		PolicyID        respjson.Field
		Used            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TimeOffListBalancesResponse) RawJSON() string { return r.JSON.raw }
func (r *TimeOffListBalancesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TimeOffListRequestsResponse struct {
	ID string `json:"id" api:"required"`
	// a string to be decoded into a Date
	CreatedAt string `json:"createdAt" api:"required"`
	// a string to be decoded into a Date
	EndAt            string  `json:"endAt" api:"required"`
	Reason           string  `json:"reason" api:"required"`
	RequestedMinutes float64 `json:"requestedMinutes" api:"required"`
	// a string to be decoded into a Date
	StartAt string `json:"startAt" api:"required"`
	// Any of "pending", "approved", "denied".
	Status TimeOffListRequestsResponseStatus `json:"status" api:"required"`
	// a string starting with "top\_"
	TimeOffPolicyID string `json:"timeOffPolicyId" api:"required"`
	// The time zone that the worker is requesting time off in.
	TimeZone string `json:"timeZone" api:"required"`
	// The id of the worker.
	WorkerID string `json:"workerId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		EndAt            respjson.Field
		Reason           respjson.Field
		RequestedMinutes respjson.Field
		StartAt          respjson.Field
		Status           respjson.Field
		TimeOffPolicyID  respjson.Field
		TimeZone         respjson.Field
		WorkerID         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TimeOffListRequestsResponse) RawJSON() string { return r.JSON.raw }
func (r *TimeOffListRequestsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TimeOffListRequestsResponseStatus string

const (
	TimeOffListRequestsResponseStatusPending  TimeOffListRequestsResponseStatus = "pending"
	TimeOffListRequestsResponseStatusApproved TimeOffListRequestsResponseStatus = "approved"
	TimeOffListRequestsResponseStatusDenied   TimeOffListRequestsResponseStatus = "denied"
)

type TimeOffListAssignmentsParams struct {
	AfterID  param.Opt[string] `query:"afterId,omitzero" json:"-"`
	BeforeID param.Opt[string] `query:"beforeId,omitzero" json:"-"`
	// a number less than or equal to 100
	Limit     param.Opt[string] `query:"limit,omitzero" json:"-"`
	PolicyIDs []string          `query:"policyIds,omitzero" json:"-"`
	WorkerIDs []string          `query:"workerIds,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TimeOffListAssignmentsParams]'s query parameters as
// `url.Values`.
func (r TimeOffListAssignmentsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TimeOffListBalancesParams struct {
	AfterID  param.Opt[string] `query:"afterId,omitzero" json:"-"`
	BeforeID param.Opt[string] `query:"beforeId,omitzero" json:"-"`
	// a string to be decoded into a Date
	EndDate param.Opt[string] `query:"endDate,omitzero" json:"-"`
	// a number less than or equal to 100
	Limit param.Opt[string] `query:"limit,omitzero" json:"-"`
	// a string to be decoded into a Date
	StartDate param.Opt[string] `query:"startDate,omitzero" json:"-"`
	PolicyIDs []string          `query:"policyIds,omitzero" json:"-"`
	WorkerIDs []string          `query:"workerIds,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TimeOffListBalancesParams]'s query parameters as
// `url.Values`.
func (r TimeOffListBalancesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TimeOffListRequestsParams struct {
	AfterID  param.Opt[string] `query:"afterId,omitzero" json:"-"`
	BeforeID param.Opt[string] `query:"beforeId,omitzero" json:"-"`
	// a string to be decoded into a Date
	EndsBefore param.Opt[string] `query:"endsBefore,omitzero" json:"-"`
	// a string to be decoded into a Date
	EndsOnOrAfter param.Opt[string] `query:"endsOnOrAfter,omitzero" json:"-"`
	// a number less than or equal to 100
	Limit param.Opt[string] `query:"limit,omitzero" json:"-"`
	// a string to be decoded into a Date
	StartsBefore param.Opt[string] `query:"startsBefore,omitzero" json:"-"`
	// a string to be decoded into a Date
	StartsOnOrAfter param.Opt[string] `query:"startsOnOrAfter,omitzero" json:"-"`
	PolicyIDs       []string          `query:"policyIds,omitzero" json:"-"`
	// Any of "pending", "approved", "denied".
	Statuses  []string `query:"statuses,omitzero" json:"-"`
	WorkerIDs []string `query:"workerIds,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TimeOffListRequestsParams]'s query parameters as
// `url.Values`.
func (r TimeOffListRequestsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
