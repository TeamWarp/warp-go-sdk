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

func TestWorkplaceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Workplaces.New(context.TODO(), warphr.WorkplaceNewParams{
		Address: warphr.F(warphr.WorkplaceNewParamsAddress{
			City:       warphr.F("city"),
			Country:    warphr.F(warphr.WorkplaceNewParamsAddressCountryUs),
			Line1:      warphr.F("x"),
			PostalCode: warphr.F("postalCode"),
			State:      warphr.F(warphr.WorkplaceNewParamsAddressStateAl),
			Line2:      warphr.F("line2"),
		}),
		Name: warphr.F("name"),
		Type: warphr.F(warphr.WorkplaceNewParamsTypeRemote),
	})
	if err != nil {
		var apierr *warphr.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkplaceUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Workplaces.Update(
		context.TODO(),
		"wkp_1234",
		warphr.WorkplaceUpdateParams{
			Name: warphr.F("name"),
		},
	)
	if err != nil {
		var apierr *warphr.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkplaceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Workplaces.List(context.TODO(), warphr.WorkplaceListParams{
		AfterID:  warphr.F("wkp_1234"),
		BeforeID: warphr.F("wkp_1234"),
		Limit:    warphr.F("limit"),
	})
	if err != nil {
		var apierr *warphr.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
