// Package pos_error error handling
package pos_error

import "errors"

// TokenError custom error
type TokenError struct {
	error
	position int
}

// GetPos getter for position in TokenError
func (err TokenError) GetPos() int {
	return err.position
}

// PosError custom error with position interface
type PosError interface {
	error
	GetPos() int
}

// New creates a new TokenError based on parameters
// message is converted into an error and stored
// typ defines the ErrType
// Position is the Position in the originating text where the error occurred
func New(message string, position int) PosError {
	err := &TokenError{
		error:    errors.New(message),
		position: position,
	}
	return err
}
