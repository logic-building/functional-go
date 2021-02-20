package fp

import (
	"errors"
	"testing"
)

func TestFilterIntPtrErr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v0 int = 0

	// Test : even number in the list
	expectedFilteredList := []*int{&v2, &v4}
	filteredList, _ := FilterIntPtrErr(isEvenIntPtrErr, []*int{&v1, &v2, &v3, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("MapFilterPtrErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	r, _ := FilterIntPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterIntPtrErr failed.")
	}

	_, err := FilterIntPtrErr(isEvenIntPtrErr, []*int{&v0})
	if err == nil {
		t.Errorf("FilterIntPtrErr failed.")
	}
}

func isEvenIntPtrErr(num *int) (bool, error) {
	if *num == 0 {
		return false, errors.New("Zero is not allowed")
	}
	return *num%2 == 0, nil
}


func TestFilterInt64PtrErr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v0 int64 = 0

	// Test : even number in the list
	expectedFilteredList := []*int64{&v2, &v4}
	filteredList, _ := FilterInt64PtrErr(isEvenInt64PtrErr, []*int64{&v1, &v2, &v3, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("MapFilterPtrErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	r, _ := FilterInt64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterInt64PtrErr failed.")
	}

	_, err := FilterInt64PtrErr(isEvenInt64PtrErr, []*int64{&v0})
	if err == nil {
		t.Errorf("FilterInt64PtrErr failed.")
	}
}

func isEvenInt64PtrErr(num *int64) (bool, error) {
	if *num == 0 {
		return false, errors.New("Zero is not allowed")
	}
	return *num%2 == 0, nil
}


func TestFilterInt32PtrErr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v0 int32 = 0

	// Test : even number in the list
	expectedFilteredList := []*int32{&v2, &v4}
	filteredList, _ := FilterInt32PtrErr(isEvenInt32PtrErr, []*int32{&v1, &v2, &v3, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("MapFilterPtrErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	r, _ := FilterInt32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterInt32PtrErr failed.")
	}

	_, err := FilterInt32PtrErr(isEvenInt32PtrErr, []*int32{&v0})
	if err == nil {
		t.Errorf("FilterInt32PtrErr failed.")
	}
}

func isEvenInt32PtrErr(num *int32) (bool, error) {
	if *num == 0 {
		return false, errors.New("Zero is not allowed")
	}
	return *num%2 == 0, nil
}


func TestFilterInt16PtrErr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v0 int16 = 0

	// Test : even number in the list
	expectedFilteredList := []*int16{&v2, &v4}
	filteredList, _ := FilterInt16PtrErr(isEvenInt16PtrErr, []*int16{&v1, &v2, &v3, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("MapFilterPtrErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	r, _ := FilterInt16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterInt16PtrErr failed.")
	}

	_, err := FilterInt16PtrErr(isEvenInt16PtrErr, []*int16{&v0})
	if err == nil {
		t.Errorf("FilterInt16PtrErr failed.")
	}
}

func isEvenInt16PtrErr(num *int16) (bool, error) {
	if *num == 0 {
		return false, errors.New("Zero is not allowed")
	}
	return *num%2 == 0, nil
}


func TestFilterInt8PtrErr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v0 int8 = 0

	// Test : even number in the list
	expectedFilteredList := []*int8{&v2, &v4}
	filteredList, _ := FilterInt8PtrErr(isEvenInt8PtrErr, []*int8{&v1, &v2, &v3, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("MapFilterPtrErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	r, _ := FilterInt8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterInt8PtrErr failed.")
	}

	_, err := FilterInt8PtrErr(isEvenInt8PtrErr, []*int8{&v0})
	if err == nil {
		t.Errorf("FilterInt8PtrErr failed.")
	}
}

func isEvenInt8PtrErr(num *int8) (bool, error) {
	if *num == 0 {
		return false, errors.New("Zero is not allowed")
	}
	return *num%2 == 0, nil
}


func TestFilterUintPtrErr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v0 uint = 0

	// Test : even number in the list
	expectedFilteredList := []*uint{&v2, &v4}
	filteredList, _ := FilterUintPtrErr(isEvenUintPtrErr, []*uint{&v1, &v2, &v3, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("MapFilterPtrErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	r, _ := FilterUintPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterUintPtrErr failed.")
	}

	_, err := FilterUintPtrErr(isEvenUintPtrErr, []*uint{&v0})
	if err == nil {
		t.Errorf("FilterUintPtrErr failed.")
	}
}

func isEvenUintPtrErr(num *uint) (bool, error) {
	if *num == 0 {
		return false, errors.New("Zero is not allowed")
	}
	return *num%2 == 0, nil
}


func TestFilterUint64PtrErr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v0 uint64 = 0

	// Test : even number in the list
	expectedFilteredList := []*uint64{&v2, &v4}
	filteredList, _ := FilterUint64PtrErr(isEvenUint64PtrErr, []*uint64{&v1, &v2, &v3, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("MapFilterPtrErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	r, _ := FilterUint64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterUint64PtrErr failed.")
	}

	_, err := FilterUint64PtrErr(isEvenUint64PtrErr, []*uint64{&v0})
	if err == nil {
		t.Errorf("FilterUint64PtrErr failed.")
	}
}

func isEvenUint64PtrErr(num *uint64) (bool, error) {
	if *num == 0 {
		return false, errors.New("Zero is not allowed")
	}
	return *num%2 == 0, nil
}


