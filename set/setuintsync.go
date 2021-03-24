package set

import "sync"

// UintSync - struct
type UintSync struct {
	nodeMap map[uint]struct{}
	sync.RWMutex
}

// NewUintSync creates set
func NewUintSync(nums []uint) *UintSync {
	s := &UintSync{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

// Add an item
func (s *UintSync) Add(num uint) *UintSync {
	s.Lock()
	defer s.Unlock()
	if s.nodeMap == nil {
		s.nodeMap = make(map[uint]struct{})
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = struct{}{}
	}
	return s
}

// Clear set
func (s *UintSync) Clear() {
	s.Lock()
	defer s.Unlock()
	s.nodeMap = make(map[uint]struct{})
}

// Remove an item
func (s *UintSync) Remove(num uint) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

// Contains - Check if item exists in set
func (s *UintSync) Contains(num uint) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.nodeMap[num]
	return ok
}

// GetList - Get set items
func (s *UintSync) GetList() []uint {
	s.RLock()
	defer s.RUnlock()
	nums := []uint{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

// Size - Get size of set
func (s *UintSync) Size() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.nodeMap)
}

// Union returns all the items that are in S or in S2
func (s *UintSync) Union(s2 *UintSync) *UintSync {
	s3 := UintSync{}
	s3.nodeMap = make(map[uint]struct{})
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
func (s *UintSync) Intersection(s2 *UintSync) *UintSync {
	s3 := UintSync{}
	s3.nodeMap = make(map[uint]struct{})
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
func (s *UintSync) Minus(s2 *UintSync) *UintSync {
	s3 := UintSync{}
	s3.nodeMap = make(map[uint]struct{})
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
func (s *UintSync) Subset(s2 *UintSync) bool {
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
func (s *UintSync) Superset(s2 *UintSync) bool {
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
