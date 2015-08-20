package api

import (
	"testing"
)

func Test_SetAdd(t *testing.T) {
	s := NewGSet()
	s.Add("k1").Add("k2").Add("k1")
	if s.Keys()[0] != "k1" {
		t.Error("Wrong keys")
	}
	if s.Keys()[1] != "k2" {
		t.Error("Wrong keys")
	}
}

func Test_SetAddX(t *testing.T) {
	var called int = 0
	var cb SimpleFunc = func(k string) { called = 1 }
	s := NewGSet()

	s.AddX("k1", cb)
	if called == 0 {
		t.Error("Callback wasnt called")
	}
	s.PrintIfNotInSet("k2")
	if s.Keys()[1] != "k2" {
		t.Error("Wrong keys: ", s.Keys())
	}
}
