package main

import (
	"fmt"
	"testing"
)

type StableMarriageTestData struct {
	testcase     string
	men, women   PreferenceMatrix
	expectedPlan MarriagePlan
	isValid      bool
}

func runTest(td StableMarriageTestData, t *testing.T) {
	var result, details = IsStableSolution(td.men, td.women, td.expectedPlan)

	if result != td.isValid {
		t.Error(fmt.Sprintf("%s expected %t but got %t. %s", td.testcase, td.isValid, result, details))
	} else {
		t.Log(fmt.Sprintf("%s passed", td.testcase))
	}
}

func TestIsStable(t *testing.T) {

	var MEN = map[string][]string{
		"A": {"C", "D"},
		"B": {"C", "D"},
	}

	var WOMEN = map[string][]string{
		"C": {"A", "B"},
		"D": {"B", "A"},
	}

	runTest(StableMarriageTestData{
		testcase: "valid test",
		men:      MEN,
		women:    WOMEN,
		expectedPlan: MarriagePlan{
			{lhs: "A", rhs: "C"},
			{lhs: "B", rhs: "D"},
		},
		isValid: true,
	},
		t)

	runTest(StableMarriageTestData{
		testcase: "invalid test",
		men:      MEN,
		women:    WOMEN,
		expectedPlan: MarriagePlan{
			{lhs: "A", rhs: "C"},
			{lhs: "B", rhs: "C"},
		},
		isValid: false,
	},
		t)

	runTest(StableMarriageTestData{
		testcase: "unstable test",
		men:      MEN,
		women:    WOMEN,
		expectedPlan: MarriagePlan{
			{lhs: "A", rhs: "D"},
			{lhs: "B", rhs: "C"},
		},
		isValid: false,
	},
		t)
}
