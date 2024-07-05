package adt

type Set struct {
	size int
} 

func NewSet() *Set {
	return &Set{0}
}

func (s *Set) IsEmpty() bool  {
	return s.size == 0
}

func (s *Set) Add(v string) {
	s.size++
}

func (s *Set) Size() int {
	return s.size
}

func (s *Set) Contains(v string) bool {
	if s.IsEmpty() {
		return false
	}

	return true
}