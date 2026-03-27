// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package warphr_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/warp-hr-go"
	"github.com/stainless-sdks/warp-hr-go/internal/testutil"
	"github.com/stainless-sdks/warp-hr-go/option"
)

func TestTimeOffPolicyGet(t *testing.T) {
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
	_, err := client.TimeOff.Policies.Get(context.TODO(), "top_1234")
	if err != nil {
		var apierr *warphr.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTimeOffPolicyListWithOptionalParams(t *testing.T) {
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
	_, err := client.TimeOff.Policies.List(context.TODO(), warphr.TimeOffPolicyListParams{
		AfterID:  warphr.String("top_1234"),
		BeforeID: warphr.String("top_1234"),
		Limit:    warphr.String("limit"),
	})
	if err != nil {
		var apierr *warphr.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
