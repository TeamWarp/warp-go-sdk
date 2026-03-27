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

// Endpoints for worker time off management. See time off requests, which workers
// are assigned to which policies, or worker remaining balances.
//
// TimeOffPolicyService contains methods and other services that help with
// interacting with the warp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTimeOffPolicyService] method instead.
type TimeOffPolicyService struct {
	Options []option.RequestOption
}

// NewTimeOffPolicyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTimeOffPolicyService(opts ...option.RequestOption) (r *TimeOffPolicyService) {
	r = &TimeOffPolicyService{}
	r.Options = opts
	return
}

// Get a specific time off policy by id
func (r *TimeOffPolicyService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *TimeOffPolicyGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/time_off/policies/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get the time off policies for your company
func (r *TimeOffPolicyService) List(ctx context.Context, query TimeOffPolicyListParams, opts ...option.RequestOption) (res *pagination.CursorPage[TimeOffPolicyListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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
	ID                  string                           `json:"id" api:"required"`
	Description         string                           `json:"description" api:"required,nullable"`
	HoursWorkedPerChunk float64                          `json:"hoursWorkedPerChunk" api:"required,nullable"`
	IsUnlimited         bool                             `json:"isUnlimited" api:"required"`
	MinutesPerChunk     float64                          `json:"minutesPerChunk" api:"required,nullable"`
	MinutesPerPeriod    float64                          `json:"minutesPerPeriod" api:"required,nullable"`
	Name                string                           `json:"name" api:"required"`
	Paid                bool                             `json:"paid" api:"required"`
	Schedule            TimeOffPolicyGetResponseSchedule `json:"schedule" api:"required"`
	// a string starting with "tot\_"
	TimeOffTypeID   string                       `json:"timeOffTypeId" api:"required"`
	TimeOffTypeName string                       `json:"timeOffTypeName" api:"required"`
	Unit            TimeOffPolicyGetResponseUnit `json:"unit" api:"required"`
	JSON            timeOffPolicyGetResponseJSON `json:"-"`
}

// timeOffPolicyGetResponseJSON contains the JSON metadata for the struct
// [TimeOffPolicyGetResponse]
type timeOffPolicyGetResponseJSON struct {
	ID                  apijson.Field
	Description         apijson.Field
	HoursWorkedPerChunk apijson.Field
	IsUnlimited         apijson.Field
	MinutesPerChunk     apijson.Field
	MinutesPerPeriod    apijson.Field
	Name                apijson.Field
	Paid                apijson.Field
	Schedule            apijson.Field
	TimeOffTypeID       apijson.Field
	TimeOffTypeName     apijson.Field
	Unit                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *TimeOffPolicyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timeOffPolicyGetResponseJSON) RawJSON() string {
	return r.raw
}

type TimeOffPolicyGetResponseSchedule string

const (
	TimeOffPolicyGetResponseSchedulePerHourWorked TimeOffPolicyGetResponseSchedule = "per_hour_worked"
	TimeOffPolicyGetResponseScheduleMonthly       TimeOffPolicyGetResponseSchedule = "monthly"
	TimeOffPolicyGetResponseScheduleYearly        TimeOffPolicyGetResponseSchedule = "yearly"
	TimeOffPolicyGetResponseScheduleUnlimited     TimeOffPolicyGetResponseSchedule = "unlimited"
)

func (r TimeOffPolicyGetResponseSchedule) IsKnown() bool {
	switch r {
	case TimeOffPolicyGetResponseSchedulePerHourWorked, TimeOffPolicyGetResponseScheduleMonthly, TimeOffPolicyGetResponseScheduleYearly, TimeOffPolicyGetResponseScheduleUnlimited:
		return true
	}
	return false
}

type TimeOffPolicyGetResponseUnit string

const (
	TimeOffPolicyGetResponseUnitHour TimeOffPolicyGetResponseUnit = "hour"
	TimeOffPolicyGetResponseUnitDay  TimeOffPolicyGetResponseUnit = "day"
)

func (r TimeOffPolicyGetResponseUnit) IsKnown() bool {
	switch r {
	case TimeOffPolicyGetResponseUnitHour, TimeOffPolicyGetResponseUnitDay:
		return true
	}
	return false
}

type TimeOffPolicyListResponse struct {
	// a string starting with "top\_"
	ID                  string                            `json:"id" api:"required"`
	Description         string                            `json:"description" api:"required,nullable"`
	HoursWorkedPerChunk float64                           `json:"hoursWorkedPerChunk" api:"required,nullable"`
	IsUnlimited         bool                              `json:"isUnlimited" api:"required"`
	MinutesPerChunk     float64                           `json:"minutesPerChunk" api:"required,nullable"`
	MinutesPerPeriod    float64                           `json:"minutesPerPeriod" api:"required,nullable"`
	Name                string                            `json:"name" api:"required"`
	Paid                bool                              `json:"paid" api:"required"`
	Schedule            TimeOffPolicyListResponseSchedule `json:"schedule" api:"required"`
	// a string starting with "tot\_"
	TimeOffTypeID   string                        `json:"timeOffTypeId" api:"required"`
	TimeOffTypeName string                        `json:"timeOffTypeName" api:"required"`
	Unit            TimeOffPolicyListResponseUnit `json:"unit" api:"required"`
	JSON            timeOffPolicyListResponseJSON `json:"-"`
}

// timeOffPolicyListResponseJSON contains the JSON metadata for the struct
// [TimeOffPolicyListResponse]
type timeOffPolicyListResponseJSON struct {
	ID                  apijson.Field
	Description         apijson.Field
	HoursWorkedPerChunk apijson.Field
	IsUnlimited         apijson.Field
	MinutesPerChunk     apijson.Field
	MinutesPerPeriod    apijson.Field
	Name                apijson.Field
	Paid                apijson.Field
	Schedule            apijson.Field
	TimeOffTypeID       apijson.Field
	TimeOffTypeName     apijson.Field
	Unit                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *TimeOffPolicyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timeOffPolicyListResponseJSON) RawJSON() string {
	return r.raw
}

type TimeOffPolicyListResponseSchedule string

const (
	TimeOffPolicyListResponseSchedulePerHourWorked TimeOffPolicyListResponseSchedule = "per_hour_worked"
	TimeOffPolicyListResponseScheduleMonthly       TimeOffPolicyListResponseSchedule = "monthly"
	TimeOffPolicyListResponseScheduleYearly        TimeOffPolicyListResponseSchedule = "yearly"
	TimeOffPolicyListResponseScheduleUnlimited     TimeOffPolicyListResponseSchedule = "unlimited"
)

func (r TimeOffPolicyListResponseSchedule) IsKnown() bool {
	switch r {
	case TimeOffPolicyListResponseSchedulePerHourWorked, TimeOffPolicyListResponseScheduleMonthly, TimeOffPolicyListResponseScheduleYearly, TimeOffPolicyListResponseScheduleUnlimited:
		return true
	}
	return false
}

type TimeOffPolicyListResponseUnit string

const (
	TimeOffPolicyListResponseUnitHour TimeOffPolicyListResponseUnit = "hour"
	TimeOffPolicyListResponseUnitDay  TimeOffPolicyListResponseUnit = "day"
)

func (r TimeOffPolicyListResponseUnit) IsKnown() bool {
	switch r {
	case TimeOffPolicyListResponseUnitHour, TimeOffPolicyListResponseUnitDay:
		return true
	}
	return false
}

type TimeOffPolicyListParams struct {
	// a string starting with "top\_"
	AfterID param.Field[string] `query:"afterId"`
	// a string starting with "top\_"
	BeforeID param.Field[string] `query:"beforeId"`
	// a number less than or equal to 100
	Limit param.Field[string] `query:"limit"`
}

// URLQuery serializes [TimeOffPolicyListParams]'s query parameters as
// `url.Values`.
func (r TimeOffPolicyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
