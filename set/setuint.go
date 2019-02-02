package set

type SetUint struct {
	nodeMap map[uint]bool
}

// Create set object
func NewUint(nums []uint) *SetUint {
	s := &SetUint{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

// Add an item
func (s *SetUint) Add(num uint) *SetUint {
	if s.nodeMap == nil {
		s.nodeMap = make(map[uint]bool)
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = true
	}
	return s
}

// Make object empty
func (s *SetUint) Clear() {
	s.nodeMap = make(map[uint]bool)
}

// Remove an item
func (s *SetUint) Remove(num uint) bool {
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

// Check if item exists in set
func (s *SetUint) Contains(num uint) bool {
	_, ok := s.nodeMap[num]
	return ok
}

// Get set items
func (s *SetUint) GetList() []uint {
	nums := []uint{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

// Get size of set
func (s *SetUint) Size() int {
	return len(s.nodeMap)
}

// Returns all the items that are in S or in S2
func (s *SetUint) Union(s2 *SetUint) *SetUint {
	s3 := SetUint{}
	s3.nodeMap = make(map[uint]bool)
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
func (s *SetUint) Intersection(s2 *SetUint) *SetUint {
	s3 := SetUint{}
	s3.nodeMap = make(map[uint]bool)
	for i := range s2.nodeMap {
		_, ok := s.nodeMap[i]
		if ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

// s.Minus(s2) : all of S but not in S2
func (s *SetUint) Minus(s2 *SetUint) *SetUint {
	s3 := SetUint{}
	s3.nodeMap = make(map[uint]bool)
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

// Checks if S is subset of S2
func (s *SetUint) Subset(s2 *SetUint) bool {
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			return false
		}
	}
	return true
}

// Checks if S is superset of S2
func (s *SetUint) Superset(s2 *SetUint) bool {
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
