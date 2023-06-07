package main

import (
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m model) updateBoot(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func getBootModel() model {
	var contents []string
	const width = 78
	vp := viewport.New(width, 20)
	vp.Style = lipgloss.NewStyle().
		PaddingRight(2)
	contents = append(contents, "Checking for updates...")

	vp.SetContent(strings.Join(contents, "\n"))
	m := model{showView: "boot", bootViewport: vp}
	return m
}