// Package domain contains small framework-independent value helpers.
package domain

import (
	"strconv"
	"strings"
	"time"
)

// ParseDuration accepts Go durations and day suffixes such as 7d.
func ParseDuration(value string) (time.Duration, error) {
	value = strings.TrimSpace(strings.ToLower(value))
	if strings.HasSuffix(value, "d") {
		days, err := strconv.Atoi(strings.TrimSuffix(value, "d"))
		if err != nil {
			return 0, err
		}
		return time.Duration(days) * 24 * time.Hour, nil
	}
	return time.ParseDuration(value)
}
