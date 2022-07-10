package formulas

import (
	"fmt"

	st "github.com/seblkma/ieq/utils/skiptree"
)

// Formula defines the arguments to setup a Scorer object
type Formula struct {
	min    float64
	max    float64
	avg    float64
	rrange float64
	ranges st.ItemSkipTree
}

// Returns the contents of a RangeNode
// Using pointer to avoid copying in case of large skiptree
func stringify(f *Formula) string {
	f.ranges.String()
	return fmt.Sprintf("Min=%g Max=%g Avg=%g Rrange=%g", f.min, f.max, f.avg, f.rrange)
}
