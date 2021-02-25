package fp

type intSlice []int
type intFunctorForMap func(int) int
type intFunctorForFilter func(int) bool

type intSlicePtr []*int
type intFunctorForMapPtr func(*int) *int
type intFunctorForFilterPtr func(*int) bool

// MakeIntSlice - creates slice for the functional method such as map, filter
func MakeIntSlice(values ...int) intSlice {
	newSlice := intSlice(values)
	return newSlice
}

// Map - applies the function(1st argument) on each item of the list and returns new list
func (slice intSlice) Map(functors ...intFunctorForMap) intSlice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = MapInt(f, tmpSlice)
	}

	return tmpSlice
}

// MakeIntSlicePtr - creates slice for the functional method such as map, filter
func MakeIntSlicePtr(values ...*int) intSlicePtr {
	newSlice := intSlicePtr(values)
	return newSlice
}

// MapPtr - applies the function(1st argument) on each item of the list and returns new list
func (slice intSlicePtr) MapPtr(functors ...intFunctorForMapPtr) intSlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = MapIntPtr(f, tmpSlice)
	}

	return tmpSlice
}

// Filter - filters list based on function passed as argument
func (slice intSlice) Filter(functors ...intFunctorForFilter) intSlice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = FilterInt(f, tmpSlice)
	}

	return tmpSlice
}

// FilterPtr - filters list based on function passed as argument
func (slice intSlicePtr) FilterPtr(functors ...intFunctorForFilterPtr) intSlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = FilterIntPtr(f, tmpSlice)
	}

	return tmpSlice
}

// Remove - removes the items from the given list based on supplied function and returns new list
func (slice intSlice) Remove(functors ...intFunctorForFilter) intSlice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = RemoveInt(f, tmpSlice)
	}

	return tmpSlice
}

// RemovePtr - removes the items from the given list based on supplied function and returns new list
func (slice intSlicePtr) RemovePtr(functors ...intFunctorForFilterPtr) intSlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = RemoveIntPtr(f, tmpSlice)
	}

	return tmpSlice
}

type int64Slice []int64
type int64FunctorForMap func(int64) int64
type int64FunctorForFilter func(int64) bool

type int64SlicePtr []*int64
type int64FunctorForMapPtr func(*int64) *int64
type int64FunctorForFilterPtr func(*int64) bool

// MakeInt64Slice - creates slice for the functional method such as map, filter
func MakeInt64Slice(values ...int64) int64Slice {
	newSlice := int64Slice(values)
	return newSlice
}

// Map - applies the function(1st argument) on each item of the list and returns new list
func (slice int64Slice) Map(functors ...int64FunctorForMap) int64Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = MapInt64(f, tmpSlice)
	}

	return tmpSlice
}

// MakeInt64SlicePtr - creates slice for the functional method such as map, filter
func MakeInt64SlicePtr(values ...*int64) int64SlicePtr {
	newSlice := int64SlicePtr(values)
	return newSlice
}

// MapPtr - applies the function(1st argument) on each item of the list and returns new list
func (slice int64SlicePtr) MapPtr(functors ...int64FunctorForMapPtr) int64SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = MapInt64Ptr(f, tmpSlice)
	}

	return tmpSlice
}

// Filter - filters list based on function passed as argument
func (slice int64Slice) Filter(functors ...int64FunctorForFilter) int64Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = FilterInt64(f, tmpSlice)
	}

	return tmpSlice
}

// FilterPtr - filters list based on function passed as argument
func (slice int64SlicePtr) FilterPtr(functors ...int64FunctorForFilterPtr) int64SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = FilterInt64Ptr(f, tmpSlice)
	}

	return tmpSlice
}

// Remove - removes the items from the given list based on supplied function and returns new list
func (slice int64Slice) Remove(functors ...int64FunctorForFilter) int64Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = RemoveInt64(f, tmpSlice)
	}

	return tmpSlice
}

// RemovePtr - removes the items from the given list based on supplied function and returns new list
func (slice int64SlicePtr) RemovePtr(functors ...int64FunctorForFilterPtr) int64SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = RemoveInt64Ptr(f, tmpSlice)
	}

	return tmpSlice
}

type int32Slice []int32
type int32FunctorForMap func(int32) int32
type int32FunctorForFilter func(int32) bool

