package main

import (
	"testing"
)

func TestIsStable(t *testing.T) {

	var MEN = map[string][]string{
		"A": {"C", "D"},
		"B": {"C", "D"},
	}

	var WOMEN = map[string][]string{
		"C": {"A", "B"},
		"D": {"B", "A"},
	}

	var expected = []Proposal{
		{party1: "A", party2: "C"},
		{party1: "B", party2: "D"},
	}

	var stable = IsStableSolution(MEN, WOMEN, expected)
	if !stable {
		t.Error("stable returned false")
	} else {
		t.Log("ok")
	}

	var invalid = []Proposal{
		{party1: "A", party2: "C"},
		{party1: "B", party2: "C"},
	}

	var isValid = IsStableSolution(MEN, WOMEN, invalid)
	if !isValid {
		t.Log("ok")
	} else {
		t.Error("isInValid returned ok")
	}

	var unstable = []Proposal{
		{party1: "A", party2: "D"},
		{party1: "B", party2: "C"},
	}

	var isStable = IsStableSolution(MEN, WOMEN, unstable)
	if !isStable {
		t.Log("ok")
	} else {
		t.Error("unstable returned ok")
	}
}
