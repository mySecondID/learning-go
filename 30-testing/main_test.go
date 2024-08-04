package main

import (
	"testing"
)

func TestMyFunc(t *testing.T) {
	actualValue := MyFunc()
	expectedValue := "hi from myfunc"
	if actualValue != expectedValue {
		t.Errorf("wrong answer, expected %s, got %s", expectedValue, actualValue)
	}
}

func TestMyMathFunc(t *testing.T) {
	actualValue := myMathFunc(-2, 3)
	expectedValue := 2
	if actualValue != expectedValue {
		t.Errorf("expected %d, got %d", expectedValue, actualValue)
	}
}
