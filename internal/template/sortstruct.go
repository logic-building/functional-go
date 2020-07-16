package template

// SortStruct is template
func SortStruct() string {
	return `
type by<FFIELD_NAME><FSTRUCT_NAME> []<STRUCT_NAME>

func (a by<FFIELD_NAME><FSTRUCT_NAME>) Len() int           { return len(a) }
func (a by<FFIELD_NAME><FSTRUCT_NAME>) Less(i, j int) bool { return a[i].<FIELD_NAME> < a[j].<FIELD_NAME> }
func (a by<FFIELD_NAME><FSTRUCT_NAME>) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// Sort<FSTRUCT_NAME>By<FFIELD_NAME> sort structs
func Sort<FSTRUCT_NAME>By<FFIELD_NAME>(list []<STRUCT_NAME>) []<STRUCT_NAME> {
	if len(list) == 0 {
		return []<STRUCT_NAME>{}
	}
	newList := make([]<STRUCT_NAME>, len(list))
	for i, item := range list {
		newList[i] = item
	}
	sort.Sort(by<FFIELD_NAME><FSTRUCT_NAME>(newList))
	return newList
}

// Sort<FSTRUCT_NAME>By<FFIELD_NAME>Ptr sorts structs
func Sort<FSTRUCT_NAME>By<FFIELD_NAME>Ptr(list []*<STRUCT_NAME>) []*<STRUCT_NAME> {
	if len(list) == 0 {
		return []*<STRUCT_NAME>{}
	}
	newList := make([]<STRUCT_NAME>, len(list))
	newListPtr := make([]*<STRUCT_NAME>, len(list))

	for i, item := range list {
		newList[i] = *item
	}
	sort.Sort(by<FFIELD_NAME><FSTRUCT_NAME>(newList))

	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}
	return newListPtr
}

type by<FFIELD_NAME><FSTRUCT_NAME>Desc []<STRUCT_NAME>

func (a by<FFIELD_NAME><FSTRUCT_NAME>Desc) Len() int           { return len(a) }
func (a by<FFIELD_NAME><FSTRUCT_NAME>Desc) Less(i, j int) bool { return a[i].<FIELD_NAME> > a[j].<FIELD_NAME> }
func (a by<FFIELD_NAME><FSTRUCT_NAME>Desc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// Sort<FSTRUCT_NAME>By<FFIELD_NAME>Desc sorts structs
func Sort<FSTRUCT_NAME>By<FFIELD_NAME>Desc(list []<STRUCT_NAME>) []<STRUCT_NAME> {
	if len(list) == 0 {
		return []<STRUCT_NAME>{}
	}
	newList := make([]<STRUCT_NAME>, len(list))
	for i, item := range list {
		newList[i] = item
	}
	sort.Sort(by<FFIELD_NAME><FSTRUCT_NAME>Desc(newList))
	return newList
}

// Sort<FSTRUCT_NAME>By<FFIELD_NAME>DescPtr sorts structs
func Sort<FSTRUCT_NAME>By<FFIELD_NAME>DescPtr(list []*<STRUCT_NAME>) []*<STRUCT_NAME> {
	if len(list) == 0 {
		return []*<STRUCT_NAME>{}
	}
	newList := make([]<STRUCT_NAME>, len(list))
	newListPtr := make([]*<STRUCT_NAME>, len(list))

	for i, item := range list {
		newList[i] = *item
	}
	sort.Sort(by<FFIELD_NAME><FSTRUCT_NAME>Desc(newList))

	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}
	return newListPtr
}`
}

