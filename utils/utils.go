package utils

import "time"

func DateToString(date time.Time) string {
	return date.Format("20060201")
}
