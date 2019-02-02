package set

type SetFloat64 struct {
	nodeMap map[float64]bool
}

// Create set object
func NewFloat64(nums []float64) *SetFloat64 {
	s := &SetFloat64{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

// Add an item
func (s *SetFloat64) Add(num float64) *SetFloat64 {
	if s.nodeMap == nil {
		s.nodeMap = make(map[float64]bool)
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = true
	}
	return s
}

// Make object empty
func (s *SetFloat64) Clear() {
	s.nodeMap = make(map[float64]bool)
}

// Remove an item
func (s *SetFloat64) Remove(num float64) bool {
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

// Check if item exists in set
func (s *SetFloat64) Contains(num float64) bool {
	_, ok := s.nodeMap[num]
	return ok
}

// Get set items
func (s *SetFloat64) GetList() []float64 {
	nums := []float64{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

// Get size of set
func (s *SetFloat64) Size() int {
	return len(s.nodeMap)
}

// Returns all the items that are in S or in S2
func (s *SetFloat64) Union(s2 *SetFloat64) *SetFloat64 {
	s3 := SetFloat64{}
	s3.nodeMap = make(map[float64]bool)
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

// Common items in S and S2
func (s *SetFloat64) Intersection(s2 *SetFloat64) *SetFloat64 {
	s3 := SetFloat64{}
	s3.nodeMap = make(map[float64]bool)
	for i := range s2.nodeMap {
		_, ok := s.nodeMap[i]
		if ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

// s.Minus(s2) : all of S but not in S2
func (s *SetFloat64) Minus(s2 *SetFloat64) *SetFloat64 {
	s3 := SetFloat64{}
	s3.nodeMap = make(map[float64]bool)
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

// Checks if S is subset of S2
func (s *SetFloat64) Subset(s2 *SetFloat64) bool {
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			return false
		}
	}
	return true
}

// Checks if S is superset of S2
func (s *SetFloat64) Superset(s2 *SetFloat64) bool {
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
