package set

import "sync"

type SetInt16Sync struct {
	nodeMap map[int16]bool
	sync.RWMutex
}

func NewInt16Sync(nums []int16) *SetInt16Sync {
	s := &SetInt16Sync{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

func (s *SetInt16Sync) Add(num int16) *SetInt16Sync {
	s.Lock()
	defer s.Unlock()
	if s.nodeMap == nil {
		s.nodeMap = make(map[int16]bool)
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = true
	}
	return s
}

func (s *SetInt16Sync) Clear() {
	s.Lock()
	defer s.Unlock()
	s.nodeMap = make(map[int16]bool)
}

func (s *SetInt16Sync) Remove(num int16) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

func (s *SetInt16Sync) Contains(num int16) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.nodeMap[num]
	return ok
}

func (s *SetInt16Sync) GetList() []int16 {
	s.RLock()
	defer s.RUnlock()
	nums := []int16{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

func (s *SetInt16Sync) Size() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.nodeMap)
}

func (s *SetInt16Sync) Union(s2 *SetInt16Sync) *SetInt16Sync {
	s3 := SetInt16Sync{}
	s3.nodeMap = make(map[int16]bool)
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

// common in A and B
func (s *SetInt16Sync) Intersection(s2 *SetInt16Sync) *SetInt16Sync {
	s3 := SetInt16Sync{}
	s3.nodeMap = make(map[int16]bool)
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

// a.Minus(b) : all of a but not in b
func (s *SetInt16Sync) Minus(s2 *SetInt16Sync) *SetInt16Sync {
	s3 := SetInt16Sync{}
	s3.nodeMap = make(map[int16]bool)
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

func (s *SetInt16Sync) Subset(s2 *SetInt16Sync) bool {
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

func (s *SetInt16Sync) Superset(s2 *SetInt16Sync) bool {
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
