package utils

import "fmt"

// Wrap adds context to the passed error. If error is nil Wrap also returns nil.
// Step allows you to simplify error checking in some cases, for example:
//
//	if err != nil {
//	    return fmt.Errorf("some context: %w", err)
//	}
//
// Can be rewritten as
//
//	return Wrap(err, "some context")
//
// Step will have the same behavior if it exists at the end of the function. It will return nil
// if there is no error, or a wrapped error if there is.
func Wrap(err error, format string, a ...any) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", fmt.Sprintf(format, a...), err)
}

func Report(err error, format string, a ...any) {
	if err == nil {
		return
	}
	fmt.Printf("%s: %s\n", fmt.Sprintf(format, a...), err)
}
