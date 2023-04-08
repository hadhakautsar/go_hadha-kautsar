package kalkulator

import (
	"testing"
)

func TestAddition(t *testing.T) {
	result := Additon(2, 3)
	if result != 5 {
		t.Errorf("Addition test failed, expected: %d, got: %d", 5, result)
	}
}

func TestSubstract(t *testing.T) {
	result := Substract(5, 2)
	if result != 3 {
		t.Errorf("Substract test failed, expected: %d, got: %d", 3, result)
	}
}

func TestMult(t *testing.T) {
	result := Mult(4, 6)
	if result != 24 {
		t.Errorf("Mult test failed, expected: %d, got: %d", 24, result)
	}
}

func TestDiv(t *testing.T) {
	result := Div(10, 2)
	if result != 5 {
		t.Errorf("Div test failed, expected: %d, got: %d", 5, result)
	}
}

func TestDivZero(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("DivZero test failed, expected panic, but got none")
		}
	}()
	Div(10, 0)
}
