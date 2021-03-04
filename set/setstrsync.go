package set

import "sync"

// StrSync - struct
type StrSync struct {
	nodeMap map[string]struct{}
	sync.RWMutex
}

// NewStrSync creates set
func NewStrSync(strList []string) *StrSync {
	s := &StrSync{}
	for _, str := range strList {
		s.Add(str)
	}
	return s
}

// Add an item
func (s *StrSync) Add(str string) *StrSync {
	s.Lock()
	defer s.Unlock()
	if s.nodeMap == nil {
		s.nodeMap = make(map[string]struct{})
	}
	_, ok := s.nodeMap[str]
	if !ok {
		s.nodeMap[str] = struct{}{}
	}
	return s
}

// Clear set
func (s *StrSync) Clear() {
	s.Lock()
	defer s.Unlock()
	s.nodeMap = make(map[string]struct{})
}

// Remove an item
func (s *StrSync) Remove(str string) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.nodeMap[str]
	if ok {
		delete(s.nodeMap, str)
	}
	return ok
}

// Contains - Check if item exists in set
func (s *StrSync) Contains(str string) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.nodeMap[str]
	return ok
}

// GetList - Get set items
func (s *StrSync) GetList() []string {
	s.RLock()
	defer s.RUnlock()
	strList := []string{}
	for i := range s.nodeMap {
		strList = append(strList, i)
	}
	return strList
}

// Size - Get size of set
func (s *StrSync) Size() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.nodeMap)
}

// Union returns all the items that are in S or in S2
func (s *StrSync) Union(s2 *StrSync) *StrSync {
	s3 := StrSync{}
	s3.nodeMap = make(map[string]struct{})
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
func (s *StrSync) Intersection(s2 *StrSync) *StrSync {
	s3 := StrSync{}
	s3.nodeMap = make(map[string]struct{})
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
func (s *StrSync) Minus(s2 *StrSync) *StrSync {
	s3 := StrSync{}
	s3.nodeMap = make(map[string]struct{})
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
func (s *StrSync) Subset(s2 *StrSync) bool {
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
func (s *StrSync) Superset(s2 *StrSync) bool {
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
