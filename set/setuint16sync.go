package set

import "sync"

// Uint16Sync - struct
type Uint16Sync struct {
	nodeMap map[uint16]bool
	sync.RWMutex
}

// NewUint16Sync creates set
func NewUint16Sync(nums []uint16) *Uint16Sync {
	s := &Uint16Sync{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

// Add an item
func (s *Uint16Sync) Add(num uint16) *Uint16Sync {
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

// Clear set
func (s *Uint16Sync) Clear() {
	s.Lock()
	defer s.Unlock()
	s.nodeMap = make(map[uint16]bool)
}

// Remove an item
func (s *Uint16Sync) Remove(num uint16) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

// Contains - Check if item exists in set
func (s *Uint16Sync) Contains(num uint16) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.nodeMap[num]
	return ok
}

// GetList - Get set items
func (s *Uint16Sync) GetList() []uint16 {
	s.RLock()
	defer s.RUnlock()
	nums := []uint16{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

// Size - Get size of set
func (s *Uint16Sync) Size() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.nodeMap)
}

// Union returns all the items that are in S or in S2
func (s *Uint16Sync) Union(s2 *Uint16Sync) *Uint16Sync {
	s3 := Uint16Sync{}
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

// Intersection returns common items in S and S2
func (s *Uint16Sync) Intersection(s2 *Uint16Sync) *Uint16Sync {
	s3 := Uint16Sync{}
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

// Minus - s.Minus(s2) : all of S but not in S2
func (s *Uint16Sync) Minus(s2 *Uint16Sync) *Uint16Sync {
	s3 := Uint16Sync{}
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

// Subset checks if S is subset of S2
func (s *Uint16Sync) Subset(s2 *Uint16Sync) bool {
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
func (s *Uint16Sync) Superset(s2 *Uint16Sync) bool {
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
