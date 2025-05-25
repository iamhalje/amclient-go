package amclient

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAlerts(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v2/alerts" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		w.Write([]byte(`[{"labels":{"alertname":"TestAlert"}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL)
	alerts, err := client.GetAlerts()
	if err != nil {
		t.Fatal(err)
	}
	if len(alerts) != 1 || alerts[0].Labels["alertname"] != "TestAlert" {
		t.Fatalf("unexpected alerts: %+v", alerts)
	}
}

func TestGetStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v2/status" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		w.Write([]byte(`{"versionInfo":{"version":"0.25.0"}}`))
	}))
	defer server.Close()

	client := NewClient(server.URL)
	status, err := client.GetStatus()
	if err != nil {
		t.Fatal(err)
	}
	if status.VersionInfo.Version != "0.25.0" {
		t.Fatalf("unexpected version: %s", status.VersionInfo.Version)
	}
}
