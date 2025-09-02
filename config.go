package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	errUnknownMode = errors.New("unknown mode value")
)

var (
	mode = map[string]float64{
		"hpy": float64(Year / time.Hour),
		"mpd": float64(Day / time.Minute),
		"hpd": float64(Day / time.Hour),
	}

	modeName = map[rune]string{
		'h': "hours",
		'm': "minutes",
	}

	flagMode   = flag.String("mode", "hpy", "set mode")
	flagConfig = flag.String("config", "", "config file path")

	commands = map[string]commandFunc{
		"show": commandShow,
		"add":  commandAdd,
		"del":  commandDel,
	}
)

func parseFlags() error {
	if s, ok := mode[*flagMode]; !ok {
		return errUnknownMode
	} else {
		scale = s
		scaleName = modeName[[]rune(*flagMode)[0]]
	}

	if bytes, err := os.ReadFile(*flagConfig); err != nil {
		return fmt.Errorf("cannot open '%s': %w", *flagConfig, err)
	} else {
		if err := json.Unmarshal(bytes, &actions); err != nil {
			return fmt.Errorf("cannot read '%s': %w", *flagConfig, err)
		}
	}

	return nil
}