type int32SlicePtr []*int32
type int32FunctorForMapPtr func(*int32) *int32
type int32FunctorForFilterPtr func(*int32) bool

// MakeInt32Slice - creates slice for the functional method such as map, filter
func MakeInt32Slice(values ...int32) int32Slice {
	newSlice := int32Slice(values)
	return newSlice
}

// Map - applies the function(1st argument) on each item of the list and returns new list
func (slice int32Slice) Map(functors ...int32FunctorForMap) int32Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = MapInt32(f, tmpSlice)
	}

	return tmpSlice
}

// MakeInt32SlicePtr - creates slice for the functional method such as map, filter
func MakeInt32SlicePtr(values ...*int32) int32SlicePtr {
	newSlice := int32SlicePtr(values)
	return newSlice
}

// MapPtr - applies the function(1st argument) on each item of the list and returns new list
func (slice int32SlicePtr) MapPtr(functors ...int32FunctorForMapPtr) int32SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = MapInt32Ptr(f, tmpSlice)
	}

	return tmpSlice
}

// Filter - filters list based on function passed as argument
func (slice int32Slice) Filter(functors ...int32FunctorForFilter) int32Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = FilterInt32(f, tmpSlice)
	}

	return tmpSlice
}

// FilterPtr - filters list based on function passed as argument
func (slice int32SlicePtr) FilterPtr(functors ...int32FunctorForFilterPtr) int32SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = FilterInt32Ptr(f, tmpSlice)
	}

	return tmpSlice
}

// Remove - removes the items from the given list based on supplied function and returns new list
func (slice int32Slice) Remove(functors ...int32FunctorForFilter) int32Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = RemoveInt32(f, tmpSlice)
	}

	return tmpSlice
}

// RemovePtr - removes the items from the given list based on supplied function and returns new list
func (slice int32SlicePtr) RemovePtr(functors ...int32FunctorForFilterPtr) int32SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = RemoveInt32Ptr(f, tmpSlice)
	}

	return tmpSlice
}

type int16Slice []int16
type int16FunctorForMap func(int16) int16
type int16FunctorForFilter func(int16) bool

type int16SlicePtr []*int16
type int16FunctorForMapPtr func(*int16) *int16
type int16FunctorForFilterPtr func(*int16) bool

// MakeInt16Slice - creates slice for the functional method such as map, filter
func MakeInt16Slice(values ...int16) int16Slice {
	newSlice := int16Slice(values)
	return newSlice
}

// Map - applies the function(1st argument) on each item of the list and returns new list
func (slice int16Slice) Map(functors ...int16FunctorForMap) int16Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = MapInt16(f, tmpSlice)
	}

	return tmpSlice
}

// MakeInt16SlicePtr - creates slice for the functional method such as map, filter
func MakeInt16SlicePtr(values ...*int16) int16SlicePtr {
	newSlice := int16SlicePtr(values)
	return newSlice
}

// MapPtr - applies the function(1st argument) on each item of the list and returns new list
func (slice int16SlicePtr) MapPtr(functors ...int16FunctorForMapPtr) int16SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = MapInt16Ptr(f, tmpSlice)
	}

	return tmpSlice
}

// Filter - filters list based on function passed as argument
func (slice int16Slice) Filter(functors ...int16FunctorForFilter) int16Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = FilterInt16(f, tmpSlice)
	}

	return tmpSlice
}

// FilterPtr - filters list based on function passed as argument
func (slice int16SlicePtr) FilterPtr(functors ...int16FunctorForFilterPtr) int16SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = FilterInt16Ptr(f, tmpSlice)
	}

	return tmpSlice
}

// Remove - removes the items from the given list based on supplied function and returns new list
func (slice int16Slice) Remove(functors ...int16FunctorForFilter) int16Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = RemoveInt16(f, tmpSlice)
	}

	return tmpSlice
}

// RemovePtr - removes the items from the given list based on supplied function and returns new list
func (slice int16SlicePtr) RemovePtr(functors ...int16FunctorForFilterPtr) int16SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = RemoveInt16Ptr(f, tmpSlice)
	}

	return tmpSlice
}

type int8Slice []int8
type int8FunctorForMap func(int8) int8
type int8FunctorForFilter func(int8) bool

