package set

// Str - struct
type Str struct {
	nodeMap map[string]struct{}
}

// NewStr creates set
func NewStr(strList []string) *Str {
	s := &Str{}
	for _, str := range strList {
		s.Add(str)
	}
	return s
}

// Add an item
func (s *Str) Add(str string) *Str {
	if s.nodeMap == nil {
		s.nodeMap = make(map[string]struct{})
	}
	_, ok := s.nodeMap[str]
	if !ok {
		s.nodeMap[str] = struct{}{}
	}
	return s
}

// Clear set
func (s *Str) Clear() {
	s.nodeMap = make(map[string]struct{})
}

// Remove an item
func (s *Str) Remove(str string) bool {
	_, ok := s.nodeMap[str]
	if ok {
		delete(s.nodeMap, str)
	}
	return ok
}

// Contains - Check if item exists in set
func (s *Str) Contains(str string) bool {
	_, ok := s.nodeMap[str]
	return ok
}

// GetList - Get set items
func (s *Str) GetList() []string {
	strList := []string{}
	for i := range s.nodeMap {
		strList = append(strList, i)
	}
	return strList
}

// Size - Get size of set
func (s *Str) Size() int {
	return len(s.nodeMap)
}

// Union returns all the items that are in S or in S2
func (s *Str) Union(s2 *Str) *Str {
	s3 := Str{}
	s3.nodeMap = make(map[string]struct{})
	for i := range s.nodeMap {
		s3.nodeMap[i] = struct{}{}
	}
	for i := range s2.nodeMap {
		_, ok := s3.nodeMap[i]
		if !ok {
			s3.nodeMap[i] = struct{}{}
		}
	}
	return &s3
}

// Intersection returns common items in S and S2
func (s *Str) Intersection(s2 *Str) *Str {
	s3 := Str{}
	s3.nodeMap = make(map[string]struct{})
	for i := range s2.nodeMap {
		_, ok := s.nodeMap[i]
		if ok {
			s3.nodeMap[i] = struct{}{}
		}
	}
	return &s3
}

// Minus - s.Minus(s2) : all of S but not in S2
func (s *Str) Minus(s2 *Str) *Str {
	s3 := Str{}
	s3.nodeMap = make(map[string]struct{})
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			s3.nodeMap[i] = struct{}{}
		}
	}
	return &s3
}

// Subset checks if S is subset of S2
func (s *Str) Subset(s2 *Str) bool {
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			return false
		}
	}
	return true
}

// Superset checks if S is superset of S2
func (s *Str) Superset(s2 *Str) bool {
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
