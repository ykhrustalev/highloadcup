package highloadcup

type IntSet struct {
	set map[int]bool
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
