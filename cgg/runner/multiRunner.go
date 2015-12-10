package runner

import (
	"runtime"
	"sync"
)

type multiRunner struct {
	f       Func
	num     int
	results []interface{}
	errors  []error
	hasRun  bool
	lk      *sync.RWMutex
	wg      *sync.WaitGroup
	*baseRunner
}

// NewMax returns a new Runner configured to use as many goroutines as the
// runtime is configured to run at once (obtained from runtime.GOMAXPROCS).
func NewMax(f Func) Runner {
	return NewMulti(f, -1)
}

// NewMulti returns a new runner with a configurable number of goroutines to be
// launched. If the number of routines is less than or equal to 0, then the
// return value from runtime.GOMAXPROCS(-1) is used instead.
func NewMulti(f Func, numRoutines int) Runner {
	if numRoutines <= 0 {
		numRoutines = runtime.GOMAXPROCS(-1)
	}
	wg := &sync.WaitGroup{}
	wg.Add(numRoutines)
	mR := &multiRunner{
		f:       f,
		num:     numRoutines,
		lk:      &sync.RWMutex{},
		wg:      wg,
		results: make([]interface{}, numRoutines),
		errors:  make([]error, numRoutines),
	}
	mR.baseRunner = &baseRunner{mR, numRoutines}
	return mR
}

func (m *multiRunner) Run() {
	// lock so we can check hasRun
	m.lk.Lock()
	if m.hasRun {
		m.lk.Unlock()
		return
	}

	// set hasRun while still locked
	m.hasRun = true
	// release lock so we can run the goroutines
	m.lk.Unlock()

	// launch all of the routines!
	for i := 0; i < m.num; i++ {
		// pass i as a parameter so we don't have to deal with it changing
		go func(j int) {
			defer m.wg.Done()
			m.results[j], m.errors[j] = m.f(j, m)
		}(i)
	}

	// return control to the caller
}

func (m *multiRunner) HasRun() bool {
	m.lk.RLock()
	defer m.lk.RUnlock()
	return m.hasRun
}

func (m *multiRunner) Wait() ([]interface{}, []error) {
	m.wg.Wait()
	m.lk.RLock()
	defer m.lk.RUnlock()
	// return m.results, m.errors
	resultsCopy := make([]interface{}, len(m.results))
	copy(resultsCopy, m.results)
	errorsCopy := make([]error, len(m.errors))
	copy(errorsCopy, m.errors)
	return resultsCopy, errorsCopy
}
