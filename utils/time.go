package utils

import "time"

const (
	DateTimeFormat         = time.RFC3339
	DateTimeFormatForInput = "02-01-2006T15:04"
	DateFormat             = "02-01-2006"
)

func ParseDateTimeForInput(v string) (time.Time, error) {
	return time.Parse(DateTimeFormatForInput, v)
}
