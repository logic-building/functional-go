package set

type SetStr struct {
	nodeMap map[string]bool
}

func NewSetStr(strList []string) *SetStr {
	s := &SetStr{}
	for _, str := range strList {
		s.Add(str)
	}
	return s
}

func (s *SetStr) Add(str string) *SetStr {
	if s.nodeMap == nil {
		s.nodeMap = make(map[string]bool)
	}
	_, ok := s.nodeMap[str]
	if !ok {
		s.nodeMap[str] = true
	}
	return s
}

func (s *SetStr) Clear() {
	s.nodeMap = make(map[string]bool)
}

func (s *SetStr) Remove(str string) bool {
	_, ok := s.nodeMap[str]
	if ok {
		delete(s.nodeMap, str)
	}
	return ok
}

func (s *SetStr) Contains(str string) bool {
	_, ok := s.nodeMap[str]
	return ok
}

func (s *SetStr) GetList() []string {
	strList := []string{}
	for i := range s.nodeMap {
		strList = append(strList, i)
	}
	return strList
}

func (s *SetStr) Size() int {
	return len(s.nodeMap)
}

func (s *SetStr) Union(s2 *SetStr) *SetStr {
	s3 := SetStr{}
	s3.nodeMap = make(map[string]bool)
	for i := range s.nodeMap {
		s3.nodeMap[i] = true
	}
	for i := range s2.nodeMap {
		_, ok := s3.nodeMap[i]
		if !ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

// common in A and B
func (s *SetStr) Intersection(s2 *SetStr) *SetStr {
	s3 := SetStr{}
	s3.nodeMap = make(map[string]bool)
	for i := range s2.nodeMap {
		_, ok := s.nodeMap[i]
		if ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

// a.Minus(b) : all of a but not in b
func (s *SetStr) Minus(s2 *SetStr) *SetStr {
	s3 := SetStr{}
	s3.nodeMap = make(map[string]bool)
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

func (s *SetStr) Subset(s2 *SetStr) bool {
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			return false
		}
	}
	return true
}

func (s *SetStr) Superset(s2 *SetStr) bool {
	for i := range s2.nodeMap {
		_, ok := s.nodeMap[i]
		if !ok {
			return false
		}
	}
	size1 := s.Size()
	size2 := s2.Size()
	return size1 == size2 || size1 > size2
}
