# Warp Go API

Complete reference of every operation, grouped by resource. See [the README](./README.md) for usage and configuration.

## Contents

- [`Benefits`](#benefits)
  - [`Benefits HealthPlans`](#benefits-healthplans)
    - [List Health Plans](#list-health-plans)
    - [Get Health Plan](#get-health-plan)
  - [`Benefits RetirementPlans`](#benefits-retirementplans)
    - [List Retirement Plans](#list-retirement-plans)
    - [Get Retirement Plan](#get-retirement-plan)
  - [`Benefits Deductions`](#benefits-deductions)
    - [List Benefit Deductions](#list-benefit-deductions)
    - [Get Benefit Deduction](#get-benefit-deduction)
- [`CustomFields`](#customfields)
  - [List Fields](#list-fields)
  - [Create Field](#create-field)
  - [Get Field](#get-field)
  - [Update Field](#update-field)
  - [Archive Field](#archive-field)
  - [Create Field Option](#create-field-option)
  - [Update Field Option](#update-field-option)
  - [Delete Unused Field Option](#delete-unused-field-option)
  - [Archive Field Option](#archive-field-option)
  - [List Field Values](#list-field-values)
  - [Set Field Value](#set-field-value)
  - [Clear Field Value](#clear-field-value)
- [`Departments`](#departments)
  - [List Departments](#list-departments)
  - [Create Department](#create-department)
  - [Update Department](#update-department)
- [`Offers`](#offers)
  - [List Offers](#list-offers)
  - [Create Offer](#create-offer)
  - [Void Offer](#void-offer)
  - [Extend Offer Deadline](#extend-offer-deadline)
  - [Resend Offer](#resend-offer)
- [`TimeOff`](#timeoff)
  - [List Time Off Assignments](#list-time-off-assignments)
  - [List Time Off Balances](#list-time-off-balances)
  - [List Time Off Requests](#list-time-off-requests)
  - [`TimeOff Policies`](#timeoff-policies)
    - [List Time Off Policies](#list-time-off-policies)
    - [Get Time Off Policy](#get-time-off-policy)
- [`Workers`](#workers)
  - [List Workers](#list-workers)
  - [Get Worker](#get-worker)
  - [Delete Worker](#delete-worker)
  - [Create Employee](#create-employee)
  - [Create Contractor](#create-contractor)
  - [Invite Worker](#invite-worker)
- [`Workplaces`](#workplaces)
  - [List Workplaces](#list-workplaces)
  - [Create Workplace](#create-workplace)
  - [Update Workplace](#update-workplace)

## Setup

```go
import (
	"context"
	"fmt"

	sdk "github.com/TeamWarp/warp-go-sdk"
)

client := sdk.NewClient()
```

## `Benefits`

### `Benefits HealthPlans`

#### List Health Plans

List company health plans. Defaults to active plans. A plan whose effectiveEndDate has elapsed is reported and filtered as terminated.

| Direction | Type |
| --- | --- |
| Request | [`BenefitHealthPlanListParams`](./benefithealthplan.go) |
| Response | [`BenefitHealthPlanListResponse`](./benefithealthplan.go) |

```go
healthPlan, err := client.Benefits.HealthPlans.List(context.Background(), sdk.BenefitHealthPlanListParams{
	Statuses: sdk.F[[]sdk.BenefitHealthPlanListParamsStatus]([]sdk.BenefitHealthPlanListParamsStatus{"active"}),
})
if err != nil {
	panic(err)
}

fmt.Println(healthPlan)
```

#### Get Health Plan

Get a publicly visible company health plan by id.

| Direction | Type |
| --- | --- |
| Response | [`BenefitHealthPlanGetResponse`](./benefithealthplan.go) |

```go
healthPlan, err := client.Benefits.HealthPlans.Get(context.Background(), "chpl_1234")
if err != nil {
	panic(err)
}

fmt.Println(healthPlan)
```

### `Benefits RetirementPlans`

#### List Retirement Plans

List company retirement plans. Defaults to active plans. A plan whose effectiveEndDate has elapsed is reported and filtered as terminated.

| Direction | Type |
| --- | --- |
| Request | [`BenefitRetirementPlanListParams`](./benefitretirementplan.go) |
| Response | [`BenefitRetirementPlanListResponse`](./benefitretirementplan.go) |

```go
retirementPlan, err := client.Benefits.RetirementPlans.List(context.Background(), sdk.BenefitRetirementPlanListParams{
	Statuses: sdk.F[[]sdk.BenefitRetirementPlanListParamsStatus]([]sdk.BenefitRetirementPlanListParamsStatus{"active"}),
})
if err != nil {
	panic(err)
}

fmt.Println(retirementPlan)
```

#### Get Retirement Plan

Get a company retirement plan by id, regardless of status.

| Direction | Type |
| --- | --- |
| Response | [`BenefitRetirementPlanGetResponse`](./benefitretirementplan.go) |

```go
retirementPlan, err := client.Benefits.RetirementPlans.Get(context.Background(), "crpl_1234")
if err != nil {
	panic(err)
}

fmt.Println(retirementPlan)
```

### `Benefits Deductions`

#### List Benefit Deductions

List current payroll benefit deductions. Defaults to active deductions. A deduction whose effectiveEndDate has elapsed is reported and filtered as terminated.

| Direction | Type |
| --- | --- |
| Request | [`BenefitDeductionListParams`](./benefitdeduction.go) |
| Response | [`BenefitDeductionListResponse`](./benefitdeduction.go) |

```go
deduction, err := client.Benefits.Deductions.List(context.Background(), sdk.BenefitDeductionListParams{
	Statuses: sdk.F[[]sdk.BenefitDeductionListParamsStatus]([]sdk.BenefitDeductionListParamsStatus{"active"}),
})
if err != nil {
	panic(err)
}

fmt.Println(deduction)
```

#### Get Benefit Deduction

Get the current version of a company benefit deduction by id.

| Direction | Type |
| --- | --- |
| Response | [`BenefitDeductionGetResponse`](./benefitdeduction.go) |

```go
deduction, err := client.Benefits.Deductions.Get(context.Background(), "pbdg_1234")
if err != nil {
	panic(err)
}

fmt.Println(deduction)
```

## `CustomFields`

### List Fields

List the custom worker field definitions your API key can read. Each field belongs to a worker-data category; fields whose category your key cannot read are omitted unless the key holds workers:custom_fields.

| Direction | Type |
| --- | --- |
| Response | [`[]CustomFieldListResponse`](./customfield.go) |

```go
customField, err := client.CustomFields.List(context.Background())
if err != nil {
	panic(err)
}

fmt.Println(customField)
```

### Create Field

Create a custom worker field definition. The field type is immutable after creation. Select and multi_select fields can include their initial options. Access to values derives from the field category; requires the workers:custom_fields permission.

| Direction | Type |
| --- | --- |
| Request | [`CustomFieldNewParams`](./customfield.go) |
| Response | [`CustomFieldNewResponse`](./customfield.go) |

```go
customField, err := client.CustomFields.New(context.Background(), sdk.CustomFieldNewParams{
	Name: sdk.F[string](""),
})
if err != nil {
	panic(err)
}

fmt.Println(customField)
```

### Get Field

Get a custom worker field definition, including its select options. Archived options may appear on existing worker values but cannot be newly selected.

| Direction | Type |
| --- | --- |
| Response | [`CustomFieldGetResponse`](./customfield.go) |

```go
customField, err := client.CustomFields.Get(context.Background(), "cf_1234")
if err != nil {
	panic(err)
}

fmt.Println(customField)
```

### Update Field

Update a custom worker field definition. The field type cannot be changed; create a new field instead. Requires the workers:custom_fields permission; changing the category, access level, or input source requires the manage level.

| Direction | Type |
| --- | --- |
| Request | [`CustomFieldUpdateParams`](./customfield.go) |
| Response | [`CustomFieldUpdateResponse`](./customfield.go) |

```go
customField, err := client.CustomFields.Update(context.Background(), "cf_1234", sdk.CustomFieldUpdateParams{})
if err != nil {
	panic(err)
}

fmt.Println(customField)
```

### Archive Field

Archive a custom worker field. Archived fields keep their existing worker values but cannot receive new ones. Requires the workers:custom_fields permission at the manage level.

| Direction | Type |
| --- | --- |
| Response | [`CustomFieldArchiveResponse`](./customfield.go) |

```go
customField, err := client.CustomFields.Archive(context.Background(), "cf_1234")
if err != nil {
	panic(err)
}

fmt.Println(customField)
```

### Create Field Option

Add an option to a select or multi_select custom worker field. The option value should be treated as stable; the label can change. Requires the workers:custom_fields permission.

| Direction | Type |
| --- | --- |
| Request | [`CustomFieldNewOptionParams`](./customfield.go) |
| Response | [`CustomFieldNewOptionResponse`](./customfield.go) |

```go
customField, err := client.CustomFields.NewOption(context.Background(), "cf_1234", sdk.CustomFieldNewOptionParams{
	Label: sdk.F[string]("x"),
	Value: sdk.F[string]("x"),
})
if err != nil {
	panic(err)
}

fmt.Println(customField)
```

### Update Field Option

Update the label or sort order of a custom worker field option. Options of archived fields cannot be edited. Requires the workers:custom_fields permission.

| Direction | Type |
| --- | --- |
| Request | [`CustomFieldUpdateOptionParams`](./customfield.go) |
| Response | [`CustomFieldUpdateOptionResponse`](./customfield.go) |

```go
customField, err := client.CustomFields.UpdateOption(context.Background(), "cfo_1234", sdk.CustomFieldUpdateOptionParams{})
if err != nil {
	panic(err)
}

fmt.Println(customField)
```

### Delete Unused Field Option

Delete a custom worker field option that is not applied to any worker. Options in use must be archived instead. Requires the workers:custom_fields permission at the manage level.

```go
err := client.CustomFields.DeleteOption(context.Background(), "cfo_1234")
if err != nil {
	panic(err)
}
```

### Archive Field Option

Archive a custom worker field option. Archived options remain on existing worker values but cannot be newly selected. Requires the workers:custom_fields permission at the manage level.

| Direction | Type |
| --- | --- |
| Response | [`CustomFieldArchiveOptionResponse`](./customfield.go) |

```go
customField, err := client.CustomFields.ArchiveOption(context.Background(), "cfo_1234")
if err != nil {
	panic(err)
}

fmt.Println(customField)
```

### List Field Values

List custom field values for workers, optionally filtered by worker or field. Values are returned only for fields whose category your API key can read.

| Direction | Type |
| --- | --- |
| Request | [`CustomFieldListValuesParams`](./customfield.go) |
| Response | [`[]CustomFieldListValuesResponse`](./customfield.go) |

```go
customField, err := client.CustomFields.ListValues(context.Background(), sdk.CustomFieldListValuesParams{})
if err != nil {
	panic(err)
}

fmt.Println(customField)
```

### Set Field Value

Create or replace a worker's value for a custom field. The value shape must match the field type, and your API key must hold write on the field's category.

| Direction | Type |
| --- | --- |
| Request | [`CustomFieldUpsertValueParams`](./customfield.go) |
| Response | [`CustomFieldUpsertValueResponse`](./customfield.go) |

```go
customField, err := client.CustomFields.UpsertValue(context.Background(), sdk.CustomFieldUpsertValueParams{
	FieldID:  sdk.F[string]("cf_1234"),
	Value:    sdk.F[sdk.CustomFieldUpsertValueParamsValueUnion](sdk.CustomFieldUpsertValueParamsValueUnion{}),
	WorkerID: sdk.F[string]("wrk_1234"),
})
if err != nil {
	panic(err)
}

fmt.Println(customField)
```

### Clear Field Value

Remove a worker's value for a custom field. Your API key must hold write on the field's category.

| Direction | Type |
| --- | --- |
| Request | [`CustomFieldClearValueParams`](./customfield.go) |

```go
err := client.CustomFields.ClearValue(context.Background(), sdk.CustomFieldClearValueParams{
	FieldID:  sdk.F[string]("cf_1234"),
	WorkerID: sdk.F[string]("wrk_1234"),
})
if err != nil {
	panic(err)
}
```

## `Departments`

### List Departments

List all departments for your company.

| Direction | Type |
| --- | --- |
| Request | [`DepartmentListParams`](./department.go) |
| Response | [`DepartmentListResponse`](./department.go) |

```go
department, err := client.Departments.List(context.Background(), sdk.DepartmentListParams{})
if err != nil {
	panic(err)
}

fmt.Println(department)
```

### Create Department

Create a new department.

| Direction | Type |
| --- | --- |
| Request | [`DepartmentNewParams`](./department.go) |
| Response | [`DepartmentNewResponse`](./department.go) |

```go
department, err := client.Departments.New(context.Background(), sdk.DepartmentNewParams{
	Name: sdk.F[string](""),
})
if err != nil {
	panic(err)
}

fmt.Println(department)
```

### Update Department

Update an existing department.

| Direction | Type |
| --- | --- |
| Request | [`DepartmentUpdateParams`](./department.go) |
| Response | [`DepartmentUpdateResponse`](./department.go) |

```go
department, err := client.Departments.Update(context.Background(), "dpt_1234", sdk.DepartmentUpdateParams{})
if err != nil {
	panic(err)
}

fmt.Println(department)
```

## `Offers`

### List Offers

List the candidate offers for your company.

| Direction | Type |
| --- | --- |
| Request | [`OfferListParams`](./offer.go) |
| Response | [`OfferListResponse`](./offer.go) |

```go
offer, err := client.Offers.List(context.Background(), sdk.OfferListParams{})
if err != nil {
	panic(err)
}

fmt.Println(offer)
```

### Create Offer

Create and send a candidate offer. The candidate receives an email with a link to the offer portal.

| Direction | Type |
| --- | --- |
| Request | [`OfferNewParams`](./offer.go) |
| Response | [`OfferNewResponse`](./offer.go) |

```go
offer, err := client.Offers.New(context.Background(), sdk.OfferNewParams{
	Candidate: sdk.F[sdk.OfferNewParamsCandidate](sdk.OfferNewParamsCandidate{
		FirstName: sdk.F[string]("x"),
		LastName:  sdk.F[string]("x"),
		Email:     sdk.F[string]("john@joinwarp.com"),
	}),
	Compensation: sdk.F[sdk.OfferNewParamsCompensation](sdk.OfferNewParamsCompensation{
		PayRate: sdk.F[float64](0),
	}),
	Position: sdk.F[sdk.OfferNewParamsPosition](sdk.OfferNewParamsPosition{
		Title:     sdk.F[string]("x"),
		StartDate: sdk.F[string]("2000-01-01"),
	}),
})
if err != nil {
	panic(err)
}

fmt.Println(offer)
```

### Void Offer

Void a previously sent offer. Only sent offers can be voided.

| Direction | Type |
| --- | --- |
| Response | [`OfferVoidResponse`](./offer.go) |

```go
offer, err := client.Offers.Void(context.Background(), "offr_1234")
if err != nil {
	panic(err)
}

fmt.Println(offer)
```

### Extend Offer Deadline

Extend the expiration deadline of a sent offer.

| Direction | Type |
| --- | --- |
| Request | [`OfferExtendDeadlineParams`](./offer.go) |
| Response | [`OfferExtendDeadlineResponse`](./offer.go) |

```go
offer, err := client.Offers.ExtendDeadline(context.Background(), "offr_1234", sdk.OfferExtendDeadlineParams{
	ExpirationTime: sdk.F[string](""),
})
if err != nil {
	panic(err)
}

fmt.Println(offer)
```

### Resend Offer

Resend the offer email to the candidate for a sent offer.

| Direction | Type |
| --- | --- |
| Response | [`OfferResendResponse`](./offer.go) |

```go
offer, err := client.Offers.Resend(context.Background(), "offr_1234")
if err != nil {
	panic(err)
}

fmt.Println(offer)
```

## `TimeOff`

### List Time Off Assignments

Time off assignments are mappings between workers and time off policies. Useful for finding out which policies a worker is assigned to, or which workers are assigned to a given policy.

| Direction | Type |
| --- | --- |
| Request | [`TimeOffListAssignmentsParams`](./timeoff.go) |
| Response | [`TimeOffListAssignmentsResponse`](./timeoff.go) |

```go
timeOff, err := client.TimeOff.ListAssignments(context.Background(), sdk.TimeOffListAssignmentsParams{})
if err != nil {
	panic(err)
}

fmt.Println(timeOff)
```

### List Time Off Balances

Get worker remaining time-off balances.

| Direction | Type |
| --- | --- |
| Request | [`TimeOffListBalancesParams`](./timeoff.go) |
| Response | [`TimeOffListBalancesResponse`](./timeoff.go) |

```go
timeOff, err := client.TimeOff.ListBalances(context.Background(), sdk.TimeOffListBalancesParams{})
if err != nil {
	panic(err)
}

fmt.Println(timeOff)
```

### List Time Off Requests

Get the time off requests that workers in your company have made.

| Direction | Type |
| --- | --- |
| Request | [`TimeOffListRequestsParams`](./timeoff.go) |
| Response | [`TimeOffListRequestsResponse`](./timeoff.go) |

```go
timeOff, err := client.TimeOff.ListRequests(context.Background(), sdk.TimeOffListRequestsParams{})
if err != nil {
	panic(err)
}

fmt.Println(timeOff)
```

### `TimeOff Policies`

#### List Time Off Policies

Get the time off policies for your company

| Direction | Type |
| --- | --- |
| Request | [`TimeOffPolicyTimeOffGetParams`](./timeoffpolicy.go) |
| Response | [`TimeOffPolicyTimeOffGetResponse`](./timeoffpolicy.go) |

```go
policy, err := client.TimeOff.Policies.TimeOffGet(context.Background(), sdk.TimeOffPolicyTimeOffGetParams{})
if err != nil {
	panic(err)
}

fmt.Println(policy)
```

#### Get Time Off Policy

Get a specific time off policy by id

| Direction | Type |
| --- | --- |
| Response | [`TimeOffPolicyTimeOffGet2Response`](./timeoffpolicy.go) |

```go
policy, err := client.TimeOff.Policies.TimeOffGet2(context.Background(), "top_1234")
if err != nil {
	panic(err)
}

fmt.Println(policy)
```

## `Workers`

### List Workers

List all workers. Workers include anyone employed by the company, whether US or international, full-time employees or contractors.

| Direction | Type |
| --- | --- |
| Request | [`WorkerListParams`](./worker.go) |
| Response | [`WorkerListResponse`](./worker.go) |

```go
worker, err := client.Workers.List(context.Background(), sdk.WorkerListParams{})
if err != nil {
	panic(err)
}

fmt.Println(worker)
```

### Get Worker

Get a specific worker by id.

| Direction | Type |
| --- | --- |
| Response | [`WorkerGetResponse`](./worker.go) |

```go
worker, err := client.Workers.Get(context.Background(), "wrk_1234")
if err != nil {
	panic(err)
}

fmt.Println(worker)
```

### Delete Worker

Delete a worker. Only workers who have not yet completed onboarding can be deleted. Active workers must be properly offboarded.

```go
err := client.Workers.Delete(context.Background(), "wrk_1234")
if err != nil {
	panic(err)
}
```

### Create Employee

Create a new US employee. The worker will be created in draft status and must be invited separately via the invite endpoint. If hiring in a state without an existing tax registration, you must specify the stateRegistration field.

| Direction | Type |
| --- | --- |
| Request | [`WorkerNewEmployeeParams`](./worker.go) |
| Response | [`WorkerNewEmployeeResponse`](./worker.go) |

```go
worker, err := client.Workers.NewEmployee(context.Background(), sdk.WorkerNewEmployeeParams{
	Compensation: sdk.F[sdk.WorkerNewEmployeeParamsCompensation](sdk.WorkerNewEmployeeParamsCompensation{
		Amount: sdk.F[float64](0),
	}),
	DepartmentID: sdk.F[string]("dpt_1234"),
	Email:        sdk.F[string]("john@joinwarp.com"),
	FirstName:    sdk.F[string](""),
	LastName:     sdk.F[string](""),
	ManagerID:    sdk.F[string]("wrk_1234"),
	Position:     sdk.F[string](""),
	StartDate:    sdk.F[string]("2000-01-01"),
	WorkLocation: sdk.F[sdk.WorkerNewEmployeeParamsWorkLocationUnion](sdk.WorkerNewEmployeeParamsWorkLocationOfficeWorkLocation{
		WorkplaceID: sdk.F[string]("wkp_1234"),
	}),
})
if err != nil {
	panic(err)
}

fmt.Println(worker)
```

### Create Contractor

Create a new contractor. The worker will be created in draft status and must be invited separately via the invite endpoint. For business contractors, the businessName field is required.

| Direction | Type |
| --- | --- |
| Request | [`WorkerNewContractorParams`](./worker.go) |
| Response | [`WorkerNewContractorResponse`](./worker.go) |

```go
worker, err := client.Workers.NewContractor(context.Background(), sdk.WorkerNewContractorParams{
	DepartmentID: sdk.F[string]("dpt_1234"),
	Email:        sdk.F[string]("john@joinwarp.com"),
	FirstName:    sdk.F[string](""),
	LastName:     sdk.F[string](""),
	ManagerID:    sdk.F[string]("wrk_1234"),
	Position:     sdk.F[string](""),
	StartDate:    sdk.F[string]("2000-01-01"),
})
if err != nil {
	panic(err)
}

fmt.Println(worker)
```

### Invite Worker

Send or resend the worker invite so they can accept and complete onboarding to Warp. If the worker has already been invited, the invite will be resent with extended validity.

| Direction | Type |
| --- | --- |
| Response | [`WorkerInviteResponse`](./worker.go) |

```go
worker, err := client.Workers.Invite(context.Background(), "wrk_1234")
if err != nil {
	panic(err)
}

fmt.Println(worker)
```

## `Workplaces`

### List Workplaces

List all workplaces for your company.

| Direction | Type |
| --- | --- |
| Request | [`WorkplaceListParams`](./workplace.go) |
| Response | [`WorkplaceListResponse`](./workplace.go) |

```go
workplace, err := client.Workplaces.List(context.Background(), sdk.WorkplaceListParams{})
if err != nil {
	panic(err)
}

fmt.Println(workplace)
```

### Create Workplace

Create a new workplace.

| Direction | Type |
| --- | --- |
| Request | [`WorkplaceNewParams`](./workplace.go) |
| Response | [`WorkplaceNewResponse`](./workplace.go) |

```go
workplace, err := client.Workplaces.New(context.Background(), sdk.WorkplaceNewParams{
	Address: sdk.F[sdk.WorkplaceNewParamsAddress](sdk.WorkplaceNewParamsAddress{
		Line1:      sdk.F[string]("x"),
		City:       sdk.F[string](""),
		PostalCode: sdk.F[string](""),
	}),
	Name: sdk.F[string](""),
})
if err != nil {
	panic(err)
}

fmt.Println(workplace)
```

### Update Workplace

Update an existing workplace.

| Direction | Type |
| --- | --- |
| Request | [`WorkplaceUpdateParams`](./workplace.go) |
| Response | [`WorkplaceUpdateResponse`](./workplace.go) |

```go
workplace, err := client.Workplaces.Update(context.Background(), "wkp_1234", sdk.WorkplaceUpdateParams{})
if err != nil {
	panic(err)
}

fmt.Println(workplace)
```
