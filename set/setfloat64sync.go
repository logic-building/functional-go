package set

import "sync"

type SetFloat64Sync struct {
	nodeMap map[float64]bool
	sync.RWMutex
}

// Create set object
func NewFloat64Sync(nums []float64) *SetFloat64Sync {
	s := &SetFloat64Sync{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

// Add an item
func (s *SetFloat64Sync) Add(num float64) *SetFloat64Sync {
	s.Lock()
	defer s.Unlock()
	if s.nodeMap == nil {
		s.nodeMap = make(map[float64]bool)
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = true
	}
	return s
}

// Make object empty
func (s *SetFloat64Sync) Clear() {
	s.Lock()
	defer s.Unlock()
	s.nodeMap = make(map[float64]bool)
}

// Remove an item
func (s *SetFloat64Sync) Remove(num float64) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

// Check if item exists in set
func (s *SetFloat64Sync) Contains(num float64) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.nodeMap[num]
	return ok
}

// Get set items
func (s *SetFloat64Sync) GetList() []float64 {
	s.RLock()
	defer s.RUnlock()
	nums := []float64{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

// Get size of set
func (s *SetFloat64Sync) Size() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.nodeMap)
}

// Returns all the items that are in S or in S2
func (s *SetFloat64Sync) Union(s2 *SetFloat64Sync) *SetFloat64Sync {
	s3 := SetFloat64Sync{}
	s3.nodeMap = make(map[float64]bool)
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

// Common items in S and S2
func (s *SetFloat64Sync) Intersection(s2 *SetFloat64Sync) *SetFloat64Sync {
	s3 := SetFloat64Sync{}
	s3.nodeMap = make(map[float64]bool)
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

// s.Minus(s2) : all of S but not in S2
func (s *SetFloat64Sync) Minus(s2 *SetFloat64Sync) *SetFloat64Sync {
	s3 := SetFloat64Sync{}
	s3.nodeMap = make(map[float64]bool)
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

// Checks if S is subset of S2
func (s *SetFloat64Sync) Subset(s2 *SetFloat64Sync) bool {
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

// Checks if S is superset of S2
func (s *SetFloat64Sync) Superset(s2 *SetFloat64Sync) bool {
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
