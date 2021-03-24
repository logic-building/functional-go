package fp

import "testing"

func TestReduceInt(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	expected := 15
	actual := ReduceInt(plusInt, list)
	if actual != expected {
		t.Errorf("ReduceInt failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int{1, 2, 3, 4, 5}
	expected = 18
	actual = ReduceInt(plusInt, list, 3)
	if actual != expected {
		t.Errorf("ReduceInt failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int{1, 2}
	expected = 3
	actual = ReduceInt(plusInt, list)
	if actual != expected {
		t.Errorf("ReduceInt failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int{1}
	expected = 1
	actual = ReduceInt(plusInt, list)
	if actual != expected {
		t.Errorf("ReduceInt failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int{}
	expected = 0
	actual = ReduceInt(plusInt, list)
	if actual != expected {
		t.Errorf("ReduceInt failed. actual=%v, expected=%v", actual, expected)
	}
}

func plusInt(num1, num2 int) int {
	return num1 + num2
}

func TestReduceInt64(t *testing.T) {
	list := []int64{1, 2, 3, 4, 5}
	var expected int64 = 15
	actual := ReduceInt64(plusInt64, list)
	if actual != expected {
		t.Errorf("ReduceInt64 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int64{1, 2, 3, 4, 5}
	expected = 18
	actual = ReduceInt64(plusInt64, list, 3)
	if actual != expected {
		t.Errorf("ReduceInt64 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int64{1, 2}
	expected = 3
	actual = ReduceInt64(plusInt64, list)
	if actual != expected {
		t.Errorf("ReduceInt64 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int64{1}
	expected = 1
	actual = ReduceInt64(plusInt64, list)
	if actual != expected {
		t.Errorf("ReduceInt64 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int64{}
	expected = 0
	actual = ReduceInt64(plusInt64, list)
	if actual != expected {
		t.Errorf("ReduceInt64 failed. actual=%v, expected=%v", actual, expected)
	}
}

func plusInt64(num1, num2 int64) int64 {
	return num1 + num2
}

func TestReduceInt32(t *testing.T) {
	list := []int32{1, 2, 3, 4, 5}
	var expected int32 = 15
	actual := ReduceInt32(plusInt32, list)
	if actual != expected {
		t.Errorf("ReduceInt32 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int32{1, 2, 3, 4, 5}
	expected = 18
	actual = ReduceInt32(plusInt32, list, 3)
	if actual != expected {
		t.Errorf("ReduceInt32 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int32{1, 2}
	expected = 3
	actual = ReduceInt32(plusInt32, list)
	if actual != expected {
		t.Errorf("ReduceInt32 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int32{1}
	expected = 1
	actual = ReduceInt32(plusInt32, list)
	if actual != expected {
		t.Errorf("ReduceInt32 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int32{}
	expected = 0
	actual = ReduceInt32(plusInt32, list)
	if actual != expected {
		t.Errorf("ReduceInt32 failed. actual=%v, expected=%v", actual, expected)
	}
}

func plusInt32(num1, num2 int32) int32 {
	return num1 + num2
}

func TestReduceInt16(t *testing.T) {
	list := []int16{1, 2, 3, 4, 5}
	var expected int16 = 15
	actual := ReduceInt16(plusInt16, list)
	if actual != expected {
		t.Errorf("ReduceInt16 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int16{1, 2, 3, 4, 5}
	expected = 18
	actual = ReduceInt16(plusInt16, list, 3)
	if actual != expected {
		t.Errorf("ReduceInt16 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int16{1, 2}
	expected = 3
	actual = ReduceInt16(plusInt16, list)
	if actual != expected {
		t.Errorf("ReduceInt16 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int16{1}
	expected = 1
	actual = ReduceInt16(plusInt16, list)
	if actual != expected {
		t.Errorf("ReduceInt16 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int16{}
	expected = 0
	actual = ReduceInt16(plusInt16, list)
	if actual != expected {
		t.Errorf("ReduceInt16 failed. actual=%v, expected=%v", actual, expected)
	}
}

func plusInt16(num1, num2 int16) int16 {
	return num1 + num2
}

func TestReduceInt8(t *testing.T) {
	list := []int8{1, 2, 3, 4, 5}
	var expected int8 = 15
	actual := ReduceInt8(plusInt8, list)
	if actual != expected {
		t.Errorf("ReduceInt8 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int8{1, 2, 3, 4, 5}
	expected = 18
	actual = ReduceInt8(plusInt8, list, 3)
	if actual != expected {
		t.Errorf("ReduceInt8 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int8{1, 2}
	expected = 3
	actual = ReduceInt8(plusInt8, list)
	if actual != expected {
		t.Errorf("ReduceInt8 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int8{1}
	expected = 1
	actual = ReduceInt8(plusInt8, list)
	if actual != expected {
		t.Errorf("ReduceInt8 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []int8{}
	expected = 0
	actual = ReduceInt8(plusInt8, list)
	if actual != expected {
		t.Errorf("ReduceInt8 failed. actual=%v, expected=%v", actual, expected)
	}
}

func plusInt8(num1, num2 int8) int8 {
	return num1 + num2
}

func TestReduceUint(t *testing.T) {
	list := []uint{1, 2, 3, 4, 5}
	var expected uint = 15
	actual := ReduceUint(plusUint, list)
	if actual != expected {
		t.Errorf("ReduceUint failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint{1, 2, 3, 4, 5}
	expected = 18
	actual = ReduceUint(plusUint, list, 3)
	if actual != expected {
		t.Errorf("ReduceUint failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint{1, 2}
	expected = 3
	actual = ReduceUint(plusUint, list)
	if actual != expected {
		t.Errorf("ReduceUint failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint{1}
	expected = 1
	actual = ReduceUint(plusUint, list)
	if actual != expected {
		t.Errorf("ReduceUint failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint{}
	expected = 0
	actual = ReduceUint(plusUint, list)
	if actual != expected {
		t.Errorf("ReduceUint failed. actual=%v, expected=%v", actual, expected)
	}
}

func plusUint(num1, num2 uint) uint {
	return num1 + num2
}

func TestReduceUint64(t *testing.T) {
	list := []uint64{1, 2, 3, 4, 5}
	var expected uint64 = 15
	actual := ReduceUint64(plusUint64, list)
	if actual != expected {
		t.Errorf("ReduceUint64 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint64{1, 2, 3, 4, 5}
	expected = 18
	actual = ReduceUint64(plusUint64, list, 3)
	if actual != expected {
		t.Errorf("ReduceUint64 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint64{1, 2}
	expected = 3
	actual = ReduceUint64(plusUint64, list)
	if actual != expected {
		t.Errorf("ReduceUint64 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint64{1}
	expected = 1
	actual = ReduceUint64(plusUint64, list)
	if actual != expected {
		t.Errorf("ReduceUint64 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint64{}
	expected = 0
	actual = ReduceUint64(plusUint64, list)
	if actual != expected {
		t.Errorf("ReduceUint64 failed. actual=%v, expected=%v", actual, expected)
	}
}

func plusUint64(num1, num2 uint64) uint64 {
	return num1 + num2
}

func TestReduceUint32(t *testing.T) {
	list := []uint32{1, 2, 3, 4, 5}
	var expected uint32 = 15
	actual := ReduceUint32(plusUint32, list)
	if actual != expected {
		t.Errorf("ReduceUint32 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint32{1, 2, 3, 4, 5}
	expected = 18
	actual = ReduceUint32(plusUint32, list, 3)
	if actual != expected {
		t.Errorf("ReduceUint32 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint32{1, 2}
	expected = 3
	actual = ReduceUint32(plusUint32, list)
	if actual != expected {
		t.Errorf("ReduceUint32 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint32{1}
	expected = 1
	actual = ReduceUint32(plusUint32, list)
	if actual != expected {
		t.Errorf("ReduceUint32 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint32{}
	expected = 0
	actual = ReduceUint32(plusUint32, list)
	if actual != expected {
		t.Errorf("ReduceUint32 failed. actual=%v, expected=%v", actual, expected)
	}
}

func plusUint32(num1, num2 uint32) uint32 {
	return num1 + num2
}

func TestReduceUint16(t *testing.T) {
	list := []uint16{1, 2, 3, 4, 5}
	var expected uint16 = 15
	actual := ReduceUint16(plusUint16, list)
	if actual != expected {
		t.Errorf("ReduceUint16 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint16{1, 2, 3, 4, 5}
	expected = 18
	actual = ReduceUint16(plusUint16, list, 3)
	if actual != expected {
		t.Errorf("ReduceUint16 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint16{1, 2}
	expected = 3
	actual = ReduceUint16(plusUint16, list)
	if actual != expected {
		t.Errorf("ReduceUint16 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint16{1}
	expected = 1
	actual = ReduceUint16(plusUint16, list)
	if actual != expected {
		t.Errorf("ReduceUint16 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint16{}
	expected = 0
	actual = ReduceUint16(plusUint16, list)
	if actual != expected {
		t.Errorf("ReduceUint16 failed. actual=%v, expected=%v", actual, expected)
	}
}

func plusUint16(num1, num2 uint16) uint16 {
	return num1 + num2
}

func TestReduceUint8(t *testing.T) {
	list := []uint8{1, 2, 3, 4, 5}
	var expected uint8 = 15
	actual := ReduceUint8(plusUint8, list)
	if actual != expected {
		t.Errorf("ReduceUint8 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint8{1, 2, 3, 4, 5}
	expected = 18
	actual = ReduceUint8(plusUint8, list, 3)
	if actual != expected {
		t.Errorf("ReduceUint8 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint8{1, 2}
	expected = 3
	actual = ReduceUint8(plusUint8, list)
	if actual != expected {
		t.Errorf("ReduceUint8 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint8{1}
	expected = 1
	actual = ReduceUint8(plusUint8, list)
	if actual != expected {
		t.Errorf("ReduceUint8 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []uint8{}
	expected = 0
	actual = ReduceUint8(plusUint8, list)
	if actual != expected {
		t.Errorf("ReduceUint8 failed. actual=%v, expected=%v", actual, expected)
	}
}

func plusUint8(num1, num2 uint8) uint8 {
	return num1 + num2
}

func TestReduceFloat64(t *testing.T) {
	list := []float64{1, 2, 3, 4, 5}
	var expected float64 = 15
	actual := ReduceFloat64(plusFloat64, list)
	if actual != expected {
		t.Errorf("ReduceFloat64 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []float64{1, 2, 3, 4, 5}
	expected = 18
	actual = ReduceFloat64(plusFloat64, list, 3)
	if actual != expected {
		t.Errorf("ReduceFloat64 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []float64{1, 2}
	expected = 3
	actual = ReduceFloat64(plusFloat64, list)
	if actual != expected {
		t.Errorf("ReduceFloat64 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []float64{1}
	expected = 1
	actual = ReduceFloat64(plusFloat64, list)
	if actual != expected {
		t.Errorf("ReduceFloat64 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []float64{}
	expected = 0
	actual = ReduceFloat64(plusFloat64, list)
	if actual != expected {
		t.Errorf("ReduceFloat64 failed. actual=%v, expected=%v", actual, expected)
	}
}

func plusFloat64(num1, num2 float64) float64 {
	return num1 + num2
}

func TestReduceFloat32(t *testing.T) {
	list := []float32{1, 2, 3, 4, 5}
	var expected float32 = 15
	actual := ReduceFloat32(plusFloat32, list)
	if actual != expected {
		t.Errorf("ReduceFloat32 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []float32{1, 2, 3, 4, 5}
	expected = 18
	actual = ReduceFloat32(plusFloat32, list, 3)
	if actual != expected {
		t.Errorf("ReduceFloat64 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []float32{1, 2}
	expected = 3
	actual = ReduceFloat32(plusFloat32, list)
	if actual != expected {
		t.Errorf("ReduceFloat64 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []float32{1}
	expected = 1
	actual = ReduceFloat32(plusFloat32, list)
	if actual != expected {
		t.Errorf("ReduceFloat64 failed. actual=%v, expected=%v", actual, expected)
	}

	list = []float32{}
	expected = 0
	actual = ReduceFloat32(plusFloat32, list)
	if actual != expected {
		t.Errorf("ReduceFloat64 failed. actual=%v, expected=%v", actual, expected)
	}
}

func plusFloat32(num1, num2 float32) float32 {
	return num1 + num2
}

func TestReduceStr(t *testing.T) {
	list := []string{"Nandeshwar", "Kumar", "Sah"}
	var expected = "NandeshwarKumarSah"
	actual := ReduceStr(plusStr, list)
	if actual != expected {
		t.Errorf("ReduceStr failed. actual=%v, expected=%v", actual, expected)
	}

	list = []string{"Nandeshwar", "Kumar", "Sah"}
	expected = "...NandeshwarKumarSah"
	actual = ReduceStr(plusStr, list, "...")
	if actual != expected {
		t.Errorf("ReduceStr failed. actual=%v, expected=%v", actual, expected)
	}

	list = []string{"Nandeshwar", "Sah"}
	expected = "NandeshwarSah"
	actual = ReduceStr(plusStr, list)
	if actual != expected {
		t.Errorf("ReduceStr failed. actual=%v, expected=%v", actual, expected)
	}

	list = []string{"Nandeshwar"}
	expected = "Nandeshwar"
	actual = ReduceStr(plusStr, list)
	if actual != expected {
		t.Errorf("ReduceStr failed. actual=%v, expected=%v", actual, expected)
	}

	list = []string{}
	expected = ""
	actual = ReduceStr(plusStr, list)
	if actual != expected {
		t.Errorf("ReduceStr failed. actual=%v, expected=%v", actual, expected)
	}
}

func plusStr(num1, num2 string) string {
	return num1 + num2
}
