package ratings

import (
	"errors"
)

// Index represents an index and its score
type Index struct {
	name  string
	score float64
}

// Rating represents rateable quality
type Rating struct {
	name      string
	weighting float64
	score     float64
	indices   []Index
}

// Setup implements interface Rateable.Setup
// Initializes the Rating with name and weighing.
func (r *Rating) Setup(n string, w float64) {
	r.name = n
	r.weighting = w
}

// AddIndex implements interface Rateable.AddIndex
// Adds an index and its score value
// Returns error if value or existing indices values already adds up to or more than 100 percent
func (r *Rating) AddIndex(n string, v float64) error {
	if v > 100 {
		//log.Fatal("Value or existing indices values already adds up to or more than 100 percent")
		return errors.New("Value must not be more than 100 percent")
	}

	sum := float64(0.0)
	for _, i := range r.indices {
		sum += i.score
	}
	if sum > 100 {
		//log.Fatal("Value or existing indices values already adds up to or more than 100 percent")
		return errors.New("Existing indices values already adds up to or more than 100 percent")
	}

	if sum+v > 100 {
		//log.Fatal("Value or existing indices values already adds up to or more than 100 percent")
		return errors.New("Value plus existing indices values already adds up to or more than 100 percent")
	}

	r.indices = append(r.indices, Index{name: n, score: v})
	return nil
}

// SetRating implements interface Rateable.SetRating
// Computes the Rating's score using its weighting.
func (r *Rating) SetRating() {
	len := float64(len(r.indices))
	//fmt.Printf("index=%s len=%g\n", r.name, len)
	if len == 0.0 {
		return
	}

	sum := float64(0.0)
	for _, i := range r.indices {
		sum += i.score
	}
	//fmt.Printf("index=%s sum=%g\n", r.name, sum)
	if sum == 0.0 {
		return
	}

	// e.g., score = ((r.Temperature + r.Humidity) / 2) * r.Rating.weighting / 100
	r.score = (sum / len) * r.weighting / 100
}

// Name implements Rateable.Name
// Returns the Rating's Name.
func (r *Rating) Name() string {
	return r.name
}

// Weighting implements Rateable.Weighting
// Returns the Rating's Weighting.
func (r *Rating) Weighting() float64 {
	return r.weighting
}

// Rate implements Rateable.Rate
// Returns the Rating's computed score.
func (r *Rating) Rate() float64 {
	return r.score
}
