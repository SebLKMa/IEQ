package ratings

import (
	"fmt"

	intf "github.com/seblkma/ieq/interfaces"
)

// Setup sets up the concrete scorer
func Setup(sc intf.Scorer, name string, min float64, max float64) {
	sc.Setup(name, min, max)
}

// ComputeScore uses the provided concrete Scorer and measured value to compute its soore
func ComputeScore(sc intf.Scorer, value float64) (score float64) {
	//fmt.Println(tempeFormula.String())
	score, ok := sc.Score(value)
	if ok {
		fmt.Printf("%s: %g   Score: %g\n", sc.Name(), value, score)
	} else {
		fmt.Println("Unable to compute score :(")
	}
	return
}

// PrintInfo is helper function to print description of a Stringer
func PrintInfo(str intf.Stringer) {
	fmt.Println(str.ToString())
}
