package set

type SetUint16 struct {
	nodeMap map[uint16]bool
}

func NewSetUint16(nums []uint16) *SetUint16 {
	s := &SetUint16{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

func (s *SetUint16) Add(num uint16) *SetUint16 {
	if s.nodeMap == nil {
		s.nodeMap = make(map[uint16]bool)
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = true
	}
	return s
}

func (s *SetUint16) Clear() {
	s.nodeMap = make(map[uint16]bool)
}

func (s *SetUint16) Remove(num uint16) bool {
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

func (s *SetUint16) Contains(num uint16) bool {
	_, ok := s.nodeMap[num]
	return ok
}

func (s *SetUint16) GetList() []uint16 {
	nums := []uint16{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

func (s *SetUint16) Size() int {
	return len(s.nodeMap)
}

func (s *SetUint16) Union(s2 *SetUint16) *SetUint16 {
	s3 := SetUint16{}
	s3.nodeMap = make(map[uint16]bool)
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
func (s *SetUint16) Intersection(s2 *SetUint16) *SetUint16 {
	s3 := SetUint16{}
	s3.nodeMap = make(map[uint16]bool)
	for i := range s2.nodeMap {
		_, ok := s.nodeMap[i]
		if ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

// a.Minus(b) : all of a but not in b
func (s *SetUint16) Minus(s2 *SetUint16) *SetUint16 {
	s3 := SetUint16{}
	s3.nodeMap = make(map[uint16]bool)
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

func (s *SetUint16) Subset(s2 *SetUint16) bool {
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			return false
		}
	}
	return true
}

func (s *SetUint16) Superset(s2 *SetUint16) bool {
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
