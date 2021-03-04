package set

// Uint - struct
type Uint struct {
	nodeMap map[uint]struct{}
}

// NewUint creates set
func NewUint(nums []uint) *Uint {
	s := &Uint{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

// Add an item
func (s *Uint) Add(num uint) *Uint {
	if s.nodeMap == nil {
		s.nodeMap = make(map[uint]struct{})
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = struct{}{}
	}
	return s
}

// Clear set
func (s *Uint) Clear() {
	s.nodeMap = make(map[uint]struct{})
}

// Remove an item
func (s *Uint) Remove(num uint) bool {
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

// Contains - Check if item exists in set
func (s *Uint) Contains(num uint) bool {
	_, ok := s.nodeMap[num]
	return ok
}

// GetList - Get set items
func (s *Uint) GetList() []uint {
	nums := []uint{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

// Size - Get size of set
func (s *Uint) Size() int {
	return len(s.nodeMap)
}

// Union returns all the items that are in S or in S2
func (s *Uint) Union(s2 *Uint) *Uint {
	s3 := Uint{}
	s3.nodeMap = make(map[uint]struct{})
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
func (s *Uint) Intersection(s2 *Uint) *Uint {
	s3 := Uint{}
	s3.nodeMap = make(map[uint]struct{})
	for i := range s2.nodeMap {
		_, ok := s.nodeMap[i]
		if ok {
			s3.nodeMap[i] = struct{}{}
		}
	}
	return &s3
}

// Minus - s.Minus(s2) : all of S but not in S2
func (s *Uint) Minus(s2 *Uint) *Uint {
	s3 := Uint{}
	s3.nodeMap = make(map[uint]struct{})
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			s3.nodeMap[i] = struct{}{}
		}
	}
	return &s3
}

// Subset checks if S is subset of S2
func (s *Uint) Subset(s2 *Uint) bool {
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			return false
		}
	}
	return true
}

// Superset checks if S is superset of S2
func (s *Uint) Superset(s2 *Uint) bool {
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
