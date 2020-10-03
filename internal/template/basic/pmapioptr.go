package basic

// PMapIOPtr is template to generate itself for different combination of data type.
func PMapIOPtr() string {
	return `
// PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMap<FINPUT_TYPE><FOUTPUT_TYPE>Ptr(f func(*<INPUT_TYPE>) *<OUTPUT_TYPE>, list []*<INPUT_TYPE>, optional ...Optional) []*<OUTPUT_TYPE> {
	if f == nil {
		return []*<OUTPUT_TYPE>{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrNoOrder(f, list, worker)
		}
	}

	return pMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrPreserveOrder(f, list, worker)
}

func pMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrPreserveOrder(f func(*<INPUT_TYPE>) *<OUTPUT_TYPE>, list []*<INPUT_TYPE>, worker int) []*<OUTPUT_TYPE> {
	chJobs := make(chan map[int]*<INPUT_TYPE>, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*<INPUT_TYPE>{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*<OUTPUT_TYPE>, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*<OUTPUT_TYPE>, chJobs chan map[int]*<INPUT_TYPE>) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]*<OUTPUT_TYPE>{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*<OUTPUT_TYPE>, len(list))
	newList := make([]*<OUTPUT_TYPE>, len(list))

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

func pMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrNoOrder(f func(*<INPUT_TYPE>) *<OUTPUT_TYPE>, list []*<INPUT_TYPE>, worker int) []*<OUTPUT_TYPE> {
	chJobs := make(chan *<INPUT_TYPE>, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *<OUTPUT_TYPE>, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *<OUTPUT_TYPE>, chJobs chan *<INPUT_TYPE>) {
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

	newList := make([]*<OUTPUT_TYPE>, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}
`
}

// PMapIOPtrErr is template to generate itself for different combination of data type.
func PMapIOPtrErr() string {
	return `
// PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr applies the function(1st argument) on each item in the list and returns a new list and error.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErr(f func(*<INPUT_TYPE>) (*<OUTPUT_TYPE>, error), list []*<INPUT_TYPE>, optional ...Optional) ([]*<OUTPUT_TYPE>, error) {
	if f == nil {
		return []*<OUTPUT_TYPE>{}, nil
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErrNoOrder(f, list, worker)
		}
	}

	return pMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErrPreserveOrder(f, list, worker)
}

func pMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErrPreserveOrder(f func(*<INPUT_TYPE>) (*<OUTPUT_TYPE>, error), list []*<INPUT_TYPE>, worker int) ([]*<OUTPUT_TYPE>, error) {
	chJobs := make(chan map[int]*<INPUT_TYPE>, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]*<INPUT_TYPE>{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]*<OUTPUT_TYPE>, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]*<OUTPUT_TYPE>, chJobs chan map[int]*<INPUT_TYPE>, errCh chan error) {
			defer wg.Done()
			if len(errCh) > 0 {
				return
			}
			for m := range chJobs {
				for k, v := range m {
					r, err := f(v)
					if err != nil {
						errCh <- err
						return
					}
					chResult <- map[int]*<OUTPUT_TYPE>{k: r}
				}
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]*<OUTPUT_TYPE>, len(list))
	newList := make([]*<OUTPUT_TYPE>, len(list))

	for m := range chResult {
		select {
		case err := <-errCh:
			return []*<OUTPUT_TYPE>{}, err
		default:
		}
		for k, v := range m {
			newListMap[k] = v
		}
	}

	if len(errCh) > 0 {
		return []*<OUTPUT_TYPE>{}, <-errCh
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList, nil
}

func pMap<FINPUT_TYPE><FOUTPUT_TYPE>PtrErrNoOrder(f func(*<INPUT_TYPE>) (*<OUTPUT_TYPE>, error), list []*<INPUT_TYPE>, worker int) ([]*<OUTPUT_TYPE>, error) {
	chJobs := make(chan *<INPUT_TYPE>, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan *<OUTPUT_TYPE>, worker/3)

	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan *<OUTPUT_TYPE>, chJobs chan *<INPUT_TYPE>, errCh chan error) {
			defer wg.Done()
			if len(errCh) > 0 {
				return
			}
			for v := range chJobs {
				r, err := f(v)
				if err != nil {
					errCh <- err
					return
				}
				chResult <- r
			}
		}(chResult, chJobs, errCh)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]*<OUTPUT_TYPE>, len(list))
	i := 0

	for v := range chResult {
		select {
		case err := <-errCh:
			return []*<OUTPUT_TYPE>{}, err
		default:
		}
		newList[i] = v
		i++
	}

	if len(errCh) > 0 {
		return []*<OUTPUT_TYPE>{}, <-errCh
	}

	return newList, nil
}
`
}
