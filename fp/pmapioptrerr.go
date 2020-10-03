package fp

import "sync"

// PMapIntInt64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntInt64PtrErr(f func(*int) (*int64, error), list []*int, optional ...Optional) ([]*int64, error) {
	if f == nil {
		return []*int64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntInt64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapIntInt64PtrErrPreserveOrder(f, list, worker)
}

func pMapIntInt64PtrErrPreserveOrder(f func(*int) (*int64, error), list []*int, worker int) ([]*int64, error) {
	chJobs := make(chan map[int]*int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int64, chJobs chan map[int]*int, errCh chan error) {
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
					chResult <- map[int]*int64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int64, len(list))
	newList := make([]*int64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapIntInt64PtrErrNoOrder(f func(*int) (*int64, error), list []*int, worker int) ([]*int64, error) {
	chJobs := make(chan *int, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int64, chJobs chan *int, errCh chan error) {
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

	newList := make([]*int64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int64{}, <-errCh
	}

	return newList, nil
}

// PMapIntInt32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntInt32PtrErr(f func(*int) (*int32, error), list []*int, optional ...Optional) ([]*int32, error) {
	if f == nil {
		return []*int32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntInt32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapIntInt32PtrErrPreserveOrder(f, list, worker)
}

func pMapIntInt32PtrErrPreserveOrder(f func(*int) (*int32, error), list []*int, worker int) ([]*int32, error) {
	chJobs := make(chan map[int]*int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int32, chJobs chan map[int]*int, errCh chan error) {
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
					chResult <- map[int]*int32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int32, len(list))
	newList := make([]*int32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapIntInt32PtrErrNoOrder(f func(*int) (*int32, error), list []*int, worker int) ([]*int32, error) {
	chJobs := make(chan *int, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int32, chJobs chan *int, errCh chan error) {
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

	newList := make([]*int32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int32{}, <-errCh
	}

	return newList, nil
}

// PMapIntInt16PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntInt16PtrErr(f func(*int) (*int16, error), list []*int, optional ...Optional) ([]*int16, error) {
	if f == nil {
		return []*int16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntInt16PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapIntInt16PtrErrPreserveOrder(f, list, worker)
}

func pMapIntInt16PtrErrPreserveOrder(f func(*int) (*int16, error), list []*int, worker int) ([]*int16, error) {
	chJobs := make(chan map[int]*int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int16, chJobs chan map[int]*int, errCh chan error) {
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
					chResult <- map[int]*int16{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int16, len(list))
	newList := make([]*int16, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int16{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int16{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapIntInt16PtrErrNoOrder(f func(*int) (*int16, error), list []*int, worker int) ([]*int16, error) {
	chJobs := make(chan *int, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int16, chJobs chan *int, errCh chan error) {
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

	newList := make([]*int16, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int16{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int16{}, <-errCh
	}

	return newList, nil
}

// PMapIntInt8PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntInt8PtrErr(f func(*int) (*int8, error), list []*int, optional ...Optional) ([]*int8, error) {
	if f == nil {
		return []*int8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntInt8PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapIntInt8PtrErrPreserveOrder(f, list, worker)
}

func pMapIntInt8PtrErrPreserveOrder(f func(*int) (*int8, error), list []*int, worker int) ([]*int8, error) {
	chJobs := make(chan map[int]*int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int8, chJobs chan map[int]*int, errCh chan error) {
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
					chResult <- map[int]*int8{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int8, len(list))
	newList := make([]*int8, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int8{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int8{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapIntInt8PtrErrNoOrder(f func(*int) (*int8, error), list []*int, worker int) ([]*int8, error) {
	chJobs := make(chan *int, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int8, chJobs chan *int, errCh chan error) {
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

	newList := make([]*int8, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int8{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int8{}, <-errCh
	}

	return newList, nil
}

// PMapIntUintPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntUintPtrErr(f func(*int) (*uint, error), list []*int, optional ...Optional) ([]*uint, error) {
	if f == nil {
		return []*uint{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntUintPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapIntUintPtrErrPreserveOrder(f, list, worker)
}

func pMapIntUintPtrErrPreserveOrder(f func(*int) (*uint, error), list []*int, worker int) ([]*uint, error) {
	chJobs := make(chan map[int]*int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint, chJobs chan map[int]*int, errCh chan error) {
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
					chResult <- map[int]*uint{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint, len(list))
	newList := make([]*uint, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapIntUintPtrErrNoOrder(f func(*int) (*uint, error), list []*int, worker int) ([]*uint, error) {
	chJobs := make(chan *int, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint, chJobs chan *int, errCh chan error) {
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

	newList := make([]*uint, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint{}, <-errCh
	}

	return newList, nil
}

// PMapIntUint64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntUint64PtrErr(f func(*int) (*uint64, error), list []*int, optional ...Optional) ([]*uint64, error) {
	if f == nil {
		return []*uint64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntUint64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapIntUint64PtrErrPreserveOrder(f, list, worker)
}

func pMapIntUint64PtrErrPreserveOrder(f func(*int) (*uint64, error), list []*int, worker int) ([]*uint64, error) {
	chJobs := make(chan map[int]*int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint64, chJobs chan map[int]*int, errCh chan error) {
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
					chResult <- map[int]*uint64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint64, len(list))
	newList := make([]*uint64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapIntUint64PtrErrNoOrder(f func(*int) (*uint64, error), list []*int, worker int) ([]*uint64, error) {
	chJobs := make(chan *int, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint64, chJobs chan *int, errCh chan error) {
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

	newList := make([]*uint64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint64{}, <-errCh
	}

	return newList, nil
}

// PMapIntUint32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntUint32PtrErr(f func(*int) (*uint32, error), list []*int, optional ...Optional) ([]*uint32, error) {
	if f == nil {
		return []*uint32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntUint32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapIntUint32PtrErrPreserveOrder(f, list, worker)
}

func pMapIntUint32PtrErrPreserveOrder(f func(*int) (*uint32, error), list []*int, worker int) ([]*uint32, error) {
	chJobs := make(chan map[int]*int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint32, chJobs chan map[int]*int, errCh chan error) {
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
					chResult <- map[int]*uint32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint32, len(list))
	newList := make([]*uint32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapIntUint32PtrErrNoOrder(f func(*int) (*uint32, error), list []*int, worker int) ([]*uint32, error) {
	chJobs := make(chan *int, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint32, chJobs chan *int, errCh chan error) {
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

	newList := make([]*uint32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint32{}, <-errCh
	}

	return newList, nil
}

// PMapIntUint16PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntUint16PtrErr(f func(*int) (*uint16, error), list []*int, optional ...Optional) ([]*uint16, error) {
	if f == nil {
		return []*uint16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntUint16PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapIntUint16PtrErrPreserveOrder(f, list, worker)
}

func pMapIntUint16PtrErrPreserveOrder(f func(*int) (*uint16, error), list []*int, worker int) ([]*uint16, error) {
	chJobs := make(chan map[int]*int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint16, chJobs chan map[int]*int, errCh chan error) {
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
					chResult <- map[int]*uint16{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint16, len(list))
	newList := make([]*uint16, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint16{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint16{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapIntUint16PtrErrNoOrder(f func(*int) (*uint16, error), list []*int, worker int) ([]*uint16, error) {
	chJobs := make(chan *int, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint16, chJobs chan *int, errCh chan error) {
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

	newList := make([]*uint16, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint16{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint16{}, <-errCh
	}

	return newList, nil
}

// PMapIntUint8PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntUint8PtrErr(f func(*int) (*uint8, error), list []*int, optional ...Optional) ([]*uint8, error) {
	if f == nil {
		return []*uint8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntUint8PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapIntUint8PtrErrPreserveOrder(f, list, worker)
}

func pMapIntUint8PtrErrPreserveOrder(f func(*int) (*uint8, error), list []*int, worker int) ([]*uint8, error) {
	chJobs := make(chan map[int]*int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint8, chJobs chan map[int]*int, errCh chan error) {
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
					chResult <- map[int]*uint8{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint8, len(list))
	newList := make([]*uint8, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint8{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint8{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapIntUint8PtrErrNoOrder(f func(*int) (*uint8, error), list []*int, worker int) ([]*uint8, error) {
	chJobs := make(chan *int, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint8, chJobs chan *int, errCh chan error) {
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

	newList := make([]*uint8, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint8{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint8{}, <-errCh
	}

	return newList, nil
}

// PMapIntStrPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntStrPtrErr(f func(*int) (*string, error), list []*int, optional ...Optional) ([]*string, error) {
	if f == nil {
		return []*string{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntStrPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapIntStrPtrErrPreserveOrder(f, list, worker)
}

func pMapIntStrPtrErrPreserveOrder(f func(*int) (*string, error), list []*int, worker int) ([]*string, error) {
	chJobs := make(chan map[int]*int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*string, chJobs chan map[int]*int, errCh chan error) {
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
					chResult <- map[int]*string{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*string, len(list))
	newList := make([]*string, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*string{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*string{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapIntStrPtrErrNoOrder(f func(*int) (*string, error), list []*int, worker int) ([]*string, error) {
	chJobs := make(chan *int, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *string, chJobs chan *int, errCh chan error) {
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

	newList := make([]*string, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*string{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*string{}, <-errCh
	}

	return newList, nil
}

// PMapIntBoolPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntBoolPtrErr(f func(*int) (*bool, error), list []*int, optional ...Optional) ([]*bool, error) {
	if f == nil {
		return []*bool{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntBoolPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapIntBoolPtrErrPreserveOrder(f, list, worker)
}

func pMapIntBoolPtrErrPreserveOrder(f func(*int) (*bool, error), list []*int, worker int) ([]*bool, error) {
	chJobs := make(chan map[int]*int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*bool, chJobs chan map[int]*int, errCh chan error) {
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
					chResult <- map[int]*bool{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*bool, len(list))
	newList := make([]*bool, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*bool{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*bool{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapIntBoolPtrErrNoOrder(f func(*int) (*bool, error), list []*int, worker int) ([]*bool, error) {
	chJobs := make(chan *int, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *bool, chJobs chan *int, errCh chan error) {
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

	newList := make([]*bool, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*bool{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*bool{}, <-errCh
	}

	return newList, nil
}

// PMapIntFloat32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntFloat32PtrErr(f func(*int) (*float32, error), list []*int, optional ...Optional) ([]*float32, error) {
	if f == nil {
		return []*float32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntFloat32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapIntFloat32PtrErrPreserveOrder(f, list, worker)
}

func pMapIntFloat32PtrErrPreserveOrder(f func(*int) (*float32, error), list []*int, worker int) ([]*float32, error) {
	chJobs := make(chan map[int]*int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float32, chJobs chan map[int]*int, errCh chan error) {
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
					chResult <- map[int]*float32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float32, len(list))
	newList := make([]*float32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*float32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*float32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapIntFloat32PtrErrNoOrder(f func(*int) (*float32, error), list []*int, worker int) ([]*float32, error) {
	chJobs := make(chan *int, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float32, chJobs chan *int, errCh chan error) {
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

	newList := make([]*float32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*float32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*float32{}, <-errCh
	}

	return newList, nil
}

// PMapIntFloat64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntFloat64PtrErr(f func(*int) (*float64, error), list []*int, optional ...Optional) ([]*float64, error) {
	if f == nil {
		return []*float64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntFloat64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapIntFloat64PtrErrPreserveOrder(f, list, worker)
}

func pMapIntFloat64PtrErrPreserveOrder(f func(*int) (*float64, error), list []*int, worker int) ([]*float64, error) {
	chJobs := make(chan map[int]*int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float64, chJobs chan map[int]*int, errCh chan error) {
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
					chResult <- map[int]*float64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float64, len(list))
	newList := make([]*float64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*float64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*float64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapIntFloat64PtrErrNoOrder(f func(*int) (*float64, error), list []*int, worker int) ([]*float64, error) {
	chJobs := make(chan *int, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float64, chJobs chan *int, errCh chan error) {
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

	newList := make([]*float64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*float64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*float64{}, <-errCh
	}

	return newList, nil
}

// PMapInt64IntPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64IntPtrErr(f func(*int64) (*int, error), list []*int64, optional ...Optional) ([]*int, error) {
	if f == nil {
		return []*int{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64IntPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt64IntPtrErrPreserveOrder(f, list, worker)
}

func pMapInt64IntPtrErrPreserveOrder(f func(*int64) (*int, error), list []*int64, worker int) ([]*int, error) {
	chJobs := make(chan map[int]*int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int, chJobs chan map[int]*int64, errCh chan error) {
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
					chResult <- map[int]*int{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int, len(list))
	newList := make([]*int, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt64IntPtrErrNoOrder(f func(*int64) (*int, error), list []*int64, worker int) ([]*int, error) {
	chJobs := make(chan *int64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int, chJobs chan *int64, errCh chan error) {
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

	newList := make([]*int, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int{}, <-errCh
	}

	return newList, nil
}

// PMapInt64Int32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Int32PtrErr(f func(*int64) (*int32, error), list []*int64, optional ...Optional) ([]*int32, error) {
	if f == nil {
		return []*int32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64Int32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt64Int32PtrErrPreserveOrder(f, list, worker)
}

func pMapInt64Int32PtrErrPreserveOrder(f func(*int64) (*int32, error), list []*int64, worker int) ([]*int32, error) {
	chJobs := make(chan map[int]*int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int32, chJobs chan map[int]*int64, errCh chan error) {
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
					chResult <- map[int]*int32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int32, len(list))
	newList := make([]*int32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt64Int32PtrErrNoOrder(f func(*int64) (*int32, error), list []*int64, worker int) ([]*int32, error) {
	chJobs := make(chan *int64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int32, chJobs chan *int64, errCh chan error) {
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

	newList := make([]*int32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int32{}, <-errCh
	}

	return newList, nil
}

// PMapInt64Int16PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Int16PtrErr(f func(*int64) (*int16, error), list []*int64, optional ...Optional) ([]*int16, error) {
	if f == nil {
		return []*int16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64Int16PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt64Int16PtrErrPreserveOrder(f, list, worker)
}

func pMapInt64Int16PtrErrPreserveOrder(f func(*int64) (*int16, error), list []*int64, worker int) ([]*int16, error) {
	chJobs := make(chan map[int]*int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int16, chJobs chan map[int]*int64, errCh chan error) {
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
					chResult <- map[int]*int16{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int16, len(list))
	newList := make([]*int16, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int16{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int16{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt64Int16PtrErrNoOrder(f func(*int64) (*int16, error), list []*int64, worker int) ([]*int16, error) {
	chJobs := make(chan *int64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int16, chJobs chan *int64, errCh chan error) {
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

	newList := make([]*int16, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int16{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int16{}, <-errCh
	}

	return newList, nil
}

// PMapInt64Int8PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Int8PtrErr(f func(*int64) (*int8, error), list []*int64, optional ...Optional) ([]*int8, error) {
	if f == nil {
		return []*int8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64Int8PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt64Int8PtrErrPreserveOrder(f, list, worker)
}

func pMapInt64Int8PtrErrPreserveOrder(f func(*int64) (*int8, error), list []*int64, worker int) ([]*int8, error) {
	chJobs := make(chan map[int]*int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int8, chJobs chan map[int]*int64, errCh chan error) {
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
					chResult <- map[int]*int8{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int8, len(list))
	newList := make([]*int8, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int8{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int8{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt64Int8PtrErrNoOrder(f func(*int64) (*int8, error), list []*int64, worker int) ([]*int8, error) {
	chJobs := make(chan *int64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int8, chJobs chan *int64, errCh chan error) {
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

	newList := make([]*int8, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int8{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int8{}, <-errCh
	}

	return newList, nil
}

// PMapInt64UintPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64UintPtrErr(f func(*int64) (*uint, error), list []*int64, optional ...Optional) ([]*uint, error) {
	if f == nil {
		return []*uint{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64UintPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt64UintPtrErrPreserveOrder(f, list, worker)
}

func pMapInt64UintPtrErrPreserveOrder(f func(*int64) (*uint, error), list []*int64, worker int) ([]*uint, error) {
	chJobs := make(chan map[int]*int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint, chJobs chan map[int]*int64, errCh chan error) {
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
					chResult <- map[int]*uint{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint, len(list))
	newList := make([]*uint, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt64UintPtrErrNoOrder(f func(*int64) (*uint, error), list []*int64, worker int) ([]*uint, error) {
	chJobs := make(chan *int64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint, chJobs chan *int64, errCh chan error) {
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

	newList := make([]*uint, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint{}, <-errCh
	}

	return newList, nil
}

// PMapInt64Uint64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Uint64PtrErr(f func(*int64) (*uint64, error), list []*int64, optional ...Optional) ([]*uint64, error) {
	if f == nil {
		return []*uint64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64Uint64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt64Uint64PtrErrPreserveOrder(f, list, worker)
}

func pMapInt64Uint64PtrErrPreserveOrder(f func(*int64) (*uint64, error), list []*int64, worker int) ([]*uint64, error) {
	chJobs := make(chan map[int]*int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint64, chJobs chan map[int]*int64, errCh chan error) {
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
					chResult <- map[int]*uint64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint64, len(list))
	newList := make([]*uint64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt64Uint64PtrErrNoOrder(f func(*int64) (*uint64, error), list []*int64, worker int) ([]*uint64, error) {
	chJobs := make(chan *int64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint64, chJobs chan *int64, errCh chan error) {
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

	newList := make([]*uint64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint64{}, <-errCh
	}

	return newList, nil
}

// PMapInt64Uint32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Uint32PtrErr(f func(*int64) (*uint32, error), list []*int64, optional ...Optional) ([]*uint32, error) {
	if f == nil {
		return []*uint32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64Uint32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt64Uint32PtrErrPreserveOrder(f, list, worker)
}

func pMapInt64Uint32PtrErrPreserveOrder(f func(*int64) (*uint32, error), list []*int64, worker int) ([]*uint32, error) {
	chJobs := make(chan map[int]*int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint32, chJobs chan map[int]*int64, errCh chan error) {
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
					chResult <- map[int]*uint32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint32, len(list))
	newList := make([]*uint32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt64Uint32PtrErrNoOrder(f func(*int64) (*uint32, error), list []*int64, worker int) ([]*uint32, error) {
	chJobs := make(chan *int64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint32, chJobs chan *int64, errCh chan error) {
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

	newList := make([]*uint32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint32{}, <-errCh
	}

	return newList, nil
}

// PMapInt64Uint16PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Uint16PtrErr(f func(*int64) (*uint16, error), list []*int64, optional ...Optional) ([]*uint16, error) {
	if f == nil {
		return []*uint16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64Uint16PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt64Uint16PtrErrPreserveOrder(f, list, worker)
}

func pMapInt64Uint16PtrErrPreserveOrder(f func(*int64) (*uint16, error), list []*int64, worker int) ([]*uint16, error) {
	chJobs := make(chan map[int]*int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint16, chJobs chan map[int]*int64, errCh chan error) {
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
					chResult <- map[int]*uint16{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint16, len(list))
	newList := make([]*uint16, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint16{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint16{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt64Uint16PtrErrNoOrder(f func(*int64) (*uint16, error), list []*int64, worker int) ([]*uint16, error) {
	chJobs := make(chan *int64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint16, chJobs chan *int64, errCh chan error) {
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

	newList := make([]*uint16, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint16{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint16{}, <-errCh
	}

	return newList, nil
}

// PMapInt64Uint8PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Uint8PtrErr(f func(*int64) (*uint8, error), list []*int64, optional ...Optional) ([]*uint8, error) {
	if f == nil {
		return []*uint8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64Uint8PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt64Uint8PtrErrPreserveOrder(f, list, worker)
}

func pMapInt64Uint8PtrErrPreserveOrder(f func(*int64) (*uint8, error), list []*int64, worker int) ([]*uint8, error) {
	chJobs := make(chan map[int]*int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint8, chJobs chan map[int]*int64, errCh chan error) {
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
					chResult <- map[int]*uint8{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint8, len(list))
	newList := make([]*uint8, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint8{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint8{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt64Uint8PtrErrNoOrder(f func(*int64) (*uint8, error), list []*int64, worker int) ([]*uint8, error) {
	chJobs := make(chan *int64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint8, chJobs chan *int64, errCh chan error) {
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

	newList := make([]*uint8, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint8{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint8{}, <-errCh
	}

	return newList, nil
}

// PMapInt64StrPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64StrPtrErr(f func(*int64) (*string, error), list []*int64, optional ...Optional) ([]*string, error) {
	if f == nil {
		return []*string{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64StrPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt64StrPtrErrPreserveOrder(f, list, worker)
}

func pMapInt64StrPtrErrPreserveOrder(f func(*int64) (*string, error), list []*int64, worker int) ([]*string, error) {
	chJobs := make(chan map[int]*int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*string, chJobs chan map[int]*int64, errCh chan error) {
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
					chResult <- map[int]*string{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*string, len(list))
	newList := make([]*string, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*string{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*string{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt64StrPtrErrNoOrder(f func(*int64) (*string, error), list []*int64, worker int) ([]*string, error) {
	chJobs := make(chan *int64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *string, chJobs chan *int64, errCh chan error) {
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

	newList := make([]*string, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*string{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*string{}, <-errCh
	}

	return newList, nil
}

// PMapInt64BoolPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64BoolPtrErr(f func(*int64) (*bool, error), list []*int64, optional ...Optional) ([]*bool, error) {
	if f == nil {
		return []*bool{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64BoolPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt64BoolPtrErrPreserveOrder(f, list, worker)
}

func pMapInt64BoolPtrErrPreserveOrder(f func(*int64) (*bool, error), list []*int64, worker int) ([]*bool, error) {
	chJobs := make(chan map[int]*int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*bool, chJobs chan map[int]*int64, errCh chan error) {
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
					chResult <- map[int]*bool{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*bool, len(list))
	newList := make([]*bool, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*bool{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*bool{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt64BoolPtrErrNoOrder(f func(*int64) (*bool, error), list []*int64, worker int) ([]*bool, error) {
	chJobs := make(chan *int64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *bool, chJobs chan *int64, errCh chan error) {
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

	newList := make([]*bool, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*bool{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*bool{}, <-errCh
	}

	return newList, nil
}

// PMapInt64Float32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Float32PtrErr(f func(*int64) (*float32, error), list []*int64, optional ...Optional) ([]*float32, error) {
	if f == nil {
		return []*float32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64Float32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt64Float32PtrErrPreserveOrder(f, list, worker)
}

func pMapInt64Float32PtrErrPreserveOrder(f func(*int64) (*float32, error), list []*int64, worker int) ([]*float32, error) {
	chJobs := make(chan map[int]*int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float32, chJobs chan map[int]*int64, errCh chan error) {
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
					chResult <- map[int]*float32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float32, len(list))
	newList := make([]*float32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*float32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*float32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt64Float32PtrErrNoOrder(f func(*int64) (*float32, error), list []*int64, worker int) ([]*float32, error) {
	chJobs := make(chan *int64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float32, chJobs chan *int64, errCh chan error) {
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

	newList := make([]*float32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*float32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*float32{}, <-errCh
	}

	return newList, nil
}

// PMapInt64Float64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Float64PtrErr(f func(*int64) (*float64, error), list []*int64, optional ...Optional) ([]*float64, error) {
	if f == nil {
		return []*float64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64Float64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt64Float64PtrErrPreserveOrder(f, list, worker)
}

func pMapInt64Float64PtrErrPreserveOrder(f func(*int64) (*float64, error), list []*int64, worker int) ([]*float64, error) {
	chJobs := make(chan map[int]*int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float64, chJobs chan map[int]*int64, errCh chan error) {
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
					chResult <- map[int]*float64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float64, len(list))
	newList := make([]*float64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*float64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*float64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt64Float64PtrErrNoOrder(f func(*int64) (*float64, error), list []*int64, worker int) ([]*float64, error) {
	chJobs := make(chan *int64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float64, chJobs chan *int64, errCh chan error) {
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

	newList := make([]*float64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*float64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*float64{}, <-errCh
	}

	return newList, nil
}

// PMapInt32IntPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32IntPtrErr(f func(*int32) (*int, error), list []*int32, optional ...Optional) ([]*int, error) {
	if f == nil {
		return []*int{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32IntPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt32IntPtrErrPreserveOrder(f, list, worker)
}

func pMapInt32IntPtrErrPreserveOrder(f func(*int32) (*int, error), list []*int32, worker int) ([]*int, error) {
	chJobs := make(chan map[int]*int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int, chJobs chan map[int]*int32, errCh chan error) {
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
					chResult <- map[int]*int{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int, len(list))
	newList := make([]*int, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt32IntPtrErrNoOrder(f func(*int32) (*int, error), list []*int32, worker int) ([]*int, error) {
	chJobs := make(chan *int32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int, chJobs chan *int32, errCh chan error) {
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

	newList := make([]*int, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int{}, <-errCh
	}

	return newList, nil
}

// PMapInt32Int64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Int64PtrErr(f func(*int32) (*int64, error), list []*int32, optional ...Optional) ([]*int64, error) {
	if f == nil {
		return []*int64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32Int64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt32Int64PtrErrPreserveOrder(f, list, worker)
}

func pMapInt32Int64PtrErrPreserveOrder(f func(*int32) (*int64, error), list []*int32, worker int) ([]*int64, error) {
	chJobs := make(chan map[int]*int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int64, chJobs chan map[int]*int32, errCh chan error) {
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
					chResult <- map[int]*int64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int64, len(list))
	newList := make([]*int64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt32Int64PtrErrNoOrder(f func(*int32) (*int64, error), list []*int32, worker int) ([]*int64, error) {
	chJobs := make(chan *int32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int64, chJobs chan *int32, errCh chan error) {
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

	newList := make([]*int64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int64{}, <-errCh
	}

	return newList, nil
}

// PMapInt32Int16PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Int16PtrErr(f func(*int32) (*int16, error), list []*int32, optional ...Optional) ([]*int16, error) {
	if f == nil {
		return []*int16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32Int16PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt32Int16PtrErrPreserveOrder(f, list, worker)
}

func pMapInt32Int16PtrErrPreserveOrder(f func(*int32) (*int16, error), list []*int32, worker int) ([]*int16, error) {
	chJobs := make(chan map[int]*int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int16, chJobs chan map[int]*int32, errCh chan error) {
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
					chResult <- map[int]*int16{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int16, len(list))
	newList := make([]*int16, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int16{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int16{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt32Int16PtrErrNoOrder(f func(*int32) (*int16, error), list []*int32, worker int) ([]*int16, error) {
	chJobs := make(chan *int32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int16, chJobs chan *int32, errCh chan error) {
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

	newList := make([]*int16, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int16{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int16{}, <-errCh
	}

	return newList, nil
}

// PMapInt32Int8PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Int8PtrErr(f func(*int32) (*int8, error), list []*int32, optional ...Optional) ([]*int8, error) {
	if f == nil {
		return []*int8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32Int8PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt32Int8PtrErrPreserveOrder(f, list, worker)
}

func pMapInt32Int8PtrErrPreserveOrder(f func(*int32) (*int8, error), list []*int32, worker int) ([]*int8, error) {
	chJobs := make(chan map[int]*int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int8, chJobs chan map[int]*int32, errCh chan error) {
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
					chResult <- map[int]*int8{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int8, len(list))
	newList := make([]*int8, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int8{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int8{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt32Int8PtrErrNoOrder(f func(*int32) (*int8, error), list []*int32, worker int) ([]*int8, error) {
	chJobs := make(chan *int32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int8, chJobs chan *int32, errCh chan error) {
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

	newList := make([]*int8, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int8{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int8{}, <-errCh
	}

	return newList, nil
}

// PMapInt32UintPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32UintPtrErr(f func(*int32) (*uint, error), list []*int32, optional ...Optional) ([]*uint, error) {
	if f == nil {
		return []*uint{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32UintPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt32UintPtrErrPreserveOrder(f, list, worker)
}

func pMapInt32UintPtrErrPreserveOrder(f func(*int32) (*uint, error), list []*int32, worker int) ([]*uint, error) {
	chJobs := make(chan map[int]*int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint, chJobs chan map[int]*int32, errCh chan error) {
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
					chResult <- map[int]*uint{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint, len(list))
	newList := make([]*uint, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt32UintPtrErrNoOrder(f func(*int32) (*uint, error), list []*int32, worker int) ([]*uint, error) {
	chJobs := make(chan *int32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint, chJobs chan *int32, errCh chan error) {
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

	newList := make([]*uint, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint{}, <-errCh
	}

	return newList, nil
}

// PMapInt32Uint64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Uint64PtrErr(f func(*int32) (*uint64, error), list []*int32, optional ...Optional) ([]*uint64, error) {
	if f == nil {
		return []*uint64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32Uint64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt32Uint64PtrErrPreserveOrder(f, list, worker)
}

func pMapInt32Uint64PtrErrPreserveOrder(f func(*int32) (*uint64, error), list []*int32, worker int) ([]*uint64, error) {
	chJobs := make(chan map[int]*int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint64, chJobs chan map[int]*int32, errCh chan error) {
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
					chResult <- map[int]*uint64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint64, len(list))
	newList := make([]*uint64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt32Uint64PtrErrNoOrder(f func(*int32) (*uint64, error), list []*int32, worker int) ([]*uint64, error) {
	chJobs := make(chan *int32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint64, chJobs chan *int32, errCh chan error) {
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

	newList := make([]*uint64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint64{}, <-errCh
	}

	return newList, nil
}

// PMapInt32Uint32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Uint32PtrErr(f func(*int32) (*uint32, error), list []*int32, optional ...Optional) ([]*uint32, error) {
	if f == nil {
		return []*uint32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32Uint32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt32Uint32PtrErrPreserveOrder(f, list, worker)
}

func pMapInt32Uint32PtrErrPreserveOrder(f func(*int32) (*uint32, error), list []*int32, worker int) ([]*uint32, error) {
	chJobs := make(chan map[int]*int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint32, chJobs chan map[int]*int32, errCh chan error) {
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
					chResult <- map[int]*uint32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint32, len(list))
	newList := make([]*uint32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt32Uint32PtrErrNoOrder(f func(*int32) (*uint32, error), list []*int32, worker int) ([]*uint32, error) {
	chJobs := make(chan *int32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint32, chJobs chan *int32, errCh chan error) {
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

	newList := make([]*uint32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint32{}, <-errCh
	}

	return newList, nil
}

// PMapInt32Uint16PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Uint16PtrErr(f func(*int32) (*uint16, error), list []*int32, optional ...Optional) ([]*uint16, error) {
	if f == nil {
		return []*uint16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32Uint16PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt32Uint16PtrErrPreserveOrder(f, list, worker)
}

func pMapInt32Uint16PtrErrPreserveOrder(f func(*int32) (*uint16, error), list []*int32, worker int) ([]*uint16, error) {
	chJobs := make(chan map[int]*int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint16, chJobs chan map[int]*int32, errCh chan error) {
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
					chResult <- map[int]*uint16{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint16, len(list))
	newList := make([]*uint16, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint16{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint16{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt32Uint16PtrErrNoOrder(f func(*int32) (*uint16, error), list []*int32, worker int) ([]*uint16, error) {
	chJobs := make(chan *int32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint16, chJobs chan *int32, errCh chan error) {
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

	newList := make([]*uint16, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint16{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint16{}, <-errCh
	}

	return newList, nil
}

// PMapInt32Uint8PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Uint8PtrErr(f func(*int32) (*uint8, error), list []*int32, optional ...Optional) ([]*uint8, error) {
	if f == nil {
		return []*uint8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32Uint8PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt32Uint8PtrErrPreserveOrder(f, list, worker)
}

func pMapInt32Uint8PtrErrPreserveOrder(f func(*int32) (*uint8, error), list []*int32, worker int) ([]*uint8, error) {
	chJobs := make(chan map[int]*int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint8, chJobs chan map[int]*int32, errCh chan error) {
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
					chResult <- map[int]*uint8{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint8, len(list))
	newList := make([]*uint8, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint8{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint8{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt32Uint8PtrErrNoOrder(f func(*int32) (*uint8, error), list []*int32, worker int) ([]*uint8, error) {
	chJobs := make(chan *int32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint8, chJobs chan *int32, errCh chan error) {
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

	newList := make([]*uint8, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint8{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint8{}, <-errCh
	}

	return newList, nil
}

// PMapInt32StrPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32StrPtrErr(f func(*int32) (*string, error), list []*int32, optional ...Optional) ([]*string, error) {
	if f == nil {
		return []*string{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32StrPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt32StrPtrErrPreserveOrder(f, list, worker)
}

func pMapInt32StrPtrErrPreserveOrder(f func(*int32) (*string, error), list []*int32, worker int) ([]*string, error) {
	chJobs := make(chan map[int]*int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*string, chJobs chan map[int]*int32, errCh chan error) {
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
					chResult <- map[int]*string{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*string, len(list))
	newList := make([]*string, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*string{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*string{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt32StrPtrErrNoOrder(f func(*int32) (*string, error), list []*int32, worker int) ([]*string, error) {
	chJobs := make(chan *int32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *string, chJobs chan *int32, errCh chan error) {
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

	newList := make([]*string, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*string{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*string{}, <-errCh
	}

	return newList, nil
}

// PMapInt32BoolPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32BoolPtrErr(f func(*int32) (*bool, error), list []*int32, optional ...Optional) ([]*bool, error) {
	if f == nil {
		return []*bool{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32BoolPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt32BoolPtrErrPreserveOrder(f, list, worker)
}

func pMapInt32BoolPtrErrPreserveOrder(f func(*int32) (*bool, error), list []*int32, worker int) ([]*bool, error) {
	chJobs := make(chan map[int]*int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*bool, chJobs chan map[int]*int32, errCh chan error) {
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
					chResult <- map[int]*bool{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*bool, len(list))
	newList := make([]*bool, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*bool{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*bool{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt32BoolPtrErrNoOrder(f func(*int32) (*bool, error), list []*int32, worker int) ([]*bool, error) {
	chJobs := make(chan *int32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *bool, chJobs chan *int32, errCh chan error) {
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

	newList := make([]*bool, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*bool{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*bool{}, <-errCh
	}

	return newList, nil
}

// PMapInt32Float32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Float32PtrErr(f func(*int32) (*float32, error), list []*int32, optional ...Optional) ([]*float32, error) {
	if f == nil {
		return []*float32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32Float32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt32Float32PtrErrPreserveOrder(f, list, worker)
}

func pMapInt32Float32PtrErrPreserveOrder(f func(*int32) (*float32, error), list []*int32, worker int) ([]*float32, error) {
	chJobs := make(chan map[int]*int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float32, chJobs chan map[int]*int32, errCh chan error) {
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
					chResult <- map[int]*float32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float32, len(list))
	newList := make([]*float32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*float32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*float32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt32Float32PtrErrNoOrder(f func(*int32) (*float32, error), list []*int32, worker int) ([]*float32, error) {
	chJobs := make(chan *int32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float32, chJobs chan *int32, errCh chan error) {
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

	newList := make([]*float32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*float32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*float32{}, <-errCh
	}

	return newList, nil
}

// PMapInt32Float64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Float64PtrErr(f func(*int32) (*float64, error), list []*int32, optional ...Optional) ([]*float64, error) {
	if f == nil {
		return []*float64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32Float64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt32Float64PtrErrPreserveOrder(f, list, worker)
}

func pMapInt32Float64PtrErrPreserveOrder(f func(*int32) (*float64, error), list []*int32, worker int) ([]*float64, error) {
	chJobs := make(chan map[int]*int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float64, chJobs chan map[int]*int32, errCh chan error) {
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
					chResult <- map[int]*float64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float64, len(list))
	newList := make([]*float64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*float64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*float64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt32Float64PtrErrNoOrder(f func(*int32) (*float64, error), list []*int32, worker int) ([]*float64, error) {
	chJobs := make(chan *int32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float64, chJobs chan *int32, errCh chan error) {
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

	newList := make([]*float64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*float64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*float64{}, <-errCh
	}

	return newList, nil
}

// PMapInt16IntPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16IntPtrErr(f func(*int16) (*int, error), list []*int16, optional ...Optional) ([]*int, error) {
	if f == nil {
		return []*int{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16IntPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt16IntPtrErrPreserveOrder(f, list, worker)
}

func pMapInt16IntPtrErrPreserveOrder(f func(*int16) (*int, error), list []*int16, worker int) ([]*int, error) {
	chJobs := make(chan map[int]*int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int, chJobs chan map[int]*int16, errCh chan error) {
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
					chResult <- map[int]*int{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int, len(list))
	newList := make([]*int, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt16IntPtrErrNoOrder(f func(*int16) (*int, error), list []*int16, worker int) ([]*int, error) {
	chJobs := make(chan *int16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int, chJobs chan *int16, errCh chan error) {
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

	newList := make([]*int, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int{}, <-errCh
	}

	return newList, nil
}

// PMapInt16Int64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Int64PtrErr(f func(*int16) (*int64, error), list []*int16, optional ...Optional) ([]*int64, error) {
	if f == nil {
		return []*int64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16Int64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt16Int64PtrErrPreserveOrder(f, list, worker)
}

func pMapInt16Int64PtrErrPreserveOrder(f func(*int16) (*int64, error), list []*int16, worker int) ([]*int64, error) {
	chJobs := make(chan map[int]*int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int64, chJobs chan map[int]*int16, errCh chan error) {
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
					chResult <- map[int]*int64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int64, len(list))
	newList := make([]*int64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt16Int64PtrErrNoOrder(f func(*int16) (*int64, error), list []*int16, worker int) ([]*int64, error) {
	chJobs := make(chan *int16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int64, chJobs chan *int16, errCh chan error) {
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

	newList := make([]*int64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int64{}, <-errCh
	}

	return newList, nil
}

// PMapInt16Int32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Int32PtrErr(f func(*int16) (*int32, error), list []*int16, optional ...Optional) ([]*int32, error) {
	if f == nil {
		return []*int32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16Int32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt16Int32PtrErrPreserveOrder(f, list, worker)
}

func pMapInt16Int32PtrErrPreserveOrder(f func(*int16) (*int32, error), list []*int16, worker int) ([]*int32, error) {
	chJobs := make(chan map[int]*int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int32, chJobs chan map[int]*int16, errCh chan error) {
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
					chResult <- map[int]*int32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int32, len(list))
	newList := make([]*int32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt16Int32PtrErrNoOrder(f func(*int16) (*int32, error), list []*int16, worker int) ([]*int32, error) {
	chJobs := make(chan *int16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int32, chJobs chan *int16, errCh chan error) {
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

	newList := make([]*int32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int32{}, <-errCh
	}

	return newList, nil
}

// PMapInt16Int8PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Int8PtrErr(f func(*int16) (*int8, error), list []*int16, optional ...Optional) ([]*int8, error) {
	if f == nil {
		return []*int8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16Int8PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt16Int8PtrErrPreserveOrder(f, list, worker)
}

func pMapInt16Int8PtrErrPreserveOrder(f func(*int16) (*int8, error), list []*int16, worker int) ([]*int8, error) {
	chJobs := make(chan map[int]*int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int8, chJobs chan map[int]*int16, errCh chan error) {
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
					chResult <- map[int]*int8{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int8, len(list))
	newList := make([]*int8, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int8{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int8{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt16Int8PtrErrNoOrder(f func(*int16) (*int8, error), list []*int16, worker int) ([]*int8, error) {
	chJobs := make(chan *int16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int8, chJobs chan *int16, errCh chan error) {
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

	newList := make([]*int8, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int8{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int8{}, <-errCh
	}

	return newList, nil
}

// PMapInt16UintPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16UintPtrErr(f func(*int16) (*uint, error), list []*int16, optional ...Optional) ([]*uint, error) {
	if f == nil {
		return []*uint{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16UintPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt16UintPtrErrPreserveOrder(f, list, worker)
}

func pMapInt16UintPtrErrPreserveOrder(f func(*int16) (*uint, error), list []*int16, worker int) ([]*uint, error) {
	chJobs := make(chan map[int]*int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint, chJobs chan map[int]*int16, errCh chan error) {
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
					chResult <- map[int]*uint{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint, len(list))
	newList := make([]*uint, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt16UintPtrErrNoOrder(f func(*int16) (*uint, error), list []*int16, worker int) ([]*uint, error) {
	chJobs := make(chan *int16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint, chJobs chan *int16, errCh chan error) {
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

	newList := make([]*uint, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint{}, <-errCh
	}

	return newList, nil
}

// PMapInt16Uint64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Uint64PtrErr(f func(*int16) (*uint64, error), list []*int16, optional ...Optional) ([]*uint64, error) {
	if f == nil {
		return []*uint64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16Uint64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt16Uint64PtrErrPreserveOrder(f, list, worker)
}

func pMapInt16Uint64PtrErrPreserveOrder(f func(*int16) (*uint64, error), list []*int16, worker int) ([]*uint64, error) {
	chJobs := make(chan map[int]*int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint64, chJobs chan map[int]*int16, errCh chan error) {
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
					chResult <- map[int]*uint64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint64, len(list))
	newList := make([]*uint64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt16Uint64PtrErrNoOrder(f func(*int16) (*uint64, error), list []*int16, worker int) ([]*uint64, error) {
	chJobs := make(chan *int16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint64, chJobs chan *int16, errCh chan error) {
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

	newList := make([]*uint64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint64{}, <-errCh
	}

	return newList, nil
}

// PMapInt16Uint32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Uint32PtrErr(f func(*int16) (*uint32, error), list []*int16, optional ...Optional) ([]*uint32, error) {
	if f == nil {
		return []*uint32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16Uint32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt16Uint32PtrErrPreserveOrder(f, list, worker)
}

func pMapInt16Uint32PtrErrPreserveOrder(f func(*int16) (*uint32, error), list []*int16, worker int) ([]*uint32, error) {
	chJobs := make(chan map[int]*int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint32, chJobs chan map[int]*int16, errCh chan error) {
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
					chResult <- map[int]*uint32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint32, len(list))
	newList := make([]*uint32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt16Uint32PtrErrNoOrder(f func(*int16) (*uint32, error), list []*int16, worker int) ([]*uint32, error) {
	chJobs := make(chan *int16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint32, chJobs chan *int16, errCh chan error) {
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

	newList := make([]*uint32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint32{}, <-errCh
	}

	return newList, nil
}

// PMapInt16Uint16PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Uint16PtrErr(f func(*int16) (*uint16, error), list []*int16, optional ...Optional) ([]*uint16, error) {
	if f == nil {
		return []*uint16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16Uint16PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt16Uint16PtrErrPreserveOrder(f, list, worker)
}

func pMapInt16Uint16PtrErrPreserveOrder(f func(*int16) (*uint16, error), list []*int16, worker int) ([]*uint16, error) {
	chJobs := make(chan map[int]*int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint16, chJobs chan map[int]*int16, errCh chan error) {
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
					chResult <- map[int]*uint16{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint16, len(list))
	newList := make([]*uint16, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint16{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint16{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt16Uint16PtrErrNoOrder(f func(*int16) (*uint16, error), list []*int16, worker int) ([]*uint16, error) {
	chJobs := make(chan *int16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint16, chJobs chan *int16, errCh chan error) {
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

	newList := make([]*uint16, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint16{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint16{}, <-errCh
	}

	return newList, nil
}

// PMapInt16Uint8PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Uint8PtrErr(f func(*int16) (*uint8, error), list []*int16, optional ...Optional) ([]*uint8, error) {
	if f == nil {
		return []*uint8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16Uint8PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt16Uint8PtrErrPreserveOrder(f, list, worker)
}

func pMapInt16Uint8PtrErrPreserveOrder(f func(*int16) (*uint8, error), list []*int16, worker int) ([]*uint8, error) {
	chJobs := make(chan map[int]*int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint8, chJobs chan map[int]*int16, errCh chan error) {
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
					chResult <- map[int]*uint8{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint8, len(list))
	newList := make([]*uint8, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint8{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint8{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt16Uint8PtrErrNoOrder(f func(*int16) (*uint8, error), list []*int16, worker int) ([]*uint8, error) {
	chJobs := make(chan *int16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint8, chJobs chan *int16, errCh chan error) {
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

	newList := make([]*uint8, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint8{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint8{}, <-errCh
	}

	return newList, nil
}

// PMapInt16StrPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16StrPtrErr(f func(*int16) (*string, error), list []*int16, optional ...Optional) ([]*string, error) {
	if f == nil {
		return []*string{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16StrPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt16StrPtrErrPreserveOrder(f, list, worker)
}

func pMapInt16StrPtrErrPreserveOrder(f func(*int16) (*string, error), list []*int16, worker int) ([]*string, error) {
	chJobs := make(chan map[int]*int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*string, chJobs chan map[int]*int16, errCh chan error) {
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
					chResult <- map[int]*string{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*string, len(list))
	newList := make([]*string, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*string{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*string{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt16StrPtrErrNoOrder(f func(*int16) (*string, error), list []*int16, worker int) ([]*string, error) {
	chJobs := make(chan *int16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *string, chJobs chan *int16, errCh chan error) {
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

	newList := make([]*string, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*string{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*string{}, <-errCh
	}

	return newList, nil
}

// PMapInt16BoolPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16BoolPtrErr(f func(*int16) (*bool, error), list []*int16, optional ...Optional) ([]*bool, error) {
	if f == nil {
		return []*bool{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16BoolPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt16BoolPtrErrPreserveOrder(f, list, worker)
}

func pMapInt16BoolPtrErrPreserveOrder(f func(*int16) (*bool, error), list []*int16, worker int) ([]*bool, error) {
	chJobs := make(chan map[int]*int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*bool, chJobs chan map[int]*int16, errCh chan error) {
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
					chResult <- map[int]*bool{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*bool, len(list))
	newList := make([]*bool, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*bool{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*bool{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt16BoolPtrErrNoOrder(f func(*int16) (*bool, error), list []*int16, worker int) ([]*bool, error) {
	chJobs := make(chan *int16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *bool, chJobs chan *int16, errCh chan error) {
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

	newList := make([]*bool, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*bool{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*bool{}, <-errCh
	}

	return newList, nil
}

// PMapInt16Float32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Float32PtrErr(f func(*int16) (*float32, error), list []*int16, optional ...Optional) ([]*float32, error) {
	if f == nil {
		return []*float32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16Float32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt16Float32PtrErrPreserveOrder(f, list, worker)
}

func pMapInt16Float32PtrErrPreserveOrder(f func(*int16) (*float32, error), list []*int16, worker int) ([]*float32, error) {
	chJobs := make(chan map[int]*int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float32, chJobs chan map[int]*int16, errCh chan error) {
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
					chResult <- map[int]*float32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float32, len(list))
	newList := make([]*float32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*float32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*float32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt16Float32PtrErrNoOrder(f func(*int16) (*float32, error), list []*int16, worker int) ([]*float32, error) {
	chJobs := make(chan *int16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float32, chJobs chan *int16, errCh chan error) {
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

	newList := make([]*float32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*float32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*float32{}, <-errCh
	}

	return newList, nil
}

// PMapInt16Float64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Float64PtrErr(f func(*int16) (*float64, error), list []*int16, optional ...Optional) ([]*float64, error) {
	if f == nil {
		return []*float64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16Float64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt16Float64PtrErrPreserveOrder(f, list, worker)
}

func pMapInt16Float64PtrErrPreserveOrder(f func(*int16) (*float64, error), list []*int16, worker int) ([]*float64, error) {
	chJobs := make(chan map[int]*int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float64, chJobs chan map[int]*int16, errCh chan error) {
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
					chResult <- map[int]*float64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float64, len(list))
	newList := make([]*float64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*float64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*float64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt16Float64PtrErrNoOrder(f func(*int16) (*float64, error), list []*int16, worker int) ([]*float64, error) {
	chJobs := make(chan *int16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float64, chJobs chan *int16, errCh chan error) {
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

	newList := make([]*float64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*float64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*float64{}, <-errCh
	}

	return newList, nil
}

// PMapInt8IntPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8IntPtrErr(f func(*int8) (*int, error), list []*int8, optional ...Optional) ([]*int, error) {
	if f == nil {
		return []*int{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8IntPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt8IntPtrErrPreserveOrder(f, list, worker)
}

func pMapInt8IntPtrErrPreserveOrder(f func(*int8) (*int, error), list []*int8, worker int) ([]*int, error) {
	chJobs := make(chan map[int]*int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int, chJobs chan map[int]*int8, errCh chan error) {
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
					chResult <- map[int]*int{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int, len(list))
	newList := make([]*int, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt8IntPtrErrNoOrder(f func(*int8) (*int, error), list []*int8, worker int) ([]*int, error) {
	chJobs := make(chan *int8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int, chJobs chan *int8, errCh chan error) {
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

	newList := make([]*int, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int{}, <-errCh
	}

	return newList, nil
}

// PMapInt8Int64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Int64PtrErr(f func(*int8) (*int64, error), list []*int8, optional ...Optional) ([]*int64, error) {
	if f == nil {
		return []*int64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8Int64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt8Int64PtrErrPreserveOrder(f, list, worker)
}

func pMapInt8Int64PtrErrPreserveOrder(f func(*int8) (*int64, error), list []*int8, worker int) ([]*int64, error) {
	chJobs := make(chan map[int]*int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int64, chJobs chan map[int]*int8, errCh chan error) {
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
					chResult <- map[int]*int64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int64, len(list))
	newList := make([]*int64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt8Int64PtrErrNoOrder(f func(*int8) (*int64, error), list []*int8, worker int) ([]*int64, error) {
	chJobs := make(chan *int8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int64, chJobs chan *int8, errCh chan error) {
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

	newList := make([]*int64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int64{}, <-errCh
	}

	return newList, nil
}

// PMapInt8Int32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Int32PtrErr(f func(*int8) (*int32, error), list []*int8, optional ...Optional) ([]*int32, error) {
	if f == nil {
		return []*int32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8Int32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt8Int32PtrErrPreserveOrder(f, list, worker)
}

func pMapInt8Int32PtrErrPreserveOrder(f func(*int8) (*int32, error), list []*int8, worker int) ([]*int32, error) {
	chJobs := make(chan map[int]*int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int32, chJobs chan map[int]*int8, errCh chan error) {
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
					chResult <- map[int]*int32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int32, len(list))
	newList := make([]*int32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt8Int32PtrErrNoOrder(f func(*int8) (*int32, error), list []*int8, worker int) ([]*int32, error) {
	chJobs := make(chan *int8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int32, chJobs chan *int8, errCh chan error) {
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

	newList := make([]*int32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int32{}, <-errCh
	}

	return newList, nil
}

// PMapInt8Int16PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Int16PtrErr(f func(*int8) (*int16, error), list []*int8, optional ...Optional) ([]*int16, error) {
	if f == nil {
		return []*int16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8Int16PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt8Int16PtrErrPreserveOrder(f, list, worker)
}

func pMapInt8Int16PtrErrPreserveOrder(f func(*int8) (*int16, error), list []*int8, worker int) ([]*int16, error) {
	chJobs := make(chan map[int]*int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int16, chJobs chan map[int]*int8, errCh chan error) {
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
					chResult <- map[int]*int16{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int16, len(list))
	newList := make([]*int16, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int16{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int16{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt8Int16PtrErrNoOrder(f func(*int8) (*int16, error), list []*int8, worker int) ([]*int16, error) {
	chJobs := make(chan *int8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int16, chJobs chan *int8, errCh chan error) {
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

	newList := make([]*int16, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int16{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int16{}, <-errCh
	}

	return newList, nil
}

// PMapInt8UintPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8UintPtrErr(f func(*int8) (*uint, error), list []*int8, optional ...Optional) ([]*uint, error) {
	if f == nil {
		return []*uint{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8UintPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt8UintPtrErrPreserveOrder(f, list, worker)
}

func pMapInt8UintPtrErrPreserveOrder(f func(*int8) (*uint, error), list []*int8, worker int) ([]*uint, error) {
	chJobs := make(chan map[int]*int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint, chJobs chan map[int]*int8, errCh chan error) {
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
					chResult <- map[int]*uint{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint, len(list))
	newList := make([]*uint, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt8UintPtrErrNoOrder(f func(*int8) (*uint, error), list []*int8, worker int) ([]*uint, error) {
	chJobs := make(chan *int8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint, chJobs chan *int8, errCh chan error) {
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

	newList := make([]*uint, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint{}, <-errCh
	}

	return newList, nil
}

// PMapInt8Uint64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Uint64PtrErr(f func(*int8) (*uint64, error), list []*int8, optional ...Optional) ([]*uint64, error) {
	if f == nil {
		return []*uint64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8Uint64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt8Uint64PtrErrPreserveOrder(f, list, worker)
}

func pMapInt8Uint64PtrErrPreserveOrder(f func(*int8) (*uint64, error), list []*int8, worker int) ([]*uint64, error) {
	chJobs := make(chan map[int]*int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint64, chJobs chan map[int]*int8, errCh chan error) {
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
					chResult <- map[int]*uint64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint64, len(list))
	newList := make([]*uint64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt8Uint64PtrErrNoOrder(f func(*int8) (*uint64, error), list []*int8, worker int) ([]*uint64, error) {
	chJobs := make(chan *int8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint64, chJobs chan *int8, errCh chan error) {
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

	newList := make([]*uint64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint64{}, <-errCh
	}

	return newList, nil
}

// PMapInt8Uint32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Uint32PtrErr(f func(*int8) (*uint32, error), list []*int8, optional ...Optional) ([]*uint32, error) {
	if f == nil {
		return []*uint32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8Uint32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt8Uint32PtrErrPreserveOrder(f, list, worker)
}

func pMapInt8Uint32PtrErrPreserveOrder(f func(*int8) (*uint32, error), list []*int8, worker int) ([]*uint32, error) {
	chJobs := make(chan map[int]*int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint32, chJobs chan map[int]*int8, errCh chan error) {
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
					chResult <- map[int]*uint32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint32, len(list))
	newList := make([]*uint32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt8Uint32PtrErrNoOrder(f func(*int8) (*uint32, error), list []*int8, worker int) ([]*uint32, error) {
	chJobs := make(chan *int8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint32, chJobs chan *int8, errCh chan error) {
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

	newList := make([]*uint32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint32{}, <-errCh
	}

	return newList, nil
}

// PMapInt8Uint16PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Uint16PtrErr(f func(*int8) (*uint16, error), list []*int8, optional ...Optional) ([]*uint16, error) {
	if f == nil {
		return []*uint16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8Uint16PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt8Uint16PtrErrPreserveOrder(f, list, worker)
}

func pMapInt8Uint16PtrErrPreserveOrder(f func(*int8) (*uint16, error), list []*int8, worker int) ([]*uint16, error) {
	chJobs := make(chan map[int]*int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint16, chJobs chan map[int]*int8, errCh chan error) {
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
					chResult <- map[int]*uint16{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint16, len(list))
	newList := make([]*uint16, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint16{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint16{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt8Uint16PtrErrNoOrder(f func(*int8) (*uint16, error), list []*int8, worker int) ([]*uint16, error) {
	chJobs := make(chan *int8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint16, chJobs chan *int8, errCh chan error) {
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

	newList := make([]*uint16, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint16{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint16{}, <-errCh
	}

	return newList, nil
}

// PMapInt8Uint8PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Uint8PtrErr(f func(*int8) (*uint8, error), list []*int8, optional ...Optional) ([]*uint8, error) {
	if f == nil {
		return []*uint8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8Uint8PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt8Uint8PtrErrPreserveOrder(f, list, worker)
}

func pMapInt8Uint8PtrErrPreserveOrder(f func(*int8) (*uint8, error), list []*int8, worker int) ([]*uint8, error) {
	chJobs := make(chan map[int]*int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint8, chJobs chan map[int]*int8, errCh chan error) {
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
					chResult <- map[int]*uint8{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint8, len(list))
	newList := make([]*uint8, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint8{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint8{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt8Uint8PtrErrNoOrder(f func(*int8) (*uint8, error), list []*int8, worker int) ([]*uint8, error) {
	chJobs := make(chan *int8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint8, chJobs chan *int8, errCh chan error) {
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

	newList := make([]*uint8, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint8{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint8{}, <-errCh
	}

	return newList, nil
}

// PMapInt8StrPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8StrPtrErr(f func(*int8) (*string, error), list []*int8, optional ...Optional) ([]*string, error) {
	if f == nil {
		return []*string{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8StrPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt8StrPtrErrPreserveOrder(f, list, worker)
}

func pMapInt8StrPtrErrPreserveOrder(f func(*int8) (*string, error), list []*int8, worker int) ([]*string, error) {
	chJobs := make(chan map[int]*int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*string, chJobs chan map[int]*int8, errCh chan error) {
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
					chResult <- map[int]*string{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*string, len(list))
	newList := make([]*string, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*string{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*string{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt8StrPtrErrNoOrder(f func(*int8) (*string, error), list []*int8, worker int) ([]*string, error) {
	chJobs := make(chan *int8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *string, chJobs chan *int8, errCh chan error) {
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

	newList := make([]*string, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*string{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*string{}, <-errCh
	}

	return newList, nil
}

// PMapInt8BoolPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8BoolPtrErr(f func(*int8) (*bool, error), list []*int8, optional ...Optional) ([]*bool, error) {
	if f == nil {
		return []*bool{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8BoolPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt8BoolPtrErrPreserveOrder(f, list, worker)
}

func pMapInt8BoolPtrErrPreserveOrder(f func(*int8) (*bool, error), list []*int8, worker int) ([]*bool, error) {
	chJobs := make(chan map[int]*int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*bool, chJobs chan map[int]*int8, errCh chan error) {
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
					chResult <- map[int]*bool{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*bool, len(list))
	newList := make([]*bool, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*bool{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*bool{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt8BoolPtrErrNoOrder(f func(*int8) (*bool, error), list []*int8, worker int) ([]*bool, error) {
	chJobs := make(chan *int8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *bool, chJobs chan *int8, errCh chan error) {
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

	newList := make([]*bool, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*bool{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*bool{}, <-errCh
	}

	return newList, nil
}

// PMapInt8Float32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Float32PtrErr(f func(*int8) (*float32, error), list []*int8, optional ...Optional) ([]*float32, error) {
	if f == nil {
		return []*float32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8Float32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt8Float32PtrErrPreserveOrder(f, list, worker)
}

func pMapInt8Float32PtrErrPreserveOrder(f func(*int8) (*float32, error), list []*int8, worker int) ([]*float32, error) {
	chJobs := make(chan map[int]*int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float32, chJobs chan map[int]*int8, errCh chan error) {
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
					chResult <- map[int]*float32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float32, len(list))
	newList := make([]*float32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*float32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*float32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt8Float32PtrErrNoOrder(f func(*int8) (*float32, error), list []*int8, worker int) ([]*float32, error) {
	chJobs := make(chan *int8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float32, chJobs chan *int8, errCh chan error) {
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

	newList := make([]*float32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*float32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*float32{}, <-errCh
	}

	return newList, nil
}

// PMapInt8Float64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Float64PtrErr(f func(*int8) (*float64, error), list []*int8, optional ...Optional) ([]*float64, error) {
	if f == nil {
		return []*float64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8Float64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapInt8Float64PtrErrPreserveOrder(f, list, worker)
}

func pMapInt8Float64PtrErrPreserveOrder(f func(*int8) (*float64, error), list []*int8, worker int) ([]*float64, error) {
	chJobs := make(chan map[int]*int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float64, chJobs chan map[int]*int8, errCh chan error) {
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
					chResult <- map[int]*float64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float64, len(list))
	newList := make([]*float64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*float64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*float64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapInt8Float64PtrErrNoOrder(f func(*int8) (*float64, error), list []*int8, worker int) ([]*float64, error) {
	chJobs := make(chan *int8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float64, chJobs chan *int8, errCh chan error) {
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

	newList := make([]*float64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*float64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*float64{}, <-errCh
	}

	return newList, nil
}

// PMapUintIntPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintIntPtrErr(f func(*uint) (*int, error), list []*uint, optional ...Optional) ([]*int, error) {
	if f == nil {
		return []*int{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintIntPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUintIntPtrErrPreserveOrder(f, list, worker)
}

func pMapUintIntPtrErrPreserveOrder(f func(*uint) (*int, error), list []*uint, worker int) ([]*int, error) {
	chJobs := make(chan map[int]*uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int, chJobs chan map[int]*uint, errCh chan error) {
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
					chResult <- map[int]*int{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int, len(list))
	newList := make([]*int, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUintIntPtrErrNoOrder(f func(*uint) (*int, error), list []*uint, worker int) ([]*int, error) {
	chJobs := make(chan *uint, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int, chJobs chan *uint, errCh chan error) {
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

	newList := make([]*int, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int{}, <-errCh
	}

	return newList, nil
}

// PMapUintInt64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintInt64PtrErr(f func(*uint) (*int64, error), list []*uint, optional ...Optional) ([]*int64, error) {
	if f == nil {
		return []*int64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintInt64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUintInt64PtrErrPreserveOrder(f, list, worker)
}

func pMapUintInt64PtrErrPreserveOrder(f func(*uint) (*int64, error), list []*uint, worker int) ([]*int64, error) {
	chJobs := make(chan map[int]*uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int64, chJobs chan map[int]*uint, errCh chan error) {
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
					chResult <- map[int]*int64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int64, len(list))
	newList := make([]*int64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUintInt64PtrErrNoOrder(f func(*uint) (*int64, error), list []*uint, worker int) ([]*int64, error) {
	chJobs := make(chan *uint, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int64, chJobs chan *uint, errCh chan error) {
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

	newList := make([]*int64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int64{}, <-errCh
	}

	return newList, nil
}

// PMapUintInt32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintInt32PtrErr(f func(*uint) (*int32, error), list []*uint, optional ...Optional) ([]*int32, error) {
	if f == nil {
		return []*int32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintInt32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUintInt32PtrErrPreserveOrder(f, list, worker)
}

func pMapUintInt32PtrErrPreserveOrder(f func(*uint) (*int32, error), list []*uint, worker int) ([]*int32, error) {
	chJobs := make(chan map[int]*uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int32, chJobs chan map[int]*uint, errCh chan error) {
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
					chResult <- map[int]*int32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int32, len(list))
	newList := make([]*int32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUintInt32PtrErrNoOrder(f func(*uint) (*int32, error), list []*uint, worker int) ([]*int32, error) {
	chJobs := make(chan *uint, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int32, chJobs chan *uint, errCh chan error) {
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

	newList := make([]*int32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int32{}, <-errCh
	}

	return newList, nil
}

// PMapUintInt16PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintInt16PtrErr(f func(*uint) (*int16, error), list []*uint, optional ...Optional) ([]*int16, error) {
	if f == nil {
		return []*int16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintInt16PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUintInt16PtrErrPreserveOrder(f, list, worker)
}

func pMapUintInt16PtrErrPreserveOrder(f func(*uint) (*int16, error), list []*uint, worker int) ([]*int16, error) {
	chJobs := make(chan map[int]*uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int16, chJobs chan map[int]*uint, errCh chan error) {
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
					chResult <- map[int]*int16{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int16, len(list))
	newList := make([]*int16, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int16{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int16{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUintInt16PtrErrNoOrder(f func(*uint) (*int16, error), list []*uint, worker int) ([]*int16, error) {
	chJobs := make(chan *uint, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int16, chJobs chan *uint, errCh chan error) {
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

	newList := make([]*int16, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int16{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int16{}, <-errCh
	}

	return newList, nil
}

// PMapUintInt8PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintInt8PtrErr(f func(*uint) (*int8, error), list []*uint, optional ...Optional) ([]*int8, error) {
	if f == nil {
		return []*int8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintInt8PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUintInt8PtrErrPreserveOrder(f, list, worker)
}

func pMapUintInt8PtrErrPreserveOrder(f func(*uint) (*int8, error), list []*uint, worker int) ([]*int8, error) {
	chJobs := make(chan map[int]*uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int8, chJobs chan map[int]*uint, errCh chan error) {
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
					chResult <- map[int]*int8{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int8, len(list))
	newList := make([]*int8, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int8{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int8{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUintInt8PtrErrNoOrder(f func(*uint) (*int8, error), list []*uint, worker int) ([]*int8, error) {
	chJobs := make(chan *uint, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int8, chJobs chan *uint, errCh chan error) {
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

	newList := make([]*int8, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int8{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int8{}, <-errCh
	}

	return newList, nil
}

// PMapUintUint64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintUint64PtrErr(f func(*uint) (*uint64, error), list []*uint, optional ...Optional) ([]*uint64, error) {
	if f == nil {
		return []*uint64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintUint64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUintUint64PtrErrPreserveOrder(f, list, worker)
}

func pMapUintUint64PtrErrPreserveOrder(f func(*uint) (*uint64, error), list []*uint, worker int) ([]*uint64, error) {
	chJobs := make(chan map[int]*uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint64, chJobs chan map[int]*uint, errCh chan error) {
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
					chResult <- map[int]*uint64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint64, len(list))
	newList := make([]*uint64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUintUint64PtrErrNoOrder(f func(*uint) (*uint64, error), list []*uint, worker int) ([]*uint64, error) {
	chJobs := make(chan *uint, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint64, chJobs chan *uint, errCh chan error) {
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

	newList := make([]*uint64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint64{}, <-errCh
	}

	return newList, nil
}

// PMapUintUint32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintUint32PtrErr(f func(*uint) (*uint32, error), list []*uint, optional ...Optional) ([]*uint32, error) {
	if f == nil {
		return []*uint32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintUint32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUintUint32PtrErrPreserveOrder(f, list, worker)
}

func pMapUintUint32PtrErrPreserveOrder(f func(*uint) (*uint32, error), list []*uint, worker int) ([]*uint32, error) {
	chJobs := make(chan map[int]*uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint32, chJobs chan map[int]*uint, errCh chan error) {
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
					chResult <- map[int]*uint32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint32, len(list))
	newList := make([]*uint32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUintUint32PtrErrNoOrder(f func(*uint) (*uint32, error), list []*uint, worker int) ([]*uint32, error) {
	chJobs := make(chan *uint, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint32, chJobs chan *uint, errCh chan error) {
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

	newList := make([]*uint32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint32{}, <-errCh
	}

	return newList, nil
}

// PMapUintUint16PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintUint16PtrErr(f func(*uint) (*uint16, error), list []*uint, optional ...Optional) ([]*uint16, error) {
	if f == nil {
		return []*uint16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintUint16PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUintUint16PtrErrPreserveOrder(f, list, worker)
}

func pMapUintUint16PtrErrPreserveOrder(f func(*uint) (*uint16, error), list []*uint, worker int) ([]*uint16, error) {
	chJobs := make(chan map[int]*uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint16, chJobs chan map[int]*uint, errCh chan error) {
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
					chResult <- map[int]*uint16{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint16, len(list))
	newList := make([]*uint16, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint16{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint16{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUintUint16PtrErrNoOrder(f func(*uint) (*uint16, error), list []*uint, worker int) ([]*uint16, error) {
	chJobs := make(chan *uint, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint16, chJobs chan *uint, errCh chan error) {
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

	newList := make([]*uint16, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint16{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint16{}, <-errCh
	}

	return newList, nil
}

// PMapUintUint8PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintUint8PtrErr(f func(*uint) (*uint8, error), list []*uint, optional ...Optional) ([]*uint8, error) {
	if f == nil {
		return []*uint8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintUint8PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUintUint8PtrErrPreserveOrder(f, list, worker)
}

func pMapUintUint8PtrErrPreserveOrder(f func(*uint) (*uint8, error), list []*uint, worker int) ([]*uint8, error) {
	chJobs := make(chan map[int]*uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint8, chJobs chan map[int]*uint, errCh chan error) {
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
					chResult <- map[int]*uint8{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint8, len(list))
	newList := make([]*uint8, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint8{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint8{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUintUint8PtrErrNoOrder(f func(*uint) (*uint8, error), list []*uint, worker int) ([]*uint8, error) {
	chJobs := make(chan *uint, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint8, chJobs chan *uint, errCh chan error) {
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

	newList := make([]*uint8, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint8{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint8{}, <-errCh
	}

	return newList, nil
}

// PMapUintStrPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintStrPtrErr(f func(*uint) (*string, error), list []*uint, optional ...Optional) ([]*string, error) {
	if f == nil {
		return []*string{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintStrPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUintStrPtrErrPreserveOrder(f, list, worker)
}

func pMapUintStrPtrErrPreserveOrder(f func(*uint) (*string, error), list []*uint, worker int) ([]*string, error) {
	chJobs := make(chan map[int]*uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*string, chJobs chan map[int]*uint, errCh chan error) {
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
					chResult <- map[int]*string{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*string, len(list))
	newList := make([]*string, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*string{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*string{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUintStrPtrErrNoOrder(f func(*uint) (*string, error), list []*uint, worker int) ([]*string, error) {
	chJobs := make(chan *uint, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *string, chJobs chan *uint, errCh chan error) {
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

	newList := make([]*string, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*string{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*string{}, <-errCh
	}

	return newList, nil
}

// PMapUintBoolPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintBoolPtrErr(f func(*uint) (*bool, error), list []*uint, optional ...Optional) ([]*bool, error) {
	if f == nil {
		return []*bool{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintBoolPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUintBoolPtrErrPreserveOrder(f, list, worker)
}

func pMapUintBoolPtrErrPreserveOrder(f func(*uint) (*bool, error), list []*uint, worker int) ([]*bool, error) {
	chJobs := make(chan map[int]*uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*bool, chJobs chan map[int]*uint, errCh chan error) {
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
					chResult <- map[int]*bool{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*bool, len(list))
	newList := make([]*bool, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*bool{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*bool{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUintBoolPtrErrNoOrder(f func(*uint) (*bool, error), list []*uint, worker int) ([]*bool, error) {
	chJobs := make(chan *uint, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *bool, chJobs chan *uint, errCh chan error) {
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

	newList := make([]*bool, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*bool{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*bool{}, <-errCh
	}

	return newList, nil
}

// PMapUintFloat32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintFloat32PtrErr(f func(*uint) (*float32, error), list []*uint, optional ...Optional) ([]*float32, error) {
	if f == nil {
		return []*float32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintFloat32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUintFloat32PtrErrPreserveOrder(f, list, worker)
}

func pMapUintFloat32PtrErrPreserveOrder(f func(*uint) (*float32, error), list []*uint, worker int) ([]*float32, error) {
	chJobs := make(chan map[int]*uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float32, chJobs chan map[int]*uint, errCh chan error) {
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
					chResult <- map[int]*float32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float32, len(list))
	newList := make([]*float32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*float32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*float32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUintFloat32PtrErrNoOrder(f func(*uint) (*float32, error), list []*uint, worker int) ([]*float32, error) {
	chJobs := make(chan *uint, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float32, chJobs chan *uint, errCh chan error) {
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

	newList := make([]*float32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*float32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*float32{}, <-errCh
	}

	return newList, nil
}

// PMapUintFloat64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintFloat64PtrErr(f func(*uint) (*float64, error), list []*uint, optional ...Optional) ([]*float64, error) {
	if f == nil {
		return []*float64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintFloat64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUintFloat64PtrErrPreserveOrder(f, list, worker)
}

func pMapUintFloat64PtrErrPreserveOrder(f func(*uint) (*float64, error), list []*uint, worker int) ([]*float64, error) {
	chJobs := make(chan map[int]*uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float64, chJobs chan map[int]*uint, errCh chan error) {
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
					chResult <- map[int]*float64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float64, len(list))
	newList := make([]*float64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*float64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*float64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUintFloat64PtrErrNoOrder(f func(*uint) (*float64, error), list []*uint, worker int) ([]*float64, error) {
	chJobs := make(chan *uint, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float64, chJobs chan *uint, errCh chan error) {
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

	newList := make([]*float64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*float64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*float64{}, <-errCh
	}

	return newList, nil
}

// PMapUint64IntPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64IntPtrErr(f func(*uint64) (*int, error), list []*uint64, optional ...Optional) ([]*int, error) {
	if f == nil {
		return []*int{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64IntPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint64IntPtrErrPreserveOrder(f, list, worker)
}

func pMapUint64IntPtrErrPreserveOrder(f func(*uint64) (*int, error), list []*uint64, worker int) ([]*int, error) {
	chJobs := make(chan map[int]*uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int, chJobs chan map[int]*uint64, errCh chan error) {
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
					chResult <- map[int]*int{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int, len(list))
	newList := make([]*int, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint64IntPtrErrNoOrder(f func(*uint64) (*int, error), list []*uint64, worker int) ([]*int, error) {
	chJobs := make(chan *uint64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int, chJobs chan *uint64, errCh chan error) {
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

	newList := make([]*int, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int{}, <-errCh
	}

	return newList, nil
}

// PMapUint64Int64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Int64PtrErr(f func(*uint64) (*int64, error), list []*uint64, optional ...Optional) ([]*int64, error) {
	if f == nil {
		return []*int64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64Int64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint64Int64PtrErrPreserveOrder(f, list, worker)
}

func pMapUint64Int64PtrErrPreserveOrder(f func(*uint64) (*int64, error), list []*uint64, worker int) ([]*int64, error) {
	chJobs := make(chan map[int]*uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int64, chJobs chan map[int]*uint64, errCh chan error) {
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
					chResult <- map[int]*int64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int64, len(list))
	newList := make([]*int64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint64Int64PtrErrNoOrder(f func(*uint64) (*int64, error), list []*uint64, worker int) ([]*int64, error) {
	chJobs := make(chan *uint64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int64, chJobs chan *uint64, errCh chan error) {
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

	newList := make([]*int64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int64{}, <-errCh
	}

	return newList, nil
}

// PMapUint64Int32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Int32PtrErr(f func(*uint64) (*int32, error), list []*uint64, optional ...Optional) ([]*int32, error) {
	if f == nil {
		return []*int32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64Int32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint64Int32PtrErrPreserveOrder(f, list, worker)
}

func pMapUint64Int32PtrErrPreserveOrder(f func(*uint64) (*int32, error), list []*uint64, worker int) ([]*int32, error) {
	chJobs := make(chan map[int]*uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int32, chJobs chan map[int]*uint64, errCh chan error) {
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
					chResult <- map[int]*int32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int32, len(list))
	newList := make([]*int32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint64Int32PtrErrNoOrder(f func(*uint64) (*int32, error), list []*uint64, worker int) ([]*int32, error) {
	chJobs := make(chan *uint64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int32, chJobs chan *uint64, errCh chan error) {
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

	newList := make([]*int32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int32{}, <-errCh
	}

	return newList, nil
}

// PMapUint64Int16PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Int16PtrErr(f func(*uint64) (*int16, error), list []*uint64, optional ...Optional) ([]*int16, error) {
	if f == nil {
		return []*int16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64Int16PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint64Int16PtrErrPreserveOrder(f, list, worker)
}

func pMapUint64Int16PtrErrPreserveOrder(f func(*uint64) (*int16, error), list []*uint64, worker int) ([]*int16, error) {
	chJobs := make(chan map[int]*uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int16, chJobs chan map[int]*uint64, errCh chan error) {
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
					chResult <- map[int]*int16{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int16, len(list))
	newList := make([]*int16, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int16{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int16{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint64Int16PtrErrNoOrder(f func(*uint64) (*int16, error), list []*uint64, worker int) ([]*int16, error) {
	chJobs := make(chan *uint64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int16, chJobs chan *uint64, errCh chan error) {
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

	newList := make([]*int16, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int16{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int16{}, <-errCh
	}

	return newList, nil
}

// PMapUint64Int8PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Int8PtrErr(f func(*uint64) (*int8, error), list []*uint64, optional ...Optional) ([]*int8, error) {
	if f == nil {
		return []*int8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64Int8PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint64Int8PtrErrPreserveOrder(f, list, worker)
}

func pMapUint64Int8PtrErrPreserveOrder(f func(*uint64) (*int8, error), list []*uint64, worker int) ([]*int8, error) {
	chJobs := make(chan map[int]*uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int8, chJobs chan map[int]*uint64, errCh chan error) {
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
					chResult <- map[int]*int8{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int8, len(list))
	newList := make([]*int8, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int8{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int8{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint64Int8PtrErrNoOrder(f func(*uint64) (*int8, error), list []*uint64, worker int) ([]*int8, error) {
	chJobs := make(chan *uint64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int8, chJobs chan *uint64, errCh chan error) {
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

	newList := make([]*int8, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int8{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int8{}, <-errCh
	}

	return newList, nil
}

// PMapUint64UintPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64UintPtrErr(f func(*uint64) (*uint, error), list []*uint64, optional ...Optional) ([]*uint, error) {
	if f == nil {
		return []*uint{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64UintPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint64UintPtrErrPreserveOrder(f, list, worker)
}

func pMapUint64UintPtrErrPreserveOrder(f func(*uint64) (*uint, error), list []*uint64, worker int) ([]*uint, error) {
	chJobs := make(chan map[int]*uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint, chJobs chan map[int]*uint64, errCh chan error) {
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
					chResult <- map[int]*uint{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint, len(list))
	newList := make([]*uint, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint64UintPtrErrNoOrder(f func(*uint64) (*uint, error), list []*uint64, worker int) ([]*uint, error) {
	chJobs := make(chan *uint64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint, chJobs chan *uint64, errCh chan error) {
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

	newList := make([]*uint, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint{}, <-errCh
	}

	return newList, nil
}

// PMapUint64Uint32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Uint32PtrErr(f func(*uint64) (*uint32, error), list []*uint64, optional ...Optional) ([]*uint32, error) {
	if f == nil {
		return []*uint32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64Uint32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint64Uint32PtrErrPreserveOrder(f, list, worker)
}

func pMapUint64Uint32PtrErrPreserveOrder(f func(*uint64) (*uint32, error), list []*uint64, worker int) ([]*uint32, error) {
	chJobs := make(chan map[int]*uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint32, chJobs chan map[int]*uint64, errCh chan error) {
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
					chResult <- map[int]*uint32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint32, len(list))
	newList := make([]*uint32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint64Uint32PtrErrNoOrder(f func(*uint64) (*uint32, error), list []*uint64, worker int) ([]*uint32, error) {
	chJobs := make(chan *uint64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint32, chJobs chan *uint64, errCh chan error) {
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

	newList := make([]*uint32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint32{}, <-errCh
	}

	return newList, nil
}

// PMapUint64Uint16PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Uint16PtrErr(f func(*uint64) (*uint16, error), list []*uint64, optional ...Optional) ([]*uint16, error) {
	if f == nil {
		return []*uint16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64Uint16PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint64Uint16PtrErrPreserveOrder(f, list, worker)
}

func pMapUint64Uint16PtrErrPreserveOrder(f func(*uint64) (*uint16, error), list []*uint64, worker int) ([]*uint16, error) {
	chJobs := make(chan map[int]*uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint16, chJobs chan map[int]*uint64, errCh chan error) {
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
					chResult <- map[int]*uint16{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint16, len(list))
	newList := make([]*uint16, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint16{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint16{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint64Uint16PtrErrNoOrder(f func(*uint64) (*uint16, error), list []*uint64, worker int) ([]*uint16, error) {
	chJobs := make(chan *uint64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint16, chJobs chan *uint64, errCh chan error) {
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

	newList := make([]*uint16, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint16{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint16{}, <-errCh
	}

	return newList, nil
}

// PMapUint64Uint8PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Uint8PtrErr(f func(*uint64) (*uint8, error), list []*uint64, optional ...Optional) ([]*uint8, error) {
	if f == nil {
		return []*uint8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64Uint8PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint64Uint8PtrErrPreserveOrder(f, list, worker)
}

func pMapUint64Uint8PtrErrPreserveOrder(f func(*uint64) (*uint8, error), list []*uint64, worker int) ([]*uint8, error) {
	chJobs := make(chan map[int]*uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint8, chJobs chan map[int]*uint64, errCh chan error) {
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
					chResult <- map[int]*uint8{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint8, len(list))
	newList := make([]*uint8, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint8{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint8{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint64Uint8PtrErrNoOrder(f func(*uint64) (*uint8, error), list []*uint64, worker int) ([]*uint8, error) {
	chJobs := make(chan *uint64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint8, chJobs chan *uint64, errCh chan error) {
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

	newList := make([]*uint8, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint8{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint8{}, <-errCh
	}

	return newList, nil
}

// PMapUint64StrPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64StrPtrErr(f func(*uint64) (*string, error), list []*uint64, optional ...Optional) ([]*string, error) {
	if f == nil {
		return []*string{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64StrPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint64StrPtrErrPreserveOrder(f, list, worker)
}

func pMapUint64StrPtrErrPreserveOrder(f func(*uint64) (*string, error), list []*uint64, worker int) ([]*string, error) {
	chJobs := make(chan map[int]*uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*string, chJobs chan map[int]*uint64, errCh chan error) {
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
					chResult <- map[int]*string{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*string, len(list))
	newList := make([]*string, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*string{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*string{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint64StrPtrErrNoOrder(f func(*uint64) (*string, error), list []*uint64, worker int) ([]*string, error) {
	chJobs := make(chan *uint64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *string, chJobs chan *uint64, errCh chan error) {
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

	newList := make([]*string, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*string{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*string{}, <-errCh
	}

	return newList, nil
}

// PMapUint64BoolPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64BoolPtrErr(f func(*uint64) (*bool, error), list []*uint64, optional ...Optional) ([]*bool, error) {
	if f == nil {
		return []*bool{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64BoolPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint64BoolPtrErrPreserveOrder(f, list, worker)
}

func pMapUint64BoolPtrErrPreserveOrder(f func(*uint64) (*bool, error), list []*uint64, worker int) ([]*bool, error) {
	chJobs := make(chan map[int]*uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*bool, chJobs chan map[int]*uint64, errCh chan error) {
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
					chResult <- map[int]*bool{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*bool, len(list))
	newList := make([]*bool, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*bool{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*bool{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint64BoolPtrErrNoOrder(f func(*uint64) (*bool, error), list []*uint64, worker int) ([]*bool, error) {
	chJobs := make(chan *uint64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *bool, chJobs chan *uint64, errCh chan error) {
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

	newList := make([]*bool, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*bool{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*bool{}, <-errCh
	}

	return newList, nil
}

// PMapUint64Float32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Float32PtrErr(f func(*uint64) (*float32, error), list []*uint64, optional ...Optional) ([]*float32, error) {
	if f == nil {
		return []*float32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64Float32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint64Float32PtrErrPreserveOrder(f, list, worker)
}

func pMapUint64Float32PtrErrPreserveOrder(f func(*uint64) (*float32, error), list []*uint64, worker int) ([]*float32, error) {
	chJobs := make(chan map[int]*uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float32, chJobs chan map[int]*uint64, errCh chan error) {
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
					chResult <- map[int]*float32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float32, len(list))
	newList := make([]*float32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*float32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*float32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint64Float32PtrErrNoOrder(f func(*uint64) (*float32, error), list []*uint64, worker int) ([]*float32, error) {
	chJobs := make(chan *uint64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float32, chJobs chan *uint64, errCh chan error) {
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

	newList := make([]*float32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*float32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*float32{}, <-errCh
	}

	return newList, nil
}

// PMapUint64Float64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Float64PtrErr(f func(*uint64) (*float64, error), list []*uint64, optional ...Optional) ([]*float64, error) {
	if f == nil {
		return []*float64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64Float64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint64Float64PtrErrPreserveOrder(f, list, worker)
}

func pMapUint64Float64PtrErrPreserveOrder(f func(*uint64) (*float64, error), list []*uint64, worker int) ([]*float64, error) {
	chJobs := make(chan map[int]*uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float64, chJobs chan map[int]*uint64, errCh chan error) {
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
					chResult <- map[int]*float64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float64, len(list))
	newList := make([]*float64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*float64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*float64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint64Float64PtrErrNoOrder(f func(*uint64) (*float64, error), list []*uint64, worker int) ([]*float64, error) {
	chJobs := make(chan *uint64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float64, chJobs chan *uint64, errCh chan error) {
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

	newList := make([]*float64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*float64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*float64{}, <-errCh
	}

	return newList, nil
}

// PMapUint32IntPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32IntPtrErr(f func(*uint32) (*int, error), list []*uint32, optional ...Optional) ([]*int, error) {
	if f == nil {
		return []*int{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32IntPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint32IntPtrErrPreserveOrder(f, list, worker)
}

func pMapUint32IntPtrErrPreserveOrder(f func(*uint32) (*int, error), list []*uint32, worker int) ([]*int, error) {
	chJobs := make(chan map[int]*uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int, chJobs chan map[int]*uint32, errCh chan error) {
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
					chResult <- map[int]*int{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int, len(list))
	newList := make([]*int, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint32IntPtrErrNoOrder(f func(*uint32) (*int, error), list []*uint32, worker int) ([]*int, error) {
	chJobs := make(chan *uint32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int, chJobs chan *uint32, errCh chan error) {
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

	newList := make([]*int, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int{}, <-errCh
	}

	return newList, nil
}

// PMapUint32Int64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Int64PtrErr(f func(*uint32) (*int64, error), list []*uint32, optional ...Optional) ([]*int64, error) {
	if f == nil {
		return []*int64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32Int64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint32Int64PtrErrPreserveOrder(f, list, worker)
}

func pMapUint32Int64PtrErrPreserveOrder(f func(*uint32) (*int64, error), list []*uint32, worker int) ([]*int64, error) {
	chJobs := make(chan map[int]*uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int64, chJobs chan map[int]*uint32, errCh chan error) {
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
					chResult <- map[int]*int64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int64, len(list))
	newList := make([]*int64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint32Int64PtrErrNoOrder(f func(*uint32) (*int64, error), list []*uint32, worker int) ([]*int64, error) {
	chJobs := make(chan *uint32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int64, chJobs chan *uint32, errCh chan error) {
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

	newList := make([]*int64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int64{}, <-errCh
	}

	return newList, nil
}

// PMapUint32Int32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Int32PtrErr(f func(*uint32) (*int32, error), list []*uint32, optional ...Optional) ([]*int32, error) {
	if f == nil {
		return []*int32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32Int32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint32Int32PtrErrPreserveOrder(f, list, worker)
}

func pMapUint32Int32PtrErrPreserveOrder(f func(*uint32) (*int32, error), list []*uint32, worker int) ([]*int32, error) {
	chJobs := make(chan map[int]*uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int32, chJobs chan map[int]*uint32, errCh chan error) {
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
					chResult <- map[int]*int32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int32, len(list))
	newList := make([]*int32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint32Int32PtrErrNoOrder(f func(*uint32) (*int32, error), list []*uint32, worker int) ([]*int32, error) {
	chJobs := make(chan *uint32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int32, chJobs chan *uint32, errCh chan error) {
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

	newList := make([]*int32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int32{}, <-errCh
	}

	return newList, nil
}

// PMapUint32Int16PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Int16PtrErr(f func(*uint32) (*int16, error), list []*uint32, optional ...Optional) ([]*int16, error) {
	if f == nil {
		return []*int16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32Int16PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint32Int16PtrErrPreserveOrder(f, list, worker)
}

func pMapUint32Int16PtrErrPreserveOrder(f func(*uint32) (*int16, error), list []*uint32, worker int) ([]*int16, error) {
	chJobs := make(chan map[int]*uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int16, chJobs chan map[int]*uint32, errCh chan error) {
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
					chResult <- map[int]*int16{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int16, len(list))
	newList := make([]*int16, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int16{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int16{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint32Int16PtrErrNoOrder(f func(*uint32) (*int16, error), list []*uint32, worker int) ([]*int16, error) {
	chJobs := make(chan *uint32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int16, chJobs chan *uint32, errCh chan error) {
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

	newList := make([]*int16, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int16{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int16{}, <-errCh
	}

	return newList, nil
}

// PMapUint32Int8PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Int8PtrErr(f func(*uint32) (*int8, error), list []*uint32, optional ...Optional) ([]*int8, error) {
	if f == nil {
		return []*int8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32Int8PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint32Int8PtrErrPreserveOrder(f, list, worker)
}

func pMapUint32Int8PtrErrPreserveOrder(f func(*uint32) (*int8, error), list []*uint32, worker int) ([]*int8, error) {
	chJobs := make(chan map[int]*uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int8, chJobs chan map[int]*uint32, errCh chan error) {
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
					chResult <- map[int]*int8{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int8, len(list))
	newList := make([]*int8, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int8{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int8{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint32Int8PtrErrNoOrder(f func(*uint32) (*int8, error), list []*uint32, worker int) ([]*int8, error) {
	chJobs := make(chan *uint32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int8, chJobs chan *uint32, errCh chan error) {
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

	newList := make([]*int8, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int8{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int8{}, <-errCh
	}

	return newList, nil
}

// PMapUint32UintPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32UintPtrErr(f func(*uint32) (*uint, error), list []*uint32, optional ...Optional) ([]*uint, error) {
	if f == nil {
		return []*uint{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32UintPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint32UintPtrErrPreserveOrder(f, list, worker)
}

func pMapUint32UintPtrErrPreserveOrder(f func(*uint32) (*uint, error), list []*uint32, worker int) ([]*uint, error) {
	chJobs := make(chan map[int]*uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint, chJobs chan map[int]*uint32, errCh chan error) {
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
					chResult <- map[int]*uint{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint, len(list))
	newList := make([]*uint, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint32UintPtrErrNoOrder(f func(*uint32) (*uint, error), list []*uint32, worker int) ([]*uint, error) {
	chJobs := make(chan *uint32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint, chJobs chan *uint32, errCh chan error) {
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

	newList := make([]*uint, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint{}, <-errCh
	}

	return newList, nil
}

// PMapUint32Uint64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Uint64PtrErr(f func(*uint32) (*uint64, error), list []*uint32, optional ...Optional) ([]*uint64, error) {
	if f == nil {
		return []*uint64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32Uint64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint32Uint64PtrErrPreserveOrder(f, list, worker)
}

func pMapUint32Uint64PtrErrPreserveOrder(f func(*uint32) (*uint64, error), list []*uint32, worker int) ([]*uint64, error) {
	chJobs := make(chan map[int]*uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint64, chJobs chan map[int]*uint32, errCh chan error) {
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
					chResult <- map[int]*uint64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint64, len(list))
	newList := make([]*uint64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint32Uint64PtrErrNoOrder(f func(*uint32) (*uint64, error), list []*uint32, worker int) ([]*uint64, error) {
	chJobs := make(chan *uint32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint64, chJobs chan *uint32, errCh chan error) {
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

	newList := make([]*uint64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint64{}, <-errCh
	}

	return newList, nil
}

// PMapUint32Uint16PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Uint16PtrErr(f func(*uint32) (*uint16, error), list []*uint32, optional ...Optional) ([]*uint16, error) {
	if f == nil {
		return []*uint16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32Uint16PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint32Uint16PtrErrPreserveOrder(f, list, worker)
}

func pMapUint32Uint16PtrErrPreserveOrder(f func(*uint32) (*uint16, error), list []*uint32, worker int) ([]*uint16, error) {
	chJobs := make(chan map[int]*uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint16, chJobs chan map[int]*uint32, errCh chan error) {
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
					chResult <- map[int]*uint16{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint16, len(list))
	newList := make([]*uint16, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint16{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint16{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint32Uint16PtrErrNoOrder(f func(*uint32) (*uint16, error), list []*uint32, worker int) ([]*uint16, error) {
	chJobs := make(chan *uint32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint16, chJobs chan *uint32, errCh chan error) {
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

	newList := make([]*uint16, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint16{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint16{}, <-errCh
	}

	return newList, nil
}

// PMapUint32Uint8PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Uint8PtrErr(f func(*uint32) (*uint8, error), list []*uint32, optional ...Optional) ([]*uint8, error) {
	if f == nil {
		return []*uint8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32Uint8PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint32Uint8PtrErrPreserveOrder(f, list, worker)
}

func pMapUint32Uint8PtrErrPreserveOrder(f func(*uint32) (*uint8, error), list []*uint32, worker int) ([]*uint8, error) {
	chJobs := make(chan map[int]*uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint8, chJobs chan map[int]*uint32, errCh chan error) {
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
					chResult <- map[int]*uint8{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint8, len(list))
	newList := make([]*uint8, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint8{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint8{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint32Uint8PtrErrNoOrder(f func(*uint32) (*uint8, error), list []*uint32, worker int) ([]*uint8, error) {
	chJobs := make(chan *uint32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint8, chJobs chan *uint32, errCh chan error) {
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

	newList := make([]*uint8, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint8{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint8{}, <-errCh
	}

	return newList, nil
}

// PMapUint32StrPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32StrPtrErr(f func(*uint32) (*string, error), list []*uint32, optional ...Optional) ([]*string, error) {
	if f == nil {
		return []*string{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32StrPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint32StrPtrErrPreserveOrder(f, list, worker)
}

func pMapUint32StrPtrErrPreserveOrder(f func(*uint32) (*string, error), list []*uint32, worker int) ([]*string, error) {
	chJobs := make(chan map[int]*uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*string, chJobs chan map[int]*uint32, errCh chan error) {
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
					chResult <- map[int]*string{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*string, len(list))
	newList := make([]*string, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*string{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*string{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint32StrPtrErrNoOrder(f func(*uint32) (*string, error), list []*uint32, worker int) ([]*string, error) {
	chJobs := make(chan *uint32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *string, chJobs chan *uint32, errCh chan error) {
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

	newList := make([]*string, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*string{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*string{}, <-errCh
	}

	return newList, nil
}

// PMapUint32BoolPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32BoolPtrErr(f func(*uint32) (*bool, error), list []*uint32, optional ...Optional) ([]*bool, error) {
	if f == nil {
		return []*bool{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32BoolPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint32BoolPtrErrPreserveOrder(f, list, worker)
}

func pMapUint32BoolPtrErrPreserveOrder(f func(*uint32) (*bool, error), list []*uint32, worker int) ([]*bool, error) {
	chJobs := make(chan map[int]*uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*bool, chJobs chan map[int]*uint32, errCh chan error) {
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
					chResult <- map[int]*bool{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*bool, len(list))
	newList := make([]*bool, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*bool{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*bool{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint32BoolPtrErrNoOrder(f func(*uint32) (*bool, error), list []*uint32, worker int) ([]*bool, error) {
	chJobs := make(chan *uint32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *bool, chJobs chan *uint32, errCh chan error) {
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

	newList := make([]*bool, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*bool{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*bool{}, <-errCh
	}

	return newList, nil
}

// PMapUint32Float32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Float32PtrErr(f func(*uint32) (*float32, error), list []*uint32, optional ...Optional) ([]*float32, error) {
	if f == nil {
		return []*float32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32Float32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint32Float32PtrErrPreserveOrder(f, list, worker)
}

func pMapUint32Float32PtrErrPreserveOrder(f func(*uint32) (*float32, error), list []*uint32, worker int) ([]*float32, error) {
	chJobs := make(chan map[int]*uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float32, chJobs chan map[int]*uint32, errCh chan error) {
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
					chResult <- map[int]*float32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float32, len(list))
	newList := make([]*float32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*float32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*float32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint32Float32PtrErrNoOrder(f func(*uint32) (*float32, error), list []*uint32, worker int) ([]*float32, error) {
	chJobs := make(chan *uint32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float32, chJobs chan *uint32, errCh chan error) {
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

	newList := make([]*float32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*float32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*float32{}, <-errCh
	}

	return newList, nil
}

// PMapUint32Float64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Float64PtrErr(f func(*uint32) (*float64, error), list []*uint32, optional ...Optional) ([]*float64, error) {
	if f == nil {
		return []*float64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32Float64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint32Float64PtrErrPreserveOrder(f, list, worker)
}

func pMapUint32Float64PtrErrPreserveOrder(f func(*uint32) (*float64, error), list []*uint32, worker int) ([]*float64, error) {
	chJobs := make(chan map[int]*uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float64, chJobs chan map[int]*uint32, errCh chan error) {
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
					chResult <- map[int]*float64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float64, len(list))
	newList := make([]*float64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*float64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*float64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint32Float64PtrErrNoOrder(f func(*uint32) (*float64, error), list []*uint32, worker int) ([]*float64, error) {
	chJobs := make(chan *uint32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float64, chJobs chan *uint32, errCh chan error) {
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

	newList := make([]*float64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*float64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*float64{}, <-errCh
	}

	return newList, nil
}

// PMapUint16IntPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16IntPtrErr(f func(*uint16) (*int, error), list []*uint16, optional ...Optional) ([]*int, error) {
	if f == nil {
		return []*int{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16IntPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint16IntPtrErrPreserveOrder(f, list, worker)
}

func pMapUint16IntPtrErrPreserveOrder(f func(*uint16) (*int, error), list []*uint16, worker int) ([]*int, error) {
	chJobs := make(chan map[int]*uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int, chJobs chan map[int]*uint16, errCh chan error) {
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
					chResult <- map[int]*int{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int, len(list))
	newList := make([]*int, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint16IntPtrErrNoOrder(f func(*uint16) (*int, error), list []*uint16, worker int) ([]*int, error) {
	chJobs := make(chan *uint16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int, chJobs chan *uint16, errCh chan error) {
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

	newList := make([]*int, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int{}, <-errCh
	}

	return newList, nil
}

// PMapUint16Int64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Int64PtrErr(f func(*uint16) (*int64, error), list []*uint16, optional ...Optional) ([]*int64, error) {
	if f == nil {
		return []*int64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16Int64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint16Int64PtrErrPreserveOrder(f, list, worker)
}

func pMapUint16Int64PtrErrPreserveOrder(f func(*uint16) (*int64, error), list []*uint16, worker int) ([]*int64, error) {
	chJobs := make(chan map[int]*uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int64, chJobs chan map[int]*uint16, errCh chan error) {
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
					chResult <- map[int]*int64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int64, len(list))
	newList := make([]*int64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint16Int64PtrErrNoOrder(f func(*uint16) (*int64, error), list []*uint16, worker int) ([]*int64, error) {
	chJobs := make(chan *uint16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int64, chJobs chan *uint16, errCh chan error) {
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

	newList := make([]*int64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int64{}, <-errCh
	}

	return newList, nil
}

// PMapUint16Int32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Int32PtrErr(f func(*uint16) (*int32, error), list []*uint16, optional ...Optional) ([]*int32, error) {
	if f == nil {
		return []*int32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16Int32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint16Int32PtrErrPreserveOrder(f, list, worker)
}

func pMapUint16Int32PtrErrPreserveOrder(f func(*uint16) (*int32, error), list []*uint16, worker int) ([]*int32, error) {
	chJobs := make(chan map[int]*uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int32, chJobs chan map[int]*uint16, errCh chan error) {
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
					chResult <- map[int]*int32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int32, len(list))
	newList := make([]*int32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint16Int32PtrErrNoOrder(f func(*uint16) (*int32, error), list []*uint16, worker int) ([]*int32, error) {
	chJobs := make(chan *uint16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int32, chJobs chan *uint16, errCh chan error) {
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

	newList := make([]*int32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int32{}, <-errCh
	}

	return newList, nil
}

// PMapUint16Int16PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Int16PtrErr(f func(*uint16) (*int16, error), list []*uint16, optional ...Optional) ([]*int16, error) {
	if f == nil {
		return []*int16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16Int16PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint16Int16PtrErrPreserveOrder(f, list, worker)
}

func pMapUint16Int16PtrErrPreserveOrder(f func(*uint16) (*int16, error), list []*uint16, worker int) ([]*int16, error) {
	chJobs := make(chan map[int]*uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int16, chJobs chan map[int]*uint16, errCh chan error) {
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
					chResult <- map[int]*int16{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int16, len(list))
	newList := make([]*int16, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int16{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int16{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint16Int16PtrErrNoOrder(f func(*uint16) (*int16, error), list []*uint16, worker int) ([]*int16, error) {
	chJobs := make(chan *uint16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int16, chJobs chan *uint16, errCh chan error) {
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

	newList := make([]*int16, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int16{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int16{}, <-errCh
	}

	return newList, nil
}

// PMapUint16Int8PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Int8PtrErr(f func(*uint16) (*int8, error), list []*uint16, optional ...Optional) ([]*int8, error) {
	if f == nil {
		return []*int8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16Int8PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint16Int8PtrErrPreserveOrder(f, list, worker)
}

func pMapUint16Int8PtrErrPreserveOrder(f func(*uint16) (*int8, error), list []*uint16, worker int) ([]*int8, error) {
	chJobs := make(chan map[int]*uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int8, chJobs chan map[int]*uint16, errCh chan error) {
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
					chResult <- map[int]*int8{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int8, len(list))
	newList := make([]*int8, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int8{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int8{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint16Int8PtrErrNoOrder(f func(*uint16) (*int8, error), list []*uint16, worker int) ([]*int8, error) {
	chJobs := make(chan *uint16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int8, chJobs chan *uint16, errCh chan error) {
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

	newList := make([]*int8, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int8{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int8{}, <-errCh
	}

	return newList, nil
}

// PMapUint16UintPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16UintPtrErr(f func(*uint16) (*uint, error), list []*uint16, optional ...Optional) ([]*uint, error) {
	if f == nil {
		return []*uint{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16UintPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint16UintPtrErrPreserveOrder(f, list, worker)
}

func pMapUint16UintPtrErrPreserveOrder(f func(*uint16) (*uint, error), list []*uint16, worker int) ([]*uint, error) {
	chJobs := make(chan map[int]*uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint, chJobs chan map[int]*uint16, errCh chan error) {
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
					chResult <- map[int]*uint{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint, len(list))
	newList := make([]*uint, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint16UintPtrErrNoOrder(f func(*uint16) (*uint, error), list []*uint16, worker int) ([]*uint, error) {
	chJobs := make(chan *uint16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint, chJobs chan *uint16, errCh chan error) {
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

	newList := make([]*uint, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint{}, <-errCh
	}

	return newList, nil
}

// PMapUint16Uint64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Uint64PtrErr(f func(*uint16) (*uint64, error), list []*uint16, optional ...Optional) ([]*uint64, error) {
	if f == nil {
		return []*uint64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16Uint64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint16Uint64PtrErrPreserveOrder(f, list, worker)
}

func pMapUint16Uint64PtrErrPreserveOrder(f func(*uint16) (*uint64, error), list []*uint16, worker int) ([]*uint64, error) {
	chJobs := make(chan map[int]*uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint64, chJobs chan map[int]*uint16, errCh chan error) {
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
					chResult <- map[int]*uint64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint64, len(list))
	newList := make([]*uint64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint16Uint64PtrErrNoOrder(f func(*uint16) (*uint64, error), list []*uint16, worker int) ([]*uint64, error) {
	chJobs := make(chan *uint16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint64, chJobs chan *uint16, errCh chan error) {
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

	newList := make([]*uint64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint64{}, <-errCh
	}

	return newList, nil
}

// PMapUint16Uint32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Uint32PtrErr(f func(*uint16) (*uint32, error), list []*uint16, optional ...Optional) ([]*uint32, error) {
	if f == nil {
		return []*uint32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16Uint32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint16Uint32PtrErrPreserveOrder(f, list, worker)
}

func pMapUint16Uint32PtrErrPreserveOrder(f func(*uint16) (*uint32, error), list []*uint16, worker int) ([]*uint32, error) {
	chJobs := make(chan map[int]*uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint32, chJobs chan map[int]*uint16, errCh chan error) {
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
					chResult <- map[int]*uint32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint32, len(list))
	newList := make([]*uint32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint16Uint32PtrErrNoOrder(f func(*uint16) (*uint32, error), list []*uint16, worker int) ([]*uint32, error) {
	chJobs := make(chan *uint16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint32, chJobs chan *uint16, errCh chan error) {
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

	newList := make([]*uint32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint32{}, <-errCh
	}

	return newList, nil
}

// PMapUint16Uint8PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Uint8PtrErr(f func(*uint16) (*uint8, error), list []*uint16, optional ...Optional) ([]*uint8, error) {
	if f == nil {
		return []*uint8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16Uint8PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint16Uint8PtrErrPreserveOrder(f, list, worker)
}

func pMapUint16Uint8PtrErrPreserveOrder(f func(*uint16) (*uint8, error), list []*uint16, worker int) ([]*uint8, error) {
	chJobs := make(chan map[int]*uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint8, chJobs chan map[int]*uint16, errCh chan error) {
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
					chResult <- map[int]*uint8{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint8, len(list))
	newList := make([]*uint8, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint8{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint8{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint16Uint8PtrErrNoOrder(f func(*uint16) (*uint8, error), list []*uint16, worker int) ([]*uint8, error) {
	chJobs := make(chan *uint16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint8, chJobs chan *uint16, errCh chan error) {
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

	newList := make([]*uint8, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint8{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint8{}, <-errCh
	}

	return newList, nil
}

// PMapUint16StrPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16StrPtrErr(f func(*uint16) (*string, error), list []*uint16, optional ...Optional) ([]*string, error) {
	if f == nil {
		return []*string{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16StrPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint16StrPtrErrPreserveOrder(f, list, worker)
}

func pMapUint16StrPtrErrPreserveOrder(f func(*uint16) (*string, error), list []*uint16, worker int) ([]*string, error) {
	chJobs := make(chan map[int]*uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*string, chJobs chan map[int]*uint16, errCh chan error) {
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
					chResult <- map[int]*string{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*string, len(list))
	newList := make([]*string, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*string{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*string{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint16StrPtrErrNoOrder(f func(*uint16) (*string, error), list []*uint16, worker int) ([]*string, error) {
	chJobs := make(chan *uint16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *string, chJobs chan *uint16, errCh chan error) {
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

	newList := make([]*string, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*string{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*string{}, <-errCh
	}

	return newList, nil
}

// PMapUint16BoolPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16BoolPtrErr(f func(*uint16) (*bool, error), list []*uint16, optional ...Optional) ([]*bool, error) {
	if f == nil {
		return []*bool{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16BoolPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint16BoolPtrErrPreserveOrder(f, list, worker)
}

func pMapUint16BoolPtrErrPreserveOrder(f func(*uint16) (*bool, error), list []*uint16, worker int) ([]*bool, error) {
	chJobs := make(chan map[int]*uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*bool, chJobs chan map[int]*uint16, errCh chan error) {
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
					chResult <- map[int]*bool{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*bool, len(list))
	newList := make([]*bool, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*bool{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*bool{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint16BoolPtrErrNoOrder(f func(*uint16) (*bool, error), list []*uint16, worker int) ([]*bool, error) {
	chJobs := make(chan *uint16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *bool, chJobs chan *uint16, errCh chan error) {
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

	newList := make([]*bool, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*bool{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*bool{}, <-errCh
	}

	return newList, nil
}

// PMapUint16Float32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Float32PtrErr(f func(*uint16) (*float32, error), list []*uint16, optional ...Optional) ([]*float32, error) {
	if f == nil {
		return []*float32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16Float32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint16Float32PtrErrPreserveOrder(f, list, worker)
}

func pMapUint16Float32PtrErrPreserveOrder(f func(*uint16) (*float32, error), list []*uint16, worker int) ([]*float32, error) {
	chJobs := make(chan map[int]*uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float32, chJobs chan map[int]*uint16, errCh chan error) {
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
					chResult <- map[int]*float32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float32, len(list))
	newList := make([]*float32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*float32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*float32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint16Float32PtrErrNoOrder(f func(*uint16) (*float32, error), list []*uint16, worker int) ([]*float32, error) {
	chJobs := make(chan *uint16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float32, chJobs chan *uint16, errCh chan error) {
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

	newList := make([]*float32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*float32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*float32{}, <-errCh
	}

	return newList, nil
}

// PMapUint16Float64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Float64PtrErr(f func(*uint16) (*float64, error), list []*uint16, optional ...Optional) ([]*float64, error) {
	if f == nil {
		return []*float64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16Float64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint16Float64PtrErrPreserveOrder(f, list, worker)
}

func pMapUint16Float64PtrErrPreserveOrder(f func(*uint16) (*float64, error), list []*uint16, worker int) ([]*float64, error) {
	chJobs := make(chan map[int]*uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float64, chJobs chan map[int]*uint16, errCh chan error) {
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
					chResult <- map[int]*float64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float64, len(list))
	newList := make([]*float64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*float64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*float64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint16Float64PtrErrNoOrder(f func(*uint16) (*float64, error), list []*uint16, worker int) ([]*float64, error) {
	chJobs := make(chan *uint16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float64, chJobs chan *uint16, errCh chan error) {
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

	newList := make([]*float64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*float64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*float64{}, <-errCh
	}

	return newList, nil
}

// PMapUint8IntPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8IntPtrErr(f func(*uint8) (*int, error), list []*uint8, optional ...Optional) ([]*int, error) {
	if f == nil {
		return []*int{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8IntPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint8IntPtrErrPreserveOrder(f, list, worker)
}

func pMapUint8IntPtrErrPreserveOrder(f func(*uint8) (*int, error), list []*uint8, worker int) ([]*int, error) {
	chJobs := make(chan map[int]*uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int, chJobs chan map[int]*uint8, errCh chan error) {
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
					chResult <- map[int]*int{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int, len(list))
	newList := make([]*int, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint8IntPtrErrNoOrder(f func(*uint8) (*int, error), list []*uint8, worker int) ([]*int, error) {
	chJobs := make(chan *uint8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int, chJobs chan *uint8, errCh chan error) {
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

	newList := make([]*int, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int{}, <-errCh
	}

	return newList, nil
}

// PMapUint8Int64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Int64PtrErr(f func(*uint8) (*int64, error), list []*uint8, optional ...Optional) ([]*int64, error) {
	if f == nil {
		return []*int64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8Int64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint8Int64PtrErrPreserveOrder(f, list, worker)
}

func pMapUint8Int64PtrErrPreserveOrder(f func(*uint8) (*int64, error), list []*uint8, worker int) ([]*int64, error) {
	chJobs := make(chan map[int]*uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int64, chJobs chan map[int]*uint8, errCh chan error) {
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
					chResult <- map[int]*int64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int64, len(list))
	newList := make([]*int64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint8Int64PtrErrNoOrder(f func(*uint8) (*int64, error), list []*uint8, worker int) ([]*int64, error) {
	chJobs := make(chan *uint8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int64, chJobs chan *uint8, errCh chan error) {
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

	newList := make([]*int64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int64{}, <-errCh
	}

	return newList, nil
}

// PMapUint8Int32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Int32PtrErr(f func(*uint8) (*int32, error), list []*uint8, optional ...Optional) ([]*int32, error) {
	if f == nil {
		return []*int32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8Int32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint8Int32PtrErrPreserveOrder(f, list, worker)
}

func pMapUint8Int32PtrErrPreserveOrder(f func(*uint8) (*int32, error), list []*uint8, worker int) ([]*int32, error) {
	chJobs := make(chan map[int]*uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int32, chJobs chan map[int]*uint8, errCh chan error) {
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
					chResult <- map[int]*int32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int32, len(list))
	newList := make([]*int32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint8Int32PtrErrNoOrder(f func(*uint8) (*int32, error), list []*uint8, worker int) ([]*int32, error) {
	chJobs := make(chan *uint8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int32, chJobs chan *uint8, errCh chan error) {
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

	newList := make([]*int32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int32{}, <-errCh
	}

	return newList, nil
}

// PMapUint8Int16PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Int16PtrErr(f func(*uint8) (*int16, error), list []*uint8, optional ...Optional) ([]*int16, error) {
	if f == nil {
		return []*int16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8Int16PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint8Int16PtrErrPreserveOrder(f, list, worker)
}

func pMapUint8Int16PtrErrPreserveOrder(f func(*uint8) (*int16, error), list []*uint8, worker int) ([]*int16, error) {
	chJobs := make(chan map[int]*uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int16, chJobs chan map[int]*uint8, errCh chan error) {
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
					chResult <- map[int]*int16{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int16, len(list))
	newList := make([]*int16, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int16{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int16{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint8Int16PtrErrNoOrder(f func(*uint8) (*int16, error), list []*uint8, worker int) ([]*int16, error) {
	chJobs := make(chan *uint8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int16, chJobs chan *uint8, errCh chan error) {
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

	newList := make([]*int16, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int16{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int16{}, <-errCh
	}

	return newList, nil
}

// PMapUint8Int8PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Int8PtrErr(f func(*uint8) (*int8, error), list []*uint8, optional ...Optional) ([]*int8, error) {
	if f == nil {
		return []*int8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8Int8PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint8Int8PtrErrPreserveOrder(f, list, worker)
}

func pMapUint8Int8PtrErrPreserveOrder(f func(*uint8) (*int8, error), list []*uint8, worker int) ([]*int8, error) {
	chJobs := make(chan map[int]*uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int8, chJobs chan map[int]*uint8, errCh chan error) {
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
					chResult <- map[int]*int8{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int8, len(list))
	newList := make([]*int8, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int8{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int8{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint8Int8PtrErrNoOrder(f func(*uint8) (*int8, error), list []*uint8, worker int) ([]*int8, error) {
	chJobs := make(chan *uint8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int8, chJobs chan *uint8, errCh chan error) {
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

	newList := make([]*int8, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int8{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int8{}, <-errCh
	}

	return newList, nil
}

// PMapUint8UintPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8UintPtrErr(f func(*uint8) (*uint, error), list []*uint8, optional ...Optional) ([]*uint, error) {
	if f == nil {
		return []*uint{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8UintPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint8UintPtrErrPreserveOrder(f, list, worker)
}

func pMapUint8UintPtrErrPreserveOrder(f func(*uint8) (*uint, error), list []*uint8, worker int) ([]*uint, error) {
	chJobs := make(chan map[int]*uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint, chJobs chan map[int]*uint8, errCh chan error) {
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
					chResult <- map[int]*uint{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint, len(list))
	newList := make([]*uint, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint8UintPtrErrNoOrder(f func(*uint8) (*uint, error), list []*uint8, worker int) ([]*uint, error) {
	chJobs := make(chan *uint8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint, chJobs chan *uint8, errCh chan error) {
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

	newList := make([]*uint, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint{}, <-errCh
	}

	return newList, nil
}

// PMapUint8Uint64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Uint64PtrErr(f func(*uint8) (*uint64, error), list []*uint8, optional ...Optional) ([]*uint64, error) {
	if f == nil {
		return []*uint64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8Uint64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint8Uint64PtrErrPreserveOrder(f, list, worker)
}

func pMapUint8Uint64PtrErrPreserveOrder(f func(*uint8) (*uint64, error), list []*uint8, worker int) ([]*uint64, error) {
	chJobs := make(chan map[int]*uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint64, chJobs chan map[int]*uint8, errCh chan error) {
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
					chResult <- map[int]*uint64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint64, len(list))
	newList := make([]*uint64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint8Uint64PtrErrNoOrder(f func(*uint8) (*uint64, error), list []*uint8, worker int) ([]*uint64, error) {
	chJobs := make(chan *uint8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint64, chJobs chan *uint8, errCh chan error) {
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

	newList := make([]*uint64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint64{}, <-errCh
	}

	return newList, nil
}

// PMapUint8Uint32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Uint32PtrErr(f func(*uint8) (*uint32, error), list []*uint8, optional ...Optional) ([]*uint32, error) {
	if f == nil {
		return []*uint32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8Uint32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint8Uint32PtrErrPreserveOrder(f, list, worker)
}

func pMapUint8Uint32PtrErrPreserveOrder(f func(*uint8) (*uint32, error), list []*uint8, worker int) ([]*uint32, error) {
	chJobs := make(chan map[int]*uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint32, chJobs chan map[int]*uint8, errCh chan error) {
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
					chResult <- map[int]*uint32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint32, len(list))
	newList := make([]*uint32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint8Uint32PtrErrNoOrder(f func(*uint8) (*uint32, error), list []*uint8, worker int) ([]*uint32, error) {
	chJobs := make(chan *uint8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint32, chJobs chan *uint8, errCh chan error) {
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

	newList := make([]*uint32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint32{}, <-errCh
	}

	return newList, nil
}

// PMapUint8Uint16PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Uint16PtrErr(f func(*uint8) (*uint16, error), list []*uint8, optional ...Optional) ([]*uint16, error) {
	if f == nil {
		return []*uint16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8Uint16PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint8Uint16PtrErrPreserveOrder(f, list, worker)
}

func pMapUint8Uint16PtrErrPreserveOrder(f func(*uint8) (*uint16, error), list []*uint8, worker int) ([]*uint16, error) {
	chJobs := make(chan map[int]*uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint16, chJobs chan map[int]*uint8, errCh chan error) {
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
					chResult <- map[int]*uint16{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint16, len(list))
	newList := make([]*uint16, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint16{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint16{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint8Uint16PtrErrNoOrder(f func(*uint8) (*uint16, error), list []*uint8, worker int) ([]*uint16, error) {
	chJobs := make(chan *uint8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint16, chJobs chan *uint8, errCh chan error) {
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

	newList := make([]*uint16, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint16{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint16{}, <-errCh
	}

	return newList, nil
}

// PMapUint8StrPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8StrPtrErr(f func(*uint8) (*string, error), list []*uint8, optional ...Optional) ([]*string, error) {
	if f == nil {
		return []*string{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8StrPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint8StrPtrErrPreserveOrder(f, list, worker)
}

func pMapUint8StrPtrErrPreserveOrder(f func(*uint8) (*string, error), list []*uint8, worker int) ([]*string, error) {
	chJobs := make(chan map[int]*uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*string, chJobs chan map[int]*uint8, errCh chan error) {
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
					chResult <- map[int]*string{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*string, len(list))
	newList := make([]*string, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*string{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*string{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint8StrPtrErrNoOrder(f func(*uint8) (*string, error), list []*uint8, worker int) ([]*string, error) {
	chJobs := make(chan *uint8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *string, chJobs chan *uint8, errCh chan error) {
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

	newList := make([]*string, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*string{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*string{}, <-errCh
	}

	return newList, nil
}

// PMapUint8BoolPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8BoolPtrErr(f func(*uint8) (*bool, error), list []*uint8, optional ...Optional) ([]*bool, error) {
	if f == nil {
		return []*bool{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8BoolPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint8BoolPtrErrPreserveOrder(f, list, worker)
}

func pMapUint8BoolPtrErrPreserveOrder(f func(*uint8) (*bool, error), list []*uint8, worker int) ([]*bool, error) {
	chJobs := make(chan map[int]*uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*bool, chJobs chan map[int]*uint8, errCh chan error) {
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
					chResult <- map[int]*bool{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*bool, len(list))
	newList := make([]*bool, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*bool{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*bool{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint8BoolPtrErrNoOrder(f func(*uint8) (*bool, error), list []*uint8, worker int) ([]*bool, error) {
	chJobs := make(chan *uint8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *bool, chJobs chan *uint8, errCh chan error) {
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

	newList := make([]*bool, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*bool{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*bool{}, <-errCh
	}

	return newList, nil
}

// PMapUint8Float32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Float32PtrErr(f func(*uint8) (*float32, error), list []*uint8, optional ...Optional) ([]*float32, error) {
	if f == nil {
		return []*float32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8Float32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint8Float32PtrErrPreserveOrder(f, list, worker)
}

func pMapUint8Float32PtrErrPreserveOrder(f func(*uint8) (*float32, error), list []*uint8, worker int) ([]*float32, error) {
	chJobs := make(chan map[int]*uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float32, chJobs chan map[int]*uint8, errCh chan error) {
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
					chResult <- map[int]*float32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float32, len(list))
	newList := make([]*float32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*float32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*float32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint8Float32PtrErrNoOrder(f func(*uint8) (*float32, error), list []*uint8, worker int) ([]*float32, error) {
	chJobs := make(chan *uint8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float32, chJobs chan *uint8, errCh chan error) {
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

	newList := make([]*float32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*float32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*float32{}, <-errCh
	}

	return newList, nil
}

// PMapUint8Float64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Float64PtrErr(f func(*uint8) (*float64, error), list []*uint8, optional ...Optional) ([]*float64, error) {
	if f == nil {
		return []*float64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8Float64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapUint8Float64PtrErrPreserveOrder(f, list, worker)
}

func pMapUint8Float64PtrErrPreserveOrder(f func(*uint8) (*float64, error), list []*uint8, worker int) ([]*float64, error) {
	chJobs := make(chan map[int]*uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float64, chJobs chan map[int]*uint8, errCh chan error) {
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
					chResult <- map[int]*float64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float64, len(list))
	newList := make([]*float64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*float64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*float64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapUint8Float64PtrErrNoOrder(f func(*uint8) (*float64, error), list []*uint8, worker int) ([]*float64, error) {
	chJobs := make(chan *uint8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float64, chJobs chan *uint8, errCh chan error) {
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

	newList := make([]*float64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*float64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*float64{}, <-errCh
	}

	return newList, nil
}

// PMapStrIntPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrIntPtrErr(f func(*string) (*int, error), list []*string, optional ...Optional) ([]*int, error) {
	if f == nil {
		return []*int{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrIntPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapStrIntPtrErrPreserveOrder(f, list, worker)
}

func pMapStrIntPtrErrPreserveOrder(f func(*string) (*int, error), list []*string, worker int) ([]*int, error) {
	chJobs := make(chan map[int]*string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int, chJobs chan map[int]*string, errCh chan error) {
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
					chResult <- map[int]*int{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int, len(list))
	newList := make([]*int, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapStrIntPtrErrNoOrder(f func(*string) (*int, error), list []*string, worker int) ([]*int, error) {
	chJobs := make(chan *string, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int, chJobs chan *string, errCh chan error) {
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

	newList := make([]*int, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int{}, <-errCh
	}

	return newList, nil
}

// PMapStrInt64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrInt64PtrErr(f func(*string) (*int64, error), list []*string, optional ...Optional) ([]*int64, error) {
	if f == nil {
		return []*int64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrInt64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapStrInt64PtrErrPreserveOrder(f, list, worker)
}

func pMapStrInt64PtrErrPreserveOrder(f func(*string) (*int64, error), list []*string, worker int) ([]*int64, error) {
	chJobs := make(chan map[int]*string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int64, chJobs chan map[int]*string, errCh chan error) {
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
					chResult <- map[int]*int64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int64, len(list))
	newList := make([]*int64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapStrInt64PtrErrNoOrder(f func(*string) (*int64, error), list []*string, worker int) ([]*int64, error) {
	chJobs := make(chan *string, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int64, chJobs chan *string, errCh chan error) {
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

	newList := make([]*int64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int64{}, <-errCh
	}

	return newList, nil
}

// PMapStrInt32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrInt32PtrErr(f func(*string) (*int32, error), list []*string, optional ...Optional) ([]*int32, error) {
	if f == nil {
		return []*int32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrInt32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapStrInt32PtrErrPreserveOrder(f, list, worker)
}

func pMapStrInt32PtrErrPreserveOrder(f func(*string) (*int32, error), list []*string, worker int) ([]*int32, error) {
	chJobs := make(chan map[int]*string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int32, chJobs chan map[int]*string, errCh chan error) {
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
					chResult <- map[int]*int32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int32, len(list))
	newList := make([]*int32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapStrInt32PtrErrNoOrder(f func(*string) (*int32, error), list []*string, worker int) ([]*int32, error) {
	chJobs := make(chan *string, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int32, chJobs chan *string, errCh chan error) {
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

	newList := make([]*int32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int32{}, <-errCh
	}

	return newList, nil
}

// PMapStrInt16PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrInt16PtrErr(f func(*string) (*int16, error), list []*string, optional ...Optional) ([]*int16, error) {
	if f == nil {
		return []*int16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrInt16PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapStrInt16PtrErrPreserveOrder(f, list, worker)
}

func pMapStrInt16PtrErrPreserveOrder(f func(*string) (*int16, error), list []*string, worker int) ([]*int16, error) {
	chJobs := make(chan map[int]*string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int16, chJobs chan map[int]*string, errCh chan error) {
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
					chResult <- map[int]*int16{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int16, len(list))
	newList := make([]*int16, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int16{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int16{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapStrInt16PtrErrNoOrder(f func(*string) (*int16, error), list []*string, worker int) ([]*int16, error) {
	chJobs := make(chan *string, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int16, chJobs chan *string, errCh chan error) {
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

	newList := make([]*int16, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int16{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int16{}, <-errCh
	}

	return newList, nil
}

// PMapStrInt8PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrInt8PtrErr(f func(*string) (*int8, error), list []*string, optional ...Optional) ([]*int8, error) {
	if f == nil {
		return []*int8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrInt8PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapStrInt8PtrErrPreserveOrder(f, list, worker)
}

func pMapStrInt8PtrErrPreserveOrder(f func(*string) (*int8, error), list []*string, worker int) ([]*int8, error) {
	chJobs := make(chan map[int]*string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int8, chJobs chan map[int]*string, errCh chan error) {
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
					chResult <- map[int]*int8{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int8, len(list))
	newList := make([]*int8, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int8{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int8{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapStrInt8PtrErrNoOrder(f func(*string) (*int8, error), list []*string, worker int) ([]*int8, error) {
	chJobs := make(chan *string, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int8, chJobs chan *string, errCh chan error) {
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

	newList := make([]*int8, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int8{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int8{}, <-errCh
	}

	return newList, nil
}

// PMapStrUintPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrUintPtrErr(f func(*string) (*uint, error), list []*string, optional ...Optional) ([]*uint, error) {
	if f == nil {
		return []*uint{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrUintPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapStrUintPtrErrPreserveOrder(f, list, worker)
}

func pMapStrUintPtrErrPreserveOrder(f func(*string) (*uint, error), list []*string, worker int) ([]*uint, error) {
	chJobs := make(chan map[int]*string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint, chJobs chan map[int]*string, errCh chan error) {
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
					chResult <- map[int]*uint{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint, len(list))
	newList := make([]*uint, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapStrUintPtrErrNoOrder(f func(*string) (*uint, error), list []*string, worker int) ([]*uint, error) {
	chJobs := make(chan *string, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint, chJobs chan *string, errCh chan error) {
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

	newList := make([]*uint, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint{}, <-errCh
	}

	return newList, nil
}

// PMapStrUint64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrUint64PtrErr(f func(*string) (*uint64, error), list []*string, optional ...Optional) ([]*uint64, error) {
	if f == nil {
		return []*uint64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrUint64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapStrUint64PtrErrPreserveOrder(f, list, worker)
}

func pMapStrUint64PtrErrPreserveOrder(f func(*string) (*uint64, error), list []*string, worker int) ([]*uint64, error) {
	chJobs := make(chan map[int]*string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint64, chJobs chan map[int]*string, errCh chan error) {
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
					chResult <- map[int]*uint64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint64, len(list))
	newList := make([]*uint64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapStrUint64PtrErrNoOrder(f func(*string) (*uint64, error), list []*string, worker int) ([]*uint64, error) {
	chJobs := make(chan *string, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint64, chJobs chan *string, errCh chan error) {
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

	newList := make([]*uint64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint64{}, <-errCh
	}

	return newList, nil
}

// PMapStrUint32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrUint32PtrErr(f func(*string) (*uint32, error), list []*string, optional ...Optional) ([]*uint32, error) {
	if f == nil {
		return []*uint32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrUint32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapStrUint32PtrErrPreserveOrder(f, list, worker)
}

func pMapStrUint32PtrErrPreserveOrder(f func(*string) (*uint32, error), list []*string, worker int) ([]*uint32, error) {
	chJobs := make(chan map[int]*string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint32, chJobs chan map[int]*string, errCh chan error) {
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
					chResult <- map[int]*uint32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint32, len(list))
	newList := make([]*uint32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapStrUint32PtrErrNoOrder(f func(*string) (*uint32, error), list []*string, worker int) ([]*uint32, error) {
	chJobs := make(chan *string, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint32, chJobs chan *string, errCh chan error) {
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

	newList := make([]*uint32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint32{}, <-errCh
	}

	return newList, nil
}

// PMapStrUint16PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrUint16PtrErr(f func(*string) (*uint16, error), list []*string, optional ...Optional) ([]*uint16, error) {
	if f == nil {
		return []*uint16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrUint16PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapStrUint16PtrErrPreserveOrder(f, list, worker)
}

func pMapStrUint16PtrErrPreserveOrder(f func(*string) (*uint16, error), list []*string, worker int) ([]*uint16, error) {
	chJobs := make(chan map[int]*string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint16, chJobs chan map[int]*string, errCh chan error) {
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
					chResult <- map[int]*uint16{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint16, len(list))
	newList := make([]*uint16, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint16{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint16{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapStrUint16PtrErrNoOrder(f func(*string) (*uint16, error), list []*string, worker int) ([]*uint16, error) {
	chJobs := make(chan *string, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint16, chJobs chan *string, errCh chan error) {
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

	newList := make([]*uint16, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint16{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint16{}, <-errCh
	}

	return newList, nil
}

// PMapStrUint8PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrUint8PtrErr(f func(*string) (*uint8, error), list []*string, optional ...Optional) ([]*uint8, error) {
	if f == nil {
		return []*uint8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrUint8PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapStrUint8PtrErrPreserveOrder(f, list, worker)
}

func pMapStrUint8PtrErrPreserveOrder(f func(*string) (*uint8, error), list []*string, worker int) ([]*uint8, error) {
	chJobs := make(chan map[int]*string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint8, chJobs chan map[int]*string, errCh chan error) {
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
					chResult <- map[int]*uint8{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint8, len(list))
	newList := make([]*uint8, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint8{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint8{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapStrUint8PtrErrNoOrder(f func(*string) (*uint8, error), list []*string, worker int) ([]*uint8, error) {
	chJobs := make(chan *string, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint8, chJobs chan *string, errCh chan error) {
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

	newList := make([]*uint8, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint8{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint8{}, <-errCh
	}

	return newList, nil
}

// PMapStrBoolPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrBoolPtrErr(f func(*string) (*bool, error), list []*string, optional ...Optional) ([]*bool, error) {
	if f == nil {
		return []*bool{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrBoolPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapStrBoolPtrErrPreserveOrder(f, list, worker)
}

func pMapStrBoolPtrErrPreserveOrder(f func(*string) (*bool, error), list []*string, worker int) ([]*bool, error) {
	chJobs := make(chan map[int]*string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*bool, chJobs chan map[int]*string, errCh chan error) {
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
					chResult <- map[int]*bool{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*bool, len(list))
	newList := make([]*bool, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*bool{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*bool{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapStrBoolPtrErrNoOrder(f func(*string) (*bool, error), list []*string, worker int) ([]*bool, error) {
	chJobs := make(chan *string, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *bool, chJobs chan *string, errCh chan error) {
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

	newList := make([]*bool, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*bool{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*bool{}, <-errCh
	}

	return newList, nil
}

// PMapStrFloat32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrFloat32PtrErr(f func(*string) (*float32, error), list []*string, optional ...Optional) ([]*float32, error) {
	if f == nil {
		return []*float32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrFloat32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapStrFloat32PtrErrPreserveOrder(f, list, worker)
}

func pMapStrFloat32PtrErrPreserveOrder(f func(*string) (*float32, error), list []*string, worker int) ([]*float32, error) {
	chJobs := make(chan map[int]*string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float32, chJobs chan map[int]*string, errCh chan error) {
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
					chResult <- map[int]*float32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float32, len(list))
	newList := make([]*float32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*float32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*float32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapStrFloat32PtrErrNoOrder(f func(*string) (*float32, error), list []*string, worker int) ([]*float32, error) {
	chJobs := make(chan *string, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float32, chJobs chan *string, errCh chan error) {
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

	newList := make([]*float32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*float32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*float32{}, <-errCh
	}

	return newList, nil
}

// PMapStrFloat64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrFloat64PtrErr(f func(*string) (*float64, error), list []*string, optional ...Optional) ([]*float64, error) {
	if f == nil {
		return []*float64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrFloat64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapStrFloat64PtrErrPreserveOrder(f, list, worker)
}

func pMapStrFloat64PtrErrPreserveOrder(f func(*string) (*float64, error), list []*string, worker int) ([]*float64, error) {
	chJobs := make(chan map[int]*string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float64, chJobs chan map[int]*string, errCh chan error) {
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
					chResult <- map[int]*float64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float64, len(list))
	newList := make([]*float64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*float64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*float64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapStrFloat64PtrErrNoOrder(f func(*string) (*float64, error), list []*string, worker int) ([]*float64, error) {
	chJobs := make(chan *string, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float64, chJobs chan *string, errCh chan error) {
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

	newList := make([]*float64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*float64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*float64{}, <-errCh
	}

	return newList, nil
}

// PMapBoolIntPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolIntPtrErr(f func(*bool) (*int, error), list []*bool, optional ...Optional) ([]*int, error) {
	if f == nil {
		return []*int{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolIntPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapBoolIntPtrErrPreserveOrder(f, list, worker)
}

func pMapBoolIntPtrErrPreserveOrder(f func(*bool) (*int, error), list []*bool, worker int) ([]*int, error) {
	chJobs := make(chan map[int]*bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int, chJobs chan map[int]*bool, errCh chan error) {
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
					chResult <- map[int]*int{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int, len(list))
	newList := make([]*int, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapBoolIntPtrErrNoOrder(f func(*bool) (*int, error), list []*bool, worker int) ([]*int, error) {
	chJobs := make(chan *bool, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int, chJobs chan *bool, errCh chan error) {
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

	newList := make([]*int, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int{}, <-errCh
	}

	return newList, nil
}

// PMapBoolInt64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolInt64PtrErr(f func(*bool) (*int64, error), list []*bool, optional ...Optional) ([]*int64, error) {
	if f == nil {
		return []*int64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolInt64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapBoolInt64PtrErrPreserveOrder(f, list, worker)
}

func pMapBoolInt64PtrErrPreserveOrder(f func(*bool) (*int64, error), list []*bool, worker int) ([]*int64, error) {
	chJobs := make(chan map[int]*bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int64, chJobs chan map[int]*bool, errCh chan error) {
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
					chResult <- map[int]*int64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int64, len(list))
	newList := make([]*int64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapBoolInt64PtrErrNoOrder(f func(*bool) (*int64, error), list []*bool, worker int) ([]*int64, error) {
	chJobs := make(chan *bool, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int64, chJobs chan *bool, errCh chan error) {
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

	newList := make([]*int64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int64{}, <-errCh
	}

	return newList, nil
}

// PMapBoolInt32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolInt32PtrErr(f func(*bool) (*int32, error), list []*bool, optional ...Optional) ([]*int32, error) {
	if f == nil {
		return []*int32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolInt32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapBoolInt32PtrErrPreserveOrder(f, list, worker)
}

func pMapBoolInt32PtrErrPreserveOrder(f func(*bool) (*int32, error), list []*bool, worker int) ([]*int32, error) {
	chJobs := make(chan map[int]*bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int32, chJobs chan map[int]*bool, errCh chan error) {
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
					chResult <- map[int]*int32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int32, len(list))
	newList := make([]*int32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapBoolInt32PtrErrNoOrder(f func(*bool) (*int32, error), list []*bool, worker int) ([]*int32, error) {
	chJobs := make(chan *bool, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int32, chJobs chan *bool, errCh chan error) {
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

	newList := make([]*int32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int32{}, <-errCh
	}

	return newList, nil
}

// PMapBoolInt16PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolInt16PtrErr(f func(*bool) (*int16, error), list []*bool, optional ...Optional) ([]*int16, error) {
	if f == nil {
		return []*int16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolInt16PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapBoolInt16PtrErrPreserveOrder(f, list, worker)
}

func pMapBoolInt16PtrErrPreserveOrder(f func(*bool) (*int16, error), list []*bool, worker int) ([]*int16, error) {
	chJobs := make(chan map[int]*bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int16, chJobs chan map[int]*bool, errCh chan error) {
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
					chResult <- map[int]*int16{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int16, len(list))
	newList := make([]*int16, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int16{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int16{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapBoolInt16PtrErrNoOrder(f func(*bool) (*int16, error), list []*bool, worker int) ([]*int16, error) {
	chJobs := make(chan *bool, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int16, chJobs chan *bool, errCh chan error) {
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

	newList := make([]*int16, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int16{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int16{}, <-errCh
	}

	return newList, nil
}

// PMapBoolInt8PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolInt8PtrErr(f func(*bool) (*int8, error), list []*bool, optional ...Optional) ([]*int8, error) {
	if f == nil {
		return []*int8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolInt8PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapBoolInt8PtrErrPreserveOrder(f, list, worker)
}

func pMapBoolInt8PtrErrPreserveOrder(f func(*bool) (*int8, error), list []*bool, worker int) ([]*int8, error) {
	chJobs := make(chan map[int]*bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int8, chJobs chan map[int]*bool, errCh chan error) {
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
					chResult <- map[int]*int8{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int8, len(list))
	newList := make([]*int8, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int8{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int8{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapBoolInt8PtrErrNoOrder(f func(*bool) (*int8, error), list []*bool, worker int) ([]*int8, error) {
	chJobs := make(chan *bool, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int8, chJobs chan *bool, errCh chan error) {
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

	newList := make([]*int8, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int8{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int8{}, <-errCh
	}

	return newList, nil
}

// PMapBoolUintPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolUintPtrErr(f func(*bool) (*uint, error), list []*bool, optional ...Optional) ([]*uint, error) {
	if f == nil {
		return []*uint{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolUintPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapBoolUintPtrErrPreserveOrder(f, list, worker)
}

func pMapBoolUintPtrErrPreserveOrder(f func(*bool) (*uint, error), list []*bool, worker int) ([]*uint, error) {
	chJobs := make(chan map[int]*bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint, chJobs chan map[int]*bool, errCh chan error) {
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
					chResult <- map[int]*uint{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint, len(list))
	newList := make([]*uint, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapBoolUintPtrErrNoOrder(f func(*bool) (*uint, error), list []*bool, worker int) ([]*uint, error) {
	chJobs := make(chan *bool, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint, chJobs chan *bool, errCh chan error) {
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

	newList := make([]*uint, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint{}, <-errCh
	}

	return newList, nil
}

// PMapBoolUint64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolUint64PtrErr(f func(*bool) (*uint64, error), list []*bool, optional ...Optional) ([]*uint64, error) {
	if f == nil {
		return []*uint64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolUint64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapBoolUint64PtrErrPreserveOrder(f, list, worker)
}

func pMapBoolUint64PtrErrPreserveOrder(f func(*bool) (*uint64, error), list []*bool, worker int) ([]*uint64, error) {
	chJobs := make(chan map[int]*bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint64, chJobs chan map[int]*bool, errCh chan error) {
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
					chResult <- map[int]*uint64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint64, len(list))
	newList := make([]*uint64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapBoolUint64PtrErrNoOrder(f func(*bool) (*uint64, error), list []*bool, worker int) ([]*uint64, error) {
	chJobs := make(chan *bool, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint64, chJobs chan *bool, errCh chan error) {
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

	newList := make([]*uint64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint64{}, <-errCh
	}

	return newList, nil
}

// PMapBoolUint32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolUint32PtrErr(f func(*bool) (*uint32, error), list []*bool, optional ...Optional) ([]*uint32, error) {
	if f == nil {
		return []*uint32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolUint32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapBoolUint32PtrErrPreserveOrder(f, list, worker)
}

func pMapBoolUint32PtrErrPreserveOrder(f func(*bool) (*uint32, error), list []*bool, worker int) ([]*uint32, error) {
	chJobs := make(chan map[int]*bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint32, chJobs chan map[int]*bool, errCh chan error) {
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
					chResult <- map[int]*uint32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint32, len(list))
	newList := make([]*uint32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapBoolUint32PtrErrNoOrder(f func(*bool) (*uint32, error), list []*bool, worker int) ([]*uint32, error) {
	chJobs := make(chan *bool, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint32, chJobs chan *bool, errCh chan error) {
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

	newList := make([]*uint32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint32{}, <-errCh
	}

	return newList, nil
}

// PMapBoolUint16PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolUint16PtrErr(f func(*bool) (*uint16, error), list []*bool, optional ...Optional) ([]*uint16, error) {
	if f == nil {
		return []*uint16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolUint16PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapBoolUint16PtrErrPreserveOrder(f, list, worker)
}

func pMapBoolUint16PtrErrPreserveOrder(f func(*bool) (*uint16, error), list []*bool, worker int) ([]*uint16, error) {
	chJobs := make(chan map[int]*bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint16, chJobs chan map[int]*bool, errCh chan error) {
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
					chResult <- map[int]*uint16{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint16, len(list))
	newList := make([]*uint16, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint16{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint16{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapBoolUint16PtrErrNoOrder(f func(*bool) (*uint16, error), list []*bool, worker int) ([]*uint16, error) {
	chJobs := make(chan *bool, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint16, chJobs chan *bool, errCh chan error) {
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

	newList := make([]*uint16, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint16{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint16{}, <-errCh
	}

	return newList, nil
}

// PMapBoolUint8PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolUint8PtrErr(f func(*bool) (*uint8, error), list []*bool, optional ...Optional) ([]*uint8, error) {
	if f == nil {
		return []*uint8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolUint8PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapBoolUint8PtrErrPreserveOrder(f, list, worker)
}

func pMapBoolUint8PtrErrPreserveOrder(f func(*bool) (*uint8, error), list []*bool, worker int) ([]*uint8, error) {
	chJobs := make(chan map[int]*bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint8, chJobs chan map[int]*bool, errCh chan error) {
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
					chResult <- map[int]*uint8{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint8, len(list))
	newList := make([]*uint8, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint8{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint8{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapBoolUint8PtrErrNoOrder(f func(*bool) (*uint8, error), list []*bool, worker int) ([]*uint8, error) {
	chJobs := make(chan *bool, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint8, chJobs chan *bool, errCh chan error) {
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

	newList := make([]*uint8, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint8{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint8{}, <-errCh
	}

	return newList, nil
}

// PMapBoolStrPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolStrPtrErr(f func(*bool) (*string, error), list []*bool, optional ...Optional) ([]*string, error) {
	if f == nil {
		return []*string{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolStrPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapBoolStrPtrErrPreserveOrder(f, list, worker)
}

func pMapBoolStrPtrErrPreserveOrder(f func(*bool) (*string, error), list []*bool, worker int) ([]*string, error) {
	chJobs := make(chan map[int]*bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*string, chJobs chan map[int]*bool, errCh chan error) {
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
					chResult <- map[int]*string{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*string, len(list))
	newList := make([]*string, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*string{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*string{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapBoolStrPtrErrNoOrder(f func(*bool) (*string, error), list []*bool, worker int) ([]*string, error) {
	chJobs := make(chan *bool, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *string, chJobs chan *bool, errCh chan error) {
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

	newList := make([]*string, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*string{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*string{}, <-errCh
	}

	return newList, nil
}

// PMapBoolFloat32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolFloat32PtrErr(f func(*bool) (*float32, error), list []*bool, optional ...Optional) ([]*float32, error) {
	if f == nil {
		return []*float32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolFloat32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapBoolFloat32PtrErrPreserveOrder(f, list, worker)
}

func pMapBoolFloat32PtrErrPreserveOrder(f func(*bool) (*float32, error), list []*bool, worker int) ([]*float32, error) {
	chJobs := make(chan map[int]*bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float32, chJobs chan map[int]*bool, errCh chan error) {
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
					chResult <- map[int]*float32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float32, len(list))
	newList := make([]*float32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*float32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*float32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapBoolFloat32PtrErrNoOrder(f func(*bool) (*float32, error), list []*bool, worker int) ([]*float32, error) {
	chJobs := make(chan *bool, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float32, chJobs chan *bool, errCh chan error) {
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

	newList := make([]*float32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*float32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*float32{}, <-errCh
	}

	return newList, nil
}

// PMapBoolFloat64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolFloat64PtrErr(f func(*bool) (*float64, error), list []*bool, optional ...Optional) ([]*float64, error) {
	if f == nil {
		return []*float64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolFloat64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapBoolFloat64PtrErrPreserveOrder(f, list, worker)
}

func pMapBoolFloat64PtrErrPreserveOrder(f func(*bool) (*float64, error), list []*bool, worker int) ([]*float64, error) {
	chJobs := make(chan map[int]*bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float64, chJobs chan map[int]*bool, errCh chan error) {
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
					chResult <- map[int]*float64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float64, len(list))
	newList := make([]*float64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*float64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*float64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapBoolFloat64PtrErrNoOrder(f func(*bool) (*float64, error), list []*bool, worker int) ([]*float64, error) {
	chJobs := make(chan *bool, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float64, chJobs chan *bool, errCh chan error) {
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

	newList := make([]*float64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*float64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*float64{}, <-errCh
	}

	return newList, nil
}

// PMapFloat32IntPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32IntPtrErr(f func(*float32) (*int, error), list []*float32, optional ...Optional) ([]*int, error) {
	if f == nil {
		return []*int{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32IntPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32IntPtrErrPreserveOrder(f, list, worker)
}

func pMapFloat32IntPtrErrPreserveOrder(f func(*float32) (*int, error), list []*float32, worker int) ([]*int, error) {
	chJobs := make(chan map[int]*float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int, chJobs chan map[int]*float32, errCh chan error) {
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
					chResult <- map[int]*int{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int, len(list))
	newList := make([]*int, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapFloat32IntPtrErrNoOrder(f func(*float32) (*int, error), list []*float32, worker int) ([]*int, error) {
	chJobs := make(chan *float32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int, chJobs chan *float32, errCh chan error) {
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

	newList := make([]*int, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int{}, <-errCh
	}

	return newList, nil
}

// PMapFloat32Int64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Int64PtrErr(f func(*float32) (*int64, error), list []*float32, optional ...Optional) ([]*int64, error) {
	if f == nil {
		return []*int64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32Int64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32Int64PtrErrPreserveOrder(f, list, worker)
}

func pMapFloat32Int64PtrErrPreserveOrder(f func(*float32) (*int64, error), list []*float32, worker int) ([]*int64, error) {
	chJobs := make(chan map[int]*float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int64, chJobs chan map[int]*float32, errCh chan error) {
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
					chResult <- map[int]*int64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int64, len(list))
	newList := make([]*int64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapFloat32Int64PtrErrNoOrder(f func(*float32) (*int64, error), list []*float32, worker int) ([]*int64, error) {
	chJobs := make(chan *float32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int64, chJobs chan *float32, errCh chan error) {
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

	newList := make([]*int64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int64{}, <-errCh
	}

	return newList, nil
}

// PMapFloat32Int32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Int32PtrErr(f func(*float32) (*int32, error), list []*float32, optional ...Optional) ([]*int32, error) {
	if f == nil {
		return []*int32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32Int32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32Int32PtrErrPreserveOrder(f, list, worker)
}

func pMapFloat32Int32PtrErrPreserveOrder(f func(*float32) (*int32, error), list []*float32, worker int) ([]*int32, error) {
	chJobs := make(chan map[int]*float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int32, chJobs chan map[int]*float32, errCh chan error) {
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
					chResult <- map[int]*int32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int32, len(list))
	newList := make([]*int32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapFloat32Int32PtrErrNoOrder(f func(*float32) (*int32, error), list []*float32, worker int) ([]*int32, error) {
	chJobs := make(chan *float32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int32, chJobs chan *float32, errCh chan error) {
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

	newList := make([]*int32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int32{}, <-errCh
	}

	return newList, nil
}

// PMapFloat32Int16PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Int16PtrErr(f func(*float32) (*int16, error), list []*float32, optional ...Optional) ([]*int16, error) {
	if f == nil {
		return []*int16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32Int16PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32Int16PtrErrPreserveOrder(f, list, worker)
}

func pMapFloat32Int16PtrErrPreserveOrder(f func(*float32) (*int16, error), list []*float32, worker int) ([]*int16, error) {
	chJobs := make(chan map[int]*float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int16, chJobs chan map[int]*float32, errCh chan error) {
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
					chResult <- map[int]*int16{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int16, len(list))
	newList := make([]*int16, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int16{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int16{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapFloat32Int16PtrErrNoOrder(f func(*float32) (*int16, error), list []*float32, worker int) ([]*int16, error) {
	chJobs := make(chan *float32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int16, chJobs chan *float32, errCh chan error) {
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

	newList := make([]*int16, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int16{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int16{}, <-errCh
	}

	return newList, nil
}

// PMapFloat32Int8PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Int8PtrErr(f func(*float32) (*int8, error), list []*float32, optional ...Optional) ([]*int8, error) {
	if f == nil {
		return []*int8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32Int8PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32Int8PtrErrPreserveOrder(f, list, worker)
}

func pMapFloat32Int8PtrErrPreserveOrder(f func(*float32) (*int8, error), list []*float32, worker int) ([]*int8, error) {
	chJobs := make(chan map[int]*float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int8, chJobs chan map[int]*float32, errCh chan error) {
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
					chResult <- map[int]*int8{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int8, len(list))
	newList := make([]*int8, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int8{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int8{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapFloat32Int8PtrErrNoOrder(f func(*float32) (*int8, error), list []*float32, worker int) ([]*int8, error) {
	chJobs := make(chan *float32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int8, chJobs chan *float32, errCh chan error) {
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

	newList := make([]*int8, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int8{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int8{}, <-errCh
	}

	return newList, nil
}

// PMapFloat32UintPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32UintPtrErr(f func(*float32) (*uint, error), list []*float32, optional ...Optional) ([]*uint, error) {
	if f == nil {
		return []*uint{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32UintPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32UintPtrErrPreserveOrder(f, list, worker)
}

func pMapFloat32UintPtrErrPreserveOrder(f func(*float32) (*uint, error), list []*float32, worker int) ([]*uint, error) {
	chJobs := make(chan map[int]*float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint, chJobs chan map[int]*float32, errCh chan error) {
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
					chResult <- map[int]*uint{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint, len(list))
	newList := make([]*uint, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapFloat32UintPtrErrNoOrder(f func(*float32) (*uint, error), list []*float32, worker int) ([]*uint, error) {
	chJobs := make(chan *float32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint, chJobs chan *float32, errCh chan error) {
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

	newList := make([]*uint, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint{}, <-errCh
	}

	return newList, nil
}

// PMapFloat32Uint64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Uint64PtrErr(f func(*float32) (*uint64, error), list []*float32, optional ...Optional) ([]*uint64, error) {
	if f == nil {
		return []*uint64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32Uint64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32Uint64PtrErrPreserveOrder(f, list, worker)
}

func pMapFloat32Uint64PtrErrPreserveOrder(f func(*float32) (*uint64, error), list []*float32, worker int) ([]*uint64, error) {
	chJobs := make(chan map[int]*float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint64, chJobs chan map[int]*float32, errCh chan error) {
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
					chResult <- map[int]*uint64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint64, len(list))
	newList := make([]*uint64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapFloat32Uint64PtrErrNoOrder(f func(*float32) (*uint64, error), list []*float32, worker int) ([]*uint64, error) {
	chJobs := make(chan *float32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint64, chJobs chan *float32, errCh chan error) {
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

	newList := make([]*uint64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint64{}, <-errCh
	}

	return newList, nil
}

// PMapFloat32Uint32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Uint32PtrErr(f func(*float32) (*uint32, error), list []*float32, optional ...Optional) ([]*uint32, error) {
	if f == nil {
		return []*uint32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32Uint32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32Uint32PtrErrPreserveOrder(f, list, worker)
}

func pMapFloat32Uint32PtrErrPreserveOrder(f func(*float32) (*uint32, error), list []*float32, worker int) ([]*uint32, error) {
	chJobs := make(chan map[int]*float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint32, chJobs chan map[int]*float32, errCh chan error) {
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
					chResult <- map[int]*uint32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint32, len(list))
	newList := make([]*uint32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapFloat32Uint32PtrErrNoOrder(f func(*float32) (*uint32, error), list []*float32, worker int) ([]*uint32, error) {
	chJobs := make(chan *float32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint32, chJobs chan *float32, errCh chan error) {
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

	newList := make([]*uint32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint32{}, <-errCh
	}

	return newList, nil
}

// PMapFloat32Uint16PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Uint16PtrErr(f func(*float32) (*uint16, error), list []*float32, optional ...Optional) ([]*uint16, error) {
	if f == nil {
		return []*uint16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32Uint16PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32Uint16PtrErrPreserveOrder(f, list, worker)
}

func pMapFloat32Uint16PtrErrPreserveOrder(f func(*float32) (*uint16, error), list []*float32, worker int) ([]*uint16, error) {
	chJobs := make(chan map[int]*float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint16, chJobs chan map[int]*float32, errCh chan error) {
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
					chResult <- map[int]*uint16{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint16, len(list))
	newList := make([]*uint16, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint16{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint16{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapFloat32Uint16PtrErrNoOrder(f func(*float32) (*uint16, error), list []*float32, worker int) ([]*uint16, error) {
	chJobs := make(chan *float32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint16, chJobs chan *float32, errCh chan error) {
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

	newList := make([]*uint16, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint16{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint16{}, <-errCh
	}

	return newList, nil
}

// PMapFloat32Uint8PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Uint8PtrErr(f func(*float32) (*uint8, error), list []*float32, optional ...Optional) ([]*uint8, error) {
	if f == nil {
		return []*uint8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32Uint8PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32Uint8PtrErrPreserveOrder(f, list, worker)
}

func pMapFloat32Uint8PtrErrPreserveOrder(f func(*float32) (*uint8, error), list []*float32, worker int) ([]*uint8, error) {
	chJobs := make(chan map[int]*float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint8, chJobs chan map[int]*float32, errCh chan error) {
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
					chResult <- map[int]*uint8{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint8, len(list))
	newList := make([]*uint8, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint8{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint8{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapFloat32Uint8PtrErrNoOrder(f func(*float32) (*uint8, error), list []*float32, worker int) ([]*uint8, error) {
	chJobs := make(chan *float32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint8, chJobs chan *float32, errCh chan error) {
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

	newList := make([]*uint8, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint8{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint8{}, <-errCh
	}

	return newList, nil
}

// PMapFloat32StrPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32StrPtrErr(f func(*float32) (*string, error), list []*float32, optional ...Optional) ([]*string, error) {
	if f == nil {
		return []*string{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32StrPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32StrPtrErrPreserveOrder(f, list, worker)
}

func pMapFloat32StrPtrErrPreserveOrder(f func(*float32) (*string, error), list []*float32, worker int) ([]*string, error) {
	chJobs := make(chan map[int]*float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*string, chJobs chan map[int]*float32, errCh chan error) {
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
					chResult <- map[int]*string{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*string, len(list))
	newList := make([]*string, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*string{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*string{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapFloat32StrPtrErrNoOrder(f func(*float32) (*string, error), list []*float32, worker int) ([]*string, error) {
	chJobs := make(chan *float32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *string, chJobs chan *float32, errCh chan error) {
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

	newList := make([]*string, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*string{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*string{}, <-errCh
	}

	return newList, nil
}

// PMapFloat32BoolPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32BoolPtrErr(f func(*float32) (*bool, error), list []*float32, optional ...Optional) ([]*bool, error) {
	if f == nil {
		return []*bool{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32BoolPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32BoolPtrErrPreserveOrder(f, list, worker)
}

func pMapFloat32BoolPtrErrPreserveOrder(f func(*float32) (*bool, error), list []*float32, worker int) ([]*bool, error) {
	chJobs := make(chan map[int]*float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*bool, chJobs chan map[int]*float32, errCh chan error) {
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
					chResult <- map[int]*bool{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*bool, len(list))
	newList := make([]*bool, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*bool{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*bool{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapFloat32BoolPtrErrNoOrder(f func(*float32) (*bool, error), list []*float32, worker int) ([]*bool, error) {
	chJobs := make(chan *float32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *bool, chJobs chan *float32, errCh chan error) {
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

	newList := make([]*bool, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*bool{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*bool{}, <-errCh
	}

	return newList, nil
}

// PMapFloat32Float64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Float64PtrErr(f func(*float32) (*float64, error), list []*float32, optional ...Optional) ([]*float64, error) {
	if f == nil {
		return []*float64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32Float64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32Float64PtrErrPreserveOrder(f, list, worker)
}

func pMapFloat32Float64PtrErrPreserveOrder(f func(*float32) (*float64, error), list []*float32, worker int) ([]*float64, error) {
	chJobs := make(chan map[int]*float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float64, chJobs chan map[int]*float32, errCh chan error) {
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
					chResult <- map[int]*float64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float64, len(list))
	newList := make([]*float64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*float64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*float64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapFloat32Float64PtrErrNoOrder(f func(*float32) (*float64, error), list []*float32, worker int) ([]*float64, error) {
	chJobs := make(chan *float32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float64, chJobs chan *float32, errCh chan error) {
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

	newList := make([]*float64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*float64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*float64{}, <-errCh
	}

	return newList, nil
}

// PMapFloat64IntPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64IntPtrErr(f func(*float64) (*int, error), list []*float64, optional ...Optional) ([]*int, error) {
	if f == nil {
		return []*int{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64IntPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64IntPtrErrPreserveOrder(f, list, worker)
}

func pMapFloat64IntPtrErrPreserveOrder(f func(*float64) (*int, error), list []*float64, worker int) ([]*int, error) {
	chJobs := make(chan map[int]*float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int, chJobs chan map[int]*float64, errCh chan error) {
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
					chResult <- map[int]*int{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int, len(list))
	newList := make([]*int, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapFloat64IntPtrErrNoOrder(f func(*float64) (*int, error), list []*float64, worker int) ([]*int, error) {
	chJobs := make(chan *float64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int, chJobs chan *float64, errCh chan error) {
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

	newList := make([]*int, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int{}, <-errCh
	}

	return newList, nil
}

// PMapFloat64Int64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Int64PtrErr(f func(*float64) (*int64, error), list []*float64, optional ...Optional) ([]*int64, error) {
	if f == nil {
		return []*int64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64Int64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64Int64PtrErrPreserveOrder(f, list, worker)
}

func pMapFloat64Int64PtrErrPreserveOrder(f func(*float64) (*int64, error), list []*float64, worker int) ([]*int64, error) {
	chJobs := make(chan map[int]*float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int64, chJobs chan map[int]*float64, errCh chan error) {
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
					chResult <- map[int]*int64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int64, len(list))
	newList := make([]*int64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapFloat64Int64PtrErrNoOrder(f func(*float64) (*int64, error), list []*float64, worker int) ([]*int64, error) {
	chJobs := make(chan *float64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int64, chJobs chan *float64, errCh chan error) {
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

	newList := make([]*int64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int64{}, <-errCh
	}

	return newList, nil
}

// PMapFloat64Int32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Int32PtrErr(f func(*float64) (*int32, error), list []*float64, optional ...Optional) ([]*int32, error) {
	if f == nil {
		return []*int32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64Int32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64Int32PtrErrPreserveOrder(f, list, worker)
}

func pMapFloat64Int32PtrErrPreserveOrder(f func(*float64) (*int32, error), list []*float64, worker int) ([]*int32, error) {
	chJobs := make(chan map[int]*float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int32, chJobs chan map[int]*float64, errCh chan error) {
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
					chResult <- map[int]*int32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int32, len(list))
	newList := make([]*int32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapFloat64Int32PtrErrNoOrder(f func(*float64) (*int32, error), list []*float64, worker int) ([]*int32, error) {
	chJobs := make(chan *float64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int32, chJobs chan *float64, errCh chan error) {
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

	newList := make([]*int32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int32{}, <-errCh
	}

	return newList, nil
}

// PMapFloat64Int16PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Int16PtrErr(f func(*float64) (*int16, error), list []*float64, optional ...Optional) ([]*int16, error) {
	if f == nil {
		return []*int16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64Int16PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64Int16PtrErrPreserveOrder(f, list, worker)
}

func pMapFloat64Int16PtrErrPreserveOrder(f func(*float64) (*int16, error), list []*float64, worker int) ([]*int16, error) {
	chJobs := make(chan map[int]*float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int16, chJobs chan map[int]*float64, errCh chan error) {
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
					chResult <- map[int]*int16{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int16, len(list))
	newList := make([]*int16, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int16{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int16{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapFloat64Int16PtrErrNoOrder(f func(*float64) (*int16, error), list []*float64, worker int) ([]*int16, error) {
	chJobs := make(chan *float64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int16, chJobs chan *float64, errCh chan error) {
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

	newList := make([]*int16, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int16{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int16{}, <-errCh
	}

	return newList, nil
}

// PMapFloat64Int8PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Int8PtrErr(f func(*float64) (*int8, error), list []*float64, optional ...Optional) ([]*int8, error) {
	if f == nil {
		return []*int8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64Int8PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64Int8PtrErrPreserveOrder(f, list, worker)
}

func pMapFloat64Int8PtrErrPreserveOrder(f func(*float64) (*int8, error), list []*float64, worker int) ([]*int8, error) {
	chJobs := make(chan map[int]*float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*int8, chJobs chan map[int]*float64, errCh chan error) {
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
					chResult <- map[int]*int8{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*int8, len(list))
	newList := make([]*int8, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*int8{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*int8{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapFloat64Int8PtrErrNoOrder(f func(*float64) (*int8, error), list []*float64, worker int) ([]*int8, error) {
	chJobs := make(chan *float64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *int8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *int8, chJobs chan *float64, errCh chan error) {
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

	newList := make([]*int8, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*int8{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*int8{}, <-errCh
	}

	return newList, nil
}

// PMapFloat64UintPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64UintPtrErr(f func(*float64) (*uint, error), list []*float64, optional ...Optional) ([]*uint, error) {
	if f == nil {
		return []*uint{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64UintPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64UintPtrErrPreserveOrder(f, list, worker)
}

func pMapFloat64UintPtrErrPreserveOrder(f func(*float64) (*uint, error), list []*float64, worker int) ([]*uint, error) {
	chJobs := make(chan map[int]*float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint, chJobs chan map[int]*float64, errCh chan error) {
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
					chResult <- map[int]*uint{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint, len(list))
	newList := make([]*uint, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapFloat64UintPtrErrNoOrder(f func(*float64) (*uint, error), list []*float64, worker int) ([]*uint, error) {
	chJobs := make(chan *float64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint, chJobs chan *float64, errCh chan error) {
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

	newList := make([]*uint, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint{}, <-errCh
	}

	return newList, nil
}

// PMapFloat64Uint64PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Uint64PtrErr(f func(*float64) (*uint64, error), list []*float64, optional ...Optional) ([]*uint64, error) {
	if f == nil {
		return []*uint64{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64Uint64PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64Uint64PtrErrPreserveOrder(f, list, worker)
}

func pMapFloat64Uint64PtrErrPreserveOrder(f func(*float64) (*uint64, error), list []*float64, worker int) ([]*uint64, error) {
	chJobs := make(chan map[int]*float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint64, chJobs chan map[int]*float64, errCh chan error) {
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
					chResult <- map[int]*uint64{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint64, len(list))
	newList := make([]*uint64, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint64{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint64{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapFloat64Uint64PtrErrNoOrder(f func(*float64) (*uint64, error), list []*float64, worker int) ([]*uint64, error) {
	chJobs := make(chan *float64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint64, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint64, chJobs chan *float64, errCh chan error) {
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

	newList := make([]*uint64, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint64{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint64{}, <-errCh
	}

	return newList, nil
}

// PMapFloat64Uint32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Uint32PtrErr(f func(*float64) (*uint32, error), list []*float64, optional ...Optional) ([]*uint32, error) {
	if f == nil {
		return []*uint32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64Uint32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64Uint32PtrErrPreserveOrder(f, list, worker)
}

func pMapFloat64Uint32PtrErrPreserveOrder(f func(*float64) (*uint32, error), list []*float64, worker int) ([]*uint32, error) {
	chJobs := make(chan map[int]*float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint32, chJobs chan map[int]*float64, errCh chan error) {
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
					chResult <- map[int]*uint32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint32, len(list))
	newList := make([]*uint32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapFloat64Uint32PtrErrNoOrder(f func(*float64) (*uint32, error), list []*float64, worker int) ([]*uint32, error) {
	chJobs := make(chan *float64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint32, chJobs chan *float64, errCh chan error) {
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

	newList := make([]*uint32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint32{}, <-errCh
	}

	return newList, nil
}

// PMapFloat64Uint16PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Uint16PtrErr(f func(*float64) (*uint16, error), list []*float64, optional ...Optional) ([]*uint16, error) {
	if f == nil {
		return []*uint16{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64Uint16PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64Uint16PtrErrPreserveOrder(f, list, worker)
}

func pMapFloat64Uint16PtrErrPreserveOrder(f func(*float64) (*uint16, error), list []*float64, worker int) ([]*uint16, error) {
	chJobs := make(chan map[int]*float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint16, chJobs chan map[int]*float64, errCh chan error) {
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
					chResult <- map[int]*uint16{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint16, len(list))
	newList := make([]*uint16, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint16{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint16{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapFloat64Uint16PtrErrNoOrder(f func(*float64) (*uint16, error), list []*float64, worker int) ([]*uint16, error) {
	chJobs := make(chan *float64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint16, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint16, chJobs chan *float64, errCh chan error) {
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

	newList := make([]*uint16, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint16{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint16{}, <-errCh
	}

	return newList, nil
}

// PMapFloat64Uint8PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Uint8PtrErr(f func(*float64) (*uint8, error), list []*float64, optional ...Optional) ([]*uint8, error) {
	if f == nil {
		return []*uint8{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64Uint8PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64Uint8PtrErrPreserveOrder(f, list, worker)
}

func pMapFloat64Uint8PtrErrPreserveOrder(f func(*float64) (*uint8, error), list []*float64, worker int) ([]*uint8, error) {
	chJobs := make(chan map[int]*float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*uint8, chJobs chan map[int]*float64, errCh chan error) {
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
					chResult <- map[int]*uint8{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*uint8, len(list))
	newList := make([]*uint8, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*uint8{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*uint8{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapFloat64Uint8PtrErrNoOrder(f func(*float64) (*uint8, error), list []*float64, worker int) ([]*uint8, error) {
	chJobs := make(chan *float64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *uint8, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *uint8, chJobs chan *float64, errCh chan error) {
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

	newList := make([]*uint8, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*uint8{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*uint8{}, <-errCh
	}

	return newList, nil
}

// PMapFloat64StrPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64StrPtrErr(f func(*float64) (*string, error), list []*float64, optional ...Optional) ([]*string, error) {
	if f == nil {
		return []*string{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64StrPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64StrPtrErrPreserveOrder(f, list, worker)
}

func pMapFloat64StrPtrErrPreserveOrder(f func(*float64) (*string, error), list []*float64, worker int) ([]*string, error) {
	chJobs := make(chan map[int]*float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*string, chJobs chan map[int]*float64, errCh chan error) {
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
					chResult <- map[int]*string{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*string, len(list))
	newList := make([]*string, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*string{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*string{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapFloat64StrPtrErrNoOrder(f func(*float64) (*string, error), list []*float64, worker int) ([]*string, error) {
	chJobs := make(chan *float64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *string, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *string, chJobs chan *float64, errCh chan error) {
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

	newList := make([]*string, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*string{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*string{}, <-errCh
	}

	return newList, nil
}

// PMapFloat64BoolPtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64BoolPtrErr(f func(*float64) (*bool, error), list []*float64, optional ...Optional) ([]*bool, error) {
	if f == nil {
		return []*bool{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64BoolPtrErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64BoolPtrErrPreserveOrder(f, list, worker)
}

func pMapFloat64BoolPtrErrPreserveOrder(f func(*float64) (*bool, error), list []*float64, worker int) ([]*bool, error) {
	chJobs := make(chan map[int]*float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*bool, chJobs chan map[int]*float64, errCh chan error) {
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
					chResult <- map[int]*bool{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*bool, len(list))
	newList := make([]*bool, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*bool{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*bool{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapFloat64BoolPtrErrNoOrder(f func(*float64) (*bool, error), list []*float64, worker int) ([]*bool, error) {
	chJobs := make(chan *float64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *bool, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *bool, chJobs chan *float64, errCh chan error) {
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

	newList := make([]*bool, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*bool{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*bool{}, <-errCh
	}

	return newList, nil
}

// PMapFloat64Float32PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Float32PtrErr(f func(*float64) (*float32, error), list []*float64, optional ...Optional) ([]*float32, error) {
	if f == nil {
		return []*float32{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64Float32PtrErrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64Float32PtrErrPreserveOrder(f, list, worker)
}

func pMapFloat64Float32PtrErrPreserveOrder(f func(*float64) (*float32, error), list []*float64, worker int) ([]*float32, error) {
	chJobs := make(chan map[int]*float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*float32, chJobs chan map[int]*float64, errCh chan error) {
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
					chResult <- map[int]*float32{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*float32, len(list))
	newList := make([]*float32, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*float32{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*float32{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMapFloat64Float32PtrErrNoOrder(f func(*float64) (*float32, error), list []*float64, worker int) ([]*float32, error) {
	chJobs := make(chan *float64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *float32, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *float32, chJobs chan *float64, errCh chan error) {
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

	newList := make([]*float32, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*float32{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*float32{}, <-errCh
	}

	return newList, nil
}
