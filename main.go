package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type tickType int64

const maxTick tickType = 1 << 24

func (t *tickType) toFloat() float64 {
	return float64(*t) / float64(maxTick)
}

type timeAction struct {
	Name  string   `json:"name"`
	Ticks tickType `json:"ticks"`
}

type commandFunc func()

var (
	actions   []timeAction

	unitName string = "hours"
	scaleFactor float64 = float64(Day) / float64(time.Hour) 
)

func commandShow() {
	left := maxTick
	fmt.Printf("Total: %0.2f %s\n", left.toFloat()*scaleFactor, unitName)
	for _, a := range actions {
		fmt.Printf("Action [%s]: %0.2f %s\n", a.Name, a.Ticks.toFloat()*scaleFactor, unitName)
		left -= a.Ticks
	}
	fmt.Printf("Remaining: %0.2f %s\n", left.toFloat()*scaleFactor, unitName)
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
		act.Ticks = tickType(ticks * float64(maxTick) / float64(scaleFactor))
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
		fmt.Printf("usage: %s [options] <command>\n", os.Args[0])
		flag.PrintDefaults()

		fmt.Println("commands:")
		for commandName := range commands {
			fmt.Println(" ", commandName)
		}

		os.Exit(1)
	}

	flag.Parse()
	if err := parseFlags(); err != nil {
		fmt.Printf("Error: %v\n", err)
		flag.Usage()
		return
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
