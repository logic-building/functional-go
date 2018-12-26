package set

type SetInt64 struct {
	nodeMap map[int64]bool
}

func NewSetInt64(nums []int64) *SetInt64 {
	s := &SetInt64{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

func (s *SetInt64) Add(num int64) *SetInt64 {
	if s.nodeMap == nil {
		s.nodeMap = make(map[int64]bool)
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = true
	}
	return s
}

func (s *SetInt64) Clear() {
	s.nodeMap = make(map[int64]bool)
}

func (s *SetInt64) Remove(num int64) bool {
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

func (s *SetInt64) Contains(num int64) bool {
	_, ok := s.nodeMap[num]
	return ok
}

func (s *SetInt64) GetList() []int64 {
	nums := []int64{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

func (s *SetInt64) Size() int {
	return len(s.nodeMap)
}

func (s *SetInt64) Union(s2 *SetInt64) *SetInt64 {
	s3 := SetInt64{}
	s3.nodeMap = make(map[int64]bool)
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
func (s *SetInt64) Intersection(s2 *SetInt64) *SetInt64 {
	s3 := SetInt64{}
	s3.nodeMap = make(map[int64]bool)
	for i := range s2.nodeMap {
		_, ok := s.nodeMap[i]
		if ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

// a.Minus(b) : all of a but not in b
func (s *SetInt64) Minus(s2 *SetInt64) *SetInt64 {
	s3 := SetInt64{}
	s3.nodeMap = make(map[int64]bool)
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

func (s *SetInt64) Subset(s2 *SetInt64) bool {
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			return false
		}
	}
	return true
}

func (s *SetInt64) Superset(s2 *SetInt64) bool {
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
