package fp

import "sync"

// PMapIntPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
func PMapIntPtr(f func(*int) *int, list []*int) []*int {
	if f == nil {
		return []*int{}
	}

	ch := make(chan map[int]*int)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int, i int, v *int) {
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

// PMapInt64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
func PMapInt64Ptr(f func(*int64) *int64, list []*int64) []*int64 {
	if f == nil {
		return []*int64{}
	}

	ch := make(chan map[int]*int64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int64, i int, v *int64) {
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

// PMapInt32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
func PMapInt32Ptr(f func(*int32) *int32, list []*int32) []*int32 {
	if f == nil {
		return []*int32{}
	}

	ch := make(chan map[int]*int32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int32, i int, v *int32) {
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

// PMapInt16Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
func PMapInt16Ptr(f func(*int16) *int16, list []*int16) []*int16 {
	if f == nil {
		return []*int16{}
	}

	ch := make(chan map[int]*int16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int16, i int, v *int16) {
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

// PMapInt8Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
func PMapInt8Ptr(f func(*int8) *int8, list []*int8) []*int8 {
	if f == nil {
		return []*int8{}
	}

	ch := make(chan map[int]*int8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*int8, i int, v *int8) {
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

// PMapUintPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
func PMapUintPtr(f func(*uint) *uint, list []*uint) []*uint {
	if f == nil {
		return []*uint{}
	}

	ch := make(chan map[int]*uint)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint, i int, v *uint) {
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

// PMapUint64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
func PMapUint64Ptr(f func(*uint64) *uint64, list []*uint64) []*uint64 {
	if f == nil {
		return []*uint64{}
	}

	ch := make(chan map[int]*uint64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint64, i int, v *uint64) {
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

// PMapUint32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
func PMapUint32Ptr(f func(*uint32) *uint32, list []*uint32) []*uint32 {
	if f == nil {
		return []*uint32{}
	}

	ch := make(chan map[int]*uint32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint32, i int, v *uint32) {
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

// PMapUint16Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
func PMapUint16Ptr(f func(*uint16) *uint16, list []*uint16) []*uint16 {
	if f == nil {
		return []*uint16{}
	}

	ch := make(chan map[int]*uint16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint16, i int, v *uint16) {
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

// PMapUint8Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
func PMapUint8Ptr(f func(*uint8) *uint8, list []*uint8) []*uint8 {
	if f == nil {
		return []*uint8{}
	}

	ch := make(chan map[int]*uint8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*uint8, i int, v *uint8) {
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

// PMapStrPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
func PMapStrPtr(f func(*string) *string, list []*string) []*string {
	if f == nil {
		return []*string{}
	}

	ch := make(chan map[int]*string)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*string, i int, v *string) {
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

// PMapBoolPtr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
func PMapBoolPtr(f func(*bool) *bool, list []*bool) []*bool {
	if f == nil {
		return []*bool{}
	}

	ch := make(chan map[int]*bool)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*bool, i int, v *bool) {
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

// PMapFloat32Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
func PMapFloat32Ptr(f func(*float32) *float32, list []*float32) []*float32 {
	if f == nil {
		return []*float32{}
	}

	ch := make(chan map[int]*float32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*float32, i int, v *float32) {
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

// PMapFloat64Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
func PMapFloat64Ptr(f func(*float64) *float64, list []*float64) []*float64 {
	if f == nil {
		return []*float64{}
	}

	ch := make(chan map[int]*float64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*float64, i int, v *float64) {
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
