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

// Endpoints for worker time off management. See time off requests, which workers
// are assigned to which policies, or worker remaining balances.
//
// TimeOffPolicyService contains methods and other services that help with
// interacting with the warp-hr API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTimeOffPolicyService] method instead.
type TimeOffPolicyService struct {
	options []option.RequestOption
}

// NewTimeOffPolicyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTimeOffPolicyService(opts ...option.RequestOption) (r TimeOffPolicyService) {
	r = TimeOffPolicyService{}
	r.options = opts
	return
}

// Get a specific time off policy by id
func (r *TimeOffPolicyService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *TimeOffPolicyGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/time_off/policies/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get the time off policies for your company
func (r *TimeOffPolicyService) List(ctx context.Context, query TimeOffPolicyListParams, opts ...option.RequestOption) (res *pagination.CursorPage[TimeOffPolicyListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/time_off/policies"
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

// Get the time off policies for your company
func (r *TimeOffPolicyService) ListAutoPaging(ctx context.Context, query TimeOffPolicyListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[TimeOffPolicyListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

type TimeOffPolicyGetResponse struct {
	// a string starting with "top\_"
	ID                  string  `json:"id" api:"required"`
	Description         string  `json:"description" api:"required"`
	HoursWorkedPerChunk float64 `json:"hoursWorkedPerChunk" api:"required"`
	IsUnlimited         bool    `json:"isUnlimited" api:"required"`
	MinutesPerChunk     float64 `json:"minutesPerChunk" api:"required"`
	MinutesPerPeriod    float64 `json:"minutesPerPeriod" api:"required"`
	Name                string  `json:"name" api:"required"`
	Paid                bool    `json:"paid" api:"required"`
	// Any of "per_hour_worked", "monthly", "yearly", "unlimited".
	Schedule TimeOffPolicyGetResponseSchedule `json:"schedule" api:"required"`
	// a string starting with "tot\_"
	TimeOffTypeID   string `json:"timeOffTypeId" api:"required"`
	TimeOffTypeName string `json:"timeOffTypeName" api:"required"`
	// Any of "hour", "day".
	Unit TimeOffPolicyGetResponseUnit `json:"unit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Description         respjson.Field
		HoursWorkedPerChunk respjson.Field
		IsUnlimited         respjson.Field
		MinutesPerChunk     respjson.Field
		MinutesPerPeriod    respjson.Field
		Name                respjson.Field
		Paid                respjson.Field
		Schedule            respjson.Field
		TimeOffTypeID       respjson.Field
		TimeOffTypeName     respjson.Field
		Unit                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TimeOffPolicyGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TimeOffPolicyGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TimeOffPolicyGetResponseSchedule string

const (
	TimeOffPolicyGetResponseSchedulePerHourWorked TimeOffPolicyGetResponseSchedule = "per_hour_worked"
	TimeOffPolicyGetResponseScheduleMonthly       TimeOffPolicyGetResponseSchedule = "monthly"
	TimeOffPolicyGetResponseScheduleYearly        TimeOffPolicyGetResponseSchedule = "yearly"
	TimeOffPolicyGetResponseScheduleUnlimited     TimeOffPolicyGetResponseSchedule = "unlimited"
)

type TimeOffPolicyGetResponseUnit string

const (
	TimeOffPolicyGetResponseUnitHour TimeOffPolicyGetResponseUnit = "hour"
	TimeOffPolicyGetResponseUnitDay  TimeOffPolicyGetResponseUnit = "day"
)

type TimeOffPolicyListResponse struct {
	// a string starting with "top\_"
	ID                  string  `json:"id" api:"required"`
	Description         string  `json:"description" api:"required"`
	HoursWorkedPerChunk float64 `json:"hoursWorkedPerChunk" api:"required"`
	IsUnlimited         bool    `json:"isUnlimited" api:"required"`
	MinutesPerChunk     float64 `json:"minutesPerChunk" api:"required"`
	MinutesPerPeriod    float64 `json:"minutesPerPeriod" api:"required"`
	Name                string  `json:"name" api:"required"`
	Paid                bool    `json:"paid" api:"required"`
	// Any of "per_hour_worked", "monthly", "yearly", "unlimited".
	Schedule TimeOffPolicyListResponseSchedule `json:"schedule" api:"required"`
	// a string starting with "tot\_"
	TimeOffTypeID   string `json:"timeOffTypeId" api:"required"`
	TimeOffTypeName string `json:"timeOffTypeName" api:"required"`
	// Any of "hour", "day".
	Unit TimeOffPolicyListResponseUnit `json:"unit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Description         respjson.Field
		HoursWorkedPerChunk respjson.Field
		IsUnlimited         respjson.Field
		MinutesPerChunk     respjson.Field
		MinutesPerPeriod    respjson.Field
		Name                respjson.Field
		Paid                respjson.Field
		Schedule            respjson.Field
		TimeOffTypeID       respjson.Field
		TimeOffTypeName     respjson.Field
		Unit                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TimeOffPolicyListResponse) RawJSON() string { return r.JSON.raw }
func (r *TimeOffPolicyListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TimeOffPolicyListResponseSchedule string

const (
	TimeOffPolicyListResponseSchedulePerHourWorked TimeOffPolicyListResponseSchedule = "per_hour_worked"
	TimeOffPolicyListResponseScheduleMonthly       TimeOffPolicyListResponseSchedule = "monthly"
	TimeOffPolicyListResponseScheduleYearly        TimeOffPolicyListResponseSchedule = "yearly"
	TimeOffPolicyListResponseScheduleUnlimited     TimeOffPolicyListResponseSchedule = "unlimited"
)

type TimeOffPolicyListResponseUnit string

const (
	TimeOffPolicyListResponseUnitHour TimeOffPolicyListResponseUnit = "hour"
	TimeOffPolicyListResponseUnitDay  TimeOffPolicyListResponseUnit = "day"
)

type TimeOffPolicyListParams struct {
	// a string starting with "top\_"
	AfterID param.Opt[string] `query:"afterId,omitzero" json:"-"`
	// a string starting with "top\_"
	BeforeID param.Opt[string] `query:"beforeId,omitzero" json:"-"`
	// a number less than or equal to 100
	Limit param.Opt[string] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TimeOffPolicyListParams]'s query parameters as
// `url.Values`.
func (r TimeOffPolicyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
