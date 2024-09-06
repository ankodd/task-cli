package utils

import "time"

func ParseDate(date time.Time) string {
	v := ""

	for _, c := range date.String() {
		if c == '.' {
			break
		}

		v += string(c)
	}

	return v
}
