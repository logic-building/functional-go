package set

import "sync"

// Float32Sync - struct
type Float32Sync struct {
	nodeMap map[float32]bool
	sync.RWMutex
}

// NewFloat32Sync creates new set
func NewFloat32Sync(nums []float32) *Float32Sync {
	s := &Float32Sync{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

// Add an item
func (s *Float32Sync) Add(num float32) *Float32Sync {
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

// Clear set
func (s *Float32Sync) Clear() {
	s.Lock()
	defer s.Unlock()
	s.nodeMap = make(map[float32]bool)
}

// Remove an item
func (s *Float32Sync) Remove(num float32) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

// Contains - Check if item exists in set
func (s *Float32Sync) Contains(num float32) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.nodeMap[num]
	return ok
}

// GetList - Get set items
func (s *Float32Sync) GetList() []float32 {
	s.RLock()
	defer s.RUnlock()
	nums := []float32{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

// Size - Get size of set
func (s *Float32Sync) Size() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.nodeMap)
}

// Union returns all the items that are in S or in S2
func (s *Float32Sync) Union(s2 *Float32Sync) *Float32Sync {
	s3 := Float32Sync{}
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

// Intersection returns common items in S and S2
func (s *Float32Sync) Intersection(s2 *Float32Sync) *Float32Sync {
	s3 := Float32Sync{}
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

// Minus - s.Minus(s2) : all of S but not in S2
func (s *Float32Sync) Minus(s2 *Float32Sync) *Float32Sync {
	s3 := Float32Sync{}
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

// Subset checks if S is subset of S2
func (s *Float32Sync) Subset(s2 *Float32Sync) bool {
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
func (s *Float32Sync) Superset(s2 *Float32Sync) bool {
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
