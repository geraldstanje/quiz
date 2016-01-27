package lcwsolver

import (
	"testing"
)

func TestLCWSolver(t *testing.T) {
	s := New()
	s.AppendWord("apple")
	s.AppendWord("applebanana")
	s.AppendWord("applestrawberrybanana")
	s.AppendWord("banana")
	s.AppendWord("bananaxbanana")
	s.AppendWord("blueberry")
	s.AppendWord("strawberry")

	result := s.GetResult()
	if result != "applestrawberrybanana" {
		t.Error("Expected applestrawberrybanana, got ", result)
	}
}
