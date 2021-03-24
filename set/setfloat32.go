package set

// Float32 struct
type Float32 struct {
	nodeMap map[float32]struct{}
}

// NewFloat32 creates set struct
func NewFloat32(nums []float32) *Float32 {
	s := &Float32{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

// Add adds given item
func (s *Float32) Add(num float32) *Float32 {
	if s.nodeMap == nil {
		s.nodeMap = make(map[float32]struct{})
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = struct{}{}
	}
	return s
}

// Clear struct
func (s *Float32) Clear() {
	s.nodeMap = make(map[float32]struct{})
}

// Remove given item
func (s *Float32) Remove(num float32) bool {
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

// Contains - Checks if item exists in set
func (s *Float32) Contains(num float32) bool {
	_, ok := s.nodeMap[num]
	return ok
}

// GetList - Get items
func (s *Float32) GetList() []float32 {
	nums := []float32{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

// Size - Get size of set
func (s *Float32) Size() int {
	return len(s.nodeMap)
}

// Union returns all the items that are in S or in S2
func (s *Float32) Union(s2 *Float32) *Float32 {
	s3 := Float32{}
	s3.nodeMap = make(map[float32]struct{})
	for i := range s.nodeMap {
		s3.nodeMap[i] = struct{}{}
	}
	for i := range s2.nodeMap {
		_, ok := s3.nodeMap[i]
		if !ok {
			s3.nodeMap[i] = struct{}{}
		}
	}
	return &s3
}

// Intersection returns common items in S and S2
func (s *Float32) Intersection(s2 *Float32) *Float32 {
	s3 := Float32{}
	s3.nodeMap = make(map[float32]struct{})
	for i := range s2.nodeMap {
		_, ok := s.nodeMap[i]
		if ok {
			s3.nodeMap[i] = struct{}{}
		}
	}
	return &s3
}

// Minus - s.Minus(s2) : all of S but not in S2
func (s *Float32) Minus(s2 *Float32) *Float32 {
	s3 := Float32{}
	s3.nodeMap = make(map[float32]struct{})
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			s3.nodeMap[i] = struct{}{}
		}
	}
	return &s3
}

// Subset checks if S is subset of S2
func (s *Float32) Subset(s2 *Float32) bool {
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			return false
		}
	}
	return true
}

// Superset checks if S is superset of S2
func (s *Float32) Superset(s2 *Float32) bool {
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