// SortStructForTimeField is template
func SortStructForTimeField() string {
	return `
type by<FFIELD_NAME><FSTRUCT_NAME> []<STRUCT_NAME>

func (a by<FFIELD_NAME><FSTRUCT_NAME>) Len() int           { return len(a) }
func (a by<FFIELD_NAME><FSTRUCT_NAME>) Less(i, j int) bool { return a[i].<FIELD_NAME>.Before(a[j].<FIELD_NAME>) }
func (a by<FFIELD_NAME><FSTRUCT_NAME>) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// Sort<FSTRUCT_NAME>By<FFIELD_NAME> sort structs
func Sort<FSTRUCT_NAME>By<FFIELD_NAME>(list []<STRUCT_NAME>) []<STRUCT_NAME> {
	if len(list) == 0 {
		return []<STRUCT_NAME>{}
	}
	newList := make([]<STRUCT_NAME>, len(list))
	for i, item := range list {
		newList[i] = item
	}
	sort.Sort(by<FFIELD_NAME><FSTRUCT_NAME>(newList))
	return newList
}

// Sort<FSTRUCT_NAME>By<FFIELD_NAME>Ptr sorts structs
func Sort<FSTRUCT_NAME>By<FFIELD_NAME>Ptr(list []*<STRUCT_NAME>) []*<STRUCT_NAME> {
	if len(list) == 0 {
		return []*<STRUCT_NAME>{}
	}
	newList := make([]<STRUCT_NAME>, len(list))
	newListPtr := make([]*<STRUCT_NAME>, len(list))

	for i, item := range list {
		newList[i] = *item
	}
	sort.Sort(by<FFIELD_NAME><FSTRUCT_NAME>(newList))

	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}
	return newListPtr
}

type by<FFIELD_NAME><FSTRUCT_NAME>Desc []<STRUCT_NAME>

func (a by<FFIELD_NAME><FSTRUCT_NAME>Desc) Len() int           { return len(a) }
func (a by<FFIELD_NAME><FSTRUCT_NAME>Desc) Less(i, j int) bool { return a[i].<FIELD_NAME>.After(a[j].<FIELD_NAME>) }
func (a by<FFIELD_NAME><FSTRUCT_NAME>Desc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// Sort<FSTRUCT_NAME>By<FFIELD_NAME>Desc sorts structs
func Sort<FSTRUCT_NAME>By<FFIELD_NAME>Desc(list []<STRUCT_NAME>) []<STRUCT_NAME> {
	if len(list) == 0 {
		return []<STRUCT_NAME>{}
	}
	newList := make([]<STRUCT_NAME>, len(list))
	for i, item := range list {
		newList[i] = item
	}
	sort.Sort(by<FFIELD_NAME><FSTRUCT_NAME>Desc(newList))
	return newList
}

// Sort<FSTRUCT_NAME>By<FFIELD_NAME>DescPtr sorts structs
func Sort<FSTRUCT_NAME>By<FFIELD_NAME>DescPtr(list []*<STRUCT_NAME>) []*<STRUCT_NAME> {
	if len(list) == 0 {
		return []*<STRUCT_NAME>{}
	}
	newList := make([]<STRUCT_NAME>, len(list))
	newListPtr := make([]*<STRUCT_NAME>, len(list))

	for i, item := range list {
		newList[i] = *item
	}
	sort.Sort(by<FFIELD_NAME><FSTRUCT_NAME>Desc(newList))

	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}
	return newListPtr
}`
}

// SortStructForTimeFieldPtr is template
func SortStructForTimeFieldPtr() string {
	return `
type by<FFIELD_NAME><FSTRUCT_NAME> []<STRUCT_NAME>

func (a by<FFIELD_NAME><FSTRUCT_NAME>) Len() int           { return len(a) }
func (a by<FFIELD_NAME><FSTRUCT_NAME>) Less(i, j int) bool {
	v1Ptr := a[i].<FIELD_NAME>
	v2Ptr := a[j].<FIELD_NAME>

	v1 := *v1Ptr
	v2 := *v2Ptr

	return v1.Before(v2)
}
func (a by<FFIELD_NAME><FSTRUCT_NAME>) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// Sort<FSTRUCT_NAME>By<FFIELD_NAME> sort structs
func Sort<FSTRUCT_NAME>By<FFIELD_NAME>(list []<STRUCT_NAME>) []<STRUCT_NAME> {
	if len(list) == 0 {
		return []<STRUCT_NAME>{}
	}
	newList := make([]<STRUCT_NAME>, len(list))
	for i, item := range list {
		newList[i] = item
	}
	sort.Sort(by<FFIELD_NAME><FSTRUCT_NAME>(newList))
	return newList
}

// Sort<FSTRUCT_NAME>By<FFIELD_NAME>Ptr sorts structs
func Sort<FSTRUCT_NAME>By<FFIELD_NAME>Ptr(list []*<STRUCT_NAME>) []*<STRUCT_NAME> {
	if len(list) == 0 {
		return []*<STRUCT_NAME>{}
	}
	newList := make([]<STRUCT_NAME>, len(list))
	newListPtr := make([]*<STRUCT_NAME>, len(list))

	for i, item := range list {
		newList[i] = *item
	}
	sort.Sort(by<FFIELD_NAME><FSTRUCT_NAME>(newList))

	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}
	return newListPtr
}

type by<FFIELD_NAME><FSTRUCT_NAME>Desc []<STRUCT_NAME>

func (a by<FFIELD_NAME><FSTRUCT_NAME>Desc) Len() int           { return len(a) }
func (a by<FFIELD_NAME><FSTRUCT_NAME>Desc) Less(i, j int) bool {
	v1Ptr := a[i].<FIELD_NAME>
	v2Ptr := a[j].<FIELD_NAME>

	v1 := *v1Ptr
	v2 := *v2Ptr

	return v1.After(v2) 
}
func (a by<FFIELD_NAME><FSTRUCT_NAME>Desc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// Sort<FSTRUCT_NAME>By<FFIELD_NAME>Desc sorts structs
func Sort<FSTRUCT_NAME>By<FFIELD_NAME>Desc(list []<STRUCT_NAME>) []<STRUCT_NAME> {
	if len(list) == 0 {
		return []<STRUCT_NAME>{}
	}
	newList := make([]<STRUCT_NAME>, len(list))
	for i, item := range list {
		newList[i] = item
	}
	sort.Sort(by<FFIELD_NAME><FSTRUCT_NAME>Desc(newList))
	return newList
}

// Sort<FSTRUCT_NAME>By<FFIELD_NAME>DescPtr sorts structs
func Sort<FSTRUCT_NAME>By<FFIELD_NAME>DescPtr(list []*<STRUCT_NAME>) []*<STRUCT_NAME> {
	if len(list) == 0 {
		return []*<STRUCT_NAME>{}
	}
	newList := make([]<STRUCT_NAME>, len(list))
	newListPtr := make([]*<STRUCT_NAME>, len(list))

	for i, item := range list {
		newList[i] = *item
	}
	sort.Sort(by<FFIELD_NAME><FSTRUCT_NAME>Desc(newList))

	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}
	return newListPtr
}`
}

