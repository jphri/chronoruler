package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type tickType int64
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
		"add": commandAdd,
		"del": commandDel,
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

func commandAdd() {
	args := flag.Args()
	if len(args) < 3 {
		fmt.Printf("usage: %s name <time-in-unit>\n", args[0])
		return
	}

	act := timeAction{}
	act.Name = args[1]

	if ticks, err := strconv.ParseFloat(args[2], 32); err != nil {
		log.Fatal(err)
	} else {
		act.Ticks = tickType(ticks * float64(maxTick) / float64(scale))
	}

	actions = append(actions, act)
	if data, err := json.Marshal(actions); err != nil {
		log.Fatal(err)
	} else {
		if err := os.WriteFile(*flagConfig, data, 0644); err != nil {
			log.Fatal(err)
		}
	}

	commandShow()
}

func commandDel() {
	args := flag.Args()
	if len(args) < 2 {
		fmt.Printf("usage: %s name\n", args[0])
		return
	}

	for idx, act := range actions {
		if act.Name == args[1] {
			actions = append(actions[0:idx], actions[(idx+1):]...)
			break
		}
	}

	if data, err := json.Marshal(actions); err != nil {
		log.Fatal(err)
	} else {
		if err := os.WriteFile(*flagConfig, data, 0644); err != nil {
			log.Fatal(err)
		}
	}

	commandShow()
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
