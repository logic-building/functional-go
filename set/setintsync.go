package set

import "sync"

// IntSync - struct
type IntSync struct {
	nodeMap map[int]struct{}
	sync.RWMutex
}

// NewIntSync creates set
func NewIntSync(nums []int) *IntSync {
	s := &IntSync{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

// Add an item
func (s *IntSync) Add(num int) *IntSync {
	s.Lock()
	defer s.Unlock()
	if s.nodeMap == nil {
		s.nodeMap = make(map[int]struct{})
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = struct{}{}
	}
	return s
}

// Clear set
func (s *IntSync) Clear() {
	s.Lock()
	defer s.Unlock()
	s.nodeMap = make(map[int]struct{})
}

// Remove an item
func (s *IntSync) Remove(num int) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

// Contains - Check if item exists in set
func (s *IntSync) Contains(num int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.nodeMap[num]
	return ok
}

// getList - Get set items
func (s *IntSync) getList() []int {
	s.RLock()
	defer s.RUnlock()
	nums := []int{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

// Size - Get size of set
func (s *IntSync) Size() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.nodeMap)
}

// Union returns all the items that are in S or in S2
func (s *IntSync) Union(s2 *IntSync) *IntSync {
	s3 := IntSync{}
	s3.nodeMap = make(map[int]struct{})
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
func (s *IntSync) Intersection(s2 *IntSync) *IntSync {
	s3 := IntSync{}
	s3.nodeMap = make(map[int]struct{})
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
func (s *IntSync) Minus(s2 *IntSync) *IntSync {
	s3 := IntSync{}
	s3.nodeMap = make(map[int]struct{})
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
func (s *IntSync) Subset(s2 *IntSync) bool {
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
func (s *IntSync) Superset(s2 *IntSync) bool {
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
