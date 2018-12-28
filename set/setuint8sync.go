package set

import "sync"

type SetUint8Sync struct {
	nodeMap map[uint8]bool
	sync.RWMutex
}

func NewUint8Sync(nums []uint8) *SetUint8Sync {
	s := &SetUint8Sync{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

func (s *SetUint8Sync) Add(num uint8) *SetUint8Sync {
	s.Lock()
	defer s.Unlock()
	if s.nodeMap == nil {
		s.nodeMap = make(map[uint8]bool)
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = true
	}
	return s
}

func (s *SetUint8Sync) Clear() {
	s.Lock()
	defer s.Unlock()
	s.nodeMap = make(map[uint8]bool)
}

func (s *SetUint8Sync) Remove(num uint8) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

func (s *SetUint8Sync) Contains(num uint8) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.nodeMap[num]
	return ok
}

func (s *SetUint8Sync) GetList() []uint8 {
	s.RLock()
	defer s.RUnlock()
	nums := []uint8{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

func (s *SetUint8Sync) Size() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.nodeMap)
}

func (s *SetUint8Sync) Union(s2 *SetUint8Sync) *SetUint8Sync {
	s3 := SetUint8Sync{}
	s3.nodeMap = make(map[uint8]bool)
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
func (s *SetUint8Sync) Intersection(s2 *SetUint8Sync) *SetUint8Sync {
	s3 := SetUint8Sync{}
	s3.nodeMap = make(map[uint8]bool)
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
func (s *SetUint8Sync) Minus(s2 *SetUint8Sync) *SetUint8Sync {
	s3 := SetUint8Sync{}
	s3.nodeMap = make(map[uint8]bool)
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

func (s *SetUint8Sync) Subset(s2 *SetUint8Sync) bool {
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

func (s *SetUint8Sync) Superset(s2 *SetUint8Sync) bool {
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
