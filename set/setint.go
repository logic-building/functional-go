package set

type SetInt struct {
	nodeMap map[int]bool
}

// Create set object
func NewInt(nums []int) *SetInt {
	s := &SetInt{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

// Add an item
func (s *SetInt) Add(num int) *SetInt {
	if s.nodeMap == nil {
		s.nodeMap = make(map[int]bool)
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = true
	}
	return s
}

// Make object empty
func (s *SetInt) Clear() {
	s.nodeMap = make(map[int]bool)
}

// Remove an item
func (s *SetInt) Remove(num int) bool {
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

// Check if item exists in set
func (s *SetInt) Contains(num int) bool {
	_, ok := s.nodeMap[num]
	return ok
}

// Get set items
func (s *SetInt) GetList() []int {
	nums := []int{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

// Get size of set
func (s *SetInt) Size() int {
	return len(s.nodeMap)
}

// Returns all the items that are in S or in S2
func (s *SetInt) Union(s2 *SetInt) *SetInt {
	s3 := SetInt{}
	s3.nodeMap = make(map[int]bool)
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
func (s *SetInt) Intersection(s2 *SetInt) *SetInt {
	s3 := SetInt{}
	s3.nodeMap = make(map[int]bool)
	for i := range s2.nodeMap {
		_, ok := s.nodeMap[i]
		if ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

// s.Minus(s2) : all of S but not in S2
func (s *SetInt) Minus(s2 *SetInt) *SetInt {
	s3 := SetInt{}
	s3.nodeMap = make(map[int]bool)
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

// Checks if S is subset of S2
func (s *SetInt) Subset(s2 *SetInt) bool {
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			return false
		}
	}
	return true
}

// Checks if S is superset of S2
func (s *SetInt) Superset(s2 *SetInt) bool {
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
