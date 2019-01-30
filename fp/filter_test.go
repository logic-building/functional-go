package fp

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

func TestFilterUint64(t *testing.T) {
	// Test : even number in the list
	expectedFilteredList := []uint64{2, 4}
	filteredList := FilterUint64(isEvenUint64, []uint64{1, 2, 3, 4})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilterUint64 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	// Test: filter all even numbers divisible by 10 in the list
	expectedFilteredList = []uint64{20, 40}
	partialIsEvenUint64 := func(num uint64) bool { return isEvenUint64DivisibleBy(num, 10) }
	filteredList = FilterUint64(partialIsEvenUint64, []uint64{20, 1, 3, 40})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilterUnt64 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterUint64(nil, nil)) > 0 {
		t.Errorf("FilterUint64 failed.")
	}
}

func isEvenUint64(num uint64) bool {
	return num%2 == 0
}

func isEvenUint64DivisibleBy(num, divisibleBy uint64) bool {
	return num%2 == 0 && num%divisibleBy == 0
}

func TestFilterUint32(t *testing.T) {
	// Test : even number in the list
	expectedFilteredList := []uint32{2, 4}
	filteredList := FilterUint32(isEvenUint32, []uint32{1, 2, 3, 4})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilterUint32 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	// Test: filter all even numbers divisible by 10 in the list
	expectedFilteredList = []uint32{20, 40}
	partialIsEvenUint32 := func(num uint32) bool { return isEvenUint32DivisibleBy(num, 10) }
	filteredList = FilterUint32(partialIsEvenUint32, []uint32{20, 1, 3, 40})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilterUint32 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterUint32(nil, nil)) > 0 {
		t.Errorf("FilterUint32 failed.")
	}
}

func isEvenUint32(num uint32) bool {
	return num%2 == 0
}

func isEvenUint32DivisibleBy(num, divisibleBy uint32) bool {
	return num%2 == 0 && num%divisibleBy == 0
}

func TestFilterUint16(t *testing.T) {
	// Test : even number in the list
	expectedFilteredList := []uint16{2, 4}
	filteredList := FilterUint16(isEvenUint16, []uint16{1, 2, 3, 4})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilterUint16 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	// Test: filter all even numbers divisible by 10 in the list
	expectedFilteredList = []uint16{20, 40}
	partialIsEvenUint16 := func(num uint16) bool { return isEvenUint16DivisibleBy(num, 10) }
	filteredList = FilterUint16(partialIsEvenUint16, []uint16{20, 1, 3, 40})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilterUint16 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterUint16(nil, nil)) > 0 {
		t.Errorf("FilterUint16 failed.")
	}
}

func isEvenUint16(num uint16) bool {
	return num%2 == 0
}

func isEvenUint16DivisibleBy(num, divisibleBy uint16) bool {
	return num%2 == 0 && num%divisibleBy == 0
}

func TestFilterUint8(t *testing.T) {
	// Test : even number in the list
	expectedFilteredList := []uint8{2, 4}
	filteredList := FilterUint8(isEvenUint8, []uint8{1, 2, 3, 4})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilterUint8 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	// Test: filter all even numbers divisible by 10 in the list
	expectedFilteredList = []uint8{20, 40}
	partialIsEvenUint8 := func(num uint8) bool { return isEvenUint8DivisibleBy(num, 10) }
	filteredList = FilterUint8(partialIsEvenUint8, []uint8{20, 1, 3, 40})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilterUint16 failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterUint8(nil, nil)) > 0 {
		t.Errorf("FilterUint16 failed.")
	}
}

func isEvenUint8(num uint8) bool {
	return num%2 == 0
}

func isEvenUint8DivisibleBy(num, divisibleBy uint8) bool {
	return num%2 == 0 && num%divisibleBy == 0
}

func TestFilterUint(t *testing.T) {
	// Test : even number in the list
	expectedFilteredList := []uint{2, 4}
	filteredList := FilterUint(isEvenUint, []uint{1, 2, 3, 4})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilterUint failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	// Test: filter all even numbers divisible by 10 in the list
	expectedFilteredList = []uint{20, 40}
	partialIsEvenUint := func(num uint) bool { return isEvenUintDivisibleBy(num, 10) }
	filteredList = FilterUint(partialIsEvenUint, []uint{20, 1, 3, 40})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilterUint failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterUint(nil, nil)) > 0 {
		t.Errorf("FilterUint16 failed.")
	}
}

func isEvenUint(num uint) bool {
	return num%2 == 0
}

func isEvenUintDivisibleBy(num, divisibleBy uint) bool {
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
	if len(FilterFloat32(nil, nil)) > 0 {
		t.Errorf("FilterFloat32 failed.")
	}
}

func isPositiveFloat32(num float32) bool {
	return num > 0
}

func TestFilterStr(t *testing.T) {
	// Test : filter all the name which length is less than 10
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
