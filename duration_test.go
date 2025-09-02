package main

import (
	"testing"
	"time"
)

func TestParseDuration(t *testing.T) {
	testStrings := []struct {
		str string
		expect time.Duration
	} {
		{ str: "1d", expect: Day },
		{ str: "1M", expect: Month },
		{ str: "1y", expect: Year },
		{ str: "1h", expect: time.Hour },
		{ str: "1m", expect: time.Minute },
		{ str: "1s", expect: time.Second },
		{ str: "1M15d", expect: Month + 15 * Day },
		{ str: "1y4M15d10h15m20s", expect: Year + 4 * Month + 15 * Day + 10 * time.Hour + 15 * time.Minute + 20 * time.Second },
		{ str: "1.5y", expect: time.Duration(float64(Year) * 1.5) },
	}
	
	for _, ts := range testStrings {
		t.Run(ts.str, func(t *testing.T) {
			r, err := ParseDuration(ts.str)
			if err != nil {
				t.Fatalf("err = %e", err)
			}

			if r != ts.expect {
				t.Fatalf("ParseDuration(%s) != %d (%d)", ts.str, ts.expect, r)
			}
		})
	}
}
