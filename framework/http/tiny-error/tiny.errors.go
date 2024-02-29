package tinyerror

import "errors"

// Create a custom error message
func New(err string) error {
	return errors.New(err)

}
