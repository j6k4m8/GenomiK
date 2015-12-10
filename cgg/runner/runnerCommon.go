package runner

// baseRunner provides a really simple implementation of some of the easy
// expected methods in the Runner interface. This allows a Runner
// implementation to embed a *baseRunner and implement several functions
// without any additional work. This all relies on a Runner implementation
// having a correct implementation of the Wait() function.
type baseRunner struct {
	r   Runner
	num int
}

func (b *baseRunner) NumRoutines() int {
	return b.num
}

func (b *baseRunner) Errors() []error {
	_, errors := b.r.Wait()
	return errors
}

func (b *baseRunner) Results() []interface{} {
	results, _ := b.r.Wait()
	return results
}

func (b *baseRunner) HadError() error {
	for _, e := range b.Errors() {
		if e != nil {
			return e
		}
	}
	return nil
}
