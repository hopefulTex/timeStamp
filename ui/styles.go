package ui

import "github.com/charmbracelet/lipgloss"

var (
	selectedColor      lipgloss.Color = lipgloss.Color("#562347")
	defaultColor       lipgloss.Color = lipgloss.Color("#262626")
	cellColor          lipgloss.Color = lipgloss.Color("#282828")
	blurredBorderColor lipgloss.Color = lipgloss.Color("#4660F6")
	focusedBorderColor lipgloss.Color = lipgloss.Color("#E060E0")
)
var (
	mainStyle     lipgloss.Style = lipgloss.NewStyle().Width(64).Border(lipgloss.RoundedBorder())
	cellStyle     lipgloss.Style = lipgloss.NewStyle().Width(2).Background(cellColor)
	selectedStyle lipgloss.Style = lipgloss.NewStyle().Background(selectedColor)
	widgetStyle   lipgloss.Style = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(blurredBorderColor)
)