func TestFilterUint32PtrErr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v0 uint32 = 0

	// Test : even number in the list
	expectedFilteredList := []*uint32{&v2, &v4}
	filteredList, _ := FilterUint32PtrErr(isEvenUint32PtrErr, []*uint32{&v1, &v2, &v3, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("MapFilterPtrErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	r, _ := FilterUint32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterUint32PtrErr failed.")
	}

	_, err := FilterUint32PtrErr(isEvenUint32PtrErr, []*uint32{&v0})
	if err == nil {
		t.Errorf("FilterUint32PtrErr failed.")
	}
}

func isEvenUint32PtrErr(num *uint32) (bool, error) {
	if *num == 0 {
		return false, errors.New("Zero is not allowed")
	}
	return *num%2 == 0, nil
}


func TestFilterUint16PtrErr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v0 uint16 = 0

	// Test : even number in the list
	expectedFilteredList := []*uint16{&v2, &v4}
	filteredList, _ := FilterUint16PtrErr(isEvenUint16PtrErr, []*uint16{&v1, &v2, &v3, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("MapFilterPtrErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	r, _ := FilterUint16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterUint16PtrErr failed.")
	}

	_, err := FilterUint16PtrErr(isEvenUint16PtrErr, []*uint16{&v0})
	if err == nil {
		t.Errorf("FilterUint16PtrErr failed.")
	}
}

func isEvenUint16PtrErr(num *uint16) (bool, error) {
	if *num == 0 {
		return false, errors.New("Zero is not allowed")
	}
	return *num%2 == 0, nil
}


func TestFilterUint8PtrErr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v0 uint8 = 0

	// Test : even number in the list
	expectedFilteredList := []*uint8{&v2, &v4}
	filteredList, _ := FilterUint8PtrErr(isEvenUint8PtrErr, []*uint8{&v1, &v2, &v3, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("MapFilterPtrErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	r, _ := FilterUint8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterUint8PtrErr failed.")
	}

	_, err := FilterUint8PtrErr(isEvenUint8PtrErr, []*uint8{&v0})
	if err == nil {
		t.Errorf("FilterUint8PtrErr failed.")
	}
}

func isEvenUint8PtrErr(num *uint8) (bool, error) {
	if *num == 0 {
		return false, errors.New("Zero is not allowed")
	}
	return *num%2 == 0, nil
}


func TestFilterStrPtrErr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v0 string = "0"

	// Test : even number in the list
	expectedFilteredList := []*string{&v2, &v4}
	filteredList, _ := FilterStrPtrErr(isEvenStrPtrErr, []*string{&v1, &v2, &v3, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("MapFilterPtrErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	r, _ := FilterStrPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterStrPtrErr failed.")
	}

	_, err := FilterStrPtrErr(isEvenStrPtrErr, []*string{&v0})
	if err == nil {
		t.Errorf("FilterStrPtrErr failed.")
	}
}

func isEvenStrPtrErr(num *string) (bool, error) {
	if *num == "0" {
		return false, errors.New("Zero is not allowed")
	} else if *num == "2" || *num == "4" {
		return true, nil
	}
	return false, nil
	
}


func TestFilterBoolPtrErr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	expectedSumList := []*bool{&vt}
	
	newList, _ := FilterBoolPtrErr(trueBoolPtrErr, []*bool{&vt})
	if *newList[0] != *expectedSumList[0]  {
		t.Errorf("FilterBoolPtrErr failed")
	}

	r, _ := FilterBoolPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterBoolPtrErr failed.")
	}

	_, err := FilterBoolPtrErr(trueBoolPtrErr, []*bool{&vf})
	if err == nil {
		t.Errorf("FilterBoolPtrErr failed.")
	}
}

func trueBoolPtrErr(num1 *bool) (bool, error) {
	if *num1 == false {
		return false, errors.New("False is not allowed")
	}
	return true, nil
}


func TestFilterFloat32PtrErr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v0 float32 = 0

	// Test : even number in the list
	expectedFilteredList := []*float32{&v2, &v4}
	filteredList, _ := FilterFloat32PtrErr(isEvenFloat32PtrErr, []*float32{&v1, &v2, &v3, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("MapFilterPtrErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	r, _ := FilterFloat32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterFloat32PtrErr failed.")
	}

	_, err := FilterFloat32PtrErr(isEvenFloat32PtrErr, []*float32{&v0})
	if err == nil {
		t.Errorf("FilterFloat32PtrErr failed.")
	}
}

func isEvenFloat32PtrErr(num *float32) (bool, error) {
	if *num == 0 {
		return false, errors.New("Zero is not allowed")
	}
	return int(*num)%2 == 0, nil
}


func TestFilterFloat64PtrErr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v0 float64 = 0

	// Test : even number in the list
	expectedFilteredList := []*float64{&v2, &v4}
	filteredList, _ := FilterFloat64PtrErr(isEvenFloat64PtrErr, []*float64{&v1, &v2, &v3, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("MapFilterPtrErr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	r, _ := FilterFloat64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("FilterFloat64PtrErr failed.")
	}

	_, err := FilterFloat64PtrErr(isEvenFloat64PtrErr, []*float64{&v0})
	if err == nil {
		t.Errorf("FilterFloat64PtrErr failed.")
	}
}

func isEvenFloat64PtrErr(num *float64) (bool, error) {
	if *num == 0 {
		return false, errors.New("Zero is not allowed")
	}
	return int(*num)%2 == 0, nil
}

