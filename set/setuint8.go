package set

type SetUint8 struct {
	nodeMap map[uint8]bool
}

// Create set object
func NewUint8(nums []uint8) *SetUint8 {
	s := &SetUint8{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

// Add an item
func (s *SetUint8) Add(num uint8) *SetUint8 {
	if s.nodeMap == nil {
		s.nodeMap = make(map[uint8]bool)
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = true
	}
	return s
}

// Make object empty
func (s *SetUint8) Clear() {
	s.nodeMap = make(map[uint8]bool)
}

// Remove an item
func (s *SetUint8) Remove(num uint8) bool {
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

// Check if item exists in set
func (s *SetUint8) Contains(num uint8) bool {
	_, ok := s.nodeMap[num]
	return ok
}

// Get set items
func (s *SetUint8) GetList() []uint8 {
	nums := []uint8{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

// Get size of set
func (s *SetUint8) Size() int {
	return len(s.nodeMap)
}

// Returns all the items that are in S or in S2
func (s *SetUint8) Union(s2 *SetUint8) *SetUint8 {
	s3 := SetUint8{}
	s3.nodeMap = make(map[uint8]bool)
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

// s.Minus(s2) : all of S but not in S2
func (s *SetUint8) Intersection(s2 *SetUint8) *SetUint8 {
	s3 := SetUint8{}
	s3.nodeMap = make(map[uint8]bool)
	for i := range s2.nodeMap {
		_, ok := s.nodeMap[i]
		if ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

// s.Minus(s2) : all of S but not in S2
func (s *SetUint8) Minus(s2 *SetUint8) *SetUint8 {
	s3 := SetUint8{}
	s3.nodeMap = make(map[uint8]bool)
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

// Checks if S is subset of S2
func (s *SetUint8) Subset(s2 *SetUint8) bool {
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			return false
		}
	}
	return true
}

// Checks if S is superset of S2
func (s *SetUint8) Superset(s2 *SetUint8) bool {
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
