package set

type SetUint16 struct {
	nodeMap map[uint16]bool
}

// Create set object
func NewUint16(nums []uint16) *SetUint16 {
	s := &SetUint16{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

// Add an item
func (s *SetUint16) Add(num uint16) *SetUint16 {
	if s.nodeMap == nil {
		s.nodeMap = make(map[uint16]bool)
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = true
	}
	return s
}

// Make object empty
func (s *SetUint16) Clear() {
	s.nodeMap = make(map[uint16]bool)
}

// Remove an item
func (s *SetUint16) Remove(num uint16) bool {
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

// Check if item exists in set
func (s *SetUint16) Contains(num uint16) bool {
	_, ok := s.nodeMap[num]
	return ok
}

// Get set items
func (s *SetUint16) GetList() []uint16 {
	nums := []uint16{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

// Get size of set
func (s *SetUint16) Size() int {
	return len(s.nodeMap)
}

// Returns all the items that are in S or in S2
func (s *SetUint16) Union(s2 *SetUint16) *SetUint16 {
	s3 := SetUint16{}
	s3.nodeMap = make(map[uint16]bool)
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
func (s *SetUint16) Intersection(s2 *SetUint16) *SetUint16 {
	s3 := SetUint16{}
	s3.nodeMap = make(map[uint16]bool)
	for i := range s2.nodeMap {
		_, ok := s.nodeMap[i]
		if ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

// s.Minus(s2) : all of S but not in S2
func (s *SetUint16) Minus(s2 *SetUint16) *SetUint16 {
	s3 := SetUint16{}
	s3.nodeMap = make(map[uint16]bool)
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

// Checks if S is subset of S2
func (s *SetUint16) Subset(s2 *SetUint16) bool {
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			return false
		}
	}
	return true
}

// Checks if S is superset of S2
func (s *SetUint16) Superset(s2 *SetUint16) bool {
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
