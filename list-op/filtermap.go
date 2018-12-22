package list_op

func FilterMapInt(fFilter func(int) bool, fMap func(int) int, list []int) []int {
	if fFilter == nil || fMap == nil {
		return []int{}
	}
	var newList []int
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

func FilterMapInt64(fFilter func(int64) bool, fMap func(int64) int64, list []int64) []int64 {
	if fFilter == nil || fMap == nil {
		return []int64{}
	}

	var newList []int64
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

func FilterMapInt32(fFilter func(int32) bool, fMap func(int32) int32, list []int32) []int32 {
	if fFilter == nil || fMap == nil {
		return []int32{}
	}
	var newList []int32
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

func FilterMapInt16(fFilter func(int16) bool, fMap func(int16) int16, list []int16) []int16 {
	if fFilter == nil || fMap == nil {
		return []int16{}
	}
	var newList []int16
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

func FilterMapInt8(fFilter func(int8) bool, fMap func(int8) int8, list []int8) []int8 {
	if fFilter == nil || fMap == nil {
		return []int8{}
	}
	var newList []int8
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

func FilterMapFloat64(fFilter func(float64) bool, fMap func(float64) float64, list []float64) []float64 {
	if fFilter == nil || fMap == nil {
		return []float64{}
	}
	var newList []float64
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

func FilterMapFloat32(fFilter func(float32) bool, fMap func(float32) float32, list []float32) []float32 {
	if fFilter == nil || fMap == nil {
		return []float32{}
	}
	var newList []float32
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

func FilterMapStr(fFilter func(string) bool, fMap func(string) string, list []string) []string {
	if fFilter == nil || fMap == nil {
		return []string{}
	}
	var newList []string
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}
