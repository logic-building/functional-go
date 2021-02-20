package fp

import "sync"

// PMapIntInt64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntInt64Err(f func(int) (int64, error), list []int, optional ...Optional) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntInt64ErrNoOrder(f, list, worker)
		}
	}

	return pMapIntInt64ErrPreserveOrder(f, list, worker)
}

func pMapIntInt64ErrPreserveOrder(f func(int) (int64, error), list []int, worker int) ([]int64, error) {
	chJobs := make(chan map[int]int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int64, chJobs chan map[int]int, errCh chan error) {
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

func pMapIntInt64ErrNoOrder(f func(int) (int64, error), list []int, worker int) ([]int64, error) {
	chJobs := make(chan int, len(list))
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

		go func(chResult chan int64, chJobs chan int, errCh chan error) {
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

// PMapIntInt32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntInt32Err(f func(int) (int32, error), list []int, optional ...Optional) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntInt32ErrNoOrder(f, list, worker)
		}
	}

	return pMapIntInt32ErrPreserveOrder(f, list, worker)
}

func pMapIntInt32ErrPreserveOrder(f func(int) (int32, error), list []int, worker int) ([]int32, error) {
	chJobs := make(chan map[int]int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int32, chJobs chan map[int]int, errCh chan error) {
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

func pMapIntInt32ErrNoOrder(f func(int) (int32, error), list []int, worker int) ([]int32, error) {
	chJobs := make(chan int, len(list))
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

		go func(chResult chan int32, chJobs chan int, errCh chan error) {
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

// PMapIntInt16Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntInt16Err(f func(int) (int16, error), list []int, optional ...Optional) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntInt16ErrNoOrder(f, list, worker)
		}
	}

	return pMapIntInt16ErrPreserveOrder(f, list, worker)
}

func pMapIntInt16ErrPreserveOrder(f func(int) (int16, error), list []int, worker int) ([]int16, error) {
	chJobs := make(chan map[int]int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int16, chJobs chan map[int]int, errCh chan error) {
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

func pMapIntInt16ErrNoOrder(f func(int) (int16, error), list []int, worker int) ([]int16, error) {
	chJobs := make(chan int, len(list))
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

		go func(chResult chan int16, chJobs chan int, errCh chan error) {
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

// PMapIntInt8Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntInt8Err(f func(int) (int8, error), list []int, optional ...Optional) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntInt8ErrNoOrder(f, list, worker)
		}
	}

	return pMapIntInt8ErrPreserveOrder(f, list, worker)
}

func pMapIntInt8ErrPreserveOrder(f func(int) (int8, error), list []int, worker int) ([]int8, error) {
	chJobs := make(chan map[int]int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int8, chJobs chan map[int]int, errCh chan error) {
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

func pMapIntInt8ErrNoOrder(f func(int) (int8, error), list []int, worker int) ([]int8, error) {
	chJobs := make(chan int, len(list))
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

		go func(chResult chan int8, chJobs chan int, errCh chan error) {
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

// PMapIntUintErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntUintErr(f func(int) (uint, error), list []int, optional ...Optional) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntUintErrNoOrder(f, list, worker)
		}
	}

	return pMapIntUintErrPreserveOrder(f, list, worker)
}

func pMapIntUintErrPreserveOrder(f func(int) (uint, error), list []int, worker int) ([]uint, error) {
	chJobs := make(chan map[int]int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint, chJobs chan map[int]int, errCh chan error) {
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

func pMapIntUintErrNoOrder(f func(int) (uint, error), list []int, worker int) ([]uint, error) {
	chJobs := make(chan int, len(list))
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

		go func(chResult chan uint, chJobs chan int, errCh chan error) {
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

// PMapIntUint64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntUint64Err(f func(int) (uint64, error), list []int, optional ...Optional) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntUint64ErrNoOrder(f, list, worker)
		}
	}

	return pMapIntUint64ErrPreserveOrder(f, list, worker)
}

func pMapIntUint64ErrPreserveOrder(f func(int) (uint64, error), list []int, worker int) ([]uint64, error) {
	chJobs := make(chan map[int]int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint64, chJobs chan map[int]int, errCh chan error) {
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

func pMapIntUint64ErrNoOrder(f func(int) (uint64, error), list []int, worker int) ([]uint64, error) {
	chJobs := make(chan int, len(list))
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

		go func(chResult chan uint64, chJobs chan int, errCh chan error) {
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

// PMapIntUint32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntUint32Err(f func(int) (uint32, error), list []int, optional ...Optional) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntUint32ErrNoOrder(f, list, worker)
		}
	}

	return pMapIntUint32ErrPreserveOrder(f, list, worker)
}

func pMapIntUint32ErrPreserveOrder(f func(int) (uint32, error), list []int, worker int) ([]uint32, error) {
	chJobs := make(chan map[int]int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint32, chJobs chan map[int]int, errCh chan error) {
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

func pMapIntUint32ErrNoOrder(f func(int) (uint32, error), list []int, worker int) ([]uint32, error) {
	chJobs := make(chan int, len(list))
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

		go func(chResult chan uint32, chJobs chan int, errCh chan error) {
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

// PMapIntUint16Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntUint16Err(f func(int) (uint16, error), list []int, optional ...Optional) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntUint16ErrNoOrder(f, list, worker)
		}
	}

	return pMapIntUint16ErrPreserveOrder(f, list, worker)
}

func pMapIntUint16ErrPreserveOrder(f func(int) (uint16, error), list []int, worker int) ([]uint16, error) {
	chJobs := make(chan map[int]int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint16, chJobs chan map[int]int, errCh chan error) {
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

func pMapIntUint16ErrNoOrder(f func(int) (uint16, error), list []int, worker int) ([]uint16, error) {
	chJobs := make(chan int, len(list))
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

		go func(chResult chan uint16, chJobs chan int, errCh chan error) {
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

// PMapIntUint8Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntUint8Err(f func(int) (uint8, error), list []int, optional ...Optional) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntUint8ErrNoOrder(f, list, worker)
		}
	}

	return pMapIntUint8ErrPreserveOrder(f, list, worker)
}

func pMapIntUint8ErrPreserveOrder(f func(int) (uint8, error), list []int, worker int) ([]uint8, error) {
	chJobs := make(chan map[int]int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint8, chJobs chan map[int]int, errCh chan error) {
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

func pMapIntUint8ErrNoOrder(f func(int) (uint8, error), list []int, worker int) ([]uint8, error) {
	chJobs := make(chan int, len(list))
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

		go func(chResult chan uint8, chJobs chan int, errCh chan error) {
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

// PMapIntStrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntStrErr(f func(int) (string, error), list []int, optional ...Optional) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntStrErrNoOrder(f, list, worker)
		}
	}

	return pMapIntStrErrPreserveOrder(f, list, worker)
}

func pMapIntStrErrPreserveOrder(f func(int) (string, error), list []int, worker int) ([]string, error) {
	chJobs := make(chan map[int]int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]string, chJobs chan map[int]int, errCh chan error) {
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

func pMapIntStrErrNoOrder(f func(int) (string, error), list []int, worker int) ([]string, error) {
	chJobs := make(chan int, len(list))
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

		go func(chResult chan string, chJobs chan int, errCh chan error) {
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

// PMapIntBoolErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntBoolErr(f func(int) (bool, error), list []int, optional ...Optional) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntBoolErrNoOrder(f, list, worker)
		}
	}

	return pMapIntBoolErrPreserveOrder(f, list, worker)
}

func pMapIntBoolErrPreserveOrder(f func(int) (bool, error), list []int, worker int) ([]bool, error) {
	chJobs := make(chan map[int]int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]bool, chJobs chan map[int]int, errCh chan error) {
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

func pMapIntBoolErrNoOrder(f func(int) (bool, error), list []int, worker int) ([]bool, error) {
	chJobs := make(chan int, len(list))
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

		go func(chResult chan bool, chJobs chan int, errCh chan error) {
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

// PMapIntFloat32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntFloat32Err(f func(int) (float32, error), list []int, optional ...Optional) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntFloat32ErrNoOrder(f, list, worker)
		}
	}

	return pMapIntFloat32ErrPreserveOrder(f, list, worker)
}

func pMapIntFloat32ErrPreserveOrder(f func(int) (float32, error), list []int, worker int) ([]float32, error) {
	chJobs := make(chan map[int]int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float32, chJobs chan map[int]int, errCh chan error) {
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

func pMapIntFloat32ErrNoOrder(f func(int) (float32, error), list []int, worker int) ([]float32, error) {
	chJobs := make(chan int, len(list))
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

		go func(chResult chan float32, chJobs chan int, errCh chan error) {
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

// PMapIntFloat64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntFloat64Err(f func(int) (float64, error), list []int, optional ...Optional) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntFloat64ErrNoOrder(f, list, worker)
		}
	}

	return pMapIntFloat64ErrPreserveOrder(f, list, worker)
}

func pMapIntFloat64ErrPreserveOrder(f func(int) (float64, error), list []int, worker int) ([]float64, error) {
	chJobs := make(chan map[int]int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float64, chJobs chan map[int]int, errCh chan error) {
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

func pMapIntFloat64ErrNoOrder(f func(int) (float64, error), list []int, worker int) ([]float64, error) {
	chJobs := make(chan int, len(list))
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

		go func(chResult chan float64, chJobs chan int, errCh chan error) {
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

// PMapInt64IntErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64IntErr(f func(int64) (int, error), list []int64, optional ...Optional) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64IntErrNoOrder(f, list, worker)
		}
	}

	return pMapInt64IntErrPreserveOrder(f, list, worker)
}

func pMapInt64IntErrPreserveOrder(f func(int64) (int, error), list []int64, worker int) ([]int, error) {
	chJobs := make(chan map[int]int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int, chJobs chan map[int]int64, errCh chan error) {
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

func pMapInt64IntErrNoOrder(f func(int64) (int, error), list []int64, worker int) ([]int, error) {
	chJobs := make(chan int64, len(list))
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

		go func(chResult chan int, chJobs chan int64, errCh chan error) {
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

// PMapInt64Int32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Int32Err(f func(int64) (int32, error), list []int64, optional ...Optional) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64Int32ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt64Int32ErrPreserveOrder(f, list, worker)
}

func pMapInt64Int32ErrPreserveOrder(f func(int64) (int32, error), list []int64, worker int) ([]int32, error) {
	chJobs := make(chan map[int]int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int32, chJobs chan map[int]int64, errCh chan error) {
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

func pMapInt64Int32ErrNoOrder(f func(int64) (int32, error), list []int64, worker int) ([]int32, error) {
	chJobs := make(chan int64, len(list))
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

		go func(chResult chan int32, chJobs chan int64, errCh chan error) {
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

// PMapInt64Int16Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Int16Err(f func(int64) (int16, error), list []int64, optional ...Optional) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64Int16ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt64Int16ErrPreserveOrder(f, list, worker)
}

func pMapInt64Int16ErrPreserveOrder(f func(int64) (int16, error), list []int64, worker int) ([]int16, error) {
	chJobs := make(chan map[int]int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int16, chJobs chan map[int]int64, errCh chan error) {
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

func pMapInt64Int16ErrNoOrder(f func(int64) (int16, error), list []int64, worker int) ([]int16, error) {
	chJobs := make(chan int64, len(list))
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

		go func(chResult chan int16, chJobs chan int64, errCh chan error) {
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

// PMapInt64Int8Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Int8Err(f func(int64) (int8, error), list []int64, optional ...Optional) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64Int8ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt64Int8ErrPreserveOrder(f, list, worker)
}

func pMapInt64Int8ErrPreserveOrder(f func(int64) (int8, error), list []int64, worker int) ([]int8, error) {
	chJobs := make(chan map[int]int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int8, chJobs chan map[int]int64, errCh chan error) {
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

func pMapInt64Int8ErrNoOrder(f func(int64) (int8, error), list []int64, worker int) ([]int8, error) {
	chJobs := make(chan int64, len(list))
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

		go func(chResult chan int8, chJobs chan int64, errCh chan error) {
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

// PMapInt64UintErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64UintErr(f func(int64) (uint, error), list []int64, optional ...Optional) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64UintErrNoOrder(f, list, worker)
		}
	}

	return pMapInt64UintErrPreserveOrder(f, list, worker)
}

func pMapInt64UintErrPreserveOrder(f func(int64) (uint, error), list []int64, worker int) ([]uint, error) {
	chJobs := make(chan map[int]int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint, chJobs chan map[int]int64, errCh chan error) {
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

func pMapInt64UintErrNoOrder(f func(int64) (uint, error), list []int64, worker int) ([]uint, error) {
	chJobs := make(chan int64, len(list))
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

		go func(chResult chan uint, chJobs chan int64, errCh chan error) {
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

// PMapInt64Uint64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Uint64Err(f func(int64) (uint64, error), list []int64, optional ...Optional) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64Uint64ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt64Uint64ErrPreserveOrder(f, list, worker)
}

func pMapInt64Uint64ErrPreserveOrder(f func(int64) (uint64, error), list []int64, worker int) ([]uint64, error) {
	chJobs := make(chan map[int]int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint64, chJobs chan map[int]int64, errCh chan error) {
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

func pMapInt64Uint64ErrNoOrder(f func(int64) (uint64, error), list []int64, worker int) ([]uint64, error) {
	chJobs := make(chan int64, len(list))
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

		go func(chResult chan uint64, chJobs chan int64, errCh chan error) {
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

// PMapInt64Uint32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Uint32Err(f func(int64) (uint32, error), list []int64, optional ...Optional) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64Uint32ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt64Uint32ErrPreserveOrder(f, list, worker)
}

func pMapInt64Uint32ErrPreserveOrder(f func(int64) (uint32, error), list []int64, worker int) ([]uint32, error) {
	chJobs := make(chan map[int]int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint32, chJobs chan map[int]int64, errCh chan error) {
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

func pMapInt64Uint32ErrNoOrder(f func(int64) (uint32, error), list []int64, worker int) ([]uint32, error) {
	chJobs := make(chan int64, len(list))
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

		go func(chResult chan uint32, chJobs chan int64, errCh chan error) {
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

// PMapInt64Uint16Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Uint16Err(f func(int64) (uint16, error), list []int64, optional ...Optional) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64Uint16ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt64Uint16ErrPreserveOrder(f, list, worker)
}

func pMapInt64Uint16ErrPreserveOrder(f func(int64) (uint16, error), list []int64, worker int) ([]uint16, error) {
	chJobs := make(chan map[int]int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint16, chJobs chan map[int]int64, errCh chan error) {
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

func pMapInt64Uint16ErrNoOrder(f func(int64) (uint16, error), list []int64, worker int) ([]uint16, error) {
	chJobs := make(chan int64, len(list))
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

		go func(chResult chan uint16, chJobs chan int64, errCh chan error) {
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

// PMapInt64Uint8Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Uint8Err(f func(int64) (uint8, error), list []int64, optional ...Optional) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64Uint8ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt64Uint8ErrPreserveOrder(f, list, worker)
}

func pMapInt64Uint8ErrPreserveOrder(f func(int64) (uint8, error), list []int64, worker int) ([]uint8, error) {
	chJobs := make(chan map[int]int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint8, chJobs chan map[int]int64, errCh chan error) {
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

func pMapInt64Uint8ErrNoOrder(f func(int64) (uint8, error), list []int64, worker int) ([]uint8, error) {
	chJobs := make(chan int64, len(list))
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

		go func(chResult chan uint8, chJobs chan int64, errCh chan error) {
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

// PMapInt64StrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64StrErr(f func(int64) (string, error), list []int64, optional ...Optional) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64StrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt64StrErrPreserveOrder(f, list, worker)
}

func pMapInt64StrErrPreserveOrder(f func(int64) (string, error), list []int64, worker int) ([]string, error) {
	chJobs := make(chan map[int]int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]string, chJobs chan map[int]int64, errCh chan error) {
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

func pMapInt64StrErrNoOrder(f func(int64) (string, error), list []int64, worker int) ([]string, error) {
	chJobs := make(chan int64, len(list))
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

		go func(chResult chan string, chJobs chan int64, errCh chan error) {
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

// PMapInt64BoolErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64BoolErr(f func(int64) (bool, error), list []int64, optional ...Optional) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64BoolErrNoOrder(f, list, worker)
		}
	}

	return pMapInt64BoolErrPreserveOrder(f, list, worker)
}

func pMapInt64BoolErrPreserveOrder(f func(int64) (bool, error), list []int64, worker int) ([]bool, error) {
	chJobs := make(chan map[int]int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]bool, chJobs chan map[int]int64, errCh chan error) {
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

func pMapInt64BoolErrNoOrder(f func(int64) (bool, error), list []int64, worker int) ([]bool, error) {
	chJobs := make(chan int64, len(list))
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

		go func(chResult chan bool, chJobs chan int64, errCh chan error) {
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

// PMapInt64Float32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Float32Err(f func(int64) (float32, error), list []int64, optional ...Optional) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64Float32ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt64Float32ErrPreserveOrder(f, list, worker)
}

func pMapInt64Float32ErrPreserveOrder(f func(int64) (float32, error), list []int64, worker int) ([]float32, error) {
	chJobs := make(chan map[int]int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float32, chJobs chan map[int]int64, errCh chan error) {
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

func pMapInt64Float32ErrNoOrder(f func(int64) (float32, error), list []int64, worker int) ([]float32, error) {
	chJobs := make(chan int64, len(list))
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

		go func(chResult chan float32, chJobs chan int64, errCh chan error) {
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

// PMapInt64Float64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Float64Err(f func(int64) (float64, error), list []int64, optional ...Optional) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64Float64ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt64Float64ErrPreserveOrder(f, list, worker)
}

func pMapInt64Float64ErrPreserveOrder(f func(int64) (float64, error), list []int64, worker int) ([]float64, error) {
	chJobs := make(chan map[int]int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float64, chJobs chan map[int]int64, errCh chan error) {
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

func pMapInt64Float64ErrNoOrder(f func(int64) (float64, error), list []int64, worker int) ([]float64, error) {
	chJobs := make(chan int64, len(list))
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

		go func(chResult chan float64, chJobs chan int64, errCh chan error) {
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

// PMapInt32IntErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32IntErr(f func(int32) (int, error), list []int32, optional ...Optional) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32IntErrNoOrder(f, list, worker)
		}
	}

	return pMapInt32IntErrPreserveOrder(f, list, worker)
}

func pMapInt32IntErrPreserveOrder(f func(int32) (int, error), list []int32, worker int) ([]int, error) {
	chJobs := make(chan map[int]int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int, chJobs chan map[int]int32, errCh chan error) {
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

func pMapInt32IntErrNoOrder(f func(int32) (int, error), list []int32, worker int) ([]int, error) {
	chJobs := make(chan int32, len(list))
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

		go func(chResult chan int, chJobs chan int32, errCh chan error) {
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

// PMapInt32Int64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Int64Err(f func(int32) (int64, error), list []int32, optional ...Optional) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32Int64ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt32Int64ErrPreserveOrder(f, list, worker)
}

func pMapInt32Int64ErrPreserveOrder(f func(int32) (int64, error), list []int32, worker int) ([]int64, error) {
	chJobs := make(chan map[int]int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int64, chJobs chan map[int]int32, errCh chan error) {
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

func pMapInt32Int64ErrNoOrder(f func(int32) (int64, error), list []int32, worker int) ([]int64, error) {
	chJobs := make(chan int32, len(list))
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

		go func(chResult chan int64, chJobs chan int32, errCh chan error) {
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

// PMapInt32Int16Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Int16Err(f func(int32) (int16, error), list []int32, optional ...Optional) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32Int16ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt32Int16ErrPreserveOrder(f, list, worker)
}

func pMapInt32Int16ErrPreserveOrder(f func(int32) (int16, error), list []int32, worker int) ([]int16, error) {
	chJobs := make(chan map[int]int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int16, chJobs chan map[int]int32, errCh chan error) {
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

func pMapInt32Int16ErrNoOrder(f func(int32) (int16, error), list []int32, worker int) ([]int16, error) {
	chJobs := make(chan int32, len(list))
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

		go func(chResult chan int16, chJobs chan int32, errCh chan error) {
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

// PMapInt32Int8Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Int8Err(f func(int32) (int8, error), list []int32, optional ...Optional) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32Int8ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt32Int8ErrPreserveOrder(f, list, worker)
}

func pMapInt32Int8ErrPreserveOrder(f func(int32) (int8, error), list []int32, worker int) ([]int8, error) {
	chJobs := make(chan map[int]int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int8, chJobs chan map[int]int32, errCh chan error) {
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

func pMapInt32Int8ErrNoOrder(f func(int32) (int8, error), list []int32, worker int) ([]int8, error) {
	chJobs := make(chan int32, len(list))
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

		go func(chResult chan int8, chJobs chan int32, errCh chan error) {
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

// PMapInt32UintErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32UintErr(f func(int32) (uint, error), list []int32, optional ...Optional) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32UintErrNoOrder(f, list, worker)
		}
	}

	return pMapInt32UintErrPreserveOrder(f, list, worker)
}

func pMapInt32UintErrPreserveOrder(f func(int32) (uint, error), list []int32, worker int) ([]uint, error) {
	chJobs := make(chan map[int]int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint, chJobs chan map[int]int32, errCh chan error) {
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

func pMapInt32UintErrNoOrder(f func(int32) (uint, error), list []int32, worker int) ([]uint, error) {
	chJobs := make(chan int32, len(list))
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

		go func(chResult chan uint, chJobs chan int32, errCh chan error) {
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

// PMapInt32Uint64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Uint64Err(f func(int32) (uint64, error), list []int32, optional ...Optional) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32Uint64ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt32Uint64ErrPreserveOrder(f, list, worker)
}

func pMapInt32Uint64ErrPreserveOrder(f func(int32) (uint64, error), list []int32, worker int) ([]uint64, error) {
	chJobs := make(chan map[int]int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint64, chJobs chan map[int]int32, errCh chan error) {
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

func pMapInt32Uint64ErrNoOrder(f func(int32) (uint64, error), list []int32, worker int) ([]uint64, error) {
	chJobs := make(chan int32, len(list))
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

		go func(chResult chan uint64, chJobs chan int32, errCh chan error) {
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

// PMapInt32Uint32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Uint32Err(f func(int32) (uint32, error), list []int32, optional ...Optional) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32Uint32ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt32Uint32ErrPreserveOrder(f, list, worker)
}

func pMapInt32Uint32ErrPreserveOrder(f func(int32) (uint32, error), list []int32, worker int) ([]uint32, error) {
	chJobs := make(chan map[int]int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint32, chJobs chan map[int]int32, errCh chan error) {
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

func pMapInt32Uint32ErrNoOrder(f func(int32) (uint32, error), list []int32, worker int) ([]uint32, error) {
	chJobs := make(chan int32, len(list))
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

		go func(chResult chan uint32, chJobs chan int32, errCh chan error) {
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

// PMapInt32Uint16Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Uint16Err(f func(int32) (uint16, error), list []int32, optional ...Optional) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32Uint16ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt32Uint16ErrPreserveOrder(f, list, worker)
}

func pMapInt32Uint16ErrPreserveOrder(f func(int32) (uint16, error), list []int32, worker int) ([]uint16, error) {
	chJobs := make(chan map[int]int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint16, chJobs chan map[int]int32, errCh chan error) {
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

func pMapInt32Uint16ErrNoOrder(f func(int32) (uint16, error), list []int32, worker int) ([]uint16, error) {
	chJobs := make(chan int32, len(list))
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

		go func(chResult chan uint16, chJobs chan int32, errCh chan error) {
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

// PMapInt32Uint8Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Uint8Err(f func(int32) (uint8, error), list []int32, optional ...Optional) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32Uint8ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt32Uint8ErrPreserveOrder(f, list, worker)
}

func pMapInt32Uint8ErrPreserveOrder(f func(int32) (uint8, error), list []int32, worker int) ([]uint8, error) {
	chJobs := make(chan map[int]int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint8, chJobs chan map[int]int32, errCh chan error) {
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

func pMapInt32Uint8ErrNoOrder(f func(int32) (uint8, error), list []int32, worker int) ([]uint8, error) {
	chJobs := make(chan int32, len(list))
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

		go func(chResult chan uint8, chJobs chan int32, errCh chan error) {
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

// PMapInt32StrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32StrErr(f func(int32) (string, error), list []int32, optional ...Optional) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32StrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt32StrErrPreserveOrder(f, list, worker)
}

func pMapInt32StrErrPreserveOrder(f func(int32) (string, error), list []int32, worker int) ([]string, error) {
	chJobs := make(chan map[int]int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]string, chJobs chan map[int]int32, errCh chan error) {
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

func pMapInt32StrErrNoOrder(f func(int32) (string, error), list []int32, worker int) ([]string, error) {
	chJobs := make(chan int32, len(list))
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

		go func(chResult chan string, chJobs chan int32, errCh chan error) {
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

// PMapInt32BoolErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32BoolErr(f func(int32) (bool, error), list []int32, optional ...Optional) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32BoolErrNoOrder(f, list, worker)
		}
	}

	return pMapInt32BoolErrPreserveOrder(f, list, worker)
}

func pMapInt32BoolErrPreserveOrder(f func(int32) (bool, error), list []int32, worker int) ([]bool, error) {
	chJobs := make(chan map[int]int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]bool, chJobs chan map[int]int32, errCh chan error) {
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

func pMapInt32BoolErrNoOrder(f func(int32) (bool, error), list []int32, worker int) ([]bool, error) {
	chJobs := make(chan int32, len(list))
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

		go func(chResult chan bool, chJobs chan int32, errCh chan error) {
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

// PMapInt32Float32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Float32Err(f func(int32) (float32, error), list []int32, optional ...Optional) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32Float32ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt32Float32ErrPreserveOrder(f, list, worker)
}

func pMapInt32Float32ErrPreserveOrder(f func(int32) (float32, error), list []int32, worker int) ([]float32, error) {
	chJobs := make(chan map[int]int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float32, chJobs chan map[int]int32, errCh chan error) {
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

func pMapInt32Float32ErrNoOrder(f func(int32) (float32, error), list []int32, worker int) ([]float32, error) {
	chJobs := make(chan int32, len(list))
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

		go func(chResult chan float32, chJobs chan int32, errCh chan error) {
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

// PMapInt32Float64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Float64Err(f func(int32) (float64, error), list []int32, optional ...Optional) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32Float64ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt32Float64ErrPreserveOrder(f, list, worker)
}

func pMapInt32Float64ErrPreserveOrder(f func(int32) (float64, error), list []int32, worker int) ([]float64, error) {
	chJobs := make(chan map[int]int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float64, chJobs chan map[int]int32, errCh chan error) {
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

func pMapInt32Float64ErrNoOrder(f func(int32) (float64, error), list []int32, worker int) ([]float64, error) {
	chJobs := make(chan int32, len(list))
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

		go func(chResult chan float64, chJobs chan int32, errCh chan error) {
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

// PMapInt16IntErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16IntErr(f func(int16) (int, error), list []int16, optional ...Optional) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16IntErrNoOrder(f, list, worker)
		}
	}

	return pMapInt16IntErrPreserveOrder(f, list, worker)
}

func pMapInt16IntErrPreserveOrder(f func(int16) (int, error), list []int16, worker int) ([]int, error) {
	chJobs := make(chan map[int]int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int, chJobs chan map[int]int16, errCh chan error) {
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

func pMapInt16IntErrNoOrder(f func(int16) (int, error), list []int16, worker int) ([]int, error) {
	chJobs := make(chan int16, len(list))
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

		go func(chResult chan int, chJobs chan int16, errCh chan error) {
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

// PMapInt16Int64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Int64Err(f func(int16) (int64, error), list []int16, optional ...Optional) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16Int64ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt16Int64ErrPreserveOrder(f, list, worker)
}

func pMapInt16Int64ErrPreserveOrder(f func(int16) (int64, error), list []int16, worker int) ([]int64, error) {
	chJobs := make(chan map[int]int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int64, chJobs chan map[int]int16, errCh chan error) {
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

func pMapInt16Int64ErrNoOrder(f func(int16) (int64, error), list []int16, worker int) ([]int64, error) {
	chJobs := make(chan int16, len(list))
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

		go func(chResult chan int64, chJobs chan int16, errCh chan error) {
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

// PMapInt16Int32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Int32Err(f func(int16) (int32, error), list []int16, optional ...Optional) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16Int32ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt16Int32ErrPreserveOrder(f, list, worker)
}

func pMapInt16Int32ErrPreserveOrder(f func(int16) (int32, error), list []int16, worker int) ([]int32, error) {
	chJobs := make(chan map[int]int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int32, chJobs chan map[int]int16, errCh chan error) {
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

func pMapInt16Int32ErrNoOrder(f func(int16) (int32, error), list []int16, worker int) ([]int32, error) {
	chJobs := make(chan int16, len(list))
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

		go func(chResult chan int32, chJobs chan int16, errCh chan error) {
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

// PMapInt16Int8Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Int8Err(f func(int16) (int8, error), list []int16, optional ...Optional) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16Int8ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt16Int8ErrPreserveOrder(f, list, worker)
}

func pMapInt16Int8ErrPreserveOrder(f func(int16) (int8, error), list []int16, worker int) ([]int8, error) {
	chJobs := make(chan map[int]int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int8, chJobs chan map[int]int16, errCh chan error) {
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

func pMapInt16Int8ErrNoOrder(f func(int16) (int8, error), list []int16, worker int) ([]int8, error) {
	chJobs := make(chan int16, len(list))
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

		go func(chResult chan int8, chJobs chan int16, errCh chan error) {
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

// PMapInt16UintErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16UintErr(f func(int16) (uint, error), list []int16, optional ...Optional) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16UintErrNoOrder(f, list, worker)
		}
	}

	return pMapInt16UintErrPreserveOrder(f, list, worker)
}

func pMapInt16UintErrPreserveOrder(f func(int16) (uint, error), list []int16, worker int) ([]uint, error) {
	chJobs := make(chan map[int]int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint, chJobs chan map[int]int16, errCh chan error) {
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

func pMapInt16UintErrNoOrder(f func(int16) (uint, error), list []int16, worker int) ([]uint, error) {
	chJobs := make(chan int16, len(list))
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

		go func(chResult chan uint, chJobs chan int16, errCh chan error) {
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

// PMapInt16Uint64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Uint64Err(f func(int16) (uint64, error), list []int16, optional ...Optional) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16Uint64ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt16Uint64ErrPreserveOrder(f, list, worker)
}

func pMapInt16Uint64ErrPreserveOrder(f func(int16) (uint64, error), list []int16, worker int) ([]uint64, error) {
	chJobs := make(chan map[int]int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint64, chJobs chan map[int]int16, errCh chan error) {
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

func pMapInt16Uint64ErrNoOrder(f func(int16) (uint64, error), list []int16, worker int) ([]uint64, error) {
	chJobs := make(chan int16, len(list))
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

		go func(chResult chan uint64, chJobs chan int16, errCh chan error) {
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

// PMapInt16Uint32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Uint32Err(f func(int16) (uint32, error), list []int16, optional ...Optional) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16Uint32ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt16Uint32ErrPreserveOrder(f, list, worker)
}

func pMapInt16Uint32ErrPreserveOrder(f func(int16) (uint32, error), list []int16, worker int) ([]uint32, error) {
	chJobs := make(chan map[int]int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint32, chJobs chan map[int]int16, errCh chan error) {
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

func pMapInt16Uint32ErrNoOrder(f func(int16) (uint32, error), list []int16, worker int) ([]uint32, error) {
	chJobs := make(chan int16, len(list))
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

		go func(chResult chan uint32, chJobs chan int16, errCh chan error) {
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

// PMapInt16Uint16Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Uint16Err(f func(int16) (uint16, error), list []int16, optional ...Optional) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16Uint16ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt16Uint16ErrPreserveOrder(f, list, worker)
}

func pMapInt16Uint16ErrPreserveOrder(f func(int16) (uint16, error), list []int16, worker int) ([]uint16, error) {
	chJobs := make(chan map[int]int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint16, chJobs chan map[int]int16, errCh chan error) {
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

func pMapInt16Uint16ErrNoOrder(f func(int16) (uint16, error), list []int16, worker int) ([]uint16, error) {
	chJobs := make(chan int16, len(list))
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

		go func(chResult chan uint16, chJobs chan int16, errCh chan error) {
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

// PMapInt16Uint8Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Uint8Err(f func(int16) (uint8, error), list []int16, optional ...Optional) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16Uint8ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt16Uint8ErrPreserveOrder(f, list, worker)
}

func pMapInt16Uint8ErrPreserveOrder(f func(int16) (uint8, error), list []int16, worker int) ([]uint8, error) {
	chJobs := make(chan map[int]int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint8, chJobs chan map[int]int16, errCh chan error) {
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

func pMapInt16Uint8ErrNoOrder(f func(int16) (uint8, error), list []int16, worker int) ([]uint8, error) {
	chJobs := make(chan int16, len(list))
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

		go func(chResult chan uint8, chJobs chan int16, errCh chan error) {
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

// PMapInt16StrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16StrErr(f func(int16) (string, error), list []int16, optional ...Optional) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16StrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt16StrErrPreserveOrder(f, list, worker)
}

func pMapInt16StrErrPreserveOrder(f func(int16) (string, error), list []int16, worker int) ([]string, error) {
	chJobs := make(chan map[int]int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]string, chJobs chan map[int]int16, errCh chan error) {
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

func pMapInt16StrErrNoOrder(f func(int16) (string, error), list []int16, worker int) ([]string, error) {
	chJobs := make(chan int16, len(list))
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

		go func(chResult chan string, chJobs chan int16, errCh chan error) {
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

// PMapInt16BoolErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16BoolErr(f func(int16) (bool, error), list []int16, optional ...Optional) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16BoolErrNoOrder(f, list, worker)
		}
	}

	return pMapInt16BoolErrPreserveOrder(f, list, worker)
}

func pMapInt16BoolErrPreserveOrder(f func(int16) (bool, error), list []int16, worker int) ([]bool, error) {
	chJobs := make(chan map[int]int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]bool, chJobs chan map[int]int16, errCh chan error) {
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

func pMapInt16BoolErrNoOrder(f func(int16) (bool, error), list []int16, worker int) ([]bool, error) {
	chJobs := make(chan int16, len(list))
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

		go func(chResult chan bool, chJobs chan int16, errCh chan error) {
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

// PMapInt16Float32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Float32Err(f func(int16) (float32, error), list []int16, optional ...Optional) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16Float32ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt16Float32ErrPreserveOrder(f, list, worker)
}

func pMapInt16Float32ErrPreserveOrder(f func(int16) (float32, error), list []int16, worker int) ([]float32, error) {
	chJobs := make(chan map[int]int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float32, chJobs chan map[int]int16, errCh chan error) {
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

func pMapInt16Float32ErrNoOrder(f func(int16) (float32, error), list []int16, worker int) ([]float32, error) {
	chJobs := make(chan int16, len(list))
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

		go func(chResult chan float32, chJobs chan int16, errCh chan error) {
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

// PMapInt16Float64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Float64Err(f func(int16) (float64, error), list []int16, optional ...Optional) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16Float64ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt16Float64ErrPreserveOrder(f, list, worker)
}

func pMapInt16Float64ErrPreserveOrder(f func(int16) (float64, error), list []int16, worker int) ([]float64, error) {
	chJobs := make(chan map[int]int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float64, chJobs chan map[int]int16, errCh chan error) {
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

func pMapInt16Float64ErrNoOrder(f func(int16) (float64, error), list []int16, worker int) ([]float64, error) {
	chJobs := make(chan int16, len(list))
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

		go func(chResult chan float64, chJobs chan int16, errCh chan error) {
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

// PMapInt8IntErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8IntErr(f func(int8) (int, error), list []int8, optional ...Optional) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8IntErrNoOrder(f, list, worker)
		}
	}

	return pMapInt8IntErrPreserveOrder(f, list, worker)
}

func pMapInt8IntErrPreserveOrder(f func(int8) (int, error), list []int8, worker int) ([]int, error) {
	chJobs := make(chan map[int]int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int, chJobs chan map[int]int8, errCh chan error) {
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

func pMapInt8IntErrNoOrder(f func(int8) (int, error), list []int8, worker int) ([]int, error) {
	chJobs := make(chan int8, len(list))
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

		go func(chResult chan int, chJobs chan int8, errCh chan error) {
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

// PMapInt8Int64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Int64Err(f func(int8) (int64, error), list []int8, optional ...Optional) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8Int64ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt8Int64ErrPreserveOrder(f, list, worker)
}

func pMapInt8Int64ErrPreserveOrder(f func(int8) (int64, error), list []int8, worker int) ([]int64, error) {
	chJobs := make(chan map[int]int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int64, chJobs chan map[int]int8, errCh chan error) {
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

func pMapInt8Int64ErrNoOrder(f func(int8) (int64, error), list []int8, worker int) ([]int64, error) {
	chJobs := make(chan int8, len(list))
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

		go func(chResult chan int64, chJobs chan int8, errCh chan error) {
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

// PMapInt8Int32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Int32Err(f func(int8) (int32, error), list []int8, optional ...Optional) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8Int32ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt8Int32ErrPreserveOrder(f, list, worker)
}

func pMapInt8Int32ErrPreserveOrder(f func(int8) (int32, error), list []int8, worker int) ([]int32, error) {
	chJobs := make(chan map[int]int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int32, chJobs chan map[int]int8, errCh chan error) {
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

func pMapInt8Int32ErrNoOrder(f func(int8) (int32, error), list []int8, worker int) ([]int32, error) {
	chJobs := make(chan int8, len(list))
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

		go func(chResult chan int32, chJobs chan int8, errCh chan error) {
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

// PMapInt8Int16Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Int16Err(f func(int8) (int16, error), list []int8, optional ...Optional) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8Int16ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt8Int16ErrPreserveOrder(f, list, worker)
}

func pMapInt8Int16ErrPreserveOrder(f func(int8) (int16, error), list []int8, worker int) ([]int16, error) {
	chJobs := make(chan map[int]int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int16, chJobs chan map[int]int8, errCh chan error) {
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

func pMapInt8Int16ErrNoOrder(f func(int8) (int16, error), list []int8, worker int) ([]int16, error) {
	chJobs := make(chan int8, len(list))
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

		go func(chResult chan int16, chJobs chan int8, errCh chan error) {
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

// PMapInt8UintErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8UintErr(f func(int8) (uint, error), list []int8, optional ...Optional) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8UintErrNoOrder(f, list, worker)
		}
	}

	return pMapInt8UintErrPreserveOrder(f, list, worker)
}

func pMapInt8UintErrPreserveOrder(f func(int8) (uint, error), list []int8, worker int) ([]uint, error) {
	chJobs := make(chan map[int]int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint, chJobs chan map[int]int8, errCh chan error) {
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

func pMapInt8UintErrNoOrder(f func(int8) (uint, error), list []int8, worker int) ([]uint, error) {
	chJobs := make(chan int8, len(list))
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

		go func(chResult chan uint, chJobs chan int8, errCh chan error) {
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

// PMapInt8Uint64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Uint64Err(f func(int8) (uint64, error), list []int8, optional ...Optional) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8Uint64ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt8Uint64ErrPreserveOrder(f, list, worker)
}

func pMapInt8Uint64ErrPreserveOrder(f func(int8) (uint64, error), list []int8, worker int) ([]uint64, error) {
	chJobs := make(chan map[int]int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint64, chJobs chan map[int]int8, errCh chan error) {
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

func pMapInt8Uint64ErrNoOrder(f func(int8) (uint64, error), list []int8, worker int) ([]uint64, error) {
	chJobs := make(chan int8, len(list))
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

		go func(chResult chan uint64, chJobs chan int8, errCh chan error) {
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

// PMapInt8Uint32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Uint32Err(f func(int8) (uint32, error), list []int8, optional ...Optional) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8Uint32ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt8Uint32ErrPreserveOrder(f, list, worker)
}

func pMapInt8Uint32ErrPreserveOrder(f func(int8) (uint32, error), list []int8, worker int) ([]uint32, error) {
	chJobs := make(chan map[int]int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint32, chJobs chan map[int]int8, errCh chan error) {
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

func pMapInt8Uint32ErrNoOrder(f func(int8) (uint32, error), list []int8, worker int) ([]uint32, error) {
	chJobs := make(chan int8, len(list))
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

		go func(chResult chan uint32, chJobs chan int8, errCh chan error) {
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

// PMapInt8Uint16Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Uint16Err(f func(int8) (uint16, error), list []int8, optional ...Optional) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8Uint16ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt8Uint16ErrPreserveOrder(f, list, worker)
}

func pMapInt8Uint16ErrPreserveOrder(f func(int8) (uint16, error), list []int8, worker int) ([]uint16, error) {
	chJobs := make(chan map[int]int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint16, chJobs chan map[int]int8, errCh chan error) {
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

func pMapInt8Uint16ErrNoOrder(f func(int8) (uint16, error), list []int8, worker int) ([]uint16, error) {
	chJobs := make(chan int8, len(list))
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

		go func(chResult chan uint16, chJobs chan int8, errCh chan error) {
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

// PMapInt8Uint8Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Uint8Err(f func(int8) (uint8, error), list []int8, optional ...Optional) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8Uint8ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt8Uint8ErrPreserveOrder(f, list, worker)
}

func pMapInt8Uint8ErrPreserveOrder(f func(int8) (uint8, error), list []int8, worker int) ([]uint8, error) {
	chJobs := make(chan map[int]int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint8, chJobs chan map[int]int8, errCh chan error) {
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

func pMapInt8Uint8ErrNoOrder(f func(int8) (uint8, error), list []int8, worker int) ([]uint8, error) {
	chJobs := make(chan int8, len(list))
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

		go func(chResult chan uint8, chJobs chan int8, errCh chan error) {
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

// PMapInt8StrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8StrErr(f func(int8) (string, error), list []int8, optional ...Optional) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8StrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt8StrErrPreserveOrder(f, list, worker)
}

func pMapInt8StrErrPreserveOrder(f func(int8) (string, error), list []int8, worker int) ([]string, error) {
	chJobs := make(chan map[int]int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]string, chJobs chan map[int]int8, errCh chan error) {
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

func pMapInt8StrErrNoOrder(f func(int8) (string, error), list []int8, worker int) ([]string, error) {
	chJobs := make(chan int8, len(list))
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

		go func(chResult chan string, chJobs chan int8, errCh chan error) {
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

// PMapInt8BoolErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8BoolErr(f func(int8) (bool, error), list []int8, optional ...Optional) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8BoolErrNoOrder(f, list, worker)
		}
	}

	return pMapInt8BoolErrPreserveOrder(f, list, worker)
}

func pMapInt8BoolErrPreserveOrder(f func(int8) (bool, error), list []int8, worker int) ([]bool, error) {
	chJobs := make(chan map[int]int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]bool, chJobs chan map[int]int8, errCh chan error) {
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

func pMapInt8BoolErrNoOrder(f func(int8) (bool, error), list []int8, worker int) ([]bool, error) {
	chJobs := make(chan int8, len(list))
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

		go func(chResult chan bool, chJobs chan int8, errCh chan error) {
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

// PMapInt8Float32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Float32Err(f func(int8) (float32, error), list []int8, optional ...Optional) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8Float32ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt8Float32ErrPreserveOrder(f, list, worker)
}

func pMapInt8Float32ErrPreserveOrder(f func(int8) (float32, error), list []int8, worker int) ([]float32, error) {
	chJobs := make(chan map[int]int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float32, chJobs chan map[int]int8, errCh chan error) {
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

func pMapInt8Float32ErrNoOrder(f func(int8) (float32, error), list []int8, worker int) ([]float32, error) {
	chJobs := make(chan int8, len(list))
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

		go func(chResult chan float32, chJobs chan int8, errCh chan error) {
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

// PMapInt8Float64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Float64Err(f func(int8) (float64, error), list []int8, optional ...Optional) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8Float64ErrNoOrder(f, list, worker)
		}
	}

	return pMapInt8Float64ErrPreserveOrder(f, list, worker)
}

func pMapInt8Float64ErrPreserveOrder(f func(int8) (float64, error), list []int8, worker int) ([]float64, error) {
	chJobs := make(chan map[int]int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float64, chJobs chan map[int]int8, errCh chan error) {
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

func pMapInt8Float64ErrNoOrder(f func(int8) (float64, error), list []int8, worker int) ([]float64, error) {
	chJobs := make(chan int8, len(list))
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

		go func(chResult chan float64, chJobs chan int8, errCh chan error) {
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

// PMapUintIntErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintIntErr(f func(uint) (int, error), list []uint, optional ...Optional) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintIntErrNoOrder(f, list, worker)
		}
	}

	return pMapUintIntErrPreserveOrder(f, list, worker)
}

func pMapUintIntErrPreserveOrder(f func(uint) (int, error), list []uint, worker int) ([]int, error) {
	chJobs := make(chan map[int]uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int, chJobs chan map[int]uint, errCh chan error) {
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

func pMapUintIntErrNoOrder(f func(uint) (int, error), list []uint, worker int) ([]int, error) {
	chJobs := make(chan uint, len(list))
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

		go func(chResult chan int, chJobs chan uint, errCh chan error) {
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

// PMapUintInt64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintInt64Err(f func(uint) (int64, error), list []uint, optional ...Optional) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintInt64ErrNoOrder(f, list, worker)
		}
	}

	return pMapUintInt64ErrPreserveOrder(f, list, worker)
}

func pMapUintInt64ErrPreserveOrder(f func(uint) (int64, error), list []uint, worker int) ([]int64, error) {
	chJobs := make(chan map[int]uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int64, chJobs chan map[int]uint, errCh chan error) {
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

func pMapUintInt64ErrNoOrder(f func(uint) (int64, error), list []uint, worker int) ([]int64, error) {
	chJobs := make(chan uint, len(list))
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

		go func(chResult chan int64, chJobs chan uint, errCh chan error) {
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

// PMapUintInt32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintInt32Err(f func(uint) (int32, error), list []uint, optional ...Optional) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintInt32ErrNoOrder(f, list, worker)
		}
	}

	return pMapUintInt32ErrPreserveOrder(f, list, worker)
}

func pMapUintInt32ErrPreserveOrder(f func(uint) (int32, error), list []uint, worker int) ([]int32, error) {
	chJobs := make(chan map[int]uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int32, chJobs chan map[int]uint, errCh chan error) {
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

func pMapUintInt32ErrNoOrder(f func(uint) (int32, error), list []uint, worker int) ([]int32, error) {
	chJobs := make(chan uint, len(list))
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

		go func(chResult chan int32, chJobs chan uint, errCh chan error) {
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

// PMapUintInt16Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintInt16Err(f func(uint) (int16, error), list []uint, optional ...Optional) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintInt16ErrNoOrder(f, list, worker)
		}
	}

	return pMapUintInt16ErrPreserveOrder(f, list, worker)
}

func pMapUintInt16ErrPreserveOrder(f func(uint) (int16, error), list []uint, worker int) ([]int16, error) {
	chJobs := make(chan map[int]uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int16, chJobs chan map[int]uint, errCh chan error) {
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

func pMapUintInt16ErrNoOrder(f func(uint) (int16, error), list []uint, worker int) ([]int16, error) {
	chJobs := make(chan uint, len(list))
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

		go func(chResult chan int16, chJobs chan uint, errCh chan error) {
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

// PMapUintInt8Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintInt8Err(f func(uint) (int8, error), list []uint, optional ...Optional) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintInt8ErrNoOrder(f, list, worker)
		}
	}

	return pMapUintInt8ErrPreserveOrder(f, list, worker)
}

func pMapUintInt8ErrPreserveOrder(f func(uint) (int8, error), list []uint, worker int) ([]int8, error) {
	chJobs := make(chan map[int]uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int8, chJobs chan map[int]uint, errCh chan error) {
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

func pMapUintInt8ErrNoOrder(f func(uint) (int8, error), list []uint, worker int) ([]int8, error) {
	chJobs := make(chan uint, len(list))
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

		go func(chResult chan int8, chJobs chan uint, errCh chan error) {
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

// PMapUintUint64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintUint64Err(f func(uint) (uint64, error), list []uint, optional ...Optional) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintUint64ErrNoOrder(f, list, worker)
		}
	}

	return pMapUintUint64ErrPreserveOrder(f, list, worker)
}

func pMapUintUint64ErrPreserveOrder(f func(uint) (uint64, error), list []uint, worker int) ([]uint64, error) {
	chJobs := make(chan map[int]uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint64, chJobs chan map[int]uint, errCh chan error) {
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

func pMapUintUint64ErrNoOrder(f func(uint) (uint64, error), list []uint, worker int) ([]uint64, error) {
	chJobs := make(chan uint, len(list))
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

		go func(chResult chan uint64, chJobs chan uint, errCh chan error) {
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

// PMapUintUint32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintUint32Err(f func(uint) (uint32, error), list []uint, optional ...Optional) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintUint32ErrNoOrder(f, list, worker)
		}
	}

	return pMapUintUint32ErrPreserveOrder(f, list, worker)
}

func pMapUintUint32ErrPreserveOrder(f func(uint) (uint32, error), list []uint, worker int) ([]uint32, error) {
	chJobs := make(chan map[int]uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint32, chJobs chan map[int]uint, errCh chan error) {
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

func pMapUintUint32ErrNoOrder(f func(uint) (uint32, error), list []uint, worker int) ([]uint32, error) {
	chJobs := make(chan uint, len(list))
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

		go func(chResult chan uint32, chJobs chan uint, errCh chan error) {
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

// PMapUintUint16Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintUint16Err(f func(uint) (uint16, error), list []uint, optional ...Optional) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintUint16ErrNoOrder(f, list, worker)
		}
	}

	return pMapUintUint16ErrPreserveOrder(f, list, worker)
}

func pMapUintUint16ErrPreserveOrder(f func(uint) (uint16, error), list []uint, worker int) ([]uint16, error) {
	chJobs := make(chan map[int]uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint16, chJobs chan map[int]uint, errCh chan error) {
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

func pMapUintUint16ErrNoOrder(f func(uint) (uint16, error), list []uint, worker int) ([]uint16, error) {
	chJobs := make(chan uint, len(list))
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

		go func(chResult chan uint16, chJobs chan uint, errCh chan error) {
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

// PMapUintUint8Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintUint8Err(f func(uint) (uint8, error), list []uint, optional ...Optional) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintUint8ErrNoOrder(f, list, worker)
		}
	}

	return pMapUintUint8ErrPreserveOrder(f, list, worker)
}

func pMapUintUint8ErrPreserveOrder(f func(uint) (uint8, error), list []uint, worker int) ([]uint8, error) {
	chJobs := make(chan map[int]uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint8, chJobs chan map[int]uint, errCh chan error) {
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

func pMapUintUint8ErrNoOrder(f func(uint) (uint8, error), list []uint, worker int) ([]uint8, error) {
	chJobs := make(chan uint, len(list))
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

		go func(chResult chan uint8, chJobs chan uint, errCh chan error) {
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

// PMapUintStrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintStrErr(f func(uint) (string, error), list []uint, optional ...Optional) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintStrErrNoOrder(f, list, worker)
		}
	}

	return pMapUintStrErrPreserveOrder(f, list, worker)
}

func pMapUintStrErrPreserveOrder(f func(uint) (string, error), list []uint, worker int) ([]string, error) {
	chJobs := make(chan map[int]uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]string, chJobs chan map[int]uint, errCh chan error) {
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

func pMapUintStrErrNoOrder(f func(uint) (string, error), list []uint, worker int) ([]string, error) {
	chJobs := make(chan uint, len(list))
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

		go func(chResult chan string, chJobs chan uint, errCh chan error) {
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

// PMapUintBoolErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintBoolErr(f func(uint) (bool, error), list []uint, optional ...Optional) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintBoolErrNoOrder(f, list, worker)
		}
	}

	return pMapUintBoolErrPreserveOrder(f, list, worker)
}

func pMapUintBoolErrPreserveOrder(f func(uint) (bool, error), list []uint, worker int) ([]bool, error) {
	chJobs := make(chan map[int]uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]bool, chJobs chan map[int]uint, errCh chan error) {
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

func pMapUintBoolErrNoOrder(f func(uint) (bool, error), list []uint, worker int) ([]bool, error) {
	chJobs := make(chan uint, len(list))
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

		go func(chResult chan bool, chJobs chan uint, errCh chan error) {
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

// PMapUintFloat32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintFloat32Err(f func(uint) (float32, error), list []uint, optional ...Optional) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintFloat32ErrNoOrder(f, list, worker)
		}
	}

	return pMapUintFloat32ErrPreserveOrder(f, list, worker)
}

func pMapUintFloat32ErrPreserveOrder(f func(uint) (float32, error), list []uint, worker int) ([]float32, error) {
	chJobs := make(chan map[int]uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float32, chJobs chan map[int]uint, errCh chan error) {
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

func pMapUintFloat32ErrNoOrder(f func(uint) (float32, error), list []uint, worker int) ([]float32, error) {
	chJobs := make(chan uint, len(list))
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

		go func(chResult chan float32, chJobs chan uint, errCh chan error) {
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

// PMapUintFloat64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintFloat64Err(f func(uint) (float64, error), list []uint, optional ...Optional) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintFloat64ErrNoOrder(f, list, worker)
		}
	}

	return pMapUintFloat64ErrPreserveOrder(f, list, worker)
}

func pMapUintFloat64ErrPreserveOrder(f func(uint) (float64, error), list []uint, worker int) ([]float64, error) {
	chJobs := make(chan map[int]uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float64, chJobs chan map[int]uint, errCh chan error) {
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

func pMapUintFloat64ErrNoOrder(f func(uint) (float64, error), list []uint, worker int) ([]float64, error) {
	chJobs := make(chan uint, len(list))
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

		go func(chResult chan float64, chJobs chan uint, errCh chan error) {
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

// PMapUint64IntErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64IntErr(f func(uint64) (int, error), list []uint64, optional ...Optional) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64IntErrNoOrder(f, list, worker)
		}
	}

	return pMapUint64IntErrPreserveOrder(f, list, worker)
}

func pMapUint64IntErrPreserveOrder(f func(uint64) (int, error), list []uint64, worker int) ([]int, error) {
	chJobs := make(chan map[int]uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int, chJobs chan map[int]uint64, errCh chan error) {
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

func pMapUint64IntErrNoOrder(f func(uint64) (int, error), list []uint64, worker int) ([]int, error) {
	chJobs := make(chan uint64, len(list))
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

		go func(chResult chan int, chJobs chan uint64, errCh chan error) {
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

// PMapUint64Int64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Int64Err(f func(uint64) (int64, error), list []uint64, optional ...Optional) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64Int64ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint64Int64ErrPreserveOrder(f, list, worker)
}

func pMapUint64Int64ErrPreserveOrder(f func(uint64) (int64, error), list []uint64, worker int) ([]int64, error) {
	chJobs := make(chan map[int]uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int64, chJobs chan map[int]uint64, errCh chan error) {
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

func pMapUint64Int64ErrNoOrder(f func(uint64) (int64, error), list []uint64, worker int) ([]int64, error) {
	chJobs := make(chan uint64, len(list))
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

		go func(chResult chan int64, chJobs chan uint64, errCh chan error) {
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

// PMapUint64Int32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Int32Err(f func(uint64) (int32, error), list []uint64, optional ...Optional) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64Int32ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint64Int32ErrPreserveOrder(f, list, worker)
}

func pMapUint64Int32ErrPreserveOrder(f func(uint64) (int32, error), list []uint64, worker int) ([]int32, error) {
	chJobs := make(chan map[int]uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int32, chJobs chan map[int]uint64, errCh chan error) {
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

func pMapUint64Int32ErrNoOrder(f func(uint64) (int32, error), list []uint64, worker int) ([]int32, error) {
	chJobs := make(chan uint64, len(list))
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

		go func(chResult chan int32, chJobs chan uint64, errCh chan error) {
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

// PMapUint64Int16Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Int16Err(f func(uint64) (int16, error), list []uint64, optional ...Optional) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64Int16ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint64Int16ErrPreserveOrder(f, list, worker)
}

func pMapUint64Int16ErrPreserveOrder(f func(uint64) (int16, error), list []uint64, worker int) ([]int16, error) {
	chJobs := make(chan map[int]uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int16, chJobs chan map[int]uint64, errCh chan error) {
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

func pMapUint64Int16ErrNoOrder(f func(uint64) (int16, error), list []uint64, worker int) ([]int16, error) {
	chJobs := make(chan uint64, len(list))
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

		go func(chResult chan int16, chJobs chan uint64, errCh chan error) {
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

// PMapUint64Int8Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Int8Err(f func(uint64) (int8, error), list []uint64, optional ...Optional) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64Int8ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint64Int8ErrPreserveOrder(f, list, worker)
}

func pMapUint64Int8ErrPreserveOrder(f func(uint64) (int8, error), list []uint64, worker int) ([]int8, error) {
	chJobs := make(chan map[int]uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int8, chJobs chan map[int]uint64, errCh chan error) {
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

func pMapUint64Int8ErrNoOrder(f func(uint64) (int8, error), list []uint64, worker int) ([]int8, error) {
	chJobs := make(chan uint64, len(list))
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

		go func(chResult chan int8, chJobs chan uint64, errCh chan error) {
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

// PMapUint64UintErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64UintErr(f func(uint64) (uint, error), list []uint64, optional ...Optional) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64UintErrNoOrder(f, list, worker)
		}
	}

	return pMapUint64UintErrPreserveOrder(f, list, worker)
}

func pMapUint64UintErrPreserveOrder(f func(uint64) (uint, error), list []uint64, worker int) ([]uint, error) {
	chJobs := make(chan map[int]uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint, chJobs chan map[int]uint64, errCh chan error) {
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

func pMapUint64UintErrNoOrder(f func(uint64) (uint, error), list []uint64, worker int) ([]uint, error) {
	chJobs := make(chan uint64, len(list))
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

		go func(chResult chan uint, chJobs chan uint64, errCh chan error) {
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

// PMapUint64Uint32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Uint32Err(f func(uint64) (uint32, error), list []uint64, optional ...Optional) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64Uint32ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint64Uint32ErrPreserveOrder(f, list, worker)
}

func pMapUint64Uint32ErrPreserveOrder(f func(uint64) (uint32, error), list []uint64, worker int) ([]uint32, error) {
	chJobs := make(chan map[int]uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint32, chJobs chan map[int]uint64, errCh chan error) {
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

func pMapUint64Uint32ErrNoOrder(f func(uint64) (uint32, error), list []uint64, worker int) ([]uint32, error) {
	chJobs := make(chan uint64, len(list))
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

		go func(chResult chan uint32, chJobs chan uint64, errCh chan error) {
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

// PMapUint64Uint16Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Uint16Err(f func(uint64) (uint16, error), list []uint64, optional ...Optional) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64Uint16ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint64Uint16ErrPreserveOrder(f, list, worker)
}

func pMapUint64Uint16ErrPreserveOrder(f func(uint64) (uint16, error), list []uint64, worker int) ([]uint16, error) {
	chJobs := make(chan map[int]uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint16, chJobs chan map[int]uint64, errCh chan error) {
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

func pMapUint64Uint16ErrNoOrder(f func(uint64) (uint16, error), list []uint64, worker int) ([]uint16, error) {
	chJobs := make(chan uint64, len(list))
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

		go func(chResult chan uint16, chJobs chan uint64, errCh chan error) {
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

// PMapUint64Uint8Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Uint8Err(f func(uint64) (uint8, error), list []uint64, optional ...Optional) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64Uint8ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint64Uint8ErrPreserveOrder(f, list, worker)
}

func pMapUint64Uint8ErrPreserveOrder(f func(uint64) (uint8, error), list []uint64, worker int) ([]uint8, error) {
	chJobs := make(chan map[int]uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint8, chJobs chan map[int]uint64, errCh chan error) {
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

func pMapUint64Uint8ErrNoOrder(f func(uint64) (uint8, error), list []uint64, worker int) ([]uint8, error) {
	chJobs := make(chan uint64, len(list))
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

		go func(chResult chan uint8, chJobs chan uint64, errCh chan error) {
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

// PMapUint64StrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64StrErr(f func(uint64) (string, error), list []uint64, optional ...Optional) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64StrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint64StrErrPreserveOrder(f, list, worker)
}

func pMapUint64StrErrPreserveOrder(f func(uint64) (string, error), list []uint64, worker int) ([]string, error) {
	chJobs := make(chan map[int]uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]string, chJobs chan map[int]uint64, errCh chan error) {
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

func pMapUint64StrErrNoOrder(f func(uint64) (string, error), list []uint64, worker int) ([]string, error) {
	chJobs := make(chan uint64, len(list))
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

		go func(chResult chan string, chJobs chan uint64, errCh chan error) {
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

// PMapUint64BoolErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64BoolErr(f func(uint64) (bool, error), list []uint64, optional ...Optional) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64BoolErrNoOrder(f, list, worker)
		}
	}

	return pMapUint64BoolErrPreserveOrder(f, list, worker)
}

func pMapUint64BoolErrPreserveOrder(f func(uint64) (bool, error), list []uint64, worker int) ([]bool, error) {
	chJobs := make(chan map[int]uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]bool, chJobs chan map[int]uint64, errCh chan error) {
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

func pMapUint64BoolErrNoOrder(f func(uint64) (bool, error), list []uint64, worker int) ([]bool, error) {
	chJobs := make(chan uint64, len(list))
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

		go func(chResult chan bool, chJobs chan uint64, errCh chan error) {
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

// PMapUint64Float32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Float32Err(f func(uint64) (float32, error), list []uint64, optional ...Optional) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64Float32ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint64Float32ErrPreserveOrder(f, list, worker)
}

func pMapUint64Float32ErrPreserveOrder(f func(uint64) (float32, error), list []uint64, worker int) ([]float32, error) {
	chJobs := make(chan map[int]uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float32, chJobs chan map[int]uint64, errCh chan error) {
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

func pMapUint64Float32ErrNoOrder(f func(uint64) (float32, error), list []uint64, worker int) ([]float32, error) {
	chJobs := make(chan uint64, len(list))
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

		go func(chResult chan float32, chJobs chan uint64, errCh chan error) {
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

// PMapUint64Float64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Float64Err(f func(uint64) (float64, error), list []uint64, optional ...Optional) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64Float64ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint64Float64ErrPreserveOrder(f, list, worker)
}

func pMapUint64Float64ErrPreserveOrder(f func(uint64) (float64, error), list []uint64, worker int) ([]float64, error) {
	chJobs := make(chan map[int]uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float64, chJobs chan map[int]uint64, errCh chan error) {
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

func pMapUint64Float64ErrNoOrder(f func(uint64) (float64, error), list []uint64, worker int) ([]float64, error) {
	chJobs := make(chan uint64, len(list))
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

		go func(chResult chan float64, chJobs chan uint64, errCh chan error) {
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

// PMapUint32IntErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32IntErr(f func(uint32) (int, error), list []uint32, optional ...Optional) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32IntErrNoOrder(f, list, worker)
		}
	}

	return pMapUint32IntErrPreserveOrder(f, list, worker)
}

func pMapUint32IntErrPreserveOrder(f func(uint32) (int, error), list []uint32, worker int) ([]int, error) {
	chJobs := make(chan map[int]uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int, chJobs chan map[int]uint32, errCh chan error) {
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

func pMapUint32IntErrNoOrder(f func(uint32) (int, error), list []uint32, worker int) ([]int, error) {
	chJobs := make(chan uint32, len(list))
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

		go func(chResult chan int, chJobs chan uint32, errCh chan error) {
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

// PMapUint32Int64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Int64Err(f func(uint32) (int64, error), list []uint32, optional ...Optional) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32Int64ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint32Int64ErrPreserveOrder(f, list, worker)
}

func pMapUint32Int64ErrPreserveOrder(f func(uint32) (int64, error), list []uint32, worker int) ([]int64, error) {
	chJobs := make(chan map[int]uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int64, chJobs chan map[int]uint32, errCh chan error) {
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

func pMapUint32Int64ErrNoOrder(f func(uint32) (int64, error), list []uint32, worker int) ([]int64, error) {
	chJobs := make(chan uint32, len(list))
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

		go func(chResult chan int64, chJobs chan uint32, errCh chan error) {
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

// PMapUint32Int32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Int32Err(f func(uint32) (int32, error), list []uint32, optional ...Optional) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32Int32ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint32Int32ErrPreserveOrder(f, list, worker)
}

func pMapUint32Int32ErrPreserveOrder(f func(uint32) (int32, error), list []uint32, worker int) ([]int32, error) {
	chJobs := make(chan map[int]uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int32, chJobs chan map[int]uint32, errCh chan error) {
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

func pMapUint32Int32ErrNoOrder(f func(uint32) (int32, error), list []uint32, worker int) ([]int32, error) {
	chJobs := make(chan uint32, len(list))
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

		go func(chResult chan int32, chJobs chan uint32, errCh chan error) {
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

// PMapUint32Int16Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Int16Err(f func(uint32) (int16, error), list []uint32, optional ...Optional) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32Int16ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint32Int16ErrPreserveOrder(f, list, worker)
}

func pMapUint32Int16ErrPreserveOrder(f func(uint32) (int16, error), list []uint32, worker int) ([]int16, error) {
	chJobs := make(chan map[int]uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int16, chJobs chan map[int]uint32, errCh chan error) {
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

func pMapUint32Int16ErrNoOrder(f func(uint32) (int16, error), list []uint32, worker int) ([]int16, error) {
	chJobs := make(chan uint32, len(list))
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

		go func(chResult chan int16, chJobs chan uint32, errCh chan error) {
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

// PMapUint32Int8Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Int8Err(f func(uint32) (int8, error), list []uint32, optional ...Optional) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32Int8ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint32Int8ErrPreserveOrder(f, list, worker)
}

func pMapUint32Int8ErrPreserveOrder(f func(uint32) (int8, error), list []uint32, worker int) ([]int8, error) {
	chJobs := make(chan map[int]uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int8, chJobs chan map[int]uint32, errCh chan error) {
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

func pMapUint32Int8ErrNoOrder(f func(uint32) (int8, error), list []uint32, worker int) ([]int8, error) {
	chJobs := make(chan uint32, len(list))
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

		go func(chResult chan int8, chJobs chan uint32, errCh chan error) {
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

// PMapUint32UintErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32UintErr(f func(uint32) (uint, error), list []uint32, optional ...Optional) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32UintErrNoOrder(f, list, worker)
		}
	}

	return pMapUint32UintErrPreserveOrder(f, list, worker)
}

func pMapUint32UintErrPreserveOrder(f func(uint32) (uint, error), list []uint32, worker int) ([]uint, error) {
	chJobs := make(chan map[int]uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint, chJobs chan map[int]uint32, errCh chan error) {
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

func pMapUint32UintErrNoOrder(f func(uint32) (uint, error), list []uint32, worker int) ([]uint, error) {
	chJobs := make(chan uint32, len(list))
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

		go func(chResult chan uint, chJobs chan uint32, errCh chan error) {
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

// PMapUint32Uint64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Uint64Err(f func(uint32) (uint64, error), list []uint32, optional ...Optional) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32Uint64ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint32Uint64ErrPreserveOrder(f, list, worker)
}

func pMapUint32Uint64ErrPreserveOrder(f func(uint32) (uint64, error), list []uint32, worker int) ([]uint64, error) {
	chJobs := make(chan map[int]uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint64, chJobs chan map[int]uint32, errCh chan error) {
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

func pMapUint32Uint64ErrNoOrder(f func(uint32) (uint64, error), list []uint32, worker int) ([]uint64, error) {
	chJobs := make(chan uint32, len(list))
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

		go func(chResult chan uint64, chJobs chan uint32, errCh chan error) {
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

// PMapUint32Uint16Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Uint16Err(f func(uint32) (uint16, error), list []uint32, optional ...Optional) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32Uint16ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint32Uint16ErrPreserveOrder(f, list, worker)
}

func pMapUint32Uint16ErrPreserveOrder(f func(uint32) (uint16, error), list []uint32, worker int) ([]uint16, error) {
	chJobs := make(chan map[int]uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint16, chJobs chan map[int]uint32, errCh chan error) {
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

func pMapUint32Uint16ErrNoOrder(f func(uint32) (uint16, error), list []uint32, worker int) ([]uint16, error) {
	chJobs := make(chan uint32, len(list))
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

		go func(chResult chan uint16, chJobs chan uint32, errCh chan error) {
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

// PMapUint32Uint8Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Uint8Err(f func(uint32) (uint8, error), list []uint32, optional ...Optional) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32Uint8ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint32Uint8ErrPreserveOrder(f, list, worker)
}

func pMapUint32Uint8ErrPreserveOrder(f func(uint32) (uint8, error), list []uint32, worker int) ([]uint8, error) {
	chJobs := make(chan map[int]uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint8, chJobs chan map[int]uint32, errCh chan error) {
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

func pMapUint32Uint8ErrNoOrder(f func(uint32) (uint8, error), list []uint32, worker int) ([]uint8, error) {
	chJobs := make(chan uint32, len(list))
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

		go func(chResult chan uint8, chJobs chan uint32, errCh chan error) {
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

// PMapUint32StrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32StrErr(f func(uint32) (string, error), list []uint32, optional ...Optional) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32StrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint32StrErrPreserveOrder(f, list, worker)
}

func pMapUint32StrErrPreserveOrder(f func(uint32) (string, error), list []uint32, worker int) ([]string, error) {
	chJobs := make(chan map[int]uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]string, chJobs chan map[int]uint32, errCh chan error) {
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

func pMapUint32StrErrNoOrder(f func(uint32) (string, error), list []uint32, worker int) ([]string, error) {
	chJobs := make(chan uint32, len(list))
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

		go func(chResult chan string, chJobs chan uint32, errCh chan error) {
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

// PMapUint32BoolErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32BoolErr(f func(uint32) (bool, error), list []uint32, optional ...Optional) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32BoolErrNoOrder(f, list, worker)
		}
	}

	return pMapUint32BoolErrPreserveOrder(f, list, worker)
}

func pMapUint32BoolErrPreserveOrder(f func(uint32) (bool, error), list []uint32, worker int) ([]bool, error) {
	chJobs := make(chan map[int]uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]bool, chJobs chan map[int]uint32, errCh chan error) {
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

func pMapUint32BoolErrNoOrder(f func(uint32) (bool, error), list []uint32, worker int) ([]bool, error) {
	chJobs := make(chan uint32, len(list))
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

		go func(chResult chan bool, chJobs chan uint32, errCh chan error) {
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

// PMapUint32Float32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Float32Err(f func(uint32) (float32, error), list []uint32, optional ...Optional) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32Float32ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint32Float32ErrPreserveOrder(f, list, worker)
}

func pMapUint32Float32ErrPreserveOrder(f func(uint32) (float32, error), list []uint32, worker int) ([]float32, error) {
	chJobs := make(chan map[int]uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float32, chJobs chan map[int]uint32, errCh chan error) {
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

func pMapUint32Float32ErrNoOrder(f func(uint32) (float32, error), list []uint32, worker int) ([]float32, error) {
	chJobs := make(chan uint32, len(list))
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

		go func(chResult chan float32, chJobs chan uint32, errCh chan error) {
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

// PMapUint32Float64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Float64Err(f func(uint32) (float64, error), list []uint32, optional ...Optional) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32Float64ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint32Float64ErrPreserveOrder(f, list, worker)
}

func pMapUint32Float64ErrPreserveOrder(f func(uint32) (float64, error), list []uint32, worker int) ([]float64, error) {
	chJobs := make(chan map[int]uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float64, chJobs chan map[int]uint32, errCh chan error) {
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

func pMapUint32Float64ErrNoOrder(f func(uint32) (float64, error), list []uint32, worker int) ([]float64, error) {
	chJobs := make(chan uint32, len(list))
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

		go func(chResult chan float64, chJobs chan uint32, errCh chan error) {
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

// PMapUint16IntErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16IntErr(f func(uint16) (int, error), list []uint16, optional ...Optional) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16IntErrNoOrder(f, list, worker)
		}
	}

	return pMapUint16IntErrPreserveOrder(f, list, worker)
}

func pMapUint16IntErrPreserveOrder(f func(uint16) (int, error), list []uint16, worker int) ([]int, error) {
	chJobs := make(chan map[int]uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int, chJobs chan map[int]uint16, errCh chan error) {
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

func pMapUint16IntErrNoOrder(f func(uint16) (int, error), list []uint16, worker int) ([]int, error) {
	chJobs := make(chan uint16, len(list))
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

		go func(chResult chan int, chJobs chan uint16, errCh chan error) {
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

// PMapUint16Int64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Int64Err(f func(uint16) (int64, error), list []uint16, optional ...Optional) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16Int64ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint16Int64ErrPreserveOrder(f, list, worker)
}

func pMapUint16Int64ErrPreserveOrder(f func(uint16) (int64, error), list []uint16, worker int) ([]int64, error) {
	chJobs := make(chan map[int]uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int64, chJobs chan map[int]uint16, errCh chan error) {
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

func pMapUint16Int64ErrNoOrder(f func(uint16) (int64, error), list []uint16, worker int) ([]int64, error) {
	chJobs := make(chan uint16, len(list))
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

		go func(chResult chan int64, chJobs chan uint16, errCh chan error) {
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

// PMapUint16Int32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Int32Err(f func(uint16) (int32, error), list []uint16, optional ...Optional) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16Int32ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint16Int32ErrPreserveOrder(f, list, worker)
}

func pMapUint16Int32ErrPreserveOrder(f func(uint16) (int32, error), list []uint16, worker int) ([]int32, error) {
	chJobs := make(chan map[int]uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int32, chJobs chan map[int]uint16, errCh chan error) {
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

func pMapUint16Int32ErrNoOrder(f func(uint16) (int32, error), list []uint16, worker int) ([]int32, error) {
	chJobs := make(chan uint16, len(list))
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

		go func(chResult chan int32, chJobs chan uint16, errCh chan error) {
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

// PMapUint16Int16Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Int16Err(f func(uint16) (int16, error), list []uint16, optional ...Optional) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16Int16ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint16Int16ErrPreserveOrder(f, list, worker)
}

func pMapUint16Int16ErrPreserveOrder(f func(uint16) (int16, error), list []uint16, worker int) ([]int16, error) {
	chJobs := make(chan map[int]uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int16, chJobs chan map[int]uint16, errCh chan error) {
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

func pMapUint16Int16ErrNoOrder(f func(uint16) (int16, error), list []uint16, worker int) ([]int16, error) {
	chJobs := make(chan uint16, len(list))
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

		go func(chResult chan int16, chJobs chan uint16, errCh chan error) {
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

// PMapUint16Int8Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Int8Err(f func(uint16) (int8, error), list []uint16, optional ...Optional) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16Int8ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint16Int8ErrPreserveOrder(f, list, worker)
}

func pMapUint16Int8ErrPreserveOrder(f func(uint16) (int8, error), list []uint16, worker int) ([]int8, error) {
	chJobs := make(chan map[int]uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int8, chJobs chan map[int]uint16, errCh chan error) {
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

func pMapUint16Int8ErrNoOrder(f func(uint16) (int8, error), list []uint16, worker int) ([]int8, error) {
	chJobs := make(chan uint16, len(list))
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

		go func(chResult chan int8, chJobs chan uint16, errCh chan error) {
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

// PMapUint16UintErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16UintErr(f func(uint16) (uint, error), list []uint16, optional ...Optional) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16UintErrNoOrder(f, list, worker)
		}
	}

	return pMapUint16UintErrPreserveOrder(f, list, worker)
}

func pMapUint16UintErrPreserveOrder(f func(uint16) (uint, error), list []uint16, worker int) ([]uint, error) {
	chJobs := make(chan map[int]uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint, chJobs chan map[int]uint16, errCh chan error) {
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

func pMapUint16UintErrNoOrder(f func(uint16) (uint, error), list []uint16, worker int) ([]uint, error) {
	chJobs := make(chan uint16, len(list))
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

		go func(chResult chan uint, chJobs chan uint16, errCh chan error) {
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

// PMapUint16Uint64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Uint64Err(f func(uint16) (uint64, error), list []uint16, optional ...Optional) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16Uint64ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint16Uint64ErrPreserveOrder(f, list, worker)
}

func pMapUint16Uint64ErrPreserveOrder(f func(uint16) (uint64, error), list []uint16, worker int) ([]uint64, error) {
	chJobs := make(chan map[int]uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint64, chJobs chan map[int]uint16, errCh chan error) {
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

func pMapUint16Uint64ErrNoOrder(f func(uint16) (uint64, error), list []uint16, worker int) ([]uint64, error) {
	chJobs := make(chan uint16, len(list))
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

		go func(chResult chan uint64, chJobs chan uint16, errCh chan error) {
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

// PMapUint16Uint32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Uint32Err(f func(uint16) (uint32, error), list []uint16, optional ...Optional) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16Uint32ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint16Uint32ErrPreserveOrder(f, list, worker)
}

func pMapUint16Uint32ErrPreserveOrder(f func(uint16) (uint32, error), list []uint16, worker int) ([]uint32, error) {
	chJobs := make(chan map[int]uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint32, chJobs chan map[int]uint16, errCh chan error) {
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

func pMapUint16Uint32ErrNoOrder(f func(uint16) (uint32, error), list []uint16, worker int) ([]uint32, error) {
	chJobs := make(chan uint16, len(list))
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

		go func(chResult chan uint32, chJobs chan uint16, errCh chan error) {
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

// PMapUint16Uint8Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Uint8Err(f func(uint16) (uint8, error), list []uint16, optional ...Optional) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16Uint8ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint16Uint8ErrPreserveOrder(f, list, worker)
}

func pMapUint16Uint8ErrPreserveOrder(f func(uint16) (uint8, error), list []uint16, worker int) ([]uint8, error) {
	chJobs := make(chan map[int]uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint8, chJobs chan map[int]uint16, errCh chan error) {
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

func pMapUint16Uint8ErrNoOrder(f func(uint16) (uint8, error), list []uint16, worker int) ([]uint8, error) {
	chJobs := make(chan uint16, len(list))
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

		go func(chResult chan uint8, chJobs chan uint16, errCh chan error) {
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

// PMapUint16StrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16StrErr(f func(uint16) (string, error), list []uint16, optional ...Optional) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16StrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint16StrErrPreserveOrder(f, list, worker)
}

func pMapUint16StrErrPreserveOrder(f func(uint16) (string, error), list []uint16, worker int) ([]string, error) {
	chJobs := make(chan map[int]uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]string, chJobs chan map[int]uint16, errCh chan error) {
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

func pMapUint16StrErrNoOrder(f func(uint16) (string, error), list []uint16, worker int) ([]string, error) {
	chJobs := make(chan uint16, len(list))
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

		go func(chResult chan string, chJobs chan uint16, errCh chan error) {
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

// PMapUint16BoolErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16BoolErr(f func(uint16) (bool, error), list []uint16, optional ...Optional) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16BoolErrNoOrder(f, list, worker)
		}
	}

	return pMapUint16BoolErrPreserveOrder(f, list, worker)
}

func pMapUint16BoolErrPreserveOrder(f func(uint16) (bool, error), list []uint16, worker int) ([]bool, error) {
	chJobs := make(chan map[int]uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]bool, chJobs chan map[int]uint16, errCh chan error) {
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

func pMapUint16BoolErrNoOrder(f func(uint16) (bool, error), list []uint16, worker int) ([]bool, error) {
	chJobs := make(chan uint16, len(list))
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

		go func(chResult chan bool, chJobs chan uint16, errCh chan error) {
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

// PMapUint16Float32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Float32Err(f func(uint16) (float32, error), list []uint16, optional ...Optional) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16Float32ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint16Float32ErrPreserveOrder(f, list, worker)
}

func pMapUint16Float32ErrPreserveOrder(f func(uint16) (float32, error), list []uint16, worker int) ([]float32, error) {
	chJobs := make(chan map[int]uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float32, chJobs chan map[int]uint16, errCh chan error) {
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

func pMapUint16Float32ErrNoOrder(f func(uint16) (float32, error), list []uint16, worker int) ([]float32, error) {
	chJobs := make(chan uint16, len(list))
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

		go func(chResult chan float32, chJobs chan uint16, errCh chan error) {
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

// PMapUint16Float64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Float64Err(f func(uint16) (float64, error), list []uint16, optional ...Optional) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16Float64ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint16Float64ErrPreserveOrder(f, list, worker)
}

func pMapUint16Float64ErrPreserveOrder(f func(uint16) (float64, error), list []uint16, worker int) ([]float64, error) {
	chJobs := make(chan map[int]uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float64, chJobs chan map[int]uint16, errCh chan error) {
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

func pMapUint16Float64ErrNoOrder(f func(uint16) (float64, error), list []uint16, worker int) ([]float64, error) {
	chJobs := make(chan uint16, len(list))
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

		go func(chResult chan float64, chJobs chan uint16, errCh chan error) {
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

// PMapUint8IntErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8IntErr(f func(uint8) (int, error), list []uint8, optional ...Optional) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8IntErrNoOrder(f, list, worker)
		}
	}

	return pMapUint8IntErrPreserveOrder(f, list, worker)
}

func pMapUint8IntErrPreserveOrder(f func(uint8) (int, error), list []uint8, worker int) ([]int, error) {
	chJobs := make(chan map[int]uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int, chJobs chan map[int]uint8, errCh chan error) {
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

func pMapUint8IntErrNoOrder(f func(uint8) (int, error), list []uint8, worker int) ([]int, error) {
	chJobs := make(chan uint8, len(list))
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

		go func(chResult chan int, chJobs chan uint8, errCh chan error) {
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

// PMapUint8Int64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Int64Err(f func(uint8) (int64, error), list []uint8, optional ...Optional) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8Int64ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint8Int64ErrPreserveOrder(f, list, worker)
}

func pMapUint8Int64ErrPreserveOrder(f func(uint8) (int64, error), list []uint8, worker int) ([]int64, error) {
	chJobs := make(chan map[int]uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int64, chJobs chan map[int]uint8, errCh chan error) {
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

func pMapUint8Int64ErrNoOrder(f func(uint8) (int64, error), list []uint8, worker int) ([]int64, error) {
	chJobs := make(chan uint8, len(list))
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

		go func(chResult chan int64, chJobs chan uint8, errCh chan error) {
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

// PMapUint8Int32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Int32Err(f func(uint8) (int32, error), list []uint8, optional ...Optional) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8Int32ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint8Int32ErrPreserveOrder(f, list, worker)
}

func pMapUint8Int32ErrPreserveOrder(f func(uint8) (int32, error), list []uint8, worker int) ([]int32, error) {
	chJobs := make(chan map[int]uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int32, chJobs chan map[int]uint8, errCh chan error) {
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

func pMapUint8Int32ErrNoOrder(f func(uint8) (int32, error), list []uint8, worker int) ([]int32, error) {
	chJobs := make(chan uint8, len(list))
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

		go func(chResult chan int32, chJobs chan uint8, errCh chan error) {
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

// PMapUint8Int16Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Int16Err(f func(uint8) (int16, error), list []uint8, optional ...Optional) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8Int16ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint8Int16ErrPreserveOrder(f, list, worker)
}

func pMapUint8Int16ErrPreserveOrder(f func(uint8) (int16, error), list []uint8, worker int) ([]int16, error) {
	chJobs := make(chan map[int]uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int16, chJobs chan map[int]uint8, errCh chan error) {
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

func pMapUint8Int16ErrNoOrder(f func(uint8) (int16, error), list []uint8, worker int) ([]int16, error) {
	chJobs := make(chan uint8, len(list))
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

		go func(chResult chan int16, chJobs chan uint8, errCh chan error) {
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

// PMapUint8Int8Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Int8Err(f func(uint8) (int8, error), list []uint8, optional ...Optional) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8Int8ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint8Int8ErrPreserveOrder(f, list, worker)
}

func pMapUint8Int8ErrPreserveOrder(f func(uint8) (int8, error), list []uint8, worker int) ([]int8, error) {
	chJobs := make(chan map[int]uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int8, chJobs chan map[int]uint8, errCh chan error) {
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

func pMapUint8Int8ErrNoOrder(f func(uint8) (int8, error), list []uint8, worker int) ([]int8, error) {
	chJobs := make(chan uint8, len(list))
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

		go func(chResult chan int8, chJobs chan uint8, errCh chan error) {
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

// PMapUint8UintErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8UintErr(f func(uint8) (uint, error), list []uint8, optional ...Optional) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8UintErrNoOrder(f, list, worker)
		}
	}

	return pMapUint8UintErrPreserveOrder(f, list, worker)
}

func pMapUint8UintErrPreserveOrder(f func(uint8) (uint, error), list []uint8, worker int) ([]uint, error) {
	chJobs := make(chan map[int]uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint, chJobs chan map[int]uint8, errCh chan error) {
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

func pMapUint8UintErrNoOrder(f func(uint8) (uint, error), list []uint8, worker int) ([]uint, error) {
	chJobs := make(chan uint8, len(list))
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

		go func(chResult chan uint, chJobs chan uint8, errCh chan error) {
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

// PMapUint8Uint64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Uint64Err(f func(uint8) (uint64, error), list []uint8, optional ...Optional) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8Uint64ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint8Uint64ErrPreserveOrder(f, list, worker)
}

func pMapUint8Uint64ErrPreserveOrder(f func(uint8) (uint64, error), list []uint8, worker int) ([]uint64, error) {
	chJobs := make(chan map[int]uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint64, chJobs chan map[int]uint8, errCh chan error) {
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

func pMapUint8Uint64ErrNoOrder(f func(uint8) (uint64, error), list []uint8, worker int) ([]uint64, error) {
	chJobs := make(chan uint8, len(list))
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

		go func(chResult chan uint64, chJobs chan uint8, errCh chan error) {
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

// PMapUint8Uint32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Uint32Err(f func(uint8) (uint32, error), list []uint8, optional ...Optional) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8Uint32ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint8Uint32ErrPreserveOrder(f, list, worker)
}

func pMapUint8Uint32ErrPreserveOrder(f func(uint8) (uint32, error), list []uint8, worker int) ([]uint32, error) {
	chJobs := make(chan map[int]uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint32, chJobs chan map[int]uint8, errCh chan error) {
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

func pMapUint8Uint32ErrNoOrder(f func(uint8) (uint32, error), list []uint8, worker int) ([]uint32, error) {
	chJobs := make(chan uint8, len(list))
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

		go func(chResult chan uint32, chJobs chan uint8, errCh chan error) {
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

// PMapUint8Uint16Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Uint16Err(f func(uint8) (uint16, error), list []uint8, optional ...Optional) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8Uint16ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint8Uint16ErrPreserveOrder(f, list, worker)
}

func pMapUint8Uint16ErrPreserveOrder(f func(uint8) (uint16, error), list []uint8, worker int) ([]uint16, error) {
	chJobs := make(chan map[int]uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint16, chJobs chan map[int]uint8, errCh chan error) {
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

func pMapUint8Uint16ErrNoOrder(f func(uint8) (uint16, error), list []uint8, worker int) ([]uint16, error) {
	chJobs := make(chan uint8, len(list))
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

		go func(chResult chan uint16, chJobs chan uint8, errCh chan error) {
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

// PMapUint8StrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8StrErr(f func(uint8) (string, error), list []uint8, optional ...Optional) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8StrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint8StrErrPreserveOrder(f, list, worker)
}

func pMapUint8StrErrPreserveOrder(f func(uint8) (string, error), list []uint8, worker int) ([]string, error) {
	chJobs := make(chan map[int]uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]string, chJobs chan map[int]uint8, errCh chan error) {
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

func pMapUint8StrErrNoOrder(f func(uint8) (string, error), list []uint8, worker int) ([]string, error) {
	chJobs := make(chan uint8, len(list))
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

		go func(chResult chan string, chJobs chan uint8, errCh chan error) {
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

// PMapUint8BoolErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8BoolErr(f func(uint8) (bool, error), list []uint8, optional ...Optional) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8BoolErrNoOrder(f, list, worker)
		}
	}

	return pMapUint8BoolErrPreserveOrder(f, list, worker)
}

func pMapUint8BoolErrPreserveOrder(f func(uint8) (bool, error), list []uint8, worker int) ([]bool, error) {
	chJobs := make(chan map[int]uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]bool, chJobs chan map[int]uint8, errCh chan error) {
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

func pMapUint8BoolErrNoOrder(f func(uint8) (bool, error), list []uint8, worker int) ([]bool, error) {
	chJobs := make(chan uint8, len(list))
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

		go func(chResult chan bool, chJobs chan uint8, errCh chan error) {
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

// PMapUint8Float32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Float32Err(f func(uint8) (float32, error), list []uint8, optional ...Optional) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8Float32ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint8Float32ErrPreserveOrder(f, list, worker)
}

func pMapUint8Float32ErrPreserveOrder(f func(uint8) (float32, error), list []uint8, worker int) ([]float32, error) {
	chJobs := make(chan map[int]uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float32, chJobs chan map[int]uint8, errCh chan error) {
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

func pMapUint8Float32ErrNoOrder(f func(uint8) (float32, error), list []uint8, worker int) ([]float32, error) {
	chJobs := make(chan uint8, len(list))
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

		go func(chResult chan float32, chJobs chan uint8, errCh chan error) {
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

// PMapUint8Float64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Float64Err(f func(uint8) (float64, error), list []uint8, optional ...Optional) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8Float64ErrNoOrder(f, list, worker)
		}
	}

	return pMapUint8Float64ErrPreserveOrder(f, list, worker)
}

func pMapUint8Float64ErrPreserveOrder(f func(uint8) (float64, error), list []uint8, worker int) ([]float64, error) {
	chJobs := make(chan map[int]uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float64, chJobs chan map[int]uint8, errCh chan error) {
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

func pMapUint8Float64ErrNoOrder(f func(uint8) (float64, error), list []uint8, worker int) ([]float64, error) {
	chJobs := make(chan uint8, len(list))
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

		go func(chResult chan float64, chJobs chan uint8, errCh chan error) {
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

// PMapStrIntErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrIntErr(f func(string) (int, error), list []string, optional ...Optional) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrIntErrNoOrder(f, list, worker)
		}
	}

	return pMapStrIntErrPreserveOrder(f, list, worker)
}

func pMapStrIntErrPreserveOrder(f func(string) (int, error), list []string, worker int) ([]int, error) {
	chJobs := make(chan map[int]string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int, chJobs chan map[int]string, errCh chan error) {
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

func pMapStrIntErrNoOrder(f func(string) (int, error), list []string, worker int) ([]int, error) {
	chJobs := make(chan string, len(list))
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

		go func(chResult chan int, chJobs chan string, errCh chan error) {
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

// PMapStrInt64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrInt64Err(f func(string) (int64, error), list []string, optional ...Optional) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrInt64ErrNoOrder(f, list, worker)
		}
	}

	return pMapStrInt64ErrPreserveOrder(f, list, worker)
}

func pMapStrInt64ErrPreserveOrder(f func(string) (int64, error), list []string, worker int) ([]int64, error) {
	chJobs := make(chan map[int]string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int64, chJobs chan map[int]string, errCh chan error) {
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

func pMapStrInt64ErrNoOrder(f func(string) (int64, error), list []string, worker int) ([]int64, error) {
	chJobs := make(chan string, len(list))
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

		go func(chResult chan int64, chJobs chan string, errCh chan error) {
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

// PMapStrInt32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrInt32Err(f func(string) (int32, error), list []string, optional ...Optional) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrInt32ErrNoOrder(f, list, worker)
		}
	}

	return pMapStrInt32ErrPreserveOrder(f, list, worker)
}

func pMapStrInt32ErrPreserveOrder(f func(string) (int32, error), list []string, worker int) ([]int32, error) {
	chJobs := make(chan map[int]string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int32, chJobs chan map[int]string, errCh chan error) {
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

func pMapStrInt32ErrNoOrder(f func(string) (int32, error), list []string, worker int) ([]int32, error) {
	chJobs := make(chan string, len(list))
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

		go func(chResult chan int32, chJobs chan string, errCh chan error) {
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

// PMapStrInt16Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrInt16Err(f func(string) (int16, error), list []string, optional ...Optional) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrInt16ErrNoOrder(f, list, worker)
		}
	}

	return pMapStrInt16ErrPreserveOrder(f, list, worker)
}

func pMapStrInt16ErrPreserveOrder(f func(string) (int16, error), list []string, worker int) ([]int16, error) {
	chJobs := make(chan map[int]string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int16, chJobs chan map[int]string, errCh chan error) {
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

func pMapStrInt16ErrNoOrder(f func(string) (int16, error), list []string, worker int) ([]int16, error) {
	chJobs := make(chan string, len(list))
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

		go func(chResult chan int16, chJobs chan string, errCh chan error) {
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

// PMapStrInt8Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrInt8Err(f func(string) (int8, error), list []string, optional ...Optional) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrInt8ErrNoOrder(f, list, worker)
		}
	}

	return pMapStrInt8ErrPreserveOrder(f, list, worker)
}

func pMapStrInt8ErrPreserveOrder(f func(string) (int8, error), list []string, worker int) ([]int8, error) {
	chJobs := make(chan map[int]string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int8, chJobs chan map[int]string, errCh chan error) {
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

func pMapStrInt8ErrNoOrder(f func(string) (int8, error), list []string, worker int) ([]int8, error) {
	chJobs := make(chan string, len(list))
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

		go func(chResult chan int8, chJobs chan string, errCh chan error) {
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

// PMapStrUintErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrUintErr(f func(string) (uint, error), list []string, optional ...Optional) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrUintErrNoOrder(f, list, worker)
		}
	}

	return pMapStrUintErrPreserveOrder(f, list, worker)
}

func pMapStrUintErrPreserveOrder(f func(string) (uint, error), list []string, worker int) ([]uint, error) {
	chJobs := make(chan map[int]string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint, chJobs chan map[int]string, errCh chan error) {
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

func pMapStrUintErrNoOrder(f func(string) (uint, error), list []string, worker int) ([]uint, error) {
	chJobs := make(chan string, len(list))
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

		go func(chResult chan uint, chJobs chan string, errCh chan error) {
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

// PMapStrUint64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrUint64Err(f func(string) (uint64, error), list []string, optional ...Optional) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrUint64ErrNoOrder(f, list, worker)
		}
	}

	return pMapStrUint64ErrPreserveOrder(f, list, worker)
}

func pMapStrUint64ErrPreserveOrder(f func(string) (uint64, error), list []string, worker int) ([]uint64, error) {
	chJobs := make(chan map[int]string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint64, chJobs chan map[int]string, errCh chan error) {
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

func pMapStrUint64ErrNoOrder(f func(string) (uint64, error), list []string, worker int) ([]uint64, error) {
	chJobs := make(chan string, len(list))
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

		go func(chResult chan uint64, chJobs chan string, errCh chan error) {
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

// PMapStrUint32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrUint32Err(f func(string) (uint32, error), list []string, optional ...Optional) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrUint32ErrNoOrder(f, list, worker)
		}
	}

	return pMapStrUint32ErrPreserveOrder(f, list, worker)
}

func pMapStrUint32ErrPreserveOrder(f func(string) (uint32, error), list []string, worker int) ([]uint32, error) {
	chJobs := make(chan map[int]string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint32, chJobs chan map[int]string, errCh chan error) {
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

func pMapStrUint32ErrNoOrder(f func(string) (uint32, error), list []string, worker int) ([]uint32, error) {
	chJobs := make(chan string, len(list))
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

		go func(chResult chan uint32, chJobs chan string, errCh chan error) {
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

// PMapStrUint16Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrUint16Err(f func(string) (uint16, error), list []string, optional ...Optional) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrUint16ErrNoOrder(f, list, worker)
		}
	}

	return pMapStrUint16ErrPreserveOrder(f, list, worker)
}

func pMapStrUint16ErrPreserveOrder(f func(string) (uint16, error), list []string, worker int) ([]uint16, error) {
	chJobs := make(chan map[int]string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint16, chJobs chan map[int]string, errCh chan error) {
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

func pMapStrUint16ErrNoOrder(f func(string) (uint16, error), list []string, worker int) ([]uint16, error) {
	chJobs := make(chan string, len(list))
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

		go func(chResult chan uint16, chJobs chan string, errCh chan error) {
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

// PMapStrUint8Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrUint8Err(f func(string) (uint8, error), list []string, optional ...Optional) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrUint8ErrNoOrder(f, list, worker)
		}
	}

	return pMapStrUint8ErrPreserveOrder(f, list, worker)
}

func pMapStrUint8ErrPreserveOrder(f func(string) (uint8, error), list []string, worker int) ([]uint8, error) {
	chJobs := make(chan map[int]string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint8, chJobs chan map[int]string, errCh chan error) {
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

func pMapStrUint8ErrNoOrder(f func(string) (uint8, error), list []string, worker int) ([]uint8, error) {
	chJobs := make(chan string, len(list))
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

		go func(chResult chan uint8, chJobs chan string, errCh chan error) {
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

// PMapStrBoolErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrBoolErr(f func(string) (bool, error), list []string, optional ...Optional) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrBoolErrNoOrder(f, list, worker)
		}
	}

	return pMapStrBoolErrPreserveOrder(f, list, worker)
}

func pMapStrBoolErrPreserveOrder(f func(string) (bool, error), list []string, worker int) ([]bool, error) {
	chJobs := make(chan map[int]string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]bool, chJobs chan map[int]string, errCh chan error) {
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

func pMapStrBoolErrNoOrder(f func(string) (bool, error), list []string, worker int) ([]bool, error) {
	chJobs := make(chan string, len(list))
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

		go func(chResult chan bool, chJobs chan string, errCh chan error) {
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

// PMapStrFloat32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrFloat32Err(f func(string) (float32, error), list []string, optional ...Optional) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrFloat32ErrNoOrder(f, list, worker)
		}
	}

	return pMapStrFloat32ErrPreserveOrder(f, list, worker)
}

func pMapStrFloat32ErrPreserveOrder(f func(string) (float32, error), list []string, worker int) ([]float32, error) {
	chJobs := make(chan map[int]string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float32, chJobs chan map[int]string, errCh chan error) {
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

func pMapStrFloat32ErrNoOrder(f func(string) (float32, error), list []string, worker int) ([]float32, error) {
	chJobs := make(chan string, len(list))
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

		go func(chResult chan float32, chJobs chan string, errCh chan error) {
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

// PMapStrFloat64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrFloat64Err(f func(string) (float64, error), list []string, optional ...Optional) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrFloat64ErrNoOrder(f, list, worker)
		}
	}

	return pMapStrFloat64ErrPreserveOrder(f, list, worker)
}

func pMapStrFloat64ErrPreserveOrder(f func(string) (float64, error), list []string, worker int) ([]float64, error) {
	chJobs := make(chan map[int]string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float64, chJobs chan map[int]string, errCh chan error) {
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

func pMapStrFloat64ErrNoOrder(f func(string) (float64, error), list []string, worker int) ([]float64, error) {
	chJobs := make(chan string, len(list))
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

		go func(chResult chan float64, chJobs chan string, errCh chan error) {
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

// PMapBoolIntErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolIntErr(f func(bool) (int, error), list []bool, optional ...Optional) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolIntErrNoOrder(f, list, worker)
		}
	}

	return pMapBoolIntErrPreserveOrder(f, list, worker)
}

func pMapBoolIntErrPreserveOrder(f func(bool) (int, error), list []bool, worker int) ([]int, error) {
	chJobs := make(chan map[int]bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int, chJobs chan map[int]bool, errCh chan error) {
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

func pMapBoolIntErrNoOrder(f func(bool) (int, error), list []bool, worker int) ([]int, error) {
	chJobs := make(chan bool, len(list))
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

		go func(chResult chan int, chJobs chan bool, errCh chan error) {
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

// PMapBoolInt64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolInt64Err(f func(bool) (int64, error), list []bool, optional ...Optional) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolInt64ErrNoOrder(f, list, worker)
		}
	}

	return pMapBoolInt64ErrPreserveOrder(f, list, worker)
}

func pMapBoolInt64ErrPreserveOrder(f func(bool) (int64, error), list []bool, worker int) ([]int64, error) {
	chJobs := make(chan map[int]bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int64, chJobs chan map[int]bool, errCh chan error) {
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

func pMapBoolInt64ErrNoOrder(f func(bool) (int64, error), list []bool, worker int) ([]int64, error) {
	chJobs := make(chan bool, len(list))
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

		go func(chResult chan int64, chJobs chan bool, errCh chan error) {
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

// PMapBoolInt32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolInt32Err(f func(bool) (int32, error), list []bool, optional ...Optional) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolInt32ErrNoOrder(f, list, worker)
		}
	}

	return pMapBoolInt32ErrPreserveOrder(f, list, worker)
}

func pMapBoolInt32ErrPreserveOrder(f func(bool) (int32, error), list []bool, worker int) ([]int32, error) {
	chJobs := make(chan map[int]bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int32, chJobs chan map[int]bool, errCh chan error) {
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

func pMapBoolInt32ErrNoOrder(f func(bool) (int32, error), list []bool, worker int) ([]int32, error) {
	chJobs := make(chan bool, len(list))
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

		go func(chResult chan int32, chJobs chan bool, errCh chan error) {
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

// PMapBoolInt16Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolInt16Err(f func(bool) (int16, error), list []bool, optional ...Optional) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolInt16ErrNoOrder(f, list, worker)
		}
	}

	return pMapBoolInt16ErrPreserveOrder(f, list, worker)
}

func pMapBoolInt16ErrPreserveOrder(f func(bool) (int16, error), list []bool, worker int) ([]int16, error) {
	chJobs := make(chan map[int]bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int16, chJobs chan map[int]bool, errCh chan error) {
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

func pMapBoolInt16ErrNoOrder(f func(bool) (int16, error), list []bool, worker int) ([]int16, error) {
	chJobs := make(chan bool, len(list))
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

		go func(chResult chan int16, chJobs chan bool, errCh chan error) {
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

// PMapBoolInt8Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolInt8Err(f func(bool) (int8, error), list []bool, optional ...Optional) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolInt8ErrNoOrder(f, list, worker)
		}
	}

	return pMapBoolInt8ErrPreserveOrder(f, list, worker)
}

func pMapBoolInt8ErrPreserveOrder(f func(bool) (int8, error), list []bool, worker int) ([]int8, error) {
	chJobs := make(chan map[int]bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int8, chJobs chan map[int]bool, errCh chan error) {
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

func pMapBoolInt8ErrNoOrder(f func(bool) (int8, error), list []bool, worker int) ([]int8, error) {
	chJobs := make(chan bool, len(list))
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

		go func(chResult chan int8, chJobs chan bool, errCh chan error) {
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

// PMapBoolUintErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolUintErr(f func(bool) (uint, error), list []bool, optional ...Optional) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolUintErrNoOrder(f, list, worker)
		}
	}

	return pMapBoolUintErrPreserveOrder(f, list, worker)
}

func pMapBoolUintErrPreserveOrder(f func(bool) (uint, error), list []bool, worker int) ([]uint, error) {
	chJobs := make(chan map[int]bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint, chJobs chan map[int]bool, errCh chan error) {
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

func pMapBoolUintErrNoOrder(f func(bool) (uint, error), list []bool, worker int) ([]uint, error) {
	chJobs := make(chan bool, len(list))
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

		go func(chResult chan uint, chJobs chan bool, errCh chan error) {
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

// PMapBoolUint64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolUint64Err(f func(bool) (uint64, error), list []bool, optional ...Optional) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolUint64ErrNoOrder(f, list, worker)
		}
	}

	return pMapBoolUint64ErrPreserveOrder(f, list, worker)
}

func pMapBoolUint64ErrPreserveOrder(f func(bool) (uint64, error), list []bool, worker int) ([]uint64, error) {
	chJobs := make(chan map[int]bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint64, chJobs chan map[int]bool, errCh chan error) {
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

func pMapBoolUint64ErrNoOrder(f func(bool) (uint64, error), list []bool, worker int) ([]uint64, error) {
	chJobs := make(chan bool, len(list))
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

		go func(chResult chan uint64, chJobs chan bool, errCh chan error) {
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

// PMapBoolUint32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolUint32Err(f func(bool) (uint32, error), list []bool, optional ...Optional) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolUint32ErrNoOrder(f, list, worker)
		}
	}

	return pMapBoolUint32ErrPreserveOrder(f, list, worker)
}

func pMapBoolUint32ErrPreserveOrder(f func(bool) (uint32, error), list []bool, worker int) ([]uint32, error) {
	chJobs := make(chan map[int]bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint32, chJobs chan map[int]bool, errCh chan error) {
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

func pMapBoolUint32ErrNoOrder(f func(bool) (uint32, error), list []bool, worker int) ([]uint32, error) {
	chJobs := make(chan bool, len(list))
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

		go func(chResult chan uint32, chJobs chan bool, errCh chan error) {
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

// PMapBoolUint16Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolUint16Err(f func(bool) (uint16, error), list []bool, optional ...Optional) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolUint16ErrNoOrder(f, list, worker)
		}
	}

	return pMapBoolUint16ErrPreserveOrder(f, list, worker)
}

func pMapBoolUint16ErrPreserveOrder(f func(bool) (uint16, error), list []bool, worker int) ([]uint16, error) {
	chJobs := make(chan map[int]bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint16, chJobs chan map[int]bool, errCh chan error) {
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

func pMapBoolUint16ErrNoOrder(f func(bool) (uint16, error), list []bool, worker int) ([]uint16, error) {
	chJobs := make(chan bool, len(list))
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

		go func(chResult chan uint16, chJobs chan bool, errCh chan error) {
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

// PMapBoolUint8Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolUint8Err(f func(bool) (uint8, error), list []bool, optional ...Optional) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolUint8ErrNoOrder(f, list, worker)
		}
	}

	return pMapBoolUint8ErrPreserveOrder(f, list, worker)
}

func pMapBoolUint8ErrPreserveOrder(f func(bool) (uint8, error), list []bool, worker int) ([]uint8, error) {
	chJobs := make(chan map[int]bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint8, chJobs chan map[int]bool, errCh chan error) {
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

func pMapBoolUint8ErrNoOrder(f func(bool) (uint8, error), list []bool, worker int) ([]uint8, error) {
	chJobs := make(chan bool, len(list))
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

		go func(chResult chan uint8, chJobs chan bool, errCh chan error) {
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

// PMapBoolStrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolStrErr(f func(bool) (string, error), list []bool, optional ...Optional) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolStrErrNoOrder(f, list, worker)
		}
	}

	return pMapBoolStrErrPreserveOrder(f, list, worker)
}

func pMapBoolStrErrPreserveOrder(f func(bool) (string, error), list []bool, worker int) ([]string, error) {
	chJobs := make(chan map[int]bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]string, chJobs chan map[int]bool, errCh chan error) {
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

func pMapBoolStrErrNoOrder(f func(bool) (string, error), list []bool, worker int) ([]string, error) {
	chJobs := make(chan bool, len(list))
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

		go func(chResult chan string, chJobs chan bool, errCh chan error) {
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

// PMapBoolFloat32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolFloat32Err(f func(bool) (float32, error), list []bool, optional ...Optional) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolFloat32ErrNoOrder(f, list, worker)
		}
	}

	return pMapBoolFloat32ErrPreserveOrder(f, list, worker)
}

func pMapBoolFloat32ErrPreserveOrder(f func(bool) (float32, error), list []bool, worker int) ([]float32, error) {
	chJobs := make(chan map[int]bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float32, chJobs chan map[int]bool, errCh chan error) {
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

func pMapBoolFloat32ErrNoOrder(f func(bool) (float32, error), list []bool, worker int) ([]float32, error) {
	chJobs := make(chan bool, len(list))
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

		go func(chResult chan float32, chJobs chan bool, errCh chan error) {
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

// PMapBoolFloat64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolFloat64Err(f func(bool) (float64, error), list []bool, optional ...Optional) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolFloat64ErrNoOrder(f, list, worker)
		}
	}

	return pMapBoolFloat64ErrPreserveOrder(f, list, worker)
}

func pMapBoolFloat64ErrPreserveOrder(f func(bool) (float64, error), list []bool, worker int) ([]float64, error) {
	chJobs := make(chan map[int]bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float64, chJobs chan map[int]bool, errCh chan error) {
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

func pMapBoolFloat64ErrNoOrder(f func(bool) (float64, error), list []bool, worker int) ([]float64, error) {
	chJobs := make(chan bool, len(list))
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

		go func(chResult chan float64, chJobs chan bool, errCh chan error) {
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

// PMapFloat32IntErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32IntErr(f func(float32) (int, error), list []float32, optional ...Optional) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32IntErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32IntErrPreserveOrder(f, list, worker)
}

func pMapFloat32IntErrPreserveOrder(f func(float32) (int, error), list []float32, worker int) ([]int, error) {
	chJobs := make(chan map[int]float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int, chJobs chan map[int]float32, errCh chan error) {
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

func pMapFloat32IntErrNoOrder(f func(float32) (int, error), list []float32, worker int) ([]int, error) {
	chJobs := make(chan float32, len(list))
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

		go func(chResult chan int, chJobs chan float32, errCh chan error) {
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

// PMapFloat32Int64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Int64Err(f func(float32) (int64, error), list []float32, optional ...Optional) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32Int64ErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32Int64ErrPreserveOrder(f, list, worker)
}

func pMapFloat32Int64ErrPreserveOrder(f func(float32) (int64, error), list []float32, worker int) ([]int64, error) {
	chJobs := make(chan map[int]float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int64, chJobs chan map[int]float32, errCh chan error) {
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

func pMapFloat32Int64ErrNoOrder(f func(float32) (int64, error), list []float32, worker int) ([]int64, error) {
	chJobs := make(chan float32, len(list))
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

		go func(chResult chan int64, chJobs chan float32, errCh chan error) {
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

// PMapFloat32Int32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Int32Err(f func(float32) (int32, error), list []float32, optional ...Optional) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32Int32ErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32Int32ErrPreserveOrder(f, list, worker)
}

func pMapFloat32Int32ErrPreserveOrder(f func(float32) (int32, error), list []float32, worker int) ([]int32, error) {
	chJobs := make(chan map[int]float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int32, chJobs chan map[int]float32, errCh chan error) {
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

func pMapFloat32Int32ErrNoOrder(f func(float32) (int32, error), list []float32, worker int) ([]int32, error) {
	chJobs := make(chan float32, len(list))
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

		go func(chResult chan int32, chJobs chan float32, errCh chan error) {
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

// PMapFloat32Int16Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Int16Err(f func(float32) (int16, error), list []float32, optional ...Optional) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32Int16ErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32Int16ErrPreserveOrder(f, list, worker)
}

func pMapFloat32Int16ErrPreserveOrder(f func(float32) (int16, error), list []float32, worker int) ([]int16, error) {
	chJobs := make(chan map[int]float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int16, chJobs chan map[int]float32, errCh chan error) {
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

func pMapFloat32Int16ErrNoOrder(f func(float32) (int16, error), list []float32, worker int) ([]int16, error) {
	chJobs := make(chan float32, len(list))
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

		go func(chResult chan int16, chJobs chan float32, errCh chan error) {
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

// PMapFloat32Int8Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Int8Err(f func(float32) (int8, error), list []float32, optional ...Optional) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32Int8ErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32Int8ErrPreserveOrder(f, list, worker)
}

func pMapFloat32Int8ErrPreserveOrder(f func(float32) (int8, error), list []float32, worker int) ([]int8, error) {
	chJobs := make(chan map[int]float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int8, chJobs chan map[int]float32, errCh chan error) {
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

func pMapFloat32Int8ErrNoOrder(f func(float32) (int8, error), list []float32, worker int) ([]int8, error) {
	chJobs := make(chan float32, len(list))
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

		go func(chResult chan int8, chJobs chan float32, errCh chan error) {
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

// PMapFloat32UintErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32UintErr(f func(float32) (uint, error), list []float32, optional ...Optional) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32UintErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32UintErrPreserveOrder(f, list, worker)
}

func pMapFloat32UintErrPreserveOrder(f func(float32) (uint, error), list []float32, worker int) ([]uint, error) {
	chJobs := make(chan map[int]float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint, chJobs chan map[int]float32, errCh chan error) {
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

func pMapFloat32UintErrNoOrder(f func(float32) (uint, error), list []float32, worker int) ([]uint, error) {
	chJobs := make(chan float32, len(list))
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

		go func(chResult chan uint, chJobs chan float32, errCh chan error) {
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

// PMapFloat32Uint64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Uint64Err(f func(float32) (uint64, error), list []float32, optional ...Optional) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32Uint64ErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32Uint64ErrPreserveOrder(f, list, worker)
}

func pMapFloat32Uint64ErrPreserveOrder(f func(float32) (uint64, error), list []float32, worker int) ([]uint64, error) {
	chJobs := make(chan map[int]float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint64, chJobs chan map[int]float32, errCh chan error) {
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

func pMapFloat32Uint64ErrNoOrder(f func(float32) (uint64, error), list []float32, worker int) ([]uint64, error) {
	chJobs := make(chan float32, len(list))
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

		go func(chResult chan uint64, chJobs chan float32, errCh chan error) {
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

// PMapFloat32Uint32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Uint32Err(f func(float32) (uint32, error), list []float32, optional ...Optional) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32Uint32ErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32Uint32ErrPreserveOrder(f, list, worker)
}

func pMapFloat32Uint32ErrPreserveOrder(f func(float32) (uint32, error), list []float32, worker int) ([]uint32, error) {
	chJobs := make(chan map[int]float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint32, chJobs chan map[int]float32, errCh chan error) {
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

func pMapFloat32Uint32ErrNoOrder(f func(float32) (uint32, error), list []float32, worker int) ([]uint32, error) {
	chJobs := make(chan float32, len(list))
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

		go func(chResult chan uint32, chJobs chan float32, errCh chan error) {
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

// PMapFloat32Uint16Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Uint16Err(f func(float32) (uint16, error), list []float32, optional ...Optional) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32Uint16ErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32Uint16ErrPreserveOrder(f, list, worker)
}

func pMapFloat32Uint16ErrPreserveOrder(f func(float32) (uint16, error), list []float32, worker int) ([]uint16, error) {
	chJobs := make(chan map[int]float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint16, chJobs chan map[int]float32, errCh chan error) {
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

func pMapFloat32Uint16ErrNoOrder(f func(float32) (uint16, error), list []float32, worker int) ([]uint16, error) {
	chJobs := make(chan float32, len(list))
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

		go func(chResult chan uint16, chJobs chan float32, errCh chan error) {
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

// PMapFloat32Uint8Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Uint8Err(f func(float32) (uint8, error), list []float32, optional ...Optional) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32Uint8ErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32Uint8ErrPreserveOrder(f, list, worker)
}

func pMapFloat32Uint8ErrPreserveOrder(f func(float32) (uint8, error), list []float32, worker int) ([]uint8, error) {
	chJobs := make(chan map[int]float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint8, chJobs chan map[int]float32, errCh chan error) {
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

func pMapFloat32Uint8ErrNoOrder(f func(float32) (uint8, error), list []float32, worker int) ([]uint8, error) {
	chJobs := make(chan float32, len(list))
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

		go func(chResult chan uint8, chJobs chan float32, errCh chan error) {
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

// PMapFloat32StrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32StrErr(f func(float32) (string, error), list []float32, optional ...Optional) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32StrErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32StrErrPreserveOrder(f, list, worker)
}

func pMapFloat32StrErrPreserveOrder(f func(float32) (string, error), list []float32, worker int) ([]string, error) {
	chJobs := make(chan map[int]float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]string, chJobs chan map[int]float32, errCh chan error) {
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

func pMapFloat32StrErrNoOrder(f func(float32) (string, error), list []float32, worker int) ([]string, error) {
	chJobs := make(chan float32, len(list))
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

		go func(chResult chan string, chJobs chan float32, errCh chan error) {
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

// PMapFloat32BoolErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32BoolErr(f func(float32) (bool, error), list []float32, optional ...Optional) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32BoolErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32BoolErrPreserveOrder(f, list, worker)
}

func pMapFloat32BoolErrPreserveOrder(f func(float32) (bool, error), list []float32, worker int) ([]bool, error) {
	chJobs := make(chan map[int]float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]bool, chJobs chan map[int]float32, errCh chan error) {
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

func pMapFloat32BoolErrNoOrder(f func(float32) (bool, error), list []float32, worker int) ([]bool, error) {
	chJobs := make(chan float32, len(list))
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

		go func(chResult chan bool, chJobs chan float32, errCh chan error) {
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

// PMapFloat32Float64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Float64Err(f func(float32) (float64, error), list []float32, optional ...Optional) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32Float64ErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32Float64ErrPreserveOrder(f, list, worker)
}

func pMapFloat32Float64ErrPreserveOrder(f func(float32) (float64, error), list []float32, worker int) ([]float64, error) {
	chJobs := make(chan map[int]float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float64, chJobs chan map[int]float32, errCh chan error) {
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

func pMapFloat32Float64ErrNoOrder(f func(float32) (float64, error), list []float32, worker int) ([]float64, error) {
	chJobs := make(chan float32, len(list))
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

		go func(chResult chan float64, chJobs chan float32, errCh chan error) {
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

// PMapFloat64IntErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64IntErr(f func(float64) (int, error), list []float64, optional ...Optional) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64IntErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64IntErrPreserveOrder(f, list, worker)
}

func pMapFloat64IntErrPreserveOrder(f func(float64) (int, error), list []float64, worker int) ([]int, error) {
	chJobs := make(chan map[int]float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int, chJobs chan map[int]float64, errCh chan error) {
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

func pMapFloat64IntErrNoOrder(f func(float64) (int, error), list []float64, worker int) ([]int, error) {
	chJobs := make(chan float64, len(list))
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

		go func(chResult chan int, chJobs chan float64, errCh chan error) {
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

// PMapFloat64Int64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Int64Err(f func(float64) (int64, error), list []float64, optional ...Optional) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64Int64ErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64Int64ErrPreserveOrder(f, list, worker)
}

func pMapFloat64Int64ErrPreserveOrder(f func(float64) (int64, error), list []float64, worker int) ([]int64, error) {
	chJobs := make(chan map[int]float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int64, chJobs chan map[int]float64, errCh chan error) {
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

func pMapFloat64Int64ErrNoOrder(f func(float64) (int64, error), list []float64, worker int) ([]int64, error) {
	chJobs := make(chan float64, len(list))
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

		go func(chResult chan int64, chJobs chan float64, errCh chan error) {
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

// PMapFloat64Int32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Int32Err(f func(float64) (int32, error), list []float64, optional ...Optional) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64Int32ErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64Int32ErrPreserveOrder(f, list, worker)
}

func pMapFloat64Int32ErrPreserveOrder(f func(float64) (int32, error), list []float64, worker int) ([]int32, error) {
	chJobs := make(chan map[int]float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int32, chJobs chan map[int]float64, errCh chan error) {
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

func pMapFloat64Int32ErrNoOrder(f func(float64) (int32, error), list []float64, worker int) ([]int32, error) {
	chJobs := make(chan float64, len(list))
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

		go func(chResult chan int32, chJobs chan float64, errCh chan error) {
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

// PMapFloat64Int16Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Int16Err(f func(float64) (int16, error), list []float64, optional ...Optional) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64Int16ErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64Int16ErrPreserveOrder(f, list, worker)
}

func pMapFloat64Int16ErrPreserveOrder(f func(float64) (int16, error), list []float64, worker int) ([]int16, error) {
	chJobs := make(chan map[int]float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int16, chJobs chan map[int]float64, errCh chan error) {
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

func pMapFloat64Int16ErrNoOrder(f func(float64) (int16, error), list []float64, worker int) ([]int16, error) {
	chJobs := make(chan float64, len(list))
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

		go func(chResult chan int16, chJobs chan float64, errCh chan error) {
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

// PMapFloat64Int8Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Int8Err(f func(float64) (int8, error), list []float64, optional ...Optional) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64Int8ErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64Int8ErrPreserveOrder(f, list, worker)
}

func pMapFloat64Int8ErrPreserveOrder(f func(float64) (int8, error), list []float64, worker int) ([]int8, error) {
	chJobs := make(chan map[int]float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int8, chJobs chan map[int]float64, errCh chan error) {
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

func pMapFloat64Int8ErrNoOrder(f func(float64) (int8, error), list []float64, worker int) ([]int8, error) {
	chJobs := make(chan float64, len(list))
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

		go func(chResult chan int8, chJobs chan float64, errCh chan error) {
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

// PMapFloat64UintErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64UintErr(f func(float64) (uint, error), list []float64, optional ...Optional) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64UintErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64UintErrPreserveOrder(f, list, worker)
}

func pMapFloat64UintErrPreserveOrder(f func(float64) (uint, error), list []float64, worker int) ([]uint, error) {
	chJobs := make(chan map[int]float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint, chJobs chan map[int]float64, errCh chan error) {
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

func pMapFloat64UintErrNoOrder(f func(float64) (uint, error), list []float64, worker int) ([]uint, error) {
	chJobs := make(chan float64, len(list))
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

		go func(chResult chan uint, chJobs chan float64, errCh chan error) {
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

// PMapFloat64Uint64Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Uint64Err(f func(float64) (uint64, error), list []float64, optional ...Optional) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64Uint64ErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64Uint64ErrPreserveOrder(f, list, worker)
}

func pMapFloat64Uint64ErrPreserveOrder(f func(float64) (uint64, error), list []float64, worker int) ([]uint64, error) {
	chJobs := make(chan map[int]float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint64, chJobs chan map[int]float64, errCh chan error) {
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

func pMapFloat64Uint64ErrNoOrder(f func(float64) (uint64, error), list []float64, worker int) ([]uint64, error) {
	chJobs := make(chan float64, len(list))
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

		go func(chResult chan uint64, chJobs chan float64, errCh chan error) {
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

// PMapFloat64Uint32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Uint32Err(f func(float64) (uint32, error), list []float64, optional ...Optional) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64Uint32ErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64Uint32ErrPreserveOrder(f, list, worker)
}

func pMapFloat64Uint32ErrPreserveOrder(f func(float64) (uint32, error), list []float64, worker int) ([]uint32, error) {
	chJobs := make(chan map[int]float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint32, chJobs chan map[int]float64, errCh chan error) {
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

func pMapFloat64Uint32ErrNoOrder(f func(float64) (uint32, error), list []float64, worker int) ([]uint32, error) {
	chJobs := make(chan float64, len(list))
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

		go func(chResult chan uint32, chJobs chan float64, errCh chan error) {
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

// PMapFloat64Uint16Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Uint16Err(f func(float64) (uint16, error), list []float64, optional ...Optional) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64Uint16ErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64Uint16ErrPreserveOrder(f, list, worker)
}

func pMapFloat64Uint16ErrPreserveOrder(f func(float64) (uint16, error), list []float64, worker int) ([]uint16, error) {
	chJobs := make(chan map[int]float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint16, chJobs chan map[int]float64, errCh chan error) {
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

func pMapFloat64Uint16ErrNoOrder(f func(float64) (uint16, error), list []float64, worker int) ([]uint16, error) {
	chJobs := make(chan float64, len(list))
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

		go func(chResult chan uint16, chJobs chan float64, errCh chan error) {
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

// PMapFloat64Uint8Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Uint8Err(f func(float64) (uint8, error), list []float64, optional ...Optional) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64Uint8ErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64Uint8ErrPreserveOrder(f, list, worker)
}

func pMapFloat64Uint8ErrPreserveOrder(f func(float64) (uint8, error), list []float64, worker int) ([]uint8, error) {
	chJobs := make(chan map[int]float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint8, chJobs chan map[int]float64, errCh chan error) {
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

func pMapFloat64Uint8ErrNoOrder(f func(float64) (uint8, error), list []float64, worker int) ([]uint8, error) {
	chJobs := make(chan float64, len(list))
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

		go func(chResult chan uint8, chJobs chan float64, errCh chan error) {
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

// PMapFloat64StrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64StrErr(f func(float64) (string, error), list []float64, optional ...Optional) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64StrErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64StrErrPreserveOrder(f, list, worker)
}

func pMapFloat64StrErrPreserveOrder(f func(float64) (string, error), list []float64, worker int) ([]string, error) {
	chJobs := make(chan map[int]float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]string, chJobs chan map[int]float64, errCh chan error) {
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

func pMapFloat64StrErrNoOrder(f func(float64) (string, error), list []float64, worker int) ([]string, error) {
	chJobs := make(chan float64, len(list))
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

		go func(chResult chan string, chJobs chan float64, errCh chan error) {
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

// PMapFloat64BoolErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64BoolErr(f func(float64) (bool, error), list []float64, optional ...Optional) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64BoolErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64BoolErrPreserveOrder(f, list, worker)
}

func pMapFloat64BoolErrPreserveOrder(f func(float64) (bool, error), list []float64, worker int) ([]bool, error) {
	chJobs := make(chan map[int]float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]bool, chJobs chan map[int]float64, errCh chan error) {
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

func pMapFloat64BoolErrNoOrder(f func(float64) (bool, error), list []float64, worker int) ([]bool, error) {
	chJobs := make(chan float64, len(list))
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

		go func(chResult chan bool, chJobs chan float64, errCh chan error) {
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

// PMapFloat64Float32Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Float32Err(f func(float64) (float32, error), list []float64, optional ...Optional) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64Float32ErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64Float32ErrPreserveOrder(f, list, worker)
}

func pMapFloat64Float32ErrPreserveOrder(f func(float64) (float32, error), list []float64, worker int) ([]float32, error) {
	chJobs := make(chan map[int]float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float32, chJobs chan map[int]float64, errCh chan error) {
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

func pMapFloat64Float32ErrNoOrder(f func(float64) (float32, error), list []float64, worker int) ([]float32, error) {
	chJobs := make(chan float64, len(list))
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

		go func(chResult chan float32, chJobs chan float64, errCh chan error) {
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
