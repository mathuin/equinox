package main

import (
	"math"
	"testing"
)

var epsilon = 0.000001 // comparing floats is hard

type testevent struct {
	year  int
	event [4]float64
}

var events = []testevent{
	{
		2006, [4]float64{2453815.268412, 2453908.018582, 2454001.669963, 2454091.516041},
	},
}

// func TestAverage(t *testing.T) {
//   for _, pair := range tests {
//     v := Average(pair.values)
//     if v != pair.average {
//       t.Error(
//         "For", pair.values,
//         "expected", pair.average,
//         "got", v,
//       )
//     }
//   }
// }

func TestEvents(t *testing.T) {
	for _, pair := range events {
		for i := 0; i < 4; i++ {
			a := eventJDE(pair.year, i)
			if math.Abs(a-pair.event[i]) > epsilon {
				t.Error(
					"For ", pair.year, "event #", i,
					"expected", pair.event[i],
					"got", a,
				)
			}
		}
	}
}

type testcross struct {
	year  int
	cross [4]float64
}

var crosses = []testcross{
	{
		2006, [4]float64{2453861.643497, 2453954.844272, 2454046.593002, 2454136.010965},
	},
}

func TestCrosses(t *testing.T) {
	for _, pair := range crosses {
		for i := 0; i < 4; i++ {
			a := cqafter(pair.year, i)
			if math.Abs(a-pair.cross[i]) > epsilon {
				t.Error(
					"For ", pair.year, "cross #", i,
					"expected", pair.cross[i],
					"got", a,
				)
			}
		}
	}
}

type testprint struct {
	value float64
	out   string
}

var prints = []testprint{
	{2453815.268412, "2006 March 20.77"},
	{2453908.018582, "2006 June 21.52"},
	{2454001.669963, "2006 September 23.17"},
	{2454091.516041, "2006 December 22.02"},
	{2453861.643497, "2006 May 6.14"},
	{2453954.844272, "2006 August 7.34"},
	{2454046.593002, "2006 November 7.09"},
	{2454136.010965, "2007 February 4.51"},
}

func TestPrints(t *testing.T) {
	for _, pair := range prints {
		a := pprintJD(pair.value)
		if a != pair.out {
			t.Error(
				"For ", pair.value,
				"expected", pair.out,
				"got", a,
			)
		}
	}
}
