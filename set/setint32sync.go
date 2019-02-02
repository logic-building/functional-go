package set

import "sync"

type SetInt32Sync struct {
	nodeMap map[int32]bool
	sync.RWMutex
}

// Create set object
func NewInt32Sync(nums []int32) *SetInt32Sync {
	s := &SetInt32Sync{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

// Add an item
func (s *SetInt32Sync) Add(num int32) *SetInt32Sync {
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

// Make object empty
func (s *SetInt32Sync) Clear() {
	s.Lock()
	defer s.Unlock()
	s.nodeMap = make(map[int32]bool)
}

// Remove an item
func (s *SetInt32Sync) Remove(num int32) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

// Check if item exists in set
func (s *SetInt32Sync) Contains(num int32) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.nodeMap[num]
	return ok
}

// Get set items
func (s *SetInt32Sync) GetList() []int32 {
	s.RLock()
	defer s.RUnlock()
	nums := []int32{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

// Get size of set
func (s *SetInt32Sync) Size() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.nodeMap)
}

// Returns all the items that are in S or in S2
func (s *SetInt32Sync) Union(s2 *SetInt32Sync) *SetInt32Sync {
	s3 := SetInt32Sync{}
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

// Common items in S and S2
func (s *SetInt32Sync) Intersection(s2 *SetInt32Sync) *SetInt32Sync {
	s3 := SetInt32Sync{}
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

// s.Minus(s2) : all of S but not in S2
func (s *SetInt32Sync) Minus(s2 *SetInt32Sync) *SetInt32Sync {
	s3 := SetInt32Sync{}
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

// Checks if S is subset of S2
func (s *SetInt32Sync) Subset(s2 *SetInt32Sync) bool {
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
func (s *SetInt32Sync) Superset(s2 *SetInt32Sync) bool {
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
