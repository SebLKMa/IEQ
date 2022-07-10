package formulas

import (
	"testing"
)

// We only need one instance of formula each
var temperatureFormula *StandardFormula
var co2Formula *MinIsGoodFormula
var lightingFormula *LightingFormula

func TestStandardFormula(t *testing.T) {
	// via pointers to interfaces
	temperatureFormula = &StandardFormula{}
	temperatureFormula.Setup("Temperature", 23, 25)
	str := temperatureFormula.ToString()
	if len(str) == 0 {
		t.Error("Standard Formula setup failure")
	}
	t.Log(str)

	// expecting 100 percent score
	value := float64(24)
	score, ok := temperatureFormula.Score(value)
	if !ok {
		t.Error("Standard Formula setup scoring error")
	}
	if score != 100 {
		t.Errorf("Standard Formula wrong score of %g for value %g", score, value)
	}

	// expecting 50 percent score
	value = float64(29)
	score, ok = temperatureFormula.Score(value)
	if !ok {
		t.Error("Standard Formula setup scoring error")
	}
	if score != 50 {
		t.Errorf("Standard Formula wrong score of %g for value %g", score, value)
	}

}

func TestMinIsGoodFormula(t *testing.T) {
	co2Formula = &MinIsGoodFormula{}
	co2Formula.Setup("CO2", 500, 700)
	str := co2Formula.ToString()
	if len(str) == 0 {
		t.Error("Min Is Good Formula setup failure")
	}
	t.Log(str)

	// expecting 100 percent score
	value := float64(700)
	score, ok := co2Formula.Score(value)
	if !ok {
		t.Error("Min Is Good Formula setup scoring error")
	}
	if score != 100 {
		t.Errorf("Min Is Good Formula wrong score of %g for value %g", score, value)
	}
	value = float64(500)
	score, ok = co2Formula.Score(value)
	if !ok {
		t.Error("Min Is Good Formula setup scoring error")
	}
	if score != 100 {
		t.Errorf("Min Is Good Formula wrong score of %g for value %g", score, value)
	}
	value = float64(0)
	score, ok = co2Formula.Score(value)
	if !ok {
		t.Error("Min Is Good Formula setup scoring error")
	}
	if score != 100 {
		t.Errorf("Min Is Good Formula wrong score of %g for value %g", score, value)
	}

	// expecting 90 percent score
	value = float64(800)
	score, ok = co2Formula.Score(value)
	if !ok {
		t.Error("Min Is Good Formula setup scoring error")
	}
	if score != 90 {
		t.Errorf("Min Is Good Formula wrong score of %g for value %g", score, value)
	}

	// expecting 50 percent score
	value = float64(1100)
	score, ok = co2Formula.Score(value)
	if !ok {
		t.Error("Min Is Good Formula setup scoring error")
	}
	if score != 50 {
		t.Errorf("Min Is Good Formula wrong score of %g for value %g", score, value)
	}

	// expecting 0 percent score
	value = float64(2000)
	score, ok = co2Formula.Score(value)
	if !ok {
		t.Error("Min Is Good Formula setup scoring error")
	}
	if score != 0 {
		t.Errorf("Min Is Good Formula wrong score of %g for value %g", score, value)
	}
}

func TestLightingFormula(t *testing.T) {
	// via pointers to interfaces
	lightingFormula = &LightingFormula{}
	lightingFormula.SetScale(1.5) // Scale must be set first
	lightingFormula.Setup("lighting", 250, 750)
	str := lightingFormula.ToString()
	if len(str) == 0 {
		t.Error("Standard Formula setup failure")
	}
	t.Log(str)

	// expecting 100 percent score
	value := float64(450)
	score, ok := lightingFormula.Score(value)
	if !ok {
		t.Error("Lighting Formula setup scoring error")
	}
	if score != 100 {
		t.Errorf("Lighting Formula wrong score of %g for value %g", score, value)
	}

	// expecting 12.5 percent score
	value = float64(20.8)
	score, ok = lightingFormula.Score(value)
	if !ok {
		t.Error("Lighting Formula setup scoring error")
	}
	if score != 12.5 {
		t.Errorf("Lighting Formula wrong score of %g for value %g", score, value)
	}
}
