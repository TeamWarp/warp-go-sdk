// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package warphr_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/TeamWarp/warp-go-sdk"
	"github.com/TeamWarp/warp-go-sdk/internal/testutil"
	"github.com/TeamWarp/warp-go-sdk/option"
)

func TestWorkerGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := warphr.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Workers.Get(context.TODO(), "wrk_1234")
	if err != nil {
		var apierr *warphr.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkerListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := warphr.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Workers.List(context.TODO(), warphr.WorkerListParams{
		AfterID:   warphr.String("wrk_1234"),
		BeforeID:  warphr.String("wrk_1234"),
		Limit:     warphr.String("limit"),
		Statuses:  []string{"draft"},
		Types:     []string{"employee"},
		WorkEmail: warphr.String("workEmail"),
	})
	if err != nil {
		var apierr *warphr.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkerDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := warphr.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Workers.Delete(context.TODO(), "wrk_1234")
	if err != nil {
		var apierr *warphr.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkerNewContractorWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := warphr.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Workers.NewContractor(context.TODO(), warphr.WorkerNewContractorParams{
		DepartmentID: "dpt_1234",
		Email:        "john@joinwarp.com",
		EntityType:   warphr.WorkerNewContractorParamsEntityTypeIndividual,
		FirstName:    "Melissa",
		LastName:     "Jones",
		ManagerID:    "wrk_1234",
		Position:     "Design Consultant",
		StartDate:    "2000-01-01",
		WorkCountry:  warphr.WorkerNewContractorParamsWorkCountryAd,
		BusinessName: warphr.String("Galt Enterprises, LLC"),
		Compensation: warphr.WorkerNewContractorParamsCompensation{
			Amount:   1,
			Currency: "USD",
			Per:      "hour",
		},
		PaySchedule: warphr.WorkerNewContractorParamsPayScheduleWeekly,
		ScopeOfWork: warphr.String("Frontend development for the customer dashboard"),
		WorkEmail:   warphr.String("john@joinwarp.com"),
	})
	if err != nil {
		var apierr *warphr.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkerNewEmployeeWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := warphr.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Workers.NewEmployee(context.TODO(), warphr.WorkerNewEmployeeParams{
		Compensation: warphr.WorkerNewEmployeeParamsCompensation{
			Amount: 1,
			Per:    "hour",
		},
		DepartmentID: "dpt_1234",
		Email:        "john@joinwarp.com",
		FirstName:    "Jonathan",
		LastName:     "Galt",
		ManagerID:    "wrk_1234",
		Position:     "Software Engineer",
		StartDate:    "2000-01-01",
		WorkLocation: warphr.WorkerNewEmployeeParamsWorkLocationUnion{
			OfWorkerNewEmployeesWorkLocationObject: &warphr.WorkerNewEmployeeParamsWorkLocationObject{
				Type:        "office",
				WorkplaceID: "wkp_1234",
			},
		},
		PaySchedule:       warphr.WorkerNewEmployeeParamsPayScheduleWeekly,
		RequireI9:         warphr.Bool(true),
		StateRegistration: warphr.WorkerNewEmployeeParamsStateRegistrationSelfManaged,
		StockOptions:      warphr.Float(10000),
		WorkEmail:         warphr.String("john@joinwarp.com"),
	})
	if err != nil {
		var apierr *warphr.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkerInvite(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := warphr.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Workers.Invite(context.TODO(), "wrk_1234")
	if err != nil {
		var apierr *warphr.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
