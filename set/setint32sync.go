package set

import "sync"

// Int32Sync - struct
type Int32Sync struct {
	nodeMap map[int32]bool
	sync.RWMutex
}

// NewInt32Sync creates set
func NewInt32Sync(nums []int32) *Int32Sync {
	s := &Int32Sync{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

// Add an item
func (s *Int32Sync) Add(num int32) *Int32Sync {
	s.Lock()
	defer s.Unlock()
	if s.nodeMap == nil {
		s.nodeMap = make(map[int32]bool)
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = true
	}
	return s
}

// Clear set
func (s *Int32Sync) Clear() {
	s.Lock()
	defer s.Unlock()
	s.nodeMap = make(map[int32]bool)
}

// Remove an item
func (s *Int32Sync) Remove(num int32) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

// Contains - Check if item exists in set
func (s *Int32Sync) Contains(num int32) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.nodeMap[num]
	return ok
}

// GetList - Get set items
func (s *Int32Sync) GetList() []int32 {
	s.RLock()
	defer s.RUnlock()
	nums := []int32{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

// Size - Get size of set
func (s *Int32Sync) Size() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.nodeMap)
}

// Union returns all the items that are in S or in S2
func (s *Int32Sync) Union(s2 *Int32Sync) *Int32Sync {
	s3 := Int32Sync{}
	s3.nodeMap = make(map[int32]bool)
	s.RLock()
	for i := range s.nodeMap {
		s3.nodeMap[i] = true
	}
	s.RUnlock()
	s2.RLock()
	for i := range s2.nodeMap {
		_, ok := s3.nodeMap[i]
		if !ok {
			s3.nodeMap[i] = true
		}
	}
	s2.RUnlock()
	return &s3
}

// Intersection returns common items in S and S2
func (s *Int32Sync) Intersection(s2 *Int32Sync) *Int32Sync {
	s3 := Int32Sync{}
	s3.nodeMap = make(map[int32]bool)
	s.RLock()
	s2.RLock()
	defer s.RUnlock()
	defer s2.RUnlock()
	for i := range s2.nodeMap {
		_, ok := s.nodeMap[i]
		if ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

// Minus - s.Minus(s2) : all of S but not in S2
func (s *Int32Sync) Minus(s2 *Int32Sync) *Int32Sync {
	s3 := Int32Sync{}
	s3.nodeMap = make(map[int32]bool)
	s.RLock()
	s2.RLock()
	defer s.RUnlock()
	defer s2.RUnlock()
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

// Subset checks if S is subset of S2
func (s *Int32Sync) Subset(s2 *Int32Sync) bool {
	s.RLock()
	s2.RLock()
	defer s.RUnlock()
	defer s2.RUnlock()
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			return false
		}
	}
	return true
}

// Superset checks if S is superset of S2
func (s *Int32Sync) Superset(s2 *Int32Sync) bool {
	s.RLock()
	s2.RLock()
	defer s.RUnlock()
	defer s2.RUnlock()
	for i := range s2.nodeMap {
		_, ok := s.nodeMap[i]
		if !ok {
			return false
		}
	}
	return true
}
