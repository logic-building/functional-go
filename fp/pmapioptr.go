package fp

import "sync"

// PMapIntInt64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntInt64Ptr(f func(*int) *int64, list []*int, optional ...Optional) []*int64 {
	if f == nil {
		return []*int64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntInt64PtrNoOrder(f, list, worker)
		}
	}

	return pMapIntInt64PtrPreserveOrder(f, list, worker)
}

func pMapIntInt64PtrPreserveOrder(f func(*int) *int64, list []*int, worker int) []*int64 {
	chJobs := make(chan map[int]*int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int64, chJobs chan map[int]*int) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int64, len(list))
	newList := make([]*int64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapIntInt64PtrNoOrder(f func(*int) *int64, list []*int, worker int) []*int64 {
	chJobs := make(chan *int, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int64, chJobs chan *int) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapIntInt32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntInt32Ptr(f func(*int) *int32, list []*int, optional ...Optional) []*int32 {
	if f == nil {
		return []*int32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntInt32PtrNoOrder(f, list, worker)
		}
	}

	return pMapIntInt32PtrPreserveOrder(f, list, worker)
}

func pMapIntInt32PtrPreserveOrder(f func(*int) *int32, list []*int, worker int) []*int32 {
	chJobs := make(chan map[int]*int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int32, chJobs chan map[int]*int) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int32, len(list))
	newList := make([]*int32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapIntInt32PtrNoOrder(f func(*int) *int32, list []*int, worker int) []*int32 {
	chJobs := make(chan *int, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int32, chJobs chan *int) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapIntInt16Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntInt16Ptr(f func(*int) *int16, list []*int, optional ...Optional) []*int16 {
	if f == nil {
		return []*int16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntInt16PtrNoOrder(f, list, worker)
		}
	}

	return pMapIntInt16PtrPreserveOrder(f, list, worker)
}

func pMapIntInt16PtrPreserveOrder(f func(*int) *int16, list []*int, worker int) []*int16 {
	chJobs := make(chan map[int]*int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int16, chJobs chan map[int]*int) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int16{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int16, len(list))
	newList := make([]*int16, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapIntInt16PtrNoOrder(f func(*int) *int16, list []*int, worker int) []*int16 {
	chJobs := make(chan *int, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int16, chJobs chan *int) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int16, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapIntInt8Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntInt8Ptr(f func(*int) *int8, list []*int, optional ...Optional) []*int8 {
	if f == nil {
		return []*int8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntInt8PtrNoOrder(f, list, worker)
		}
	}

	return pMapIntInt8PtrPreserveOrder(f, list, worker)
}

func pMapIntInt8PtrPreserveOrder(f func(*int) *int8, list []*int, worker int) []*int8 {
	chJobs := make(chan map[int]*int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int8, chJobs chan map[int]*int) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int8{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int8, len(list))
	newList := make([]*int8, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapIntInt8PtrNoOrder(f func(*int) *int8, list []*int, worker int) []*int8 {
	chJobs := make(chan *int, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int8, chJobs chan *int) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int8, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapIntUintPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntUintPtr(f func(*int) *uint, list []*int, optional ...Optional) []*uint {
	if f == nil {
		return []*uint{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntUintPtrNoOrder(f, list, worker)
		}
	}

	return pMapIntUintPtrPreserveOrder(f, list, worker)
}

func pMapIntUintPtrPreserveOrder(f func(*int) *uint, list []*int, worker int) []*uint {
	chJobs := make(chan map[int]*int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint, chJobs chan map[int]*int) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint, len(list))
	newList := make([]*uint, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapIntUintPtrNoOrder(f func(*int) *uint, list []*int, worker int) []*uint {
	chJobs := make(chan *int, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint, chJobs chan *int) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapIntUint64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntUint64Ptr(f func(*int) *uint64, list []*int, optional ...Optional) []*uint64 {
	if f == nil {
		return []*uint64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntUint64PtrNoOrder(f, list, worker)
		}
	}

	return pMapIntUint64PtrPreserveOrder(f, list, worker)
}

func pMapIntUint64PtrPreserveOrder(f func(*int) *uint64, list []*int, worker int) []*uint64 {
	chJobs := make(chan map[int]*int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint64, chJobs chan map[int]*int) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint64, len(list))
	newList := make([]*uint64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapIntUint64PtrNoOrder(f func(*int) *uint64, list []*int, worker int) []*uint64 {
	chJobs := make(chan *int, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint64, chJobs chan *int) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapIntUint32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntUint32Ptr(f func(*int) *uint32, list []*int, optional ...Optional) []*uint32 {
	if f == nil {
		return []*uint32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntUint32PtrNoOrder(f, list, worker)
		}
	}

	return pMapIntUint32PtrPreserveOrder(f, list, worker)
}

func pMapIntUint32PtrPreserveOrder(f func(*int) *uint32, list []*int, worker int) []*uint32 {
	chJobs := make(chan map[int]*int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint32, chJobs chan map[int]*int) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint32, len(list))
	newList := make([]*uint32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapIntUint32PtrNoOrder(f func(*int) *uint32, list []*int, worker int) []*uint32 {
	chJobs := make(chan *int, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint32, chJobs chan *int) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapIntUint16Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntUint16Ptr(f func(*int) *uint16, list []*int, optional ...Optional) []*uint16 {
	if f == nil {
		return []*uint16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntUint16PtrNoOrder(f, list, worker)
		}
	}

	return pMapIntUint16PtrPreserveOrder(f, list, worker)
}

func pMapIntUint16PtrPreserveOrder(f func(*int) *uint16, list []*int, worker int) []*uint16 {
	chJobs := make(chan map[int]*int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint16, chJobs chan map[int]*int) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint16{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint16, len(list))
	newList := make([]*uint16, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapIntUint16PtrNoOrder(f func(*int) *uint16, list []*int, worker int) []*uint16 {
	chJobs := make(chan *int, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint16, chJobs chan *int) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint16, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapIntUint8Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntUint8Ptr(f func(*int) *uint8, list []*int, optional ...Optional) []*uint8 {
	if f == nil {
		return []*uint8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntUint8PtrNoOrder(f, list, worker)
		}
	}

	return pMapIntUint8PtrPreserveOrder(f, list, worker)
}

func pMapIntUint8PtrPreserveOrder(f func(*int) *uint8, list []*int, worker int) []*uint8 {
	chJobs := make(chan map[int]*int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint8, chJobs chan map[int]*int) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint8{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint8, len(list))
	newList := make([]*uint8, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapIntUint8PtrNoOrder(f func(*int) *uint8, list []*int, worker int) []*uint8 {
	chJobs := make(chan *int, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint8, chJobs chan *int) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint8, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapIntStrPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntStrPtr(f func(*int) *string, list []*int, optional ...Optional) []*string {
	if f == nil {
		return []*string{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntStrPtrNoOrder(f, list, worker)
		}
	}

	return pMapIntStrPtrPreserveOrder(f, list, worker)
}

func pMapIntStrPtrPreserveOrder(f func(*int) *string, list []*int, worker int) []*string {
	chJobs := make(chan map[int]*int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*string, chJobs chan map[int]*int) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*string{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*string, len(list))
	newList := make([]*string, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapIntStrPtrNoOrder(f func(*int) *string, list []*int, worker int) []*string {
	chJobs := make(chan *int, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *string, chJobs chan *int) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*string, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapIntBoolPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntBoolPtr(f func(*int) *bool, list []*int, optional ...Optional) []*bool {
	if f == nil {
		return []*bool{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntBoolPtrNoOrder(f, list, worker)
		}
	}

	return pMapIntBoolPtrPreserveOrder(f, list, worker)
}

func pMapIntBoolPtrPreserveOrder(f func(*int) *bool, list []*int, worker int) []*bool {
	chJobs := make(chan map[int]*int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*bool, chJobs chan map[int]*int) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*bool{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*bool, len(list))
	newList := make([]*bool, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapIntBoolPtrNoOrder(f func(*int) *bool, list []*int, worker int) []*bool {
	chJobs := make(chan *int, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *bool, chJobs chan *int) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*bool, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapIntFloat32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntFloat32Ptr(f func(*int) *float32, list []*int, optional ...Optional) []*float32 {
	if f == nil {
		return []*float32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntFloat32PtrNoOrder(f, list, worker)
		}
	}

	return pMapIntFloat32PtrPreserveOrder(f, list, worker)
}

func pMapIntFloat32PtrPreserveOrder(f func(*int) *float32, list []*int, worker int) []*float32 {
	chJobs := make(chan map[int]*int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float32, chJobs chan map[int]*int) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*float32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float32, len(list))
	newList := make([]*float32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapIntFloat32PtrNoOrder(f func(*int) *float32, list []*int, worker int) []*float32 {
	chJobs := make(chan *int, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float32, chJobs chan *int) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*float32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapIntFloat64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntFloat64Ptr(f func(*int) *float64, list []*int, optional ...Optional) []*float64 {
	if f == nil {
		return []*float64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntFloat64PtrNoOrder(f, list, worker)
		}
	}

	return pMapIntFloat64PtrPreserveOrder(f, list, worker)
}

func pMapIntFloat64PtrPreserveOrder(f func(*int) *float64, list []*int, worker int) []*float64 {
	chJobs := make(chan map[int]*int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float64, chJobs chan map[int]*int) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*float64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float64, len(list))
	newList := make([]*float64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapIntFloat64PtrNoOrder(f func(*int) *float64, list []*int, worker int) []*float64 {
	chJobs := make(chan *int, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float64, chJobs chan *int) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*float64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt64IntPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64IntPtr(f func(*int64) *int, list []*int64, optional ...Optional) []*int {
	if f == nil {
		return []*int{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64IntPtrNoOrder(f, list, worker)
		}
	}

	return pMapInt64IntPtrPreserveOrder(f, list, worker)
}

func pMapInt64IntPtrPreserveOrder(f func(*int64) *int, list []*int64, worker int) []*int {
	chJobs := make(chan map[int]*int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int, chJobs chan map[int]*int64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int, len(list))
	newList := make([]*int, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt64IntPtrNoOrder(f func(*int64) *int, list []*int64, worker int) []*int {
	chJobs := make(chan *int64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int, chJobs chan *int64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt64Int32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Int32Ptr(f func(*int64) *int32, list []*int64, optional ...Optional) []*int32 {
	if f == nil {
		return []*int32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64Int32PtrNoOrder(f, list, worker)
		}
	}

	return pMapInt64Int32PtrPreserveOrder(f, list, worker)
}

func pMapInt64Int32PtrPreserveOrder(f func(*int64) *int32, list []*int64, worker int) []*int32 {
	chJobs := make(chan map[int]*int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int32, chJobs chan map[int]*int64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int32, len(list))
	newList := make([]*int32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt64Int32PtrNoOrder(f func(*int64) *int32, list []*int64, worker int) []*int32 {
	chJobs := make(chan *int64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int32, chJobs chan *int64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt64Int16Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Int16Ptr(f func(*int64) *int16, list []*int64, optional ...Optional) []*int16 {
	if f == nil {
		return []*int16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64Int16PtrNoOrder(f, list, worker)
		}
	}

	return pMapInt64Int16PtrPreserveOrder(f, list, worker)
}

func pMapInt64Int16PtrPreserveOrder(f func(*int64) *int16, list []*int64, worker int) []*int16 {
	chJobs := make(chan map[int]*int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int16, chJobs chan map[int]*int64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int16{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int16, len(list))
	newList := make([]*int16, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt64Int16PtrNoOrder(f func(*int64) *int16, list []*int64, worker int) []*int16 {
	chJobs := make(chan *int64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int16, chJobs chan *int64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int16, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt64Int8Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Int8Ptr(f func(*int64) *int8, list []*int64, optional ...Optional) []*int8 {
	if f == nil {
		return []*int8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64Int8PtrNoOrder(f, list, worker)
		}
	}

	return pMapInt64Int8PtrPreserveOrder(f, list, worker)
}

func pMapInt64Int8PtrPreserveOrder(f func(*int64) *int8, list []*int64, worker int) []*int8 {
	chJobs := make(chan map[int]*int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int8, chJobs chan map[int]*int64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int8{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int8, len(list))
	newList := make([]*int8, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt64Int8PtrNoOrder(f func(*int64) *int8, list []*int64, worker int) []*int8 {
	chJobs := make(chan *int64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int8, chJobs chan *int64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int8, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt64UintPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64UintPtr(f func(*int64) *uint, list []*int64, optional ...Optional) []*uint {
	if f == nil {
		return []*uint{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64UintPtrNoOrder(f, list, worker)
		}
	}

	return pMapInt64UintPtrPreserveOrder(f, list, worker)
}

func pMapInt64UintPtrPreserveOrder(f func(*int64) *uint, list []*int64, worker int) []*uint {
	chJobs := make(chan map[int]*int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint, chJobs chan map[int]*int64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint, len(list))
	newList := make([]*uint, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt64UintPtrNoOrder(f func(*int64) *uint, list []*int64, worker int) []*uint {
	chJobs := make(chan *int64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint, chJobs chan *int64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt64Uint64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Uint64Ptr(f func(*int64) *uint64, list []*int64, optional ...Optional) []*uint64 {
	if f == nil {
		return []*uint64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64Uint64PtrNoOrder(f, list, worker)
		}
	}

	return pMapInt64Uint64PtrPreserveOrder(f, list, worker)
}

func pMapInt64Uint64PtrPreserveOrder(f func(*int64) *uint64, list []*int64, worker int) []*uint64 {
	chJobs := make(chan map[int]*int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint64, chJobs chan map[int]*int64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint64, len(list))
	newList := make([]*uint64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt64Uint64PtrNoOrder(f func(*int64) *uint64, list []*int64, worker int) []*uint64 {
	chJobs := make(chan *int64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint64, chJobs chan *int64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt64Uint32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Uint32Ptr(f func(*int64) *uint32, list []*int64, optional ...Optional) []*uint32 {
	if f == nil {
		return []*uint32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64Uint32PtrNoOrder(f, list, worker)
		}
	}

	return pMapInt64Uint32PtrPreserveOrder(f, list, worker)
}

func pMapInt64Uint32PtrPreserveOrder(f func(*int64) *uint32, list []*int64, worker int) []*uint32 {
	chJobs := make(chan map[int]*int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint32, chJobs chan map[int]*int64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint32, len(list))
	newList := make([]*uint32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt64Uint32PtrNoOrder(f func(*int64) *uint32, list []*int64, worker int) []*uint32 {
	chJobs := make(chan *int64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint32, chJobs chan *int64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt64Uint16Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Uint16Ptr(f func(*int64) *uint16, list []*int64, optional ...Optional) []*uint16 {
	if f == nil {
		return []*uint16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64Uint16PtrNoOrder(f, list, worker)
		}
	}

	return pMapInt64Uint16PtrPreserveOrder(f, list, worker)
}

func pMapInt64Uint16PtrPreserveOrder(f func(*int64) *uint16, list []*int64, worker int) []*uint16 {
	chJobs := make(chan map[int]*int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint16, chJobs chan map[int]*int64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint16{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint16, len(list))
	newList := make([]*uint16, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt64Uint16PtrNoOrder(f func(*int64) *uint16, list []*int64, worker int) []*uint16 {
	chJobs := make(chan *int64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint16, chJobs chan *int64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint16, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt64Uint8Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Uint8Ptr(f func(*int64) *uint8, list []*int64, optional ...Optional) []*uint8 {
	if f == nil {
		return []*uint8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64Uint8PtrNoOrder(f, list, worker)
		}
	}

	return pMapInt64Uint8PtrPreserveOrder(f, list, worker)
}

func pMapInt64Uint8PtrPreserveOrder(f func(*int64) *uint8, list []*int64, worker int) []*uint8 {
	chJobs := make(chan map[int]*int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint8, chJobs chan map[int]*int64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint8{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint8, len(list))
	newList := make([]*uint8, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt64Uint8PtrNoOrder(f func(*int64) *uint8, list []*int64, worker int) []*uint8 {
	chJobs := make(chan *int64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint8, chJobs chan *int64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint8, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt64StrPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64StrPtr(f func(*int64) *string, list []*int64, optional ...Optional) []*string {
	if f == nil {
		return []*string{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64StrPtrNoOrder(f, list, worker)
		}
	}

	return pMapInt64StrPtrPreserveOrder(f, list, worker)
}

func pMapInt64StrPtrPreserveOrder(f func(*int64) *string, list []*int64, worker int) []*string {
	chJobs := make(chan map[int]*int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*string, chJobs chan map[int]*int64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*string{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*string, len(list))
	newList := make([]*string, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt64StrPtrNoOrder(f func(*int64) *string, list []*int64, worker int) []*string {
	chJobs := make(chan *int64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *string, chJobs chan *int64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*string, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt64BoolPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64BoolPtr(f func(*int64) *bool, list []*int64, optional ...Optional) []*bool {
	if f == nil {
		return []*bool{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64BoolPtrNoOrder(f, list, worker)
		}
	}

	return pMapInt64BoolPtrPreserveOrder(f, list, worker)
}

func pMapInt64BoolPtrPreserveOrder(f func(*int64) *bool, list []*int64, worker int) []*bool {
	chJobs := make(chan map[int]*int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*bool, chJobs chan map[int]*int64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*bool{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*bool, len(list))
	newList := make([]*bool, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt64BoolPtrNoOrder(f func(*int64) *bool, list []*int64, worker int) []*bool {
	chJobs := make(chan *int64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *bool, chJobs chan *int64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*bool, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt64Float32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Float32Ptr(f func(*int64) *float32, list []*int64, optional ...Optional) []*float32 {
	if f == nil {
		return []*float32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64Float32PtrNoOrder(f, list, worker)
		}
	}

	return pMapInt64Float32PtrPreserveOrder(f, list, worker)
}

func pMapInt64Float32PtrPreserveOrder(f func(*int64) *float32, list []*int64, worker int) []*float32 {
	chJobs := make(chan map[int]*int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float32, chJobs chan map[int]*int64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*float32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float32, len(list))
	newList := make([]*float32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt64Float32PtrNoOrder(f func(*int64) *float32, list []*int64, worker int) []*float32 {
	chJobs := make(chan *int64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float32, chJobs chan *int64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*float32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt64Float64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Float64Ptr(f func(*int64) *float64, list []*int64, optional ...Optional) []*float64 {
	if f == nil {
		return []*float64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64Float64PtrNoOrder(f, list, worker)
		}
	}

	return pMapInt64Float64PtrPreserveOrder(f, list, worker)
}

func pMapInt64Float64PtrPreserveOrder(f func(*int64) *float64, list []*int64, worker int) []*float64 {
	chJobs := make(chan map[int]*int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float64, chJobs chan map[int]*int64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*float64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float64, len(list))
	newList := make([]*float64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt64Float64PtrNoOrder(f func(*int64) *float64, list []*int64, worker int) []*float64 {
	chJobs := make(chan *int64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float64, chJobs chan *int64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*float64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt32IntPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32IntPtr(f func(*int32) *int, list []*int32, optional ...Optional) []*int {
	if f == nil {
		return []*int{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32IntPtrNoOrder(f, list, worker)
		}
	}

	return pMapInt32IntPtrPreserveOrder(f, list, worker)
}

func pMapInt32IntPtrPreserveOrder(f func(*int32) *int, list []*int32, worker int) []*int {
	chJobs := make(chan map[int]*int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int, chJobs chan map[int]*int32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int, len(list))
	newList := make([]*int, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt32IntPtrNoOrder(f func(*int32) *int, list []*int32, worker int) []*int {
	chJobs := make(chan *int32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int, chJobs chan *int32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt32Int64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Int64Ptr(f func(*int32) *int64, list []*int32, optional ...Optional) []*int64 {
	if f == nil {
		return []*int64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32Int64PtrNoOrder(f, list, worker)
		}
	}

	return pMapInt32Int64PtrPreserveOrder(f, list, worker)
}

func pMapInt32Int64PtrPreserveOrder(f func(*int32) *int64, list []*int32, worker int) []*int64 {
	chJobs := make(chan map[int]*int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int64, chJobs chan map[int]*int32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int64, len(list))
	newList := make([]*int64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt32Int64PtrNoOrder(f func(*int32) *int64, list []*int32, worker int) []*int64 {
	chJobs := make(chan *int32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int64, chJobs chan *int32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt32Int16Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Int16Ptr(f func(*int32) *int16, list []*int32, optional ...Optional) []*int16 {
	if f == nil {
		return []*int16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32Int16PtrNoOrder(f, list, worker)
		}
	}

	return pMapInt32Int16PtrPreserveOrder(f, list, worker)
}

func pMapInt32Int16PtrPreserveOrder(f func(*int32) *int16, list []*int32, worker int) []*int16 {
	chJobs := make(chan map[int]*int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int16, chJobs chan map[int]*int32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int16{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int16, len(list))
	newList := make([]*int16, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt32Int16PtrNoOrder(f func(*int32) *int16, list []*int32, worker int) []*int16 {
	chJobs := make(chan *int32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int16, chJobs chan *int32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int16, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt32Int8Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Int8Ptr(f func(*int32) *int8, list []*int32, optional ...Optional) []*int8 {
	if f == nil {
		return []*int8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32Int8PtrNoOrder(f, list, worker)
		}
	}

	return pMapInt32Int8PtrPreserveOrder(f, list, worker)
}

func pMapInt32Int8PtrPreserveOrder(f func(*int32) *int8, list []*int32, worker int) []*int8 {
	chJobs := make(chan map[int]*int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int8, chJobs chan map[int]*int32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int8{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int8, len(list))
	newList := make([]*int8, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt32Int8PtrNoOrder(f func(*int32) *int8, list []*int32, worker int) []*int8 {
	chJobs := make(chan *int32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int8, chJobs chan *int32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int8, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt32UintPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32UintPtr(f func(*int32) *uint, list []*int32, optional ...Optional) []*uint {
	if f == nil {
		return []*uint{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32UintPtrNoOrder(f, list, worker)
		}
	}

	return pMapInt32UintPtrPreserveOrder(f, list, worker)
}

func pMapInt32UintPtrPreserveOrder(f func(*int32) *uint, list []*int32, worker int) []*uint {
	chJobs := make(chan map[int]*int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint, chJobs chan map[int]*int32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint, len(list))
	newList := make([]*uint, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt32UintPtrNoOrder(f func(*int32) *uint, list []*int32, worker int) []*uint {
	chJobs := make(chan *int32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint, chJobs chan *int32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt32Uint64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Uint64Ptr(f func(*int32) *uint64, list []*int32, optional ...Optional) []*uint64 {
	if f == nil {
		return []*uint64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32Uint64PtrNoOrder(f, list, worker)
		}
	}

	return pMapInt32Uint64PtrPreserveOrder(f, list, worker)
}

func pMapInt32Uint64PtrPreserveOrder(f func(*int32) *uint64, list []*int32, worker int) []*uint64 {
	chJobs := make(chan map[int]*int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint64, chJobs chan map[int]*int32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint64, len(list))
	newList := make([]*uint64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt32Uint64PtrNoOrder(f func(*int32) *uint64, list []*int32, worker int) []*uint64 {
	chJobs := make(chan *int32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint64, chJobs chan *int32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt32Uint32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Uint32Ptr(f func(*int32) *uint32, list []*int32, optional ...Optional) []*uint32 {
	if f == nil {
		return []*uint32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32Uint32PtrNoOrder(f, list, worker)
		}
	}

	return pMapInt32Uint32PtrPreserveOrder(f, list, worker)
}

func pMapInt32Uint32PtrPreserveOrder(f func(*int32) *uint32, list []*int32, worker int) []*uint32 {
	chJobs := make(chan map[int]*int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint32, chJobs chan map[int]*int32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint32, len(list))
	newList := make([]*uint32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt32Uint32PtrNoOrder(f func(*int32) *uint32, list []*int32, worker int) []*uint32 {
	chJobs := make(chan *int32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint32, chJobs chan *int32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt32Uint16Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Uint16Ptr(f func(*int32) *uint16, list []*int32, optional ...Optional) []*uint16 {
	if f == nil {
		return []*uint16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32Uint16PtrNoOrder(f, list, worker)
		}
	}

	return pMapInt32Uint16PtrPreserveOrder(f, list, worker)
}

func pMapInt32Uint16PtrPreserveOrder(f func(*int32) *uint16, list []*int32, worker int) []*uint16 {
	chJobs := make(chan map[int]*int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint16, chJobs chan map[int]*int32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint16{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint16, len(list))
	newList := make([]*uint16, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt32Uint16PtrNoOrder(f func(*int32) *uint16, list []*int32, worker int) []*uint16 {
	chJobs := make(chan *int32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint16, chJobs chan *int32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint16, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt32Uint8Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Uint8Ptr(f func(*int32) *uint8, list []*int32, optional ...Optional) []*uint8 {
	if f == nil {
		return []*uint8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32Uint8PtrNoOrder(f, list, worker)
		}
	}

	return pMapInt32Uint8PtrPreserveOrder(f, list, worker)
}

func pMapInt32Uint8PtrPreserveOrder(f func(*int32) *uint8, list []*int32, worker int) []*uint8 {
	chJobs := make(chan map[int]*int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint8, chJobs chan map[int]*int32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint8{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint8, len(list))
	newList := make([]*uint8, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt32Uint8PtrNoOrder(f func(*int32) *uint8, list []*int32, worker int) []*uint8 {
	chJobs := make(chan *int32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint8, chJobs chan *int32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint8, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt32StrPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32StrPtr(f func(*int32) *string, list []*int32, optional ...Optional) []*string {
	if f == nil {
		return []*string{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32StrPtrNoOrder(f, list, worker)
		}
	}

	return pMapInt32StrPtrPreserveOrder(f, list, worker)
}

func pMapInt32StrPtrPreserveOrder(f func(*int32) *string, list []*int32, worker int) []*string {
	chJobs := make(chan map[int]*int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*string, chJobs chan map[int]*int32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*string{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*string, len(list))
	newList := make([]*string, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt32StrPtrNoOrder(f func(*int32) *string, list []*int32, worker int) []*string {
	chJobs := make(chan *int32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *string, chJobs chan *int32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*string, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt32BoolPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32BoolPtr(f func(*int32) *bool, list []*int32, optional ...Optional) []*bool {
	if f == nil {
		return []*bool{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32BoolPtrNoOrder(f, list, worker)
		}
	}

	return pMapInt32BoolPtrPreserveOrder(f, list, worker)
}

func pMapInt32BoolPtrPreserveOrder(f func(*int32) *bool, list []*int32, worker int) []*bool {
	chJobs := make(chan map[int]*int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*bool, chJobs chan map[int]*int32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*bool{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*bool, len(list))
	newList := make([]*bool, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt32BoolPtrNoOrder(f func(*int32) *bool, list []*int32, worker int) []*bool {
	chJobs := make(chan *int32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *bool, chJobs chan *int32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*bool, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt32Float32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Float32Ptr(f func(*int32) *float32, list []*int32, optional ...Optional) []*float32 {
	if f == nil {
		return []*float32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32Float32PtrNoOrder(f, list, worker)
		}
	}

	return pMapInt32Float32PtrPreserveOrder(f, list, worker)
}

func pMapInt32Float32PtrPreserveOrder(f func(*int32) *float32, list []*int32, worker int) []*float32 {
	chJobs := make(chan map[int]*int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float32, chJobs chan map[int]*int32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*float32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float32, len(list))
	newList := make([]*float32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt32Float32PtrNoOrder(f func(*int32) *float32, list []*int32, worker int) []*float32 {
	chJobs := make(chan *int32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float32, chJobs chan *int32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*float32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt32Float64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Float64Ptr(f func(*int32) *float64, list []*int32, optional ...Optional) []*float64 {
	if f == nil {
		return []*float64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32Float64PtrNoOrder(f, list, worker)
		}
	}

	return pMapInt32Float64PtrPreserveOrder(f, list, worker)
}

func pMapInt32Float64PtrPreserveOrder(f func(*int32) *float64, list []*int32, worker int) []*float64 {
	chJobs := make(chan map[int]*int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float64, chJobs chan map[int]*int32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*float64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float64, len(list))
	newList := make([]*float64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt32Float64PtrNoOrder(f func(*int32) *float64, list []*int32, worker int) []*float64 {
	chJobs := make(chan *int32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float64, chJobs chan *int32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*float64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt16IntPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16IntPtr(f func(*int16) *int, list []*int16, optional ...Optional) []*int {
	if f == nil {
		return []*int{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16IntPtrNoOrder(f, list, worker)
		}
	}

	return pMapInt16IntPtrPreserveOrder(f, list, worker)
}

func pMapInt16IntPtrPreserveOrder(f func(*int16) *int, list []*int16, worker int) []*int {
	chJobs := make(chan map[int]*int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int, chJobs chan map[int]*int16) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int, len(list))
	newList := make([]*int, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt16IntPtrNoOrder(f func(*int16) *int, list []*int16, worker int) []*int {
	chJobs := make(chan *int16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int, chJobs chan *int16) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt16Int64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Int64Ptr(f func(*int16) *int64, list []*int16, optional ...Optional) []*int64 {
	if f == nil {
		return []*int64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16Int64PtrNoOrder(f, list, worker)
		}
	}

	return pMapInt16Int64PtrPreserveOrder(f, list, worker)
}

func pMapInt16Int64PtrPreserveOrder(f func(*int16) *int64, list []*int16, worker int) []*int64 {
	chJobs := make(chan map[int]*int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int64, chJobs chan map[int]*int16) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int64, len(list))
	newList := make([]*int64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt16Int64PtrNoOrder(f func(*int16) *int64, list []*int16, worker int) []*int64 {
	chJobs := make(chan *int16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int64, chJobs chan *int16) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt16Int32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Int32Ptr(f func(*int16) *int32, list []*int16, optional ...Optional) []*int32 {
	if f == nil {
		return []*int32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16Int32PtrNoOrder(f, list, worker)
		}
	}

	return pMapInt16Int32PtrPreserveOrder(f, list, worker)
}

func pMapInt16Int32PtrPreserveOrder(f func(*int16) *int32, list []*int16, worker int) []*int32 {
	chJobs := make(chan map[int]*int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int32, chJobs chan map[int]*int16) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int32, len(list))
	newList := make([]*int32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt16Int32PtrNoOrder(f func(*int16) *int32, list []*int16, worker int) []*int32 {
	chJobs := make(chan *int16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int32, chJobs chan *int16) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt16Int8Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Int8Ptr(f func(*int16) *int8, list []*int16, optional ...Optional) []*int8 {
	if f == nil {
		return []*int8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16Int8PtrNoOrder(f, list, worker)
		}
	}

	return pMapInt16Int8PtrPreserveOrder(f, list, worker)
}

func pMapInt16Int8PtrPreserveOrder(f func(*int16) *int8, list []*int16, worker int) []*int8 {
	chJobs := make(chan map[int]*int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int8, chJobs chan map[int]*int16) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int8{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int8, len(list))
	newList := make([]*int8, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt16Int8PtrNoOrder(f func(*int16) *int8, list []*int16, worker int) []*int8 {
	chJobs := make(chan *int16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int8, chJobs chan *int16) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int8, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt16UintPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16UintPtr(f func(*int16) *uint, list []*int16, optional ...Optional) []*uint {
	if f == nil {
		return []*uint{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16UintPtrNoOrder(f, list, worker)
		}
	}

	return pMapInt16UintPtrPreserveOrder(f, list, worker)
}

func pMapInt16UintPtrPreserveOrder(f func(*int16) *uint, list []*int16, worker int) []*uint {
	chJobs := make(chan map[int]*int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint, chJobs chan map[int]*int16) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint, len(list))
	newList := make([]*uint, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt16UintPtrNoOrder(f func(*int16) *uint, list []*int16, worker int) []*uint {
	chJobs := make(chan *int16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint, chJobs chan *int16) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt16Uint64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Uint64Ptr(f func(*int16) *uint64, list []*int16, optional ...Optional) []*uint64 {
	if f == nil {
		return []*uint64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16Uint64PtrNoOrder(f, list, worker)
		}
	}

	return pMapInt16Uint64PtrPreserveOrder(f, list, worker)
}

func pMapInt16Uint64PtrPreserveOrder(f func(*int16) *uint64, list []*int16, worker int) []*uint64 {
	chJobs := make(chan map[int]*int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint64, chJobs chan map[int]*int16) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint64, len(list))
	newList := make([]*uint64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt16Uint64PtrNoOrder(f func(*int16) *uint64, list []*int16, worker int) []*uint64 {
	chJobs := make(chan *int16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint64, chJobs chan *int16) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt16Uint32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Uint32Ptr(f func(*int16) *uint32, list []*int16, optional ...Optional) []*uint32 {
	if f == nil {
		return []*uint32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16Uint32PtrNoOrder(f, list, worker)
		}
	}

	return pMapInt16Uint32PtrPreserveOrder(f, list, worker)
}

func pMapInt16Uint32PtrPreserveOrder(f func(*int16) *uint32, list []*int16, worker int) []*uint32 {
	chJobs := make(chan map[int]*int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint32, chJobs chan map[int]*int16) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint32, len(list))
	newList := make([]*uint32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt16Uint32PtrNoOrder(f func(*int16) *uint32, list []*int16, worker int) []*uint32 {
	chJobs := make(chan *int16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint32, chJobs chan *int16) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt16Uint16Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Uint16Ptr(f func(*int16) *uint16, list []*int16, optional ...Optional) []*uint16 {
	if f == nil {
		return []*uint16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16Uint16PtrNoOrder(f, list, worker)
		}
	}

	return pMapInt16Uint16PtrPreserveOrder(f, list, worker)
}

func pMapInt16Uint16PtrPreserveOrder(f func(*int16) *uint16, list []*int16, worker int) []*uint16 {
	chJobs := make(chan map[int]*int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint16, chJobs chan map[int]*int16) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint16{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint16, len(list))
	newList := make([]*uint16, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt16Uint16PtrNoOrder(f func(*int16) *uint16, list []*int16, worker int) []*uint16 {
	chJobs := make(chan *int16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint16, chJobs chan *int16) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint16, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt16Uint8Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Uint8Ptr(f func(*int16) *uint8, list []*int16, optional ...Optional) []*uint8 {
	if f == nil {
		return []*uint8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16Uint8PtrNoOrder(f, list, worker)
		}
	}

	return pMapInt16Uint8PtrPreserveOrder(f, list, worker)
}

func pMapInt16Uint8PtrPreserveOrder(f func(*int16) *uint8, list []*int16, worker int) []*uint8 {
	chJobs := make(chan map[int]*int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint8, chJobs chan map[int]*int16) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint8{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint8, len(list))
	newList := make([]*uint8, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt16Uint8PtrNoOrder(f func(*int16) *uint8, list []*int16, worker int) []*uint8 {
	chJobs := make(chan *int16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint8, chJobs chan *int16) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint8, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt16StrPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16StrPtr(f func(*int16) *string, list []*int16, optional ...Optional) []*string {
	if f == nil {
		return []*string{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16StrPtrNoOrder(f, list, worker)
		}
	}

	return pMapInt16StrPtrPreserveOrder(f, list, worker)
}

func pMapInt16StrPtrPreserveOrder(f func(*int16) *string, list []*int16, worker int) []*string {
	chJobs := make(chan map[int]*int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*string, chJobs chan map[int]*int16) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*string{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*string, len(list))
	newList := make([]*string, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt16StrPtrNoOrder(f func(*int16) *string, list []*int16, worker int) []*string {
	chJobs := make(chan *int16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *string, chJobs chan *int16) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*string, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt16BoolPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16BoolPtr(f func(*int16) *bool, list []*int16, optional ...Optional) []*bool {
	if f == nil {
		return []*bool{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16BoolPtrNoOrder(f, list, worker)
		}
	}

	return pMapInt16BoolPtrPreserveOrder(f, list, worker)
}

func pMapInt16BoolPtrPreserveOrder(f func(*int16) *bool, list []*int16, worker int) []*bool {
	chJobs := make(chan map[int]*int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*bool, chJobs chan map[int]*int16) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*bool{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*bool, len(list))
	newList := make([]*bool, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt16BoolPtrNoOrder(f func(*int16) *bool, list []*int16, worker int) []*bool {
	chJobs := make(chan *int16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *bool, chJobs chan *int16) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*bool, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt16Float32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Float32Ptr(f func(*int16) *float32, list []*int16, optional ...Optional) []*float32 {
	if f == nil {
		return []*float32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16Float32PtrNoOrder(f, list, worker)
		}
	}

	return pMapInt16Float32PtrPreserveOrder(f, list, worker)
}

func pMapInt16Float32PtrPreserveOrder(f func(*int16) *float32, list []*int16, worker int) []*float32 {
	chJobs := make(chan map[int]*int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float32, chJobs chan map[int]*int16) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*float32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float32, len(list))
	newList := make([]*float32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt16Float32PtrNoOrder(f func(*int16) *float32, list []*int16, worker int) []*float32 {
	chJobs := make(chan *int16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float32, chJobs chan *int16) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*float32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt16Float64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Float64Ptr(f func(*int16) *float64, list []*int16, optional ...Optional) []*float64 {
	if f == nil {
		return []*float64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16Float64PtrNoOrder(f, list, worker)
		}
	}

	return pMapInt16Float64PtrPreserveOrder(f, list, worker)
}

func pMapInt16Float64PtrPreserveOrder(f func(*int16) *float64, list []*int16, worker int) []*float64 {
	chJobs := make(chan map[int]*int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float64, chJobs chan map[int]*int16) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*float64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float64, len(list))
	newList := make([]*float64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt16Float64PtrNoOrder(f func(*int16) *float64, list []*int16, worker int) []*float64 {
	chJobs := make(chan *int16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float64, chJobs chan *int16) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*float64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt8IntPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8IntPtr(f func(*int8) *int, list []*int8, optional ...Optional) []*int {
	if f == nil {
		return []*int{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8IntPtrNoOrder(f, list, worker)
		}
	}

	return pMapInt8IntPtrPreserveOrder(f, list, worker)
}

func pMapInt8IntPtrPreserveOrder(f func(*int8) *int, list []*int8, worker int) []*int {
	chJobs := make(chan map[int]*int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int, chJobs chan map[int]*int8) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int, len(list))
	newList := make([]*int, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt8IntPtrNoOrder(f func(*int8) *int, list []*int8, worker int) []*int {
	chJobs := make(chan *int8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int, chJobs chan *int8) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt8Int64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Int64Ptr(f func(*int8) *int64, list []*int8, optional ...Optional) []*int64 {
	if f == nil {
		return []*int64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8Int64PtrNoOrder(f, list, worker)
		}
	}

	return pMapInt8Int64PtrPreserveOrder(f, list, worker)
}

func pMapInt8Int64PtrPreserveOrder(f func(*int8) *int64, list []*int8, worker int) []*int64 {
	chJobs := make(chan map[int]*int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int64, chJobs chan map[int]*int8) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int64, len(list))
	newList := make([]*int64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt8Int64PtrNoOrder(f func(*int8) *int64, list []*int8, worker int) []*int64 {
	chJobs := make(chan *int8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int64, chJobs chan *int8) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt8Int32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Int32Ptr(f func(*int8) *int32, list []*int8, optional ...Optional) []*int32 {
	if f == nil {
		return []*int32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8Int32PtrNoOrder(f, list, worker)
		}
	}

	return pMapInt8Int32PtrPreserveOrder(f, list, worker)
}

func pMapInt8Int32PtrPreserveOrder(f func(*int8) *int32, list []*int8, worker int) []*int32 {
	chJobs := make(chan map[int]*int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int32, chJobs chan map[int]*int8) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int32, len(list))
	newList := make([]*int32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt8Int32PtrNoOrder(f func(*int8) *int32, list []*int8, worker int) []*int32 {
	chJobs := make(chan *int8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int32, chJobs chan *int8) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt8Int16Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Int16Ptr(f func(*int8) *int16, list []*int8, optional ...Optional) []*int16 {
	if f == nil {
		return []*int16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8Int16PtrNoOrder(f, list, worker)
		}
	}

	return pMapInt8Int16PtrPreserveOrder(f, list, worker)
}

func pMapInt8Int16PtrPreserveOrder(f func(*int8) *int16, list []*int8, worker int) []*int16 {
	chJobs := make(chan map[int]*int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int16, chJobs chan map[int]*int8) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int16{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int16, len(list))
	newList := make([]*int16, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt8Int16PtrNoOrder(f func(*int8) *int16, list []*int8, worker int) []*int16 {
	chJobs := make(chan *int8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int16, chJobs chan *int8) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int16, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt8UintPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8UintPtr(f func(*int8) *uint, list []*int8, optional ...Optional) []*uint {
	if f == nil {
		return []*uint{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8UintPtrNoOrder(f, list, worker)
		}
	}

	return pMapInt8UintPtrPreserveOrder(f, list, worker)
}

func pMapInt8UintPtrPreserveOrder(f func(*int8) *uint, list []*int8, worker int) []*uint {
	chJobs := make(chan map[int]*int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint, chJobs chan map[int]*int8) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint, len(list))
	newList := make([]*uint, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt8UintPtrNoOrder(f func(*int8) *uint, list []*int8, worker int) []*uint {
	chJobs := make(chan *int8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint, chJobs chan *int8) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt8Uint64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Uint64Ptr(f func(*int8) *uint64, list []*int8, optional ...Optional) []*uint64 {
	if f == nil {
		return []*uint64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8Uint64PtrNoOrder(f, list, worker)
		}
	}

	return pMapInt8Uint64PtrPreserveOrder(f, list, worker)
}

func pMapInt8Uint64PtrPreserveOrder(f func(*int8) *uint64, list []*int8, worker int) []*uint64 {
	chJobs := make(chan map[int]*int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint64, chJobs chan map[int]*int8) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint64, len(list))
	newList := make([]*uint64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt8Uint64PtrNoOrder(f func(*int8) *uint64, list []*int8, worker int) []*uint64 {
	chJobs := make(chan *int8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint64, chJobs chan *int8) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt8Uint32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Uint32Ptr(f func(*int8) *uint32, list []*int8, optional ...Optional) []*uint32 {
	if f == nil {
		return []*uint32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8Uint32PtrNoOrder(f, list, worker)
		}
	}

	return pMapInt8Uint32PtrPreserveOrder(f, list, worker)
}

func pMapInt8Uint32PtrPreserveOrder(f func(*int8) *uint32, list []*int8, worker int) []*uint32 {
	chJobs := make(chan map[int]*int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint32, chJobs chan map[int]*int8) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint32, len(list))
	newList := make([]*uint32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt8Uint32PtrNoOrder(f func(*int8) *uint32, list []*int8, worker int) []*uint32 {
	chJobs := make(chan *int8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint32, chJobs chan *int8) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt8Uint16Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Uint16Ptr(f func(*int8) *uint16, list []*int8, optional ...Optional) []*uint16 {
	if f == nil {
		return []*uint16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8Uint16PtrNoOrder(f, list, worker)
		}
	}

	return pMapInt8Uint16PtrPreserveOrder(f, list, worker)
}

func pMapInt8Uint16PtrPreserveOrder(f func(*int8) *uint16, list []*int8, worker int) []*uint16 {
	chJobs := make(chan map[int]*int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint16, chJobs chan map[int]*int8) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint16{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint16, len(list))
	newList := make([]*uint16, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt8Uint16PtrNoOrder(f func(*int8) *uint16, list []*int8, worker int) []*uint16 {
	chJobs := make(chan *int8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint16, chJobs chan *int8) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint16, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt8Uint8Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Uint8Ptr(f func(*int8) *uint8, list []*int8, optional ...Optional) []*uint8 {
	if f == nil {
		return []*uint8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8Uint8PtrNoOrder(f, list, worker)
		}
	}

	return pMapInt8Uint8PtrPreserveOrder(f, list, worker)
}

func pMapInt8Uint8PtrPreserveOrder(f func(*int8) *uint8, list []*int8, worker int) []*uint8 {
	chJobs := make(chan map[int]*int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint8, chJobs chan map[int]*int8) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint8{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint8, len(list))
	newList := make([]*uint8, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt8Uint8PtrNoOrder(f func(*int8) *uint8, list []*int8, worker int) []*uint8 {
	chJobs := make(chan *int8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint8, chJobs chan *int8) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint8, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt8StrPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8StrPtr(f func(*int8) *string, list []*int8, optional ...Optional) []*string {
	if f == nil {
		return []*string{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8StrPtrNoOrder(f, list, worker)
		}
	}

	return pMapInt8StrPtrPreserveOrder(f, list, worker)
}

func pMapInt8StrPtrPreserveOrder(f func(*int8) *string, list []*int8, worker int) []*string {
	chJobs := make(chan map[int]*int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*string, chJobs chan map[int]*int8) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*string{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*string, len(list))
	newList := make([]*string, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt8StrPtrNoOrder(f func(*int8) *string, list []*int8, worker int) []*string {
	chJobs := make(chan *int8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *string, chJobs chan *int8) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*string, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt8BoolPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8BoolPtr(f func(*int8) *bool, list []*int8, optional ...Optional) []*bool {
	if f == nil {
		return []*bool{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8BoolPtrNoOrder(f, list, worker)
		}
	}

	return pMapInt8BoolPtrPreserveOrder(f, list, worker)
}

func pMapInt8BoolPtrPreserveOrder(f func(*int8) *bool, list []*int8, worker int) []*bool {
	chJobs := make(chan map[int]*int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*bool, chJobs chan map[int]*int8) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*bool{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*bool, len(list))
	newList := make([]*bool, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt8BoolPtrNoOrder(f func(*int8) *bool, list []*int8, worker int) []*bool {
	chJobs := make(chan *int8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *bool, chJobs chan *int8) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*bool, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt8Float32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Float32Ptr(f func(*int8) *float32, list []*int8, optional ...Optional) []*float32 {
	if f == nil {
		return []*float32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8Float32PtrNoOrder(f, list, worker)
		}
	}

	return pMapInt8Float32PtrPreserveOrder(f, list, worker)
}

func pMapInt8Float32PtrPreserveOrder(f func(*int8) *float32, list []*int8, worker int) []*float32 {
	chJobs := make(chan map[int]*int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float32, chJobs chan map[int]*int8) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*float32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float32, len(list))
	newList := make([]*float32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt8Float32PtrNoOrder(f func(*int8) *float32, list []*int8, worker int) []*float32 {
	chJobs := make(chan *int8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float32, chJobs chan *int8) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*float32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt8Float64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Float64Ptr(f func(*int8) *float64, list []*int8, optional ...Optional) []*float64 {
	if f == nil {
		return []*float64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8Float64PtrNoOrder(f, list, worker)
		}
	}

	return pMapInt8Float64PtrPreserveOrder(f, list, worker)
}

func pMapInt8Float64PtrPreserveOrder(f func(*int8) *float64, list []*int8, worker int) []*float64 {
	chJobs := make(chan map[int]*int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float64, chJobs chan map[int]*int8) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*float64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float64, len(list))
	newList := make([]*float64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt8Float64PtrNoOrder(f func(*int8) *float64, list []*int8, worker int) []*float64 {
	chJobs := make(chan *int8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float64, chJobs chan *int8) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*float64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUintIntPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintIntPtr(f func(*uint) *int, list []*uint, optional ...Optional) []*int {
	if f == nil {
		return []*int{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintIntPtrNoOrder(f, list, worker)
		}
	}

	return pMapUintIntPtrPreserveOrder(f, list, worker)
}

func pMapUintIntPtrPreserveOrder(f func(*uint) *int, list []*uint, worker int) []*int {
	chJobs := make(chan map[int]*uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int, chJobs chan map[int]*uint) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int, len(list))
	newList := make([]*int, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUintIntPtrNoOrder(f func(*uint) *int, list []*uint, worker int) []*int {
	chJobs := make(chan *uint, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int, chJobs chan *uint) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUintInt64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintInt64Ptr(f func(*uint) *int64, list []*uint, optional ...Optional) []*int64 {
	if f == nil {
		return []*int64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintInt64PtrNoOrder(f, list, worker)
		}
	}

	return pMapUintInt64PtrPreserveOrder(f, list, worker)
}

func pMapUintInt64PtrPreserveOrder(f func(*uint) *int64, list []*uint, worker int) []*int64 {
	chJobs := make(chan map[int]*uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int64, chJobs chan map[int]*uint) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int64, len(list))
	newList := make([]*int64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUintInt64PtrNoOrder(f func(*uint) *int64, list []*uint, worker int) []*int64 {
	chJobs := make(chan *uint, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int64, chJobs chan *uint) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUintInt32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintInt32Ptr(f func(*uint) *int32, list []*uint, optional ...Optional) []*int32 {
	if f == nil {
		return []*int32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintInt32PtrNoOrder(f, list, worker)
		}
	}

	return pMapUintInt32PtrPreserveOrder(f, list, worker)
}

func pMapUintInt32PtrPreserveOrder(f func(*uint) *int32, list []*uint, worker int) []*int32 {
	chJobs := make(chan map[int]*uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int32, chJobs chan map[int]*uint) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int32, len(list))
	newList := make([]*int32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUintInt32PtrNoOrder(f func(*uint) *int32, list []*uint, worker int) []*int32 {
	chJobs := make(chan *uint, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int32, chJobs chan *uint) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUintInt16Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintInt16Ptr(f func(*uint) *int16, list []*uint, optional ...Optional) []*int16 {
	if f == nil {
		return []*int16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintInt16PtrNoOrder(f, list, worker)
		}
	}

	return pMapUintInt16PtrPreserveOrder(f, list, worker)
}

func pMapUintInt16PtrPreserveOrder(f func(*uint) *int16, list []*uint, worker int) []*int16 {
	chJobs := make(chan map[int]*uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int16, chJobs chan map[int]*uint) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int16{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int16, len(list))
	newList := make([]*int16, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUintInt16PtrNoOrder(f func(*uint) *int16, list []*uint, worker int) []*int16 {
	chJobs := make(chan *uint, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int16, chJobs chan *uint) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int16, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUintInt8Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintInt8Ptr(f func(*uint) *int8, list []*uint, optional ...Optional) []*int8 {
	if f == nil {
		return []*int8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintInt8PtrNoOrder(f, list, worker)
		}
	}

	return pMapUintInt8PtrPreserveOrder(f, list, worker)
}

func pMapUintInt8PtrPreserveOrder(f func(*uint) *int8, list []*uint, worker int) []*int8 {
	chJobs := make(chan map[int]*uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int8, chJobs chan map[int]*uint) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int8{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int8, len(list))
	newList := make([]*int8, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUintInt8PtrNoOrder(f func(*uint) *int8, list []*uint, worker int) []*int8 {
	chJobs := make(chan *uint, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int8, chJobs chan *uint) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int8, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUintUint64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintUint64Ptr(f func(*uint) *uint64, list []*uint, optional ...Optional) []*uint64 {
	if f == nil {
		return []*uint64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintUint64PtrNoOrder(f, list, worker)
		}
	}

	return pMapUintUint64PtrPreserveOrder(f, list, worker)
}

func pMapUintUint64PtrPreserveOrder(f func(*uint) *uint64, list []*uint, worker int) []*uint64 {
	chJobs := make(chan map[int]*uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint64, chJobs chan map[int]*uint) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint64, len(list))
	newList := make([]*uint64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUintUint64PtrNoOrder(f func(*uint) *uint64, list []*uint, worker int) []*uint64 {
	chJobs := make(chan *uint, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint64, chJobs chan *uint) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUintUint32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintUint32Ptr(f func(*uint) *uint32, list []*uint, optional ...Optional) []*uint32 {
	if f == nil {
		return []*uint32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintUint32PtrNoOrder(f, list, worker)
		}
	}

	return pMapUintUint32PtrPreserveOrder(f, list, worker)
}

func pMapUintUint32PtrPreserveOrder(f func(*uint) *uint32, list []*uint, worker int) []*uint32 {
	chJobs := make(chan map[int]*uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint32, chJobs chan map[int]*uint) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint32, len(list))
	newList := make([]*uint32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUintUint32PtrNoOrder(f func(*uint) *uint32, list []*uint, worker int) []*uint32 {
	chJobs := make(chan *uint, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint32, chJobs chan *uint) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUintUint16Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintUint16Ptr(f func(*uint) *uint16, list []*uint, optional ...Optional) []*uint16 {
	if f == nil {
		return []*uint16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintUint16PtrNoOrder(f, list, worker)
		}
	}

	return pMapUintUint16PtrPreserveOrder(f, list, worker)
}

func pMapUintUint16PtrPreserveOrder(f func(*uint) *uint16, list []*uint, worker int) []*uint16 {
	chJobs := make(chan map[int]*uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint16, chJobs chan map[int]*uint) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint16{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint16, len(list))
	newList := make([]*uint16, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUintUint16PtrNoOrder(f func(*uint) *uint16, list []*uint, worker int) []*uint16 {
	chJobs := make(chan *uint, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint16, chJobs chan *uint) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint16, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUintUint8Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintUint8Ptr(f func(*uint) *uint8, list []*uint, optional ...Optional) []*uint8 {
	if f == nil {
		return []*uint8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintUint8PtrNoOrder(f, list, worker)
		}
	}

	return pMapUintUint8PtrPreserveOrder(f, list, worker)
}

func pMapUintUint8PtrPreserveOrder(f func(*uint) *uint8, list []*uint, worker int) []*uint8 {
	chJobs := make(chan map[int]*uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint8, chJobs chan map[int]*uint) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint8{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint8, len(list))
	newList := make([]*uint8, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUintUint8PtrNoOrder(f func(*uint) *uint8, list []*uint, worker int) []*uint8 {
	chJobs := make(chan *uint, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint8, chJobs chan *uint) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint8, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUintStrPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintStrPtr(f func(*uint) *string, list []*uint, optional ...Optional) []*string {
	if f == nil {
		return []*string{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintStrPtrNoOrder(f, list, worker)
		}
	}

	return pMapUintStrPtrPreserveOrder(f, list, worker)
}

func pMapUintStrPtrPreserveOrder(f func(*uint) *string, list []*uint, worker int) []*string {
	chJobs := make(chan map[int]*uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*string, chJobs chan map[int]*uint) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*string{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*string, len(list))
	newList := make([]*string, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUintStrPtrNoOrder(f func(*uint) *string, list []*uint, worker int) []*string {
	chJobs := make(chan *uint, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *string, chJobs chan *uint) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*string, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUintBoolPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintBoolPtr(f func(*uint) *bool, list []*uint, optional ...Optional) []*bool {
	if f == nil {
		return []*bool{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintBoolPtrNoOrder(f, list, worker)
		}
	}

	return pMapUintBoolPtrPreserveOrder(f, list, worker)
}

func pMapUintBoolPtrPreserveOrder(f func(*uint) *bool, list []*uint, worker int) []*bool {
	chJobs := make(chan map[int]*uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*bool, chJobs chan map[int]*uint) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*bool{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*bool, len(list))
	newList := make([]*bool, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUintBoolPtrNoOrder(f func(*uint) *bool, list []*uint, worker int) []*bool {
	chJobs := make(chan *uint, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *bool, chJobs chan *uint) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*bool, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUintFloat32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintFloat32Ptr(f func(*uint) *float32, list []*uint, optional ...Optional) []*float32 {
	if f == nil {
		return []*float32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintFloat32PtrNoOrder(f, list, worker)
		}
	}

	return pMapUintFloat32PtrPreserveOrder(f, list, worker)
}

func pMapUintFloat32PtrPreserveOrder(f func(*uint) *float32, list []*uint, worker int) []*float32 {
	chJobs := make(chan map[int]*uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float32, chJobs chan map[int]*uint) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*float32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float32, len(list))
	newList := make([]*float32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUintFloat32PtrNoOrder(f func(*uint) *float32, list []*uint, worker int) []*float32 {
	chJobs := make(chan *uint, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float32, chJobs chan *uint) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*float32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUintFloat64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintFloat64Ptr(f func(*uint) *float64, list []*uint, optional ...Optional) []*float64 {
	if f == nil {
		return []*float64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintFloat64PtrNoOrder(f, list, worker)
		}
	}

	return pMapUintFloat64PtrPreserveOrder(f, list, worker)
}

func pMapUintFloat64PtrPreserveOrder(f func(*uint) *float64, list []*uint, worker int) []*float64 {
	chJobs := make(chan map[int]*uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float64, chJobs chan map[int]*uint) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*float64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float64, len(list))
	newList := make([]*float64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUintFloat64PtrNoOrder(f func(*uint) *float64, list []*uint, worker int) []*float64 {
	chJobs := make(chan *uint, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float64, chJobs chan *uint) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*float64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint64IntPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64IntPtr(f func(*uint64) *int, list []*uint64, optional ...Optional) []*int {
	if f == nil {
		return []*int{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64IntPtrNoOrder(f, list, worker)
		}
	}

	return pMapUint64IntPtrPreserveOrder(f, list, worker)
}

func pMapUint64IntPtrPreserveOrder(f func(*uint64) *int, list []*uint64, worker int) []*int {
	chJobs := make(chan map[int]*uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int, chJobs chan map[int]*uint64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int, len(list))
	newList := make([]*int, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint64IntPtrNoOrder(f func(*uint64) *int, list []*uint64, worker int) []*int {
	chJobs := make(chan *uint64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int, chJobs chan *uint64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint64Int64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Int64Ptr(f func(*uint64) *int64, list []*uint64, optional ...Optional) []*int64 {
	if f == nil {
		return []*int64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64Int64PtrNoOrder(f, list, worker)
		}
	}

	return pMapUint64Int64PtrPreserveOrder(f, list, worker)
}

func pMapUint64Int64PtrPreserveOrder(f func(*uint64) *int64, list []*uint64, worker int) []*int64 {
	chJobs := make(chan map[int]*uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int64, chJobs chan map[int]*uint64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int64, len(list))
	newList := make([]*int64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint64Int64PtrNoOrder(f func(*uint64) *int64, list []*uint64, worker int) []*int64 {
	chJobs := make(chan *uint64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int64, chJobs chan *uint64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint64Int32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Int32Ptr(f func(*uint64) *int32, list []*uint64, optional ...Optional) []*int32 {
	if f == nil {
		return []*int32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64Int32PtrNoOrder(f, list, worker)
		}
	}

	return pMapUint64Int32PtrPreserveOrder(f, list, worker)
}

func pMapUint64Int32PtrPreserveOrder(f func(*uint64) *int32, list []*uint64, worker int) []*int32 {
	chJobs := make(chan map[int]*uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int32, chJobs chan map[int]*uint64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int32, len(list))
	newList := make([]*int32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint64Int32PtrNoOrder(f func(*uint64) *int32, list []*uint64, worker int) []*int32 {
	chJobs := make(chan *uint64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int32, chJobs chan *uint64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint64Int16Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Int16Ptr(f func(*uint64) *int16, list []*uint64, optional ...Optional) []*int16 {
	if f == nil {
		return []*int16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64Int16PtrNoOrder(f, list, worker)
		}
	}

	return pMapUint64Int16PtrPreserveOrder(f, list, worker)
}

func pMapUint64Int16PtrPreserveOrder(f func(*uint64) *int16, list []*uint64, worker int) []*int16 {
	chJobs := make(chan map[int]*uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int16, chJobs chan map[int]*uint64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int16{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int16, len(list))
	newList := make([]*int16, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint64Int16PtrNoOrder(f func(*uint64) *int16, list []*uint64, worker int) []*int16 {
	chJobs := make(chan *uint64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int16, chJobs chan *uint64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int16, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint64Int8Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Int8Ptr(f func(*uint64) *int8, list []*uint64, optional ...Optional) []*int8 {
	if f == nil {
		return []*int8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64Int8PtrNoOrder(f, list, worker)
		}
	}

	return pMapUint64Int8PtrPreserveOrder(f, list, worker)
}

func pMapUint64Int8PtrPreserveOrder(f func(*uint64) *int8, list []*uint64, worker int) []*int8 {
	chJobs := make(chan map[int]*uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int8, chJobs chan map[int]*uint64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int8{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int8, len(list))
	newList := make([]*int8, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint64Int8PtrNoOrder(f func(*uint64) *int8, list []*uint64, worker int) []*int8 {
	chJobs := make(chan *uint64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int8, chJobs chan *uint64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int8, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint64UintPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64UintPtr(f func(*uint64) *uint, list []*uint64, optional ...Optional) []*uint {
	if f == nil {
		return []*uint{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64UintPtrNoOrder(f, list, worker)
		}
	}

	return pMapUint64UintPtrPreserveOrder(f, list, worker)
}

func pMapUint64UintPtrPreserveOrder(f func(*uint64) *uint, list []*uint64, worker int) []*uint {
	chJobs := make(chan map[int]*uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint, chJobs chan map[int]*uint64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint, len(list))
	newList := make([]*uint, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint64UintPtrNoOrder(f func(*uint64) *uint, list []*uint64, worker int) []*uint {
	chJobs := make(chan *uint64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint, chJobs chan *uint64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint64Uint32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Uint32Ptr(f func(*uint64) *uint32, list []*uint64, optional ...Optional) []*uint32 {
	if f == nil {
		return []*uint32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64Uint32PtrNoOrder(f, list, worker)
		}
	}

	return pMapUint64Uint32PtrPreserveOrder(f, list, worker)
}

func pMapUint64Uint32PtrPreserveOrder(f func(*uint64) *uint32, list []*uint64, worker int) []*uint32 {
	chJobs := make(chan map[int]*uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint32, chJobs chan map[int]*uint64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint32, len(list))
	newList := make([]*uint32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint64Uint32PtrNoOrder(f func(*uint64) *uint32, list []*uint64, worker int) []*uint32 {
	chJobs := make(chan *uint64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint32, chJobs chan *uint64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint64Uint16Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Uint16Ptr(f func(*uint64) *uint16, list []*uint64, optional ...Optional) []*uint16 {
	if f == nil {
		return []*uint16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64Uint16PtrNoOrder(f, list, worker)
		}
	}

	return pMapUint64Uint16PtrPreserveOrder(f, list, worker)
}

func pMapUint64Uint16PtrPreserveOrder(f func(*uint64) *uint16, list []*uint64, worker int) []*uint16 {
	chJobs := make(chan map[int]*uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint16, chJobs chan map[int]*uint64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint16{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint16, len(list))
	newList := make([]*uint16, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint64Uint16PtrNoOrder(f func(*uint64) *uint16, list []*uint64, worker int) []*uint16 {
	chJobs := make(chan *uint64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint16, chJobs chan *uint64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint16, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint64Uint8Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Uint8Ptr(f func(*uint64) *uint8, list []*uint64, optional ...Optional) []*uint8 {
	if f == nil {
		return []*uint8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64Uint8PtrNoOrder(f, list, worker)
		}
	}

	return pMapUint64Uint8PtrPreserveOrder(f, list, worker)
}

func pMapUint64Uint8PtrPreserveOrder(f func(*uint64) *uint8, list []*uint64, worker int) []*uint8 {
	chJobs := make(chan map[int]*uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint8, chJobs chan map[int]*uint64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint8{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint8, len(list))
	newList := make([]*uint8, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint64Uint8PtrNoOrder(f func(*uint64) *uint8, list []*uint64, worker int) []*uint8 {
	chJobs := make(chan *uint64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint8, chJobs chan *uint64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint8, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint64StrPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64StrPtr(f func(*uint64) *string, list []*uint64, optional ...Optional) []*string {
	if f == nil {
		return []*string{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64StrPtrNoOrder(f, list, worker)
		}
	}

	return pMapUint64StrPtrPreserveOrder(f, list, worker)
}

func pMapUint64StrPtrPreserveOrder(f func(*uint64) *string, list []*uint64, worker int) []*string {
	chJobs := make(chan map[int]*uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*string, chJobs chan map[int]*uint64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*string{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*string, len(list))
	newList := make([]*string, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint64StrPtrNoOrder(f func(*uint64) *string, list []*uint64, worker int) []*string {
	chJobs := make(chan *uint64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *string, chJobs chan *uint64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*string, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint64BoolPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64BoolPtr(f func(*uint64) *bool, list []*uint64, optional ...Optional) []*bool {
	if f == nil {
		return []*bool{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64BoolPtrNoOrder(f, list, worker)
		}
	}

	return pMapUint64BoolPtrPreserveOrder(f, list, worker)
}

func pMapUint64BoolPtrPreserveOrder(f func(*uint64) *bool, list []*uint64, worker int) []*bool {
	chJobs := make(chan map[int]*uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*bool, chJobs chan map[int]*uint64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*bool{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*bool, len(list))
	newList := make([]*bool, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint64BoolPtrNoOrder(f func(*uint64) *bool, list []*uint64, worker int) []*bool {
	chJobs := make(chan *uint64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *bool, chJobs chan *uint64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*bool, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint64Float32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Float32Ptr(f func(*uint64) *float32, list []*uint64, optional ...Optional) []*float32 {
	if f == nil {
		return []*float32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64Float32PtrNoOrder(f, list, worker)
		}
	}

	return pMapUint64Float32PtrPreserveOrder(f, list, worker)
}

func pMapUint64Float32PtrPreserveOrder(f func(*uint64) *float32, list []*uint64, worker int) []*float32 {
	chJobs := make(chan map[int]*uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float32, chJobs chan map[int]*uint64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*float32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float32, len(list))
	newList := make([]*float32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint64Float32PtrNoOrder(f func(*uint64) *float32, list []*uint64, worker int) []*float32 {
	chJobs := make(chan *uint64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float32, chJobs chan *uint64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*float32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint64Float64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Float64Ptr(f func(*uint64) *float64, list []*uint64, optional ...Optional) []*float64 {
	if f == nil {
		return []*float64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64Float64PtrNoOrder(f, list, worker)
		}
	}

	return pMapUint64Float64PtrPreserveOrder(f, list, worker)
}

func pMapUint64Float64PtrPreserveOrder(f func(*uint64) *float64, list []*uint64, worker int) []*float64 {
	chJobs := make(chan map[int]*uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float64, chJobs chan map[int]*uint64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*float64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float64, len(list))
	newList := make([]*float64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint64Float64PtrNoOrder(f func(*uint64) *float64, list []*uint64, worker int) []*float64 {
	chJobs := make(chan *uint64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float64, chJobs chan *uint64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*float64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint32IntPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32IntPtr(f func(*uint32) *int, list []*uint32, optional ...Optional) []*int {
	if f == nil {
		return []*int{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32IntPtrNoOrder(f, list, worker)
		}
	}

	return pMapUint32IntPtrPreserveOrder(f, list, worker)
}

func pMapUint32IntPtrPreserveOrder(f func(*uint32) *int, list []*uint32, worker int) []*int {
	chJobs := make(chan map[int]*uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int, chJobs chan map[int]*uint32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int, len(list))
	newList := make([]*int, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint32IntPtrNoOrder(f func(*uint32) *int, list []*uint32, worker int) []*int {
	chJobs := make(chan *uint32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int, chJobs chan *uint32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint32Int64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Int64Ptr(f func(*uint32) *int64, list []*uint32, optional ...Optional) []*int64 {
	if f == nil {
		return []*int64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32Int64PtrNoOrder(f, list, worker)
		}
	}

	return pMapUint32Int64PtrPreserveOrder(f, list, worker)
}

func pMapUint32Int64PtrPreserveOrder(f func(*uint32) *int64, list []*uint32, worker int) []*int64 {
	chJobs := make(chan map[int]*uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int64, chJobs chan map[int]*uint32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int64, len(list))
	newList := make([]*int64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint32Int64PtrNoOrder(f func(*uint32) *int64, list []*uint32, worker int) []*int64 {
	chJobs := make(chan *uint32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int64, chJobs chan *uint32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint32Int32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Int32Ptr(f func(*uint32) *int32, list []*uint32, optional ...Optional) []*int32 {
	if f == nil {
		return []*int32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32Int32PtrNoOrder(f, list, worker)
		}
	}

	return pMapUint32Int32PtrPreserveOrder(f, list, worker)
}

func pMapUint32Int32PtrPreserveOrder(f func(*uint32) *int32, list []*uint32, worker int) []*int32 {
	chJobs := make(chan map[int]*uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int32, chJobs chan map[int]*uint32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int32, len(list))
	newList := make([]*int32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint32Int32PtrNoOrder(f func(*uint32) *int32, list []*uint32, worker int) []*int32 {
	chJobs := make(chan *uint32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int32, chJobs chan *uint32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint32Int16Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Int16Ptr(f func(*uint32) *int16, list []*uint32, optional ...Optional) []*int16 {
	if f == nil {
		return []*int16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32Int16PtrNoOrder(f, list, worker)
		}
	}

	return pMapUint32Int16PtrPreserveOrder(f, list, worker)
}

func pMapUint32Int16PtrPreserveOrder(f func(*uint32) *int16, list []*uint32, worker int) []*int16 {
	chJobs := make(chan map[int]*uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int16, chJobs chan map[int]*uint32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int16{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int16, len(list))
	newList := make([]*int16, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint32Int16PtrNoOrder(f func(*uint32) *int16, list []*uint32, worker int) []*int16 {
	chJobs := make(chan *uint32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int16, chJobs chan *uint32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int16, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint32Int8Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Int8Ptr(f func(*uint32) *int8, list []*uint32, optional ...Optional) []*int8 {
	if f == nil {
		return []*int8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32Int8PtrNoOrder(f, list, worker)
		}
	}

	return pMapUint32Int8PtrPreserveOrder(f, list, worker)
}

func pMapUint32Int8PtrPreserveOrder(f func(*uint32) *int8, list []*uint32, worker int) []*int8 {
	chJobs := make(chan map[int]*uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int8, chJobs chan map[int]*uint32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int8{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int8, len(list))
	newList := make([]*int8, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint32Int8PtrNoOrder(f func(*uint32) *int8, list []*uint32, worker int) []*int8 {
	chJobs := make(chan *uint32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int8, chJobs chan *uint32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int8, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint32UintPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32UintPtr(f func(*uint32) *uint, list []*uint32, optional ...Optional) []*uint {
	if f == nil {
		return []*uint{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32UintPtrNoOrder(f, list, worker)
		}
	}

	return pMapUint32UintPtrPreserveOrder(f, list, worker)
}

func pMapUint32UintPtrPreserveOrder(f func(*uint32) *uint, list []*uint32, worker int) []*uint {
	chJobs := make(chan map[int]*uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint, chJobs chan map[int]*uint32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint, len(list))
	newList := make([]*uint, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint32UintPtrNoOrder(f func(*uint32) *uint, list []*uint32, worker int) []*uint {
	chJobs := make(chan *uint32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint, chJobs chan *uint32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint32Uint64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Uint64Ptr(f func(*uint32) *uint64, list []*uint32, optional ...Optional) []*uint64 {
	if f == nil {
		return []*uint64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32Uint64PtrNoOrder(f, list, worker)
		}
	}

	return pMapUint32Uint64PtrPreserveOrder(f, list, worker)
}

func pMapUint32Uint64PtrPreserveOrder(f func(*uint32) *uint64, list []*uint32, worker int) []*uint64 {
	chJobs := make(chan map[int]*uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint64, chJobs chan map[int]*uint32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint64, len(list))
	newList := make([]*uint64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint32Uint64PtrNoOrder(f func(*uint32) *uint64, list []*uint32, worker int) []*uint64 {
	chJobs := make(chan *uint32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint64, chJobs chan *uint32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint32Uint16Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Uint16Ptr(f func(*uint32) *uint16, list []*uint32, optional ...Optional) []*uint16 {
	if f == nil {
		return []*uint16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32Uint16PtrNoOrder(f, list, worker)
		}
	}

	return pMapUint32Uint16PtrPreserveOrder(f, list, worker)
}

func pMapUint32Uint16PtrPreserveOrder(f func(*uint32) *uint16, list []*uint32, worker int) []*uint16 {
	chJobs := make(chan map[int]*uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint16, chJobs chan map[int]*uint32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint16{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint16, len(list))
	newList := make([]*uint16, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint32Uint16PtrNoOrder(f func(*uint32) *uint16, list []*uint32, worker int) []*uint16 {
	chJobs := make(chan *uint32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint16, chJobs chan *uint32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint16, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint32Uint8Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Uint8Ptr(f func(*uint32) *uint8, list []*uint32, optional ...Optional) []*uint8 {
	if f == nil {
		return []*uint8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32Uint8PtrNoOrder(f, list, worker)
		}
	}

	return pMapUint32Uint8PtrPreserveOrder(f, list, worker)
}

func pMapUint32Uint8PtrPreserveOrder(f func(*uint32) *uint8, list []*uint32, worker int) []*uint8 {
	chJobs := make(chan map[int]*uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint8, chJobs chan map[int]*uint32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint8{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint8, len(list))
	newList := make([]*uint8, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint32Uint8PtrNoOrder(f func(*uint32) *uint8, list []*uint32, worker int) []*uint8 {
	chJobs := make(chan *uint32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint8, chJobs chan *uint32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint8, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint32StrPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32StrPtr(f func(*uint32) *string, list []*uint32, optional ...Optional) []*string {
	if f == nil {
		return []*string{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32StrPtrNoOrder(f, list, worker)
		}
	}

	return pMapUint32StrPtrPreserveOrder(f, list, worker)
}

func pMapUint32StrPtrPreserveOrder(f func(*uint32) *string, list []*uint32, worker int) []*string {
	chJobs := make(chan map[int]*uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*string, chJobs chan map[int]*uint32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*string{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*string, len(list))
	newList := make([]*string, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint32StrPtrNoOrder(f func(*uint32) *string, list []*uint32, worker int) []*string {
	chJobs := make(chan *uint32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *string, chJobs chan *uint32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*string, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint32BoolPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32BoolPtr(f func(*uint32) *bool, list []*uint32, optional ...Optional) []*bool {
	if f == nil {
		return []*bool{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32BoolPtrNoOrder(f, list, worker)
		}
	}

	return pMapUint32BoolPtrPreserveOrder(f, list, worker)
}

func pMapUint32BoolPtrPreserveOrder(f func(*uint32) *bool, list []*uint32, worker int) []*bool {
	chJobs := make(chan map[int]*uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*bool, chJobs chan map[int]*uint32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*bool{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*bool, len(list))
	newList := make([]*bool, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint32BoolPtrNoOrder(f func(*uint32) *bool, list []*uint32, worker int) []*bool {
	chJobs := make(chan *uint32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *bool, chJobs chan *uint32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*bool, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint32Float32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Float32Ptr(f func(*uint32) *float32, list []*uint32, optional ...Optional) []*float32 {
	if f == nil {
		return []*float32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32Float32PtrNoOrder(f, list, worker)
		}
	}

	return pMapUint32Float32PtrPreserveOrder(f, list, worker)
}

func pMapUint32Float32PtrPreserveOrder(f func(*uint32) *float32, list []*uint32, worker int) []*float32 {
	chJobs := make(chan map[int]*uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float32, chJobs chan map[int]*uint32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*float32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float32, len(list))
	newList := make([]*float32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint32Float32PtrNoOrder(f func(*uint32) *float32, list []*uint32, worker int) []*float32 {
	chJobs := make(chan *uint32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float32, chJobs chan *uint32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*float32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint32Float64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Float64Ptr(f func(*uint32) *float64, list []*uint32, optional ...Optional) []*float64 {
	if f == nil {
		return []*float64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32Float64PtrNoOrder(f, list, worker)
		}
	}

	return pMapUint32Float64PtrPreserveOrder(f, list, worker)
}

func pMapUint32Float64PtrPreserveOrder(f func(*uint32) *float64, list []*uint32, worker int) []*float64 {
	chJobs := make(chan map[int]*uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float64, chJobs chan map[int]*uint32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*float64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float64, len(list))
	newList := make([]*float64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint32Float64PtrNoOrder(f func(*uint32) *float64, list []*uint32, worker int) []*float64 {
	chJobs := make(chan *uint32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float64, chJobs chan *uint32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*float64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint16IntPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16IntPtr(f func(*uint16) *int, list []*uint16, optional ...Optional) []*int {
	if f == nil {
		return []*int{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16IntPtrNoOrder(f, list, worker)
		}
	}

	return pMapUint16IntPtrPreserveOrder(f, list, worker)
}

func pMapUint16IntPtrPreserveOrder(f func(*uint16) *int, list []*uint16, worker int) []*int {
	chJobs := make(chan map[int]*uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int, chJobs chan map[int]*uint16) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int, len(list))
	newList := make([]*int, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint16IntPtrNoOrder(f func(*uint16) *int, list []*uint16, worker int) []*int {
	chJobs := make(chan *uint16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int, chJobs chan *uint16) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint16Int64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Int64Ptr(f func(*uint16) *int64, list []*uint16, optional ...Optional) []*int64 {
	if f == nil {
		return []*int64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16Int64PtrNoOrder(f, list, worker)
		}
	}

	return pMapUint16Int64PtrPreserveOrder(f, list, worker)
}

func pMapUint16Int64PtrPreserveOrder(f func(*uint16) *int64, list []*uint16, worker int) []*int64 {
	chJobs := make(chan map[int]*uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int64, chJobs chan map[int]*uint16) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int64, len(list))
	newList := make([]*int64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint16Int64PtrNoOrder(f func(*uint16) *int64, list []*uint16, worker int) []*int64 {
	chJobs := make(chan *uint16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int64, chJobs chan *uint16) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint16Int32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Int32Ptr(f func(*uint16) *int32, list []*uint16, optional ...Optional) []*int32 {
	if f == nil {
		return []*int32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16Int32PtrNoOrder(f, list, worker)
		}
	}

	return pMapUint16Int32PtrPreserveOrder(f, list, worker)
}

func pMapUint16Int32PtrPreserveOrder(f func(*uint16) *int32, list []*uint16, worker int) []*int32 {
	chJobs := make(chan map[int]*uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int32, chJobs chan map[int]*uint16) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int32, len(list))
	newList := make([]*int32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint16Int32PtrNoOrder(f func(*uint16) *int32, list []*uint16, worker int) []*int32 {
	chJobs := make(chan *uint16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int32, chJobs chan *uint16) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint16Int16Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Int16Ptr(f func(*uint16) *int16, list []*uint16, optional ...Optional) []*int16 {
	if f == nil {
		return []*int16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16Int16PtrNoOrder(f, list, worker)
		}
	}

	return pMapUint16Int16PtrPreserveOrder(f, list, worker)
}

func pMapUint16Int16PtrPreserveOrder(f func(*uint16) *int16, list []*uint16, worker int) []*int16 {
	chJobs := make(chan map[int]*uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int16, chJobs chan map[int]*uint16) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int16{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int16, len(list))
	newList := make([]*int16, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint16Int16PtrNoOrder(f func(*uint16) *int16, list []*uint16, worker int) []*int16 {
	chJobs := make(chan *uint16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int16, chJobs chan *uint16) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int16, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint16Int8Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Int8Ptr(f func(*uint16) *int8, list []*uint16, optional ...Optional) []*int8 {
	if f == nil {
		return []*int8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16Int8PtrNoOrder(f, list, worker)
		}
	}

	return pMapUint16Int8PtrPreserveOrder(f, list, worker)
}

func pMapUint16Int8PtrPreserveOrder(f func(*uint16) *int8, list []*uint16, worker int) []*int8 {
	chJobs := make(chan map[int]*uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int8, chJobs chan map[int]*uint16) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int8{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int8, len(list))
	newList := make([]*int8, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint16Int8PtrNoOrder(f func(*uint16) *int8, list []*uint16, worker int) []*int8 {
	chJobs := make(chan *uint16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int8, chJobs chan *uint16) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int8, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint16UintPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16UintPtr(f func(*uint16) *uint, list []*uint16, optional ...Optional) []*uint {
	if f == nil {
		return []*uint{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16UintPtrNoOrder(f, list, worker)
		}
	}

	return pMapUint16UintPtrPreserveOrder(f, list, worker)
}

func pMapUint16UintPtrPreserveOrder(f func(*uint16) *uint, list []*uint16, worker int) []*uint {
	chJobs := make(chan map[int]*uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint, chJobs chan map[int]*uint16) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint, len(list))
	newList := make([]*uint, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint16UintPtrNoOrder(f func(*uint16) *uint, list []*uint16, worker int) []*uint {
	chJobs := make(chan *uint16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint, chJobs chan *uint16) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint16Uint64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Uint64Ptr(f func(*uint16) *uint64, list []*uint16, optional ...Optional) []*uint64 {
	if f == nil {
		return []*uint64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16Uint64PtrNoOrder(f, list, worker)
		}
	}

	return pMapUint16Uint64PtrPreserveOrder(f, list, worker)
}

func pMapUint16Uint64PtrPreserveOrder(f func(*uint16) *uint64, list []*uint16, worker int) []*uint64 {
	chJobs := make(chan map[int]*uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint64, chJobs chan map[int]*uint16) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint64, len(list))
	newList := make([]*uint64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint16Uint64PtrNoOrder(f func(*uint16) *uint64, list []*uint16, worker int) []*uint64 {
	chJobs := make(chan *uint16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint64, chJobs chan *uint16) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint16Uint32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Uint32Ptr(f func(*uint16) *uint32, list []*uint16, optional ...Optional) []*uint32 {
	if f == nil {
		return []*uint32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16Uint32PtrNoOrder(f, list, worker)
		}
	}

	return pMapUint16Uint32PtrPreserveOrder(f, list, worker)
}

func pMapUint16Uint32PtrPreserveOrder(f func(*uint16) *uint32, list []*uint16, worker int) []*uint32 {
	chJobs := make(chan map[int]*uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint32, chJobs chan map[int]*uint16) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint32, len(list))
	newList := make([]*uint32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint16Uint32PtrNoOrder(f func(*uint16) *uint32, list []*uint16, worker int) []*uint32 {
	chJobs := make(chan *uint16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint32, chJobs chan *uint16) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint16Uint8Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Uint8Ptr(f func(*uint16) *uint8, list []*uint16, optional ...Optional) []*uint8 {
	if f == nil {
		return []*uint8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16Uint8PtrNoOrder(f, list, worker)
		}
	}

	return pMapUint16Uint8PtrPreserveOrder(f, list, worker)
}

func pMapUint16Uint8PtrPreserveOrder(f func(*uint16) *uint8, list []*uint16, worker int) []*uint8 {
	chJobs := make(chan map[int]*uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint8, chJobs chan map[int]*uint16) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint8{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint8, len(list))
	newList := make([]*uint8, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint16Uint8PtrNoOrder(f func(*uint16) *uint8, list []*uint16, worker int) []*uint8 {
	chJobs := make(chan *uint16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint8, chJobs chan *uint16) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint8, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint16StrPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16StrPtr(f func(*uint16) *string, list []*uint16, optional ...Optional) []*string {
	if f == nil {
		return []*string{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16StrPtrNoOrder(f, list, worker)
		}
	}

	return pMapUint16StrPtrPreserveOrder(f, list, worker)
}

func pMapUint16StrPtrPreserveOrder(f func(*uint16) *string, list []*uint16, worker int) []*string {
	chJobs := make(chan map[int]*uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*string, chJobs chan map[int]*uint16) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*string{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*string, len(list))
	newList := make([]*string, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint16StrPtrNoOrder(f func(*uint16) *string, list []*uint16, worker int) []*string {
	chJobs := make(chan *uint16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *string, chJobs chan *uint16) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*string, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint16BoolPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16BoolPtr(f func(*uint16) *bool, list []*uint16, optional ...Optional) []*bool {
	if f == nil {
		return []*bool{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16BoolPtrNoOrder(f, list, worker)
		}
	}

	return pMapUint16BoolPtrPreserveOrder(f, list, worker)
}

func pMapUint16BoolPtrPreserveOrder(f func(*uint16) *bool, list []*uint16, worker int) []*bool {
	chJobs := make(chan map[int]*uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*bool, chJobs chan map[int]*uint16) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*bool{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*bool, len(list))
	newList := make([]*bool, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint16BoolPtrNoOrder(f func(*uint16) *bool, list []*uint16, worker int) []*bool {
	chJobs := make(chan *uint16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *bool, chJobs chan *uint16) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*bool, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint16Float32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Float32Ptr(f func(*uint16) *float32, list []*uint16, optional ...Optional) []*float32 {
	if f == nil {
		return []*float32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16Float32PtrNoOrder(f, list, worker)
		}
	}

	return pMapUint16Float32PtrPreserveOrder(f, list, worker)
}

func pMapUint16Float32PtrPreserveOrder(f func(*uint16) *float32, list []*uint16, worker int) []*float32 {
	chJobs := make(chan map[int]*uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float32, chJobs chan map[int]*uint16) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*float32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float32, len(list))
	newList := make([]*float32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint16Float32PtrNoOrder(f func(*uint16) *float32, list []*uint16, worker int) []*float32 {
	chJobs := make(chan *uint16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float32, chJobs chan *uint16) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*float32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint16Float64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Float64Ptr(f func(*uint16) *float64, list []*uint16, optional ...Optional) []*float64 {
	if f == nil {
		return []*float64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16Float64PtrNoOrder(f, list, worker)
		}
	}

	return pMapUint16Float64PtrPreserveOrder(f, list, worker)
}

func pMapUint16Float64PtrPreserveOrder(f func(*uint16) *float64, list []*uint16, worker int) []*float64 {
	chJobs := make(chan map[int]*uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float64, chJobs chan map[int]*uint16) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*float64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float64, len(list))
	newList := make([]*float64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint16Float64PtrNoOrder(f func(*uint16) *float64, list []*uint16, worker int) []*float64 {
	chJobs := make(chan *uint16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float64, chJobs chan *uint16) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*float64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint8IntPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8IntPtr(f func(*uint8) *int, list []*uint8, optional ...Optional) []*int {
	if f == nil {
		return []*int{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8IntPtrNoOrder(f, list, worker)
		}
	}

	return pMapUint8IntPtrPreserveOrder(f, list, worker)
}

func pMapUint8IntPtrPreserveOrder(f func(*uint8) *int, list []*uint8, worker int) []*int {
	chJobs := make(chan map[int]*uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int, chJobs chan map[int]*uint8) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int, len(list))
	newList := make([]*int, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint8IntPtrNoOrder(f func(*uint8) *int, list []*uint8, worker int) []*int {
	chJobs := make(chan *uint8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int, chJobs chan *uint8) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint8Int64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Int64Ptr(f func(*uint8) *int64, list []*uint8, optional ...Optional) []*int64 {
	if f == nil {
		return []*int64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8Int64PtrNoOrder(f, list, worker)
		}
	}

	return pMapUint8Int64PtrPreserveOrder(f, list, worker)
}

func pMapUint8Int64PtrPreserveOrder(f func(*uint8) *int64, list []*uint8, worker int) []*int64 {
	chJobs := make(chan map[int]*uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int64, chJobs chan map[int]*uint8) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int64, len(list))
	newList := make([]*int64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint8Int64PtrNoOrder(f func(*uint8) *int64, list []*uint8, worker int) []*int64 {
	chJobs := make(chan *uint8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int64, chJobs chan *uint8) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint8Int32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Int32Ptr(f func(*uint8) *int32, list []*uint8, optional ...Optional) []*int32 {
	if f == nil {
		return []*int32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8Int32PtrNoOrder(f, list, worker)
		}
	}

	return pMapUint8Int32PtrPreserveOrder(f, list, worker)
}

func pMapUint8Int32PtrPreserveOrder(f func(*uint8) *int32, list []*uint8, worker int) []*int32 {
	chJobs := make(chan map[int]*uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int32, chJobs chan map[int]*uint8) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int32, len(list))
	newList := make([]*int32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint8Int32PtrNoOrder(f func(*uint8) *int32, list []*uint8, worker int) []*int32 {
	chJobs := make(chan *uint8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int32, chJobs chan *uint8) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint8Int16Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Int16Ptr(f func(*uint8) *int16, list []*uint8, optional ...Optional) []*int16 {
	if f == nil {
		return []*int16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8Int16PtrNoOrder(f, list, worker)
		}
	}

	return pMapUint8Int16PtrPreserveOrder(f, list, worker)
}

func pMapUint8Int16PtrPreserveOrder(f func(*uint8) *int16, list []*uint8, worker int) []*int16 {
	chJobs := make(chan map[int]*uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int16, chJobs chan map[int]*uint8) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int16{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int16, len(list))
	newList := make([]*int16, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint8Int16PtrNoOrder(f func(*uint8) *int16, list []*uint8, worker int) []*int16 {
	chJobs := make(chan *uint8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int16, chJobs chan *uint8) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int16, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint8Int8Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Int8Ptr(f func(*uint8) *int8, list []*uint8, optional ...Optional) []*int8 {
	if f == nil {
		return []*int8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8Int8PtrNoOrder(f, list, worker)
		}
	}

	return pMapUint8Int8PtrPreserveOrder(f, list, worker)
}

func pMapUint8Int8PtrPreserveOrder(f func(*uint8) *int8, list []*uint8, worker int) []*int8 {
	chJobs := make(chan map[int]*uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int8, chJobs chan map[int]*uint8) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int8{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int8, len(list))
	newList := make([]*int8, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint8Int8PtrNoOrder(f func(*uint8) *int8, list []*uint8, worker int) []*int8 {
	chJobs := make(chan *uint8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int8, chJobs chan *uint8) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int8, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint8UintPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8UintPtr(f func(*uint8) *uint, list []*uint8, optional ...Optional) []*uint {
	if f == nil {
		return []*uint{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8UintPtrNoOrder(f, list, worker)
		}
	}

	return pMapUint8UintPtrPreserveOrder(f, list, worker)
}

func pMapUint8UintPtrPreserveOrder(f func(*uint8) *uint, list []*uint8, worker int) []*uint {
	chJobs := make(chan map[int]*uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint, chJobs chan map[int]*uint8) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint, len(list))
	newList := make([]*uint, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint8UintPtrNoOrder(f func(*uint8) *uint, list []*uint8, worker int) []*uint {
	chJobs := make(chan *uint8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint, chJobs chan *uint8) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint8Uint64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Uint64Ptr(f func(*uint8) *uint64, list []*uint8, optional ...Optional) []*uint64 {
	if f == nil {
		return []*uint64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8Uint64PtrNoOrder(f, list, worker)
		}
	}

	return pMapUint8Uint64PtrPreserveOrder(f, list, worker)
}

func pMapUint8Uint64PtrPreserveOrder(f func(*uint8) *uint64, list []*uint8, worker int) []*uint64 {
	chJobs := make(chan map[int]*uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint64, chJobs chan map[int]*uint8) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint64, len(list))
	newList := make([]*uint64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint8Uint64PtrNoOrder(f func(*uint8) *uint64, list []*uint8, worker int) []*uint64 {
	chJobs := make(chan *uint8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint64, chJobs chan *uint8) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint8Uint32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Uint32Ptr(f func(*uint8) *uint32, list []*uint8, optional ...Optional) []*uint32 {
	if f == nil {
		return []*uint32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8Uint32PtrNoOrder(f, list, worker)
		}
	}

	return pMapUint8Uint32PtrPreserveOrder(f, list, worker)
}

func pMapUint8Uint32PtrPreserveOrder(f func(*uint8) *uint32, list []*uint8, worker int) []*uint32 {
	chJobs := make(chan map[int]*uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint32, chJobs chan map[int]*uint8) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint32, len(list))
	newList := make([]*uint32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint8Uint32PtrNoOrder(f func(*uint8) *uint32, list []*uint8, worker int) []*uint32 {
	chJobs := make(chan *uint8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint32, chJobs chan *uint8) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint8Uint16Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Uint16Ptr(f func(*uint8) *uint16, list []*uint8, optional ...Optional) []*uint16 {
	if f == nil {
		return []*uint16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8Uint16PtrNoOrder(f, list, worker)
		}
	}

	return pMapUint8Uint16PtrPreserveOrder(f, list, worker)
}

func pMapUint8Uint16PtrPreserveOrder(f func(*uint8) *uint16, list []*uint8, worker int) []*uint16 {
	chJobs := make(chan map[int]*uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint16, chJobs chan map[int]*uint8) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint16{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint16, len(list))
	newList := make([]*uint16, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint8Uint16PtrNoOrder(f func(*uint8) *uint16, list []*uint8, worker int) []*uint16 {
	chJobs := make(chan *uint8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint16, chJobs chan *uint8) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint16, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint8StrPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8StrPtr(f func(*uint8) *string, list []*uint8, optional ...Optional) []*string {
	if f == nil {
		return []*string{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8StrPtrNoOrder(f, list, worker)
		}
	}

	return pMapUint8StrPtrPreserveOrder(f, list, worker)
}

func pMapUint8StrPtrPreserveOrder(f func(*uint8) *string, list []*uint8, worker int) []*string {
	chJobs := make(chan map[int]*uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*string, chJobs chan map[int]*uint8) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*string{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*string, len(list))
	newList := make([]*string, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint8StrPtrNoOrder(f func(*uint8) *string, list []*uint8, worker int) []*string {
	chJobs := make(chan *uint8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *string, chJobs chan *uint8) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*string, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint8BoolPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8BoolPtr(f func(*uint8) *bool, list []*uint8, optional ...Optional) []*bool {
	if f == nil {
		return []*bool{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8BoolPtrNoOrder(f, list, worker)
		}
	}

	return pMapUint8BoolPtrPreserveOrder(f, list, worker)
}

func pMapUint8BoolPtrPreserveOrder(f func(*uint8) *bool, list []*uint8, worker int) []*bool {
	chJobs := make(chan map[int]*uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*bool, chJobs chan map[int]*uint8) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*bool{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*bool, len(list))
	newList := make([]*bool, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint8BoolPtrNoOrder(f func(*uint8) *bool, list []*uint8, worker int) []*bool {
	chJobs := make(chan *uint8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *bool, chJobs chan *uint8) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*bool, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint8Float32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Float32Ptr(f func(*uint8) *float32, list []*uint8, optional ...Optional) []*float32 {
	if f == nil {
		return []*float32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8Float32PtrNoOrder(f, list, worker)
		}
	}

	return pMapUint8Float32PtrPreserveOrder(f, list, worker)
}

func pMapUint8Float32PtrPreserveOrder(f func(*uint8) *float32, list []*uint8, worker int) []*float32 {
	chJobs := make(chan map[int]*uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float32, chJobs chan map[int]*uint8) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*float32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float32, len(list))
	newList := make([]*float32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint8Float32PtrNoOrder(f func(*uint8) *float32, list []*uint8, worker int) []*float32 {
	chJobs := make(chan *uint8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float32, chJobs chan *uint8) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*float32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint8Float64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Float64Ptr(f func(*uint8) *float64, list []*uint8, optional ...Optional) []*float64 {
	if f == nil {
		return []*float64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8Float64PtrNoOrder(f, list, worker)
		}
	}

	return pMapUint8Float64PtrPreserveOrder(f, list, worker)
}

func pMapUint8Float64PtrPreserveOrder(f func(*uint8) *float64, list []*uint8, worker int) []*float64 {
	chJobs := make(chan map[int]*uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float64, chJobs chan map[int]*uint8) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*float64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float64, len(list))
	newList := make([]*float64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint8Float64PtrNoOrder(f func(*uint8) *float64, list []*uint8, worker int) []*float64 {
	chJobs := make(chan *uint8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float64, chJobs chan *uint8) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*float64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapStrIntPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrIntPtr(f func(*string) *int, list []*string, optional ...Optional) []*int {
	if f == nil {
		return []*int{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrIntPtrNoOrder(f, list, worker)
		}
	}

	return pMapStrIntPtrPreserveOrder(f, list, worker)
}

func pMapStrIntPtrPreserveOrder(f func(*string) *int, list []*string, worker int) []*int {
	chJobs := make(chan map[int]*string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int, chJobs chan map[int]*string) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int, len(list))
	newList := make([]*int, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapStrIntPtrNoOrder(f func(*string) *int, list []*string, worker int) []*int {
	chJobs := make(chan *string, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int, chJobs chan *string) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapStrInt64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrInt64Ptr(f func(*string) *int64, list []*string, optional ...Optional) []*int64 {
	if f == nil {
		return []*int64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrInt64PtrNoOrder(f, list, worker)
		}
	}

	return pMapStrInt64PtrPreserveOrder(f, list, worker)
}

func pMapStrInt64PtrPreserveOrder(f func(*string) *int64, list []*string, worker int) []*int64 {
	chJobs := make(chan map[int]*string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int64, chJobs chan map[int]*string) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int64, len(list))
	newList := make([]*int64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapStrInt64PtrNoOrder(f func(*string) *int64, list []*string, worker int) []*int64 {
	chJobs := make(chan *string, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int64, chJobs chan *string) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapStrInt32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrInt32Ptr(f func(*string) *int32, list []*string, optional ...Optional) []*int32 {
	if f == nil {
		return []*int32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrInt32PtrNoOrder(f, list, worker)
		}
	}

	return pMapStrInt32PtrPreserveOrder(f, list, worker)
}

func pMapStrInt32PtrPreserveOrder(f func(*string) *int32, list []*string, worker int) []*int32 {
	chJobs := make(chan map[int]*string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int32, chJobs chan map[int]*string) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int32, len(list))
	newList := make([]*int32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapStrInt32PtrNoOrder(f func(*string) *int32, list []*string, worker int) []*int32 {
	chJobs := make(chan *string, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int32, chJobs chan *string) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapStrInt16Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrInt16Ptr(f func(*string) *int16, list []*string, optional ...Optional) []*int16 {
	if f == nil {
		return []*int16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrInt16PtrNoOrder(f, list, worker)
		}
	}

	return pMapStrInt16PtrPreserveOrder(f, list, worker)
}

func pMapStrInt16PtrPreserveOrder(f func(*string) *int16, list []*string, worker int) []*int16 {
	chJobs := make(chan map[int]*string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int16, chJobs chan map[int]*string) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int16{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int16, len(list))
	newList := make([]*int16, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapStrInt16PtrNoOrder(f func(*string) *int16, list []*string, worker int) []*int16 {
	chJobs := make(chan *string, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int16, chJobs chan *string) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int16, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapStrInt8Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrInt8Ptr(f func(*string) *int8, list []*string, optional ...Optional) []*int8 {
	if f == nil {
		return []*int8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrInt8PtrNoOrder(f, list, worker)
		}
	}

	return pMapStrInt8PtrPreserveOrder(f, list, worker)
}

func pMapStrInt8PtrPreserveOrder(f func(*string) *int8, list []*string, worker int) []*int8 {
	chJobs := make(chan map[int]*string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int8, chJobs chan map[int]*string) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int8{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int8, len(list))
	newList := make([]*int8, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapStrInt8PtrNoOrder(f func(*string) *int8, list []*string, worker int) []*int8 {
	chJobs := make(chan *string, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int8, chJobs chan *string) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int8, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapStrUintPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrUintPtr(f func(*string) *uint, list []*string, optional ...Optional) []*uint {
	if f == nil {
		return []*uint{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrUintPtrNoOrder(f, list, worker)
		}
	}

	return pMapStrUintPtrPreserveOrder(f, list, worker)
}

func pMapStrUintPtrPreserveOrder(f func(*string) *uint, list []*string, worker int) []*uint {
	chJobs := make(chan map[int]*string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint, chJobs chan map[int]*string) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint, len(list))
	newList := make([]*uint, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapStrUintPtrNoOrder(f func(*string) *uint, list []*string, worker int) []*uint {
	chJobs := make(chan *string, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint, chJobs chan *string) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapStrUint64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrUint64Ptr(f func(*string) *uint64, list []*string, optional ...Optional) []*uint64 {
	if f == nil {
		return []*uint64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrUint64PtrNoOrder(f, list, worker)
		}
	}

	return pMapStrUint64PtrPreserveOrder(f, list, worker)
}

func pMapStrUint64PtrPreserveOrder(f func(*string) *uint64, list []*string, worker int) []*uint64 {
	chJobs := make(chan map[int]*string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint64, chJobs chan map[int]*string) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint64, len(list))
	newList := make([]*uint64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapStrUint64PtrNoOrder(f func(*string) *uint64, list []*string, worker int) []*uint64 {
	chJobs := make(chan *string, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint64, chJobs chan *string) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapStrUint32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrUint32Ptr(f func(*string) *uint32, list []*string, optional ...Optional) []*uint32 {
	if f == nil {
		return []*uint32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrUint32PtrNoOrder(f, list, worker)
		}
	}

	return pMapStrUint32PtrPreserveOrder(f, list, worker)
}

func pMapStrUint32PtrPreserveOrder(f func(*string) *uint32, list []*string, worker int) []*uint32 {
	chJobs := make(chan map[int]*string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint32, chJobs chan map[int]*string) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint32, len(list))
	newList := make([]*uint32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapStrUint32PtrNoOrder(f func(*string) *uint32, list []*string, worker int) []*uint32 {
	chJobs := make(chan *string, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint32, chJobs chan *string) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapStrUint16Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrUint16Ptr(f func(*string) *uint16, list []*string, optional ...Optional) []*uint16 {
	if f == nil {
		return []*uint16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrUint16PtrNoOrder(f, list, worker)
		}
	}

	return pMapStrUint16PtrPreserveOrder(f, list, worker)
}

func pMapStrUint16PtrPreserveOrder(f func(*string) *uint16, list []*string, worker int) []*uint16 {
	chJobs := make(chan map[int]*string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint16, chJobs chan map[int]*string) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint16{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint16, len(list))
	newList := make([]*uint16, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapStrUint16PtrNoOrder(f func(*string) *uint16, list []*string, worker int) []*uint16 {
	chJobs := make(chan *string, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint16, chJobs chan *string) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint16, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapStrUint8Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrUint8Ptr(f func(*string) *uint8, list []*string, optional ...Optional) []*uint8 {
	if f == nil {
		return []*uint8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrUint8PtrNoOrder(f, list, worker)
		}
	}

	return pMapStrUint8PtrPreserveOrder(f, list, worker)
}

func pMapStrUint8PtrPreserveOrder(f func(*string) *uint8, list []*string, worker int) []*uint8 {
	chJobs := make(chan map[int]*string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint8, chJobs chan map[int]*string) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint8{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint8, len(list))
	newList := make([]*uint8, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapStrUint8PtrNoOrder(f func(*string) *uint8, list []*string, worker int) []*uint8 {
	chJobs := make(chan *string, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint8, chJobs chan *string) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint8, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapStrBoolPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrBoolPtr(f func(*string) *bool, list []*string, optional ...Optional) []*bool {
	if f == nil {
		return []*bool{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrBoolPtrNoOrder(f, list, worker)
		}
	}

	return pMapStrBoolPtrPreserveOrder(f, list, worker)
}

func pMapStrBoolPtrPreserveOrder(f func(*string) *bool, list []*string, worker int) []*bool {
	chJobs := make(chan map[int]*string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*bool, chJobs chan map[int]*string) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*bool{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*bool, len(list))
	newList := make([]*bool, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapStrBoolPtrNoOrder(f func(*string) *bool, list []*string, worker int) []*bool {
	chJobs := make(chan *string, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *bool, chJobs chan *string) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*bool, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapStrFloat32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrFloat32Ptr(f func(*string) *float32, list []*string, optional ...Optional) []*float32 {
	if f == nil {
		return []*float32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrFloat32PtrNoOrder(f, list, worker)
		}
	}

	return pMapStrFloat32PtrPreserveOrder(f, list, worker)
}

func pMapStrFloat32PtrPreserveOrder(f func(*string) *float32, list []*string, worker int) []*float32 {
	chJobs := make(chan map[int]*string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float32, chJobs chan map[int]*string) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*float32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float32, len(list))
	newList := make([]*float32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapStrFloat32PtrNoOrder(f func(*string) *float32, list []*string, worker int) []*float32 {
	chJobs := make(chan *string, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float32, chJobs chan *string) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*float32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapStrFloat64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrFloat64Ptr(f func(*string) *float64, list []*string, optional ...Optional) []*float64 {
	if f == nil {
		return []*float64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrFloat64PtrNoOrder(f, list, worker)
		}
	}

	return pMapStrFloat64PtrPreserveOrder(f, list, worker)
}

func pMapStrFloat64PtrPreserveOrder(f func(*string) *float64, list []*string, worker int) []*float64 {
	chJobs := make(chan map[int]*string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float64, chJobs chan map[int]*string) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*float64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float64, len(list))
	newList := make([]*float64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapStrFloat64PtrNoOrder(f func(*string) *float64, list []*string, worker int) []*float64 {
	chJobs := make(chan *string, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float64, chJobs chan *string) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*float64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapBoolIntPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolIntPtr(f func(*bool) *int, list []*bool, optional ...Optional) []*int {
	if f == nil {
		return []*int{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolIntPtrNoOrder(f, list, worker)
		}
	}

	return pMapBoolIntPtrPreserveOrder(f, list, worker)
}

func pMapBoolIntPtrPreserveOrder(f func(*bool) *int, list []*bool, worker int) []*int {
	chJobs := make(chan map[int]*bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int, chJobs chan map[int]*bool) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int, len(list))
	newList := make([]*int, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapBoolIntPtrNoOrder(f func(*bool) *int, list []*bool, worker int) []*int {
	chJobs := make(chan *bool, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int, chJobs chan *bool) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapBoolInt64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolInt64Ptr(f func(*bool) *int64, list []*bool, optional ...Optional) []*int64 {
	if f == nil {
		return []*int64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolInt64PtrNoOrder(f, list, worker)
		}
	}

	return pMapBoolInt64PtrPreserveOrder(f, list, worker)
}

func pMapBoolInt64PtrPreserveOrder(f func(*bool) *int64, list []*bool, worker int) []*int64 {
	chJobs := make(chan map[int]*bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int64, chJobs chan map[int]*bool) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int64, len(list))
	newList := make([]*int64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapBoolInt64PtrNoOrder(f func(*bool) *int64, list []*bool, worker int) []*int64 {
	chJobs := make(chan *bool, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int64, chJobs chan *bool) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapBoolInt32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolInt32Ptr(f func(*bool) *int32, list []*bool, optional ...Optional) []*int32 {
	if f == nil {
		return []*int32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolInt32PtrNoOrder(f, list, worker)
		}
	}

	return pMapBoolInt32PtrPreserveOrder(f, list, worker)
}

func pMapBoolInt32PtrPreserveOrder(f func(*bool) *int32, list []*bool, worker int) []*int32 {
	chJobs := make(chan map[int]*bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int32, chJobs chan map[int]*bool) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int32, len(list))
	newList := make([]*int32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapBoolInt32PtrNoOrder(f func(*bool) *int32, list []*bool, worker int) []*int32 {
	chJobs := make(chan *bool, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int32, chJobs chan *bool) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapBoolInt16Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolInt16Ptr(f func(*bool) *int16, list []*bool, optional ...Optional) []*int16 {
	if f == nil {
		return []*int16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolInt16PtrNoOrder(f, list, worker)
		}
	}

	return pMapBoolInt16PtrPreserveOrder(f, list, worker)
}

func pMapBoolInt16PtrPreserveOrder(f func(*bool) *int16, list []*bool, worker int) []*int16 {
	chJobs := make(chan map[int]*bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int16, chJobs chan map[int]*bool) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int16{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int16, len(list))
	newList := make([]*int16, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapBoolInt16PtrNoOrder(f func(*bool) *int16, list []*bool, worker int) []*int16 {
	chJobs := make(chan *bool, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int16, chJobs chan *bool) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int16, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapBoolInt8Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolInt8Ptr(f func(*bool) *int8, list []*bool, optional ...Optional) []*int8 {
	if f == nil {
		return []*int8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolInt8PtrNoOrder(f, list, worker)
		}
	}

	return pMapBoolInt8PtrPreserveOrder(f, list, worker)
}

func pMapBoolInt8PtrPreserveOrder(f func(*bool) *int8, list []*bool, worker int) []*int8 {
	chJobs := make(chan map[int]*bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int8, chJobs chan map[int]*bool) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int8{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int8, len(list))
	newList := make([]*int8, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapBoolInt8PtrNoOrder(f func(*bool) *int8, list []*bool, worker int) []*int8 {
	chJobs := make(chan *bool, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int8, chJobs chan *bool) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int8, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapBoolUintPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolUintPtr(f func(*bool) *uint, list []*bool, optional ...Optional) []*uint {
	if f == nil {
		return []*uint{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolUintPtrNoOrder(f, list, worker)
		}
	}

	return pMapBoolUintPtrPreserveOrder(f, list, worker)
}

func pMapBoolUintPtrPreserveOrder(f func(*bool) *uint, list []*bool, worker int) []*uint {
	chJobs := make(chan map[int]*bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint, chJobs chan map[int]*bool) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint, len(list))
	newList := make([]*uint, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapBoolUintPtrNoOrder(f func(*bool) *uint, list []*bool, worker int) []*uint {
	chJobs := make(chan *bool, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint, chJobs chan *bool) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapBoolUint64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolUint64Ptr(f func(*bool) *uint64, list []*bool, optional ...Optional) []*uint64 {
	if f == nil {
		return []*uint64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolUint64PtrNoOrder(f, list, worker)
		}
	}

	return pMapBoolUint64PtrPreserveOrder(f, list, worker)
}

func pMapBoolUint64PtrPreserveOrder(f func(*bool) *uint64, list []*bool, worker int) []*uint64 {
	chJobs := make(chan map[int]*bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint64, chJobs chan map[int]*bool) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint64, len(list))
	newList := make([]*uint64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapBoolUint64PtrNoOrder(f func(*bool) *uint64, list []*bool, worker int) []*uint64 {
	chJobs := make(chan *bool, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint64, chJobs chan *bool) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapBoolUint32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolUint32Ptr(f func(*bool) *uint32, list []*bool, optional ...Optional) []*uint32 {
	if f == nil {
		return []*uint32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolUint32PtrNoOrder(f, list, worker)
		}
	}

	return pMapBoolUint32PtrPreserveOrder(f, list, worker)
}

func pMapBoolUint32PtrPreserveOrder(f func(*bool) *uint32, list []*bool, worker int) []*uint32 {
	chJobs := make(chan map[int]*bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint32, chJobs chan map[int]*bool) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint32, len(list))
	newList := make([]*uint32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapBoolUint32PtrNoOrder(f func(*bool) *uint32, list []*bool, worker int) []*uint32 {
	chJobs := make(chan *bool, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint32, chJobs chan *bool) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapBoolUint16Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolUint16Ptr(f func(*bool) *uint16, list []*bool, optional ...Optional) []*uint16 {
	if f == nil {
		return []*uint16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolUint16PtrNoOrder(f, list, worker)
		}
	}

	return pMapBoolUint16PtrPreserveOrder(f, list, worker)
}

func pMapBoolUint16PtrPreserveOrder(f func(*bool) *uint16, list []*bool, worker int) []*uint16 {
	chJobs := make(chan map[int]*bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint16, chJobs chan map[int]*bool) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint16{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint16, len(list))
	newList := make([]*uint16, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapBoolUint16PtrNoOrder(f func(*bool) *uint16, list []*bool, worker int) []*uint16 {
	chJobs := make(chan *bool, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint16, chJobs chan *bool) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint16, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapBoolUint8Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolUint8Ptr(f func(*bool) *uint8, list []*bool, optional ...Optional) []*uint8 {
	if f == nil {
		return []*uint8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolUint8PtrNoOrder(f, list, worker)
		}
	}

	return pMapBoolUint8PtrPreserveOrder(f, list, worker)
}

func pMapBoolUint8PtrPreserveOrder(f func(*bool) *uint8, list []*bool, worker int) []*uint8 {
	chJobs := make(chan map[int]*bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint8, chJobs chan map[int]*bool) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint8{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint8, len(list))
	newList := make([]*uint8, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapBoolUint8PtrNoOrder(f func(*bool) *uint8, list []*bool, worker int) []*uint8 {
	chJobs := make(chan *bool, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint8, chJobs chan *bool) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint8, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapBoolStrPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolStrPtr(f func(*bool) *string, list []*bool, optional ...Optional) []*string {
	if f == nil {
		return []*string{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolStrPtrNoOrder(f, list, worker)
		}
	}

	return pMapBoolStrPtrPreserveOrder(f, list, worker)
}

func pMapBoolStrPtrPreserveOrder(f func(*bool) *string, list []*bool, worker int) []*string {
	chJobs := make(chan map[int]*bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*string, chJobs chan map[int]*bool) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*string{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*string, len(list))
	newList := make([]*string, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapBoolStrPtrNoOrder(f func(*bool) *string, list []*bool, worker int) []*string {
	chJobs := make(chan *bool, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *string, chJobs chan *bool) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*string, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapBoolFloat32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolFloat32Ptr(f func(*bool) *float32, list []*bool, optional ...Optional) []*float32 {
	if f == nil {
		return []*float32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolFloat32PtrNoOrder(f, list, worker)
		}
	}

	return pMapBoolFloat32PtrPreserveOrder(f, list, worker)
}

func pMapBoolFloat32PtrPreserveOrder(f func(*bool) *float32, list []*bool, worker int) []*float32 {
	chJobs := make(chan map[int]*bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float32, chJobs chan map[int]*bool) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*float32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float32, len(list))
	newList := make([]*float32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapBoolFloat32PtrNoOrder(f func(*bool) *float32, list []*bool, worker int) []*float32 {
	chJobs := make(chan *bool, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float32, chJobs chan *bool) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*float32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapBoolFloat64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolFloat64Ptr(f func(*bool) *float64, list []*bool, optional ...Optional) []*float64 {
	if f == nil {
		return []*float64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolFloat64PtrNoOrder(f, list, worker)
		}
	}

	return pMapBoolFloat64PtrPreserveOrder(f, list, worker)
}

func pMapBoolFloat64PtrPreserveOrder(f func(*bool) *float64, list []*bool, worker int) []*float64 {
	chJobs := make(chan map[int]*bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float64, chJobs chan map[int]*bool) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*float64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float64, len(list))
	newList := make([]*float64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapBoolFloat64PtrNoOrder(f func(*bool) *float64, list []*bool, worker int) []*float64 {
	chJobs := make(chan *bool, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float64, chJobs chan *bool) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*float64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapFloat32IntPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32IntPtr(f func(*float32) *int, list []*float32, optional ...Optional) []*int {
	if f == nil {
		return []*int{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32IntPtrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32IntPtrPreserveOrder(f, list, worker)
}

func pMapFloat32IntPtrPreserveOrder(f func(*float32) *int, list []*float32, worker int) []*int {
	chJobs := make(chan map[int]*float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int, chJobs chan map[int]*float32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int, len(list))
	newList := make([]*int, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapFloat32IntPtrNoOrder(f func(*float32) *int, list []*float32, worker int) []*int {
	chJobs := make(chan *float32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int, chJobs chan *float32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapFloat32Int64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Int64Ptr(f func(*float32) *int64, list []*float32, optional ...Optional) []*int64 {
	if f == nil {
		return []*int64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32Int64PtrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32Int64PtrPreserveOrder(f, list, worker)
}

func pMapFloat32Int64PtrPreserveOrder(f func(*float32) *int64, list []*float32, worker int) []*int64 {
	chJobs := make(chan map[int]*float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int64, chJobs chan map[int]*float32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int64, len(list))
	newList := make([]*int64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapFloat32Int64PtrNoOrder(f func(*float32) *int64, list []*float32, worker int) []*int64 {
	chJobs := make(chan *float32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int64, chJobs chan *float32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapFloat32Int32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Int32Ptr(f func(*float32) *int32, list []*float32, optional ...Optional) []*int32 {
	if f == nil {
		return []*int32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32Int32PtrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32Int32PtrPreserveOrder(f, list, worker)
}

func pMapFloat32Int32PtrPreserveOrder(f func(*float32) *int32, list []*float32, worker int) []*int32 {
	chJobs := make(chan map[int]*float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int32, chJobs chan map[int]*float32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int32, len(list))
	newList := make([]*int32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapFloat32Int32PtrNoOrder(f func(*float32) *int32, list []*float32, worker int) []*int32 {
	chJobs := make(chan *float32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int32, chJobs chan *float32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapFloat32Int16Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Int16Ptr(f func(*float32) *int16, list []*float32, optional ...Optional) []*int16 {
	if f == nil {
		return []*int16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32Int16PtrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32Int16PtrPreserveOrder(f, list, worker)
}

func pMapFloat32Int16PtrPreserveOrder(f func(*float32) *int16, list []*float32, worker int) []*int16 {
	chJobs := make(chan map[int]*float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int16, chJobs chan map[int]*float32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int16{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int16, len(list))
	newList := make([]*int16, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapFloat32Int16PtrNoOrder(f func(*float32) *int16, list []*float32, worker int) []*int16 {
	chJobs := make(chan *float32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int16, chJobs chan *float32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int16, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapFloat32Int8Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Int8Ptr(f func(*float32) *int8, list []*float32, optional ...Optional) []*int8 {
	if f == nil {
		return []*int8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32Int8PtrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32Int8PtrPreserveOrder(f, list, worker)
}

func pMapFloat32Int8PtrPreserveOrder(f func(*float32) *int8, list []*float32, worker int) []*int8 {
	chJobs := make(chan map[int]*float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int8, chJobs chan map[int]*float32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int8{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int8, len(list))
	newList := make([]*int8, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapFloat32Int8PtrNoOrder(f func(*float32) *int8, list []*float32, worker int) []*int8 {
	chJobs := make(chan *float32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int8, chJobs chan *float32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int8, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapFloat32UintPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32UintPtr(f func(*float32) *uint, list []*float32, optional ...Optional) []*uint {
	if f == nil {
		return []*uint{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32UintPtrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32UintPtrPreserveOrder(f, list, worker)
}

func pMapFloat32UintPtrPreserveOrder(f func(*float32) *uint, list []*float32, worker int) []*uint {
	chJobs := make(chan map[int]*float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint, chJobs chan map[int]*float32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint, len(list))
	newList := make([]*uint, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapFloat32UintPtrNoOrder(f func(*float32) *uint, list []*float32, worker int) []*uint {
	chJobs := make(chan *float32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint, chJobs chan *float32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapFloat32Uint64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Uint64Ptr(f func(*float32) *uint64, list []*float32, optional ...Optional) []*uint64 {
	if f == nil {
		return []*uint64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32Uint64PtrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32Uint64PtrPreserveOrder(f, list, worker)
}

func pMapFloat32Uint64PtrPreserveOrder(f func(*float32) *uint64, list []*float32, worker int) []*uint64 {
	chJobs := make(chan map[int]*float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint64, chJobs chan map[int]*float32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint64, len(list))
	newList := make([]*uint64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapFloat32Uint64PtrNoOrder(f func(*float32) *uint64, list []*float32, worker int) []*uint64 {
	chJobs := make(chan *float32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint64, chJobs chan *float32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapFloat32Uint32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Uint32Ptr(f func(*float32) *uint32, list []*float32, optional ...Optional) []*uint32 {
	if f == nil {
		return []*uint32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32Uint32PtrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32Uint32PtrPreserveOrder(f, list, worker)
}

func pMapFloat32Uint32PtrPreserveOrder(f func(*float32) *uint32, list []*float32, worker int) []*uint32 {
	chJobs := make(chan map[int]*float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint32, chJobs chan map[int]*float32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint32, len(list))
	newList := make([]*uint32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapFloat32Uint32PtrNoOrder(f func(*float32) *uint32, list []*float32, worker int) []*uint32 {
	chJobs := make(chan *float32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint32, chJobs chan *float32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapFloat32Uint16Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Uint16Ptr(f func(*float32) *uint16, list []*float32, optional ...Optional) []*uint16 {
	if f == nil {
		return []*uint16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32Uint16PtrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32Uint16PtrPreserveOrder(f, list, worker)
}

func pMapFloat32Uint16PtrPreserveOrder(f func(*float32) *uint16, list []*float32, worker int) []*uint16 {
	chJobs := make(chan map[int]*float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint16, chJobs chan map[int]*float32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint16{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint16, len(list))
	newList := make([]*uint16, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapFloat32Uint16PtrNoOrder(f func(*float32) *uint16, list []*float32, worker int) []*uint16 {
	chJobs := make(chan *float32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint16, chJobs chan *float32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint16, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapFloat32Uint8Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Uint8Ptr(f func(*float32) *uint8, list []*float32, optional ...Optional) []*uint8 {
	if f == nil {
		return []*uint8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32Uint8PtrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32Uint8PtrPreserveOrder(f, list, worker)
}

func pMapFloat32Uint8PtrPreserveOrder(f func(*float32) *uint8, list []*float32, worker int) []*uint8 {
	chJobs := make(chan map[int]*float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint8, chJobs chan map[int]*float32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint8{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint8, len(list))
	newList := make([]*uint8, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapFloat32Uint8PtrNoOrder(f func(*float32) *uint8, list []*float32, worker int) []*uint8 {
	chJobs := make(chan *float32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint8, chJobs chan *float32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint8, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapFloat32StrPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32StrPtr(f func(*float32) *string, list []*float32, optional ...Optional) []*string {
	if f == nil {
		return []*string{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32StrPtrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32StrPtrPreserveOrder(f, list, worker)
}

func pMapFloat32StrPtrPreserveOrder(f func(*float32) *string, list []*float32, worker int) []*string {
	chJobs := make(chan map[int]*float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*string, chJobs chan map[int]*float32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*string{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*string, len(list))
	newList := make([]*string, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapFloat32StrPtrNoOrder(f func(*float32) *string, list []*float32, worker int) []*string {
	chJobs := make(chan *float32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *string, chJobs chan *float32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*string, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapFloat32BoolPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32BoolPtr(f func(*float32) *bool, list []*float32, optional ...Optional) []*bool {
	if f == nil {
		return []*bool{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32BoolPtrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32BoolPtrPreserveOrder(f, list, worker)
}

func pMapFloat32BoolPtrPreserveOrder(f func(*float32) *bool, list []*float32, worker int) []*bool {
	chJobs := make(chan map[int]*float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*bool, chJobs chan map[int]*float32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*bool{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*bool, len(list))
	newList := make([]*bool, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapFloat32BoolPtrNoOrder(f func(*float32) *bool, list []*float32, worker int) []*bool {
	chJobs := make(chan *float32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *bool, chJobs chan *float32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*bool, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapFloat32Float64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Float64Ptr(f func(*float32) *float64, list []*float32, optional ...Optional) []*float64 {
	if f == nil {
		return []*float64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32Float64PtrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32Float64PtrPreserveOrder(f, list, worker)
}

func pMapFloat32Float64PtrPreserveOrder(f func(*float32) *float64, list []*float32, worker int) []*float64 {
	chJobs := make(chan map[int]*float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float64, chJobs chan map[int]*float32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*float64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float64, len(list))
	newList := make([]*float64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapFloat32Float64PtrNoOrder(f func(*float32) *float64, list []*float32, worker int) []*float64 {
	chJobs := make(chan *float32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float64, chJobs chan *float32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*float64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapFloat64IntPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64IntPtr(f func(*float64) *int, list []*float64, optional ...Optional) []*int {
	if f == nil {
		return []*int{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64IntPtrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64IntPtrPreserveOrder(f, list, worker)
}

func pMapFloat64IntPtrPreserveOrder(f func(*float64) *int, list []*float64, worker int) []*int {
	chJobs := make(chan map[int]*float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int, chJobs chan map[int]*float64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int, len(list))
	newList := make([]*int, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapFloat64IntPtrNoOrder(f func(*float64) *int, list []*float64, worker int) []*int {
	chJobs := make(chan *float64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int, chJobs chan *float64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapFloat64Int64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Int64Ptr(f func(*float64) *int64, list []*float64, optional ...Optional) []*int64 {
	if f == nil {
		return []*int64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64Int64PtrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64Int64PtrPreserveOrder(f, list, worker)
}

func pMapFloat64Int64PtrPreserveOrder(f func(*float64) *int64, list []*float64, worker int) []*int64 {
	chJobs := make(chan map[int]*float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int64, chJobs chan map[int]*float64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int64, len(list))
	newList := make([]*int64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapFloat64Int64PtrNoOrder(f func(*float64) *int64, list []*float64, worker int) []*int64 {
	chJobs := make(chan *float64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int64, chJobs chan *float64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapFloat64Int32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Int32Ptr(f func(*float64) *int32, list []*float64, optional ...Optional) []*int32 {
	if f == nil {
		return []*int32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64Int32PtrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64Int32PtrPreserveOrder(f, list, worker)
}

func pMapFloat64Int32PtrPreserveOrder(f func(*float64) *int32, list []*float64, worker int) []*int32 {
	chJobs := make(chan map[int]*float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int32, chJobs chan map[int]*float64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int32, len(list))
	newList := make([]*int32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapFloat64Int32PtrNoOrder(f func(*float64) *int32, list []*float64, worker int) []*int32 {
	chJobs := make(chan *float64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int32, chJobs chan *float64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapFloat64Int16Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Int16Ptr(f func(*float64) *int16, list []*float64, optional ...Optional) []*int16 {
	if f == nil {
		return []*int16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64Int16PtrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64Int16PtrPreserveOrder(f, list, worker)
}

func pMapFloat64Int16PtrPreserveOrder(f func(*float64) *int16, list []*float64, worker int) []*int16 {
	chJobs := make(chan map[int]*float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int16, chJobs chan map[int]*float64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int16{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int16, len(list))
	newList := make([]*int16, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapFloat64Int16PtrNoOrder(f func(*float64) *int16, list []*float64, worker int) []*int16 {
	chJobs := make(chan *float64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int16, chJobs chan *float64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int16, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapFloat64Int8Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Int8Ptr(f func(*float64) *int8, list []*float64, optional ...Optional) []*int8 {
	if f == nil {
		return []*int8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64Int8PtrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64Int8PtrPreserveOrder(f, list, worker)
}

func pMapFloat64Int8PtrPreserveOrder(f func(*float64) *int8, list []*float64, worker int) []*int8 {
	chJobs := make(chan map[int]*float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int8, chJobs chan map[int]*float64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*int8{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int8, len(list))
	newList := make([]*int8, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapFloat64Int8PtrNoOrder(f func(*float64) *int8, list []*float64, worker int) []*int8 {
	chJobs := make(chan *float64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int8, chJobs chan *float64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*int8, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapFloat64UintPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64UintPtr(f func(*float64) *uint, list []*float64, optional ...Optional) []*uint {
	if f == nil {
		return []*uint{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64UintPtrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64UintPtrPreserveOrder(f, list, worker)
}

func pMapFloat64UintPtrPreserveOrder(f func(*float64) *uint, list []*float64, worker int) []*uint {
	chJobs := make(chan map[int]*float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint, chJobs chan map[int]*float64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint, len(list))
	newList := make([]*uint, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapFloat64UintPtrNoOrder(f func(*float64) *uint, list []*float64, worker int) []*uint {
	chJobs := make(chan *float64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint, chJobs chan *float64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapFloat64Uint64Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Uint64Ptr(f func(*float64) *uint64, list []*float64, optional ...Optional) []*uint64 {
	if f == nil {
		return []*uint64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64Uint64PtrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64Uint64PtrPreserveOrder(f, list, worker)
}

func pMapFloat64Uint64PtrPreserveOrder(f func(*float64) *uint64, list []*float64, worker int) []*uint64 {
	chJobs := make(chan map[int]*float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint64, chJobs chan map[int]*float64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint64, len(list))
	newList := make([]*uint64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapFloat64Uint64PtrNoOrder(f func(*float64) *uint64, list []*float64, worker int) []*uint64 {
	chJobs := make(chan *float64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint64, chJobs chan *float64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapFloat64Uint32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Uint32Ptr(f func(*float64) *uint32, list []*float64, optional ...Optional) []*uint32 {
	if f == nil {
		return []*uint32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64Uint32PtrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64Uint32PtrPreserveOrder(f, list, worker)
}

func pMapFloat64Uint32PtrPreserveOrder(f func(*float64) *uint32, list []*float64, worker int) []*uint32 {
	chJobs := make(chan map[int]*float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint32, chJobs chan map[int]*float64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint32, len(list))
	newList := make([]*uint32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapFloat64Uint32PtrNoOrder(f func(*float64) *uint32, list []*float64, worker int) []*uint32 {
	chJobs := make(chan *float64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint32, chJobs chan *float64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapFloat64Uint16Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Uint16Ptr(f func(*float64) *uint16, list []*float64, optional ...Optional) []*uint16 {
	if f == nil {
		return []*uint16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64Uint16PtrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64Uint16PtrPreserveOrder(f, list, worker)
}

func pMapFloat64Uint16PtrPreserveOrder(f func(*float64) *uint16, list []*float64, worker int) []*uint16 {
	chJobs := make(chan map[int]*float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint16, chJobs chan map[int]*float64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint16{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint16, len(list))
	newList := make([]*uint16, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapFloat64Uint16PtrNoOrder(f func(*float64) *uint16, list []*float64, worker int) []*uint16 {
	chJobs := make(chan *float64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint16, chJobs chan *float64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint16, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapFloat64Uint8Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Uint8Ptr(f func(*float64) *uint8, list []*float64, optional ...Optional) []*uint8 {
	if f == nil {
		return []*uint8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64Uint8PtrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64Uint8PtrPreserveOrder(f, list, worker)
}

func pMapFloat64Uint8PtrPreserveOrder(f func(*float64) *uint8, list []*float64, worker int) []*uint8 {
	chJobs := make(chan map[int]*float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint8, chJobs chan map[int]*float64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*uint8{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint8, len(list))
	newList := make([]*uint8, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapFloat64Uint8PtrNoOrder(f func(*float64) *uint8, list []*float64, worker int) []*uint8 {
	chJobs := make(chan *float64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint8, chJobs chan *float64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*uint8, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapFloat64StrPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64StrPtr(f func(*float64) *string, list []*float64, optional ...Optional) []*string {
	if f == nil {
		return []*string{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64StrPtrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64StrPtrPreserveOrder(f, list, worker)
}

func pMapFloat64StrPtrPreserveOrder(f func(*float64) *string, list []*float64, worker int) []*string {
	chJobs := make(chan map[int]*float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*string, chJobs chan map[int]*float64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*string{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*string, len(list))
	newList := make([]*string, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapFloat64StrPtrNoOrder(f func(*float64) *string, list []*float64, worker int) []*string {
	chJobs := make(chan *float64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *string, chJobs chan *float64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*string, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapFloat64BoolPtr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64BoolPtr(f func(*float64) *bool, list []*float64, optional ...Optional) []*bool {
	if f == nil {
		return []*bool{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64BoolPtrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64BoolPtrPreserveOrder(f, list, worker)
}

func pMapFloat64BoolPtrPreserveOrder(f func(*float64) *bool, list []*float64, worker int) []*bool {
	chJobs := make(chan map[int]*float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*bool, chJobs chan map[int]*float64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*bool{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*bool, len(list))
	newList := make([]*bool, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapFloat64BoolPtrNoOrder(f func(*float64) *bool, list []*float64, worker int) []*bool {
	chJobs := make(chan *float64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *bool, chJobs chan *float64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*bool, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapFloat64Float32Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Float32Ptr(f func(*float64) *float32, list []*float64, optional ...Optional) []*float32 {
	if f == nil {
		return []*float32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64Float32PtrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64Float32PtrPreserveOrder(f, list, worker)
}

func pMapFloat64Float32PtrPreserveOrder(f func(*float64) *float32, list []*float64, worker int) []*float32 {
	chJobs := make(chan map[int]*float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float32, chJobs chan map[int]*float64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*float32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float32, len(list))
	newList := make([]*float32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapFloat64Float32PtrNoOrder(f func(*float64) *float32, list []*float64, worker int) []*float32 {
	chJobs := make(chan *float64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float32, chJobs chan *float64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*float32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}
