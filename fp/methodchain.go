package fp

type intSlice []int
type intFunctor func(int) int

// MakeIntSlice - creates slice for the functional method such as map, filter
func MakeIntSlice(values ...int) intSlice {
	newSlice := intSlice(values)
	return newSlice
}

func mapCoreint(f intFunctor, slice intSlice) intSlice {
	newSlice := make(intSlice, len(slice))
	for i, v := range slice {
		newSlice[i] = f(v)
	}
	return newSlice
}

// Map - applies the function(1st argument) on each item of the list and returns new list
func (slice intSlice) Map(functors ...intFunctor) intSlice {

	tmpSlice := slice
	
	for _, f := range functors {
		tmpSlice = mapCoreint(f, tmpSlice)
	}

	return tmpSlice
}

type int64Slice []int64
type int64Functor func(int64) int64

// MakeInt64Slice - creates slice for the functional method such as map, filter
func MakeInt64Slice(values ...int64) int64Slice {
	newSlice := int64Slice(values)
	return newSlice
}

func mapCoreint64(f int64Functor, slice int64Slice) int64Slice {
	newSlice := make(int64Slice, len(slice))
	for i, v := range slice {
		newSlice[i] = f(v)
	}
	return newSlice
}

// Map - applies the function(1st argument) on each item of the list and returns new list
func (slice int64Slice) Map(functors ...int64Functor) int64Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		tmpSlice = mapCoreint64(f, tmpSlice)
	}

	return tmpSlice
}

type int32Slice []int32
type int32Functor func(int32) int32

// MakeInt32Slice - creates slice for the functional method such as map, filter
func MakeInt32Slice(values ...int32) int32Slice {
	newSlice := int32Slice(values)
	return newSlice
}

func mapCoreint32(f int32Functor, slice int32Slice) int32Slice {
	newSlice := make(int32Slice, len(slice))
	for i, v := range slice {
		newSlice[i] = f(v)
	}
	return newSlice
}

// Map - applies the function(1st argument) on each item of the list and returns new list
func (slice int32Slice) Map(functors ...int32Functor) int32Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		tmpSlice = mapCoreint32(f, tmpSlice)
	}

	return tmpSlice
}

type int16Slice []int16
type int16Functor func(int16) int16

// MakeInt16Slice - creates slice for the functional method such as map, filter
func MakeInt16Slice(values ...int16) int16Slice {
	newSlice := int16Slice(values)
	return newSlice
}

func mapCoreint16(f int16Functor, slice int16Slice) int16Slice {
	newSlice := make(int16Slice, len(slice))
	for i, v := range slice {
		newSlice[i] = f(v)
	}
	return newSlice
}

// Map - applies the function(1st argument) on each item of the list and returns new list
func (slice int16Slice) Map(functors ...int16Functor) int16Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		tmpSlice = mapCoreint16(f, tmpSlice)
	}

	return tmpSlice
}

type int8Slice []int8
type int8Functor func(int8) int8

// MakeInt8Slice - creates slice for the functional method such as map, filter
func MakeInt8Slice(values ...int8) int8Slice {
	newSlice := int8Slice(values)
	return newSlice
}

func mapCoreint8(f int8Functor, slice int8Slice) int8Slice {
	newSlice := make(int8Slice, len(slice))
	for i, v := range slice {
		newSlice[i] = f(v)
	}
	return newSlice
}

// Map - applies the function(1st argument) on each item of the list and returns new list
func (slice int8Slice) Map(functors ...int8Functor) int8Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		tmpSlice = mapCoreint8(f, tmpSlice)
	}

	return tmpSlice
}

type uintSlice []uint
type uintFunctor func(uint) uint

// MakeUintSlice - creates slice for the functional method such as map, filter
func MakeUintSlice(values ...uint) uintSlice {
	newSlice := uintSlice(values)
	return newSlice
}

func mapCoreuint(f uintFunctor, slice uintSlice) uintSlice {
	newSlice := make(uintSlice, len(slice))
	for i, v := range slice {
		newSlice[i] = f(v)
	}
	return newSlice
}

// Map - applies the function(1st argument) on each item of the list and returns new list
func (slice uintSlice) Map(functors ...uintFunctor) uintSlice {

	tmpSlice := slice
	
	for _, f := range functors {
		tmpSlice = mapCoreuint(f, tmpSlice)
	}

	return tmpSlice
}

type uint64Slice []uint64
type uint64Functor func(uint64) uint64

// MakeUint64Slice - creates slice for the functional method such as map, filter
func MakeUint64Slice(values ...uint64) uint64Slice {
	newSlice := uint64Slice(values)
	return newSlice
}

func mapCoreuint64(f uint64Functor, slice uint64Slice) uint64Slice {
	newSlice := make(uint64Slice, len(slice))
	for i, v := range slice {
		newSlice[i] = f(v)
	}
	return newSlice
}

// Map - applies the function(1st argument) on each item of the list and returns new list
func (slice uint64Slice) Map(functors ...uint64Functor) uint64Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		tmpSlice = mapCoreuint64(f, tmpSlice)
	}

	return tmpSlice
}

type uint32Slice []uint32
type uint32Functor func(uint32) uint32