type int8SlicePtr []*int8
type int8FunctorForMapPtr func(*int8) *int8
type int8FunctorForFilterPtr func(*int8) bool

// MakeInt8Slice - creates slice for the functional method such as map, filter
func MakeInt8Slice(values ...int8) int8Slice {
	newSlice := int8Slice(values)
	return newSlice
}

// Map - applies the function(1st argument) on each item of the list and returns new list
func (slice int8Slice) Map(functors ...int8FunctorForMap) int8Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = MapInt8(f, tmpSlice)
	}

	return tmpSlice
}

// MakeInt8SlicePtr - creates slice for the functional method such as map, filter
func MakeInt8SlicePtr(values ...*int8) int8SlicePtr {
	newSlice := int8SlicePtr(values)
	return newSlice
}

// MapPtr - applies the function(1st argument) on each item of the list and returns new list
func (slice int8SlicePtr) MapPtr(functors ...int8FunctorForMapPtr) int8SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = MapInt8Ptr(f, tmpSlice)
	}

	return tmpSlice
}

// Filter - filters list based on function passed as argument
func (slice int8Slice) Filter(functors ...int8FunctorForFilter) int8Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = FilterInt8(f, tmpSlice)
	}

	return tmpSlice
}

// FilterPtr - filters list based on function passed as argument
func (slice int8SlicePtr) FilterPtr(functors ...int8FunctorForFilterPtr) int8SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = FilterInt8Ptr(f, tmpSlice)
	}

	return tmpSlice
}

// Remove - removes the items from the given list based on supplied function and returns new list
func (slice int8Slice) Remove(functors ...int8FunctorForFilter) int8Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = RemoveInt8(f, tmpSlice)
	}

	return tmpSlice
}

// RemovePtr - removes the items from the given list based on supplied function and returns new list
func (slice int8SlicePtr) RemovePtr(functors ...int8FunctorForFilterPtr) int8SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = RemoveInt8Ptr(f, tmpSlice)
	}

	return tmpSlice
}

type uintSlice []uint
type uintFunctorForMap func(uint) uint
type uintFunctorForFilter func(uint) bool

type uintSlicePtr []*uint
type uintFunctorForMapPtr func(*uint) *uint
type uintFunctorForFilterPtr func(*uint) bool

// MakeUintSlice - creates slice for the functional method such as map, filter
func MakeUintSlice(values ...uint) uintSlice {
	newSlice := uintSlice(values)
	return newSlice
}

// Map - applies the function(1st argument) on each item of the list and returns new list
func (slice uintSlice) Map(functors ...uintFunctorForMap) uintSlice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = MapUint(f, tmpSlice)
	}

	return tmpSlice
}

// MakeUintSlicePtr - creates slice for the functional method such as map, filter
func MakeUintSlicePtr(values ...*uint) uintSlicePtr {
	newSlice := uintSlicePtr(values)
	return newSlice
}

// MapPtr - applies the function(1st argument) on each item of the list and returns new list
func (slice uintSlicePtr) MapPtr(functors ...uintFunctorForMapPtr) uintSlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = MapUintPtr(f, tmpSlice)
	}

	return tmpSlice
}

// Filter - filters list based on function passed as argument
func (slice uintSlice) Filter(functors ...uintFunctorForFilter) uintSlice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = FilterUint(f, tmpSlice)
	}

	return tmpSlice
}

// FilterPtr - filters list based on function passed as argument
func (slice uintSlicePtr) FilterPtr(functors ...uintFunctorForFilterPtr) uintSlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = FilterUintPtr(f, tmpSlice)
	}

	return tmpSlice
}

// Remove - removes the items from the given list based on supplied function and returns new list
func (slice uintSlice) Remove(functors ...uintFunctorForFilter) uintSlice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = RemoveUint(f, tmpSlice)
	}

	return tmpSlice
}

// RemovePtr - removes the items from the given list based on supplied function and returns new list
func (slice uintSlicePtr) RemovePtr(functors ...uintFunctorForFilterPtr) uintSlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = RemoveUintPtr(f, tmpSlice)
	}

	return tmpSlice
}

type uint64Slice []uint64
type uint64FunctorForMap func(uint64) uint64
type uint64FunctorForFilter func(uint64) bool

