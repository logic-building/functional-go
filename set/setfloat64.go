package set

// Float64 - struct
type Float64 struct {
	nodeMap map[float64]bool
}

// NewFloat64 creates new set
func NewFloat64(nums []float64) *Float64 {
	s := &Float64{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

// Add an item
func (s *Float64) Add(num float64) *Float64 {
	if s.nodeMap == nil {
		s.nodeMap = make(map[float64]bool)
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = true
	}
	return s
}

// Clear set
func (s *Float64) Clear() {
	s.nodeMap = make(map[float64]bool)
}

// Remove given item
func (s *Float64) Remove(num float64) bool {
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

// Contains - Check if item exists in set
func (s *Float64) Contains(num float64) bool {
	_, ok := s.nodeMap[num]
	return ok
}

// GetList - Get set items
func (s *Float64) GetList() []float64 {
	nums := []float64{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

// Size - Get size of set
func (s *Float64) Size() int {
	return len(s.nodeMap)
}

// Union returns all the items that are in S or in S2
func (s *Float64) Union(s2 *Float64) *Float64 {
	s3 := Float64{}
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

// Intersection returns common items in S and S2
func (s *Float64) Intersection(s2 *Float64) *Float64 {
	s3 := Float64{}
	s3.nodeMap = make(map[float64]bool)
	for i := range s2.nodeMap {
		_, ok := s.nodeMap[i]
		if ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

// Minus - s.Minus(s2) : all of S but not in S2
func (s *Float64) Minus(s2 *Float64) *Float64 {
	s3 := Float64{}
	s3.nodeMap = make(map[float64]bool)
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

// Subset checks if S is subset of S2
func (s *Float64) Subset(s2 *Float64) bool {
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			return false
		}
	}
	return true
}

// Superset checks if S is superset of S2
func (s *Float64) Superset(s2 *Float64) bool {
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
