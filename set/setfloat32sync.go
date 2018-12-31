package set

import "sync"

type SetFloat32Sync struct {
	nodeMap map[float32]bool
	sync.RWMutex
}

func NewFloat32Sync(nums []float32) *SetFloat32Sync {
	s := &SetFloat32Sync{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

func (s *SetFloat32Sync) Add(num float32) *SetFloat32Sync {
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

func (s *SetFloat32Sync) Clear() {
	s.Lock()
	defer s.Unlock()
	s.nodeMap = make(map[float32]bool)
}

func (s *SetFloat32Sync) Remove(num float32) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

func (s *SetFloat32Sync) Contains(num float32) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.nodeMap[num]
	return ok
}

func (s *SetFloat32Sync) GetList() []float32 {
	s.RLock()
	defer s.RUnlock()
	nums := []float32{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

func (s *SetFloat32Sync) Size() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.nodeMap)
}

func (s *SetFloat32Sync) Union(s2 *SetFloat32Sync) *SetFloat32Sync {
	s3 := SetFloat32Sync{}
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

// common in A and B
func (s *SetFloat32Sync) Intersection(s2 *SetFloat32Sync) *SetFloat32Sync {
	s3 := SetFloat32Sync{}
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

// a.Minus(b) : all of a but not in b
func (s *SetFloat32Sync) Minus(s2 *SetFloat32Sync) *SetFloat32Sync {
	s3 := SetFloat32Sync{}
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

func (s *SetFloat32Sync) Subset(s2 *SetFloat32Sync) bool {
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

func (s *SetFloat32Sync) Superset(s2 *SetFloat32Sync) bool {
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
