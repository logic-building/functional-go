package set

import "sync"

type SetUint16Sync struct {
	nodeMap map[uint16]bool
	sync.RWMutex
}

// Create set object
func NewUint16Sync(nums []uint16) *SetUint16Sync {
	s := &SetUint16Sync{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

// Add an item
func (s *SetUint16Sync) Add(num uint16) *SetUint16Sync {
	s.Lock()
	defer s.Unlock()
	if s.nodeMap == nil {
		s.nodeMap = make(map[uint16]bool)
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = true
	}
	return s
}

// Make object empty
func (s *SetUint16Sync) Clear() {
	s.Lock()
	defer s.Unlock()
	s.nodeMap = make(map[uint16]bool)
}

// Remove an item
func (s *SetUint16Sync) Remove(num uint16) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

// Check if item exists in set
func (s *SetUint16Sync) Contains(num uint16) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.nodeMap[num]
	return ok
}

// Get set items
func (s *SetUint16Sync) GetList() []uint16 {
	s.RLock()
	defer s.RUnlock()
	nums := []uint16{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

// Get size of set
func (s *SetUint16Sync) Size() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.nodeMap)
}

// Returns all the items that are in S or in S2
func (s *SetUint16Sync) Union(s2 *SetUint16Sync) *SetUint16Sync {
	s3 := SetUint16Sync{}
	s3.nodeMap = make(map[uint16]bool)
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
func (s *SetUint16Sync) Intersection(s2 *SetUint16Sync) *SetUint16Sync {
	s3 := SetUint16Sync{}
	s3.nodeMap = make(map[uint16]bool)
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
func (s *SetUint16Sync) Minus(s2 *SetUint16Sync) *SetUint16Sync {
	s3 := SetUint16Sync{}
	s3.nodeMap = make(map[uint16]bool)
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
func (s *SetUint16Sync) Subset(s2 *SetUint16Sync) bool {
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
func (s *SetUint16Sync) Superset(s2 *SetUint16Sync) bool {
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
