package main

import (
	"fmt"
)

type PreferenceMatrix map[string][]string

type Proposal struct {
	lhs string
	rhs string
}

type MarriagePlan []Proposal

/*
var MEN = map[string][]string{
	"A": {"C", "D"},
	"B": {"C", "D"},
}

var WOMEN = map[string][]string{
	"C": {"A", "B"},
	"D": {"B", "A"},
}
*/

var MEN = map[string][]string{
	"abe":  {"abi", "eve", "cath", "ivy", "jan", "dee", "fay", "bea", "hope", "gay"},
	"bob":  {"cath", "hope", "abi", "dee", "eve", "fay", "bea", "jan", "ivy", "gay"},
	"col":  {"hope", "eve", "abi", "dee", "bea", "fay", "ivy", "gay", "cath", "jan"},
	"dan":  {"ivy", "fay", "dee", "gay", "hope", "eve", "jan", "bea", "cath", "abi"},
	"ed":   {"jan", "dee", "bea", "cath", "fay", "eve", "abi", "ivy", "hope", "gay"},
	"fred": {"bea", "abi", "dee", "gay", "eve", "ivy", "cath", "jan", "hope", "fay"},
	"gav":  {"gay", "eve", "ivy", "bea", "cath", "abi", "dee", "hope", "jan", "fay"},
	"hal":  {"abi", "eve", "hope", "fay", "ivy", "cath", "jan", "bea", "gay", "dee"},
	"ian":  {"hope", "cath", "dee", "gay", "bea", "abi", "fay", "ivy", "jan", "eve"},
	"jon":  {"abi", "fay", "jan", "gay", "eve", "bea", "dee", "cath", "ivy", "hope"},
}

var WOMEN = map[string][]string{
	"abi":  {"bob", "fred", "jon", "gav", "ian", "abe", "dan", "ed", "col", "hal"},
	"bea":  {"bob", "abe", "col", "fred", "gav", "dan", "ian", "ed", "jon", "hal"},
	"cath": {"fred", "bob", "ed", "gav", "hal", "col", "ian", "abe", "dan", "jon"},
	"dee":  {"fred", "jon", "col", "abe", "ian", "hal", "gav", "dan", "bob", "ed"},
	"eve":  {"jon", "hal", "fred", "dan", "abe", "gav", "col", "ed", "ian", "bob"},
	"fay":  {"bob", "abe", "ed", "ian", "jon", "dan", "fred", "gav", "col", "hal"},
	"gay":  {"jon", "gav", "hal", "fred", "bob", "abe", "col", "ed", "dan", "ian"},
	"hope": {"gav", "jon", "bob", "abe", "ian", "dan", "hal", "ed", "col", "fred"},
	"ivy":  {"ian", "col", "hal", "gav", "fred", "bob", "abe", "ed", "jon", "dan"},
	"jan":  {"ed", "hal", "gav", "abe", "bob", "jon", "col", "ian", "fred", "dan"},
}

func main() {
	var plan = Solve(MEN, WOMEN)
	fmt.Printf("The man-optimal plan:\n%s\n", plan)
	ok, reason := IsStableSolution(MEN, WOMEN, *plan)
	if ok {
		fmt.Println(fmt.Sprintf("stable solution found. %s", reason))
	} else {
		fmt.Println(fmt.Sprintf("unstable / invalid solution found. %s", reason))
	}

	plan = Solve(WOMEN, MEN)
	fmt.Printf("The women-optimal plan:\n%s\n", plan)
	ok, reason = IsStableSolution(WOMEN, MEN, *plan)
	if ok {
		fmt.Println(fmt.Sprintf("stable solution found. %s", reason))
	} else {
		fmt.Println(fmt.Sprintf("unstable / invalid solution found. %s", reason))
	}
}

type PreferenceCursor struct {
	subject         string
	nextChoiceIndex int
	currentChoice   string
	priorities      []string
}

const NO_CHOICE = "NO_CHOICE"

func (cursor *PreferenceCursor) init(subject string, priorities []string) {
	cursor.subject = subject
	cursor.nextChoiceIndex = 0
	cursor.currentChoice = NO_CHOICE
	cursor.priorities = priorities
}

func (cursor *PreferenceCursor) isFree() bool {
	return cursor.currentChoice == NO_CHOICE
}

func (cursor *PreferenceCursor) peekChoice() (string, bool) {
	if cursor.nextChoiceIndex < len(cursor.priorities) {
		return cursor.priorities[cursor.nextChoiceIndex], true
	} else {
		return NO_CHOICE, false
	}

}

func (cursor *PreferenceCursor) nextChoice() (string, bool) {
	choice, ok := cursor.peekChoice()
	if ok {
		cursor.nextChoiceIndex++
		return choice, true
	} else {
		return NO_CHOICE, false
	}
}

func (cursor *PreferenceCursor) prefers(newChoice string, existingChoice string) bool {
	i := indexOf(cursor.priorities, newChoice)
	j := indexOf(cursor.priorities, existingChoice)
	return i < j
}

func engage(left *PreferenceCursor, right *PreferenceCursor) {
	left.currentChoice = right.subject
	right.currentChoice = left.subject
}

type Worklist []PreferenceCursor

