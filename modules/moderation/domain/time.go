// Package domain contains moderation rules that do not depend on Discord or storage.
package domain

import (
	"strconv"
	"time"
)

// SnowflakeTime returns the creation time encoded in a Discord snowflake ID.
func SnowflakeTime(id string) time.Time {
	n, _ := strconv.ParseInt(id, 10, 64)
	ms := (n >> 22) + 1420070400000
	return time.Unix(ms/1000, (ms%1000)*1e6)
}
