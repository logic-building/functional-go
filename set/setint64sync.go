package set

import "sync"

// Int64Sync - struct
type Int64Sync struct {
	nodeMap map[int64]struct{}
	sync.RWMutex
}

// NewInt64Sync creates set
func NewInt64Sync(nums []int64) *Int64Sync {
	s := &Int64Sync{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

// Add an item
func (s *Int64Sync) Add(num int64) *Int64Sync {
	s.Lock()
	defer s.Unlock()
	if s.nodeMap == nil {
		s.nodeMap = make(map[int64]struct{})
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = struct{}{}
	}
	return s
}

// Clear set
func (s *Int64Sync) Clear() {
	s.Lock()
	defer s.Unlock()
	s.nodeMap = make(map[int64]struct{})
}

// Remove an item
func (s *Int64Sync) Remove(num int64) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

// Contains - Check if item exists in set
func (s *Int64Sync) Contains(num int64) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.nodeMap[num]
	return ok
}

// GetList - Get set items
func (s *Int64Sync) GetList() []int64 {
	s.RLock()
	defer s.RUnlock()
	nums := []int64{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

// Size - Get size of set
func (s *Int64Sync) Size() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.nodeMap)
}

// Union returns all the items that are in S or in S2
func (s *Int64Sync) Union(s2 *Int64Sync) *Int64Sync {
	s3 := Int64Sync{}
	s3.nodeMap = make(map[int64]struct{})
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
func (s *Int64Sync) Intersection(s2 *Int64Sync) *Int64Sync {
	s3 := Int64Sync{}
	s3.nodeMap = make(map[int64]struct{})
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
func (s *Int64Sync) Minus(s2 *Int64Sync) *Int64Sync {
	s3 := Int64Sync{}
	s3.nodeMap = make(map[int64]struct{})
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
func (s *Int64Sync) Subset(s2 *Int64Sync) bool {
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
func (s *Int64Sync) Superset(s2 *Int64Sync) bool {
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
