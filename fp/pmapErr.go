package fp

import "sync"

// PMapIntErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntErr(f func(int) (int, error), list []int, optional ...Optional) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntErrNoOrder(f, list, worker)
		}
	}

	return pMapIntErrPreserveOrder(f, list, worker)
}

func pMapIntErrPreserveOrder(f func(int) (int, error), list []int, worker int) ([]int, error) {
	chJobs := make(chan map[int]int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int, chJobs chan map[int]int, errCh chan error) {
			defer wg.Done()
			if len(errCh) > 0 {
				return
			}
			for m := range chJobs {
				for k, v := range m {
					r, err := f(v)
					if err != nil {
						errCh <- err
						return
					}
					chResult <- map[int]int{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]int, len(list))
	newList := make([]int, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []int{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []int{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapIntErrNoOrder(f func(int) (int, error), list []int, worker int) ([]int, error) {
	chJobs := make(chan int, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan int, chJobs chan int, errCh chan error) {
			defer wg.Done()
			if len(errCh) > 0 {
				return
			}
			for v := range chJobs {
				r, err := f(v)
				if err != nil {
					errCh <- err
					return
				}
				chResult <- r
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]int, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []int{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []int{}, <-errCh
	}

	return newList, nil
}

// PMapInt64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Err(f func(int64) (int64, error), list []int64, optional ...Optional) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt64ErrPreserveOrder(f, list, worker)
}

func pMapInt64ErrPreserveOrder(f func(int64) (int64, error), list []int64, worker int) ([]int64, error) {
	chJobs := make(chan map[int]int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int64, chJobs chan map[int]int64, errCh chan error) {
			defer wg.Done()
			if len(errCh) > 0 {
				return
			}
			for m := range chJobs {
				for k, v := range m {
					r, err := f(v)
					if err != nil {
						errCh <- err
						return
					}
					chResult <- map[int]int64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]int64, len(list))
	newList := make([]int64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []int64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []int64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt64ErrNoOrder(f func(int64) (int64, error), list []int64, worker int) ([]int64, error) {
	chJobs := make(chan int64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan int64, chJobs chan int64, errCh chan error) {
			defer wg.Done()
			if len(errCh) > 0 {
				return
			}
			for v := range chJobs {
				r, err := f(v)
				if err != nil {
					errCh <- err
					return
				}
				chResult <- r
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]int64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []int64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []int64{}, <-errCh
	}

	return newList, nil
}

// PMapInt32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Err(f func(int32) (int32, error), list []int32, optional ...Optional) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt32ErrPreserveOrder(f, list, worker)
}

func pMapInt32ErrPreserveOrder(f func(int32) (int32, error), list []int32, worker int) ([]int32, error) {
	chJobs := make(chan map[int]int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int32, chJobs chan map[int]int32, errCh chan error) {
			defer wg.Done()
			if len(errCh) > 0 {
				return
			}
			for m := range chJobs {
				for k, v := range m {
					r, err := f(v)
					if err != nil {
						errCh <- err
						return
					}
					chResult <- map[int]int32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]int32, len(list))
	newList := make([]int32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []int32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []int32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt32ErrNoOrder(f func(int32) (int32, error), list []int32, worker int) ([]int32, error) {
	chJobs := make(chan int32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan int32, chJobs chan int32, errCh chan error) {
			defer wg.Done()
			if len(errCh) > 0 {
				return
			}
			for v := range chJobs {
				r, err := f(v)
				if err != nil {
					errCh <- err
					return
				}
				chResult <- r
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]int32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []int32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []int32{}, <-errCh
	}

	return newList, nil
}

// PMapInt16Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Err(f func(int16) (int16, error), list []int16, optional ...Optional) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt16ErrPreserveOrder(f, list, worker)
}

func pMapInt16ErrPreserveOrder(f func(int16) (int16, error), list []int16, worker int) ([]int16, error) {
	chJobs := make(chan map[int]int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int16, chJobs chan map[int]int16, errCh chan error) {
			defer wg.Done()
			if len(errCh) > 0 {
				return
			}
			for m := range chJobs {
				for k, v := range m {
					r, err := f(v)
					if err != nil {
						errCh <- err
						return
					}
					chResult <- map[int]int16{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]int16, len(list))
	newList := make([]int16, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []int16{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []int16{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt16ErrNoOrder(f func(int16) (int16, error), list []int16, worker int) ([]int16, error) {
	chJobs := make(chan int16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan int16, chJobs chan int16, errCh chan error) {
			defer wg.Done()
			if len(errCh) > 0 {
				return
			}
			for v := range chJobs {
				r, err := f(v)
				if err != nil {
					errCh <- err
					return
				}
				chResult <- r
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]int16, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []int16{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []int16{}, <-errCh
	}

	return newList, nil
}

// PMapInt8Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Err(f func(int8) (int8, error), list []int8, optional ...Optional) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt8ErrPreserveOrder(f, list, worker)
}

func pMapInt8ErrPreserveOrder(f func(int8) (int8, error), list []int8, worker int) ([]int8, error) {
	chJobs := make(chan map[int]int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int8, chJobs chan map[int]int8, errCh chan error) {
			defer wg.Done()
			if len(errCh) > 0 {
				return
			}
			for m := range chJobs {
				for k, v := range m {
					r, err := f(v)
					if err != nil {
						errCh <- err
						return
					}
					chResult <- map[int]int8{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]int8, len(list))
	newList := make([]int8, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []int8{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []int8{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt8ErrNoOrder(f func(int8) (int8, error), list []int8, worker int) ([]int8, error) {
	chJobs := make(chan int8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan int8, chJobs chan int8, errCh chan error) {
			defer wg.Done()
			if len(errCh) > 0 {
				return
			}
			for v := range chJobs {
				r, err := f(v)
				if err != nil {
					errCh <- err
					return
				}
				chResult <- r
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]int8, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []int8{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []int8{}, <-errCh
	}

	return newList, nil
}

// PMapUintErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintErr(f func(uint) (uint, error), list []uint, optional ...Optional) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintErrNoOrder(f, list, worker)
		}
	}

	return pMapUintErrPreserveOrder(f, list, worker)
}

func pMapUintErrPreserveOrder(f func(uint) (uint, error), list []uint, worker int) ([]uint, error) {
	chJobs := make(chan map[int]uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint, chJobs chan map[int]uint, errCh chan error) {
			defer wg.Done()
			if len(errCh) > 0 {
				return
			}
			for m := range chJobs {
				for k, v := range m {
					r, err := f(v)
					if err != nil {
						errCh <- err
						return
					}
					chResult <- map[int]uint{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]uint, len(list))
	newList := make([]uint, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []uint{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []uint{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUintErrNoOrder(f func(uint) (uint, error), list []uint, worker int) ([]uint, error) {
	chJobs := make(chan uint, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan uint, chJobs chan uint, errCh chan error) {
			defer wg.Done()
			if len(errCh) > 0 {
				return
			}
			for v := range chJobs {
				r, err := f(v)
				if err != nil {
					errCh <- err
					return
				}
				chResult <- r
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]uint, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []uint{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []uint{}, <-errCh
	}

	return newList, nil
}

// PMapUint64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Err(f func(uint64) (uint64, error), list []uint64, optional ...Optional) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint64ErrPreserveOrder(f, list, worker)
}

func pMapUint64ErrPreserveOrder(f func(uint64) (uint64, error), list []uint64, worker int) ([]uint64, error) {
	chJobs := make(chan map[int]uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint64, chJobs chan map[int]uint64, errCh chan error) {
			defer wg.Done()
			if len(errCh) > 0 {
				return
			}
			for m := range chJobs {
				for k, v := range m {
					r, err := f(v)
					if err != nil {
						errCh <- err
						return
					}
					chResult <- map[int]uint64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]uint64, len(list))
	newList := make([]uint64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []uint64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []uint64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint64ErrNoOrder(f func(uint64) (uint64, error), list []uint64, worker int) ([]uint64, error) {
	chJobs := make(chan uint64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan uint64, chJobs chan uint64, errCh chan error) {
			defer wg.Done()
			if len(errCh) > 0 {
				return
			}
			for v := range chJobs {
				r, err := f(v)
				if err != nil {
					errCh <- err
					return
				}
				chResult <- r
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]uint64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []uint64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []uint64{}, <-errCh
	}

	return newList, nil
}

// PMapUint32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Err(f func(uint32) (uint32, error), list []uint32, optional ...Optional) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint32ErrPreserveOrder(f, list, worker)
}

func pMapUint32ErrPreserveOrder(f func(uint32) (uint32, error), list []uint32, worker int) ([]uint32, error) {
	chJobs := make(chan map[int]uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint32, chJobs chan map[int]uint32, errCh chan error) {
			defer wg.Done()
			if len(errCh) > 0 {
				return
			}
			for m := range chJobs {
				for k, v := range m {
					r, err := f(v)
					if err != nil {
						errCh <- err
						return
					}
					chResult <- map[int]uint32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]uint32, len(list))
	newList := make([]uint32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []uint32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []uint32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint32ErrNoOrder(f func(uint32) (uint32, error), list []uint32, worker int) ([]uint32, error) {
	chJobs := make(chan uint32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan uint32, chJobs chan uint32, errCh chan error) {
			defer wg.Done()
			if len(errCh) > 0 {
				return
			}
			for v := range chJobs {
				r, err := f(v)
				if err != nil {
					errCh <- err
					return
				}
				chResult <- r
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]uint32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []uint32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []uint32{}, <-errCh
	}

	return newList, nil
}

// PMapUint16Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Err(f func(uint16) (uint16, error), list []uint16, optional ...Optional) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint16ErrPreserveOrder(f, list, worker)
}

func pMapUint16ErrPreserveOrder(f func(uint16) (uint16, error), list []uint16, worker int) ([]uint16, error) {
	chJobs := make(chan map[int]uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint16, chJobs chan map[int]uint16, errCh chan error) {
			defer wg.Done()
			if len(errCh) > 0 {
				return
			}
			for m := range chJobs {
				for k, v := range m {
					r, err := f(v)
					if err != nil {
						errCh <- err
						return
					}
					chResult <- map[int]uint16{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]uint16, len(list))
	newList := make([]uint16, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []uint16{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []uint16{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint16ErrNoOrder(f func(uint16) (uint16, error), list []uint16, worker int) ([]uint16, error) {
	chJobs := make(chan uint16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan uint16, chJobs chan uint16, errCh chan error) {
			defer wg.Done()
			if len(errCh) > 0 {
				return
			}
			for v := range chJobs {
				r, err := f(v)
				if err != nil {
					errCh <- err
					return
				}
				chResult <- r
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]uint16, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []uint16{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []uint16{}, <-errCh
	}

	return newList, nil
}

// PMapUint8Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Err(f func(uint8) (uint8, error), list []uint8, optional ...Optional) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint8ErrPreserveOrder(f, list, worker)
}

func pMapUint8ErrPreserveOrder(f func(uint8) (uint8, error), list []uint8, worker int) ([]uint8, error) {
	chJobs := make(chan map[int]uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint8, chJobs chan map[int]uint8, errCh chan error) {
			defer wg.Done()
			if len(errCh) > 0 {
				return
			}
			for m := range chJobs {
				for k, v := range m {
					r, err := f(v)
					if err != nil {
						errCh <- err
						return
					}
					chResult <- map[int]uint8{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]uint8, len(list))
	newList := make([]uint8, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []uint8{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []uint8{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint8ErrNoOrder(f func(uint8) (uint8, error), list []uint8, worker int) ([]uint8, error) {
	chJobs := make(chan uint8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan uint8, chJobs chan uint8, errCh chan error) {
			defer wg.Done()
			if len(errCh) > 0 {
				return
			}
			for v := range chJobs {
				r, err := f(v)
				if err != nil {
					errCh <- err
					return
				}
				chResult <- r
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]uint8, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []uint8{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []uint8{}, <-errCh
	}

	return newList, nil
}

// PMapStrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrErr(f func(string) (string, error), list []string, optional ...Optional) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrErrNoOrder(f, list, worker)
		}
	}

	return pMapStrErrPreserveOrder(f, list, worker)
}

func pMapStrErrPreserveOrder(f func(string) (string, error), list []string, worker int) ([]string, error) {
	chJobs := make(chan map[int]string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]string, chJobs chan map[int]string, errCh chan error) {
			defer wg.Done()
			if len(errCh) > 0 {
				return
			}
			for m := range chJobs {
				for k, v := range m {
					r, err := f(v)
					if err != nil {
						errCh <- err
						return
					}
					chResult <- map[int]string{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]string, len(list))
	newList := make([]string, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []string{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []string{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapStrErrNoOrder(f func(string) (string, error), list []string, worker int) ([]string, error) {
	chJobs := make(chan string, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan string, chJobs chan string, errCh chan error) {
			defer wg.Done()
			if len(errCh) > 0 {
				return
			}
			for v := range chJobs {
				r, err := f(v)
				if err != nil {
					errCh <- err
					return
				}
				chResult <- r
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]string, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []string{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []string{}, <-errCh
	}

	return newList, nil
}

// PMapBoolErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolErr(f func(bool) (bool, error), list []bool, optional ...Optional) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolErrNoOrder(f, list, worker)
		}
	}

	return pMapBoolErrPreserveOrder(f, list, worker)
}

func pMapBoolErrPreserveOrder(f func(bool) (bool, error), list []bool, worker int) ([]bool, error) {
	chJobs := make(chan map[int]bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]bool, chJobs chan map[int]bool, errCh chan error) {
			defer wg.Done()
			if len(errCh) > 0 {
				return
			}
			for m := range chJobs {
				for k, v := range m {
					r, err := f(v)
					if err != nil {
						errCh <- err
						return
					}
					chResult <- map[int]bool{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]bool, len(list))
	newList := make([]bool, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []bool{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []bool{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapBoolErrNoOrder(f func(bool) (bool, error), list []bool, worker int) ([]bool, error) {
	chJobs := make(chan bool, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan bool, chJobs chan bool, errCh chan error) {
			defer wg.Done()
			if len(errCh) > 0 {
				return
			}
			for v := range chJobs {
				r, err := f(v)
				if err != nil {
					errCh <- err
					return
				}
				chResult <- r
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]bool, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []bool{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []bool{}, <-errCh
	}

	return newList, nil
}

// PMapFloat32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Err(f func(float32) (float32, error), list []float32, optional ...Optional) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32ErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32ErrPreserveOrder(f, list, worker)
}

func pMapFloat32ErrPreserveOrder(f func(float32) (float32, error), list []float32, worker int) ([]float32, error) {
	chJobs := make(chan map[int]float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float32, chJobs chan map[int]float32, errCh chan error) {
			defer wg.Done()
			if len(errCh) > 0 {
				return
			}
			for m := range chJobs {
				for k, v := range m {
					r, err := f(v)
					if err != nil {
						errCh <- err
						return
					}
					chResult <- map[int]float32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]float32, len(list))
	newList := make([]float32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []float32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []float32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapFloat32ErrNoOrder(f func(float32) (float32, error), list []float32, worker int) ([]float32, error) {
	chJobs := make(chan float32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan float32, chJobs chan float32, errCh chan error) {
			defer wg.Done()
			if len(errCh) > 0 {
				return
			}
			for v := range chJobs {
				r, err := f(v)
				if err != nil {
					errCh <- err
					return
				}
				chResult <- r
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]float32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []float32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []float32{}, <-errCh
	}

	return newList, nil
}

// PMapFloat64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Err(f func(float64) (float64, error), list []float64, optional ...Optional) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64ErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64ErrPreserveOrder(f, list, worker)
}

func pMapFloat64ErrPreserveOrder(f func(float64) (float64, error), list []float64, worker int) ([]float64, error) {
	chJobs := make(chan map[int]float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float64, chJobs chan map[int]float64, errCh chan error) {
			defer wg.Done()
			if len(errCh) > 0 {
				return
			}
			for m := range chJobs {
				for k, v := range m {
					r, err := f(v)
					if err != nil {
						errCh <- err
						return
					}
					chResult <- map[int]float64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]float64, len(list))
	newList := make([]float64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []float64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []float64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapFloat64ErrNoOrder(f func(float64) (float64, error), list []float64, worker int) ([]float64, error) {
	chJobs := make(chan float64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan float64, chJobs chan float64, errCh chan error) {
			defer wg.Done()
			if len(errCh) > 0 {
				return
			}
			for v := range chJobs {
				r, err := f(v)
				if err != nil {
					errCh <- err
					return
				}
				chResult <- r
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]float64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []float64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []float64{}, <-errCh
	}

	return newList, nil
}
