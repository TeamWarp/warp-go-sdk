# Warp Go API

Complete reference of every operation, grouped by resource. See [the README](./README.md) for usage and configuration.

## Contents

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
- [`Departments`](#departments)
  - [List departments](#list-departments)
  - [Create department](#create-department)
  - [Update department](#update-department)
- [`Workplaces`](#workplaces)
  - [List workplaces](#list-workplaces)
  - [Create workplace](#create-workplace)
  - [Update workplace](#update-workplace)

## Setup

```go
import (
	"context"
	"fmt"

	sdk "github.com/marclave/warp-go-sdk"
)

client := sdk.NewClient()
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
| Request | [`TimeOffPolicyListParams`](./timeoffpolicy.go) |
| Response | [`TimeOffPolicyListResponse`](./timeoffpolicy.go) |

```go
policy, err := client.TimeOff.Policies.List(context.Background(), sdk.TimeOffPolicyListParams{})
if err != nil {
	panic(err)
}
fmt.Println(policy)
```

#### Get time off policy

Get a specific time off policy by id

| Direction | Type |
| --- | --- |
| Response | [`TimeOffPolicyGetResponse`](./timeoffpolicy.go) |

```go
policy, err := client.TimeOff.Policies.Get(context.Background(), "top_1234")
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
	WorkLocation: sdk.F[sdk.WorkerNewEmployeeParamsWorkLocationUnion](sdk.WorkerNewEmployeeParamsWorkLocationUnion{}),
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
