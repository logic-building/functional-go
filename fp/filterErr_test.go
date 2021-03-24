package fp

import (
	"errors"
	"testing"
)

func TestFilterIntErr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v0 int = 0

	// Test : even number in the list
	expectedFilteredList := []int{v2, v4}
	filteredList, _ := FilterIntErr(isEvenIntErr, []int{v1, v2, v3, v4})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilterErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	r, _ := FilterIntErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterIntErr failed.")
	}

	_, err := FilterIntErr(isEvenIntErr, []int{v0})
	if err == nil {
		t.Errorf("FilterIntPtrErr failed.")
	}
}

func isEvenIntErr(num int) (bool, error) {
	if num == 0 {
		return false, errors.New("Zero is not allowed")
	}
	return num%2 == 0, nil
}

func TestFilterInt64Err(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v0 int64 = 0

	// Test : even number in the list
	expectedFilteredList := []int64{v2, v4}
	filteredList, _ := FilterInt64Err(isEvenInt64Err, []int64{v1, v2, v3, v4})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilterErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	r, _ := FilterInt64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterInt64Err failed.")
	}

	_, err := FilterInt64Err(isEvenInt64Err, []int64{v0})
	if err == nil {
		t.Errorf("FilterInt64PtrErr failed.")
	}
}

func isEvenInt64Err(num int64) (bool, error) {
	if num == 0 {
		return false, errors.New("Zero is not allowed")
	}
	return num%2 == 0, nil
}

func TestFilterInt32Err(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v0 int32 = 0

	// Test : even number in the list
	expectedFilteredList := []int32{v2, v4}
	filteredList, _ := FilterInt32Err(isEvenInt32Err, []int32{v1, v2, v3, v4})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilterErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	r, _ := FilterInt32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterInt32Err failed.")
	}

	_, err := FilterInt32Err(isEvenInt32Err, []int32{v0})
	if err == nil {
		t.Errorf("FilterInt32PtrErr failed.")
	}
}

func isEvenInt32Err(num int32) (bool, error) {
	if num == 0 {
		return false, errors.New("Zero is not allowed")
	}
	return num%2 == 0, nil
}

func TestFilterInt16Err(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v0 int16 = 0

	// Test : even number in the list
	expectedFilteredList := []int16{v2, v4}
	filteredList, _ := FilterInt16Err(isEvenInt16Err, []int16{v1, v2, v3, v4})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilterErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	r, _ := FilterInt16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterInt16Err failed.")
	}

	_, err := FilterInt16Err(isEvenInt16Err, []int16{v0})
	if err == nil {
		t.Errorf("FilterInt16PtrErr failed.")
	}
}

func isEvenInt16Err(num int16) (bool, error) {
	if num == 0 {
		return false, errors.New("Zero is not allowed")
	}
	return num%2 == 0, nil
}

func TestFilterInt8Err(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v0 int8 = 0

	// Test : even number in the list
	expectedFilteredList := []int8{v2, v4}
	filteredList, _ := FilterInt8Err(isEvenInt8Err, []int8{v1, v2, v3, v4})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilterErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	r, _ := FilterInt8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterInt8Err failed.")
	}

	_, err := FilterInt8Err(isEvenInt8Err, []int8{v0})
	if err == nil {
		t.Errorf("FilterInt8PtrErr failed.")
	}
}

func isEvenInt8Err(num int8) (bool, error) {
	if num == 0 {
		return false, errors.New("Zero is not allowed")
	}
	return num%2 == 0, nil
}

func TestFilterUintErr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v0 uint = 0

	// Test : even number in the list
	expectedFilteredList := []uint{v2, v4}
	filteredList, _ := FilterUintErr(isEvenUintErr, []uint{v1, v2, v3, v4})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilterErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	r, _ := FilterUintErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterUintErr failed.")
	}

	_, err := FilterUintErr(isEvenUintErr, []uint{v0})
	if err == nil {
		t.Errorf("FilterUintPtrErr failed.")
	}
}

func isEvenUintErr(num uint) (bool, error) {
	if num == 0 {
		return false, errors.New("Zero is not allowed")
	}
	return num%2 == 0, nil
}

func TestFilterUint64Err(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v0 uint64 = 0

	// Test : even number in the list
	expectedFilteredList := []uint64{v2, v4}
	filteredList, _ := FilterUint64Err(isEvenUint64Err, []uint64{v1, v2, v3, v4})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilterErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	r, _ := FilterUint64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterUint64Err failed.")
	}

	_, err := FilterUint64Err(isEvenUint64Err, []uint64{v0})
	if err == nil {
		t.Errorf("FilterUint64PtrErr failed.")
	}
}

func isEvenUint64Err(num uint64) (bool, error) {
	if num == 0 {
		return false, errors.New("Zero is not allowed")
	}
	return num%2 == 0, nil
}

