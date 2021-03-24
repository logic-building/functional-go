package set

// Uint16 - struct
type Uint16 struct {
	nodeMap map[uint16]struct{}
}

// NewUint16 creates set
func NewUint16(nums []uint16) *Uint16 {
	s := &Uint16{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

// Add an item
func (s *Uint16) Add(num uint16) *Uint16 {
	if s.nodeMap == nil {
		s.nodeMap = make(map[uint16]struct{})
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = struct{}{}
	}
	return s
}

// Clear set
func (s *Uint16) Clear() {
	s.nodeMap = make(map[uint16]struct{})
}

// Remove an item
func (s *Uint16) Remove(num uint16) bool {
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

// Contains - Check if item exists in set
func (s *Uint16) Contains(num uint16) bool {
	_, ok := s.nodeMap[num]
	return ok
}

// GetList - Get set items
func (s *Uint16) GetList() []uint16 {
	nums := []uint16{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

// Size - Get size of set
func (s *Uint16) Size() int {
	return len(s.nodeMap)
}

// Union returns all the items that are in S or in S2
func (s *Uint16) Union(s2 *Uint16) *Uint16 {
	s3 := Uint16{}
	s3.nodeMap = make(map[uint16]struct{})
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
func (s *Uint16) Intersection(s2 *Uint16) *Uint16 {
	s3 := Uint16{}
	s3.nodeMap = make(map[uint16]struct{})
	for i := range s2.nodeMap {
		_, ok := s.nodeMap[i]
		if ok {
			s3.nodeMap[i] = struct{}{}
		}
	}
	return &s3
}

// Minus - s.Minus(s2) : all of S but not in S2
func (s *Uint16) Minus(s2 *Uint16) *Uint16 {
	s3 := Uint16{}
	s3.nodeMap = make(map[uint16]struct{})
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			s3.nodeMap[i] = struct{}{}
		}
	}
	return &s3
}

// Subset checks if S is subset of S2
func (s *Uint16) Subset(s2 *Uint16) bool {
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			return false
		}
	}
	return true
}

// Superset checks if S is superset of S2
func (s *Uint16) Superset(s2 *Uint16) bool {
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
