package ui

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	datePicker   DateModel
	formatPicker FormatModel

	selectedTime time.Time
	mode         rune
	copied       bool

	size   tea.WindowSizeMsg
	zone   int
	format int

	quitting bool
}

func New() *tea.Program {
	return tea.NewProgram(initialModel())
}

func initialModel() Model {
	// d := ?.New()
	f := initialFormatModel()
	d := initialDateModel()
	d.Focus()
	return Model{
		datePicker:   d,
		formatPicker: f,
		zone:         0,
		format:       0,
		copied:       false,
		quitting:     false,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.quitting {
		return m, tea.Quit
	}
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.QuitMsg:
		m.quitting = true
		return m, nil
	case tea.WindowSizeMsg:
		m.size = msg
		mainStyle = mainStyle.Width(m.size.Width - 2)
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+q", "esc":
			m.quitting = true
			return m, nil
		case "tab":
			if m.datePicker.IsFocused() {
				m.formatPicker.Focus()
				m.datePicker.Blur()
				m.formatPicker.SetTime(m.datePicker.TimeStamp())
			} else {
				m.formatPicker.Blur()
				m.datePicker.Focus()
			}

		case "enter":
			if m.datePicker.IsFocused() {
				m.formatPicker.Focus()
				m.datePicker.Blur()
				m.formatPicker.SetTime(m.datePicker.TimeStamp())
			} else {
				m.copied = true
				// clipboardCopy(timeString(m.selectedTime, m.mode))
			}
		}
	}

	m.datePicker, cmd = m.datePicker.Update(msg)
	cmds = append(cmds, cmd)
	m.formatPicker, cmd = m.formatPicker.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

var borderVertical string = lipgloss.NormalBorder().Right

func (m Model) View() string {
	if m.quitting {
		return "\n"
	}
	var view strings.Builder

	if m.datePicker.IsFocused() {
		widgetStyle = widgetStyle.BorderForeground(focusedBorderColor)
	}

	panel1 := widgetStyle.Render(m.datePicker.View())

	// view.WriteString(widgetStyle.Render(m.datePicker.View()))
	// view.WriteString("\n\n")

	if m.formatPicker.IsFocused() {
		widgetStyle = widgetStyle.BorderForeground(focusedBorderColor)
	} else {
		if m.datePicker.IsFocused() {
			widgetStyle = widgetStyle.BorderForeground(blurredBorderColor)
		}
	}

	panel2 := widgetStyle.Render(m.formatPicker.View())

	// view.WriteString(widgetStyle.Render(m.formatPicker.View()))
	// view.WriteRune('\n')

	widgetStyle = widgetStyle.BorderForeground(blurredBorderColor)

	height := max(lipgloss.Height(panel1), lipgloss.Height(panel2))
	middleBorder := ""
	for i := 0; i < height; i++ {
		middleBorder += borderVertical + "\n"
	}

	view.WriteString(lipgloss.JoinHorizontal(lipgloss.Center, panel1, middleBorder, panel2))

	viewString := mainStyle.Render(view.String()) + "\n"

	if m.datePicker.err != nil {
		viewString += m.datePicker.Error()
	}

	if m.copied {
		viewString += "Copied to clipboard\n"
	}
	return viewString
}
