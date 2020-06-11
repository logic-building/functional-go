package template

// Pmap is template to generate function(Pmap) for user defined data type
func Pmap() string {
	return `
func PMap<CONDITIONAL_TYPE>(f func(<TYPE>) <TYPE>, list []<TYPE>) []<TYPE> {
	if f == nil {
		return []<TYPE>{}
	}

	ch := make(chan map[int]<TYPE>)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]<TYPE>, i int, v <TYPE>) {
			defer wg.Done()
			ch <- map[int]<TYPE>{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]<TYPE>, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}
`
}

// PmapPtr is template to generate function(Pmap) for user defined data type
func PmapPtr() string {
	return `
// PMap<CONDITIONAL_TYPE>Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
func PMap<CONDITIONAL_TYPE>Ptr(f func(*<TYPE>) *<TYPE>, list []*<TYPE>) []*<TYPE> {
	if f == nil {
		return []*<TYPE>{}
	}

	ch := make(chan map[int]*<TYPE>)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*<TYPE>, i int, v *<TYPE>) {
			defer wg.Done()
			ch <- map[int]*<TYPE>{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*<TYPE>, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}
`
}
