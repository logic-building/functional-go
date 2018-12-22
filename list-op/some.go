package list_op

func SomeInt(num int, list []int) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

func SomeInt64(num int64, list []int64) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

func SomeInt32(num int32, list []int32) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

func SomeInt16(num int16, list []int16) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

func SomeInt8(num int8, list []int8) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

func SomeFloat64(num float64, list []float64) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

func SomeFloat32(num float32, list []float32) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

func SomeStr(num string, list []string) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}
