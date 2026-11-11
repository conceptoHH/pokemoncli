package main

import (
	"testing"
	"reflect"
)

func TestCleanInput(t *testing.T) {

	type test struct {
		input string;
		expected []string;
	}

	tests := []test {
		{input: "  hello world   ", expected: []string{"hello","world"}},
		{input: "cosascosas", expected: []string{"cosascosas"}},
		{input: " tipo singular acaro-no", expected: []string{"tipo","singular","acaro-no"}},
		{input: "MAYUSCULAS", expected: []string{"mayusculas"}},
	}

	for _, ts := range tests {
		got := cleanInput(ts.input)
		if !reflect.DeepEqual(ts.expected, got){
			t.Fatalf("expected: %v, got: %v",ts.expected, got)
		}
	}
}
