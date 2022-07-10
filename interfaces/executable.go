package interfaces

// Executable represents an executable task
type Executable interface {
	Execute() error
}
