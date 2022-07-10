package ratings

import (
	"testing"
)

func TestStandardRating(t *testing.T) {
	// components weightings set to all 25
	thermalRating := Rating{}
	thermalRating.Setup("Thermal", 25)

	iaqRating := Rating{}
	iaqRating.Setup("IAQ", 25)

	lightingRating := Rating{}
	lightingRating.Setup("Lighting", 25)

	noiseRating := Rating{}
	noiseRating.Setup("Noise", 25)

	// all scores 100 percent
	thermalRating.AddIndex("Temperature", 50)
	thermalRating.AddIndex("Humidity", 50)
	iaqRating.AddIndex("CO2", 30)
	iaqRating.AddIndex("VOC", 30)
	iaqRating.AddIndex("PM25", 40)
	lightingRating.AddIndex("Lighting", 100)
	noiseRating.AddIndex("Noise", 100)

	thermalRating.SetRating()
	iaqRating.SetRating()
	lightingRating.SetRating()
	noiseRating.SetRating()

	// expecting overall score also 100 percent
	ieqRating := IEQRating{}
	ieqRating.Setup("Overall IEQ", 1.0)
	ieqRating.AddIndex(thermalRating.Name(), thermalRating.Rate())
	ieqRating.AddIndex(iaqRating.Name(), iaqRating.Rate())
	ieqRating.AddIndex(lightingRating.Name(), lightingRating.Rate())
	ieqRating.AddIndex(noiseRating.Name(), noiseRating.Rate())
	ieqRating.SetRating()

	total := float64(0)
	for _, i := range ieqRating.Indices() {
		t.Logf("%v \n", i)
		total += i.score
	}
	if total != ieqRating.Rate() {
		t.Errorf("Unexpected Overall Rating %g\n", ieqRating.Rate())
	}
	t.Logf("%s Rating: %g\n", ieqRating.Name(), ieqRating.Rate())
}

func TestRatingsSetupError(t *testing.T) {
	// Test existing list of indices already adds up to 100 percent

	// components weightings set to all 25
	thermalRating := Rating{}
	thermalRating.Setup("Thermal", 25)

	iaqRating := Rating{}
	iaqRating.Setup("IAQ", 25)

	lightingRating := Rating{}
	lightingRating.Setup("Lighting", 25)

	noiseRating := Rating{}
	noiseRating.Setup("Noise", 25)

	// all scores 100 percent
	err := thermalRating.AddIndex("Temperature", 50)
	if err != nil {
		t.Error(err.Error())
	}
	err = thermalRating.AddIndex("Humidity", 50)
	if err != nil {
		t.Error(err.Error())
	}
	err = iaqRating.AddIndex("CO2", 30)
	if err != nil {
		t.Error(err.Error())
	}
	err = iaqRating.AddIndex("VOC", 30)
	if err != nil {
		t.Error(err.Error())
	}
	err = iaqRating.AddIndex("PM25", 50)
	if err == nil {
		t.Error("Expecting AddIndex error 30 + 30 + 50 exceeds 100")
	}

	t.Logf("OK we expect error: %s \n", err.Error())
}
