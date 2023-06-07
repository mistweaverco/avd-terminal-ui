package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type dashboardItem struct {
	title, desc string
}

func (i dashboardItem) Title() string       { return i.title }
func (i dashboardItem) Description() string { return i.desc }
func (i dashboardItem) FilterValue() string { return i.title }

func (m model) updateDashboard(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			selectedItem := m.dashboardList.SelectedItem().(dashboardItem).title
			tea.Println("Selected item: " + selectedItem)
		}
	case tea.WindowSizeMsg:
		h, v := getScreenDefaultStyles().GetFrameSize()
		globals.vh = msg.Height - v
		globals.vw = msg.Width - h
		m.dashboardList.SetSize(globals.vw, globals.vh)
	}
	var cmd tea.Cmd
	m.dashboardList, cmd = m.dashboardList.Update(msg)
	return m, cmd
}

func getDashboardModel() model {
	items := []list.Item{
		dashboardItem{title: "Pixel XL", desc: "Yummie"},
		dashboardItem{title: "Pixel XXL", desc: "Turbo Yummie"},
	}

	m := model{showView: "dashboard"}
	m.dashboardList = list.New(items, list.NewDefaultDelegate(), 0, 0)
	m.dashboardList.Title = "Dashboard"
	if globals.vw > 0 && globals.vh > 0 {
		m.dashboardList.SetSize(globals.vw, globals.vh)
	}
	return m
}