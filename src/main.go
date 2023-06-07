package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/pkg/browser"
)

var ANDROID_HOME string = os.ExpandEnv("$ANDROID_HOME")

var teaProgram *tea.Program

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "escape", "q":
			return m, tea.Quit
		}
	}

	switch m.showView {
	case "boot":
		return m.updateBoot(msg)
	case "dashboard":
		return m.updateDashboard(msg)
	default:
		return m, nil
	}
}

func (m model) View() string {
	s := getScreenDefaultStyles()
	switch m.showView {
	case "boot":
		return s.Render(m.bootViewport.View())
	case "dashboard":
		return s.Render(m.dashboardList.View())
	default:
		return "View not fould"
	}
}

func main() {
	var err error

	browser.Stdout = ioutil.Discard
	browser.Stderr = ioutil.Discard

	if err != nil {
		panic(err)
	}

	m := getDashboardModel()

	teaProgram = tea.NewProgram(m, tea.WithAltScreen())

	if err := teaProgram.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}