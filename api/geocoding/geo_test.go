package geocoding

import (
	"context"
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	// TestNewErrorTimeout tests the newErrorTimeout function.
	t.Run("TestNewErrorTimeout", func(t *testing.T) {
		err := newErrorTimeout("timeout", nil)
		if err.Code != http.StatusRequestTimeout {
			t.Errorf("expected error code %d, got %d", http.StatusRequestTimeout, err.Code)
		}
		if err.Message != "geocoding: timeout" {
			t.Errorf("expected error message %s, got %s", "geocoding: timeout", err.Message)
		}
		if err.Err != nil {
			t.Errorf("expected error to be nil, got %v", err.Err)
		}
	})

	// TestErrIsTimeout tests the ErrIsTimeout function.
	t.Run("a nil context in params", func(t *testing.T) {
		long, lat, err := New(nil, "jakarta")
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if err.Error() != ErrNilContext.Error() {
			t.Errorf("expected error %s, got %s", ErrNilContext.Error(), err.Error())
		}
		if long != 0 {
			t.Errorf("expected long 0, got %f", long)
		}
		if lat != 0 {
			t.Errorf("expected lat 0, got %f", lat)
		}
	})

	// test if there is no query in params
	t.Run("test if there is no query in the query params", func(t *testing.T) {
		ctx := context.Background()
		long, lat, err := New(ctx, "")
		if err == nil {
			t.Errorf("expected a query parameter, got an empty string")
		}
		if err.Error() != ErrEmptyQuery.Error() {
			t.Errorf("expected error %s, got %s", ErrEmptyQuery.Error(), err.Error())
		}
		if long != 0 {
			t.Errorf("expected long 0, got %f", long)
		}
		if lat != 0 {
			t.Errorf("expected lat 0, got %f", lat)
		}
	})
}