type uint64SlicePtr []*uint64
type uint64FunctorForMapPtr func(*uint64) *uint64
type uint64FunctorForFilterPtr func(*uint64) bool

// MakeUint64Slice - creates slice for the functional method such as map, filter
func MakeUint64Slice(values ...uint64) uint64Slice {
	newSlice := uint64Slice(values)
	return newSlice
}

// Map - applies the function(1st argument) on each item of the list and returns new list
func (slice uint64Slice) Map(functors ...uint64FunctorForMap) uint64Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = MapUint64(f, tmpSlice)
	}

	return tmpSlice
}

// MakeUint64SlicePtr - creates slice for the functional method such as map, filter
func MakeUint64SlicePtr(values ...*uint64) uint64SlicePtr {
	newSlice := uint64SlicePtr(values)
	return newSlice
}

// MapPtr - applies the function(1st argument) on each item of the list and returns new list
func (slice uint64SlicePtr) MapPtr(functors ...uint64FunctorForMapPtr) uint64SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = MapUint64Ptr(f, tmpSlice)
	}

	return tmpSlice
}

// Filter - filters list based on function passed as argument
func (slice uint64Slice) Filter(functors ...uint64FunctorForFilter) uint64Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = FilterUint64(f, tmpSlice)
	}

	return tmpSlice
}

// FilterPtr - filters list based on function passed as argument
func (slice uint64SlicePtr) FilterPtr(functors ...uint64FunctorForFilterPtr) uint64SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = FilterUint64Ptr(f, tmpSlice)
	}

	return tmpSlice
}

// Remove - removes the items from the given list based on supplied function and returns new list
func (slice uint64Slice) Remove(functors ...uint64FunctorForFilter) uint64Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = RemoveUint64(f, tmpSlice)
	}

	return tmpSlice
}

// RemovePtr - removes the items from the given list based on supplied function and returns new list
func (slice uint64SlicePtr) RemovePtr(functors ...uint64FunctorForFilterPtr) uint64SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = RemoveUint64Ptr(f, tmpSlice)
	}

	return tmpSlice
}

type uint32Slice []uint32
type uint32FunctorForMap func(uint32) uint32
type uint32FunctorForFilter func(uint32) bool

type uint32SlicePtr []*uint32
type uint32FunctorForMapPtr func(*uint32) *uint32
type uint32FunctorForFilterPtr func(*uint32) bool

// MakeUint32Slice - creates slice for the functional method such as map, filter
func MakeUint32Slice(values ...uint32) uint32Slice {
	newSlice := uint32Slice(values)
	return newSlice
}

// Map - applies the function(1st argument) on each item of the list and returns new list
func (slice uint32Slice) Map(functors ...uint32FunctorForMap) uint32Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = MapUint32(f, tmpSlice)
	}

	return tmpSlice
}

// MakeUint32SlicePtr - creates slice for the functional method such as map, filter
func MakeUint32SlicePtr(values ...*uint32) uint32SlicePtr {
	newSlice := uint32SlicePtr(values)
	return newSlice
}

// MapPtr - applies the function(1st argument) on each item of the list and returns new list
func (slice uint32SlicePtr) MapPtr(functors ...uint32FunctorForMapPtr) uint32SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = MapUint32Ptr(f, tmpSlice)
	}

	return tmpSlice
}

// Filter - filters list based on function passed as argument
func (slice uint32Slice) Filter(functors ...uint32FunctorForFilter) uint32Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = FilterUint32(f, tmpSlice)
	}

	return tmpSlice
}

// FilterPtr - filters list based on function passed as argument
func (slice uint32SlicePtr) FilterPtr(functors ...uint32FunctorForFilterPtr) uint32SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = FilterUint32Ptr(f, tmpSlice)
	}

	return tmpSlice
}

// Remove - removes the items from the given list based on supplied function and returns new list
func (slice uint32Slice) Remove(functors ...uint32FunctorForFilter) uint32Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = RemoveUint32(f, tmpSlice)
	}

	return tmpSlice
}

// RemovePtr - removes the items from the given list based on supplied function and returns new list
func (slice uint32SlicePtr) RemovePtr(functors ...uint32FunctorForFilterPtr) uint32SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = RemoveUint32Ptr(f, tmpSlice)
	}

	return tmpSlice
}

type uint16Slice []uint16
type uint16FunctorForMap func(uint16) uint16
type uint16FunctorForFilter func(uint16) bool

