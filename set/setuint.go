package set

type SetUint struct {
	nodeMap map[uint]bool
}

func NewSetUint(nums []uint) *SetUint {
	s := &SetUint{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

func (s *SetUint) Add(num uint) *SetUint {
	if s.nodeMap == nil {
		s.nodeMap = make(map[uint]bool)
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = true
	}
	return s
}

func (s *SetUint) Clear() {
	s.nodeMap = make(map[uint]bool)
}

func (s *SetUint) Remove(num uint) bool {
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

func (s *SetUint) Contains(num uint) bool {
	_, ok := s.nodeMap[num]
	return ok
}

func (s *SetUint) GetList() []uint {
	nums := []uint{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

func (s *SetUint) Size() int {
	return len(s.nodeMap)
}

func (s *SetUint) Union(s2 *SetUint) *SetUint {
	s3 := SetUint{}
	s3.nodeMap = make(map[uint]bool)
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
func (s *SetUint) Intersection(s2 *SetUint) *SetUint {
	s3 := SetUint{}
	s3.nodeMap = make(map[uint]bool)
	for i := range s2.nodeMap {
		_, ok := s.nodeMap[i]
		if ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

// a.Minus(b) : all of a but not in b
func (s *SetUint) Minus(s2 *SetUint) *SetUint {
	s3 := SetUint{}
	s3.nodeMap = make(map[uint]bool)
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

func (s *SetUint) Subset(s2 *SetUint) bool {
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			return false
		}
	}
	return true
}

func (s *SetUint) Superset(s2 *SetUint) bool {
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