// MakeUint32Slice - creates slice for the functional method such as map, filter
func MakeUint32Slice(values ...uint32) uint32Slice {
	newSlice := uint32Slice(values)
	return newSlice
}

func mapCoreuint32(f uint32Functor, slice uint32Slice) uint32Slice {
	newSlice := make(uint32Slice, len(slice))
	for i, v := range slice {
		newSlice[i] = f(v)
	}
	return newSlice
}

// Map - applies the function(1st argument) on each item of the list and returns new list
func (slice uint32Slice) Map(functors ...uint32Functor) uint32Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		tmpSlice = mapCoreuint32(f, tmpSlice)
	}

	return tmpSlice
}

type uint16Slice []uint16
type uint16Functor func(uint16) uint16

// MakeUint16Slice - creates slice for the functional method such as map, filter
func MakeUint16Slice(values ...uint16) uint16Slice {
	newSlice := uint16Slice(values)
	return newSlice
}

func mapCoreuint16(f uint16Functor, slice uint16Slice) uint16Slice {
	newSlice := make(uint16Slice, len(slice))
	for i, v := range slice {
		newSlice[i] = f(v)
	}
	return newSlice
}

// Map - applies the function(1st argument) on each item of the list and returns new list
func (slice uint16Slice) Map(functors ...uint16Functor) uint16Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		tmpSlice = mapCoreuint16(f, tmpSlice)
	}

	return tmpSlice
}

type uint8Slice []uint8
type uint8Functor func(uint8) uint8

// MakeUint8Slice - creates slice for the functional method such as map, filter
func MakeUint8Slice(values ...uint8) uint8Slice {
	newSlice := uint8Slice(values)
	return newSlice
}

func mapCoreuint8(f uint8Functor, slice uint8Slice) uint8Slice {
	newSlice := make(uint8Slice, len(slice))
	for i, v := range slice {
		newSlice[i] = f(v)
	}
	return newSlice
}

// Map - applies the function(1st argument) on each item of the list and returns new list
func (slice uint8Slice) Map(functors ...uint8Functor) uint8Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		tmpSlice = mapCoreuint8(f, tmpSlice)
	}

	return tmpSlice
}

type stringSlice []string
type stringFunctor func(string) string

// MakeStrSlice - creates slice for the functional method such as map, filter
func MakeStrSlice(values ...string) stringSlice {
	newSlice := stringSlice(values)
	return newSlice
}

func mapCorestring(f stringFunctor, slice stringSlice) stringSlice {
	newSlice := make(stringSlice, len(slice))
	for i, v := range slice {
		newSlice[i] = f(v)
	}
	return newSlice
}

// Map - applies the function(1st argument) on each item of the list and returns new list
func (slice stringSlice) Map(functors ...stringFunctor) stringSlice {

	tmpSlice := slice
	
	for _, f := range functors {
		tmpSlice = mapCorestring(f, tmpSlice)
	}

	return tmpSlice
}

type boolSlice []bool
type boolFunctor func(bool) bool

// MakeBoolSlice - creates slice for the functional method such as map, filter
func MakeBoolSlice(values ...bool) boolSlice {
	newSlice := boolSlice(values)
	return newSlice
}

func mapCorebool(f boolFunctor, slice boolSlice) boolSlice {
	newSlice := make(boolSlice, len(slice))
	for i, v := range slice {
		newSlice[i] = f(v)
	}
	return newSlice
}

// Map - applies the function(1st argument) on each item of the list and returns new list
func (slice boolSlice) Map(functors ...boolFunctor) boolSlice {

	tmpSlice := slice
	
	for _, f := range functors {
		tmpSlice = mapCorebool(f, tmpSlice)
	}

	return tmpSlice
}

type float32Slice []float32
type float32Functor func(float32) float32

// MakeFloat32Slice - creates slice for the functional method such as map, filter
func MakeFloat32Slice(values ...float32) float32Slice {
	newSlice := float32Slice(values)
	return newSlice
}

func mapCorefloat32(f float32Functor, slice float32Slice) float32Slice {
	newSlice := make(float32Slice, len(slice))
	for i, v := range slice {
		newSlice[i] = f(v)
	}
	return newSlice
}

// Map - applies the function(1st argument) on each item of the list and returns new list
func (slice float32Slice) Map(functors ...float32Functor) float32Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		tmpSlice = mapCorefloat32(f, tmpSlice)
	}

	return tmpSlice
}

type float64Slice []float64
type float64Functor func(float64) float64

// MakeFloat64Slice - creates slice for the functional method such as map, filter
func MakeFloat64Slice(values ...float64) float64Slice {
	newSlice := float64Slice(values)
	return newSlice
}

func mapCorefloat64(f float64Functor, slice float64Slice) float64Slice {
	newSlice := make(float64Slice, len(slice))
	for i, v := range slice {
		newSlice[i] = f(v)
	}
	return newSlice
}

// Map - applies the function(1st argument) on each item of the list and returns new list
func (slice float64Slice) Map(functors ...float64Functor) float64Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		tmpSlice = mapCorefloat64(f, tmpSlice)
	}

	return tmpSlice
}
