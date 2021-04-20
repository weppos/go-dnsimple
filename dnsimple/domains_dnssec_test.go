package dnsimple

import (
	"context"
	"io"
	"net/http"
	"reflect"
	"testing"
)

func TestDnssecPath(t *testing.T) {
	if want, got := "/1010/domains/example.com/dnssec", dnssecPath("1010", "example.com"); want != got {
		t.Errorf("dnssecPath(%v) = %v, want %v", "", got, want)
	}
}

func TestDomainsService_EnableDnssec(t *testing.T) {
	setupMockServer()
	defer teardownMockServer()

	mux.HandleFunc("/v2/1010/domains/example.com/dnssec", func(w http.ResponseWriter, r *http.Request) {
		httpResponse := httpResponseFixture(t, "/api/enableDnssec/success.http")

		testMethod(t, r, "POST")
		testHeaders(t, r)

		w.WriteHeader(httpResponse.StatusCode)
	})

	accountID := "1010"

	_, err := client.Domains.EnableDnssec(context.Background(), accountID, "example.com")
	if err != nil {
		t.Fatalf("Domains.EnableDnssec() returned error: %v", err)
	}
}

func TestDomainsService_DisableDnssec(t *testing.T) {
	setupMockServer()
	defer teardownMockServer()

	mux.HandleFunc("/v2/1010/domains/example.com/dnssec", func(w http.ResponseWriter, r *http.Request) {
		httpResponse := httpResponseFixture(t, "/api/disableDnssec/success.http")

		testMethod(t, r, "DELETE")
		testHeaders(t, r)

		w.WriteHeader(httpResponse.StatusCode)
	})

	accountID := "1010"

	_, err := client.Domains.DisableDnssec(context.Background(), accountID, "example.com")
	if err != nil {
		t.Fatalf("Domains.DisableDnssec() returned error: %v", err)
	}
}

func TestDomainsService_GetDnssec(t *testing.T) {
	setupMockServer()
	defer teardownMockServer()

	mux.HandleFunc("/v2/1010/domains/example.com/dnssec", func(w http.ResponseWriter, r *http.Request) {
		httpResponse := httpResponseFixture(t, "/api/getDnssec/success.http")

		testMethod(t, r, "GET")
		testHeaders(t, r)

		w.WriteHeader(httpResponse.StatusCode)
		_, _ = io.Copy(w, httpResponse.Body)
	})

	dnssecResponse, err := client.Domains.GetDnssec(context.Background(), "1010", "example.com")
	if err != nil {
		t.Errorf("Domains.GetDnssec() returned error: %v", err)
	}

	dnssec := dnssecResponse.Data
	wantSingle := &Dnssec{Enabled: true}

	if !reflect.DeepEqual(dnssec, wantSingle) {
		t.Fatalf("Domains.GetDnssec() returned %+v, want %+v", dnssec, wantSingle)
	}
}
