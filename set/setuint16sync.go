package set

import "sync"

type SetUint16Sync struct {
	nodeMap map[uint16]bool
	sync.RWMutex
}

func NewUint16Sync(nums []uint16) *SetUint16Sync {
	s := &SetUint16Sync{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

func (s *SetUint16Sync) Add(num uint16) *SetUint16Sync {
	s.Lock()
	defer s.Unlock()
	if s.nodeMap == nil {
		s.nodeMap = make(map[uint16]bool)
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = true
	}
	return s
}

func (s *SetUint16Sync) Clear() {
	s.Lock()
	defer s.Unlock()
	s.nodeMap = make(map[uint16]bool)
}

func (s *SetUint16Sync) Remove(num uint16) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

func (s *SetUint16Sync) Contains(num uint16) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.nodeMap[num]
	return ok
}

func (s *SetUint16Sync) GetList() []uint16 {
	s.RLock()
	defer s.RUnlock()
	nums := []uint16{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

func (s *SetUint16Sync) Size() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.nodeMap)
}

func (s *SetUint16Sync) Union(s2 *SetUint16Sync) *SetUint16Sync {
	s3 := SetUint16Sync{}
	s3.nodeMap = make(map[uint16]bool)
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
func (s *SetUint16Sync) Intersection(s2 *SetUint16Sync) *SetUint16Sync {
	s3 := SetUint16Sync{}
	s3.nodeMap = make(map[uint16]bool)
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
func (s *SetUint16Sync) Minus(s2 *SetUint16Sync) *SetUint16Sync {
	s3 := SetUint16Sync{}
	s3.nodeMap = make(map[uint16]bool)
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

func (s *SetUint16Sync) Subset(s2 *SetUint16Sync) bool {
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

func (s *SetUint16Sync) Superset(s2 *SetUint16Sync) bool {
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
