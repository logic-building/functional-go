package set

type SetInt8 struct {
	nodeMap map[int8]bool
}

func NewSetInt8(nums []int8) *SetInt8 {
	s := &SetInt8{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

func (s *SetInt8) Add(num int8) *SetInt8 {
	if s.nodeMap == nil {
		s.nodeMap = make(map[int8]bool)
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = true
	}
	return s
}

func (s *SetInt8) Clear() {
	s.nodeMap = make(map[int8]bool)
}

func (s *SetInt8) Remove(num int8) bool {
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

func (s *SetInt8) Contains(num int8) bool {
	_, ok := s.nodeMap[num]
	return ok
}

func (s *SetInt8) GetList() []int8 {
	nums := []int8{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

func (s *SetInt8) Size() int {
	return len(s.nodeMap)
}

func (s *SetInt8) Union(s2 *SetInt8) *SetInt8 {
	s3 := SetInt8{}
	s3.nodeMap = make(map[int8]bool)
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
func (s *SetInt8) Intersection(s2 *SetInt8) *SetInt8 {
	s3 := SetInt8{}
	s3.nodeMap = make(map[int8]bool)
	for i := range s2.nodeMap {
		_, ok := s.nodeMap[i]
		if ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

// a.Minus(b) : all of a but not in b
func (s *SetInt8) Minus(s2 *SetInt8) *SetInt8 {
	s3 := SetInt8{}
	s3.nodeMap = make(map[int8]bool)
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

func (s *SetInt8) Subset(s2 *SetInt8) bool {
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			return false
		}
	}
	return true
}

func (s *SetInt8) Superset(s2 *SetInt8) bool {
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
