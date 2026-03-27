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
		AfterID:   warphr.F("wrk_1234"),
		BeforeID:  warphr.F("wrk_1234"),
		Limit:     warphr.F("limit"),
		Statuses:  warphr.F([]warphr.WorkerListParamsStatus{warphr.WorkerListParamsStatusDraft}),
		Types:     warphr.F([]warphr.WorkerListParamsType{warphr.WorkerListParamsTypeEmployee}),
		WorkEmail: warphr.F("workEmail"),
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
		DepartmentID: warphr.F("dpt_1234"),
		Email:        warphr.F("john@joinwarp.com"),
		EntityType:   warphr.F(warphr.WorkerNewContractorParamsEntityTypeIndividual),
		FirstName:    warphr.F("Melissa"),
		LastName:     warphr.F("Jones"),
		ManagerID:    warphr.F("wrk_1234"),
		Position:     warphr.F("Design Consultant"),
		StartDate:    warphr.F("2000-01-01"),
		WorkCountry:  warphr.F(warphr.WorkerNewContractorParamsWorkCountryAd),
		BusinessName: warphr.F("Galt Enterprises, LLC"),
		Compensation: warphr.F(warphr.WorkerNewContractorParamsCompensation{
			Amount:   warphr.F(1.000000),
			Currency: warphr.F(warphr.WorkerNewContractorParamsCompensationCurrencyUsd),
			Per:      warphr.F(warphr.WorkerNewContractorParamsCompensationPerHour),
		}),
		PaySchedule: warphr.F(warphr.WorkerNewContractorParamsPayScheduleWeekly),
		ScopeOfWork: warphr.F("Frontend development for the customer dashboard"),
		WorkEmail:   warphr.F("john@joinwarp.com"),
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
		Compensation: warphr.F(warphr.WorkerNewEmployeeParamsCompensation{
			Amount: warphr.F(1.000000),
			Per:    warphr.F(warphr.WorkerNewEmployeeParamsCompensationPerHour),
		}),
		DepartmentID: warphr.F("dpt_1234"),
		Email:        warphr.F("john@joinwarp.com"),
		FirstName:    warphr.F("Jonathan"),
		LastName:     warphr.F("Galt"),
		ManagerID:    warphr.F("wrk_1234"),
		Position:     warphr.F("Software Engineer"),
		StartDate:    warphr.F("2000-01-01"),
		WorkLocation: warphr.F[warphr.WorkerNewEmployeeParamsWorkLocationUnion](warphr.WorkerNewEmployeeParamsWorkLocationOfficeWorkLocation{
			Type:        warphr.F(warphr.WorkerNewEmployeeParamsWorkLocationOfficeWorkLocationTypeOffice),
			WorkplaceID: warphr.F("wkp_1234"),
		}),
		PaySchedule:       warphr.F(warphr.WorkerNewEmployeeParamsPayScheduleWeekly),
		RequireI9:         warphr.F(true),
		StateRegistration: warphr.F(warphr.WorkerNewEmployeeParamsStateRegistrationSelfManaged),
		StockOptions:      warphr.F(10000.000000),
		WorkEmail:         warphr.F("john@joinwarp.com"),
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
