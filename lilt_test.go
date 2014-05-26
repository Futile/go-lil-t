package lilt

import "testing"

func TestLilT(t *testing.T) {
	If, IfNot := New(t)

	If(false).Errorf("omg")

	IfNot(true).Errorf("omg")
}
