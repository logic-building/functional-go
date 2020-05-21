package fp

import "sync"

// PMapIntInt64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int output type: *int64
//	2. List
//
// Returns
//	New List of type *int64
//	Empty list if all arguments are nil or either one is nil
func PMapIntInt64Ptr(f func(*int) *int64, list []*int) []*int64 {
	if f == nil {
		return []*int64{}
	}

	ch := make(chan map[int]*int64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int64, i int, v *int) {
			defer wg.Done()
			ch <- map[int]*int64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapIntInt32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int output type: *int32
//	2. List
//
// Returns
//	New List of type *int32
//	Empty list if all arguments are nil or either one is nil
func PMapIntInt32Ptr(f func(*int) *int32, list []*int) []*int32 {
	if f == nil {
		return []*int32{}
	}

	ch := make(chan map[int]*int32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int32, i int, v *int) {
			defer wg.Done()
			ch <- map[int]*int32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapIntInt16Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int output type: *int16
//	2. List
//
// Returns
//	New List of type *int16
//	Empty list if all arguments are nil or either one is nil
func PMapIntInt16Ptr(f func(*int) *int16, list []*int) []*int16 {
	if f == nil {
		return []*int16{}
	}

	ch := make(chan map[int]*int16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int16, i int, v *int) {
			defer wg.Done()
			ch <- map[int]*int16{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapIntInt8Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int output type: *int8
//	2. List
//
// Returns
//	New List of type *int8
//	Empty list if all arguments are nil or either one is nil
func PMapIntInt8Ptr(f func(*int) *int8, list []*int) []*int8 {
	if f == nil {
		return []*int8{}
	}

	ch := make(chan map[int]*int8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int8, i int, v *int) {
			defer wg.Done()
			ch <- map[int]*int8{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapIntUintPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int output type: *uint
//	2. List
//
// Returns
//	New List of type *uint
//	Empty list if all arguments are nil or either one is nil
func PMapIntUintPtr(f func(*int) *uint, list []*int) []*uint {
	if f == nil {
		return []*uint{}
	}

	ch := make(chan map[int]*uint)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint, i int, v *int) {
			defer wg.Done()
			ch <- map[int]*uint{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapIntUint64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int output type: *uint64
//	2. List
//
// Returns
//	New List of type *uint64
//	Empty list if all arguments are nil or either one is nil
func PMapIntUint64Ptr(f func(*int) *uint64, list []*int) []*uint64 {
	if f == nil {
		return []*uint64{}
	}

	ch := make(chan map[int]*uint64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint64, i int, v *int) {
			defer wg.Done()
			ch <- map[int]*uint64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapIntUint32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int output type: *uint32
//	2. List
//
// Returns
//	New List of type *uint32
//	Empty list if all arguments are nil or either one is nil
func PMapIntUint32Ptr(f func(*int) *uint32, list []*int) []*uint32 {
	if f == nil {
		return []*uint32{}
	}

	ch := make(chan map[int]*uint32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint32, i int, v *int) {
			defer wg.Done()
			ch <- map[int]*uint32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapIntUint16Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int output type: *uint16
//	2. List
//
// Returns
//	New List of type *uint16
//	Empty list if all arguments are nil or either one is nil
func PMapIntUint16Ptr(f func(*int) *uint16, list []*int) []*uint16 {
	if f == nil {
		return []*uint16{}
	}

	ch := make(chan map[int]*uint16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint16, i int, v *int) {
			defer wg.Done()
			ch <- map[int]*uint16{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapIntUint8Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int output type: *uint8
//	2. List
//
// Returns
//	New List of type *uint8
//	Empty list if all arguments are nil or either one is nil
func PMapIntUint8Ptr(f func(*int) *uint8, list []*int) []*uint8 {
	if f == nil {
		return []*uint8{}
	}

	ch := make(chan map[int]*uint8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint8, i int, v *int) {
			defer wg.Done()
			ch <- map[int]*uint8{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapIntStrPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int output type: *string
//	2. List
//
// Returns
//	New List of type *string
//	Empty list if all arguments are nil or either one is nil
func PMapIntStrPtr(f func(*int) *string, list []*int) []*string {
	if f == nil {
		return []*string{}
	}

	ch := make(chan map[int]*string)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*string, i int, v *int) {
			defer wg.Done()
			ch <- map[int]*string{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*string, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapIntBoolPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int output type: *bool
//	2. List
//
// Returns
//	New List of type *bool
//	Empty list if all arguments are nil or either one is nil
func PMapIntBoolPtr(f func(*int) *bool, list []*int) []*bool {
	if f == nil {
		return []*bool{}
	}

	ch := make(chan map[int]*bool)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*bool, i int, v *int) {
			defer wg.Done()
			ch <- map[int]*bool{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*bool, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapIntFloat32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int output type: *float32
//	2. List
//
// Returns
//	New List of type *float32
//	Empty list if all arguments are nil or either one is nil
func PMapIntFloat32Ptr(f func(*int) *float32, list []*int) []*float32 {
	if f == nil {
		return []*float32{}
	}

	ch := make(chan map[int]*float32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*float32, i int, v *int) {
			defer wg.Done()
			ch <- map[int]*float32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*float32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapIntFloat64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int output type: *float64
//	2. List
//
// Returns
//	New List of type *float64
//	Empty list if all arguments are nil or either one is nil
func PMapIntFloat64Ptr(f func(*int) *float64, list []*int) []*float64 {
	if f == nil {
		return []*float64{}
	}

	ch := make(chan map[int]*float64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*float64, i int, v *int) {
			defer wg.Done()
			ch <- map[int]*float64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*float64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt64IntPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int64 output type: *int
//	2. List
//
// Returns
//	New List of type *int
//	Empty list if all arguments are nil or either one is nil
func PMapInt64IntPtr(f func(*int64) *int, list []*int64) []*int {
	if f == nil {
		return []*int{}
	}

	ch := make(chan map[int]*int)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int, i int, v *int64) {
			defer wg.Done()
			ch <- map[int]*int{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt64Int32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int64 output type: *int32
//	2. List
//
// Returns
//	New List of type *int32
//	Empty list if all arguments are nil or either one is nil
func PMapInt64Int32Ptr(f func(*int64) *int32, list []*int64) []*int32 {
	if f == nil {
		return []*int32{}
	}

	ch := make(chan map[int]*int32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int32, i int, v *int64) {
			defer wg.Done()
			ch <- map[int]*int32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt64Int16Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int64 output type: *int16
//	2. List
//
// Returns
//	New List of type *int16
//	Empty list if all arguments are nil or either one is nil
func PMapInt64Int16Ptr(f func(*int64) *int16, list []*int64) []*int16 {
	if f == nil {
		return []*int16{}
	}

	ch := make(chan map[int]*int16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int16, i int, v *int64) {
			defer wg.Done()
			ch <- map[int]*int16{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt64Int8Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int64 output type: *int8
//	2. List
//
// Returns
//	New List of type *int8
//	Empty list if all arguments are nil or either one is nil
func PMapInt64Int8Ptr(f func(*int64) *int8, list []*int64) []*int8 {
	if f == nil {
		return []*int8{}
	}

	ch := make(chan map[int]*int8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int8, i int, v *int64) {
			defer wg.Done()
			ch <- map[int]*int8{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt64UintPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int64 output type: *uint
//	2. List
//
// Returns
//	New List of type *uint
//	Empty list if all arguments are nil or either one is nil
func PMapInt64UintPtr(f func(*int64) *uint, list []*int64) []*uint {
	if f == nil {
		return []*uint{}
	}

	ch := make(chan map[int]*uint)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint, i int, v *int64) {
			defer wg.Done()
			ch <- map[int]*uint{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt64Uint64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int64 output type: *uint64
//	2. List
//
// Returns
//	New List of type *uint64
//	Empty list if all arguments are nil or either one is nil
func PMapInt64Uint64Ptr(f func(*int64) *uint64, list []*int64) []*uint64 {
	if f == nil {
		return []*uint64{}
	}

	ch := make(chan map[int]*uint64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint64, i int, v *int64) {
			defer wg.Done()
			ch <- map[int]*uint64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt64Uint32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int64 output type: *uint32
//	2. List
//
// Returns
//	New List of type *uint32
//	Empty list if all arguments are nil or either one is nil
func PMapInt64Uint32Ptr(f func(*int64) *uint32, list []*int64) []*uint32 {
	if f == nil {
		return []*uint32{}
	}

	ch := make(chan map[int]*uint32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint32, i int, v *int64) {
			defer wg.Done()
			ch <- map[int]*uint32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt64Uint16Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int64 output type: *uint16
//	2. List
//
// Returns
//	New List of type *uint16
//	Empty list if all arguments are nil or either one is nil
func PMapInt64Uint16Ptr(f func(*int64) *uint16, list []*int64) []*uint16 {
	if f == nil {
		return []*uint16{}
	}

	ch := make(chan map[int]*uint16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint16, i int, v *int64) {
			defer wg.Done()
			ch <- map[int]*uint16{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt64Uint8Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int64 output type: *uint8
//	2. List
//
// Returns
//	New List of type *uint8
//	Empty list if all arguments are nil or either one is nil
func PMapInt64Uint8Ptr(f func(*int64) *uint8, list []*int64) []*uint8 {
	if f == nil {
		return []*uint8{}
	}

	ch := make(chan map[int]*uint8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint8, i int, v *int64) {
			defer wg.Done()
			ch <- map[int]*uint8{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt64StrPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int64 output type: *string
//	2. List
//
// Returns
//	New List of type *string
//	Empty list if all arguments are nil or either one is nil
func PMapInt64StrPtr(f func(*int64) *string, list []*int64) []*string {
	if f == nil {
		return []*string{}
	}

	ch := make(chan map[int]*string)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*string, i int, v *int64) {
			defer wg.Done()
			ch <- map[int]*string{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*string, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt64BoolPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int64 output type: *bool
//	2. List
//
// Returns
//	New List of type *bool
//	Empty list if all arguments are nil or either one is nil
func PMapInt64BoolPtr(f func(*int64) *bool, list []*int64) []*bool {
	if f == nil {
		return []*bool{}
	}

	ch := make(chan map[int]*bool)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*bool, i int, v *int64) {
			defer wg.Done()
			ch <- map[int]*bool{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*bool, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt64Float32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int64 output type: *float32
//	2. List
//
// Returns
//	New List of type *float32
//	Empty list if all arguments are nil or either one is nil
func PMapInt64Float32Ptr(f func(*int64) *float32, list []*int64) []*float32 {
	if f == nil {
		return []*float32{}
	}

	ch := make(chan map[int]*float32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*float32, i int, v *int64) {
			defer wg.Done()
			ch <- map[int]*float32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*float32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt64Float64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int64 output type: *float64
//	2. List
//
// Returns
//	New List of type *float64
//	Empty list if all arguments are nil or either one is nil
func PMapInt64Float64Ptr(f func(*int64) *float64, list []*int64) []*float64 {
	if f == nil {
		return []*float64{}
	}

	ch := make(chan map[int]*float64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*float64, i int, v *int64) {
			defer wg.Done()
			ch <- map[int]*float64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*float64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt32IntPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int32 output type: *int
//	2. List
//
// Returns
//	New List of type *int
//	Empty list if all arguments are nil or either one is nil
func PMapInt32IntPtr(f func(*int32) *int, list []*int32) []*int {
	if f == nil {
		return []*int{}
	}

	ch := make(chan map[int]*int)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int, i int, v *int32) {
			defer wg.Done()
			ch <- map[int]*int{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt32Int64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int32 output type: *int64
//	2. List
//
// Returns
//	New List of type *int64
//	Empty list if all arguments are nil or either one is nil
func PMapInt32Int64Ptr(f func(*int32) *int64, list []*int32) []*int64 {
	if f == nil {
		return []*int64{}
	}

	ch := make(chan map[int]*int64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int64, i int, v *int32) {
			defer wg.Done()
			ch <- map[int]*int64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt32Int16Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int32 output type: *int16
//	2. List
//
// Returns
//	New List of type *int16
//	Empty list if all arguments are nil or either one is nil
func PMapInt32Int16Ptr(f func(*int32) *int16, list []*int32) []*int16 {
	if f == nil {
		return []*int16{}
	}

	ch := make(chan map[int]*int16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int16, i int, v *int32) {
			defer wg.Done()
			ch <- map[int]*int16{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt32Int8Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int32 output type: *int8
//	2. List
//
// Returns
//	New List of type *int8
//	Empty list if all arguments are nil or either one is nil
func PMapInt32Int8Ptr(f func(*int32) *int8, list []*int32) []*int8 {
	if f == nil {
		return []*int8{}
	}

	ch := make(chan map[int]*int8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int8, i int, v *int32) {
			defer wg.Done()
			ch <- map[int]*int8{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt32UintPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int32 output type: *uint
//	2. List
//
// Returns
//	New List of type *uint
//	Empty list if all arguments are nil or either one is nil
func PMapInt32UintPtr(f func(*int32) *uint, list []*int32) []*uint {
	if f == nil {
		return []*uint{}
	}

	ch := make(chan map[int]*uint)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint, i int, v *int32) {
			defer wg.Done()
			ch <- map[int]*uint{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt32Uint64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int32 output type: *uint64
//	2. List
//
// Returns
//	New List of type *uint64
//	Empty list if all arguments are nil or either one is nil
func PMapInt32Uint64Ptr(f func(*int32) *uint64, list []*int32) []*uint64 {
	if f == nil {
		return []*uint64{}
	}

	ch := make(chan map[int]*uint64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint64, i int, v *int32) {
			defer wg.Done()
			ch <- map[int]*uint64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt32Uint32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int32 output type: *uint32
//	2. List
//
// Returns
//	New List of type *uint32
//	Empty list if all arguments are nil or either one is nil
func PMapInt32Uint32Ptr(f func(*int32) *uint32, list []*int32) []*uint32 {
	if f == nil {
		return []*uint32{}
	}

	ch := make(chan map[int]*uint32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint32, i int, v *int32) {
			defer wg.Done()
			ch <- map[int]*uint32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt32Uint16Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int32 output type: *uint16
//	2. List
//
// Returns
//	New List of type *uint16
//	Empty list if all arguments are nil or either one is nil
func PMapInt32Uint16Ptr(f func(*int32) *uint16, list []*int32) []*uint16 {
	if f == nil {
		return []*uint16{}
	}

	ch := make(chan map[int]*uint16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint16, i int, v *int32) {
			defer wg.Done()
			ch <- map[int]*uint16{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt32Uint8Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int32 output type: *uint8
//	2. List
//
// Returns
//	New List of type *uint8
//	Empty list if all arguments are nil or either one is nil
func PMapInt32Uint8Ptr(f func(*int32) *uint8, list []*int32) []*uint8 {
	if f == nil {
		return []*uint8{}
	}

	ch := make(chan map[int]*uint8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint8, i int, v *int32) {
			defer wg.Done()
			ch <- map[int]*uint8{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt32StrPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int32 output type: *string
//	2. List
//
// Returns
//	New List of type *string
//	Empty list if all arguments are nil or either one is nil
func PMapInt32StrPtr(f func(*int32) *string, list []*int32) []*string {
	if f == nil {
		return []*string{}
	}

	ch := make(chan map[int]*string)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*string, i int, v *int32) {
			defer wg.Done()
			ch <- map[int]*string{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*string, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt32BoolPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int32 output type: *bool
//	2. List
//
// Returns
//	New List of type *bool
//	Empty list if all arguments are nil or either one is nil
func PMapInt32BoolPtr(f func(*int32) *bool, list []*int32) []*bool {
	if f == nil {
		return []*bool{}
	}

	ch := make(chan map[int]*bool)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*bool, i int, v *int32) {
			defer wg.Done()
			ch <- map[int]*bool{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*bool, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt32Float32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int32 output type: *float32
//	2. List
//
// Returns
//	New List of type *float32
//	Empty list if all arguments are nil or either one is nil
func PMapInt32Float32Ptr(f func(*int32) *float32, list []*int32) []*float32 {
	if f == nil {
		return []*float32{}
	}

	ch := make(chan map[int]*float32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*float32, i int, v *int32) {
			defer wg.Done()
			ch <- map[int]*float32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*float32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt32Float64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int32 output type: *float64
//	2. List
//
// Returns
//	New List of type *float64
//	Empty list if all arguments are nil or either one is nil
func PMapInt32Float64Ptr(f func(*int32) *float64, list []*int32) []*float64 {
	if f == nil {
		return []*float64{}
	}

	ch := make(chan map[int]*float64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*float64, i int, v *int32) {
			defer wg.Done()
			ch <- map[int]*float64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*float64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt16IntPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int16 output type: *int
//	2. List
//
// Returns
//	New List of type *int
//	Empty list if all arguments are nil or either one is nil
func PMapInt16IntPtr(f func(*int16) *int, list []*int16) []*int {
	if f == nil {
		return []*int{}
	}

	ch := make(chan map[int]*int)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int, i int, v *int16) {
			defer wg.Done()
			ch <- map[int]*int{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt16Int64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int16 output type: *int64
//	2. List
//
// Returns
//	New List of type *int64
//	Empty list if all arguments are nil or either one is nil
func PMapInt16Int64Ptr(f func(*int16) *int64, list []*int16) []*int64 {
	if f == nil {
		return []*int64{}
	}

	ch := make(chan map[int]*int64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int64, i int, v *int16) {
			defer wg.Done()
			ch <- map[int]*int64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt16Int32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int16 output type: *int32
//	2. List
//
// Returns
//	New List of type *int32
//	Empty list if all arguments are nil or either one is nil
func PMapInt16Int32Ptr(f func(*int16) *int32, list []*int16) []*int32 {
	if f == nil {
		return []*int32{}
	}

	ch := make(chan map[int]*int32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int32, i int, v *int16) {
			defer wg.Done()
			ch <- map[int]*int32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt16Int8Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int16 output type: *int8
//	2. List
//
// Returns
//	New List of type *int8
//	Empty list if all arguments are nil or either one is nil
func PMapInt16Int8Ptr(f func(*int16) *int8, list []*int16) []*int8 {
	if f == nil {
		return []*int8{}
	}

	ch := make(chan map[int]*int8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int8, i int, v *int16) {
			defer wg.Done()
			ch <- map[int]*int8{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt16UintPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int16 output type: *uint
//	2. List
//
// Returns
//	New List of type *uint
//	Empty list if all arguments are nil or either one is nil
func PMapInt16UintPtr(f func(*int16) *uint, list []*int16) []*uint {
	if f == nil {
		return []*uint{}
	}

	ch := make(chan map[int]*uint)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint, i int, v *int16) {
			defer wg.Done()
			ch <- map[int]*uint{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt16Uint64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int16 output type: *uint64
//	2. List
//
// Returns
//	New List of type *uint64
//	Empty list if all arguments are nil or either one is nil
func PMapInt16Uint64Ptr(f func(*int16) *uint64, list []*int16) []*uint64 {
	if f == nil {
		return []*uint64{}
	}

	ch := make(chan map[int]*uint64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint64, i int, v *int16) {
			defer wg.Done()
			ch <- map[int]*uint64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt16Uint32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int16 output type: *uint32
//	2. List
//
// Returns
//	New List of type *uint32
//	Empty list if all arguments are nil or either one is nil
func PMapInt16Uint32Ptr(f func(*int16) *uint32, list []*int16) []*uint32 {
	if f == nil {
		return []*uint32{}
	}

	ch := make(chan map[int]*uint32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint32, i int, v *int16) {
			defer wg.Done()
			ch <- map[int]*uint32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt16Uint16Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int16 output type: *uint16
//	2. List
//
// Returns
//	New List of type *uint16
//	Empty list if all arguments are nil or either one is nil
func PMapInt16Uint16Ptr(f func(*int16) *uint16, list []*int16) []*uint16 {
	if f == nil {
		return []*uint16{}
	}

	ch := make(chan map[int]*uint16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint16, i int, v *int16) {
			defer wg.Done()
			ch <- map[int]*uint16{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt16Uint8Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int16 output type: *uint8
//	2. List
//
// Returns
//	New List of type *uint8
//	Empty list if all arguments are nil or either one is nil
func PMapInt16Uint8Ptr(f func(*int16) *uint8, list []*int16) []*uint8 {
	if f == nil {
		return []*uint8{}
	}

	ch := make(chan map[int]*uint8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint8, i int, v *int16) {
			defer wg.Done()
			ch <- map[int]*uint8{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt16StrPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int16 output type: *string
//	2. List
//
// Returns
//	New List of type *string
//	Empty list if all arguments are nil or either one is nil
func PMapInt16StrPtr(f func(*int16) *string, list []*int16) []*string {
	if f == nil {
		return []*string{}
	}

	ch := make(chan map[int]*string)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*string, i int, v *int16) {
			defer wg.Done()
			ch <- map[int]*string{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*string, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt16BoolPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int16 output type: *bool
//	2. List
//
// Returns
//	New List of type *bool
//	Empty list if all arguments are nil or either one is nil
func PMapInt16BoolPtr(f func(*int16) *bool, list []*int16) []*bool {
	if f == nil {
		return []*bool{}
	}

	ch := make(chan map[int]*bool)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*bool, i int, v *int16) {
			defer wg.Done()
			ch <- map[int]*bool{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*bool, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt16Float32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int16 output type: *float32
//	2. List
//
// Returns
//	New List of type *float32
//	Empty list if all arguments are nil or either one is nil
func PMapInt16Float32Ptr(f func(*int16) *float32, list []*int16) []*float32 {
	if f == nil {
		return []*float32{}
	}

	ch := make(chan map[int]*float32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*float32, i int, v *int16) {
			defer wg.Done()
			ch <- map[int]*float32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*float32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt16Float64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int16 output type: *float64
//	2. List
//
// Returns
//	New List of type *float64
//	Empty list if all arguments are nil or either one is nil
func PMapInt16Float64Ptr(f func(*int16) *float64, list []*int16) []*float64 {
	if f == nil {
		return []*float64{}
	}

	ch := make(chan map[int]*float64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*float64, i int, v *int16) {
			defer wg.Done()
			ch <- map[int]*float64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*float64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt8IntPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int8 output type: *int
//	2. List
//
// Returns
//	New List of type *int
//	Empty list if all arguments are nil or either one is nil
func PMapInt8IntPtr(f func(*int8) *int, list []*int8) []*int {
	if f == nil {
		return []*int{}
	}

	ch := make(chan map[int]*int)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int, i int, v *int8) {
			defer wg.Done()
			ch <- map[int]*int{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt8Int64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int8 output type: *int64
//	2. List
//
// Returns
//	New List of type *int64
//	Empty list if all arguments are nil or either one is nil
func PMapInt8Int64Ptr(f func(*int8) *int64, list []*int8) []*int64 {
	if f == nil {
		return []*int64{}
	}

	ch := make(chan map[int]*int64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int64, i int, v *int8) {
			defer wg.Done()
			ch <- map[int]*int64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt8Int32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int8 output type: *int32
//	2. List
//
// Returns
//	New List of type *int32
//	Empty list if all arguments are nil or either one is nil
func PMapInt8Int32Ptr(f func(*int8) *int32, list []*int8) []*int32 {
	if f == nil {
		return []*int32{}
	}

	ch := make(chan map[int]*int32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int32, i int, v *int8) {
			defer wg.Done()
			ch <- map[int]*int32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt8Int16Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int8 output type: *int16
//	2. List
//
// Returns
//	New List of type *int16
//	Empty list if all arguments are nil or either one is nil
func PMapInt8Int16Ptr(f func(*int8) *int16, list []*int8) []*int16 {
	if f == nil {
		return []*int16{}
	}

	ch := make(chan map[int]*int16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int16, i int, v *int8) {
			defer wg.Done()
			ch <- map[int]*int16{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt8UintPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int8 output type: *uint
//	2. List
//
// Returns
//	New List of type *uint
//	Empty list if all arguments are nil or either one is nil
func PMapInt8UintPtr(f func(*int8) *uint, list []*int8) []*uint {
	if f == nil {
		return []*uint{}
	}

	ch := make(chan map[int]*uint)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint, i int, v *int8) {
			defer wg.Done()
			ch <- map[int]*uint{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt8Uint64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int8 output type: *uint64
//	2. List
//
// Returns
//	New List of type *uint64
//	Empty list if all arguments are nil or either one is nil
func PMapInt8Uint64Ptr(f func(*int8) *uint64, list []*int8) []*uint64 {
	if f == nil {
		return []*uint64{}
	}

	ch := make(chan map[int]*uint64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint64, i int, v *int8) {
			defer wg.Done()
			ch <- map[int]*uint64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt8Uint32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int8 output type: *uint32
//	2. List
//
// Returns
//	New List of type *uint32
//	Empty list if all arguments are nil or either one is nil
func PMapInt8Uint32Ptr(f func(*int8) *uint32, list []*int8) []*uint32 {
	if f == nil {
		return []*uint32{}
	}

	ch := make(chan map[int]*uint32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint32, i int, v *int8) {
			defer wg.Done()
			ch <- map[int]*uint32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt8Uint16Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int8 output type: *uint16
//	2. List
//
// Returns
//	New List of type *uint16
//	Empty list if all arguments are nil or either one is nil
func PMapInt8Uint16Ptr(f func(*int8) *uint16, list []*int8) []*uint16 {
	if f == nil {
		return []*uint16{}
	}

	ch := make(chan map[int]*uint16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint16, i int, v *int8) {
			defer wg.Done()
			ch <- map[int]*uint16{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt8Uint8Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int8 output type: *uint8
//	2. List
//
// Returns
//	New List of type *uint8
//	Empty list if all arguments are nil or either one is nil
func PMapInt8Uint8Ptr(f func(*int8) *uint8, list []*int8) []*uint8 {
	if f == nil {
		return []*uint8{}
	}

	ch := make(chan map[int]*uint8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint8, i int, v *int8) {
			defer wg.Done()
			ch <- map[int]*uint8{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt8StrPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int8 output type: *string
//	2. List
//
// Returns
//	New List of type *string
//	Empty list if all arguments are nil or either one is nil
func PMapInt8StrPtr(f func(*int8) *string, list []*int8) []*string {
	if f == nil {
		return []*string{}
	}

	ch := make(chan map[int]*string)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*string, i int, v *int8) {
			defer wg.Done()
			ch <- map[int]*string{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*string, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt8BoolPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int8 output type: *bool
//	2. List
//
// Returns
//	New List of type *bool
//	Empty list if all arguments are nil or either one is nil
func PMapInt8BoolPtr(f func(*int8) *bool, list []*int8) []*bool {
	if f == nil {
		return []*bool{}
	}

	ch := make(chan map[int]*bool)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*bool, i int, v *int8) {
			defer wg.Done()
			ch <- map[int]*bool{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*bool, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt8Float32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int8 output type: *float32
//	2. List
//
// Returns
//	New List of type *float32
//	Empty list if all arguments are nil or either one is nil
func PMapInt8Float32Ptr(f func(*int8) *float32, list []*int8) []*float32 {
	if f == nil {
		return []*float32{}
	}

	ch := make(chan map[int]*float32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*float32, i int, v *int8) {
			defer wg.Done()
			ch <- map[int]*float32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*float32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapInt8Float64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *int8 output type: *float64
//	2. List
//
// Returns
//	New List of type *float64
//	Empty list if all arguments are nil or either one is nil
func PMapInt8Float64Ptr(f func(*int8) *float64, list []*int8) []*float64 {
	if f == nil {
		return []*float64{}
	}

	ch := make(chan map[int]*float64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*float64, i int, v *int8) {
			defer wg.Done()
			ch <- map[int]*float64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*float64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUintIntPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint output type: *int
//	2. List
//
// Returns
//	New List of type *int
//	Empty list if all arguments are nil or either one is nil
func PMapUintIntPtr(f func(*uint) *int, list []*uint) []*int {
	if f == nil {
		return []*int{}
	}

	ch := make(chan map[int]*int)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int, i int, v *uint) {
			defer wg.Done()
			ch <- map[int]*int{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUintInt64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint output type: *int64
//	2. List
//
// Returns
//	New List of type *int64
//	Empty list if all arguments are nil or either one is nil
func PMapUintInt64Ptr(f func(*uint) *int64, list []*uint) []*int64 {
	if f == nil {
		return []*int64{}
	}

	ch := make(chan map[int]*int64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int64, i int, v *uint) {
			defer wg.Done()
			ch <- map[int]*int64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUintInt32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint output type: *int32
//	2. List
//
// Returns
//	New List of type *int32
//	Empty list if all arguments are nil or either one is nil
func PMapUintInt32Ptr(f func(*uint) *int32, list []*uint) []*int32 {
	if f == nil {
		return []*int32{}
	}

	ch := make(chan map[int]*int32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int32, i int, v *uint) {
			defer wg.Done()
			ch <- map[int]*int32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUintInt16Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint output type: *int16
//	2. List
//
// Returns
//	New List of type *int16
//	Empty list if all arguments are nil or either one is nil
func PMapUintInt16Ptr(f func(*uint) *int16, list []*uint) []*int16 {
	if f == nil {
		return []*int16{}
	}

	ch := make(chan map[int]*int16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int16, i int, v *uint) {
			defer wg.Done()
			ch <- map[int]*int16{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUintInt8Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint output type: *int8
//	2. List
//
// Returns
//	New List of type *int8
//	Empty list if all arguments are nil or either one is nil
func PMapUintInt8Ptr(f func(*uint) *int8, list []*uint) []*int8 {
	if f == nil {
		return []*int8{}
	}

	ch := make(chan map[int]*int8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int8, i int, v *uint) {
			defer wg.Done()
			ch <- map[int]*int8{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUintUint64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint output type: *uint64
//	2. List
//
// Returns
//	New List of type *uint64
//	Empty list if all arguments are nil or either one is nil
func PMapUintUint64Ptr(f func(*uint) *uint64, list []*uint) []*uint64 {
	if f == nil {
		return []*uint64{}
	}

	ch := make(chan map[int]*uint64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint64, i int, v *uint) {
			defer wg.Done()
			ch <- map[int]*uint64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUintUint32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint output type: *uint32
//	2. List
//
// Returns
//	New List of type *uint32
//	Empty list if all arguments are nil or either one is nil
func PMapUintUint32Ptr(f func(*uint) *uint32, list []*uint) []*uint32 {
	if f == nil {
		return []*uint32{}
	}

	ch := make(chan map[int]*uint32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint32, i int, v *uint) {
			defer wg.Done()
			ch <- map[int]*uint32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUintUint16Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint output type: *uint16
//	2. List
//
// Returns
//	New List of type *uint16
//	Empty list if all arguments are nil or either one is nil
func PMapUintUint16Ptr(f func(*uint) *uint16, list []*uint) []*uint16 {
	if f == nil {
		return []*uint16{}
	}

	ch := make(chan map[int]*uint16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint16, i int, v *uint) {
			defer wg.Done()
			ch <- map[int]*uint16{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUintUint8Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint output type: *uint8
//	2. List
//
// Returns
//	New List of type *uint8
//	Empty list if all arguments are nil or either one is nil
func PMapUintUint8Ptr(f func(*uint) *uint8, list []*uint) []*uint8 {
	if f == nil {
		return []*uint8{}
	}

	ch := make(chan map[int]*uint8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint8, i int, v *uint) {
			defer wg.Done()
			ch <- map[int]*uint8{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUintStrPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint output type: *string
//	2. List
//
// Returns
//	New List of type *string
//	Empty list if all arguments are nil or either one is nil
func PMapUintStrPtr(f func(*uint) *string, list []*uint) []*string {
	if f == nil {
		return []*string{}
	}

	ch := make(chan map[int]*string)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*string, i int, v *uint) {
			defer wg.Done()
			ch <- map[int]*string{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*string, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUintBoolPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint output type: *bool
//	2. List
//
// Returns
//	New List of type *bool
//	Empty list if all arguments are nil or either one is nil
func PMapUintBoolPtr(f func(*uint) *bool, list []*uint) []*bool {
	if f == nil {
		return []*bool{}
	}

	ch := make(chan map[int]*bool)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*bool, i int, v *uint) {
			defer wg.Done()
			ch <- map[int]*bool{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*bool, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUintFloat32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint output type: *float32
//	2. List
//
// Returns
//	New List of type *float32
//	Empty list if all arguments are nil or either one is nil
func PMapUintFloat32Ptr(f func(*uint) *float32, list []*uint) []*float32 {
	if f == nil {
		return []*float32{}
	}

	ch := make(chan map[int]*float32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*float32, i int, v *uint) {
			defer wg.Done()
			ch <- map[int]*float32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*float32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUintFloat64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint output type: *float64
//	2. List
//
// Returns
//	New List of type *float64
//	Empty list if all arguments are nil or either one is nil
func PMapUintFloat64Ptr(f func(*uint) *float64, list []*uint) []*float64 {
	if f == nil {
		return []*float64{}
	}

	ch := make(chan map[int]*float64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*float64, i int, v *uint) {
			defer wg.Done()
			ch <- map[int]*float64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*float64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint64IntPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint64 output type: *int
//	2. List
//
// Returns
//	New List of type *int
//	Empty list if all arguments are nil or either one is nil
func PMapUint64IntPtr(f func(*uint64) *int, list []*uint64) []*int {
	if f == nil {
		return []*int{}
	}

	ch := make(chan map[int]*int)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int, i int, v *uint64) {
			defer wg.Done()
			ch <- map[int]*int{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint64Int64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint64 output type: *int64
//	2. List
//
// Returns
//	New List of type *int64
//	Empty list if all arguments are nil or either one is nil
func PMapUint64Int64Ptr(f func(*uint64) *int64, list []*uint64) []*int64 {
	if f == nil {
		return []*int64{}
	}

	ch := make(chan map[int]*int64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int64, i int, v *uint64) {
			defer wg.Done()
			ch <- map[int]*int64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint64Int32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint64 output type: *int32
//	2. List
//
// Returns
//	New List of type *int32
//	Empty list if all arguments are nil or either one is nil
func PMapUint64Int32Ptr(f func(*uint64) *int32, list []*uint64) []*int32 {
	if f == nil {
		return []*int32{}
	}

	ch := make(chan map[int]*int32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int32, i int, v *uint64) {
			defer wg.Done()
			ch <- map[int]*int32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint64Int16Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint64 output type: *int16
//	2. List
//
// Returns
//	New List of type *int16
//	Empty list if all arguments are nil or either one is nil
func PMapUint64Int16Ptr(f func(*uint64) *int16, list []*uint64) []*int16 {
	if f == nil {
		return []*int16{}
	}

	ch := make(chan map[int]*int16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int16, i int, v *uint64) {
			defer wg.Done()
			ch <- map[int]*int16{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint64Int8Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint64 output type: *int8
//	2. List
//
// Returns
//	New List of type *int8
//	Empty list if all arguments are nil or either one is nil
func PMapUint64Int8Ptr(f func(*uint64) *int8, list []*uint64) []*int8 {
	if f == nil {
		return []*int8{}
	}

	ch := make(chan map[int]*int8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int8, i int, v *uint64) {
			defer wg.Done()
			ch <- map[int]*int8{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint64UintPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint64 output type: *uint
//	2. List
//
// Returns
//	New List of type *uint
//	Empty list if all arguments are nil or either one is nil
func PMapUint64UintPtr(f func(*uint64) *uint, list []*uint64) []*uint {
	if f == nil {
		return []*uint{}
	}

	ch := make(chan map[int]*uint)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint, i int, v *uint64) {
			defer wg.Done()
			ch <- map[int]*uint{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint64Uint32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint64 output type: *uint32
//	2. List
//
// Returns
//	New List of type *uint32
//	Empty list if all arguments are nil or either one is nil
func PMapUint64Uint32Ptr(f func(*uint64) *uint32, list []*uint64) []*uint32 {
	if f == nil {
		return []*uint32{}
	}

	ch := make(chan map[int]*uint32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint32, i int, v *uint64) {
			defer wg.Done()
			ch <- map[int]*uint32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint64Uint16Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint64 output type: *uint16
//	2. List
//
// Returns
//	New List of type *uint16
//	Empty list if all arguments are nil or either one is nil
func PMapUint64Uint16Ptr(f func(*uint64) *uint16, list []*uint64) []*uint16 {
	if f == nil {
		return []*uint16{}
	}

	ch := make(chan map[int]*uint16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint16, i int, v *uint64) {
			defer wg.Done()
			ch <- map[int]*uint16{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint64Uint8Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint64 output type: *uint8
//	2. List
//
// Returns
//	New List of type *uint8
//	Empty list if all arguments are nil or either one is nil
func PMapUint64Uint8Ptr(f func(*uint64) *uint8, list []*uint64) []*uint8 {
	if f == nil {
		return []*uint8{}
	}

	ch := make(chan map[int]*uint8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint8, i int, v *uint64) {
			defer wg.Done()
			ch <- map[int]*uint8{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint64StrPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint64 output type: *string
//	2. List
//
// Returns
//	New List of type *string
//	Empty list if all arguments are nil or either one is nil
func PMapUint64StrPtr(f func(*uint64) *string, list []*uint64) []*string {
	if f == nil {
		return []*string{}
	}

	ch := make(chan map[int]*string)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*string, i int, v *uint64) {
			defer wg.Done()
			ch <- map[int]*string{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*string, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint64BoolPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint64 output type: *bool
//	2. List
//
// Returns
//	New List of type *bool
//	Empty list if all arguments are nil or either one is nil
func PMapUint64BoolPtr(f func(*uint64) *bool, list []*uint64) []*bool {
	if f == nil {
		return []*bool{}
	}

	ch := make(chan map[int]*bool)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*bool, i int, v *uint64) {
			defer wg.Done()
			ch <- map[int]*bool{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*bool, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint64Float32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint64 output type: *float32
//	2. List
//
// Returns
//	New List of type *float32
//	Empty list if all arguments are nil or either one is nil
func PMapUint64Float32Ptr(f func(*uint64) *float32, list []*uint64) []*float32 {
	if f == nil {
		return []*float32{}
	}

	ch := make(chan map[int]*float32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*float32, i int, v *uint64) {
			defer wg.Done()
			ch <- map[int]*float32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*float32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint64Float64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint64 output type: *float64
//	2. List
//
// Returns
//	New List of type *float64
//	Empty list if all arguments are nil or either one is nil
func PMapUint64Float64Ptr(f func(*uint64) *float64, list []*uint64) []*float64 {
	if f == nil {
		return []*float64{}
	}

	ch := make(chan map[int]*float64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*float64, i int, v *uint64) {
			defer wg.Done()
			ch <- map[int]*float64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*float64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint32IntPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint32 output type: *int
//	2. List
//
// Returns
//	New List of type *int
//	Empty list if all arguments are nil or either one is nil
func PMapUint32IntPtr(f func(*uint32) *int, list []*uint32) []*int {
	if f == nil {
		return []*int{}
	}

	ch := make(chan map[int]*int)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int, i int, v *uint32) {
			defer wg.Done()
			ch <- map[int]*int{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint32Int64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint32 output type: *int64
//	2. List
//
// Returns
//	New List of type *int64
//	Empty list if all arguments are nil or either one is nil
func PMapUint32Int64Ptr(f func(*uint32) *int64, list []*uint32) []*int64 {
	if f == nil {
		return []*int64{}
	}

	ch := make(chan map[int]*int64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int64, i int, v *uint32) {
			defer wg.Done()
			ch <- map[int]*int64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint32Int32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint32 output type: *int32
//	2. List
//
// Returns
//	New List of type *int32
//	Empty list if all arguments are nil or either one is nil
func PMapUint32Int32Ptr(f func(*uint32) *int32, list []*uint32) []*int32 {
	if f == nil {
		return []*int32{}
	}

	ch := make(chan map[int]*int32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int32, i int, v *uint32) {
			defer wg.Done()
			ch <- map[int]*int32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint32Int16Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint32 output type: *int16
//	2. List
//
// Returns
//	New List of type *int16
//	Empty list if all arguments are nil or either one is nil
func PMapUint32Int16Ptr(f func(*uint32) *int16, list []*uint32) []*int16 {
	if f == nil {
		return []*int16{}
	}

	ch := make(chan map[int]*int16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int16, i int, v *uint32) {
			defer wg.Done()
			ch <- map[int]*int16{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint32Int8Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint32 output type: *int8
//	2. List
//
// Returns
//	New List of type *int8
//	Empty list if all arguments are nil or either one is nil
func PMapUint32Int8Ptr(f func(*uint32) *int8, list []*uint32) []*int8 {
	if f == nil {
		return []*int8{}
	}

	ch := make(chan map[int]*int8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int8, i int, v *uint32) {
			defer wg.Done()
			ch <- map[int]*int8{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint32UintPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint32 output type: *uint
//	2. List
//
// Returns
//	New List of type *uint
//	Empty list if all arguments are nil or either one is nil
func PMapUint32UintPtr(f func(*uint32) *uint, list []*uint32) []*uint {
	if f == nil {
		return []*uint{}
	}

	ch := make(chan map[int]*uint)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint, i int, v *uint32) {
			defer wg.Done()
			ch <- map[int]*uint{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint32Uint64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint32 output type: *uint64
//	2. List
//
// Returns
//	New List of type *uint64
//	Empty list if all arguments are nil or either one is nil
func PMapUint32Uint64Ptr(f func(*uint32) *uint64, list []*uint32) []*uint64 {
	if f == nil {
		return []*uint64{}
	}

	ch := make(chan map[int]*uint64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint64, i int, v *uint32) {
			defer wg.Done()
			ch <- map[int]*uint64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint32Uint16Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint32 output type: *uint16
//	2. List
//
// Returns
//	New List of type *uint16
//	Empty list if all arguments are nil or either one is nil
func PMapUint32Uint16Ptr(f func(*uint32) *uint16, list []*uint32) []*uint16 {
	if f == nil {
		return []*uint16{}
	}

	ch := make(chan map[int]*uint16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint16, i int, v *uint32) {
			defer wg.Done()
			ch <- map[int]*uint16{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint32Uint8Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint32 output type: *uint8
//	2. List
//
// Returns
//	New List of type *uint8
//	Empty list if all arguments are nil or either one is nil
func PMapUint32Uint8Ptr(f func(*uint32) *uint8, list []*uint32) []*uint8 {
	if f == nil {
		return []*uint8{}
	}

	ch := make(chan map[int]*uint8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint8, i int, v *uint32) {
			defer wg.Done()
			ch <- map[int]*uint8{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint32StrPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint32 output type: *string
//	2. List
//
// Returns
//	New List of type *string
//	Empty list if all arguments are nil or either one is nil
func PMapUint32StrPtr(f func(*uint32) *string, list []*uint32) []*string {
	if f == nil {
		return []*string{}
	}

	ch := make(chan map[int]*string)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*string, i int, v *uint32) {
			defer wg.Done()
			ch <- map[int]*string{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*string, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint32BoolPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint32 output type: *bool
//	2. List
//
// Returns
//	New List of type *bool
//	Empty list if all arguments are nil or either one is nil
func PMapUint32BoolPtr(f func(*uint32) *bool, list []*uint32) []*bool {
	if f == nil {
		return []*bool{}
	}

	ch := make(chan map[int]*bool)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*bool, i int, v *uint32) {
			defer wg.Done()
			ch <- map[int]*bool{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*bool, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint32Float32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint32 output type: *float32
//	2. List
//
// Returns
//	New List of type *float32
//	Empty list if all arguments are nil or either one is nil
func PMapUint32Float32Ptr(f func(*uint32) *float32, list []*uint32) []*float32 {
	if f == nil {
		return []*float32{}
	}

	ch := make(chan map[int]*float32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*float32, i int, v *uint32) {
			defer wg.Done()
			ch <- map[int]*float32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*float32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint32Float64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint32 output type: *float64
//	2. List
//
// Returns
//	New List of type *float64
//	Empty list if all arguments are nil or either one is nil
func PMapUint32Float64Ptr(f func(*uint32) *float64, list []*uint32) []*float64 {
	if f == nil {
		return []*float64{}
	}

	ch := make(chan map[int]*float64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*float64, i int, v *uint32) {
			defer wg.Done()
			ch <- map[int]*float64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*float64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint16IntPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint16 output type: *int
//	2. List
//
// Returns
//	New List of type *int
//	Empty list if all arguments are nil or either one is nil
func PMapUint16IntPtr(f func(*uint16) *int, list []*uint16) []*int {
	if f == nil {
		return []*int{}
	}

	ch := make(chan map[int]*int)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int, i int, v *uint16) {
			defer wg.Done()
			ch <- map[int]*int{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint16Int64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint16 output type: *int64
//	2. List
//
// Returns
//	New List of type *int64
//	Empty list if all arguments are nil or either one is nil
func PMapUint16Int64Ptr(f func(*uint16) *int64, list []*uint16) []*int64 {
	if f == nil {
		return []*int64{}
	}

	ch := make(chan map[int]*int64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int64, i int, v *uint16) {
			defer wg.Done()
			ch <- map[int]*int64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint16Int32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint16 output type: *int32
//	2. List
//
// Returns
//	New List of type *int32
//	Empty list if all arguments are nil or either one is nil
func PMapUint16Int32Ptr(f func(*uint16) *int32, list []*uint16) []*int32 {
	if f == nil {
		return []*int32{}
	}

	ch := make(chan map[int]*int32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int32, i int, v *uint16) {
			defer wg.Done()
			ch <- map[int]*int32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint16Int16Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint16 output type: *int16
//	2. List
//
// Returns
//	New List of type *int16
//	Empty list if all arguments are nil or either one is nil
func PMapUint16Int16Ptr(f func(*uint16) *int16, list []*uint16) []*int16 {
	if f == nil {
		return []*int16{}
	}

	ch := make(chan map[int]*int16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int16, i int, v *uint16) {
			defer wg.Done()
			ch <- map[int]*int16{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint16Int8Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint16 output type: *int8
//	2. List
//
// Returns
//	New List of type *int8
//	Empty list if all arguments are nil or either one is nil
func PMapUint16Int8Ptr(f func(*uint16) *int8, list []*uint16) []*int8 {
	if f == nil {
		return []*int8{}
	}

	ch := make(chan map[int]*int8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int8, i int, v *uint16) {
			defer wg.Done()
			ch <- map[int]*int8{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint16UintPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint16 output type: *uint
//	2. List
//
// Returns
//	New List of type *uint
//	Empty list if all arguments are nil or either one is nil
func PMapUint16UintPtr(f func(*uint16) *uint, list []*uint16) []*uint {
	if f == nil {
		return []*uint{}
	}

	ch := make(chan map[int]*uint)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint, i int, v *uint16) {
			defer wg.Done()
			ch <- map[int]*uint{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint16Uint64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint16 output type: *uint64
//	2. List
//
// Returns
//	New List of type *uint64
//	Empty list if all arguments are nil or either one is nil
func PMapUint16Uint64Ptr(f func(*uint16) *uint64, list []*uint16) []*uint64 {
	if f == nil {
		return []*uint64{}
	}

	ch := make(chan map[int]*uint64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint64, i int, v *uint16) {
			defer wg.Done()
			ch <- map[int]*uint64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint16Uint32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint16 output type: *uint32
//	2. List
//
// Returns
//	New List of type *uint32
//	Empty list if all arguments are nil or either one is nil
func PMapUint16Uint32Ptr(f func(*uint16) *uint32, list []*uint16) []*uint32 {
	if f == nil {
		return []*uint32{}
	}

	ch := make(chan map[int]*uint32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint32, i int, v *uint16) {
			defer wg.Done()
			ch <- map[int]*uint32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint16Uint8Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint16 output type: *uint8
//	2. List
//
// Returns
//	New List of type *uint8
//	Empty list if all arguments are nil or either one is nil
func PMapUint16Uint8Ptr(f func(*uint16) *uint8, list []*uint16) []*uint8 {
	if f == nil {
		return []*uint8{}
	}

	ch := make(chan map[int]*uint8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint8, i int, v *uint16) {
			defer wg.Done()
			ch <- map[int]*uint8{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint16StrPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint16 output type: *string
//	2. List
//
// Returns
//	New List of type *string
//	Empty list if all arguments are nil or either one is nil
func PMapUint16StrPtr(f func(*uint16) *string, list []*uint16) []*string {
	if f == nil {
		return []*string{}
	}

	ch := make(chan map[int]*string)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*string, i int, v *uint16) {
			defer wg.Done()
			ch <- map[int]*string{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*string, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint16BoolPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint16 output type: *bool
//	2. List
//
// Returns
//	New List of type *bool
//	Empty list if all arguments are nil or either one is nil
func PMapUint16BoolPtr(f func(*uint16) *bool, list []*uint16) []*bool {
	if f == nil {
		return []*bool{}
	}

	ch := make(chan map[int]*bool)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*bool, i int, v *uint16) {
			defer wg.Done()
			ch <- map[int]*bool{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*bool, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint16Float32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint16 output type: *float32
//	2. List
//
// Returns
//	New List of type *float32
//	Empty list if all arguments are nil or either one is nil
func PMapUint16Float32Ptr(f func(*uint16) *float32, list []*uint16) []*float32 {
	if f == nil {
		return []*float32{}
	}

	ch := make(chan map[int]*float32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*float32, i int, v *uint16) {
			defer wg.Done()
			ch <- map[int]*float32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*float32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint16Float64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint16 output type: *float64
//	2. List
//
// Returns
//	New List of type *float64
//	Empty list if all arguments are nil or either one is nil
func PMapUint16Float64Ptr(f func(*uint16) *float64, list []*uint16) []*float64 {
	if f == nil {
		return []*float64{}
	}

	ch := make(chan map[int]*float64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*float64, i int, v *uint16) {
			defer wg.Done()
			ch <- map[int]*float64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*float64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint8IntPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint8 output type: *int
//	2. List
//
// Returns
//	New List of type *int
//	Empty list if all arguments are nil or either one is nil
func PMapUint8IntPtr(f func(*uint8) *int, list []*uint8) []*int {
	if f == nil {
		return []*int{}
	}

	ch := make(chan map[int]*int)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int, i int, v *uint8) {
			defer wg.Done()
			ch <- map[int]*int{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint8Int64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint8 output type: *int64
//	2. List
//
// Returns
//	New List of type *int64
//	Empty list if all arguments are nil or either one is nil
func PMapUint8Int64Ptr(f func(*uint8) *int64, list []*uint8) []*int64 {
	if f == nil {
		return []*int64{}
	}

	ch := make(chan map[int]*int64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int64, i int, v *uint8) {
			defer wg.Done()
			ch <- map[int]*int64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint8Int32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint8 output type: *int32
//	2. List
//
// Returns
//	New List of type *int32
//	Empty list if all arguments are nil or either one is nil
func PMapUint8Int32Ptr(f func(*uint8) *int32, list []*uint8) []*int32 {
	if f == nil {
		return []*int32{}
	}

	ch := make(chan map[int]*int32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int32, i int, v *uint8) {
			defer wg.Done()
			ch <- map[int]*int32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint8Int16Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint8 output type: *int16
//	2. List
//
// Returns
//	New List of type *int16
//	Empty list if all arguments are nil or either one is nil
func PMapUint8Int16Ptr(f func(*uint8) *int16, list []*uint8) []*int16 {
	if f == nil {
		return []*int16{}
	}

	ch := make(chan map[int]*int16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int16, i int, v *uint8) {
			defer wg.Done()
			ch <- map[int]*int16{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint8Int8Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint8 output type: *int8
//	2. List
//
// Returns
//	New List of type *int8
//	Empty list if all arguments are nil or either one is nil
func PMapUint8Int8Ptr(f func(*uint8) *int8, list []*uint8) []*int8 {
	if f == nil {
		return []*int8{}
	}

	ch := make(chan map[int]*int8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int8, i int, v *uint8) {
			defer wg.Done()
			ch <- map[int]*int8{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint8UintPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint8 output type: *uint
//	2. List
//
// Returns
//	New List of type *uint
//	Empty list if all arguments are nil or either one is nil
func PMapUint8UintPtr(f func(*uint8) *uint, list []*uint8) []*uint {
	if f == nil {
		return []*uint{}
	}

	ch := make(chan map[int]*uint)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint, i int, v *uint8) {
			defer wg.Done()
			ch <- map[int]*uint{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint8Uint64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint8 output type: *uint64
//	2. List
//
// Returns
//	New List of type *uint64
//	Empty list if all arguments are nil or either one is nil
func PMapUint8Uint64Ptr(f func(*uint8) *uint64, list []*uint8) []*uint64 {
	if f == nil {
		return []*uint64{}
	}

	ch := make(chan map[int]*uint64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint64, i int, v *uint8) {
			defer wg.Done()
			ch <- map[int]*uint64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint8Uint32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint8 output type: *uint32
//	2. List
//
// Returns
//	New List of type *uint32
//	Empty list if all arguments are nil or either one is nil
func PMapUint8Uint32Ptr(f func(*uint8) *uint32, list []*uint8) []*uint32 {
	if f == nil {
		return []*uint32{}
	}

	ch := make(chan map[int]*uint32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint32, i int, v *uint8) {
			defer wg.Done()
			ch <- map[int]*uint32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint8Uint16Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint8 output type: *uint16
//	2. List
//
// Returns
//	New List of type *uint16
//	Empty list if all arguments are nil or either one is nil
func PMapUint8Uint16Ptr(f func(*uint8) *uint16, list []*uint8) []*uint16 {
	if f == nil {
		return []*uint16{}
	}

	ch := make(chan map[int]*uint16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint16, i int, v *uint8) {
			defer wg.Done()
			ch <- map[int]*uint16{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint8StrPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint8 output type: *string
//	2. List
//
// Returns
//	New List of type *string
//	Empty list if all arguments are nil or either one is nil
func PMapUint8StrPtr(f func(*uint8) *string, list []*uint8) []*string {
	if f == nil {
		return []*string{}
	}

	ch := make(chan map[int]*string)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*string, i int, v *uint8) {
			defer wg.Done()
			ch <- map[int]*string{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*string, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint8BoolPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint8 output type: *bool
//	2. List
//
// Returns
//	New List of type *bool
//	Empty list if all arguments are nil or either one is nil
func PMapUint8BoolPtr(f func(*uint8) *bool, list []*uint8) []*bool {
	if f == nil {
		return []*bool{}
	}

	ch := make(chan map[int]*bool)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*bool, i int, v *uint8) {
			defer wg.Done()
			ch <- map[int]*bool{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*bool, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint8Float32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint8 output type: *float32
//	2. List
//
// Returns
//	New List of type *float32
//	Empty list if all arguments are nil or either one is nil
func PMapUint8Float32Ptr(f func(*uint8) *float32, list []*uint8) []*float32 {
	if f == nil {
		return []*float32{}
	}

	ch := make(chan map[int]*float32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*float32, i int, v *uint8) {
			defer wg.Done()
			ch <- map[int]*float32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*float32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapUint8Float64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *uint8 output type: *float64
//	2. List
//
// Returns
//	New List of type *float64
//	Empty list if all arguments are nil or either one is nil
func PMapUint8Float64Ptr(f func(*uint8) *float64, list []*uint8) []*float64 {
	if f == nil {
		return []*float64{}
	}

	ch := make(chan map[int]*float64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*float64, i int, v *uint8) {
			defer wg.Done()
			ch <- map[int]*float64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*float64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapStrIntPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *string output type: *int
//	2. List
//
// Returns
//	New List of type *int
//	Empty list if all arguments are nil or either one is nil
func PMapStrIntPtr(f func(*string) *int, list []*string) []*int {
	if f == nil {
		return []*int{}
	}

	ch := make(chan map[int]*int)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int, i int, v *string) {
			defer wg.Done()
			ch <- map[int]*int{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapStrInt64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *string output type: *int64
//	2. List
//
// Returns
//	New List of type *int64
//	Empty list if all arguments are nil or either one is nil
func PMapStrInt64Ptr(f func(*string) *int64, list []*string) []*int64 {
	if f == nil {
		return []*int64{}
	}

	ch := make(chan map[int]*int64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int64, i int, v *string) {
			defer wg.Done()
			ch <- map[int]*int64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapStrInt32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *string output type: *int32
//	2. List
//
// Returns
//	New List of type *int32
//	Empty list if all arguments are nil or either one is nil
func PMapStrInt32Ptr(f func(*string) *int32, list []*string) []*int32 {
	if f == nil {
		return []*int32{}
	}

	ch := make(chan map[int]*int32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int32, i int, v *string) {
			defer wg.Done()
			ch <- map[int]*int32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapStrInt16Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *string output type: *int16
//	2. List
//
// Returns
//	New List of type *int16
//	Empty list if all arguments are nil or either one is nil
func PMapStrInt16Ptr(f func(*string) *int16, list []*string) []*int16 {
	if f == nil {
		return []*int16{}
	}

	ch := make(chan map[int]*int16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int16, i int, v *string) {
			defer wg.Done()
			ch <- map[int]*int16{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapStrInt8Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *string output type: *int8
//	2. List
//
// Returns
//	New List of type *int8
//	Empty list if all arguments are nil or either one is nil
func PMapStrInt8Ptr(f func(*string) *int8, list []*string) []*int8 {
	if f == nil {
		return []*int8{}
	}

	ch := make(chan map[int]*int8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int8, i int, v *string) {
			defer wg.Done()
			ch <- map[int]*int8{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapStrUintPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *string output type: *uint
//	2. List
//
// Returns
//	New List of type *uint
//	Empty list if all arguments are nil or either one is nil
func PMapStrUintPtr(f func(*string) *uint, list []*string) []*uint {
	if f == nil {
		return []*uint{}
	}

	ch := make(chan map[int]*uint)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint, i int, v *string) {
			defer wg.Done()
			ch <- map[int]*uint{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapStrUint64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *string output type: *uint64
//	2. List
//
// Returns
//	New List of type *uint64
//	Empty list if all arguments are nil or either one is nil
func PMapStrUint64Ptr(f func(*string) *uint64, list []*string) []*uint64 {
	if f == nil {
		return []*uint64{}
	}

	ch := make(chan map[int]*uint64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint64, i int, v *string) {
			defer wg.Done()
			ch <- map[int]*uint64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapStrUint32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *string output type: *uint32
//	2. List
//
// Returns
//	New List of type *uint32
//	Empty list if all arguments are nil or either one is nil
func PMapStrUint32Ptr(f func(*string) *uint32, list []*string) []*uint32 {
	if f == nil {
		return []*uint32{}
	}

	ch := make(chan map[int]*uint32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint32, i int, v *string) {
			defer wg.Done()
			ch <- map[int]*uint32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapStrUint16Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *string output type: *uint16
//	2. List
//
// Returns
//	New List of type *uint16
//	Empty list if all arguments are nil or either one is nil
func PMapStrUint16Ptr(f func(*string) *uint16, list []*string) []*uint16 {
	if f == nil {
		return []*uint16{}
	}

	ch := make(chan map[int]*uint16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint16, i int, v *string) {
			defer wg.Done()
			ch <- map[int]*uint16{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapStrUint8Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *string output type: *uint8
//	2. List
//
// Returns
//	New List of type *uint8
//	Empty list if all arguments are nil or either one is nil
func PMapStrUint8Ptr(f func(*string) *uint8, list []*string) []*uint8 {
	if f == nil {
		return []*uint8{}
	}

	ch := make(chan map[int]*uint8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint8, i int, v *string) {
			defer wg.Done()
			ch <- map[int]*uint8{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapStrBoolPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *string output type: *bool
//	2. List
//
// Returns
//	New List of type *bool
//	Empty list if all arguments are nil or either one is nil
func PMapStrBoolPtr(f func(*string) *bool, list []*string) []*bool {
	if f == nil {
		return []*bool{}
	}

	ch := make(chan map[int]*bool)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*bool, i int, v *string) {
			defer wg.Done()
			ch <- map[int]*bool{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*bool, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapStrFloat32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *string output type: *float32
//	2. List
//
// Returns
//	New List of type *float32
//	Empty list if all arguments are nil or either one is nil
func PMapStrFloat32Ptr(f func(*string) *float32, list []*string) []*float32 {
	if f == nil {
		return []*float32{}
	}

	ch := make(chan map[int]*float32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*float32, i int, v *string) {
			defer wg.Done()
			ch <- map[int]*float32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*float32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapStrFloat64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *string output type: *float64
//	2. List
//
// Returns
//	New List of type *float64
//	Empty list if all arguments are nil or either one is nil
func PMapStrFloat64Ptr(f func(*string) *float64, list []*string) []*float64 {
	if f == nil {
		return []*float64{}
	}

	ch := make(chan map[int]*float64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*float64, i int, v *string) {
			defer wg.Done()
			ch <- map[int]*float64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*float64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapBoolIntPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *bool output type: *int
//	2. List
//
// Returns
//	New List of type *int
//	Empty list if all arguments are nil or either one is nil
func PMapBoolIntPtr(f func(*bool) *int, list []*bool) []*int {
	if f == nil {
		return []*int{}
	}

	ch := make(chan map[int]*int)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int, i int, v *bool) {
			defer wg.Done()
			ch <- map[int]*int{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapBoolInt64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *bool output type: *int64
//	2. List
//
// Returns
//	New List of type *int64
//	Empty list if all arguments are nil or either one is nil
func PMapBoolInt64Ptr(f func(*bool) *int64, list []*bool) []*int64 {
	if f == nil {
		return []*int64{}
	}

	ch := make(chan map[int]*int64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int64, i int, v *bool) {
			defer wg.Done()
			ch <- map[int]*int64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapBoolInt32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *bool output type: *int32
//	2. List
//
// Returns
//	New List of type *int32
//	Empty list if all arguments are nil or either one is nil
func PMapBoolInt32Ptr(f func(*bool) *int32, list []*bool) []*int32 {
	if f == nil {
		return []*int32{}
	}

	ch := make(chan map[int]*int32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int32, i int, v *bool) {
			defer wg.Done()
			ch <- map[int]*int32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapBoolInt16Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *bool output type: *int16
//	2. List
//
// Returns
//	New List of type *int16
//	Empty list if all arguments are nil or either one is nil
func PMapBoolInt16Ptr(f func(*bool) *int16, list []*bool) []*int16 {
	if f == nil {
		return []*int16{}
	}

	ch := make(chan map[int]*int16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int16, i int, v *bool) {
			defer wg.Done()
			ch <- map[int]*int16{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapBoolInt8Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *bool output type: *int8
//	2. List
//
// Returns
//	New List of type *int8
//	Empty list if all arguments are nil or either one is nil
func PMapBoolInt8Ptr(f func(*bool) *int8, list []*bool) []*int8 {
	if f == nil {
		return []*int8{}
	}

	ch := make(chan map[int]*int8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int8, i int, v *bool) {
			defer wg.Done()
			ch <- map[int]*int8{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapBoolUintPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *bool output type: *uint
//	2. List
//
// Returns
//	New List of type *uint
//	Empty list if all arguments are nil or either one is nil
func PMapBoolUintPtr(f func(*bool) *uint, list []*bool) []*uint {
	if f == nil {
		return []*uint{}
	}

	ch := make(chan map[int]*uint)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint, i int, v *bool) {
			defer wg.Done()
			ch <- map[int]*uint{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapBoolUint64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *bool output type: *uint64
//	2. List
//
// Returns
//	New List of type *uint64
//	Empty list if all arguments are nil or either one is nil
func PMapBoolUint64Ptr(f func(*bool) *uint64, list []*bool) []*uint64 {
	if f == nil {
		return []*uint64{}
	}

	ch := make(chan map[int]*uint64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint64, i int, v *bool) {
			defer wg.Done()
			ch <- map[int]*uint64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapBoolUint32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *bool output type: *uint32
//	2. List
//
// Returns
//	New List of type *uint32
//	Empty list if all arguments are nil or either one is nil
func PMapBoolUint32Ptr(f func(*bool) *uint32, list []*bool) []*uint32 {
	if f == nil {
		return []*uint32{}
	}

	ch := make(chan map[int]*uint32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint32, i int, v *bool) {
			defer wg.Done()
			ch <- map[int]*uint32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapBoolUint16Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *bool output type: *uint16
//	2. List
//
// Returns
//	New List of type *uint16
//	Empty list if all arguments are nil or either one is nil
func PMapBoolUint16Ptr(f func(*bool) *uint16, list []*bool) []*uint16 {
	if f == nil {
		return []*uint16{}
	}

	ch := make(chan map[int]*uint16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint16, i int, v *bool) {
			defer wg.Done()
			ch <- map[int]*uint16{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapBoolUint8Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *bool output type: *uint8
//	2. List
//
// Returns
//	New List of type *uint8
//	Empty list if all arguments are nil or either one is nil
func PMapBoolUint8Ptr(f func(*bool) *uint8, list []*bool) []*uint8 {
	if f == nil {
		return []*uint8{}
	}

	ch := make(chan map[int]*uint8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint8, i int, v *bool) {
			defer wg.Done()
			ch <- map[int]*uint8{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapBoolStrPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *bool output type: *string
//	2. List
//
// Returns
//	New List of type *string
//	Empty list if all arguments are nil or either one is nil
func PMapBoolStrPtr(f func(*bool) *string, list []*bool) []*string {
	if f == nil {
		return []*string{}
	}

	ch := make(chan map[int]*string)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*string, i int, v *bool) {
			defer wg.Done()
			ch <- map[int]*string{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*string, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapBoolFloat32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *bool output type: *float32
//	2. List
//
// Returns
//	New List of type *float32
//	Empty list if all arguments are nil or either one is nil
func PMapBoolFloat32Ptr(f func(*bool) *float32, list []*bool) []*float32 {
	if f == nil {
		return []*float32{}
	}

	ch := make(chan map[int]*float32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*float32, i int, v *bool) {
			defer wg.Done()
			ch <- map[int]*float32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*float32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapBoolFloat64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *bool output type: *float64
//	2. List
//
// Returns
//	New List of type *float64
//	Empty list if all arguments are nil or either one is nil
func PMapBoolFloat64Ptr(f func(*bool) *float64, list []*bool) []*float64 {
	if f == nil {
		return []*float64{}
	}

	ch := make(chan map[int]*float64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*float64, i int, v *bool) {
			defer wg.Done()
			ch <- map[int]*float64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*float64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapFloat32IntPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *float32 output type: *int
//	2. List
//
// Returns
//	New List of type *int
//	Empty list if all arguments are nil or either one is nil
func PMapFloat32IntPtr(f func(*float32) *int, list []*float32) []*int {
	if f == nil {
		return []*int{}
	}

	ch := make(chan map[int]*int)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int, i int, v *float32) {
			defer wg.Done()
			ch <- map[int]*int{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapFloat32Int64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *float32 output type: *int64
//	2. List
//
// Returns
//	New List of type *int64
//	Empty list if all arguments are nil or either one is nil
func PMapFloat32Int64Ptr(f func(*float32) *int64, list []*float32) []*int64 {
	if f == nil {
		return []*int64{}
	}

	ch := make(chan map[int]*int64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int64, i int, v *float32) {
			defer wg.Done()
			ch <- map[int]*int64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapFloat32Int32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *float32 output type: *int32
//	2. List
//
// Returns
//	New List of type *int32
//	Empty list if all arguments are nil or either one is nil
func PMapFloat32Int32Ptr(f func(*float32) *int32, list []*float32) []*int32 {
	if f == nil {
		return []*int32{}
	}

	ch := make(chan map[int]*int32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int32, i int, v *float32) {
			defer wg.Done()
			ch <- map[int]*int32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapFloat32Int16Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *float32 output type: *int16
//	2. List
//
// Returns
//	New List of type *int16
//	Empty list if all arguments are nil or either one is nil
func PMapFloat32Int16Ptr(f func(*float32) *int16, list []*float32) []*int16 {
	if f == nil {
		return []*int16{}
	}

	ch := make(chan map[int]*int16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int16, i int, v *float32) {
			defer wg.Done()
			ch <- map[int]*int16{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapFloat32Int8Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *float32 output type: *int8
//	2. List
//
// Returns
//	New List of type *int8
//	Empty list if all arguments are nil or either one is nil
func PMapFloat32Int8Ptr(f func(*float32) *int8, list []*float32) []*int8 {
	if f == nil {
		return []*int8{}
	}

	ch := make(chan map[int]*int8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int8, i int, v *float32) {
			defer wg.Done()
			ch <- map[int]*int8{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapFloat32UintPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *float32 output type: *uint
//	2. List
//
// Returns
//	New List of type *uint
//	Empty list if all arguments are nil or either one is nil
func PMapFloat32UintPtr(f func(*float32) *uint, list []*float32) []*uint {
	if f == nil {
		return []*uint{}
	}

	ch := make(chan map[int]*uint)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint, i int, v *float32) {
			defer wg.Done()
			ch <- map[int]*uint{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapFloat32Uint64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *float32 output type: *uint64
//	2. List
//
// Returns
//	New List of type *uint64
//	Empty list if all arguments are nil or either one is nil
func PMapFloat32Uint64Ptr(f func(*float32) *uint64, list []*float32) []*uint64 {
	if f == nil {
		return []*uint64{}
	}

	ch := make(chan map[int]*uint64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint64, i int, v *float32) {
			defer wg.Done()
			ch <- map[int]*uint64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapFloat32Uint32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *float32 output type: *uint32
//	2. List
//
// Returns
//	New List of type *uint32
//	Empty list if all arguments are nil or either one is nil
func PMapFloat32Uint32Ptr(f func(*float32) *uint32, list []*float32) []*uint32 {
	if f == nil {
		return []*uint32{}
	}

	ch := make(chan map[int]*uint32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint32, i int, v *float32) {
			defer wg.Done()
			ch <- map[int]*uint32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapFloat32Uint16Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *float32 output type: *uint16
//	2. List
//
// Returns
//	New List of type *uint16
//	Empty list if all arguments are nil or either one is nil
func PMapFloat32Uint16Ptr(f func(*float32) *uint16, list []*float32) []*uint16 {
	if f == nil {
		return []*uint16{}
	}

	ch := make(chan map[int]*uint16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint16, i int, v *float32) {
			defer wg.Done()
			ch <- map[int]*uint16{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapFloat32Uint8Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *float32 output type: *uint8
//	2. List
//
// Returns
//	New List of type *uint8
//	Empty list if all arguments are nil or either one is nil
func PMapFloat32Uint8Ptr(f func(*float32) *uint8, list []*float32) []*uint8 {
	if f == nil {
		return []*uint8{}
	}

	ch := make(chan map[int]*uint8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint8, i int, v *float32) {
			defer wg.Done()
			ch <- map[int]*uint8{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapFloat32StrPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *float32 output type: *string
//	2. List
//
// Returns
//	New List of type *string
//	Empty list if all arguments are nil or either one is nil
func PMapFloat32StrPtr(f func(*float32) *string, list []*float32) []*string {
	if f == nil {
		return []*string{}
	}

	ch := make(chan map[int]*string)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*string, i int, v *float32) {
			defer wg.Done()
			ch <- map[int]*string{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*string, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapFloat32BoolPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *float32 output type: *bool
//	2. List
//
// Returns
//	New List of type *bool
//	Empty list if all arguments are nil or either one is nil
func PMapFloat32BoolPtr(f func(*float32) *bool, list []*float32) []*bool {
	if f == nil {
		return []*bool{}
	}

	ch := make(chan map[int]*bool)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*bool, i int, v *float32) {
			defer wg.Done()
			ch <- map[int]*bool{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*bool, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapFloat32Float64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *float32 output type: *float64
//	2. List
//
// Returns
//	New List of type *float64
//	Empty list if all arguments are nil or either one is nil
func PMapFloat32Float64Ptr(f func(*float32) *float64, list []*float32) []*float64 {
	if f == nil {
		return []*float64{}
	}

	ch := make(chan map[int]*float64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*float64, i int, v *float32) {
			defer wg.Done()
			ch <- map[int]*float64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*float64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapFloat64IntPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *float64 output type: *int
//	2. List
//
// Returns
//	New List of type *int
//	Empty list if all arguments are nil or either one is nil
func PMapFloat64IntPtr(f func(*float64) *int, list []*float64) []*int {
	if f == nil {
		return []*int{}
	}

	ch := make(chan map[int]*int)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int, i int, v *float64) {
			defer wg.Done()
			ch <- map[int]*int{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapFloat64Int64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *float64 output type: *int64
//	2. List
//
// Returns
//	New List of type *int64
//	Empty list if all arguments are nil or either one is nil
func PMapFloat64Int64Ptr(f func(*float64) *int64, list []*float64) []*int64 {
	if f == nil {
		return []*int64{}
	}

	ch := make(chan map[int]*int64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int64, i int, v *float64) {
			defer wg.Done()
			ch <- map[int]*int64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapFloat64Int32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *float64 output type: *int32
//	2. List
//
// Returns
//	New List of type *int32
//	Empty list if all arguments are nil or either one is nil
func PMapFloat64Int32Ptr(f func(*float64) *int32, list []*float64) []*int32 {
	if f == nil {
		return []*int32{}
	}

	ch := make(chan map[int]*int32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int32, i int, v *float64) {
			defer wg.Done()
			ch <- map[int]*int32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapFloat64Int16Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *float64 output type: *int16
//	2. List
//
// Returns
//	New List of type *int16
//	Empty list if all arguments are nil or either one is nil
func PMapFloat64Int16Ptr(f func(*float64) *int16, list []*float64) []*int16 {
	if f == nil {
		return []*int16{}
	}

	ch := make(chan map[int]*int16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int16, i int, v *float64) {
			defer wg.Done()
			ch <- map[int]*int16{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapFloat64Int8Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *float64 output type: *int8
//	2. List
//
// Returns
//	New List of type *int8
//	Empty list if all arguments are nil or either one is nil
func PMapFloat64Int8Ptr(f func(*float64) *int8, list []*float64) []*int8 {
	if f == nil {
		return []*int8{}
	}

	ch := make(chan map[int]*int8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int8, i int, v *float64) {
			defer wg.Done()
			ch <- map[int]*int8{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*int8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapFloat64UintPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *float64 output type: *uint
//	2. List
//
// Returns
//	New List of type *uint
//	Empty list if all arguments are nil or either one is nil
func PMapFloat64UintPtr(f func(*float64) *uint, list []*float64) []*uint {
	if f == nil {
		return []*uint{}
	}

	ch := make(chan map[int]*uint)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint, i int, v *float64) {
			defer wg.Done()
			ch <- map[int]*uint{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapFloat64Uint64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *float64 output type: *uint64
//	2. List
//
// Returns
//	New List of type *uint64
//	Empty list if all arguments are nil or either one is nil
func PMapFloat64Uint64Ptr(f func(*float64) *uint64, list []*float64) []*uint64 {
	if f == nil {
		return []*uint64{}
	}

	ch := make(chan map[int]*uint64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint64, i int, v *float64) {
			defer wg.Done()
			ch <- map[int]*uint64{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapFloat64Uint32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *float64 output type: *uint32
//	2. List
//
// Returns
//	New List of type *uint32
//	Empty list if all arguments are nil or either one is nil
func PMapFloat64Uint32Ptr(f func(*float64) *uint32, list []*float64) []*uint32 {
	if f == nil {
		return []*uint32{}
	}

	ch := make(chan map[int]*uint32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint32, i int, v *float64) {
			defer wg.Done()
			ch <- map[int]*uint32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapFloat64Uint16Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *float64 output type: *uint16
//	2. List
//
// Returns
//	New List of type *uint16
//	Empty list if all arguments are nil or either one is nil
func PMapFloat64Uint16Ptr(f func(*float64) *uint16, list []*float64) []*uint16 {
	if f == nil {
		return []*uint16{}
	}

	ch := make(chan map[int]*uint16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint16, i int, v *float64) {
			defer wg.Done()
			ch <- map[int]*uint16{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapFloat64Uint8Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *float64 output type: *uint8
//	2. List
//
// Returns
//	New List of type *uint8
//	Empty list if all arguments are nil or either one is nil
func PMapFloat64Uint8Ptr(f func(*float64) *uint8, list []*float64) []*uint8 {
	if f == nil {
		return []*uint8{}
	}

	ch := make(chan map[int]*uint8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint8, i int, v *float64) {
			defer wg.Done()
			ch <- map[int]*uint8{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*uint8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapFloat64StrPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *float64 output type: *string
//	2. List
//
// Returns
//	New List of type *string
//	Empty list if all arguments are nil or either one is nil
func PMapFloat64StrPtr(f func(*float64) *string, list []*float64) []*string {
	if f == nil {
		return []*string{}
	}

	ch := make(chan map[int]*string)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*string, i int, v *float64) {
			defer wg.Done()
			ch <- map[int]*string{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*string, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapFloat64BoolPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *float64 output type: *bool
//	2. List
//
// Returns
//	New List of type *bool
//	Empty list if all arguments are nil or either one is nil
func PMapFloat64BoolPtr(f func(*float64) *bool, list []*float64) []*bool {
	if f == nil {
		return []*bool{}
	}

	ch := make(chan map[int]*bool)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*bool, i int, v *float64) {
			defer wg.Done()
			ch <- map[int]*bool{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*bool, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// PMapFloat64Float32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: *float64 output type: *float32
//	2. List
//
// Returns
//	New List of type *float32
//	Empty list if all arguments are nil or either one is nil
func PMapFloat64Float32Ptr(f func(*float64) *float32, list []*float64) []*float32 {
	if f == nil {
		return []*float32{}
	}

	ch := make(chan map[int]*float32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*float32, i int, v *float64) {
			defer wg.Done()
			ch <- map[int]*float32{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*float32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}
