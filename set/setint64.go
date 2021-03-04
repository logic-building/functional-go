package set

// Int64 - struct
type Int64 struct {
	nodeMap map[int64]struct{}
}

// NewInt64 creates set
func NewInt64(nums []int64) *Int64 {
	s := &Int64{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

// Add an item
func (s *Int64) Add(num int64) *Int64 {
	if s.nodeMap == nil {
		s.nodeMap = make(map[int64]struct{})
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = struct{}{}
	}
	return s
}

// Clear set
func (s *Int64) Clear() {
	s.nodeMap = make(map[int64]struct{})
}

// Remove an item
func (s *Int64) Remove(num int64) bool {
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

// Contains - Check if item exists in set
func (s *Int64) Contains(num int64) bool {
	_, ok := s.nodeMap[num]
	return ok
}

// GetList - Get set items
func (s *Int64) GetList() []int64 {
	nums := []int64{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

// Size - Get size of set
func (s *Int64) Size() int {
	return len(s.nodeMap)
}

// Union returns all the items that are in S or in S2
func (s *Int64) Union(s2 *Int64) *Int64 {
	s3 := Int64{}
	s3.nodeMap = make(map[int64]struct{})
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
func (s *Int64) Intersection(s2 *Int64) *Int64 {
	s3 := Int64{}
	s3.nodeMap = make(map[int64]struct{})
	for i := range s2.nodeMap {
		_, ok := s.nodeMap[i]
		if ok {
			s3.nodeMap[i] = struct{}{}
		}
	}
	return &s3
}

// Minus - s.Minus(s2) : all of S but not in S2
func (s *Int64) Minus(s2 *Int64) *Int64 {
	s3 := Int64{}
	s3.nodeMap = make(map[int64]struct{})
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			s3.nodeMap[i] = struct{}{}
		}
	}
	return &s3
}

// Subset checks if S is subset of S2
func (s *Int64) Subset(s2 *Int64) bool {
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			return false
		}
	}
	return true
}

// Superset checks if S is superset of S2
func (s *Int64) Superset(s2 *Int64) bool {
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
