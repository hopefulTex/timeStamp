package ui

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var formats []rune = []rune{
	't', 'T', 'd', 'D', 'f', 'F', 'R',
}

type FormatModel struct {
	isFocused bool
	index     int

	timestamp time.Time

	TimeStrings []string

	Width int
	Style lipgloss.Style
}

func initialFormatModel() FormatModel {
	return FormatModel{
		isFocused:   false,
		timestamp:   time.Now(),
		TimeStrings: formatArrayView(time.Now()),
		index:       0,
	}
}

func (m *FormatModel) Focus() {
	m.isFocused = true
}

func (m *FormatModel) Blur() {
	m.isFocused = false
}

func (m FormatModel) IsFocused() bool {
	return m.isFocused
}

func (m *FormatModel) SetTime(timestamp time.Time) {
	m.timestamp = timestamp
	m.TimeStrings = formatArrayView(m.timestamp)
}

func (m FormatModel) Update(msg tea.Msg) (FormatModel, tea.Cmd) {
	if !m.isFocused {
		return m, nil
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if m.index > 0 {
				m.index--
			}
		case "down":
			if m.index < len(formats)-1 {
				m.index++
			}
		}
	}

	return m, nil
}

func (m FormatModel) View() string {
	var view strings.Builder
	for i, str := range m.TimeStrings {
		if i == m.index {
			m.Style = m.Style.Background(selectedColor)
		}
		view.WriteString(m.Style.Render(str))
		view.WriteRune('\n')
		m.Style = m.Style.Background(defaultColor)
	}
	return m.Style.Render(strings.TrimLeft(view.String(), "\n"))
}

func formatArrayView(timestamp time.Time) []string {
	var views []string = []string{
		// short time      - hr:min
		TimeView(timestamp, 't'),
		TimeString(timestamp, 't'),
		// long time 	   - hr:min:sec
		TimeView(timestamp, 'T'),
		TimeString(timestamp, 'T'),
		// short date	   - day/month/year
		TimeView(timestamp, 'd'),
		TimeString(timestamp, 'd'),
		// long date	   - day monthName year
		TimeView(timestamp, 'D'),
		TimeString(timestamp, 'D'),
		// short date/time - day monthName year hr:min
		TimeView(timestamp, 'f'),
		TimeString(timestamp, 'f'),
		// long date/time  - dayName, day monthName year hr:min
		TimeView(timestamp, 'F'),
		TimeString(timestamp, 'F'),
		// relative time   - x time ago / in x time
		TimeView(timestamp, 'R'),
		TimeString(timestamp, 'R'),
	}

	viewWidth := 0
	stringWidth := 0
	for i, str := range views {
		if i%2 == 0 {
			if viewWidth < len(str) {
				viewWidth = len(str)
			}
		} else {
			if stringWidth < len(str) {
				stringWidth = len(str)
			}
		}
	}

	viewStyle := lipgloss.NewStyle().Width(viewWidth)
	stringStyle := viewStyle.Copy().Width(stringWidth)

	var pairs []string
	var tmp string
	var sep string = " " + lipgloss.NormalBorder().Right + " "

	for i, str := range views {
		if i%2 == 0 {
			tmp = viewStyle.Render(str)
		} else {
			tmp += sep + stringStyle.Render(str)
			pairs = append(pairs, tmp)
		}
	}

	return pairs
}
