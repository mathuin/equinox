package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"

	"github.com/soniakeys/meeus/julian"
	"github.com/soniakeys/meeus/solstice"
)

// datefromJD in Python is julian.JDToCalendar(JD) here

func eventJDE(year, event int) (jd float64) {
	switch event {
	case 0:
		return solstice.March(year)
	case 1:
		return solstice.June(year)
	case 2:
		return solstice.September(year)
	case 3:
		return solstice.December(year)
	default:
		panic(fmt.Sprintf("unexpected event value %d", event))
	}
}

func cqafter(year, event int) (jd float64) {
	low := eventJDE(year, event)
	var high float64
	if event == 3 {
		high = eventJDE(year+1, 0)
	} else {
		high = eventJDE(year, event+1)
	}
	return (low + high) / 2
}

func pprintJD(jd float64) string {
	y, m, d := julian.JDToCalendar(jd)
	return fmt.Sprintf("%d %s %.2f", y, time.Month(m), d)
}

func main() {
	curyear, err := strconv.Atoi(time.Now().Format("2006"))
	if err != nil {
		fmt.Println(err)
	}
	var year = flag.Int("year", curyear, "year for which to calculate cross quarter dates")
	flag.Parse()

	extremes := [...]string{"March", "June", "September", "December"}
	var tag string

	for event, name := range extremes {
		if event%2 == 0 {
			tag = "equinox"
		} else {
			tag = "solstice"
		}
		fmt.Printf("%s %s is: %s\n", name, tag, pprintJD(eventJDE(*year, event)))
	}

	for event := range extremes {
		fmt.Printf("Cross quarter #%v for year %v is: %s\n", event, *year, pprintJD(cqafter(*year, event)))
	}
}
