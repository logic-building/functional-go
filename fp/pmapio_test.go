package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestPmapIntInt64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3, 4}
	newList := PMapIntInt64(plusOneIntInt64, []int{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapIntInt64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapIntInt64(nil, nil)) > 0 {
		t.Errorf("PMapIntInt64 failed")
	}

	if len(PMapIntInt64(nil, []int{})) > 0 {
		t.Errorf("PMapIntInt64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int64{2, 3, 4}
	newList = PMapIntInt64(plusOneIntInt64, []int{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapIntInt64 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneIntInt64(num int) int64 {
	return int64(num + 1)
}

func TestPmapIntInt32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3, 4}
	newList := PMapIntInt32(plusOneIntInt32, []int{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapIntInt32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapIntInt32(nil, nil)) > 0 {
		t.Errorf("PMapIntInt32 failed")
	}

	if len(PMapIntInt32(nil, []int{})) > 0 {
		t.Errorf("PMapIntInt32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int32{2, 3, 4}
	newList = PMapIntInt32(plusOneIntInt32, []int{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapIntInt32 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneIntInt32(num int) int32 {
	return int32(num + 1)
}

func TestPmapIntInt16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3, 4}
	newList := PMapIntInt16(plusOneIntInt16, []int{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapIntInt16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapIntInt16(nil, nil)) > 0 {
		t.Errorf("PMapIntInt16 failed")
	}

	if len(PMapIntInt16(nil, []int{})) > 0 {
		t.Errorf("PMapIntInt16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int16{2, 3, 4}
	newList = PMapIntInt16(plusOneIntInt16, []int{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapIntInt16 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneIntInt16(num int) int16 {
	return int16(num + 1)
}

func TestPmapIntInt8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3, 4}
	newList := PMapIntInt8(plusOneIntInt8, []int{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapIntInt8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapIntInt8(nil, nil)) > 0 {
		t.Errorf("PMapIntInt8 failed")
	}

	if len(PMapIntInt8(nil, []int{})) > 0 {
		t.Errorf("PMapIntInt8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int8{2, 3, 4}
	newList = PMapIntInt8(plusOneIntInt8, []int{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapIntInt8 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneIntInt8(num int) int8 {
	return int8(num + 1)
}

func TestPmapIntUint(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3, 4}
	newList := PMapIntUint(plusOneIntUint, []int{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapIntUint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapIntUint(nil, nil)) > 0 {
		t.Errorf("PMapIntUint failed")
	}

	if len(PMapIntUint(nil, []int{})) > 0 {
		t.Errorf("PMapIntUint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint{2, 3, 4}
	newList = PMapIntUint(plusOneIntUint, []int{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapIntUint failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneIntUint(num int) uint {
	return uint(num + 1)
}

func TestPmapIntUint64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3, 4}
	newList := PMapIntUint64(plusOneIntUint64, []int{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapIntUint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapIntUint64(nil, nil)) > 0 {
		t.Errorf("PMapIntUint64 failed")
	}

	if len(PMapIntUint64(nil, []int{})) > 0 {
		t.Errorf("PMapIntUint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint64{2, 3, 4}
	newList = PMapIntUint64(plusOneIntUint64, []int{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapIntUint64 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneIntUint64(num int) uint64 {
	return uint64(num + 1)
}

func TestPmapIntUint32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3, 4}
	newList := PMapIntUint32(plusOneIntUint32, []int{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapIntUint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapIntUint32(nil, nil)) > 0 {
		t.Errorf("PMapIntUint32 failed")
	}

	if len(PMapIntUint32(nil, []int{})) > 0 {
		t.Errorf("PMapIntUint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint32{2, 3, 4}
	newList = PMapIntUint32(plusOneIntUint32, []int{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapIntUint32 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneIntUint32(num int) uint32 {
	return uint32(num + 1)
}

func TestPmapIntUint16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3, 4}
	newList := PMapIntUint16(plusOneIntUint16, []int{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapIntUint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapIntUint16(nil, nil)) > 0 {
		t.Errorf("PMapIntUint16 failed")
	}

	if len(PMapIntUint16(nil, []int{})) > 0 {
		t.Errorf("PMapIntUint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint16{2, 3, 4}
	newList = PMapIntUint16(plusOneIntUint16, []int{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapIntUint16 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneIntUint16(num int) uint16 {
	return uint16(num + 1)
}

func TestPmapIntUint8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3, 4}
	newList := PMapIntUint8(plusOneIntUint8, []int{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapIntUint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapIntUint8(nil, nil)) > 0 {
		t.Errorf("PMapIntUint8 failed")
	}

	if len(PMapIntUint8(nil, []int{})) > 0 {
		t.Errorf("PMapIntUint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint8{2, 3, 4}
	newList = PMapIntUint8(plusOneIntUint8, []int{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapIntUint8 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneIntUint8(num int) uint8 {
	return uint8(num + 1)
}

func TestPmapIntStr(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList := PMapIntStr(someLogicIntStr, []int{10})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapIntStr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapIntStr(nil, nil)) > 0 {
		t.Errorf("PMapIntStr failed")
	}

	if len(PMapIntStr(nil, []int{})) > 0 {
		t.Errorf("PMapIntStr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []string{"10"}
	newList = PMapIntStr(someLogicIntStr, []int{10, 20}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != "0" {
		t.Errorf("PMapIntStr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicIntStr(num int) string {
	if num == 10 {
		return string("10")
	}
	return "0"
}

func TestPmapIntBool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := PMapIntBool(someLogicIntBool, []int{10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapIntBool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapIntBool(nil, nil)) > 0 {
		t.Errorf("PMapIntBool failed")
	}

	if len(PMapIntBool(nil, []int{})) > 0 {
		t.Errorf("PMapIntBool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []bool{true, false}
	newList = PMapIntBool(someLogicIntBool, []int{10, 0}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != false {
		t.Errorf("PMapIntBool failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicIntBool(num int) bool {
	if num > 0 {
		return true
	}
	return false
}

func TestPmapIntFloat32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3, 4}
	newList := PMapIntFloat32(plusOneIntFloat32, []int{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapIntFloat32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapIntFloat32(nil, nil)) > 0 {
		t.Errorf("PMapIntFloat32 failed")
	}

	if len(PMapIntFloat32(nil, []int{})) > 0 {
		t.Errorf("PMapIntFloat32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []float32{2, 3, 4}
	newList = PMapIntFloat32(plusOneIntFloat32, []int{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapIntFloat32 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneIntFloat32(num int) float32 {
	return float32(num + 1)
}

func TestPmapIntFloat64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3, 4}
	newList := PMapIntFloat64(plusOneIntFloat64, []int{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapIntFloat64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapIntFloat64(nil, nil)) > 0 {
		t.Errorf("PMapIntFloat64 failed")
	}

	if len(PMapIntFloat64(nil, []int{})) > 0 {
		t.Errorf("PMapIntFloat64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []float64{2, 3, 4}
	newList = PMapIntFloat64(plusOneIntFloat64, []int{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapIntFloat64 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneIntFloat64(num int) float64 {
	return float64(num + 1)
}

func TestPmapInt64Int(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3, 4}
	newList := PMapInt64Int(plusOneInt64Int, []int64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt64Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt64Int(nil, nil)) > 0 {
		t.Errorf("PMapInt64Int failed")
	}

	if len(PMapInt64Int(nil, []int64{})) > 0 {
		t.Errorf("PMapInt64Int failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int{2, 3, 4}
	newList = PMapInt64Int(plusOneInt64Int, []int64{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt64Int failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt64Int(num int64) int {
	return int(num + 1)
}

func TestPmapInt64Int32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3, 4}
	newList := PMapInt64Int32(plusOneInt64Int32, []int64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt64Int32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt64Int32(nil, nil)) > 0 {
		t.Errorf("PMapInt64Int32 failed")
	}

	if len(PMapInt64Int32(nil, []int64{})) > 0 {
		t.Errorf("PMapInt64Int32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int32{2, 3, 4}
	newList = PMapInt64Int32(plusOneInt64Int32, []int64{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt64Int32 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt64Int32(num int64) int32 {
	return int32(num + 1)
}

func TestPmapInt64Int16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3, 4}
	newList := PMapInt64Int16(plusOneInt64Int16, []int64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt64Int16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt64Int16(nil, nil)) > 0 {
		t.Errorf("PMapInt64Int16 failed")
	}

	if len(PMapInt64Int16(nil, []int64{})) > 0 {
		t.Errorf("PMapInt64Int16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int16{2, 3, 4}
	newList = PMapInt64Int16(plusOneInt64Int16, []int64{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt64Int16 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt64Int16(num int64) int16 {
	return int16(num + 1)
}

func TestPmapInt64Int8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3, 4}
	newList := PMapInt64Int8(plusOneInt64Int8, []int64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt64Int8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt64Int8(nil, nil)) > 0 {
		t.Errorf("PMapInt64Int8 failed")
	}

	if len(PMapInt64Int8(nil, []int64{})) > 0 {
		t.Errorf("PMapInt64Int8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int8{2, 3, 4}
	newList = PMapInt64Int8(plusOneInt64Int8, []int64{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt64Int8 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt64Int8(num int64) int8 {
	return int8(num + 1)
}

func TestPmapInt64Uint(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3, 4}
	newList := PMapInt64Uint(plusOneInt64Uint, []int64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt64Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt64Uint(nil, nil)) > 0 {
		t.Errorf("PMapInt64Uint failed")
	}

	if len(PMapInt64Uint(nil, []int64{})) > 0 {
		t.Errorf("PMapInt64Uint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint{2, 3, 4}
	newList = PMapInt64Uint(plusOneInt64Uint, []int64{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt64Uint failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt64Uint(num int64) uint {
	return uint(num + 1)
}

func TestPmapInt64Uint64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3, 4}
	newList := PMapInt64Uint64(plusOneInt64Uint64, []int64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt64Uint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt64Uint64(nil, nil)) > 0 {
		t.Errorf("PMapInt64Uint64 failed")
	}

	if len(PMapInt64Uint64(nil, []int64{})) > 0 {
		t.Errorf("PMapInt64Uint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint64{2, 3, 4}
	newList = PMapInt64Uint64(plusOneInt64Uint64, []int64{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt64Uint64 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt64Uint64(num int64) uint64 {
	return uint64(num + 1)
}

func TestPmapInt64Uint32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3, 4}
	newList := PMapInt64Uint32(plusOneInt64Uint32, []int64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt64Uint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt64Uint32(nil, nil)) > 0 {
		t.Errorf("PMapInt64Uint32 failed")
	}

	if len(PMapInt64Uint32(nil, []int64{})) > 0 {
		t.Errorf("PMapInt64Uint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint32{2, 3, 4}
	newList = PMapInt64Uint32(plusOneInt64Uint32, []int64{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt64Uint32 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt64Uint32(num int64) uint32 {
	return uint32(num + 1)
}

func TestPmapInt64Uint16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3, 4}
	newList := PMapInt64Uint16(plusOneInt64Uint16, []int64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt64Uint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt64Uint16(nil, nil)) > 0 {
		t.Errorf("PMapInt64Uint16 failed")
	}

	if len(PMapInt64Uint16(nil, []int64{})) > 0 {
		t.Errorf("PMapInt64Uint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint16{2, 3, 4}
	newList = PMapInt64Uint16(plusOneInt64Uint16, []int64{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt64Uint16 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt64Uint16(num int64) uint16 {
	return uint16(num + 1)
}

func TestPmapInt64Uint8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3, 4}
	newList := PMapInt64Uint8(plusOneInt64Uint8, []int64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt64Uint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt64Uint8(nil, nil)) > 0 {
		t.Errorf("PMapInt64Uint8 failed")
	}

	if len(PMapInt64Uint8(nil, []int64{})) > 0 {
		t.Errorf("PMapInt64Uint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint8{2, 3, 4}
	newList = PMapInt64Uint8(plusOneInt64Uint8, []int64{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt64Uint8 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt64Uint8(num int64) uint8 {
	return uint8(num + 1)
}

func TestPmapInt64Str(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList := PMapInt64Str(someLogicInt64Str, []int64{10})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapInt64Str failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt64Str(nil, nil)) > 0 {
		t.Errorf("PMapInt64Str failed")
	}

	if len(PMapInt64Str(nil, []int64{})) > 0 {
		t.Errorf("PMapInt64Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []string{"10"}
	newList = PMapInt64Str(someLogicInt64Str, []int64{10, 20}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != "0" {
		t.Errorf("PMapInt64Str failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicInt64Str(num int64) string {
	if num == 10 {
		return string("10")
	}
	return "0"
}

func TestPmapInt64Bool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := PMapInt64Bool(someLogicInt64Bool, []int64{10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt64Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt64Bool(nil, nil)) > 0 {
		t.Errorf("PMapInt64Bool failed")
	}

	if len(PMapInt64Bool(nil, []int64{})) > 0 {
		t.Errorf("PMapInt64Bool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []bool{true, false}
	newList = PMapInt64Bool(someLogicInt64Bool, []int64{10, 0}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != false {
		t.Errorf("PMapInt64Bool failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicInt64Bool(num int64) bool {
	if num > 0 {
		return true
	}
	return false
}

func TestPmapInt64Float32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3, 4}
	newList := PMapInt64Float32(plusOneInt64Float32, []int64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt64Float32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt64Float32(nil, nil)) > 0 {
		t.Errorf("PMapInt64Float32 failed")
	}

	if len(PMapInt64Float32(nil, []int64{})) > 0 {
		t.Errorf("PMapInt64Float32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []float32{2, 3, 4}
	newList = PMapInt64Float32(plusOneInt64Float32, []int64{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt64Float32 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt64Float32(num int64) float32 {
	return float32(num + 1)
}

func TestPmapInt64Float64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3, 4}
	newList := PMapInt64Float64(plusOneInt64Float64, []int64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt64Float64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt64Float64(nil, nil)) > 0 {
		t.Errorf("PMapInt64Float64 failed")
	}

	if len(PMapInt64Float64(nil, []int64{})) > 0 {
		t.Errorf("PMapInt64Float64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []float64{2, 3, 4}
	newList = PMapInt64Float64(plusOneInt64Float64, []int64{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt64Float64 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt64Float64(num int64) float64 {
	return float64(num + 1)
}

func TestPmapInt32Int(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3, 4}
	newList := PMapInt32Int(plusOneInt32Int, []int32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt32Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt32Int(nil, nil)) > 0 {
		t.Errorf("PMapInt32Int failed")
	}

	if len(PMapInt32Int(nil, []int32{})) > 0 {
		t.Errorf("PMapInt32Int failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int{2, 3, 4}
	newList = PMapInt32Int(plusOneInt32Int, []int32{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt32Int failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt32Int(num int32) int {
	return int(num + 1)
}

func TestPmapInt32Int64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3, 4}
	newList := PMapInt32Int64(plusOneInt32Int64, []int32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt32Int64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt32Int64(nil, nil)) > 0 {
		t.Errorf("PMapInt32Int64 failed")
	}

	if len(PMapInt32Int64(nil, []int32{})) > 0 {
		t.Errorf("PMapInt32Int64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int64{2, 3, 4}
	newList = PMapInt32Int64(plusOneInt32Int64, []int32{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt32Int64 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt32Int64(num int32) int64 {
	return int64(num + 1)
}

func TestPmapInt32Int16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3, 4}
	newList := PMapInt32Int16(plusOneInt32Int16, []int32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt32Int16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt32Int16(nil, nil)) > 0 {
		t.Errorf("PMapInt32Int16 failed")
	}

	if len(PMapInt32Int16(nil, []int32{})) > 0 {
		t.Errorf("PMapInt32Int16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int16{2, 3, 4}
	newList = PMapInt32Int16(plusOneInt32Int16, []int32{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt32Int16 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt32Int16(num int32) int16 {
	return int16(num + 1)
}

func TestPmapInt32Int8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3, 4}
	newList := PMapInt32Int8(plusOneInt32Int8, []int32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt32Int8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt32Int8(nil, nil)) > 0 {
		t.Errorf("PMapInt32Int8 failed")
	}

	if len(PMapInt32Int8(nil, []int32{})) > 0 {
		t.Errorf("PMapInt32Int8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int8{2, 3, 4}
	newList = PMapInt32Int8(plusOneInt32Int8, []int32{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt32Int8 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt32Int8(num int32) int8 {
	return int8(num + 1)
}

func TestPmapInt32Uint(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3, 4}
	newList := PMapInt32Uint(plusOneInt32Uint, []int32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt32Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt32Uint(nil, nil)) > 0 {
		t.Errorf("PMapInt32Uint failed")
	}

	if len(PMapInt32Uint(nil, []int32{})) > 0 {
		t.Errorf("PMapInt32Uint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint{2, 3, 4}
	newList = PMapInt32Uint(plusOneInt32Uint, []int32{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt32Uint failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt32Uint(num int32) uint {
	return uint(num + 1)
}

func TestPmapInt32Uint64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3, 4}
	newList := PMapInt32Uint64(plusOneInt32Uint64, []int32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt32Uint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt32Uint64(nil, nil)) > 0 {
		t.Errorf("PMapInt32Uint64 failed")
	}

	if len(PMapInt32Uint64(nil, []int32{})) > 0 {
		t.Errorf("PMapInt32Uint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint64{2, 3, 4}
	newList = PMapInt32Uint64(plusOneInt32Uint64, []int32{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt32Uint64 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt32Uint64(num int32) uint64 {
	return uint64(num + 1)
}

func TestPmapInt32Uint32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3, 4}
	newList := PMapInt32Uint32(plusOneInt32Uint32, []int32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt32Uint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt32Uint32(nil, nil)) > 0 {
		t.Errorf("PMapInt32Uint32 failed")
	}

	if len(PMapInt32Uint32(nil, []int32{})) > 0 {
		t.Errorf("PMapInt32Uint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint32{2, 3, 4}
	newList = PMapInt32Uint32(plusOneInt32Uint32, []int32{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt32Uint32 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt32Uint32(num int32) uint32 {
	return uint32(num + 1)
}

func TestPmapInt32Uint16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3, 4}
	newList := PMapInt32Uint16(plusOneInt32Uint16, []int32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt32Uint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt32Uint16(nil, nil)) > 0 {
		t.Errorf("PMapInt32Uint16 failed")
	}

	if len(PMapInt32Uint16(nil, []int32{})) > 0 {
		t.Errorf("PMapInt32Uint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint16{2, 3, 4}
	newList = PMapInt32Uint16(plusOneInt32Uint16, []int32{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt32Uint16 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt32Uint16(num int32) uint16 {
	return uint16(num + 1)
}

func TestPmapInt32Uint8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3, 4}
	newList := PMapInt32Uint8(plusOneInt32Uint8, []int32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt32Uint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt32Uint8(nil, nil)) > 0 {
		t.Errorf("PMapInt32Uint8 failed")
	}

	if len(PMapInt32Uint8(nil, []int32{})) > 0 {
		t.Errorf("PMapInt32Uint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint8{2, 3, 4}
	newList = PMapInt32Uint8(plusOneInt32Uint8, []int32{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt32Uint8 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt32Uint8(num int32) uint8 {
	return uint8(num + 1)
}

func TestPmapInt32Str(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList := PMapInt32Str(someLogicInt32Str, []int32{10})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapInt32Str failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt32Str(nil, nil)) > 0 {
		t.Errorf("PMapInt32Str failed")
	}

	if len(PMapInt32Str(nil, []int32{})) > 0 {
		t.Errorf("PMapInt32Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []string{"10"}
	newList = PMapInt32Str(someLogicInt32Str, []int32{10, 20}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != "0" {
		t.Errorf("PMapInt32Str failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicInt32Str(num int32) string {
	if num == 10 {
		return string("10")
	}
	return "0"
}

func TestPmapInt32Bool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := PMapInt32Bool(someLogicInt32Bool, []int32{10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt32Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt32Bool(nil, nil)) > 0 {
		t.Errorf("PMapInt32Bool failed")
	}

	if len(PMapInt32Bool(nil, []int32{})) > 0 {
		t.Errorf("PMapInt32Bool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []bool{true, false}
	newList = PMapInt32Bool(someLogicInt32Bool, []int32{10, 0}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != false {
		t.Errorf("PMapInt32Bool failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicInt32Bool(num int32) bool {
	if num > 0 {
		return true
	}
	return false
}

func TestPmapInt32Float32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3, 4}
	newList := PMapInt32Float32(plusOneInt32Float32, []int32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt32Float32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt32Float32(nil, nil)) > 0 {
		t.Errorf("PMapInt32Float32 failed")
	}

	if len(PMapInt32Float32(nil, []int32{})) > 0 {
		t.Errorf("PMapInt32Float32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []float32{2, 3, 4}
	newList = PMapInt32Float32(plusOneInt32Float32, []int32{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt32Float32 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt32Float32(num int32) float32 {
	return float32(num + 1)
}

func TestPmapInt32Float64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3, 4}
	newList := PMapInt32Float64(plusOneInt32Float64, []int32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt32Float64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt32Float64(nil, nil)) > 0 {
		t.Errorf("PMapInt32Float64 failed")
	}

	if len(PMapInt32Float64(nil, []int32{})) > 0 {
		t.Errorf("PMapInt32Float64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []float64{2, 3, 4}
	newList = PMapInt32Float64(plusOneInt32Float64, []int32{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt32Float64 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt32Float64(num int32) float64 {
	return float64(num + 1)
}

func TestPmapInt16Int(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3, 4}
	newList := PMapInt16Int(plusOneInt16Int, []int16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt16Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt16Int(nil, nil)) > 0 {
		t.Errorf("PMapInt16Int failed")
	}

	if len(PMapInt16Int(nil, []int16{})) > 0 {
		t.Errorf("PMapInt16Int failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int{2, 3, 4}
	newList = PMapInt16Int(plusOneInt16Int, []int16{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt16Int failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt16Int(num int16) int {
	return int(num + 1)
}

func TestPmapInt16Int64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3, 4}
	newList := PMapInt16Int64(plusOneInt16Int64, []int16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt16Int64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt16Int64(nil, nil)) > 0 {
		t.Errorf("PMapInt16Int64 failed")
	}

	if len(PMapInt16Int64(nil, []int16{})) > 0 {
		t.Errorf("PMapInt16Int64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int64{2, 3, 4}
	newList = PMapInt16Int64(plusOneInt16Int64, []int16{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt16Int64 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt16Int64(num int16) int64 {
	return int64(num + 1)
}

func TestPmapInt16Int32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3, 4}
	newList := PMapInt16Int32(plusOneInt16Int32, []int16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt16Int32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt16Int32(nil, nil)) > 0 {
		t.Errorf("PMapInt16Int32 failed")
	}

	if len(PMapInt16Int32(nil, []int16{})) > 0 {
		t.Errorf("PMapInt16Int32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int32{2, 3, 4}
	newList = PMapInt16Int32(plusOneInt16Int32, []int16{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt16Int32 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt16Int32(num int16) int32 {
	return int32(num + 1)
}

func TestPmapInt16Int8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3, 4}
	newList := PMapInt16Int8(plusOneInt16Int8, []int16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt16Int8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt16Int8(nil, nil)) > 0 {
		t.Errorf("PMapInt16Int8 failed")
	}

	if len(PMapInt16Int8(nil, []int16{})) > 0 {
		t.Errorf("PMapInt16Int8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int8{2, 3, 4}
	newList = PMapInt16Int8(plusOneInt16Int8, []int16{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt16Int8 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt16Int8(num int16) int8 {
	return int8(num + 1)
}

func TestPmapInt16Uint(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3, 4}
	newList := PMapInt16Uint(plusOneInt16Uint, []int16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt16Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt16Uint(nil, nil)) > 0 {
		t.Errorf("PMapInt16Uint failed")
	}

	if len(PMapInt16Uint(nil, []int16{})) > 0 {
		t.Errorf("PMapInt16Uint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint{2, 3, 4}
	newList = PMapInt16Uint(plusOneInt16Uint, []int16{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt16Uint failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt16Uint(num int16) uint {
	return uint(num + 1)
}

func TestPmapInt16Uint64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3, 4}
	newList := PMapInt16Uint64(plusOneInt16Uint64, []int16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt16Uint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt16Uint64(nil, nil)) > 0 {
		t.Errorf("PMapInt16Uint64 failed")
	}

	if len(PMapInt16Uint64(nil, []int16{})) > 0 {
		t.Errorf("PMapInt16Uint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint64{2, 3, 4}
	newList = PMapInt16Uint64(plusOneInt16Uint64, []int16{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt16Uint64 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt16Uint64(num int16) uint64 {
	return uint64(num + 1)
}

func TestPmapInt16Uint32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3, 4}
	newList := PMapInt16Uint32(plusOneInt16Uint32, []int16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt16Uint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt16Uint32(nil, nil)) > 0 {
		t.Errorf("PMapInt16Uint32 failed")
	}

	if len(PMapInt16Uint32(nil, []int16{})) > 0 {
		t.Errorf("PMapInt16Uint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint32{2, 3, 4}
	newList = PMapInt16Uint32(plusOneInt16Uint32, []int16{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt16Uint32 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt16Uint32(num int16) uint32 {
	return uint32(num + 1)
}

func TestPmapInt16Uint16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3, 4}
	newList := PMapInt16Uint16(plusOneInt16Uint16, []int16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt16Uint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt16Uint16(nil, nil)) > 0 {
		t.Errorf("PMapInt16Uint16 failed")
	}

	if len(PMapInt16Uint16(nil, []int16{})) > 0 {
		t.Errorf("PMapInt16Uint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint16{2, 3, 4}
	newList = PMapInt16Uint16(plusOneInt16Uint16, []int16{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt16Uint16 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt16Uint16(num int16) uint16 {
	return uint16(num + 1)
}

func TestPmapInt16Uint8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3, 4}
	newList := PMapInt16Uint8(plusOneInt16Uint8, []int16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt16Uint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt16Uint8(nil, nil)) > 0 {
		t.Errorf("PMapInt16Uint8 failed")
	}

	if len(PMapInt16Uint8(nil, []int16{})) > 0 {
		t.Errorf("PMapInt16Uint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint8{2, 3, 4}
	newList = PMapInt16Uint8(plusOneInt16Uint8, []int16{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt16Uint8 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt16Uint8(num int16) uint8 {
	return uint8(num + 1)
}

func TestPmapInt16Str(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList := PMapInt16Str(someLogicInt16Str, []int16{10})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapInt16Str failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt16Str(nil, nil)) > 0 {
		t.Errorf("PMapInt16Str failed")
	}

	if len(PMapInt16Str(nil, []int16{})) > 0 {
		t.Errorf("PMapInt16Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []string{"10"}
	newList = PMapInt16Str(someLogicInt16Str, []int16{10, 20}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != "0" {
		t.Errorf("PMapInt16Str failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicInt16Str(num int16) string {
	if num == 10 {
		return string("10")
	}
	return "0"
}

func TestPmapInt16Bool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := PMapInt16Bool(someLogicInt16Bool, []int16{10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt16Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt16Bool(nil, nil)) > 0 {
		t.Errorf("PMapInt16Bool failed")
	}

	if len(PMapInt16Bool(nil, []int16{})) > 0 {
		t.Errorf("PMapInt16Bool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []bool{true, false}
	newList = PMapInt16Bool(someLogicInt16Bool, []int16{10, 0}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != false {
		t.Errorf("PMapInt16Bool failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicInt16Bool(num int16) bool {
	if num > 0 {
		return true
	}
	return false
}

func TestPmapInt16Float32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3, 4}
	newList := PMapInt16Float32(plusOneInt16Float32, []int16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt16Float32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt16Float32(nil, nil)) > 0 {
		t.Errorf("PMapInt16Float32 failed")
	}

	if len(PMapInt16Float32(nil, []int16{})) > 0 {
		t.Errorf("PMapInt16Float32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []float32{2, 3, 4}
	newList = PMapInt16Float32(plusOneInt16Float32, []int16{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt16Float32 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt16Float32(num int16) float32 {
	return float32(num + 1)
}

func TestPmapInt16Float64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3, 4}
	newList := PMapInt16Float64(plusOneInt16Float64, []int16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt16Float64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt16Float64(nil, nil)) > 0 {
		t.Errorf("PMapInt16Float64 failed")
	}

	if len(PMapInt16Float64(nil, []int16{})) > 0 {
		t.Errorf("PMapInt16Float64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []float64{2, 3, 4}
	newList = PMapInt16Float64(plusOneInt16Float64, []int16{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt16Float64 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt16Float64(num int16) float64 {
	return float64(num + 1)
}

func TestPmapInt8Int(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3, 4}
	newList := PMapInt8Int(plusOneInt8Int, []int8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt8Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt8Int(nil, nil)) > 0 {
		t.Errorf("PMapInt8Int failed")
	}

	if len(PMapInt8Int(nil, []int8{})) > 0 {
		t.Errorf("PMapInt8Int failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int{2, 3, 4}
	newList = PMapInt8Int(plusOneInt8Int, []int8{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt8Int failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt8Int(num int8) int {
	return int(num + 1)
}

func TestPmapInt8Int64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3, 4}
	newList := PMapInt8Int64(plusOneInt8Int64, []int8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt8Int64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt8Int64(nil, nil)) > 0 {
		t.Errorf("PMapInt8Int64 failed")
	}

	if len(PMapInt8Int64(nil, []int8{})) > 0 {
		t.Errorf("PMapInt8Int64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int64{2, 3, 4}
	newList = PMapInt8Int64(plusOneInt8Int64, []int8{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt8Int64 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt8Int64(num int8) int64 {
	return int64(num + 1)
}

func TestPmapInt8Int32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3, 4}
	newList := PMapInt8Int32(plusOneInt8Int32, []int8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt8Int32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt8Int32(nil, nil)) > 0 {
		t.Errorf("PMapInt8Int32 failed")
	}

	if len(PMapInt8Int32(nil, []int8{})) > 0 {
		t.Errorf("PMapInt8Int32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int32{2, 3, 4}
	newList = PMapInt8Int32(plusOneInt8Int32, []int8{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt8Int32 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt8Int32(num int8) int32 {
	return int32(num + 1)
}

func TestPmapInt8Int16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3, 4}
	newList := PMapInt8Int16(plusOneInt8Int16, []int8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt8Int16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt8Int16(nil, nil)) > 0 {
		t.Errorf("PMapInt8Int16 failed")
	}

	if len(PMapInt8Int16(nil, []int8{})) > 0 {
		t.Errorf("PMapInt8Int16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int16{2, 3, 4}
	newList = PMapInt8Int16(plusOneInt8Int16, []int8{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt8Int16 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt8Int16(num int8) int16 {
	return int16(num + 1)
}

func TestPmapInt8Uint(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3, 4}
	newList := PMapInt8Uint(plusOneInt8Uint, []int8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt8Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt8Uint(nil, nil)) > 0 {
		t.Errorf("PMapInt8Uint failed")
	}

	if len(PMapInt8Uint(nil, []int8{})) > 0 {
		t.Errorf("PMapInt8Uint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint{2, 3, 4}
	newList = PMapInt8Uint(plusOneInt8Uint, []int8{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt8Uint failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt8Uint(num int8) uint {
	return uint(num + 1)
}

func TestPmapInt8Uint64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3, 4}
	newList := PMapInt8Uint64(plusOneInt8Uint64, []int8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt8Uint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt8Uint64(nil, nil)) > 0 {
		t.Errorf("PMapInt8Uint64 failed")
	}

	if len(PMapInt8Uint64(nil, []int8{})) > 0 {
		t.Errorf("PMapInt8Uint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint64{2, 3, 4}
	newList = PMapInt8Uint64(plusOneInt8Uint64, []int8{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt8Uint64 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt8Uint64(num int8) uint64 {
	return uint64(num + 1)
}

func TestPmapInt8Uint32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3, 4}
	newList := PMapInt8Uint32(plusOneInt8Uint32, []int8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt8Uint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt8Uint32(nil, nil)) > 0 {
		t.Errorf("PMapInt8Uint32 failed")
	}

	if len(PMapInt8Uint32(nil, []int8{})) > 0 {
		t.Errorf("PMapInt8Uint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint32{2, 3, 4}
	newList = PMapInt8Uint32(plusOneInt8Uint32, []int8{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt8Uint32 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt8Uint32(num int8) uint32 {
	return uint32(num + 1)
}

func TestPmapInt8Uint16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3, 4}
	newList := PMapInt8Uint16(plusOneInt8Uint16, []int8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt8Uint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt8Uint16(nil, nil)) > 0 {
		t.Errorf("PMapInt8Uint16 failed")
	}

	if len(PMapInt8Uint16(nil, []int8{})) > 0 {
		t.Errorf("PMapInt8Uint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint16{2, 3, 4}
	newList = PMapInt8Uint16(plusOneInt8Uint16, []int8{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt8Uint16 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt8Uint16(num int8) uint16 {
	return uint16(num + 1)
}

func TestPmapInt8Uint8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3, 4}
	newList := PMapInt8Uint8(plusOneInt8Uint8, []int8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt8Uint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt8Uint8(nil, nil)) > 0 {
		t.Errorf("PMapInt8Uint8 failed")
	}

	if len(PMapInt8Uint8(nil, []int8{})) > 0 {
		t.Errorf("PMapInt8Uint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint8{2, 3, 4}
	newList = PMapInt8Uint8(plusOneInt8Uint8, []int8{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt8Uint8 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt8Uint8(num int8) uint8 {
	return uint8(num + 1)
}

func TestPmapInt8Str(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList := PMapInt8Str(someLogicInt8Str, []int8{10})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapInt8Str failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt8Str(nil, nil)) > 0 {
		t.Errorf("PMapInt8Str failed")
	}

	if len(PMapInt8Str(nil, []int8{})) > 0 {
		t.Errorf("PMapInt8Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []string{"10"}
	newList = PMapInt8Str(someLogicInt8Str, []int8{10, 20}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != "0" {
		t.Errorf("PMapInt8Str failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicInt8Str(num int8) string {
	if num == 10 {
		return string("10")
	}
	return "0"
}

func TestPmapInt8Bool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := PMapInt8Bool(someLogicInt8Bool, []int8{10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapInt8Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt8Bool(nil, nil)) > 0 {
		t.Errorf("PMapInt8Bool failed")
	}

	if len(PMapInt8Bool(nil, []int8{})) > 0 {
		t.Errorf("PMapInt8Bool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []bool{true, false}
	newList = PMapInt8Bool(someLogicInt8Bool, []int8{10, 0}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != false {
		t.Errorf("PMapInt8Bool failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicInt8Bool(num int8) bool {
	if num > 0 {
		return true
	}
	return false
}

func TestPmapInt8Float32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3, 4}
	newList := PMapInt8Float32(plusOneInt8Float32, []int8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt8Float32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt8Float32(nil, nil)) > 0 {
		t.Errorf("PMapInt8Float32 failed")
	}

	if len(PMapInt8Float32(nil, []int8{})) > 0 {
		t.Errorf("PMapInt8Float32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []float32{2, 3, 4}
	newList = PMapInt8Float32(plusOneInt8Float32, []int8{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt8Float32 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt8Float32(num int8) float32 {
	return float32(num + 1)
}

func TestPmapInt8Float64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3, 4}
	newList := PMapInt8Float64(plusOneInt8Float64, []int8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapInt8Float64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt8Float64(nil, nil)) > 0 {
		t.Errorf("PMapInt8Float64 failed")
	}

	if len(PMapInt8Float64(nil, []int8{})) > 0 {
		t.Errorf("PMapInt8Float64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []float64{2, 3, 4}
	newList = PMapInt8Float64(plusOneInt8Float64, []int8{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt8Float64 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneInt8Float64(num int8) float64 {
	return float64(num + 1)
}

func TestPmapUintInt(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3, 4}
	newList := PMapUintInt(plusOneUintInt, []uint{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUintInt failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUintInt(nil, nil)) > 0 {
		t.Errorf("PMapUintInt failed")
	}

	if len(PMapUintInt(nil, []uint{})) > 0 {
		t.Errorf("PMapUintInt failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int{2, 3, 4}
	newList = PMapUintInt(plusOneUintInt, []uint{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUintInt failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUintInt(num uint) int {
	return int(num + 1)
}

func TestPmapUintInt64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3, 4}
	newList := PMapUintInt64(plusOneUintInt64, []uint{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUintInt64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUintInt64(nil, nil)) > 0 {
		t.Errorf("PMapUintInt64 failed")
	}

	if len(PMapUintInt64(nil, []uint{})) > 0 {
		t.Errorf("PMapUintInt64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int64{2, 3, 4}
	newList = PMapUintInt64(plusOneUintInt64, []uint{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUintInt64 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUintInt64(num uint) int64 {
	return int64(num + 1)
}

func TestPmapUintInt32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3, 4}
	newList := PMapUintInt32(plusOneUintInt32, []uint{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUintInt32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUintInt32(nil, nil)) > 0 {
		t.Errorf("PMapUintInt32 failed")
	}

	if len(PMapUintInt32(nil, []uint{})) > 0 {
		t.Errorf("PMapUintInt32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int32{2, 3, 4}
	newList = PMapUintInt32(plusOneUintInt32, []uint{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUintInt32 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUintInt32(num uint) int32 {
	return int32(num + 1)
}

func TestPmapUintInt16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3, 4}
	newList := PMapUintInt16(plusOneUintInt16, []uint{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUintInt16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUintInt16(nil, nil)) > 0 {
		t.Errorf("PMapUintInt16 failed")
	}

	if len(PMapUintInt16(nil, []uint{})) > 0 {
		t.Errorf("PMapUintInt16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int16{2, 3, 4}
	newList = PMapUintInt16(plusOneUintInt16, []uint{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUintInt16 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUintInt16(num uint) int16 {
	return int16(num + 1)
}

func TestPmapUintInt8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3, 4}
	newList := PMapUintInt8(plusOneUintInt8, []uint{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUintInt8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUintInt8(nil, nil)) > 0 {
		t.Errorf("PMapUintInt8 failed")
	}

	if len(PMapUintInt8(nil, []uint{})) > 0 {
		t.Errorf("PMapUintInt8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int8{2, 3, 4}
	newList = PMapUintInt8(plusOneUintInt8, []uint{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUintInt8 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUintInt8(num uint) int8 {
	return int8(num + 1)
}

func TestPmapUintUint64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3, 4}
	newList := PMapUintUint64(plusOneUintUint64, []uint{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUintUint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUintUint64(nil, nil)) > 0 {
		t.Errorf("PMapUintUint64 failed")
	}

	if len(PMapUintUint64(nil, []uint{})) > 0 {
		t.Errorf("PMapUintUint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint64{2, 3, 4}
	newList = PMapUintUint64(plusOneUintUint64, []uint{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUintUint64 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUintUint64(num uint) uint64 {
	return uint64(num + 1)
}

func TestPmapUintUint32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3, 4}
	newList := PMapUintUint32(plusOneUintUint32, []uint{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUintUint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUintUint32(nil, nil)) > 0 {
		t.Errorf("PMapUintUint32 failed")
	}

	if len(PMapUintUint32(nil, []uint{})) > 0 {
		t.Errorf("PMapUintUint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint32{2, 3, 4}
	newList = PMapUintUint32(plusOneUintUint32, []uint{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUintUint32 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUintUint32(num uint) uint32 {
	return uint32(num + 1)
}

func TestPmapUintUint16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3, 4}
	newList := PMapUintUint16(plusOneUintUint16, []uint{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUintUint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUintUint16(nil, nil)) > 0 {
		t.Errorf("PMapUintUint16 failed")
	}

	if len(PMapUintUint16(nil, []uint{})) > 0 {
		t.Errorf("PMapUintUint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint16{2, 3, 4}
	newList = PMapUintUint16(plusOneUintUint16, []uint{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUintUint16 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUintUint16(num uint) uint16 {
	return uint16(num + 1)
}

func TestPmapUintUint8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3, 4}
	newList := PMapUintUint8(plusOneUintUint8, []uint{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUintUint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUintUint8(nil, nil)) > 0 {
		t.Errorf("PMapUintUint8 failed")
	}

	if len(PMapUintUint8(nil, []uint{})) > 0 {
		t.Errorf("PMapUintUint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint8{2, 3, 4}
	newList = PMapUintUint8(plusOneUintUint8, []uint{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUintUint8 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUintUint8(num uint) uint8 {
	return uint8(num + 1)
}

func TestPmapUintStr(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList := PMapUintStr(someLogicUintStr, []uint{10})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapUintStr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUintStr(nil, nil)) > 0 {
		t.Errorf("PMapUintStr failed")
	}

	if len(PMapUintStr(nil, []uint{})) > 0 {
		t.Errorf("PMapUintStr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []string{"10"}
	newList = PMapUintStr(someLogicUintStr, []uint{10, 20}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != "0" {
		t.Errorf("PMapUintStr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicUintStr(num uint) string {
	if num == 10 {
		return string("10")
	}
	return "0"
}

func TestPmapUintBool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := PMapUintBool(someLogicUintBool, []uint{10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUintBool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUintBool(nil, nil)) > 0 {
		t.Errorf("PMapUintBool failed")
	}

	if len(PMapUintBool(nil, []uint{})) > 0 {
		t.Errorf("PMapUintBool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []bool{true, false}
	newList = PMapUintBool(someLogicUintBool, []uint{10, 0}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != false {
		t.Errorf("PMapUintBool failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicUintBool(num uint) bool {
	if num > 0 {
		return true
	}
	return false
}

func TestPmapUintFloat32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3, 4}
	newList := PMapUintFloat32(plusOneUintFloat32, []uint{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUintFloat32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUintFloat32(nil, nil)) > 0 {
		t.Errorf("PMapUintFloat32 failed")
	}

	if len(PMapUintFloat32(nil, []uint{})) > 0 {
		t.Errorf("PMapUintFloat32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []float32{2, 3, 4}
	newList = PMapUintFloat32(plusOneUintFloat32, []uint{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUintFloat32 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUintFloat32(num uint) float32 {
	return float32(num + 1)
}

func TestPmapUintFloat64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3, 4}
	newList := PMapUintFloat64(plusOneUintFloat64, []uint{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUintFloat64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUintFloat64(nil, nil)) > 0 {
		t.Errorf("PMapUintFloat64 failed")
	}

	if len(PMapUintFloat64(nil, []uint{})) > 0 {
		t.Errorf("PMapUintFloat64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []float64{2, 3, 4}
	newList = PMapUintFloat64(plusOneUintFloat64, []uint{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUintFloat64 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUintFloat64(num uint) float64 {
	return float64(num + 1)
}

func TestPmapUint64Int(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3, 4}
	newList := PMapUint64Int(plusOneUint64Int, []uint64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint64Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint64Int(nil, nil)) > 0 {
		t.Errorf("PMapUint64Int failed")
	}

	if len(PMapUint64Int(nil, []uint64{})) > 0 {
		t.Errorf("PMapUint64Int failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int{2, 3, 4}
	newList = PMapUint64Int(plusOneUint64Int, []uint64{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint64Int failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint64Int(num uint64) int {
	return int(num + 1)
}

func TestPmapUint64Int64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3, 4}
	newList := PMapUint64Int64(plusOneUint64Int64, []uint64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint64Int64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint64Int64(nil, nil)) > 0 {
		t.Errorf("PMapUint64Int64 failed")
	}

	if len(PMapUint64Int64(nil, []uint64{})) > 0 {
		t.Errorf("PMapUint64Int64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int64{2, 3, 4}
	newList = PMapUint64Int64(plusOneUint64Int64, []uint64{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint64Int64 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint64Int64(num uint64) int64 {
	return int64(num + 1)
}

func TestPmapUint64Int32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3, 4}
	newList := PMapUint64Int32(plusOneUint64Int32, []uint64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint64Int32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint64Int32(nil, nil)) > 0 {
		t.Errorf("PMapUint64Int32 failed")
	}

	if len(PMapUint64Int32(nil, []uint64{})) > 0 {
		t.Errorf("PMapUint64Int32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int32{2, 3, 4}
	newList = PMapUint64Int32(plusOneUint64Int32, []uint64{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint64Int32 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint64Int32(num uint64) int32 {
	return int32(num + 1)
}

func TestPmapUint64Int16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3, 4}
	newList := PMapUint64Int16(plusOneUint64Int16, []uint64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint64Int16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint64Int16(nil, nil)) > 0 {
		t.Errorf("PMapUint64Int16 failed")
	}

	if len(PMapUint64Int16(nil, []uint64{})) > 0 {
		t.Errorf("PMapUint64Int16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int16{2, 3, 4}
	newList = PMapUint64Int16(plusOneUint64Int16, []uint64{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint64Int16 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint64Int16(num uint64) int16 {
	return int16(num + 1)
}

func TestPmapUint64Int8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3, 4}
	newList := PMapUint64Int8(plusOneUint64Int8, []uint64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint64Int8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint64Int8(nil, nil)) > 0 {
		t.Errorf("PMapUint64Int8 failed")
	}

	if len(PMapUint64Int8(nil, []uint64{})) > 0 {
		t.Errorf("PMapUint64Int8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int8{2, 3, 4}
	newList = PMapUint64Int8(plusOneUint64Int8, []uint64{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint64Int8 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint64Int8(num uint64) int8 {
	return int8(num + 1)
}

func TestPmapUint64Uint(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3, 4}
	newList := PMapUint64Uint(plusOneUint64Uint, []uint64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint64Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint64Uint(nil, nil)) > 0 {
		t.Errorf("PMapUint64Uint failed")
	}

	if len(PMapUint64Uint(nil, []uint64{})) > 0 {
		t.Errorf("PMapUint64Uint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint{2, 3, 4}
	newList = PMapUint64Uint(plusOneUint64Uint, []uint64{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint64Uint failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint64Uint(num uint64) uint {
	return uint(num + 1)
}

func TestPmapUint64Uint32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3, 4}
	newList := PMapUint64Uint32(plusOneUint64Uint32, []uint64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint64Uint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint64Uint32(nil, nil)) > 0 {
		t.Errorf("PMapUint64Uint32 failed")
	}

	if len(PMapUint64Uint32(nil, []uint64{})) > 0 {
		t.Errorf("PMapUint64Uint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint32{2, 3, 4}
	newList = PMapUint64Uint32(plusOneUint64Uint32, []uint64{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint64Uint32 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint64Uint32(num uint64) uint32 {
	return uint32(num + 1)
}

func TestPmapUint64Uint16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3, 4}
	newList := PMapUint64Uint16(plusOneUint64Uint16, []uint64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint64Uint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint64Uint16(nil, nil)) > 0 {
		t.Errorf("PMapUint64Uint16 failed")
	}

	if len(PMapUint64Uint16(nil, []uint64{})) > 0 {
		t.Errorf("PMapUint64Uint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint16{2, 3, 4}
	newList = PMapUint64Uint16(plusOneUint64Uint16, []uint64{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint64Uint16 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint64Uint16(num uint64) uint16 {
	return uint16(num + 1)
}

func TestPmapUint64Uint8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3, 4}
	newList := PMapUint64Uint8(plusOneUint64Uint8, []uint64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint64Uint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint64Uint8(nil, nil)) > 0 {
		t.Errorf("PMapUint64Uint8 failed")
	}

	if len(PMapUint64Uint8(nil, []uint64{})) > 0 {
		t.Errorf("PMapUint64Uint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint8{2, 3, 4}
	newList = PMapUint64Uint8(plusOneUint64Uint8, []uint64{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint64Uint8 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint64Uint8(num uint64) uint8 {
	return uint8(num + 1)
}

func TestPmapUint64Str(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList := PMapUint64Str(someLogicUint64Str, []uint64{10})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapUint64Str failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint64Str(nil, nil)) > 0 {
		t.Errorf("PMapUint64Str failed")
	}

	if len(PMapUint64Str(nil, []uint64{})) > 0 {
		t.Errorf("PMapUint64Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []string{"10"}
	newList = PMapUint64Str(someLogicUint64Str, []uint64{10, 20}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != "0" {
		t.Errorf("PMapUint64Str failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicUint64Str(num uint64) string {
	if num == 10 {
		return string("10")
	}
	return "0"
}

func TestPmapUint64Bool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := PMapUint64Bool(someLogicUint64Bool, []uint64{10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint64Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint64Bool(nil, nil)) > 0 {
		t.Errorf("PMapUint64Bool failed")
	}

	if len(PMapUint64Bool(nil, []uint64{})) > 0 {
		t.Errorf("PMapUint64Bool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []bool{true, false}
	newList = PMapUint64Bool(someLogicUint64Bool, []uint64{10, 0}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != false {
		t.Errorf("PMapUint64Bool failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicUint64Bool(num uint64) bool {
	if num > 0 {
		return true
	}
	return false
}

func TestPmapUint64Float32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3, 4}
	newList := PMapUint64Float32(plusOneUint64Float32, []uint64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint64Float32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint64Float32(nil, nil)) > 0 {
		t.Errorf("PMapUint64Float32 failed")
	}

	if len(PMapUint64Float32(nil, []uint64{})) > 0 {
		t.Errorf("PMapUint64Float32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []float32{2, 3, 4}
	newList = PMapUint64Float32(plusOneUint64Float32, []uint64{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint64Float32 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint64Float32(num uint64) float32 {
	return float32(num + 1)
}

func TestPmapUint64Float64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3, 4}
	newList := PMapUint64Float64(plusOneUint64Float64, []uint64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint64Float64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint64Float64(nil, nil)) > 0 {
		t.Errorf("PMapUint64Float64 failed")
	}

	if len(PMapUint64Float64(nil, []uint64{})) > 0 {
		t.Errorf("PMapUint64Float64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []float64{2, 3, 4}
	newList = PMapUint64Float64(plusOneUint64Float64, []uint64{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint64Float64 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint64Float64(num uint64) float64 {
	return float64(num + 1)
}

func TestPmapUint32Int(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3, 4}
	newList := PMapUint32Int(plusOneUint32Int, []uint32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint32Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint32Int(nil, nil)) > 0 {
		t.Errorf("PMapUint32Int failed")
	}

	if len(PMapUint32Int(nil, []uint32{})) > 0 {
		t.Errorf("PMapUint32Int failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int{2, 3, 4}
	newList = PMapUint32Int(plusOneUint32Int, []uint32{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint32Int failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint32Int(num uint32) int {
	return int(num + 1)
}

func TestPmapUint32Int64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3, 4}
	newList := PMapUint32Int64(plusOneUint32Int64, []uint32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint32Int64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint32Int64(nil, nil)) > 0 {
		t.Errorf("PMapUint32Int64 failed")
	}

	if len(PMapUint32Int64(nil, []uint32{})) > 0 {
		t.Errorf("PMapUint32Int64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int64{2, 3, 4}
	newList = PMapUint32Int64(plusOneUint32Int64, []uint32{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint32Int64 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint32Int64(num uint32) int64 {
	return int64(num + 1)
}

func TestPmapUint32Int32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3, 4}
	newList := PMapUint32Int32(plusOneUint32Int32, []uint32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint32Int32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint32Int32(nil, nil)) > 0 {
		t.Errorf("PMapUint32Int32 failed")
	}

	if len(PMapUint32Int32(nil, []uint32{})) > 0 {
		t.Errorf("PMapUint32Int32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int32{2, 3, 4}
	newList = PMapUint32Int32(plusOneUint32Int32, []uint32{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint32Int32 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint32Int32(num uint32) int32 {
	return int32(num + 1)
}

func TestPmapUint32Int16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3, 4}
	newList := PMapUint32Int16(plusOneUint32Int16, []uint32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint32Int16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint32Int16(nil, nil)) > 0 {
		t.Errorf("PMapUint32Int16 failed")
	}

	if len(PMapUint32Int16(nil, []uint32{})) > 0 {
		t.Errorf("PMapUint32Int16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int16{2, 3, 4}
	newList = PMapUint32Int16(plusOneUint32Int16, []uint32{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint32Int16 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint32Int16(num uint32) int16 {
	return int16(num + 1)
}

func TestPmapUint32Int8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3, 4}
	newList := PMapUint32Int8(plusOneUint32Int8, []uint32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint32Int8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint32Int8(nil, nil)) > 0 {
		t.Errorf("PMapUint32Int8 failed")
	}

	if len(PMapUint32Int8(nil, []uint32{})) > 0 {
		t.Errorf("PMapUint32Int8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int8{2, 3, 4}
	newList = PMapUint32Int8(plusOneUint32Int8, []uint32{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint32Int8 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint32Int8(num uint32) int8 {
	return int8(num + 1)
}

func TestPmapUint32Uint(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3, 4}
	newList := PMapUint32Uint(plusOneUint32Uint, []uint32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint32Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint32Uint(nil, nil)) > 0 {
		t.Errorf("PMapUint32Uint failed")
	}

	if len(PMapUint32Uint(nil, []uint32{})) > 0 {
		t.Errorf("PMapUint32Uint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint{2, 3, 4}
	newList = PMapUint32Uint(plusOneUint32Uint, []uint32{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint32Uint failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint32Uint(num uint32) uint {
	return uint(num + 1)
}

func TestPmapUint32Uint64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3, 4}
	newList := PMapUint32Uint64(plusOneUint32Uint64, []uint32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint32Uint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint32Uint64(nil, nil)) > 0 {
		t.Errorf("PMapUint32Uint64 failed")
	}

	if len(PMapUint32Uint64(nil, []uint32{})) > 0 {
		t.Errorf("PMapUint32Uint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint64{2, 3, 4}
	newList = PMapUint32Uint64(plusOneUint32Uint64, []uint32{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint32Uint64 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint32Uint64(num uint32) uint64 {
	return uint64(num + 1)
}

func TestPmapUint32Uint16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3, 4}
	newList := PMapUint32Uint16(plusOneUint32Uint16, []uint32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint32Uint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint32Uint16(nil, nil)) > 0 {
		t.Errorf("PMapUint32Uint16 failed")
	}

	if len(PMapUint32Uint16(nil, []uint32{})) > 0 {
		t.Errorf("PMapUint32Uint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint16{2, 3, 4}
	newList = PMapUint32Uint16(plusOneUint32Uint16, []uint32{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint32Uint16 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint32Uint16(num uint32) uint16 {
	return uint16(num + 1)
}

func TestPmapUint32Uint8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3, 4}
	newList := PMapUint32Uint8(plusOneUint32Uint8, []uint32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint32Uint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint32Uint8(nil, nil)) > 0 {
		t.Errorf("PMapUint32Uint8 failed")
	}

	if len(PMapUint32Uint8(nil, []uint32{})) > 0 {
		t.Errorf("PMapUint32Uint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint8{2, 3, 4}
	newList = PMapUint32Uint8(plusOneUint32Uint8, []uint32{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint32Uint8 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint32Uint8(num uint32) uint8 {
	return uint8(num + 1)
}

func TestPmapUint32Str(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList := PMapUint32Str(someLogicUint32Str, []uint32{10})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapUint32Str failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint32Str(nil, nil)) > 0 {
		t.Errorf("PMapUint32Str failed")
	}

	if len(PMapUint32Str(nil, []uint32{})) > 0 {
		t.Errorf("PMapUint32Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []string{"10"}
	newList = PMapUint32Str(someLogicUint32Str, []uint32{10, 20}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != "0" {
		t.Errorf("PMapUint32Str failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicUint32Str(num uint32) string {
	if num == 10 {
		return string("10")
	}
	return "0"
}

func TestPmapUint32Bool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := PMapUint32Bool(someLogicUint32Bool, []uint32{10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint32Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint32Bool(nil, nil)) > 0 {
		t.Errorf("PMapUint32Bool failed")
	}

	if len(PMapUint32Bool(nil, []uint32{})) > 0 {
		t.Errorf("PMapUint32Bool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []bool{true, false}
	newList = PMapUint32Bool(someLogicUint32Bool, []uint32{10, 0}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != false {
		t.Errorf("PMapUint32Bool failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicUint32Bool(num uint32) bool {
	if num > 0 {
		return true
	}
	return false
}

func TestPmapUint32Float32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3, 4}
	newList := PMapUint32Float32(plusOneUint32Float32, []uint32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint32Float32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint32Float32(nil, nil)) > 0 {
		t.Errorf("PMapUint32Float32 failed")
	}

	if len(PMapUint32Float32(nil, []uint32{})) > 0 {
		t.Errorf("PMapUint32Float32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []float32{2, 3, 4}
	newList = PMapUint32Float32(plusOneUint32Float32, []uint32{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint32Float32 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint32Float32(num uint32) float32 {
	return float32(num + 1)
}

func TestPmapUint32Float64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3, 4}
	newList := PMapUint32Float64(plusOneUint32Float64, []uint32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint32Float64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint32Float64(nil, nil)) > 0 {
		t.Errorf("PMapUint32Float64 failed")
	}

	if len(PMapUint32Float64(nil, []uint32{})) > 0 {
		t.Errorf("PMapUint32Float64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []float64{2, 3, 4}
	newList = PMapUint32Float64(plusOneUint32Float64, []uint32{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint32Float64 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint32Float64(num uint32) float64 {
	return float64(num + 1)
}

func TestPmapUint16Int(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3, 4}
	newList := PMapUint16Int(plusOneUint16Int, []uint16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint16Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint16Int(nil, nil)) > 0 {
		t.Errorf("PMapUint16Int failed")
	}

	if len(PMapUint16Int(nil, []uint16{})) > 0 {
		t.Errorf("PMapUint16Int failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int{2, 3, 4}
	newList = PMapUint16Int(plusOneUint16Int, []uint16{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint16Int failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint16Int(num uint16) int {
	return int(num + 1)
}

func TestPmapUint16Int64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3, 4}
	newList := PMapUint16Int64(plusOneUint16Int64, []uint16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint16Int64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint16Int64(nil, nil)) > 0 {
		t.Errorf("PMapUint16Int64 failed")
	}

	if len(PMapUint16Int64(nil, []uint16{})) > 0 {
		t.Errorf("PMapUint16Int64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int64{2, 3, 4}
	newList = PMapUint16Int64(plusOneUint16Int64, []uint16{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint16Int64 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint16Int64(num uint16) int64 {
	return int64(num + 1)
}

func TestPmapUint16Int32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3, 4}
	newList := PMapUint16Int32(plusOneUint16Int32, []uint16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint16Int32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint16Int32(nil, nil)) > 0 {
		t.Errorf("PMapUint16Int32 failed")
	}

	if len(PMapUint16Int32(nil, []uint16{})) > 0 {
		t.Errorf("PMapUint16Int32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int32{2, 3, 4}
	newList = PMapUint16Int32(plusOneUint16Int32, []uint16{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint16Int32 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint16Int32(num uint16) int32 {
	return int32(num + 1)
}

func TestPmapUint16Int16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3, 4}
	newList := PMapUint16Int16(plusOneUint16Int16, []uint16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint16Int16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint16Int16(nil, nil)) > 0 {
		t.Errorf("PMapUint16Int16 failed")
	}

	if len(PMapUint16Int16(nil, []uint16{})) > 0 {
		t.Errorf("PMapUint16Int16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int16{2, 3, 4}
	newList = PMapUint16Int16(plusOneUint16Int16, []uint16{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint16Int16 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint16Int16(num uint16) int16 {
	return int16(num + 1)
}

func TestPmapUint16Int8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3, 4}
	newList := PMapUint16Int8(plusOneUint16Int8, []uint16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint16Int8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint16Int8(nil, nil)) > 0 {
		t.Errorf("PMapUint16Int8 failed")
	}

	if len(PMapUint16Int8(nil, []uint16{})) > 0 {
		t.Errorf("PMapUint16Int8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int8{2, 3, 4}
	newList = PMapUint16Int8(plusOneUint16Int8, []uint16{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint16Int8 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint16Int8(num uint16) int8 {
	return int8(num + 1)
}

func TestPmapUint16Uint(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3, 4}
	newList := PMapUint16Uint(plusOneUint16Uint, []uint16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint16Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint16Uint(nil, nil)) > 0 {
		t.Errorf("PMapUint16Uint failed")
	}

	if len(PMapUint16Uint(nil, []uint16{})) > 0 {
		t.Errorf("PMapUint16Uint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint{2, 3, 4}
	newList = PMapUint16Uint(plusOneUint16Uint, []uint16{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint16Uint failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint16Uint(num uint16) uint {
	return uint(num + 1)
}

func TestPmapUint16Uint64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3, 4}
	newList := PMapUint16Uint64(plusOneUint16Uint64, []uint16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint16Uint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint16Uint64(nil, nil)) > 0 {
		t.Errorf("PMapUint16Uint64 failed")
	}

	if len(PMapUint16Uint64(nil, []uint16{})) > 0 {
		t.Errorf("PMapUint16Uint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint64{2, 3, 4}
	newList = PMapUint16Uint64(plusOneUint16Uint64, []uint16{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint16Uint64 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint16Uint64(num uint16) uint64 {
	return uint64(num + 1)
}

func TestPmapUint16Uint32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3, 4}
	newList := PMapUint16Uint32(plusOneUint16Uint32, []uint16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint16Uint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint16Uint32(nil, nil)) > 0 {
		t.Errorf("PMapUint16Uint32 failed")
	}

	if len(PMapUint16Uint32(nil, []uint16{})) > 0 {
		t.Errorf("PMapUint16Uint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint32{2, 3, 4}
	newList = PMapUint16Uint32(plusOneUint16Uint32, []uint16{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint16Uint32 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint16Uint32(num uint16) uint32 {
	return uint32(num + 1)
}

func TestPmapUint16Uint8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3, 4}
	newList := PMapUint16Uint8(plusOneUint16Uint8, []uint16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint16Uint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint16Uint8(nil, nil)) > 0 {
		t.Errorf("PMapUint16Uint8 failed")
	}

	if len(PMapUint16Uint8(nil, []uint16{})) > 0 {
		t.Errorf("PMapUint16Uint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint8{2, 3, 4}
	newList = PMapUint16Uint8(plusOneUint16Uint8, []uint16{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint16Uint8 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint16Uint8(num uint16) uint8 {
	return uint8(num + 1)
}

func TestPmapUint16Str(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList := PMapUint16Str(someLogicUint16Str, []uint16{10})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapUint16Str failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint16Str(nil, nil)) > 0 {
		t.Errorf("PMapUint16Str failed")
	}

	if len(PMapUint16Str(nil, []uint16{})) > 0 {
		t.Errorf("PMapUint16Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []string{"10"}
	newList = PMapUint16Str(someLogicUint16Str, []uint16{10, 20}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != "0" {
		t.Errorf("PMapUint16Str failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicUint16Str(num uint16) string {
	if num == 10 {
		return string("10")
	}
	return "0"
}

func TestPmapUint16Bool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := PMapUint16Bool(someLogicUint16Bool, []uint16{10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint16Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint16Bool(nil, nil)) > 0 {
		t.Errorf("PMapUint16Bool failed")
	}

	if len(PMapUint16Bool(nil, []uint16{})) > 0 {
		t.Errorf("PMapUint16Bool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []bool{true, false}
	newList = PMapUint16Bool(someLogicUint16Bool, []uint16{10, 0}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != false {
		t.Errorf("PMapUint16Bool failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicUint16Bool(num uint16) bool {
	if num > 0 {
		return true
	}
	return false
}

func TestPmapUint16Float32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3, 4}
	newList := PMapUint16Float32(plusOneUint16Float32, []uint16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint16Float32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint16Float32(nil, nil)) > 0 {
		t.Errorf("PMapUint16Float32 failed")
	}

	if len(PMapUint16Float32(nil, []uint16{})) > 0 {
		t.Errorf("PMapUint16Float32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []float32{2, 3, 4}
	newList = PMapUint16Float32(plusOneUint16Float32, []uint16{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint16Float32 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint16Float32(num uint16) float32 {
	return float32(num + 1)
}

func TestPmapUint16Float64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3, 4}
	newList := PMapUint16Float64(plusOneUint16Float64, []uint16{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint16Float64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint16Float64(nil, nil)) > 0 {
		t.Errorf("PMapUint16Float64 failed")
	}

	if len(PMapUint16Float64(nil, []uint16{})) > 0 {
		t.Errorf("PMapUint16Float64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []float64{2, 3, 4}
	newList = PMapUint16Float64(plusOneUint16Float64, []uint16{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint16Float64 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint16Float64(num uint16) float64 {
	return float64(num + 1)
}

func TestPmapUint8Int(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3, 4}
	newList := PMapUint8Int(plusOneUint8Int, []uint8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint8Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint8Int(nil, nil)) > 0 {
		t.Errorf("PMapUint8Int failed")
	}

	if len(PMapUint8Int(nil, []uint8{})) > 0 {
		t.Errorf("PMapUint8Int failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int{2, 3, 4}
	newList = PMapUint8Int(plusOneUint8Int, []uint8{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint8Int failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint8Int(num uint8) int {
	return int(num + 1)
}

func TestPmapUint8Int64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3, 4}
	newList := PMapUint8Int64(plusOneUint8Int64, []uint8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint8Int64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint8Int64(nil, nil)) > 0 {
		t.Errorf("PMapUint8Int64 failed")
	}

	if len(PMapUint8Int64(nil, []uint8{})) > 0 {
		t.Errorf("PMapUint8Int64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int64{2, 3, 4}
	newList = PMapUint8Int64(plusOneUint8Int64, []uint8{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint8Int64 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint8Int64(num uint8) int64 {
	return int64(num + 1)
}

func TestPmapUint8Int32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3, 4}
	newList := PMapUint8Int32(plusOneUint8Int32, []uint8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint8Int32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint8Int32(nil, nil)) > 0 {
		t.Errorf("PMapUint8Int32 failed")
	}

	if len(PMapUint8Int32(nil, []uint8{})) > 0 {
		t.Errorf("PMapUint8Int32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int32{2, 3, 4}
	newList = PMapUint8Int32(plusOneUint8Int32, []uint8{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint8Int32 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint8Int32(num uint8) int32 {
	return int32(num + 1)
}

func TestPmapUint8Int16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3, 4}
	newList := PMapUint8Int16(plusOneUint8Int16, []uint8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint8Int16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint8Int16(nil, nil)) > 0 {
		t.Errorf("PMapUint8Int16 failed")
	}

	if len(PMapUint8Int16(nil, []uint8{})) > 0 {
		t.Errorf("PMapUint8Int16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int16{2, 3, 4}
	newList = PMapUint8Int16(plusOneUint8Int16, []uint8{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint8Int16 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint8Int16(num uint8) int16 {
	return int16(num + 1)
}

func TestPmapUint8Int8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3, 4}
	newList := PMapUint8Int8(plusOneUint8Int8, []uint8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint8Int8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint8Int8(nil, nil)) > 0 {
		t.Errorf("PMapUint8Int8 failed")
	}

	if len(PMapUint8Int8(nil, []uint8{})) > 0 {
		t.Errorf("PMapUint8Int8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int8{2, 3, 4}
	newList = PMapUint8Int8(plusOneUint8Int8, []uint8{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint8Int8 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint8Int8(num uint8) int8 {
	return int8(num + 1)
}

func TestPmapUint8Uint(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3, 4}
	newList := PMapUint8Uint(plusOneUint8Uint, []uint8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint8Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint8Uint(nil, nil)) > 0 {
		t.Errorf("PMapUint8Uint failed")
	}

	if len(PMapUint8Uint(nil, []uint8{})) > 0 {
		t.Errorf("PMapUint8Uint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint{2, 3, 4}
	newList = PMapUint8Uint(plusOneUint8Uint, []uint8{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint8Uint failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint8Uint(num uint8) uint {
	return uint(num + 1)
}

func TestPmapUint8Uint64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3, 4}
	newList := PMapUint8Uint64(plusOneUint8Uint64, []uint8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint8Uint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint8Uint64(nil, nil)) > 0 {
		t.Errorf("PMapUint8Uint64 failed")
	}

	if len(PMapUint8Uint64(nil, []uint8{})) > 0 {
		t.Errorf("PMapUint8Uint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint64{2, 3, 4}
	newList = PMapUint8Uint64(plusOneUint8Uint64, []uint8{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint8Uint64 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint8Uint64(num uint8) uint64 {
	return uint64(num + 1)
}

func TestPmapUint8Uint32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3, 4}
	newList := PMapUint8Uint32(plusOneUint8Uint32, []uint8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint8Uint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint8Uint32(nil, nil)) > 0 {
		t.Errorf("PMapUint8Uint32 failed")
	}

	if len(PMapUint8Uint32(nil, []uint8{})) > 0 {
		t.Errorf("PMapUint8Uint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint32{2, 3, 4}
	newList = PMapUint8Uint32(plusOneUint8Uint32, []uint8{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint8Uint32 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint8Uint32(num uint8) uint32 {
	return uint32(num + 1)
}

func TestPmapUint8Uint16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3, 4}
	newList := PMapUint8Uint16(plusOneUint8Uint16, []uint8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint8Uint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint8Uint16(nil, nil)) > 0 {
		t.Errorf("PMapUint8Uint16 failed")
	}

	if len(PMapUint8Uint16(nil, []uint8{})) > 0 {
		t.Errorf("PMapUint8Uint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint16{2, 3, 4}
	newList = PMapUint8Uint16(plusOneUint8Uint16, []uint8{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint8Uint16 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint8Uint16(num uint8) uint16 {
	return uint16(num + 1)
}

func TestPmapUint8Str(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList := PMapUint8Str(someLogicUint8Str, []uint8{10})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapUint8Str failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint8Str(nil, nil)) > 0 {
		t.Errorf("PMapUint8Str failed")
	}

	if len(PMapUint8Str(nil, []uint8{})) > 0 {
		t.Errorf("PMapUint8Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []string{"10"}
	newList = PMapUint8Str(someLogicUint8Str, []uint8{10, 20}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != "0" {
		t.Errorf("PMapUint8Str failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicUint8Str(num uint8) string {
	if num == 10 {
		return string("10")
	}
	return "0"
}

func TestPmapUint8Bool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := PMapUint8Bool(someLogicUint8Bool, []uint8{10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapUint8Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint8Bool(nil, nil)) > 0 {
		t.Errorf("PMapUint8Bool failed")
	}

	if len(PMapUint8Bool(nil, []uint8{})) > 0 {
		t.Errorf("PMapUint8Bool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []bool{true, false}
	newList = PMapUint8Bool(someLogicUint8Bool, []uint8{10, 0}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != false {
		t.Errorf("PMapUint8Bool failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicUint8Bool(num uint8) bool {
	if num > 0 {
		return true
	}
	return false
}

func TestPmapUint8Float32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3, 4}
	newList := PMapUint8Float32(plusOneUint8Float32, []uint8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint8Float32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint8Float32(nil, nil)) > 0 {
		t.Errorf("PMapUint8Float32 failed")
	}

	if len(PMapUint8Float32(nil, []uint8{})) > 0 {
		t.Errorf("PMapUint8Float32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []float32{2, 3, 4}
	newList = PMapUint8Float32(plusOneUint8Float32, []uint8{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint8Float32 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint8Float32(num uint8) float32 {
	return float32(num + 1)
}

func TestPmapUint8Float64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3, 4}
	newList := PMapUint8Float64(plusOneUint8Float64, []uint8{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapUint8Float64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint8Float64(nil, nil)) > 0 {
		t.Errorf("PMapUint8Float64 failed")
	}

	if len(PMapUint8Float64(nil, []uint8{})) > 0 {
		t.Errorf("PMapUint8Float64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []float64{2, 3, 4}
	newList = PMapUint8Float64(plusOneUint8Float64, []uint8{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint8Float64 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneUint8Float64(num uint8) float64 {
	return float64(num + 1)
}

func TestPmapStrInt(t *testing.T) {
	// Test : someLogic
	expectedList := []int{10}
	newList := PMapStrInt(someLogicStrInt, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapStrInt failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapStrInt(nil, nil)) > 0 {
		t.Errorf("PMapStrInt failed")
	}

	if len(PMapStrInt(nil, []string{})) > 0 {
		t.Errorf("PMapStrInt failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int{10}
	newList = PMapStrInt(someLogicStrInt, []string{"ten", "one"}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != 0 {
		t.Errorf("PMapStrInt failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicStrInt(num string) int {
	if num == "ten" {
		return int(10)
	}
	return 0
}

func TestPmapStrInt64(t *testing.T) {
	// Test : someLogic
	expectedList := []int64{10}
	newList := PMapStrInt64(someLogicStrInt64, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapStrInt64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapStrInt64(nil, nil)) > 0 {
		t.Errorf("PMapStrInt64 failed")
	}

	if len(PMapStrInt64(nil, []string{})) > 0 {
		t.Errorf("PMapStrInt64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int64{10}
	newList = PMapStrInt64(someLogicStrInt64, []string{"ten", "one"}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != 0 {
		t.Errorf("PMapStrInt64 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicStrInt64(num string) int64 {
	if num == "ten" {
		return int64(10)
	}
	return 0
}

func TestPmapStrInt32(t *testing.T) {
	// Test : someLogic
	expectedList := []int32{10}
	newList := PMapStrInt32(someLogicStrInt32, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapStrInt32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapStrInt32(nil, nil)) > 0 {
		t.Errorf("PMapStrInt32 failed")
	}

	if len(PMapStrInt32(nil, []string{})) > 0 {
		t.Errorf("PMapStrInt32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int32{10}
	newList = PMapStrInt32(someLogicStrInt32, []string{"ten", "one"}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != 0 {
		t.Errorf("PMapStrInt32 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicStrInt32(num string) int32 {
	if num == "ten" {
		return int32(10)
	}
	return 0
}

func TestPmapStrInt16(t *testing.T) {
	// Test : someLogic
	expectedList := []int16{10}
	newList := PMapStrInt16(someLogicStrInt16, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapStrInt16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapStrInt16(nil, nil)) > 0 {
		t.Errorf("PMapStrInt16 failed")
	}

	if len(PMapStrInt16(nil, []string{})) > 0 {
		t.Errorf("PMapStrInt16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int16{10}
	newList = PMapStrInt16(someLogicStrInt16, []string{"ten", "one"}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != 0 {
		t.Errorf("PMapStrInt16 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicStrInt16(num string) int16 {
	if num == "ten" {
		return int16(10)
	}
	return 0
}

func TestPmapStrInt8(t *testing.T) {
	// Test : someLogic
	expectedList := []int8{10}
	newList := PMapStrInt8(someLogicStrInt8, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapStrInt8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapStrInt8(nil, nil)) > 0 {
		t.Errorf("PMapStrInt8 failed")
	}

	if len(PMapStrInt8(nil, []string{})) > 0 {
		t.Errorf("PMapStrInt8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int8{10}
	newList = PMapStrInt8(someLogicStrInt8, []string{"ten", "one"}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != 0 {
		t.Errorf("PMapStrInt8 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicStrInt8(num string) int8 {
	if num == "ten" {
		return int8(10)
	}
	return 0
}

func TestPmapStrUint(t *testing.T) {
	// Test : someLogic
	expectedList := []uint{10}
	newList := PMapStrUint(someLogicStrUint, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapStrUint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapStrUint(nil, nil)) > 0 {
		t.Errorf("PMapStrUint failed")
	}

	if len(PMapStrUint(nil, []string{})) > 0 {
		t.Errorf("PMapStrUint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint{10}
	newList = PMapStrUint(someLogicStrUint, []string{"ten", "one"}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != 0 {
		t.Errorf("PMapStrUint failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicStrUint(num string) uint {
	if num == "ten" {
		return uint(10)
	}
	return 0
}

func TestPmapStrUint64(t *testing.T) {
	// Test : someLogic
	expectedList := []uint64{10}
	newList := PMapStrUint64(someLogicStrUint64, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapStrUint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapStrUint64(nil, nil)) > 0 {
		t.Errorf("PMapStrUint64 failed")
	}

	if len(PMapStrUint64(nil, []string{})) > 0 {
		t.Errorf("PMapStrUint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint64{10}
	newList = PMapStrUint64(someLogicStrUint64, []string{"ten", "one"}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != 0 {
		t.Errorf("PMapStrUint64 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicStrUint64(num string) uint64 {
	if num == "ten" {
		return uint64(10)
	}
	return 0
}

func TestPmapStrUint32(t *testing.T) {
	// Test : someLogic
	expectedList := []uint32{10}
	newList := PMapStrUint32(someLogicStrUint32, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapStrUint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapStrUint32(nil, nil)) > 0 {
		t.Errorf("PMapStrUint32 failed")
	}

	if len(PMapStrUint32(nil, []string{})) > 0 {
		t.Errorf("PMapStrUint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint32{10}
	newList = PMapStrUint32(someLogicStrUint32, []string{"ten", "one"}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != 0 {
		t.Errorf("PMapStrUint32 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicStrUint32(num string) uint32 {
	if num == "ten" {
		return uint32(10)
	}
	return 0
}

func TestPmapStrUint16(t *testing.T) {
	// Test : someLogic
	expectedList := []uint16{10}
	newList := PMapStrUint16(someLogicStrUint16, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapStrUint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapStrUint16(nil, nil)) > 0 {
		t.Errorf("PMapStrUint16 failed")
	}

	if len(PMapStrUint16(nil, []string{})) > 0 {
		t.Errorf("PMapStrUint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint16{10}
	newList = PMapStrUint16(someLogicStrUint16, []string{"ten", "one"}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != 0 {
		t.Errorf("PMapStrUint16 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicStrUint16(num string) uint16 {
	if num == "ten" {
		return uint16(10)
	}
	return 0
}

func TestPmapStrUint8(t *testing.T) {
	// Test : someLogic
	expectedList := []uint8{10}
	newList := PMapStrUint8(someLogicStrUint8, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapStrUint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapStrUint8(nil, nil)) > 0 {
		t.Errorf("PMapStrUint8 failed")
	}

	if len(PMapStrUint8(nil, []string{})) > 0 {
		t.Errorf("PMapStrUint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint8{10}
	newList = PMapStrUint8(someLogicStrUint8, []string{"ten", "one"}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != 0 {
		t.Errorf("PMapStrUint8 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicStrUint8(num string) uint8 {
	if num == "ten" {
		return uint8(10)
	}
	return 0
}

func TestPmapStrBool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := PMapStrBool(someLogicStrBool, []string{"10", "0"})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapStrBool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapStrBool(nil, nil)) > 0 {
		t.Errorf("PMapStrBool failed")
	}

	if len(PMapStrBool(nil, []string{})) > 0 {
		t.Errorf("PMapStrBool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []bool{true, false}
	newList = PMapStrBool(someLogicStrBool, []string{"10", "0"}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != false {
		t.Errorf("PMapStrBool failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicStrBool(num string) bool {
	if num != "0" {
		return true
	}
	return false
}

func TestPmapStrFloat32(t *testing.T) {
	// Test : someLogic
	expectedList := []float32{10}
	newList := PMapStrFloat32(someLogicStrFloat32, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapStrFloat32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapStrFloat32(nil, nil)) > 0 {
		t.Errorf("PMapStrFloat32 failed")
	}

	if len(PMapStrFloat32(nil, []string{})) > 0 {
		t.Errorf("PMapStrFloat32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []float32{10}
	newList = PMapStrFloat32(someLogicStrFloat32, []string{"ten", "one"}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != 0 {
		t.Errorf("PMapStrFloat32 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicStrFloat32(num string) float32 {
	if num == "ten" {
		return float32(10)
	}
	return 0
}

func TestPmapStrFloat64(t *testing.T) {
	// Test : someLogic
	expectedList := []float64{10}
	newList := PMapStrFloat64(someLogicStrFloat64, []string{"ten"})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapStrFloat64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapStrFloat64(nil, nil)) > 0 {
		t.Errorf("PMapStrFloat64 failed")
	}

	if len(PMapStrFloat64(nil, []string{})) > 0 {
		t.Errorf("PMapStrFloat64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []float64{10}
	newList = PMapStrFloat64(someLogicStrFloat64, []string{"ten", "one"}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != 0 {
		t.Errorf("PMapStrFloat64 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicStrFloat64(num string) float64 {
	if num == "ten" {
		return float64(10)
	}
	return 0
}

func TestPmapBoolInt(t *testing.T) {
	// Test : someLogic
	expectedList := []int{10, 0}
	newList := PMapBoolInt(someLogicBoolInt, []bool{true, false})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolInt failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapBoolInt(nil, nil)) > 0 {
		t.Errorf("PMapBoolInt failed")
	}

	if len(PMapBoolInt(nil, []bool{})) > 0 {
		t.Errorf("PMapBoolInt failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int{10, 0}
	newList = PMapBoolInt(someLogicBoolInt, []bool{true, false}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolInt failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicBoolInt(num bool) int {
	if num == true {
		return 10
	}
	return 0
}

func TestPmapBoolInt64(t *testing.T) {
	// Test : someLogic
	expectedList := []int64{10, 0}
	newList := PMapBoolInt64(someLogicBoolInt64, []bool{true, false})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolInt64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapBoolInt64(nil, nil)) > 0 {
		t.Errorf("PMapBoolInt64 failed")
	}

	if len(PMapBoolInt64(nil, []bool{})) > 0 {
		t.Errorf("PMapBoolInt64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int64{10, 0}
	newList = PMapBoolInt64(someLogicBoolInt64, []bool{true, false}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolInt64 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicBoolInt64(num bool) int64 {
	if num == true {
		return 10
	}
	return 0
}

func TestPmapBoolInt32(t *testing.T) {
	// Test : someLogic
	expectedList := []int32{10, 0}
	newList := PMapBoolInt32(someLogicBoolInt32, []bool{true, false})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolInt32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapBoolInt32(nil, nil)) > 0 {
		t.Errorf("PMapBoolInt32 failed")
	}

	if len(PMapBoolInt32(nil, []bool{})) > 0 {
		t.Errorf("PMapBoolInt32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int32{10, 0}
	newList = PMapBoolInt32(someLogicBoolInt32, []bool{true, false}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolInt32 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicBoolInt32(num bool) int32 {
	if num == true {
		return 10
	}
	return 0
}

func TestPmapBoolInt16(t *testing.T) {
	// Test : someLogic
	expectedList := []int16{10, 0}
	newList := PMapBoolInt16(someLogicBoolInt16, []bool{true, false})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolInt16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapBoolInt16(nil, nil)) > 0 {
		t.Errorf("PMapBoolInt16 failed")
	}

	if len(PMapBoolInt16(nil, []bool{})) > 0 {
		t.Errorf("PMapBoolInt16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int16{10, 0}
	newList = PMapBoolInt16(someLogicBoolInt16, []bool{true, false}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolInt16 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicBoolInt16(num bool) int16 {
	if num == true {
		return 10
	}
	return 0
}

func TestPmapBoolInt8(t *testing.T) {
	// Test : someLogic
	expectedList := []int8{10, 0}
	newList := PMapBoolInt8(someLogicBoolInt8, []bool{true, false})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolInt8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapBoolInt8(nil, nil)) > 0 {
		t.Errorf("PMapBoolInt8 failed")
	}

	if len(PMapBoolInt8(nil, []bool{})) > 0 {
		t.Errorf("PMapBoolInt8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int8{10, 0}
	newList = PMapBoolInt8(someLogicBoolInt8, []bool{true, false}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolInt8 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicBoolInt8(num bool) int8 {
	if num == true {
		return 10
	}
	return 0
}

func TestPmapBoolUint(t *testing.T) {
	// Test : someLogic
	expectedList := []uint{10, 0}
	newList := PMapBoolUint(someLogicBoolUint, []bool{true, false})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolUint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapBoolUint(nil, nil)) > 0 {
		t.Errorf("PMapBoolUint failed")
	}

	if len(PMapBoolUint(nil, []bool{})) > 0 {
		t.Errorf("PMapBoolUint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint{10, 0}
	newList = PMapBoolUint(someLogicBoolUint, []bool{true, false}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolUint failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicBoolUint(num bool) uint {
	if num == true {
		return 10
	}
	return 0
}

func TestPmapBoolUint64(t *testing.T) {
	// Test : someLogic
	expectedList := []uint64{10, 0}
	newList := PMapBoolUint64(someLogicBoolUint64, []bool{true, false})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolUint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapBoolUint64(nil, nil)) > 0 {
		t.Errorf("PMapBoolUint64 failed")
	}

	if len(PMapBoolUint64(nil, []bool{})) > 0 {
		t.Errorf("PMapBoolUint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint64{10, 0}
	newList = PMapBoolUint64(someLogicBoolUint64, []bool{true, false}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolUint64 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicBoolUint64(num bool) uint64 {
	if num == true {
		return 10
	}
	return 0
}

func TestPmapBoolUint32(t *testing.T) {
	// Test : someLogic
	expectedList := []uint32{10, 0}
	newList := PMapBoolUint32(someLogicBoolUint32, []bool{true, false})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolUint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapBoolUint32(nil, nil)) > 0 {
		t.Errorf("PMapBoolUint32 failed")
	}

	if len(PMapBoolUint32(nil, []bool{})) > 0 {
		t.Errorf("PMapBoolUint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint32{10, 0}
	newList = PMapBoolUint32(someLogicBoolUint32, []bool{true, false}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolUint32 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicBoolUint32(num bool) uint32 {
	if num == true {
		return 10
	}
	return 0
}

func TestPmapBoolUint16(t *testing.T) {
	// Test : someLogic
	expectedList := []uint16{10, 0}
	newList := PMapBoolUint16(someLogicBoolUint16, []bool{true, false})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolUint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapBoolUint16(nil, nil)) > 0 {
		t.Errorf("PMapBoolUint16 failed")
	}

	if len(PMapBoolUint16(nil, []bool{})) > 0 {
		t.Errorf("PMapBoolUint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint16{10, 0}
	newList = PMapBoolUint16(someLogicBoolUint16, []bool{true, false}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolUint16 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicBoolUint16(num bool) uint16 {
	if num == true {
		return 10
	}
	return 0
}

func TestPmapBoolUint8(t *testing.T) {
	// Test : someLogic
	expectedList := []uint8{10, 0}
	newList := PMapBoolUint8(someLogicBoolUint8, []bool{true, false})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolUint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapBoolUint8(nil, nil)) > 0 {
		t.Errorf("PMapBoolUint8 failed")
	}

	if len(PMapBoolUint8(nil, []bool{})) > 0 {
		t.Errorf("PMapBoolUint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint8{10, 0}
	newList = PMapBoolUint8(someLogicBoolUint8, []bool{true, false}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolUint8 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicBoolUint8(num bool) uint8 {
	if num == true {
		return 10
	}
	return 0
}

func TestPmapBoolStr(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10", "0"}
	newList := PMapBoolStr(someLogicBoolStr, []bool{true, false})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolStr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapBoolStr(nil, nil)) > 0 {
		t.Errorf("PMapBoolStr failed")
	}

	if len(PMapBoolStr(nil, []bool{})) > 0 {
		t.Errorf("PMapBoolStr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []string{"10", "0"}
	newList = PMapBoolStr(someLogicBoolStr, []bool{true, false}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolStr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicBoolStr(num bool) string {
	if num == true {
		return "10"
	}
	return "0"
}

func TestPmapBoolFloat32(t *testing.T) {
	// Test : someLogic
	expectedList := []float32{10, 0}
	newList := PMapBoolFloat32(someLogicBoolFloat32, []bool{true, false})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolFloat32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapBoolFloat32(nil, nil)) > 0 {
		t.Errorf("PMapBoolFloat32 failed")
	}

	if len(PMapBoolFloat32(nil, []bool{})) > 0 {
		t.Errorf("PMapBoolFloat32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []float32{10, 0}
	newList = PMapBoolFloat32(someLogicBoolFloat32, []bool{true, false}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolFloat32 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicBoolFloat32(num bool) float32 {
	if num == true {
		return 10
	}
	return 0
}

func TestPmapBoolFloat64(t *testing.T) {
	// Test : someLogic
	expectedList := []float64{10, 0}
	newList := PMapBoolFloat64(someLogicBoolFloat64, []bool{true, false})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolFloat64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapBoolFloat64(nil, nil)) > 0 {
		t.Errorf("PMapBoolFloat64 failed")
	}

	if len(PMapBoolFloat64(nil, []bool{})) > 0 {
		t.Errorf("PMapBoolFloat64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []float64{10, 0}
	newList = PMapBoolFloat64(someLogicBoolFloat64, []bool{true, false}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapBoolFloat64 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicBoolFloat64(num bool) float64 {
	if num == true {
		return 10
	}
	return 0
}

func TestPmapFloat32Int(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3, 4}
	newList := PMapFloat32Int(plusOneFloat32Int, []float32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapFloat32Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat32Int(nil, nil)) > 0 {
		t.Errorf("PMapFloat32Int failed")
	}

	if len(PMapFloat32Int(nil, []float32{})) > 0 {
		t.Errorf("PMapFloat32Int failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int{2, 3, 4}
	newList = PMapFloat32Int(plusOneFloat32Int, []float32{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat32Int failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneFloat32Int(num float32) int {
	return int(num + 1)
}

func TestPmapFloat32Int64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3, 4}
	newList := PMapFloat32Int64(plusOneFloat32Int64, []float32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapFloat32Int64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat32Int64(nil, nil)) > 0 {
		t.Errorf("PMapFloat32Int64 failed")
	}

	if len(PMapFloat32Int64(nil, []float32{})) > 0 {
		t.Errorf("PMapFloat32Int64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int64{2, 3, 4}
	newList = PMapFloat32Int64(plusOneFloat32Int64, []float32{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat32Int64 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneFloat32Int64(num float32) int64 {
	return int64(num + 1)
}

func TestPmapFloat32Int32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3, 4}
	newList := PMapFloat32Int32(plusOneFloat32Int32, []float32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapFloat32Int32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat32Int32(nil, nil)) > 0 {
		t.Errorf("PMapFloat32Int32 failed")
	}

	if len(PMapFloat32Int32(nil, []float32{})) > 0 {
		t.Errorf("PMapFloat32Int32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int32{2, 3, 4}
	newList = PMapFloat32Int32(plusOneFloat32Int32, []float32{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat32Int32 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneFloat32Int32(num float32) int32 {
	return int32(num + 1)
}

func TestPmapFloat32Int16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3, 4}
	newList := PMapFloat32Int16(plusOneFloat32Int16, []float32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapFloat32Int16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat32Int16(nil, nil)) > 0 {
		t.Errorf("PMapFloat32Int16 failed")
	}

	if len(PMapFloat32Int16(nil, []float32{})) > 0 {
		t.Errorf("PMapFloat32Int16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int16{2, 3, 4}
	newList = PMapFloat32Int16(plusOneFloat32Int16, []float32{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat32Int16 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneFloat32Int16(num float32) int16 {
	return int16(num + 1)
}

func TestPmapFloat32Int8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3, 4}
	newList := PMapFloat32Int8(plusOneFloat32Int8, []float32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapFloat32Int8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat32Int8(nil, nil)) > 0 {
		t.Errorf("PMapFloat32Int8 failed")
	}

	if len(PMapFloat32Int8(nil, []float32{})) > 0 {
		t.Errorf("PMapFloat32Int8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int8{2, 3, 4}
	newList = PMapFloat32Int8(plusOneFloat32Int8, []float32{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat32Int8 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneFloat32Int8(num float32) int8 {
	return int8(num + 1)
}

func TestPmapFloat32Uint(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3, 4}
	newList := PMapFloat32Uint(plusOneFloat32Uint, []float32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapFloat32Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat32Uint(nil, nil)) > 0 {
		t.Errorf("PMapFloat32Uint failed")
	}

	if len(PMapFloat32Uint(nil, []float32{})) > 0 {
		t.Errorf("PMapFloat32Uint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint{2, 3, 4}
	newList = PMapFloat32Uint(plusOneFloat32Uint, []float32{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat32Uint failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneFloat32Uint(num float32) uint {
	return uint(num + 1)
}

func TestPmapFloat32Uint64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3, 4}
	newList := PMapFloat32Uint64(plusOneFloat32Uint64, []float32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapFloat32Uint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat32Uint64(nil, nil)) > 0 {
		t.Errorf("PMapFloat32Uint64 failed")
	}

	if len(PMapFloat32Uint64(nil, []float32{})) > 0 {
		t.Errorf("PMapFloat32Uint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint64{2, 3, 4}
	newList = PMapFloat32Uint64(plusOneFloat32Uint64, []float32{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat32Uint64 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneFloat32Uint64(num float32) uint64 {
	return uint64(num + 1)
}

func TestPmapFloat32Uint32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3, 4}
	newList := PMapFloat32Uint32(plusOneFloat32Uint32, []float32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapFloat32Uint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat32Uint32(nil, nil)) > 0 {
		t.Errorf("PMapFloat32Uint32 failed")
	}

	if len(PMapFloat32Uint32(nil, []float32{})) > 0 {
		t.Errorf("PMapFloat32Uint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint32{2, 3, 4}
	newList = PMapFloat32Uint32(plusOneFloat32Uint32, []float32{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat32Uint32 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneFloat32Uint32(num float32) uint32 {
	return uint32(num + 1)
}

func TestPmapFloat32Uint16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3, 4}
	newList := PMapFloat32Uint16(plusOneFloat32Uint16, []float32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapFloat32Uint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat32Uint16(nil, nil)) > 0 {
		t.Errorf("PMapFloat32Uint16 failed")
	}

	if len(PMapFloat32Uint16(nil, []float32{})) > 0 {
		t.Errorf("PMapFloat32Uint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint16{2, 3, 4}
	newList = PMapFloat32Uint16(plusOneFloat32Uint16, []float32{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat32Uint16 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneFloat32Uint16(num float32) uint16 {
	return uint16(num + 1)
}

func TestPmapFloat32Uint8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3, 4}
	newList := PMapFloat32Uint8(plusOneFloat32Uint8, []float32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapFloat32Uint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat32Uint8(nil, nil)) > 0 {
		t.Errorf("PMapFloat32Uint8 failed")
	}

	if len(PMapFloat32Uint8(nil, []float32{})) > 0 {
		t.Errorf("PMapFloat32Uint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint8{2, 3, 4}
	newList = PMapFloat32Uint8(plusOneFloat32Uint8, []float32{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat32Uint8 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneFloat32Uint8(num float32) uint8 {
	return uint8(num + 1)
}

func TestPmapFloat32Str(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList := PMapFloat32Str(someLogicFloat32Str, []float32{10})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapFloat32Str failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat32Str(nil, nil)) > 0 {
		t.Errorf("PMapFloat32Str failed")
	}

	if len(PMapFloat32Str(nil, []float32{})) > 0 {
		t.Errorf("PMapFloat32Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []string{"10"}
	newList = PMapFloat32Str(someLogicFloat32Str, []float32{10, 20}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != "0" {
		t.Errorf("PMapFloat32Str failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicFloat32Str(num float32) string {
	if num == 10 {
		return string("10")
	}
	return "0"
}

func TestPmapFloat32Bool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := PMapFloat32Bool(someLogicFloat32Bool, []float32{10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapFloat32Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat32Bool(nil, nil)) > 0 {
		t.Errorf("PMapFloat32Bool failed")
	}

	if len(PMapFloat32Bool(nil, []float32{})) > 0 {
		t.Errorf("PMapFloat32Bool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []bool{true, false}
	newList = PMapFloat32Bool(someLogicFloat32Bool, []float32{10, 0}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != false {
		t.Errorf("PMapFloat32Bool failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicFloat32Bool(num float32) bool {
	if num > 0 {
		return true
	}
	return false
}

func TestPmapFloat32Float64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float64{2, 3, 4}
	newList := PMapFloat32Float64(plusOneFloat32Float64, []float32{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapFloat32Float64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat32Float64(nil, nil)) > 0 {
		t.Errorf("PMapFloat32Float64 failed")
	}

	if len(PMapFloat32Float64(nil, []float32{})) > 0 {
		t.Errorf("PMapFloat32Float64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []float64{2, 3, 4}
	newList = PMapFloat32Float64(plusOneFloat32Float64, []float32{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat32Float64 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneFloat32Float64(num float32) float64 {
	return float64(num + 1)
}

func TestPmapFloat64Int(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int{2, 3, 4}
	newList := PMapFloat64Int(plusOneFloat64Int, []float64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapFloat64Int failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat64Int(nil, nil)) > 0 {
		t.Errorf("PMapFloat64Int failed")
	}

	if len(PMapFloat64Int(nil, []float64{})) > 0 {
		t.Errorf("PMapFloat64Int failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int{2, 3, 4}
	newList = PMapFloat64Int(plusOneFloat64Int, []float64{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat64Int failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneFloat64Int(num float64) int {
	return int(num + 1)
}

func TestPmapFloat64Int64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int64{2, 3, 4}
	newList := PMapFloat64Int64(plusOneFloat64Int64, []float64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapFloat64Int64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat64Int64(nil, nil)) > 0 {
		t.Errorf("PMapFloat64Int64 failed")
	}

	if len(PMapFloat64Int64(nil, []float64{})) > 0 {
		t.Errorf("PMapFloat64Int64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int64{2, 3, 4}
	newList = PMapFloat64Int64(plusOneFloat64Int64, []float64{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat64Int64 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneFloat64Int64(num float64) int64 {
	return int64(num + 1)
}

func TestPmapFloat64Int32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int32{2, 3, 4}
	newList := PMapFloat64Int32(plusOneFloat64Int32, []float64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapFloat64Int32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat64Int32(nil, nil)) > 0 {
		t.Errorf("PMapFloat64Int32 failed")
	}

	if len(PMapFloat64Int32(nil, []float64{})) > 0 {
		t.Errorf("PMapFloat64Int32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int32{2, 3, 4}
	newList = PMapFloat64Int32(plusOneFloat64Int32, []float64{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat64Int32 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneFloat64Int32(num float64) int32 {
	return int32(num + 1)
}

func TestPmapFloat64Int16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int16{2, 3, 4}
	newList := PMapFloat64Int16(plusOneFloat64Int16, []float64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapFloat64Int16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat64Int16(nil, nil)) > 0 {
		t.Errorf("PMapFloat64Int16 failed")
	}

	if len(PMapFloat64Int16(nil, []float64{})) > 0 {
		t.Errorf("PMapFloat64Int16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int16{2, 3, 4}
	newList = PMapFloat64Int16(plusOneFloat64Int16, []float64{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat64Int16 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneFloat64Int16(num float64) int16 {
	return int16(num + 1)
}

func TestPmapFloat64Int8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []int8{2, 3, 4}
	newList := PMapFloat64Int8(plusOneFloat64Int8, []float64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapFloat64Int8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat64Int8(nil, nil)) > 0 {
		t.Errorf("PMapFloat64Int8 failed")
	}

	if len(PMapFloat64Int8(nil, []float64{})) > 0 {
		t.Errorf("PMapFloat64Int8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []int8{2, 3, 4}
	newList = PMapFloat64Int8(plusOneFloat64Int8, []float64{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat64Int8 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneFloat64Int8(num float64) int8 {
	return int8(num + 1)
}

func TestPmapFloat64Uint(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint{2, 3, 4}
	newList := PMapFloat64Uint(plusOneFloat64Uint, []float64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapFloat64Uint failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat64Uint(nil, nil)) > 0 {
		t.Errorf("PMapFloat64Uint failed")
	}

	if len(PMapFloat64Uint(nil, []float64{})) > 0 {
		t.Errorf("PMapFloat64Uint failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint{2, 3, 4}
	newList = PMapFloat64Uint(plusOneFloat64Uint, []float64{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat64Uint failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneFloat64Uint(num float64) uint {
	return uint(num + 1)
}

func TestPmapFloat64Uint64(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint64{2, 3, 4}
	newList := PMapFloat64Uint64(plusOneFloat64Uint64, []float64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapFloat64Uint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat64Uint64(nil, nil)) > 0 {
		t.Errorf("PMapFloat64Uint64 failed")
	}

	if len(PMapFloat64Uint64(nil, []float64{})) > 0 {
		t.Errorf("PMapFloat64Uint64 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint64{2, 3, 4}
	newList = PMapFloat64Uint64(plusOneFloat64Uint64, []float64{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat64Uint64 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneFloat64Uint64(num float64) uint64 {
	return uint64(num + 1)
}

func TestPmapFloat64Uint32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint32{2, 3, 4}
	newList := PMapFloat64Uint32(plusOneFloat64Uint32, []float64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapFloat64Uint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat64Uint32(nil, nil)) > 0 {
		t.Errorf("PMapFloat64Uint32 failed")
	}

	if len(PMapFloat64Uint32(nil, []float64{})) > 0 {
		t.Errorf("PMapFloat64Uint32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint32{2, 3, 4}
	newList = PMapFloat64Uint32(plusOneFloat64Uint32, []float64{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat64Uint32 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneFloat64Uint32(num float64) uint32 {
	return uint32(num + 1)
}

func TestPmapFloat64Uint16(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint16{2, 3, 4}
	newList := PMapFloat64Uint16(plusOneFloat64Uint16, []float64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapFloat64Uint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat64Uint16(nil, nil)) > 0 {
		t.Errorf("PMapFloat64Uint16 failed")
	}

	if len(PMapFloat64Uint16(nil, []float64{})) > 0 {
		t.Errorf("PMapFloat64Uint16 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint16{2, 3, 4}
	newList = PMapFloat64Uint16(plusOneFloat64Uint16, []float64{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat64Uint16 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneFloat64Uint16(num float64) uint16 {
	return uint16(num + 1)
}

func TestPmapFloat64Uint8(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []uint8{2, 3, 4}
	newList := PMapFloat64Uint8(plusOneFloat64Uint8, []float64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapFloat64Uint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat64Uint8(nil, nil)) > 0 {
		t.Errorf("PMapFloat64Uint8 failed")
	}

	if len(PMapFloat64Uint8(nil, []float64{})) > 0 {
		t.Errorf("PMapFloat64Uint8 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []uint8{2, 3, 4}
	newList = PMapFloat64Uint8(plusOneFloat64Uint8, []float64{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat64Uint8 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneFloat64Uint8(num float64) uint8 {
	return uint8(num + 1)
}

func TestPmapFloat64Str(t *testing.T) {
	// Test : someLogic
	expectedList := []string{"10"}
	newList := PMapFloat64Str(someLogicFloat64Str, []float64{10})

	if newList[0] != expectedList[0] {
		t.Errorf("PMapFloat64Str failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat64Str(nil, nil)) > 0 {
		t.Errorf("PMapFloat64Str failed")
	}

	if len(PMapFloat64Str(nil, []float64{})) > 0 {
		t.Errorf("PMapFloat64Str failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []string{"10"}
	newList = PMapFloat64Str(someLogicFloat64Str, []float64{10, 20}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != "0" {
		t.Errorf("PMapFloat64Str failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicFloat64Str(num float64) string {
	if num == 10 {
		return string("10")
	}
	return "0"
}

func TestPmapFloat64Bool(t *testing.T) {
	// Test : someLogic
	expectedList := []bool{true, false}
	newList := PMapFloat64Bool(someLogicFloat64Bool, []float64{10, 0})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] {
		t.Errorf("PMapFloat64Bool failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat64Bool(nil, nil)) > 0 {
		t.Errorf("PMapFloat64Bool failed")
	}

	if len(PMapFloat64Bool(nil, []float64{})) > 0 {
		t.Errorf("PMapFloat64Bool failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []bool{true, false}
	newList = PMapFloat64Bool(someLogicFloat64Bool, []float64{10, 0}, Optional{FixedPool: 1, RandomOrder: true})

	if newList[0] != expectedList[0] || newList[1] != false {
		t.Errorf("PMapFloat64Bool failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func someLogicFloat64Bool(num float64) bool {
	if num > 0 {
		return true
	}
	return false
}

func TestPmapFloat64Float32(t *testing.T) {
	// Test : add 1 to the list
	expectedList := []float32{2, 3, 4}
	newList := PMapFloat64Float32(plusOneFloat64Float32, []float64{1, 2, 3})

	if newList[0] != expectedList[0] || newList[1] != expectedList[1] || newList[2] != expectedList[2] {
		t.Errorf("PMapFloat64Float32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat64Float32(nil, nil)) > 0 {
		t.Errorf("PMapFloat64Float32 failed")
	}

	if len(PMapFloat64Float32(nil, []float64{})) > 0 {
		t.Errorf("PMapFloat64Float32 failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []float32{2, 3, 4}
	newList = PMapFloat64Float32(plusOneFloat64Float32, []float64{1, 2, 3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if v == x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat64Float32 failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func plusOneFloat64Float32(num float64) float32 {
	return float32(num + 1)
}
