package set

// Int32 - struct
type Int32 struct {
	nodeMap map[int32]bool
}

// NewInt32 creates set
func NewInt32(nums []int32) *Int32 {
	s := &Int32{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

// Add an item
func (s *Int32) Add(num int32) *Int32 {
	if s.nodeMap == nil {
		s.nodeMap = make(map[int32]bool)
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = true
	}
	return s
}

// Clear set
func (s *Int32) Clear() {
	s.nodeMap = make(map[int32]bool)
}

// Remove an item
func (s *Int32) Remove(num int32) bool {
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

// Contains - Check if item exists in set
func (s *Int32) Contains(num int32) bool {
	_, ok := s.nodeMap[num]
	return ok
}

// GetList - Get set items
func (s *Int32) GetList() []int32 {
	nums := []int32{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

// Size - Get size of set
func (s *Int32) Size() int {
	return len(s.nodeMap)
}

// Union returns all the items that are in S or in S2
func (s *Int32) Union(s2 *Int32) *Int32 {
	s3 := Int32{}
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

// Intersection returns common items in S and S2
func (s *Int32) Intersection(s2 *Int32) *Int32 {
	s3 := Int32{}
	s3.nodeMap = make(map[int32]bool)
	for i := range s2.nodeMap {
		_, ok := s.nodeMap[i]
		if ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

// Minus - s.Minus(s2) : all of S but not in S2
func (s *Int32) Minus(s2 *Int32) *Int32 {
	s3 := Int32{}
	s3.nodeMap = make(map[int32]bool)
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

// Subset checks if S is subset of S2
func (s *Int32) Subset(s2 *Int32) bool {
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			return false
		}
	}
	return true
}

// Superset checks if S is superset of S2
func (s *Int32) Superset(s2 *Int32) bool {
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
