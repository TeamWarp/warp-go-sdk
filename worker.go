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
	Options []option.RequestOption
}

// NewWorkerService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWorkerService(opts ...option.RequestOption) (r *WorkerService) {
	r = &WorkerService{}
	r.Options = opts
	return
}

// Get a specific worker by id.
func (r *WorkerService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *WorkerGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List all workers. Workers include anyone employed by the company, whether US or
// international, full-time employees or contractors.
func (r *WorkerService) List(ctx context.Context, query WorkerListParams, opts ...option.RequestOption) (res *pagination.CursorPage[WorkerListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/workers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Create a new contractor. The worker will be created in draft status and must be
// invited separately via the invite endpoint. For business contractors, the
// businessName field is required.
func (r *WorkerService) NewContractor(ctx context.Context, body WorkerNewContractorParams, opts ...option.RequestOption) (res *WorkerNewContractorResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/workers/contractor"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Create a new US employee. The worker will be created in draft status and must be
// invited separately via the invite endpoint. If hiring in a state without an
// existing tax registration, you must specify the stateRegistration field.
func (r *WorkerService) NewEmployee(ctx context.Context, body WorkerNewEmployeeParams, opts ...option.RequestOption) (res *WorkerNewEmployeeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/workers/employee"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Send or resend the worker invite so they can accept and complete onboarding to
// Warp. If the worker has already been invited, the invite will be resent with
// extended validity.
func (r *WorkerService) Invite(ctx context.Context, id string, opts ...option.RequestOption) (res *WorkerInviteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workers/%s/invite", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type WorkerGetResponse struct {
	// The id of the worker.
	ID           string `json:"id" api:"required"`
	BusinessName string `json:"businessName" api:"required,nullable"`
	// The department the worker belongs to, or null if unassigned.
	Department WorkerGetResponseDepartment `json:"department" api:"required,nullable"`
	// The "ui" name of a worker. If it's a business contractor business name is used.
	// Otherwise we default to preferred name, then first-last.
	DisplayName string `json:"displayName" api:"required"`
	// An email with a reasonably valid regex (shamelessly taken from zod)
	Email string `json:"email" api:"required"`
	// A date string in the form YYYY-MM-DD
	EndDate       string `json:"endDate" api:"required,nullable"`
	FirstName     string `json:"firstName" api:"required"`
	IsBusiness    bool   `json:"isBusiness" api:"required,nullable"`
	LastName      string `json:"lastName" api:"required"`
	Position      string `json:"position" api:"required"`
	PreferredName string `json:"preferredName" api:"required,nullable"`
	// A date string in the form YYYY-MM-DD
	StartDate string                  `json:"startDate" api:"required"`
	Status    WorkerGetResponseStatus `json:"status" api:"required"`
	// The IANA timezone of the worker (e.g., America/New_York).
	TimeZone string                `json:"timeZone" api:"required,nullable"`
	Type     WorkerGetResponseType `json:"type" api:"required"`
	// An email with a reasonably valid regex (shamelessly taken from zod)
	WorkEmail string                `json:"workEmail" api:"required,nullable"`
	JSON      workerGetResponseJSON `json:"-"`
}

// workerGetResponseJSON contains the JSON metadata for the struct
// [WorkerGetResponse]
type workerGetResponseJSON struct {
	ID            apijson.Field
	BusinessName  apijson.Field
	Department    apijson.Field
	DisplayName   apijson.Field
	Email         apijson.Field
	EndDate       apijson.Field
	FirstName     apijson.Field
	IsBusiness    apijson.Field
	LastName      apijson.Field
	Position      apijson.Field
	PreferredName apijson.Field
	StartDate     apijson.Field
	Status        apijson.Field
	TimeZone      apijson.Field
	Type          apijson.Field
	WorkEmail     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkerGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerGetResponseJSON) RawJSON() string {
	return r.raw
}

// The department the worker belongs to, or null if unassigned.
type WorkerGetResponseDepartment struct {
	// The unique public id of the department
	ID   string                          `json:"id" api:"required"`
	Name string                          `json:"name" api:"required"`
	JSON workerGetResponseDepartmentJSON `json:"-"`
}

// workerGetResponseDepartmentJSON contains the JSON metadata for the struct
// [WorkerGetResponseDepartment]
type workerGetResponseDepartmentJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerGetResponseDepartment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerGetResponseDepartmentJSON) RawJSON() string {
	return r.raw
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

func (r WorkerGetResponseStatus) IsKnown() bool {
	switch r {
	case WorkerGetResponseStatusDraft, WorkerGetResponseStatusInvited, WorkerGetResponseStatusOnboarding, WorkerGetResponseStatusActive, WorkerGetResponseStatusOffboarding, WorkerGetResponseStatusInactive:
		return true
	}
	return false
}

type WorkerGetResponseType string

const (
	WorkerGetResponseTypeEmployee   WorkerGetResponseType = "employee"
	WorkerGetResponseTypeContractor WorkerGetResponseType = "contractor"
)

func (r WorkerGetResponseType) IsKnown() bool {
	switch r {
	case WorkerGetResponseTypeEmployee, WorkerGetResponseTypeContractor:
		return true
	}
	return false
}

type WorkerListResponse struct {
	// The id of the worker.
	ID           string `json:"id" api:"required"`
	BusinessName string `json:"businessName" api:"required,nullable"`
	// The department the worker belongs to, or null if unassigned.
	Department WorkerListResponseDepartment `json:"department" api:"required,nullable"`
	// The "ui" name of a worker. If it's a business contractor business name is used.
	// Otherwise we default to preferred name, then first-last.
	DisplayName string `json:"displayName" api:"required"`
	// An email with a reasonably valid regex (shamelessly taken from zod)
	Email string `json:"email" api:"required"`
	// A date string in the form YYYY-MM-DD
	EndDate       string `json:"endDate" api:"required,nullable"`
	FirstName     string `json:"firstName" api:"required"`
	IsBusiness    bool   `json:"isBusiness" api:"required,nullable"`
	LastName      string `json:"lastName" api:"required"`
	Position      string `json:"position" api:"required"`
	PreferredName string `json:"preferredName" api:"required,nullable"`
	// A date string in the form YYYY-MM-DD
	StartDate string                   `json:"startDate" api:"required"`
	Status    WorkerListResponseStatus `json:"status" api:"required"`
	// The IANA timezone of the worker (e.g., America/New_York).
	TimeZone string                 `json:"timeZone" api:"required,nullable"`
	Type     WorkerListResponseType `json:"type" api:"required"`
	// An email with a reasonably valid regex (shamelessly taken from zod)
	WorkEmail string                 `json:"workEmail" api:"required,nullable"`
	JSON      workerListResponseJSON `json:"-"`
}

// workerListResponseJSON contains the JSON metadata for the struct
// [WorkerListResponse]
type workerListResponseJSON struct {
	ID            apijson.Field
	BusinessName  apijson.Field
	Department    apijson.Field
	DisplayName   apijson.Field
	Email         apijson.Field
	EndDate       apijson.Field
	FirstName     apijson.Field
	IsBusiness    apijson.Field
	LastName      apijson.Field
	Position      apijson.Field
	PreferredName apijson.Field
	StartDate     apijson.Field
	Status        apijson.Field
	TimeZone      apijson.Field
	Type          apijson.Field
	WorkEmail     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkerListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerListResponseJSON) RawJSON() string {
	return r.raw
}

// The department the worker belongs to, or null if unassigned.
type WorkerListResponseDepartment struct {
	// The unique public id of the department
	ID   string                           `json:"id" api:"required"`
	Name string                           `json:"name" api:"required"`
	JSON workerListResponseDepartmentJSON `json:"-"`
}

// workerListResponseDepartmentJSON contains the JSON metadata for the struct
// [WorkerListResponseDepartment]
type workerListResponseDepartmentJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerListResponseDepartment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerListResponseDepartmentJSON) RawJSON() string {
	return r.raw
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

func (r WorkerListResponseStatus) IsKnown() bool {
	switch r {
	case WorkerListResponseStatusDraft, WorkerListResponseStatusInvited, WorkerListResponseStatusOnboarding, WorkerListResponseStatusActive, WorkerListResponseStatusOffboarding, WorkerListResponseStatusInactive:
		return true
	}
	return false
}

type WorkerListResponseType string

const (
	WorkerListResponseTypeEmployee   WorkerListResponseType = "employee"
	WorkerListResponseTypeContractor WorkerListResponseType = "contractor"
)

func (r WorkerListResponseType) IsKnown() bool {
	switch r {
	case WorkerListResponseTypeEmployee, WorkerListResponseTypeContractor:
		return true
	}
	return false
}

type WorkerNewContractorResponse struct {
	// The id of the worker.
	ID           string `json:"id" api:"required"`
	BusinessName string `json:"businessName" api:"required,nullable"`
	// The department the worker belongs to, or null if unassigned.
	Department WorkerNewContractorResponseDepartment `json:"department" api:"required,nullable"`
	// The "ui" name of a worker. If it's a business contractor business name is used.
	// Otherwise we default to preferred name, then first-last.
	DisplayName string `json:"displayName" api:"required"`
	// An email with a reasonably valid regex (shamelessly taken from zod)
	Email string `json:"email" api:"required"`
	// A date string in the form YYYY-MM-DD
	EndDate       string `json:"endDate" api:"required,nullable"`
	FirstName     string `json:"firstName" api:"required"`
	IsBusiness    bool   `json:"isBusiness" api:"required,nullable"`
	LastName      string `json:"lastName" api:"required"`
	Position      string `json:"position" api:"required"`
	PreferredName string `json:"preferredName" api:"required,nullable"`
	// A date string in the form YYYY-MM-DD
	StartDate string                            `json:"startDate" api:"required"`
	Status    WorkerNewContractorResponseStatus `json:"status" api:"required"`
	// The IANA timezone of the worker (e.g., America/New_York).
	TimeZone string                          `json:"timeZone" api:"required,nullable"`
	Type     WorkerNewContractorResponseType `json:"type" api:"required"`
	// An email with a reasonably valid regex (shamelessly taken from zod)
	WorkEmail string                          `json:"workEmail" api:"required,nullable"`
	JSON      workerNewContractorResponseJSON `json:"-"`
}

// workerNewContractorResponseJSON contains the JSON metadata for the struct
// [WorkerNewContractorResponse]
type workerNewContractorResponseJSON struct {
	ID            apijson.Field
	BusinessName  apijson.Field
	Department    apijson.Field
	DisplayName   apijson.Field
	Email         apijson.Field
	EndDate       apijson.Field
	FirstName     apijson.Field
	IsBusiness    apijson.Field
	LastName      apijson.Field
	Position      apijson.Field
	PreferredName apijson.Field
	StartDate     apijson.Field
	Status        apijson.Field
	TimeZone      apijson.Field
	Type          apijson.Field
	WorkEmail     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkerNewContractorResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerNewContractorResponseJSON) RawJSON() string {
	return r.raw
}

// The department the worker belongs to, or null if unassigned.
type WorkerNewContractorResponseDepartment struct {
	// The unique public id of the department
	ID   string                                    `json:"id" api:"required"`
	Name string                                    `json:"name" api:"required"`
	JSON workerNewContractorResponseDepartmentJSON `json:"-"`
}

// workerNewContractorResponseDepartmentJSON contains the JSON metadata for the
// struct [WorkerNewContractorResponseDepartment]
type workerNewContractorResponseDepartmentJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerNewContractorResponseDepartment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerNewContractorResponseDepartmentJSON) RawJSON() string {
	return r.raw
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

func (r WorkerNewContractorResponseStatus) IsKnown() bool {
	switch r {
	case WorkerNewContractorResponseStatusDraft, WorkerNewContractorResponseStatusInvited, WorkerNewContractorResponseStatusOnboarding, WorkerNewContractorResponseStatusActive, WorkerNewContractorResponseStatusOffboarding, WorkerNewContractorResponseStatusInactive:
		return true
	}
	return false
}

type WorkerNewContractorResponseType string

const (
	WorkerNewContractorResponseTypeEmployee   WorkerNewContractorResponseType = "employee"
	WorkerNewContractorResponseTypeContractor WorkerNewContractorResponseType = "contractor"
)

func (r WorkerNewContractorResponseType) IsKnown() bool {
	switch r {
	case WorkerNewContractorResponseTypeEmployee, WorkerNewContractorResponseTypeContractor:
		return true
	}
	return false
}

type WorkerNewEmployeeResponse struct {
	// The id of the worker.
	ID           string `json:"id" api:"required"`
	BusinessName string `json:"businessName" api:"required,nullable"`
	// The department the worker belongs to, or null if unassigned.
	Department WorkerNewEmployeeResponseDepartment `json:"department" api:"required,nullable"`
	// The "ui" name of a worker. If it's a business contractor business name is used.
	// Otherwise we default to preferred name, then first-last.
	DisplayName string `json:"displayName" api:"required"`
	// An email with a reasonably valid regex (shamelessly taken from zod)
	Email string `json:"email" api:"required"`
	// A date string in the form YYYY-MM-DD
	EndDate       string `json:"endDate" api:"required,nullable"`
	FirstName     string `json:"firstName" api:"required"`
	IsBusiness    bool   `json:"isBusiness" api:"required,nullable"`
	LastName      string `json:"lastName" api:"required"`
	Position      string `json:"position" api:"required"`
	PreferredName string `json:"preferredName" api:"required,nullable"`
	// A date string in the form YYYY-MM-DD
	StartDate string                          `json:"startDate" api:"required"`
	Status    WorkerNewEmployeeResponseStatus `json:"status" api:"required"`
	// The IANA timezone of the worker (e.g., America/New_York).
	TimeZone string                        `json:"timeZone" api:"required,nullable"`
	Type     WorkerNewEmployeeResponseType `json:"type" api:"required"`
	// An email with a reasonably valid regex (shamelessly taken from zod)
	WorkEmail string                        `json:"workEmail" api:"required,nullable"`
	JSON      workerNewEmployeeResponseJSON `json:"-"`
}

// workerNewEmployeeResponseJSON contains the JSON metadata for the struct
// [WorkerNewEmployeeResponse]
type workerNewEmployeeResponseJSON struct {
	ID            apijson.Field
	BusinessName  apijson.Field
	Department    apijson.Field
	DisplayName   apijson.Field
	Email         apijson.Field
	EndDate       apijson.Field
	FirstName     apijson.Field
	IsBusiness    apijson.Field
	LastName      apijson.Field
	Position      apijson.Field
	PreferredName apijson.Field
	StartDate     apijson.Field
	Status        apijson.Field
	TimeZone      apijson.Field
	Type          apijson.Field
	WorkEmail     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkerNewEmployeeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerNewEmployeeResponseJSON) RawJSON() string {
	return r.raw
}

// The department the worker belongs to, or null if unassigned.
type WorkerNewEmployeeResponseDepartment struct {
	// The unique public id of the department
	ID   string                                  `json:"id" api:"required"`
	Name string                                  `json:"name" api:"required"`
	JSON workerNewEmployeeResponseDepartmentJSON `json:"-"`
}

// workerNewEmployeeResponseDepartmentJSON contains the JSON metadata for the
// struct [WorkerNewEmployeeResponseDepartment]
type workerNewEmployeeResponseDepartmentJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerNewEmployeeResponseDepartment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerNewEmployeeResponseDepartmentJSON) RawJSON() string {
	return r.raw
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

func (r WorkerNewEmployeeResponseStatus) IsKnown() bool {
	switch r {
	case WorkerNewEmployeeResponseStatusDraft, WorkerNewEmployeeResponseStatusInvited, WorkerNewEmployeeResponseStatusOnboarding, WorkerNewEmployeeResponseStatusActive, WorkerNewEmployeeResponseStatusOffboarding, WorkerNewEmployeeResponseStatusInactive:
		return true
	}
	return false
}

type WorkerNewEmployeeResponseType string

const (
	WorkerNewEmployeeResponseTypeEmployee   WorkerNewEmployeeResponseType = "employee"
	WorkerNewEmployeeResponseTypeContractor WorkerNewEmployeeResponseType = "contractor"
)

func (r WorkerNewEmployeeResponseType) IsKnown() bool {
	switch r {
	case WorkerNewEmployeeResponseTypeEmployee, WorkerNewEmployeeResponseTypeContractor:
		return true
	}
	return false
}

type WorkerInviteResponse struct {
	// The id of the worker.
	ID           string `json:"id" api:"required"`
	BusinessName string `json:"businessName" api:"required,nullable"`
	// The department the worker belongs to, or null if unassigned.
	Department WorkerInviteResponseDepartment `json:"department" api:"required,nullable"`
	// The "ui" name of a worker. If it's a business contractor business name is used.
	// Otherwise we default to preferred name, then first-last.
	DisplayName string `json:"displayName" api:"required"`
	// An email with a reasonably valid regex (shamelessly taken from zod)
	Email string `json:"email" api:"required"`
	// A date string in the form YYYY-MM-DD
	EndDate       string `json:"endDate" api:"required,nullable"`
	FirstName     string `json:"firstName" api:"required"`
	IsBusiness    bool   `json:"isBusiness" api:"required,nullable"`
	LastName      string `json:"lastName" api:"required"`
	Position      string `json:"position" api:"required"`
	PreferredName string `json:"preferredName" api:"required,nullable"`
	// A date string in the form YYYY-MM-DD
	StartDate string                     `json:"startDate" api:"required"`
	Status    WorkerInviteResponseStatus `json:"status" api:"required"`
	// The IANA timezone of the worker (e.g., America/New_York).
	TimeZone string                   `json:"timeZone" api:"required,nullable"`
	Type     WorkerInviteResponseType `json:"type" api:"required"`
	// An email with a reasonably valid regex (shamelessly taken from zod)
	WorkEmail string                   `json:"workEmail" api:"required,nullable"`
	JSON      workerInviteResponseJSON `json:"-"`
}

// workerInviteResponseJSON contains the JSON metadata for the struct
// [WorkerInviteResponse]
type workerInviteResponseJSON struct {
	ID            apijson.Field
	BusinessName  apijson.Field
	Department    apijson.Field
	DisplayName   apijson.Field
	Email         apijson.Field
	EndDate       apijson.Field
	FirstName     apijson.Field
	IsBusiness    apijson.Field
	LastName      apijson.Field
	Position      apijson.Field
	PreferredName apijson.Field
	StartDate     apijson.Field
	Status        apijson.Field
	TimeZone      apijson.Field
	Type          apijson.Field
	WorkEmail     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkerInviteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerInviteResponseJSON) RawJSON() string {
	return r.raw
}

// The department the worker belongs to, or null if unassigned.
type WorkerInviteResponseDepartment struct {
	// The unique public id of the department
	ID   string                             `json:"id" api:"required"`
	Name string                             `json:"name" api:"required"`
	JSON workerInviteResponseDepartmentJSON `json:"-"`
}

// workerInviteResponseDepartmentJSON contains the JSON metadata for the struct
// [WorkerInviteResponseDepartment]
type workerInviteResponseDepartmentJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerInviteResponseDepartment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerInviteResponseDepartmentJSON) RawJSON() string {
	return r.raw
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

func (r WorkerInviteResponseStatus) IsKnown() bool {
	switch r {
	case WorkerInviteResponseStatusDraft, WorkerInviteResponseStatusInvited, WorkerInviteResponseStatusOnboarding, WorkerInviteResponseStatusActive, WorkerInviteResponseStatusOffboarding, WorkerInviteResponseStatusInactive:
		return true
	}
	return false
}

type WorkerInviteResponseType string

const (
	WorkerInviteResponseTypeEmployee   WorkerInviteResponseType = "employee"
	WorkerInviteResponseTypeContractor WorkerInviteResponseType = "contractor"
)

func (r WorkerInviteResponseType) IsKnown() bool {
	switch r {
	case WorkerInviteResponseTypeEmployee, WorkerInviteResponseTypeContractor:
		return true
	}
	return false
}

type WorkerListParams struct {
	// The id of the worker.
	AfterID param.Field[string] `query:"afterId"`
	// The id of the worker.
	BeforeID param.Field[string] `query:"beforeId"`
	// a number less than or equal to 100
	Limit     param.Field[string]                   `query:"limit"`
	Statuses  param.Field[[]WorkerListParamsStatus] `query:"statuses"`
	Types     param.Field[[]WorkerListParamsType]   `query:"types"`
	WorkEmail param.Field[string]                   `query:"workEmail"`
}

// URLQuery serializes [WorkerListParams]'s query parameters as `url.Values`.
func (r WorkerListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkerListParamsStatus string

const (
	WorkerListParamsStatusDraft       WorkerListParamsStatus = "draft"
	WorkerListParamsStatusInvited     WorkerListParamsStatus = "invited"
	WorkerListParamsStatusOnboarding  WorkerListParamsStatus = "onboarding"
	WorkerListParamsStatusActive      WorkerListParamsStatus = "active"
	WorkerListParamsStatusOffboarding WorkerListParamsStatus = "offboarding"
	WorkerListParamsStatusInactive    WorkerListParamsStatus = "inactive"
)

func (r WorkerListParamsStatus) IsKnown() bool {
	switch r {
	case WorkerListParamsStatusDraft, WorkerListParamsStatusInvited, WorkerListParamsStatusOnboarding, WorkerListParamsStatusActive, WorkerListParamsStatusOffboarding, WorkerListParamsStatusInactive:
		return true
	}
	return false
}

type WorkerListParamsType string

const (
	WorkerListParamsTypeEmployee   WorkerListParamsType = "employee"
	WorkerListParamsTypeContractor WorkerListParamsType = "contractor"
)

func (r WorkerListParamsType) IsKnown() bool {
	switch r {
	case WorkerListParamsTypeEmployee, WorkerListParamsTypeContractor:
		return true
	}
	return false
}

type WorkerNewContractorParams struct {
	// The department to assign this contractor to.
	DepartmentID param.Field[string] `json:"departmentId" api:"required"`
	// Personal email address. The invite will be sent here.
	Email param.Field[string] `json:"email" api:"required"`
	// Whether the contractor is an individual person or a business entity.
	EntityType param.Field[WorkerNewContractorParamsEntityType] `json:"entityType" api:"required"`
	// a non empty string
	FirstName param.Field[string] `json:"firstName" api:"required"`
	// a non empty string
	LastName param.Field[string] `json:"lastName" api:"required"`
	// The worker id of this contractor's direct manager.
	ManagerID param.Field[string] `json:"managerId" api:"required"`
	// The contractor's role or job title.
	Position param.Field[string] `json:"position" api:"required"`
	// A date string in the form YYYY-MM-DD
	StartDate   param.Field[string]                               `json:"startDate" api:"required"`
	WorkCountry param.Field[WorkerNewContractorParamsWorkCountry] `json:"workCountry" api:"required"`
	// Required when entityType is "business". The legal name of the contractor's
	// business.
	BusinessName param.Field[string] `json:"businessName"`
	// The pay rate for the contractor. Leave this blank if you'd like to pay this
	// contractor on-demand or via invoicing.
	Compensation param.Field[WorkerNewContractorParamsCompensation] `json:"compensation"`
	// The contractor's pay schedule. Must be a pay schedule that the company has
	// configured.
	PaySchedule param.Field[WorkerNewContractorParamsPaySchedule] `json:"paySchedule"`
	// A description of the work the contractor will perform.
	ScopeOfWork param.Field[string] `json:"scopeOfWork"`
	// An email with a reasonably valid regex (shamelessly taken from zod)
	WorkEmail param.Field[string] `json:"workEmail"`
}

func (r WorkerNewContractorParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the contractor is an individual person or a business entity.
type WorkerNewContractorParamsEntityType string

const (
	WorkerNewContractorParamsEntityTypeIndividual WorkerNewContractorParamsEntityType = "individual"
	WorkerNewContractorParamsEntityTypeBusiness   WorkerNewContractorParamsEntityType = "business"
)

func (r WorkerNewContractorParamsEntityType) IsKnown() bool {
	switch r {
	case WorkerNewContractorParamsEntityTypeIndividual, WorkerNewContractorParamsEntityTypeBusiness:
		return true
	}
	return false
}

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

func (r WorkerNewContractorParamsWorkCountry) IsKnown() bool {
	switch r {
	case WorkerNewContractorParamsWorkCountryAd, WorkerNewContractorParamsWorkCountryAe, WorkerNewContractorParamsWorkCountryAf, WorkerNewContractorParamsWorkCountryAg, WorkerNewContractorParamsWorkCountryAI, WorkerNewContractorParamsWorkCountryAl, WorkerNewContractorParamsWorkCountryAm, WorkerNewContractorParamsWorkCountryAo, WorkerNewContractorParamsWorkCountryAq, WorkerNewContractorParamsWorkCountryAr, WorkerNewContractorParamsWorkCountryAs, WorkerNewContractorParamsWorkCountryAt, WorkerNewContractorParamsWorkCountryAu, WorkerNewContractorParamsWorkCountryAw, WorkerNewContractorParamsWorkCountryAx, WorkerNewContractorParamsWorkCountryAz, WorkerNewContractorParamsWorkCountryBa, WorkerNewContractorParamsWorkCountryBb, WorkerNewContractorParamsWorkCountryBd, WorkerNewContractorParamsWorkCountryBe, WorkerNewContractorParamsWorkCountryBf, WorkerNewContractorParamsWorkCountryBg, WorkerNewContractorParamsWorkCountryBh, WorkerNewContractorParamsWorkCountryBi, WorkerNewContractorParamsWorkCountryBj, WorkerNewContractorParamsWorkCountryBl, WorkerNewContractorParamsWorkCountryBm, WorkerNewContractorParamsWorkCountryBn, WorkerNewContractorParamsWorkCountryBo, WorkerNewContractorParamsWorkCountryBq, WorkerNewContractorParamsWorkCountryBr, WorkerNewContractorParamsWorkCountryBs, WorkerNewContractorParamsWorkCountryBt, WorkerNewContractorParamsWorkCountryBv, WorkerNewContractorParamsWorkCountryBw, WorkerNewContractorParamsWorkCountryBy, WorkerNewContractorParamsWorkCountryBz, WorkerNewContractorParamsWorkCountryCa, WorkerNewContractorParamsWorkCountryCc, WorkerNewContractorParamsWorkCountryCd, WorkerNewContractorParamsWorkCountryCf, WorkerNewContractorParamsWorkCountryCg, WorkerNewContractorParamsWorkCountryCh, WorkerNewContractorParamsWorkCountryCi, WorkerNewContractorParamsWorkCountryCk, WorkerNewContractorParamsWorkCountryCl, WorkerNewContractorParamsWorkCountryCm, WorkerNewContractorParamsWorkCountryCn, WorkerNewContractorParamsWorkCountryCo, WorkerNewContractorParamsWorkCountryCr, WorkerNewContractorParamsWorkCountryCu, WorkerNewContractorParamsWorkCountryCv, WorkerNewContractorParamsWorkCountryCw, WorkerNewContractorParamsWorkCountryCx, WorkerNewContractorParamsWorkCountryCy, WorkerNewContractorParamsWorkCountryCz, WorkerNewContractorParamsWorkCountryDe, WorkerNewContractorParamsWorkCountryDj, WorkerNewContractorParamsWorkCountryDk, WorkerNewContractorParamsWorkCountryDm, WorkerNewContractorParamsWorkCountryDo, WorkerNewContractorParamsWorkCountryDz, WorkerNewContractorParamsWorkCountryEc, WorkerNewContractorParamsWorkCountryEe, WorkerNewContractorParamsWorkCountryEg, WorkerNewContractorParamsWorkCountryEh, WorkerNewContractorParamsWorkCountryEr, WorkerNewContractorParamsWorkCountryEs, WorkerNewContractorParamsWorkCountryEt, WorkerNewContractorParamsWorkCountryFi, WorkerNewContractorParamsWorkCountryFj, WorkerNewContractorParamsWorkCountryFk, WorkerNewContractorParamsWorkCountryFm, WorkerNewContractorParamsWorkCountryFo, WorkerNewContractorParamsWorkCountryFr, WorkerNewContractorParamsWorkCountryGa, WorkerNewContractorParamsWorkCountryGB, WorkerNewContractorParamsWorkCountryGd, WorkerNewContractorParamsWorkCountryGe, WorkerNewContractorParamsWorkCountryGf, WorkerNewContractorParamsWorkCountryGg, WorkerNewContractorParamsWorkCountryGh, WorkerNewContractorParamsWorkCountryGi, WorkerNewContractorParamsWorkCountryGl, WorkerNewContractorParamsWorkCountryGm, WorkerNewContractorParamsWorkCountryGn, WorkerNewContractorParamsWorkCountryGp, WorkerNewContractorParamsWorkCountryGq, WorkerNewContractorParamsWorkCountryGr, WorkerNewContractorParamsWorkCountryGs, WorkerNewContractorParamsWorkCountryGt, WorkerNewContractorParamsWorkCountryGu, WorkerNewContractorParamsWorkCountryGw, WorkerNewContractorParamsWorkCountryGy, WorkerNewContractorParamsWorkCountryHk, WorkerNewContractorParamsWorkCountryHm, WorkerNewContractorParamsWorkCountryHn, WorkerNewContractorParamsWorkCountryHr, WorkerNewContractorParamsWorkCountryHt, WorkerNewContractorParamsWorkCountryHu, WorkerNewContractorParamsWorkCountryID, WorkerNewContractorParamsWorkCountryIe, WorkerNewContractorParamsWorkCountryIl, WorkerNewContractorParamsWorkCountryIm, WorkerNewContractorParamsWorkCountryIn, WorkerNewContractorParamsWorkCountryIo, WorkerNewContractorParamsWorkCountryIq, WorkerNewContractorParamsWorkCountryIr, WorkerNewContractorParamsWorkCountryIs, WorkerNewContractorParamsWorkCountryIt, WorkerNewContractorParamsWorkCountryJe, WorkerNewContractorParamsWorkCountryJm, WorkerNewContractorParamsWorkCountryJo, WorkerNewContractorParamsWorkCountryJp, WorkerNewContractorParamsWorkCountryKe, WorkerNewContractorParamsWorkCountryKg, WorkerNewContractorParamsWorkCountryKh, WorkerNewContractorParamsWorkCountryKi, WorkerNewContractorParamsWorkCountryKm, WorkerNewContractorParamsWorkCountryKn, WorkerNewContractorParamsWorkCountryKp, WorkerNewContractorParamsWorkCountryKr, WorkerNewContractorParamsWorkCountryKw, WorkerNewContractorParamsWorkCountryKy, WorkerNewContractorParamsWorkCountryKz, WorkerNewContractorParamsWorkCountryLa, WorkerNewContractorParamsWorkCountryLb, WorkerNewContractorParamsWorkCountryLc, WorkerNewContractorParamsWorkCountryLi, WorkerNewContractorParamsWorkCountryLk, WorkerNewContractorParamsWorkCountryLr, WorkerNewContractorParamsWorkCountryLs, WorkerNewContractorParamsWorkCountryLt, WorkerNewContractorParamsWorkCountryLu, WorkerNewContractorParamsWorkCountryLv, WorkerNewContractorParamsWorkCountryLy, WorkerNewContractorParamsWorkCountryMa, WorkerNewContractorParamsWorkCountryMc, WorkerNewContractorParamsWorkCountryMd, WorkerNewContractorParamsWorkCountryMe, WorkerNewContractorParamsWorkCountryMf, WorkerNewContractorParamsWorkCountryMg, WorkerNewContractorParamsWorkCountryMh, WorkerNewContractorParamsWorkCountryMk, WorkerNewContractorParamsWorkCountryMl, WorkerNewContractorParamsWorkCountryMm, WorkerNewContractorParamsWorkCountryMn, WorkerNewContractorParamsWorkCountryMo, WorkerNewContractorParamsWorkCountryMp, WorkerNewContractorParamsWorkCountryMq, WorkerNewContractorParamsWorkCountryMr, WorkerNewContractorParamsWorkCountryMs, WorkerNewContractorParamsWorkCountryMt, WorkerNewContractorParamsWorkCountryMu, WorkerNewContractorParamsWorkCountryMv, WorkerNewContractorParamsWorkCountryMw, WorkerNewContractorParamsWorkCountryMx, WorkerNewContractorParamsWorkCountryMy, WorkerNewContractorParamsWorkCountryMz, WorkerNewContractorParamsWorkCountryNa, WorkerNewContractorParamsWorkCountryNc, WorkerNewContractorParamsWorkCountryNe, WorkerNewContractorParamsWorkCountryNf, WorkerNewContractorParamsWorkCountryNg, WorkerNewContractorParamsWorkCountryNi, WorkerNewContractorParamsWorkCountryNl, WorkerNewContractorParamsWorkCountryNo, WorkerNewContractorParamsWorkCountryNp, WorkerNewContractorParamsWorkCountryNr, WorkerNewContractorParamsWorkCountryNu, WorkerNewContractorParamsWorkCountryNz, WorkerNewContractorParamsWorkCountryOm, WorkerNewContractorParamsWorkCountryPa, WorkerNewContractorParamsWorkCountryPe, WorkerNewContractorParamsWorkCountryPf, WorkerNewContractorParamsWorkCountryPg, WorkerNewContractorParamsWorkCountryPh, WorkerNewContractorParamsWorkCountryPk, WorkerNewContractorParamsWorkCountryPl, WorkerNewContractorParamsWorkCountryPm, WorkerNewContractorParamsWorkCountryPn, WorkerNewContractorParamsWorkCountryPr, WorkerNewContractorParamsWorkCountryPs, WorkerNewContractorParamsWorkCountryPt, WorkerNewContractorParamsWorkCountryPw, WorkerNewContractorParamsWorkCountryPy, WorkerNewContractorParamsWorkCountryQa, WorkerNewContractorParamsWorkCountryRe, WorkerNewContractorParamsWorkCountryRo, WorkerNewContractorParamsWorkCountryRs, WorkerNewContractorParamsWorkCountryRu, WorkerNewContractorParamsWorkCountryRw, WorkerNewContractorParamsWorkCountrySa, WorkerNewContractorParamsWorkCountrySb, WorkerNewContractorParamsWorkCountrySc, WorkerNewContractorParamsWorkCountrySd, WorkerNewContractorParamsWorkCountrySe, WorkerNewContractorParamsWorkCountrySg, WorkerNewContractorParamsWorkCountrySh, WorkerNewContractorParamsWorkCountrySi, WorkerNewContractorParamsWorkCountrySj, WorkerNewContractorParamsWorkCountrySk, WorkerNewContractorParamsWorkCountrySl, WorkerNewContractorParamsWorkCountrySm, WorkerNewContractorParamsWorkCountrySn, WorkerNewContractorParamsWorkCountrySo, WorkerNewContractorParamsWorkCountrySr, WorkerNewContractorParamsWorkCountrySS, WorkerNewContractorParamsWorkCountrySt, WorkerNewContractorParamsWorkCountrySv, WorkerNewContractorParamsWorkCountrySx, WorkerNewContractorParamsWorkCountrySy, WorkerNewContractorParamsWorkCountrySz, WorkerNewContractorParamsWorkCountryTc, WorkerNewContractorParamsWorkCountryTd, WorkerNewContractorParamsWorkCountryTf, WorkerNewContractorParamsWorkCountryTg, WorkerNewContractorParamsWorkCountryTh, WorkerNewContractorParamsWorkCountryTj, WorkerNewContractorParamsWorkCountryTk, WorkerNewContractorParamsWorkCountryTl, WorkerNewContractorParamsWorkCountryTm, WorkerNewContractorParamsWorkCountryTn, WorkerNewContractorParamsWorkCountryTo, WorkerNewContractorParamsWorkCountryTr, WorkerNewContractorParamsWorkCountryTt, WorkerNewContractorParamsWorkCountryTv, WorkerNewContractorParamsWorkCountryTw, WorkerNewContractorParamsWorkCountryTz, WorkerNewContractorParamsWorkCountryUa, WorkerNewContractorParamsWorkCountryUg, WorkerNewContractorParamsWorkCountryUm, WorkerNewContractorParamsWorkCountryUs, WorkerNewContractorParamsWorkCountryUy, WorkerNewContractorParamsWorkCountryUz, WorkerNewContractorParamsWorkCountryVa, WorkerNewContractorParamsWorkCountryVc, WorkerNewContractorParamsWorkCountryVe, WorkerNewContractorParamsWorkCountryVg, WorkerNewContractorParamsWorkCountryVi, WorkerNewContractorParamsWorkCountryVn, WorkerNewContractorParamsWorkCountryVu, WorkerNewContractorParamsWorkCountryWf, WorkerNewContractorParamsWorkCountryWs, WorkerNewContractorParamsWorkCountryXk, WorkerNewContractorParamsWorkCountryYe, WorkerNewContractorParamsWorkCountryYt, WorkerNewContractorParamsWorkCountryZa, WorkerNewContractorParamsWorkCountryZm, WorkerNewContractorParamsWorkCountryZw:
		return true
	}
	return false
}

// The pay rate for the contractor. Leave this blank if you'd like to pay this
// contractor on-demand or via invoicing.
type WorkerNewContractorParamsCompensation struct {
	// a positive number
	Amount   param.Field[float64]                                       `json:"amount" api:"required"`
	Currency param.Field[WorkerNewContractorParamsCompensationCurrency] `json:"currency" api:"required"`
	// The pay period for the compensation amount.
	Per param.Field[WorkerNewContractorParamsCompensationPer] `json:"per" api:"required"`
}

func (r WorkerNewContractorParamsCompensation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerNewContractorParamsCompensationCurrency string

const (
	WorkerNewContractorParamsCompensationCurrencyUsd WorkerNewContractorParamsCompensationCurrency = "USD"
	WorkerNewContractorParamsCompensationCurrencyAud WorkerNewContractorParamsCompensationCurrency = "AUD"
	WorkerNewContractorParamsCompensationCurrencyBgn WorkerNewContractorParamsCompensationCurrency = "BGN"
	WorkerNewContractorParamsCompensationCurrencyBrl WorkerNewContractorParamsCompensationCurrency = "BRL"
	WorkerNewContractorParamsCompensationCurrencyCad WorkerNewContractorParamsCompensationCurrency = "CAD"
	WorkerNewContractorParamsCompensationCurrencyChf WorkerNewContractorParamsCompensationCurrency = "CHF"
	WorkerNewContractorParamsCompensationCurrencyCzk WorkerNewContractorParamsCompensationCurrency = "CZK"
	WorkerNewContractorParamsCompensationCurrencyDkk WorkerNewContractorParamsCompensationCurrency = "DKK"
	WorkerNewContractorParamsCompensationCurrencyEur WorkerNewContractorParamsCompensationCurrency = "EUR"
	WorkerNewContractorParamsCompensationCurrencyGbp WorkerNewContractorParamsCompensationCurrency = "GBP"
	WorkerNewContractorParamsCompensationCurrencyHkd WorkerNewContractorParamsCompensationCurrency = "HKD"
	WorkerNewContractorParamsCompensationCurrencyHuf WorkerNewContractorParamsCompensationCurrency = "HUF"
	WorkerNewContractorParamsCompensationCurrencyIdr WorkerNewContractorParamsCompensationCurrency = "IDR"
	WorkerNewContractorParamsCompensationCurrencyInr WorkerNewContractorParamsCompensationCurrency = "INR"
	WorkerNewContractorParamsCompensationCurrencyJpy WorkerNewContractorParamsCompensationCurrency = "JPY"
	WorkerNewContractorParamsCompensationCurrencyMyr WorkerNewContractorParamsCompensationCurrency = "MYR"
	WorkerNewContractorParamsCompensationCurrencyNok WorkerNewContractorParamsCompensationCurrency = "NOK"
	WorkerNewContractorParamsCompensationCurrencyNzd WorkerNewContractorParamsCompensationCurrency = "NZD"
	WorkerNewContractorParamsCompensationCurrencyCny WorkerNewContractorParamsCompensationCurrency = "CNY"
	WorkerNewContractorParamsCompensationCurrencyPln WorkerNewContractorParamsCompensationCurrency = "PLN"
	WorkerNewContractorParamsCompensationCurrencyRon WorkerNewContractorParamsCompensationCurrency = "RON"
	WorkerNewContractorParamsCompensationCurrencyTry WorkerNewContractorParamsCompensationCurrency = "TRY"
	WorkerNewContractorParamsCompensationCurrencySek WorkerNewContractorParamsCompensationCurrency = "SEK"
	WorkerNewContractorParamsCompensationCurrencySgd WorkerNewContractorParamsCompensationCurrency = "SGD"
	WorkerNewContractorParamsCompensationCurrencyAed WorkerNewContractorParamsCompensationCurrency = "AED"
	WorkerNewContractorParamsCompensationCurrencyArs WorkerNewContractorParamsCompensationCurrency = "ARS"
	WorkerNewContractorParamsCompensationCurrencyBdt WorkerNewContractorParamsCompensationCurrency = "BDT"
	WorkerNewContractorParamsCompensationCurrencyBwp WorkerNewContractorParamsCompensationCurrency = "BWP"
	WorkerNewContractorParamsCompensationCurrencyClp WorkerNewContractorParamsCompensationCurrency = "CLP"
	WorkerNewContractorParamsCompensationCurrencyCop WorkerNewContractorParamsCompensationCurrency = "COP"
	WorkerNewContractorParamsCompensationCurrencyCrc WorkerNewContractorParamsCompensationCurrency = "CRC"
	WorkerNewContractorParamsCompensationCurrencyEgp WorkerNewContractorParamsCompensationCurrency = "EGP"
	WorkerNewContractorParamsCompensationCurrencyFjd WorkerNewContractorParamsCompensationCurrency = "FJD"
	WorkerNewContractorParamsCompensationCurrencyGel WorkerNewContractorParamsCompensationCurrency = "GEL"
	WorkerNewContractorParamsCompensationCurrencyGhs WorkerNewContractorParamsCompensationCurrency = "GHS"
	WorkerNewContractorParamsCompensationCurrencyIls WorkerNewContractorParamsCompensationCurrency = "ILS"
	WorkerNewContractorParamsCompensationCurrencyKes WorkerNewContractorParamsCompensationCurrency = "KES"
	WorkerNewContractorParamsCompensationCurrencyKrw WorkerNewContractorParamsCompensationCurrency = "KRW"
	WorkerNewContractorParamsCompensationCurrencyLkr WorkerNewContractorParamsCompensationCurrency = "LKR"
	WorkerNewContractorParamsCompensationCurrencyMad WorkerNewContractorParamsCompensationCurrency = "MAD"
	WorkerNewContractorParamsCompensationCurrencyMxn WorkerNewContractorParamsCompensationCurrency = "MXN"
	WorkerNewContractorParamsCompensationCurrencyNpr WorkerNewContractorParamsCompensationCurrency = "NPR"
	WorkerNewContractorParamsCompensationCurrencyPhp WorkerNewContractorParamsCompensationCurrency = "PHP"
	WorkerNewContractorParamsCompensationCurrencyPkr WorkerNewContractorParamsCompensationCurrency = "PKR"
	WorkerNewContractorParamsCompensationCurrencyThb WorkerNewContractorParamsCompensationCurrency = "THB"
	WorkerNewContractorParamsCompensationCurrencyUah WorkerNewContractorParamsCompensationCurrency = "UAH"
	WorkerNewContractorParamsCompensationCurrencyUgx WorkerNewContractorParamsCompensationCurrency = "UGX"
	WorkerNewContractorParamsCompensationCurrencyUyu WorkerNewContractorParamsCompensationCurrency = "UYU"
	WorkerNewContractorParamsCompensationCurrencyVnd WorkerNewContractorParamsCompensationCurrency = "VND"
	WorkerNewContractorParamsCompensationCurrencyZar WorkerNewContractorParamsCompensationCurrency = "ZAR"
	WorkerNewContractorParamsCompensationCurrencyZmw WorkerNewContractorParamsCompensationCurrency = "ZMW"
	WorkerNewContractorParamsCompensationCurrencyTnd WorkerNewContractorParamsCompensationCurrency = "TND"
	WorkerNewContractorParamsCompensationCurrencyNgn WorkerNewContractorParamsCompensationCurrency = "NGN"
	WorkerNewContractorParamsCompensationCurrencyRsd WorkerNewContractorParamsCompensationCurrency = "RSD"
	WorkerNewContractorParamsCompensationCurrencyTwd WorkerNewContractorParamsCompensationCurrency = "TWD"
	WorkerNewContractorParamsCompensationCurrencyGtq WorkerNewContractorParamsCompensationCurrency = "GTQ"
	WorkerNewContractorParamsCompensationCurrencyHnl WorkerNewContractorParamsCompensationCurrency = "HNL"
	WorkerNewContractorParamsCompensationCurrencyDop WorkerNewContractorParamsCompensationCurrency = "DOP"
	WorkerNewContractorParamsCompensationCurrencySar WorkerNewContractorParamsCompensationCurrency = "SAR"
	WorkerNewContractorParamsCompensationCurrencyXaf WorkerNewContractorParamsCompensationCurrency = "XAF"
	WorkerNewContractorParamsCompensationCurrencyPen WorkerNewContractorParamsCompensationCurrency = "PEN"
)

func (r WorkerNewContractorParamsCompensationCurrency) IsKnown() bool {
	switch r {
	case WorkerNewContractorParamsCompensationCurrencyUsd, WorkerNewContractorParamsCompensationCurrencyAud, WorkerNewContractorParamsCompensationCurrencyBgn, WorkerNewContractorParamsCompensationCurrencyBrl, WorkerNewContractorParamsCompensationCurrencyCad, WorkerNewContractorParamsCompensationCurrencyChf, WorkerNewContractorParamsCompensationCurrencyCzk, WorkerNewContractorParamsCompensationCurrencyDkk, WorkerNewContractorParamsCompensationCurrencyEur, WorkerNewContractorParamsCompensationCurrencyGbp, WorkerNewContractorParamsCompensationCurrencyHkd, WorkerNewContractorParamsCompensationCurrencyHuf, WorkerNewContractorParamsCompensationCurrencyIdr, WorkerNewContractorParamsCompensationCurrencyInr, WorkerNewContractorParamsCompensationCurrencyJpy, WorkerNewContractorParamsCompensationCurrencyMyr, WorkerNewContractorParamsCompensationCurrencyNok, WorkerNewContractorParamsCompensationCurrencyNzd, WorkerNewContractorParamsCompensationCurrencyCny, WorkerNewContractorParamsCompensationCurrencyPln, WorkerNewContractorParamsCompensationCurrencyRon, WorkerNewContractorParamsCompensationCurrencyTry, WorkerNewContractorParamsCompensationCurrencySek, WorkerNewContractorParamsCompensationCurrencySgd, WorkerNewContractorParamsCompensationCurrencyAed, WorkerNewContractorParamsCompensationCurrencyArs, WorkerNewContractorParamsCompensationCurrencyBdt, WorkerNewContractorParamsCompensationCurrencyBwp, WorkerNewContractorParamsCompensationCurrencyClp, WorkerNewContractorParamsCompensationCurrencyCop, WorkerNewContractorParamsCompensationCurrencyCrc, WorkerNewContractorParamsCompensationCurrencyEgp, WorkerNewContractorParamsCompensationCurrencyFjd, WorkerNewContractorParamsCompensationCurrencyGel, WorkerNewContractorParamsCompensationCurrencyGhs, WorkerNewContractorParamsCompensationCurrencyIls, WorkerNewContractorParamsCompensationCurrencyKes, WorkerNewContractorParamsCompensationCurrencyKrw, WorkerNewContractorParamsCompensationCurrencyLkr, WorkerNewContractorParamsCompensationCurrencyMad, WorkerNewContractorParamsCompensationCurrencyMxn, WorkerNewContractorParamsCompensationCurrencyNpr, WorkerNewContractorParamsCompensationCurrencyPhp, WorkerNewContractorParamsCompensationCurrencyPkr, WorkerNewContractorParamsCompensationCurrencyThb, WorkerNewContractorParamsCompensationCurrencyUah, WorkerNewContractorParamsCompensationCurrencyUgx, WorkerNewContractorParamsCompensationCurrencyUyu, WorkerNewContractorParamsCompensationCurrencyVnd, WorkerNewContractorParamsCompensationCurrencyZar, WorkerNewContractorParamsCompensationCurrencyZmw, WorkerNewContractorParamsCompensationCurrencyTnd, WorkerNewContractorParamsCompensationCurrencyNgn, WorkerNewContractorParamsCompensationCurrencyRsd, WorkerNewContractorParamsCompensationCurrencyTwd, WorkerNewContractorParamsCompensationCurrencyGtq, WorkerNewContractorParamsCompensationCurrencyHnl, WorkerNewContractorParamsCompensationCurrencyDop, WorkerNewContractorParamsCompensationCurrencySar, WorkerNewContractorParamsCompensationCurrencyXaf, WorkerNewContractorParamsCompensationCurrencyPen:
		return true
	}
	return false
}

// The pay period for the compensation amount.
type WorkerNewContractorParamsCompensationPer string

const (
	WorkerNewContractorParamsCompensationPerHour  WorkerNewContractorParamsCompensationPer = "hour"
	WorkerNewContractorParamsCompensationPerYear  WorkerNewContractorParamsCompensationPer = "year"
	WorkerNewContractorParamsCompensationPerMonth WorkerNewContractorParamsCompensationPer = "month"
	WorkerNewContractorParamsCompensationPerWeek  WorkerNewContractorParamsCompensationPer = "week"
)

func (r WorkerNewContractorParamsCompensationPer) IsKnown() bool {
	switch r {
	case WorkerNewContractorParamsCompensationPerHour, WorkerNewContractorParamsCompensationPerYear, WorkerNewContractorParamsCompensationPerMonth, WorkerNewContractorParamsCompensationPerWeek:
		return true
	}
	return false
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

func (r WorkerNewContractorParamsPaySchedule) IsKnown() bool {
	switch r {
	case WorkerNewContractorParamsPayScheduleWeekly, WorkerNewContractorParamsPayScheduleBiweekly, WorkerNewContractorParamsPayScheduleMonthly, WorkerNewContractorParamsPayScheduleSemimonthly, WorkerNewContractorParamsPayScheduleQuarterly, WorkerNewContractorParamsPayScheduleAnnually:
		return true
	}
	return false
}

type WorkerNewEmployeeParams struct {
	// The employee's base compensation.
	Compensation param.Field[WorkerNewEmployeeParamsCompensation] `json:"compensation" api:"required"`
	// The department to assign this employee to.
	DepartmentID param.Field[string] `json:"departmentId" api:"required"`
	// Personal email address. The invite will be sent here.
	Email param.Field[string] `json:"email" api:"required"`
	// a non empty string
	FirstName param.Field[string] `json:"firstName" api:"required"`
	// a non empty string
	LastName param.Field[string] `json:"lastName" api:"required"`
	// The worker id of this employee's direct manager.
	ManagerID param.Field[string] `json:"managerId" api:"required"`
	// The employee's job title.
	Position param.Field[string] `json:"position" api:"required"`
	// A date string in the form YYYY-MM-DD
	StartDate param.Field[string] `json:"startDate" api:"required"`
	// Where the employee will work. Either an existing company workplace or a remote
	// US state.
	WorkLocation param.Field[WorkerNewEmployeeParamsWorkLocationUnion] `json:"workLocation" api:"required"`
	// The employee's pay schedule. Must be a pay schedule that the company has
	// configured.
	PaySchedule param.Field[WorkerNewEmployeeParamsPaySchedule] `json:"paySchedule"`
	// Whether the employee is required to complete I-9 work authorization. Set to
	// false if the employee has already been verified off-platform. Defaults to true.
	RequireI9 param.Field[bool] `json:"requireI9"`
	// How state tax registration is handled for this employee's work state. Required
	// when hiring in a state where your company doesn't have an existing registration.
	// Use 'self_managed' if you've already registered in this state, or 'warp_managed'
	// for Warp to handle registration on your behalf.
	StateRegistration param.Field[WorkerNewEmployeeParamsStateRegistration] `json:"stateRegistration"`
	// a non-negative number
	StockOptions param.Field[float64] `json:"stockOptions"`
	// An email with a reasonably valid regex (shamelessly taken from zod)
	WorkEmail param.Field[string] `json:"workEmail"`
}

func (r WorkerNewEmployeeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The employee's base compensation.
type WorkerNewEmployeeParamsCompensation struct {
	// a positive number
	Amount param.Field[float64] `json:"amount" api:"required"`
	// Whether the amount is per hour or per year.
	Per param.Field[WorkerNewEmployeeParamsCompensationPer] `json:"per" api:"required"`
}

func (r WorkerNewEmployeeParamsCompensation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the amount is per hour or per year.
type WorkerNewEmployeeParamsCompensationPer string

const (
	WorkerNewEmployeeParamsCompensationPerHour WorkerNewEmployeeParamsCompensationPer = "hour"
	WorkerNewEmployeeParamsCompensationPerYear WorkerNewEmployeeParamsCompensationPer = "year"
)

func (r WorkerNewEmployeeParamsCompensationPer) IsKnown() bool {
	switch r {
	case WorkerNewEmployeeParamsCompensationPerHour, WorkerNewEmployeeParamsCompensationPerYear:
		return true
	}
	return false
}

// Where the employee will work. Either an existing company workplace or a remote
// US state.
type WorkerNewEmployeeParamsWorkLocation struct {
	Type param.Field[WorkerNewEmployeeParamsWorkLocationType] `json:"type" api:"required"`
	// The US state where the remote employee works. Required for tax purposes.
	State param.Field[WorkerNewEmployeeParamsWorkLocationState] `json:"state"`
	// Public workplace identifier
	WorkplaceID param.Field[string] `json:"workplaceId"`
}

func (r WorkerNewEmployeeParamsWorkLocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerNewEmployeeParamsWorkLocation) implementsWorkerNewEmployeeParamsWorkLocationUnion() {}

// Where the employee will work. Either an existing company workplace or a remote
// US state.
//
// Satisfied by [WorkerNewEmployeeParamsWorkLocationOfficeWorkLocation],
// [WorkerNewEmployeeParamsWorkLocationRemoteWorkLocation],
// [WorkerNewEmployeeParamsWorkLocation].
type WorkerNewEmployeeParamsWorkLocationUnion interface {
	implementsWorkerNewEmployeeParamsWorkLocationUnion()
}

// Employee works from a company workplace.
type WorkerNewEmployeeParamsWorkLocationOfficeWorkLocation struct {
	Type param.Field[WorkerNewEmployeeParamsWorkLocationOfficeWorkLocationType] `json:"type" api:"required"`
	// Public workplace identifier
	WorkplaceID param.Field[string] `json:"workplaceId" api:"required"`
}

func (r WorkerNewEmployeeParamsWorkLocationOfficeWorkLocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerNewEmployeeParamsWorkLocationOfficeWorkLocation) implementsWorkerNewEmployeeParamsWorkLocationUnion() {
}

type WorkerNewEmployeeParamsWorkLocationOfficeWorkLocationType string

const (
	WorkerNewEmployeeParamsWorkLocationOfficeWorkLocationTypeOffice WorkerNewEmployeeParamsWorkLocationOfficeWorkLocationType = "office"
)

func (r WorkerNewEmployeeParamsWorkLocationOfficeWorkLocationType) IsKnown() bool {
	switch r {
	case WorkerNewEmployeeParamsWorkLocationOfficeWorkLocationTypeOffice:
		return true
	}
	return false
}

// Employee works remotely from a US state.
type WorkerNewEmployeeParamsWorkLocationRemoteWorkLocation struct {
	// The US state where the remote employee works. Required for tax purposes.
	State param.Field[WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState] `json:"state" api:"required"`
	Type  param.Field[WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationType]  `json:"type" api:"required"`
}

func (r WorkerNewEmployeeParamsWorkLocationRemoteWorkLocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerNewEmployeeParamsWorkLocationRemoteWorkLocation) implementsWorkerNewEmployeeParamsWorkLocationUnion() {
}

// The US state where the remote employee works. Required for tax purposes.
type WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState string

const (
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateAl WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "AL"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateAk WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "AK"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateAz WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "AZ"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateAr WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "AR"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateCa WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "CA"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateCo WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "CO"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateCt WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "CT"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateDc WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "DC"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateDe WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "DE"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateFl WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "FL"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateGa WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "GA"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateHi WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "HI"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateID WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "ID"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateIl WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "IL"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateIn WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "IN"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateIa WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "IA"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateKs WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "KS"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateKy WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "KY"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateLa WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "LA"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateMe WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "ME"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateMd WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "MD"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateMa WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "MA"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateMi WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "MI"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateMn WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "MN"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateMs WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "MS"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateMo WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "MO"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateMt WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "MT"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateNe WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "NE"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateNv WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "NV"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateNh WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "NH"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateNj WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "NJ"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateNm WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "NM"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateNy WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "NY"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateNc WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "NC"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateNd WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "ND"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateOh WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "OH"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateOk WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "OK"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateOr WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "OR"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStatePa WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "PA"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateRi WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "RI"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateSc WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "SC"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateSd WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "SD"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateTn WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "TN"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateTx WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "TX"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateUt WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "UT"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateVt WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "VT"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateVa WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "VA"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateWa WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "WA"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateWv WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "WV"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateWi WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "WI"
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateWy WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState = "WY"
)

func (r WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationState) IsKnown() bool {
	switch r {
	case WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateAl, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateAk, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateAz, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateAr, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateCa, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateCo, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateCt, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateDc, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateDe, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateFl, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateGa, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateHi, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateID, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateIl, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateIn, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateIa, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateKs, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateKy, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateLa, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateMe, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateMd, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateMa, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateMi, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateMn, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateMs, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateMo, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateMt, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateNe, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateNv, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateNh, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateNj, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateNm, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateNy, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateNc, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateNd, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateOh, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateOk, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateOr, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStatePa, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateRi, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateSc, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateSd, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateTn, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateTx, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateUt, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateVt, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateVa, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateWa, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateWv, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateWi, WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationStateWy:
		return true
	}
	return false
}

type WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationType string

const (
	WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationTypeRemote WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationType = "remote"
)

func (r WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationType) IsKnown() bool {
	switch r {
	case WorkerNewEmployeeParamsWorkLocationRemoteWorkLocationTypeRemote:
		return true
	}
	return false
}

type WorkerNewEmployeeParamsWorkLocationType string

const (
	WorkerNewEmployeeParamsWorkLocationTypeOffice WorkerNewEmployeeParamsWorkLocationType = "office"
	WorkerNewEmployeeParamsWorkLocationTypeRemote WorkerNewEmployeeParamsWorkLocationType = "remote"
)

func (r WorkerNewEmployeeParamsWorkLocationType) IsKnown() bool {
	switch r {
	case WorkerNewEmployeeParamsWorkLocationTypeOffice, WorkerNewEmployeeParamsWorkLocationTypeRemote:
		return true
	}
	return false
}

// The US state where the remote employee works. Required for tax purposes.
type WorkerNewEmployeeParamsWorkLocationState string

const (
	WorkerNewEmployeeParamsWorkLocationStateAl WorkerNewEmployeeParamsWorkLocationState = "AL"
	WorkerNewEmployeeParamsWorkLocationStateAk WorkerNewEmployeeParamsWorkLocationState = "AK"
	WorkerNewEmployeeParamsWorkLocationStateAz WorkerNewEmployeeParamsWorkLocationState = "AZ"
	WorkerNewEmployeeParamsWorkLocationStateAr WorkerNewEmployeeParamsWorkLocationState = "AR"
	WorkerNewEmployeeParamsWorkLocationStateCa WorkerNewEmployeeParamsWorkLocationState = "CA"
	WorkerNewEmployeeParamsWorkLocationStateCo WorkerNewEmployeeParamsWorkLocationState = "CO"
	WorkerNewEmployeeParamsWorkLocationStateCt WorkerNewEmployeeParamsWorkLocationState = "CT"
	WorkerNewEmployeeParamsWorkLocationStateDc WorkerNewEmployeeParamsWorkLocationState = "DC"
	WorkerNewEmployeeParamsWorkLocationStateDe WorkerNewEmployeeParamsWorkLocationState = "DE"
	WorkerNewEmployeeParamsWorkLocationStateFl WorkerNewEmployeeParamsWorkLocationState = "FL"
	WorkerNewEmployeeParamsWorkLocationStateGa WorkerNewEmployeeParamsWorkLocationState = "GA"
	WorkerNewEmployeeParamsWorkLocationStateHi WorkerNewEmployeeParamsWorkLocationState = "HI"
	WorkerNewEmployeeParamsWorkLocationStateID WorkerNewEmployeeParamsWorkLocationState = "ID"
	WorkerNewEmployeeParamsWorkLocationStateIl WorkerNewEmployeeParamsWorkLocationState = "IL"
	WorkerNewEmployeeParamsWorkLocationStateIn WorkerNewEmployeeParamsWorkLocationState = "IN"
	WorkerNewEmployeeParamsWorkLocationStateIa WorkerNewEmployeeParamsWorkLocationState = "IA"
	WorkerNewEmployeeParamsWorkLocationStateKs WorkerNewEmployeeParamsWorkLocationState = "KS"
	WorkerNewEmployeeParamsWorkLocationStateKy WorkerNewEmployeeParamsWorkLocationState = "KY"
	WorkerNewEmployeeParamsWorkLocationStateLa WorkerNewEmployeeParamsWorkLocationState = "LA"
	WorkerNewEmployeeParamsWorkLocationStateMe WorkerNewEmployeeParamsWorkLocationState = "ME"
	WorkerNewEmployeeParamsWorkLocationStateMd WorkerNewEmployeeParamsWorkLocationState = "MD"
	WorkerNewEmployeeParamsWorkLocationStateMa WorkerNewEmployeeParamsWorkLocationState = "MA"
	WorkerNewEmployeeParamsWorkLocationStateMi WorkerNewEmployeeParamsWorkLocationState = "MI"
	WorkerNewEmployeeParamsWorkLocationStateMn WorkerNewEmployeeParamsWorkLocationState = "MN"
	WorkerNewEmployeeParamsWorkLocationStateMs WorkerNewEmployeeParamsWorkLocationState = "MS"
	WorkerNewEmployeeParamsWorkLocationStateMo WorkerNewEmployeeParamsWorkLocationState = "MO"
	WorkerNewEmployeeParamsWorkLocationStateMt WorkerNewEmployeeParamsWorkLocationState = "MT"
	WorkerNewEmployeeParamsWorkLocationStateNe WorkerNewEmployeeParamsWorkLocationState = "NE"
	WorkerNewEmployeeParamsWorkLocationStateNv WorkerNewEmployeeParamsWorkLocationState = "NV"
	WorkerNewEmployeeParamsWorkLocationStateNh WorkerNewEmployeeParamsWorkLocationState = "NH"
	WorkerNewEmployeeParamsWorkLocationStateNj WorkerNewEmployeeParamsWorkLocationState = "NJ"
	WorkerNewEmployeeParamsWorkLocationStateNm WorkerNewEmployeeParamsWorkLocationState = "NM"
	WorkerNewEmployeeParamsWorkLocationStateNy WorkerNewEmployeeParamsWorkLocationState = "NY"
	WorkerNewEmployeeParamsWorkLocationStateNc WorkerNewEmployeeParamsWorkLocationState = "NC"
	WorkerNewEmployeeParamsWorkLocationStateNd WorkerNewEmployeeParamsWorkLocationState = "ND"
	WorkerNewEmployeeParamsWorkLocationStateOh WorkerNewEmployeeParamsWorkLocationState = "OH"
	WorkerNewEmployeeParamsWorkLocationStateOk WorkerNewEmployeeParamsWorkLocationState = "OK"
	WorkerNewEmployeeParamsWorkLocationStateOr WorkerNewEmployeeParamsWorkLocationState = "OR"
	WorkerNewEmployeeParamsWorkLocationStatePa WorkerNewEmployeeParamsWorkLocationState = "PA"
	WorkerNewEmployeeParamsWorkLocationStateRi WorkerNewEmployeeParamsWorkLocationState = "RI"
	WorkerNewEmployeeParamsWorkLocationStateSc WorkerNewEmployeeParamsWorkLocationState = "SC"
	WorkerNewEmployeeParamsWorkLocationStateSd WorkerNewEmployeeParamsWorkLocationState = "SD"
	WorkerNewEmployeeParamsWorkLocationStateTn WorkerNewEmployeeParamsWorkLocationState = "TN"
	WorkerNewEmployeeParamsWorkLocationStateTx WorkerNewEmployeeParamsWorkLocationState = "TX"
	WorkerNewEmployeeParamsWorkLocationStateUt WorkerNewEmployeeParamsWorkLocationState = "UT"
	WorkerNewEmployeeParamsWorkLocationStateVt WorkerNewEmployeeParamsWorkLocationState = "VT"
	WorkerNewEmployeeParamsWorkLocationStateVa WorkerNewEmployeeParamsWorkLocationState = "VA"
	WorkerNewEmployeeParamsWorkLocationStateWa WorkerNewEmployeeParamsWorkLocationState = "WA"
	WorkerNewEmployeeParamsWorkLocationStateWv WorkerNewEmployeeParamsWorkLocationState = "WV"
	WorkerNewEmployeeParamsWorkLocationStateWi WorkerNewEmployeeParamsWorkLocationState = "WI"
	WorkerNewEmployeeParamsWorkLocationStateWy WorkerNewEmployeeParamsWorkLocationState = "WY"
)

func (r WorkerNewEmployeeParamsWorkLocationState) IsKnown() bool {
	switch r {
	case WorkerNewEmployeeParamsWorkLocationStateAl, WorkerNewEmployeeParamsWorkLocationStateAk, WorkerNewEmployeeParamsWorkLocationStateAz, WorkerNewEmployeeParamsWorkLocationStateAr, WorkerNewEmployeeParamsWorkLocationStateCa, WorkerNewEmployeeParamsWorkLocationStateCo, WorkerNewEmployeeParamsWorkLocationStateCt, WorkerNewEmployeeParamsWorkLocationStateDc, WorkerNewEmployeeParamsWorkLocationStateDe, WorkerNewEmployeeParamsWorkLocationStateFl, WorkerNewEmployeeParamsWorkLocationStateGa, WorkerNewEmployeeParamsWorkLocationStateHi, WorkerNewEmployeeParamsWorkLocationStateID, WorkerNewEmployeeParamsWorkLocationStateIl, WorkerNewEmployeeParamsWorkLocationStateIn, WorkerNewEmployeeParamsWorkLocationStateIa, WorkerNewEmployeeParamsWorkLocationStateKs, WorkerNewEmployeeParamsWorkLocationStateKy, WorkerNewEmployeeParamsWorkLocationStateLa, WorkerNewEmployeeParamsWorkLocationStateMe, WorkerNewEmployeeParamsWorkLocationStateMd, WorkerNewEmployeeParamsWorkLocationStateMa, WorkerNewEmployeeParamsWorkLocationStateMi, WorkerNewEmployeeParamsWorkLocationStateMn, WorkerNewEmployeeParamsWorkLocationStateMs, WorkerNewEmployeeParamsWorkLocationStateMo, WorkerNewEmployeeParamsWorkLocationStateMt, WorkerNewEmployeeParamsWorkLocationStateNe, WorkerNewEmployeeParamsWorkLocationStateNv, WorkerNewEmployeeParamsWorkLocationStateNh, WorkerNewEmployeeParamsWorkLocationStateNj, WorkerNewEmployeeParamsWorkLocationStateNm, WorkerNewEmployeeParamsWorkLocationStateNy, WorkerNewEmployeeParamsWorkLocationStateNc, WorkerNewEmployeeParamsWorkLocationStateNd, WorkerNewEmployeeParamsWorkLocationStateOh, WorkerNewEmployeeParamsWorkLocationStateOk, WorkerNewEmployeeParamsWorkLocationStateOr, WorkerNewEmployeeParamsWorkLocationStatePa, WorkerNewEmployeeParamsWorkLocationStateRi, WorkerNewEmployeeParamsWorkLocationStateSc, WorkerNewEmployeeParamsWorkLocationStateSd, WorkerNewEmployeeParamsWorkLocationStateTn, WorkerNewEmployeeParamsWorkLocationStateTx, WorkerNewEmployeeParamsWorkLocationStateUt, WorkerNewEmployeeParamsWorkLocationStateVt, WorkerNewEmployeeParamsWorkLocationStateVa, WorkerNewEmployeeParamsWorkLocationStateWa, WorkerNewEmployeeParamsWorkLocationStateWv, WorkerNewEmployeeParamsWorkLocationStateWi, WorkerNewEmployeeParamsWorkLocationStateWy:
		return true
	}
	return false
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

func (r WorkerNewEmployeeParamsPaySchedule) IsKnown() bool {
	switch r {
	case WorkerNewEmployeeParamsPayScheduleWeekly, WorkerNewEmployeeParamsPayScheduleBiweekly, WorkerNewEmployeeParamsPayScheduleMonthly, WorkerNewEmployeeParamsPayScheduleSemimonthly, WorkerNewEmployeeParamsPayScheduleQuarterly, WorkerNewEmployeeParamsPayScheduleAnnually:
		return true
	}
	return false
}

// How state tax registration is handled for this employee's work state. Required
// when hiring in a state where your company doesn't have an existing registration.
// Use 'self_managed' if you've already registered in this state, or 'warp_managed'
// for Warp to handle registration on your behalf.
type WorkerNewEmployeeParamsStateRegistration string

const (
	WorkerNewEmployeeParamsStateRegistrationSelfManaged WorkerNewEmployeeParamsStateRegistration = "self_managed"
	WorkerNewEmployeeParamsStateRegistrationWarpManaged WorkerNewEmployeeParamsStateRegistration = "warp_managed"
)

func (r WorkerNewEmployeeParamsStateRegistration) IsKnown() bool {
	switch r {
	case WorkerNewEmployeeParamsStateRegistrationSelfManaged, WorkerNewEmployeeParamsStateRegistrationWarpManaged:
		return true
	}
	return false
}
