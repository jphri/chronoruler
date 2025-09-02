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
	errUnknownUnit = errors.New("unknown unit")
)

var (
	unitsConfig = map[rune]struct{
		unitName string
		unitScale time.Duration
	} {
		'y': { "year", Year },
		'M': { "months", Month },
		'd': { "days", Day },
		'h': { "hours", time.Hour },
		'm': { "minutes", time.Minute },
	}

	flagConfig = flag.String("config", "", "config file path")

	flagUnit  = flag.String("unit", "h", "unit used")
	flagScale = flag.String("scale", "1y", "scale")

	commands = map[string]commandFunc{
		"show": commandShow,
		"add":  commandAdd,
		"del":  commandDel,
	}
)

func parseFlags() error {
	var (
		scaleTime time.Duration
		err error
	)

	if scaleTime, err = ParseDuration(*flagScale); err != nil {
		return err
	}

	rFlagUnit := []rune(*flagUnit)
	if len(rFlagUnit) == 0 {
		return errUnknownUnit
	}

	if u, ok := unitsConfig[rFlagUnit[0]]; !ok {
		return errUnknownUnit
	} else {
		unitName = u.unitName
		scaleFactor = float64(scaleTime) / float64(u.unitScale)
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
