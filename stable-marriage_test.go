package main

import (
	"fmt"
	"testing"
)

var SIMILAR_MEN = map[string][]string{
	"A": {"C", "D"},
	"B": {"C", "D"},
}

var SIMILAR_WOMEN = map[string][]string{
	"C": {"A", "B"},
	"D": {"A", "B"},
}
var OPPOSITE_MEN = map[string][]string{
	"A": {"C", "D"},
	"B": {"D", "C"},
}

var OPPOSITE_WOMEN = map[string][]string{
	"C": {"A", "B"},
	"D": {"B", "A"},
}

var OPPOSITE_WOMEN_2 = map[string][]string{
	"C": {"B", "A"},
	"D": {"A", "B"},
}

type StableMarriageSolutionTestCase struct {
	testcase   string
	men, women PreferenceMatrix
	isValid    bool
}

func runSolutionTest(tc StableMarriageSolutionTestCase, t *testing.T) {
	var foundPlan = Solve(tc.men, tc.women)
	var result, details = IsStableSolution(tc.men, tc.women, *foundPlan)

	if result != tc.isValid {
		t.Error(fmt.Sprintf("%s expected %t but got %t. %s", tc.testcase, tc.isValid, result, details))
	} else {
		t.Log(fmt.Sprintf("%s passed", tc.testcase))
	}
}

func TestSolve(t *testing.T) {
	runSolutionTest(StableMarriageSolutionTestCase{
		testcase: "ss",
		men:      SIMILAR_MEN,
		women:    SIMILAR_WOMEN,
		isValid:  true,
	},
		t)
	runSolutionTest(StableMarriageSolutionTestCase{
		testcase: "so2",
		men:      SIMILAR_MEN,
		women:    OPPOSITE_WOMEN_2,
		isValid:  true,
	},
		t)
	runSolutionTest(StableMarriageSolutionTestCase{
		testcase: "oo2",
		men:      OPPOSITE_MEN,
		women:    OPPOSITE_WOMEN_2,
		isValid:  true,
	},
		t)
}

type IsStableSolutionTestCase struct {
	testcase     string
	men, women   PreferenceMatrix
	expectedPlan MarriagePlan
	isValid      bool
}

func runIsStableTest(td IsStableSolutionTestCase, t *testing.T) {
	var result, details = IsStableSolution(td.men, td.women, td.expectedPlan)

	if result != td.isValid {
		t.Error(fmt.Sprintf("%s expected %t but got %t. %s", td.testcase, td.isValid, result, details))
	} else {
		t.Log(fmt.Sprintf("%s passed", td.testcase))
	}
}

func TestIsStable(t *testing.T) {

	runIsStableTest(IsStableSolutionTestCase{
		testcase: "valid test",
		men:      SIMILAR_MEN,
		women:    OPPOSITE_WOMEN,
		expectedPlan: MarriagePlan{
			{lhs: "A", rhs: "C"},
			{lhs: "B", rhs: "D"},
		},
		isValid: true,
	},
		t)

	runIsStableTest(IsStableSolutionTestCase{
		testcase: "invalid test",
		men:      SIMILAR_MEN,
		women:    OPPOSITE_WOMEN,
		expectedPlan: MarriagePlan{
			{lhs: "A", rhs: "C"},
			{lhs: "B", rhs: "C"},
		},
		isValid: false,
	},
		t)

	runIsStableTest(IsStableSolutionTestCase{
		testcase: "invalid test",
		men:      SIMILAR_MEN,
		women:    OPPOSITE_WOMEN,
		expectedPlan: MarriagePlan{
			{lhs: "A", rhs: "C"},
			{lhs: "B", rhs: "C"},
		},
		isValid: false,
	},
		t)

	runIsStableTest(IsStableSolutionTestCase{
		testcase: "unstable test",
		men:      SIMILAR_MEN,
		women:    OPPOSITE_WOMEN,
		expectedPlan: MarriagePlan{
			{lhs: "A", rhs: "D"},
			{lhs: "B", rhs: "C"},
		},
		isValid: false,
	},
		t)
}
