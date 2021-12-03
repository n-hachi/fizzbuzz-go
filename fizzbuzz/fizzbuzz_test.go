package fizzbuzz

import "testing"

func TestFizzbuzz(t *testing.T) {
	// Check number string
	n1 := uint64(1)
	s1 := "1"
	if s1 != Fizzbuzz(n1) {
		t.Errorf("Fizzbuzz(1) = %v but expected %v\n", Fizzbuzz(n1), s1)
	}
}
