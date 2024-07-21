package configs

// Config interface for every config
type Config interface {

	// Validate function makes necessary fields validation
	Validate() error
}
