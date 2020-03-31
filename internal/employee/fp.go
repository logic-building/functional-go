// Code generated by 'gofp'. DO NOT EDIT.
package employee
import "sync" 

func Map(f func(Employee) Employee, list []Employee) []Employee {
	if f == nil {
		return []Employee{}
	}
	newList := make([]Employee, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapPtr(f func(*Employee) *Employee, list []*Employee) []*Employee {
	if f == nil {
		return []*Employee{}
	}
	newList := make([]*Employee, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func Filter(f func(Employee) bool, list []Employee) []Employee {
	if f == nil {
		return []Employee{}
	}
	var newList []Employee
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

func FilterPtr(f func(*Employee) bool, list []*Employee) []*Employee {
	if f == nil {
		return []*Employee{}
	}
	var newList []*Employee
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

func Remove(f func(Employee) bool, list []Employee) []Employee {
	if f == nil {
		return []Employee{}
	}
	var newList []Employee
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

func Some(f func(Employee) bool, list []Employee) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}

func SomePtr(f func(*Employee) bool, list []*Employee) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}

func Every(f func(Employee) bool, list []Employee) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

func DropWhile(f func(Employee) bool, list []Employee) []Employee {
	if f == nil {
		return []Employee{}
	}
	var newList []Employee
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]Employee, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

func TakeWhile(f func(Employee) bool, list []Employee) []Employee {
	if f == nil {
		return []Employee{}
	}
	var newList []Employee
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}

func PMap(f func(Employee) Employee, list []Employee) []Employee {
	if f == nil {
		return []Employee{}
	}

	ch := make(chan map[int]Employee)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]Employee, i int, v Employee) {
			defer wg.Done()
			ch <- map[int]Employee{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]Employee, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

func FilterMap(fFilter func(Employee) bool, fMap func(Employee) Employee, list []Employee) []Employee {
	if fFilter == nil || fMap == nil {
		return []Employee{}
	}
	var newList []Employee
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

func FilterMapPtr(fFilter func(*Employee) bool, fMap func(*Employee) *Employee, list []*Employee) []*Employee {
	if fFilter == nil || fMap == nil {
		return []*Employee{}
	}
	var newList []*Employee
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

func Rest(l []Employee) []Employee {
	if l == nil {
		return []Employee{}
	}

	len := len(l)
	if len == 0 || len == 1 {
		return []Employee{}
	}

	newList := make([]Employee, len-1)

	for i, v := range l[1:] {
		newList[i] = v
	}

	return newList
}

func Reduce(f func(Employee, Employee) Employee, list []Employee, initializer ...Employee) Employee {
	var init Employee 
	lenList := len(list)

	if initializer != nil {
		init = initializer[0]
	} else if lenList > 0 {
		init = list[0]
		if lenList == 1 {
			return list[0]
		}
		if lenList >= 2 {
			list = list[1:]
		}
	}
	
	if lenList == 0 {
		return init
	}
	r := f(init, list[0])
	return Reduce(f, list[1:], r)
}

// DropLast drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLast(list []Employee) []Employee {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []Employee{}
	}

	newList := make([]Employee, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

func MapTeacher(f func(Teacher) Teacher, list []Teacher) []Teacher {
	if f == nil {
		return []Teacher{}
	}
	newList := make([]Teacher, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapTeacherPtr(f func(*Teacher) *Teacher, list []*Teacher) []*Teacher {
	if f == nil {
		return []*Teacher{}
	}
	newList := make([]*Teacher, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func FilterTeacher(f func(Teacher) bool, list []Teacher) []Teacher {
	if f == nil {
		return []Teacher{}
	}
	var newList []Teacher
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

func FilterTeacherPtr(f func(*Teacher) bool, list []*Teacher) []*Teacher {
	if f == nil {
		return []*Teacher{}
	}
	var newList []*Teacher
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

func RemoveTeacher(f func(Teacher) bool, list []Teacher) []Teacher {
	if f == nil {
		return []Teacher{}
	}
	var newList []Teacher
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

func SomeTeacher(f func(Teacher) bool, list []Teacher) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}

func SomeTeacherPtr(f func(*Teacher) bool, list []*Teacher) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}

func EveryTeacher(f func(Teacher) bool, list []Teacher) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

func DropWhileTeacher(f func(Teacher) bool, list []Teacher) []Teacher {
	if f == nil {
		return []Teacher{}
	}
	var newList []Teacher
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]Teacher, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

func TakeWhileTeacher(f func(Teacher) bool, list []Teacher) []Teacher {
	if f == nil {
		return []Teacher{}
	}
	var newList []Teacher
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}

func PMapTeacher(f func(Teacher) Teacher, list []Teacher) []Teacher {
	if f == nil {
		return []Teacher{}
	}

	ch := make(chan map[int]Teacher)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]Teacher, i int, v Teacher) {
			defer wg.Done()
			ch <- map[int]Teacher{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]Teacher, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

func FilterMapTeacher(fFilter func(Teacher) bool, fMap func(Teacher) Teacher, list []Teacher) []Teacher {
	if fFilter == nil || fMap == nil {
		return []Teacher{}
	}
	var newList []Teacher
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

func FilterMapTeacherPtr(fFilter func(*Teacher) bool, fMap func(*Teacher) *Teacher, list []*Teacher) []*Teacher {
	if fFilter == nil || fMap == nil {
		return []*Teacher{}
	}
	var newList []*Teacher
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

func RestTeacher(l []Teacher) []Teacher {
	if l == nil {
		return []Teacher{}
	}

	len := len(l)
	if len == 0 || len == 1 {
		return []Teacher{}
	}

	newList := make([]Teacher, len-1)

	for i, v := range l[1:] {
		newList[i] = v
	}

	return newList
}

func ReduceTeacher(f func(Teacher, Teacher) Teacher, list []Teacher, initializer ...Teacher) Teacher {
	var init Teacher 
	lenList := len(list)

	if initializer != nil {
		init = initializer[0]
	} else if lenList > 0 {
		init = list[0]
		if lenList == 1 {
			return list[0]
		}
		if lenList >= 2 {
			list = list[1:]
		}
	}
	
	if lenList == 0 {
		return init
	}
	r := f(init, list[0])
	return ReduceTeacher(f, list[1:], r)
}

// DropLastTeacher drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastTeacher(list []Teacher) []Teacher {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []Teacher{}
	}

	newList := make([]Teacher, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}


// MapEmployeeTeacher takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapEmployeeTeacher(f func(Employee) Teacher, list []Employee) []Teacher {
	if f == nil {
		return []Teacher{}
	}
	newList := make([]Teacher, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapEmployeeTeacher takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapEmployeeTeacherPtr(f func(*Employee) *Teacher, list []*Employee) []*Teacher {
	if f == nil {
		return []*Teacher{}
	}
	newList := make([]*Teacher, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// PMapEmployeeTeacher applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: Employee output type: Teacher
//	2. List
//
// Returns
//	New List of type Teacher
//	Empty list if all arguments are nil or either one is nil
func PMapEmployeeTeacher(f func(Employee) Teacher, list []Employee) []Teacher {
	if f == nil {
		return []Teacher{}
	}

	ch := make(chan map[int]Teacher)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]Teacher, i int, v Employee) {
			defer wg.Done()
			ch <- map[int]Teacher{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]Teacher, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// FilterMapEmployeeTeacher filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - Employee and returns true/false.
//	2. Function: takes Employee as argument and returns Teacher
// 	3. List of type Employee
//
// Returns:
//	New List of type Teacher
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapEmployeeTeacher(fFilter func(Employee) bool, fMap func(Employee) Teacher, list []Employee) []Teacher {
	if fFilter == nil || fMap == nil {
		return []Teacher{}
	}
	var newList []Teacher
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// MapEmployeeInt takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapEmployeeInt(f func(Employee) int, list []Employee) []int {
	if f == nil {
		return []int{}
	}
	newList := make([]int, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapEmployeeInt takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapEmployeeIntPtr(f func(*Employee) *int, list []*Employee) []*int {
	if f == nil {
		return []*int{}
	}
	newList := make([]*int, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// PMapEmployeeInt applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: Employee output type: int
//	2. List
//
// Returns
//	New List of type int
//	Empty list if all arguments are nil or either one is nil
func PMapEmployeeInt(f func(Employee) int, list []Employee) []int {
	if f == nil {
		return []int{}
	}

	ch := make(chan map[int]int)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int, i int, v Employee) {
			defer wg.Done()
			ch <- map[int]int{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]int, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// FilterMapEmployeeInt filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - Employee and returns true/false.
//	2. Function: takes Employee as argument and returns int
// 	3. List of type Employee
//
// Returns:
//	New List of type int
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapEmployeeInt(fFilter func(Employee) bool, fMap func(Employee) int, list []Employee) []int {
	if fFilter == nil || fMap == nil {
		return []int{}
	}
	var newList []int
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// MapEmployeeStr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapEmployeeStr(f func(Employee) string, list []Employee) []string {
	if f == nil {
		return []string{}
	}
	newList := make([]string, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapEmployeeStr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapEmployeeStrPtr(f func(*Employee) *string, list []*Employee) []*string {
	if f == nil {
		return []*string{}
	}
	newList := make([]*string, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// PMapEmployeeStr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: Employee output type: string
//	2. List
//
// Returns
//	New List of type string
//	Empty list if all arguments are nil or either one is nil
func PMapEmployeeStr(f func(Employee) string, list []Employee) []string {
	if f == nil {
		return []string{}
	}

	ch := make(chan map[int]string)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]string, i int, v Employee) {
			defer wg.Done()
			ch <- map[int]string{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]string, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// FilterMapEmployeeStr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - Employee and returns true/false.
//	2. Function: takes Employee as argument and returns string
// 	3. List of type Employee
//
// Returns:
//	New List of type string
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapEmployeeStr(fFilter func(Employee) bool, fMap func(Employee) string, list []Employee) []string {
	if fFilter == nil || fMap == nil {
		return []string{}
	}
	var newList []string
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// MapTeacherEmployee takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapTeacherEmployee(f func(Teacher) Employee, list []Teacher) []Employee {
	if f == nil {
		return []Employee{}
	}
	newList := make([]Employee, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapTeacherEmployee takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapTeacherEmployeePtr(f func(*Teacher) *Employee, list []*Teacher) []*Employee {
	if f == nil {
		return []*Employee{}
	}
	newList := make([]*Employee, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// PMapTeacherEmployee applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: Teacher output type: Employee
//	2. List
//
// Returns
//	New List of type Employee
//	Empty list if all arguments are nil or either one is nil
func PMapTeacherEmployee(f func(Teacher) Employee, list []Teacher) []Employee {
	if f == nil {
		return []Employee{}
	}

	ch := make(chan map[int]Employee)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]Employee, i int, v Teacher) {
			defer wg.Done()
			ch <- map[int]Employee{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]Employee, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// FilterMapTeacherEmployee filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - Teacher and returns true/false.
//	2. Function: takes Teacher as argument and returns Employee
// 	3. List of type Teacher
//
// Returns:
//	New List of type Employee
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapTeacherEmployee(fFilter func(Teacher) bool, fMap func(Teacher) Employee, list []Teacher) []Employee {
	if fFilter == nil || fMap == nil {
		return []Employee{}
	}
	var newList []Employee
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// MapTeacherInt takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapTeacherInt(f func(Teacher) int, list []Teacher) []int {
	if f == nil {
		return []int{}
	}
	newList := make([]int, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapTeacherInt takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapTeacherIntPtr(f func(*Teacher) *int, list []*Teacher) []*int {
	if f == nil {
		return []*int{}
	}
	newList := make([]*int, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// PMapTeacherInt applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: Teacher output type: int
//	2. List
//
// Returns
//	New List of type int
//	Empty list if all arguments are nil or either one is nil
func PMapTeacherInt(f func(Teacher) int, list []Teacher) []int {
	if f == nil {
		return []int{}
	}

	ch := make(chan map[int]int)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int, i int, v Teacher) {
			defer wg.Done()
			ch <- map[int]int{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]int, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// FilterMapTeacherInt filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - Teacher and returns true/false.
//	2. Function: takes Teacher as argument and returns int
// 	3. List of type Teacher
//
// Returns:
//	New List of type int
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapTeacherInt(fFilter func(Teacher) bool, fMap func(Teacher) int, list []Teacher) []int {
	if fFilter == nil || fMap == nil {
		return []int{}
	}
	var newList []int
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// MapTeacherStr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapTeacherStr(f func(Teacher) string, list []Teacher) []string {
	if f == nil {
		return []string{}
	}
	newList := make([]string, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapTeacherStr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapTeacherStrPtr(f func(*Teacher) *string, list []*Teacher) []*string {
	if f == nil {
		return []*string{}
	}
	newList := make([]*string, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// PMapTeacherStr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: Teacher output type: string
//	2. List
//
// Returns
//	New List of type string
//	Empty list if all arguments are nil or either one is nil
func PMapTeacherStr(f func(Teacher) string, list []Teacher) []string {
	if f == nil {
		return []string{}
	}

	ch := make(chan map[int]string)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]string, i int, v Teacher) {
			defer wg.Done()
			ch <- map[int]string{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]string, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// FilterMapTeacherStr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - Teacher and returns true/false.
//	2. Function: takes Teacher as argument and returns string
// 	3. List of type Teacher
//
// Returns:
//	New List of type string
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapTeacherStr(fFilter func(Teacher) bool, fMap func(Teacher) string, list []Teacher) []string {
	if fFilter == nil || fMap == nil {
		return []string{}
	}
	var newList []string
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// MapIntEmployee takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapIntEmployee(f func(int) Employee, list []int) []Employee {
	if f == nil {
		return []Employee{}
	}
	newList := make([]Employee, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapIntEmployee takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapIntEmployeePtr(f func(*int) *Employee, list []*int) []*Employee {
	if f == nil {
		return []*Employee{}
	}
	newList := make([]*Employee, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// PMapIntEmployee applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int output type: Employee
//	2. List
//
// Returns
//	New List of type Employee
//	Empty list if all arguments are nil or either one is nil
func PMapIntEmployee(f func(int) Employee, list []int) []Employee {
	if f == nil {
		return []Employee{}
	}

	ch := make(chan map[int]Employee)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]Employee, i int, v int) {
			defer wg.Done()
			ch <- map[int]Employee{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]Employee, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// FilterMapIntEmployee filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int and returns true/false.
//	2. Function: takes int as argument and returns Employee
// 	3. List of type int
//
// Returns:
//	New List of type Employee
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntEmployee(fFilter func(int) bool, fMap func(int) Employee, list []int) []Employee {
	if fFilter == nil || fMap == nil {
		return []Employee{}
	}
	var newList []Employee
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// MapIntTeacher takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapIntTeacher(f func(int) Teacher, list []int) []Teacher {
	if f == nil {
		return []Teacher{}
	}
	newList := make([]Teacher, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapIntTeacher takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapIntTeacherPtr(f func(*int) *Teacher, list []*int) []*Teacher {
	if f == nil {
		return []*Teacher{}
	}
	newList := make([]*Teacher, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// PMapIntTeacher applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int output type: Teacher
//	2. List
//
// Returns
//	New List of type Teacher
//	Empty list if all arguments are nil or either one is nil
func PMapIntTeacher(f func(int) Teacher, list []int) []Teacher {
	if f == nil {
		return []Teacher{}
	}

	ch := make(chan map[int]Teacher)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]Teacher, i int, v int) {
			defer wg.Done()
			ch <- map[int]Teacher{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]Teacher, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// FilterMapIntTeacher filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int and returns true/false.
//	2. Function: takes int as argument and returns Teacher
// 	3. List of type int
//
// Returns:
//	New List of type Teacher
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntTeacher(fFilter func(int) bool, fMap func(int) Teacher, list []int) []Teacher {
	if fFilter == nil || fMap == nil {
		return []Teacher{}
	}
	var newList []Teacher
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// MapStrEmployee takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapStrEmployee(f func(string) Employee, list []string) []Employee {
	if f == nil {
		return []Employee{}
	}
	newList := make([]Employee, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapStrEmployee takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapStrEmployeePtr(f func(*string) *Employee, list []*string) []*Employee {
	if f == nil {
		return []*Employee{}
	}
	newList := make([]*Employee, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// PMapStrEmployee applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: string output type: Employee
//	2. List
//
// Returns
//	New List of type Employee
//	Empty list if all arguments are nil or either one is nil
func PMapStrEmployee(f func(string) Employee, list []string) []Employee {
	if f == nil {
		return []Employee{}
	}

	ch := make(chan map[int]Employee)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]Employee, i int, v string) {
			defer wg.Done()
			ch <- map[int]Employee{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]Employee, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// FilterMapStrEmployee filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - string and returns true/false.
//	2. Function: takes string as argument and returns Employee
// 	3. List of type string
//
// Returns:
//	New List of type Employee
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrEmployee(fFilter func(string) bool, fMap func(string) Employee, list []string) []Employee {
	if fFilter == nil || fMap == nil {
		return []Employee{}
	}
	var newList []Employee
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// MapStrTeacher takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapStrTeacher(f func(string) Teacher, list []string) []Teacher {
	if f == nil {
		return []Teacher{}
	}
	newList := make([]Teacher, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapStrTeacher takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapStrTeacherPtr(f func(*string) *Teacher, list []*string) []*Teacher {
	if f == nil {
		return []*Teacher{}
	}
	newList := make([]*Teacher, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// PMapStrTeacher applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: string output type: Teacher
//	2. List
//
// Returns
//	New List of type Teacher
//	Empty list if all arguments are nil or either one is nil
func PMapStrTeacher(f func(string) Teacher, list []string) []Teacher {
	if f == nil {
		return []Teacher{}
	}

	ch := make(chan map[int]Teacher)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]Teacher, i int, v string) {
			defer wg.Done()
			ch <- map[int]Teacher{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]Teacher, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// FilterMapStrTeacher filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - string and returns true/false.
//	2. Function: takes string as argument and returns Teacher
// 	3. List of type string
//
// Returns:
//	New List of type Teacher
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrTeacher(fFilter func(string) bool, fMap func(string) Teacher, list []string) []Teacher {
	if fFilter == nil || fMap == nil {
		return []Teacher{}
	}
	var newList []Teacher
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

