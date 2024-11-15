package unittest_test

import (
	. "go_hack/sample_unittest"
	"testing"
)

func TestAdd(t *testing.T) {
	testTable := []struct {
		a, b, expected int
	}{
		{1, 1, 2},
		{3, 1, 4},
		{0, 0, 0},
	}

	for _, test := range testTable {
		res := Add(test.a, test.b)
		if res != test.expected {
			t.Logf("a: %d, b: %d", test.a, test.b)
			t.Logf("Expected %d, got %d", test.expected, res)
			t.Fail()
		}
	}

}

func TestSubHierarchy(t *testing.T) {
	// run by hierarchy (parallel)
	t.Run("a:positive", func(t *testing.T) {
		a := 10
		t.Run("b:positive", func(t *testing.T) {
			b := -5
			expect := a + b
			res := Add(a, b)
			if res != expect {
				t.Logf("a: %d, b: %d", a, b)
				t.Logf("Expected %d, got %d", expect, res)
				t.Fail()
			}
		})
		t.Run("b:negative", func(t *testing.T) {
			b := -5
			expect := a + b
			res := Add(a, b)
			if res != expect {
				t.Logf("a: %d, b: %d", a, b)
				t.Logf("Expected %d, got %d", expect, res)
				t.Fail()
			}
		})
	})
}
