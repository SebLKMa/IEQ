package formulas

import (
	"fmt"
)

// LightingFormula defines a type to contain the lighting formula and implement Score interface
type LightingFormula struct {
	name    string
	formula *Formula
	scale   float64
}

// SetScale sets the scale for Setup to generate ranges.
// SetScale has to be called before Setup.
// If not called, the default scale is set to 1.
func (f *LightingFormula) SetScale(sc float64) {
	f.scale = sc
}

// Setup implements interface Scorer.Setup()
// Setup will populate the scores ranges used by Score().
// Please note that SetScale has to be called before Setup.
// If scale is not set, the default scale is set to 1.
func (f *LightingFormula) Setup(n string, min float64, max float64) {
	// TODO - make this formula unexported, export a constructor with required input for scale.
	if f.scale <= 0 {
		f.scale = 1
	}
	f.formula = &Formula{}
	f.name = n
	f.formula.min = min
	f.formula.max = max
	f.formula.avg = (f.formula.min + f.formula.max) / 2
	f.formula.rrange = (f.formula.max - f.formula.avg) / f.formula.avg * 100

	// generate the relative ranges as a binary skiptree
	//scale := 1.5
	incrementalfrom := f.formula.avg // 500.0
	incrementalto := incrementalfrom * f.scale
	decrementalfrom := f.formula.avg // 500.0
	decrementalto := decrementalfrom / f.scale
	high := 100.0
	chunks := 10
	diff := 12.5 // interval
	score := high
	//fmt.Printf("From:%g To:%g Score:%g\n", incrementalfrom, incrementalfrom, float64(score))
	f.formula.ranges.Insert(incrementalfrom, incrementalto, float64(score))
	for i := 1; i < chunks; i++ {

		//fmt.Printf("From:%g To:%g Score:%g From:%g To:%g\n", incrementalfrom, incrementalto, float64(score), decrementalfrom, decrementalto)
		f.formula.ranges.Insert(incrementalfrom, incrementalto, float64(score)) // from - to
		f.formula.ranges.Insert(decrementalto, decrementalfrom, float64(score)) // to - from
		score = score - diff
		incrementalfrom = incrementalto
		incrementalto *= f.scale
		decrementalfrom = decrementalto
		decrementalto /= f.scale
	}
}

// Score implements interface Scorer.Score()
func (f *LightingFormula) Score(v float64) (result float64, ok bool) {

	result, ok = f.formula.ranges.Search(v)
	if !ok {
		//fmt.Printf("Lighting result:%g\n", result)
		// score is 0 when lux value is out of ranges
		return result, true
	}

	return
}

// Name implements interface Scorer.Name()
func (f *LightingFormula) Name() string {
	return f.name
}

// ToString implements interface Stringer.ToString()
func (f *LightingFormula) ToString() string {
	return fmt.Sprintf("Formula:%s %s Scale:%g", f.name, stringify(f.formula), f.scale)
}
