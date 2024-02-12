package ui

// for use when model uses classical time.Time timestamps

// func cellString(index, kind int, t time.Time) string {
// 	var num int
// 	var s string
// 	if index == kind {
// 		cellStyle = cellStyle.Background(selectedColor)
// 	}

// 	switch kind {
// 	case 0:
// 		num = t.Hour()
// 	case 1:
// 		num = t.Minute()
// 	case 2:
// 		num = t.Second()
// 	case 3:
// 		num = t.Day()
// 	case 4:
// 		num = int(t.Month())
// 	case 5:
// 		cellStyle = cellStyle.Width(4)
// 		s = cellStyle.Render(fmt.Sprintf("%04d", t.Year()))
// 		cellStyle = cellStyle.Background(cellColor).Width(2)
// 		return s
// 	default:
// 		s = ""
// 	}

// 	s = cellStyle.Render(fmt.Sprintf("%02d", num))

// 	cellStyle = cellStyle.Background(cellColor)
// 	return s
// }

// func (m DateModel) View() string {
// 	var view strings.Builder

// 	view.WriteString(cellString(m.index, 0, m.Timestamp))
// 	view.WriteRune(':')
// 	view.WriteString(cellString(m.index, 1, m.Timestamp))
// 	view.WriteRune(':')
// 	view.WriteString(cellString(m.index, 2, m.Timestamp))
// 	view.WriteRune('\n')

// 	view.WriteString(cellString(m.index, 3, m.Timestamp))
// 	view.WriteRune('/')
// 	view.WriteString(cellString(m.index, 4, m.Timestamp))
// 	view.WriteRune('/')
// 	view.WriteString(cellString(m.index, 5, m.Timestamp))

// 	view.WriteRune('\n')
// 	view.WriteString(fmt.Sprintf("index: %d\n", m.index))
// 	return m.Style.Render(view.String())
// }
