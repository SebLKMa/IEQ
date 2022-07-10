package formulas

import (
	"fmt"
)

// StandardFormula defines a type to contain the standard formula and implement Score interface
type StandardFormula struct {
	name    string
	formula *Formula
}

// Setup implements interface Scorer.Setup()
// Setup will populate the scores ranges used by Score().
func (f *StandardFormula) Setup(n string, min float64, max float64) {
	f.formula = &Formula{}
	f.name = n
	f.formula.min = min
	f.formula.max = max
	f.formula.avg = (f.formula.min + f.formula.max) / 2
	f.formula.rrange = (f.formula.max - f.formula.avg) / f.formula.avg * 100

	// generate the relative ranges as a binary skiptree
	from := 0.0
	to := f.formula.rrange
	high := 100
	low := -10
	chunks := 11
	diff := (high - low) / chunks
	score := high
	for i := 1; i < chunks; i++ {
		//fmt.Printf("From:%g To:%g Score:%g\n", from, to, float64(score))
		f.formula.ranges.Insert(from, to, float64(score))
		score = score - diff
		from = to
		to += f.formula.rrange
	}
	f.formula.ranges.Insert(from, to, float64(score))
}

// Score implements interface Scorer.Score()
func (f *StandardFormula) Score(v float64) (result float64, ok bool) {

	// in target range
	if v >= f.formula.min && v <= f.formula.max {
		return 100, true
	}

	// higher than best average
	if v > f.formula.avg {
		rr := (v - f.formula.avg) / f.formula.avg * 100
		result, ok = f.formula.ranges.Search(rr)
	} else if v < f.formula.avg {
		rr := (f.formula.avg - v) / f.formula.avg * 100
		result, ok = f.formula.ranges.Search(rr)
	} else {
		return 0.0, false
	}
	return
}

// Name implements interface Scorer.Name()
func (f *StandardFormula) Name() string {
	return f.name
}

// ToString implements interface Stringer.ToString()
func (f *StandardFormula) ToString() string {
	return fmt.Sprintf("Formula:%s %s", f.name, stringify(f.formula))
}
