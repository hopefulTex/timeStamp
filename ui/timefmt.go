package ui

import (
	"fmt"
	"time"
)

func TimeView(timestamp time.Time, mode rune) string {
	var s string
	switch mode {
	case 't': // Short Time
		s = "15:04"
	case 'T': // Long Time
		s = time.TimeOnly
	case 'd': // Short Date
		s = "02/01/2006"
	case 'D': // Long Date
		s = "02 January 2006"
	case 'F': // Long Date/Time
		s = "Monday, January 2006 15:04"
	case 'R': // Relative Time
		return RelativeTimeString(timestamp)
	default: // Short Date/Time
		s = "2 Jan 2006 15:04"
	}

	stamp := timestamp.Format(s)
	return stamp
}

func TimeString(timestamp time.Time, mode rune) string {
	return fmt.Sprintf("<t:%d:%c>", timestamp.Unix(), mode)
}

func RelativeTimeString(timestamp time.Time) string {
	durText := ""
	duration := time.Since(timestamp)
	sec := duration.Seconds()

	if sec < 0 {
		sec *= -1
	}

	if sec >= 60 {
		sec /= 60
		if sec >= 60 {
			sec /= 60
			if sec >= 24 {
				sec /= 24
				if sec >= 30 {
					sec /= 30
					if sec > 12 {
						sec /= 12
						durText = "years"
					} else {
						durText = "months"
					}
				} else {
					durText = "days"
				}
			} else {
				durText = "hours"
			}
		} else {
			durText = "minutes"
		}
	} else {
		durText = "seconds"
	}

	if int(sec) == 1 {
		durText = durText[:len(durText)-1]
	}

	durText = fmt.Sprintf("%d %s", int(sec), durText)

	if duration.Seconds() > 0 {
		durText += " ago"
	} else {
		durText = "in " + durText
	}

	return durText
}
