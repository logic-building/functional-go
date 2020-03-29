package fp

import (
	"testing"
)

func TestTakeWhileInt(t *testing.T) {
	// Test : Take the numbers as long as condition match
	expectedNewList := []int{4, 2, 4}
	NewList := TakeWhileInt(isEven, []int{4, 2, 4, 7, 5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("TakeWhileInt failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []int{40}
	partialIsEvenDivisibleBy := func(num int) bool { return isEvenDivisibleBy(num, 10) }
	NewList = TakeWhileInt(partialIsEvenDivisibleBy, []int{40})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("TakeWhileInt failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(TakeWhileInt(nil, nil)) > 0 {
		t.Errorf("TakeWhileInt failed.")
	}

	if len(TakeWhileInt(nil, []int{})) > 0 {
		t.Errorf("TakeWhileInt failed.")
	}
}

func TestTakeWhileInt64(t *testing.T) {
	// Test : Take the numbers as long as condition match
	expectedNewList := []int64{4, 2, 4}
	NewList := TakeWhileInt64(isEvenInt64, []int64{4, 2, 4, 7, 5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("TakeWhileInt64 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []int64{40}
	partialIsEvenDivisibleBy := func(num int64) bool { return isEvenInt64DivisibleBy(num, 10) }
	NewList = TakeWhileInt64(partialIsEvenDivisibleBy, []int64{40})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("TakeWhileInt64 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(TakeWhileInt64(nil, nil)) > 0 {
		t.Errorf("TakeWhileInt64 failed.")
	}

	if len(TakeWhileInt64(nil, []int64{})) > 0 {
		t.Errorf("TakeWhileInt64 failed.")
	}
}

func TestTakeWhileInt32(t *testing.T) {
	// Test : Take the numbers as long as condition match
	expectedNewList := []int32{4, 2, 4}
	NewList := TakeWhileInt32(isEvenInt32, []int32{4, 2, 4, 7, 5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("TakeWhileInt32 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []int32{40}
	partialIsEvenDivisibleBy := func(num int32) bool { return isEvenInt32DivisibleBy(num, 10) }
	NewList = TakeWhileInt32(partialIsEvenDivisibleBy, []int32{40})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("TakeWhileInt32 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(TakeWhileInt32(nil, nil)) > 0 {
		t.Errorf("TakeWhileInt32 failed.")
	}

	if len(TakeWhileInt32(nil, []int32{})) > 0 {
		t.Errorf("TakeWhileInt32 failed.")
	}
}

func TestTakeWhileInt16(t *testing.T) {
	// Test : Take the numbers as long as condition match
	expectedNewList := []int16{4, 2, 4}
	NewList := TakeWhileInt16(isEvenInt16, []int16{4, 2, 4, 7, 5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("TakeWhileInt32 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []int16{40}
	partialIsEvenDivisibleBy := func(num int16) bool { return isEvenInt16DivisibleBy(num, 10) }
	NewList = TakeWhileInt16(partialIsEvenDivisibleBy, []int16{40})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("TakeWhileInt16 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(TakeWhileInt16(nil, nil)) > 0 {
		t.Errorf("TakeWhileInt16 failed.")
	}

	if len(TakeWhileInt16(nil, []int16{})) > 0 {
		t.Errorf("TakeWhileInt16 failed.")
	}
}

func TestTakeWhileInt8(t *testing.T) {
	// Test : Take the numbers as long as condition match
	expectedNewList := []int8{4, 2, 4}
	NewList := TakeWhileInt8(isEvenInt8, []int8{4, 2, 4, 7, 5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("TakeWhileInt8 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []int8{40}
	partialIsEvenDivisibleBy := func(num int8) bool { return isEvenInt8DivisibleBy(num, 10) }
	NewList = TakeWhileInt8(partialIsEvenDivisibleBy, []int8{40})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("TakeWhileInt8 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(TakeWhileInt8(nil, nil)) > 0 {
		t.Errorf("TakeWhileInt8 failed.")
	}

	if len(TakeWhileInt8(nil, []int8{})) > 0 {
		t.Errorf("TakeWhileInt8 failed.")
	}
}

func TestTakeWhileUInt(t *testing.T) {
	// Test : Take the numbers as long as condition match
	expectedNewList := []uint{4, 2, 4}
	NewList := TakeWhileUint(isEvenUint, []uint{4, 2, 4, 7, 5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("TakeWhileUint failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []uint{40}
	partialIsEvenDivisibleBy := func(num uint) bool { return isEvenUintDivisibleBy(num, 10) }
	NewList = TakeWhileUint(partialIsEvenDivisibleBy, []uint{40})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("TakeWhileUint failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(TakeWhileUint(nil, nil)) > 0 {
		t.Errorf("TakeWhileUint failed.")
	}

	if len(TakeWhileUint(nil, []uint{})) > 0 {
		t.Errorf("TakeWhileUint failed.")
	}
}

func TestTakeWhileUInt64(t *testing.T) {
	// Test : Take the numbers as long as condition match
	expectedNewList := []uint64{4, 2, 4}
	NewList := TakeWhileUint64(isEvenUint64, []uint64{4, 2, 4, 7, 5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("TakeWhileUint64 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []uint64{40}
	partialIsEvenDivisibleBy := func(num uint64) bool { return isEvenUint64DivisibleBy(num, 10) }
	NewList = TakeWhileUint64(partialIsEvenDivisibleBy, []uint64{40})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("TakeWhileUint64 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(TakeWhileUint64(nil, nil)) > 0 {
		t.Errorf("TakeWhileUint64 failed.")
	}

	if len(TakeWhileUint64(nil, []uint64{})) > 0 {
		t.Errorf("TakeWhileUint64 failed.")
	}
}

func TestTakeWhileUInt32(t *testing.T) {
	// Test : Take the numbers as long as condition match
	expectedNewList := []uint32{4, 2, 4}
	NewList := TakeWhileUint32(isEvenUint32, []uint32{4, 2, 4, 7, 5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("TakeWhileUint32 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []uint32{40}
	partialIsEvenDivisibleBy := func(num uint32) bool { return isEvenUint32DivisibleBy(num, 10) }
	NewList = TakeWhileUint32(partialIsEvenDivisibleBy, []uint32{40})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("TakeWhileUint32 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(TakeWhileUint32(nil, nil)) > 0 {
		t.Errorf("TakeWhileUint32 failed.")
	}

	if len(TakeWhileUint32(nil, []uint32{})) > 0 {
		t.Errorf("TakeWhileUint32 failed.")
	}
}

func TestTakeWhileUInt16(t *testing.T) {
	// Test : Take the numbers as long as condition match
	expectedNewList := []uint16{4, 2, 4}
	NewList := TakeWhileUint16(isEvenUint16, []uint16{4, 2, 4, 7, 5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("TakeWhileUint16 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []uint16{40}
	partialIsEvenDivisibleBy := func(num uint16) bool { return isEvenUint16DivisibleBy(num, 10) }
	NewList = TakeWhileUint16(partialIsEvenDivisibleBy, []uint16{40})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("TakeWhileUint16 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(TakeWhileUint16(nil, nil)) > 0 {
		t.Errorf("TakeWhileUint16 failed.")
	}

	if len(TakeWhileUint16(nil, []uint16{})) > 0 {
		t.Errorf("TakeWhileUint16 failed.")
	}
}

func TestTakeWhileUInt8(t *testing.T) {
	// Test : Take the numbers as long as condition match
	expectedNewList := []uint8{4, 2, 4}
	NewList := TakeWhileUint8(isEvenUint8, []uint8{4, 2, 4, 7, 5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("TakeWhileUint8 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []uint8{40}
	partialIsEvenDivisibleBy := func(num uint8) bool { return isEvenUint8DivisibleBy(num, 10) }
	NewList = TakeWhileUint8(partialIsEvenDivisibleBy, []uint8{40})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("TakeWhileUint8 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(TakeWhileUint8(nil, nil)) > 0 {
		t.Errorf("TakeWhileUint8 failed.")
	}

	if len(TakeWhileUint8(nil, []uint8{})) > 0 {
		t.Errorf("TakeWhileUint8 failed.")
	}
}

func TestTakeWhileFloat64(t *testing.T) {
	// Test : Take the numbers as long as condition match
	expectedNewList := []float64{4, 2, 3}
	NewList := TakeWhileFloat64(isPositiveFloat64, []float64{4, 2, 3, -1, 4, 5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("TakeWhileFloat64 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []float64{4, 2, 3}
	NewList = TakeWhileFloat64(isPositiveFloat64, []float64{4, 2, 3})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("TakeWhileFloat64 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(TakeWhileFloat64(nil, nil)) > 0 {
		t.Errorf("TakeWhileFloat64 failed.")
	}

	if len(TakeWhileFloat64(nil, []float64{})) > 0 {
		t.Errorf("TakeWhileFloat64 failed.")
	}
}

func TestTakeWhileFloat32(t *testing.T) {
	// Test : Take the numbers as long as condition match
	expectedNewList := []float32{4, 2, 3}
	NewList := TakeWhileFloat32(isPositiveFloat32, []float32{4, 2, 3, -1, 4, 5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("TakeWhileFloat32 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []float32{4, 2, 3}
	NewList = TakeWhileFloat32(isPositiveFloat32, []float32{4, 2, 3})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("TakeWhileFloat32 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(TakeWhileFloat32(nil, nil)) > 0 {
		t.Errorf("TakeWhileFloat32 failed.")
	}

	if len(TakeWhileFloat32(nil, []float32{})) > 0 {
		t.Errorf("TakeWhileFloat64 failed.")
	}
}

func TestTakeWhileStr(t *testing.T) {
	// Test : Take the numbers as long as condition match
	expectedNewList := []string{"Ram", "Shyam", "Radha"}
	NewList := TakeWhileStr(
		func(str string) bool {
			return len(str) < 10
		},
		[]string{"Ram", "Shyam", "Radha", "Nandeshwar", "ShyamSundar", "Hari Shankar"})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("TakeWhileStr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []string{"Ram"}
	NewList = TakeWhileStr(
		func(str string) bool {
			return len(str) < 10
		},
		[]string{"Ram", "Shyam", "Radha"})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("TakeWhileStr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(TakeWhileStr(nil, nil)) > 0 {
		t.Errorf("TakeWhileStr failed.")
	}

	if len(TakeWhileStr(nil, []string{})) > 0 {
		t.Errorf("TakeWhileStr failed.")
	}
}

func TestTakeWhileBool(t *testing.T) {
	// Test : Take the numbers as long as condition match
	var vt bool = true
	var vf bool = false

	expectedNewList := []bool{vt, vt, vf}
	NewList := TakeWhileBool(func(v bool) bool { return v == true }, []bool{vt, vt, vf, vf, vf})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("TakeWhileBool failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []bool{vt}
	NewList = TakeWhileBool(func(v bool) bool { return v == true }, []bool{vt})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("TakeWhileBoolPtr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(TakeWhileBoolPtr(nil, nil)) > 0 {
		t.Errorf("TakeWhileBoolPtr failed.")
	}

	if len(TakeWhileBool(nil, []bool{})) > 0 {
		t.Errorf("TakeWhileBoolPtr failed.")
	}
}
