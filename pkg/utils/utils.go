package utils

import (
	"strings"
	"time"
)

func CreateDurIntoString(d time.Duration) string {
	s := d.String()
	if strings.HasSuffix(s, "m0 s") {
		s = s[:len(s)-2]
	}

	if strings.HasSuffix(s, "h0 m") {
		s = s[:len(s)-2]
	}

	return s
}
