package set

type SetFloat32 struct {
	nodeMap map[float32]bool
}

func NewSetFloat32(nums []float32) *SetFloat32 {
	s := &SetFloat32{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

func (s *SetFloat32) Add(num float32) *SetFloat32 {
	if s.nodeMap == nil {
		s.nodeMap = make(map[float32]bool)
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = true
	}
	return s
}

func (s *SetFloat32) Clear() {
	s.nodeMap = make(map[float32]bool)
}

func (s *SetFloat32) Remove(num float32) bool {
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

func (s *SetFloat32) Contains(num float32) bool {
	_, ok := s.nodeMap[num]
	return ok
}

func (s *SetFloat32) GetList() []float32 {
	nums := []float32{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

func (s *SetFloat32) Size() int {
	return len(s.nodeMap)
}

func (s *SetFloat32) Union(s2 *SetFloat32) *SetFloat32 {
	s3 := SetFloat32{}
	s3.nodeMap = make(map[float32]bool)
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
func (s *SetFloat32) Intersection(s2 *SetFloat32) *SetFloat32 {
	s3 := SetFloat32{}
	s3.nodeMap = make(map[float32]bool)
	for i := range s2.nodeMap {
		_, ok := s.nodeMap[i]
		if ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

// a.Minus(b) : all of a but not in b
func (s *SetFloat32) Minus(s2 *SetFloat32) *SetFloat32 {
	s3 := SetFloat32{}
	s3.nodeMap = make(map[float32]bool)
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

func (s *SetFloat32) Subset(s2 *SetFloat32) bool {
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			return false
		}
	}
	return true
}

func (s *SetFloat32) Superset(s2 *SetFloat32) bool {
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
