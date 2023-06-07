package main

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
)

type model struct {
	showView      string
	dashboardList list.Model
	bootViewport  viewport.Model
	altscreen     bool
}