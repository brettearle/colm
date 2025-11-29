package logger_test

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"

	"github.com/brettearle/colm/internal/logger"
)

func Test(t *testing.T) {
	t.Run("logger should put 'hello' in msg key", func(t *testing.T) {
		var buf bytes.Buffer
		var got io.Writer = &buf
		log := logger.New(got)

		log.Info("hello")

		var m map[string]any
		json.Unmarshal(buf.Bytes(), &m)

		expected := "hello"
		if m["msg"] != expected {
			t.Errorf("expected msg %q got %q", expected, m["msg"])
		}
	})
}
