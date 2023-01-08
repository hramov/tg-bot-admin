package utils

import "testing"

func TestEqualSlice_Equal(t *testing.T) {
	s1 := []uint8{1, 2, 3}
	s2 := []uint8{1, 2, 3}
	e := EqualSlice(s1, s2)
	if !e {
		t.Errorf("should be equal")
	}
}

func TestEqualSlice_NotEqual(t *testing.T) {
	s1 := []uint8{1, 2, 3}
	s2 := []uint8{1, 2, 4}
	e := EqualSlice(s1, s2)
	if e {
		t.Errorf("should not be equal")
	}
}

func TestEqualSlice_NotEqual_DifferentLength(t *testing.T) {
	s1 := []uint8{1, 2, 3}
	s2 := []uint8{1, 2, 3, 4}
	e := EqualSlice(s1, s2)
	if e {
		t.Errorf("should not be equal")
	}
}
