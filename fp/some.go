package fp

import "strings"

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

func SomeUint64(num uint64, list []uint64) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

func SomeUint32(num uint32, list []uint32) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

func SomeUint16(num uint16, list []uint16) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

func SomeUint8(num uint8, list []uint8) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

func SomeUint(num uint, list []uint) bool {
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

func SomeStrIgnoreCase(str string, list []string) bool {
	strLowerCase := strings.ToLower(str)
	for _, v := range list {
		if strings.ToLower(v) == strLowerCase {
			return true
		}
	}
	return false
}
