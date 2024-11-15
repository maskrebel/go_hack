package unittest_test

import (
	. "go_hack/sample_unittest"
	"testing"
)

func TestAdd(t *testing.T) {
	res := Add(1, 2)
	if res != 3 {
		t.Logf("Expected 3, got %d", res)
		t.Fail()
	}
}
