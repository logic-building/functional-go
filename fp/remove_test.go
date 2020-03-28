package fp

import "testing"

func TestRemoveInt(t *testing.T) {
	// Test : even number in the list
	expectedNewList := []int{1, 3}
	NewList := RemoveInt(isEven, []int{1, 2, 3, 4})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveInt failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []int{1, 3}
	partialIsEven := func(num int) bool { return isEvenDivisibleBy(num, 10) }
	NewList = RemoveInt(partialIsEven, []int{20, 1, 3, 40})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveInt failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(RemoveInt(nil, nil)) > 0 {
		t.Errorf("RemoveInt failed.")
	}
}

func TestRemoveInt64(t *testing.T) {
	// Test : even number in the list
	expectedNewList := []int64{1, 3}
	NewList := RemoveInt64(isEvenInt64, []int64{1, 2, 3, 4})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveInt64 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []int64{1, 3}
	partialIsEvenInt64 := func(num int64) bool { return isEvenInt64DivisibleBy(num, 10) }
	NewList = RemoveInt64(partialIsEvenInt64, []int64{20, 1, 3, 40})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveInt64 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
	if len(RemoveInt64(nil, nil)) > 0 {
		t.Errorf("RemoveInt64 failed.")
	}
}

func BenchmarkRemoveInt64(b *testing.B) {
	b.N = iterations
	list := make([]int64, size)
	var i int64
	for i = 0; i < size; i++ {
		list[i] = i
	}
	for i := 0; i < b.N; i++ {
		RemoveInt64(isEvenInt64, list)
	}
}

func TestRemoveInt32(t *testing.T) {
	// Test : even number in the list
	expectedNewList := []int32{1, 3}
	NewList := RemoveInt32(isEvenInt32, []int32{1, 2, 3, 4})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveInt32 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []int32{1, 3}
	partialIsEvenInt32 := func(num int32) bool { return isEvenInt32DivisibleBy(num, 10) }
	NewList = RemoveInt32(partialIsEvenInt32, []int32{20, 1, 3, 40})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveInt32 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
	if len(RemoveInt32(nil, nil)) > 0 {
		t.Errorf("RemoveInt32 failed.")
	}
}

func TestRemoveInt16(t *testing.T) {
	// Test : even number in the list
	expectedNewList := []int16{1, 3}
	NewList := RemoveInt16(isEvenInt16, []int16{1, 2, 3, 4})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveInt16 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []int16{1, 3}
	partialIsEven := func(num int16) bool { return isEvenInt16DivisibleBy(num, 10) }
	NewList = RemoveInt16(partialIsEven, []int16{20, 1, 3, 40})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveInt16 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
	if len(RemoveInt16(nil, nil)) > 0 {
		t.Errorf("RemoveInt16 failed.")
	}
}

func TestRemoveInt8(t *testing.T) {
	// Test : even number in the list
	expectedNewList := []int8{1, 3}
	NewList := RemoveInt8(isEvenInt8, []int8{1, 2, 3, 4})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveInt8 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []int8{1, 3}
	partialIsEvenInt8 := func(num int8) bool { return isEvenInt8DivisibleBy(num, 10) }
	NewList = RemoveInt8(partialIsEvenInt8, []int8{20, 1, 3, 40})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveInt8 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
	if len(RemoveInt8(nil, nil)) > 0 {
		t.Errorf("RemoveInt8 failed.")
	}
}

func TestRemoveUint64(t *testing.T) {
	// Test : even number in the list
	expectedNewList := []uint64{1, 3}
	NewList := RemoveUint64(isEvenUint64, []uint64{1, 2, 3, 4})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveUint64 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []uint64{1, 3}
	partialIsEvenUint64 := func(num uint64) bool { return isEvenUint64DivisibleBy(num, 10) }
	NewList = RemoveUint64(partialIsEvenUint64, []uint64{20, 1, 3, 40})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveUnt64 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
	if len(RemoveUint64(nil, nil)) > 0 {
		t.Errorf("RemoveUint64 failed.")
	}
}

