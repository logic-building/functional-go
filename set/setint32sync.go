package set

import "sync"

type SetInt32Sync struct {
	nodeMap map[int32]bool
	sync.RWMutex
}

func NewInt32Sync(nums []int32) *SetInt32Sync {
	s := &SetInt32Sync{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

func (s *SetInt32Sync) Add(num int32) *SetInt32Sync {
	s.Lock()
	defer s.Unlock()
	if s.nodeMap == nil {
		s.nodeMap = make(map[int32]bool)
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = true
	}
	return s
}

func (s *SetInt32Sync) Clear() {
	s.Lock()
	defer s.Unlock()
	s.nodeMap = make(map[int32]bool)
}

func (s *SetInt32Sync) Remove(num int32) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

func (s *SetInt32Sync) Contains(num int32) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.nodeMap[num]
	return ok
}

func (s *SetInt32Sync) GetList() []int32 {
	s.RLock()
	defer s.RUnlock()
	nums := []int32{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

func (s *SetInt32Sync) Size() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.nodeMap)
}

func (s *SetInt32Sync) Union(s2 *SetInt32Sync) *SetInt32Sync {
	s3 := SetInt32Sync{}
	s3.nodeMap = make(map[int32]bool)
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
func (s *SetInt32Sync) Intersection(s2 *SetInt32Sync) *SetInt32Sync {
	s3 := SetInt32Sync{}
	s3.nodeMap = make(map[int32]bool)
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
func (s *SetInt32Sync) Minus(s2 *SetInt32Sync) *SetInt32Sync {
	s3 := SetInt32Sync{}
	s3.nodeMap = make(map[int32]bool)
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

func (s *SetInt32Sync) Subset(s2 *SetInt32Sync) bool {
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

func (s *SetInt32Sync) Superset(s2 *SetInt32Sync) bool {
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
