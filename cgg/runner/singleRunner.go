package runner

import "sync"

type singleRunner struct {
	f      Func
	result interface{}
	e      error
	hasRun bool
	lk     *sync.RWMutex
	wg     *sync.WaitGroup
	*baseRunner
}

// NewSingle returns a Runner that is optimized for running on only a single
// goroutine. This is eqivalent to using a Multi Runner with one goroutine
// in all other ways and all functions are still safe to call from multiple
// goroutines.
func NewSingle(f Func) Runner {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	sR := &singleRunner{
		lk: &sync.RWMutex{},
		wg: wg,
		f:  f,
	}
	sR.baseRunner = &baseRunner{sR, 1}
	return sR
}

func (s *singleRunner) Run() {
	// lock so we can check hasRun
	s.lk.Lock()
	if s.hasRun {
		s.lk.Unlock()
		return
	}

	// defer the "done" call on the waitgroup
	defer s.wg.Done()

	s.hasRun = true
	// unlock now that we have set hasRun but done == false
	s.lk.Unlock()

	// compute result while unlocked
	s.result, s.e = s.f(0, s)
}

func (s *singleRunner) HasRun() bool {
	s.lk.RLock()
	defer s.lk.RUnlock()
	return s.hasRun
}

func (s *singleRunner) Wait() ([]interface{}, []error) {
	s.wg.Wait()
	s.lk.RLock()
	defer s.lk.RUnlock()
	return []interface{}{s.result}, []error{s.e}
}