func createWorklist(pm PreferenceMatrix) Worklist {
	var worklist = make([]PreferenceCursor, len(pm))

	var i = 0
	for k, v := range pm {
		worklist[i].init(k, v)
		// fmt.Println(fmt.Sprintf("address of cursor for %s: %p", k, &worklist[i]))
		i++
	}
	return worklist
}

func (worklist Worklist) nextFreeCursor() int {
	for idx, value := range worklist {
		if value.isFree() {
			return idx
		}
	}
	return -1
}

var NO_CURSOR_FOUND = PreferenceCursor{}

func cursorFor0(worklist Worklist, subject string) (*PreferenceCursor, bool) {
	// TODO: figure out why this returns pointers to a copy :-()
	for _, value := range worklist {
		if value.subject == subject {
			return &value, true
		}
	}
	return &NO_CURSOR_FOUND, false
}

func cursorFor(worklist Worklist, subject string) (int, bool) {
	// returning indexes works around a problem with the worklist
	// the incoming Worklist ([]PreferenceCursor) isn't sharing the
	// underlying array with the caller so the & returns the wrong address
	for idx, value := range worklist {
		if value.subject == subject {
			return idx, true
		}
	}
	return -1, false
}

func Solve(lhs PreferenceMatrix, rhs PreferenceMatrix) *MarriagePlan {
	var lhsWorklist = createWorklist(lhs)
	var rhsWorklist = createWorklist(rhs)

	for {
		idx := lhsWorklist.nextFreeCursor()
		if idx < 0 {
			break
		}
		lhsCursor := &lhsWorklist[idx]
		lhsSubject := lhsCursor.subject
		nextChoiceSubject, ok := lhsCursor.nextChoice()
		if !ok {
			panic(fmt.Sprintf("ran out of options for %s", lhsSubject))
		}
		nextChoiceIndex, found := cursorFor(rhsWorklist, nextChoiceSubject)
		if !found {
			panic(fmt.Sprintf("Cannot find a cursor for %s in %p", nextChoiceSubject, &rhsWorklist))
		}
		nextChoiceCursor := &rhsWorklist[nextChoiceIndex]

		if nextChoiceCursor.isFree() {
			fmt.Printf("Engage %s with %s\n", lhsSubject, nextChoiceSubject)
			engage(lhsCursor, nextChoiceCursor)
		} else { // some pair (m', w) already exists
			rhsCurrentChoice := nextChoiceCursor.currentChoice
			// is lhs a better match for nextChoice than her current choice
			if nextChoiceCursor.prefers(lhsSubject, rhsCurrentChoice) {
				otherCursorIndex, _ := cursorFor(lhsWorklist, rhsCurrentChoice)
				otherCursor := &lhsWorklist[otherCursorIndex]
				fmt.Printf("Clear engagement %s with %s\n", rhsCurrentChoice, nextChoiceSubject)
				otherCursor.currentChoice = NO_CHOICE
				fmt.Printf("Engage %s with %s\n", lhsSubject, nextChoiceSubject)
				engage(lhsCursor, nextChoiceCursor)
			} else {
				// go around the loop again and try their next choice
			}
		}
	}

	result := make(MarriagePlan, len(lhsWorklist))
	for idx, cursor := range lhsWorklist {
		proposal := &result[idx]
		proposal.lhs = cursor.subject
		proposal.rhs = cursor.currentChoice
	}

	return &result
}

func indexOf(slice []string, item string) int {
	for i := range slice {
		if slice[i] == item {
			return i
		}
	}
	return -1
}

func buildRejectList(lhs string, rhs string, prefs PreferenceMatrix) []string {
	var idx = indexOf(prefs[lhs], rhs)
	if idx < 0 {
		return make([]string, 0)
	} else {
		return prefs[lhs][0:idx]
	}
}

func isCompleteSolution(lhsPrefs PreferenceMatrix, rhsPrefs PreferenceMatrix, plan MarriagePlan) bool {
	for left := range lhsPrefs {
		count := 0
		for _, prop := range plan {
			if prop.lhs == left {
				count++
			}
		}
		if count != 1 {
			return false
		}
	}

	for party := range rhsPrefs {
		count := 0
		for _, prop := range plan {
			if prop.rhs == party {
				count++
			}
		}
		if count != 1 {
			return false
		}
	}
	return true
}

func IsStableSolution(lhsPrefs PreferenceMatrix, rhsPrefs PreferenceMatrix, plan MarriagePlan) (bool, string) {

	if !isCompleteSolution(lhsPrefs, rhsPrefs, plan) {
		return false, "not complete"
	}

	var lhsRejects = make(map[string][]string)
	var rhsRejects = make(map[string][]string)

	for _, proposal := range plan {
		var left = proposal.lhs
		var right = proposal.rhs
		lhsRejects[left] = buildRejectList(left, right, lhsPrefs)
		rhsRejects[right] = buildRejectList(right, left, rhsPrefs)
	}

	for _, proposal := range plan {
		var left = proposal.lhs
		var leftRejects = lhsRejects[left]
		for _, preferedChoice := range leftRejects {
			if indexOf(rhsRejects[preferedChoice], left) >= 0 {
				return false, fmt.Sprintf("%s would rather be with %s and vice versa", left, preferedChoice)
			}
		}
	}

	return true, "ok"
}
