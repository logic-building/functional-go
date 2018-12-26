package set

import "sync"

type SetUintSync struct {
	nodeMap map[uint]bool
	sync.RWMutex
}

func NewSetUintSync(nums []uint) *SetUintSync {
	s := &SetUintSync{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

func (s *SetUintSync) Add(num uint) *SetUintSync {
	s.Lock()
	defer s.Unlock()
	if s.nodeMap == nil {
		s.nodeMap = make(map[uint]bool)
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = true
	}
	return s
}

func (s *SetUintSync) Clear() {
	s.Lock()
	defer s.Unlock()
	s.nodeMap = make(map[uint]bool)
}

func (s *SetUintSync) Remove(num uint) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

func (s *SetUintSync) Contains(num uint) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.nodeMap[num]
	return ok
}

func (s *SetUintSync) GetList() []uint {
	s.RLock()
	defer s.RUnlock()
	nums := []uint{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

func (s *SetUintSync) Size() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.nodeMap)
}

func (s *SetUintSync) Union(s2 *SetUintSync) *SetUintSync {
	s3 := SetUintSync{}
	s3.nodeMap = make(map[uint]bool)
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
func (s *SetUintSync) Intersection(s2 *SetUintSync) *SetUintSync {
	s3 := SetUintSync{}
	s3.nodeMap = make(map[uint]bool)
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
func (s *SetUintSync) Minus(s2 *SetUintSync) *SetUintSync {
	s3 := SetUintSync{}
	s3.nodeMap = make(map[uint]bool)
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

func (s *SetUintSync) Subset(s2 *SetUintSync) bool {
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

func (s *SetUintSync) Superset(s2 *SetUintSync) bool {
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
