# Warp Go API

Complete reference of every operation, grouped by resource. See [the README](./README.md) for usage and configuration.

## Contents

- [`CustomWorkerFields`](#customworkerfields)
  - [List custom worker fields](#list-custom-worker-fields)
  - [Create custom worker field](#create-custom-worker-field)
  - [Get custom worker field](#get-custom-worker-field)
  - [Update custom worker field](#update-custom-worker-field)
  - [Archive custom worker field](#archive-custom-worker-field)
  - [Create field option](#create-field-option)
  - [Update field option](#update-field-option)
  - [Delete unused field option](#delete-unused-field-option)
  - [Archive field option](#archive-field-option)
  - [List worker custom field values](#list-worker-custom-field-values)
  - [Set worker custom field value](#set-worker-custom-field-value)
  - [Clear worker custom field value](#clear-worker-custom-field-value)
- [`Departments`](#departments)
  - [List departments](#list-departments)
  - [Create department](#create-department)
  - [Update department](#update-department)
- [`Offers`](#offers)
  - [List offers](#list-offers)
  - [Create offer](#create-offer)
  - [Void offer](#void-offer)
  - [Extend offer deadline](#extend-offer-deadline)
  - [Resend offer](#resend-offer)
- [`TimeOff`](#timeoff)
  - [List time off assignments](#list-time-off-assignments)
  - [List time off balances](#list-time-off-balances)
  - [List time off requests](#list-time-off-requests)
  - [`TimeOff Policies`](#timeoff-policies)
    - [List time off policies](#list-time-off-policies)
    - [Get time off policy](#get-time-off-policy)
- [`Workers`](#workers)
  - [List workers](#list-workers)
  - [Get worker](#get-worker)
  - [Delete worker](#delete-worker)
  - [Create employee](#create-employee)
  - [Create contractor](#create-contractor)
  - [Invite worker](#invite-worker)
- [`Workplaces`](#workplaces)
  - [List workplaces](#list-workplaces)
  - [Create workplace](#create-workplace)
  - [Update workplace](#update-workplace)

## Setup

```go
import (
	"context"
	"fmt"

	sdk "github.com/TeamWarp/warp-sdk-go"
)

client := sdk.NewClient()
```

## `CustomWorkerFields`

### List custom worker fields

List the custom worker field definitions your API key can read. Each field belongs to a worker-data category; fields whose category your key cannot read are omitted unless the key holds workers:custom_fields.

| Direction | Type |
| --- | --- |
| Response | [`[]CustomWorkerFieldListResponse`](./customworkerfield.go) |

```go
customWorkerField, err := client.CustomWorkerFields.List(context.Background())
if err != nil {
	panic(err)
}
fmt.Println(customWorkerField)
```

### Create custom worker field

Create a custom worker field definition. The field type is immutable after creation. Select and multi_select fields can include their initial options. Access to values derives from the field category; requires the workers:custom_fields permission.

| Direction | Type |
| --- | --- |
| Request | [`CustomWorkerFieldNewParams`](./customworkerfield.go) |
| Response | [`CustomWorkerFieldNewResponse`](./customworkerfield.go) |

```go
customWorkerField, err := client.CustomWorkerFields.New(context.Background(), sdk.CustomWorkerFieldNewParams{
	Name: sdk.F[string](""),
})
if err != nil {
	panic(err)
}
fmt.Println(customWorkerField)
```

### Get custom worker field

Get a custom worker field definition, including its select options. Archived options may appear on existing worker values but cannot be newly selected.

| Direction | Type |
| --- | --- |
| Response | [`CustomWorkerFieldGetResponse`](./customworkerfield.go) |

```go
customWorkerField, err := client.CustomWorkerFields.Get(context.Background(), "cf_1234")
if err != nil {
	panic(err)
}
fmt.Println(customWorkerField)
```

### Update custom worker field

Update a custom worker field definition. The field type cannot be changed; create a new field instead. Requires the workers:custom_fields permission; changing the category, access level, or input source requires the manage level.

| Direction | Type |
| --- | --- |
| Request | [`CustomWorkerFieldUpdateParams`](./customworkerfield.go) |
| Response | [`CustomWorkerFieldUpdateResponse`](./customworkerfield.go) |

```go
customWorkerField, err := client.CustomWorkerFields.Update(context.Background(), "cf_1234", sdk.CustomWorkerFieldUpdateParams{})
if err != nil {
	panic(err)
}
fmt.Println(customWorkerField)
```

### Archive custom worker field

Archive a custom worker field. Archived fields keep their existing worker values but cannot receive new ones. Requires the workers:custom_fields permission at the manage level.

| Direction | Type |
| --- | --- |
| Response | [`CustomWorkerFieldArchiveResponse`](./customworkerfield.go) |

```go
customWorkerField, err := client.CustomWorkerFields.Archive(context.Background(), "cf_1234")
if err != nil {
	panic(err)
}
fmt.Println(customWorkerField)
```

### Create field option

Add an option to a select or multi_select custom worker field. The option value should be treated as stable; the label can change. Requires the workers:custom_fields permission.

| Direction | Type |
| --- | --- |
| Request | [`CustomWorkerFieldNewOptionParams`](./customworkerfield.go) |
| Response | [`CustomWorkerFieldNewOptionResponse`](./customworkerfield.go) |

```go
customWorkerField, err := client.CustomWorkerFields.NewOption(context.Background(), "cf_1234", sdk.CustomWorkerFieldNewOptionParams{
	Label: sdk.F[string]("x"),
	Value: sdk.F[string]("x"),
})
if err != nil {
	panic(err)
}
fmt.Println(customWorkerField)
```

### Update field option

Update the label or sort order of a custom worker field option. Options of archived fields cannot be edited. Requires the workers:custom_fields permission.

| Direction | Type |
| --- | --- |
| Request | [`CustomWorkerFieldUpdateOptionParams`](./customworkerfield.go) |
| Response | [`CustomWorkerFieldUpdateOptionResponse`](./customworkerfield.go) |

```go
customWorkerField, err := client.CustomWorkerFields.UpdateOption(context.Background(), "cfo_1234", sdk.CustomWorkerFieldUpdateOptionParams{})
if err != nil {
	panic(err)
}
fmt.Println(customWorkerField)
```

### Delete unused field option

Delete a custom worker field option that is not applied to any worker. Options in use must be archived instead. Requires the workers:custom_fields permission at the manage level.

```go
err := client.CustomWorkerFields.DeleteOption(context.Background(), "cfo_1234")
if err != nil {
	panic(err)
}
```

### Archive field option

Archive a custom worker field option. Archived options remain on existing worker values but cannot be newly selected. Requires the workers:custom_fields permission at the manage level.

| Direction | Type |
| --- | --- |
| Response | [`CustomWorkerFieldArchiveOptionResponse`](./customworkerfield.go) |

```go
customWorkerField, err := client.CustomWorkerFields.ArchiveOption(context.Background(), "cfo_1234")
if err != nil {
	panic(err)
}
fmt.Println(customWorkerField)
```

### List worker custom field values

List custom field values for workers, optionally filtered by worker or field. Values are returned only for fields whose category your API key can read.

| Direction | Type |
| --- | --- |
| Request | [`CustomWorkerFieldListValuesParams`](./customworkerfield.go) |
| Response | [`[]CustomWorkerFieldListValuesResponse`](./customworkerfield.go) |

```go
customWorkerField, err := client.CustomWorkerFields.ListValues(context.Background(), sdk.CustomWorkerFieldListValuesParams{})
if err != nil {
	panic(err)
}
fmt.Println(customWorkerField)
```

### Set worker custom field value

Create or replace a worker's value for a custom field. The value shape must match the field type, and your API key must hold write on the field's category.

| Direction | Type |
| --- | --- |
| Request | [`CustomWorkerFieldUpsertValueParams`](./customworkerfield.go) |
| Response | [`CustomWorkerFieldUpsertValueResponse`](./customworkerfield.go) |

```go
customWorkerField, err := client.CustomWorkerFields.UpsertValue(context.Background(), sdk.CustomWorkerFieldUpsertValueParams{
	FieldID: sdk.F[string]("cf_1234"),
	Value: sdk.F[sdk.CustomWorkerFieldUpsertValueParamsValueUnion](sdk.CustomWorkerFieldUpsertValueParamsValueUnion{}),
	WorkerID: sdk.F[string]("wrk_1234"),
})
if err != nil {
	panic(err)
}
fmt.Println(customWorkerField)
```

### Clear worker custom field value

Remove a worker's value for a custom field. Your API key must hold write on the field's category.

| Direction | Type |
| --- | --- |
| Request | [`CustomWorkerFieldClearValueParams`](./customworkerfield.go) |

```go
err := client.CustomWorkerFields.ClearValue(context.Background(), sdk.CustomWorkerFieldClearValueParams{
	FieldID: sdk.F[string]("cf_1234"),
	WorkerID: sdk.F[string]("wrk_1234"),
})
if err != nil {
	panic(err)
}
```

## `Departments`

### List departments

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

### Create department

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

### Update department

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

### List offers

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

### Create offer

Create and send a candidate offer. The candidate receives an email with a link to the offer portal.

| Direction | Type |
| --- | --- |
| Request | [`OfferNewParams`](./offer.go) |
| Response | [`OfferNewResponse`](./offer.go) |

```go
offer, err := client.Offers.New(context.Background(), sdk.OfferNewParams{
	Candidate: sdk.F[sdk.OfferNewParamsCandidate](sdk.OfferNewParamsCandidate{
		FirstName: sdk.F[string]("x"),
		LastName: sdk.F[string]("x"),
		Email: sdk.F[string]("john@joinwarp.com"),
	}),
	Compensation: sdk.F[sdk.OfferNewParamsCompensation](sdk.OfferNewParamsCompensation{
		PayRate: sdk.F[float64](0),
	}),
	Position: sdk.F[sdk.OfferNewParamsPosition](sdk.OfferNewParamsPosition{
		Title: sdk.F[string]("x"),
		StartDate: sdk.F[string]("2000-01-01"),
	}),
})
if err != nil {
	panic(err)
}
fmt.Println(offer)
```

### Void offer

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

### Extend offer deadline

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

### Resend offer

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

### List time off assignments

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

### List time off balances

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

### List time off requests

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

#### List time off policies

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

#### Get time off policy

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

### List workers

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

### Get worker

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

### Delete worker

Delete a worker. Only workers who have not yet completed onboarding can be deleted. Active workers must be properly offboarded.

```go
err := client.Workers.Delete(context.Background(), "wrk_1234")
if err != nil {
	panic(err)
}
```

### Create employee

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
	Email: sdk.F[string]("john@joinwarp.com"),
	FirstName: sdk.F[string](""),
	LastName: sdk.F[string](""),
	ManagerID: sdk.F[string]("wrk_1234"),
	Position: sdk.F[string](""),
	StartDate: sdk.F[string]("2000-01-01"),
	WorkLocation: sdk.F[sdk.WorkerNewEmployeeParamsWorkLocationUnion](sdk.WorkerNewEmployeeParamsWorkLocationOfficeWorkLocation{
		WorkplaceID: sdk.F[string]("wkp_1234"),
	}),
})
if err != nil {
	panic(err)
}
fmt.Println(worker)
```

### Create contractor

Create a new contractor. The worker will be created in draft status and must be invited separately via the invite endpoint. For business contractors, the businessName field is required.

| Direction | Type |
| --- | --- |
| Request | [`WorkerNewContractorParams`](./worker.go) |
| Response | [`WorkerNewContractorResponse`](./worker.go) |

```go
worker, err := client.Workers.NewContractor(context.Background(), sdk.WorkerNewContractorParams{
	DepartmentID: sdk.F[string]("dpt_1234"),
	Email: sdk.F[string]("john@joinwarp.com"),
	FirstName: sdk.F[string](""),
	LastName: sdk.F[string](""),
	ManagerID: sdk.F[string]("wrk_1234"),
	Position: sdk.F[string](""),
	StartDate: sdk.F[string]("2000-01-01"),
})
if err != nil {
	panic(err)
}
fmt.Println(worker)
```

### Invite worker

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

### List workplaces

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

### Create workplace

Create a new workplace.

| Direction | Type |
| --- | --- |
| Request | [`WorkplaceNewParams`](./workplace.go) |
| Response | [`WorkplaceNewResponse`](./workplace.go) |

```go
workplace, err := client.Workplaces.New(context.Background(), sdk.WorkplaceNewParams{
	Address: sdk.F[sdk.WorkplaceNewParamsAddress](sdk.WorkplaceNewParamsAddress{
		Line1: sdk.F[string]("x"),
		City: sdk.F[string](""),
		PostalCode: sdk.F[string](""),
	}),
	Name: sdk.F[string](""),
})
if err != nil {
	panic(err)
}
fmt.Println(workplace)
```

### Update workplace

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
