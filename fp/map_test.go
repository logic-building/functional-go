package fp

import (
	"strconv"
	"strings"
	"testing"
)

const size = 1000
const iterations = 1000 * 1000

func TestMapInt(t *testing.T) {
	// Test : square the list
	expectedSquareList := []int{1, 4, 9}
	squareList := MapInt(squareInt, []int{1, 2, 3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("MapInt failed")
	}

	// Test: add 5 to each item in the list
	expectedSumList := []int{6, 7, 8}
	partialAddInt64 := func(num int) int {
		return addInt(5, num)
	}
	sumList := MapInt(partialAddInt64, []int{1, 2, 3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("MapInt failed")
	}

	if len(MapInt(nil, nil)) > 0 {
		t.Errorf("MapInt failed.")
	}
}

func squareInt(num int) int {
	return num * num
}

func addInt(num1, num2 int) int {
	return num1 + num2
}

func BenchmarkMapInt64_PassedMethod_1_Arg(b *testing.B) {
	b.N = iterations
	list := make([]int64, size)
	var i int64
	for i = 0; i < size; i++ {
		list[i] = i
	}
	for i := 0; i < b.N; i++ {
		MapInt64(multiply2, list)
	}
}

func BenchmarkMapInt64_PassedMethod_2_Arg(b *testing.B) {
	b.N = iterations
	list := make([]int64, size)
	var i int64
	for i = 0; i < size; i++ {
		list[i] = i
	}
	partialMultiplyBy := func(num int64) int64 { return multiplyNumBy(num, 2) }
	for i := 0; i < b.N; i++ {
		MapInt64(partialMultiplyBy, list)
	}
}

func multiply2(num int64) int64 {
	return num * 2
}

func multiplyNumBy(num, multiplyBy int64) int64 {
	return num * multiplyBy
}

func TestMapInt64(t *testing.T) {
	// Test : square the list
	expectedSquareList := []int64{1, 4, 9}
	squareList := MapInt64(squareInt64, []int64{1, 2, 3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("MapInt64 failed")
	}

	// Test: add 5 to each item in the list
	expectedSumList := []int64{6, 7, 8}
	partialAddInt64 := func(num int64) int64 {
		return addInt64(5, num)
	}
	sumList := MapInt64(partialAddInt64, []int64{1, 2, 3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("MapInt64 failed")
	}

	if len(MapInt64(nil, nil)) > 0 {
		t.Errorf("MapInt64 failed.")
	}
}

func squareInt64(num int64) int64 {
	return num * num
}

func addInt64(num1, num2 int64) int64 {
	return num1 + num2
}

func TestMapInt32(t *testing.T) {
	// Test : square the list
	expectedSquareList := []int32{1, 4, 9}
	squareList := MapInt32(squareInt32, []int32{1, 2, 3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("MapInt32 failed")
	}

	// Test: add 5 to each item in the list
	expectedSumList := []int32{6, 7, 8}
	partialAddInt32 := func(num int32) int32 {
		return addInt32(5, num)
	}
	sumList := MapInt32(partialAddInt32, []int32{1, 2, 3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("MapInt32 failed")
	}
	if len(MapInt32(nil, nil)) > 0 {
		t.Errorf("MapInt32 failed.")
	}
}

func squareInt32(num int32) int32 {
	return num * num
}

func addInt32(num1, num2 int32) int32 {
	return num1 + num2
}

func TestMapInt16(t *testing.T) {
	// Test : square the list
	expectedSquareList := []int16{1, 4, 9}
	squareList := MapInt16(squareInt16, []int16{1, 2, 3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("MapInt16 failed")
	}

	// Test: add 5 to each item in the list
	expectedSumList := []int16{6, 7, 8}
	partialAddInt16 := func(num int16) int16 {
		return addInt16(5, num)
	}
	sumList := MapInt16(partialAddInt16, []int16{1, 2, 3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("MapInt16 failed")
	}
	if len(MapInt16(nil, nil)) > 0 {
		t.Errorf("MapInt16 failed.")
	}
}

func squareInt16(num int16) int16 {
	return num * num
}

func addInt16(num1, num2 int16) int16 {
	return num1 + num2
}

func TestMapInt8(t *testing.T) {
	// Test : square the list
	expectedSquareList := []int8{1, 4, 9}
	squareList := MapInt8(squareInt8, []int8{1, 2, 3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("MapInt8 failed")
	}

	// Test: add 5 to each item in the list
	expectedSumList := []int8{6, 7, 8}
	partialAddInt8 := func(num int8) int8 {
		return addInt8(5, num)
	}
	sumList := MapInt8(partialAddInt8, []int8{1, 2, 3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("MapInt8 failed")
	}
	if len(MapInt8(nil, nil)) > 0 {
		t.Errorf("MapInt8 failed.")
	}
}

func squareInt8(num int8) int8 {
	return num * num
}

func addInt8(num1, num2 int8) int8 {
	return num1 + num2
}

func TestMapUint64(t *testing.T) {
	// Test : square the list
	expectedSquareList := []uint64{1, 4, 9}
	squareList := MapUint64(squareUint64, []uint64{1, 2, 3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("MapUint64 failed")
	}

	// Test: add 5 to each item in the list
	expectedSumList := []uint64{6, 7, 8}
	partialAddUint64 := func(num uint64) uint64 {
		return addUint64(5, num)
	}
	sumList := MapUint64(partialAddUint64, []uint64{1, 2, 3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("MapUInt64 failed")
	}
	if len(MapUint64(nil, nil)) > 0 {
		t.Errorf("MapUint32 failed.")
	}
}

func squareUint64(num uint64) uint64 {
	return num * num
}

func addUint64(num1, num2 uint64) uint64 {
	return num1 + num2
}

func TestMapUint32(t *testing.T) {
	// Test : square the list
	expectedSquareList := []uint32{1, 4, 9}
	squareList := MapUint32(squareUint32, []uint32{1, 2, 3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("MapUint32 failed")
	}

	// Test: add 5 to each item in the list
	expectedSumList := []uint32{6, 7, 8}
	partialAddUint32 := func(num uint32) uint32 {
		return addUint32(5, num)
	}
	sumList := MapUint32(partialAddUint32, []uint32{1, 2, 3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("MapUInt32 failed")
	}
	if len(MapUint32(nil, nil)) > 0 {
		t.Errorf("MapUint32 failed.")
	}
}

func squareUint32(num uint32) uint32 {
	return num * num
}

func addUint32(num1, num2 uint32) uint32 {
	return num1 + num2
}

func TestMapUint16(t *testing.T) {
	// Test : square the list
	expectedSquareList := []uint16{1, 4, 9}
	squareList := MapUint16(squareUint16, []uint16{1, 2, 3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("MapUint16 failed")
	}

	// Test: add 5 to each item in the list
	expectedSumList := []uint16{6, 7, 8}
	partialAddUint16 := func(num uint16) uint16 {
		return addUint16(5, num)
	}
	sumList := MapUint16(partialAddUint16, []uint16{1, 2, 3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("MapUInt16 failed")
	}
	if len(MapUint16(nil, nil)) > 0 {
		t.Errorf("MapUint16 failed.")
	}
}

func squareUint16(num uint16) uint16 {
	return num * num
}

func addUint16(num1, num2 uint16) uint16 {
	return num1 + num2
}

func TestMapUint8(t *testing.T) {
	// Test : square the list
	expectedSquareList := []uint8{1, 4, 9}
	squareList := MapUint8(squareUint8, []uint8{1, 2, 3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("MapUint8 failed")
	}

	// Test: add 5 to each item in the list
	expectedSumList := []uint8{6, 7, 8}
	partialAddUint8 := func(num uint8) uint8 {
		return addUint8(5, num)
	}
	sumList := MapUint8(partialAddUint8, []uint8{1, 2, 3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("MapUInt8 failed")
	}
	if len(MapUint8(nil, nil)) > 0 {
		t.Errorf("MapUint8 failed.")
	}
}

func squareUint8(num uint8) uint8 {
	return num * num
}

func addUint8(num1, num2 uint8) uint8 {
	return num1 + num2
}

func TestMapUint(t *testing.T) {
	// Test : square the list
	expectedSquareList := []uint{1, 4, 9}
	squareList := MapUint(squareUint, []uint{1, 2, 3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("MapUint failed")
	}

	// Test: add 5 to each item in the list
	expectedSumList := []uint{6, 7, 8}
	partialAddUint := func(num uint) uint {
		return addUint(5, num)
	}
	sumList := MapUint(partialAddUint, []uint{1, 2, 3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("MapUInt failed")
	}
	if len(MapUint(nil, nil)) > 0 {
		t.Errorf("MapUint failed.")
	}
}

func squareUint(num uint) uint {
	return num * num
}

func addUint(num1, num2 uint) uint {
	return num1 + num2
}

func TestMapFloat64(t *testing.T) {
	// Test : square the list
	expectedSquareList := []float64{1.2100000000000002, 4.840000000000001, 10.889999999999999}
	squareList := MapFloat64(squareFloat64, []float64{1.1, 2.2, 3.3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("MapFloat64 failed")
	}

	// Test: add 5 to each item in the list
	expectedSumList := []float64{6.1, 7.2, 8.3}
	partialAddFloat64 := func(num float64) float64 {
		return addFloat64(5, num)
	}
	sumList := MapFloat64(partialAddFloat64, []float64{1.1, 2.2, 3.3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("MapFloat64 failed")
	}
	if len(MapFloat64(nil, nil)) > 0 {
		t.Errorf("MapFloat64 failed.")
	}
}

func squareFloat64(num float64) float64 {
	return num * num
}

func addFloat64(num1, num2 float64) float64 {
	return num1 + num2
}

func TestMapFloat32(t *testing.T) {
	// Test : square the list
	expectedSquareList := []float32{1.2100000000000002, 4.840000000000001, 10.889999}
	squareList := MapFloat32(squareFloat32, []float32{1.1, 2.2, 3.3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("MapFloat32 failed. actual_list=%v, expected_list=%v", squareList, expectedSquareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []float32{6.1, 7.2, 8.3}
	partialAddFloat32 := func(num float32) float32 {
		return addFloat32(5, num)
	}
	sumList := MapFloat32(partialAddFloat32, []float32{1.1, 2.2, 3.3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("MapFloat32 failed")
	}
	if len(MapFloat32(nil, nil)) > 0 {
		t.Errorf("MapFloat32 failed.")
	}
}

func squareFloat32(num float32) float32 {
	return num * num
}

func addFloat32(num1, num2 float32) float32 {
	return num1 + num2
}

func TestMapStr(t *testing.T) {
	// Test : change string to upper case in the list
	expectedSquareList := []string{"GOVINDA", "GOPAL", "SHYAM"}
	strList := MapStr(strings.ToUpper, []string{"govinda", "gopal", "shyam"})

	if strList[0] != expectedSquareList[0] || strList[1] != expectedSquareList[1] || strList[2] != expectedSquareList[2] {
		t.Errorf("TestMapStr. actual_list=%v, expected_list=%v", strList, expectedSquareList)
	}

	// Test : prepend each string in the list
	expectedSquareList = []string{"Name: Govinda", "Name: Gopal", "Name: Shyam"}

	partialPrependStr := func(str string) string { return prependStr(str, "Name:") }
	strList = MapStr(partialPrependStr, []string{"Govinda", "Gopal", "Shyam"})

	if strList[0] != expectedSquareList[0] || strList[1] != expectedSquareList[1] || strList[2] != expectedSquareList[2] {
		t.Errorf("TestMapStr failed. actual_list=%v, expected_list=%v", strList, expectedSquareList)
	}
	if len(MapStr(nil, nil)) > 0 {
		t.Errorf("MapStr failed.")
	}
}

func BenchmarkMapStr(b *testing.B) {
	b.N = iterations
	list := make([]string, size)
	for i := 0; i < size; i++ {
		list[i] = strconv.Itoa(i)
	}

	partialPrependStr := func(str string) string { return prependStr(str, "Name:") }
	for i := 0; i < b.N; i++ {
		MapStr(partialPrependStr, list)
	}
}

func prependStr(str, sbuString string) string {
	return sbuString + " " + str
}
