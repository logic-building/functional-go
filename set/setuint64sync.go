package set

import "sync"

// Uint64Sync - struct
type Uint64Sync struct {
	nodeMap map[uint64]bool
	sync.RWMutex
}

// NewUint64Sync creates set
func NewUint64Sync(nums []uint64) *Uint64Sync {
	s := &Uint64Sync{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

// Add an item
func (s *Uint64Sync) Add(num uint64) *Uint64Sync {
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

// Clear set
func (s *Uint64Sync) Clear() {
	s.Lock()
	defer s.Unlock()
	s.nodeMap = make(map[uint64]bool)
}

// Remove an item
func (s *Uint64Sync) Remove(num uint64) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

// Contains - Check if item exists in set
func (s *Uint64Sync) Contains(num uint64) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.nodeMap[num]
	return ok
}

// GetList - Get set items
func (s *Uint64Sync) GetList() []uint64 {
	s.RLock()
	defer s.RUnlock()
	nums := []uint64{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

// Size - Get size of set
func (s *Uint64Sync) Size() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.nodeMap)
}

// Union returns all the items that are in S or in S2
func (s *Uint64Sync) Union(s2 *Uint64Sync) *Uint64Sync {
	s3 := Uint64Sync{}
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

// Intersection returns common items in S and S2
func (s *Uint64Sync) Intersection(s2 *Uint64Sync) *Uint64Sync {
	s3 := Uint64Sync{}
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

// Minus - s.Minus(s2) : all of S but not in S2
func (s *Uint64Sync) Minus(s2 *Uint64Sync) *Uint64Sync {
	s3 := Uint64Sync{}
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

// Subset checks if S is subset of S2
func (s *Uint64Sync) Subset(s2 *Uint64Sync) bool {
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
func (s *Uint64Sync) Superset(s2 *Uint64Sync) bool {
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
