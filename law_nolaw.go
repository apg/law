// +build nolaw

package law

type Predicate func() bool

func Require(desc string, x bool) {

}

func RequireEq(desc string, x interface{}, y interface{}) {

}

func RequirePred(desc string, p Predicate) {

}

func Ensure(desc string, x bool) {

}

func EnsureEq(desc string, x bool) {

}

func EnsurePred(desc string, p Predicate) {

}
