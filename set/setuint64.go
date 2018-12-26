package set

type SetUint64 struct {
	nodeMap map[uint64]bool
}

func NewSetUint64(nums []uint64) *SetUint64 {
	s := &SetUint64{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

func (s *SetUint64) Add(num uint64) *SetUint64 {
	if s.nodeMap == nil {
		s.nodeMap = make(map[uint64]bool)
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = true
	}
	return s
}

func (s *SetUint64) Clear() {
	s.nodeMap = make(map[uint64]bool)
}

func (s *SetUint64) Remove(num uint64) bool {
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

func (s *SetUint64) Contains(num uint64) bool {
	_, ok := s.nodeMap[num]
	return ok
}

func (s *SetUint64) GetList() []uint64 {
	nums := []uint64{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

func (s *SetUint64) Size() int {
	return len(s.nodeMap)
}

func (s *SetUint64) Union(s2 *SetUint64) *SetUint64 {
	s3 := SetUint64{}
	s3.nodeMap = make(map[uint64]bool)
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
func (s *SetUint64) Intersection(s2 *SetUint64) *SetUint64 {
	s3 := SetUint64{}
	s3.nodeMap = make(map[uint64]bool)
	for i := range s2.nodeMap {
		_, ok := s.nodeMap[i]
		if ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

// a.Minus(b) : all of a but not in b
func (s *SetUint64) Minus(s2 *SetUint64) *SetUint64 {
	s3 := SetUint64{}
	s3.nodeMap = make(map[uint64]bool)
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

func (s *SetUint64) Subset(s2 *SetUint64) bool {
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			return false
		}
	}
	return true
}

func (s *SetUint64) Superset(s2 *SetUint64) bool {
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
