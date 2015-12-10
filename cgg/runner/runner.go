// Package runner defines an interface and implementation of a Runner which
// provides a generic way to split the computation of a function amoung
// an arbitrary number of goroutines. In addition, it provides all the tools
// needed in order to synchronize access to the results and wait on the
// computation's completion.
package runner

// Func defines a standard function type for a function that is
// parallelized by a Runner.
type Func func(i int, r Runner) (interface{}, error)

// Runner defines an interface for a general case method to parallelize (or
// make concurrent depending on parameters/resources available) a arbitrary
// function that matches an expected interface. All methods should be safe to
// call concurrently in multiple goroutines.
//
// In many cases one could obtain slightly more efficiency by doing this on a
// case-by-case basis but much of the work would be repeated.
type Runner interface {
	// HasRun provides a concurrent-safe view into whether the Run() function
	// has been called on this Runner. This will return true after Run() is
	// called and potentially before the Run method's execution has completed.
	HasRun() bool

	// Run will start the execution of the supplied function split amoung the
	// desired number of goroutines. This is safe to call more than once as
	// all subsequent calls will simply return.
	//
	// Typically a call to Run() is not blocking although this is left up to
	// the implementation.
	Run()

	// NumRoutines returns the configured number of routines that the Runner
	// should use upon a call to Run().
	NumRoutines() int

	// Results waits for a call to Run to finish and then returns a copy of
	// the slice of results. This slice will have length equal to the return of
	// NumRoutines().
	//
	// Note: The return value is guarenteed to be a copy of the slice
	// (different backing array) but it is NOT guarneteed that the values
	// in the array themselves will be copies (especially if they themselves
	// are slices).
	Results() []interface{}

	// Errors waits for a call to Run to finish and then returns a copy of the
	// slice of errors (including nil errors).
	Errors() []error

	// Wait blocks until a call to Run() and all related computations have
	// completed. It will then return a copy of the results and errors (see
	// comments on Results() and Errors()).
	Wait() ([]interface{}, []error)

	// HadError waits for a call to Run to finish and then returns an arbitrary
	// error if one occurred. If all entries in the errors are nil then this
	// returns nil itself.
	HadError() error
}

// New returns a new Runner with the desired number of goroutines. The
// returned Runner implementation may vary depending on this number.
func New(f Func, numRoutines int) Runner {
	if numRoutines == 1 {
		return NewSingle(f)
	}
	return NewMulti(f, numRoutines)
}
