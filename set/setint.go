package set

type SetInt struct {
	nodeMap map[int]bool
}

func NewSetInt(nums []int) *SetInt {
	s := &SetInt{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

func (s *SetInt) Add(num int) *SetInt {
	if s.nodeMap == nil {
		s.nodeMap = make(map[int]bool)
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = true
	}
	return s
}

func (s *SetInt) Clear() {
	s.nodeMap = make(map[int]bool)
}

func (s *SetInt) Remove(num int) bool {
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

func (s *SetInt) Contains(num int) bool {
	_, ok := s.nodeMap[num]
	return ok
}

func (s *SetInt) getList() []int {
	nums := []int{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

func (s *SetInt) Size() int {
	return len(s.nodeMap)
}

func (s *SetInt) Union(s2 *SetInt) *SetInt {
	s3 := SetInt{}
	s3.nodeMap = make(map[int]bool)
	for i := range s.nodeMap {
		s3.nodeMap[i] = true
	}
	for i := range s2.nodeMap {
		_, ok := s3.nodeMap[i]
		if !ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

// common in A and B
func (s *SetInt) Intersection(s2 *SetInt) *SetInt {
	s3 := SetInt{}
	s3.nodeMap = make(map[int]bool)
	for i := range s2.nodeMap {
		_, ok := s.nodeMap[i]
		if ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

// a.Minus(b) : all of a but not in b
func (s *SetInt) Minus(s2 *SetInt) *SetInt {
	s3 := SetInt{}
	s3.nodeMap = make(map[int]bool)
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

func (s *SetInt) Subset(s2 *SetInt) bool {
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			return false
		}
	}
	return true
}

func (s *SetInt) Superset(s2 *SetInt) bool {
	for i := range s2.nodeMap {
		_, ok := s.nodeMap[i]
		if !ok {
			return false
		}
	}
	size1 := s.Size()
	size2 := s2.Size()
	return size1 == size2 || size1 > size2
}