func TestRemoveUint32(t *testing.T) {
	// Test : even number in the list
	expectedNewList := []uint32{1, 3}
	NewList := RemoveUint32(isEvenUint32, []uint32{1, 2, 3, 4})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveUint32 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []uint32{1, 3}
	partialIsEvenUint32 := func(num uint32) bool { return isEvenUint32DivisibleBy(num, 10) }
	NewList = RemoveUint32(partialIsEvenUint32, []uint32{20, 1, 3, 40})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveUint32 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
	if len(RemoveUint32(nil, nil)) > 0 {
		t.Errorf("RemoveUint32 failed.")
	}
}

func TestRemoveUint16(t *testing.T) {
	// Test : even number in the list
	expectedNewList := []uint16{1, 3}
	NewList := RemoveUint16(isEvenUint16, []uint16{1, 2, 3, 4})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveUint16 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []uint16{1, 3}
	partialIsEvenUint16 := func(num uint16) bool { return isEvenUint16DivisibleBy(num, 10) }
	NewList = RemoveUint16(partialIsEvenUint16, []uint16{20, 1, 3, 40})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveUint16 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
	if len(RemoveUint16(nil, nil)) > 0 {
		t.Errorf("RemoveUint16 failed.")
	}
}

func TestRemoveUint8(t *testing.T) {
	// Test : even number in the list
	expectedNewList := []uint8{1, 3}
	NewList := RemoveUint8(isEvenUint8, []uint8{1, 2, 3, 4})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveUint8 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []uint8{1, 3}
	partialIsEvenUint8 := func(num uint8) bool { return isEvenUint8DivisibleBy(num, 10) }
	NewList = RemoveUint8(partialIsEvenUint8, []uint8{20, 1, 3, 40})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveUint16 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
	if len(RemoveUint8(nil, nil)) > 0 {
		t.Errorf("RemoveUint16 failed.")
	}
}

func TestRemoveUint(t *testing.T) {
	// Test : even number in the list
	expectedNewList := []uint{1, 3}
	NewList := RemoveUint(isEvenUint, []uint{1, 2, 3, 4})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveUint failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	// Test: Remove all even numbers divisible by 10 in the list
	expectedNewList = []uint{1, 3}
	partialIsEvenUint := func(num uint) bool { return isEvenUintDivisibleBy(num, 10) }
	NewList = RemoveUint(partialIsEvenUint, []uint{20, 1, 3, 40})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveUint failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
	if len(RemoveUint(nil, nil)) > 0 {
		t.Errorf("RemoveUint16 failed.")
	}
}

func TestRemoveFloat64(t *testing.T) {
	// Test : Remove all the positive numbers in the list
	expectedNewList := []float64{0, -2}
	NewList := RemoveFloat64(isPositiveFloat64, []float64{0, -2, 2, 40.50})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveFloat64 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
	if len(RemoveFloat64(nil, nil)) > 0 {
		t.Errorf("RemoveFloat64 failed.")
	}
}

func TestRemoveFloat32(t *testing.T) {
	// Test : Remove all the positive numbers in the list
	expectedNewList := []float32{0, -2}
	NewList := RemoveFloat32(isPositiveFloat32, []float32{0, -2, 2, 40.50})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("RemoveFloat32 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
	if len(RemoveFloat32(nil, nil)) > 0 {
		t.Errorf("RemoveFloat32 failed.")
	}
}

func TestRemoveStr(t *testing.T) {
	// Test : Remove all the positive numbers in the list
	expectedNewList := []string{"Nandeshwar"}
	NewList := RemoveStr(isStringLengthLessThan10, []string{"gopal", "govinda", "Nandeshwar"})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("FiltrStr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
	if len(RemoveStr(nil, nil)) > 0 {
		t.Errorf("RemoveStr failed.")
	}
}

func TestRemoveBool(t *testing.T) {
	var vt bool = true
	var vf bool = false

	expectedNewList := []bool{vt}
	NewList := RemoveBool(func(v bool) bool { return v == false }, []bool{vt, vf, vf})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("RemoveBool failed. Expected New list=%v, actual list=%v", expectedNewList[0], NewList[0])
	}

	if len(RemoveBoolPtr(nil, nil)) > 0 {
		t.Errorf("RemoveBoolPtr failed.")
	}
}
