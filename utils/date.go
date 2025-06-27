package utils

import "time"

func ParseStringToDate(date string) *time.Time {
	if date == "" {
		return nil
	}

	layouts := []string{
		"2006-01-02",
		"2006-01-02 15:04:05",
		"20026-01-02T15:04:05.000Z07:00",
		time.RFC3339,
		time.RFC3339Nano,
	}

	for _, layout := range layouts {
		if result, err := time.Parse(layout, date); err == nil {
			return &result
		}
	}

	return nil
}
