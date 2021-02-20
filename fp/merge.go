package fp

// MergeInt takes two inputs: map[int]int and map[int]int and merge two maps and returns a new map[int]int.
func MergeInt(map1, map2 map[int]int) map[int]int {
	if map1 == nil && map2 == nil {
		return map[int]int{}
	}

	newMap := make(map[int]int)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeIntInt64 takes two inputs: map[int]int64 and map[int]int64 and merge two maps and returns a new map[int]int64.
func MergeIntInt64(map1, map2 map[int]int64) map[int]int64 {
	if map1 == nil && map2 == nil {
		return map[int]int64{}
	}

	newMap := make(map[int]int64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeIntInt32 takes two inputs: map[int]int32 and map[int]int32 and merge two maps and returns a new map[int]int32.
func MergeIntInt32(map1, map2 map[int]int32) map[int]int32 {
	if map1 == nil && map2 == nil {
		return map[int]int32{}
	}

	newMap := make(map[int]int32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeIntInt16 takes two inputs: map[int]int16 and map[int]int16 and merge two maps and returns a new map[int]int16.
func MergeIntInt16(map1, map2 map[int]int16) map[int]int16 {
	if map1 == nil && map2 == nil {
		return map[int]int16{}
	}

	newMap := make(map[int]int16)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeIntInt8 takes two inputs: map[int]int8 and map[int]int8 and merge two maps and returns a new map[int]int8.
func MergeIntInt8(map1, map2 map[int]int8) map[int]int8 {
	if map1 == nil && map2 == nil {
		return map[int]int8{}
	}

	newMap := make(map[int]int8)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeIntUint takes two inputs: map[int]uint and map[int]uint and merge two maps and returns a new map[int]uint.
func MergeIntUint(map1, map2 map[int]uint) map[int]uint {
	if map1 == nil && map2 == nil {
		return map[int]uint{}
	}

	newMap := make(map[int]uint)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeIntUint64 takes two inputs: map[int]uint64 and map[int]uint64 and merge two maps and returns a new map[int]uint64.
func MergeIntUint64(map1, map2 map[int]uint64) map[int]uint64 {
	if map1 == nil && map2 == nil {
		return map[int]uint64{}
	}

	newMap := make(map[int]uint64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeIntUint32 takes two inputs: map[int]uint32 and map[int]uint32 and merge two maps and returns a new map[int]uint32.
func MergeIntUint32(map1, map2 map[int]uint32) map[int]uint32 {
	if map1 == nil && map2 == nil {
		return map[int]uint32{}
	}

	newMap := make(map[int]uint32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeIntUint16 takes two inputs: map[int]uint16 and map[int]uint16 and merge two maps and returns a new map[int]uint16.
func MergeIntUint16(map1, map2 map[int]uint16) map[int]uint16 {
	if map1 == nil && map2 == nil {
		return map[int]uint16{}
	}

	newMap := make(map[int]uint16)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeIntUint8 takes two inputs: map[int]uint8 and map[int]uint8 and merge two maps and returns a new map[int]uint8.
func MergeIntUint8(map1, map2 map[int]uint8) map[int]uint8 {
	if map1 == nil && map2 == nil {
		return map[int]uint8{}
	}

	newMap := make(map[int]uint8)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeIntStr takes two inputs: map[int]string and map[int]string and merge two maps and returns a new map[int]string.
func MergeIntStr(map1, map2 map[int]string) map[int]string {
	if map1 == nil && map2 == nil {
		return map[int]string{}
	}

	newMap := make(map[int]string)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeIntBool takes two inputs: map[int]bool and map[int]bool and merge two maps and returns a new map[int]bool.
func MergeIntBool(map1, map2 map[int]bool) map[int]bool {
	if map1 == nil && map2 == nil {
		return map[int]bool{}
	}

	newMap := make(map[int]bool)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeIntFloat32 takes two inputs: map[int]float32 and map[int]float32 and merge two maps and returns a new map[int]float32.
func MergeIntFloat32(map1, map2 map[int]float32) map[int]float32 {
	if map1 == nil && map2 == nil {
		return map[int]float32{}
	}

	newMap := make(map[int]float32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeIntFloat64 takes two inputs: map[int]float64 and map[int]float64 and merge two maps and returns a new map[int]float64.
func MergeIntFloat64(map1, map2 map[int]float64) map[int]float64 {
	if map1 == nil && map2 == nil {
		return map[int]float64{}
	}

	newMap := make(map[int]float64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt64Int takes two inputs: map[int64]int and map[int64]int and merge two maps and returns a new map[int64]int.
func MergeInt64Int(map1, map2 map[int64]int) map[int64]int {
	if map1 == nil && map2 == nil {
		return map[int64]int{}
	}

	newMap := make(map[int64]int)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt64 takes two inputs: map[int64]int64 and map[int64]int64 and merge two maps and returns a new map[int64]int64.
func MergeInt64(map1, map2 map[int64]int64) map[int64]int64 {
	if map1 == nil && map2 == nil {
		return map[int64]int64{}
	}

	newMap := make(map[int64]int64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt64Int32 takes two inputs: map[int64]int32 and map[int64]int32 and merge two maps and returns a new map[int64]int32.
func MergeInt64Int32(map1, map2 map[int64]int32) map[int64]int32 {
	if map1 == nil && map2 == nil {
		return map[int64]int32{}
	}

	newMap := make(map[int64]int32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt64Int16 takes two inputs: map[int64]int16 and map[int64]int16 and merge two maps and returns a new map[int64]int16.
func MergeInt64Int16(map1, map2 map[int64]int16) map[int64]int16 {
	if map1 == nil && map2 == nil {
		return map[int64]int16{}
	}

	newMap := make(map[int64]int16)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt64Int8 takes two inputs: map[int64]int8 and map[int64]int8 and merge two maps and returns a new map[int64]int8.
func MergeInt64Int8(map1, map2 map[int64]int8) map[int64]int8 {
	if map1 == nil && map2 == nil {
		return map[int64]int8{}
	}

	newMap := make(map[int64]int8)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt64Uint takes two inputs: map[int64]uint and map[int64]uint and merge two maps and returns a new map[int64]uint.
func MergeInt64Uint(map1, map2 map[int64]uint) map[int64]uint {
	if map1 == nil && map2 == nil {
		return map[int64]uint{}
	}

	newMap := make(map[int64]uint)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt64Uint64 takes two inputs: map[int64]uint64 and map[int64]uint64 and merge two maps and returns a new map[int64]uint64.
func MergeInt64Uint64(map1, map2 map[int64]uint64) map[int64]uint64 {
	if map1 == nil && map2 == nil {
		return map[int64]uint64{}
	}

	newMap := make(map[int64]uint64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt64Uint32 takes two inputs: map[int64]uint32 and map[int64]uint32 and merge two maps and returns a new map[int64]uint32.
func MergeInt64Uint32(map1, map2 map[int64]uint32) map[int64]uint32 {
	if map1 == nil && map2 == nil {
		return map[int64]uint32{}
	}

	newMap := make(map[int64]uint32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt64Uint16 takes two inputs: map[int64]uint16 and map[int64]uint16 and merge two maps and returns a new map[int64]uint16.
func MergeInt64Uint16(map1, map2 map[int64]uint16) map[int64]uint16 {
	if map1 == nil && map2 == nil {
		return map[int64]uint16{}
	}

	newMap := make(map[int64]uint16)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt64Uint8 takes two inputs: map[int64]uint8 and map[int64]uint8 and merge two maps and returns a new map[int64]uint8.
func MergeInt64Uint8(map1, map2 map[int64]uint8) map[int64]uint8 {
	if map1 == nil && map2 == nil {
		return map[int64]uint8{}
	}

	newMap := make(map[int64]uint8)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt64Str takes two inputs: map[int64]string and map[int64]string and merge two maps and returns a new map[int64]string.
func MergeInt64Str(map1, map2 map[int64]string) map[int64]string {
	if map1 == nil && map2 == nil {
		return map[int64]string{}
	}

	newMap := make(map[int64]string)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt64Bool takes two inputs: map[int64]bool and map[int64]bool and merge two maps and returns a new map[int64]bool.
func MergeInt64Bool(map1, map2 map[int64]bool) map[int64]bool {
	if map1 == nil && map2 == nil {
		return map[int64]bool{}
	}

	newMap := make(map[int64]bool)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt64Float32 takes two inputs: map[int64]float32 and map[int64]float32 and merge two maps and returns a new map[int64]float32.
func MergeInt64Float32(map1, map2 map[int64]float32) map[int64]float32 {
	if map1 == nil && map2 == nil {
		return map[int64]float32{}
	}

	newMap := make(map[int64]float32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt64Float64 takes two inputs: map[int64]float64 and map[int64]float64 and merge two maps and returns a new map[int64]float64.
func MergeInt64Float64(map1, map2 map[int64]float64) map[int64]float64 {
	if map1 == nil && map2 == nil {
		return map[int64]float64{}
	}

	newMap := make(map[int64]float64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt32Int takes two inputs: map[int32]int and map[int32]int and merge two maps and returns a new map[int32]int.
func MergeInt32Int(map1, map2 map[int32]int) map[int32]int {
	if map1 == nil && map2 == nil {
		return map[int32]int{}
	}

	newMap := make(map[int32]int)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt32Int64 takes two inputs: map[int32]int64 and map[int32]int64 and merge two maps and returns a new map[int32]int64.
func MergeInt32Int64(map1, map2 map[int32]int64) map[int32]int64 {
	if map1 == nil && map2 == nil {
		return map[int32]int64{}
	}

	newMap := make(map[int32]int64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt32 takes two inputs: map[int32]int32 and map[int32]int32 and merge two maps and returns a new map[int32]int32.
func MergeInt32(map1, map2 map[int32]int32) map[int32]int32 {
	if map1 == nil && map2 == nil {
		return map[int32]int32{}
	}

	newMap := make(map[int32]int32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt32Int16 takes two inputs: map[int32]int16 and map[int32]int16 and merge two maps and returns a new map[int32]int16.
func MergeInt32Int16(map1, map2 map[int32]int16) map[int32]int16 {
	if map1 == nil && map2 == nil {
		return map[int32]int16{}
	}

	newMap := make(map[int32]int16)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt32Int8 takes two inputs: map[int32]int8 and map[int32]int8 and merge two maps and returns a new map[int32]int8.
func MergeInt32Int8(map1, map2 map[int32]int8) map[int32]int8 {
	if map1 == nil && map2 == nil {
		return map[int32]int8{}
	}

	newMap := make(map[int32]int8)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt32Uint takes two inputs: map[int32]uint and map[int32]uint and merge two maps and returns a new map[int32]uint.
func MergeInt32Uint(map1, map2 map[int32]uint) map[int32]uint {
	if map1 == nil && map2 == nil {
		return map[int32]uint{}
	}

	newMap := make(map[int32]uint)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt32Uint64 takes two inputs: map[int32]uint64 and map[int32]uint64 and merge two maps and returns a new map[int32]uint64.
func MergeInt32Uint64(map1, map2 map[int32]uint64) map[int32]uint64 {
	if map1 == nil && map2 == nil {
		return map[int32]uint64{}
	}

	newMap := make(map[int32]uint64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt32Uint32 takes two inputs: map[int32]uint32 and map[int32]uint32 and merge two maps and returns a new map[int32]uint32.
func MergeInt32Uint32(map1, map2 map[int32]uint32) map[int32]uint32 {
	if map1 == nil && map2 == nil {
		return map[int32]uint32{}
	}

	newMap := make(map[int32]uint32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt32Uint16 takes two inputs: map[int32]uint16 and map[int32]uint16 and merge two maps and returns a new map[int32]uint16.
func MergeInt32Uint16(map1, map2 map[int32]uint16) map[int32]uint16 {
	if map1 == nil && map2 == nil {
		return map[int32]uint16{}
	}

	newMap := make(map[int32]uint16)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt32Uint8 takes two inputs: map[int32]uint8 and map[int32]uint8 and merge two maps and returns a new map[int32]uint8.
func MergeInt32Uint8(map1, map2 map[int32]uint8) map[int32]uint8 {
	if map1 == nil && map2 == nil {
		return map[int32]uint8{}
	}

	newMap := make(map[int32]uint8)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt32Str takes two inputs: map[int32]string and map[int32]string and merge two maps and returns a new map[int32]string.
func MergeInt32Str(map1, map2 map[int32]string) map[int32]string {
	if map1 == nil && map2 == nil {
		return map[int32]string{}
	}

	newMap := make(map[int32]string)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt32Bool takes two inputs: map[int32]bool and map[int32]bool and merge two maps and returns a new map[int32]bool.
func MergeInt32Bool(map1, map2 map[int32]bool) map[int32]bool {
	if map1 == nil && map2 == nil {
		return map[int32]bool{}
	}

	newMap := make(map[int32]bool)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt32Float32 takes two inputs: map[int32]float32 and map[int32]float32 and merge two maps and returns a new map[int32]float32.
func MergeInt32Float32(map1, map2 map[int32]float32) map[int32]float32 {
	if map1 == nil && map2 == nil {
		return map[int32]float32{}
	}

	newMap := make(map[int32]float32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt32Float64 takes two inputs: map[int32]float64 and map[int32]float64 and merge two maps and returns a new map[int32]float64.
func MergeInt32Float64(map1, map2 map[int32]float64) map[int32]float64 {
	if map1 == nil && map2 == nil {
		return map[int32]float64{}
	}

	newMap := make(map[int32]float64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt16Int takes two inputs: map[int16]int and map[int16]int and merge two maps and returns a new map[int16]int.
func MergeInt16Int(map1, map2 map[int16]int) map[int16]int {
	if map1 == nil && map2 == nil {
		return map[int16]int{}
	}

	newMap := make(map[int16]int)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt16Int64 takes two inputs: map[int16]int64 and map[int16]int64 and merge two maps and returns a new map[int16]int64.
func MergeInt16Int64(map1, map2 map[int16]int64) map[int16]int64 {
	if map1 == nil && map2 == nil {
		return map[int16]int64{}
	}

	newMap := make(map[int16]int64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt16Int32 takes two inputs: map[int16]int32 and map[int16]int32 and merge two maps and returns a new map[int16]int32.
func MergeInt16Int32(map1, map2 map[int16]int32) map[int16]int32 {
	if map1 == nil && map2 == nil {
		return map[int16]int32{}
	}

	newMap := make(map[int16]int32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt16 takes two inputs: map[int16]int16 and map[int16]int16 and merge two maps and returns a new map[int16]int16.
func MergeInt16(map1, map2 map[int16]int16) map[int16]int16 {
	if map1 == nil && map2 == nil {
		return map[int16]int16{}
	}

	newMap := make(map[int16]int16)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt16Int8 takes two inputs: map[int16]int8 and map[int16]int8 and merge two maps and returns a new map[int16]int8.
func MergeInt16Int8(map1, map2 map[int16]int8) map[int16]int8 {
	if map1 == nil && map2 == nil {
		return map[int16]int8{}
	}

	newMap := make(map[int16]int8)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt16Uint takes two inputs: map[int16]uint and map[int16]uint and merge two maps and returns a new map[int16]uint.
func MergeInt16Uint(map1, map2 map[int16]uint) map[int16]uint {
	if map1 == nil && map2 == nil {
		return map[int16]uint{}
	}

	newMap := make(map[int16]uint)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt16Uint64 takes two inputs: map[int16]uint64 and map[int16]uint64 and merge two maps and returns a new map[int16]uint64.
func MergeInt16Uint64(map1, map2 map[int16]uint64) map[int16]uint64 {
	if map1 == nil && map2 == nil {
		return map[int16]uint64{}
	}

	newMap := make(map[int16]uint64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt16Uint32 takes two inputs: map[int16]uint32 and map[int16]uint32 and merge two maps and returns a new map[int16]uint32.
func MergeInt16Uint32(map1, map2 map[int16]uint32) map[int16]uint32 {
	if map1 == nil && map2 == nil {
		return map[int16]uint32{}
	}

	newMap := make(map[int16]uint32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt16Uint16 takes two inputs: map[int16]uint16 and map[int16]uint16 and merge two maps and returns a new map[int16]uint16.
func MergeInt16Uint16(map1, map2 map[int16]uint16) map[int16]uint16 {
	if map1 == nil && map2 == nil {
		return map[int16]uint16{}
	}

	newMap := make(map[int16]uint16)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt16Uint8 takes two inputs: map[int16]uint8 and map[int16]uint8 and merge two maps and returns a new map[int16]uint8.
func MergeInt16Uint8(map1, map2 map[int16]uint8) map[int16]uint8 {
	if map1 == nil && map2 == nil {
		return map[int16]uint8{}
	}

	newMap := make(map[int16]uint8)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt16Str takes two inputs: map[int16]string and map[int16]string and merge two maps and returns a new map[int16]string.
func MergeInt16Str(map1, map2 map[int16]string) map[int16]string {
	if map1 == nil && map2 == nil {
		return map[int16]string{}
	}

	newMap := make(map[int16]string)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt16Bool takes two inputs: map[int16]bool and map[int16]bool and merge two maps and returns a new map[int16]bool.
func MergeInt16Bool(map1, map2 map[int16]bool) map[int16]bool {
	if map1 == nil && map2 == nil {
		return map[int16]bool{}
	}

	newMap := make(map[int16]bool)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt16Float32 takes two inputs: map[int16]float32 and map[int16]float32 and merge two maps and returns a new map[int16]float32.
func MergeInt16Float32(map1, map2 map[int16]float32) map[int16]float32 {
	if map1 == nil && map2 == nil {
		return map[int16]float32{}
	}

	newMap := make(map[int16]float32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt16Float64 takes two inputs: map[int16]float64 and map[int16]float64 and merge two maps and returns a new map[int16]float64.
func MergeInt16Float64(map1, map2 map[int16]float64) map[int16]float64 {
	if map1 == nil && map2 == nil {
		return map[int16]float64{}
	}

	newMap := make(map[int16]float64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt8Int takes two inputs: map[int8]int and map[int8]int and merge two maps and returns a new map[int8]int.
func MergeInt8Int(map1, map2 map[int8]int) map[int8]int {
	if map1 == nil && map2 == nil {
		return map[int8]int{}
	}

	newMap := make(map[int8]int)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt8Int64 takes two inputs: map[int8]int64 and map[int8]int64 and merge two maps and returns a new map[int8]int64.
func MergeInt8Int64(map1, map2 map[int8]int64) map[int8]int64 {
	if map1 == nil && map2 == nil {
		return map[int8]int64{}
	}

	newMap := make(map[int8]int64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt8Int32 takes two inputs: map[int8]int32 and map[int8]int32 and merge two maps and returns a new map[int8]int32.
func MergeInt8Int32(map1, map2 map[int8]int32) map[int8]int32 {
	if map1 == nil && map2 == nil {
		return map[int8]int32{}
	}

	newMap := make(map[int8]int32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt8Int16 takes two inputs: map[int8]int16 and map[int8]int16 and merge two maps and returns a new map[int8]int16.
func MergeInt8Int16(map1, map2 map[int8]int16) map[int8]int16 {
	if map1 == nil && map2 == nil {
		return map[int8]int16{}
	}

	newMap := make(map[int8]int16)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt8 takes two inputs: map[int8]int8 and map[int8]int8 and merge two maps and returns a new map[int8]int8.
func MergeInt8(map1, map2 map[int8]int8) map[int8]int8 {
	if map1 == nil && map2 == nil {
		return map[int8]int8{}
	}

	newMap := make(map[int8]int8)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt8Uint takes two inputs: map[int8]uint and map[int8]uint and merge two maps and returns a new map[int8]uint.
func MergeInt8Uint(map1, map2 map[int8]uint) map[int8]uint {
	if map1 == nil && map2 == nil {
		return map[int8]uint{}
	}

	newMap := make(map[int8]uint)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt8Uint64 takes two inputs: map[int8]uint64 and map[int8]uint64 and merge two maps and returns a new map[int8]uint64.
func MergeInt8Uint64(map1, map2 map[int8]uint64) map[int8]uint64 {
	if map1 == nil && map2 == nil {
		return map[int8]uint64{}
	}

	newMap := make(map[int8]uint64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt8Uint32 takes two inputs: map[int8]uint32 and map[int8]uint32 and merge two maps and returns a new map[int8]uint32.
func MergeInt8Uint32(map1, map2 map[int8]uint32) map[int8]uint32 {
	if map1 == nil && map2 == nil {
		return map[int8]uint32{}
	}

	newMap := make(map[int8]uint32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt8Uint16 takes two inputs: map[int8]uint16 and map[int8]uint16 and merge two maps and returns a new map[int8]uint16.
func MergeInt8Uint16(map1, map2 map[int8]uint16) map[int8]uint16 {
	if map1 == nil && map2 == nil {
		return map[int8]uint16{}
	}

	newMap := make(map[int8]uint16)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt8Uint8 takes two inputs: map[int8]uint8 and map[int8]uint8 and merge two maps and returns a new map[int8]uint8.
func MergeInt8Uint8(map1, map2 map[int8]uint8) map[int8]uint8 {
	if map1 == nil && map2 == nil {
		return map[int8]uint8{}
	}

	newMap := make(map[int8]uint8)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt8Str takes two inputs: map[int8]string and map[int8]string and merge two maps and returns a new map[int8]string.
func MergeInt8Str(map1, map2 map[int8]string) map[int8]string {
	if map1 == nil && map2 == nil {
		return map[int8]string{}
	}

	newMap := make(map[int8]string)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt8Bool takes two inputs: map[int8]bool and map[int8]bool and merge two maps and returns a new map[int8]bool.
func MergeInt8Bool(map1, map2 map[int8]bool) map[int8]bool {
	if map1 == nil && map2 == nil {
		return map[int8]bool{}
	}

	newMap := make(map[int8]bool)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt8Float32 takes two inputs: map[int8]float32 and map[int8]float32 and merge two maps and returns a new map[int8]float32.
func MergeInt8Float32(map1, map2 map[int8]float32) map[int8]float32 {
	if map1 == nil && map2 == nil {
		return map[int8]float32{}
	}

	newMap := make(map[int8]float32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeInt8Float64 takes two inputs: map[int8]float64 and map[int8]float64 and merge two maps and returns a new map[int8]float64.
func MergeInt8Float64(map1, map2 map[int8]float64) map[int8]float64 {
	if map1 == nil && map2 == nil {
		return map[int8]float64{}
	}

	newMap := make(map[int8]float64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUintInt takes two inputs: map[uint]int and map[uint]int and merge two maps and returns a new map[uint]int.
func MergeUintInt(map1, map2 map[uint]int) map[uint]int {
	if map1 == nil && map2 == nil {
		return map[uint]int{}
	}

	newMap := make(map[uint]int)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUintInt64 takes two inputs: map[uint]int64 and map[uint]int64 and merge two maps and returns a new map[uint]int64.
func MergeUintInt64(map1, map2 map[uint]int64) map[uint]int64 {
	if map1 == nil && map2 == nil {
		return map[uint]int64{}
	}

	newMap := make(map[uint]int64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUintInt32 takes two inputs: map[uint]int32 and map[uint]int32 and merge two maps and returns a new map[uint]int32.
func MergeUintInt32(map1, map2 map[uint]int32) map[uint]int32 {
	if map1 == nil && map2 == nil {
		return map[uint]int32{}
	}

	newMap := make(map[uint]int32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUintInt16 takes two inputs: map[uint]int16 and map[uint]int16 and merge two maps and returns a new map[uint]int16.
func MergeUintInt16(map1, map2 map[uint]int16) map[uint]int16 {
	if map1 == nil && map2 == nil {
		return map[uint]int16{}
	}

	newMap := make(map[uint]int16)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUintInt8 takes two inputs: map[uint]int8 and map[uint]int8 and merge two maps and returns a new map[uint]int8.
func MergeUintInt8(map1, map2 map[uint]int8) map[uint]int8 {
	if map1 == nil && map2 == nil {
		return map[uint]int8{}
	}

	newMap := make(map[uint]int8)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint takes two inputs: map[uint]uint and map[uint]uint and merge two maps and returns a new map[uint]uint.
func MergeUint(map1, map2 map[uint]uint) map[uint]uint {
	if map1 == nil && map2 == nil {
		return map[uint]uint{}
	}

	newMap := make(map[uint]uint)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUintUint64 takes two inputs: map[uint]uint64 and map[uint]uint64 and merge two maps and returns a new map[uint]uint64.
func MergeUintUint64(map1, map2 map[uint]uint64) map[uint]uint64 {
	if map1 == nil && map2 == nil {
		return map[uint]uint64{}
	}

	newMap := make(map[uint]uint64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUintUint32 takes two inputs: map[uint]uint32 and map[uint]uint32 and merge two maps and returns a new map[uint]uint32.
func MergeUintUint32(map1, map2 map[uint]uint32) map[uint]uint32 {
	if map1 == nil && map2 == nil {
		return map[uint]uint32{}
	}

	newMap := make(map[uint]uint32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUintUint16 takes two inputs: map[uint]uint16 and map[uint]uint16 and merge two maps and returns a new map[uint]uint16.
func MergeUintUint16(map1, map2 map[uint]uint16) map[uint]uint16 {
	if map1 == nil && map2 == nil {
		return map[uint]uint16{}
	}

	newMap := make(map[uint]uint16)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUintUint8 takes two inputs: map[uint]uint8 and map[uint]uint8 and merge two maps and returns a new map[uint]uint8.
func MergeUintUint8(map1, map2 map[uint]uint8) map[uint]uint8 {
	if map1 == nil && map2 == nil {
		return map[uint]uint8{}
	}

	newMap := make(map[uint]uint8)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUintStr takes two inputs: map[uint]string and map[uint]string and merge two maps and returns a new map[uint]string.
func MergeUintStr(map1, map2 map[uint]string) map[uint]string {
	if map1 == nil && map2 == nil {
		return map[uint]string{}
	}

	newMap := make(map[uint]string)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUintBool takes two inputs: map[uint]bool and map[uint]bool and merge two maps and returns a new map[uint]bool.
func MergeUintBool(map1, map2 map[uint]bool) map[uint]bool {
	if map1 == nil && map2 == nil {
		return map[uint]bool{}
	}

	newMap := make(map[uint]bool)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUintFloat32 takes two inputs: map[uint]float32 and map[uint]float32 and merge two maps and returns a new map[uint]float32.
func MergeUintFloat32(map1, map2 map[uint]float32) map[uint]float32 {
	if map1 == nil && map2 == nil {
		return map[uint]float32{}
	}

	newMap := make(map[uint]float32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUintFloat64 takes two inputs: map[uint]float64 and map[uint]float64 and merge two maps and returns a new map[uint]float64.
func MergeUintFloat64(map1, map2 map[uint]float64) map[uint]float64 {
	if map1 == nil && map2 == nil {
		return map[uint]float64{}
	}

	newMap := make(map[uint]float64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint64Int takes two inputs: map[uint64]int and map[uint64]int and merge two maps and returns a new map[uint64]int.
func MergeUint64Int(map1, map2 map[uint64]int) map[uint64]int {
	if map1 == nil && map2 == nil {
		return map[uint64]int{}
	}

	newMap := make(map[uint64]int)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint64Int64 takes two inputs: map[uint64]int64 and map[uint64]int64 and merge two maps and returns a new map[uint64]int64.
func MergeUint64Int64(map1, map2 map[uint64]int64) map[uint64]int64 {
	if map1 == nil && map2 == nil {
		return map[uint64]int64{}
	}

	newMap := make(map[uint64]int64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint64Int32 takes two inputs: map[uint64]int32 and map[uint64]int32 and merge two maps and returns a new map[uint64]int32.
func MergeUint64Int32(map1, map2 map[uint64]int32) map[uint64]int32 {
	if map1 == nil && map2 == nil {
		return map[uint64]int32{}
	}

	newMap := make(map[uint64]int32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint64Int16 takes two inputs: map[uint64]int16 and map[uint64]int16 and merge two maps and returns a new map[uint64]int16.
func MergeUint64Int16(map1, map2 map[uint64]int16) map[uint64]int16 {
	if map1 == nil && map2 == nil {
		return map[uint64]int16{}
	}

	newMap := make(map[uint64]int16)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint64Int8 takes two inputs: map[uint64]int8 and map[uint64]int8 and merge two maps and returns a new map[uint64]int8.
func MergeUint64Int8(map1, map2 map[uint64]int8) map[uint64]int8 {
	if map1 == nil && map2 == nil {
		return map[uint64]int8{}
	}

	newMap := make(map[uint64]int8)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint64Uint takes two inputs: map[uint64]uint and map[uint64]uint and merge two maps and returns a new map[uint64]uint.
func MergeUint64Uint(map1, map2 map[uint64]uint) map[uint64]uint {
	if map1 == nil && map2 == nil {
		return map[uint64]uint{}
	}

	newMap := make(map[uint64]uint)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint64 takes two inputs: map[uint64]uint64 and map[uint64]uint64 and merge two maps and returns a new map[uint64]uint64.
func MergeUint64(map1, map2 map[uint64]uint64) map[uint64]uint64 {
	if map1 == nil && map2 == nil {
		return map[uint64]uint64{}
	}

	newMap := make(map[uint64]uint64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint64Uint32 takes two inputs: map[uint64]uint32 and map[uint64]uint32 and merge two maps and returns a new map[uint64]uint32.
func MergeUint64Uint32(map1, map2 map[uint64]uint32) map[uint64]uint32 {
	if map1 == nil && map2 == nil {
		return map[uint64]uint32{}
	}

	newMap := make(map[uint64]uint32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint64Uint16 takes two inputs: map[uint64]uint16 and map[uint64]uint16 and merge two maps and returns a new map[uint64]uint16.
func MergeUint64Uint16(map1, map2 map[uint64]uint16) map[uint64]uint16 {
	if map1 == nil && map2 == nil {
		return map[uint64]uint16{}
	}

	newMap := make(map[uint64]uint16)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint64Uint8 takes two inputs: map[uint64]uint8 and map[uint64]uint8 and merge two maps and returns a new map[uint64]uint8.
func MergeUint64Uint8(map1, map2 map[uint64]uint8) map[uint64]uint8 {
	if map1 == nil && map2 == nil {
		return map[uint64]uint8{}
	}

	newMap := make(map[uint64]uint8)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint64Str takes two inputs: map[uint64]string and map[uint64]string and merge two maps and returns a new map[uint64]string.
func MergeUint64Str(map1, map2 map[uint64]string) map[uint64]string {
	if map1 == nil && map2 == nil {
		return map[uint64]string{}
	}

	newMap := make(map[uint64]string)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint64Bool takes two inputs: map[uint64]bool and map[uint64]bool and merge two maps and returns a new map[uint64]bool.
func MergeUint64Bool(map1, map2 map[uint64]bool) map[uint64]bool {
	if map1 == nil && map2 == nil {
		return map[uint64]bool{}
	}

	newMap := make(map[uint64]bool)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint64Float32 takes two inputs: map[uint64]float32 and map[uint64]float32 and merge two maps and returns a new map[uint64]float32.
func MergeUint64Float32(map1, map2 map[uint64]float32) map[uint64]float32 {
	if map1 == nil && map2 == nil {
		return map[uint64]float32{}
	}

	newMap := make(map[uint64]float32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint64Float64 takes two inputs: map[uint64]float64 and map[uint64]float64 and merge two maps and returns a new map[uint64]float64.
func MergeUint64Float64(map1, map2 map[uint64]float64) map[uint64]float64 {
	if map1 == nil && map2 == nil {
		return map[uint64]float64{}
	}

	newMap := make(map[uint64]float64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint32Int takes two inputs: map[uint32]int and map[uint32]int and merge two maps and returns a new map[uint32]int.
func MergeUint32Int(map1, map2 map[uint32]int) map[uint32]int {
	if map1 == nil && map2 == nil {
		return map[uint32]int{}
	}

	newMap := make(map[uint32]int)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint32Int64 takes two inputs: map[uint32]int64 and map[uint32]int64 and merge two maps and returns a new map[uint32]int64.
func MergeUint32Int64(map1, map2 map[uint32]int64) map[uint32]int64 {
	if map1 == nil && map2 == nil {
		return map[uint32]int64{}
	}

	newMap := make(map[uint32]int64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint32Int32 takes two inputs: map[uint32]int32 and map[uint32]int32 and merge two maps and returns a new map[uint32]int32.
func MergeUint32Int32(map1, map2 map[uint32]int32) map[uint32]int32 {
	if map1 == nil && map2 == nil {
		return map[uint32]int32{}
	}

	newMap := make(map[uint32]int32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint32Int16 takes two inputs: map[uint32]int16 and map[uint32]int16 and merge two maps and returns a new map[uint32]int16.
func MergeUint32Int16(map1, map2 map[uint32]int16) map[uint32]int16 {
	if map1 == nil && map2 == nil {
		return map[uint32]int16{}
	}

	newMap := make(map[uint32]int16)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint32Int8 takes two inputs: map[uint32]int8 and map[uint32]int8 and merge two maps and returns a new map[uint32]int8.
func MergeUint32Int8(map1, map2 map[uint32]int8) map[uint32]int8 {
	if map1 == nil && map2 == nil {
		return map[uint32]int8{}
	}

	newMap := make(map[uint32]int8)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint32Uint takes two inputs: map[uint32]uint and map[uint32]uint and merge two maps and returns a new map[uint32]uint.
func MergeUint32Uint(map1, map2 map[uint32]uint) map[uint32]uint {
	if map1 == nil && map2 == nil {
		return map[uint32]uint{}
	}

	newMap := make(map[uint32]uint)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint32Uint64 takes two inputs: map[uint32]uint64 and map[uint32]uint64 and merge two maps and returns a new map[uint32]uint64.
func MergeUint32Uint64(map1, map2 map[uint32]uint64) map[uint32]uint64 {
	if map1 == nil && map2 == nil {
		return map[uint32]uint64{}
	}

	newMap := make(map[uint32]uint64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint32 takes two inputs: map[uint32]uint32 and map[uint32]uint32 and merge two maps and returns a new map[uint32]uint32.
func MergeUint32(map1, map2 map[uint32]uint32) map[uint32]uint32 {
	if map1 == nil && map2 == nil {
		return map[uint32]uint32{}
	}

	newMap := make(map[uint32]uint32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint32Uint16 takes two inputs: map[uint32]uint16 and map[uint32]uint16 and merge two maps and returns a new map[uint32]uint16.
func MergeUint32Uint16(map1, map2 map[uint32]uint16) map[uint32]uint16 {
	if map1 == nil && map2 == nil {
		return map[uint32]uint16{}
	}

	newMap := make(map[uint32]uint16)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint32Uint8 takes two inputs: map[uint32]uint8 and map[uint32]uint8 and merge two maps and returns a new map[uint32]uint8.
func MergeUint32Uint8(map1, map2 map[uint32]uint8) map[uint32]uint8 {
	if map1 == nil && map2 == nil {
		return map[uint32]uint8{}
	}

	newMap := make(map[uint32]uint8)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint32Str takes two inputs: map[uint32]string and map[uint32]string and merge two maps and returns a new map[uint32]string.
func MergeUint32Str(map1, map2 map[uint32]string) map[uint32]string {
	if map1 == nil && map2 == nil {
		return map[uint32]string{}
	}

	newMap := make(map[uint32]string)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint32Bool takes two inputs: map[uint32]bool and map[uint32]bool and merge two maps and returns a new map[uint32]bool.
func MergeUint32Bool(map1, map2 map[uint32]bool) map[uint32]bool {
	if map1 == nil && map2 == nil {
		return map[uint32]bool{}
	}

	newMap := make(map[uint32]bool)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint32Float32 takes two inputs: map[uint32]float32 and map[uint32]float32 and merge two maps and returns a new map[uint32]float32.
func MergeUint32Float32(map1, map2 map[uint32]float32) map[uint32]float32 {
	if map1 == nil && map2 == nil {
		return map[uint32]float32{}
	}

	newMap := make(map[uint32]float32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint32Float64 takes two inputs: map[uint32]float64 and map[uint32]float64 and merge two maps and returns a new map[uint32]float64.
func MergeUint32Float64(map1, map2 map[uint32]float64) map[uint32]float64 {
	if map1 == nil && map2 == nil {
		return map[uint32]float64{}
	}

	newMap := make(map[uint32]float64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint16Int takes two inputs: map[uint16]int and map[uint16]int and merge two maps and returns a new map[uint16]int.
func MergeUint16Int(map1, map2 map[uint16]int) map[uint16]int {
	if map1 == nil && map2 == nil {
		return map[uint16]int{}
	}

	newMap := make(map[uint16]int)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint16Int64 takes two inputs: map[uint16]int64 and map[uint16]int64 and merge two maps and returns a new map[uint16]int64.
func MergeUint16Int64(map1, map2 map[uint16]int64) map[uint16]int64 {
	if map1 == nil && map2 == nil {
		return map[uint16]int64{}
	}

	newMap := make(map[uint16]int64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint16Int32 takes two inputs: map[uint16]int32 and map[uint16]int32 and merge two maps and returns a new map[uint16]int32.
func MergeUint16Int32(map1, map2 map[uint16]int32) map[uint16]int32 {
	if map1 == nil && map2 == nil {
		return map[uint16]int32{}
	}

	newMap := make(map[uint16]int32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint16Int16 takes two inputs: map[uint16]int16 and map[uint16]int16 and merge two maps and returns a new map[uint16]int16.
func MergeUint16Int16(map1, map2 map[uint16]int16) map[uint16]int16 {
	if map1 == nil && map2 == nil {
		return map[uint16]int16{}
	}

	newMap := make(map[uint16]int16)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint16Int8 takes two inputs: map[uint16]int8 and map[uint16]int8 and merge two maps and returns a new map[uint16]int8.
func MergeUint16Int8(map1, map2 map[uint16]int8) map[uint16]int8 {
	if map1 == nil && map2 == nil {
		return map[uint16]int8{}
	}

	newMap := make(map[uint16]int8)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint16Uint takes two inputs: map[uint16]uint and map[uint16]uint and merge two maps and returns a new map[uint16]uint.
func MergeUint16Uint(map1, map2 map[uint16]uint) map[uint16]uint {
	if map1 == nil && map2 == nil {
		return map[uint16]uint{}
	}

	newMap := make(map[uint16]uint)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint16Uint64 takes two inputs: map[uint16]uint64 and map[uint16]uint64 and merge two maps and returns a new map[uint16]uint64.
func MergeUint16Uint64(map1, map2 map[uint16]uint64) map[uint16]uint64 {
	if map1 == nil && map2 == nil {
		return map[uint16]uint64{}
	}

	newMap := make(map[uint16]uint64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint16Uint32 takes two inputs: map[uint16]uint32 and map[uint16]uint32 and merge two maps and returns a new map[uint16]uint32.
func MergeUint16Uint32(map1, map2 map[uint16]uint32) map[uint16]uint32 {
	if map1 == nil && map2 == nil {
		return map[uint16]uint32{}
	}

	newMap := make(map[uint16]uint32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint16 takes two inputs: map[uint16]uint16 and map[uint16]uint16 and merge two maps and returns a new map[uint16]uint16.
func MergeUint16(map1, map2 map[uint16]uint16) map[uint16]uint16 {
	if map1 == nil && map2 == nil {
		return map[uint16]uint16{}
	}

	newMap := make(map[uint16]uint16)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint16Uint8 takes two inputs: map[uint16]uint8 and map[uint16]uint8 and merge two maps and returns a new map[uint16]uint8.
func MergeUint16Uint8(map1, map2 map[uint16]uint8) map[uint16]uint8 {
	if map1 == nil && map2 == nil {
		return map[uint16]uint8{}
	}

	newMap := make(map[uint16]uint8)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint16Str takes two inputs: map[uint16]string and map[uint16]string and merge two maps and returns a new map[uint16]string.
func MergeUint16Str(map1, map2 map[uint16]string) map[uint16]string {
	if map1 == nil && map2 == nil {
		return map[uint16]string{}
	}

	newMap := make(map[uint16]string)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint16Bool takes two inputs: map[uint16]bool and map[uint16]bool and merge two maps and returns a new map[uint16]bool.
func MergeUint16Bool(map1, map2 map[uint16]bool) map[uint16]bool {
	if map1 == nil && map2 == nil {
		return map[uint16]bool{}
	}

	newMap := make(map[uint16]bool)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint16Float32 takes two inputs: map[uint16]float32 and map[uint16]float32 and merge two maps and returns a new map[uint16]float32.
func MergeUint16Float32(map1, map2 map[uint16]float32) map[uint16]float32 {
	if map1 == nil && map2 == nil {
		return map[uint16]float32{}
	}

	newMap := make(map[uint16]float32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint16Float64 takes two inputs: map[uint16]float64 and map[uint16]float64 and merge two maps and returns a new map[uint16]float64.
func MergeUint16Float64(map1, map2 map[uint16]float64) map[uint16]float64 {
	if map1 == nil && map2 == nil {
		return map[uint16]float64{}
	}

	newMap := make(map[uint16]float64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint8Int takes two inputs: map[uint8]int and map[uint8]int and merge two maps and returns a new map[uint8]int.
func MergeUint8Int(map1, map2 map[uint8]int) map[uint8]int {
	if map1 == nil && map2 == nil {
		return map[uint8]int{}
	}

	newMap := make(map[uint8]int)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint8Int64 takes two inputs: map[uint8]int64 and map[uint8]int64 and merge two maps and returns a new map[uint8]int64.
func MergeUint8Int64(map1, map2 map[uint8]int64) map[uint8]int64 {
	if map1 == nil && map2 == nil {
		return map[uint8]int64{}
	}

	newMap := make(map[uint8]int64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint8Int32 takes two inputs: map[uint8]int32 and map[uint8]int32 and merge two maps and returns a new map[uint8]int32.
func MergeUint8Int32(map1, map2 map[uint8]int32) map[uint8]int32 {
	if map1 == nil && map2 == nil {
		return map[uint8]int32{}
	}

	newMap := make(map[uint8]int32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint8Int16 takes two inputs: map[uint8]int16 and map[uint8]int16 and merge two maps and returns a new map[uint8]int16.
func MergeUint8Int16(map1, map2 map[uint8]int16) map[uint8]int16 {
	if map1 == nil && map2 == nil {
		return map[uint8]int16{}
	}

	newMap := make(map[uint8]int16)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint8Int8 takes two inputs: map[uint8]int8 and map[uint8]int8 and merge two maps and returns a new map[uint8]int8.
func MergeUint8Int8(map1, map2 map[uint8]int8) map[uint8]int8 {
	if map1 == nil && map2 == nil {
		return map[uint8]int8{}
	}

	newMap := make(map[uint8]int8)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint8Uint takes two inputs: map[uint8]uint and map[uint8]uint and merge two maps and returns a new map[uint8]uint.
func MergeUint8Uint(map1, map2 map[uint8]uint) map[uint8]uint {
	if map1 == nil && map2 == nil {
		return map[uint8]uint{}
	}

	newMap := make(map[uint8]uint)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint8Uint64 takes two inputs: map[uint8]uint64 and map[uint8]uint64 and merge two maps and returns a new map[uint8]uint64.
func MergeUint8Uint64(map1, map2 map[uint8]uint64) map[uint8]uint64 {
	if map1 == nil && map2 == nil {
		return map[uint8]uint64{}
	}

	newMap := make(map[uint8]uint64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint8Uint32 takes two inputs: map[uint8]uint32 and map[uint8]uint32 and merge two maps and returns a new map[uint8]uint32.
func MergeUint8Uint32(map1, map2 map[uint8]uint32) map[uint8]uint32 {
	if map1 == nil && map2 == nil {
		return map[uint8]uint32{}
	}

	newMap := make(map[uint8]uint32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint8Uint16 takes two inputs: map[uint8]uint16 and map[uint8]uint16 and merge two maps and returns a new map[uint8]uint16.
func MergeUint8Uint16(map1, map2 map[uint8]uint16) map[uint8]uint16 {
	if map1 == nil && map2 == nil {
		return map[uint8]uint16{}
	}

	newMap := make(map[uint8]uint16)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint8 takes two inputs: map[uint8]uint8 and map[uint8]uint8 and merge two maps and returns a new map[uint8]uint8.
func MergeUint8(map1, map2 map[uint8]uint8) map[uint8]uint8 {
	if map1 == nil && map2 == nil {
		return map[uint8]uint8{}
	}

	newMap := make(map[uint8]uint8)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint8Str takes two inputs: map[uint8]string and map[uint8]string and merge two maps and returns a new map[uint8]string.
func MergeUint8Str(map1, map2 map[uint8]string) map[uint8]string {
	if map1 == nil && map2 == nil {
		return map[uint8]string{}
	}

	newMap := make(map[uint8]string)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint8Bool takes two inputs: map[uint8]bool and map[uint8]bool and merge two maps and returns a new map[uint8]bool.
func MergeUint8Bool(map1, map2 map[uint8]bool) map[uint8]bool {
	if map1 == nil && map2 == nil {
		return map[uint8]bool{}
	}

	newMap := make(map[uint8]bool)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint8Float32 takes two inputs: map[uint8]float32 and map[uint8]float32 and merge two maps and returns a new map[uint8]float32.
func MergeUint8Float32(map1, map2 map[uint8]float32) map[uint8]float32 {
	if map1 == nil && map2 == nil {
		return map[uint8]float32{}
	}

	newMap := make(map[uint8]float32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeUint8Float64 takes two inputs: map[uint8]float64 and map[uint8]float64 and merge two maps and returns a new map[uint8]float64.
func MergeUint8Float64(map1, map2 map[uint8]float64) map[uint8]float64 {
	if map1 == nil && map2 == nil {
		return map[uint8]float64{}
	}

	newMap := make(map[uint8]float64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeStrInt takes two inputs: map[string]int and map[string]int and merge two maps and returns a new map[string]int.
func MergeStrInt(map1, map2 map[string]int) map[string]int {
	if map1 == nil && map2 == nil {
		return map[string]int{}
	}

	newMap := make(map[string]int)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeStrInt64 takes two inputs: map[string]int64 and map[string]int64 and merge two maps and returns a new map[string]int64.
func MergeStrInt64(map1, map2 map[string]int64) map[string]int64 {
	if map1 == nil && map2 == nil {
		return map[string]int64{}
	}

	newMap := make(map[string]int64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeStrInt32 takes two inputs: map[string]int32 and map[string]int32 and merge two maps and returns a new map[string]int32.
func MergeStrInt32(map1, map2 map[string]int32) map[string]int32 {
	if map1 == nil && map2 == nil {
		return map[string]int32{}
	}

	newMap := make(map[string]int32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeStrInt16 takes two inputs: map[string]int16 and map[string]int16 and merge two maps and returns a new map[string]int16.
func MergeStrInt16(map1, map2 map[string]int16) map[string]int16 {
	if map1 == nil && map2 == nil {
		return map[string]int16{}
	}

	newMap := make(map[string]int16)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeStrInt8 takes two inputs: map[string]int8 and map[string]int8 and merge two maps and returns a new map[string]int8.
func MergeStrInt8(map1, map2 map[string]int8) map[string]int8 {
	if map1 == nil && map2 == nil {
		return map[string]int8{}
	}

	newMap := make(map[string]int8)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeStrUint takes two inputs: map[string]uint and map[string]uint and merge two maps and returns a new map[string]uint.
func MergeStrUint(map1, map2 map[string]uint) map[string]uint {
	if map1 == nil && map2 == nil {
		return map[string]uint{}
	}

	newMap := make(map[string]uint)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeStrUint64 takes two inputs: map[string]uint64 and map[string]uint64 and merge two maps and returns a new map[string]uint64.
func MergeStrUint64(map1, map2 map[string]uint64) map[string]uint64 {
	if map1 == nil && map2 == nil {
		return map[string]uint64{}
	}

	newMap := make(map[string]uint64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeStrUint32 takes two inputs: map[string]uint32 and map[string]uint32 and merge two maps and returns a new map[string]uint32.
func MergeStrUint32(map1, map2 map[string]uint32) map[string]uint32 {
	if map1 == nil && map2 == nil {
		return map[string]uint32{}
	}

	newMap := make(map[string]uint32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeStrUint16 takes two inputs: map[string]uint16 and map[string]uint16 and merge two maps and returns a new map[string]uint16.
func MergeStrUint16(map1, map2 map[string]uint16) map[string]uint16 {
	if map1 == nil && map2 == nil {
		return map[string]uint16{}
	}

	newMap := make(map[string]uint16)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeStrUint8 takes two inputs: map[string]uint8 and map[string]uint8 and merge two maps and returns a new map[string]uint8.
func MergeStrUint8(map1, map2 map[string]uint8) map[string]uint8 {
	if map1 == nil && map2 == nil {
		return map[string]uint8{}
	}

	newMap := make(map[string]uint8)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeStr takes two inputs: map[string]string and map[string]string and merge two maps and returns a new map[string]string.
func MergeStr(map1, map2 map[string]string) map[string]string {
	if map1 == nil && map2 == nil {
		return map[string]string{}
	}

	newMap := make(map[string]string)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeStrBool takes two inputs: map[string]bool and map[string]bool and merge two maps and returns a new map[string]bool.
func MergeStrBool(map1, map2 map[string]bool) map[string]bool {
	if map1 == nil && map2 == nil {
		return map[string]bool{}
	}

	newMap := make(map[string]bool)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeStrFloat32 takes two inputs: map[string]float32 and map[string]float32 and merge two maps and returns a new map[string]float32.
func MergeStrFloat32(map1, map2 map[string]float32) map[string]float32 {
	if map1 == nil && map2 == nil {
		return map[string]float32{}
	}

	newMap := make(map[string]float32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeStrFloat64 takes two inputs: map[string]float64 and map[string]float64 and merge two maps and returns a new map[string]float64.
func MergeStrFloat64(map1, map2 map[string]float64) map[string]float64 {
	if map1 == nil && map2 == nil {
		return map[string]float64{}
	}

	newMap := make(map[string]float64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeBoolInt takes two inputs: map[bool]int and map[bool]int and merge two maps and returns a new map[bool]int.
func MergeBoolInt(map1, map2 map[bool]int) map[bool]int {
	if map1 == nil && map2 == nil {
		return map[bool]int{}
	}

	newMap := make(map[bool]int)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeBoolInt64 takes two inputs: map[bool]int64 and map[bool]int64 and merge two maps and returns a new map[bool]int64.
func MergeBoolInt64(map1, map2 map[bool]int64) map[bool]int64 {
	if map1 == nil && map2 == nil {
		return map[bool]int64{}
	}

	newMap := make(map[bool]int64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeBoolInt32 takes two inputs: map[bool]int32 and map[bool]int32 and merge two maps and returns a new map[bool]int32.
func MergeBoolInt32(map1, map2 map[bool]int32) map[bool]int32 {
	if map1 == nil && map2 == nil {
		return map[bool]int32{}
	}

	newMap := make(map[bool]int32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeBoolInt16 takes two inputs: map[bool]int16 and map[bool]int16 and merge two maps and returns a new map[bool]int16.
func MergeBoolInt16(map1, map2 map[bool]int16) map[bool]int16 {
	if map1 == nil && map2 == nil {
		return map[bool]int16{}
	}

	newMap := make(map[bool]int16)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeBoolInt8 takes two inputs: map[bool]int8 and map[bool]int8 and merge two maps and returns a new map[bool]int8.
func MergeBoolInt8(map1, map2 map[bool]int8) map[bool]int8 {
	if map1 == nil && map2 == nil {
		return map[bool]int8{}
	}

	newMap := make(map[bool]int8)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeBoolUint takes two inputs: map[bool]uint and map[bool]uint and merge two maps and returns a new map[bool]uint.
func MergeBoolUint(map1, map2 map[bool]uint) map[bool]uint {
	if map1 == nil && map2 == nil {
		return map[bool]uint{}
	}

	newMap := make(map[bool]uint)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeBoolUint64 takes two inputs: map[bool]uint64 and map[bool]uint64 and merge two maps and returns a new map[bool]uint64.
func MergeBoolUint64(map1, map2 map[bool]uint64) map[bool]uint64 {
	if map1 == nil && map2 == nil {
		return map[bool]uint64{}
	}

	newMap := make(map[bool]uint64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeBoolUint32 takes two inputs: map[bool]uint32 and map[bool]uint32 and merge two maps and returns a new map[bool]uint32.
func MergeBoolUint32(map1, map2 map[bool]uint32) map[bool]uint32 {
	if map1 == nil && map2 == nil {
		return map[bool]uint32{}
	}

	newMap := make(map[bool]uint32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeBoolUint16 takes two inputs: map[bool]uint16 and map[bool]uint16 and merge two maps and returns a new map[bool]uint16.
func MergeBoolUint16(map1, map2 map[bool]uint16) map[bool]uint16 {
	if map1 == nil && map2 == nil {
		return map[bool]uint16{}
	}

	newMap := make(map[bool]uint16)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeBoolUint8 takes two inputs: map[bool]uint8 and map[bool]uint8 and merge two maps and returns a new map[bool]uint8.
func MergeBoolUint8(map1, map2 map[bool]uint8) map[bool]uint8 {
	if map1 == nil && map2 == nil {
		return map[bool]uint8{}
	}

	newMap := make(map[bool]uint8)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeBoolStr takes two inputs: map[bool]string and map[bool]string and merge two maps and returns a new map[bool]string.
func MergeBoolStr(map1, map2 map[bool]string) map[bool]string {
	if map1 == nil && map2 == nil {
		return map[bool]string{}
	}

	newMap := make(map[bool]string)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeBool takes two inputs: map[bool]bool and map[bool]bool and merge two maps and returns a new map[bool]bool.
func MergeBool(map1, map2 map[bool]bool) map[bool]bool {
	if map1 == nil && map2 == nil {
		return map[bool]bool{}
	}

	newMap := make(map[bool]bool)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeBoolFloat32 takes two inputs: map[bool]float32 and map[bool]float32 and merge two maps and returns a new map[bool]float32.
func MergeBoolFloat32(map1, map2 map[bool]float32) map[bool]float32 {
	if map1 == nil && map2 == nil {
		return map[bool]float32{}
	}

	newMap := make(map[bool]float32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeBoolFloat64 takes two inputs: map[bool]float64 and map[bool]float64 and merge two maps and returns a new map[bool]float64.
func MergeBoolFloat64(map1, map2 map[bool]float64) map[bool]float64 {
	if map1 == nil && map2 == nil {
		return map[bool]float64{}
	}

	newMap := make(map[bool]float64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeFloat32Int takes two inputs: map[float32]int and map[float32]int and merge two maps and returns a new map[float32]int.
func MergeFloat32Int(map1, map2 map[float32]int) map[float32]int {
	if map1 == nil && map2 == nil {
		return map[float32]int{}
	}

	newMap := make(map[float32]int)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeFloat32Int64 takes two inputs: map[float32]int64 and map[float32]int64 and merge two maps and returns a new map[float32]int64.
func MergeFloat32Int64(map1, map2 map[float32]int64) map[float32]int64 {
	if map1 == nil && map2 == nil {
		return map[float32]int64{}
	}

	newMap := make(map[float32]int64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeFloat32Int32 takes two inputs: map[float32]int32 and map[float32]int32 and merge two maps and returns a new map[float32]int32.
func MergeFloat32Int32(map1, map2 map[float32]int32) map[float32]int32 {
	if map1 == nil && map2 == nil {
		return map[float32]int32{}
	}

	newMap := make(map[float32]int32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeFloat32Int16 takes two inputs: map[float32]int16 and map[float32]int16 and merge two maps and returns a new map[float32]int16.
func MergeFloat32Int16(map1, map2 map[float32]int16) map[float32]int16 {
	if map1 == nil && map2 == nil {
		return map[float32]int16{}
	}

	newMap := make(map[float32]int16)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeFloat32Int8 takes two inputs: map[float32]int8 and map[float32]int8 and merge two maps and returns a new map[float32]int8.
func MergeFloat32Int8(map1, map2 map[float32]int8) map[float32]int8 {
	if map1 == nil && map2 == nil {
		return map[float32]int8{}
	}

	newMap := make(map[float32]int8)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeFloat32Uint takes two inputs: map[float32]uint and map[float32]uint and merge two maps and returns a new map[float32]uint.
func MergeFloat32Uint(map1, map2 map[float32]uint) map[float32]uint {
	if map1 == nil && map2 == nil {
		return map[float32]uint{}
	}

	newMap := make(map[float32]uint)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeFloat32Uint64 takes two inputs: map[float32]uint64 and map[float32]uint64 and merge two maps and returns a new map[float32]uint64.
func MergeFloat32Uint64(map1, map2 map[float32]uint64) map[float32]uint64 {
	if map1 == nil && map2 == nil {
		return map[float32]uint64{}
	}

	newMap := make(map[float32]uint64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeFloat32Uint32 takes two inputs: map[float32]uint32 and map[float32]uint32 and merge two maps and returns a new map[float32]uint32.
func MergeFloat32Uint32(map1, map2 map[float32]uint32) map[float32]uint32 {
	if map1 == nil && map2 == nil {
		return map[float32]uint32{}
	}

	newMap := make(map[float32]uint32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeFloat32Uint16 takes two inputs: map[float32]uint16 and map[float32]uint16 and merge two maps and returns a new map[float32]uint16.
func MergeFloat32Uint16(map1, map2 map[float32]uint16) map[float32]uint16 {
	if map1 == nil && map2 == nil {
		return map[float32]uint16{}
	}

	newMap := make(map[float32]uint16)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeFloat32Uint8 takes two inputs: map[float32]uint8 and map[float32]uint8 and merge two maps and returns a new map[float32]uint8.
func MergeFloat32Uint8(map1, map2 map[float32]uint8) map[float32]uint8 {
	if map1 == nil && map2 == nil {
		return map[float32]uint8{}
	}

	newMap := make(map[float32]uint8)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeFloat32Str takes two inputs: map[float32]string and map[float32]string and merge two maps and returns a new map[float32]string.
func MergeFloat32Str(map1, map2 map[float32]string) map[float32]string {
	if map1 == nil && map2 == nil {
		return map[float32]string{}
	}

	newMap := make(map[float32]string)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeFloat32Bool takes two inputs: map[float32]bool and map[float32]bool and merge two maps and returns a new map[float32]bool.
func MergeFloat32Bool(map1, map2 map[float32]bool) map[float32]bool {
	if map1 == nil && map2 == nil {
		return map[float32]bool{}
	}

	newMap := make(map[float32]bool)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeFloat32 takes two inputs: map[float32]float32 and map[float32]float32 and merge two maps and returns a new map[float32]float32.
func MergeFloat32(map1, map2 map[float32]float32) map[float32]float32 {
	if map1 == nil && map2 == nil {
		return map[float32]float32{}
	}

	newMap := make(map[float32]float32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeFloat32Float64 takes two inputs: map[float32]float64 and map[float32]float64 and merge two maps and returns a new map[float32]float64.
func MergeFloat32Float64(map1, map2 map[float32]float64) map[float32]float64 {
	if map1 == nil && map2 == nil {
		return map[float32]float64{}
	}

	newMap := make(map[float32]float64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeFloat64Int takes two inputs: map[float64]int and map[float64]int and merge two maps and returns a new map[float64]int.
func MergeFloat64Int(map1, map2 map[float64]int) map[float64]int {
	if map1 == nil && map2 == nil {
		return map[float64]int{}
	}

	newMap := make(map[float64]int)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeFloat64Int64 takes two inputs: map[float64]int64 and map[float64]int64 and merge two maps and returns a new map[float64]int64.
func MergeFloat64Int64(map1, map2 map[float64]int64) map[float64]int64 {
	if map1 == nil && map2 == nil {
		return map[float64]int64{}
	}

	newMap := make(map[float64]int64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeFloat64Int32 takes two inputs: map[float64]int32 and map[float64]int32 and merge two maps and returns a new map[float64]int32.
func MergeFloat64Int32(map1, map2 map[float64]int32) map[float64]int32 {
	if map1 == nil && map2 == nil {
		return map[float64]int32{}
	}

	newMap := make(map[float64]int32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeFloat64Int16 takes two inputs: map[float64]int16 and map[float64]int16 and merge two maps and returns a new map[float64]int16.
func MergeFloat64Int16(map1, map2 map[float64]int16) map[float64]int16 {
	if map1 == nil && map2 == nil {
		return map[float64]int16{}
	}

	newMap := make(map[float64]int16)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeFloat64Int8 takes two inputs: map[float64]int8 and map[float64]int8 and merge two maps and returns a new map[float64]int8.
func MergeFloat64Int8(map1, map2 map[float64]int8) map[float64]int8 {
	if map1 == nil && map2 == nil {
		return map[float64]int8{}
	}

	newMap := make(map[float64]int8)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeFloat64Uint takes two inputs: map[float64]uint and map[float64]uint and merge two maps and returns a new map[float64]uint.
func MergeFloat64Uint(map1, map2 map[float64]uint) map[float64]uint {
	if map1 == nil && map2 == nil {
		return map[float64]uint{}
	}

	newMap := make(map[float64]uint)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeFloat64Uint64 takes two inputs: map[float64]uint64 and map[float64]uint64 and merge two maps and returns a new map[float64]uint64.
func MergeFloat64Uint64(map1, map2 map[float64]uint64) map[float64]uint64 {
	if map1 == nil && map2 == nil {
		return map[float64]uint64{}
	}

	newMap := make(map[float64]uint64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeFloat64Uint32 takes two inputs: map[float64]uint32 and map[float64]uint32 and merge two maps and returns a new map[float64]uint32.
func MergeFloat64Uint32(map1, map2 map[float64]uint32) map[float64]uint32 {
	if map1 == nil && map2 == nil {
		return map[float64]uint32{}
	}

	newMap := make(map[float64]uint32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeFloat64Uint16 takes two inputs: map[float64]uint16 and map[float64]uint16 and merge two maps and returns a new map[float64]uint16.
func MergeFloat64Uint16(map1, map2 map[float64]uint16) map[float64]uint16 {
	if map1 == nil && map2 == nil {
		return map[float64]uint16{}
	}

	newMap := make(map[float64]uint16)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeFloat64Uint8 takes two inputs: map[float64]uint8 and map[float64]uint8 and merge two maps and returns a new map[float64]uint8.
func MergeFloat64Uint8(map1, map2 map[float64]uint8) map[float64]uint8 {
	if map1 == nil && map2 == nil {
		return map[float64]uint8{}
	}

	newMap := make(map[float64]uint8)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeFloat64Str takes two inputs: map[float64]string and map[float64]string and merge two maps and returns a new map[float64]string.
func MergeFloat64Str(map1, map2 map[float64]string) map[float64]string {
	if map1 == nil && map2 == nil {
		return map[float64]string{}
	}

	newMap := make(map[float64]string)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeFloat64Bool takes two inputs: map[float64]bool and map[float64]bool and merge two maps and returns a new map[float64]bool.
func MergeFloat64Bool(map1, map2 map[float64]bool) map[float64]bool {
	if map1 == nil && map2 == nil {
		return map[float64]bool{}
	}

	newMap := make(map[float64]bool)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeFloat64Float32 takes two inputs: map[float64]float32 and map[float64]float32 and merge two maps and returns a new map[float64]float32.
func MergeFloat64Float32(map1, map2 map[float64]float32) map[float64]float32 {
	if map1 == nil && map2 == nil {
		return map[float64]float32{}
	}

	newMap := make(map[float64]float32)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}

// MergeFloat64 takes two inputs: map[float64]float64 and map[float64]float64 and merge two maps and returns a new map[float64]float64.
func MergeFloat64(map1, map2 map[float64]float64) map[float64]float64 {
	if map1 == nil && map2 == nil {
		return map[float64]float64{}
	}

	newMap := make(map[float64]float64)

	if map1 == nil {
		for k, v := range map2 {
			newMap[k] = v
		}
		return newMap
	}

	if map2 == nil {
		for k, v := range map1 {
			newMap[k] = v
		}
		return newMap
	}

	for k, v := range map1 {
		newMap[k] = v
	}

	for k, v := range map2 {
		newMap[k] = v
	}

	return newMap
}
