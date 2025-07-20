package testhelpers

import (
	"reflect"
	"testing"
)

// AssertDeepEqual uses reflect deep equal to make sure got and want is equal
func AssertDeepEqual[T any](t testing.TB, got, want T) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("epect %v, got %v", want, got)
	}
}
