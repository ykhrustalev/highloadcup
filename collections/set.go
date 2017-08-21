package collections

type IntSet struct {
	set map[int]bool
}

var EmptyIntSet = IntSet{}

func NewIntSet() *IntSet {
	return &IntSet{
		set: make(map[int]bool),
	}
}

func (s *IntSet) Add(i int) bool {
	found := s.Contains(i)
	s.set[i] = true
	return !found
}

func (s *IntSet) Contains(i int) bool {
	_, found := s.set[i]
	return found
}

func (s *IntSet) Values() []int {
	res := make([]int, len(s.set))
	for k := range s.set {
		res = append(res, k)
	}
	return res
}

func (s *IntSet) Copy() *IntSet {
	c := NewIntSet()

	for k := range s.set {
		c.Add(k)
	}

	return c
}
