package set

import "sync"

type SetFloat64Sync struct {
	nodeMap map[float64]bool
	sync.RWMutex
}

func NewFloat64Sync(nums []float64) *SetFloat64Sync {
	s := &SetFloat64Sync{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

func (s *SetFloat64Sync) Add(num float64) *SetFloat64Sync {
	s.Lock()
	defer s.Unlock()
	if s.nodeMap == nil {
		s.nodeMap = make(map[float64]bool)
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = true
	}
	return s
}

func (s *SetFloat64Sync) Clear() {
	s.Lock()
	defer s.Unlock()
	s.nodeMap = make(map[float64]bool)
}

func (s *SetFloat64Sync) Remove(num float64) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

func (s *SetFloat64Sync) Contains(num float64) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.nodeMap[num]
	return ok
}

func (s *SetFloat64Sync) GetList() []float64 {
	s.RLock()
	defer s.RUnlock()
	nums := []float64{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

func (s *SetFloat64Sync) Size() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.nodeMap)
}

func (s *SetFloat64Sync) Union(s2 *SetFloat64Sync) *SetFloat64Sync {
	s3 := SetFloat64Sync{}
	s3.nodeMap = make(map[float64]bool)
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
func (s *SetFloat64Sync) Intersection(s2 *SetFloat64Sync) *SetFloat64Sync {
	s3 := SetFloat64Sync{}
	s3.nodeMap = make(map[float64]bool)
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
func (s *SetFloat64Sync) Minus(s2 *SetFloat64Sync) *SetFloat64Sync {
	s3 := SetFloat64Sync{}
	s3.nodeMap = make(map[float64]bool)
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

func (s *SetFloat64Sync) Subset(s2 *SetFloat64Sync) bool {
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

func (s *SetFloat64Sync) Superset(s2 *SetFloat64Sync) bool {
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
