package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

type tickType int32
const maxTick tickType = 1 << 24

func (t *tickType) toFloat() float32 {
	return float32(*t) / float32(maxTick)
}

type timeAction struct {
	Name string `json:"name"`
	Ticks tickType `json:"ticks"`
}

type commandFunc func()

var (
	actions []timeAction
	scaleName string
	scale float32
)

var (
	mode = map[string]float32 {
		"hpy": 24 * 365,
		"mpd": 24 * 60,
		"hpd": 24,
	}

	modeName = map[rune]string {
		'h': "hours",
		'm': "minutes",
	}

	flagMode = flag.String("mode", "hpy", "set mode")
	flagConfig = flag.String("config", "", "config file path")

	commands = map[string]commandFunc {
		"show": commandShow,
	}
)


func commandShow() {
	left := maxTick
	fmt.Printf("Total: %0.2f %s\n", left.toFloat() * scale, scaleName)
	for _, a := range actions {
		fmt.Printf("Action [%s]: %0.2f %s\n", a.Name, a.Ticks.toFloat() * scale, scaleName)
		left -= a.Ticks
	}
	fmt.Printf("Remaining: %0.2f %s\n", left.toFloat() * scale, scaleName)
}

func main() {
	flag.Usage = func() {
		fmt.Printf("usage: %s [options] <command>\n", os.Args[0]);
		flag.PrintDefaults()

		fmt.Println("commands:")
		for commandName := range commands {
			fmt.Println(" ", commandName)
		}
		
		os.Exit(1)
	}

	flag.Parse()
	if s, ok := mode[*flagMode]; !ok {
		flag.Usage()
		return
	} else {
		scale = s
		scaleName = modeName[[]rune(*flagMode)[0]]
	}

	if *flagConfig == "" {
		flag.Usage()
		return
	} else {
		bytes, err := os.ReadFile(*flagConfig)
		if err != nil {
			log.Fatal(err)
			return
		}
		if err := json.Unmarshal(bytes, &actions); err != nil {
			log.Fatal(err)
			return
		}
	}

	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		return
	}

	if f, ok := commands[args[0]]; ok {
		f()
	} else {
		flag.Usage()
	}
}
