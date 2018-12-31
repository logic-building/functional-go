package set

import "sync"

type SetUint32Sync struct {
	nodeMap map[uint32]bool
	sync.RWMutex
}

func NewUint32Sync(nums []uint32) *SetUint32Sync {
	s := &SetUint32Sync{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

func (s *SetUint32Sync) Add(num uint32) *SetUint32Sync {
	s.Lock()
	defer s.Unlock()
	if s.nodeMap == nil {
		s.nodeMap = make(map[uint32]bool)
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = true
	}
	return s
}

func (s *SetUint32Sync) Clear() {
	s.Lock()
	defer s.Unlock()
	s.nodeMap = make(map[uint32]bool)
}

func (s *SetUint32Sync) Remove(num uint32) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

func (s *SetUint32Sync) Contains(num uint32) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.nodeMap[num]
	return ok
}

func (s *SetUint32Sync) GetList() []uint32 {
	s.RLock()
	defer s.RUnlock()
	nums := []uint32{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

func (s *SetUint32Sync) Size() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.nodeMap)
}

func (s *SetUint32Sync) Union(s2 *SetUint32Sync) *SetUint32Sync {
	s3 := SetUint32Sync{}
	s3.nodeMap = make(map[uint32]bool)
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
func (s *SetUint32Sync) Intersection(s2 *SetUint32Sync) *SetUint32Sync {
	s3 := SetUint32Sync{}
	s3.nodeMap = make(map[uint32]bool)
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
func (s *SetUint32Sync) Minus(s2 *SetUint32Sync) *SetUint32Sync {
	s3 := SetUint32Sync{}
	s3.nodeMap = make(map[uint32]bool)
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

func (s *SetUint32Sync) Subset(s2 *SetUint32Sync) bool {
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

func (s *SetUint32Sync) Superset(s2 *SetUint32Sync) bool {
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
