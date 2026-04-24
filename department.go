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

// Endpoints for department management. Create, list, and update departments within
// your company.
//
// DepartmentService contains methods and other services that help with interacting
// with the warp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDepartmentService] method instead.
type DepartmentService struct {
	Options []option.RequestOption
}

// NewDepartmentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDepartmentService(opts ...option.RequestOption) (r *DepartmentService) {
	r = &DepartmentService{}
	r.Options = opts
	return
}

// Create a new department.
func (r *DepartmentService) New(ctx context.Context, body DepartmentNewParams, opts ...option.RequestOption) (res *DepartmentNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/departments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update an existing department.
func (r *DepartmentService) Update(ctx context.Context, id string, body DepartmentUpdateParams, opts ...option.RequestOption) (res *DepartmentUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/departments/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List all departments for your company.
func (r *DepartmentService) List(ctx context.Context, query DepartmentListParams, opts ...option.RequestOption) (res *pagination.CursorPage[DepartmentListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/departments"
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

// List all departments for your company.
func (r *DepartmentService) ListAutoPaging(ctx context.Context, query DepartmentListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[DepartmentListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

type DepartmentNewResponse struct {
	// The unique public id of the department
	ID string `json:"id" api:"required"`
	// a string to be decoded into a Date
	CreatedAt string                    `json:"createdAt" api:"required"`
	Name      string                    `json:"name" api:"required"`
	JSON      departmentNewResponseJSON `json:"-"`
}

// departmentNewResponseJSON contains the JSON metadata for the struct
// [DepartmentNewResponse]
type departmentNewResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DepartmentNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r departmentNewResponseJSON) RawJSON() string {
	return r.raw
}

type DepartmentUpdateResponse struct {
	// The unique public id of the department
	ID string `json:"id" api:"required"`
	// a string to be decoded into a Date
	CreatedAt string                       `json:"createdAt" api:"required"`
	Name      string                       `json:"name" api:"required"`
	JSON      departmentUpdateResponseJSON `json:"-"`
}

// departmentUpdateResponseJSON contains the JSON metadata for the struct
// [DepartmentUpdateResponse]
type departmentUpdateResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DepartmentUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r departmentUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type DepartmentListResponse struct {
	// The unique public id of the department
	ID string `json:"id" api:"required"`
	// a string to be decoded into a Date
	CreatedAt string                     `json:"createdAt" api:"required"`
	Name      string                     `json:"name" api:"required"`
	JSON      departmentListResponseJSON `json:"-"`
}

// departmentListResponseJSON contains the JSON metadata for the struct
// [DepartmentListResponse]
type departmentListResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DepartmentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r departmentListResponseJSON) RawJSON() string {
	return r.raw
}

type DepartmentNewParams struct {
	// a non empty string
	Name param.Field[string] `json:"name" api:"required"`
}

func (r DepartmentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DepartmentUpdateParams struct {
	Name param.Field[string] `json:"name"`
}

func (r DepartmentUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DepartmentListParams struct {
	// The unique public id of the department
	AfterID param.Field[string] `query:"afterId"`
	// The unique public id of the department
	BeforeID param.Field[string] `query:"beforeId"`
	// a number less than or equal to 100
	Limit param.Field[string] `query:"limit"`
}

// URLQuery serializes [DepartmentListParams]'s query parameters as `url.Values`.
func (r DepartmentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
