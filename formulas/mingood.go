package formulas

import (
	"fmt"
)

// MinIsGoodFormula defines a type to contain the minimum is good formula and implement Score interface
type MinIsGoodFormula struct {
	mingood StandardFormula
}

// Setup implements interface Scorer.Setup().
// Setup will populate the scores ranges used by Score().
func (f *MinIsGoodFormula) Setup(n string, min float64, max float64) {
	f.mingood.formula = &Formula{}
	f.mingood.name = n
	f.mingood.formula.min = min
	f.mingood.formula.max = max
	f.mingood.formula.avg = (f.mingood.formula.min + f.mingood.formula.max) / 2
	f.mingood.formula.rrange = (f.mingood.formula.max - f.mingood.formula.avg) / f.mingood.formula.avg * 100

	// generate the relative ranges as a binary skiptree
	from := 0.0
	to := f.mingood.formula.rrange
	high := 100
	low := -10
	chunks := 11
	diff := (high - low) / chunks
	score := high
	for i := 1; i < chunks; i++ {
		//fmt.Printf("From:%g To:%g Score:%g\n", from, to, float64(score))
		f.mingood.formula.ranges.Insert(from, to, float64(score))
		score = score - diff
		from = to
		to += f.mingood.formula.rrange
	}
	f.mingood.formula.ranges.Insert(from, to, float64(score))
}

// Score implements interface Scorer.Score()
func (f *MinIsGoodFormula) Score(v float64) (result float64, ok bool) {

	// for minimum is good formula, anything less than target max is good
	if v <= f.mingood.formula.max {
		return 100, true
	}

	// higher than best average
	if v > f.mingood.formula.avg {
		rr := (v - f.mingood.formula.avg) / f.mingood.formula.avg * 100
		result, ok = f.mingood.formula.ranges.Search(rr)
		if !ok {
			return 0.0, true // zero score for out of generated ranges
		}
	} else {
		return 0.0, false // should not reach here
	}
	return
}

// Name implements interface Scorer.Name()
func (f *MinIsGoodFormula) Name() string {
	return f.mingood.name
}

// ToString implements interface Stringer.ToString()
func (f *MinIsGoodFormula) ToString() string {
	return fmt.Sprintf("Formula:%s %s", f.mingood.name, stringify(f.mingood.formula))
}
