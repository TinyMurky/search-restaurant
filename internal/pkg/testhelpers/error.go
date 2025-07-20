package testhelpers

import (
	"errors"
	"testing"
)

// AssertNoError assert err 應為 nil
func AssertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("expect nil, got error: %v", err)
	}
}

// AssertHasError assert err 應存在
func AssertHasError(t testing.TB, err error) {
	t.Helper()

	if err == nil {
		t.Error("expect err, got nil")
	}
}

// AssertAsError use errors.As to make sure error is certain type
func AssertAsError(t testing.TB, got error, wantErrorType any) {
	t.Helper()

	if !errors.As(got, wantErrorType) {
		t.Errorf("expect error with type %+v, got error %+v", wantErrorType, got)
	}
}
