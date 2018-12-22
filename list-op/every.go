package list_op

func EveryInt(f func(int) bool, list []int) bool {
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

func EveryInt64(f func(int64) bool, list []int64) bool {
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

func EveryInt32(f func(int32) bool, list []int32) bool {
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

func EveryInt16(f func(int16) bool, list []int16) bool {
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

func EveryInt8(f func(int8) bool, list []int8) bool {
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

func EveryFloat64(f func(float64) bool, list []float64) bool {
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

func EveryFloat32(f func(float32) bool, list []float32) bool {
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

func EveryStr(f func(string) bool, list []string) bool {
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}
