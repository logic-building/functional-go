package set

import "sync"

type SetStrSync struct {
	nodeMap map[string]bool
	sync.RWMutex
}

// Create set object
func NewStrSync(strList []string) *SetStrSync {
	s := &SetStrSync{}
	for _, str := range strList {
		s.Add(str)
	}
	return s
}

// Add an item
func (s *SetStrSync) Add(str string) *SetStrSync {
	s.Lock()
	defer s.Unlock()
	if s.nodeMap == nil {
		s.nodeMap = make(map[string]bool)
	}
	_, ok := s.nodeMap[str]
	if !ok {
		s.nodeMap[str] = true
	}
	return s
}

// Make object empty
func (s *SetStrSync) Clear() {
	s.Lock()
	defer s.Unlock()
	s.nodeMap = make(map[string]bool)
}

// Remove an item
func (s *SetStrSync) Remove(str string) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.nodeMap[str]
	if ok {
		delete(s.nodeMap, str)
	}
	return ok
}

// Check if item exists in set
func (s *SetStrSync) Contains(str string) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.nodeMap[str]
	return ok
}

// Get set items
func (s *SetStrSync) GetList() []string {
	s.RLock()
	defer s.RUnlock()
	strList := []string{}
	for i := range s.nodeMap {
		strList = append(strList, i)
	}
	return strList
}

// Get size of set
func (s *SetStrSync) Size() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.nodeMap)
}

// Returns all the items that are in S or in S2
func (s *SetStrSync) Union(s2 *SetStrSync) *SetStrSync {
	s3 := SetStrSync{}
	s3.nodeMap = make(map[string]bool)
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
func (s *SetStrSync) Intersection(s2 *SetStrSync) *SetStrSync {
	s3 := SetStrSync{}
	s3.nodeMap = make(map[string]bool)
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
func (s *SetStrSync) Minus(s2 *SetStrSync) *SetStrSync {
	s3 := SetStrSync{}
	s3.nodeMap = make(map[string]bool)
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
func (s *SetStrSync) Subset(s2 *SetStrSync) bool {
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
func (s *SetStrSync) Superset(s2 *SetStrSync) bool {
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
