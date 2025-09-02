package main

import (
	"errors"
	"time"
)

var (
	Day   = time.Hour * 24
	Month = Day * 30
	Year  = Day * 365
)

var (
	errParseDuration = errors.New("duration parse failed")
)

func durationParseFloat(s string) (r float64, state string, err error) {
	state = s

	for len(state) > 0 {
		c := state[0]
		if c < '0' || c > '9' {
			break
		}
		r = r*10.0 + float64(c-'0')
		state = state[1:]
	}

	if state[0] == '.' {
		state = state[1:]
		i := 10.0
		for len(state) > 0 {
			c := state[0]
			if c < '0' || c > '9' {
				break
			}
			r = r + float64(c-'0')/i
			state = state[1:]
			i *= 10.0
		}
	}

	if state == s {
		err = errParseDuration
	}
	return
}

func durationParseMode(s string) (r time.Duration, state string, err error) {
	state = s
	c := state[0]

	switch c {
	case 'y':
		r = Year

	case 'M':
		r = Month

	case 'd':
		r = Day

	case 'h':
		r = time.Hour

	case 'm':
		r = time.Minute

	case 's':
		r = time.Second

	default:
		err = errParseDuration
		return
	}

	state = state[1:]
	return
}

func ParseDuration(s string) (time.Duration, error) {
	var err error
	t := time.Duration(0)

	// (<float><timeMode>)+

	state := s
	for len(state) > 0 {
		var (
			r    float64
			mode time.Duration
		)

		r, state, err = durationParseFloat(state)
		if err != nil {
			return time.Duration(0), err
		}

		mode, state, err = durationParseMode(state)
		if err != nil {
			return time.Duration(0), err
		}

		t += time.Duration(r * float64(mode))
	}

	return t, nil
}

