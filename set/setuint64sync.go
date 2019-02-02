package set

import "sync"

type SetUint64Sync struct {
	nodeMap map[uint64]bool
	sync.RWMutex
}

// Create set object
func NewUint64Sync(nums []uint64) *SetUint64Sync {
	s := &SetUint64Sync{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

// Add an item
func (s *SetUint64Sync) Add(num uint64) *SetUint64Sync {
	s.Lock()
	defer s.Unlock()
	if s.nodeMap == nil {
		s.nodeMap = make(map[uint64]bool)
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = true
	}
	return s
}

// Make object empty
func (s *SetUint64Sync) Clear() {
	s.Lock()
	defer s.Unlock()
	s.nodeMap = make(map[uint64]bool)
}

// Remove an item
func (s *SetUint64Sync) Remove(num uint64) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

// Check if item exists in set
func (s *SetUint64Sync) Contains(num uint64) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.nodeMap[num]
	return ok
}

// Get set items
func (s *SetUint64Sync) GetList() []uint64 {
	s.RLock()
	defer s.RUnlock()
	nums := []uint64{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

// Get size of set
func (s *SetUint64Sync) Size() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.nodeMap)
}

// Returns all the items that are in S or in S2
func (s *SetUint64Sync) Union(s2 *SetUint64Sync) *SetUint64Sync {
	s3 := SetUint64Sync{}
	s3.nodeMap = make(map[uint64]bool)
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
func (s *SetUint64Sync) Intersection(s2 *SetUint64Sync) *SetUint64Sync {
	s3 := SetUint64Sync{}
	s3.nodeMap = make(map[uint64]bool)
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
func (s *SetUint64Sync) Minus(s2 *SetUint64Sync) *SetUint64Sync {
	s3 := SetUint64Sync{}
	s3.nodeMap = make(map[uint64]bool)
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
func (s *SetUint64Sync) Subset(s2 *SetUint64Sync) bool {
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
func (s *SetUint64Sync) Superset(s2 *SetUint64Sync) bool {
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
