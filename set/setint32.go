package set

type SetInt32 struct {
	nodeMap map[int32]bool
}

func NewSetInt32(nums []int32) *SetInt32 {
	s := &SetInt32{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

func (s *SetInt32) Add(num int32) *SetInt32 {
	if s.nodeMap == nil {
		s.nodeMap = make(map[int32]bool)
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = true
	}
	return s
}

func (s *SetInt32) Clear() {
	s.nodeMap = make(map[int32]bool)
}

func (s *SetInt32) Remove(num int32) bool {
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

func (s *SetInt32) Contains(num int32) bool {
	_, ok := s.nodeMap[num]
	return ok
}

func (s *SetInt32) GetList() []int32 {
	nums := []int32{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

func (s *SetInt32) Size() int {
	return len(s.nodeMap)
}

func (s *SetInt32) Union(s2 *SetInt32) *SetInt32 {
	s3 := SetInt32{}
	s3.nodeMap = make(map[int32]bool)
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
func (s *SetInt32) Intersection(s2 *SetInt32) *SetInt32 {
	s3 := SetInt32{}
	s3.nodeMap = make(map[int32]bool)
	for i := range s2.nodeMap {
		_, ok := s.nodeMap[i]
		if ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

// a.Minus(b) : all of a but not in b
func (s *SetInt32) Minus(s2 *SetInt32) *SetInt32 {
	s3 := SetInt32{}
	s3.nodeMap = make(map[int32]bool)
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

func (s *SetInt32) Subset(s2 *SetInt32) bool {
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			return false
		}
	}
	return true
}

func (s *SetInt32) Superset(s2 *SetInt32) bool {
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
