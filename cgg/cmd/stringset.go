package cmd

// StringSet provides a quick implementation of a Set for strings.
import "sync"

// This may or may not be concurrent safe depending on how it is created.
type StringSet struct {
	contents map[string]int8
	lk       *sync.RWMutex
}

// NewStringSet returns a new string set intended for use in a single
// goroutine.
func NewStringSet() *StringSet {
	return &StringSet{
		contents: make(map[string]int8),
	}
}

// NewSafeStringSet returns a new string set that is guarded against concurrent
// use.
func NewSafeStringSet() *StringSet {
	return &StringSet{
		contents: make(map[string]int8),
		lk:       &sync.RWMutex{},
	}
}

// IsSafe returns whether or not the StringSet is safe for concurrent use.
func (s *StringSet) IsSafe() bool {
	return s.lk != nil
}

// AddContains checks if the value is in the StringSet and then adds it if it
// is not. If it was in the set beforehand then false is returned.
func (s *StringSet) AddContains(value string) bool {
	if s.IsSafe() {
		s.lk.Lock()
		defer s.lk.Unlock()
	}
	if _, exists := s.contents[value]; exists {
		return false
	}
	s.contents[value] = 0
	return true
}

// Add adds the given value to the string set.
func (s *StringSet) Add(value string) {
	if s.IsSafe() {
		s.lk.Lock()
		defer s.lk.Unlock()
	}
	s.contents[value] = 0
}

// Remove removes the given value from the string set if it was present.
func (s *StringSet) Remove(value string) {
	if s.IsSafe() {
		s.lk.Lock()
		defer s.lk.Unlock()
	}
	delete(s.contents, value)
}

// Contains returns true if the given value exists in the string set.
func (s *StringSet) Contains(value string) bool {
	s.lk.RLock()
	defer s.lk.RUnlock()
	_, exists := s.contents[value]
	return exists
}

// ToSlice returns a copy of the StringSet in the form of a slice of strings.
func (s *StringSet) ToSlice() []string {
	if s.IsSafe() {
		s.lk.RLock()
		defer s.lk.RUnlock()
	}
	cSlice := make([]string, 0, len(s.contents))
	for k := range s.contents {
		cSlice = append(cSlice, k)
	}
	return cSlice
}
