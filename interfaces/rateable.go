package interfaces

// Rateable represents objects that can be rated based on a weighting
type Rateable interface {
	Setup(n string, w float64)
	AddIndex(n string, v float64) error
	SetRating()
	Name() string
	Weighting() float64
	Rate() float64
}
