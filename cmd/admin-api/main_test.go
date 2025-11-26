package main_test

import (
	"context"
	"net/http"
	"os"
	"testing"
	"time"

	m "github.com/brettearle/colm/cmd/admin-api"
	"github.com/brettearle/colm/cmd/admin-api/internal/testutil"
)

func Test(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	t.Cleanup(cancel)
	go m.Run(ctx, m.Config{
		Host: "0.0.0.0",
		Port: "8081",
	}, os.Stderr)

	// Time for server ready
	testutil.WaitForReady(ctx, 2*time.Second, "http://0.0.0.0:8081/health")

	t.Run("/health returns 200", func(t *testing.T) {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://0.0.0.0:8081/health", nil)
		if err != nil {
			t.Errorf("Failed to create request")
		}
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Errorf("Error sending request")
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			t.Errorf("got %d want %d", resp.StatusCode, http.StatusOK)
		}
	})
}
