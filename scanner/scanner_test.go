package main

import "testing"

func TestScanner(t *testing.T) {
	if err := Scanner(); err != nil {
		t.Errorf("Expected nil, got: %v", err)
	}
}

func BenchScanner(b *testing.B) {

}
