package main

type Globals struct {
	vw                    int
	vh                    int
	selectedIssueIndex    int
	selectedIssueNumber   int
	selectedDashboardItem string
	selectedPullNumber    int
	runID                 int
}

var globals Globals
