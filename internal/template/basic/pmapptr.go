package basic

// PMap applies the function(1st argument) on each item of the list and returns new list
func PMap() string {
	return `
// PMap<FTYPE> applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMap<FTYPE>(f func(<TYPE>) <TYPE>, list []<TYPE>, optional ...Optional) []<TYPE> {
	if f == nil {
		return []<TYPE>{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMap<FTYPE>NoOrder(f, list, worker)
		}
	}

	return pMap<FTYPE>PreserveOrder(f, list, worker)
}

func pMap<FTYPE>PreserveOrder(f func(<TYPE>) <TYPE>, list []<TYPE>, worker int) []<TYPE> {
	chJobs := make(chan map[int]<TYPE>, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]<TYPE>{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]<TYPE>, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]<TYPE>, chJobs chan map[int]<TYPE>) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]<TYPE>{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]<TYPE>, len(list))
	newList := make([]<TYPE>, len(list))

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

func pMap<FTYPE>NoOrder(f func(<TYPE>) <TYPE>, list []<TYPE>, worker int) []<TYPE> {
	chJobs := make(chan <TYPE>, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan <TYPE>, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan <TYPE>, chJobs chan <TYPE>) {
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

	newList := make([]<TYPE>, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}
`
}

// PMapPtr applies the function(1st argument) on each item of the list and returns new list
func PMapPtr() string {
	return `
// PMap<FTYPE>Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMap<FTYPE>Ptr(f func(*<TYPE>) *<TYPE>, list []*<TYPE>, optional ...Optional) []*<TYPE> {
	if f == nil {
		return []*<TYPE>{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMap<FTYPE>PtrNoOrder(f, list, worker)
		}
	}

	return pMap<FTYPE>PtrPreserveOrder(f, list, worker)
}

func pMap<FTYPE>PtrPreserveOrder(f func(*<TYPE>) *<TYPE>, list []*<TYPE>, worker int) []*<TYPE> {
	chJobs := make(chan map[int]*<TYPE>, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*<TYPE>{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*<TYPE>, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*<TYPE>, chJobs chan map[int]*<TYPE>) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*<TYPE>{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*<TYPE>, len(list))
	newList := make([]*<TYPE>, len(list))

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

func pMap<FTYPE>PtrNoOrder(f func(*<TYPE>) *<TYPE>, list []*<TYPE>, worker int) []*<TYPE> {
	chJobs := make(chan *<TYPE>, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *<TYPE>, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *<TYPE>, chJobs chan *<TYPE>) {
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

	newList := make([]*<TYPE>, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}
`
}

// PMapPtrErr applies the function(1st argument) on each item of the list and returns new list
func PMapPtrErr() string {
	return `
// PMap<FTYPE>PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMap<FTYPE>PtrErr(f func(*<TYPE>) (*<TYPE>, error), list []*<TYPE>, optional ...Optional) ([]*<TYPE>, error) {
	if f == nil {
		return []*<TYPE>{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMap<FTYPE>PtrErrNoOrder(f, list, worker)
		}
	}

	return pMap<FTYPE>PtrErrPreserveOrder(f, list, worker)
}

func pMap<FTYPE>PtrErrPreserveOrder(f func(*<TYPE>) (*<TYPE>, error), list []*<TYPE>, worker int) ([]*<TYPE>, error) {
	chJobs := make(chan map[int]*<TYPE>, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*<TYPE>{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*<TYPE>, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*<TYPE>, chJobs chan map[int]*<TYPE>, errCh chan error) {
			defer wg.Done()
			if len(errCh) > 0 {
				return
			}
			for m := range chJobs {
				for k, v := range m {
					r, err := f(v)
					if err != nil {
						errCh <- err
						return
					}
					chResult <- map[int]*<TYPE>{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*<TYPE>, len(list))
	newList := make([]*<TYPE>, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*<TYPE>{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*<TYPE>{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMap<FTYPE>PtrErrNoOrder(f func(*<TYPE>) (*<TYPE>, error), list []*<TYPE>, worker int) ([]*<TYPE>, error) {
	chJobs := make(chan *<TYPE>, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *<TYPE>, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *<TYPE>, chJobs chan *<TYPE>, errCh chan error) {
			defer wg.Done()
			if len(errCh) > 0 {
				return
			}
			for v := range chJobs {
				r, err := f(v)
				if err != nil {
					errCh <- err
					return
				}
				chResult <- r
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*<TYPE>, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*<TYPE>{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*<TYPE>{}, <-errCh
	}

	return newList, nil
}
`
}

// PMapErr applies the function(1st argument) on each item of the list and returns new list
func PMapErr() string {
	return `
// PMap<FTYPE>Err applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMap<FTYPE>Err(f func(<TYPE>) (<TYPE>, error), list []<TYPE>, optional ...Optional) ([]<TYPE>, error) {
	if f == nil {
		return []<TYPE>{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMap<FTYPE>ErrNoOrder(f, list, worker)
		}
	}

	return pMap<FTYPE>ErrPreserveOrder(f, list, worker)
}

func pMap<FTYPE>ErrPreserveOrder(f func(<TYPE>) (<TYPE>, error), list []<TYPE>, worker int) ([]<TYPE>, error) {
	chJobs := make(chan map[int]<TYPE>, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]<TYPE>{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]<TYPE>, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]<TYPE>, chJobs chan map[int]<TYPE>, errCh chan error) {
			defer wg.Done()
			if len(errCh) > 0 {
				return
			}
			for m := range chJobs {
				for k, v := range m {
					r, err := f(v)
					if err != nil {
						errCh <- err
						return
					}
					chResult <- map[int]<TYPE>{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]<TYPE>, len(list))
	newList := make([]<TYPE>, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []<TYPE>{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []<TYPE>{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMap<FTYPE>ErrNoOrder(f func(<TYPE>) (<TYPE>, error), list []<TYPE>, worker int) ([]<TYPE>, error) {
	chJobs := make(chan <TYPE>, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan <TYPE>, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan <TYPE>, chJobs chan <TYPE>, errCh chan error) {
			defer wg.Done()
			if len(errCh) > 0 {
				return
			}
			for v := range chJobs {
				r, err := f(v)
				if err != nil {
					errCh <- err
					return
				}
				chResult <- r
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]<TYPE>, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []<TYPE>{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []<TYPE>{}, <-errCh
	}

	return newList, nil
}
`
}
