package set

import "sync"

type SetFloat32Sync struct {
	nodeMap map[float32]bool
	sync.RWMutex
}

// Create set object
func NewFloat32Sync(nums []float32) *SetFloat32Sync {
	s := &SetFloat32Sync{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

// Add an item
func (s *SetFloat32Sync) Add(num float32) *SetFloat32Sync {
	s.Lock()
	defer s.Unlock()
	if s.nodeMap == nil {
		s.nodeMap = make(map[float32]bool)
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = true
	}
	return s
}

// Make object empty
func (s *SetFloat32Sync) Clear() {
	s.Lock()
	defer s.Unlock()
	s.nodeMap = make(map[float32]bool)
}

// Remove an item
func (s *SetFloat32Sync) Remove(num float32) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

// Check if item exists in set
func (s *SetFloat32Sync) Contains(num float32) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.nodeMap[num]
	return ok
}

// Get set items
func (s *SetFloat32Sync) GetList() []float32 {
	s.RLock()
	defer s.RUnlock()
	nums := []float32{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

// Get size of set
func (s *SetFloat32Sync) Size() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.nodeMap)
}

// Returns all the items that are in S or in S2
func (s *SetFloat32Sync) Union(s2 *SetFloat32Sync) *SetFloat32Sync {
	s3 := SetFloat32Sync{}
	s3.nodeMap = make(map[float32]bool)
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
func (s *SetFloat32Sync) Intersection(s2 *SetFloat32Sync) *SetFloat32Sync {
	s3 := SetFloat32Sync{}
	s3.nodeMap = make(map[float32]bool)
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
func (s *SetFloat32Sync) Minus(s2 *SetFloat32Sync) *SetFloat32Sync {
	s3 := SetFloat32Sync{}
	s3.nodeMap = make(map[float32]bool)
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
func (s *SetFloat32Sync) Subset(s2 *SetFloat32Sync) bool {
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
func (s *SetFloat32Sync) Superset(s2 *SetFloat32Sync) bool {
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