type uint16SlicePtr []*uint16
type uint16FunctorForMapPtr func(*uint16) *uint16
type uint16FunctorForFilterPtr func(*uint16) bool

// MakeUint16Slice - creates slice for the functional method such as map, filter
func MakeUint16Slice(values ...uint16) uint16Slice {
	newSlice := uint16Slice(values)
	return newSlice
}

// Map - applies the function(1st argument) on each item of the list and returns new list
func (slice uint16Slice) Map(functors ...uint16FunctorForMap) uint16Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = MapUint16(f, tmpSlice)
	}

	return tmpSlice
}

// MakeUint16SlicePtr - creates slice for the functional method such as map, filter
func MakeUint16SlicePtr(values ...*uint16) uint16SlicePtr {
	newSlice := uint16SlicePtr(values)
	return newSlice
}

// MapPtr - applies the function(1st argument) on each item of the list and returns new list
func (slice uint16SlicePtr) MapPtr(functors ...uint16FunctorForMapPtr) uint16SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = MapUint16Ptr(f, tmpSlice)
	}

	return tmpSlice
}

// Filter - filters list based on function passed as argument
func (slice uint16Slice) Filter(functors ...uint16FunctorForFilter) uint16Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = FilterUint16(f, tmpSlice)
	}

	return tmpSlice
}

// FilterPtr - filters list based on function passed as argument
func (slice uint16SlicePtr) FilterPtr(functors ...uint16FunctorForFilterPtr) uint16SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = FilterUint16Ptr(f, tmpSlice)
	}

	return tmpSlice
}

// Remove - removes the items from the given list based on supplied function and returns new list
func (slice uint16Slice) Remove(functors ...uint16FunctorForFilter) uint16Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = RemoveUint16(f, tmpSlice)
	}

	return tmpSlice
}

// RemovePtr - removes the items from the given list based on supplied function and returns new list
func (slice uint16SlicePtr) RemovePtr(functors ...uint16FunctorForFilterPtr) uint16SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = RemoveUint16Ptr(f, tmpSlice)
	}

	return tmpSlice
}

type uint8Slice []uint8
type uint8FunctorForMap func(uint8) uint8
type uint8FunctorForFilter func(uint8) bool

type uint8SlicePtr []*uint8
type uint8FunctorForMapPtr func(*uint8) *uint8
type uint8FunctorForFilterPtr func(*uint8) bool

// MakeUint8Slice - creates slice for the functional method such as map, filter
func MakeUint8Slice(values ...uint8) uint8Slice {
	newSlice := uint8Slice(values)
	return newSlice
}

// Map - applies the function(1st argument) on each item of the list and returns new list
func (slice uint8Slice) Map(functors ...uint8FunctorForMap) uint8Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = MapUint8(f, tmpSlice)
	}

	return tmpSlice
}

// MakeUint8SlicePtr - creates slice for the functional method such as map, filter
func MakeUint8SlicePtr(values ...*uint8) uint8SlicePtr {
	newSlice := uint8SlicePtr(values)
	return newSlice
}

// MapPtr - applies the function(1st argument) on each item of the list and returns new list
func (slice uint8SlicePtr) MapPtr(functors ...uint8FunctorForMapPtr) uint8SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = MapUint8Ptr(f, tmpSlice)
	}

	return tmpSlice
}

// Filter - filters list based on function passed as argument
func (slice uint8Slice) Filter(functors ...uint8FunctorForFilter) uint8Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = FilterUint8(f, tmpSlice)
	}

	return tmpSlice
}

// FilterPtr - filters list based on function passed as argument
func (slice uint8SlicePtr) FilterPtr(functors ...uint8FunctorForFilterPtr) uint8SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = FilterUint8Ptr(f, tmpSlice)
	}

	return tmpSlice
}

// Remove - removes the items from the given list based on supplied function and returns new list
func (slice uint8Slice) Remove(functors ...uint8FunctorForFilter) uint8Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = RemoveUint8(f, tmpSlice)
	}

	return tmpSlice
}

// RemovePtr - removes the items from the given list based on supplied function and returns new list
func (slice uint8SlicePtr) RemovePtr(functors ...uint8FunctorForFilterPtr) uint8SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = RemoveUint8Ptr(f, tmpSlice)
	}

	return tmpSlice
}

