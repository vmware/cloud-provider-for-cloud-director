package testingsdk

var exists = struct{}{}

type set struct {
	m map[string]struct{}
}

func NewSet() *set {
	s := &set{}
	s.m = make(map[string]struct{})
	return s
}

func (s *set) Add(value string) {
	s.m[value] = exists
}

func (s *set) Contains(value string) bool {
	_, c := s.m[value]
	return c
}

func (s *set) Size() int {
	return len(s.m)
}
