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

func TestTimeOffListAssignmentsWithOptionalParams(t *testing.T) {
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
	_, err := client.TimeOff.ListAssignments(context.TODO(), warphr.TimeOffListAssignmentsParams{
		AfterID:   warphr.String("afterId"),
		BeforeID:  warphr.String("beforeId"),
		Limit:     warphr.String("limit"),
		PolicyIDs: []string{"top_1234"},
		WorkerIDs: []string{"wrk_1234"},
	})
	if err != nil {
		var apierr *warphr.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTimeOffListBalancesWithOptionalParams(t *testing.T) {
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
	_, err := client.TimeOff.ListBalances(context.TODO(), warphr.TimeOffListBalancesParams{
		AfterID:   warphr.String("afterId"),
		BeforeID:  warphr.String("beforeId"),
		EndDate:   warphr.String("endDate"),
		Limit:     warphr.String("limit"),
		PolicyIDs: []string{"top_1234"},
		StartDate: warphr.String("startDate"),
		WorkerIDs: []string{"wrk_1234"},
	})
	if err != nil {
		var apierr *warphr.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTimeOffListRequestsWithOptionalParams(t *testing.T) {
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
	_, err := client.TimeOff.ListRequests(context.TODO(), warphr.TimeOffListRequestsParams{
		AfterID:         warphr.String("afterId"),
		BeforeID:        warphr.String("beforeId"),
		EndsBefore:      warphr.String("endsBefore"),
		EndsOnOrAfter:   warphr.String("endsOnOrAfter"),
		Limit:           warphr.String("limit"),
		PolicyIDs:       []string{"top_1234"},
		StartsBefore:    warphr.String("startsBefore"),
		StartsOnOrAfter: warphr.String("startsOnOrAfter"),
		Statuses:        []string{"pending"},
		WorkerIDs:       []string{"wrk_1234"},
	})
	if err != nil {
		var apierr *warphr.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
