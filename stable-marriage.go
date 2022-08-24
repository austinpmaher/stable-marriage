package main

import (
	"fmt"
)

type PreferenceMatrix map[string][]string

type Proposal struct {
	party1 string
	party2 string
}

type MarriagePlan []Proposal

var MEN = map[string][]string{
	"A": {"C", "D"},
	"B": {"C", "D"},
}

var WOMEN = map[string][]string{
	"C": {"A", "B"},
	"D": {"B", "A"},
}

func main() {
	var plan = solve(MEN, WOMEN)
	IsStableSolution(MEN, WOMEN, plan)
	fmt.Println("done")
}

func solve(a PreferenceMatrix, b PreferenceMatrix) MarriagePlan {
	return nil
}

func indexOf(slice []string, item string) int {
	for i := range slice {
		if slice[i] == item {
			return i
		}
	}
	return -1
}

func buildRejectList(subject string, option string, prefs PreferenceMatrix) []string {
	var idx = indexOf(prefs[subject], option)
	if idx < 0 {
		return make([]string, 0)
	} else {
		return prefs[subject][0:idx]
	}
}

func isCompleteSolution(p1Prefs PreferenceMatrix, p2Prefs PreferenceMatrix, plan MarriagePlan) bool {
	for party := range p1Prefs {
		count := 0
		for _, p := range plan {
			if p.party1 == party {
				count++
			}
		}
		if count != 1 {
			return false
		}
	}

	for party := range p2Prefs {
		count := 0
		for _, p := range plan {
			if p.party2 == party {
				count++
			}
		}
		if count != 1 {
			return false
		}
	}
	return true
}

func IsStableSolution(p1Prefs PreferenceMatrix, p2Prefs PreferenceMatrix, plan MarriagePlan) bool {

	if !isCompleteSolution(p1Prefs, p2Prefs, plan) {
		return false
	}

	var p1Rejects = make(map[string][]string)
	var p2Rejects = make(map[string][]string)

	for _, proposal := range plan {
		var p1 = proposal.party1
		var p2 = proposal.party2
		p1Rejects[p1] = buildRejectList(p1, p2, p1Prefs)
		p2Rejects[p2] = buildRejectList(p2, p1, p2Prefs)
	}

	for _, proposal := range plan {
		var p1 = proposal.party1
		var p1Rejects = p1Rejects[p1]
		for _, preferedChoice := range p1Rejects {
			if indexOf(p2Rejects[preferedChoice], p1) >= 0 {
				return false
			}
		}
	}

	return true
}
