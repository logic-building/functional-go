package list_op

import "testing"

func TestFilterInt(t *testing.T) {
	// Test : even number in the list
	expectedFilteredList := []int{2, 4}
	filteredList := FilterInt(isEven, []int{1, 2, 3, 4})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilter failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	// Test: filter all even numbers divisible by 10 in the list
	expectedFilteredList = []int{20, 40}
	partialIsEven := func(num int) bool { return isEvenDivisibleBy(num, 10) }
	filteredList = FilterInt(partialIsEven, []int{20, 1, 3, 40})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilter failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	if len(FilterInt(nil, nil)) > 0 {
		t.Errorf("FilterInt failed.")
	}
}

func isEven(num int) bool {
	return num%2 == 0
}

func isEvenDivisibleBy(num, divisibleBy int) bool {
	return num%2 == 0 && num%divisibleBy == 0
}

func TestFilterInt64(t *testing.T) {
	// Test : even number in the list
	expectedFilteredList := []int64{2, 4}
	filteredList := FilterInt64(isEvenInt64, []int64{1, 2, 3, 4})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilterInt64 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	// Test: filter all even numbers divisible by 10 in the list
	expectedFilteredList = []int64{20, 40}
	partialIsEvenInt64 := func(num int64) bool { return isEvenInt64DivisibleBy(num, 10) }
	filteredList = FilterInt64(partialIsEvenInt64, []int64{20, 1, 3, 40})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilterInt64 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterInt64(nil, nil)) > 0 {
		t.Errorf("FilterInt64 failed.")
	}
}

func BenchmarkFilterInt64(b *testing.B) {
	b.N = iterations
	list := make([]int64, size)
	var i int64
	for i = 0; i < size; i++ {
		list[i] = i
	}
	for i := 0; i < b.N; i++ {
		FilterInt64(isEvenInt64, list)
	}
}

func isEvenInt64(num int64) bool {
	return num%2 == 0
}

func isEvenInt64DivisibleBy(num, divisibleBy int64) bool {
	return num%2 == 0 && num%divisibleBy == 0
}

func TestFilterInt32(t *testing.T) {
	// Test : even number in the list
	expectedFilteredList := []int32{2, 4}
	filteredList := FilterInt32(isEvenInt32, []int32{1, 2, 3, 4})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilterInt32 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	// Test: filter all even numbers divisible by 10 in the list
	expectedFilteredList = []int32{20, 40}
	partialIsEvenInt32 := func(num int32) bool { return isEvenInt32DivisibleBy(num, 10) }
	filteredList = FilterInt32(partialIsEvenInt32, []int32{20, 1, 3, 40})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilterInt32 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterInt32(nil, nil)) > 0 {
		t.Errorf("FilterInt32 failed.")
	}
}

func isEvenInt32(num int32) bool {
	return num%2 == 0
}

func isEvenInt32DivisibleBy(num, divisibleBy int32) bool {
	return num%2 == 0 && num%divisibleBy == 0
}

func TestFilterInt16(t *testing.T) {
	// Test : even number in the list
	expectedFilteredList := []int16{2, 4}
	filteredList := FilterInt16(isEvenInt16, []int16{1, 2, 3, 4})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilterInt16 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	// Test: filter all even numbers divisible by 10 in the list
	expectedFilteredList = []int16{20, 40}
	partialIsEven := func(num int16) bool { return isEvenInt16DivisibleBy(num, 10) }
	filteredList = FilterInt16(partialIsEven, []int16{20, 1, 3, 40})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilterInt16 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterInt16(nil, nil)) > 0 {
		t.Errorf("FilterInt16 failed.")
	}
}

func isEvenInt16(num int16) bool {
	return num%2 == 0
}

func isEvenInt16DivisibleBy(num, divisibleBy int16) bool {
	return num%2 == 0 && num%divisibleBy == 0
}

func TestFilterInt8(t *testing.T) {
	// Test : even number in the list
	expectedFilteredList := []int8{2, 4}
	filteredList := FilterInt8(isEvenInt8, []int8{1, 2, 3, 4})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilterInt8 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	// Test: filter all even numbers divisible by 10 in the list
	expectedFilteredList = []int8{20, 40}
	partialIsEvenInt8 := func(num int8) bool { return isEvenInt8DivisibleBy(num, 10) }
	filteredList = FilterInt8(partialIsEvenInt8, []int8{20, 1, 3, 40})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilterInt8 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterInt8(nil, nil)) > 0 {
		t.Errorf("FilterInt8 failed.")
	}
}

func isEvenInt8(num int8) bool {
	return num%2 == 0
}

func isEvenInt8DivisibleBy(num, divisibleBy int8) bool {
	return num%2 == 0 && num%divisibleBy == 0
}

func TestFilterFloat64(t *testing.T) {
	// Test : filter all the positive numbers in the list
	expectedFilteredList := []float64{2, 40.50}
	filteredList := FilterFloat64(isPositiveFloat64, []float64{0, -2, 2, 40.50})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilterFloat64 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterFloat64(nil, nil)) > 0 {
		t.Errorf("FilterFloat64 failed.")
	}
}

func isPositiveFloat64(num float64) bool {
	return num > 0
}

func TestFilterFloat32(t *testing.T) {
	// Test : filter all the positive numbers in the list
	expectedFilteredList := []float32{2, 40.50}
	filteredList := FilterFloat32(isPositiveFloat32, []float32{0, -2, 2, 40.50})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilterFloat32 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
}

func isPositiveFloat32(num float32) bool {
	return num > 0
}

func TestFilterStr(t *testing.T) {
	// Test : filter all the positive numbers in the list
	expectedFilteredList := []string{"gopal", "govinda"}
	filteredList := FilterStr(isStringLengthLessThan10, []string{"gopal", "govinda", "Nandeshwar"})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("FiltrStr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterStr(nil, nil)) > 0 {
		t.Errorf("FilterStr failed.")
	}
}

func isStringLengthLessThan10(str string) bool {
	return len(str) < 10
}