// SortStructPtr is template
func SortStructPtr() string {
	return `
type by<FFIELD_NAME><FSTRUCT_NAME> []<STRUCT_NAME>

func (a by<FFIELD_NAME><FSTRUCT_NAME>) Len() int           { return len(a) }
func (a by<FFIELD_NAME><FSTRUCT_NAME>) Less(i, j int) bool {
	v1 := a[i].<FIELD_NAME>
	v2 := a[j].<FIELD_NAME>
	return *v1 < *v2
}
func (a by<FFIELD_NAME><FSTRUCT_NAME>) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// Sort<FSTRUCT_NAME>By<FFIELD_NAME> sort structs
func Sort<FSTRUCT_NAME>By<FFIELD_NAME>(list []<STRUCT_NAME>) []<STRUCT_NAME> {
	if len(list) == 0 {
		return []<STRUCT_NAME>{}
	}
	newList := make([]<STRUCT_NAME>, len(list))
	for i, item := range list {
		newList[i] = item
	}
	sort.Sort(by<FFIELD_NAME><FSTRUCT_NAME>(newList))
	return newList
}

// Sort<FSTRUCT_NAME>By<FFIELD_NAME>Ptr sorts structs
func Sort<FSTRUCT_NAME>By<FFIELD_NAME>Ptr(list []*<STRUCT_NAME>) []*<STRUCT_NAME> {
	if len(list) == 0 {
		return []*<STRUCT_NAME>{}
	}
	newList := make([]<STRUCT_NAME>, len(list))
	newListPtr := make([]*<STRUCT_NAME>, len(list))

	for i, item := range list {
		newList[i] = *item
	}
	sort.Sort(by<FFIELD_NAME><FSTRUCT_NAME>(newList))

	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}
	return newListPtr
}

type by<FFIELD_NAME><FSTRUCT_NAME>Desc []<STRUCT_NAME>

func (a by<FFIELD_NAME><FSTRUCT_NAME>Desc) Len() int           { return len(a) }
func (a by<FFIELD_NAME><FSTRUCT_NAME>Desc) Less(i, j int) bool {
	v1 := a[i].<FIELD_NAME>
	v2 := a[j].<FIELD_NAME>

	return *v1 > *v2
}
func (a by<FFIELD_NAME><FSTRUCT_NAME>Desc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// Sort<FSTRUCT_NAME>By<FFIELD_NAME>Desc sorts structs
func Sort<FSTRUCT_NAME>By<FFIELD_NAME>Desc(list []<STRUCT_NAME>) []<STRUCT_NAME> {
	if len(list) == 0 {
		return []<STRUCT_NAME>{}
	}
	newList := make([]<STRUCT_NAME>, len(list))
	for i, item := range list {
		newList[i] = item
	}
	sort.Sort(by<FFIELD_NAME><FSTRUCT_NAME>Desc(newList))
	return newList
}

// Sort<FSTRUCT_NAME>By<FFIELD_NAME>DescPtr sorts structs
func Sort<FSTRUCT_NAME>By<FFIELD_NAME>DescPtr(list []*<STRUCT_NAME>) []*<STRUCT_NAME> {
	if len(list) == 0 {
		return []*<STRUCT_NAME>{}
	}
	newList := make([]<STRUCT_NAME>, len(list))
	newListPtr := make([]*<STRUCT_NAME>, len(list))

	for i, item := range list {
		newList[i] = *item
	}
	sort.Sort(by<FFIELD_NAME><FSTRUCT_NAME>Desc(newList))

	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}
	return newListPtr
}`
}
