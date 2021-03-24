package set

import "sync"

// Uint32Sync - struct
type Uint32Sync struct {
	nodeMap map[uint32]struct{}
	sync.RWMutex
}

// NewUint32Sync create set
func NewUint32Sync(nums []uint32) *Uint32Sync {
	s := &Uint32Sync{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

// Add an item
func (s *Uint32Sync) Add(num uint32) *Uint32Sync {
	s.Lock()
	defer s.Unlock()
	if s.nodeMap == nil {
		s.nodeMap = make(map[uint32]struct{})
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = struct{}{}
	}
	return s
}

// Clear set
func (s *Uint32Sync) Clear() {
	s.Lock()
	defer s.Unlock()
	s.nodeMap = make(map[uint32]struct{})
}

// Remove an item
func (s *Uint32Sync) Remove(num uint32) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

// Contains - Check if item exists in set
func (s *Uint32Sync) Contains(num uint32) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.nodeMap[num]
	return ok
}

// GetList - Get set items
func (s *Uint32Sync) GetList() []uint32 {
	s.RLock()
	defer s.RUnlock()
	nums := []uint32{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

// Size - Get size of set
func (s *Uint32Sync) Size() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.nodeMap)
}

// Union returns all the items that are in S or in S2
func (s *Uint32Sync) Union(s2 *Uint32Sync) *Uint32Sync {
	s3 := Uint32Sync{}
	s3.nodeMap = make(map[uint32]struct{})
	s.RLock()
	for i := range s.nodeMap {
		s3.nodeMap[i] = struct{}{}
	}
	s.RUnlock()
	s2.RLock()
	for i := range s2.nodeMap {
		_, ok := s3.nodeMap[i]
		if !ok {
			s3.nodeMap[i] = struct{}{}
		}
	}
	s2.RUnlock()
	return &s3
}

// Intersection returns common items in S and S2
func (s *Uint32Sync) Intersection(s2 *Uint32Sync) *Uint32Sync {
	s3 := Uint32Sync{}
	s3.nodeMap = make(map[uint32]struct{})
	s.RLock()
	s2.RLock()
	defer s.RUnlock()
	defer s2.RUnlock()
	for i := range s2.nodeMap {
		_, ok := s.nodeMap[i]
		if ok {
			s3.nodeMap[i] = struct{}{}
		}
	}
	return &s3
}

// Minus - s.Minus(s2) : all of S but not in S2
func (s *Uint32Sync) Minus(s2 *Uint32Sync) *Uint32Sync {
	s3 := Uint32Sync{}
	s3.nodeMap = make(map[uint32]struct{})
	s.RLock()
	s2.RLock()
	defer s.RUnlock()
	defer s2.RUnlock()
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			s3.nodeMap[i] = struct{}{}
		}
	}
	return &s3
}

// Subset checks if S is subset of S2
func (s *Uint32Sync) Subset(s2 *Uint32Sync) bool {
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
func (s *Uint32Sync) Superset(s2 *Uint32Sync) bool {
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
