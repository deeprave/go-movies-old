package test

import (
	"github.com/google/go-cmp/cmp"
	"go/types"
	"strings"
	"testing"
)

func difference(value1 interface{}, value2 interface{}, opts ...cmp.Option) string {
	return cmp.Diff(value1, value2, opts...)
}

func contains[V interface{} | types.Basic](array []V, value V, opts ...cmp.Option) bool {
	for _, v := range array {
		if cmp.Equal(v, value, opts...) {
			return true
		}
	}
	return false
}

//goland:noinspection GoUnusedExportedFunction
func ShouldBeTrue(t *testing.T, condition bool, format string, v ...any) bool {
	if condition {
		return true
	}
	t.Error(format, v)
	return false
}

//goland:noinspection GoUnusedExportedFunction
func ShouldBeFalse(t *testing.T, condition bool, format string, v ...any) bool {
	if !condition {
		return true
	}
	t.Error(format, v)
	return false
}

func ShouldBeEqual(t *testing.T, value1 interface{}, value2 interface{}, opts ...cmp.Option) bool {
	if cmp.Equal(value1, value2, opts...) {
		return true
	}
	t.Errorf("Failed: expecting equality:\n"+
		"\tExpected: '%v'\n"+
		"\tGot: '%v'\n\t--- diff ---\n%v",
		value1, value2, difference(value1, value2, opts...))
	return false
}

//goland:noinspection GoUnusedExportedFunction
func ShouldNotBeEqual(t *testing.T, value1 interface{}, value2 interface{}, opts ...cmp.Option) bool {
	if !cmp.Equal(value1, value2, opts...) {
		return true
	}
	t.Errorf("Failed: expecing inequality but found equal '%v'", value1)
	return false
}

func ShouldBeInArray[V interface{} | types.Basic](t *testing.T, array []V, value V, opts ...cmp.Option) bool {
	if contains(array, value, opts...) {
		return true
	}
	t.Errorf("Value '%v' not found in array", value)
	return false
}

//goland:noinspection GoUnusedExportedFunction
func ShouldNotBeInArray[V interface{} | string](t *testing.T, array []V, value V, opts ...cmp.Option) bool {
	if !contains(array, value, opts...) {
		return true
	}
	t.Errorf("Value '%v' found in array", value)
	return false
}

//goland:noinspection GoUnusedExportedFunction
func ShouldBeSubstring(t *testing.T, str string, sub string) bool {
	if strings.Contains(str, sub) {
		return true
	}
	t.Errorf("'%s' is not within '%s", sub, str)
	return false
}

//goland:noinspection GoUnusedExportedFunction
func ShouldNotBeSubString(t *testing.T, str string, sub string) bool {
	if !strings.Contains(str, sub) {
		return true
	}
	t.Errorf("'%s' is within '%s", sub, str)
	return false
}

func ShouldBeNoError(t *testing.T, err error, format string, v ...any) bool {
	if err != nil {
		t.Errorf(format, v...)
		return false
	}
	return true
}

//goland:noinspection GoUnusedExportedFunction
func ShouldBeError(t *testing.T, err error, format string, v ...any) bool {
	if err == nil {
		t.Errorf(format, v...)
		return false
	}
	return true
}
