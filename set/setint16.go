package set

// Int16 - struct
type Int16 struct {
	nodeMap map[int16]struct{}
}

// NewInt16 creates
func NewInt16(nums []int16) *Int16 {
	s := &Int16{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

// Add an item
func (s *Int16) Add(num int16) *Int16 {
	if s.nodeMap == nil {
		s.nodeMap = make(map[int16]struct{})
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = struct{}{}
	}
	return s
}

// Clear set
func (s *Int16) Clear() {
	s.nodeMap = make(map[int16]struct{})
}

// Remove an item
func (s *Int16) Remove(num int16) bool {
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

// Contains - check if item exists in set
func (s *Int16) Contains(num int16) bool {
	_, ok := s.nodeMap[num]
	return ok
}

// GetList - Get set items
func (s *Int16) GetList() []int16 {
	nums := []int16{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

// Size - GetList - Get size of set
func (s *Int16) Size() int {
	return len(s.nodeMap)
}

// Union returns all the items that are in S or in S2
func (s *Int16) Union(s2 *Int16) *Int16 {
	s3 := Int16{}
	s3.nodeMap = make(map[int16]struct{})
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
func (s *Int16) Intersection(s2 *Int16) *Int16 {
	s3 := Int16{}
	s3.nodeMap = make(map[int16]struct{})
	for i := range s2.nodeMap {
		_, ok := s.nodeMap[i]
		if ok {
			s3.nodeMap[i] = struct{}{}
		}
	}
	return &s3
}

// Minus - s.Minus(s2) : all of S but not in S2
func (s *Int16) Minus(s2 *Int16) *Int16 {
	s3 := Int16{}
	s3.nodeMap = make(map[int16]struct{})
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			s3.nodeMap[i] = struct{}{}
		}
	}
	return &s3
}

// Subset if checks if S is subset of S2
func (s *Int16) Subset(s2 *Int16) bool {
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			return false
		}
	}
	return true
}

// Superset checks if S is superset of S2
func (s *Int16) Superset(s2 *Int16) bool {
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
