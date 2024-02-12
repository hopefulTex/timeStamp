package ui

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TimeBlock struct {
	Year   int
	Month  int
	Day    int
	Hour   int
	Minute int
	Second int
}

func (t *TimeBlock) FromTime(timestamp time.Time) {
	t.Year = timestamp.Year()
	t.Month = int(timestamp.Month())
	t.Day = timestamp.Day()

	t.Hour = timestamp.Hour()
	t.Minute = timestamp.Minute()
	t.Second = timestamp.Second()
}

type DateModel struct {
	isFocused bool
	index     int

	input textinput.Model

	timeArray []*int
	Time      TimeBlock

	err error

	typing    bool
	hasInited bool
	Width     int
	Style     lipgloss.Style
}

func initialDateModel() DateModel {
	t := time.Now()
	tme := TimeBlock{}
	tme.FromTime(t)

	input := textinput.New()
	input.CharLimit = 2
	input.Prompt = ""
	input.Width = 2
	input.SetValue(fmt.Sprintf("%02d", tme.Hour))
	input.CursorEnd()

	//input.Validate = validateFunc

	return DateModel{
		isFocused: false,
		Time:      tme,
		input:     input,
		hasInited: false,
		typing:    false,
		index:     0,
	}
}

func (m *DateModel) Focus() {
	m.isFocused = true
	m.input.Focus()
}

func (m *DateModel) Blur() {
	m.isFocused = false
	m.typing = false
	num, err := strconv.Atoi(strings.TrimSpace(m.input.Value()))
	if err != nil {
		m.err = err
	} else {
		*m.timeArray[m.index] = num
		m.err = fmt.Errorf("set timeArray[%d] to %d", m.index, num)
	}
	format := "%02d"
	if m.index == 5 {
		format = "%04d"
		m.input.CharLimit = 4
		m.input.Width = 4
	} else {
		m.input.CharLimit = 2
		m.input.Width = 2
	}

	m.input.SetValue(fmt.Sprintf(format, *m.timeArray[m.index]))
	m.input.Blur()
}

func (m DateModel) IsFocused() bool {
	return m.isFocused
}

func (m DateModel) Error() string {
	return m.err.Error()
}

func (m DateModel) Update(msg tea.Msg) (DateModel, tea.Cmd) {
	var cmd tea.Cmd

	// init but with pointers so we cant do it in the constructor
	if !m.hasInited {
		m.hasInited = true
		m.timeArray = []*int{
			&m.Time.Hour,
			&m.Time.Minute,
			&m.Time.Second,
			&m.Time.Day,
			&m.Time.Month,
			&m.Time.Year,
		}
	}
	if !m.isFocused {
		return m, nil
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "left", "right", "up", "down":

			m.typing = false

			num, err := strconv.Atoi(strings.TrimSpace(m.input.Value()))
			if err != nil {
				m.err = err
			} else {
				*m.timeArray[m.index] = num
				m.err = fmt.Errorf("set timeArray[%d] to %d", m.index, num)
			}

			switch msg.String() {

			case "left":
				if m.index > 0 {
					m.index--
				}
			case "right":
				if m.index < 5 {
					m.index++
				}
			case "up":
				m.index -= 3
				if m.index < 0 {
					m.index = 0
				}
			case "down":
				m.index += 3
				if m.index > 5 {
					m.index = 5
				}
			}
			format := "%02d"
			if m.index == 5 {
				format = "%04d"
				m.input.CharLimit = 4
				m.input.Width = 4
			} else {
				m.input.CharLimit = 2
				m.input.Width = 2
			}

			m.input.SetValue(fmt.Sprintf(format, *m.timeArray[m.index]))

		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			if !m.typing {
				m.input.SetValue("")
				m.typing = true
			}
			v := m.input.Value()
			if len(v) == m.input.Width {
				m.input.SetValue(v[1:] + msg.String())
			}
			fallthrough
		default:
			m.input, cmd = m.input.Update(msg)
		}
	}

	return m, cmd
}
func (m DateModel) View() string {
	var view strings.Builder

	m.input.CursorEnd()

	view.WriteString("Time\n")

	for i, val := range m.timeArray {
		if i != m.index {
			view.WriteString(cellString(m.index, i, *val))
		} else {
			view.WriteString(selectedStyle.Render(m.input.View()[:m.input.Width]))
		}

		switch {
		case i < 2:
			view.WriteRune(':')
		case i == 2, i == 5:
			if i == 2 {
				view.WriteString("\nDate")
			}
			view.WriteRune('\n')
		case i > 2:
			view.WriteRune('/')
		}
	}

	s := strings.TrimSpace(view.String())

	return m.Style.Render(s)
}

func cellString(index, kind int, t int) string {
	var s string
	if index == kind {
		cellStyle = cellStyle.Background(selectedColor)
	}

	if kind == 5 {
		cellStyle = cellStyle.Width(4)
		s = cellStyle.Render(fmt.Sprintf("%04d", t))
		cellStyle = cellStyle.Background(cellColor).Width(2)
		return s
	}

	s = cellStyle.Render(fmt.Sprintf("%02d", t))

	cellStyle = cellStyle.Background(cellColor)
	return s
}

func (m *DateModel) TimeStamp() time.Time {
	s := m.String()
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		m.err = err
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		return time.Unix(0, 0)
	}
	return t
}

func (m DateModel) String() string {
	_, offset := time.Now().Zone()
	// seconds to hours
	offset /= (60 * 60)
	negative := offset < 0

	if negative {
		offset *= -1
	}

	offsetString := fmt.Sprintf("%02d:00", offset)
	if !negative {
		offsetString = "+" + offsetString
	} else {
		offsetString = "-" + offsetString
	}

	// (year)-(month)-(day)T(hour):(minute):(second)(timezoneoffset):00
	s := fmt.Sprintf("%04d-%02d-%02dT%02d:%02d:%02d%s",
		*m.timeArray[5], *m.timeArray[4], *m.timeArray[3],
		*m.timeArray[0], *m.timeArray[1], *m.timeArray[2],
		offsetString)

	return s
}
