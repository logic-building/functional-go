package set

// Uint8 - struct
type Uint8 struct {
	nodeMap map[uint8]struct{}
}

// NewUint8 creates set
func NewUint8(nums []uint8) *Uint8 {
	s := &Uint8{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

// Add an item
func (s *Uint8) Add(num uint8) *Uint8 {
	if s.nodeMap == nil {
		s.nodeMap = make(map[uint8]struct{})
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = struct{}{}
	}
	return s
}

// Clear set
func (s *Uint8) Clear() {
	s.nodeMap = make(map[uint8]struct{})
}

// Remove an item
func (s *Uint8) Remove(num uint8) bool {
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

// Contains - Check if item exists in set
func (s *Uint8) Contains(num uint8) bool {
	_, ok := s.nodeMap[num]
	return ok
}

// GetList - Get set items
func (s *Uint8) GetList() []uint8 {
	nums := []uint8{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

// Size - Get size of set
func (s *Uint8) Size() int {
	return len(s.nodeMap)
}

// Union returns all the items that are in S or in S2
func (s *Uint8) Union(s2 *Uint8) *Uint8 {
	s3 := Uint8{}
	s3.nodeMap = make(map[uint8]struct{})
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
func (s *Uint8) Intersection(s2 *Uint8) *Uint8 {
	s3 := Uint8{}
	s3.nodeMap = make(map[uint8]struct{})
	for i := range s2.nodeMap {
		_, ok := s.nodeMap[i]
		if ok {
			s3.nodeMap[i] = struct{}{}
		}
	}
	return &s3
}

// Minus - s.Minus(s2) : all of S but not in S2
func (s *Uint8) Minus(s2 *Uint8) *Uint8 {
	s3 := Uint8{}
	s3.nodeMap = make(map[uint8]struct{})
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			s3.nodeMap[i] = struct{}{}
		}
	}
	return &s3
}

// Subset checks if S is subset of S2
func (s *Uint8) Subset(s2 *Uint8) bool {
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			return false
		}
	}
	return true
}

// Superset checks if S is superset of S2
func (s *Uint8) Superset(s2 *Uint8) bool {
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
