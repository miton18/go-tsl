package tsl

import (
	"math"
	"strings"
	"time"
)

// NilTime is used to not specify a time
var NilTime = time.Time{}

// ToMicroSeconds return a time in micro seconds
func toMicroSeconds(t time.Time) int64 {
	return int64(math.Trunc(float64(t.UnixNano() / 1000)))
}

func shortDur(d time.Duration) string {
	s := d.String()
	if strings.HasSuffix(s, "m0s") {
		s = s[:len(s)-2]
	}
	if strings.HasSuffix(s, "h0m") {
		s = s[:len(s)-2]
	}
	return s
}

// Eq look for exact match of the label value
func Eq(labelValue string) string {
	return "=" + labelValue
}

// NotEq look for label value not equal to `value`
func NotEq(labelValue string) string {
	return "!=" + labelValue
}

// Like regex to match of the label value
func Like(labelValue string) string {
	return "~" + labelValue
}

// NotLike regex to not match the label value
func NotLike(value string) string {
	return "!~" + value
}

func toStringArray(a []string) string {
	s := "["
	for i, e := range a {
		s += "\"" + e + "\""
		if i != len(a) {

		}
	}
	return s + "]"
}