type stringSlice []string
type stringFunctorForMap func(string) string
type stringFunctorForFilter func(string) bool

type stringSlicePtr []*string
type stringFunctorForMapPtr func(*string) *string
type stringFunctorForFilterPtr func(*string) bool

// MakeStrSlice - creates slice for the functional method such as map, filter
func MakeStrSlice(values ...string) stringSlice {
	newSlice := stringSlice(values)
	return newSlice
}

// Map - applies the function(1st argument) on each item of the list and returns new list
func (slice stringSlice) Map(functors ...stringFunctorForMap) stringSlice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = MapStr(f, tmpSlice)
	}

	return tmpSlice
}

// MakeStrSlicePtr - creates slice for the functional method such as map, filter
func MakeStrSlicePtr(values ...*string) stringSlicePtr {
	newSlice := stringSlicePtr(values)
	return newSlice
}

// MapPtr - applies the function(1st argument) on each item of the list and returns new list
func (slice stringSlicePtr) MapPtr(functors ...stringFunctorForMapPtr) stringSlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = MapStrPtr(f, tmpSlice)
	}

	return tmpSlice
}

// Filter - filters list based on function passed as argument
func (slice stringSlice) Filter(functors ...stringFunctorForFilter) stringSlice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = FilterStr(f, tmpSlice)
	}

	return tmpSlice
}

// FilterPtr - filters list based on function passed as argument
func (slice stringSlicePtr) FilterPtr(functors ...stringFunctorForFilterPtr) stringSlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = FilterStrPtr(f, tmpSlice)
	}

	return tmpSlice
}

// Remove - removes the items from the given list based on supplied function and returns new list
func (slice stringSlice) Remove(functors ...stringFunctorForFilter) stringSlice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = RemoveStr(f, tmpSlice)
	}

	return tmpSlice
}

// RemovePtr - removes the items from the given list based on supplied function and returns new list
func (slice stringSlicePtr) RemovePtr(functors ...stringFunctorForFilterPtr) stringSlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = RemoveStrPtr(f, tmpSlice)
	}

	return tmpSlice
}

type boolSlice []bool
type boolFunctorForMap func(bool) bool
type boolFunctorForFilter func(bool) bool

type boolSlicePtr []*bool
type boolFunctorForMapPtr func(*bool) *bool
type boolFunctorForFilterPtr func(*bool) bool

// MakeBoolSlice - creates slice for the functional method such as map, filter
func MakeBoolSlice(values ...bool) boolSlice {
	newSlice := boolSlice(values)
	return newSlice
}

// Map - applies the function(1st argument) on each item of the list and returns new list
func (slice boolSlice) Map(functors ...boolFunctorForMap) boolSlice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = MapBool(f, tmpSlice)
	}

	return tmpSlice
}

// MakeBoolSlicePtr - creates slice for the functional method such as map, filter
func MakeBoolSlicePtr(values ...*bool) boolSlicePtr {
	newSlice := boolSlicePtr(values)
	return newSlice
}

// MapPtr - applies the function(1st argument) on each item of the list and returns new list
func (slice boolSlicePtr) MapPtr(functors ...boolFunctorForMapPtr) boolSlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = MapBoolPtr(f, tmpSlice)
	}

	return tmpSlice
}

// Filter - filters list based on function passed as argument
func (slice boolSlice) Filter(functors ...boolFunctorForFilter) boolSlice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = FilterBool(f, tmpSlice)
	}

	return tmpSlice
}

// FilterPtr - filters list based on function passed as argument
func (slice boolSlicePtr) FilterPtr(functors ...boolFunctorForFilterPtr) boolSlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = FilterBoolPtr(f, tmpSlice)
	}

	return tmpSlice
}

// Remove - removes the items from the given list based on supplied function and returns new list
func (slice boolSlice) Remove(functors ...boolFunctorForFilter) boolSlice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = RemoveBool(f, tmpSlice)
	}

	return tmpSlice
}

// RemovePtr - removes the items from the given list based on supplied function and returns new list
func (slice boolSlicePtr) RemovePtr(functors ...boolFunctorForFilterPtr) boolSlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = RemoveBoolPtr(f, tmpSlice)
	}

	return tmpSlice
}

