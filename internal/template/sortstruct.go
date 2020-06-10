package template

// SortStruct
func SortStruct() string {
	return `
type by<FFIELD_NAME> []<STRUCT_NAME>

func (a by<FFIELD_NAME>) Len() int           { return len(a) }
func (a by<FFIELD_NAME>) Less(i, j int) bool { return a[i].<FIELD_NAME> < a[j].<FIELD_NAME> }
func (a by<FFIELD_NAME>) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func Sort<FSTRUCT_NAME>By<FFIELD_NAME>(list []<STRUCT_NAME>) []<STRUCT_NAME> {
	if len(list) == 0 {
		return []<STRUCT_NAME>{}
	}
	newList := make([]<STRUCT_NAME>, len(list))
	for i, item := range list {
		newList[i] = item
	}
	sort.Sort(by<FFIELD_NAME>(newList))
	return newList
}

func Sort<FSTRUCT_NAME>By<FFIELD_NAME>Ptr(list []*<STRUCT_NAME>) []*<STRUCT_NAME> {
	if len(list) == 0 {
		return []*<STRUCT_NAME>{}
	}
	newList := make([]<STRUCT_NAME>, len(list))
	newListPtr := make([]*<STRUCT_NAME>, len(list))

	for i, item := range list {
		newList[i] = *item
	}
	sort.Sort(by<FFIELD_NAME>(newList))

	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}
	return newListPtr
}

type by<FFIELD_NAME>Desc []<STRUCT_NAME>

func (a by<FFIELD_NAME>Desc) Len() int           { return len(a) }
func (a by<FFIELD_NAME>Desc) Less(i, j int) bool { return a[i].<FIELD_NAME> > a[j].<FIELD_NAME> }
func (a by<FFIELD_NAME>Desc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func Sort<FSTRUCT_NAME>By<FFIELD_NAME>Desc(list []<STRUCT_NAME>) []<STRUCT_NAME> {
	if len(list) == 0 {
		return []<STRUCT_NAME>{}
	}
	newList := make([]<STRUCT_NAME>, len(list))
	for i, item := range list {
		newList[i] = item
	}
	sort.Sort(by<FFIELD_NAME>Desc(newList))
	return newList
}

func Sort<FSTRUCT_NAME>By<FFIELD_NAME>DescPtr(list []*<STRUCT_NAME>) []*<STRUCT_NAME> {
	if len(list) == 0 {
		return []*<STRUCT_NAME>{}
	}
	newList := make([]<STRUCT_NAME>, len(list))
	newListPtr := make([]*<STRUCT_NAME>, len(list))

	for i, item := range list {
		newList[i] = *item
	}
	sort.Sort(by<FFIELD_NAME>Desc(newList))

	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}
	return newListPtr
}`
}
