package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestFilterIntPtr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v10 int = 10
	var v20 int = 20
	var v40 int = 40


	// Test : even number in the list
	expectedFilteredList := []*int{&v2, &v4}
	filteredList := FilterIntPtr(isEvenIntPtr, []*int{&v1, &v2, &v3, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("MapFilter failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	// Test: filter all even numbers divisible by 10 in the list
	expectedFilteredList = []*int{&v20, &v40}
	partialIsEven := func(num *int) bool { return isEvenDivisibleByIntPtr(num, &v10) }
	filteredList = FilterIntPtr(partialIsEven, []*int{&v20, &v1, &v3, &v40})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilter failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	if len(FilterIntPtr(nil, nil)) > 0 {
		t.Errorf("FilterInt failed.")
	}
}

func isEvenIntPtr(num *int) bool {
	return *num%2 == 0
}

func isEvenDivisibleByIntPtr(num, divisibleBy *int) bool {
	return *num%2 == 0 && *num % *divisibleBy == 0
}


func TestFilterInt64Ptr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v10 int64 = 10
	var v20 int64 = 20
	var v40 int64 = 40


	// Test : even number in the list
	expectedFilteredList := []*int64{&v2, &v4}
	filteredList := FilterInt64Ptr(isEvenInt64Ptr, []*int64{&v1, &v2, &v3, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("MapFilter failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	// Test: filter all even numbers divisible by 10 in the list
	expectedFilteredList = []*int64{&v20, &v40}
	partialIsEven := func(num *int64) bool { return isEvenDivisibleByInt64Ptr(num, &v10) }
	filteredList = FilterInt64Ptr(partialIsEven, []*int64{&v20, &v1, &v3, &v40})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilter failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	if len(FilterInt64Ptr(nil, nil)) > 0 {
		t.Errorf("FilterInt failed.")
	}
}

func isEvenInt64Ptr(num *int64) bool {
	return *num%2 == 0
}

func isEvenDivisibleByInt64Ptr(num, divisibleBy *int64) bool {
	return *num%2 == 0 && *num % *divisibleBy == 0
}


func TestFilterInt32Ptr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v10 int32 = 10
	var v20 int32 = 20
	var v40 int32 = 40


	// Test : even number in the list
	expectedFilteredList := []*int32{&v2, &v4}
	filteredList := FilterInt32Ptr(isEvenInt32Ptr, []*int32{&v1, &v2, &v3, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("MapFilter failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	// Test: filter all even numbers divisible by 10 in the list
	expectedFilteredList = []*int32{&v20, &v40}
	partialIsEven := func(num *int32) bool { return isEvenDivisibleByInt32Ptr(num, &v10) }
	filteredList = FilterInt32Ptr(partialIsEven, []*int32{&v20, &v1, &v3, &v40})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilter failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	if len(FilterInt32Ptr(nil, nil)) > 0 {
		t.Errorf("FilterInt failed.")
	}
}

func isEvenInt32Ptr(num *int32) bool {
	return *num%2 == 0
}

func isEvenDivisibleByInt32Ptr(num, divisibleBy *int32) bool {
	return *num%2 == 0 && *num % *divisibleBy == 0
}


func TestFilterInt16Ptr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v10 int16 = 10
	var v20 int16 = 20
	var v40 int16 = 40


	// Test : even number in the list
	expectedFilteredList := []*int16{&v2, &v4}
	filteredList := FilterInt16Ptr(isEvenInt16Ptr, []*int16{&v1, &v2, &v3, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("MapFilter failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	// Test: filter all even numbers divisible by 10 in the list
	expectedFilteredList = []*int16{&v20, &v40}
	partialIsEven := func(num *int16) bool { return isEvenDivisibleByInt16Ptr(num, &v10) }
	filteredList = FilterInt16Ptr(partialIsEven, []*int16{&v20, &v1, &v3, &v40})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilter failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	if len(FilterInt16Ptr(nil, nil)) > 0 {
		t.Errorf("FilterInt failed.")
	}
}

func isEvenInt16Ptr(num *int16) bool {
	return *num%2 == 0
}

func isEvenDivisibleByInt16Ptr(num, divisibleBy *int16) bool {
	return *num%2 == 0 && *num % *divisibleBy == 0
}


func TestFilterInt8Ptr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v10 int8 = 10
	var v20 int8 = 20
	var v40 int8 = 40


	// Test : even number in the list
	expectedFilteredList := []*int8{&v2, &v4}
	filteredList := FilterInt8Ptr(isEvenInt8Ptr, []*int8{&v1, &v2, &v3, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("MapFilter failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	// Test: filter all even numbers divisible by 10 in the list
	expectedFilteredList = []*int8{&v20, &v40}
	partialIsEven := func(num *int8) bool { return isEvenDivisibleByInt8Ptr(num, &v10) }
	filteredList = FilterInt8Ptr(partialIsEven, []*int8{&v20, &v1, &v3, &v40})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilter failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	if len(FilterInt8Ptr(nil, nil)) > 0 {
		t.Errorf("FilterInt failed.")
	}
}

func isEvenInt8Ptr(num *int8) bool {
	return *num%2 == 0
}

func isEvenDivisibleByInt8Ptr(num, divisibleBy *int8) bool {
	return *num%2 == 0 && *num % *divisibleBy == 0
}


func TestFilterUintPtr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v10 uint = 10
	var v20 uint = 20
	var v40 uint = 40


	// Test : even number in the list
	expectedFilteredList := []*uint{&v2, &v4}
	filteredList := FilterUintPtr(isEvenUintPtr, []*uint{&v1, &v2, &v3, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("MapFilter failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	// Test: filter all even numbers divisible by 10 in the list
	expectedFilteredList = []*uint{&v20, &v40}
	partialIsEven := func(num *uint) bool { return isEvenDivisibleByUintPtr(num, &v10) }
	filteredList = FilterUintPtr(partialIsEven, []*uint{&v20, &v1, &v3, &v40})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilter failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	if len(FilterUintPtr(nil, nil)) > 0 {
		t.Errorf("FilterInt failed.")
	}
}

func isEvenUintPtr(num *uint) bool {
	return *num%2 == 0
}

func isEvenDivisibleByUintPtr(num, divisibleBy *uint) bool {
	return *num%2 == 0 && *num % *divisibleBy == 0
}


func TestFilterUint64Ptr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v10 uint64 = 10
	var v20 uint64 = 20
	var v40 uint64 = 40


	// Test : even number in the list
	expectedFilteredList := []*uint64{&v2, &v4}
	filteredList := FilterUint64Ptr(isEvenUint64Ptr, []*uint64{&v1, &v2, &v3, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("MapFilter failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	// Test: filter all even numbers divisible by 10 in the list
	expectedFilteredList = []*uint64{&v20, &v40}
	partialIsEven := func(num *uint64) bool { return isEvenDivisibleByUint64Ptr(num, &v10) }
	filteredList = FilterUint64Ptr(partialIsEven, []*uint64{&v20, &v1, &v3, &v40})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilter failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	if len(FilterUint64Ptr(nil, nil)) > 0 {
		t.Errorf("FilterInt failed.")
	}
}

func isEvenUint64Ptr(num *uint64) bool {
	return *num%2 == 0
}

func isEvenDivisibleByUint64Ptr(num, divisibleBy *uint64) bool {
	return *num%2 == 0 && *num % *divisibleBy == 0
}


func TestFilterUint32Ptr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v10 uint32 = 10
	var v20 uint32 = 20
	var v40 uint32 = 40


	// Test : even number in the list
	expectedFilteredList := []*uint32{&v2, &v4}
	filteredList := FilterUint32Ptr(isEvenUint32Ptr, []*uint32{&v1, &v2, &v3, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("MapFilter failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	// Test: filter all even numbers divisible by 10 in the list
	expectedFilteredList = []*uint32{&v20, &v40}
	partialIsEven := func(num *uint32) bool { return isEvenDivisibleByUint32Ptr(num, &v10) }
	filteredList = FilterUint32Ptr(partialIsEven, []*uint32{&v20, &v1, &v3, &v40})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilter failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	if len(FilterUint32Ptr(nil, nil)) > 0 {
		t.Errorf("FilterInt failed.")
	}
}

func isEvenUint32Ptr(num *uint32) bool {
	return *num%2 == 0
}

func isEvenDivisibleByUint32Ptr(num, divisibleBy *uint32) bool {
	return *num%2 == 0 && *num % *divisibleBy == 0
}


func TestFilterUint16Ptr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v10 uint16 = 10
	var v20 uint16 = 20
	var v40 uint16 = 40


	// Test : even number in the list
	expectedFilteredList := []*uint16{&v2, &v4}
	filteredList := FilterUint16Ptr(isEvenUint16Ptr, []*uint16{&v1, &v2, &v3, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("MapFilter failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	// Test: filter all even numbers divisible by 10 in the list
	expectedFilteredList = []*uint16{&v20, &v40}
	partialIsEven := func(num *uint16) bool { return isEvenDivisibleByUint16Ptr(num, &v10) }
	filteredList = FilterUint16Ptr(partialIsEven, []*uint16{&v20, &v1, &v3, &v40})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilter failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	if len(FilterUint16Ptr(nil, nil)) > 0 {
		t.Errorf("FilterInt failed.")
	}
}

func isEvenUint16Ptr(num *uint16) bool {
	return *num%2 == 0
}

func isEvenDivisibleByUint16Ptr(num, divisibleBy *uint16) bool {
	return *num%2 == 0 && *num % *divisibleBy == 0
}


func TestFilterUint8Ptr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v10 uint8 = 10
	var v20 uint8 = 20
	var v40 uint8 = 40


	// Test : even number in the list
	expectedFilteredList := []*uint8{&v2, &v4}
	filteredList := FilterUint8Ptr(isEvenUint8Ptr, []*uint8{&v1, &v2, &v3, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("MapFilter failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	// Test: filter all even numbers divisible by 10 in the list
	expectedFilteredList = []*uint8{&v20, &v40}
	partialIsEven := func(num *uint8) bool { return isEvenDivisibleByUint8Ptr(num, &v10) }
	filteredList = FilterUint8Ptr(partialIsEven, []*uint8{&v20, &v1, &v3, &v40})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilter failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	if len(FilterUint8Ptr(nil, nil)) > 0 {
		t.Errorf("FilterInt failed.")
	}
}

func isEvenUint8Ptr(num *uint8) bool {
	return *num%2 == 0
}

func isEvenDivisibleByUint8Ptr(num, divisibleBy *uint8) bool {
	return *num%2 == 0 && *num % *divisibleBy == 0
}


func TestFilterStrPtr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"

	// Test : even number in the list
	expectedFilteredList := []*string{&v2, &v4}
	filteredList := FilterStrPtr(isEvenStrPtr, []*string{&v1, &v2, &v3, &v4})

	if *filteredList[0] != *expectedFilteredList[0] {
		t.Errorf("MapFilter failed. Expected filtered list=%v, actual list=%v", *expectedFilteredList[0], *filteredList[0])
	}

	if len(FilterStrPtr(nil, nil)) > 0 {
		t.Errorf("FilterInt failed.")
        t.Errorf(reflect.String.String())
	}
}

func isEvenStrPtr(num *string) bool {
	if *num == "2" || *num == "4" || *num == "6" || *num == "8" || *num == "10" {
		return true
	}
	return false
}


func TestFilterBoolPtr(t *testing.T) {
	var vt bool = true

	expectedSumList := []*bool{&vt}
	
	newList := FilterBoolPtr(trueBoolPtr, []*bool{&vt})
	if *newList[0] != *expectedSumList[0]  {
		t.Errorf("FilterBoolPtr failed")
	}

	if len(FilterBoolPtr(nil, nil)) > 0 {
		t.Errorf("MapBoolPtr failed.")
	}
}

func trueBoolPtr(num1 *bool) bool {
	return true
}


func TestFilterFloat32Ptr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v10 float32 = 10
	var v20 float32 = 20
	var v40 float32 = 40


	// Test : even number in the list
	expectedFilteredList := []*float32{&v2, &v4}
	filteredList := FilterFloat32Ptr(isEvenFloat32Ptr, []*float32{&v1, &v2, &v3, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("MapFilter failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	// Test: filter all even numbers divisible by 10 in the list
	expectedFilteredList = []*float32{&v20, &v40}
	partialIsEven := func(num *float32) bool { return isEvenDivisibleByFloat32Ptr(num, &v10) }
	filteredList = FilterFloat32Ptr(partialIsEven, []*float32{&v20, &v1, &v3, &v40})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilter failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	if len(FilterFloat32Ptr(nil, nil)) > 0 {
		t.Errorf("FilterInt failed.")
	}
}

func isEvenFloat32Ptr(num *float32) bool {
		return int(*num)%2 == 0
	}
	
	func isEvenDivisibleByFloat32Ptr(num, divisibleBy *float32) bool {
		return int(*num)%2 == 0 && int(*num) % int(*divisibleBy) == 0
	}


func TestFilterFloat64Ptr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v10 float64 = 10
	var v20 float64 = 20
	var v40 float64 = 40


	// Test : even number in the list
	expectedFilteredList := []*float64{&v2, &v4}
	filteredList := FilterFloat64Ptr(isEvenFloat64Ptr, []*float64{&v1, &v2, &v3, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("MapFilter failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	// Test: filter all even numbers divisible by 10 in the list
	expectedFilteredList = []*float64{&v20, &v40}
	partialIsEven := func(num *float64) bool { return isEvenDivisibleByFloat64Ptr(num, &v10) }
	filteredList = FilterFloat64Ptr(partialIsEven, []*float64{&v20, &v1, &v3, &v40})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilter failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	if len(FilterFloat64Ptr(nil, nil)) > 0 {
		t.Errorf("FilterInt failed.")
	}
}

func isEvenFloat64Ptr(num *float64) bool {
	return int(*num)%2 == 0
}

func isEvenDivisibleByFloat64Ptr(num, divisibleBy *float64) bool {
	return int(*num)%2 == 0 && int(*num) % int(*divisibleBy) == 0
}


