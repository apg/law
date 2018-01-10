package law

import (
	"strings"
	"testing"
)

func failedContract() {
	Require("x > 0", 0 > 0)
}

func TestRequire(t *testing.T) {
	exs := []string{
		"REQUIRE FAILED",
		"x > 0",
		"law.failedContract",
		"Call stack:",
		"law_test.go",
	}

	testContractFailure(t, func() {
		failedContract()
	}, exs)
}

func TestRequirePred(t *testing.T) {
	exs := []string{
		"REQUIRE FAILED",
		"thing is sorted",
		"law.TestRequirePred",
		"Call stack:",
		"law_test.go",
	}

	testContractFailure(t, func() {
		RequirePred("thing is sorted", func() bool { return false })
	}, exs)

}

func TestRequireEq(t *testing.T) {
	exs := []string{
		"REQUIRE FAILED",
		"1 == 0",
		"Call stack:",
		"law_test.go",
	}

	testContractFailure(t, func() {
		RequireEq("1 == 0", 1, 0)
	}, exs)

}

func TestEnsure(t *testing.T) {
	exs := []string{
		"ENSURE FAILED",
		"x > 0",
		"Call stack:",
		"law_test.go",
	}

	testContractFailure(t, func() {
		Ensure("x > 0", 0 > 0)
	}, exs)
}

func TestEnsurePred(t *testing.T) {
	exs := []string{
		"ENSURE FAILED",
		"thing is sorted",
		"Call stack:",
		"law_test.go",
	}

	testContractFailure(t, func() {
		EnsurePred("thing is sorted", func() bool { return false })
	}, exs)

}

func TestEnsureEq(t *testing.T) {
	exs := []string{
		"ENSURE FAILED",
		"1 == 0",
		"Call stack:",
		"law_test.go",
	}

	testContractFailure(t, func() {
		EnsureEq("1 == 0", 1, 0)
	}, exs)

}

func testContractFailure(t *testing.T, f func(), es []string) {
	err := capturePanic(f)
	if err == nil {
		t.Fatal("got nil, expected assertion failure")
	}

	msg := err.Error()

	for _, s := range es {
		if !strings.Contains(msg, s) {
			t.Fatalf("expected %q in error message, not found:\n%s", s, msg)
		}
	}
}

func capturePanic(f func()) (err error) {
	defer func() {
		x := recover()
		if e, ok := x.(error); ok {
			err = e
		}
	}()

	f()

	return
}
