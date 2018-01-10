// +build !nolaw

package law

import (
	"fmt"
	"reflect"
)

const (
	RequireType = "REQUIRE"
	EnsureType  = "ENSURE"
)

type Predicate func() bool

func Require(desc string, x bool) {
	if !x {
		fail(RequireType, fmt.Sprintf("%s: was not true", desc))
	}
}

func RequireEq(desc string, x, y interface{}) {
	if !reflect.DeepEqual(x, y) {
		fail(RequireType,
			fmt.Sprintf("%s: %v was not equal to %v", desc, x, y))
	}
}

func RequirePred(desc string, p Predicate) {
	if !p() {
		fail(RequireType,
			fmt.Sprintf("%s: provided predicate failed", desc))
	}
}

func Ensure(desc string, x bool) {
	if !x {
		fail(EnsureType, fmt.Sprintf("%s: was not true", desc))
	}
}

func EnsureEq(desc string, x, y interface{}) {
	if !reflect.DeepEqual(x, y) {
		fail(EnsureType,
			fmt.Sprintf("%s: %v was not equal to %v", desc, x, y))
	}
}

func EnsurePred(desc string, p Predicate) {
	if !p() {
		fail(EnsureType,
			fmt.Sprintf("%s: provided predicate failed", desc))
	}
}

func fail(typ, desc string) {
	panic(newAssertionError(typ, desc))
}
