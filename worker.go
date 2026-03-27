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

// Endpoints for worker management. "Workers" include anyone employed by your
// company, whether US or international, full-time employees or contractors.
//
// WorkerService contains methods and other services that help with interacting
// with the warp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkerService] method instead.
type WorkerService struct {
	options []option.RequestOption
}

// NewWorkerService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWorkerService(opts ...option.RequestOption) (r WorkerService) {
	r = WorkerService{}
	r.options = opts
	return
}

// Get a specific worker by id.
func (r *WorkerService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *WorkerGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workers/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List all workers. Workers include anyone employed by the company, whether US or
// international, full-time employees or contractors.
func (r *WorkerService) List(ctx context.Context, query WorkerListParams, opts ...option.RequestOption) (res *pagination.CursorPage[WorkerListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/workers"
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

// List all workers. Workers include anyone employed by the company, whether US or
// international, full-time employees or contractors.
func (r *WorkerService) ListAutoPaging(ctx context.Context, query WorkerListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[WorkerListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a worker. Only workers who have not yet completed onboarding can be
// deleted. Active workers must be properly offboarded.
func (r *WorkerService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/workers/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Create a new contractor. The worker will be created in draft status and must be
// invited separately via the invite endpoint. For business contractors, the
// businessName field is required.
func (r *WorkerService) NewContractor(ctx context.Context, body WorkerNewContractorParams, opts ...option.RequestOption) (res *WorkerNewContractorResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/workers/contractor"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Create a new US employee. The worker will be created in draft status and must be
// invited separately via the invite endpoint. If hiring in a state without an
// existing tax registration, you must specify the stateRegistration field.
func (r *WorkerService) NewEmployee(ctx context.Context, body WorkerNewEmployeeParams, opts ...option.RequestOption) (res *WorkerNewEmployeeResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/workers/employee"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Send or resend the worker invite so they can accept and complete onboarding to
// Warp. If the worker has already been invited, the invite will be resent with
// extended validity.
func (r *WorkerService) Invite(ctx context.Context, id string, opts ...option.RequestOption) (res *WorkerInviteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workers/%s/invite", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type WorkerGetResponse struct {
	// The id of the worker.
	ID           string `json:"id" api:"required"`
	BusinessName string `json:"businessName" api:"required"`
	// The department the worker belongs to, or null if unassigned.
	Department WorkerGetResponseDepartment `json:"department" api:"required"`
	// The "ui" name of a worker. If it's a business contractor business name is used.
	// Otherwise we default to preferred name, then first-last.
	DisplayName string `json:"displayName" api:"required"`
	// An email with a reasonably valid regex (shamelessly taken from zod)
	Email string `json:"email" api:"required"`
	// A date string in the form YYYY-MM-DD
	EndDate       string `json:"endDate" api:"required"`
	FirstName     string `json:"firstName" api:"required"`
	IsBusiness    bool   `json:"isBusiness" api:"required"`
	LastName      string `json:"lastName" api:"required"`
	Position      string `json:"position" api:"required"`
	PreferredName string `json:"preferredName" api:"required"`
	// A date string in the form YYYY-MM-DD
	StartDate string `json:"startDate" api:"required"`
	// Any of "draft", "invited", "onboarding", "active", "offboarding", "inactive".
	Status WorkerGetResponseStatus `json:"status" api:"required"`
	// The IANA timezone of the worker (e.g., America/New_York).
	TimeZone string `json:"timeZone" api:"required"`
	// Any of "employee", "contractor".
	Type WorkerGetResponseType `json:"type" api:"required"`
	// An email with a reasonably valid regex (shamelessly taken from zod)
	WorkEmail string `json:"workEmail" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		BusinessName  respjson.Field
		Department    respjson.Field
		DisplayName   respjson.Field
		Email         respjson.Field
		EndDate       respjson.Field
		FirstName     respjson.Field
		IsBusiness    respjson.Field
		LastName      respjson.Field
		Position      respjson.Field
		PreferredName respjson.Field
		StartDate     respjson.Field
		Status        respjson.Field
		TimeZone      respjson.Field
		Type          respjson.Field
		WorkEmail     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkerGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkerGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The department the worker belongs to, or null if unassigned.
type WorkerGetResponseDepartment struct {
	// The unique public id of the department
	ID   string `json:"id" api:"required"`
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkerGetResponseDepartment) RawJSON() string { return r.JSON.raw }
func (r *WorkerGetResponseDepartment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerGetResponseStatus string

const (
	WorkerGetResponseStatusDraft       WorkerGetResponseStatus = "draft"
	WorkerGetResponseStatusInvited     WorkerGetResponseStatus = "invited"
	WorkerGetResponseStatusOnboarding  WorkerGetResponseStatus = "onboarding"
	WorkerGetResponseStatusActive      WorkerGetResponseStatus = "active"
	WorkerGetResponseStatusOffboarding WorkerGetResponseStatus = "offboarding"
	WorkerGetResponseStatusInactive    WorkerGetResponseStatus = "inactive"
)

type WorkerGetResponseType string

const (
	WorkerGetResponseTypeEmployee   WorkerGetResponseType = "employee"
	WorkerGetResponseTypeContractor WorkerGetResponseType = "contractor"
)

type WorkerListResponse struct {
	// The id of the worker.
	ID           string `json:"id" api:"required"`
	BusinessName string `json:"businessName" api:"required"`
	// The department the worker belongs to, or null if unassigned.
	Department WorkerListResponseDepartment `json:"department" api:"required"`
	// The "ui" name of a worker. If it's a business contractor business name is used.
	// Otherwise we default to preferred name, then first-last.
	DisplayName string `json:"displayName" api:"required"`
	// An email with a reasonably valid regex (shamelessly taken from zod)
	Email string `json:"email" api:"required"`
	// A date string in the form YYYY-MM-DD
	EndDate       string `json:"endDate" api:"required"`
	FirstName     string `json:"firstName" api:"required"`
	IsBusiness    bool   `json:"isBusiness" api:"required"`
	LastName      string `json:"lastName" api:"required"`
	Position      string `json:"position" api:"required"`
	PreferredName string `json:"preferredName" api:"required"`
	// A date string in the form YYYY-MM-DD
	StartDate string `json:"startDate" api:"required"`
	// Any of "draft", "invited", "onboarding", "active", "offboarding", "inactive".
	Status WorkerListResponseStatus `json:"status" api:"required"`
	// The IANA timezone of the worker (e.g., America/New_York).
	TimeZone string `json:"timeZone" api:"required"`
	// Any of "employee", "contractor".
	Type WorkerListResponseType `json:"type" api:"required"`
	// An email with a reasonably valid regex (shamelessly taken from zod)
	WorkEmail string `json:"workEmail" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		BusinessName  respjson.Field
		Department    respjson.Field
		DisplayName   respjson.Field
		Email         respjson.Field
		EndDate       respjson.Field
		FirstName     respjson.Field
		IsBusiness    respjson.Field
		LastName      respjson.Field
		Position      respjson.Field
		PreferredName respjson.Field
		StartDate     respjson.Field
		Status        respjson.Field
		TimeZone      respjson.Field
		Type          respjson.Field
		WorkEmail     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkerListResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkerListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The department the worker belongs to, or null if unassigned.
type WorkerListResponseDepartment struct {
	// The unique public id of the department
	ID   string `json:"id" api:"required"`
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkerListResponseDepartment) RawJSON() string { return r.JSON.raw }
func (r *WorkerListResponseDepartment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerListResponseStatus string

const (
	WorkerListResponseStatusDraft       WorkerListResponseStatus = "draft"
	WorkerListResponseStatusInvited     WorkerListResponseStatus = "invited"
	WorkerListResponseStatusOnboarding  WorkerListResponseStatus = "onboarding"
	WorkerListResponseStatusActive      WorkerListResponseStatus = "active"
	WorkerListResponseStatusOffboarding WorkerListResponseStatus = "offboarding"
	WorkerListResponseStatusInactive    WorkerListResponseStatus = "inactive"
)

type WorkerListResponseType string

const (
	WorkerListResponseTypeEmployee   WorkerListResponseType = "employee"
	WorkerListResponseTypeContractor WorkerListResponseType = "contractor"
)

type WorkerNewContractorResponse struct {
	// The id of the worker.
	ID           string `json:"id" api:"required"`
	BusinessName string `json:"businessName" api:"required"`
	// The department the worker belongs to, or null if unassigned.
	Department WorkerNewContractorResponseDepartment `json:"department" api:"required"`
	// The "ui" name of a worker. If it's a business contractor business name is used.
	// Otherwise we default to preferred name, then first-last.
	DisplayName string `json:"displayName" api:"required"`
	// An email with a reasonably valid regex (shamelessly taken from zod)
	Email string `json:"email" api:"required"`
	// A date string in the form YYYY-MM-DD
	EndDate       string `json:"endDate" api:"required"`
	FirstName     string `json:"firstName" api:"required"`
	IsBusiness    bool   `json:"isBusiness" api:"required"`
	LastName      string `json:"lastName" api:"required"`
	Position      string `json:"position" api:"required"`
	PreferredName string `json:"preferredName" api:"required"`
	// A date string in the form YYYY-MM-DD
	StartDate string `json:"startDate" api:"required"`
	// Any of "draft", "invited", "onboarding", "active", "offboarding", "inactive".
	Status WorkerNewContractorResponseStatus `json:"status" api:"required"`
	// The IANA timezone of the worker (e.g., America/New_York).
	TimeZone string `json:"timeZone" api:"required"`
	// Any of "employee", "contractor".
	Type WorkerNewContractorResponseType `json:"type" api:"required"`
	// An email with a reasonably valid regex (shamelessly taken from zod)
	WorkEmail string `json:"workEmail" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		BusinessName  respjson.Field
		Department    respjson.Field
		DisplayName   respjson.Field
		Email         respjson.Field
		EndDate       respjson.Field
		FirstName     respjson.Field
		IsBusiness    respjson.Field
		LastName      respjson.Field
		Position      respjson.Field
		PreferredName respjson.Field
		StartDate     respjson.Field
		Status        respjson.Field
		TimeZone      respjson.Field
		Type          respjson.Field
		WorkEmail     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkerNewContractorResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkerNewContractorResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The department the worker belongs to, or null if unassigned.
type WorkerNewContractorResponseDepartment struct {
	// The unique public id of the department
	ID   string `json:"id" api:"required"`
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkerNewContractorResponseDepartment) RawJSON() string { return r.JSON.raw }
func (r *WorkerNewContractorResponseDepartment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerNewContractorResponseStatus string

const (
	WorkerNewContractorResponseStatusDraft       WorkerNewContractorResponseStatus = "draft"
	WorkerNewContractorResponseStatusInvited     WorkerNewContractorResponseStatus = "invited"
	WorkerNewContractorResponseStatusOnboarding  WorkerNewContractorResponseStatus = "onboarding"
	WorkerNewContractorResponseStatusActive      WorkerNewContractorResponseStatus = "active"
	WorkerNewContractorResponseStatusOffboarding WorkerNewContractorResponseStatus = "offboarding"
	WorkerNewContractorResponseStatusInactive    WorkerNewContractorResponseStatus = "inactive"
)

type WorkerNewContractorResponseType string

const (
	WorkerNewContractorResponseTypeEmployee   WorkerNewContractorResponseType = "employee"
	WorkerNewContractorResponseTypeContractor WorkerNewContractorResponseType = "contractor"
)

type WorkerNewEmployeeResponse struct {
	// The id of the worker.
	ID           string `json:"id" api:"required"`
	BusinessName string `json:"businessName" api:"required"`
	// The department the worker belongs to, or null if unassigned.
	Department WorkerNewEmployeeResponseDepartment `json:"department" api:"required"`
	// The "ui" name of a worker. If it's a business contractor business name is used.
	// Otherwise we default to preferred name, then first-last.
	DisplayName string `json:"displayName" api:"required"`
	// An email with a reasonably valid regex (shamelessly taken from zod)
	Email string `json:"email" api:"required"`
	// A date string in the form YYYY-MM-DD
	EndDate       string `json:"endDate" api:"required"`
	FirstName     string `json:"firstName" api:"required"`
	IsBusiness    bool   `json:"isBusiness" api:"required"`
	LastName      string `json:"lastName" api:"required"`
	Position      string `json:"position" api:"required"`
	PreferredName string `json:"preferredName" api:"required"`
	// A date string in the form YYYY-MM-DD
	StartDate string `json:"startDate" api:"required"`
	// Any of "draft", "invited", "onboarding", "active", "offboarding", "inactive".
	Status WorkerNewEmployeeResponseStatus `json:"status" api:"required"`
	// The IANA timezone of the worker (e.g., America/New_York).
	TimeZone string `json:"timeZone" api:"required"`
	// Any of "employee", "contractor".
	Type WorkerNewEmployeeResponseType `json:"type" api:"required"`
	// An email with a reasonably valid regex (shamelessly taken from zod)
	WorkEmail string `json:"workEmail" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		BusinessName  respjson.Field
		Department    respjson.Field
		DisplayName   respjson.Field
		Email         respjson.Field
		EndDate       respjson.Field
		FirstName     respjson.Field
		IsBusiness    respjson.Field
		LastName      respjson.Field
		Position      respjson.Field
		PreferredName respjson.Field
		StartDate     respjson.Field
		Status        respjson.Field
		TimeZone      respjson.Field
		Type          respjson.Field
		WorkEmail     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkerNewEmployeeResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkerNewEmployeeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The department the worker belongs to, or null if unassigned.
type WorkerNewEmployeeResponseDepartment struct {
	// The unique public id of the department
	ID   string `json:"id" api:"required"`
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkerNewEmployeeResponseDepartment) RawJSON() string { return r.JSON.raw }
func (r *WorkerNewEmployeeResponseDepartment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerNewEmployeeResponseStatus string

const (
	WorkerNewEmployeeResponseStatusDraft       WorkerNewEmployeeResponseStatus = "draft"
	WorkerNewEmployeeResponseStatusInvited     WorkerNewEmployeeResponseStatus = "invited"
	WorkerNewEmployeeResponseStatusOnboarding  WorkerNewEmployeeResponseStatus = "onboarding"
	WorkerNewEmployeeResponseStatusActive      WorkerNewEmployeeResponseStatus = "active"
	WorkerNewEmployeeResponseStatusOffboarding WorkerNewEmployeeResponseStatus = "offboarding"
	WorkerNewEmployeeResponseStatusInactive    WorkerNewEmployeeResponseStatus = "inactive"
)

type WorkerNewEmployeeResponseType string

const (
	WorkerNewEmployeeResponseTypeEmployee   WorkerNewEmployeeResponseType = "employee"
	WorkerNewEmployeeResponseTypeContractor WorkerNewEmployeeResponseType = "contractor"
)

type WorkerInviteResponse struct {
	// The id of the worker.
	ID           string `json:"id" api:"required"`
	BusinessName string `json:"businessName" api:"required"`
	// The department the worker belongs to, or null if unassigned.
	Department WorkerInviteResponseDepartment `json:"department" api:"required"`
	// The "ui" name of a worker. If it's a business contractor business name is used.
	// Otherwise we default to preferred name, then first-last.
	DisplayName string `json:"displayName" api:"required"`
	// An email with a reasonably valid regex (shamelessly taken from zod)
	Email string `json:"email" api:"required"`
	// A date string in the form YYYY-MM-DD
	EndDate       string `json:"endDate" api:"required"`
	FirstName     string `json:"firstName" api:"required"`
	IsBusiness    bool   `json:"isBusiness" api:"required"`
	LastName      string `json:"lastName" api:"required"`
	Position      string `json:"position" api:"required"`
	PreferredName string `json:"preferredName" api:"required"`
	// A date string in the form YYYY-MM-DD
	StartDate string `json:"startDate" api:"required"`
	// Any of "draft", "invited", "onboarding", "active", "offboarding", "inactive".
	Status WorkerInviteResponseStatus `json:"status" api:"required"`
	// The IANA timezone of the worker (e.g., America/New_York).
	TimeZone string `json:"timeZone" api:"required"`
	// Any of "employee", "contractor".
	Type WorkerInviteResponseType `json:"type" api:"required"`
	// An email with a reasonably valid regex (shamelessly taken from zod)
	WorkEmail string `json:"workEmail" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		BusinessName  respjson.Field
		Department    respjson.Field
		DisplayName   respjson.Field
		Email         respjson.Field
		EndDate       respjson.Field
		FirstName     respjson.Field
		IsBusiness    respjson.Field
		LastName      respjson.Field
		Position      respjson.Field
		PreferredName respjson.Field
		StartDate     respjson.Field
		Status        respjson.Field
		TimeZone      respjson.Field
		Type          respjson.Field
		WorkEmail     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkerInviteResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkerInviteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The department the worker belongs to, or null if unassigned.
type WorkerInviteResponseDepartment struct {
	// The unique public id of the department
	ID   string `json:"id" api:"required"`
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkerInviteResponseDepartment) RawJSON() string { return r.JSON.raw }
func (r *WorkerInviteResponseDepartment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerInviteResponseStatus string

const (
	WorkerInviteResponseStatusDraft       WorkerInviteResponseStatus = "draft"
	WorkerInviteResponseStatusInvited     WorkerInviteResponseStatus = "invited"
	WorkerInviteResponseStatusOnboarding  WorkerInviteResponseStatus = "onboarding"
	WorkerInviteResponseStatusActive      WorkerInviteResponseStatus = "active"
	WorkerInviteResponseStatusOffboarding WorkerInviteResponseStatus = "offboarding"
	WorkerInviteResponseStatusInactive    WorkerInviteResponseStatus = "inactive"
)

type WorkerInviteResponseType string

const (
	WorkerInviteResponseTypeEmployee   WorkerInviteResponseType = "employee"
	WorkerInviteResponseTypeContractor WorkerInviteResponseType = "contractor"
)

type WorkerListParams struct {
	// The id of the worker.
	AfterID param.Opt[string] `query:"afterId,omitzero" json:"-"`
	// The id of the worker.
	BeforeID param.Opt[string] `query:"beforeId,omitzero" json:"-"`
	// a number less than or equal to 100
	Limit     param.Opt[string] `query:"limit,omitzero" json:"-"`
	WorkEmail param.Opt[string] `query:"workEmail,omitzero" json:"-"`
	// Any of "draft", "invited", "onboarding", "active", "offboarding", "inactive".
	Statuses []string `query:"statuses,omitzero" json:"-"`
	// Any of "employee", "contractor".
	Types []string `query:"types,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkerListParams]'s query parameters as `url.Values`.
func (r WorkerListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkerNewContractorParams struct {
	// The department to assign this contractor to.
	DepartmentID string `json:"departmentId" api:"required"`
	// Personal email address. The invite will be sent here.
	Email string `json:"email" api:"required"`
	// Whether the contractor is an individual person or a business entity.
	//
	// Any of "individual", "business".
	EntityType WorkerNewContractorParamsEntityType `json:"entityType,omitzero" api:"required"`
	// a non empty string
	FirstName string `json:"firstName" api:"required"`
	// a non empty string
	LastName string `json:"lastName" api:"required"`
	// The worker id of this contractor's direct manager.
	ManagerID string `json:"managerId" api:"required"`
	// The contractor's role or job title.
	Position string `json:"position" api:"required"`
	// A date string in the form YYYY-MM-DD
	StartDate string `json:"startDate" api:"required"`
	// Any of "AD", "AE", "AF", "AG", "AI", "AL", "AM", "AO", "AQ", "AR", "AS", "AT",
	// "AU", "AW", "AX", "AZ", "BA", "BB", "BD", "BE", "BF", "BG", "BH", "BI", "BJ",
	// "BL", "BM", "BN", "BO", "BQ", "BR", "BS", "BT", "BV", "BW", "BY", "BZ", "CA",
	// "CC", "CD", "CF", "CG", "CH", "CI", "CK", "CL", "CM", "CN", "CO", "CR", "CU",
	// "CV", "CW", "CX", "CY", "CZ", "DE", "DJ", "DK", "DM", "DO", "DZ", "EC", "EE",
	// "EG", "EH", "ER", "ES", "ET", "FI", "FJ", "FK", "FM", "FO", "FR", "GA", "GB",
	// "GD", "GE", "GF", "GG", "GH", "GI", "GL", "GM", "GN", "GP", "GQ", "GR", "GS",
	// "GT", "GU", "GW", "GY", "HK", "HM", "HN", "HR", "HT", "HU", "ID", "IE", "IL",
	// "IM", "IN", "IO", "IQ", "IR", "IS", "IT", "JE", "JM", "JO", "JP", "KE", "KG",
	// "KH", "KI", "KM", "KN", "KP", "KR", "KW", "KY", "KZ", "LA", "LB", "LC", "LI",
	// "LK", "LR", "LS", "LT", "LU", "LV", "LY", "MA", "MC", "MD", "ME", "MF", "MG",
	// "MH", "MK", "ML", "MM", "MN", "MO", "MP", "MQ", "MR", "MS", "MT", "MU", "MV",
	// "MW", "MX", "MY", "MZ", "NA", "NC", "NE", "NF", "NG", "NI", "NL", "NO", "NP",
	// "NR", "NU", "NZ", "OM", "PA", "PE", "PF", "PG", "PH", "PK", "PL", "PM", "PN",
	// "PR", "PS", "PT", "PW", "PY", "QA", "RE", "RO", "RS", "RU", "RW", "SA", "SB",
	// "SC", "SD", "SE", "SG", "SH", "SI", "SJ", "SK", "SL", "SM", "SN", "SO", "SR",
	// "SS", "ST", "SV", "SX", "SY", "SZ", "TC", "TD", "TF", "TG", "TH", "TJ", "TK",
	// "TL", "TM", "TN", "TO", "TR", "TT", "TV", "TW", "TZ", "UA", "UG", "UM", "US",
	// "UY", "UZ", "VA", "VC", "VE", "VG", "VI", "VN", "VU", "WF", "WS", "XK", "YE",
	// "YT", "ZA", "ZM", "ZW".
	WorkCountry WorkerNewContractorParamsWorkCountry `json:"workCountry,omitzero" api:"required"`
	// A description of the work the contractor will perform.
	ScopeOfWork param.Opt[string] `json:"scopeOfWork,omitzero"`
	// An email with a reasonably valid regex (shamelessly taken from zod)
	WorkEmail param.Opt[string] `json:"workEmail,omitzero"`
	// Required when entityType is "business". The legal name of the contractor's
	// business.
	BusinessName param.Opt[string] `json:"businessName,omitzero"`
	// The pay rate for the contractor. Leave this blank if you'd like to pay this
	// contractor on-demand or via invoicing.
	Compensation WorkerNewContractorParamsCompensation `json:"compensation,omitzero"`
	// The contractor's pay schedule. Must be a pay schedule that the company has
	// configured.
	//
	// Any of "weekly", "biweekly", "monthly", "semimonthly", "quarterly", "annually".
	PaySchedule WorkerNewContractorParamsPaySchedule `json:"paySchedule,omitzero"`
	paramObj
}

func (r WorkerNewContractorParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkerNewContractorParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkerNewContractorParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the contractor is an individual person or a business entity.
type WorkerNewContractorParamsEntityType string

const (
	WorkerNewContractorParamsEntityTypeIndividual WorkerNewContractorParamsEntityType = "individual"
	WorkerNewContractorParamsEntityTypeBusiness   WorkerNewContractorParamsEntityType = "business"
)

type WorkerNewContractorParamsWorkCountry string

const (
	WorkerNewContractorParamsWorkCountryAd WorkerNewContractorParamsWorkCountry = "AD"
	WorkerNewContractorParamsWorkCountryAe WorkerNewContractorParamsWorkCountry = "AE"
	WorkerNewContractorParamsWorkCountryAf WorkerNewContractorParamsWorkCountry = "AF"
	WorkerNewContractorParamsWorkCountryAg WorkerNewContractorParamsWorkCountry = "AG"
	WorkerNewContractorParamsWorkCountryAI WorkerNewContractorParamsWorkCountry = "AI"
	WorkerNewContractorParamsWorkCountryAl WorkerNewContractorParamsWorkCountry = "AL"
	WorkerNewContractorParamsWorkCountryAm WorkerNewContractorParamsWorkCountry = "AM"
	WorkerNewContractorParamsWorkCountryAo WorkerNewContractorParamsWorkCountry = "AO"
	WorkerNewContractorParamsWorkCountryAq WorkerNewContractorParamsWorkCountry = "AQ"
	WorkerNewContractorParamsWorkCountryAr WorkerNewContractorParamsWorkCountry = "AR"
	WorkerNewContractorParamsWorkCountryAs WorkerNewContractorParamsWorkCountry = "AS"
	WorkerNewContractorParamsWorkCountryAt WorkerNewContractorParamsWorkCountry = "AT"
	WorkerNewContractorParamsWorkCountryAu WorkerNewContractorParamsWorkCountry = "AU"
	WorkerNewContractorParamsWorkCountryAw WorkerNewContractorParamsWorkCountry = "AW"
	WorkerNewContractorParamsWorkCountryAx WorkerNewContractorParamsWorkCountry = "AX"
	WorkerNewContractorParamsWorkCountryAz WorkerNewContractorParamsWorkCountry = "AZ"
	WorkerNewContractorParamsWorkCountryBa WorkerNewContractorParamsWorkCountry = "BA"
	WorkerNewContractorParamsWorkCountryBb WorkerNewContractorParamsWorkCountry = "BB"
	WorkerNewContractorParamsWorkCountryBd WorkerNewContractorParamsWorkCountry = "BD"
	WorkerNewContractorParamsWorkCountryBe WorkerNewContractorParamsWorkCountry = "BE"
	WorkerNewContractorParamsWorkCountryBf WorkerNewContractorParamsWorkCountry = "BF"
	WorkerNewContractorParamsWorkCountryBg WorkerNewContractorParamsWorkCountry = "BG"
	WorkerNewContractorParamsWorkCountryBh WorkerNewContractorParamsWorkCountry = "BH"
	WorkerNewContractorParamsWorkCountryBi WorkerNewContractorParamsWorkCountry = "BI"
	WorkerNewContractorParamsWorkCountryBj WorkerNewContractorParamsWorkCountry = "BJ"
	WorkerNewContractorParamsWorkCountryBl WorkerNewContractorParamsWorkCountry = "BL"
	WorkerNewContractorParamsWorkCountryBm WorkerNewContractorParamsWorkCountry = "BM"
	WorkerNewContractorParamsWorkCountryBn WorkerNewContractorParamsWorkCountry = "BN"
	WorkerNewContractorParamsWorkCountryBo WorkerNewContractorParamsWorkCountry = "BO"
	WorkerNewContractorParamsWorkCountryBq WorkerNewContractorParamsWorkCountry = "BQ"
	WorkerNewContractorParamsWorkCountryBr WorkerNewContractorParamsWorkCountry = "BR"
	WorkerNewContractorParamsWorkCountryBs WorkerNewContractorParamsWorkCountry = "BS"
	WorkerNewContractorParamsWorkCountryBt WorkerNewContractorParamsWorkCountry = "BT"
	WorkerNewContractorParamsWorkCountryBv WorkerNewContractorParamsWorkCountry = "BV"
	WorkerNewContractorParamsWorkCountryBw WorkerNewContractorParamsWorkCountry = "BW"
	WorkerNewContractorParamsWorkCountryBy WorkerNewContractorParamsWorkCountry = "BY"
	WorkerNewContractorParamsWorkCountryBz WorkerNewContractorParamsWorkCountry = "BZ"
	WorkerNewContractorParamsWorkCountryCa WorkerNewContractorParamsWorkCountry = "CA"
	WorkerNewContractorParamsWorkCountryCc WorkerNewContractorParamsWorkCountry = "CC"
	WorkerNewContractorParamsWorkCountryCd WorkerNewContractorParamsWorkCountry = "CD"
	WorkerNewContractorParamsWorkCountryCf WorkerNewContractorParamsWorkCountry = "CF"
	WorkerNewContractorParamsWorkCountryCg WorkerNewContractorParamsWorkCountry = "CG"
	WorkerNewContractorParamsWorkCountryCh WorkerNewContractorParamsWorkCountry = "CH"
	WorkerNewContractorParamsWorkCountryCi WorkerNewContractorParamsWorkCountry = "CI"
	WorkerNewContractorParamsWorkCountryCk WorkerNewContractorParamsWorkCountry = "CK"
	WorkerNewContractorParamsWorkCountryCl WorkerNewContractorParamsWorkCountry = "CL"
	WorkerNewContractorParamsWorkCountryCm WorkerNewContractorParamsWorkCountry = "CM"
	WorkerNewContractorParamsWorkCountryCn WorkerNewContractorParamsWorkCountry = "CN"
	WorkerNewContractorParamsWorkCountryCo WorkerNewContractorParamsWorkCountry = "CO"
	WorkerNewContractorParamsWorkCountryCr WorkerNewContractorParamsWorkCountry = "CR"
	WorkerNewContractorParamsWorkCountryCu WorkerNewContractorParamsWorkCountry = "CU"
	WorkerNewContractorParamsWorkCountryCv WorkerNewContractorParamsWorkCountry = "CV"
	WorkerNewContractorParamsWorkCountryCw WorkerNewContractorParamsWorkCountry = "CW"
	WorkerNewContractorParamsWorkCountryCx WorkerNewContractorParamsWorkCountry = "CX"
	WorkerNewContractorParamsWorkCountryCy WorkerNewContractorParamsWorkCountry = "CY"
	WorkerNewContractorParamsWorkCountryCz WorkerNewContractorParamsWorkCountry = "CZ"
	WorkerNewContractorParamsWorkCountryDe WorkerNewContractorParamsWorkCountry = "DE"
	WorkerNewContractorParamsWorkCountryDj WorkerNewContractorParamsWorkCountry = "DJ"
	WorkerNewContractorParamsWorkCountryDk WorkerNewContractorParamsWorkCountry = "DK"
	WorkerNewContractorParamsWorkCountryDm WorkerNewContractorParamsWorkCountry = "DM"
	WorkerNewContractorParamsWorkCountryDo WorkerNewContractorParamsWorkCountry = "DO"
	WorkerNewContractorParamsWorkCountryDz WorkerNewContractorParamsWorkCountry = "DZ"
	WorkerNewContractorParamsWorkCountryEc WorkerNewContractorParamsWorkCountry = "EC"
	WorkerNewContractorParamsWorkCountryEe WorkerNewContractorParamsWorkCountry = "EE"
	WorkerNewContractorParamsWorkCountryEg WorkerNewContractorParamsWorkCountry = "EG"
	WorkerNewContractorParamsWorkCountryEh WorkerNewContractorParamsWorkCountry = "EH"
	WorkerNewContractorParamsWorkCountryEr WorkerNewContractorParamsWorkCountry = "ER"
	WorkerNewContractorParamsWorkCountryEs WorkerNewContractorParamsWorkCountry = "ES"
	WorkerNewContractorParamsWorkCountryEt WorkerNewContractorParamsWorkCountry = "ET"
	WorkerNewContractorParamsWorkCountryFi WorkerNewContractorParamsWorkCountry = "FI"
	WorkerNewContractorParamsWorkCountryFj WorkerNewContractorParamsWorkCountry = "FJ"
	WorkerNewContractorParamsWorkCountryFk WorkerNewContractorParamsWorkCountry = "FK"
	WorkerNewContractorParamsWorkCountryFm WorkerNewContractorParamsWorkCountry = "FM"
	WorkerNewContractorParamsWorkCountryFo WorkerNewContractorParamsWorkCountry = "FO"
	WorkerNewContractorParamsWorkCountryFr WorkerNewContractorParamsWorkCountry = "FR"
	WorkerNewContractorParamsWorkCountryGa WorkerNewContractorParamsWorkCountry = "GA"
	WorkerNewContractorParamsWorkCountryGB WorkerNewContractorParamsWorkCountry = "GB"
	WorkerNewContractorParamsWorkCountryGd WorkerNewContractorParamsWorkCountry = "GD"
	WorkerNewContractorParamsWorkCountryGe WorkerNewContractorParamsWorkCountry = "GE"
	WorkerNewContractorParamsWorkCountryGf WorkerNewContractorParamsWorkCountry = "GF"
	WorkerNewContractorParamsWorkCountryGg WorkerNewContractorParamsWorkCountry = "GG"
	WorkerNewContractorParamsWorkCountryGh WorkerNewContractorParamsWorkCountry = "GH"
	WorkerNewContractorParamsWorkCountryGi WorkerNewContractorParamsWorkCountry = "GI"
	WorkerNewContractorParamsWorkCountryGl WorkerNewContractorParamsWorkCountry = "GL"
	WorkerNewContractorParamsWorkCountryGm WorkerNewContractorParamsWorkCountry = "GM"
	WorkerNewContractorParamsWorkCountryGn WorkerNewContractorParamsWorkCountry = "GN"
	WorkerNewContractorParamsWorkCountryGp WorkerNewContractorParamsWorkCountry = "GP"
	WorkerNewContractorParamsWorkCountryGq WorkerNewContractorParamsWorkCountry = "GQ"
	WorkerNewContractorParamsWorkCountryGr WorkerNewContractorParamsWorkCountry = "GR"
	WorkerNewContractorParamsWorkCountryGs WorkerNewContractorParamsWorkCountry = "GS"
	WorkerNewContractorParamsWorkCountryGt WorkerNewContractorParamsWorkCountry = "GT"
	WorkerNewContractorParamsWorkCountryGu WorkerNewContractorParamsWorkCountry = "GU"
	WorkerNewContractorParamsWorkCountryGw WorkerNewContractorParamsWorkCountry = "GW"
	WorkerNewContractorParamsWorkCountryGy WorkerNewContractorParamsWorkCountry = "GY"
	WorkerNewContractorParamsWorkCountryHk WorkerNewContractorParamsWorkCountry = "HK"
	WorkerNewContractorParamsWorkCountryHm WorkerNewContractorParamsWorkCountry = "HM"
	WorkerNewContractorParamsWorkCountryHn WorkerNewContractorParamsWorkCountry = "HN"
	WorkerNewContractorParamsWorkCountryHr WorkerNewContractorParamsWorkCountry = "HR"
	WorkerNewContractorParamsWorkCountryHt WorkerNewContractorParamsWorkCountry = "HT"
	WorkerNewContractorParamsWorkCountryHu WorkerNewContractorParamsWorkCountry = "HU"
	WorkerNewContractorParamsWorkCountryID WorkerNewContractorParamsWorkCountry = "ID"
	WorkerNewContractorParamsWorkCountryIe WorkerNewContractorParamsWorkCountry = "IE"
	WorkerNewContractorParamsWorkCountryIl WorkerNewContractorParamsWorkCountry = "IL"
	WorkerNewContractorParamsWorkCountryIm WorkerNewContractorParamsWorkCountry = "IM"
	WorkerNewContractorParamsWorkCountryIn WorkerNewContractorParamsWorkCountry = "IN"
	WorkerNewContractorParamsWorkCountryIo WorkerNewContractorParamsWorkCountry = "IO"
	WorkerNewContractorParamsWorkCountryIq WorkerNewContractorParamsWorkCountry = "IQ"
	WorkerNewContractorParamsWorkCountryIr WorkerNewContractorParamsWorkCountry = "IR"
	WorkerNewContractorParamsWorkCountryIs WorkerNewContractorParamsWorkCountry = "IS"
	WorkerNewContractorParamsWorkCountryIt WorkerNewContractorParamsWorkCountry = "IT"
	WorkerNewContractorParamsWorkCountryJe WorkerNewContractorParamsWorkCountry = "JE"
	WorkerNewContractorParamsWorkCountryJm WorkerNewContractorParamsWorkCountry = "JM"
	WorkerNewContractorParamsWorkCountryJo WorkerNewContractorParamsWorkCountry = "JO"
	WorkerNewContractorParamsWorkCountryJp WorkerNewContractorParamsWorkCountry = "JP"
	WorkerNewContractorParamsWorkCountryKe WorkerNewContractorParamsWorkCountry = "KE"
	WorkerNewContractorParamsWorkCountryKg WorkerNewContractorParamsWorkCountry = "KG"
	WorkerNewContractorParamsWorkCountryKh WorkerNewContractorParamsWorkCountry = "KH"
	WorkerNewContractorParamsWorkCountryKi WorkerNewContractorParamsWorkCountry = "KI"
	WorkerNewContractorParamsWorkCountryKm WorkerNewContractorParamsWorkCountry = "KM"
	WorkerNewContractorParamsWorkCountryKn WorkerNewContractorParamsWorkCountry = "KN"
	WorkerNewContractorParamsWorkCountryKp WorkerNewContractorParamsWorkCountry = "KP"
	WorkerNewContractorParamsWorkCountryKr WorkerNewContractorParamsWorkCountry = "KR"
	WorkerNewContractorParamsWorkCountryKw WorkerNewContractorParamsWorkCountry = "KW"
	WorkerNewContractorParamsWorkCountryKy WorkerNewContractorParamsWorkCountry = "KY"
	WorkerNewContractorParamsWorkCountryKz WorkerNewContractorParamsWorkCountry = "KZ"
	WorkerNewContractorParamsWorkCountryLa WorkerNewContractorParamsWorkCountry = "LA"
	WorkerNewContractorParamsWorkCountryLb WorkerNewContractorParamsWorkCountry = "LB"
	WorkerNewContractorParamsWorkCountryLc WorkerNewContractorParamsWorkCountry = "LC"
	WorkerNewContractorParamsWorkCountryLi WorkerNewContractorParamsWorkCountry = "LI"
	WorkerNewContractorParamsWorkCountryLk WorkerNewContractorParamsWorkCountry = "LK"
	WorkerNewContractorParamsWorkCountryLr WorkerNewContractorParamsWorkCountry = "LR"
	WorkerNewContractorParamsWorkCountryLs WorkerNewContractorParamsWorkCountry = "LS"
	WorkerNewContractorParamsWorkCountryLt WorkerNewContractorParamsWorkCountry = "LT"
	WorkerNewContractorParamsWorkCountryLu WorkerNewContractorParamsWorkCountry = "LU"
	WorkerNewContractorParamsWorkCountryLv WorkerNewContractorParamsWorkCountry = "LV"
	WorkerNewContractorParamsWorkCountryLy WorkerNewContractorParamsWorkCountry = "LY"
	WorkerNewContractorParamsWorkCountryMa WorkerNewContractorParamsWorkCountry = "MA"
	WorkerNewContractorParamsWorkCountryMc WorkerNewContractorParamsWorkCountry = "MC"
	WorkerNewContractorParamsWorkCountryMd WorkerNewContractorParamsWorkCountry = "MD"
	WorkerNewContractorParamsWorkCountryMe WorkerNewContractorParamsWorkCountry = "ME"
	WorkerNewContractorParamsWorkCountryMf WorkerNewContractorParamsWorkCountry = "MF"
	WorkerNewContractorParamsWorkCountryMg WorkerNewContractorParamsWorkCountry = "MG"
	WorkerNewContractorParamsWorkCountryMh WorkerNewContractorParamsWorkCountry = "MH"
	WorkerNewContractorParamsWorkCountryMk WorkerNewContractorParamsWorkCountry = "MK"
	WorkerNewContractorParamsWorkCountryMl WorkerNewContractorParamsWorkCountry = "ML"
	WorkerNewContractorParamsWorkCountryMm WorkerNewContractorParamsWorkCountry = "MM"
	WorkerNewContractorParamsWorkCountryMn WorkerNewContractorParamsWorkCountry = "MN"
	WorkerNewContractorParamsWorkCountryMo WorkerNewContractorParamsWorkCountry = "MO"
	WorkerNewContractorParamsWorkCountryMp WorkerNewContractorParamsWorkCountry = "MP"
	WorkerNewContractorParamsWorkCountryMq WorkerNewContractorParamsWorkCountry = "MQ"
	WorkerNewContractorParamsWorkCountryMr WorkerNewContractorParamsWorkCountry = "MR"
	WorkerNewContractorParamsWorkCountryMs WorkerNewContractorParamsWorkCountry = "MS"
	WorkerNewContractorParamsWorkCountryMt WorkerNewContractorParamsWorkCountry = "MT"
	WorkerNewContractorParamsWorkCountryMu WorkerNewContractorParamsWorkCountry = "MU"
	WorkerNewContractorParamsWorkCountryMv WorkerNewContractorParamsWorkCountry = "MV"
	WorkerNewContractorParamsWorkCountryMw WorkerNewContractorParamsWorkCountry = "MW"
	WorkerNewContractorParamsWorkCountryMx WorkerNewContractorParamsWorkCountry = "MX"
	WorkerNewContractorParamsWorkCountryMy WorkerNewContractorParamsWorkCountry = "MY"
	WorkerNewContractorParamsWorkCountryMz WorkerNewContractorParamsWorkCountry = "MZ"
	WorkerNewContractorParamsWorkCountryNa WorkerNewContractorParamsWorkCountry = "NA"
	WorkerNewContractorParamsWorkCountryNc WorkerNewContractorParamsWorkCountry = "NC"
	WorkerNewContractorParamsWorkCountryNe WorkerNewContractorParamsWorkCountry = "NE"
	WorkerNewContractorParamsWorkCountryNf WorkerNewContractorParamsWorkCountry = "NF"
	WorkerNewContractorParamsWorkCountryNg WorkerNewContractorParamsWorkCountry = "NG"
	WorkerNewContractorParamsWorkCountryNi WorkerNewContractorParamsWorkCountry = "NI"
	WorkerNewContractorParamsWorkCountryNl WorkerNewContractorParamsWorkCountry = "NL"
	WorkerNewContractorParamsWorkCountryNo WorkerNewContractorParamsWorkCountry = "NO"
	WorkerNewContractorParamsWorkCountryNp WorkerNewContractorParamsWorkCountry = "NP"
	WorkerNewContractorParamsWorkCountryNr WorkerNewContractorParamsWorkCountry = "NR"
	WorkerNewContractorParamsWorkCountryNu WorkerNewContractorParamsWorkCountry = "NU"
	WorkerNewContractorParamsWorkCountryNz WorkerNewContractorParamsWorkCountry = "NZ"
	WorkerNewContractorParamsWorkCountryOm WorkerNewContractorParamsWorkCountry = "OM"
	WorkerNewContractorParamsWorkCountryPa WorkerNewContractorParamsWorkCountry = "PA"
	WorkerNewContractorParamsWorkCountryPe WorkerNewContractorParamsWorkCountry = "PE"
	WorkerNewContractorParamsWorkCountryPf WorkerNewContractorParamsWorkCountry = "PF"
	WorkerNewContractorParamsWorkCountryPg WorkerNewContractorParamsWorkCountry = "PG"
	WorkerNewContractorParamsWorkCountryPh WorkerNewContractorParamsWorkCountry = "PH"
	WorkerNewContractorParamsWorkCountryPk WorkerNewContractorParamsWorkCountry = "PK"
	WorkerNewContractorParamsWorkCountryPl WorkerNewContractorParamsWorkCountry = "PL"
	WorkerNewContractorParamsWorkCountryPm WorkerNewContractorParamsWorkCountry = "PM"
	WorkerNewContractorParamsWorkCountryPn WorkerNewContractorParamsWorkCountry = "PN"
	WorkerNewContractorParamsWorkCountryPr WorkerNewContractorParamsWorkCountry = "PR"
	WorkerNewContractorParamsWorkCountryPs WorkerNewContractorParamsWorkCountry = "PS"
	WorkerNewContractorParamsWorkCountryPt WorkerNewContractorParamsWorkCountry = "PT"
	WorkerNewContractorParamsWorkCountryPw WorkerNewContractorParamsWorkCountry = "PW"
	WorkerNewContractorParamsWorkCountryPy WorkerNewContractorParamsWorkCountry = "PY"
	WorkerNewContractorParamsWorkCountryQa WorkerNewContractorParamsWorkCountry = "QA"
	WorkerNewContractorParamsWorkCountryRe WorkerNewContractorParamsWorkCountry = "RE"
	WorkerNewContractorParamsWorkCountryRo WorkerNewContractorParamsWorkCountry = "RO"
	WorkerNewContractorParamsWorkCountryRs WorkerNewContractorParamsWorkCountry = "RS"
	WorkerNewContractorParamsWorkCountryRu WorkerNewContractorParamsWorkCountry = "RU"
	WorkerNewContractorParamsWorkCountryRw WorkerNewContractorParamsWorkCountry = "RW"
	WorkerNewContractorParamsWorkCountrySa WorkerNewContractorParamsWorkCountry = "SA"
	WorkerNewContractorParamsWorkCountrySb WorkerNewContractorParamsWorkCountry = "SB"
	WorkerNewContractorParamsWorkCountrySc WorkerNewContractorParamsWorkCountry = "SC"
	WorkerNewContractorParamsWorkCountrySd WorkerNewContractorParamsWorkCountry = "SD"
	WorkerNewContractorParamsWorkCountrySe WorkerNewContractorParamsWorkCountry = "SE"
	WorkerNewContractorParamsWorkCountrySg WorkerNewContractorParamsWorkCountry = "SG"
	WorkerNewContractorParamsWorkCountrySh WorkerNewContractorParamsWorkCountry = "SH"
	WorkerNewContractorParamsWorkCountrySi WorkerNewContractorParamsWorkCountry = "SI"
	WorkerNewContractorParamsWorkCountrySj WorkerNewContractorParamsWorkCountry = "SJ"
	WorkerNewContractorParamsWorkCountrySk WorkerNewContractorParamsWorkCountry = "SK"
	WorkerNewContractorParamsWorkCountrySl WorkerNewContractorParamsWorkCountry = "SL"
	WorkerNewContractorParamsWorkCountrySm WorkerNewContractorParamsWorkCountry = "SM"
	WorkerNewContractorParamsWorkCountrySn WorkerNewContractorParamsWorkCountry = "SN"
	WorkerNewContractorParamsWorkCountrySo WorkerNewContractorParamsWorkCountry = "SO"
	WorkerNewContractorParamsWorkCountrySr WorkerNewContractorParamsWorkCountry = "SR"
	WorkerNewContractorParamsWorkCountrySS WorkerNewContractorParamsWorkCountry = "SS"
	WorkerNewContractorParamsWorkCountrySt WorkerNewContractorParamsWorkCountry = "ST"
	WorkerNewContractorParamsWorkCountrySv WorkerNewContractorParamsWorkCountry = "SV"
	WorkerNewContractorParamsWorkCountrySx WorkerNewContractorParamsWorkCountry = "SX"
	WorkerNewContractorParamsWorkCountrySy WorkerNewContractorParamsWorkCountry = "SY"
	WorkerNewContractorParamsWorkCountrySz WorkerNewContractorParamsWorkCountry = "SZ"
	WorkerNewContractorParamsWorkCountryTc WorkerNewContractorParamsWorkCountry = "TC"
	WorkerNewContractorParamsWorkCountryTd WorkerNewContractorParamsWorkCountry = "TD"
	WorkerNewContractorParamsWorkCountryTf WorkerNewContractorParamsWorkCountry = "TF"
	WorkerNewContractorParamsWorkCountryTg WorkerNewContractorParamsWorkCountry = "TG"
	WorkerNewContractorParamsWorkCountryTh WorkerNewContractorParamsWorkCountry = "TH"
	WorkerNewContractorParamsWorkCountryTj WorkerNewContractorParamsWorkCountry = "TJ"
	WorkerNewContractorParamsWorkCountryTk WorkerNewContractorParamsWorkCountry = "TK"
	WorkerNewContractorParamsWorkCountryTl WorkerNewContractorParamsWorkCountry = "TL"
	WorkerNewContractorParamsWorkCountryTm WorkerNewContractorParamsWorkCountry = "TM"
	WorkerNewContractorParamsWorkCountryTn WorkerNewContractorParamsWorkCountry = "TN"
	WorkerNewContractorParamsWorkCountryTo WorkerNewContractorParamsWorkCountry = "TO"
	WorkerNewContractorParamsWorkCountryTr WorkerNewContractorParamsWorkCountry = "TR"
	WorkerNewContractorParamsWorkCountryTt WorkerNewContractorParamsWorkCountry = "TT"
	WorkerNewContractorParamsWorkCountryTv WorkerNewContractorParamsWorkCountry = "TV"
	WorkerNewContractorParamsWorkCountryTw WorkerNewContractorParamsWorkCountry = "TW"
	WorkerNewContractorParamsWorkCountryTz WorkerNewContractorParamsWorkCountry = "TZ"
	WorkerNewContractorParamsWorkCountryUa WorkerNewContractorParamsWorkCountry = "UA"
	WorkerNewContractorParamsWorkCountryUg WorkerNewContractorParamsWorkCountry = "UG"
	WorkerNewContractorParamsWorkCountryUm WorkerNewContractorParamsWorkCountry = "UM"
	WorkerNewContractorParamsWorkCountryUs WorkerNewContractorParamsWorkCountry = "US"
	WorkerNewContractorParamsWorkCountryUy WorkerNewContractorParamsWorkCountry = "UY"
	WorkerNewContractorParamsWorkCountryUz WorkerNewContractorParamsWorkCountry = "UZ"
	WorkerNewContractorParamsWorkCountryVa WorkerNewContractorParamsWorkCountry = "VA"
	WorkerNewContractorParamsWorkCountryVc WorkerNewContractorParamsWorkCountry = "VC"
	WorkerNewContractorParamsWorkCountryVe WorkerNewContractorParamsWorkCountry = "VE"
	WorkerNewContractorParamsWorkCountryVg WorkerNewContractorParamsWorkCountry = "VG"
	WorkerNewContractorParamsWorkCountryVi WorkerNewContractorParamsWorkCountry = "VI"
	WorkerNewContractorParamsWorkCountryVn WorkerNewContractorParamsWorkCountry = "VN"
	WorkerNewContractorParamsWorkCountryVu WorkerNewContractorParamsWorkCountry = "VU"
	WorkerNewContractorParamsWorkCountryWf WorkerNewContractorParamsWorkCountry = "WF"
	WorkerNewContractorParamsWorkCountryWs WorkerNewContractorParamsWorkCountry = "WS"
	WorkerNewContractorParamsWorkCountryXk WorkerNewContractorParamsWorkCountry = "XK"
	WorkerNewContractorParamsWorkCountryYe WorkerNewContractorParamsWorkCountry = "YE"
	WorkerNewContractorParamsWorkCountryYt WorkerNewContractorParamsWorkCountry = "YT"
	WorkerNewContractorParamsWorkCountryZa WorkerNewContractorParamsWorkCountry = "ZA"
	WorkerNewContractorParamsWorkCountryZm WorkerNewContractorParamsWorkCountry = "ZM"
	WorkerNewContractorParamsWorkCountryZw WorkerNewContractorParamsWorkCountry = "ZW"
)

// The pay rate for the contractor. Leave this blank if you'd like to pay this
// contractor on-demand or via invoicing.
//
// The properties Amount, Currency, Per are required.
type WorkerNewContractorParamsCompensation struct {
	// a positive number
	Amount float64 `json:"amount" api:"required"`
	// Any of "USD", "AUD", "BGN", "BRL", "CAD", "CHF", "CZK", "DKK", "EUR", "GBP",
	// "HKD", "HUF", "IDR", "INR", "JPY", "MYR", "NOK", "NZD", "CNY", "PLN", "RON",
	// "TRY", "SEK", "SGD", "AED", "ARS", "BDT", "BWP", "CLP", "COP", "CRC", "EGP",
	// "FJD", "GEL", "GHS", "ILS", "KES", "KRW", "LKR", "MAD", "MXN", "NPR", "PHP",
	// "PKR", "THB", "UAH", "UGX", "UYU", "VND", "ZAR", "ZMW", "TND", "NGN", "RSD",
	// "TWD", "GTQ", "HNL", "DOP", "SAR", "XAF", "PEN".
	Currency string `json:"currency,omitzero" api:"required"`
	// The pay period for the compensation amount.
	//
	// Any of "hour", "year", "month", "week".
	Per string `json:"per,omitzero" api:"required"`
	paramObj
}

func (r WorkerNewContractorParamsCompensation) MarshalJSON() (data []byte, err error) {
	type shadow WorkerNewContractorParamsCompensation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkerNewContractorParamsCompensation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkerNewContractorParamsCompensation](
		"currency", "USD", "AUD", "BGN", "BRL", "CAD", "CHF", "CZK", "DKK", "EUR", "GBP", "HKD", "HUF", "IDR", "INR", "JPY", "MYR", "NOK", "NZD", "CNY", "PLN", "RON", "TRY", "SEK", "SGD", "AED", "ARS", "BDT", "BWP", "CLP", "COP", "CRC", "EGP", "FJD", "GEL", "GHS", "ILS", "KES", "KRW", "LKR", "MAD", "MXN", "NPR", "PHP", "PKR", "THB", "UAH", "UGX", "UYU", "VND", "ZAR", "ZMW", "TND", "NGN", "RSD", "TWD", "GTQ", "HNL", "DOP", "SAR", "XAF", "PEN",
	)
	apijson.RegisterFieldValidator[WorkerNewContractorParamsCompensation](
		"per", "hour", "year", "month", "week",
	)
}

// The contractor's pay schedule. Must be a pay schedule that the company has
// configured.
type WorkerNewContractorParamsPaySchedule string

const (
	WorkerNewContractorParamsPayScheduleWeekly      WorkerNewContractorParamsPaySchedule = "weekly"
	WorkerNewContractorParamsPayScheduleBiweekly    WorkerNewContractorParamsPaySchedule = "biweekly"
	WorkerNewContractorParamsPayScheduleMonthly     WorkerNewContractorParamsPaySchedule = "monthly"
	WorkerNewContractorParamsPayScheduleSemimonthly WorkerNewContractorParamsPaySchedule = "semimonthly"
	WorkerNewContractorParamsPayScheduleQuarterly   WorkerNewContractorParamsPaySchedule = "quarterly"
	WorkerNewContractorParamsPayScheduleAnnually    WorkerNewContractorParamsPaySchedule = "annually"
)

type WorkerNewEmployeeParams struct {
	// The employee's base compensation.
	Compensation WorkerNewEmployeeParamsCompensation `json:"compensation,omitzero" api:"required"`
	// The department to assign this employee to.
	DepartmentID string `json:"departmentId" api:"required"`
	// Personal email address. The invite will be sent here.
	Email string `json:"email" api:"required"`
	// a non empty string
	FirstName string `json:"firstName" api:"required"`
	// a non empty string
	LastName string `json:"lastName" api:"required"`
	// The worker id of this employee's direct manager.
	ManagerID string `json:"managerId" api:"required"`
	// The employee's job title.
	Position string `json:"position" api:"required"`
	// A date string in the form YYYY-MM-DD
	StartDate string `json:"startDate" api:"required"`
	// Where the employee will work. Either an existing company workplace or a remote
	// US state.
	WorkLocation WorkerNewEmployeeParamsWorkLocationUnion `json:"workLocation,omitzero" api:"required"`
	// a non-negative number
	StockOptions param.Opt[float64] `json:"stockOptions,omitzero"`
	// An email with a reasonably valid regex (shamelessly taken from zod)
	WorkEmail param.Opt[string] `json:"workEmail,omitzero"`
	// Whether the employee is required to complete I-9 work authorization. Set to
	// false if the employee has already been verified off-platform. Defaults to true.
	RequireI9 param.Opt[bool] `json:"requireI9,omitzero"`
	// The employee's pay schedule. Must be a pay schedule that the company has
	// configured.
	//
	// Any of "weekly", "biweekly", "monthly", "semimonthly", "quarterly", "annually".
	PaySchedule WorkerNewEmployeeParamsPaySchedule `json:"paySchedule,omitzero"`
	// How state tax registration is handled for this employee's work state. Required
	// when hiring in a state where your company doesn't have an existing registration.
	// Use 'self_managed' if you've already registered in this state, or 'warp_managed'
	// for Warp to handle registration on your behalf.
	//
	// Any of "self_managed", "warp_managed".
	StateRegistration WorkerNewEmployeeParamsStateRegistration `json:"stateRegistration,omitzero"`
	paramObj
}

func (r WorkerNewEmployeeParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkerNewEmployeeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkerNewEmployeeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The employee's base compensation.
//
// The properties Amount, Per are required.
type WorkerNewEmployeeParamsCompensation struct {
	// a positive number
	Amount float64 `json:"amount" api:"required"`
	// Whether the amount is per hour or per year.
	//
	// Any of "hour", "year".
	Per string `json:"per,omitzero" api:"required"`
	paramObj
}

func (r WorkerNewEmployeeParamsCompensation) MarshalJSON() (data []byte, err error) {
	type shadow WorkerNewEmployeeParamsCompensation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkerNewEmployeeParamsCompensation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkerNewEmployeeParamsCompensation](
		"per", "hour", "year",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WorkerNewEmployeeParamsWorkLocationUnion struct {
	OfWorkerNewEmployeesWorkLocationOfficeWorkLocation *WorkerNewEmployeeParamsWorkLocationOfficeWorkLocation `json:",omitzero,inline"`
	OfWorkerNewEmployeesWorkLocationRemoteWorkLocation *WorkerNewEmployeeParamsWorkLocationRemoteWorkLocation `json:",omitzero,inline"`
	paramUnion
}

func (u WorkerNewEmployeeParamsWorkLocationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfWorkerNewEmployeesWorkLocationOfficeWorkLocation, u.OfWorkerNewEmployeesWorkLocationRemoteWorkLocation)
}
func (u *WorkerNewEmployeeParamsWorkLocationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Employee works from a company workplace.
//
// The properties Type, WorkplaceID are required.
type WorkerNewEmployeeParamsWorkLocationOfficeWorkLocation struct {
	// Any of "office".
	Type string `json:"type,omitzero" api:"required"`
	// Public workplace identifier
	WorkplaceID string `json:"workplaceId" api:"required"`
	paramObj
}

func (r WorkerNewEmployeeParamsWorkLocationOfficeWorkLocation) MarshalJSON() (data []byte, err error) {
	type shadow WorkerNewEmployeeParamsWorkLocationOfficeWorkLocation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkerNewEmployeeParamsWorkLocationOfficeWorkLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkerNewEmployeeParamsWorkLocationOfficeWorkLocation](
		"type", "office",
	)
}

// Employee works remotely from a US state.
//
// The properties State, Type are required.
type WorkerNewEmployeeParamsWorkLocationRemoteWorkLocation struct {
	// The US state where the remote employee works. Required for tax purposes.
	//
	// Any of "AL", "AK", "AZ", "AR", "CA", "CO", "CT", "DC", "DE", "FL", "GA", "HI",
	// "ID", "IL", "IN", "IA", "KS", "KY", "LA", "ME", "MD", "MA", "MI", "MN", "MS",
	// "MO", "MT", "NE", "NV", "NH", "NJ", "NM", "NY", "NC", "ND", "OH", "OK", "OR",
	// "PA", "RI", "SC", "SD", "TN", "TX", "UT", "VT", "VA", "WA", "WV", "WI", "WY".
	State string `json:"state,omitzero" api:"required"`
	// Any of "remote".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r WorkerNewEmployeeParamsWorkLocationRemoteWorkLocation) MarshalJSON() (data []byte, err error) {
	type shadow WorkerNewEmployeeParamsWorkLocationRemoteWorkLocation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkerNewEmployeeParamsWorkLocationRemoteWorkLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkerNewEmployeeParamsWorkLocationRemoteWorkLocation](
		"state", "AL", "AK", "AZ", "AR", "CA", "CO", "CT", "DC", "DE", "FL", "GA", "HI", "ID", "IL", "IN", "IA", "KS", "KY", "LA", "ME", "MD", "MA", "MI", "MN", "MS", "MO", "MT", "NE", "NV", "NH", "NJ", "NM", "NY", "NC", "ND", "OH", "OK", "OR", "PA", "RI", "SC", "SD", "TN", "TX", "UT", "VT", "VA", "WA", "WV", "WI", "WY",
	)
	apijson.RegisterFieldValidator[WorkerNewEmployeeParamsWorkLocationRemoteWorkLocation](
		"type", "remote",
	)
}

// The employee's pay schedule. Must be a pay schedule that the company has
// configured.
type WorkerNewEmployeeParamsPaySchedule string

const (
	WorkerNewEmployeeParamsPayScheduleWeekly      WorkerNewEmployeeParamsPaySchedule = "weekly"
	WorkerNewEmployeeParamsPayScheduleBiweekly    WorkerNewEmployeeParamsPaySchedule = "biweekly"
	WorkerNewEmployeeParamsPayScheduleMonthly     WorkerNewEmployeeParamsPaySchedule = "monthly"
	WorkerNewEmployeeParamsPayScheduleSemimonthly WorkerNewEmployeeParamsPaySchedule = "semimonthly"
	WorkerNewEmployeeParamsPayScheduleQuarterly   WorkerNewEmployeeParamsPaySchedule = "quarterly"
	WorkerNewEmployeeParamsPayScheduleAnnually    WorkerNewEmployeeParamsPaySchedule = "annually"
)

// How state tax registration is handled for this employee's work state. Required
// when hiring in a state where your company doesn't have an existing registration.
// Use 'self_managed' if you've already registered in this state, or 'warp_managed'
// for Warp to handle registration on your behalf.
type WorkerNewEmployeeParamsStateRegistration string

const (
	WorkerNewEmployeeParamsStateRegistrationSelfManaged WorkerNewEmployeeParamsStateRegistration = "self_managed"
	WorkerNewEmployeeParamsStateRegistrationWarpManaged WorkerNewEmployeeParamsStateRegistration = "warp_managed"
)
