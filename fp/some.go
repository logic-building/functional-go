package fp

func SomeInt(f func(int) bool, list []int) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}

func SomeInt64(f func(int64) bool, list []int64) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}

func SomeInt32(f func(int32) bool, list []int32) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}

func SomeInt16(f func(int16) bool, list []int16) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}

func SomeInt8(f func(int8) bool, list []int8) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}

func SomeUint64(f func(uint64) bool, list []uint64) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}

func SomeUint32(f func(uint32) bool, list []uint32) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}

func SomeUint16(f func(uint16) bool, list []uint16) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}

func SomeUint8(f func(uint8) bool, list []uint8) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}

func SomeUint(f func(uint) bool, list []uint) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}

func SomeFloat64(f func(float64) bool, list []float64) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}

func SomeFloat32(f func(float32) bool, list []float32) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}

func SomeStr(f func(string) bool, list []string) bool {
	if f == nil {
		return false
	}

	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}
