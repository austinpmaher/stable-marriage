package main

import (
	"fmt"
)

type PreferenceMatrix map[string][]string

type MarriagePlan map[string]int

var MEN = map[string][]string{
	"A": {"C", "D"},
	"B": {"C", "D"},
}

var WOMEN = map[string][]string{
	"C": {"A", "B"},
	"D": {"B", "A"},
}

func main() {
	isValidSolution(solve(MEN, WOMEN))
	fmt.Println("done")
}

func solve(a PreferenceMatrix, b PreferenceMatrix) MarriagePlan {
	return nil
}

func isValidSolution(plan MarriagePlan) bool {
	return false
}
