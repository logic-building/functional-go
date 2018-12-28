package set

import "sync"

type SetUint64Sync struct {
	nodeMap map[uint64]bool
	sync.RWMutex
}

func NewUint64Sync(nums []uint64) *SetUint64Sync {
	s := &SetUint64Sync{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

func (s *SetUint64Sync) Add(num uint64) *SetUint64Sync {
	s.Lock()
	defer s.Unlock()
	if s.nodeMap == nil {
		s.nodeMap = make(map[uint64]bool)
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = true
	}
	return s
}

func (s *SetUint64Sync) Clear() {
	s.Lock()
	defer s.Unlock()
	s.nodeMap = make(map[uint64]bool)
}

func (s *SetUint64Sync) Remove(num uint64) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

func (s *SetUint64Sync) Contains(num uint64) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.nodeMap[num]
	return ok
}

func (s *SetUint64Sync) GetList() []uint64 {
	s.RLock()
	defer s.RUnlock()
	nums := []uint64{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

func (s *SetUint64Sync) Size() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.nodeMap)
}

func (s *SetUint64Sync) Union(s2 *SetUint64Sync) *SetUint64Sync {
	s3 := SetUint64Sync{}
	s3.nodeMap = make(map[uint64]bool)
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
func (s *SetUint64Sync) Intersection(s2 *SetUint64Sync) *SetUint64Sync {
	s3 := SetUint64Sync{}
	s3.nodeMap = make(map[uint64]bool)
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
func (s *SetUint64Sync) Minus(s2 *SetUint64Sync) *SetUint64Sync {
	s3 := SetUint64Sync{}
	s3.nodeMap = make(map[uint64]bool)
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

func (s *SetUint64Sync) Subset(s2 *SetUint64Sync) bool {
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

func (s *SetUint64Sync) Superset(s2 *SetUint64Sync) bool {
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