func TestFilterUint32Err(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v0 uint32 = 0

	// Test : even number in the list
	expectedFilteredList := []uint32{v2, v4}
	filteredList, _ := FilterUint32Err(isEvenUint32Err, []uint32{v1, v2, v3, v4})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilterErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	r, _ := FilterUint32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterUint32Err failed.")
	}

	_, err := FilterUint32Err(isEvenUint32Err, []uint32{v0})
	if err == nil {
		t.Errorf("FilterUint32PtrErr failed.")
	}
}

func isEvenUint32Err(num uint32) (bool, error) {
	if num == 0 {
		return false, errors.New("Zero is not allowed")
	}
	return num%2 == 0, nil
}

func TestFilterUint16Err(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v0 uint16 = 0

	// Test : even number in the list
	expectedFilteredList := []uint16{v2, v4}
	filteredList, _ := FilterUint16Err(isEvenUint16Err, []uint16{v1, v2, v3, v4})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilterErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	r, _ := FilterUint16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterUint16Err failed.")
	}

	_, err := FilterUint16Err(isEvenUint16Err, []uint16{v0})
	if err == nil {
		t.Errorf("FilterUint16PtrErr failed.")
	}
}

func isEvenUint16Err(num uint16) (bool, error) {
	if num == 0 {
		return false, errors.New("Zero is not allowed")
	}
	return num%2 == 0, nil
}

func TestFilterUint8Err(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v0 uint8 = 0

	// Test : even number in the list
	expectedFilteredList := []uint8{v2, v4}
	filteredList, _ := FilterUint8Err(isEvenUint8Err, []uint8{v1, v2, v3, v4})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilterErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	r, _ := FilterUint8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterUint8Err failed.")
	}

	_, err := FilterUint8Err(isEvenUint8Err, []uint8{v0})
	if err == nil {
		t.Errorf("FilterUint8PtrErr failed.")
	}
}

func isEvenUint8Err(num uint8) (bool, error) {
	if num == 0 {
		return false, errors.New("Zero is not allowed")
	}
	return num%2 == 0, nil
}

func TestFilterStrErr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v0 string = "0"

	// Test : even number in the list
	expectedFilteredList := []string{v2, v4}
	filteredList, _ := FilterStrErr(isEvenStrErr, []string{v1, v2, v3, v4})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilterErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	r, _ := FilterStrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterStrErr failed.")
	}

	_, err := FilterStrErr(isEvenStrErr, []string{v0})
	if err == nil {
		t.Errorf("FilterStrPtrErr failed.")
	}
}

func isEvenStrErr(num string) (bool, error) {
	if num == "0" {
		return false, errors.New("Zero is not allowed")
	} else if num == "2" || num == "4" {
		return true, nil
	}
	return false, nil

}

func TestFilterBoolErr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	expectedSumList := []bool{vt}

	newList, _ := FilterBoolErr(trueBoolErr, []bool{vt})
	if newList[0] != expectedSumList[0] {
		t.Errorf("FilterBoolErr failed")
	}

	r, _ := FilterBoolErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterBoolErr failed.")
	}

	_, err := FilterBoolErr(trueBoolErr, []bool{vf})
	if err == nil {
		t.Errorf("FilterBoolErr failed.")
	}
}

func trueBoolErr(num1 bool) (bool, error) {
	if num1 == false {
		return false, errors.New("False is not allowed")
	}
	return true, nil
}

func TestFilterFloat32Err(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v0 float32 = 0

	// Test : even number in the list
	expectedFilteredList := []float32{v2, v4}
	filteredList, _ := FilterFloat32Err(isEvenFloat32Err, []float32{v1, v2, v3, v4})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilterErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	r, _ := FilterFloat32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterFloat32Err failed.")
	}

	_, err := FilterFloat32Err(isEvenFloat32Err, []float32{v0})
	if err == nil {
		t.Errorf("FilterFloat32PtrErr failed.")
	}
}

func isEvenFloat32Err(num float32) (bool, error) {
	if num == 0 {
		return false, errors.New("Zero is not allowed")
	}
	return int(num)%2 == 0, nil
}

func TestFilterFloat64Err(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v0 float64 = 0

	// Test : even number in the list
	expectedFilteredList := []float64{v2, v4}
	filteredList, _ := FilterFloat64Err(isEvenFloat64Err, []float64{v1, v2, v3, v4})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilterErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	r, _ := FilterFloat64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterFloat64Err failed.")
	}

	_, err := FilterFloat64Err(isEvenFloat64Err, []float64{v0})
	if err == nil {
		t.Errorf("FilterFloat64PtrErr failed.")
	}
}

func isEvenFloat64Err(num float64) (bool, error) {
	if num == 0 {
		return false, errors.New("Zero is not allowed")
	}
	return int(num)%2 == 0, nil
}
