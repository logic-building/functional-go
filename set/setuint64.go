package set

// Uint64 - struct
type Uint64 struct {
	nodeMap map[uint64]bool
}

// NewUint64 creates set
func NewUint64(nums []uint64) *Uint64 {
	s := &Uint64{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

// Add an item
func (s *Uint64) Add(num uint64) *Uint64 {
	if s.nodeMap == nil {
		s.nodeMap = make(map[uint64]bool)
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = true
	}
	return s
}

// Clear set
func (s *Uint64) Clear() {
	s.nodeMap = make(map[uint64]bool)
}

// Remove an item
func (s *Uint64) Remove(num uint64) bool {
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

// Contains - Check if item exists in set
func (s *Uint64) Contains(num uint64) bool {
	_, ok := s.nodeMap[num]
	return ok
}

// GetList - Get set items
func (s *Uint64) GetList() []uint64 {
	nums := []uint64{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

// Size - Get size of set
func (s *Uint64) Size() int {
	return len(s.nodeMap)
}

// Union returns all the items that are in S or in S2
func (s *Uint64) Union(s2 *Uint64) *Uint64 {
	s3 := Uint64{}
	s3.nodeMap = make(map[uint64]bool)
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
func (s *Uint64) Intersection(s2 *Uint64) *Uint64 {
	s3 := Uint64{}
	s3.nodeMap = make(map[uint64]bool)
	for i := range s2.nodeMap {
		_, ok := s.nodeMap[i]
		if ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

// Minus - s.Minus(s2) : all of S but not in S2
func (s *Uint64) Minus(s2 *Uint64) *Uint64 {
	s3 := Uint64{}
	s3.nodeMap = make(map[uint64]bool)
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

// Subset checks if S is subset of S2
func (s *Uint64) Subset(s2 *Uint64) bool {
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			return false
		}
	}
	return true
}

// Superset checks if S is superset of S2
func (s *Uint64) Superset(s2 *Uint64) bool {
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
