package set

import "sync"

type SetIntSync struct {
	nodeMap map[int]bool
	sync.RWMutex
}

func NewSetIntSync(nums []int) *SetIntSync {
	s := &SetIntSync{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

func (s *SetIntSync) Add(num int) *SetIntSync {
	s.Lock()
	defer s.Unlock()
	if s.nodeMap == nil {
		s.nodeMap = make(map[int]bool)
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = true
	}
	return s
}

func (s *SetIntSync) Clear() {
	s.Lock()
	defer s.Unlock()
	s.nodeMap = make(map[int]bool)
}

func (s *SetIntSync) Remove(num int) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

func (s *SetIntSync) Contains(num int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.nodeMap[num]
	return ok
}

func (s *SetIntSync) getList() []int {
	s.RLock()
	defer s.RUnlock()
	nums := []int{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

func (s *SetIntSync) Size() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.nodeMap)
}

func (s *SetIntSync) Union(s2 *SetIntSync) *SetIntSync {
	s3 := SetIntSync{}
	s3.nodeMap = make(map[int]bool)
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
func (s *SetIntSync) Intersection(s2 *SetIntSync) *SetIntSync {
	s3 := SetIntSync{}
	s3.nodeMap = make(map[int]bool)
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
func (s *SetIntSync) Minus(s2 *SetIntSync) *SetIntSync {
	s3 := SetIntSync{}
	s3.nodeMap = make(map[int]bool)
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

func (s *SetIntSync) Subset(s2 *SetIntSync) bool {
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

func (s *SetIntSync) Superset(s2 *SetIntSync) bool {
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
