package taxjar

import (
	"fmt"
	"log"
	"os"
	"testing"
)

var (
	testAPIKey = os.Getenv("TAXJAR_TEST_API_KEY")
	testClient *Client
)

func TestNew(t *testing.T) {
	var err error
	type testcase struct {
		apiKey string
		want   error
	}

	tcs := []testcase{
		{
			apiKey: testAPIKey,
			want:   nil,
		},
		{
			apiKey: "",
			want:   ErrEmptyAPIKey,
		},
	}

	for _, tc := range tcs {
		if _, err = New(tc.apiKey); err != tc.want {
			t.Fatalf("invalid error, expected %v and recieved %v", tc.want, err)
		}
	}
}

func TestClient_GetTaxes(t *testing.T) {
	var (
		c   *Client
		err error
	)

	if c, err = New(testAPIKey); err != nil {
		t.Fatal(err)
	}

	type testcase struct {
		req     TaxesRequest
		want    TaxesResponse
		wantErr error
	}

	tcs := []testcase{
		{
			req: TaxesRequest{
				FromCountry: "US",
				ToCountry:   "US",
				ToState:     "WA",
				ToZip:       "98686",
				Amount:      100,
				Shipping:    10,
			},
			want: TaxesResponse{
				AmountToCollect: 9.24,
			},
		},
		{
			req: TaxesRequest{
				FromCountry: "US",
				ToCountry:   "US",
				ToState:     "OR",
				ToZip:       "97035",
				Amount:      100,
				Shipping:    10,
			},
			want: TaxesResponse{
				AmountToCollect: 0,
			},
		},
	}

	for _, tc := range tcs {
		var resp *TaxesResponse
		resp, err = c.GetTaxes(tc.req)
		if toErrString(tc.wantErr) != toErrString(err) {
			t.Fatalf("invalid error, expected <%v and received <%v>", tc.wantErr, err)
		}

		if err != nil {
			continue
		}

		if tc.want.AmountToCollect != resp.AmountToCollect {
			t.Fatalf("invalid response: expected amount to be collected to be %f and received %f", tc.want.AmountToCollect, resp.AmountToCollect)
		}
	}
}

func ExampleNew() {
	var err error
	if testClient, err = New("[Shippo API Key]"); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Shippo Client has been initialized! %v\n", testClient)
}

func ExampleClient_GetTaxes() {
	var err error
	if testClient, err = New("[Shippo API Key]"); err != nil {
		log.Fatal(err)
	}

	var req TaxesRequest
	req.FromCountry = "US"
	req.ToCountry = "US"
	req.ToState = "OR"
	req.ToZip = "97035"
	req.Amount = 100
	req.Shipping = 10

	var resp *TaxesResponse
	if resp, err = testClient.GetTaxes(req); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Taxes response! %+v\n", resp)
}

func toErrString(err error) string {
	if err == nil {
		return ""
	}

	return err.Error()
}
