package fp

import (
	"testing"
)

func TestDropWhileInt(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails
	expectedNewList := []int{3, 4, 5}
	NewList := DropWhileInt(isEven, []int{4, 2, 3, 4, 5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhileInt failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	NewList = DropWhileInt(isEven, []int{4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileInt failed")
	}

	expectedNewList = []int{4}
	partialIsEvenDivisibleBy := func(num int) bool { return isEvenDivisibleBy(num, 10) }
	NewList = DropWhileInt(partialIsEvenDivisibleBy, []int{20, 40, 4})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("DropWhileInt failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(DropWhileInt(nil, nil)) > 0 {
		t.Errorf("DropWhileInt failed.")
	}

	if len(DropWhileInt(nil, []int{})) > 0 {
		t.Errorf("DropWhileInt failed.")
	}
}

func TestDropWhileInt64(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails
	expectedNewList := []int64{3, 4, 5}
	NewList := DropWhileInt64(isEvenInt64, []int64{4, 2, 3, 4, 5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhileInt64 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []int64{4}
	partialIsEvenDivisibleBy := func(num int64) bool { return isEvenInt64DivisibleBy(num, 10) }
	NewList = DropWhileInt64(partialIsEvenDivisibleBy, []int64{20, 40, 4})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("DropWhileInt64 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(DropWhileInt64(nil, nil)) > 0 {
		t.Errorf("DropWhileInt64 failed.")
	}

	if len(DropWhileInt64(nil, []int64{})) > 0 {
		t.Errorf("DropWhileInt64 failed.")
	}

	NewList = DropWhileInt64(isEvenInt64, []int64{4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileInt64 failed")
	}
}

func TestDropWhileInt32(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails
	expectedNewList := []int32{3, 4, 5}
	NewList := DropWhileInt32(isEvenInt32, []int32{4, 2, 3, 4, 5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhileInt32 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []int32{4}
	partialIsEvenDivisibleBy := func(num int32) bool { return isEvenInt32DivisibleBy(num, 10) }
	NewList = DropWhileInt32(partialIsEvenDivisibleBy, []int32{20, 40, 4})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("DropWhileInt32 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(DropWhileInt32(nil, nil)) > 0 {
		t.Errorf("DropWhileInt32 failed.")
	}

	if len(DropWhileInt32(nil, []int32{})) > 0 {
		t.Errorf("DropWhileInt32 failed.")
	}

	NewList = DropWhileInt32(isEvenInt32, []int32{4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileInt32 failed")
	}
}

func TestDropWhileInt16(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails
	expectedNewList := []int16{3, 4, 5}
	NewList := DropWhileInt16(isEvenInt16, []int16{4, 2, 3, 4, 5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhileInt32 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []int16{4}
	partialIsEvenDivisibleBy := func(num int16) bool { return isEvenInt16DivisibleBy(num, 10) }
	NewList = DropWhileInt16(partialIsEvenDivisibleBy, []int16{20, 40, 4})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("DropWhileInt16 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(DropWhileInt16(nil, nil)) > 0 {
		t.Errorf("DropWhileInt16 failed.")
	}

	if len(DropWhileInt16(nil, []int16{})) > 0 {
		t.Errorf("DropWhileInt16 failed.")
	}

	NewList = DropWhileInt16(isEvenInt16, []int16{4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileInt16 failed")
	}

}

func TestDropWhileInt8(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails
	expectedNewList := []int8{3, 4, 5}
	NewList := DropWhileInt8(isEvenInt8, []int8{4, 2, 3, 4, 5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhileInt8 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []int8{4}
	partialIsEvenDivisibleBy := func(num int8) bool { return isEvenInt8DivisibleBy(num, 10) }
	NewList = DropWhileInt8(partialIsEvenDivisibleBy, []int8{20, 40, 4})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("DropWhileInt8 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(DropWhileInt8(nil, nil)) > 0 {
		t.Errorf("DropWhileInt8 failed.")
	}

	if len(DropWhileInt8(nil, []int8{})) > 0 {
		t.Errorf("DropWhileInt8 failed.")
	}

	NewList = DropWhileInt8(isEvenInt8, []int8{4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileInt8 failed")
	}
}

func TestDropWhileUInt(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails
	expectedNewList := []uint{3, 4, 5}
	NewList := DropWhileUint(isEvenUint, []uint{4, 2, 3, 4, 5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhileUint failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []uint{4}
	partialIsEvenDivisibleBy := func(num uint) bool { return isEvenUintDivisibleBy(num, 10) }
	NewList = DropWhileUint(partialIsEvenDivisibleBy, []uint{20, 40, 4})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("DropWhileUint failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(DropWhileUint(nil, nil)) > 0 {
		t.Errorf("DropWhileUint failed.")
	}

	if len(DropWhileUint(nil, []uint{})) > 0 {
		t.Errorf("DropWhileUint failed.")
	}

	NewList = DropWhileUint(isEvenUint, []uint{4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileUint failed")
	}
}

func TestDropWhileUInt64(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails
	expectedNewList := []uint64{3, 4, 5}
	NewList := DropWhileUint64(isEvenUint64, []uint64{4, 2, 3, 4, 5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhileUint64 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []uint64{4}
	partialIsEvenDivisibleBy := func(num uint64) bool { return isEvenUint64DivisibleBy(num, 10) }
	NewList = DropWhileUint64(partialIsEvenDivisibleBy, []uint64{20, 40, 4})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("DropWhileUint64 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(DropWhileUint64(nil, nil)) > 0 {
		t.Errorf("DropWhileUint64 failed.")
	}

	if len(DropWhileUint64(nil, []uint64{})) > 0 {
		t.Errorf("DropWhileUint64 failed.")
	}

	NewList = DropWhileUint64(isEvenUint64, []uint64{4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileUint64 failed")
	}
}

func TestDropWhileUInt32(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails
	expectedNewList := []uint32{3, 4, 5}
	NewList := DropWhileUint32(isEvenUint32, []uint32{4, 2, 3, 4, 5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhileUint32 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []uint32{4}
	partialIsEvenDivisibleBy := func(num uint32) bool { return isEvenUint32DivisibleBy(num, 10) }
	NewList = DropWhileUint32(partialIsEvenDivisibleBy, []uint32{20, 40, 4})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("DropWhileUint32 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(DropWhileUint32(nil, nil)) > 0 {
		t.Errorf("DropWhileUint32 failed.")
	}

	if len(DropWhileUint32(nil, []uint32{})) > 0 {
		t.Errorf("DropWhileUint32 failed.")
	}

	NewList = DropWhileUint32(isEvenUint32, []uint32{4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileUint32 failed")
	}
}

func TestDropWhileUInt16(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails
	expectedNewList := []uint16{3, 4, 5}
	NewList := DropWhileUint16(isEvenUint16, []uint16{4, 2, 3, 4, 5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhileUint16 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []uint16{4}
	partialIsEvenDivisibleBy := func(num uint16) bool { return isEvenUint16DivisibleBy(num, 10) }
	NewList = DropWhileUint16(partialIsEvenDivisibleBy, []uint16{20, 40, 4})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("DropWhileUint16 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(DropWhileUint16(nil, nil)) > 0 {
		t.Errorf("DropWhileUint16 failed.")
	}

	if len(DropWhileUint16(nil, []uint16{})) > 0 {
		t.Errorf("DropWhileUint16 failed.")
	}

	NewList = DropWhileUint16(isEvenUint16, []uint16{4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileUint16 failed")
	}
}

func TestDropWhileUInt8(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails
	expectedNewList := []uint8{3, 4, 5}
	NewList := DropWhileUint8(isEvenUint8, []uint8{4, 2, 3, 4, 5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhileUint8 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	expectedNewList = []uint8{4}
	partialIsEvenDivisibleBy := func(num uint8) bool { return isEvenUint8DivisibleBy(num, 10) }
	NewList = DropWhileUint8(partialIsEvenDivisibleBy, []uint8{20, 40, 4})

	if NewList[0] != expectedNewList[0] {
		t.Errorf("DropWhileUint8 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(DropWhileUint8(nil, nil)) > 0 {
		t.Errorf("DropWhileUint8 failed.")
	}

	if len(DropWhileUint8(nil, []uint8{})) > 0 {
		t.Errorf("DropWhileUint8 failed.")
	}

	NewList = DropWhileUint8(isEvenUint8, []uint8{4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileUint8 failed")
	}
}

func TestDropWhileFloat64(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails
	expectedNewList := []float64{-1, 4, 5}
	NewList := DropWhileFloat64(isPositiveFloat64, []float64{4, 2, 3, -1, 4, 5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhileFloat64 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(DropWhileFloat64(nil, nil)) > 0 {
		t.Errorf("DropWhileFloat64 failed.")
	}

	if len(DropWhileFloat64(nil, []float64{})) > 0 {
		t.Errorf("DropWhileFloat64 failed.")
	}

	NewList = DropWhileFloat64(isPositiveFloat64, []float64{4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileFloat64 failed")
	}
}

func TestDropWhileFloat32(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails
	expectedNewList := []float32{-1, 4, 5}
	NewList := DropWhileFloat32(isPositiveFloat32, []float32{4, 2, 3, -1, 4, 5})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhileFloat32 failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(DropWhileFloat32(nil, nil)) > 0 {
		t.Errorf("DropWhileFloat32 failed.")
	}

	if len(DropWhileFloat32(nil, []float32{})) > 0 {
		t.Errorf("DropWhileFloat64 failed.")
	}

	NewList = DropWhileFloat32(isPositiveFloat32, []float32{4})
	if len(NewList) != 0 {
		t.Errorf("DropWhileFloat32 failed")
	}
}

func TestDropWhileStr(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails
	expectedNewList := []string{"Nandeshwar", "ShyamSundar", "Hari Shankar"}
	NewList := DropWhileStr(
		func(str string) bool {
			return len(str) < 10
		},
		[]string{"Ram", "Shyam", "Nandeshwar", "ShyamSundar", "Hari Shankar"})

	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] || NewList[2] != expectedNewList[2] {
		t.Errorf("DropWhileStr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(DropWhileStr(nil, nil)) > 0 {
		t.Errorf("DropWhileStr failed.")
	}

	if len(DropWhileStr(nil, []string{})) > 0 {
		t.Errorf("DropWhileStr failed.")
	}

	if len(DropWhileStr(
		func(str string) bool {
			return len(str) < 10
		}, []string{"Ram", "Shyam"})) > 0 {
		t.Errorf("DropWhileStr failed.")
	}
}

func TestDropWhileBool(t *testing.T) {
	var vt bool = true
	var vf bool = false

	isTrueBool := func(num bool) bool {
		return num == true
	}

	expectedNewList := []bool{vf, vt}
	NewList := DropWhileBool(isTrueBool, []bool{vt, vf, vt})
	if NewList[0] != expectedNewList[0] || NewList[1] != expectedNewList[1] {
		t.Errorf("DropWhileBool failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}
}
