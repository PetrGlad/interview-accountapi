package form3

import "testing"

func TestDoSomething(t *testing.T) {
	cutoff := 12
	if DoSomething(cutoff) <= cutoff {
		t.Errorf("DoSomething(%d) = %d, want something larger than %d", cutoff, cutoff, cutoff)
	}
}
