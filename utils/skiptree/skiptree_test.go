package skiptree

import (
	"fmt"
	"math"
	"testing"
)

var st ItemSkipTree

func fillSkipTree(st *ItemSkipTree) {
	// manually balance the tree by insert sequence :)

	st.Insert(20.83333333, 25, 50)
	st.Insert(25, 29.16666667, 40)
	st.Insert(29.16666667, 33.33333333, 30)
	st.Insert(33.33333333, 37.5, 20)
	st.Insert(37.5, 41.66666667, 10)
	st.Insert(41.66666667, 45.83333333, 0)

	st.Insert(0, 4.166666667, 100)
	st.Insert(4.166666667, 8.333333333, 90)
	st.Insert(8.333333333, 12.5, 80)
	st.Insert(12.5, 16.66666667, 70)
	st.Insert(16.66666667, 20.83333333, 60)
}
func TestInsertSkipTree(t *testing.T) {
	fillSkipTree(&st)
	st.String()
}

func TestSearchSkipTree(t *testing.T) {
	value := 38.5
	score, found := st.Search(value)
	fmt.Printf("Value:%g Score:%g\n", value, score)
	if !found || score != 10 {
		t.Errorf("search not working")
	}

	value = 30.139999389648438
	score, found = st.Search(value)
	fmt.Printf("Value:%g Score:%g\n", value, score)
	if !found || score != 30 {
		t.Errorf("search not working")
	}
}

func TestSearchSkipTreeRounded(t *testing.T) {
	value := math.Round(0)
	score, found := st.Search(value)
	fmt.Printf("Value:%g Score:%g\n", value, score)
	if !found || score != 100 {
		t.Errorf("search not working")
	}

	value = math.Round(18.48)
	score, found = st.Search(value)
	fmt.Printf("Value:%g Score:%g\n", value, score)
	if !found || score != 60 {
		t.Errorf("search not working")
	}

	value = math.Round(38.5)
	score, found = st.Search(value)
	fmt.Printf("Value:%g Score:%g\n", value, score)
	if !found || score != 10 {
		t.Errorf("search not working")
	}

	value = math.Round(30.139999389648438*100) / 100 // to round to 2 decimal places
	score, found = st.Search(value)
	fmt.Printf("Value:%g Score:%g\n", value, score)
	if !found || score != 30 {
		t.Errorf("search not working")
	}
}

func TestRemoveSkipTree(t *testing.T) {
	st.Remove(37.5)

	// should not be able to find this in ranges
	score, found := st.Search(math.Round(38.5))
	if found || score == 10 {
		t.Errorf("Remove not working")
	}
	//st.String()
}
