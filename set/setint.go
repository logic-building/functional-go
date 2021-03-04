package set

// Int - struct
type Int struct {
	nodeMap map[int]struct{}
}

// NewInt creates set
func NewInt(nums []int) *Int {
	s := &Int{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

// Add an item
func (s *Int) Add(num int) *Int {
	if s.nodeMap == nil {
		s.nodeMap = make(map[int]struct{})
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = struct{}{}
	}
	return s
}

// Clear set
func (s *Int) Clear() {
	s.nodeMap = make(map[int]struct{})
}

// Remove an item
func (s *Int) Remove(num int) bool {
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

// Contains - Check if item exists in set
func (s *Int) Contains(num int) bool {
	_, ok := s.nodeMap[num]
	return ok
}

// GetList - Get set items
func (s *Int) GetList() []int {
	nums := []int{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

// Size - Get size of set
func (s *Int) Size() int {
	return len(s.nodeMap)
}

// Union returns all the items that are in S or in S2
func (s *Int) Union(s2 *Int) *Int {
	s3 := Int{}
	s3.nodeMap = make(map[int]struct{})
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

// Intersection - Common items in S and S2
func (s *Int) Intersection(s2 *Int) *Int {
	s3 := Int{}
	s3.nodeMap = make(map[int]struct{})
	for i := range s2.nodeMap {
		_, ok := s.nodeMap[i]
		if ok {
			s3.nodeMap[i] = struct{}{}
		}
	}
	return &s3
}

// Minus - s.Minus(s2) : all of S but not in S2
func (s *Int) Minus(s2 *Int) *Int {
	s3 := Int{}
	s3.nodeMap = make(map[int]struct{})
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			s3.nodeMap[i] = struct{}{}
		}
	}
	return &s3
}

// Subset checks if S is subset of S2
func (s *Int) Subset(s2 *Int) bool {
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			return false
		}
	}
	return true
}

// Superset checks if S is superset of S2
func (s *Int) Superset(s2 *Int) bool {
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