type float32Slice []float32
type float32FunctorForMap func(float32) float32
type float32FunctorForFilter func(float32) bool

type float32SlicePtr []*float32
type float32FunctorForMapPtr func(*float32) *float32
type float32FunctorForFilterPtr func(*float32) bool

// MakeFloat32Slice - creates slice for the functional method such as map, filter
func MakeFloat32Slice(values ...float32) float32Slice {
	newSlice := float32Slice(values)
	return newSlice
}

// Map - applies the function(1st argument) on each item of the list and returns new list
func (slice float32Slice) Map(functors ...float32FunctorForMap) float32Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = MapFloat32(f, tmpSlice)
	}

	return tmpSlice
}

// MakeFloat32SlicePtr - creates slice for the functional method such as map, filter
func MakeFloat32SlicePtr(values ...*float32) float32SlicePtr {
	newSlice := float32SlicePtr(values)
	return newSlice
}

// MapPtr - applies the function(1st argument) on each item of the list and returns new list
func (slice float32SlicePtr) MapPtr(functors ...float32FunctorForMapPtr) float32SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = MapFloat32Ptr(f, tmpSlice)
	}

	return tmpSlice
}

// Filter - filters list based on function passed as argument
func (slice float32Slice) Filter(functors ...float32FunctorForFilter) float32Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = FilterFloat32(f, tmpSlice)
	}

	return tmpSlice
}

// FilterPtr - filters list based on function passed as argument
func (slice float32SlicePtr) FilterPtr(functors ...float32FunctorForFilterPtr) float32SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = FilterFloat32Ptr(f, tmpSlice)
	}

	return tmpSlice
}

// Remove - removes the items from the given list based on supplied function and returns new list
func (slice float32Slice) Remove(functors ...float32FunctorForFilter) float32Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = RemoveFloat32(f, tmpSlice)
	}

	return tmpSlice
}

// RemovePtr - removes the items from the given list based on supplied function and returns new list
func (slice float32SlicePtr) RemovePtr(functors ...float32FunctorForFilterPtr) float32SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = RemoveFloat32Ptr(f, tmpSlice)
	}

	return tmpSlice
}

type float64Slice []float64
type float64FunctorForMap func(float64) float64
type float64FunctorForFilter func(float64) bool

type float64SlicePtr []*float64
type float64FunctorForMapPtr func(*float64) *float64
type float64FunctorForFilterPtr func(*float64) bool

// MakeFloat64Slice - creates slice for the functional method such as map, filter
func MakeFloat64Slice(values ...float64) float64Slice {
	newSlice := float64Slice(values)
	return newSlice
}

// Map - applies the function(1st argument) on each item of the list and returns new list
func (slice float64Slice) Map(functors ...float64FunctorForMap) float64Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = MapFloat64(f, tmpSlice)
	}

	return tmpSlice
}

// MakeFloat64SlicePtr - creates slice for the functional method such as map, filter
func MakeFloat64SlicePtr(values ...*float64) float64SlicePtr {
	newSlice := float64SlicePtr(values)
	return newSlice
}

// MapPtr - applies the function(1st argument) on each item of the list and returns new list
func (slice float64SlicePtr) MapPtr(functors ...float64FunctorForMapPtr) float64SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = MapFloat64Ptr(f, tmpSlice)
	}

	return tmpSlice
}

// Filter - filters list based on function passed as argument
func (slice float64Slice) Filter(functors ...float64FunctorForFilter) float64Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = FilterFloat64(f, tmpSlice)
	}

	return tmpSlice
}

// FilterPtr - filters list based on function passed as argument
func (slice float64SlicePtr) FilterPtr(functors ...float64FunctorForFilterPtr) float64SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = FilterFloat64Ptr(f, tmpSlice)
	}

	return tmpSlice
}

// Remove - removes the items from the given list based on supplied function and returns new list
func (slice float64Slice) Remove(functors ...float64FunctorForFilter) float64Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = RemoveFloat64(f, tmpSlice)
	}

	return tmpSlice
}

// RemovePtr - removes the items from the given list based on supplied function and returns new list
func (slice float64SlicePtr) RemovePtr(functors ...float64FunctorForFilterPtr) float64SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = RemoveFloat64Ptr(f, tmpSlice)
	}

	return tmpSlice
}
