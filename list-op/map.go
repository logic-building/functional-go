package list_op

func MapInt(f func(int) int, list []int) []int {
	newList := make([]int, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt64(f func(int64) int64, list []int64) []int64 {
	newList := make([]int64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt32(f func(int32) int32, list []int32) []int32 {
	newList := make([]int32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt8(f func(int8) int8, list []int8) []int8 {
	newList := make([]int8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt16(f func(int16) int16, list []int16) []int16 {
	newList := make([]int16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapFloat64(f func(float64) float64, list []float64) []float64 {
	newList := make([]float64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapFloat32(f func(float32) float32, list []float32) []float32 {
	newList := make([]float32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapStr(f func(string) string, list []string) []string {
	newList := make([]string, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}
