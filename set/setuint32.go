package set

type SetUint32 struct {
	nodeMap map[uint32]bool
}

func NewUint32(nums []uint32) *SetUint32 {
	s := &SetUint32{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

func (s *SetUint32) Add(num uint32) *SetUint32 {
	if s.nodeMap == nil {
		s.nodeMap = make(map[uint32]bool)
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = true
	}
	return s
}

func (s *SetUint32) Clear() {
	s.nodeMap = make(map[uint32]bool)
}

func (s *SetUint32) Remove(num uint32) bool {
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

func (s *SetUint32) Contains(num uint32) bool {
	_, ok := s.nodeMap[num]
	return ok
}

func (s *SetUint32) GetList() []uint32 {
	nums := []uint32{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

func (s *SetUint32) Size() int {
	return len(s.nodeMap)
}

func (s *SetUint32) Union(s2 *SetUint32) *SetUint32 {
	s3 := SetUint32{}
	s3.nodeMap = make(map[uint32]bool)
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
func (s *SetUint32) Intersection(s2 *SetUint32) *SetUint32 {
	s3 := SetUint32{}
	s3.nodeMap = make(map[uint32]bool)
	for i := range s2.nodeMap {
		_, ok := s.nodeMap[i]
		if ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

// a.Minus(b) : all of a but not in b
func (s *SetUint32) Minus(s2 *SetUint32) *SetUint32 {
	s3 := SetUint32{}
	s3.nodeMap = make(map[uint32]bool)
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

func (s *SetUint32) Subset(s2 *SetUint32) bool {
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			return false
		}
	}
	return true
}

func (s *SetUint32) Superset(s2 *SetUint32) bool {
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
