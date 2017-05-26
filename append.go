package multierror

import "fmt"

// Append is a helper function that will append more errors
// onto an Error in order to create a larger multi-error.
//
// If err is not a multierror.Error, then it will be turned into
// one. If any of the errs are multierr.Error, they will be flattened
// one level into err.
func Append(err error, errs ...error) *Error {
	switch err := err.(type) {
	case *Error:
		// Typed nils can reach here, so initialize if we are nil
		if err == nil {
			err = new(Error)
		}

		// Go through each error and flatten
		for _, e := range errs {
			switch e := e.(type) {
			case *Error:
				if e != nil {
					err.Errors = append(err.Errors, e.Errors...)
				}
			default:
				if e != nil {
					err.Errors = append(err.Errors, e)
				}
			}
		}

		// if none of the errors are non-nil
		// return a nil value
		if len(err.Errors) == 0 {
			return nil
		}

		return err
	default:
		newErrs := make([]error, 0, len(errs)+1)
		for _, e := range errs {
			if e != nil {
				newErrs = append(newErrs, e)
			}
		}

		// if none of the errors are non-nil
		// return a nil value
		if len(newErrs) == 0 {
			err = nil
			fmt.Println("EHRE?")
			return nil
		}

		return Append(&Error{}, newErrs...)
	}
}
