package main_test

import (
	"context"
	"net/http"
	"os"
	"testing"
	"time"

	m "github.com/brettearle/colm/cmd/colm-broker"
	"github.com/brettearle/colm/cmd/colm-broker/internal/testutil"
)

func Test(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	t.Cleanup(cancel)
	go m.Run(ctx, m.Config{
		Host: "0.0.0.0",
		Port: "8080",
	}, os.Stderr)

	// Time for server ready
	testutil.WaitForReady(ctx, 2*time.Second, "http://0.0.0.0:8080/health")

	t.Run("/health returns 200", func(t *testing.T) {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://0.0.0.0:8080/health", nil)
		if err != nil {
			t.Errorf("Failed to create request")
		}
		client := http.Client{}
		res, err := client.Do(req)
		if err != nil {
			t.Errorf("Error sending request")
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("got %d want %d", res.StatusCode, http.StatusOK)
		}
	})
}
