package set

type SetInt16 struct {
	nodeMap map[int16]bool
}

func NewSetInt16(nums []int16) *SetInt16 {
	s := &SetInt16{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

func (s *SetInt16) Add(num int16) *SetInt16 {
	if s.nodeMap == nil {
		s.nodeMap = make(map[int16]bool)
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = true
	}
	return s
}

func (s *SetInt16) Clear() {
	s.nodeMap = make(map[int16]bool)
}

func (s *SetInt16) Remove(num int16) bool {
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

func (s *SetInt16) Contains(num int16) bool {
	_, ok := s.nodeMap[num]
	return ok
}

func (s *SetInt16) GetList() []int16 {
	nums := []int16{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

func (s *SetInt16) Size() int {
	return len(s.nodeMap)
}

func (s *SetInt16) Union(s2 *SetInt16) *SetInt16 {
	s3 := SetInt16{}
	s3.nodeMap = make(map[int16]bool)
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
func (s *SetInt16) Intersection(s2 *SetInt16) *SetInt16 {
	s3 := SetInt16{}
	s3.nodeMap = make(map[int16]bool)
	for i := range s2.nodeMap {
		_, ok := s.nodeMap[i]
		if ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

// a.Minus(b) : all of a but not in b
func (s *SetInt16) Minus(s2 *SetInt16) *SetInt16 {
	s3 := SetInt16{}
	s3.nodeMap = make(map[int16]bool)
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

func (s *SetInt16) Subset(s2 *SetInt16) bool {
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			return false
		}
	}
	return true
}

func (s *SetInt16) Superset(s2 *SetInt16) bool {
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
