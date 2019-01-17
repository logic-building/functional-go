package fp

import (
	"strings"
	"testing"
)

func TestPmapInt(t *testing.T) {
	// Test : square the list
	expectedSquareList := []int{1, 4, 9}
	squareList := PmapInt(squareInt, []int{1, 2, 3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("PmapInt failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []int{6, 7, 8}
	partialAddInt := func(num int) int {
		return addInt(5, num)
	}
	sumList := PmapInt(partialAddInt, []int{1, 2, 3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("PmapInt failed.expected=%v, actual=%v", expectedSumList, sumList)
	}

	if len(PmapInt(nil, nil)) > 0 {
		t.Errorf("PmapInt failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	if len(PmapInt(nil, []int{})) > 0 {
		t.Errorf("PmapInt failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func TestPmapInt64(t *testing.T) {
	// Test : square the list
	expectedSquareList := []int64{1, 4, 9}
	squareList := PmapInt64(squareInt64, []int64{1, 2, 3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("PmapInt64 failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []int64{6, 7, 8}
	partialAddInt := func(num int64) int64 {
		return addInt64(5, num)
	}
	sumList := PmapInt64(partialAddInt, []int64{1, 2, 3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("PmapInt64 failed.expected=%v, actual=%v", expectedSumList, sumList)
	}

	if len(PmapInt64(nil, nil)) > 0 {
		t.Errorf("PmapInt64 failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	if len(PmapInt64(nil, []int64{})) > 0 {
		t.Errorf("PmapInt64 failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func TestPmapInt32(t *testing.T) {
	// Test : square the list
	expectedSquareList := []int32{1, 4, 9}
	squareList := PmapInt32(squareInt32, []int32{1, 2, 3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("PmapInt32 failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []int64{6, 7, 8}
	partialAddInt := func(num int64) int64 {
		return addInt64(5, num)
	}
	sumList := PmapInt64(partialAddInt, []int64{1, 2, 3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("PmapInt32 failed.expected=%v, actual=%v", expectedSumList, sumList)
	}

	if len(PmapInt32(nil, nil)) > 0 {
		t.Errorf("PmapInt32 failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	if len(PmapInt32(nil, []int32{})) > 0 {
		t.Errorf("PmapInt32 failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func TestPmapInt16(t *testing.T) {
	// Test : square the list
	expectedSquareList := []int16{1, 4, 9}
	squareList := PmapInt16(squareInt16, []int16{1, 2, 3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("PmapInt16 failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []int16{6, 7, 8}
	partialAddInt := func(num int16) int16 {
		return addInt16(5, num)
	}
	sumList := PmapInt16(partialAddInt, []int16{1, 2, 3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("PmapInt16 failed.expected=%v, actual=%v", expectedSumList, sumList)
	}

	if len(PmapInt16(nil, nil)) > 0 {
		t.Errorf("PmapInt16 failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	if len(PmapInt16(nil, []int16{})) > 0 {
		t.Errorf("PmapInt16 failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func TestPmapInt8(t *testing.T) {
	// Test : square the list
	expectedSquareList := []int8{1, 4, 9}
	squareList := PmapInt8(squareInt8, []int8{1, 2, 3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("PmapInt8 failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []int8{6, 7, 8}
	partialAddInt := func(num int8) int8 {
		return addInt8(5, num)
	}
	sumList := PmapInt8(partialAddInt, []int8{1, 2, 3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("PmapInt8 failed.expected=%v, actual=%v", expectedSumList, sumList)
	}

	if len(PmapInt8(nil, nil)) > 0 {
		t.Errorf("PmapInt8 failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	if len(PmapInt8(nil, []int8{})) > 0 {
		t.Errorf("PmapInt8 failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func TestPmapUint(t *testing.T) {
	// Test : square the list
	expectedSquareList := []uint{1, 4, 9}
	squareList := PmapUint(squareUint, []uint{1, 2, 3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("PmapUint failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []uint{6, 7, 8}
	partialAddInt := func(num uint) uint {
		return addUint(5, num)
	}
	sumList := PmapUint(partialAddInt, []uint{1, 2, 3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("PmapUint failed.expected=%v, actual=%v", expectedSumList, sumList)
	}

	if len(PmapUint(nil, nil)) > 0 {
		t.Errorf("PmapUint failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	if len(PmapUint(nil, []uint{})) > 0 {
		t.Errorf("PmapUint failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func TestPmapUint64(t *testing.T) {
	// Test : square the list
	expectedSquareList := []uint64{1, 4, 9}
	squareList := PmapUint64(squareUint64, []uint64{1, 2, 3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("PmapUint64 failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []uint64{6, 7, 8}
	partialAddInt := func(num uint64) uint64 {
		return addUint64(5, num)
	}
	sumList := PmapUint64(partialAddInt, []uint64{1, 2, 3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("PmapUint64 failed.expected=%v, actual=%v", expectedSumList, sumList)
	}

	if len(PmapUint64(nil, nil)) > 0 {
		t.Errorf("PmapUint64 failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	if len(PmapUint64(nil, []uint64{})) > 0 {
		t.Errorf("PmapUint64 failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func TestPmapUint32(t *testing.T) {
	// Test : square the list
	expectedSquareList := []uint32{1, 4, 9}
	squareList := PmapUint32(squareUint32, []uint32{1, 2, 3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("PmapUint32 failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []uint32{6, 7, 8}
	partialAddInt := func(num uint32) uint32 {
		return addUint32(5, num)
	}
	sumList := PmapUint32(partialAddInt, []uint32{1, 2, 3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("PmapUint32 failed.expected=%v, actual=%v", expectedSumList, sumList)
	}

	if len(PmapUint32(nil, nil)) > 0 {
		t.Errorf("PmapUint32 failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	if len(PmapUint32(nil, []uint32{})) > 0 {
		t.Errorf("PmapUint32 failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func TestPmapUint16(t *testing.T) {
	// Test : square the list
	expectedSquareList := []uint16{1, 4, 9}
	squareList := PmapUint16(squareUint16, []uint16{1, 2, 3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("PmapUint16 failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []uint16{6, 7, 8}
	partialAddInt := func(num uint16) uint16 {
		return addUint16(5, num)
	}
	sumList := PmapUint16(partialAddInt, []uint16{1, 2, 3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("PmapUint16 failed.expected=%v, actual=%v", expectedSumList, sumList)
	}

	if len(PmapUint16(nil, nil)) > 0 {
		t.Errorf("PmapUint16 failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	if len(PmapUint16(nil, []uint16{})) > 0 {
		t.Errorf("PmapUint16 failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func TestPmapUint8(t *testing.T) {
	// Test : square the list
	expectedSquareList := []uint8{1, 4, 9}
	squareList := PmapUint8(squareUint8, []uint8{1, 2, 3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("PmapUint8 failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []uint8{6, 7, 8}
	partialAddInt := func(num uint8) uint8 {
		return addUint8(5, num)
	}
	sumList := PmapUint8(partialAddInt, []uint8{1, 2, 3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("PmapUint8 failed.expected=%v, actual=%v", expectedSumList, sumList)
	}

	if len(PmapUint8(nil, nil)) > 0 {
		t.Errorf("PmapUint8 failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	if len(PmapUint8(nil, []uint8{})) > 0 {
		t.Errorf("PmapUint8 failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func TestPmapFloat64(t *testing.T) {
	// Test : square the list
	expectedSquareList := []float64{1.2100000000000002, 4.840000000000001, 10.889999999999999}
	squareList := PmapFloat64(squareFloat64, []float64{1.1, 2.2, 3.3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("PmapFloat64 failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []float64{6.1, 7.2, 8.3}
	partialAddFloat64 := func(num float64) float64 {
		return addFloat64(5, num)
	}
	sumList := PmapFloat64(partialAddFloat64, []float64{1.1, 2.2, 3.3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("PmapFloat64 failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}
	if len(PmapFloat64(nil, nil)) > 0 {
		t.Errorf("PmapFloat64 failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}
	if len(PmapFloat64(nil, []float64{})) > 0 {
		t.Errorf("PmapFloat64 failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func TestPmapFloat32(t *testing.T) {
	// Test : square the list
	expectedSquareList := []float32{1.2100000000000002, 4.840000000000001, 10.889999}
	squareList := PmapFloat32(squareFloat32, []float32{1.1, 2.2, 3.3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("PmapFloat32 failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []float32{6.1, 7.2, 8.3}
	partialAddFloat64 := func(num float32) float32 {
		return addFloat32(5, num)
	}
	sumList := PmapFloat32(partialAddFloat64, []float32{1.1, 2.2, 3.3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("PmapFloat32 failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}
	if len(PmapFloat32(nil, nil)) > 0 {
		t.Errorf("PmapFloat32 failed")
	}
	if len(PmapFloat32(nil, []float32{})) > 0 {
		t.Errorf("PmapFloat32 failed")
	}
}

func TestPmapStr(t *testing.T) {
	// Test : change string to upper case in the list
	expectedSquareList := []string{"GOVINDA", "GOPAL", "SHYAM"}
	strList := PmapStr(strings.ToUpper, []string{"govinda", "gopal", "shyam"})

	if strList[0] != expectedSquareList[0] || strList[1] != expectedSquareList[1] || strList[2] != expectedSquareList[2] {
		t.Errorf("TestMapStr. actual_list=%v, expected_list=%v", strList, expectedSquareList)
	}

	// Test : prepend each string in the list
	expectedSquareList = []string{"Name: Govinda", "Name: Gopal", "Name: Shyam"}

	partialPrependStr := func(str string) string { return prependStr(str, "Name:") }
	strList = PmapStr(partialPrependStr, []string{"Govinda", "Gopal", "Shyam"})

	if strList[0] != expectedSquareList[0] || strList[1] != expectedSquareList[1] || strList[2] != expectedSquareList[2] {
		t.Errorf("TestMapStr failed. actual_list=%v, expected_list=%v", strList, expectedSquareList)
	}
	if len(PmapStr(nil, nil)) > 0 {
		t.Errorf("MapStr failed.")
	}
	if len(PmapStr(nil, []string{})) > 0 {
		t.Errorf("MapStr failed.")
	}
}
