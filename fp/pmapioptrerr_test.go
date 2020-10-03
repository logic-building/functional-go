package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestPmapIntInt64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3

	expectedList := []*int64{&vo2, &vo3}
	newList, _ := PMapIntInt64PtrErr(plusOneIntInt64PtrErr, []*int{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapIntInt64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapIntInt64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapIntInt64PtrErr failed")
	}

	r, _ = PMapIntInt64PtrErr(nil, []*int{})
	if len(r) > 0 {
		t.Errorf("PMapIntInt64PtrErr failed")
	}
	_, err := PMapIntInt64PtrErr(plusOneIntInt64PtrErr, []*int{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapIntInt64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapIntInt64PtrErr(plusOneIntInt64PtrErr, []*int{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntInt64PtrErr failed")
	}

	_, err = PMapIntInt64PtrErr(plusOneIntInt64PtrErr, []*int{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapIntInt64PtrErr failed")
	}

	_, err = PMapIntInt64PtrErr(plusOneIntInt64PtrErr, []*int{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntInt64PtrErr failed")
	}
	_, err = PMapIntInt64PtrErr(plusOneIntInt64PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntInt64PtrErr failed")
	}

	expectedList = []*int64{&vo2, &vo3}
	newList, _ = PMapIntInt64PtrErr(plusOneIntInt64PtrErr, []*int{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapIntInt64Ptr failed")
	}

	_, err = PMapIntInt64PtrErr(plusOneIntInt64PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntInt64PtrErr failed")
	}
	_, err = PMapIntInt64PtrErr(plusOneIntInt64PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntInt64PtrErr failed")
	}

	_, err = PMapIntInt64PtrErr(plusOneIntInt64PtrErr, []*int{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntInt64PtrErr failed")
	}

	_, err = PMapIntInt64PtrErr(plusOneIntInt64PtrErr, []*int{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntInt64PtrErr failed")
	}

	_, err = PMapIntInt64PtrErr(plusOneIntInt64PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntInt64PtrErr failed")
	}
}

func TestPmapIntInt32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3

	expectedList := []*int32{&vo2, &vo3}
	newList, _ := PMapIntInt32PtrErr(plusOneIntInt32PtrErr, []*int{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapIntInt32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapIntInt32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapIntInt32PtrErr failed")
	}

	r, _ = PMapIntInt32PtrErr(nil, []*int{})
	if len(r) > 0 {
		t.Errorf("PMapIntInt32PtrErr failed")
	}
	_, err := PMapIntInt32PtrErr(plusOneIntInt32PtrErr, []*int{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapIntInt32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapIntInt32PtrErr(plusOneIntInt32PtrErr, []*int{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntInt32PtrErr failed")
	}

	_, err = PMapIntInt32PtrErr(plusOneIntInt32PtrErr, []*int{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapIntInt32PtrErr failed")
	}

	_, err = PMapIntInt32PtrErr(plusOneIntInt32PtrErr, []*int{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntInt32PtrErr failed")
	}
	_, err = PMapIntInt32PtrErr(plusOneIntInt32PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntInt32PtrErr failed")
	}

	expectedList = []*int32{&vo2, &vo3}
	newList, _ = PMapIntInt32PtrErr(plusOneIntInt32PtrErr, []*int{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapIntInt32Ptr failed")
	}

	_, err = PMapIntInt32PtrErr(plusOneIntInt32PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntInt32PtrErr failed")
	}
	_, err = PMapIntInt32PtrErr(plusOneIntInt32PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntInt32PtrErr failed")
	}

	_, err = PMapIntInt32PtrErr(plusOneIntInt32PtrErr, []*int{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntInt32PtrErr failed")
	}

	_, err = PMapIntInt32PtrErr(plusOneIntInt32PtrErr, []*int{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntInt32PtrErr failed")
	}

	_, err = PMapIntInt32PtrErr(plusOneIntInt32PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntInt32PtrErr failed")
	}
}

func TestPmapIntInt16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3

	expectedList := []*int16{&vo2, &vo3}
	newList, _ := PMapIntInt16PtrErr(plusOneIntInt16PtrErr, []*int{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapIntInt16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapIntInt16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapIntInt16PtrErr failed")
	}

	r, _ = PMapIntInt16PtrErr(nil, []*int{})
	if len(r) > 0 {
		t.Errorf("PMapIntInt16PtrErr failed")
	}
	_, err := PMapIntInt16PtrErr(plusOneIntInt16PtrErr, []*int{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapIntInt16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapIntInt16PtrErr(plusOneIntInt16PtrErr, []*int{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntInt16PtrErr failed")
	}

	_, err = PMapIntInt16PtrErr(plusOneIntInt16PtrErr, []*int{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapIntInt16PtrErr failed")
	}

	_, err = PMapIntInt16PtrErr(plusOneIntInt16PtrErr, []*int{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntInt16PtrErr failed")
	}
	_, err = PMapIntInt16PtrErr(plusOneIntInt16PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntInt16PtrErr failed")
	}

	expectedList = []*int16{&vo2, &vo3}
	newList, _ = PMapIntInt16PtrErr(plusOneIntInt16PtrErr, []*int{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapIntInt16Ptr failed")
	}

	_, err = PMapIntInt16PtrErr(plusOneIntInt16PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntInt16PtrErr failed")
	}
	_, err = PMapIntInt16PtrErr(plusOneIntInt16PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntInt16PtrErr failed")
	}

	_, err = PMapIntInt16PtrErr(plusOneIntInt16PtrErr, []*int{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntInt16PtrErr failed")
	}

	_, err = PMapIntInt16PtrErr(plusOneIntInt16PtrErr, []*int{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntInt16PtrErr failed")
	}

	_, err = PMapIntInt16PtrErr(plusOneIntInt16PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntInt16PtrErr failed")
	}
}

func TestPmapIntInt8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3

	expectedList := []*int8{&vo2, &vo3}
	newList, _ := PMapIntInt8PtrErr(plusOneIntInt8PtrErr, []*int{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapIntInt8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapIntInt8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapIntInt8PtrErr failed")
	}

	r, _ = PMapIntInt8PtrErr(nil, []*int{})
	if len(r) > 0 {
		t.Errorf("PMapIntInt8PtrErr failed")
	}
	_, err := PMapIntInt8PtrErr(plusOneIntInt8PtrErr, []*int{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapIntInt8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapIntInt8PtrErr(plusOneIntInt8PtrErr, []*int{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntInt8PtrErr failed")
	}

	_, err = PMapIntInt8PtrErr(plusOneIntInt8PtrErr, []*int{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapIntInt8PtrErr failed")
	}

	_, err = PMapIntInt8PtrErr(plusOneIntInt8PtrErr, []*int{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntInt8PtrErr failed")
	}
	_, err = PMapIntInt8PtrErr(plusOneIntInt8PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntInt8PtrErr failed")
	}

	expectedList = []*int8{&vo2, &vo3}
	newList, _ = PMapIntInt8PtrErr(plusOneIntInt8PtrErr, []*int{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapIntInt8Ptr failed")
	}

	_, err = PMapIntInt8PtrErr(plusOneIntInt8PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntInt8PtrErr failed")
	}
	_, err = PMapIntInt8PtrErr(plusOneIntInt8PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntInt8PtrErr failed")
	}

	_, err = PMapIntInt8PtrErr(plusOneIntInt8PtrErr, []*int{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntInt8PtrErr failed")
	}

	_, err = PMapIntInt8PtrErr(plusOneIntInt8PtrErr, []*int{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntInt8PtrErr failed")
	}

	_, err = PMapIntInt8PtrErr(plusOneIntInt8PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntInt8PtrErr failed")
	}
}

func TestPmapIntUintPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3

	expectedList := []*uint{&vo2, &vo3}
	newList, _ := PMapIntUintPtrErr(plusOneIntUintPtrErr, []*int{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapIntUintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapIntUintPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapIntUintPtrErr failed")
	}

	r, _ = PMapIntUintPtrErr(nil, []*int{})
	if len(r) > 0 {
		t.Errorf("PMapIntUintPtrErr failed")
	}
	_, err := PMapIntUintPtrErr(plusOneIntUintPtrErr, []*int{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapIntUintPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapIntUintPtrErr(plusOneIntUintPtrErr, []*int{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntUintPtrErr failed")
	}

	_, err = PMapIntUintPtrErr(plusOneIntUintPtrErr, []*int{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapIntUintPtrErr failed")
	}

	_, err = PMapIntUintPtrErr(plusOneIntUintPtrErr, []*int{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntUintPtrErr failed")
	}
	_, err = PMapIntUintPtrErr(plusOneIntUintPtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntUintPtrErr failed")
	}

	expectedList = []*uint{&vo2, &vo3}
	newList, _ = PMapIntUintPtrErr(plusOneIntUintPtrErr, []*int{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapIntUintPtr failed")
	}

	_, err = PMapIntUintPtrErr(plusOneIntUintPtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntUintPtrErr failed")
	}
	_, err = PMapIntUintPtrErr(plusOneIntUintPtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntUintPtrErr failed")
	}

	_, err = PMapIntUintPtrErr(plusOneIntUintPtrErr, []*int{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntUintPtrErr failed")
	}

	_, err = PMapIntUintPtrErr(plusOneIntUintPtrErr, []*int{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntUintPtrErr failed")
	}

	_, err = PMapIntUintPtrErr(plusOneIntUintPtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntUintPtrErr failed")
	}
}

func TestPmapIntUint64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3

	expectedList := []*uint64{&vo2, &vo3}
	newList, _ := PMapIntUint64PtrErr(plusOneIntUint64PtrErr, []*int{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapIntUint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapIntUint64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapIntUint64PtrErr failed")
	}

	r, _ = PMapIntUint64PtrErr(nil, []*int{})
	if len(r) > 0 {
		t.Errorf("PMapIntUint64PtrErr failed")
	}
	_, err := PMapIntUint64PtrErr(plusOneIntUint64PtrErr, []*int{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapIntUint64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapIntUint64PtrErr(plusOneIntUint64PtrErr, []*int{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntUint64PtrErr failed")
	}

	_, err = PMapIntUint64PtrErr(plusOneIntUint64PtrErr, []*int{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapIntUint64PtrErr failed")
	}

	_, err = PMapIntUint64PtrErr(plusOneIntUint64PtrErr, []*int{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntUint64PtrErr failed")
	}
	_, err = PMapIntUint64PtrErr(plusOneIntUint64PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntUint64PtrErr failed")
	}

	expectedList = []*uint64{&vo2, &vo3}
	newList, _ = PMapIntUint64PtrErr(plusOneIntUint64PtrErr, []*int{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapIntUint64Ptr failed")
	}

	_, err = PMapIntUint64PtrErr(plusOneIntUint64PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntUint64PtrErr failed")
	}
	_, err = PMapIntUint64PtrErr(plusOneIntUint64PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntUint64PtrErr failed")
	}

	_, err = PMapIntUint64PtrErr(plusOneIntUint64PtrErr, []*int{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntUint64PtrErr failed")
	}

	_, err = PMapIntUint64PtrErr(plusOneIntUint64PtrErr, []*int{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntUint64PtrErr failed")
	}

	_, err = PMapIntUint64PtrErr(plusOneIntUint64PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntUint64PtrErr failed")
	}
}

func TestPmapIntUint32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3

	expectedList := []*uint32{&vo2, &vo3}
	newList, _ := PMapIntUint32PtrErr(plusOneIntUint32PtrErr, []*int{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapIntUint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapIntUint32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapIntUint32PtrErr failed")
	}

	r, _ = PMapIntUint32PtrErr(nil, []*int{})
	if len(r) > 0 {
		t.Errorf("PMapIntUint32PtrErr failed")
	}
	_, err := PMapIntUint32PtrErr(plusOneIntUint32PtrErr, []*int{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapIntUint32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapIntUint32PtrErr(plusOneIntUint32PtrErr, []*int{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntUint32PtrErr failed")
	}

	_, err = PMapIntUint32PtrErr(plusOneIntUint32PtrErr, []*int{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapIntUint32PtrErr failed")
	}

	_, err = PMapIntUint32PtrErr(plusOneIntUint32PtrErr, []*int{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntUint32PtrErr failed")
	}
	_, err = PMapIntUint32PtrErr(plusOneIntUint32PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntUint32PtrErr failed")
	}

	expectedList = []*uint32{&vo2, &vo3}
	newList, _ = PMapIntUint32PtrErr(plusOneIntUint32PtrErr, []*int{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapIntUint32Ptr failed")
	}

	_, err = PMapIntUint32PtrErr(plusOneIntUint32PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntUint32PtrErr failed")
	}
	_, err = PMapIntUint32PtrErr(plusOneIntUint32PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntUint32PtrErr failed")
	}

	_, err = PMapIntUint32PtrErr(plusOneIntUint32PtrErr, []*int{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntUint32PtrErr failed")
	}

	_, err = PMapIntUint32PtrErr(plusOneIntUint32PtrErr, []*int{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntUint32PtrErr failed")
	}

	_, err = PMapIntUint32PtrErr(plusOneIntUint32PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntUint32PtrErr failed")
	}
}

func TestPmapIntUint16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3

	expectedList := []*uint16{&vo2, &vo3}
	newList, _ := PMapIntUint16PtrErr(plusOneIntUint16PtrErr, []*int{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapIntUint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapIntUint16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapIntUint16PtrErr failed")
	}

	r, _ = PMapIntUint16PtrErr(nil, []*int{})
	if len(r) > 0 {
		t.Errorf("PMapIntUint16PtrErr failed")
	}
	_, err := PMapIntUint16PtrErr(plusOneIntUint16PtrErr, []*int{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapIntUint16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapIntUint16PtrErr(plusOneIntUint16PtrErr, []*int{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntUint16PtrErr failed")
	}

	_, err = PMapIntUint16PtrErr(plusOneIntUint16PtrErr, []*int{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapIntUint16PtrErr failed")
	}

	_, err = PMapIntUint16PtrErr(plusOneIntUint16PtrErr, []*int{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntUint16PtrErr failed")
	}
	_, err = PMapIntUint16PtrErr(plusOneIntUint16PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntUint16PtrErr failed")
	}

	expectedList = []*uint16{&vo2, &vo3}
	newList, _ = PMapIntUint16PtrErr(plusOneIntUint16PtrErr, []*int{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapIntUint16Ptr failed")
	}

	_, err = PMapIntUint16PtrErr(plusOneIntUint16PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntUint16PtrErr failed")
	}
	_, err = PMapIntUint16PtrErr(plusOneIntUint16PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntUint16PtrErr failed")
	}

	_, err = PMapIntUint16PtrErr(plusOneIntUint16PtrErr, []*int{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntUint16PtrErr failed")
	}

	_, err = PMapIntUint16PtrErr(plusOneIntUint16PtrErr, []*int{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntUint16PtrErr failed")
	}

	_, err = PMapIntUint16PtrErr(plusOneIntUint16PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntUint16PtrErr failed")
	}
}

func TestPmapIntUint8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3

	expectedList := []*uint8{&vo2, &vo3}
	newList, _ := PMapIntUint8PtrErr(plusOneIntUint8PtrErr, []*int{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapIntUint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapIntUint8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapIntUint8PtrErr failed")
	}

	r, _ = PMapIntUint8PtrErr(nil, []*int{})
	if len(r) > 0 {
		t.Errorf("PMapIntUint8PtrErr failed")
	}
	_, err := PMapIntUint8PtrErr(plusOneIntUint8PtrErr, []*int{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapIntUint8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapIntUint8PtrErr(plusOneIntUint8PtrErr, []*int{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntUint8PtrErr failed")
	}

	_, err = PMapIntUint8PtrErr(plusOneIntUint8PtrErr, []*int{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapIntUint8PtrErr failed")
	}

	_, err = PMapIntUint8PtrErr(plusOneIntUint8PtrErr, []*int{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntUint8PtrErr failed")
	}
	_, err = PMapIntUint8PtrErr(plusOneIntUint8PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntUint8PtrErr failed")
	}

	expectedList = []*uint8{&vo2, &vo3}
	newList, _ = PMapIntUint8PtrErr(plusOneIntUint8PtrErr, []*int{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapIntUint8Ptr failed")
	}

	_, err = PMapIntUint8PtrErr(plusOneIntUint8PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntUint8PtrErr failed")
	}
	_, err = PMapIntUint8PtrErr(plusOneIntUint8PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntUint8PtrErr failed")
	}

	_, err = PMapIntUint8PtrErr(plusOneIntUint8PtrErr, []*int{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntUint8PtrErr failed")
	}

	_, err = PMapIntUint8PtrErr(plusOneIntUint8PtrErr, []*int{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntUint8PtrErr failed")
	}

	_, err = PMapIntUint8PtrErr(plusOneIntUint8PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntUint8PtrErr failed")
	}
}

func TestPmapIntStrPtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 int = 10
	var vi0 int

	expectedList := []*string{&vo10}
	newList, _ := PMapIntStrPtrErr(someLogicIntStrPtrErr, []*int{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapIntStrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapIntStrPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapIntStrPtrErr failed")
	}

	r, _ = PMapIntStrPtrErr(nil, []*int{})
	if len(r) > 0 {
		t.Errorf("PMapIntStrPtrErr failed")
	}
	_, err := PMapIntStrPtrErr(someLogicIntStrPtrErr, []*int{&vi10, &vi0})
	if err == nil {
		t.Errorf("PMapIntStrPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapIntStrPtrErr(someLogicIntStrPtrErr, []*int{&vi10, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntStrPtrErr failed")
	}

	_, err = PMapIntStrPtrErr(someLogicIntStrPtrErr, []*int{&vi0, &vi0, &vi10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapIntStrPtrErr failed")
	}

	_, err = PMapIntStrPtrErr(someLogicIntStrPtrErr, []*int{&vi0, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntStrPtrErr failed")
	}
	_, err = PMapIntStrPtrErr(someLogicIntStrPtrErr, []*int{&vi10, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntStrPtrErr failed")
	}

	expectedList = []*string{&vo10, &vo10}
	newList, _ = PMapIntStrPtrErr(someLogicIntStrPtrErr, []*int{&vi10, &vi10}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapIntStrPtr failed")
	}

	_, err = PMapIntStrPtrErr(someLogicIntStrPtrErr, []*int{&vi10, &vi10, &vi0}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntStrPtrErr failed")
	}
	_, err = PMapIntStrPtrErr(someLogicIntStrPtrErr, []*int{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntStrPtrErr failed")
	}

	_, err = PMapIntStrPtrErr(someLogicIntStrPtrErr, []*int{&vi0, &vi0, &vi10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntStrPtrErr failed")
	}

	_, err = PMapIntStrPtrErr(someLogicIntStrPtrErr, []*int{&vi0, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntStrPtrErr failed")
	}

	_, err = PMapIntStrPtrErr(someLogicIntStrPtrErr, []*int{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntStrPtrErr failed")
	}
}

func TestPmapIntBoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 int = 10
	var v0 int
	var v3 int = 3

	expectedList := []*bool{&vt, &vf}
	newList, _ := PMapIntBoolPtrErr(someLogicIntBoolPtrErr, []*int{&v10, &v0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapIntBoolPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapIntBoolPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapIntBoolPtrErr failed")
	}

	r, _ = PMapIntBoolPtrErr(nil, []*int{})
	if len(r) > 0 {
		t.Errorf("PMapIntBoolPtrErr failed")
	}
	_, err := PMapIntBoolPtrErr(someLogicIntBoolPtrErr, []*int{&v10, &v3})
	if err == nil {
		t.Errorf("PMapIntBoolPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapIntBoolPtrErr(someLogicIntBoolPtrErr, []*int{&v10, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntBoolPtrErr failed")
	}

	_, err = PMapIntBoolPtrErr(someLogicIntBoolPtrErr, []*int{&v3, &v3, &v10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapIntBoolPtrErr failed")
	}

	_, err = PMapIntBoolPtrErr(someLogicIntBoolPtrErr, []*int{&v3, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntBoolErr failed")
	}
	_, err = PMapIntBoolPtrErr(someLogicIntBoolPtrErr, []*int{&v10, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntBoolPtrErr failed")
	}

	expectedList = []*bool{&vt, &vt}
	newList, _ = PMapIntBoolPtrErr(someLogicIntBoolPtrErr, []*int{&v10, &v10}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapIntBoolPtr failed")
	}

	_, err = PMapIntBoolPtrErr(someLogicIntBoolPtrErr, []*int{&v10, &v10, &v3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntBoolErr failed")
	}
	_, err = PMapIntBoolPtrErr(someLogicIntBoolPtrErr, []*int{&v10, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntBoolPtrErr failed")
	}

	_, err = PMapIntBoolPtrErr(someLogicIntBoolPtrErr, []*int{&v3, &v3, &v10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntBoolPtrErr failed")
	}

	_, err = PMapIntBoolPtrErr(someLogicIntBoolPtrErr, []*int{&v3, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntBoolPtrErr failed")
	}

	_, err = PMapIntBoolPtrErr(someLogicIntBoolPtrErr, []*int{&v10, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntBoolPtrErr failed")
	}
}

func TestPmapIntFloat32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3

	expectedList := []*float32{&vo2, &vo3}
	newList, _ := PMapIntFloat32PtrErr(plusOneIntFloat32PtrErr, []*int{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapIntFloat32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapIntFloat32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapIntFloat32PtrErr failed")
	}

	r, _ = PMapIntFloat32PtrErr(nil, []*int{})
	if len(r) > 0 {
		t.Errorf("PMapIntFloat32PtrErr failed")
	}
	_, err := PMapIntFloat32PtrErr(plusOneIntFloat32PtrErr, []*int{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapIntFloat32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapIntFloat32PtrErr(plusOneIntFloat32PtrErr, []*int{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntFloat32PtrErr failed")
	}

	_, err = PMapIntFloat32PtrErr(plusOneIntFloat32PtrErr, []*int{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapIntFloat32PtrErr failed")
	}

	_, err = PMapIntFloat32PtrErr(plusOneIntFloat32PtrErr, []*int{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntFloat32PtrErr failed")
	}
	_, err = PMapIntFloat32PtrErr(plusOneIntFloat32PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntFloat32PtrErr failed")
	}

	expectedList = []*float32{&vo2, &vo3}
	newList, _ = PMapIntFloat32PtrErr(plusOneIntFloat32PtrErr, []*int{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapIntFloat32Ptr failed")
	}

	_, err = PMapIntFloat32PtrErr(plusOneIntFloat32PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntFloat32PtrErr failed")
	}
	_, err = PMapIntFloat32PtrErr(plusOneIntFloat32PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntFloat32PtrErr failed")
	}

	_, err = PMapIntFloat32PtrErr(plusOneIntFloat32PtrErr, []*int{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntFloat32PtrErr failed")
	}

	_, err = PMapIntFloat32PtrErr(plusOneIntFloat32PtrErr, []*int{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntFloat32PtrErr failed")
	}

	_, err = PMapIntFloat32PtrErr(plusOneIntFloat32PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntFloat32PtrErr failed")
	}
}

func TestPmapIntFloat64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3

	expectedList := []*float64{&vo2, &vo3}
	newList, _ := PMapIntFloat64PtrErr(plusOneIntFloat64PtrErr, []*int{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapIntFloat64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapIntFloat64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapIntFloat64PtrErr failed")
	}

	r, _ = PMapIntFloat64PtrErr(nil, []*int{})
	if len(r) > 0 {
		t.Errorf("PMapIntFloat64PtrErr failed")
	}
	_, err := PMapIntFloat64PtrErr(plusOneIntFloat64PtrErr, []*int{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapIntFloat64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapIntFloat64PtrErr(plusOneIntFloat64PtrErr, []*int{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntFloat64PtrErr failed")
	}

	_, err = PMapIntFloat64PtrErr(plusOneIntFloat64PtrErr, []*int{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapIntFloat64PtrErr failed")
	}

	_, err = PMapIntFloat64PtrErr(plusOneIntFloat64PtrErr, []*int{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntFloat64PtrErr failed")
	}
	_, err = PMapIntFloat64PtrErr(plusOneIntFloat64PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntFloat64PtrErr failed")
	}

	expectedList = []*float64{&vo2, &vo3}
	newList, _ = PMapIntFloat64PtrErr(plusOneIntFloat64PtrErr, []*int{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapIntFloat64Ptr failed")
	}

	_, err = PMapIntFloat64PtrErr(plusOneIntFloat64PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntFloat64PtrErr failed")
	}
	_, err = PMapIntFloat64PtrErr(plusOneIntFloat64PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntFloat64PtrErr failed")
	}

	_, err = PMapIntFloat64PtrErr(plusOneIntFloat64PtrErr, []*int{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntFloat64PtrErr failed")
	}

	_, err = PMapIntFloat64PtrErr(plusOneIntFloat64PtrErr, []*int{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntFloat64PtrErr failed")
	}

	_, err = PMapIntFloat64PtrErr(plusOneIntFloat64PtrErr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapIntFloat64PtrErr failed")
	}
}

func TestPmapInt64IntPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3

	expectedList := []*int{&vo2, &vo3}
	newList, _ := PMapInt64IntPtrErr(plusOneInt64IntPtrErr, []*int64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt64IntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt64IntPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt64IntPtrErr failed")
	}

	r, _ = PMapInt64IntPtrErr(nil, []*int64{})
	if len(r) > 0 {
		t.Errorf("PMapInt64IntPtrErr failed")
	}
	_, err := PMapInt64IntPtrErr(plusOneInt64IntPtrErr, []*int64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt64IntPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt64IntPtrErr(plusOneInt64IntPtrErr, []*int64{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64IntPtrErr failed")
	}

	_, err = PMapInt64IntPtrErr(plusOneInt64IntPtrErr, []*int64{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt64IntPtrErr failed")
	}

	_, err = PMapInt64IntPtrErr(plusOneInt64IntPtrErr, []*int64{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64IntPtrErr failed")
	}
	_, err = PMapInt64IntPtrErr(plusOneInt64IntPtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64IntPtrErr failed")
	}

	expectedList = []*int{&vo2, &vo3}
	newList, _ = PMapInt64IntPtrErr(plusOneInt64IntPtrErr, []*int64{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt64IntPtr failed")
	}

	_, err = PMapInt64IntPtrErr(plusOneInt64IntPtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64IntPtrErr failed")
	}
	_, err = PMapInt64IntPtrErr(plusOneInt64IntPtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64IntPtrErr failed")
	}

	_, err = PMapInt64IntPtrErr(plusOneInt64IntPtrErr, []*int64{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64IntPtrErr failed")
	}

	_, err = PMapInt64IntPtrErr(plusOneInt64IntPtrErr, []*int64{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64IntPtrErr failed")
	}

	_, err = PMapInt64IntPtrErr(plusOneInt64IntPtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64IntPtrErr failed")
	}
}

func TestPmapInt64Int32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3

	expectedList := []*int32{&vo2, &vo3}
	newList, _ := PMapInt64Int32PtrErr(plusOneInt64Int32PtrErr, []*int64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt64Int32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt64Int32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt64Int32PtrErr failed")
	}

	r, _ = PMapInt64Int32PtrErr(nil, []*int64{})
	if len(r) > 0 {
		t.Errorf("PMapInt64Int32PtrErr failed")
	}
	_, err := PMapInt64Int32PtrErr(plusOneInt64Int32PtrErr, []*int64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt64Int32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt64Int32PtrErr(plusOneInt64Int32PtrErr, []*int64{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64Int32PtrErr failed")
	}

	_, err = PMapInt64Int32PtrErr(plusOneInt64Int32PtrErr, []*int64{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt64Int32PtrErr failed")
	}

	_, err = PMapInt64Int32PtrErr(plusOneInt64Int32PtrErr, []*int64{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64Int32PtrErr failed")
	}
	_, err = PMapInt64Int32PtrErr(plusOneInt64Int32PtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64Int32PtrErr failed")
	}

	expectedList = []*int32{&vo2, &vo3}
	newList, _ = PMapInt64Int32PtrErr(plusOneInt64Int32PtrErr, []*int64{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt64Int32Ptr failed")
	}

	_, err = PMapInt64Int32PtrErr(plusOneInt64Int32PtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Int32PtrErr failed")
	}
	_, err = PMapInt64Int32PtrErr(plusOneInt64Int32PtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Int32PtrErr failed")
	}

	_, err = PMapInt64Int32PtrErr(plusOneInt64Int32PtrErr, []*int64{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Int32PtrErr failed")
	}

	_, err = PMapInt64Int32PtrErr(plusOneInt64Int32PtrErr, []*int64{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Int32PtrErr failed")
	}

	_, err = PMapInt64Int32PtrErr(plusOneInt64Int32PtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Int32PtrErr failed")
	}
}

func TestPmapInt64Int16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3

	expectedList := []*int16{&vo2, &vo3}
	newList, _ := PMapInt64Int16PtrErr(plusOneInt64Int16PtrErr, []*int64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt64Int16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt64Int16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt64Int16PtrErr failed")
	}

	r, _ = PMapInt64Int16PtrErr(nil, []*int64{})
	if len(r) > 0 {
		t.Errorf("PMapInt64Int16PtrErr failed")
	}
	_, err := PMapInt64Int16PtrErr(plusOneInt64Int16PtrErr, []*int64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt64Int16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt64Int16PtrErr(plusOneInt64Int16PtrErr, []*int64{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64Int16PtrErr failed")
	}

	_, err = PMapInt64Int16PtrErr(plusOneInt64Int16PtrErr, []*int64{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt64Int16PtrErr failed")
	}

	_, err = PMapInt64Int16PtrErr(plusOneInt64Int16PtrErr, []*int64{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64Int16PtrErr failed")
	}
	_, err = PMapInt64Int16PtrErr(plusOneInt64Int16PtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64Int16PtrErr failed")
	}

	expectedList = []*int16{&vo2, &vo3}
	newList, _ = PMapInt64Int16PtrErr(plusOneInt64Int16PtrErr, []*int64{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt64Int16Ptr failed")
	}

	_, err = PMapInt64Int16PtrErr(plusOneInt64Int16PtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Int16PtrErr failed")
	}
	_, err = PMapInt64Int16PtrErr(plusOneInt64Int16PtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Int16PtrErr failed")
	}

	_, err = PMapInt64Int16PtrErr(plusOneInt64Int16PtrErr, []*int64{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Int16PtrErr failed")
	}

	_, err = PMapInt64Int16PtrErr(plusOneInt64Int16PtrErr, []*int64{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Int16PtrErr failed")
	}

	_, err = PMapInt64Int16PtrErr(plusOneInt64Int16PtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Int16PtrErr failed")
	}
}

func TestPmapInt64Int8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3

	expectedList := []*int8{&vo2, &vo3}
	newList, _ := PMapInt64Int8PtrErr(plusOneInt64Int8PtrErr, []*int64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt64Int8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt64Int8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt64Int8PtrErr failed")
	}

	r, _ = PMapInt64Int8PtrErr(nil, []*int64{})
	if len(r) > 0 {
		t.Errorf("PMapInt64Int8PtrErr failed")
	}
	_, err := PMapInt64Int8PtrErr(plusOneInt64Int8PtrErr, []*int64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt64Int8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt64Int8PtrErr(plusOneInt64Int8PtrErr, []*int64{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64Int8PtrErr failed")
	}

	_, err = PMapInt64Int8PtrErr(plusOneInt64Int8PtrErr, []*int64{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt64Int8PtrErr failed")
	}

	_, err = PMapInt64Int8PtrErr(plusOneInt64Int8PtrErr, []*int64{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64Int8PtrErr failed")
	}
	_, err = PMapInt64Int8PtrErr(plusOneInt64Int8PtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64Int8PtrErr failed")
	}

	expectedList = []*int8{&vo2, &vo3}
	newList, _ = PMapInt64Int8PtrErr(plusOneInt64Int8PtrErr, []*int64{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt64Int8Ptr failed")
	}

	_, err = PMapInt64Int8PtrErr(plusOneInt64Int8PtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Int8PtrErr failed")
	}
	_, err = PMapInt64Int8PtrErr(plusOneInt64Int8PtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Int8PtrErr failed")
	}

	_, err = PMapInt64Int8PtrErr(plusOneInt64Int8PtrErr, []*int64{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Int8PtrErr failed")
	}

	_, err = PMapInt64Int8PtrErr(plusOneInt64Int8PtrErr, []*int64{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Int8PtrErr failed")
	}

	_, err = PMapInt64Int8PtrErr(plusOneInt64Int8PtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Int8PtrErr failed")
	}
}

func TestPmapInt64UintPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3

	expectedList := []*uint{&vo2, &vo3}
	newList, _ := PMapInt64UintPtrErr(plusOneInt64UintPtrErr, []*int64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt64UintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt64UintPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt64UintPtrErr failed")
	}

	r, _ = PMapInt64UintPtrErr(nil, []*int64{})
	if len(r) > 0 {
		t.Errorf("PMapInt64UintPtrErr failed")
	}
	_, err := PMapInt64UintPtrErr(plusOneInt64UintPtrErr, []*int64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt64UintPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt64UintPtrErr(plusOneInt64UintPtrErr, []*int64{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64UintPtrErr failed")
	}

	_, err = PMapInt64UintPtrErr(plusOneInt64UintPtrErr, []*int64{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt64UintPtrErr failed")
	}

	_, err = PMapInt64UintPtrErr(plusOneInt64UintPtrErr, []*int64{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64UintPtrErr failed")
	}
	_, err = PMapInt64UintPtrErr(plusOneInt64UintPtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64UintPtrErr failed")
	}

	expectedList = []*uint{&vo2, &vo3}
	newList, _ = PMapInt64UintPtrErr(plusOneInt64UintPtrErr, []*int64{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt64UintPtr failed")
	}

	_, err = PMapInt64UintPtrErr(plusOneInt64UintPtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64UintPtrErr failed")
	}
	_, err = PMapInt64UintPtrErr(plusOneInt64UintPtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64UintPtrErr failed")
	}

	_, err = PMapInt64UintPtrErr(plusOneInt64UintPtrErr, []*int64{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64UintPtrErr failed")
	}

	_, err = PMapInt64UintPtrErr(plusOneInt64UintPtrErr, []*int64{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64UintPtrErr failed")
	}

	_, err = PMapInt64UintPtrErr(plusOneInt64UintPtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64UintPtrErr failed")
	}
}

func TestPmapInt64Uint64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3

	expectedList := []*uint64{&vo2, &vo3}
	newList, _ := PMapInt64Uint64PtrErr(plusOneInt64Uint64PtrErr, []*int64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt64Uint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt64Uint64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt64Uint64PtrErr failed")
	}

	r, _ = PMapInt64Uint64PtrErr(nil, []*int64{})
	if len(r) > 0 {
		t.Errorf("PMapInt64Uint64PtrErr failed")
	}
	_, err := PMapInt64Uint64PtrErr(plusOneInt64Uint64PtrErr, []*int64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt64Uint64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt64Uint64PtrErr(plusOneInt64Uint64PtrErr, []*int64{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64Uint64PtrErr failed")
	}

	_, err = PMapInt64Uint64PtrErr(plusOneInt64Uint64PtrErr, []*int64{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt64Uint64PtrErr failed")
	}

	_, err = PMapInt64Uint64PtrErr(plusOneInt64Uint64PtrErr, []*int64{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64Uint64PtrErr failed")
	}
	_, err = PMapInt64Uint64PtrErr(plusOneInt64Uint64PtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64Uint64PtrErr failed")
	}

	expectedList = []*uint64{&vo2, &vo3}
	newList, _ = PMapInt64Uint64PtrErr(plusOneInt64Uint64PtrErr, []*int64{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt64Uint64Ptr failed")
	}

	_, err = PMapInt64Uint64PtrErr(plusOneInt64Uint64PtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Uint64PtrErr failed")
	}
	_, err = PMapInt64Uint64PtrErr(plusOneInt64Uint64PtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Uint64PtrErr failed")
	}

	_, err = PMapInt64Uint64PtrErr(plusOneInt64Uint64PtrErr, []*int64{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Uint64PtrErr failed")
	}

	_, err = PMapInt64Uint64PtrErr(plusOneInt64Uint64PtrErr, []*int64{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Uint64PtrErr failed")
	}

	_, err = PMapInt64Uint64PtrErr(plusOneInt64Uint64PtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Uint64PtrErr failed")
	}
}

func TestPmapInt64Uint32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3

	expectedList := []*uint32{&vo2, &vo3}
	newList, _ := PMapInt64Uint32PtrErr(plusOneInt64Uint32PtrErr, []*int64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt64Uint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt64Uint32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt64Uint32PtrErr failed")
	}

	r, _ = PMapInt64Uint32PtrErr(nil, []*int64{})
	if len(r) > 0 {
		t.Errorf("PMapInt64Uint32PtrErr failed")
	}
	_, err := PMapInt64Uint32PtrErr(plusOneInt64Uint32PtrErr, []*int64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt64Uint32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt64Uint32PtrErr(plusOneInt64Uint32PtrErr, []*int64{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64Uint32PtrErr failed")
	}

	_, err = PMapInt64Uint32PtrErr(plusOneInt64Uint32PtrErr, []*int64{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt64Uint32PtrErr failed")
	}

	_, err = PMapInt64Uint32PtrErr(plusOneInt64Uint32PtrErr, []*int64{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64Uint32PtrErr failed")
	}
	_, err = PMapInt64Uint32PtrErr(plusOneInt64Uint32PtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64Uint32PtrErr failed")
	}

	expectedList = []*uint32{&vo2, &vo3}
	newList, _ = PMapInt64Uint32PtrErr(plusOneInt64Uint32PtrErr, []*int64{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt64Uint32Ptr failed")
	}

	_, err = PMapInt64Uint32PtrErr(plusOneInt64Uint32PtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Uint32PtrErr failed")
	}
	_, err = PMapInt64Uint32PtrErr(plusOneInt64Uint32PtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Uint32PtrErr failed")
	}

	_, err = PMapInt64Uint32PtrErr(plusOneInt64Uint32PtrErr, []*int64{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Uint32PtrErr failed")
	}

	_, err = PMapInt64Uint32PtrErr(plusOneInt64Uint32PtrErr, []*int64{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Uint32PtrErr failed")
	}

	_, err = PMapInt64Uint32PtrErr(plusOneInt64Uint32PtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Uint32PtrErr failed")
	}
}

func TestPmapInt64Uint16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3

	expectedList := []*uint16{&vo2, &vo3}
	newList, _ := PMapInt64Uint16PtrErr(plusOneInt64Uint16PtrErr, []*int64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt64Uint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt64Uint16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt64Uint16PtrErr failed")
	}

	r, _ = PMapInt64Uint16PtrErr(nil, []*int64{})
	if len(r) > 0 {
		t.Errorf("PMapInt64Uint16PtrErr failed")
	}
	_, err := PMapInt64Uint16PtrErr(plusOneInt64Uint16PtrErr, []*int64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt64Uint16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt64Uint16PtrErr(plusOneInt64Uint16PtrErr, []*int64{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64Uint16PtrErr failed")
	}

	_, err = PMapInt64Uint16PtrErr(plusOneInt64Uint16PtrErr, []*int64{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt64Uint16PtrErr failed")
	}

	_, err = PMapInt64Uint16PtrErr(plusOneInt64Uint16PtrErr, []*int64{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64Uint16PtrErr failed")
	}
	_, err = PMapInt64Uint16PtrErr(plusOneInt64Uint16PtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64Uint16PtrErr failed")
	}

	expectedList = []*uint16{&vo2, &vo3}
	newList, _ = PMapInt64Uint16PtrErr(plusOneInt64Uint16PtrErr, []*int64{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt64Uint16Ptr failed")
	}

	_, err = PMapInt64Uint16PtrErr(plusOneInt64Uint16PtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Uint16PtrErr failed")
	}
	_, err = PMapInt64Uint16PtrErr(plusOneInt64Uint16PtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Uint16PtrErr failed")
	}

	_, err = PMapInt64Uint16PtrErr(plusOneInt64Uint16PtrErr, []*int64{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Uint16PtrErr failed")
	}

	_, err = PMapInt64Uint16PtrErr(plusOneInt64Uint16PtrErr, []*int64{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Uint16PtrErr failed")
	}

	_, err = PMapInt64Uint16PtrErr(plusOneInt64Uint16PtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Uint16PtrErr failed")
	}
}

func TestPmapInt64Uint8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3

	expectedList := []*uint8{&vo2, &vo3}
	newList, _ := PMapInt64Uint8PtrErr(plusOneInt64Uint8PtrErr, []*int64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt64Uint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt64Uint8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt64Uint8PtrErr failed")
	}

	r, _ = PMapInt64Uint8PtrErr(nil, []*int64{})
	if len(r) > 0 {
		t.Errorf("PMapInt64Uint8PtrErr failed")
	}
	_, err := PMapInt64Uint8PtrErr(plusOneInt64Uint8PtrErr, []*int64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt64Uint8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt64Uint8PtrErr(plusOneInt64Uint8PtrErr, []*int64{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64Uint8PtrErr failed")
	}

	_, err = PMapInt64Uint8PtrErr(plusOneInt64Uint8PtrErr, []*int64{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt64Uint8PtrErr failed")
	}

	_, err = PMapInt64Uint8PtrErr(plusOneInt64Uint8PtrErr, []*int64{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64Uint8PtrErr failed")
	}
	_, err = PMapInt64Uint8PtrErr(plusOneInt64Uint8PtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64Uint8PtrErr failed")
	}

	expectedList = []*uint8{&vo2, &vo3}
	newList, _ = PMapInt64Uint8PtrErr(plusOneInt64Uint8PtrErr, []*int64{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt64Uint8Ptr failed")
	}

	_, err = PMapInt64Uint8PtrErr(plusOneInt64Uint8PtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Uint8PtrErr failed")
	}
	_, err = PMapInt64Uint8PtrErr(plusOneInt64Uint8PtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Uint8PtrErr failed")
	}

	_, err = PMapInt64Uint8PtrErr(plusOneInt64Uint8PtrErr, []*int64{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Uint8PtrErr failed")
	}

	_, err = PMapInt64Uint8PtrErr(plusOneInt64Uint8PtrErr, []*int64{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Uint8PtrErr failed")
	}

	_, err = PMapInt64Uint8PtrErr(plusOneInt64Uint8PtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Uint8PtrErr failed")
	}
}

func TestPmapInt64StrPtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 int64 = 10
	var vi0 int64

	expectedList := []*string{&vo10}
	newList, _ := PMapInt64StrPtrErr(someLogicInt64StrPtrErr, []*int64{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapInt64StrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt64StrPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt64StrPtrErr failed")
	}

	r, _ = PMapInt64StrPtrErr(nil, []*int64{})
	if len(r) > 0 {
		t.Errorf("PMapInt64StrPtrErr failed")
	}
	_, err := PMapInt64StrPtrErr(someLogicInt64StrPtrErr, []*int64{&vi10, &vi0})
	if err == nil {
		t.Errorf("PMapInt64StrPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt64StrPtrErr(someLogicInt64StrPtrErr, []*int64{&vi10, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64StrPtrErr failed")
	}

	_, err = PMapInt64StrPtrErr(someLogicInt64StrPtrErr, []*int64{&vi0, &vi0, &vi10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt64StrPtrErr failed")
	}

	_, err = PMapInt64StrPtrErr(someLogicInt64StrPtrErr, []*int64{&vi0, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64StrPtrErr failed")
	}
	_, err = PMapInt64StrPtrErr(someLogicInt64StrPtrErr, []*int64{&vi10, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64StrPtrErr failed")
	}

	expectedList = []*string{&vo10, &vo10}
	newList, _ = PMapInt64StrPtrErr(someLogicInt64StrPtrErr, []*int64{&vi10, &vi10}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt64StrPtr failed")
	}

	_, err = PMapInt64StrPtrErr(someLogicInt64StrPtrErr, []*int64{&vi10, &vi10, &vi0}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64StrPtrErr failed")
	}
	_, err = PMapInt64StrPtrErr(someLogicInt64StrPtrErr, []*int64{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64StrPtrErr failed")
	}

	_, err = PMapInt64StrPtrErr(someLogicInt64StrPtrErr, []*int64{&vi0, &vi0, &vi10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64StrPtrErr failed")
	}

	_, err = PMapInt64StrPtrErr(someLogicInt64StrPtrErr, []*int64{&vi0, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64StrPtrErr failed")
	}

	_, err = PMapInt64StrPtrErr(someLogicInt64StrPtrErr, []*int64{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64StrPtrErr failed")
	}
}

func TestPmapInt64BoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 int64 = 10
	var v0 int64
	var v3 int64 = 3

	expectedList := []*bool{&vt, &vf}
	newList, _ := PMapInt64BoolPtrErr(someLogicInt64BoolPtrErr, []*int64{&v10, &v0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt64BoolPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt64BoolPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt64BoolPtrErr failed")
	}

	r, _ = PMapInt64BoolPtrErr(nil, []*int64{})
	if len(r) > 0 {
		t.Errorf("PMapInt64BoolPtrErr failed")
	}
	_, err := PMapInt64BoolPtrErr(someLogicInt64BoolPtrErr, []*int64{&v10, &v3})
	if err == nil {
		t.Errorf("PMapInt64BoolPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt64BoolPtrErr(someLogicInt64BoolPtrErr, []*int64{&v10, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64BoolPtrErr failed")
	}

	_, err = PMapInt64BoolPtrErr(someLogicInt64BoolPtrErr, []*int64{&v3, &v3, &v10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt64BoolPtrErr failed")
	}

	_, err = PMapInt64BoolPtrErr(someLogicInt64BoolPtrErr, []*int64{&v3, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64BoolErr failed")
	}
	_, err = PMapInt64BoolPtrErr(someLogicInt64BoolPtrErr, []*int64{&v10, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64BoolPtrErr failed")
	}

	expectedList = []*bool{&vt, &vt}
	newList, _ = PMapInt64BoolPtrErr(someLogicInt64BoolPtrErr, []*int64{&v10, &v10}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt64BoolPtr failed")
	}

	_, err = PMapInt64BoolPtrErr(someLogicInt64BoolPtrErr, []*int64{&v10, &v10, &v3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64BoolErr failed")
	}
	_, err = PMapInt64BoolPtrErr(someLogicInt64BoolPtrErr, []*int64{&v10, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64BoolPtrErr failed")
	}

	_, err = PMapInt64BoolPtrErr(someLogicInt64BoolPtrErr, []*int64{&v3, &v3, &v10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64BoolPtrErr failed")
	}

	_, err = PMapInt64BoolPtrErr(someLogicInt64BoolPtrErr, []*int64{&v3, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64BoolPtrErr failed")
	}

	_, err = PMapInt64BoolPtrErr(someLogicInt64BoolPtrErr, []*int64{&v10, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64BoolPtrErr failed")
	}
}

func TestPmapInt64Float32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3

	expectedList := []*float32{&vo2, &vo3}
	newList, _ := PMapInt64Float32PtrErr(plusOneInt64Float32PtrErr, []*int64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt64Float32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt64Float32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt64Float32PtrErr failed")
	}

	r, _ = PMapInt64Float32PtrErr(nil, []*int64{})
	if len(r) > 0 {
		t.Errorf("PMapInt64Float32PtrErr failed")
	}
	_, err := PMapInt64Float32PtrErr(plusOneInt64Float32PtrErr, []*int64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt64Float32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt64Float32PtrErr(plusOneInt64Float32PtrErr, []*int64{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64Float32PtrErr failed")
	}

	_, err = PMapInt64Float32PtrErr(plusOneInt64Float32PtrErr, []*int64{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt64Float32PtrErr failed")
	}

	_, err = PMapInt64Float32PtrErr(plusOneInt64Float32PtrErr, []*int64{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64Float32PtrErr failed")
	}
	_, err = PMapInt64Float32PtrErr(plusOneInt64Float32PtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64Float32PtrErr failed")
	}

	expectedList = []*float32{&vo2, &vo3}
	newList, _ = PMapInt64Float32PtrErr(plusOneInt64Float32PtrErr, []*int64{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt64Float32Ptr failed")
	}

	_, err = PMapInt64Float32PtrErr(plusOneInt64Float32PtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Float32PtrErr failed")
	}
	_, err = PMapInt64Float32PtrErr(plusOneInt64Float32PtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Float32PtrErr failed")
	}

	_, err = PMapInt64Float32PtrErr(plusOneInt64Float32PtrErr, []*int64{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Float32PtrErr failed")
	}

	_, err = PMapInt64Float32PtrErr(plusOneInt64Float32PtrErr, []*int64{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Float32PtrErr failed")
	}

	_, err = PMapInt64Float32PtrErr(plusOneInt64Float32PtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Float32PtrErr failed")
	}
}

func TestPmapInt64Float64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3

	expectedList := []*float64{&vo2, &vo3}
	newList, _ := PMapInt64Float64PtrErr(plusOneInt64Float64PtrErr, []*int64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt64Float64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt64Float64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt64Float64PtrErr failed")
	}

	r, _ = PMapInt64Float64PtrErr(nil, []*int64{})
	if len(r) > 0 {
		t.Errorf("PMapInt64Float64PtrErr failed")
	}
	_, err := PMapInt64Float64PtrErr(plusOneInt64Float64PtrErr, []*int64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt64Float64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt64Float64PtrErr(plusOneInt64Float64PtrErr, []*int64{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64Float64PtrErr failed")
	}

	_, err = PMapInt64Float64PtrErr(plusOneInt64Float64PtrErr, []*int64{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt64Float64PtrErr failed")
	}

	_, err = PMapInt64Float64PtrErr(plusOneInt64Float64PtrErr, []*int64{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64Float64PtrErr failed")
	}
	_, err = PMapInt64Float64PtrErr(plusOneInt64Float64PtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64Float64PtrErr failed")
	}

	expectedList = []*float64{&vo2, &vo3}
	newList, _ = PMapInt64Float64PtrErr(plusOneInt64Float64PtrErr, []*int64{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt64Float64Ptr failed")
	}

	_, err = PMapInt64Float64PtrErr(plusOneInt64Float64PtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Float64PtrErr failed")
	}
	_, err = PMapInt64Float64PtrErr(plusOneInt64Float64PtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Float64PtrErr failed")
	}

	_, err = PMapInt64Float64PtrErr(plusOneInt64Float64PtrErr, []*int64{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Float64PtrErr failed")
	}

	_, err = PMapInt64Float64PtrErr(plusOneInt64Float64PtrErr, []*int64{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Float64PtrErr failed")
	}

	_, err = PMapInt64Float64PtrErr(plusOneInt64Float64PtrErr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt64Float64PtrErr failed")
	}
}

func TestPmapInt32IntPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3

	expectedList := []*int{&vo2, &vo3}
	newList, _ := PMapInt32IntPtrErr(plusOneInt32IntPtrErr, []*int32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt32IntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt32IntPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt32IntPtrErr failed")
	}

	r, _ = PMapInt32IntPtrErr(nil, []*int32{})
	if len(r) > 0 {
		t.Errorf("PMapInt32IntPtrErr failed")
	}
	_, err := PMapInt32IntPtrErr(plusOneInt32IntPtrErr, []*int32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt32IntPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt32IntPtrErr(plusOneInt32IntPtrErr, []*int32{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32IntPtrErr failed")
	}

	_, err = PMapInt32IntPtrErr(plusOneInt32IntPtrErr, []*int32{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt32IntPtrErr failed")
	}

	_, err = PMapInt32IntPtrErr(plusOneInt32IntPtrErr, []*int32{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32IntPtrErr failed")
	}
	_, err = PMapInt32IntPtrErr(plusOneInt32IntPtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32IntPtrErr failed")
	}

	expectedList = []*int{&vo2, &vo3}
	newList, _ = PMapInt32IntPtrErr(plusOneInt32IntPtrErr, []*int32{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt32IntPtr failed")
	}

	_, err = PMapInt32IntPtrErr(plusOneInt32IntPtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32IntPtrErr failed")
	}
	_, err = PMapInt32IntPtrErr(plusOneInt32IntPtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32IntPtrErr failed")
	}

	_, err = PMapInt32IntPtrErr(plusOneInt32IntPtrErr, []*int32{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32IntPtrErr failed")
	}

	_, err = PMapInt32IntPtrErr(plusOneInt32IntPtrErr, []*int32{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32IntPtrErr failed")
	}

	_, err = PMapInt32IntPtrErr(plusOneInt32IntPtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32IntPtrErr failed")
	}
}

func TestPmapInt32Int64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3

	expectedList := []*int64{&vo2, &vo3}
	newList, _ := PMapInt32Int64PtrErr(plusOneInt32Int64PtrErr, []*int32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt32Int64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt32Int64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt32Int64PtrErr failed")
	}

	r, _ = PMapInt32Int64PtrErr(nil, []*int32{})
	if len(r) > 0 {
		t.Errorf("PMapInt32Int64PtrErr failed")
	}
	_, err := PMapInt32Int64PtrErr(plusOneInt32Int64PtrErr, []*int32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt32Int64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt32Int64PtrErr(plusOneInt32Int64PtrErr, []*int32{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32Int64PtrErr failed")
	}

	_, err = PMapInt32Int64PtrErr(plusOneInt32Int64PtrErr, []*int32{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt32Int64PtrErr failed")
	}

	_, err = PMapInt32Int64PtrErr(plusOneInt32Int64PtrErr, []*int32{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32Int64PtrErr failed")
	}
	_, err = PMapInt32Int64PtrErr(plusOneInt32Int64PtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32Int64PtrErr failed")
	}

	expectedList = []*int64{&vo2, &vo3}
	newList, _ = PMapInt32Int64PtrErr(plusOneInt32Int64PtrErr, []*int32{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt32Int64Ptr failed")
	}

	_, err = PMapInt32Int64PtrErr(plusOneInt32Int64PtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Int64PtrErr failed")
	}
	_, err = PMapInt32Int64PtrErr(plusOneInt32Int64PtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Int64PtrErr failed")
	}

	_, err = PMapInt32Int64PtrErr(plusOneInt32Int64PtrErr, []*int32{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Int64PtrErr failed")
	}

	_, err = PMapInt32Int64PtrErr(plusOneInt32Int64PtrErr, []*int32{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Int64PtrErr failed")
	}

	_, err = PMapInt32Int64PtrErr(plusOneInt32Int64PtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Int64PtrErr failed")
	}
}

func TestPmapInt32Int16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3

	expectedList := []*int16{&vo2, &vo3}
	newList, _ := PMapInt32Int16PtrErr(plusOneInt32Int16PtrErr, []*int32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt32Int16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt32Int16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt32Int16PtrErr failed")
	}

	r, _ = PMapInt32Int16PtrErr(nil, []*int32{})
	if len(r) > 0 {
		t.Errorf("PMapInt32Int16PtrErr failed")
	}
	_, err := PMapInt32Int16PtrErr(plusOneInt32Int16PtrErr, []*int32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt32Int16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt32Int16PtrErr(plusOneInt32Int16PtrErr, []*int32{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32Int16PtrErr failed")
	}

	_, err = PMapInt32Int16PtrErr(plusOneInt32Int16PtrErr, []*int32{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt32Int16PtrErr failed")
	}

	_, err = PMapInt32Int16PtrErr(plusOneInt32Int16PtrErr, []*int32{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32Int16PtrErr failed")
	}
	_, err = PMapInt32Int16PtrErr(plusOneInt32Int16PtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32Int16PtrErr failed")
	}

	expectedList = []*int16{&vo2, &vo3}
	newList, _ = PMapInt32Int16PtrErr(plusOneInt32Int16PtrErr, []*int32{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt32Int16Ptr failed")
	}

	_, err = PMapInt32Int16PtrErr(plusOneInt32Int16PtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Int16PtrErr failed")
	}
	_, err = PMapInt32Int16PtrErr(plusOneInt32Int16PtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Int16PtrErr failed")
	}

	_, err = PMapInt32Int16PtrErr(plusOneInt32Int16PtrErr, []*int32{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Int16PtrErr failed")
	}

	_, err = PMapInt32Int16PtrErr(plusOneInt32Int16PtrErr, []*int32{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Int16PtrErr failed")
	}

	_, err = PMapInt32Int16PtrErr(plusOneInt32Int16PtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Int16PtrErr failed")
	}
}

func TestPmapInt32Int8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3

	expectedList := []*int8{&vo2, &vo3}
	newList, _ := PMapInt32Int8PtrErr(plusOneInt32Int8PtrErr, []*int32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt32Int8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt32Int8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt32Int8PtrErr failed")
	}

	r, _ = PMapInt32Int8PtrErr(nil, []*int32{})
	if len(r) > 0 {
		t.Errorf("PMapInt32Int8PtrErr failed")
	}
	_, err := PMapInt32Int8PtrErr(plusOneInt32Int8PtrErr, []*int32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt32Int8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt32Int8PtrErr(plusOneInt32Int8PtrErr, []*int32{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32Int8PtrErr failed")
	}

	_, err = PMapInt32Int8PtrErr(plusOneInt32Int8PtrErr, []*int32{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt32Int8PtrErr failed")
	}

	_, err = PMapInt32Int8PtrErr(plusOneInt32Int8PtrErr, []*int32{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32Int8PtrErr failed")
	}
	_, err = PMapInt32Int8PtrErr(plusOneInt32Int8PtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32Int8PtrErr failed")
	}

	expectedList = []*int8{&vo2, &vo3}
	newList, _ = PMapInt32Int8PtrErr(plusOneInt32Int8PtrErr, []*int32{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt32Int8Ptr failed")
	}

	_, err = PMapInt32Int8PtrErr(plusOneInt32Int8PtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Int8PtrErr failed")
	}
	_, err = PMapInt32Int8PtrErr(plusOneInt32Int8PtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Int8PtrErr failed")
	}

	_, err = PMapInt32Int8PtrErr(plusOneInt32Int8PtrErr, []*int32{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Int8PtrErr failed")
	}

	_, err = PMapInt32Int8PtrErr(plusOneInt32Int8PtrErr, []*int32{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Int8PtrErr failed")
	}

	_, err = PMapInt32Int8PtrErr(plusOneInt32Int8PtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Int8PtrErr failed")
	}
}

func TestPmapInt32UintPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3

	expectedList := []*uint{&vo2, &vo3}
	newList, _ := PMapInt32UintPtrErr(plusOneInt32UintPtrErr, []*int32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt32UintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt32UintPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt32UintPtrErr failed")
	}

	r, _ = PMapInt32UintPtrErr(nil, []*int32{})
	if len(r) > 0 {
		t.Errorf("PMapInt32UintPtrErr failed")
	}
	_, err := PMapInt32UintPtrErr(plusOneInt32UintPtrErr, []*int32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt32UintPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt32UintPtrErr(plusOneInt32UintPtrErr, []*int32{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32UintPtrErr failed")
	}

	_, err = PMapInt32UintPtrErr(plusOneInt32UintPtrErr, []*int32{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt32UintPtrErr failed")
	}

	_, err = PMapInt32UintPtrErr(plusOneInt32UintPtrErr, []*int32{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32UintPtrErr failed")
	}
	_, err = PMapInt32UintPtrErr(plusOneInt32UintPtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32UintPtrErr failed")
	}

	expectedList = []*uint{&vo2, &vo3}
	newList, _ = PMapInt32UintPtrErr(plusOneInt32UintPtrErr, []*int32{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt32UintPtr failed")
	}

	_, err = PMapInt32UintPtrErr(plusOneInt32UintPtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32UintPtrErr failed")
	}
	_, err = PMapInt32UintPtrErr(plusOneInt32UintPtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32UintPtrErr failed")
	}

	_, err = PMapInt32UintPtrErr(plusOneInt32UintPtrErr, []*int32{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32UintPtrErr failed")
	}

	_, err = PMapInt32UintPtrErr(plusOneInt32UintPtrErr, []*int32{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32UintPtrErr failed")
	}

	_, err = PMapInt32UintPtrErr(plusOneInt32UintPtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32UintPtrErr failed")
	}
}

func TestPmapInt32Uint64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3

	expectedList := []*uint64{&vo2, &vo3}
	newList, _ := PMapInt32Uint64PtrErr(plusOneInt32Uint64PtrErr, []*int32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt32Uint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt32Uint64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt32Uint64PtrErr failed")
	}

	r, _ = PMapInt32Uint64PtrErr(nil, []*int32{})
	if len(r) > 0 {
		t.Errorf("PMapInt32Uint64PtrErr failed")
	}
	_, err := PMapInt32Uint64PtrErr(plusOneInt32Uint64PtrErr, []*int32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt32Uint64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt32Uint64PtrErr(plusOneInt32Uint64PtrErr, []*int32{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32Uint64PtrErr failed")
	}

	_, err = PMapInt32Uint64PtrErr(plusOneInt32Uint64PtrErr, []*int32{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt32Uint64PtrErr failed")
	}

	_, err = PMapInt32Uint64PtrErr(plusOneInt32Uint64PtrErr, []*int32{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32Uint64PtrErr failed")
	}
	_, err = PMapInt32Uint64PtrErr(plusOneInt32Uint64PtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32Uint64PtrErr failed")
	}

	expectedList = []*uint64{&vo2, &vo3}
	newList, _ = PMapInt32Uint64PtrErr(plusOneInt32Uint64PtrErr, []*int32{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt32Uint64Ptr failed")
	}

	_, err = PMapInt32Uint64PtrErr(plusOneInt32Uint64PtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Uint64PtrErr failed")
	}
	_, err = PMapInt32Uint64PtrErr(plusOneInt32Uint64PtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Uint64PtrErr failed")
	}

	_, err = PMapInt32Uint64PtrErr(plusOneInt32Uint64PtrErr, []*int32{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Uint64PtrErr failed")
	}

	_, err = PMapInt32Uint64PtrErr(plusOneInt32Uint64PtrErr, []*int32{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Uint64PtrErr failed")
	}

	_, err = PMapInt32Uint64PtrErr(plusOneInt32Uint64PtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Uint64PtrErr failed")
	}
}

func TestPmapInt32Uint32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3

	expectedList := []*uint32{&vo2, &vo3}
	newList, _ := PMapInt32Uint32PtrErr(plusOneInt32Uint32PtrErr, []*int32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt32Uint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt32Uint32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt32Uint32PtrErr failed")
	}

	r, _ = PMapInt32Uint32PtrErr(nil, []*int32{})
	if len(r) > 0 {
		t.Errorf("PMapInt32Uint32PtrErr failed")
	}
	_, err := PMapInt32Uint32PtrErr(plusOneInt32Uint32PtrErr, []*int32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt32Uint32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt32Uint32PtrErr(plusOneInt32Uint32PtrErr, []*int32{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32Uint32PtrErr failed")
	}

	_, err = PMapInt32Uint32PtrErr(plusOneInt32Uint32PtrErr, []*int32{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt32Uint32PtrErr failed")
	}

	_, err = PMapInt32Uint32PtrErr(plusOneInt32Uint32PtrErr, []*int32{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32Uint32PtrErr failed")
	}
	_, err = PMapInt32Uint32PtrErr(plusOneInt32Uint32PtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32Uint32PtrErr failed")
	}

	expectedList = []*uint32{&vo2, &vo3}
	newList, _ = PMapInt32Uint32PtrErr(plusOneInt32Uint32PtrErr, []*int32{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt32Uint32Ptr failed")
	}

	_, err = PMapInt32Uint32PtrErr(plusOneInt32Uint32PtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Uint32PtrErr failed")
	}
	_, err = PMapInt32Uint32PtrErr(plusOneInt32Uint32PtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Uint32PtrErr failed")
	}

	_, err = PMapInt32Uint32PtrErr(plusOneInt32Uint32PtrErr, []*int32{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Uint32PtrErr failed")
	}

	_, err = PMapInt32Uint32PtrErr(plusOneInt32Uint32PtrErr, []*int32{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Uint32PtrErr failed")
	}

	_, err = PMapInt32Uint32PtrErr(plusOneInt32Uint32PtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Uint32PtrErr failed")
	}
}

func TestPmapInt32Uint16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3

	expectedList := []*uint16{&vo2, &vo3}
	newList, _ := PMapInt32Uint16PtrErr(plusOneInt32Uint16PtrErr, []*int32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt32Uint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt32Uint16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt32Uint16PtrErr failed")
	}

	r, _ = PMapInt32Uint16PtrErr(nil, []*int32{})
	if len(r) > 0 {
		t.Errorf("PMapInt32Uint16PtrErr failed")
	}
	_, err := PMapInt32Uint16PtrErr(plusOneInt32Uint16PtrErr, []*int32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt32Uint16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt32Uint16PtrErr(plusOneInt32Uint16PtrErr, []*int32{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32Uint16PtrErr failed")
	}

	_, err = PMapInt32Uint16PtrErr(plusOneInt32Uint16PtrErr, []*int32{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt32Uint16PtrErr failed")
	}

	_, err = PMapInt32Uint16PtrErr(plusOneInt32Uint16PtrErr, []*int32{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32Uint16PtrErr failed")
	}
	_, err = PMapInt32Uint16PtrErr(plusOneInt32Uint16PtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32Uint16PtrErr failed")
	}

	expectedList = []*uint16{&vo2, &vo3}
	newList, _ = PMapInt32Uint16PtrErr(plusOneInt32Uint16PtrErr, []*int32{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt32Uint16Ptr failed")
	}

	_, err = PMapInt32Uint16PtrErr(plusOneInt32Uint16PtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Uint16PtrErr failed")
	}
	_, err = PMapInt32Uint16PtrErr(plusOneInt32Uint16PtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Uint16PtrErr failed")
	}

	_, err = PMapInt32Uint16PtrErr(plusOneInt32Uint16PtrErr, []*int32{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Uint16PtrErr failed")
	}

	_, err = PMapInt32Uint16PtrErr(plusOneInt32Uint16PtrErr, []*int32{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Uint16PtrErr failed")
	}

	_, err = PMapInt32Uint16PtrErr(plusOneInt32Uint16PtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Uint16PtrErr failed")
	}
}

func TestPmapInt32Uint8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3

	expectedList := []*uint8{&vo2, &vo3}
	newList, _ := PMapInt32Uint8PtrErr(plusOneInt32Uint8PtrErr, []*int32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt32Uint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt32Uint8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt32Uint8PtrErr failed")
	}

	r, _ = PMapInt32Uint8PtrErr(nil, []*int32{})
	if len(r) > 0 {
		t.Errorf("PMapInt32Uint8PtrErr failed")
	}
	_, err := PMapInt32Uint8PtrErr(plusOneInt32Uint8PtrErr, []*int32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt32Uint8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt32Uint8PtrErr(plusOneInt32Uint8PtrErr, []*int32{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32Uint8PtrErr failed")
	}

	_, err = PMapInt32Uint8PtrErr(plusOneInt32Uint8PtrErr, []*int32{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt32Uint8PtrErr failed")
	}

	_, err = PMapInt32Uint8PtrErr(plusOneInt32Uint8PtrErr, []*int32{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32Uint8PtrErr failed")
	}
	_, err = PMapInt32Uint8PtrErr(plusOneInt32Uint8PtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32Uint8PtrErr failed")
	}

	expectedList = []*uint8{&vo2, &vo3}
	newList, _ = PMapInt32Uint8PtrErr(plusOneInt32Uint8PtrErr, []*int32{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt32Uint8Ptr failed")
	}

	_, err = PMapInt32Uint8PtrErr(plusOneInt32Uint8PtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Uint8PtrErr failed")
	}
	_, err = PMapInt32Uint8PtrErr(plusOneInt32Uint8PtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Uint8PtrErr failed")
	}

	_, err = PMapInt32Uint8PtrErr(plusOneInt32Uint8PtrErr, []*int32{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Uint8PtrErr failed")
	}

	_, err = PMapInt32Uint8PtrErr(plusOneInt32Uint8PtrErr, []*int32{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Uint8PtrErr failed")
	}

	_, err = PMapInt32Uint8PtrErr(plusOneInt32Uint8PtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Uint8PtrErr failed")
	}
}

func TestPmapInt32StrPtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 int32 = 10
	var vi0 int32

	expectedList := []*string{&vo10}
	newList, _ := PMapInt32StrPtrErr(someLogicInt32StrPtrErr, []*int32{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapInt32StrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt32StrPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt32StrPtrErr failed")
	}

	r, _ = PMapInt32StrPtrErr(nil, []*int32{})
	if len(r) > 0 {
		t.Errorf("PMapInt32StrPtrErr failed")
	}
	_, err := PMapInt32StrPtrErr(someLogicInt32StrPtrErr, []*int32{&vi10, &vi0})
	if err == nil {
		t.Errorf("PMapInt32StrPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt32StrPtrErr(someLogicInt32StrPtrErr, []*int32{&vi10, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32StrPtrErr failed")
	}

	_, err = PMapInt32StrPtrErr(someLogicInt32StrPtrErr, []*int32{&vi0, &vi0, &vi10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt32StrPtrErr failed")
	}

	_, err = PMapInt32StrPtrErr(someLogicInt32StrPtrErr, []*int32{&vi0, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32StrPtrErr failed")
	}
	_, err = PMapInt32StrPtrErr(someLogicInt32StrPtrErr, []*int32{&vi10, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32StrPtrErr failed")
	}

	expectedList = []*string{&vo10, &vo10}
	newList, _ = PMapInt32StrPtrErr(someLogicInt32StrPtrErr, []*int32{&vi10, &vi10}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt32StrPtr failed")
	}

	_, err = PMapInt32StrPtrErr(someLogicInt32StrPtrErr, []*int32{&vi10, &vi10, &vi0}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32StrPtrErr failed")
	}
	_, err = PMapInt32StrPtrErr(someLogicInt32StrPtrErr, []*int32{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32StrPtrErr failed")
	}

	_, err = PMapInt32StrPtrErr(someLogicInt32StrPtrErr, []*int32{&vi0, &vi0, &vi10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32StrPtrErr failed")
	}

	_, err = PMapInt32StrPtrErr(someLogicInt32StrPtrErr, []*int32{&vi0, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32StrPtrErr failed")
	}

	_, err = PMapInt32StrPtrErr(someLogicInt32StrPtrErr, []*int32{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32StrPtrErr failed")
	}
}

func TestPmapInt32BoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 int32 = 10
	var v0 int32
	var v3 int32 = 3

	expectedList := []*bool{&vt, &vf}
	newList, _ := PMapInt32BoolPtrErr(someLogicInt32BoolPtrErr, []*int32{&v10, &v0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt32BoolPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt32BoolPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt32BoolPtrErr failed")
	}

	r, _ = PMapInt32BoolPtrErr(nil, []*int32{})
	if len(r) > 0 {
		t.Errorf("PMapInt32BoolPtrErr failed")
	}
	_, err := PMapInt32BoolPtrErr(someLogicInt32BoolPtrErr, []*int32{&v10, &v3})
	if err == nil {
		t.Errorf("PMapInt32BoolPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt32BoolPtrErr(someLogicInt32BoolPtrErr, []*int32{&v10, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32BoolPtrErr failed")
	}

	_, err = PMapInt32BoolPtrErr(someLogicInt32BoolPtrErr, []*int32{&v3, &v3, &v10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt32BoolPtrErr failed")
	}

	_, err = PMapInt32BoolPtrErr(someLogicInt32BoolPtrErr, []*int32{&v3, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32BoolErr failed")
	}
	_, err = PMapInt32BoolPtrErr(someLogicInt32BoolPtrErr, []*int32{&v10, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32BoolPtrErr failed")
	}

	expectedList = []*bool{&vt, &vt}
	newList, _ = PMapInt32BoolPtrErr(someLogicInt32BoolPtrErr, []*int32{&v10, &v10}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt32BoolPtr failed")
	}

	_, err = PMapInt32BoolPtrErr(someLogicInt32BoolPtrErr, []*int32{&v10, &v10, &v3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32BoolErr failed")
	}
	_, err = PMapInt32BoolPtrErr(someLogicInt32BoolPtrErr, []*int32{&v10, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32BoolPtrErr failed")
	}

	_, err = PMapInt32BoolPtrErr(someLogicInt32BoolPtrErr, []*int32{&v3, &v3, &v10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32BoolPtrErr failed")
	}

	_, err = PMapInt32BoolPtrErr(someLogicInt32BoolPtrErr, []*int32{&v3, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32BoolPtrErr failed")
	}

	_, err = PMapInt32BoolPtrErr(someLogicInt32BoolPtrErr, []*int32{&v10, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32BoolPtrErr failed")
	}
}

func TestPmapInt32Float32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3

	expectedList := []*float32{&vo2, &vo3}
	newList, _ := PMapInt32Float32PtrErr(plusOneInt32Float32PtrErr, []*int32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt32Float32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt32Float32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt32Float32PtrErr failed")
	}

	r, _ = PMapInt32Float32PtrErr(nil, []*int32{})
	if len(r) > 0 {
		t.Errorf("PMapInt32Float32PtrErr failed")
	}
	_, err := PMapInt32Float32PtrErr(plusOneInt32Float32PtrErr, []*int32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt32Float32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt32Float32PtrErr(plusOneInt32Float32PtrErr, []*int32{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32Float32PtrErr failed")
	}

	_, err = PMapInt32Float32PtrErr(plusOneInt32Float32PtrErr, []*int32{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt32Float32PtrErr failed")
	}

	_, err = PMapInt32Float32PtrErr(plusOneInt32Float32PtrErr, []*int32{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32Float32PtrErr failed")
	}
	_, err = PMapInt32Float32PtrErr(plusOneInt32Float32PtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32Float32PtrErr failed")
	}

	expectedList = []*float32{&vo2, &vo3}
	newList, _ = PMapInt32Float32PtrErr(plusOneInt32Float32PtrErr, []*int32{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt32Float32Ptr failed")
	}

	_, err = PMapInt32Float32PtrErr(plusOneInt32Float32PtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Float32PtrErr failed")
	}
	_, err = PMapInt32Float32PtrErr(plusOneInt32Float32PtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Float32PtrErr failed")
	}

	_, err = PMapInt32Float32PtrErr(plusOneInt32Float32PtrErr, []*int32{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Float32PtrErr failed")
	}

	_, err = PMapInt32Float32PtrErr(plusOneInt32Float32PtrErr, []*int32{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Float32PtrErr failed")
	}

	_, err = PMapInt32Float32PtrErr(plusOneInt32Float32PtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Float32PtrErr failed")
	}
}

func TestPmapInt32Float64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3

	expectedList := []*float64{&vo2, &vo3}
	newList, _ := PMapInt32Float64PtrErr(plusOneInt32Float64PtrErr, []*int32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt32Float64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt32Float64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt32Float64PtrErr failed")
	}

	r, _ = PMapInt32Float64PtrErr(nil, []*int32{})
	if len(r) > 0 {
		t.Errorf("PMapInt32Float64PtrErr failed")
	}
	_, err := PMapInt32Float64PtrErr(plusOneInt32Float64PtrErr, []*int32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt32Float64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt32Float64PtrErr(plusOneInt32Float64PtrErr, []*int32{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32Float64PtrErr failed")
	}

	_, err = PMapInt32Float64PtrErr(plusOneInt32Float64PtrErr, []*int32{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt32Float64PtrErr failed")
	}

	_, err = PMapInt32Float64PtrErr(plusOneInt32Float64PtrErr, []*int32{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32Float64PtrErr failed")
	}
	_, err = PMapInt32Float64PtrErr(plusOneInt32Float64PtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32Float64PtrErr failed")
	}

	expectedList = []*float64{&vo2, &vo3}
	newList, _ = PMapInt32Float64PtrErr(plusOneInt32Float64PtrErr, []*int32{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt32Float64Ptr failed")
	}

	_, err = PMapInt32Float64PtrErr(plusOneInt32Float64PtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Float64PtrErr failed")
	}
	_, err = PMapInt32Float64PtrErr(plusOneInt32Float64PtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Float64PtrErr failed")
	}

	_, err = PMapInt32Float64PtrErr(plusOneInt32Float64PtrErr, []*int32{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Float64PtrErr failed")
	}

	_, err = PMapInt32Float64PtrErr(plusOneInt32Float64PtrErr, []*int32{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Float64PtrErr failed")
	}

	_, err = PMapInt32Float64PtrErr(plusOneInt32Float64PtrErr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt32Float64PtrErr failed")
	}
}

func TestPmapInt16IntPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3

	expectedList := []*int{&vo2, &vo3}
	newList, _ := PMapInt16IntPtrErr(plusOneInt16IntPtrErr, []*int16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt16IntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt16IntPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt16IntPtrErr failed")
	}

	r, _ = PMapInt16IntPtrErr(nil, []*int16{})
	if len(r) > 0 {
		t.Errorf("PMapInt16IntPtrErr failed")
	}
	_, err := PMapInt16IntPtrErr(plusOneInt16IntPtrErr, []*int16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt16IntPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt16IntPtrErr(plusOneInt16IntPtrErr, []*int16{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16IntPtrErr failed")
	}

	_, err = PMapInt16IntPtrErr(plusOneInt16IntPtrErr, []*int16{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt16IntPtrErr failed")
	}

	_, err = PMapInt16IntPtrErr(plusOneInt16IntPtrErr, []*int16{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16IntPtrErr failed")
	}
	_, err = PMapInt16IntPtrErr(plusOneInt16IntPtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16IntPtrErr failed")
	}

	expectedList = []*int{&vo2, &vo3}
	newList, _ = PMapInt16IntPtrErr(plusOneInt16IntPtrErr, []*int16{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt16IntPtr failed")
	}

	_, err = PMapInt16IntPtrErr(plusOneInt16IntPtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16IntPtrErr failed")
	}
	_, err = PMapInt16IntPtrErr(plusOneInt16IntPtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16IntPtrErr failed")
	}

	_, err = PMapInt16IntPtrErr(plusOneInt16IntPtrErr, []*int16{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16IntPtrErr failed")
	}

	_, err = PMapInt16IntPtrErr(plusOneInt16IntPtrErr, []*int16{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16IntPtrErr failed")
	}

	_, err = PMapInt16IntPtrErr(plusOneInt16IntPtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16IntPtrErr failed")
	}
}

func TestPmapInt16Int64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3

	expectedList := []*int64{&vo2, &vo3}
	newList, _ := PMapInt16Int64PtrErr(plusOneInt16Int64PtrErr, []*int16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt16Int64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt16Int64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt16Int64PtrErr failed")
	}

	r, _ = PMapInt16Int64PtrErr(nil, []*int16{})
	if len(r) > 0 {
		t.Errorf("PMapInt16Int64PtrErr failed")
	}
	_, err := PMapInt16Int64PtrErr(plusOneInt16Int64PtrErr, []*int16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt16Int64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt16Int64PtrErr(plusOneInt16Int64PtrErr, []*int16{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16Int64PtrErr failed")
	}

	_, err = PMapInt16Int64PtrErr(plusOneInt16Int64PtrErr, []*int16{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt16Int64PtrErr failed")
	}

	_, err = PMapInt16Int64PtrErr(plusOneInt16Int64PtrErr, []*int16{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16Int64PtrErr failed")
	}
	_, err = PMapInt16Int64PtrErr(plusOneInt16Int64PtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16Int64PtrErr failed")
	}

	expectedList = []*int64{&vo2, &vo3}
	newList, _ = PMapInt16Int64PtrErr(plusOneInt16Int64PtrErr, []*int16{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt16Int64Ptr failed")
	}

	_, err = PMapInt16Int64PtrErr(plusOneInt16Int64PtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Int64PtrErr failed")
	}
	_, err = PMapInt16Int64PtrErr(plusOneInt16Int64PtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Int64PtrErr failed")
	}

	_, err = PMapInt16Int64PtrErr(plusOneInt16Int64PtrErr, []*int16{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Int64PtrErr failed")
	}

	_, err = PMapInt16Int64PtrErr(plusOneInt16Int64PtrErr, []*int16{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Int64PtrErr failed")
	}

	_, err = PMapInt16Int64PtrErr(plusOneInt16Int64PtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Int64PtrErr failed")
	}
}

func TestPmapInt16Int32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3

	expectedList := []*int32{&vo2, &vo3}
	newList, _ := PMapInt16Int32PtrErr(plusOneInt16Int32PtrErr, []*int16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt16Int32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt16Int32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt16Int32PtrErr failed")
	}

	r, _ = PMapInt16Int32PtrErr(nil, []*int16{})
	if len(r) > 0 {
		t.Errorf("PMapInt16Int32PtrErr failed")
	}
	_, err := PMapInt16Int32PtrErr(plusOneInt16Int32PtrErr, []*int16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt16Int32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt16Int32PtrErr(plusOneInt16Int32PtrErr, []*int16{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16Int32PtrErr failed")
	}

	_, err = PMapInt16Int32PtrErr(plusOneInt16Int32PtrErr, []*int16{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt16Int32PtrErr failed")
	}

	_, err = PMapInt16Int32PtrErr(plusOneInt16Int32PtrErr, []*int16{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16Int32PtrErr failed")
	}
	_, err = PMapInt16Int32PtrErr(plusOneInt16Int32PtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16Int32PtrErr failed")
	}

	expectedList = []*int32{&vo2, &vo3}
	newList, _ = PMapInt16Int32PtrErr(plusOneInt16Int32PtrErr, []*int16{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt16Int32Ptr failed")
	}

	_, err = PMapInt16Int32PtrErr(plusOneInt16Int32PtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Int32PtrErr failed")
	}
	_, err = PMapInt16Int32PtrErr(plusOneInt16Int32PtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Int32PtrErr failed")
	}

	_, err = PMapInt16Int32PtrErr(plusOneInt16Int32PtrErr, []*int16{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Int32PtrErr failed")
	}

	_, err = PMapInt16Int32PtrErr(plusOneInt16Int32PtrErr, []*int16{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Int32PtrErr failed")
	}

	_, err = PMapInt16Int32PtrErr(plusOneInt16Int32PtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Int32PtrErr failed")
	}
}

func TestPmapInt16Int8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3

	expectedList := []*int8{&vo2, &vo3}
	newList, _ := PMapInt16Int8PtrErr(plusOneInt16Int8PtrErr, []*int16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt16Int8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt16Int8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt16Int8PtrErr failed")
	}

	r, _ = PMapInt16Int8PtrErr(nil, []*int16{})
	if len(r) > 0 {
		t.Errorf("PMapInt16Int8PtrErr failed")
	}
	_, err := PMapInt16Int8PtrErr(plusOneInt16Int8PtrErr, []*int16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt16Int8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt16Int8PtrErr(plusOneInt16Int8PtrErr, []*int16{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16Int8PtrErr failed")
	}

	_, err = PMapInt16Int8PtrErr(plusOneInt16Int8PtrErr, []*int16{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt16Int8PtrErr failed")
	}

	_, err = PMapInt16Int8PtrErr(plusOneInt16Int8PtrErr, []*int16{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16Int8PtrErr failed")
	}
	_, err = PMapInt16Int8PtrErr(plusOneInt16Int8PtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16Int8PtrErr failed")
	}

	expectedList = []*int8{&vo2, &vo3}
	newList, _ = PMapInt16Int8PtrErr(plusOneInt16Int8PtrErr, []*int16{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt16Int8Ptr failed")
	}

	_, err = PMapInt16Int8PtrErr(plusOneInt16Int8PtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Int8PtrErr failed")
	}
	_, err = PMapInt16Int8PtrErr(plusOneInt16Int8PtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Int8PtrErr failed")
	}

	_, err = PMapInt16Int8PtrErr(plusOneInt16Int8PtrErr, []*int16{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Int8PtrErr failed")
	}

	_, err = PMapInt16Int8PtrErr(plusOneInt16Int8PtrErr, []*int16{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Int8PtrErr failed")
	}

	_, err = PMapInt16Int8PtrErr(plusOneInt16Int8PtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Int8PtrErr failed")
	}
}

func TestPmapInt16UintPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3

	expectedList := []*uint{&vo2, &vo3}
	newList, _ := PMapInt16UintPtrErr(plusOneInt16UintPtrErr, []*int16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt16UintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt16UintPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt16UintPtrErr failed")
	}

	r, _ = PMapInt16UintPtrErr(nil, []*int16{})
	if len(r) > 0 {
		t.Errorf("PMapInt16UintPtrErr failed")
	}
	_, err := PMapInt16UintPtrErr(plusOneInt16UintPtrErr, []*int16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt16UintPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt16UintPtrErr(plusOneInt16UintPtrErr, []*int16{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16UintPtrErr failed")
	}

	_, err = PMapInt16UintPtrErr(plusOneInt16UintPtrErr, []*int16{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt16UintPtrErr failed")
	}

	_, err = PMapInt16UintPtrErr(plusOneInt16UintPtrErr, []*int16{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16UintPtrErr failed")
	}
	_, err = PMapInt16UintPtrErr(plusOneInt16UintPtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16UintPtrErr failed")
	}

	expectedList = []*uint{&vo2, &vo3}
	newList, _ = PMapInt16UintPtrErr(plusOneInt16UintPtrErr, []*int16{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt16UintPtr failed")
	}

	_, err = PMapInt16UintPtrErr(plusOneInt16UintPtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16UintPtrErr failed")
	}
	_, err = PMapInt16UintPtrErr(plusOneInt16UintPtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16UintPtrErr failed")
	}

	_, err = PMapInt16UintPtrErr(plusOneInt16UintPtrErr, []*int16{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16UintPtrErr failed")
	}

	_, err = PMapInt16UintPtrErr(plusOneInt16UintPtrErr, []*int16{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16UintPtrErr failed")
	}

	_, err = PMapInt16UintPtrErr(plusOneInt16UintPtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16UintPtrErr failed")
	}
}

func TestPmapInt16Uint64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3

	expectedList := []*uint64{&vo2, &vo3}
	newList, _ := PMapInt16Uint64PtrErr(plusOneInt16Uint64PtrErr, []*int16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt16Uint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt16Uint64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt16Uint64PtrErr failed")
	}

	r, _ = PMapInt16Uint64PtrErr(nil, []*int16{})
	if len(r) > 0 {
		t.Errorf("PMapInt16Uint64PtrErr failed")
	}
	_, err := PMapInt16Uint64PtrErr(plusOneInt16Uint64PtrErr, []*int16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt16Uint64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt16Uint64PtrErr(plusOneInt16Uint64PtrErr, []*int16{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16Uint64PtrErr failed")
	}

	_, err = PMapInt16Uint64PtrErr(plusOneInt16Uint64PtrErr, []*int16{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt16Uint64PtrErr failed")
	}

	_, err = PMapInt16Uint64PtrErr(plusOneInt16Uint64PtrErr, []*int16{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16Uint64PtrErr failed")
	}
	_, err = PMapInt16Uint64PtrErr(plusOneInt16Uint64PtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16Uint64PtrErr failed")
	}

	expectedList = []*uint64{&vo2, &vo3}
	newList, _ = PMapInt16Uint64PtrErr(plusOneInt16Uint64PtrErr, []*int16{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt16Uint64Ptr failed")
	}

	_, err = PMapInt16Uint64PtrErr(plusOneInt16Uint64PtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Uint64PtrErr failed")
	}
	_, err = PMapInt16Uint64PtrErr(plusOneInt16Uint64PtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Uint64PtrErr failed")
	}

	_, err = PMapInt16Uint64PtrErr(plusOneInt16Uint64PtrErr, []*int16{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Uint64PtrErr failed")
	}

	_, err = PMapInt16Uint64PtrErr(plusOneInt16Uint64PtrErr, []*int16{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Uint64PtrErr failed")
	}

	_, err = PMapInt16Uint64PtrErr(plusOneInt16Uint64PtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Uint64PtrErr failed")
	}
}

func TestPmapInt16Uint32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3

	expectedList := []*uint32{&vo2, &vo3}
	newList, _ := PMapInt16Uint32PtrErr(plusOneInt16Uint32PtrErr, []*int16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt16Uint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt16Uint32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt16Uint32PtrErr failed")
	}

	r, _ = PMapInt16Uint32PtrErr(nil, []*int16{})
	if len(r) > 0 {
		t.Errorf("PMapInt16Uint32PtrErr failed")
	}
	_, err := PMapInt16Uint32PtrErr(plusOneInt16Uint32PtrErr, []*int16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt16Uint32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt16Uint32PtrErr(plusOneInt16Uint32PtrErr, []*int16{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16Uint32PtrErr failed")
	}

	_, err = PMapInt16Uint32PtrErr(plusOneInt16Uint32PtrErr, []*int16{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt16Uint32PtrErr failed")
	}

	_, err = PMapInt16Uint32PtrErr(plusOneInt16Uint32PtrErr, []*int16{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16Uint32PtrErr failed")
	}
	_, err = PMapInt16Uint32PtrErr(plusOneInt16Uint32PtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16Uint32PtrErr failed")
	}

	expectedList = []*uint32{&vo2, &vo3}
	newList, _ = PMapInt16Uint32PtrErr(plusOneInt16Uint32PtrErr, []*int16{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt16Uint32Ptr failed")
	}

	_, err = PMapInt16Uint32PtrErr(plusOneInt16Uint32PtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Uint32PtrErr failed")
	}
	_, err = PMapInt16Uint32PtrErr(plusOneInt16Uint32PtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Uint32PtrErr failed")
	}

	_, err = PMapInt16Uint32PtrErr(plusOneInt16Uint32PtrErr, []*int16{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Uint32PtrErr failed")
	}

	_, err = PMapInt16Uint32PtrErr(plusOneInt16Uint32PtrErr, []*int16{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Uint32PtrErr failed")
	}

	_, err = PMapInt16Uint32PtrErr(plusOneInt16Uint32PtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Uint32PtrErr failed")
	}
}

func TestPmapInt16Uint16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3

	expectedList := []*uint16{&vo2, &vo3}
	newList, _ := PMapInt16Uint16PtrErr(plusOneInt16Uint16PtrErr, []*int16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt16Uint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt16Uint16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt16Uint16PtrErr failed")
	}

	r, _ = PMapInt16Uint16PtrErr(nil, []*int16{})
	if len(r) > 0 {
		t.Errorf("PMapInt16Uint16PtrErr failed")
	}
	_, err := PMapInt16Uint16PtrErr(plusOneInt16Uint16PtrErr, []*int16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt16Uint16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt16Uint16PtrErr(plusOneInt16Uint16PtrErr, []*int16{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16Uint16PtrErr failed")
	}

	_, err = PMapInt16Uint16PtrErr(plusOneInt16Uint16PtrErr, []*int16{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt16Uint16PtrErr failed")
	}

	_, err = PMapInt16Uint16PtrErr(plusOneInt16Uint16PtrErr, []*int16{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16Uint16PtrErr failed")
	}
	_, err = PMapInt16Uint16PtrErr(plusOneInt16Uint16PtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16Uint16PtrErr failed")
	}

	expectedList = []*uint16{&vo2, &vo3}
	newList, _ = PMapInt16Uint16PtrErr(plusOneInt16Uint16PtrErr, []*int16{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt16Uint16Ptr failed")
	}

	_, err = PMapInt16Uint16PtrErr(plusOneInt16Uint16PtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Uint16PtrErr failed")
	}
	_, err = PMapInt16Uint16PtrErr(plusOneInt16Uint16PtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Uint16PtrErr failed")
	}

	_, err = PMapInt16Uint16PtrErr(plusOneInt16Uint16PtrErr, []*int16{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Uint16PtrErr failed")
	}

	_, err = PMapInt16Uint16PtrErr(plusOneInt16Uint16PtrErr, []*int16{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Uint16PtrErr failed")
	}

	_, err = PMapInt16Uint16PtrErr(plusOneInt16Uint16PtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Uint16PtrErr failed")
	}
}

func TestPmapInt16Uint8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3

	expectedList := []*uint8{&vo2, &vo3}
	newList, _ := PMapInt16Uint8PtrErr(plusOneInt16Uint8PtrErr, []*int16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt16Uint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt16Uint8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt16Uint8PtrErr failed")
	}

	r, _ = PMapInt16Uint8PtrErr(nil, []*int16{})
	if len(r) > 0 {
		t.Errorf("PMapInt16Uint8PtrErr failed")
	}
	_, err := PMapInt16Uint8PtrErr(plusOneInt16Uint8PtrErr, []*int16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt16Uint8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt16Uint8PtrErr(plusOneInt16Uint8PtrErr, []*int16{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16Uint8PtrErr failed")
	}

	_, err = PMapInt16Uint8PtrErr(plusOneInt16Uint8PtrErr, []*int16{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt16Uint8PtrErr failed")
	}

	_, err = PMapInt16Uint8PtrErr(plusOneInt16Uint8PtrErr, []*int16{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16Uint8PtrErr failed")
	}
	_, err = PMapInt16Uint8PtrErr(plusOneInt16Uint8PtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16Uint8PtrErr failed")
	}

	expectedList = []*uint8{&vo2, &vo3}
	newList, _ = PMapInt16Uint8PtrErr(plusOneInt16Uint8PtrErr, []*int16{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt16Uint8Ptr failed")
	}

	_, err = PMapInt16Uint8PtrErr(plusOneInt16Uint8PtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Uint8PtrErr failed")
	}
	_, err = PMapInt16Uint8PtrErr(plusOneInt16Uint8PtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Uint8PtrErr failed")
	}

	_, err = PMapInt16Uint8PtrErr(plusOneInt16Uint8PtrErr, []*int16{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Uint8PtrErr failed")
	}

	_, err = PMapInt16Uint8PtrErr(plusOneInt16Uint8PtrErr, []*int16{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Uint8PtrErr failed")
	}

	_, err = PMapInt16Uint8PtrErr(plusOneInt16Uint8PtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Uint8PtrErr failed")
	}
}

func TestPmapInt16StrPtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 int16 = 10
	var vi0 int16

	expectedList := []*string{&vo10}
	newList, _ := PMapInt16StrPtrErr(someLogicInt16StrPtrErr, []*int16{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapInt16StrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt16StrPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt16StrPtrErr failed")
	}

	r, _ = PMapInt16StrPtrErr(nil, []*int16{})
	if len(r) > 0 {
		t.Errorf("PMapInt16StrPtrErr failed")
	}
	_, err := PMapInt16StrPtrErr(someLogicInt16StrPtrErr, []*int16{&vi10, &vi0})
	if err == nil {
		t.Errorf("PMapInt16StrPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt16StrPtrErr(someLogicInt16StrPtrErr, []*int16{&vi10, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16StrPtrErr failed")
	}

	_, err = PMapInt16StrPtrErr(someLogicInt16StrPtrErr, []*int16{&vi0, &vi0, &vi10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt16StrPtrErr failed")
	}

	_, err = PMapInt16StrPtrErr(someLogicInt16StrPtrErr, []*int16{&vi0, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16StrPtrErr failed")
	}
	_, err = PMapInt16StrPtrErr(someLogicInt16StrPtrErr, []*int16{&vi10, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16StrPtrErr failed")
	}

	expectedList = []*string{&vo10, &vo10}
	newList, _ = PMapInt16StrPtrErr(someLogicInt16StrPtrErr, []*int16{&vi10, &vi10}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt16StrPtr failed")
	}

	_, err = PMapInt16StrPtrErr(someLogicInt16StrPtrErr, []*int16{&vi10, &vi10, &vi0}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16StrPtrErr failed")
	}
	_, err = PMapInt16StrPtrErr(someLogicInt16StrPtrErr, []*int16{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16StrPtrErr failed")
	}

	_, err = PMapInt16StrPtrErr(someLogicInt16StrPtrErr, []*int16{&vi0, &vi0, &vi10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16StrPtrErr failed")
	}

	_, err = PMapInt16StrPtrErr(someLogicInt16StrPtrErr, []*int16{&vi0, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16StrPtrErr failed")
	}

	_, err = PMapInt16StrPtrErr(someLogicInt16StrPtrErr, []*int16{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16StrPtrErr failed")
	}
}

func TestPmapInt16BoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 int16 = 10
	var v0 int16
	var v3 int16 = 3

	expectedList := []*bool{&vt, &vf}
	newList, _ := PMapInt16BoolPtrErr(someLogicInt16BoolPtrErr, []*int16{&v10, &v0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt16BoolPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt16BoolPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt16BoolPtrErr failed")
	}

	r, _ = PMapInt16BoolPtrErr(nil, []*int16{})
	if len(r) > 0 {
		t.Errorf("PMapInt16BoolPtrErr failed")
	}
	_, err := PMapInt16BoolPtrErr(someLogicInt16BoolPtrErr, []*int16{&v10, &v3})
	if err == nil {
		t.Errorf("PMapInt16BoolPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt16BoolPtrErr(someLogicInt16BoolPtrErr, []*int16{&v10, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16BoolPtrErr failed")
	}

	_, err = PMapInt16BoolPtrErr(someLogicInt16BoolPtrErr, []*int16{&v3, &v3, &v10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt16BoolPtrErr failed")
	}

	_, err = PMapInt16BoolPtrErr(someLogicInt16BoolPtrErr, []*int16{&v3, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16BoolErr failed")
	}
	_, err = PMapInt16BoolPtrErr(someLogicInt16BoolPtrErr, []*int16{&v10, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16BoolPtrErr failed")
	}

	expectedList = []*bool{&vt, &vt}
	newList, _ = PMapInt16BoolPtrErr(someLogicInt16BoolPtrErr, []*int16{&v10, &v10}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt16BoolPtr failed")
	}

	_, err = PMapInt16BoolPtrErr(someLogicInt16BoolPtrErr, []*int16{&v10, &v10, &v3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16BoolErr failed")
	}
	_, err = PMapInt16BoolPtrErr(someLogicInt16BoolPtrErr, []*int16{&v10, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16BoolPtrErr failed")
	}

	_, err = PMapInt16BoolPtrErr(someLogicInt16BoolPtrErr, []*int16{&v3, &v3, &v10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16BoolPtrErr failed")
	}

	_, err = PMapInt16BoolPtrErr(someLogicInt16BoolPtrErr, []*int16{&v3, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16BoolPtrErr failed")
	}

	_, err = PMapInt16BoolPtrErr(someLogicInt16BoolPtrErr, []*int16{&v10, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16BoolPtrErr failed")
	}
}

func TestPmapInt16Float32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3

	expectedList := []*float32{&vo2, &vo3}
	newList, _ := PMapInt16Float32PtrErr(plusOneInt16Float32PtrErr, []*int16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt16Float32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt16Float32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt16Float32PtrErr failed")
	}

	r, _ = PMapInt16Float32PtrErr(nil, []*int16{})
	if len(r) > 0 {
		t.Errorf("PMapInt16Float32PtrErr failed")
	}
	_, err := PMapInt16Float32PtrErr(plusOneInt16Float32PtrErr, []*int16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt16Float32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt16Float32PtrErr(plusOneInt16Float32PtrErr, []*int16{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16Float32PtrErr failed")
	}

	_, err = PMapInt16Float32PtrErr(plusOneInt16Float32PtrErr, []*int16{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt16Float32PtrErr failed")
	}

	_, err = PMapInt16Float32PtrErr(plusOneInt16Float32PtrErr, []*int16{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16Float32PtrErr failed")
	}
	_, err = PMapInt16Float32PtrErr(plusOneInt16Float32PtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16Float32PtrErr failed")
	}

	expectedList = []*float32{&vo2, &vo3}
	newList, _ = PMapInt16Float32PtrErr(plusOneInt16Float32PtrErr, []*int16{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt16Float32Ptr failed")
	}

	_, err = PMapInt16Float32PtrErr(plusOneInt16Float32PtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Float32PtrErr failed")
	}
	_, err = PMapInt16Float32PtrErr(plusOneInt16Float32PtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Float32PtrErr failed")
	}

	_, err = PMapInt16Float32PtrErr(plusOneInt16Float32PtrErr, []*int16{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Float32PtrErr failed")
	}

	_, err = PMapInt16Float32PtrErr(plusOneInt16Float32PtrErr, []*int16{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Float32PtrErr failed")
	}

	_, err = PMapInt16Float32PtrErr(plusOneInt16Float32PtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Float32PtrErr failed")
	}
}

func TestPmapInt16Float64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3

	expectedList := []*float64{&vo2, &vo3}
	newList, _ := PMapInt16Float64PtrErr(plusOneInt16Float64PtrErr, []*int16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt16Float64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt16Float64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt16Float64PtrErr failed")
	}

	r, _ = PMapInt16Float64PtrErr(nil, []*int16{})
	if len(r) > 0 {
		t.Errorf("PMapInt16Float64PtrErr failed")
	}
	_, err := PMapInt16Float64PtrErr(plusOneInt16Float64PtrErr, []*int16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt16Float64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt16Float64PtrErr(plusOneInt16Float64PtrErr, []*int16{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16Float64PtrErr failed")
	}

	_, err = PMapInt16Float64PtrErr(plusOneInt16Float64PtrErr, []*int16{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt16Float64PtrErr failed")
	}

	_, err = PMapInt16Float64PtrErr(plusOneInt16Float64PtrErr, []*int16{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16Float64PtrErr failed")
	}
	_, err = PMapInt16Float64PtrErr(plusOneInt16Float64PtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16Float64PtrErr failed")
	}

	expectedList = []*float64{&vo2, &vo3}
	newList, _ = PMapInt16Float64PtrErr(plusOneInt16Float64PtrErr, []*int16{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt16Float64Ptr failed")
	}

	_, err = PMapInt16Float64PtrErr(plusOneInt16Float64PtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Float64PtrErr failed")
	}
	_, err = PMapInt16Float64PtrErr(plusOneInt16Float64PtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Float64PtrErr failed")
	}

	_, err = PMapInt16Float64PtrErr(plusOneInt16Float64PtrErr, []*int16{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Float64PtrErr failed")
	}

	_, err = PMapInt16Float64PtrErr(plusOneInt16Float64PtrErr, []*int16{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Float64PtrErr failed")
	}

	_, err = PMapInt16Float64PtrErr(plusOneInt16Float64PtrErr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt16Float64PtrErr failed")
	}
}

func TestPmapInt8IntPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3

	expectedList := []*int{&vo2, &vo3}
	newList, _ := PMapInt8IntPtrErr(plusOneInt8IntPtrErr, []*int8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt8IntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt8IntPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt8IntPtrErr failed")
	}

	r, _ = PMapInt8IntPtrErr(nil, []*int8{})
	if len(r) > 0 {
		t.Errorf("PMapInt8IntPtrErr failed")
	}
	_, err := PMapInt8IntPtrErr(plusOneInt8IntPtrErr, []*int8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt8IntPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt8IntPtrErr(plusOneInt8IntPtrErr, []*int8{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8IntPtrErr failed")
	}

	_, err = PMapInt8IntPtrErr(plusOneInt8IntPtrErr, []*int8{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt8IntPtrErr failed")
	}

	_, err = PMapInt8IntPtrErr(plusOneInt8IntPtrErr, []*int8{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8IntPtrErr failed")
	}
	_, err = PMapInt8IntPtrErr(plusOneInt8IntPtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8IntPtrErr failed")
	}

	expectedList = []*int{&vo2, &vo3}
	newList, _ = PMapInt8IntPtrErr(plusOneInt8IntPtrErr, []*int8{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt8IntPtr failed")
	}

	_, err = PMapInt8IntPtrErr(plusOneInt8IntPtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8IntPtrErr failed")
	}
	_, err = PMapInt8IntPtrErr(plusOneInt8IntPtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8IntPtrErr failed")
	}

	_, err = PMapInt8IntPtrErr(plusOneInt8IntPtrErr, []*int8{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8IntPtrErr failed")
	}

	_, err = PMapInt8IntPtrErr(plusOneInt8IntPtrErr, []*int8{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8IntPtrErr failed")
	}

	_, err = PMapInt8IntPtrErr(plusOneInt8IntPtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8IntPtrErr failed")
	}
}

func TestPmapInt8Int64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3

	expectedList := []*int64{&vo2, &vo3}
	newList, _ := PMapInt8Int64PtrErr(plusOneInt8Int64PtrErr, []*int8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt8Int64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt8Int64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt8Int64PtrErr failed")
	}

	r, _ = PMapInt8Int64PtrErr(nil, []*int8{})
	if len(r) > 0 {
		t.Errorf("PMapInt8Int64PtrErr failed")
	}
	_, err := PMapInt8Int64PtrErr(plusOneInt8Int64PtrErr, []*int8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt8Int64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt8Int64PtrErr(plusOneInt8Int64PtrErr, []*int8{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8Int64PtrErr failed")
	}

	_, err = PMapInt8Int64PtrErr(plusOneInt8Int64PtrErr, []*int8{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt8Int64PtrErr failed")
	}

	_, err = PMapInt8Int64PtrErr(plusOneInt8Int64PtrErr, []*int8{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8Int64PtrErr failed")
	}
	_, err = PMapInt8Int64PtrErr(plusOneInt8Int64PtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8Int64PtrErr failed")
	}

	expectedList = []*int64{&vo2, &vo3}
	newList, _ = PMapInt8Int64PtrErr(plusOneInt8Int64PtrErr, []*int8{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt8Int64Ptr failed")
	}

	_, err = PMapInt8Int64PtrErr(plusOneInt8Int64PtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Int64PtrErr failed")
	}
	_, err = PMapInt8Int64PtrErr(plusOneInt8Int64PtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Int64PtrErr failed")
	}

	_, err = PMapInt8Int64PtrErr(plusOneInt8Int64PtrErr, []*int8{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Int64PtrErr failed")
	}

	_, err = PMapInt8Int64PtrErr(plusOneInt8Int64PtrErr, []*int8{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Int64PtrErr failed")
	}

	_, err = PMapInt8Int64PtrErr(plusOneInt8Int64PtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Int64PtrErr failed")
	}
}

func TestPmapInt8Int32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3

	expectedList := []*int32{&vo2, &vo3}
	newList, _ := PMapInt8Int32PtrErr(plusOneInt8Int32PtrErr, []*int8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt8Int32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt8Int32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt8Int32PtrErr failed")
	}

	r, _ = PMapInt8Int32PtrErr(nil, []*int8{})
	if len(r) > 0 {
		t.Errorf("PMapInt8Int32PtrErr failed")
	}
	_, err := PMapInt8Int32PtrErr(plusOneInt8Int32PtrErr, []*int8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt8Int32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt8Int32PtrErr(plusOneInt8Int32PtrErr, []*int8{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8Int32PtrErr failed")
	}

	_, err = PMapInt8Int32PtrErr(plusOneInt8Int32PtrErr, []*int8{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt8Int32PtrErr failed")
	}

	_, err = PMapInt8Int32PtrErr(plusOneInt8Int32PtrErr, []*int8{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8Int32PtrErr failed")
	}
	_, err = PMapInt8Int32PtrErr(plusOneInt8Int32PtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8Int32PtrErr failed")
	}

	expectedList = []*int32{&vo2, &vo3}
	newList, _ = PMapInt8Int32PtrErr(plusOneInt8Int32PtrErr, []*int8{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt8Int32Ptr failed")
	}

	_, err = PMapInt8Int32PtrErr(plusOneInt8Int32PtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Int32PtrErr failed")
	}
	_, err = PMapInt8Int32PtrErr(plusOneInt8Int32PtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Int32PtrErr failed")
	}

	_, err = PMapInt8Int32PtrErr(plusOneInt8Int32PtrErr, []*int8{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Int32PtrErr failed")
	}

	_, err = PMapInt8Int32PtrErr(plusOneInt8Int32PtrErr, []*int8{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Int32PtrErr failed")
	}

	_, err = PMapInt8Int32PtrErr(plusOneInt8Int32PtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Int32PtrErr failed")
	}
}

func TestPmapInt8Int16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3

	expectedList := []*int16{&vo2, &vo3}
	newList, _ := PMapInt8Int16PtrErr(plusOneInt8Int16PtrErr, []*int8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt8Int16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt8Int16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt8Int16PtrErr failed")
	}

	r, _ = PMapInt8Int16PtrErr(nil, []*int8{})
	if len(r) > 0 {
		t.Errorf("PMapInt8Int16PtrErr failed")
	}
	_, err := PMapInt8Int16PtrErr(plusOneInt8Int16PtrErr, []*int8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt8Int16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt8Int16PtrErr(plusOneInt8Int16PtrErr, []*int8{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8Int16PtrErr failed")
	}

	_, err = PMapInt8Int16PtrErr(plusOneInt8Int16PtrErr, []*int8{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt8Int16PtrErr failed")
	}

	_, err = PMapInt8Int16PtrErr(plusOneInt8Int16PtrErr, []*int8{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8Int16PtrErr failed")
	}
	_, err = PMapInt8Int16PtrErr(plusOneInt8Int16PtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8Int16PtrErr failed")
	}

	expectedList = []*int16{&vo2, &vo3}
	newList, _ = PMapInt8Int16PtrErr(plusOneInt8Int16PtrErr, []*int8{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt8Int16Ptr failed")
	}

	_, err = PMapInt8Int16PtrErr(plusOneInt8Int16PtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Int16PtrErr failed")
	}
	_, err = PMapInt8Int16PtrErr(plusOneInt8Int16PtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Int16PtrErr failed")
	}

	_, err = PMapInt8Int16PtrErr(plusOneInt8Int16PtrErr, []*int8{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Int16PtrErr failed")
	}

	_, err = PMapInt8Int16PtrErr(plusOneInt8Int16PtrErr, []*int8{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Int16PtrErr failed")
	}

	_, err = PMapInt8Int16PtrErr(plusOneInt8Int16PtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Int16PtrErr failed")
	}
}

func TestPmapInt8UintPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3

	expectedList := []*uint{&vo2, &vo3}
	newList, _ := PMapInt8UintPtrErr(plusOneInt8UintPtrErr, []*int8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt8UintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt8UintPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt8UintPtrErr failed")
	}

	r, _ = PMapInt8UintPtrErr(nil, []*int8{})
	if len(r) > 0 {
		t.Errorf("PMapInt8UintPtrErr failed")
	}
	_, err := PMapInt8UintPtrErr(plusOneInt8UintPtrErr, []*int8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt8UintPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt8UintPtrErr(plusOneInt8UintPtrErr, []*int8{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8UintPtrErr failed")
	}

	_, err = PMapInt8UintPtrErr(plusOneInt8UintPtrErr, []*int8{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt8UintPtrErr failed")
	}

	_, err = PMapInt8UintPtrErr(plusOneInt8UintPtrErr, []*int8{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8UintPtrErr failed")
	}
	_, err = PMapInt8UintPtrErr(plusOneInt8UintPtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8UintPtrErr failed")
	}

	expectedList = []*uint{&vo2, &vo3}
	newList, _ = PMapInt8UintPtrErr(plusOneInt8UintPtrErr, []*int8{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt8UintPtr failed")
	}

	_, err = PMapInt8UintPtrErr(plusOneInt8UintPtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8UintPtrErr failed")
	}
	_, err = PMapInt8UintPtrErr(plusOneInt8UintPtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8UintPtrErr failed")
	}

	_, err = PMapInt8UintPtrErr(plusOneInt8UintPtrErr, []*int8{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8UintPtrErr failed")
	}

	_, err = PMapInt8UintPtrErr(plusOneInt8UintPtrErr, []*int8{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8UintPtrErr failed")
	}

	_, err = PMapInt8UintPtrErr(plusOneInt8UintPtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8UintPtrErr failed")
	}
}

func TestPmapInt8Uint64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3

	expectedList := []*uint64{&vo2, &vo3}
	newList, _ := PMapInt8Uint64PtrErr(plusOneInt8Uint64PtrErr, []*int8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt8Uint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt8Uint64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt8Uint64PtrErr failed")
	}

	r, _ = PMapInt8Uint64PtrErr(nil, []*int8{})
	if len(r) > 0 {
		t.Errorf("PMapInt8Uint64PtrErr failed")
	}
	_, err := PMapInt8Uint64PtrErr(plusOneInt8Uint64PtrErr, []*int8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt8Uint64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt8Uint64PtrErr(plusOneInt8Uint64PtrErr, []*int8{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8Uint64PtrErr failed")
	}

	_, err = PMapInt8Uint64PtrErr(plusOneInt8Uint64PtrErr, []*int8{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt8Uint64PtrErr failed")
	}

	_, err = PMapInt8Uint64PtrErr(plusOneInt8Uint64PtrErr, []*int8{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8Uint64PtrErr failed")
	}
	_, err = PMapInt8Uint64PtrErr(plusOneInt8Uint64PtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8Uint64PtrErr failed")
	}

	expectedList = []*uint64{&vo2, &vo3}
	newList, _ = PMapInt8Uint64PtrErr(plusOneInt8Uint64PtrErr, []*int8{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt8Uint64Ptr failed")
	}

	_, err = PMapInt8Uint64PtrErr(plusOneInt8Uint64PtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Uint64PtrErr failed")
	}
	_, err = PMapInt8Uint64PtrErr(plusOneInt8Uint64PtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Uint64PtrErr failed")
	}

	_, err = PMapInt8Uint64PtrErr(plusOneInt8Uint64PtrErr, []*int8{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Uint64PtrErr failed")
	}

	_, err = PMapInt8Uint64PtrErr(plusOneInt8Uint64PtrErr, []*int8{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Uint64PtrErr failed")
	}

	_, err = PMapInt8Uint64PtrErr(plusOneInt8Uint64PtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Uint64PtrErr failed")
	}
}

func TestPmapInt8Uint32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3

	expectedList := []*uint32{&vo2, &vo3}
	newList, _ := PMapInt8Uint32PtrErr(plusOneInt8Uint32PtrErr, []*int8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt8Uint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt8Uint32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt8Uint32PtrErr failed")
	}

	r, _ = PMapInt8Uint32PtrErr(nil, []*int8{})
	if len(r) > 0 {
		t.Errorf("PMapInt8Uint32PtrErr failed")
	}
	_, err := PMapInt8Uint32PtrErr(plusOneInt8Uint32PtrErr, []*int8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt8Uint32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt8Uint32PtrErr(plusOneInt8Uint32PtrErr, []*int8{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8Uint32PtrErr failed")
	}

	_, err = PMapInt8Uint32PtrErr(plusOneInt8Uint32PtrErr, []*int8{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt8Uint32PtrErr failed")
	}

	_, err = PMapInt8Uint32PtrErr(plusOneInt8Uint32PtrErr, []*int8{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8Uint32PtrErr failed")
	}
	_, err = PMapInt8Uint32PtrErr(plusOneInt8Uint32PtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8Uint32PtrErr failed")
	}

	expectedList = []*uint32{&vo2, &vo3}
	newList, _ = PMapInt8Uint32PtrErr(plusOneInt8Uint32PtrErr, []*int8{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt8Uint32Ptr failed")
	}

	_, err = PMapInt8Uint32PtrErr(plusOneInt8Uint32PtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Uint32PtrErr failed")
	}
	_, err = PMapInt8Uint32PtrErr(plusOneInt8Uint32PtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Uint32PtrErr failed")
	}

	_, err = PMapInt8Uint32PtrErr(plusOneInt8Uint32PtrErr, []*int8{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Uint32PtrErr failed")
	}

	_, err = PMapInt8Uint32PtrErr(plusOneInt8Uint32PtrErr, []*int8{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Uint32PtrErr failed")
	}

	_, err = PMapInt8Uint32PtrErr(plusOneInt8Uint32PtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Uint32PtrErr failed")
	}
}

func TestPmapInt8Uint16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3

	expectedList := []*uint16{&vo2, &vo3}
	newList, _ := PMapInt8Uint16PtrErr(plusOneInt8Uint16PtrErr, []*int8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt8Uint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt8Uint16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt8Uint16PtrErr failed")
	}

	r, _ = PMapInt8Uint16PtrErr(nil, []*int8{})
	if len(r) > 0 {
		t.Errorf("PMapInt8Uint16PtrErr failed")
	}
	_, err := PMapInt8Uint16PtrErr(plusOneInt8Uint16PtrErr, []*int8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt8Uint16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt8Uint16PtrErr(plusOneInt8Uint16PtrErr, []*int8{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8Uint16PtrErr failed")
	}

	_, err = PMapInt8Uint16PtrErr(plusOneInt8Uint16PtrErr, []*int8{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt8Uint16PtrErr failed")
	}

	_, err = PMapInt8Uint16PtrErr(plusOneInt8Uint16PtrErr, []*int8{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8Uint16PtrErr failed")
	}
	_, err = PMapInt8Uint16PtrErr(plusOneInt8Uint16PtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8Uint16PtrErr failed")
	}

	expectedList = []*uint16{&vo2, &vo3}
	newList, _ = PMapInt8Uint16PtrErr(plusOneInt8Uint16PtrErr, []*int8{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt8Uint16Ptr failed")
	}

	_, err = PMapInt8Uint16PtrErr(plusOneInt8Uint16PtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Uint16PtrErr failed")
	}
	_, err = PMapInt8Uint16PtrErr(plusOneInt8Uint16PtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Uint16PtrErr failed")
	}

	_, err = PMapInt8Uint16PtrErr(plusOneInt8Uint16PtrErr, []*int8{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Uint16PtrErr failed")
	}

	_, err = PMapInt8Uint16PtrErr(plusOneInt8Uint16PtrErr, []*int8{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Uint16PtrErr failed")
	}

	_, err = PMapInt8Uint16PtrErr(plusOneInt8Uint16PtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Uint16PtrErr failed")
	}
}

func TestPmapInt8Uint8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3

	expectedList := []*uint8{&vo2, &vo3}
	newList, _ := PMapInt8Uint8PtrErr(plusOneInt8Uint8PtrErr, []*int8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt8Uint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt8Uint8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt8Uint8PtrErr failed")
	}

	r, _ = PMapInt8Uint8PtrErr(nil, []*int8{})
	if len(r) > 0 {
		t.Errorf("PMapInt8Uint8PtrErr failed")
	}
	_, err := PMapInt8Uint8PtrErr(plusOneInt8Uint8PtrErr, []*int8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt8Uint8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt8Uint8PtrErr(plusOneInt8Uint8PtrErr, []*int8{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8Uint8PtrErr failed")
	}

	_, err = PMapInt8Uint8PtrErr(plusOneInt8Uint8PtrErr, []*int8{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt8Uint8PtrErr failed")
	}

	_, err = PMapInt8Uint8PtrErr(plusOneInt8Uint8PtrErr, []*int8{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8Uint8PtrErr failed")
	}
	_, err = PMapInt8Uint8PtrErr(plusOneInt8Uint8PtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8Uint8PtrErr failed")
	}

	expectedList = []*uint8{&vo2, &vo3}
	newList, _ = PMapInt8Uint8PtrErr(plusOneInt8Uint8PtrErr, []*int8{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt8Uint8Ptr failed")
	}

	_, err = PMapInt8Uint8PtrErr(plusOneInt8Uint8PtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Uint8PtrErr failed")
	}
	_, err = PMapInt8Uint8PtrErr(plusOneInt8Uint8PtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Uint8PtrErr failed")
	}

	_, err = PMapInt8Uint8PtrErr(plusOneInt8Uint8PtrErr, []*int8{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Uint8PtrErr failed")
	}

	_, err = PMapInt8Uint8PtrErr(plusOneInt8Uint8PtrErr, []*int8{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Uint8PtrErr failed")
	}

	_, err = PMapInt8Uint8PtrErr(plusOneInt8Uint8PtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Uint8PtrErr failed")
	}
}

func TestPmapInt8StrPtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 int8 = 10
	var vi0 int8

	expectedList := []*string{&vo10}
	newList, _ := PMapInt8StrPtrErr(someLogicInt8StrPtrErr, []*int8{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapInt8StrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt8StrPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt8StrPtrErr failed")
	}

	r, _ = PMapInt8StrPtrErr(nil, []*int8{})
	if len(r) > 0 {
		t.Errorf("PMapInt8StrPtrErr failed")
	}
	_, err := PMapInt8StrPtrErr(someLogicInt8StrPtrErr, []*int8{&vi10, &vi0})
	if err == nil {
		t.Errorf("PMapInt8StrPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt8StrPtrErr(someLogicInt8StrPtrErr, []*int8{&vi10, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8StrPtrErr failed")
	}

	_, err = PMapInt8StrPtrErr(someLogicInt8StrPtrErr, []*int8{&vi0, &vi0, &vi10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt8StrPtrErr failed")
	}

	_, err = PMapInt8StrPtrErr(someLogicInt8StrPtrErr, []*int8{&vi0, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8StrPtrErr failed")
	}
	_, err = PMapInt8StrPtrErr(someLogicInt8StrPtrErr, []*int8{&vi10, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8StrPtrErr failed")
	}

	expectedList = []*string{&vo10, &vo10}
	newList, _ = PMapInt8StrPtrErr(someLogicInt8StrPtrErr, []*int8{&vi10, &vi10}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt8StrPtr failed")
	}

	_, err = PMapInt8StrPtrErr(someLogicInt8StrPtrErr, []*int8{&vi10, &vi10, &vi0}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8StrPtrErr failed")
	}
	_, err = PMapInt8StrPtrErr(someLogicInt8StrPtrErr, []*int8{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8StrPtrErr failed")
	}

	_, err = PMapInt8StrPtrErr(someLogicInt8StrPtrErr, []*int8{&vi0, &vi0, &vi10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8StrPtrErr failed")
	}

	_, err = PMapInt8StrPtrErr(someLogicInt8StrPtrErr, []*int8{&vi0, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8StrPtrErr failed")
	}

	_, err = PMapInt8StrPtrErr(someLogicInt8StrPtrErr, []*int8{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8StrPtrErr failed")
	}
}

func TestPmapInt8BoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 int8 = 10
	var v0 int8
	var v3 int8 = 3

	expectedList := []*bool{&vt, &vf}
	newList, _ := PMapInt8BoolPtrErr(someLogicInt8BoolPtrErr, []*int8{&v10, &v0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt8BoolPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt8BoolPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt8BoolPtrErr failed")
	}

	r, _ = PMapInt8BoolPtrErr(nil, []*int8{})
	if len(r) > 0 {
		t.Errorf("PMapInt8BoolPtrErr failed")
	}
	_, err := PMapInt8BoolPtrErr(someLogicInt8BoolPtrErr, []*int8{&v10, &v3})
	if err == nil {
		t.Errorf("PMapInt8BoolPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt8BoolPtrErr(someLogicInt8BoolPtrErr, []*int8{&v10, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8BoolPtrErr failed")
	}

	_, err = PMapInt8BoolPtrErr(someLogicInt8BoolPtrErr, []*int8{&v3, &v3, &v10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt8BoolPtrErr failed")
	}

	_, err = PMapInt8BoolPtrErr(someLogicInt8BoolPtrErr, []*int8{&v3, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8BoolErr failed")
	}
	_, err = PMapInt8BoolPtrErr(someLogicInt8BoolPtrErr, []*int8{&v10, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8BoolPtrErr failed")
	}

	expectedList = []*bool{&vt, &vt}
	newList, _ = PMapInt8BoolPtrErr(someLogicInt8BoolPtrErr, []*int8{&v10, &v10}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt8BoolPtr failed")
	}

	_, err = PMapInt8BoolPtrErr(someLogicInt8BoolPtrErr, []*int8{&v10, &v10, &v3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8BoolErr failed")
	}
	_, err = PMapInt8BoolPtrErr(someLogicInt8BoolPtrErr, []*int8{&v10, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8BoolPtrErr failed")
	}

	_, err = PMapInt8BoolPtrErr(someLogicInt8BoolPtrErr, []*int8{&v3, &v3, &v10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8BoolPtrErr failed")
	}

	_, err = PMapInt8BoolPtrErr(someLogicInt8BoolPtrErr, []*int8{&v3, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8BoolPtrErr failed")
	}

	_, err = PMapInt8BoolPtrErr(someLogicInt8BoolPtrErr, []*int8{&v10, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8BoolPtrErr failed")
	}
}

func TestPmapInt8Float32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3

	expectedList := []*float32{&vo2, &vo3}
	newList, _ := PMapInt8Float32PtrErr(plusOneInt8Float32PtrErr, []*int8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt8Float32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt8Float32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt8Float32PtrErr failed")
	}

	r, _ = PMapInt8Float32PtrErr(nil, []*int8{})
	if len(r) > 0 {
		t.Errorf("PMapInt8Float32PtrErr failed")
	}
	_, err := PMapInt8Float32PtrErr(plusOneInt8Float32PtrErr, []*int8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt8Float32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt8Float32PtrErr(plusOneInt8Float32PtrErr, []*int8{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8Float32PtrErr failed")
	}

	_, err = PMapInt8Float32PtrErr(plusOneInt8Float32PtrErr, []*int8{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt8Float32PtrErr failed")
	}

	_, err = PMapInt8Float32PtrErr(plusOneInt8Float32PtrErr, []*int8{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8Float32PtrErr failed")
	}
	_, err = PMapInt8Float32PtrErr(plusOneInt8Float32PtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8Float32PtrErr failed")
	}

	expectedList = []*float32{&vo2, &vo3}
	newList, _ = PMapInt8Float32PtrErr(plusOneInt8Float32PtrErr, []*int8{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt8Float32Ptr failed")
	}

	_, err = PMapInt8Float32PtrErr(plusOneInt8Float32PtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Float32PtrErr failed")
	}
	_, err = PMapInt8Float32PtrErr(plusOneInt8Float32PtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Float32PtrErr failed")
	}

	_, err = PMapInt8Float32PtrErr(plusOneInt8Float32PtrErr, []*int8{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Float32PtrErr failed")
	}

	_, err = PMapInt8Float32PtrErr(plusOneInt8Float32PtrErr, []*int8{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Float32PtrErr failed")
	}

	_, err = PMapInt8Float32PtrErr(plusOneInt8Float32PtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Float32PtrErr failed")
	}
}

func TestPmapInt8Float64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3

	expectedList := []*float64{&vo2, &vo3}
	newList, _ := PMapInt8Float64PtrErr(plusOneInt8Float64PtrErr, []*int8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt8Float64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapInt8Float64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt8Float64PtrErr failed")
	}

	r, _ = PMapInt8Float64PtrErr(nil, []*int8{})
	if len(r) > 0 {
		t.Errorf("PMapInt8Float64PtrErr failed")
	}
	_, err := PMapInt8Float64PtrErr(plusOneInt8Float64PtrErr, []*int8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapInt8Float64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapInt8Float64PtrErr(plusOneInt8Float64PtrErr, []*int8{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8Float64PtrErr failed")
	}

	_, err = PMapInt8Float64PtrErr(plusOneInt8Float64PtrErr, []*int8{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapInt8Float64PtrErr failed")
	}

	_, err = PMapInt8Float64PtrErr(plusOneInt8Float64PtrErr, []*int8{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8Float64PtrErr failed")
	}
	_, err = PMapInt8Float64PtrErr(plusOneInt8Float64PtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8Float64PtrErr failed")
	}

	expectedList = []*float64{&vo2, &vo3}
	newList, _ = PMapInt8Float64PtrErr(plusOneInt8Float64PtrErr, []*int8{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapInt8Float64Ptr failed")
	}

	_, err = PMapInt8Float64PtrErr(plusOneInt8Float64PtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Float64PtrErr failed")
	}
	_, err = PMapInt8Float64PtrErr(plusOneInt8Float64PtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Float64PtrErr failed")
	}

	_, err = PMapInt8Float64PtrErr(plusOneInt8Float64PtrErr, []*int8{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Float64PtrErr failed")
	}

	_, err = PMapInt8Float64PtrErr(plusOneInt8Float64PtrErr, []*int8{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Float64PtrErr failed")
	}

	_, err = PMapInt8Float64PtrErr(plusOneInt8Float64PtrErr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapInt8Float64PtrErr failed")
	}
}

func TestPmapUintIntPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3

	expectedList := []*int{&vo2, &vo3}
	newList, _ := PMapUintIntPtrErr(plusOneUintIntPtrErr, []*uint{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUintIntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUintIntPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUintIntPtrErr failed")
	}

	r, _ = PMapUintIntPtrErr(nil, []*uint{})
	if len(r) > 0 {
		t.Errorf("PMapUintIntPtrErr failed")
	}
	_, err := PMapUintIntPtrErr(plusOneUintIntPtrErr, []*uint{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUintIntPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUintIntPtrErr(plusOneUintIntPtrErr, []*uint{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintIntPtrErr failed")
	}

	_, err = PMapUintIntPtrErr(plusOneUintIntPtrErr, []*uint{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUintIntPtrErr failed")
	}

	_, err = PMapUintIntPtrErr(plusOneUintIntPtrErr, []*uint{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintIntPtrErr failed")
	}
	_, err = PMapUintIntPtrErr(plusOneUintIntPtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintIntPtrErr failed")
	}

	expectedList = []*int{&vo2, &vo3}
	newList, _ = PMapUintIntPtrErr(plusOneUintIntPtrErr, []*uint{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUintIntPtr failed")
	}

	_, err = PMapUintIntPtrErr(plusOneUintIntPtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintIntPtrErr failed")
	}
	_, err = PMapUintIntPtrErr(plusOneUintIntPtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintIntPtrErr failed")
	}

	_, err = PMapUintIntPtrErr(plusOneUintIntPtrErr, []*uint{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintIntPtrErr failed")
	}

	_, err = PMapUintIntPtrErr(plusOneUintIntPtrErr, []*uint{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintIntPtrErr failed")
	}

	_, err = PMapUintIntPtrErr(plusOneUintIntPtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintIntPtrErr failed")
	}
}

func TestPmapUintInt64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3

	expectedList := []*int64{&vo2, &vo3}
	newList, _ := PMapUintInt64PtrErr(plusOneUintInt64PtrErr, []*uint{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUintInt64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUintInt64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUintInt64PtrErr failed")
	}

	r, _ = PMapUintInt64PtrErr(nil, []*uint{})
	if len(r) > 0 {
		t.Errorf("PMapUintInt64PtrErr failed")
	}
	_, err := PMapUintInt64PtrErr(plusOneUintInt64PtrErr, []*uint{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUintInt64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUintInt64PtrErr(plusOneUintInt64PtrErr, []*uint{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintInt64PtrErr failed")
	}

	_, err = PMapUintInt64PtrErr(plusOneUintInt64PtrErr, []*uint{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUintInt64PtrErr failed")
	}

	_, err = PMapUintInt64PtrErr(plusOneUintInt64PtrErr, []*uint{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintInt64PtrErr failed")
	}
	_, err = PMapUintInt64PtrErr(plusOneUintInt64PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintInt64PtrErr failed")
	}

	expectedList = []*int64{&vo2, &vo3}
	newList, _ = PMapUintInt64PtrErr(plusOneUintInt64PtrErr, []*uint{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUintInt64Ptr failed")
	}

	_, err = PMapUintInt64PtrErr(plusOneUintInt64PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintInt64PtrErr failed")
	}
	_, err = PMapUintInt64PtrErr(plusOneUintInt64PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintInt64PtrErr failed")
	}

	_, err = PMapUintInt64PtrErr(plusOneUintInt64PtrErr, []*uint{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintInt64PtrErr failed")
	}

	_, err = PMapUintInt64PtrErr(plusOneUintInt64PtrErr, []*uint{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintInt64PtrErr failed")
	}

	_, err = PMapUintInt64PtrErr(plusOneUintInt64PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintInt64PtrErr failed")
	}
}

func TestPmapUintInt32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3

	expectedList := []*int32{&vo2, &vo3}
	newList, _ := PMapUintInt32PtrErr(plusOneUintInt32PtrErr, []*uint{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUintInt32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUintInt32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUintInt32PtrErr failed")
	}

	r, _ = PMapUintInt32PtrErr(nil, []*uint{})
	if len(r) > 0 {
		t.Errorf("PMapUintInt32PtrErr failed")
	}
	_, err := PMapUintInt32PtrErr(plusOneUintInt32PtrErr, []*uint{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUintInt32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUintInt32PtrErr(plusOneUintInt32PtrErr, []*uint{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintInt32PtrErr failed")
	}

	_, err = PMapUintInt32PtrErr(plusOneUintInt32PtrErr, []*uint{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUintInt32PtrErr failed")
	}

	_, err = PMapUintInt32PtrErr(plusOneUintInt32PtrErr, []*uint{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintInt32PtrErr failed")
	}
	_, err = PMapUintInt32PtrErr(plusOneUintInt32PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintInt32PtrErr failed")
	}

	expectedList = []*int32{&vo2, &vo3}
	newList, _ = PMapUintInt32PtrErr(plusOneUintInt32PtrErr, []*uint{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUintInt32Ptr failed")
	}

	_, err = PMapUintInt32PtrErr(plusOneUintInt32PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintInt32PtrErr failed")
	}
	_, err = PMapUintInt32PtrErr(plusOneUintInt32PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintInt32PtrErr failed")
	}

	_, err = PMapUintInt32PtrErr(plusOneUintInt32PtrErr, []*uint{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintInt32PtrErr failed")
	}

	_, err = PMapUintInt32PtrErr(plusOneUintInt32PtrErr, []*uint{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintInt32PtrErr failed")
	}

	_, err = PMapUintInt32PtrErr(plusOneUintInt32PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintInt32PtrErr failed")
	}
}

func TestPmapUintInt16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3

	expectedList := []*int16{&vo2, &vo3}
	newList, _ := PMapUintInt16PtrErr(plusOneUintInt16PtrErr, []*uint{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUintInt16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUintInt16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUintInt16PtrErr failed")
	}

	r, _ = PMapUintInt16PtrErr(nil, []*uint{})
	if len(r) > 0 {
		t.Errorf("PMapUintInt16PtrErr failed")
	}
	_, err := PMapUintInt16PtrErr(plusOneUintInt16PtrErr, []*uint{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUintInt16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUintInt16PtrErr(plusOneUintInt16PtrErr, []*uint{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintInt16PtrErr failed")
	}

	_, err = PMapUintInt16PtrErr(plusOneUintInt16PtrErr, []*uint{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUintInt16PtrErr failed")
	}

	_, err = PMapUintInt16PtrErr(plusOneUintInt16PtrErr, []*uint{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintInt16PtrErr failed")
	}
	_, err = PMapUintInt16PtrErr(plusOneUintInt16PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintInt16PtrErr failed")
	}

	expectedList = []*int16{&vo2, &vo3}
	newList, _ = PMapUintInt16PtrErr(plusOneUintInt16PtrErr, []*uint{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUintInt16Ptr failed")
	}

	_, err = PMapUintInt16PtrErr(plusOneUintInt16PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintInt16PtrErr failed")
	}
	_, err = PMapUintInt16PtrErr(plusOneUintInt16PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintInt16PtrErr failed")
	}

	_, err = PMapUintInt16PtrErr(plusOneUintInt16PtrErr, []*uint{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintInt16PtrErr failed")
	}

	_, err = PMapUintInt16PtrErr(plusOneUintInt16PtrErr, []*uint{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintInt16PtrErr failed")
	}

	_, err = PMapUintInt16PtrErr(plusOneUintInt16PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintInt16PtrErr failed")
	}
}

func TestPmapUintInt8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3

	expectedList := []*int8{&vo2, &vo3}
	newList, _ := PMapUintInt8PtrErr(plusOneUintInt8PtrErr, []*uint{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUintInt8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUintInt8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUintInt8PtrErr failed")
	}

	r, _ = PMapUintInt8PtrErr(nil, []*uint{})
	if len(r) > 0 {
		t.Errorf("PMapUintInt8PtrErr failed")
	}
	_, err := PMapUintInt8PtrErr(plusOneUintInt8PtrErr, []*uint{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUintInt8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUintInt8PtrErr(plusOneUintInt8PtrErr, []*uint{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintInt8PtrErr failed")
	}

	_, err = PMapUintInt8PtrErr(plusOneUintInt8PtrErr, []*uint{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUintInt8PtrErr failed")
	}

	_, err = PMapUintInt8PtrErr(plusOneUintInt8PtrErr, []*uint{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintInt8PtrErr failed")
	}
	_, err = PMapUintInt8PtrErr(plusOneUintInt8PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintInt8PtrErr failed")
	}

	expectedList = []*int8{&vo2, &vo3}
	newList, _ = PMapUintInt8PtrErr(plusOneUintInt8PtrErr, []*uint{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUintInt8Ptr failed")
	}

	_, err = PMapUintInt8PtrErr(plusOneUintInt8PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintInt8PtrErr failed")
	}
	_, err = PMapUintInt8PtrErr(plusOneUintInt8PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintInt8PtrErr failed")
	}

	_, err = PMapUintInt8PtrErr(plusOneUintInt8PtrErr, []*uint{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintInt8PtrErr failed")
	}

	_, err = PMapUintInt8PtrErr(plusOneUintInt8PtrErr, []*uint{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintInt8PtrErr failed")
	}

	_, err = PMapUintInt8PtrErr(plusOneUintInt8PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintInt8PtrErr failed")
	}
}

func TestPmapUintUint64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3

	expectedList := []*uint64{&vo2, &vo3}
	newList, _ := PMapUintUint64PtrErr(plusOneUintUint64PtrErr, []*uint{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUintUint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUintUint64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUintUint64PtrErr failed")
	}

	r, _ = PMapUintUint64PtrErr(nil, []*uint{})
	if len(r) > 0 {
		t.Errorf("PMapUintUint64PtrErr failed")
	}
	_, err := PMapUintUint64PtrErr(plusOneUintUint64PtrErr, []*uint{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUintUint64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUintUint64PtrErr(plusOneUintUint64PtrErr, []*uint{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintUint64PtrErr failed")
	}

	_, err = PMapUintUint64PtrErr(plusOneUintUint64PtrErr, []*uint{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUintUint64PtrErr failed")
	}

	_, err = PMapUintUint64PtrErr(plusOneUintUint64PtrErr, []*uint{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintUint64PtrErr failed")
	}
	_, err = PMapUintUint64PtrErr(plusOneUintUint64PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintUint64PtrErr failed")
	}

	expectedList = []*uint64{&vo2, &vo3}
	newList, _ = PMapUintUint64PtrErr(plusOneUintUint64PtrErr, []*uint{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUintUint64Ptr failed")
	}

	_, err = PMapUintUint64PtrErr(plusOneUintUint64PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintUint64PtrErr failed")
	}
	_, err = PMapUintUint64PtrErr(plusOneUintUint64PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintUint64PtrErr failed")
	}

	_, err = PMapUintUint64PtrErr(plusOneUintUint64PtrErr, []*uint{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintUint64PtrErr failed")
	}

	_, err = PMapUintUint64PtrErr(plusOneUintUint64PtrErr, []*uint{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintUint64PtrErr failed")
	}

	_, err = PMapUintUint64PtrErr(plusOneUintUint64PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintUint64PtrErr failed")
	}
}

func TestPmapUintUint32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3

	expectedList := []*uint32{&vo2, &vo3}
	newList, _ := PMapUintUint32PtrErr(plusOneUintUint32PtrErr, []*uint{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUintUint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUintUint32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUintUint32PtrErr failed")
	}

	r, _ = PMapUintUint32PtrErr(nil, []*uint{})
	if len(r) > 0 {
		t.Errorf("PMapUintUint32PtrErr failed")
	}
	_, err := PMapUintUint32PtrErr(plusOneUintUint32PtrErr, []*uint{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUintUint32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUintUint32PtrErr(plusOneUintUint32PtrErr, []*uint{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintUint32PtrErr failed")
	}

	_, err = PMapUintUint32PtrErr(plusOneUintUint32PtrErr, []*uint{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUintUint32PtrErr failed")
	}

	_, err = PMapUintUint32PtrErr(plusOneUintUint32PtrErr, []*uint{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintUint32PtrErr failed")
	}
	_, err = PMapUintUint32PtrErr(plusOneUintUint32PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintUint32PtrErr failed")
	}

	expectedList = []*uint32{&vo2, &vo3}
	newList, _ = PMapUintUint32PtrErr(plusOneUintUint32PtrErr, []*uint{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUintUint32Ptr failed")
	}

	_, err = PMapUintUint32PtrErr(plusOneUintUint32PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintUint32PtrErr failed")
	}
	_, err = PMapUintUint32PtrErr(plusOneUintUint32PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintUint32PtrErr failed")
	}

	_, err = PMapUintUint32PtrErr(plusOneUintUint32PtrErr, []*uint{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintUint32PtrErr failed")
	}

	_, err = PMapUintUint32PtrErr(plusOneUintUint32PtrErr, []*uint{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintUint32PtrErr failed")
	}

	_, err = PMapUintUint32PtrErr(plusOneUintUint32PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintUint32PtrErr failed")
	}
}

func TestPmapUintUint16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3

	expectedList := []*uint16{&vo2, &vo3}
	newList, _ := PMapUintUint16PtrErr(plusOneUintUint16PtrErr, []*uint{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUintUint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUintUint16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUintUint16PtrErr failed")
	}

	r, _ = PMapUintUint16PtrErr(nil, []*uint{})
	if len(r) > 0 {
		t.Errorf("PMapUintUint16PtrErr failed")
	}
	_, err := PMapUintUint16PtrErr(plusOneUintUint16PtrErr, []*uint{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUintUint16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUintUint16PtrErr(plusOneUintUint16PtrErr, []*uint{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintUint16PtrErr failed")
	}

	_, err = PMapUintUint16PtrErr(plusOneUintUint16PtrErr, []*uint{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUintUint16PtrErr failed")
	}

	_, err = PMapUintUint16PtrErr(plusOneUintUint16PtrErr, []*uint{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintUint16PtrErr failed")
	}
	_, err = PMapUintUint16PtrErr(plusOneUintUint16PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintUint16PtrErr failed")
	}

	expectedList = []*uint16{&vo2, &vo3}
	newList, _ = PMapUintUint16PtrErr(plusOneUintUint16PtrErr, []*uint{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUintUint16Ptr failed")
	}

	_, err = PMapUintUint16PtrErr(plusOneUintUint16PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintUint16PtrErr failed")
	}
	_, err = PMapUintUint16PtrErr(plusOneUintUint16PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintUint16PtrErr failed")
	}

	_, err = PMapUintUint16PtrErr(plusOneUintUint16PtrErr, []*uint{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintUint16PtrErr failed")
	}

	_, err = PMapUintUint16PtrErr(plusOneUintUint16PtrErr, []*uint{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintUint16PtrErr failed")
	}

	_, err = PMapUintUint16PtrErr(plusOneUintUint16PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintUint16PtrErr failed")
	}
}

func TestPmapUintUint8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3

	expectedList := []*uint8{&vo2, &vo3}
	newList, _ := PMapUintUint8PtrErr(plusOneUintUint8PtrErr, []*uint{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUintUint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUintUint8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUintUint8PtrErr failed")
	}

	r, _ = PMapUintUint8PtrErr(nil, []*uint{})
	if len(r) > 0 {
		t.Errorf("PMapUintUint8PtrErr failed")
	}
	_, err := PMapUintUint8PtrErr(plusOneUintUint8PtrErr, []*uint{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUintUint8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUintUint8PtrErr(plusOneUintUint8PtrErr, []*uint{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintUint8PtrErr failed")
	}

	_, err = PMapUintUint8PtrErr(plusOneUintUint8PtrErr, []*uint{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUintUint8PtrErr failed")
	}

	_, err = PMapUintUint8PtrErr(plusOneUintUint8PtrErr, []*uint{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintUint8PtrErr failed")
	}
	_, err = PMapUintUint8PtrErr(plusOneUintUint8PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintUint8PtrErr failed")
	}

	expectedList = []*uint8{&vo2, &vo3}
	newList, _ = PMapUintUint8PtrErr(plusOneUintUint8PtrErr, []*uint{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUintUint8Ptr failed")
	}

	_, err = PMapUintUint8PtrErr(plusOneUintUint8PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintUint8PtrErr failed")
	}
	_, err = PMapUintUint8PtrErr(plusOneUintUint8PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintUint8PtrErr failed")
	}

	_, err = PMapUintUint8PtrErr(plusOneUintUint8PtrErr, []*uint{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintUint8PtrErr failed")
	}

	_, err = PMapUintUint8PtrErr(plusOneUintUint8PtrErr, []*uint{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintUint8PtrErr failed")
	}

	_, err = PMapUintUint8PtrErr(plusOneUintUint8PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintUint8PtrErr failed")
	}
}

func TestPmapUintStrPtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 uint = 10
	var vi0 uint

	expectedList := []*string{&vo10}
	newList, _ := PMapUintStrPtrErr(someLogicUintStrPtrErr, []*uint{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapUintStrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUintStrPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUintStrPtrErr failed")
	}

	r, _ = PMapUintStrPtrErr(nil, []*uint{})
	if len(r) > 0 {
		t.Errorf("PMapUintStrPtrErr failed")
	}
	_, err := PMapUintStrPtrErr(someLogicUintStrPtrErr, []*uint{&vi10, &vi0})
	if err == nil {
		t.Errorf("PMapUintStrPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUintStrPtrErr(someLogicUintStrPtrErr, []*uint{&vi10, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintStrPtrErr failed")
	}

	_, err = PMapUintStrPtrErr(someLogicUintStrPtrErr, []*uint{&vi0, &vi0, &vi10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUintStrPtrErr failed")
	}

	_, err = PMapUintStrPtrErr(someLogicUintStrPtrErr, []*uint{&vi0, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintStrPtrErr failed")
	}
	_, err = PMapUintStrPtrErr(someLogicUintStrPtrErr, []*uint{&vi10, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintStrPtrErr failed")
	}

	expectedList = []*string{&vo10, &vo10}
	newList, _ = PMapUintStrPtrErr(someLogicUintStrPtrErr, []*uint{&vi10, &vi10}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUintStrPtr failed")
	}

	_, err = PMapUintStrPtrErr(someLogicUintStrPtrErr, []*uint{&vi10, &vi10, &vi0}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintStrPtrErr failed")
	}
	_, err = PMapUintStrPtrErr(someLogicUintStrPtrErr, []*uint{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintStrPtrErr failed")
	}

	_, err = PMapUintStrPtrErr(someLogicUintStrPtrErr, []*uint{&vi0, &vi0, &vi10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintStrPtrErr failed")
	}

	_, err = PMapUintStrPtrErr(someLogicUintStrPtrErr, []*uint{&vi0, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintStrPtrErr failed")
	}

	_, err = PMapUintStrPtrErr(someLogicUintStrPtrErr, []*uint{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintStrPtrErr failed")
	}
}

func TestPmapUintBoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 uint = 10
	var v0 uint
	var v3 uint = 3

	expectedList := []*bool{&vt, &vf}
	newList, _ := PMapUintBoolPtrErr(someLogicUintBoolPtrErr, []*uint{&v10, &v0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUintBoolPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUintBoolPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUintBoolPtrErr failed")
	}

	r, _ = PMapUintBoolPtrErr(nil, []*uint{})
	if len(r) > 0 {
		t.Errorf("PMapUintBoolPtrErr failed")
	}
	_, err := PMapUintBoolPtrErr(someLogicUintBoolPtrErr, []*uint{&v10, &v3})
	if err == nil {
		t.Errorf("PMapUintBoolPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUintBoolPtrErr(someLogicUintBoolPtrErr, []*uint{&v10, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintBoolPtrErr failed")
	}

	_, err = PMapUintBoolPtrErr(someLogicUintBoolPtrErr, []*uint{&v3, &v3, &v10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUintBoolPtrErr failed")
	}

	_, err = PMapUintBoolPtrErr(someLogicUintBoolPtrErr, []*uint{&v3, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintBoolErr failed")
	}
	_, err = PMapUintBoolPtrErr(someLogicUintBoolPtrErr, []*uint{&v10, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintBoolPtrErr failed")
	}

	expectedList = []*bool{&vt, &vt}
	newList, _ = PMapUintBoolPtrErr(someLogicUintBoolPtrErr, []*uint{&v10, &v10}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUintBoolPtr failed")
	}

	_, err = PMapUintBoolPtrErr(someLogicUintBoolPtrErr, []*uint{&v10, &v10, &v3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintBoolErr failed")
	}
	_, err = PMapUintBoolPtrErr(someLogicUintBoolPtrErr, []*uint{&v10, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintBoolPtrErr failed")
	}

	_, err = PMapUintBoolPtrErr(someLogicUintBoolPtrErr, []*uint{&v3, &v3, &v10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintBoolPtrErr failed")
	}

	_, err = PMapUintBoolPtrErr(someLogicUintBoolPtrErr, []*uint{&v3, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintBoolPtrErr failed")
	}

	_, err = PMapUintBoolPtrErr(someLogicUintBoolPtrErr, []*uint{&v10, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintBoolPtrErr failed")
	}
}

func TestPmapUintFloat32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3

	expectedList := []*float32{&vo2, &vo3}
	newList, _ := PMapUintFloat32PtrErr(plusOneUintFloat32PtrErr, []*uint{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUintFloat32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUintFloat32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUintFloat32PtrErr failed")
	}

	r, _ = PMapUintFloat32PtrErr(nil, []*uint{})
	if len(r) > 0 {
		t.Errorf("PMapUintFloat32PtrErr failed")
	}
	_, err := PMapUintFloat32PtrErr(plusOneUintFloat32PtrErr, []*uint{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUintFloat32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUintFloat32PtrErr(plusOneUintFloat32PtrErr, []*uint{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintFloat32PtrErr failed")
	}

	_, err = PMapUintFloat32PtrErr(plusOneUintFloat32PtrErr, []*uint{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUintFloat32PtrErr failed")
	}

	_, err = PMapUintFloat32PtrErr(plusOneUintFloat32PtrErr, []*uint{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintFloat32PtrErr failed")
	}
	_, err = PMapUintFloat32PtrErr(plusOneUintFloat32PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintFloat32PtrErr failed")
	}

	expectedList = []*float32{&vo2, &vo3}
	newList, _ = PMapUintFloat32PtrErr(plusOneUintFloat32PtrErr, []*uint{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUintFloat32Ptr failed")
	}

	_, err = PMapUintFloat32PtrErr(plusOneUintFloat32PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintFloat32PtrErr failed")
	}
	_, err = PMapUintFloat32PtrErr(plusOneUintFloat32PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintFloat32PtrErr failed")
	}

	_, err = PMapUintFloat32PtrErr(plusOneUintFloat32PtrErr, []*uint{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintFloat32PtrErr failed")
	}

	_, err = PMapUintFloat32PtrErr(plusOneUintFloat32PtrErr, []*uint{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintFloat32PtrErr failed")
	}

	_, err = PMapUintFloat32PtrErr(plusOneUintFloat32PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintFloat32PtrErr failed")
	}
}

func TestPmapUintFloat64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3

	expectedList := []*float64{&vo2, &vo3}
	newList, _ := PMapUintFloat64PtrErr(plusOneUintFloat64PtrErr, []*uint{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUintFloat64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUintFloat64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUintFloat64PtrErr failed")
	}

	r, _ = PMapUintFloat64PtrErr(nil, []*uint{})
	if len(r) > 0 {
		t.Errorf("PMapUintFloat64PtrErr failed")
	}
	_, err := PMapUintFloat64PtrErr(plusOneUintFloat64PtrErr, []*uint{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUintFloat64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUintFloat64PtrErr(plusOneUintFloat64PtrErr, []*uint{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintFloat64PtrErr failed")
	}

	_, err = PMapUintFloat64PtrErr(plusOneUintFloat64PtrErr, []*uint{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUintFloat64PtrErr failed")
	}

	_, err = PMapUintFloat64PtrErr(plusOneUintFloat64PtrErr, []*uint{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintFloat64PtrErr failed")
	}
	_, err = PMapUintFloat64PtrErr(plusOneUintFloat64PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintFloat64PtrErr failed")
	}

	expectedList = []*float64{&vo2, &vo3}
	newList, _ = PMapUintFloat64PtrErr(plusOneUintFloat64PtrErr, []*uint{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUintFloat64Ptr failed")
	}

	_, err = PMapUintFloat64PtrErr(plusOneUintFloat64PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintFloat64PtrErr failed")
	}
	_, err = PMapUintFloat64PtrErr(plusOneUintFloat64PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintFloat64PtrErr failed")
	}

	_, err = PMapUintFloat64PtrErr(plusOneUintFloat64PtrErr, []*uint{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintFloat64PtrErr failed")
	}

	_, err = PMapUintFloat64PtrErr(plusOneUintFloat64PtrErr, []*uint{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintFloat64PtrErr failed")
	}

	_, err = PMapUintFloat64PtrErr(plusOneUintFloat64PtrErr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUintFloat64PtrErr failed")
	}
}

func TestPmapUint64IntPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3

	expectedList := []*int{&vo2, &vo3}
	newList, _ := PMapUint64IntPtrErr(plusOneUint64IntPtrErr, []*uint64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint64IntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint64IntPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint64IntPtrErr failed")
	}

	r, _ = PMapUint64IntPtrErr(nil, []*uint64{})
	if len(r) > 0 {
		t.Errorf("PMapUint64IntPtrErr failed")
	}
	_, err := PMapUint64IntPtrErr(plusOneUint64IntPtrErr, []*uint64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint64IntPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint64IntPtrErr(plusOneUint64IntPtrErr, []*uint64{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64IntPtrErr failed")
	}

	_, err = PMapUint64IntPtrErr(plusOneUint64IntPtrErr, []*uint64{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint64IntPtrErr failed")
	}

	_, err = PMapUint64IntPtrErr(plusOneUint64IntPtrErr, []*uint64{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64IntPtrErr failed")
	}
	_, err = PMapUint64IntPtrErr(plusOneUint64IntPtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64IntPtrErr failed")
	}

	expectedList = []*int{&vo2, &vo3}
	newList, _ = PMapUint64IntPtrErr(plusOneUint64IntPtrErr, []*uint64{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint64IntPtr failed")
	}

	_, err = PMapUint64IntPtrErr(plusOneUint64IntPtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64IntPtrErr failed")
	}
	_, err = PMapUint64IntPtrErr(plusOneUint64IntPtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64IntPtrErr failed")
	}

	_, err = PMapUint64IntPtrErr(plusOneUint64IntPtrErr, []*uint64{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64IntPtrErr failed")
	}

	_, err = PMapUint64IntPtrErr(plusOneUint64IntPtrErr, []*uint64{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64IntPtrErr failed")
	}

	_, err = PMapUint64IntPtrErr(plusOneUint64IntPtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64IntPtrErr failed")
	}
}

func TestPmapUint64Int64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3

	expectedList := []*int64{&vo2, &vo3}
	newList, _ := PMapUint64Int64PtrErr(plusOneUint64Int64PtrErr, []*uint64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint64Int64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint64Int64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint64Int64PtrErr failed")
	}

	r, _ = PMapUint64Int64PtrErr(nil, []*uint64{})
	if len(r) > 0 {
		t.Errorf("PMapUint64Int64PtrErr failed")
	}
	_, err := PMapUint64Int64PtrErr(plusOneUint64Int64PtrErr, []*uint64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint64Int64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint64Int64PtrErr(plusOneUint64Int64PtrErr, []*uint64{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64Int64PtrErr failed")
	}

	_, err = PMapUint64Int64PtrErr(plusOneUint64Int64PtrErr, []*uint64{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint64Int64PtrErr failed")
	}

	_, err = PMapUint64Int64PtrErr(plusOneUint64Int64PtrErr, []*uint64{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64Int64PtrErr failed")
	}
	_, err = PMapUint64Int64PtrErr(plusOneUint64Int64PtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64Int64PtrErr failed")
	}

	expectedList = []*int64{&vo2, &vo3}
	newList, _ = PMapUint64Int64PtrErr(plusOneUint64Int64PtrErr, []*uint64{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint64Int64Ptr failed")
	}

	_, err = PMapUint64Int64PtrErr(plusOneUint64Int64PtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Int64PtrErr failed")
	}
	_, err = PMapUint64Int64PtrErr(plusOneUint64Int64PtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Int64PtrErr failed")
	}

	_, err = PMapUint64Int64PtrErr(plusOneUint64Int64PtrErr, []*uint64{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Int64PtrErr failed")
	}

	_, err = PMapUint64Int64PtrErr(plusOneUint64Int64PtrErr, []*uint64{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Int64PtrErr failed")
	}

	_, err = PMapUint64Int64PtrErr(plusOneUint64Int64PtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Int64PtrErr failed")
	}
}

func TestPmapUint64Int32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3

	expectedList := []*int32{&vo2, &vo3}
	newList, _ := PMapUint64Int32PtrErr(plusOneUint64Int32PtrErr, []*uint64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint64Int32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint64Int32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint64Int32PtrErr failed")
	}

	r, _ = PMapUint64Int32PtrErr(nil, []*uint64{})
	if len(r) > 0 {
		t.Errorf("PMapUint64Int32PtrErr failed")
	}
	_, err := PMapUint64Int32PtrErr(plusOneUint64Int32PtrErr, []*uint64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint64Int32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint64Int32PtrErr(plusOneUint64Int32PtrErr, []*uint64{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64Int32PtrErr failed")
	}

	_, err = PMapUint64Int32PtrErr(plusOneUint64Int32PtrErr, []*uint64{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint64Int32PtrErr failed")
	}

	_, err = PMapUint64Int32PtrErr(plusOneUint64Int32PtrErr, []*uint64{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64Int32PtrErr failed")
	}
	_, err = PMapUint64Int32PtrErr(plusOneUint64Int32PtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64Int32PtrErr failed")
	}

	expectedList = []*int32{&vo2, &vo3}
	newList, _ = PMapUint64Int32PtrErr(plusOneUint64Int32PtrErr, []*uint64{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint64Int32Ptr failed")
	}

	_, err = PMapUint64Int32PtrErr(plusOneUint64Int32PtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Int32PtrErr failed")
	}
	_, err = PMapUint64Int32PtrErr(plusOneUint64Int32PtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Int32PtrErr failed")
	}

	_, err = PMapUint64Int32PtrErr(plusOneUint64Int32PtrErr, []*uint64{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Int32PtrErr failed")
	}

	_, err = PMapUint64Int32PtrErr(plusOneUint64Int32PtrErr, []*uint64{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Int32PtrErr failed")
	}

	_, err = PMapUint64Int32PtrErr(plusOneUint64Int32PtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Int32PtrErr failed")
	}
}

func TestPmapUint64Int16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3

	expectedList := []*int16{&vo2, &vo3}
	newList, _ := PMapUint64Int16PtrErr(plusOneUint64Int16PtrErr, []*uint64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint64Int16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint64Int16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint64Int16PtrErr failed")
	}

	r, _ = PMapUint64Int16PtrErr(nil, []*uint64{})
	if len(r) > 0 {
		t.Errorf("PMapUint64Int16PtrErr failed")
	}
	_, err := PMapUint64Int16PtrErr(plusOneUint64Int16PtrErr, []*uint64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint64Int16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint64Int16PtrErr(plusOneUint64Int16PtrErr, []*uint64{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64Int16PtrErr failed")
	}

	_, err = PMapUint64Int16PtrErr(plusOneUint64Int16PtrErr, []*uint64{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint64Int16PtrErr failed")
	}

	_, err = PMapUint64Int16PtrErr(plusOneUint64Int16PtrErr, []*uint64{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64Int16PtrErr failed")
	}
	_, err = PMapUint64Int16PtrErr(plusOneUint64Int16PtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64Int16PtrErr failed")
	}

	expectedList = []*int16{&vo2, &vo3}
	newList, _ = PMapUint64Int16PtrErr(plusOneUint64Int16PtrErr, []*uint64{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint64Int16Ptr failed")
	}

	_, err = PMapUint64Int16PtrErr(plusOneUint64Int16PtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Int16PtrErr failed")
	}
	_, err = PMapUint64Int16PtrErr(plusOneUint64Int16PtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Int16PtrErr failed")
	}

	_, err = PMapUint64Int16PtrErr(plusOneUint64Int16PtrErr, []*uint64{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Int16PtrErr failed")
	}

	_, err = PMapUint64Int16PtrErr(plusOneUint64Int16PtrErr, []*uint64{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Int16PtrErr failed")
	}

	_, err = PMapUint64Int16PtrErr(plusOneUint64Int16PtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Int16PtrErr failed")
	}
}

func TestPmapUint64Int8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3

	expectedList := []*int8{&vo2, &vo3}
	newList, _ := PMapUint64Int8PtrErr(plusOneUint64Int8PtrErr, []*uint64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint64Int8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint64Int8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint64Int8PtrErr failed")
	}

	r, _ = PMapUint64Int8PtrErr(nil, []*uint64{})
	if len(r) > 0 {
		t.Errorf("PMapUint64Int8PtrErr failed")
	}
	_, err := PMapUint64Int8PtrErr(plusOneUint64Int8PtrErr, []*uint64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint64Int8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint64Int8PtrErr(plusOneUint64Int8PtrErr, []*uint64{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64Int8PtrErr failed")
	}

	_, err = PMapUint64Int8PtrErr(plusOneUint64Int8PtrErr, []*uint64{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint64Int8PtrErr failed")
	}

	_, err = PMapUint64Int8PtrErr(plusOneUint64Int8PtrErr, []*uint64{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64Int8PtrErr failed")
	}
	_, err = PMapUint64Int8PtrErr(plusOneUint64Int8PtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64Int8PtrErr failed")
	}

	expectedList = []*int8{&vo2, &vo3}
	newList, _ = PMapUint64Int8PtrErr(plusOneUint64Int8PtrErr, []*uint64{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint64Int8Ptr failed")
	}

	_, err = PMapUint64Int8PtrErr(plusOneUint64Int8PtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Int8PtrErr failed")
	}
	_, err = PMapUint64Int8PtrErr(plusOneUint64Int8PtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Int8PtrErr failed")
	}

	_, err = PMapUint64Int8PtrErr(plusOneUint64Int8PtrErr, []*uint64{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Int8PtrErr failed")
	}

	_, err = PMapUint64Int8PtrErr(plusOneUint64Int8PtrErr, []*uint64{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Int8PtrErr failed")
	}

	_, err = PMapUint64Int8PtrErr(plusOneUint64Int8PtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Int8PtrErr failed")
	}
}

func TestPmapUint64UintPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3

	expectedList := []*uint{&vo2, &vo3}
	newList, _ := PMapUint64UintPtrErr(plusOneUint64UintPtrErr, []*uint64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint64UintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint64UintPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint64UintPtrErr failed")
	}

	r, _ = PMapUint64UintPtrErr(nil, []*uint64{})
	if len(r) > 0 {
		t.Errorf("PMapUint64UintPtrErr failed")
	}
	_, err := PMapUint64UintPtrErr(plusOneUint64UintPtrErr, []*uint64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint64UintPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint64UintPtrErr(plusOneUint64UintPtrErr, []*uint64{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64UintPtrErr failed")
	}

	_, err = PMapUint64UintPtrErr(plusOneUint64UintPtrErr, []*uint64{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint64UintPtrErr failed")
	}

	_, err = PMapUint64UintPtrErr(plusOneUint64UintPtrErr, []*uint64{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64UintPtrErr failed")
	}
	_, err = PMapUint64UintPtrErr(plusOneUint64UintPtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64UintPtrErr failed")
	}

	expectedList = []*uint{&vo2, &vo3}
	newList, _ = PMapUint64UintPtrErr(plusOneUint64UintPtrErr, []*uint64{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint64UintPtr failed")
	}

	_, err = PMapUint64UintPtrErr(plusOneUint64UintPtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64UintPtrErr failed")
	}
	_, err = PMapUint64UintPtrErr(plusOneUint64UintPtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64UintPtrErr failed")
	}

	_, err = PMapUint64UintPtrErr(plusOneUint64UintPtrErr, []*uint64{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64UintPtrErr failed")
	}

	_, err = PMapUint64UintPtrErr(plusOneUint64UintPtrErr, []*uint64{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64UintPtrErr failed")
	}

	_, err = PMapUint64UintPtrErr(plusOneUint64UintPtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64UintPtrErr failed")
	}
}

func TestPmapUint64Uint32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3

	expectedList := []*uint32{&vo2, &vo3}
	newList, _ := PMapUint64Uint32PtrErr(plusOneUint64Uint32PtrErr, []*uint64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint64Uint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint64Uint32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint64Uint32PtrErr failed")
	}

	r, _ = PMapUint64Uint32PtrErr(nil, []*uint64{})
	if len(r) > 0 {
		t.Errorf("PMapUint64Uint32PtrErr failed")
	}
	_, err := PMapUint64Uint32PtrErr(plusOneUint64Uint32PtrErr, []*uint64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint64Uint32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint64Uint32PtrErr(plusOneUint64Uint32PtrErr, []*uint64{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64Uint32PtrErr failed")
	}

	_, err = PMapUint64Uint32PtrErr(plusOneUint64Uint32PtrErr, []*uint64{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint64Uint32PtrErr failed")
	}

	_, err = PMapUint64Uint32PtrErr(plusOneUint64Uint32PtrErr, []*uint64{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64Uint32PtrErr failed")
	}
	_, err = PMapUint64Uint32PtrErr(plusOneUint64Uint32PtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64Uint32PtrErr failed")
	}

	expectedList = []*uint32{&vo2, &vo3}
	newList, _ = PMapUint64Uint32PtrErr(plusOneUint64Uint32PtrErr, []*uint64{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint64Uint32Ptr failed")
	}

	_, err = PMapUint64Uint32PtrErr(plusOneUint64Uint32PtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Uint32PtrErr failed")
	}
	_, err = PMapUint64Uint32PtrErr(plusOneUint64Uint32PtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Uint32PtrErr failed")
	}

	_, err = PMapUint64Uint32PtrErr(plusOneUint64Uint32PtrErr, []*uint64{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Uint32PtrErr failed")
	}

	_, err = PMapUint64Uint32PtrErr(plusOneUint64Uint32PtrErr, []*uint64{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Uint32PtrErr failed")
	}

	_, err = PMapUint64Uint32PtrErr(plusOneUint64Uint32PtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Uint32PtrErr failed")
	}
}

func TestPmapUint64Uint16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3

	expectedList := []*uint16{&vo2, &vo3}
	newList, _ := PMapUint64Uint16PtrErr(plusOneUint64Uint16PtrErr, []*uint64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint64Uint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint64Uint16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint64Uint16PtrErr failed")
	}

	r, _ = PMapUint64Uint16PtrErr(nil, []*uint64{})
	if len(r) > 0 {
		t.Errorf("PMapUint64Uint16PtrErr failed")
	}
	_, err := PMapUint64Uint16PtrErr(plusOneUint64Uint16PtrErr, []*uint64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint64Uint16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint64Uint16PtrErr(plusOneUint64Uint16PtrErr, []*uint64{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64Uint16PtrErr failed")
	}

	_, err = PMapUint64Uint16PtrErr(plusOneUint64Uint16PtrErr, []*uint64{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint64Uint16PtrErr failed")
	}

	_, err = PMapUint64Uint16PtrErr(plusOneUint64Uint16PtrErr, []*uint64{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64Uint16PtrErr failed")
	}
	_, err = PMapUint64Uint16PtrErr(plusOneUint64Uint16PtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64Uint16PtrErr failed")
	}

	expectedList = []*uint16{&vo2, &vo3}
	newList, _ = PMapUint64Uint16PtrErr(plusOneUint64Uint16PtrErr, []*uint64{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint64Uint16Ptr failed")
	}

	_, err = PMapUint64Uint16PtrErr(plusOneUint64Uint16PtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Uint16PtrErr failed")
	}
	_, err = PMapUint64Uint16PtrErr(plusOneUint64Uint16PtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Uint16PtrErr failed")
	}

	_, err = PMapUint64Uint16PtrErr(plusOneUint64Uint16PtrErr, []*uint64{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Uint16PtrErr failed")
	}

	_, err = PMapUint64Uint16PtrErr(plusOneUint64Uint16PtrErr, []*uint64{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Uint16PtrErr failed")
	}

	_, err = PMapUint64Uint16PtrErr(plusOneUint64Uint16PtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Uint16PtrErr failed")
	}
}

func TestPmapUint64Uint8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3

	expectedList := []*uint8{&vo2, &vo3}
	newList, _ := PMapUint64Uint8PtrErr(plusOneUint64Uint8PtrErr, []*uint64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint64Uint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint64Uint8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint64Uint8PtrErr failed")
	}

	r, _ = PMapUint64Uint8PtrErr(nil, []*uint64{})
	if len(r) > 0 {
		t.Errorf("PMapUint64Uint8PtrErr failed")
	}
	_, err := PMapUint64Uint8PtrErr(plusOneUint64Uint8PtrErr, []*uint64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint64Uint8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint64Uint8PtrErr(plusOneUint64Uint8PtrErr, []*uint64{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64Uint8PtrErr failed")
	}

	_, err = PMapUint64Uint8PtrErr(plusOneUint64Uint8PtrErr, []*uint64{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint64Uint8PtrErr failed")
	}

	_, err = PMapUint64Uint8PtrErr(plusOneUint64Uint8PtrErr, []*uint64{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64Uint8PtrErr failed")
	}
	_, err = PMapUint64Uint8PtrErr(plusOneUint64Uint8PtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64Uint8PtrErr failed")
	}

	expectedList = []*uint8{&vo2, &vo3}
	newList, _ = PMapUint64Uint8PtrErr(plusOneUint64Uint8PtrErr, []*uint64{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint64Uint8Ptr failed")
	}

	_, err = PMapUint64Uint8PtrErr(plusOneUint64Uint8PtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Uint8PtrErr failed")
	}
	_, err = PMapUint64Uint8PtrErr(plusOneUint64Uint8PtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Uint8PtrErr failed")
	}

	_, err = PMapUint64Uint8PtrErr(plusOneUint64Uint8PtrErr, []*uint64{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Uint8PtrErr failed")
	}

	_, err = PMapUint64Uint8PtrErr(plusOneUint64Uint8PtrErr, []*uint64{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Uint8PtrErr failed")
	}

	_, err = PMapUint64Uint8PtrErr(plusOneUint64Uint8PtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Uint8PtrErr failed")
	}
}

func TestPmapUint64StrPtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 uint64 = 10
	var vi0 uint64

	expectedList := []*string{&vo10}
	newList, _ := PMapUint64StrPtrErr(someLogicUint64StrPtrErr, []*uint64{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapUint64StrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint64StrPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint64StrPtrErr failed")
	}

	r, _ = PMapUint64StrPtrErr(nil, []*uint64{})
	if len(r) > 0 {
		t.Errorf("PMapUint64StrPtrErr failed")
	}
	_, err := PMapUint64StrPtrErr(someLogicUint64StrPtrErr, []*uint64{&vi10, &vi0})
	if err == nil {
		t.Errorf("PMapUint64StrPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint64StrPtrErr(someLogicUint64StrPtrErr, []*uint64{&vi10, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64StrPtrErr failed")
	}

	_, err = PMapUint64StrPtrErr(someLogicUint64StrPtrErr, []*uint64{&vi0, &vi0, &vi10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint64StrPtrErr failed")
	}

	_, err = PMapUint64StrPtrErr(someLogicUint64StrPtrErr, []*uint64{&vi0, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64StrPtrErr failed")
	}
	_, err = PMapUint64StrPtrErr(someLogicUint64StrPtrErr, []*uint64{&vi10, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64StrPtrErr failed")
	}

	expectedList = []*string{&vo10, &vo10}
	newList, _ = PMapUint64StrPtrErr(someLogicUint64StrPtrErr, []*uint64{&vi10, &vi10}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint64StrPtr failed")
	}

	_, err = PMapUint64StrPtrErr(someLogicUint64StrPtrErr, []*uint64{&vi10, &vi10, &vi0}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64StrPtrErr failed")
	}
	_, err = PMapUint64StrPtrErr(someLogicUint64StrPtrErr, []*uint64{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64StrPtrErr failed")
	}

	_, err = PMapUint64StrPtrErr(someLogicUint64StrPtrErr, []*uint64{&vi0, &vi0, &vi10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64StrPtrErr failed")
	}

	_, err = PMapUint64StrPtrErr(someLogicUint64StrPtrErr, []*uint64{&vi0, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64StrPtrErr failed")
	}

	_, err = PMapUint64StrPtrErr(someLogicUint64StrPtrErr, []*uint64{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64StrPtrErr failed")
	}
}

func TestPmapUint64BoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 uint64 = 10
	var v0 uint64
	var v3 uint64 = 3

	expectedList := []*bool{&vt, &vf}
	newList, _ := PMapUint64BoolPtrErr(someLogicUint64BoolPtrErr, []*uint64{&v10, &v0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint64BoolPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint64BoolPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint64BoolPtrErr failed")
	}

	r, _ = PMapUint64BoolPtrErr(nil, []*uint64{})
	if len(r) > 0 {
		t.Errorf("PMapUint64BoolPtrErr failed")
	}
	_, err := PMapUint64BoolPtrErr(someLogicUint64BoolPtrErr, []*uint64{&v10, &v3})
	if err == nil {
		t.Errorf("PMapUint64BoolPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint64BoolPtrErr(someLogicUint64BoolPtrErr, []*uint64{&v10, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64BoolPtrErr failed")
	}

	_, err = PMapUint64BoolPtrErr(someLogicUint64BoolPtrErr, []*uint64{&v3, &v3, &v10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint64BoolPtrErr failed")
	}

	_, err = PMapUint64BoolPtrErr(someLogicUint64BoolPtrErr, []*uint64{&v3, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64BoolErr failed")
	}
	_, err = PMapUint64BoolPtrErr(someLogicUint64BoolPtrErr, []*uint64{&v10, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64BoolPtrErr failed")
	}

	expectedList = []*bool{&vt, &vt}
	newList, _ = PMapUint64BoolPtrErr(someLogicUint64BoolPtrErr, []*uint64{&v10, &v10}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint64BoolPtr failed")
	}

	_, err = PMapUint64BoolPtrErr(someLogicUint64BoolPtrErr, []*uint64{&v10, &v10, &v3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64BoolErr failed")
	}
	_, err = PMapUint64BoolPtrErr(someLogicUint64BoolPtrErr, []*uint64{&v10, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64BoolPtrErr failed")
	}

	_, err = PMapUint64BoolPtrErr(someLogicUint64BoolPtrErr, []*uint64{&v3, &v3, &v10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64BoolPtrErr failed")
	}

	_, err = PMapUint64BoolPtrErr(someLogicUint64BoolPtrErr, []*uint64{&v3, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64BoolPtrErr failed")
	}

	_, err = PMapUint64BoolPtrErr(someLogicUint64BoolPtrErr, []*uint64{&v10, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64BoolPtrErr failed")
	}
}

func TestPmapUint64Float32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3

	expectedList := []*float32{&vo2, &vo3}
	newList, _ := PMapUint64Float32PtrErr(plusOneUint64Float32PtrErr, []*uint64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint64Float32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint64Float32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint64Float32PtrErr failed")
	}

	r, _ = PMapUint64Float32PtrErr(nil, []*uint64{})
	if len(r) > 0 {
		t.Errorf("PMapUint64Float32PtrErr failed")
	}
	_, err := PMapUint64Float32PtrErr(plusOneUint64Float32PtrErr, []*uint64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint64Float32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint64Float32PtrErr(plusOneUint64Float32PtrErr, []*uint64{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64Float32PtrErr failed")
	}

	_, err = PMapUint64Float32PtrErr(plusOneUint64Float32PtrErr, []*uint64{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint64Float32PtrErr failed")
	}

	_, err = PMapUint64Float32PtrErr(plusOneUint64Float32PtrErr, []*uint64{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64Float32PtrErr failed")
	}
	_, err = PMapUint64Float32PtrErr(plusOneUint64Float32PtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64Float32PtrErr failed")
	}

	expectedList = []*float32{&vo2, &vo3}
	newList, _ = PMapUint64Float32PtrErr(plusOneUint64Float32PtrErr, []*uint64{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint64Float32Ptr failed")
	}

	_, err = PMapUint64Float32PtrErr(plusOneUint64Float32PtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Float32PtrErr failed")
	}
	_, err = PMapUint64Float32PtrErr(plusOneUint64Float32PtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Float32PtrErr failed")
	}

	_, err = PMapUint64Float32PtrErr(plusOneUint64Float32PtrErr, []*uint64{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Float32PtrErr failed")
	}

	_, err = PMapUint64Float32PtrErr(plusOneUint64Float32PtrErr, []*uint64{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Float32PtrErr failed")
	}

	_, err = PMapUint64Float32PtrErr(plusOneUint64Float32PtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Float32PtrErr failed")
	}
}

func TestPmapUint64Float64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3

	expectedList := []*float64{&vo2, &vo3}
	newList, _ := PMapUint64Float64PtrErr(plusOneUint64Float64PtrErr, []*uint64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint64Float64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint64Float64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint64Float64PtrErr failed")
	}

	r, _ = PMapUint64Float64PtrErr(nil, []*uint64{})
	if len(r) > 0 {
		t.Errorf("PMapUint64Float64PtrErr failed")
	}
	_, err := PMapUint64Float64PtrErr(plusOneUint64Float64PtrErr, []*uint64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint64Float64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint64Float64PtrErr(plusOneUint64Float64PtrErr, []*uint64{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64Float64PtrErr failed")
	}

	_, err = PMapUint64Float64PtrErr(plusOneUint64Float64PtrErr, []*uint64{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint64Float64PtrErr failed")
	}

	_, err = PMapUint64Float64PtrErr(plusOneUint64Float64PtrErr, []*uint64{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64Float64PtrErr failed")
	}
	_, err = PMapUint64Float64PtrErr(plusOneUint64Float64PtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64Float64PtrErr failed")
	}

	expectedList = []*float64{&vo2, &vo3}
	newList, _ = PMapUint64Float64PtrErr(plusOneUint64Float64PtrErr, []*uint64{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint64Float64Ptr failed")
	}

	_, err = PMapUint64Float64PtrErr(plusOneUint64Float64PtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Float64PtrErr failed")
	}
	_, err = PMapUint64Float64PtrErr(plusOneUint64Float64PtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Float64PtrErr failed")
	}

	_, err = PMapUint64Float64PtrErr(plusOneUint64Float64PtrErr, []*uint64{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Float64PtrErr failed")
	}

	_, err = PMapUint64Float64PtrErr(plusOneUint64Float64PtrErr, []*uint64{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Float64PtrErr failed")
	}

	_, err = PMapUint64Float64PtrErr(plusOneUint64Float64PtrErr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint64Float64PtrErr failed")
	}
}

func TestPmapUint32IntPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3

	expectedList := []*int{&vo2, &vo3}
	newList, _ := PMapUint32IntPtrErr(plusOneUint32IntPtrErr, []*uint32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint32IntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint32IntPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint32IntPtrErr failed")
	}

	r, _ = PMapUint32IntPtrErr(nil, []*uint32{})
	if len(r) > 0 {
		t.Errorf("PMapUint32IntPtrErr failed")
	}
	_, err := PMapUint32IntPtrErr(plusOneUint32IntPtrErr, []*uint32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint32IntPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint32IntPtrErr(plusOneUint32IntPtrErr, []*uint32{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32IntPtrErr failed")
	}

	_, err = PMapUint32IntPtrErr(plusOneUint32IntPtrErr, []*uint32{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint32IntPtrErr failed")
	}

	_, err = PMapUint32IntPtrErr(plusOneUint32IntPtrErr, []*uint32{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32IntPtrErr failed")
	}
	_, err = PMapUint32IntPtrErr(plusOneUint32IntPtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32IntPtrErr failed")
	}

	expectedList = []*int{&vo2, &vo3}
	newList, _ = PMapUint32IntPtrErr(plusOneUint32IntPtrErr, []*uint32{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint32IntPtr failed")
	}

	_, err = PMapUint32IntPtrErr(plusOneUint32IntPtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32IntPtrErr failed")
	}
	_, err = PMapUint32IntPtrErr(plusOneUint32IntPtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32IntPtrErr failed")
	}

	_, err = PMapUint32IntPtrErr(plusOneUint32IntPtrErr, []*uint32{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32IntPtrErr failed")
	}

	_, err = PMapUint32IntPtrErr(plusOneUint32IntPtrErr, []*uint32{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32IntPtrErr failed")
	}

	_, err = PMapUint32IntPtrErr(plusOneUint32IntPtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32IntPtrErr failed")
	}
}

func TestPmapUint32Int64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3

	expectedList := []*int64{&vo2, &vo3}
	newList, _ := PMapUint32Int64PtrErr(plusOneUint32Int64PtrErr, []*uint32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint32Int64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint32Int64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint32Int64PtrErr failed")
	}

	r, _ = PMapUint32Int64PtrErr(nil, []*uint32{})
	if len(r) > 0 {
		t.Errorf("PMapUint32Int64PtrErr failed")
	}
	_, err := PMapUint32Int64PtrErr(plusOneUint32Int64PtrErr, []*uint32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint32Int64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint32Int64PtrErr(plusOneUint32Int64PtrErr, []*uint32{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32Int64PtrErr failed")
	}

	_, err = PMapUint32Int64PtrErr(plusOneUint32Int64PtrErr, []*uint32{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint32Int64PtrErr failed")
	}

	_, err = PMapUint32Int64PtrErr(plusOneUint32Int64PtrErr, []*uint32{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32Int64PtrErr failed")
	}
	_, err = PMapUint32Int64PtrErr(plusOneUint32Int64PtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32Int64PtrErr failed")
	}

	expectedList = []*int64{&vo2, &vo3}
	newList, _ = PMapUint32Int64PtrErr(plusOneUint32Int64PtrErr, []*uint32{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint32Int64Ptr failed")
	}

	_, err = PMapUint32Int64PtrErr(plusOneUint32Int64PtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Int64PtrErr failed")
	}
	_, err = PMapUint32Int64PtrErr(plusOneUint32Int64PtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Int64PtrErr failed")
	}

	_, err = PMapUint32Int64PtrErr(plusOneUint32Int64PtrErr, []*uint32{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Int64PtrErr failed")
	}

	_, err = PMapUint32Int64PtrErr(plusOneUint32Int64PtrErr, []*uint32{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Int64PtrErr failed")
	}

	_, err = PMapUint32Int64PtrErr(plusOneUint32Int64PtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Int64PtrErr failed")
	}
}

func TestPmapUint32Int32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3

	expectedList := []*int32{&vo2, &vo3}
	newList, _ := PMapUint32Int32PtrErr(plusOneUint32Int32PtrErr, []*uint32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint32Int32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint32Int32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint32Int32PtrErr failed")
	}

	r, _ = PMapUint32Int32PtrErr(nil, []*uint32{})
	if len(r) > 0 {
		t.Errorf("PMapUint32Int32PtrErr failed")
	}
	_, err := PMapUint32Int32PtrErr(plusOneUint32Int32PtrErr, []*uint32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint32Int32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint32Int32PtrErr(plusOneUint32Int32PtrErr, []*uint32{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32Int32PtrErr failed")
	}

	_, err = PMapUint32Int32PtrErr(plusOneUint32Int32PtrErr, []*uint32{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint32Int32PtrErr failed")
	}

	_, err = PMapUint32Int32PtrErr(plusOneUint32Int32PtrErr, []*uint32{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32Int32PtrErr failed")
	}
	_, err = PMapUint32Int32PtrErr(plusOneUint32Int32PtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32Int32PtrErr failed")
	}

	expectedList = []*int32{&vo2, &vo3}
	newList, _ = PMapUint32Int32PtrErr(plusOneUint32Int32PtrErr, []*uint32{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint32Int32Ptr failed")
	}

	_, err = PMapUint32Int32PtrErr(plusOneUint32Int32PtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Int32PtrErr failed")
	}
	_, err = PMapUint32Int32PtrErr(plusOneUint32Int32PtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Int32PtrErr failed")
	}

	_, err = PMapUint32Int32PtrErr(plusOneUint32Int32PtrErr, []*uint32{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Int32PtrErr failed")
	}

	_, err = PMapUint32Int32PtrErr(plusOneUint32Int32PtrErr, []*uint32{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Int32PtrErr failed")
	}

	_, err = PMapUint32Int32PtrErr(plusOneUint32Int32PtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Int32PtrErr failed")
	}
}

func TestPmapUint32Int16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3

	expectedList := []*int16{&vo2, &vo3}
	newList, _ := PMapUint32Int16PtrErr(plusOneUint32Int16PtrErr, []*uint32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint32Int16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint32Int16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint32Int16PtrErr failed")
	}

	r, _ = PMapUint32Int16PtrErr(nil, []*uint32{})
	if len(r) > 0 {
		t.Errorf("PMapUint32Int16PtrErr failed")
	}
	_, err := PMapUint32Int16PtrErr(plusOneUint32Int16PtrErr, []*uint32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint32Int16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint32Int16PtrErr(plusOneUint32Int16PtrErr, []*uint32{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32Int16PtrErr failed")
	}

	_, err = PMapUint32Int16PtrErr(plusOneUint32Int16PtrErr, []*uint32{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint32Int16PtrErr failed")
	}

	_, err = PMapUint32Int16PtrErr(plusOneUint32Int16PtrErr, []*uint32{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32Int16PtrErr failed")
	}
	_, err = PMapUint32Int16PtrErr(plusOneUint32Int16PtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32Int16PtrErr failed")
	}

	expectedList = []*int16{&vo2, &vo3}
	newList, _ = PMapUint32Int16PtrErr(plusOneUint32Int16PtrErr, []*uint32{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint32Int16Ptr failed")
	}

	_, err = PMapUint32Int16PtrErr(plusOneUint32Int16PtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Int16PtrErr failed")
	}
	_, err = PMapUint32Int16PtrErr(plusOneUint32Int16PtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Int16PtrErr failed")
	}

	_, err = PMapUint32Int16PtrErr(plusOneUint32Int16PtrErr, []*uint32{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Int16PtrErr failed")
	}

	_, err = PMapUint32Int16PtrErr(plusOneUint32Int16PtrErr, []*uint32{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Int16PtrErr failed")
	}

	_, err = PMapUint32Int16PtrErr(plusOneUint32Int16PtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Int16PtrErr failed")
	}
}

func TestPmapUint32Int8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3

	expectedList := []*int8{&vo2, &vo3}
	newList, _ := PMapUint32Int8PtrErr(plusOneUint32Int8PtrErr, []*uint32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint32Int8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint32Int8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint32Int8PtrErr failed")
	}

	r, _ = PMapUint32Int8PtrErr(nil, []*uint32{})
	if len(r) > 0 {
		t.Errorf("PMapUint32Int8PtrErr failed")
	}
	_, err := PMapUint32Int8PtrErr(plusOneUint32Int8PtrErr, []*uint32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint32Int8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint32Int8PtrErr(plusOneUint32Int8PtrErr, []*uint32{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32Int8PtrErr failed")
	}

	_, err = PMapUint32Int8PtrErr(plusOneUint32Int8PtrErr, []*uint32{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint32Int8PtrErr failed")
	}

	_, err = PMapUint32Int8PtrErr(plusOneUint32Int8PtrErr, []*uint32{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32Int8PtrErr failed")
	}
	_, err = PMapUint32Int8PtrErr(plusOneUint32Int8PtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32Int8PtrErr failed")
	}

	expectedList = []*int8{&vo2, &vo3}
	newList, _ = PMapUint32Int8PtrErr(plusOneUint32Int8PtrErr, []*uint32{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint32Int8Ptr failed")
	}

	_, err = PMapUint32Int8PtrErr(plusOneUint32Int8PtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Int8PtrErr failed")
	}
	_, err = PMapUint32Int8PtrErr(plusOneUint32Int8PtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Int8PtrErr failed")
	}

	_, err = PMapUint32Int8PtrErr(plusOneUint32Int8PtrErr, []*uint32{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Int8PtrErr failed")
	}

	_, err = PMapUint32Int8PtrErr(plusOneUint32Int8PtrErr, []*uint32{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Int8PtrErr failed")
	}

	_, err = PMapUint32Int8PtrErr(plusOneUint32Int8PtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Int8PtrErr failed")
	}
}

func TestPmapUint32UintPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3

	expectedList := []*uint{&vo2, &vo3}
	newList, _ := PMapUint32UintPtrErr(plusOneUint32UintPtrErr, []*uint32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint32UintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint32UintPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint32UintPtrErr failed")
	}

	r, _ = PMapUint32UintPtrErr(nil, []*uint32{})
	if len(r) > 0 {
		t.Errorf("PMapUint32UintPtrErr failed")
	}
	_, err := PMapUint32UintPtrErr(plusOneUint32UintPtrErr, []*uint32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint32UintPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint32UintPtrErr(plusOneUint32UintPtrErr, []*uint32{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32UintPtrErr failed")
	}

	_, err = PMapUint32UintPtrErr(plusOneUint32UintPtrErr, []*uint32{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint32UintPtrErr failed")
	}

	_, err = PMapUint32UintPtrErr(plusOneUint32UintPtrErr, []*uint32{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32UintPtrErr failed")
	}
	_, err = PMapUint32UintPtrErr(plusOneUint32UintPtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32UintPtrErr failed")
	}

	expectedList = []*uint{&vo2, &vo3}
	newList, _ = PMapUint32UintPtrErr(plusOneUint32UintPtrErr, []*uint32{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint32UintPtr failed")
	}

	_, err = PMapUint32UintPtrErr(plusOneUint32UintPtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32UintPtrErr failed")
	}
	_, err = PMapUint32UintPtrErr(plusOneUint32UintPtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32UintPtrErr failed")
	}

	_, err = PMapUint32UintPtrErr(plusOneUint32UintPtrErr, []*uint32{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32UintPtrErr failed")
	}

	_, err = PMapUint32UintPtrErr(plusOneUint32UintPtrErr, []*uint32{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32UintPtrErr failed")
	}

	_, err = PMapUint32UintPtrErr(plusOneUint32UintPtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32UintPtrErr failed")
	}
}

func TestPmapUint32Uint64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3

	expectedList := []*uint64{&vo2, &vo3}
	newList, _ := PMapUint32Uint64PtrErr(plusOneUint32Uint64PtrErr, []*uint32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint32Uint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint32Uint64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint32Uint64PtrErr failed")
	}

	r, _ = PMapUint32Uint64PtrErr(nil, []*uint32{})
	if len(r) > 0 {
		t.Errorf("PMapUint32Uint64PtrErr failed")
	}
	_, err := PMapUint32Uint64PtrErr(plusOneUint32Uint64PtrErr, []*uint32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint32Uint64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint32Uint64PtrErr(plusOneUint32Uint64PtrErr, []*uint32{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32Uint64PtrErr failed")
	}

	_, err = PMapUint32Uint64PtrErr(plusOneUint32Uint64PtrErr, []*uint32{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint32Uint64PtrErr failed")
	}

	_, err = PMapUint32Uint64PtrErr(plusOneUint32Uint64PtrErr, []*uint32{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32Uint64PtrErr failed")
	}
	_, err = PMapUint32Uint64PtrErr(plusOneUint32Uint64PtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32Uint64PtrErr failed")
	}

	expectedList = []*uint64{&vo2, &vo3}
	newList, _ = PMapUint32Uint64PtrErr(plusOneUint32Uint64PtrErr, []*uint32{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint32Uint64Ptr failed")
	}

	_, err = PMapUint32Uint64PtrErr(plusOneUint32Uint64PtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Uint64PtrErr failed")
	}
	_, err = PMapUint32Uint64PtrErr(plusOneUint32Uint64PtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Uint64PtrErr failed")
	}

	_, err = PMapUint32Uint64PtrErr(plusOneUint32Uint64PtrErr, []*uint32{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Uint64PtrErr failed")
	}

	_, err = PMapUint32Uint64PtrErr(plusOneUint32Uint64PtrErr, []*uint32{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Uint64PtrErr failed")
	}

	_, err = PMapUint32Uint64PtrErr(plusOneUint32Uint64PtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Uint64PtrErr failed")
	}
}

func TestPmapUint32Uint16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3

	expectedList := []*uint16{&vo2, &vo3}
	newList, _ := PMapUint32Uint16PtrErr(plusOneUint32Uint16PtrErr, []*uint32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint32Uint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint32Uint16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint32Uint16PtrErr failed")
	}

	r, _ = PMapUint32Uint16PtrErr(nil, []*uint32{})
	if len(r) > 0 {
		t.Errorf("PMapUint32Uint16PtrErr failed")
	}
	_, err := PMapUint32Uint16PtrErr(plusOneUint32Uint16PtrErr, []*uint32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint32Uint16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint32Uint16PtrErr(plusOneUint32Uint16PtrErr, []*uint32{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32Uint16PtrErr failed")
	}

	_, err = PMapUint32Uint16PtrErr(plusOneUint32Uint16PtrErr, []*uint32{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint32Uint16PtrErr failed")
	}

	_, err = PMapUint32Uint16PtrErr(plusOneUint32Uint16PtrErr, []*uint32{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32Uint16PtrErr failed")
	}
	_, err = PMapUint32Uint16PtrErr(plusOneUint32Uint16PtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32Uint16PtrErr failed")
	}

	expectedList = []*uint16{&vo2, &vo3}
	newList, _ = PMapUint32Uint16PtrErr(plusOneUint32Uint16PtrErr, []*uint32{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint32Uint16Ptr failed")
	}

	_, err = PMapUint32Uint16PtrErr(plusOneUint32Uint16PtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Uint16PtrErr failed")
	}
	_, err = PMapUint32Uint16PtrErr(plusOneUint32Uint16PtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Uint16PtrErr failed")
	}

	_, err = PMapUint32Uint16PtrErr(plusOneUint32Uint16PtrErr, []*uint32{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Uint16PtrErr failed")
	}

	_, err = PMapUint32Uint16PtrErr(plusOneUint32Uint16PtrErr, []*uint32{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Uint16PtrErr failed")
	}

	_, err = PMapUint32Uint16PtrErr(plusOneUint32Uint16PtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Uint16PtrErr failed")
	}
}

func TestPmapUint32Uint8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3

	expectedList := []*uint8{&vo2, &vo3}
	newList, _ := PMapUint32Uint8PtrErr(plusOneUint32Uint8PtrErr, []*uint32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint32Uint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint32Uint8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint32Uint8PtrErr failed")
	}

	r, _ = PMapUint32Uint8PtrErr(nil, []*uint32{})
	if len(r) > 0 {
		t.Errorf("PMapUint32Uint8PtrErr failed")
	}
	_, err := PMapUint32Uint8PtrErr(plusOneUint32Uint8PtrErr, []*uint32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint32Uint8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint32Uint8PtrErr(plusOneUint32Uint8PtrErr, []*uint32{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32Uint8PtrErr failed")
	}

	_, err = PMapUint32Uint8PtrErr(plusOneUint32Uint8PtrErr, []*uint32{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint32Uint8PtrErr failed")
	}

	_, err = PMapUint32Uint8PtrErr(plusOneUint32Uint8PtrErr, []*uint32{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32Uint8PtrErr failed")
	}
	_, err = PMapUint32Uint8PtrErr(plusOneUint32Uint8PtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32Uint8PtrErr failed")
	}

	expectedList = []*uint8{&vo2, &vo3}
	newList, _ = PMapUint32Uint8PtrErr(plusOneUint32Uint8PtrErr, []*uint32{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint32Uint8Ptr failed")
	}

	_, err = PMapUint32Uint8PtrErr(plusOneUint32Uint8PtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Uint8PtrErr failed")
	}
	_, err = PMapUint32Uint8PtrErr(plusOneUint32Uint8PtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Uint8PtrErr failed")
	}

	_, err = PMapUint32Uint8PtrErr(plusOneUint32Uint8PtrErr, []*uint32{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Uint8PtrErr failed")
	}

	_, err = PMapUint32Uint8PtrErr(plusOneUint32Uint8PtrErr, []*uint32{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Uint8PtrErr failed")
	}

	_, err = PMapUint32Uint8PtrErr(plusOneUint32Uint8PtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Uint8PtrErr failed")
	}
}

func TestPmapUint32StrPtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 uint32 = 10
	var vi0 uint32

	expectedList := []*string{&vo10}
	newList, _ := PMapUint32StrPtrErr(someLogicUint32StrPtrErr, []*uint32{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapUint32StrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint32StrPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint32StrPtrErr failed")
	}

	r, _ = PMapUint32StrPtrErr(nil, []*uint32{})
	if len(r) > 0 {
		t.Errorf("PMapUint32StrPtrErr failed")
	}
	_, err := PMapUint32StrPtrErr(someLogicUint32StrPtrErr, []*uint32{&vi10, &vi0})
	if err == nil {
		t.Errorf("PMapUint32StrPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint32StrPtrErr(someLogicUint32StrPtrErr, []*uint32{&vi10, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32StrPtrErr failed")
	}

	_, err = PMapUint32StrPtrErr(someLogicUint32StrPtrErr, []*uint32{&vi0, &vi0, &vi10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint32StrPtrErr failed")
	}

	_, err = PMapUint32StrPtrErr(someLogicUint32StrPtrErr, []*uint32{&vi0, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32StrPtrErr failed")
	}
	_, err = PMapUint32StrPtrErr(someLogicUint32StrPtrErr, []*uint32{&vi10, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32StrPtrErr failed")
	}

	expectedList = []*string{&vo10, &vo10}
	newList, _ = PMapUint32StrPtrErr(someLogicUint32StrPtrErr, []*uint32{&vi10, &vi10}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint32StrPtr failed")
	}

	_, err = PMapUint32StrPtrErr(someLogicUint32StrPtrErr, []*uint32{&vi10, &vi10, &vi0}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32StrPtrErr failed")
	}
	_, err = PMapUint32StrPtrErr(someLogicUint32StrPtrErr, []*uint32{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32StrPtrErr failed")
	}

	_, err = PMapUint32StrPtrErr(someLogicUint32StrPtrErr, []*uint32{&vi0, &vi0, &vi10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32StrPtrErr failed")
	}

	_, err = PMapUint32StrPtrErr(someLogicUint32StrPtrErr, []*uint32{&vi0, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32StrPtrErr failed")
	}

	_, err = PMapUint32StrPtrErr(someLogicUint32StrPtrErr, []*uint32{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32StrPtrErr failed")
	}
}

func TestPmapUint32BoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 uint32 = 10
	var v0 uint32
	var v3 uint32 = 3

	expectedList := []*bool{&vt, &vf}
	newList, _ := PMapUint32BoolPtrErr(someLogicUint32BoolPtrErr, []*uint32{&v10, &v0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint32BoolPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint32BoolPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint32BoolPtrErr failed")
	}

	r, _ = PMapUint32BoolPtrErr(nil, []*uint32{})
	if len(r) > 0 {
		t.Errorf("PMapUint32BoolPtrErr failed")
	}
	_, err := PMapUint32BoolPtrErr(someLogicUint32BoolPtrErr, []*uint32{&v10, &v3})
	if err == nil {
		t.Errorf("PMapUint32BoolPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint32BoolPtrErr(someLogicUint32BoolPtrErr, []*uint32{&v10, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32BoolPtrErr failed")
	}

	_, err = PMapUint32BoolPtrErr(someLogicUint32BoolPtrErr, []*uint32{&v3, &v3, &v10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint32BoolPtrErr failed")
	}

	_, err = PMapUint32BoolPtrErr(someLogicUint32BoolPtrErr, []*uint32{&v3, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32BoolErr failed")
	}
	_, err = PMapUint32BoolPtrErr(someLogicUint32BoolPtrErr, []*uint32{&v10, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32BoolPtrErr failed")
	}

	expectedList = []*bool{&vt, &vt}
	newList, _ = PMapUint32BoolPtrErr(someLogicUint32BoolPtrErr, []*uint32{&v10, &v10}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint32BoolPtr failed")
	}

	_, err = PMapUint32BoolPtrErr(someLogicUint32BoolPtrErr, []*uint32{&v10, &v10, &v3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32BoolErr failed")
	}
	_, err = PMapUint32BoolPtrErr(someLogicUint32BoolPtrErr, []*uint32{&v10, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32BoolPtrErr failed")
	}

	_, err = PMapUint32BoolPtrErr(someLogicUint32BoolPtrErr, []*uint32{&v3, &v3, &v10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32BoolPtrErr failed")
	}

	_, err = PMapUint32BoolPtrErr(someLogicUint32BoolPtrErr, []*uint32{&v3, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32BoolPtrErr failed")
	}

	_, err = PMapUint32BoolPtrErr(someLogicUint32BoolPtrErr, []*uint32{&v10, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32BoolPtrErr failed")
	}
}

func TestPmapUint32Float32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3

	expectedList := []*float32{&vo2, &vo3}
	newList, _ := PMapUint32Float32PtrErr(plusOneUint32Float32PtrErr, []*uint32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint32Float32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint32Float32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint32Float32PtrErr failed")
	}

	r, _ = PMapUint32Float32PtrErr(nil, []*uint32{})
	if len(r) > 0 {
		t.Errorf("PMapUint32Float32PtrErr failed")
	}
	_, err := PMapUint32Float32PtrErr(plusOneUint32Float32PtrErr, []*uint32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint32Float32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint32Float32PtrErr(plusOneUint32Float32PtrErr, []*uint32{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32Float32PtrErr failed")
	}

	_, err = PMapUint32Float32PtrErr(plusOneUint32Float32PtrErr, []*uint32{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint32Float32PtrErr failed")
	}

	_, err = PMapUint32Float32PtrErr(plusOneUint32Float32PtrErr, []*uint32{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32Float32PtrErr failed")
	}
	_, err = PMapUint32Float32PtrErr(plusOneUint32Float32PtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32Float32PtrErr failed")
	}

	expectedList = []*float32{&vo2, &vo3}
	newList, _ = PMapUint32Float32PtrErr(plusOneUint32Float32PtrErr, []*uint32{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint32Float32Ptr failed")
	}

	_, err = PMapUint32Float32PtrErr(plusOneUint32Float32PtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Float32PtrErr failed")
	}
	_, err = PMapUint32Float32PtrErr(plusOneUint32Float32PtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Float32PtrErr failed")
	}

	_, err = PMapUint32Float32PtrErr(plusOneUint32Float32PtrErr, []*uint32{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Float32PtrErr failed")
	}

	_, err = PMapUint32Float32PtrErr(plusOneUint32Float32PtrErr, []*uint32{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Float32PtrErr failed")
	}

	_, err = PMapUint32Float32PtrErr(plusOneUint32Float32PtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Float32PtrErr failed")
	}
}

func TestPmapUint32Float64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3

	expectedList := []*float64{&vo2, &vo3}
	newList, _ := PMapUint32Float64PtrErr(plusOneUint32Float64PtrErr, []*uint32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint32Float64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint32Float64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint32Float64PtrErr failed")
	}

	r, _ = PMapUint32Float64PtrErr(nil, []*uint32{})
	if len(r) > 0 {
		t.Errorf("PMapUint32Float64PtrErr failed")
	}
	_, err := PMapUint32Float64PtrErr(plusOneUint32Float64PtrErr, []*uint32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint32Float64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint32Float64PtrErr(plusOneUint32Float64PtrErr, []*uint32{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32Float64PtrErr failed")
	}

	_, err = PMapUint32Float64PtrErr(plusOneUint32Float64PtrErr, []*uint32{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint32Float64PtrErr failed")
	}

	_, err = PMapUint32Float64PtrErr(plusOneUint32Float64PtrErr, []*uint32{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32Float64PtrErr failed")
	}
	_, err = PMapUint32Float64PtrErr(plusOneUint32Float64PtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32Float64PtrErr failed")
	}

	expectedList = []*float64{&vo2, &vo3}
	newList, _ = PMapUint32Float64PtrErr(plusOneUint32Float64PtrErr, []*uint32{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint32Float64Ptr failed")
	}

	_, err = PMapUint32Float64PtrErr(plusOneUint32Float64PtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Float64PtrErr failed")
	}
	_, err = PMapUint32Float64PtrErr(plusOneUint32Float64PtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Float64PtrErr failed")
	}

	_, err = PMapUint32Float64PtrErr(plusOneUint32Float64PtrErr, []*uint32{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Float64PtrErr failed")
	}

	_, err = PMapUint32Float64PtrErr(plusOneUint32Float64PtrErr, []*uint32{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Float64PtrErr failed")
	}

	_, err = PMapUint32Float64PtrErr(plusOneUint32Float64PtrErr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint32Float64PtrErr failed")
	}
}

func TestPmapUint16IntPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3

	expectedList := []*int{&vo2, &vo3}
	newList, _ := PMapUint16IntPtrErr(plusOneUint16IntPtrErr, []*uint16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint16IntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint16IntPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint16IntPtrErr failed")
	}

	r, _ = PMapUint16IntPtrErr(nil, []*uint16{})
	if len(r) > 0 {
		t.Errorf("PMapUint16IntPtrErr failed")
	}
	_, err := PMapUint16IntPtrErr(plusOneUint16IntPtrErr, []*uint16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint16IntPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint16IntPtrErr(plusOneUint16IntPtrErr, []*uint16{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16IntPtrErr failed")
	}

	_, err = PMapUint16IntPtrErr(plusOneUint16IntPtrErr, []*uint16{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint16IntPtrErr failed")
	}

	_, err = PMapUint16IntPtrErr(plusOneUint16IntPtrErr, []*uint16{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16IntPtrErr failed")
	}
	_, err = PMapUint16IntPtrErr(plusOneUint16IntPtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16IntPtrErr failed")
	}

	expectedList = []*int{&vo2, &vo3}
	newList, _ = PMapUint16IntPtrErr(plusOneUint16IntPtrErr, []*uint16{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint16IntPtr failed")
	}

	_, err = PMapUint16IntPtrErr(plusOneUint16IntPtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16IntPtrErr failed")
	}
	_, err = PMapUint16IntPtrErr(plusOneUint16IntPtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16IntPtrErr failed")
	}

	_, err = PMapUint16IntPtrErr(plusOneUint16IntPtrErr, []*uint16{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16IntPtrErr failed")
	}

	_, err = PMapUint16IntPtrErr(plusOneUint16IntPtrErr, []*uint16{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16IntPtrErr failed")
	}

	_, err = PMapUint16IntPtrErr(plusOneUint16IntPtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16IntPtrErr failed")
	}
}

func TestPmapUint16Int64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3

	expectedList := []*int64{&vo2, &vo3}
	newList, _ := PMapUint16Int64PtrErr(plusOneUint16Int64PtrErr, []*uint16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint16Int64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint16Int64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint16Int64PtrErr failed")
	}

	r, _ = PMapUint16Int64PtrErr(nil, []*uint16{})
	if len(r) > 0 {
		t.Errorf("PMapUint16Int64PtrErr failed")
	}
	_, err := PMapUint16Int64PtrErr(plusOneUint16Int64PtrErr, []*uint16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint16Int64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint16Int64PtrErr(plusOneUint16Int64PtrErr, []*uint16{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16Int64PtrErr failed")
	}

	_, err = PMapUint16Int64PtrErr(plusOneUint16Int64PtrErr, []*uint16{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint16Int64PtrErr failed")
	}

	_, err = PMapUint16Int64PtrErr(plusOneUint16Int64PtrErr, []*uint16{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16Int64PtrErr failed")
	}
	_, err = PMapUint16Int64PtrErr(plusOneUint16Int64PtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16Int64PtrErr failed")
	}

	expectedList = []*int64{&vo2, &vo3}
	newList, _ = PMapUint16Int64PtrErr(plusOneUint16Int64PtrErr, []*uint16{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint16Int64Ptr failed")
	}

	_, err = PMapUint16Int64PtrErr(plusOneUint16Int64PtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Int64PtrErr failed")
	}
	_, err = PMapUint16Int64PtrErr(plusOneUint16Int64PtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Int64PtrErr failed")
	}

	_, err = PMapUint16Int64PtrErr(plusOneUint16Int64PtrErr, []*uint16{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Int64PtrErr failed")
	}

	_, err = PMapUint16Int64PtrErr(plusOneUint16Int64PtrErr, []*uint16{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Int64PtrErr failed")
	}

	_, err = PMapUint16Int64PtrErr(plusOneUint16Int64PtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Int64PtrErr failed")
	}
}

func TestPmapUint16Int32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3

	expectedList := []*int32{&vo2, &vo3}
	newList, _ := PMapUint16Int32PtrErr(plusOneUint16Int32PtrErr, []*uint16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint16Int32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint16Int32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint16Int32PtrErr failed")
	}

	r, _ = PMapUint16Int32PtrErr(nil, []*uint16{})
	if len(r) > 0 {
		t.Errorf("PMapUint16Int32PtrErr failed")
	}
	_, err := PMapUint16Int32PtrErr(plusOneUint16Int32PtrErr, []*uint16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint16Int32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint16Int32PtrErr(plusOneUint16Int32PtrErr, []*uint16{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16Int32PtrErr failed")
	}

	_, err = PMapUint16Int32PtrErr(plusOneUint16Int32PtrErr, []*uint16{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint16Int32PtrErr failed")
	}

	_, err = PMapUint16Int32PtrErr(plusOneUint16Int32PtrErr, []*uint16{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16Int32PtrErr failed")
	}
	_, err = PMapUint16Int32PtrErr(plusOneUint16Int32PtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16Int32PtrErr failed")
	}

	expectedList = []*int32{&vo2, &vo3}
	newList, _ = PMapUint16Int32PtrErr(plusOneUint16Int32PtrErr, []*uint16{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint16Int32Ptr failed")
	}

	_, err = PMapUint16Int32PtrErr(plusOneUint16Int32PtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Int32PtrErr failed")
	}
	_, err = PMapUint16Int32PtrErr(plusOneUint16Int32PtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Int32PtrErr failed")
	}

	_, err = PMapUint16Int32PtrErr(plusOneUint16Int32PtrErr, []*uint16{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Int32PtrErr failed")
	}

	_, err = PMapUint16Int32PtrErr(plusOneUint16Int32PtrErr, []*uint16{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Int32PtrErr failed")
	}

	_, err = PMapUint16Int32PtrErr(plusOneUint16Int32PtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Int32PtrErr failed")
	}
}

func TestPmapUint16Int16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3

	expectedList := []*int16{&vo2, &vo3}
	newList, _ := PMapUint16Int16PtrErr(plusOneUint16Int16PtrErr, []*uint16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint16Int16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint16Int16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint16Int16PtrErr failed")
	}

	r, _ = PMapUint16Int16PtrErr(nil, []*uint16{})
	if len(r) > 0 {
		t.Errorf("PMapUint16Int16PtrErr failed")
	}
	_, err := PMapUint16Int16PtrErr(plusOneUint16Int16PtrErr, []*uint16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint16Int16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint16Int16PtrErr(plusOneUint16Int16PtrErr, []*uint16{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16Int16PtrErr failed")
	}

	_, err = PMapUint16Int16PtrErr(plusOneUint16Int16PtrErr, []*uint16{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint16Int16PtrErr failed")
	}

	_, err = PMapUint16Int16PtrErr(plusOneUint16Int16PtrErr, []*uint16{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16Int16PtrErr failed")
	}
	_, err = PMapUint16Int16PtrErr(plusOneUint16Int16PtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16Int16PtrErr failed")
	}

	expectedList = []*int16{&vo2, &vo3}
	newList, _ = PMapUint16Int16PtrErr(plusOneUint16Int16PtrErr, []*uint16{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint16Int16Ptr failed")
	}

	_, err = PMapUint16Int16PtrErr(plusOneUint16Int16PtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Int16PtrErr failed")
	}
	_, err = PMapUint16Int16PtrErr(plusOneUint16Int16PtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Int16PtrErr failed")
	}

	_, err = PMapUint16Int16PtrErr(plusOneUint16Int16PtrErr, []*uint16{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Int16PtrErr failed")
	}

	_, err = PMapUint16Int16PtrErr(plusOneUint16Int16PtrErr, []*uint16{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Int16PtrErr failed")
	}

	_, err = PMapUint16Int16PtrErr(plusOneUint16Int16PtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Int16PtrErr failed")
	}
}

func TestPmapUint16Int8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3

	expectedList := []*int8{&vo2, &vo3}
	newList, _ := PMapUint16Int8PtrErr(plusOneUint16Int8PtrErr, []*uint16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint16Int8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint16Int8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint16Int8PtrErr failed")
	}

	r, _ = PMapUint16Int8PtrErr(nil, []*uint16{})
	if len(r) > 0 {
		t.Errorf("PMapUint16Int8PtrErr failed")
	}
	_, err := PMapUint16Int8PtrErr(plusOneUint16Int8PtrErr, []*uint16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint16Int8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint16Int8PtrErr(plusOneUint16Int8PtrErr, []*uint16{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16Int8PtrErr failed")
	}

	_, err = PMapUint16Int8PtrErr(plusOneUint16Int8PtrErr, []*uint16{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint16Int8PtrErr failed")
	}

	_, err = PMapUint16Int8PtrErr(plusOneUint16Int8PtrErr, []*uint16{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16Int8PtrErr failed")
	}
	_, err = PMapUint16Int8PtrErr(plusOneUint16Int8PtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16Int8PtrErr failed")
	}

	expectedList = []*int8{&vo2, &vo3}
	newList, _ = PMapUint16Int8PtrErr(plusOneUint16Int8PtrErr, []*uint16{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint16Int8Ptr failed")
	}

	_, err = PMapUint16Int8PtrErr(plusOneUint16Int8PtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Int8PtrErr failed")
	}
	_, err = PMapUint16Int8PtrErr(plusOneUint16Int8PtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Int8PtrErr failed")
	}

	_, err = PMapUint16Int8PtrErr(plusOneUint16Int8PtrErr, []*uint16{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Int8PtrErr failed")
	}

	_, err = PMapUint16Int8PtrErr(plusOneUint16Int8PtrErr, []*uint16{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Int8PtrErr failed")
	}

	_, err = PMapUint16Int8PtrErr(plusOneUint16Int8PtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Int8PtrErr failed")
	}
}

func TestPmapUint16UintPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3

	expectedList := []*uint{&vo2, &vo3}
	newList, _ := PMapUint16UintPtrErr(plusOneUint16UintPtrErr, []*uint16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint16UintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint16UintPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint16UintPtrErr failed")
	}

	r, _ = PMapUint16UintPtrErr(nil, []*uint16{})
	if len(r) > 0 {
		t.Errorf("PMapUint16UintPtrErr failed")
	}
	_, err := PMapUint16UintPtrErr(plusOneUint16UintPtrErr, []*uint16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint16UintPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint16UintPtrErr(plusOneUint16UintPtrErr, []*uint16{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16UintPtrErr failed")
	}

	_, err = PMapUint16UintPtrErr(plusOneUint16UintPtrErr, []*uint16{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint16UintPtrErr failed")
	}

	_, err = PMapUint16UintPtrErr(plusOneUint16UintPtrErr, []*uint16{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16UintPtrErr failed")
	}
	_, err = PMapUint16UintPtrErr(plusOneUint16UintPtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16UintPtrErr failed")
	}

	expectedList = []*uint{&vo2, &vo3}
	newList, _ = PMapUint16UintPtrErr(plusOneUint16UintPtrErr, []*uint16{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint16UintPtr failed")
	}

	_, err = PMapUint16UintPtrErr(plusOneUint16UintPtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16UintPtrErr failed")
	}
	_, err = PMapUint16UintPtrErr(plusOneUint16UintPtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16UintPtrErr failed")
	}

	_, err = PMapUint16UintPtrErr(plusOneUint16UintPtrErr, []*uint16{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16UintPtrErr failed")
	}

	_, err = PMapUint16UintPtrErr(plusOneUint16UintPtrErr, []*uint16{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16UintPtrErr failed")
	}

	_, err = PMapUint16UintPtrErr(plusOneUint16UintPtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16UintPtrErr failed")
	}
}

func TestPmapUint16Uint64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3

	expectedList := []*uint64{&vo2, &vo3}
	newList, _ := PMapUint16Uint64PtrErr(plusOneUint16Uint64PtrErr, []*uint16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint16Uint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint16Uint64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint16Uint64PtrErr failed")
	}

	r, _ = PMapUint16Uint64PtrErr(nil, []*uint16{})
	if len(r) > 0 {
		t.Errorf("PMapUint16Uint64PtrErr failed")
	}
	_, err := PMapUint16Uint64PtrErr(plusOneUint16Uint64PtrErr, []*uint16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint16Uint64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint16Uint64PtrErr(plusOneUint16Uint64PtrErr, []*uint16{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16Uint64PtrErr failed")
	}

	_, err = PMapUint16Uint64PtrErr(plusOneUint16Uint64PtrErr, []*uint16{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint16Uint64PtrErr failed")
	}

	_, err = PMapUint16Uint64PtrErr(plusOneUint16Uint64PtrErr, []*uint16{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16Uint64PtrErr failed")
	}
	_, err = PMapUint16Uint64PtrErr(plusOneUint16Uint64PtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16Uint64PtrErr failed")
	}

	expectedList = []*uint64{&vo2, &vo3}
	newList, _ = PMapUint16Uint64PtrErr(plusOneUint16Uint64PtrErr, []*uint16{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint16Uint64Ptr failed")
	}

	_, err = PMapUint16Uint64PtrErr(plusOneUint16Uint64PtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Uint64PtrErr failed")
	}
	_, err = PMapUint16Uint64PtrErr(plusOneUint16Uint64PtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Uint64PtrErr failed")
	}

	_, err = PMapUint16Uint64PtrErr(plusOneUint16Uint64PtrErr, []*uint16{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Uint64PtrErr failed")
	}

	_, err = PMapUint16Uint64PtrErr(plusOneUint16Uint64PtrErr, []*uint16{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Uint64PtrErr failed")
	}

	_, err = PMapUint16Uint64PtrErr(plusOneUint16Uint64PtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Uint64PtrErr failed")
	}
}

func TestPmapUint16Uint32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3

	expectedList := []*uint32{&vo2, &vo3}
	newList, _ := PMapUint16Uint32PtrErr(plusOneUint16Uint32PtrErr, []*uint16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint16Uint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint16Uint32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint16Uint32PtrErr failed")
	}

	r, _ = PMapUint16Uint32PtrErr(nil, []*uint16{})
	if len(r) > 0 {
		t.Errorf("PMapUint16Uint32PtrErr failed")
	}
	_, err := PMapUint16Uint32PtrErr(plusOneUint16Uint32PtrErr, []*uint16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint16Uint32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint16Uint32PtrErr(plusOneUint16Uint32PtrErr, []*uint16{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16Uint32PtrErr failed")
	}

	_, err = PMapUint16Uint32PtrErr(plusOneUint16Uint32PtrErr, []*uint16{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint16Uint32PtrErr failed")
	}

	_, err = PMapUint16Uint32PtrErr(plusOneUint16Uint32PtrErr, []*uint16{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16Uint32PtrErr failed")
	}
	_, err = PMapUint16Uint32PtrErr(plusOneUint16Uint32PtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16Uint32PtrErr failed")
	}

	expectedList = []*uint32{&vo2, &vo3}
	newList, _ = PMapUint16Uint32PtrErr(plusOneUint16Uint32PtrErr, []*uint16{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint16Uint32Ptr failed")
	}

	_, err = PMapUint16Uint32PtrErr(plusOneUint16Uint32PtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Uint32PtrErr failed")
	}
	_, err = PMapUint16Uint32PtrErr(plusOneUint16Uint32PtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Uint32PtrErr failed")
	}

	_, err = PMapUint16Uint32PtrErr(plusOneUint16Uint32PtrErr, []*uint16{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Uint32PtrErr failed")
	}

	_, err = PMapUint16Uint32PtrErr(plusOneUint16Uint32PtrErr, []*uint16{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Uint32PtrErr failed")
	}

	_, err = PMapUint16Uint32PtrErr(plusOneUint16Uint32PtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Uint32PtrErr failed")
	}
}

func TestPmapUint16Uint8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3

	expectedList := []*uint8{&vo2, &vo3}
	newList, _ := PMapUint16Uint8PtrErr(plusOneUint16Uint8PtrErr, []*uint16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint16Uint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint16Uint8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint16Uint8PtrErr failed")
	}

	r, _ = PMapUint16Uint8PtrErr(nil, []*uint16{})
	if len(r) > 0 {
		t.Errorf("PMapUint16Uint8PtrErr failed")
	}
	_, err := PMapUint16Uint8PtrErr(plusOneUint16Uint8PtrErr, []*uint16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint16Uint8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint16Uint8PtrErr(plusOneUint16Uint8PtrErr, []*uint16{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16Uint8PtrErr failed")
	}

	_, err = PMapUint16Uint8PtrErr(plusOneUint16Uint8PtrErr, []*uint16{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint16Uint8PtrErr failed")
	}

	_, err = PMapUint16Uint8PtrErr(plusOneUint16Uint8PtrErr, []*uint16{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16Uint8PtrErr failed")
	}
	_, err = PMapUint16Uint8PtrErr(plusOneUint16Uint8PtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16Uint8PtrErr failed")
	}

	expectedList = []*uint8{&vo2, &vo3}
	newList, _ = PMapUint16Uint8PtrErr(plusOneUint16Uint8PtrErr, []*uint16{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint16Uint8Ptr failed")
	}

	_, err = PMapUint16Uint8PtrErr(plusOneUint16Uint8PtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Uint8PtrErr failed")
	}
	_, err = PMapUint16Uint8PtrErr(plusOneUint16Uint8PtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Uint8PtrErr failed")
	}

	_, err = PMapUint16Uint8PtrErr(plusOneUint16Uint8PtrErr, []*uint16{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Uint8PtrErr failed")
	}

	_, err = PMapUint16Uint8PtrErr(plusOneUint16Uint8PtrErr, []*uint16{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Uint8PtrErr failed")
	}

	_, err = PMapUint16Uint8PtrErr(plusOneUint16Uint8PtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Uint8PtrErr failed")
	}
}

func TestPmapUint16StrPtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 uint16 = 10
	var vi0 uint16

	expectedList := []*string{&vo10}
	newList, _ := PMapUint16StrPtrErr(someLogicUint16StrPtrErr, []*uint16{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapUint16StrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint16StrPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint16StrPtrErr failed")
	}

	r, _ = PMapUint16StrPtrErr(nil, []*uint16{})
	if len(r) > 0 {
		t.Errorf("PMapUint16StrPtrErr failed")
	}
	_, err := PMapUint16StrPtrErr(someLogicUint16StrPtrErr, []*uint16{&vi10, &vi0})
	if err == nil {
		t.Errorf("PMapUint16StrPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint16StrPtrErr(someLogicUint16StrPtrErr, []*uint16{&vi10, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16StrPtrErr failed")
	}

	_, err = PMapUint16StrPtrErr(someLogicUint16StrPtrErr, []*uint16{&vi0, &vi0, &vi10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint16StrPtrErr failed")
	}

	_, err = PMapUint16StrPtrErr(someLogicUint16StrPtrErr, []*uint16{&vi0, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16StrPtrErr failed")
	}
	_, err = PMapUint16StrPtrErr(someLogicUint16StrPtrErr, []*uint16{&vi10, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16StrPtrErr failed")
	}

	expectedList = []*string{&vo10, &vo10}
	newList, _ = PMapUint16StrPtrErr(someLogicUint16StrPtrErr, []*uint16{&vi10, &vi10}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint16StrPtr failed")
	}

	_, err = PMapUint16StrPtrErr(someLogicUint16StrPtrErr, []*uint16{&vi10, &vi10, &vi0}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16StrPtrErr failed")
	}
	_, err = PMapUint16StrPtrErr(someLogicUint16StrPtrErr, []*uint16{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16StrPtrErr failed")
	}

	_, err = PMapUint16StrPtrErr(someLogicUint16StrPtrErr, []*uint16{&vi0, &vi0, &vi10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16StrPtrErr failed")
	}

	_, err = PMapUint16StrPtrErr(someLogicUint16StrPtrErr, []*uint16{&vi0, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16StrPtrErr failed")
	}

	_, err = PMapUint16StrPtrErr(someLogicUint16StrPtrErr, []*uint16{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16StrPtrErr failed")
	}
}

func TestPmapUint16BoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 uint16 = 10
	var v0 uint16
	var v3 uint16 = 3

	expectedList := []*bool{&vt, &vf}
	newList, _ := PMapUint16BoolPtrErr(someLogicUint16BoolPtrErr, []*uint16{&v10, &v0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint16BoolPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint16BoolPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint16BoolPtrErr failed")
	}

	r, _ = PMapUint16BoolPtrErr(nil, []*uint16{})
	if len(r) > 0 {
		t.Errorf("PMapUint16BoolPtrErr failed")
	}
	_, err := PMapUint16BoolPtrErr(someLogicUint16BoolPtrErr, []*uint16{&v10, &v3})
	if err == nil {
		t.Errorf("PMapUint16BoolPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint16BoolPtrErr(someLogicUint16BoolPtrErr, []*uint16{&v10, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16BoolPtrErr failed")
	}

	_, err = PMapUint16BoolPtrErr(someLogicUint16BoolPtrErr, []*uint16{&v3, &v3, &v10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint16BoolPtrErr failed")
	}

	_, err = PMapUint16BoolPtrErr(someLogicUint16BoolPtrErr, []*uint16{&v3, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16BoolErr failed")
	}
	_, err = PMapUint16BoolPtrErr(someLogicUint16BoolPtrErr, []*uint16{&v10, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16BoolPtrErr failed")
	}

	expectedList = []*bool{&vt, &vt}
	newList, _ = PMapUint16BoolPtrErr(someLogicUint16BoolPtrErr, []*uint16{&v10, &v10}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint16BoolPtr failed")
	}

	_, err = PMapUint16BoolPtrErr(someLogicUint16BoolPtrErr, []*uint16{&v10, &v10, &v3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16BoolErr failed")
	}
	_, err = PMapUint16BoolPtrErr(someLogicUint16BoolPtrErr, []*uint16{&v10, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16BoolPtrErr failed")
	}

	_, err = PMapUint16BoolPtrErr(someLogicUint16BoolPtrErr, []*uint16{&v3, &v3, &v10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16BoolPtrErr failed")
	}

	_, err = PMapUint16BoolPtrErr(someLogicUint16BoolPtrErr, []*uint16{&v3, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16BoolPtrErr failed")
	}

	_, err = PMapUint16BoolPtrErr(someLogicUint16BoolPtrErr, []*uint16{&v10, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16BoolPtrErr failed")
	}
}

func TestPmapUint16Float32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3

	expectedList := []*float32{&vo2, &vo3}
	newList, _ := PMapUint16Float32PtrErr(plusOneUint16Float32PtrErr, []*uint16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint16Float32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint16Float32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint16Float32PtrErr failed")
	}

	r, _ = PMapUint16Float32PtrErr(nil, []*uint16{})
	if len(r) > 0 {
		t.Errorf("PMapUint16Float32PtrErr failed")
	}
	_, err := PMapUint16Float32PtrErr(plusOneUint16Float32PtrErr, []*uint16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint16Float32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint16Float32PtrErr(plusOneUint16Float32PtrErr, []*uint16{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16Float32PtrErr failed")
	}

	_, err = PMapUint16Float32PtrErr(plusOneUint16Float32PtrErr, []*uint16{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint16Float32PtrErr failed")
	}

	_, err = PMapUint16Float32PtrErr(plusOneUint16Float32PtrErr, []*uint16{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16Float32PtrErr failed")
	}
	_, err = PMapUint16Float32PtrErr(plusOneUint16Float32PtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16Float32PtrErr failed")
	}

	expectedList = []*float32{&vo2, &vo3}
	newList, _ = PMapUint16Float32PtrErr(plusOneUint16Float32PtrErr, []*uint16{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint16Float32Ptr failed")
	}

	_, err = PMapUint16Float32PtrErr(plusOneUint16Float32PtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Float32PtrErr failed")
	}
	_, err = PMapUint16Float32PtrErr(plusOneUint16Float32PtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Float32PtrErr failed")
	}

	_, err = PMapUint16Float32PtrErr(plusOneUint16Float32PtrErr, []*uint16{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Float32PtrErr failed")
	}

	_, err = PMapUint16Float32PtrErr(plusOneUint16Float32PtrErr, []*uint16{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Float32PtrErr failed")
	}

	_, err = PMapUint16Float32PtrErr(plusOneUint16Float32PtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Float32PtrErr failed")
	}
}

func TestPmapUint16Float64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3

	expectedList := []*float64{&vo2, &vo3}
	newList, _ := PMapUint16Float64PtrErr(plusOneUint16Float64PtrErr, []*uint16{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint16Float64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint16Float64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint16Float64PtrErr failed")
	}

	r, _ = PMapUint16Float64PtrErr(nil, []*uint16{})
	if len(r) > 0 {
		t.Errorf("PMapUint16Float64PtrErr failed")
	}
	_, err := PMapUint16Float64PtrErr(plusOneUint16Float64PtrErr, []*uint16{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint16Float64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint16Float64PtrErr(plusOneUint16Float64PtrErr, []*uint16{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16Float64PtrErr failed")
	}

	_, err = PMapUint16Float64PtrErr(plusOneUint16Float64PtrErr, []*uint16{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint16Float64PtrErr failed")
	}

	_, err = PMapUint16Float64PtrErr(plusOneUint16Float64PtrErr, []*uint16{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16Float64PtrErr failed")
	}
	_, err = PMapUint16Float64PtrErr(plusOneUint16Float64PtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16Float64PtrErr failed")
	}

	expectedList = []*float64{&vo2, &vo3}
	newList, _ = PMapUint16Float64PtrErr(plusOneUint16Float64PtrErr, []*uint16{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint16Float64Ptr failed")
	}

	_, err = PMapUint16Float64PtrErr(plusOneUint16Float64PtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Float64PtrErr failed")
	}
	_, err = PMapUint16Float64PtrErr(plusOneUint16Float64PtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Float64PtrErr failed")
	}

	_, err = PMapUint16Float64PtrErr(plusOneUint16Float64PtrErr, []*uint16{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Float64PtrErr failed")
	}

	_, err = PMapUint16Float64PtrErr(plusOneUint16Float64PtrErr, []*uint16{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Float64PtrErr failed")
	}

	_, err = PMapUint16Float64PtrErr(plusOneUint16Float64PtrErr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint16Float64PtrErr failed")
	}
}

func TestPmapUint8IntPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3

	expectedList := []*int{&vo2, &vo3}
	newList, _ := PMapUint8IntPtrErr(plusOneUint8IntPtrErr, []*uint8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint8IntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint8IntPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint8IntPtrErr failed")
	}

	r, _ = PMapUint8IntPtrErr(nil, []*uint8{})
	if len(r) > 0 {
		t.Errorf("PMapUint8IntPtrErr failed")
	}
	_, err := PMapUint8IntPtrErr(plusOneUint8IntPtrErr, []*uint8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint8IntPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint8IntPtrErr(plusOneUint8IntPtrErr, []*uint8{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8IntPtrErr failed")
	}

	_, err = PMapUint8IntPtrErr(plusOneUint8IntPtrErr, []*uint8{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint8IntPtrErr failed")
	}

	_, err = PMapUint8IntPtrErr(plusOneUint8IntPtrErr, []*uint8{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8IntPtrErr failed")
	}
	_, err = PMapUint8IntPtrErr(plusOneUint8IntPtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8IntPtrErr failed")
	}

	expectedList = []*int{&vo2, &vo3}
	newList, _ = PMapUint8IntPtrErr(plusOneUint8IntPtrErr, []*uint8{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint8IntPtr failed")
	}

	_, err = PMapUint8IntPtrErr(plusOneUint8IntPtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8IntPtrErr failed")
	}
	_, err = PMapUint8IntPtrErr(plusOneUint8IntPtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8IntPtrErr failed")
	}

	_, err = PMapUint8IntPtrErr(plusOneUint8IntPtrErr, []*uint8{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8IntPtrErr failed")
	}

	_, err = PMapUint8IntPtrErr(plusOneUint8IntPtrErr, []*uint8{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8IntPtrErr failed")
	}

	_, err = PMapUint8IntPtrErr(plusOneUint8IntPtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8IntPtrErr failed")
	}
}

func TestPmapUint8Int64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3

	expectedList := []*int64{&vo2, &vo3}
	newList, _ := PMapUint8Int64PtrErr(plusOneUint8Int64PtrErr, []*uint8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint8Int64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint8Int64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint8Int64PtrErr failed")
	}

	r, _ = PMapUint8Int64PtrErr(nil, []*uint8{})
	if len(r) > 0 {
		t.Errorf("PMapUint8Int64PtrErr failed")
	}
	_, err := PMapUint8Int64PtrErr(plusOneUint8Int64PtrErr, []*uint8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint8Int64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint8Int64PtrErr(plusOneUint8Int64PtrErr, []*uint8{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8Int64PtrErr failed")
	}

	_, err = PMapUint8Int64PtrErr(plusOneUint8Int64PtrErr, []*uint8{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint8Int64PtrErr failed")
	}

	_, err = PMapUint8Int64PtrErr(plusOneUint8Int64PtrErr, []*uint8{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8Int64PtrErr failed")
	}
	_, err = PMapUint8Int64PtrErr(plusOneUint8Int64PtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8Int64PtrErr failed")
	}

	expectedList = []*int64{&vo2, &vo3}
	newList, _ = PMapUint8Int64PtrErr(plusOneUint8Int64PtrErr, []*uint8{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint8Int64Ptr failed")
	}

	_, err = PMapUint8Int64PtrErr(plusOneUint8Int64PtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Int64PtrErr failed")
	}
	_, err = PMapUint8Int64PtrErr(plusOneUint8Int64PtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Int64PtrErr failed")
	}

	_, err = PMapUint8Int64PtrErr(plusOneUint8Int64PtrErr, []*uint8{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Int64PtrErr failed")
	}

	_, err = PMapUint8Int64PtrErr(plusOneUint8Int64PtrErr, []*uint8{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Int64PtrErr failed")
	}

	_, err = PMapUint8Int64PtrErr(plusOneUint8Int64PtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Int64PtrErr failed")
	}
}

func TestPmapUint8Int32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3

	expectedList := []*int32{&vo2, &vo3}
	newList, _ := PMapUint8Int32PtrErr(plusOneUint8Int32PtrErr, []*uint8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint8Int32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint8Int32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint8Int32PtrErr failed")
	}

	r, _ = PMapUint8Int32PtrErr(nil, []*uint8{})
	if len(r) > 0 {
		t.Errorf("PMapUint8Int32PtrErr failed")
	}
	_, err := PMapUint8Int32PtrErr(plusOneUint8Int32PtrErr, []*uint8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint8Int32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint8Int32PtrErr(plusOneUint8Int32PtrErr, []*uint8{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8Int32PtrErr failed")
	}

	_, err = PMapUint8Int32PtrErr(plusOneUint8Int32PtrErr, []*uint8{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint8Int32PtrErr failed")
	}

	_, err = PMapUint8Int32PtrErr(plusOneUint8Int32PtrErr, []*uint8{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8Int32PtrErr failed")
	}
	_, err = PMapUint8Int32PtrErr(plusOneUint8Int32PtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8Int32PtrErr failed")
	}

	expectedList = []*int32{&vo2, &vo3}
	newList, _ = PMapUint8Int32PtrErr(plusOneUint8Int32PtrErr, []*uint8{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint8Int32Ptr failed")
	}

	_, err = PMapUint8Int32PtrErr(plusOneUint8Int32PtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Int32PtrErr failed")
	}
	_, err = PMapUint8Int32PtrErr(plusOneUint8Int32PtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Int32PtrErr failed")
	}

	_, err = PMapUint8Int32PtrErr(plusOneUint8Int32PtrErr, []*uint8{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Int32PtrErr failed")
	}

	_, err = PMapUint8Int32PtrErr(plusOneUint8Int32PtrErr, []*uint8{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Int32PtrErr failed")
	}

	_, err = PMapUint8Int32PtrErr(plusOneUint8Int32PtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Int32PtrErr failed")
	}
}

func TestPmapUint8Int16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3

	expectedList := []*int16{&vo2, &vo3}
	newList, _ := PMapUint8Int16PtrErr(plusOneUint8Int16PtrErr, []*uint8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint8Int16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint8Int16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint8Int16PtrErr failed")
	}

	r, _ = PMapUint8Int16PtrErr(nil, []*uint8{})
	if len(r) > 0 {
		t.Errorf("PMapUint8Int16PtrErr failed")
	}
	_, err := PMapUint8Int16PtrErr(plusOneUint8Int16PtrErr, []*uint8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint8Int16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint8Int16PtrErr(plusOneUint8Int16PtrErr, []*uint8{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8Int16PtrErr failed")
	}

	_, err = PMapUint8Int16PtrErr(plusOneUint8Int16PtrErr, []*uint8{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint8Int16PtrErr failed")
	}

	_, err = PMapUint8Int16PtrErr(plusOneUint8Int16PtrErr, []*uint8{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8Int16PtrErr failed")
	}
	_, err = PMapUint8Int16PtrErr(plusOneUint8Int16PtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8Int16PtrErr failed")
	}

	expectedList = []*int16{&vo2, &vo3}
	newList, _ = PMapUint8Int16PtrErr(plusOneUint8Int16PtrErr, []*uint8{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint8Int16Ptr failed")
	}

	_, err = PMapUint8Int16PtrErr(plusOneUint8Int16PtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Int16PtrErr failed")
	}
	_, err = PMapUint8Int16PtrErr(plusOneUint8Int16PtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Int16PtrErr failed")
	}

	_, err = PMapUint8Int16PtrErr(plusOneUint8Int16PtrErr, []*uint8{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Int16PtrErr failed")
	}

	_, err = PMapUint8Int16PtrErr(plusOneUint8Int16PtrErr, []*uint8{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Int16PtrErr failed")
	}

	_, err = PMapUint8Int16PtrErr(plusOneUint8Int16PtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Int16PtrErr failed")
	}
}

func TestPmapUint8Int8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3

	expectedList := []*int8{&vo2, &vo3}
	newList, _ := PMapUint8Int8PtrErr(plusOneUint8Int8PtrErr, []*uint8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint8Int8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint8Int8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint8Int8PtrErr failed")
	}

	r, _ = PMapUint8Int8PtrErr(nil, []*uint8{})
	if len(r) > 0 {
		t.Errorf("PMapUint8Int8PtrErr failed")
	}
	_, err := PMapUint8Int8PtrErr(plusOneUint8Int8PtrErr, []*uint8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint8Int8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint8Int8PtrErr(plusOneUint8Int8PtrErr, []*uint8{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8Int8PtrErr failed")
	}

	_, err = PMapUint8Int8PtrErr(plusOneUint8Int8PtrErr, []*uint8{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint8Int8PtrErr failed")
	}

	_, err = PMapUint8Int8PtrErr(plusOneUint8Int8PtrErr, []*uint8{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8Int8PtrErr failed")
	}
	_, err = PMapUint8Int8PtrErr(plusOneUint8Int8PtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8Int8PtrErr failed")
	}

	expectedList = []*int8{&vo2, &vo3}
	newList, _ = PMapUint8Int8PtrErr(plusOneUint8Int8PtrErr, []*uint8{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint8Int8Ptr failed")
	}

	_, err = PMapUint8Int8PtrErr(plusOneUint8Int8PtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Int8PtrErr failed")
	}
	_, err = PMapUint8Int8PtrErr(plusOneUint8Int8PtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Int8PtrErr failed")
	}

	_, err = PMapUint8Int8PtrErr(plusOneUint8Int8PtrErr, []*uint8{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Int8PtrErr failed")
	}

	_, err = PMapUint8Int8PtrErr(plusOneUint8Int8PtrErr, []*uint8{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Int8PtrErr failed")
	}

	_, err = PMapUint8Int8PtrErr(plusOneUint8Int8PtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Int8PtrErr failed")
	}
}

func TestPmapUint8UintPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3

	expectedList := []*uint{&vo2, &vo3}
	newList, _ := PMapUint8UintPtrErr(plusOneUint8UintPtrErr, []*uint8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint8UintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint8UintPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint8UintPtrErr failed")
	}

	r, _ = PMapUint8UintPtrErr(nil, []*uint8{})
	if len(r) > 0 {
		t.Errorf("PMapUint8UintPtrErr failed")
	}
	_, err := PMapUint8UintPtrErr(plusOneUint8UintPtrErr, []*uint8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint8UintPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint8UintPtrErr(plusOneUint8UintPtrErr, []*uint8{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8UintPtrErr failed")
	}

	_, err = PMapUint8UintPtrErr(plusOneUint8UintPtrErr, []*uint8{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint8UintPtrErr failed")
	}

	_, err = PMapUint8UintPtrErr(plusOneUint8UintPtrErr, []*uint8{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8UintPtrErr failed")
	}
	_, err = PMapUint8UintPtrErr(plusOneUint8UintPtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8UintPtrErr failed")
	}

	expectedList = []*uint{&vo2, &vo3}
	newList, _ = PMapUint8UintPtrErr(plusOneUint8UintPtrErr, []*uint8{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint8UintPtr failed")
	}

	_, err = PMapUint8UintPtrErr(plusOneUint8UintPtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8UintPtrErr failed")
	}
	_, err = PMapUint8UintPtrErr(plusOneUint8UintPtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8UintPtrErr failed")
	}

	_, err = PMapUint8UintPtrErr(plusOneUint8UintPtrErr, []*uint8{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8UintPtrErr failed")
	}

	_, err = PMapUint8UintPtrErr(plusOneUint8UintPtrErr, []*uint8{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8UintPtrErr failed")
	}

	_, err = PMapUint8UintPtrErr(plusOneUint8UintPtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8UintPtrErr failed")
	}
}

func TestPmapUint8Uint64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3

	expectedList := []*uint64{&vo2, &vo3}
	newList, _ := PMapUint8Uint64PtrErr(plusOneUint8Uint64PtrErr, []*uint8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint8Uint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint8Uint64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint8Uint64PtrErr failed")
	}

	r, _ = PMapUint8Uint64PtrErr(nil, []*uint8{})
	if len(r) > 0 {
		t.Errorf("PMapUint8Uint64PtrErr failed")
	}
	_, err := PMapUint8Uint64PtrErr(plusOneUint8Uint64PtrErr, []*uint8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint8Uint64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint8Uint64PtrErr(plusOneUint8Uint64PtrErr, []*uint8{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8Uint64PtrErr failed")
	}

	_, err = PMapUint8Uint64PtrErr(plusOneUint8Uint64PtrErr, []*uint8{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint8Uint64PtrErr failed")
	}

	_, err = PMapUint8Uint64PtrErr(plusOneUint8Uint64PtrErr, []*uint8{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8Uint64PtrErr failed")
	}
	_, err = PMapUint8Uint64PtrErr(plusOneUint8Uint64PtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8Uint64PtrErr failed")
	}

	expectedList = []*uint64{&vo2, &vo3}
	newList, _ = PMapUint8Uint64PtrErr(plusOneUint8Uint64PtrErr, []*uint8{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint8Uint64Ptr failed")
	}

	_, err = PMapUint8Uint64PtrErr(plusOneUint8Uint64PtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Uint64PtrErr failed")
	}
	_, err = PMapUint8Uint64PtrErr(plusOneUint8Uint64PtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Uint64PtrErr failed")
	}

	_, err = PMapUint8Uint64PtrErr(plusOneUint8Uint64PtrErr, []*uint8{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Uint64PtrErr failed")
	}

	_, err = PMapUint8Uint64PtrErr(plusOneUint8Uint64PtrErr, []*uint8{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Uint64PtrErr failed")
	}

	_, err = PMapUint8Uint64PtrErr(plusOneUint8Uint64PtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Uint64PtrErr failed")
	}
}

func TestPmapUint8Uint32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3

	expectedList := []*uint32{&vo2, &vo3}
	newList, _ := PMapUint8Uint32PtrErr(plusOneUint8Uint32PtrErr, []*uint8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint8Uint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint8Uint32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint8Uint32PtrErr failed")
	}

	r, _ = PMapUint8Uint32PtrErr(nil, []*uint8{})
	if len(r) > 0 {
		t.Errorf("PMapUint8Uint32PtrErr failed")
	}
	_, err := PMapUint8Uint32PtrErr(plusOneUint8Uint32PtrErr, []*uint8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint8Uint32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint8Uint32PtrErr(plusOneUint8Uint32PtrErr, []*uint8{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8Uint32PtrErr failed")
	}

	_, err = PMapUint8Uint32PtrErr(plusOneUint8Uint32PtrErr, []*uint8{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint8Uint32PtrErr failed")
	}

	_, err = PMapUint8Uint32PtrErr(plusOneUint8Uint32PtrErr, []*uint8{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8Uint32PtrErr failed")
	}
	_, err = PMapUint8Uint32PtrErr(plusOneUint8Uint32PtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8Uint32PtrErr failed")
	}

	expectedList = []*uint32{&vo2, &vo3}
	newList, _ = PMapUint8Uint32PtrErr(plusOneUint8Uint32PtrErr, []*uint8{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint8Uint32Ptr failed")
	}

	_, err = PMapUint8Uint32PtrErr(plusOneUint8Uint32PtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Uint32PtrErr failed")
	}
	_, err = PMapUint8Uint32PtrErr(plusOneUint8Uint32PtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Uint32PtrErr failed")
	}

	_, err = PMapUint8Uint32PtrErr(plusOneUint8Uint32PtrErr, []*uint8{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Uint32PtrErr failed")
	}

	_, err = PMapUint8Uint32PtrErr(plusOneUint8Uint32PtrErr, []*uint8{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Uint32PtrErr failed")
	}

	_, err = PMapUint8Uint32PtrErr(plusOneUint8Uint32PtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Uint32PtrErr failed")
	}
}

func TestPmapUint8Uint16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3

	expectedList := []*uint16{&vo2, &vo3}
	newList, _ := PMapUint8Uint16PtrErr(plusOneUint8Uint16PtrErr, []*uint8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint8Uint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint8Uint16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint8Uint16PtrErr failed")
	}

	r, _ = PMapUint8Uint16PtrErr(nil, []*uint8{})
	if len(r) > 0 {
		t.Errorf("PMapUint8Uint16PtrErr failed")
	}
	_, err := PMapUint8Uint16PtrErr(plusOneUint8Uint16PtrErr, []*uint8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint8Uint16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint8Uint16PtrErr(plusOneUint8Uint16PtrErr, []*uint8{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8Uint16PtrErr failed")
	}

	_, err = PMapUint8Uint16PtrErr(plusOneUint8Uint16PtrErr, []*uint8{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint8Uint16PtrErr failed")
	}

	_, err = PMapUint8Uint16PtrErr(plusOneUint8Uint16PtrErr, []*uint8{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8Uint16PtrErr failed")
	}
	_, err = PMapUint8Uint16PtrErr(plusOneUint8Uint16PtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8Uint16PtrErr failed")
	}

	expectedList = []*uint16{&vo2, &vo3}
	newList, _ = PMapUint8Uint16PtrErr(plusOneUint8Uint16PtrErr, []*uint8{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint8Uint16Ptr failed")
	}

	_, err = PMapUint8Uint16PtrErr(plusOneUint8Uint16PtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Uint16PtrErr failed")
	}
	_, err = PMapUint8Uint16PtrErr(plusOneUint8Uint16PtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Uint16PtrErr failed")
	}

	_, err = PMapUint8Uint16PtrErr(plusOneUint8Uint16PtrErr, []*uint8{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Uint16PtrErr failed")
	}

	_, err = PMapUint8Uint16PtrErr(plusOneUint8Uint16PtrErr, []*uint8{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Uint16PtrErr failed")
	}

	_, err = PMapUint8Uint16PtrErr(plusOneUint8Uint16PtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Uint16PtrErr failed")
	}
}

func TestPmapUint8StrPtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 uint8 = 10
	var vi0 uint8

	expectedList := []*string{&vo10}
	newList, _ := PMapUint8StrPtrErr(someLogicUint8StrPtrErr, []*uint8{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapUint8StrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint8StrPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint8StrPtrErr failed")
	}

	r, _ = PMapUint8StrPtrErr(nil, []*uint8{})
	if len(r) > 0 {
		t.Errorf("PMapUint8StrPtrErr failed")
	}
	_, err := PMapUint8StrPtrErr(someLogicUint8StrPtrErr, []*uint8{&vi10, &vi0})
	if err == nil {
		t.Errorf("PMapUint8StrPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint8StrPtrErr(someLogicUint8StrPtrErr, []*uint8{&vi10, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8StrPtrErr failed")
	}

	_, err = PMapUint8StrPtrErr(someLogicUint8StrPtrErr, []*uint8{&vi0, &vi0, &vi10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint8StrPtrErr failed")
	}

	_, err = PMapUint8StrPtrErr(someLogicUint8StrPtrErr, []*uint8{&vi0, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8StrPtrErr failed")
	}
	_, err = PMapUint8StrPtrErr(someLogicUint8StrPtrErr, []*uint8{&vi10, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8StrPtrErr failed")
	}

	expectedList = []*string{&vo10, &vo10}
	newList, _ = PMapUint8StrPtrErr(someLogicUint8StrPtrErr, []*uint8{&vi10, &vi10}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint8StrPtr failed")
	}

	_, err = PMapUint8StrPtrErr(someLogicUint8StrPtrErr, []*uint8{&vi10, &vi10, &vi0}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8StrPtrErr failed")
	}
	_, err = PMapUint8StrPtrErr(someLogicUint8StrPtrErr, []*uint8{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8StrPtrErr failed")
	}

	_, err = PMapUint8StrPtrErr(someLogicUint8StrPtrErr, []*uint8{&vi0, &vi0, &vi10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8StrPtrErr failed")
	}

	_, err = PMapUint8StrPtrErr(someLogicUint8StrPtrErr, []*uint8{&vi0, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8StrPtrErr failed")
	}

	_, err = PMapUint8StrPtrErr(someLogicUint8StrPtrErr, []*uint8{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8StrPtrErr failed")
	}
}

func TestPmapUint8BoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 uint8 = 10
	var v0 uint8
	var v3 uint8 = 3

	expectedList := []*bool{&vt, &vf}
	newList, _ := PMapUint8BoolPtrErr(someLogicUint8BoolPtrErr, []*uint8{&v10, &v0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint8BoolPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint8BoolPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint8BoolPtrErr failed")
	}

	r, _ = PMapUint8BoolPtrErr(nil, []*uint8{})
	if len(r) > 0 {
		t.Errorf("PMapUint8BoolPtrErr failed")
	}
	_, err := PMapUint8BoolPtrErr(someLogicUint8BoolPtrErr, []*uint8{&v10, &v3})
	if err == nil {
		t.Errorf("PMapUint8BoolPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint8BoolPtrErr(someLogicUint8BoolPtrErr, []*uint8{&v10, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8BoolPtrErr failed")
	}

	_, err = PMapUint8BoolPtrErr(someLogicUint8BoolPtrErr, []*uint8{&v3, &v3, &v10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint8BoolPtrErr failed")
	}

	_, err = PMapUint8BoolPtrErr(someLogicUint8BoolPtrErr, []*uint8{&v3, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8BoolErr failed")
	}
	_, err = PMapUint8BoolPtrErr(someLogicUint8BoolPtrErr, []*uint8{&v10, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8BoolPtrErr failed")
	}

	expectedList = []*bool{&vt, &vt}
	newList, _ = PMapUint8BoolPtrErr(someLogicUint8BoolPtrErr, []*uint8{&v10, &v10}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint8BoolPtr failed")
	}

	_, err = PMapUint8BoolPtrErr(someLogicUint8BoolPtrErr, []*uint8{&v10, &v10, &v3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8BoolErr failed")
	}
	_, err = PMapUint8BoolPtrErr(someLogicUint8BoolPtrErr, []*uint8{&v10, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8BoolPtrErr failed")
	}

	_, err = PMapUint8BoolPtrErr(someLogicUint8BoolPtrErr, []*uint8{&v3, &v3, &v10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8BoolPtrErr failed")
	}

	_, err = PMapUint8BoolPtrErr(someLogicUint8BoolPtrErr, []*uint8{&v3, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8BoolPtrErr failed")
	}

	_, err = PMapUint8BoolPtrErr(someLogicUint8BoolPtrErr, []*uint8{&v10, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8BoolPtrErr failed")
	}
}

func TestPmapUint8Float32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3

	expectedList := []*float32{&vo2, &vo3}
	newList, _ := PMapUint8Float32PtrErr(plusOneUint8Float32PtrErr, []*uint8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint8Float32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint8Float32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint8Float32PtrErr failed")
	}

	r, _ = PMapUint8Float32PtrErr(nil, []*uint8{})
	if len(r) > 0 {
		t.Errorf("PMapUint8Float32PtrErr failed")
	}
	_, err := PMapUint8Float32PtrErr(plusOneUint8Float32PtrErr, []*uint8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint8Float32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint8Float32PtrErr(plusOneUint8Float32PtrErr, []*uint8{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8Float32PtrErr failed")
	}

	_, err = PMapUint8Float32PtrErr(plusOneUint8Float32PtrErr, []*uint8{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint8Float32PtrErr failed")
	}

	_, err = PMapUint8Float32PtrErr(plusOneUint8Float32PtrErr, []*uint8{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8Float32PtrErr failed")
	}
	_, err = PMapUint8Float32PtrErr(plusOneUint8Float32PtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8Float32PtrErr failed")
	}

	expectedList = []*float32{&vo2, &vo3}
	newList, _ = PMapUint8Float32PtrErr(plusOneUint8Float32PtrErr, []*uint8{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint8Float32Ptr failed")
	}

	_, err = PMapUint8Float32PtrErr(plusOneUint8Float32PtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Float32PtrErr failed")
	}
	_, err = PMapUint8Float32PtrErr(plusOneUint8Float32PtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Float32PtrErr failed")
	}

	_, err = PMapUint8Float32PtrErr(plusOneUint8Float32PtrErr, []*uint8{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Float32PtrErr failed")
	}

	_, err = PMapUint8Float32PtrErr(plusOneUint8Float32PtrErr, []*uint8{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Float32PtrErr failed")
	}

	_, err = PMapUint8Float32PtrErr(plusOneUint8Float32PtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Float32PtrErr failed")
	}
}

func TestPmapUint8Float64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3

	expectedList := []*float64{&vo2, &vo3}
	newList, _ := PMapUint8Float64PtrErr(plusOneUint8Float64PtrErr, []*uint8{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint8Float64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapUint8Float64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint8Float64PtrErr failed")
	}

	r, _ = PMapUint8Float64PtrErr(nil, []*uint8{})
	if len(r) > 0 {
		t.Errorf("PMapUint8Float64PtrErr failed")
	}
	_, err := PMapUint8Float64PtrErr(plusOneUint8Float64PtrErr, []*uint8{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapUint8Float64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapUint8Float64PtrErr(plusOneUint8Float64PtrErr, []*uint8{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8Float64PtrErr failed")
	}

	_, err = PMapUint8Float64PtrErr(plusOneUint8Float64PtrErr, []*uint8{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapUint8Float64PtrErr failed")
	}

	_, err = PMapUint8Float64PtrErr(plusOneUint8Float64PtrErr, []*uint8{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8Float64PtrErr failed")
	}
	_, err = PMapUint8Float64PtrErr(plusOneUint8Float64PtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8Float64PtrErr failed")
	}

	expectedList = []*float64{&vo2, &vo3}
	newList, _ = PMapUint8Float64PtrErr(plusOneUint8Float64PtrErr, []*uint8{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapUint8Float64Ptr failed")
	}

	_, err = PMapUint8Float64PtrErr(plusOneUint8Float64PtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Float64PtrErr failed")
	}
	_, err = PMapUint8Float64PtrErr(plusOneUint8Float64PtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Float64PtrErr failed")
	}

	_, err = PMapUint8Float64PtrErr(plusOneUint8Float64PtrErr, []*uint8{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Float64PtrErr failed")
	}

	_, err = PMapUint8Float64PtrErr(plusOneUint8Float64PtrErr, []*uint8{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Float64PtrErr failed")
	}

	_, err = PMapUint8Float64PtrErr(plusOneUint8Float64PtrErr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapUint8Float64PtrErr failed")
	}
}

func TestPmapStrIntPtrErr(t *testing.T) {
	// Test : someLogic
	// Test : someLogic
	var vo10 int = 10

	var vi10 string = "10"
	var vi0 string = "0"

	expectedList := []*int{&vo10}
	newList, _ := PMapStrIntPtrErr(someLogicStrIntPtrErr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapStrIntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapStrIntPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapStrIntPtrErr failed")
	}

	r, _ = PMapStrIntPtrErr(nil, []*string{})
	if len(r) > 0 {
		t.Errorf("PMapStrIntPtrErr failed")
	}
	_, err := PMapStrIntPtrErr(someLogicStrIntPtrErr, []*string{&vi10, &vi0})
	if err == nil {
		t.Errorf("PMapStrIntPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapStrIntPtrErr(someLogicStrIntPtrErr, []*string{&vi0, &vi0, &vi10}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrIntPtrErr failed")
	}

	_, err = PMapStrIntPtrErr(someLogicStrIntPtrErr, []*string{&vi0, &vi0, &vi10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapStrIntPtrErr failed")
	}

	_, err = PMapStrIntPtrErr(someLogicStrIntPtrErr, []*string{&vi0, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrIntPtrErr failed")
	}
	_, err = PMapStrIntPtrErr(someLogicStrIntPtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrIntPtrErr failed")
	}

	expectedList = []*int{&vo10, &vo10}
	newList, _ = PMapStrIntPtrErr(someLogicStrIntPtrErr, []*string{&vi10, &vi10}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapStrIntPtr failed")
	}

	_, err = PMapStrIntPtrErr(someLogicStrIntPtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrIntPtrErr failed")
	}
	_, err = PMapStrIntPtrErr(someLogicStrIntPtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrIntPtrErr failed")
	}

	_, err = PMapStrIntPtrErr(someLogicStrIntPtrErr, []*string{&vi0, &vi0, &vi10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrIntPtrErr failed")
	}

	_, err = PMapStrIntPtrErr(someLogicStrIntPtrErr, []*string{&vi0, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrIntPtrErr failed")
	}

	_, err = PMapStrIntPtrErr(someLogicStrIntPtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrIntPtrErr failed")
	}
}

func TestPmapStrInt64PtrErr(t *testing.T) {
	// Test : someLogic
	// Test : someLogic
	var vo10 int64 = 10

	var vi10 string = "10"
	var vi0 string = "0"

	expectedList := []*int64{&vo10}
	newList, _ := PMapStrInt64PtrErr(someLogicStrInt64PtrErr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapStrInt64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapStrInt64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapStrInt64PtrErr failed")
	}

	r, _ = PMapStrInt64PtrErr(nil, []*string{})
	if len(r) > 0 {
		t.Errorf("PMapStrInt64PtrErr failed")
	}
	_, err := PMapStrInt64PtrErr(someLogicStrInt64PtrErr, []*string{&vi10, &vi0})
	if err == nil {
		t.Errorf("PMapStrInt64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapStrInt64PtrErr(someLogicStrInt64PtrErr, []*string{&vi0, &vi0, &vi10}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrInt64PtrErr failed")
	}

	_, err = PMapStrInt64PtrErr(someLogicStrInt64PtrErr, []*string{&vi0, &vi0, &vi10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapStrInt64PtrErr failed")
	}

	_, err = PMapStrInt64PtrErr(someLogicStrInt64PtrErr, []*string{&vi0, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrInt64PtrErr failed")
	}
	_, err = PMapStrInt64PtrErr(someLogicStrInt64PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrInt64PtrErr failed")
	}

	expectedList = []*int64{&vo10, &vo10}
	newList, _ = PMapStrInt64PtrErr(someLogicStrInt64PtrErr, []*string{&vi10, &vi10}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapStrInt64Ptr failed")
	}

	_, err = PMapStrInt64PtrErr(someLogicStrInt64PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrInt64PtrErr failed")
	}
	_, err = PMapStrInt64PtrErr(someLogicStrInt64PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrInt64PtrErr failed")
	}

	_, err = PMapStrInt64PtrErr(someLogicStrInt64PtrErr, []*string{&vi0, &vi0, &vi10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrInt64PtrErr failed")
	}

	_, err = PMapStrInt64PtrErr(someLogicStrInt64PtrErr, []*string{&vi0, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrInt64PtrErr failed")
	}

	_, err = PMapStrInt64PtrErr(someLogicStrInt64PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrInt64PtrErr failed")
	}
}

func TestPmapStrInt32PtrErr(t *testing.T) {
	// Test : someLogic
	// Test : someLogic
	var vo10 int32 = 10

	var vi10 string = "10"
	var vi0 string = "0"

	expectedList := []*int32{&vo10}
	newList, _ := PMapStrInt32PtrErr(someLogicStrInt32PtrErr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapStrInt32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapStrInt32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapStrInt32PtrErr failed")
	}

	r, _ = PMapStrInt32PtrErr(nil, []*string{})
	if len(r) > 0 {
		t.Errorf("PMapStrInt32PtrErr failed")
	}
	_, err := PMapStrInt32PtrErr(someLogicStrInt32PtrErr, []*string{&vi10, &vi0})
	if err == nil {
		t.Errorf("PMapStrInt32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapStrInt32PtrErr(someLogicStrInt32PtrErr, []*string{&vi0, &vi0, &vi10}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrInt32PtrErr failed")
	}

	_, err = PMapStrInt32PtrErr(someLogicStrInt32PtrErr, []*string{&vi0, &vi0, &vi10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapStrInt32PtrErr failed")
	}

	_, err = PMapStrInt32PtrErr(someLogicStrInt32PtrErr, []*string{&vi0, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrInt32PtrErr failed")
	}
	_, err = PMapStrInt32PtrErr(someLogicStrInt32PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrInt32PtrErr failed")
	}

	expectedList = []*int32{&vo10, &vo10}
	newList, _ = PMapStrInt32PtrErr(someLogicStrInt32PtrErr, []*string{&vi10, &vi10}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapStrInt32Ptr failed")
	}

	_, err = PMapStrInt32PtrErr(someLogicStrInt32PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrInt32PtrErr failed")
	}
	_, err = PMapStrInt32PtrErr(someLogicStrInt32PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrInt32PtrErr failed")
	}

	_, err = PMapStrInt32PtrErr(someLogicStrInt32PtrErr, []*string{&vi0, &vi0, &vi10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrInt32PtrErr failed")
	}

	_, err = PMapStrInt32PtrErr(someLogicStrInt32PtrErr, []*string{&vi0, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrInt32PtrErr failed")
	}

	_, err = PMapStrInt32PtrErr(someLogicStrInt32PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrInt32PtrErr failed")
	}
}

func TestPmapStrInt16PtrErr(t *testing.T) {
	// Test : someLogic
	// Test : someLogic
	var vo10 int16 = 10

	var vi10 string = "10"
	var vi0 string = "0"

	expectedList := []*int16{&vo10}
	newList, _ := PMapStrInt16PtrErr(someLogicStrInt16PtrErr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapStrInt16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapStrInt16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapStrInt16PtrErr failed")
	}

	r, _ = PMapStrInt16PtrErr(nil, []*string{})
	if len(r) > 0 {
		t.Errorf("PMapStrInt16PtrErr failed")
	}
	_, err := PMapStrInt16PtrErr(someLogicStrInt16PtrErr, []*string{&vi10, &vi0})
	if err == nil {
		t.Errorf("PMapStrInt16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapStrInt16PtrErr(someLogicStrInt16PtrErr, []*string{&vi0, &vi0, &vi10}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrInt16PtrErr failed")
	}

	_, err = PMapStrInt16PtrErr(someLogicStrInt16PtrErr, []*string{&vi0, &vi0, &vi10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapStrInt16PtrErr failed")
	}

	_, err = PMapStrInt16PtrErr(someLogicStrInt16PtrErr, []*string{&vi0, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrInt16PtrErr failed")
	}
	_, err = PMapStrInt16PtrErr(someLogicStrInt16PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrInt16PtrErr failed")
	}

	expectedList = []*int16{&vo10, &vo10}
	newList, _ = PMapStrInt16PtrErr(someLogicStrInt16PtrErr, []*string{&vi10, &vi10}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapStrInt16Ptr failed")
	}

	_, err = PMapStrInt16PtrErr(someLogicStrInt16PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrInt16PtrErr failed")
	}
	_, err = PMapStrInt16PtrErr(someLogicStrInt16PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrInt16PtrErr failed")
	}

	_, err = PMapStrInt16PtrErr(someLogicStrInt16PtrErr, []*string{&vi0, &vi0, &vi10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrInt16PtrErr failed")
	}

	_, err = PMapStrInt16PtrErr(someLogicStrInt16PtrErr, []*string{&vi0, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrInt16PtrErr failed")
	}

	_, err = PMapStrInt16PtrErr(someLogicStrInt16PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrInt16PtrErr failed")
	}
}

func TestPmapStrInt8PtrErr(t *testing.T) {
	// Test : someLogic
	// Test : someLogic
	var vo10 int8 = 10

	var vi10 string = "10"
	var vi0 string = "0"

	expectedList := []*int8{&vo10}
	newList, _ := PMapStrInt8PtrErr(someLogicStrInt8PtrErr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapStrInt8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapStrInt8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapStrInt8PtrErr failed")
	}

	r, _ = PMapStrInt8PtrErr(nil, []*string{})
	if len(r) > 0 {
		t.Errorf("PMapStrInt8PtrErr failed")
	}
	_, err := PMapStrInt8PtrErr(someLogicStrInt8PtrErr, []*string{&vi10, &vi0})
	if err == nil {
		t.Errorf("PMapStrInt8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapStrInt8PtrErr(someLogicStrInt8PtrErr, []*string{&vi0, &vi0, &vi10}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrInt8PtrErr failed")
	}

	_, err = PMapStrInt8PtrErr(someLogicStrInt8PtrErr, []*string{&vi0, &vi0, &vi10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapStrInt8PtrErr failed")
	}

	_, err = PMapStrInt8PtrErr(someLogicStrInt8PtrErr, []*string{&vi0, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrInt8PtrErr failed")
	}
	_, err = PMapStrInt8PtrErr(someLogicStrInt8PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrInt8PtrErr failed")
	}

	expectedList = []*int8{&vo10, &vo10}
	newList, _ = PMapStrInt8PtrErr(someLogicStrInt8PtrErr, []*string{&vi10, &vi10}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapStrInt8Ptr failed")
	}

	_, err = PMapStrInt8PtrErr(someLogicStrInt8PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrInt8PtrErr failed")
	}
	_, err = PMapStrInt8PtrErr(someLogicStrInt8PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrInt8PtrErr failed")
	}

	_, err = PMapStrInt8PtrErr(someLogicStrInt8PtrErr, []*string{&vi0, &vi0, &vi10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrInt8PtrErr failed")
	}

	_, err = PMapStrInt8PtrErr(someLogicStrInt8PtrErr, []*string{&vi0, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrInt8PtrErr failed")
	}

	_, err = PMapStrInt8PtrErr(someLogicStrInt8PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrInt8PtrErr failed")
	}
}

func TestPmapStrUintPtrErr(t *testing.T) {
	// Test : someLogic
	// Test : someLogic
	var vo10 uint = 10

	var vi10 string = "10"
	var vi0 string = "0"

	expectedList := []*uint{&vo10}
	newList, _ := PMapStrUintPtrErr(someLogicStrUintPtrErr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapStrUintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapStrUintPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapStrUintPtrErr failed")
	}

	r, _ = PMapStrUintPtrErr(nil, []*string{})
	if len(r) > 0 {
		t.Errorf("PMapStrUintPtrErr failed")
	}
	_, err := PMapStrUintPtrErr(someLogicStrUintPtrErr, []*string{&vi10, &vi0})
	if err == nil {
		t.Errorf("PMapStrUintPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapStrUintPtrErr(someLogicStrUintPtrErr, []*string{&vi0, &vi0, &vi10}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrUintPtrErr failed")
	}

	_, err = PMapStrUintPtrErr(someLogicStrUintPtrErr, []*string{&vi0, &vi0, &vi10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapStrUintPtrErr failed")
	}

	_, err = PMapStrUintPtrErr(someLogicStrUintPtrErr, []*string{&vi0, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrUintPtrErr failed")
	}
	_, err = PMapStrUintPtrErr(someLogicStrUintPtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrUintPtrErr failed")
	}

	expectedList = []*uint{&vo10, &vo10}
	newList, _ = PMapStrUintPtrErr(someLogicStrUintPtrErr, []*string{&vi10, &vi10}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapStrUintPtr failed")
	}

	_, err = PMapStrUintPtrErr(someLogicStrUintPtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrUintPtrErr failed")
	}
	_, err = PMapStrUintPtrErr(someLogicStrUintPtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrUintPtrErr failed")
	}

	_, err = PMapStrUintPtrErr(someLogicStrUintPtrErr, []*string{&vi0, &vi0, &vi10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrUintPtrErr failed")
	}

	_, err = PMapStrUintPtrErr(someLogicStrUintPtrErr, []*string{&vi0, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrUintPtrErr failed")
	}

	_, err = PMapStrUintPtrErr(someLogicStrUintPtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrUintPtrErr failed")
	}
}

func TestPmapStrUint64PtrErr(t *testing.T) {
	// Test : someLogic
	// Test : someLogic
	var vo10 uint64 = 10

	var vi10 string = "10"
	var vi0 string = "0"

	expectedList := []*uint64{&vo10}
	newList, _ := PMapStrUint64PtrErr(someLogicStrUint64PtrErr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapStrUint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapStrUint64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapStrUint64PtrErr failed")
	}

	r, _ = PMapStrUint64PtrErr(nil, []*string{})
	if len(r) > 0 {
		t.Errorf("PMapStrUint64PtrErr failed")
	}
	_, err := PMapStrUint64PtrErr(someLogicStrUint64PtrErr, []*string{&vi10, &vi0})
	if err == nil {
		t.Errorf("PMapStrUint64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapStrUint64PtrErr(someLogicStrUint64PtrErr, []*string{&vi0, &vi0, &vi10}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrUint64PtrErr failed")
	}

	_, err = PMapStrUint64PtrErr(someLogicStrUint64PtrErr, []*string{&vi0, &vi0, &vi10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapStrUint64PtrErr failed")
	}

	_, err = PMapStrUint64PtrErr(someLogicStrUint64PtrErr, []*string{&vi0, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrUint64PtrErr failed")
	}
	_, err = PMapStrUint64PtrErr(someLogicStrUint64PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrUint64PtrErr failed")
	}

	expectedList = []*uint64{&vo10, &vo10}
	newList, _ = PMapStrUint64PtrErr(someLogicStrUint64PtrErr, []*string{&vi10, &vi10}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapStrUint64Ptr failed")
	}

	_, err = PMapStrUint64PtrErr(someLogicStrUint64PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrUint64PtrErr failed")
	}
	_, err = PMapStrUint64PtrErr(someLogicStrUint64PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrUint64PtrErr failed")
	}

	_, err = PMapStrUint64PtrErr(someLogicStrUint64PtrErr, []*string{&vi0, &vi0, &vi10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrUint64PtrErr failed")
	}

	_, err = PMapStrUint64PtrErr(someLogicStrUint64PtrErr, []*string{&vi0, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrUint64PtrErr failed")
	}

	_, err = PMapStrUint64PtrErr(someLogicStrUint64PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrUint64PtrErr failed")
	}
}

func TestPmapStrUint32PtrErr(t *testing.T) {
	// Test : someLogic
	// Test : someLogic
	var vo10 uint32 = 10

	var vi10 string = "10"
	var vi0 string = "0"

	expectedList := []*uint32{&vo10}
	newList, _ := PMapStrUint32PtrErr(someLogicStrUint32PtrErr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapStrUint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapStrUint32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapStrUint32PtrErr failed")
	}

	r, _ = PMapStrUint32PtrErr(nil, []*string{})
	if len(r) > 0 {
		t.Errorf("PMapStrUint32PtrErr failed")
	}
	_, err := PMapStrUint32PtrErr(someLogicStrUint32PtrErr, []*string{&vi10, &vi0})
	if err == nil {
		t.Errorf("PMapStrUint32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapStrUint32PtrErr(someLogicStrUint32PtrErr, []*string{&vi0, &vi0, &vi10}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrUint32PtrErr failed")
	}

	_, err = PMapStrUint32PtrErr(someLogicStrUint32PtrErr, []*string{&vi0, &vi0, &vi10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapStrUint32PtrErr failed")
	}

	_, err = PMapStrUint32PtrErr(someLogicStrUint32PtrErr, []*string{&vi0, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrUint32PtrErr failed")
	}
	_, err = PMapStrUint32PtrErr(someLogicStrUint32PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrUint32PtrErr failed")
	}

	expectedList = []*uint32{&vo10, &vo10}
	newList, _ = PMapStrUint32PtrErr(someLogicStrUint32PtrErr, []*string{&vi10, &vi10}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapStrUint32Ptr failed")
	}

	_, err = PMapStrUint32PtrErr(someLogicStrUint32PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrUint32PtrErr failed")
	}
	_, err = PMapStrUint32PtrErr(someLogicStrUint32PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrUint32PtrErr failed")
	}

	_, err = PMapStrUint32PtrErr(someLogicStrUint32PtrErr, []*string{&vi0, &vi0, &vi10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrUint32PtrErr failed")
	}

	_, err = PMapStrUint32PtrErr(someLogicStrUint32PtrErr, []*string{&vi0, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrUint32PtrErr failed")
	}

	_, err = PMapStrUint32PtrErr(someLogicStrUint32PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrUint32PtrErr failed")
	}
}

func TestPmapStrUint16PtrErr(t *testing.T) {
	// Test : someLogic
	// Test : someLogic
	var vo10 uint16 = 10

	var vi10 string = "10"
	var vi0 string = "0"

	expectedList := []*uint16{&vo10}
	newList, _ := PMapStrUint16PtrErr(someLogicStrUint16PtrErr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapStrUint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapStrUint16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapStrUint16PtrErr failed")
	}

	r, _ = PMapStrUint16PtrErr(nil, []*string{})
	if len(r) > 0 {
		t.Errorf("PMapStrUint16PtrErr failed")
	}
	_, err := PMapStrUint16PtrErr(someLogicStrUint16PtrErr, []*string{&vi10, &vi0})
	if err == nil {
		t.Errorf("PMapStrUint16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapStrUint16PtrErr(someLogicStrUint16PtrErr, []*string{&vi0, &vi0, &vi10}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrUint16PtrErr failed")
	}

	_, err = PMapStrUint16PtrErr(someLogicStrUint16PtrErr, []*string{&vi0, &vi0, &vi10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapStrUint16PtrErr failed")
	}

	_, err = PMapStrUint16PtrErr(someLogicStrUint16PtrErr, []*string{&vi0, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrUint16PtrErr failed")
	}
	_, err = PMapStrUint16PtrErr(someLogicStrUint16PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrUint16PtrErr failed")
	}

	expectedList = []*uint16{&vo10, &vo10}
	newList, _ = PMapStrUint16PtrErr(someLogicStrUint16PtrErr, []*string{&vi10, &vi10}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapStrUint16Ptr failed")
	}

	_, err = PMapStrUint16PtrErr(someLogicStrUint16PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrUint16PtrErr failed")
	}
	_, err = PMapStrUint16PtrErr(someLogicStrUint16PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrUint16PtrErr failed")
	}

	_, err = PMapStrUint16PtrErr(someLogicStrUint16PtrErr, []*string{&vi0, &vi0, &vi10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrUint16PtrErr failed")
	}

	_, err = PMapStrUint16PtrErr(someLogicStrUint16PtrErr, []*string{&vi0, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrUint16PtrErr failed")
	}

	_, err = PMapStrUint16PtrErr(someLogicStrUint16PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrUint16PtrErr failed")
	}
}

func TestPmapStrUint8PtrErr(t *testing.T) {
	// Test : someLogic
	// Test : someLogic
	var vo10 uint8 = 10

	var vi10 string = "10"
	var vi0 string = "0"

	expectedList := []*uint8{&vo10}
	newList, _ := PMapStrUint8PtrErr(someLogicStrUint8PtrErr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapStrUint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapStrUint8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapStrUint8PtrErr failed")
	}

	r, _ = PMapStrUint8PtrErr(nil, []*string{})
	if len(r) > 0 {
		t.Errorf("PMapStrUint8PtrErr failed")
	}
	_, err := PMapStrUint8PtrErr(someLogicStrUint8PtrErr, []*string{&vi10, &vi0})
	if err == nil {
		t.Errorf("PMapStrUint8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapStrUint8PtrErr(someLogicStrUint8PtrErr, []*string{&vi0, &vi0, &vi10}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrUint8PtrErr failed")
	}

	_, err = PMapStrUint8PtrErr(someLogicStrUint8PtrErr, []*string{&vi0, &vi0, &vi10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapStrUint8PtrErr failed")
	}

	_, err = PMapStrUint8PtrErr(someLogicStrUint8PtrErr, []*string{&vi0, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrUint8PtrErr failed")
	}
	_, err = PMapStrUint8PtrErr(someLogicStrUint8PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrUint8PtrErr failed")
	}

	expectedList = []*uint8{&vo10, &vo10}
	newList, _ = PMapStrUint8PtrErr(someLogicStrUint8PtrErr, []*string{&vi10, &vi10}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapStrUint8Ptr failed")
	}

	_, err = PMapStrUint8PtrErr(someLogicStrUint8PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrUint8PtrErr failed")
	}
	_, err = PMapStrUint8PtrErr(someLogicStrUint8PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrUint8PtrErr failed")
	}

	_, err = PMapStrUint8PtrErr(someLogicStrUint8PtrErr, []*string{&vi0, &vi0, &vi10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrUint8PtrErr failed")
	}

	_, err = PMapStrUint8PtrErr(someLogicStrUint8PtrErr, []*string{&vi0, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrUint8PtrErr failed")
	}

	_, err = PMapStrUint8PtrErr(someLogicStrUint8PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrUint8PtrErr failed")
	}
}

func TestPmapStrBoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 string = "10"
	var v0 string = "0"
	var v3 string = "3"

	expectedList := []*bool{&vt, &vf}
	newList, _ := PMapStrBoolPtrErr(someLogicStrBoolPtrErr, []*string{&v10, &v0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapStrBoolPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapStrBoolPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapStrBoolPtrErr failed")
	}

	r, _ = PMapStrBoolPtrErr(nil, []*string{})
	if len(r) > 0 {
		t.Errorf("PMapStrBoolPtrErr failed")
	}
	_, err := PMapStrBoolPtrErr(someLogicStrBoolPtrErr, []*string{&v10, &v0, &v3})
	if err == nil {
		t.Errorf("PMapStrBoolPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapStrBoolPtrErr(someLogicStrBoolPtrErr, []*string{&v10, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrBoolPtrErr failed")
	}

	_, err = PMapStrBoolPtrErr(someLogicStrBoolPtrErr, []*string{&v3, &v3, &v10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapStrBoolPtrErr failed")
	}

	_, err = PMapStrBoolPtrErr(someLogicStrBoolPtrErr, []*string{&v3, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrBoolPtrErr failed")
	}
	_, err = PMapStrBoolPtrErr(someLogicStrBoolPtrErr, []*string{&v10, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrBoolPtrErr failed")
	}

	expectedList = []*bool{&vt, &vt}
	newList, _ = PMapStrBoolPtrErr(someLogicStrBoolPtrErr, []*string{&v10, &v10}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapStrBoolPtr failed")
	}

	_, err = PMapStrBoolPtrErr(someLogicStrBoolPtrErr, []*string{&v10, &v10, &v3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrBoolPtrErr failed")
	}
	_, err = PMapStrBoolPtrErr(someLogicStrBoolPtrErr, []*string{&v10, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrBoolPtrErr failed")
	}

	_, err = PMapStrBoolPtrErr(someLogicStrBoolPtrErr, []*string{&v3, &v3, &v10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrBoolPtrErr failed")
	}

	_, err = PMapStrBoolPtrErr(someLogicStrBoolPtrErr, []*string{&v3, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrBoolPtrErr failed")
	}

	_, err = PMapStrBoolPtrErr(someLogicStrBoolPtrErr, []*string{&v10, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrBoolErr failed")
	}
}

func TestPmapStrFloat32PtrErr(t *testing.T) {
	// Test : someLogic
	// Test : someLogic
	var vo10 float32 = 10

	var vi10 string = "10"
	var vi0 string = "0"

	expectedList := []*float32{&vo10}
	newList, _ := PMapStrFloat32PtrErr(someLogicStrFloat32PtrErr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapStrFloat32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapStrFloat32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapStrFloat32PtrErr failed")
	}

	r, _ = PMapStrFloat32PtrErr(nil, []*string{})
	if len(r) > 0 {
		t.Errorf("PMapStrFloat32PtrErr failed")
	}
	_, err := PMapStrFloat32PtrErr(someLogicStrFloat32PtrErr, []*string{&vi10, &vi0})
	if err == nil {
		t.Errorf("PMapStrFloat32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapStrFloat32PtrErr(someLogicStrFloat32PtrErr, []*string{&vi0, &vi0, &vi10}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrFloat32PtrErr failed")
	}

	_, err = PMapStrFloat32PtrErr(someLogicStrFloat32PtrErr, []*string{&vi0, &vi0, &vi10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapStrFloat32PtrErr failed")
	}

	_, err = PMapStrFloat32PtrErr(someLogicStrFloat32PtrErr, []*string{&vi0, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrFloat32PtrErr failed")
	}
	_, err = PMapStrFloat32PtrErr(someLogicStrFloat32PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrFloat32PtrErr failed")
	}

	expectedList = []*float32{&vo10, &vo10}
	newList, _ = PMapStrFloat32PtrErr(someLogicStrFloat32PtrErr, []*string{&vi10, &vi10}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapStrFloat32Ptr failed")
	}

	_, err = PMapStrFloat32PtrErr(someLogicStrFloat32PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrFloat32PtrErr failed")
	}
	_, err = PMapStrFloat32PtrErr(someLogicStrFloat32PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrFloat32PtrErr failed")
	}

	_, err = PMapStrFloat32PtrErr(someLogicStrFloat32PtrErr, []*string{&vi0, &vi0, &vi10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrFloat32PtrErr failed")
	}

	_, err = PMapStrFloat32PtrErr(someLogicStrFloat32PtrErr, []*string{&vi0, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrFloat32PtrErr failed")
	}

	_, err = PMapStrFloat32PtrErr(someLogicStrFloat32PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrFloat32PtrErr failed")
	}
}

func TestPmapStrFloat64PtrErr(t *testing.T) {
	// Test : someLogic
	// Test : someLogic
	var vo10 float64 = 10

	var vi10 string = "10"
	var vi0 string = "0"

	expectedList := []*float64{&vo10}
	newList, _ := PMapStrFloat64PtrErr(someLogicStrFloat64PtrErr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapStrFloat64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapStrFloat64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapStrFloat64PtrErr failed")
	}

	r, _ = PMapStrFloat64PtrErr(nil, []*string{})
	if len(r) > 0 {
		t.Errorf("PMapStrFloat64PtrErr failed")
	}
	_, err := PMapStrFloat64PtrErr(someLogicStrFloat64PtrErr, []*string{&vi10, &vi0})
	if err == nil {
		t.Errorf("PMapStrFloat64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapStrFloat64PtrErr(someLogicStrFloat64PtrErr, []*string{&vi0, &vi0, &vi10}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrFloat64PtrErr failed")
	}

	_, err = PMapStrFloat64PtrErr(someLogicStrFloat64PtrErr, []*string{&vi0, &vi0, &vi10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapStrFloat64PtrErr failed")
	}

	_, err = PMapStrFloat64PtrErr(someLogicStrFloat64PtrErr, []*string{&vi0, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrFloat64PtrErr failed")
	}
	_, err = PMapStrFloat64PtrErr(someLogicStrFloat64PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapStrFloat64PtrErr failed")
	}

	expectedList = []*float64{&vo10, &vo10}
	newList, _ = PMapStrFloat64PtrErr(someLogicStrFloat64PtrErr, []*string{&vi10, &vi10}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapStrFloat64Ptr failed")
	}

	_, err = PMapStrFloat64PtrErr(someLogicStrFloat64PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrFloat64PtrErr failed")
	}
	_, err = PMapStrFloat64PtrErr(someLogicStrFloat64PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrFloat64PtrErr failed")
	}

	_, err = PMapStrFloat64PtrErr(someLogicStrFloat64PtrErr, []*string{&vi0, &vi0, &vi10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrFloat64PtrErr failed")
	}

	_, err = PMapStrFloat64PtrErr(someLogicStrFloat64PtrErr, []*string{&vi0, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrFloat64PtrErr failed")
	}

	_, err = PMapStrFloat64PtrErr(someLogicStrFloat64PtrErr, []*string{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapStrFloat64PtrErr failed")
	}
}

func TestPmapBoolIntPtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 int = 10
	var v0 int = 0

	expectedList := []*int{&v10, &v0}
	newList, _ := PMapBoolIntPtrErr(someLogicBoolIntPtrErr, []*bool{&vt, &vt})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolIntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapBoolIntPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapBoolIntPtrErr failed")
	}

	r, _ = PMapBoolIntPtrErr(nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("PMapBoolIntPtrErr failed")
	}
	_, err := PMapBoolIntPtrErr(someLogicBoolIntPtrErr, []*bool{&vt, &vf})
	if err == nil {
		t.Errorf("PMapBoolIntPtrErr failed")
	}
	reflect.TypeOf("Ram") // Leaving it here to make use of import reflect

	_, err = PMapBoolIntPtrErr(someLogicBoolIntPtrErr, []*bool{&vt, &vt, &vf, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolIntPtrErr failed")
	}

	_, err = PMapBoolIntPtrErr(someLogicBoolIntPtrErr, []*bool{&vf, &vf, &vt}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapBoolIntPtrErr failed")
	}

	_, err = PMapBoolIntPtrErr(someLogicBoolIntPtrErr, []*bool{&vf, &vt, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolIntPtrErr failed")
	}
	_, err = PMapBoolIntPtrErr(someLogicBoolIntPtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolIntPtrErr failed")
	}

	expectedList = []*int{&v10, &v10}
	newList, _ = PMapBoolIntPtrErr(someLogicBoolIntPtrErr, []*bool{&vt, &vt}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapBoolIntPtr failed")
	}

	_, err = PMapBoolIntPtrErr(someLogicBoolIntPtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolIntPtrErr failed")
	}
	_, err = PMapBoolIntPtrErr(someLogicBoolIntPtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolIntPtrErr failed")
	}

	_, err = PMapBoolIntPtrErr(someLogicBoolIntPtrErr, []*bool{&vf, &vf, &vt}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolIntPtrErr failed")
	}

	_, err = PMapBoolIntPtrErr(someLogicBoolIntPtrErr, []*bool{&vf, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolIntPtrErr failed")
	}

	_, err = PMapBoolIntPtrErr(someLogicBoolIntPtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolIntPtrErr failed")
	}
}

func TestPmapBoolInt64PtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 int64 = 10
	var v0 int64 = 0

	expectedList := []*int64{&v10, &v0}
	newList, _ := PMapBoolInt64PtrErr(someLogicBoolInt64PtrErr, []*bool{&vt, &vt})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolInt64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapBoolInt64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapBoolInt64PtrErr failed")
	}

	r, _ = PMapBoolInt64PtrErr(nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("PMapBoolInt64PtrErr failed")
	}
	_, err := PMapBoolInt64PtrErr(someLogicBoolInt64PtrErr, []*bool{&vt, &vf})
	if err == nil {
		t.Errorf("PMapBoolInt64PtrErr failed")
	}
	reflect.TypeOf("Ram") // Leaving it here to make use of import reflect

	_, err = PMapBoolInt64PtrErr(someLogicBoolInt64PtrErr, []*bool{&vt, &vt, &vf, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolInt64PtrErr failed")
	}

	_, err = PMapBoolInt64PtrErr(someLogicBoolInt64PtrErr, []*bool{&vf, &vf, &vt}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapBoolInt64PtrErr failed")
	}

	_, err = PMapBoolInt64PtrErr(someLogicBoolInt64PtrErr, []*bool{&vf, &vt, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolInt64PtrErr failed")
	}
	_, err = PMapBoolInt64PtrErr(someLogicBoolInt64PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolInt64PtrErr failed")
	}

	expectedList = []*int64{&v10, &v10}
	newList, _ = PMapBoolInt64PtrErr(someLogicBoolInt64PtrErr, []*bool{&vt, &vt}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapBoolInt64Ptr failed")
	}

	_, err = PMapBoolInt64PtrErr(someLogicBoolInt64PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolInt64PtrErr failed")
	}
	_, err = PMapBoolInt64PtrErr(someLogicBoolInt64PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolInt64PtrErr failed")
	}

	_, err = PMapBoolInt64PtrErr(someLogicBoolInt64PtrErr, []*bool{&vf, &vf, &vt}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolInt64PtrErr failed")
	}

	_, err = PMapBoolInt64PtrErr(someLogicBoolInt64PtrErr, []*bool{&vf, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolInt64PtrErr failed")
	}

	_, err = PMapBoolInt64PtrErr(someLogicBoolInt64PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolInt64PtrErr failed")
	}
}

func TestPmapBoolInt32PtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 int32 = 10
	var v0 int32 = 0

	expectedList := []*int32{&v10, &v0}
	newList, _ := PMapBoolInt32PtrErr(someLogicBoolInt32PtrErr, []*bool{&vt, &vt})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolInt32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapBoolInt32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapBoolInt32PtrErr failed")
	}

	r, _ = PMapBoolInt32PtrErr(nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("PMapBoolInt32PtrErr failed")
	}
	_, err := PMapBoolInt32PtrErr(someLogicBoolInt32PtrErr, []*bool{&vt, &vf})
	if err == nil {
		t.Errorf("PMapBoolInt32PtrErr failed")
	}
	reflect.TypeOf("Ram") // Leaving it here to make use of import reflect

	_, err = PMapBoolInt32PtrErr(someLogicBoolInt32PtrErr, []*bool{&vt, &vt, &vf, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolInt32PtrErr failed")
	}

	_, err = PMapBoolInt32PtrErr(someLogicBoolInt32PtrErr, []*bool{&vf, &vf, &vt}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapBoolInt32PtrErr failed")
	}

	_, err = PMapBoolInt32PtrErr(someLogicBoolInt32PtrErr, []*bool{&vf, &vt, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolInt32PtrErr failed")
	}
	_, err = PMapBoolInt32PtrErr(someLogicBoolInt32PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolInt32PtrErr failed")
	}

	expectedList = []*int32{&v10, &v10}
	newList, _ = PMapBoolInt32PtrErr(someLogicBoolInt32PtrErr, []*bool{&vt, &vt}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapBoolInt32Ptr failed")
	}

	_, err = PMapBoolInt32PtrErr(someLogicBoolInt32PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolInt32PtrErr failed")
	}
	_, err = PMapBoolInt32PtrErr(someLogicBoolInt32PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolInt32PtrErr failed")
	}

	_, err = PMapBoolInt32PtrErr(someLogicBoolInt32PtrErr, []*bool{&vf, &vf, &vt}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolInt32PtrErr failed")
	}

	_, err = PMapBoolInt32PtrErr(someLogicBoolInt32PtrErr, []*bool{&vf, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolInt32PtrErr failed")
	}

	_, err = PMapBoolInt32PtrErr(someLogicBoolInt32PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolInt32PtrErr failed")
	}
}

func TestPmapBoolInt16PtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 int16 = 10
	var v0 int16 = 0

	expectedList := []*int16{&v10, &v0}
	newList, _ := PMapBoolInt16PtrErr(someLogicBoolInt16PtrErr, []*bool{&vt, &vt})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolInt16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapBoolInt16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapBoolInt16PtrErr failed")
	}

	r, _ = PMapBoolInt16PtrErr(nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("PMapBoolInt16PtrErr failed")
	}
	_, err := PMapBoolInt16PtrErr(someLogicBoolInt16PtrErr, []*bool{&vt, &vf})
	if err == nil {
		t.Errorf("PMapBoolInt16PtrErr failed")
	}
	reflect.TypeOf("Ram") // Leaving it here to make use of import reflect

	_, err = PMapBoolInt16PtrErr(someLogicBoolInt16PtrErr, []*bool{&vt, &vt, &vf, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolInt16PtrErr failed")
	}

	_, err = PMapBoolInt16PtrErr(someLogicBoolInt16PtrErr, []*bool{&vf, &vf, &vt}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapBoolInt16PtrErr failed")
	}

	_, err = PMapBoolInt16PtrErr(someLogicBoolInt16PtrErr, []*bool{&vf, &vt, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolInt16PtrErr failed")
	}
	_, err = PMapBoolInt16PtrErr(someLogicBoolInt16PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolInt16PtrErr failed")
	}

	expectedList = []*int16{&v10, &v10}
	newList, _ = PMapBoolInt16PtrErr(someLogicBoolInt16PtrErr, []*bool{&vt, &vt}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapBoolInt16Ptr failed")
	}

	_, err = PMapBoolInt16PtrErr(someLogicBoolInt16PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolInt16PtrErr failed")
	}
	_, err = PMapBoolInt16PtrErr(someLogicBoolInt16PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolInt16PtrErr failed")
	}

	_, err = PMapBoolInt16PtrErr(someLogicBoolInt16PtrErr, []*bool{&vf, &vf, &vt}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolInt16PtrErr failed")
	}

	_, err = PMapBoolInt16PtrErr(someLogicBoolInt16PtrErr, []*bool{&vf, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolInt16PtrErr failed")
	}

	_, err = PMapBoolInt16PtrErr(someLogicBoolInt16PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolInt16PtrErr failed")
	}
}

func TestPmapBoolInt8PtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 int8 = 10
	var v0 int8 = 0

	expectedList := []*int8{&v10, &v0}
	newList, _ := PMapBoolInt8PtrErr(someLogicBoolInt8PtrErr, []*bool{&vt, &vt})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolInt8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapBoolInt8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapBoolInt8PtrErr failed")
	}

	r, _ = PMapBoolInt8PtrErr(nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("PMapBoolInt8PtrErr failed")
	}
	_, err := PMapBoolInt8PtrErr(someLogicBoolInt8PtrErr, []*bool{&vt, &vf})
	if err == nil {
		t.Errorf("PMapBoolInt8PtrErr failed")
	}
	reflect.TypeOf("Ram") // Leaving it here to make use of import reflect

	_, err = PMapBoolInt8PtrErr(someLogicBoolInt8PtrErr, []*bool{&vt, &vt, &vf, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolInt8PtrErr failed")
	}

	_, err = PMapBoolInt8PtrErr(someLogicBoolInt8PtrErr, []*bool{&vf, &vf, &vt}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapBoolInt8PtrErr failed")
	}

	_, err = PMapBoolInt8PtrErr(someLogicBoolInt8PtrErr, []*bool{&vf, &vt, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolInt8PtrErr failed")
	}
	_, err = PMapBoolInt8PtrErr(someLogicBoolInt8PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolInt8PtrErr failed")
	}

	expectedList = []*int8{&v10, &v10}
	newList, _ = PMapBoolInt8PtrErr(someLogicBoolInt8PtrErr, []*bool{&vt, &vt}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapBoolInt8Ptr failed")
	}

	_, err = PMapBoolInt8PtrErr(someLogicBoolInt8PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolInt8PtrErr failed")
	}
	_, err = PMapBoolInt8PtrErr(someLogicBoolInt8PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolInt8PtrErr failed")
	}

	_, err = PMapBoolInt8PtrErr(someLogicBoolInt8PtrErr, []*bool{&vf, &vf, &vt}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolInt8PtrErr failed")
	}

	_, err = PMapBoolInt8PtrErr(someLogicBoolInt8PtrErr, []*bool{&vf, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolInt8PtrErr failed")
	}

	_, err = PMapBoolInt8PtrErr(someLogicBoolInt8PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolInt8PtrErr failed")
	}
}

func TestPmapBoolUintPtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 uint = 10
	var v0 uint = 0

	expectedList := []*uint{&v10, &v0}
	newList, _ := PMapBoolUintPtrErr(someLogicBoolUintPtrErr, []*bool{&vt, &vt})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolUintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapBoolUintPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapBoolUintPtrErr failed")
	}

	r, _ = PMapBoolUintPtrErr(nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("PMapBoolUintPtrErr failed")
	}
	_, err := PMapBoolUintPtrErr(someLogicBoolUintPtrErr, []*bool{&vt, &vf})
	if err == nil {
		t.Errorf("PMapBoolUintPtrErr failed")
	}
	reflect.TypeOf("Ram") // Leaving it here to make use of import reflect

	_, err = PMapBoolUintPtrErr(someLogicBoolUintPtrErr, []*bool{&vt, &vt, &vf, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolUintPtrErr failed")
	}

	_, err = PMapBoolUintPtrErr(someLogicBoolUintPtrErr, []*bool{&vf, &vf, &vt}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapBoolUintPtrErr failed")
	}

	_, err = PMapBoolUintPtrErr(someLogicBoolUintPtrErr, []*bool{&vf, &vt, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolUintPtrErr failed")
	}
	_, err = PMapBoolUintPtrErr(someLogicBoolUintPtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolUintPtrErr failed")
	}

	expectedList = []*uint{&v10, &v10}
	newList, _ = PMapBoolUintPtrErr(someLogicBoolUintPtrErr, []*bool{&vt, &vt}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapBoolUintPtr failed")
	}

	_, err = PMapBoolUintPtrErr(someLogicBoolUintPtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolUintPtrErr failed")
	}
	_, err = PMapBoolUintPtrErr(someLogicBoolUintPtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolUintPtrErr failed")
	}

	_, err = PMapBoolUintPtrErr(someLogicBoolUintPtrErr, []*bool{&vf, &vf, &vt}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolUintPtrErr failed")
	}

	_, err = PMapBoolUintPtrErr(someLogicBoolUintPtrErr, []*bool{&vf, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolUintPtrErr failed")
	}

	_, err = PMapBoolUintPtrErr(someLogicBoolUintPtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolUintPtrErr failed")
	}
}

func TestPmapBoolUint64PtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 uint64 = 10
	var v0 uint64 = 0

	expectedList := []*uint64{&v10, &v0}
	newList, _ := PMapBoolUint64PtrErr(someLogicBoolUint64PtrErr, []*bool{&vt, &vt})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolUint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapBoolUint64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapBoolUint64PtrErr failed")
	}

	r, _ = PMapBoolUint64PtrErr(nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("PMapBoolUint64PtrErr failed")
	}
	_, err := PMapBoolUint64PtrErr(someLogicBoolUint64PtrErr, []*bool{&vt, &vf})
	if err == nil {
		t.Errorf("PMapBoolUint64PtrErr failed")
	}
	reflect.TypeOf("Ram") // Leaving it here to make use of import reflect

	_, err = PMapBoolUint64PtrErr(someLogicBoolUint64PtrErr, []*bool{&vt, &vt, &vf, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolUint64PtrErr failed")
	}

	_, err = PMapBoolUint64PtrErr(someLogicBoolUint64PtrErr, []*bool{&vf, &vf, &vt}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapBoolUint64PtrErr failed")
	}

	_, err = PMapBoolUint64PtrErr(someLogicBoolUint64PtrErr, []*bool{&vf, &vt, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolUint64PtrErr failed")
	}
	_, err = PMapBoolUint64PtrErr(someLogicBoolUint64PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolUint64PtrErr failed")
	}

	expectedList = []*uint64{&v10, &v10}
	newList, _ = PMapBoolUint64PtrErr(someLogicBoolUint64PtrErr, []*bool{&vt, &vt}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapBoolUint64Ptr failed")
	}

	_, err = PMapBoolUint64PtrErr(someLogicBoolUint64PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolUint64PtrErr failed")
	}
	_, err = PMapBoolUint64PtrErr(someLogicBoolUint64PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolUint64PtrErr failed")
	}

	_, err = PMapBoolUint64PtrErr(someLogicBoolUint64PtrErr, []*bool{&vf, &vf, &vt}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolUint64PtrErr failed")
	}

	_, err = PMapBoolUint64PtrErr(someLogicBoolUint64PtrErr, []*bool{&vf, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolUint64PtrErr failed")
	}

	_, err = PMapBoolUint64PtrErr(someLogicBoolUint64PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolUint64PtrErr failed")
	}
}

func TestPmapBoolUint32PtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 uint32 = 10
	var v0 uint32 = 0

	expectedList := []*uint32{&v10, &v0}
	newList, _ := PMapBoolUint32PtrErr(someLogicBoolUint32PtrErr, []*bool{&vt, &vt})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolUint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapBoolUint32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapBoolUint32PtrErr failed")
	}

	r, _ = PMapBoolUint32PtrErr(nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("PMapBoolUint32PtrErr failed")
	}
	_, err := PMapBoolUint32PtrErr(someLogicBoolUint32PtrErr, []*bool{&vt, &vf})
	if err == nil {
		t.Errorf("PMapBoolUint32PtrErr failed")
	}
	reflect.TypeOf("Ram") // Leaving it here to make use of import reflect

	_, err = PMapBoolUint32PtrErr(someLogicBoolUint32PtrErr, []*bool{&vt, &vt, &vf, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolUint32PtrErr failed")
	}

	_, err = PMapBoolUint32PtrErr(someLogicBoolUint32PtrErr, []*bool{&vf, &vf, &vt}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapBoolUint32PtrErr failed")
	}

	_, err = PMapBoolUint32PtrErr(someLogicBoolUint32PtrErr, []*bool{&vf, &vt, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolUint32PtrErr failed")
	}
	_, err = PMapBoolUint32PtrErr(someLogicBoolUint32PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolUint32PtrErr failed")
	}

	expectedList = []*uint32{&v10, &v10}
	newList, _ = PMapBoolUint32PtrErr(someLogicBoolUint32PtrErr, []*bool{&vt, &vt}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapBoolUint32Ptr failed")
	}

	_, err = PMapBoolUint32PtrErr(someLogicBoolUint32PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolUint32PtrErr failed")
	}
	_, err = PMapBoolUint32PtrErr(someLogicBoolUint32PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolUint32PtrErr failed")
	}

	_, err = PMapBoolUint32PtrErr(someLogicBoolUint32PtrErr, []*bool{&vf, &vf, &vt}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolUint32PtrErr failed")
	}

	_, err = PMapBoolUint32PtrErr(someLogicBoolUint32PtrErr, []*bool{&vf, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolUint32PtrErr failed")
	}

	_, err = PMapBoolUint32PtrErr(someLogicBoolUint32PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolUint32PtrErr failed")
	}
}

func TestPmapBoolUint16PtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 uint16 = 10
	var v0 uint16 = 0

	expectedList := []*uint16{&v10, &v0}
	newList, _ := PMapBoolUint16PtrErr(someLogicBoolUint16PtrErr, []*bool{&vt, &vt})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolUint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapBoolUint16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapBoolUint16PtrErr failed")
	}

	r, _ = PMapBoolUint16PtrErr(nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("PMapBoolUint16PtrErr failed")
	}
	_, err := PMapBoolUint16PtrErr(someLogicBoolUint16PtrErr, []*bool{&vt, &vf})
	if err == nil {
		t.Errorf("PMapBoolUint16PtrErr failed")
	}
	reflect.TypeOf("Ram") // Leaving it here to make use of import reflect

	_, err = PMapBoolUint16PtrErr(someLogicBoolUint16PtrErr, []*bool{&vt, &vt, &vf, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolUint16PtrErr failed")
	}

	_, err = PMapBoolUint16PtrErr(someLogicBoolUint16PtrErr, []*bool{&vf, &vf, &vt}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapBoolUint16PtrErr failed")
	}

	_, err = PMapBoolUint16PtrErr(someLogicBoolUint16PtrErr, []*bool{&vf, &vt, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolUint16PtrErr failed")
	}
	_, err = PMapBoolUint16PtrErr(someLogicBoolUint16PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolUint16PtrErr failed")
	}

	expectedList = []*uint16{&v10, &v10}
	newList, _ = PMapBoolUint16PtrErr(someLogicBoolUint16PtrErr, []*bool{&vt, &vt}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapBoolUint16Ptr failed")
	}

	_, err = PMapBoolUint16PtrErr(someLogicBoolUint16PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolUint16PtrErr failed")
	}
	_, err = PMapBoolUint16PtrErr(someLogicBoolUint16PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolUint16PtrErr failed")
	}

	_, err = PMapBoolUint16PtrErr(someLogicBoolUint16PtrErr, []*bool{&vf, &vf, &vt}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolUint16PtrErr failed")
	}

	_, err = PMapBoolUint16PtrErr(someLogicBoolUint16PtrErr, []*bool{&vf, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolUint16PtrErr failed")
	}

	_, err = PMapBoolUint16PtrErr(someLogicBoolUint16PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolUint16PtrErr failed")
	}
}

func TestPmapBoolUint8PtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 uint8 = 10
	var v0 uint8 = 0

	expectedList := []*uint8{&v10, &v0}
	newList, _ := PMapBoolUint8PtrErr(someLogicBoolUint8PtrErr, []*bool{&vt, &vt})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolUint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapBoolUint8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapBoolUint8PtrErr failed")
	}

	r, _ = PMapBoolUint8PtrErr(nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("PMapBoolUint8PtrErr failed")
	}
	_, err := PMapBoolUint8PtrErr(someLogicBoolUint8PtrErr, []*bool{&vt, &vf})
	if err == nil {
		t.Errorf("PMapBoolUint8PtrErr failed")
	}
	reflect.TypeOf("Ram") // Leaving it here to make use of import reflect

	_, err = PMapBoolUint8PtrErr(someLogicBoolUint8PtrErr, []*bool{&vt, &vt, &vf, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolUint8PtrErr failed")
	}

	_, err = PMapBoolUint8PtrErr(someLogicBoolUint8PtrErr, []*bool{&vf, &vf, &vt}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapBoolUint8PtrErr failed")
	}

	_, err = PMapBoolUint8PtrErr(someLogicBoolUint8PtrErr, []*bool{&vf, &vt, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolUint8PtrErr failed")
	}
	_, err = PMapBoolUint8PtrErr(someLogicBoolUint8PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolUint8PtrErr failed")
	}

	expectedList = []*uint8{&v10, &v10}
	newList, _ = PMapBoolUint8PtrErr(someLogicBoolUint8PtrErr, []*bool{&vt, &vt}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapBoolUint8Ptr failed")
	}

	_, err = PMapBoolUint8PtrErr(someLogicBoolUint8PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolUint8PtrErr failed")
	}
	_, err = PMapBoolUint8PtrErr(someLogicBoolUint8PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolUint8PtrErr failed")
	}

	_, err = PMapBoolUint8PtrErr(someLogicBoolUint8PtrErr, []*bool{&vf, &vf, &vt}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolUint8PtrErr failed")
	}

	_, err = PMapBoolUint8PtrErr(someLogicBoolUint8PtrErr, []*bool{&vf, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolUint8PtrErr failed")
	}

	_, err = PMapBoolUint8PtrErr(someLogicBoolUint8PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolUint8PtrErr failed")
	}
}

func TestPmapBoolStrPtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 string = "10"
	var v0 string = "0"

	expectedList := []*string{&v10, &v0}
	newList, _ := PMapBoolStrPtrErr(someLogicBoolStrPtrErr, []*bool{&vt, &vt})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolStrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapBoolStrPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapBoolStrPtrErr failed")
	}

	r, _ = PMapBoolStrPtrErr(nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("PMapBoolStrPtrErr failed")
	}
	_, err := PMapBoolStrPtrErr(someLogicBoolStrPtrErr, []*bool{&vt, &vf})
	if err == nil {
		t.Errorf("PMapBoolStrPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapBoolStrPtrErr(someLogicBoolStrPtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolStrPtrErr failed")
	}

	_, err = PMapBoolStrPtrErr(someLogicBoolStrPtrErr, []*bool{&vf, &vf, &vt}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapBoolStrPtrErr failed")
	}

	_, err = PMapBoolStrPtrErr(someLogicBoolStrPtrErr, []*bool{&vf, &vt, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolStrPtrErr failed")
	}
	_, err = PMapBoolStrPtrErr(someLogicBoolStrPtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolStrPtrErr failed")
	}

	expectedList = []*string{&v10, &v10}
	newList, _ = PMapBoolStrPtrErr(someLogicBoolStrPtrErr, []*bool{&vt, &vt}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapBoolStrPtrErr failed")
	}

	_, err = PMapBoolStrPtrErr(someLogicBoolStrPtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolStrPtrErr failed")
	}
	_, err = PMapBoolStrPtrErr(someLogicBoolStrPtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolStrPtrErr failed")
	}

	_, err = PMapBoolStrPtrErr(someLogicBoolStrPtrErr, []*bool{&vf, &vf, &vt}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolStrPtrErr failed")
	}

	_, err = PMapBoolStrPtrErr(someLogicBoolStrPtrErr, []*bool{&vf, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolStrPtrErr failed")
	}

	_, err = PMapBoolStrPtrErr(someLogicBoolStrPtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolStrPtrErr failed")
	}
}

func TestPmapBoolFloat32PtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 float32 = 10
	var v0 float32 = 0

	expectedList := []*float32{&v10, &v0}
	newList, _ := PMapBoolFloat32PtrErr(someLogicBoolFloat32PtrErr, []*bool{&vt, &vt})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolFloat32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapBoolFloat32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapBoolFloat32PtrErr failed")
	}

	r, _ = PMapBoolFloat32PtrErr(nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("PMapBoolFloat32PtrErr failed")
	}
	_, err := PMapBoolFloat32PtrErr(someLogicBoolFloat32PtrErr, []*bool{&vt, &vf})
	if err == nil {
		t.Errorf("PMapBoolFloat32PtrErr failed")
	}
	reflect.TypeOf("Ram") // Leaving it here to make use of import reflect

	_, err = PMapBoolFloat32PtrErr(someLogicBoolFloat32PtrErr, []*bool{&vt, &vt, &vf, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolFloat32PtrErr failed")
	}

	_, err = PMapBoolFloat32PtrErr(someLogicBoolFloat32PtrErr, []*bool{&vf, &vf, &vt}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapBoolFloat32PtrErr failed")
	}

	_, err = PMapBoolFloat32PtrErr(someLogicBoolFloat32PtrErr, []*bool{&vf, &vt, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolFloat32PtrErr failed")
	}
	_, err = PMapBoolFloat32PtrErr(someLogicBoolFloat32PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolFloat32PtrErr failed")
	}

	expectedList = []*float32{&v10, &v10}
	newList, _ = PMapBoolFloat32PtrErr(someLogicBoolFloat32PtrErr, []*bool{&vt, &vt}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapBoolFloat32Ptr failed")
	}

	_, err = PMapBoolFloat32PtrErr(someLogicBoolFloat32PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolFloat32PtrErr failed")
	}
	_, err = PMapBoolFloat32PtrErr(someLogicBoolFloat32PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolFloat32PtrErr failed")
	}

	_, err = PMapBoolFloat32PtrErr(someLogicBoolFloat32PtrErr, []*bool{&vf, &vf, &vt}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolFloat32PtrErr failed")
	}

	_, err = PMapBoolFloat32PtrErr(someLogicBoolFloat32PtrErr, []*bool{&vf, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolFloat32PtrErr failed")
	}

	_, err = PMapBoolFloat32PtrErr(someLogicBoolFloat32PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolFloat32PtrErr failed")
	}
}

func TestPmapBoolFloat64PtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 float64 = 10
	var v0 float64 = 0

	expectedList := []*float64{&v10, &v0}
	newList, _ := PMapBoolFloat64PtrErr(someLogicBoolFloat64PtrErr, []*bool{&vt, &vt})

	if *newList[0] != *expectedList[0] && *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolFloat64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapBoolFloat64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapBoolFloat64PtrErr failed")
	}

	r, _ = PMapBoolFloat64PtrErr(nil, []*bool{})
	if len(r) > 0 {
		t.Errorf("PMapBoolFloat64PtrErr failed")
	}
	_, err := PMapBoolFloat64PtrErr(someLogicBoolFloat64PtrErr, []*bool{&vt, &vf})
	if err == nil {
		t.Errorf("PMapBoolFloat64PtrErr failed")
	}
	reflect.TypeOf("Ram") // Leaving it here to make use of import reflect

	_, err = PMapBoolFloat64PtrErr(someLogicBoolFloat64PtrErr, []*bool{&vt, &vt, &vf, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolFloat64PtrErr failed")
	}

	_, err = PMapBoolFloat64PtrErr(someLogicBoolFloat64PtrErr, []*bool{&vf, &vf, &vt}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapBoolFloat64PtrErr failed")
	}

	_, err = PMapBoolFloat64PtrErr(someLogicBoolFloat64PtrErr, []*bool{&vf, &vt, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolFloat64PtrErr failed")
	}
	_, err = PMapBoolFloat64PtrErr(someLogicBoolFloat64PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapBoolFloat64PtrErr failed")
	}

	expectedList = []*float64{&v10, &v10}
	newList, _ = PMapBoolFloat64PtrErr(someLogicBoolFloat64PtrErr, []*bool{&vt, &vt}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapBoolFloat64Ptr failed")
	}

	_, err = PMapBoolFloat64PtrErr(someLogicBoolFloat64PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolFloat64PtrErr failed")
	}
	_, err = PMapBoolFloat64PtrErr(someLogicBoolFloat64PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolFloat64PtrErr failed")
	}

	_, err = PMapBoolFloat64PtrErr(someLogicBoolFloat64PtrErr, []*bool{&vf, &vf, &vt}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolFloat64PtrErr failed")
	}

	_, err = PMapBoolFloat64PtrErr(someLogicBoolFloat64PtrErr, []*bool{&vf, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolFloat64PtrErr failed")
	}

	_, err = PMapBoolFloat64PtrErr(someLogicBoolFloat64PtrErr, []*bool{&vt, &vt, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolFloat64PtrErr failed")
	}
}

func TestPmapFloat32IntPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3

	expectedList := []*int{&vo2, &vo3}
	newList, _ := PMapFloat32IntPtrErr(plusOneFloat32IntPtrErr, []*float32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapFloat32IntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat32IntPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat32IntPtrErr failed")
	}

	r, _ = PMapFloat32IntPtrErr(nil, []*float32{})
	if len(r) > 0 {
		t.Errorf("PMapFloat32IntPtrErr failed")
	}
	_, err := PMapFloat32IntPtrErr(plusOneFloat32IntPtrErr, []*float32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapFloat32IntPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapFloat32IntPtrErr(plusOneFloat32IntPtrErr, []*float32{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32IntPtrErr failed")
	}

	_, err = PMapFloat32IntPtrErr(plusOneFloat32IntPtrErr, []*float32{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapFloat32IntPtrErr failed")
	}

	_, err = PMapFloat32IntPtrErr(plusOneFloat32IntPtrErr, []*float32{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32IntPtrErr failed")
	}
	_, err = PMapFloat32IntPtrErr(plusOneFloat32IntPtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32IntPtrErr failed")
	}

	expectedList = []*int{&vo2, &vo3}
	newList, _ = PMapFloat32IntPtrErr(plusOneFloat32IntPtrErr, []*float32{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapFloat32IntPtr failed")
	}

	_, err = PMapFloat32IntPtrErr(plusOneFloat32IntPtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32IntPtrErr failed")
	}
	_, err = PMapFloat32IntPtrErr(plusOneFloat32IntPtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32IntPtrErr failed")
	}

	_, err = PMapFloat32IntPtrErr(plusOneFloat32IntPtrErr, []*float32{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32IntPtrErr failed")
	}

	_, err = PMapFloat32IntPtrErr(plusOneFloat32IntPtrErr, []*float32{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32IntPtrErr failed")
	}

	_, err = PMapFloat32IntPtrErr(plusOneFloat32IntPtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32IntPtrErr failed")
	}
}

func TestPmapFloat32Int64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3

	expectedList := []*int64{&vo2, &vo3}
	newList, _ := PMapFloat32Int64PtrErr(plusOneFloat32Int64PtrErr, []*float32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapFloat32Int64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat32Int64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat32Int64PtrErr failed")
	}

	r, _ = PMapFloat32Int64PtrErr(nil, []*float32{})
	if len(r) > 0 {
		t.Errorf("PMapFloat32Int64PtrErr failed")
	}
	_, err := PMapFloat32Int64PtrErr(plusOneFloat32Int64PtrErr, []*float32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapFloat32Int64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapFloat32Int64PtrErr(plusOneFloat32Int64PtrErr, []*float32{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32Int64PtrErr failed")
	}

	_, err = PMapFloat32Int64PtrErr(plusOneFloat32Int64PtrErr, []*float32{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapFloat32Int64PtrErr failed")
	}

	_, err = PMapFloat32Int64PtrErr(plusOneFloat32Int64PtrErr, []*float32{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32Int64PtrErr failed")
	}
	_, err = PMapFloat32Int64PtrErr(plusOneFloat32Int64PtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32Int64PtrErr failed")
	}

	expectedList = []*int64{&vo2, &vo3}
	newList, _ = PMapFloat32Int64PtrErr(plusOneFloat32Int64PtrErr, []*float32{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapFloat32Int64Ptr failed")
	}

	_, err = PMapFloat32Int64PtrErr(plusOneFloat32Int64PtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Int64PtrErr failed")
	}
	_, err = PMapFloat32Int64PtrErr(plusOneFloat32Int64PtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Int64PtrErr failed")
	}

	_, err = PMapFloat32Int64PtrErr(plusOneFloat32Int64PtrErr, []*float32{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Int64PtrErr failed")
	}

	_, err = PMapFloat32Int64PtrErr(plusOneFloat32Int64PtrErr, []*float32{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Int64PtrErr failed")
	}

	_, err = PMapFloat32Int64PtrErr(plusOneFloat32Int64PtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Int64PtrErr failed")
	}
}

func TestPmapFloat32Int32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3

	expectedList := []*int32{&vo2, &vo3}
	newList, _ := PMapFloat32Int32PtrErr(plusOneFloat32Int32PtrErr, []*float32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapFloat32Int32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat32Int32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat32Int32PtrErr failed")
	}

	r, _ = PMapFloat32Int32PtrErr(nil, []*float32{})
	if len(r) > 0 {
		t.Errorf("PMapFloat32Int32PtrErr failed")
	}
	_, err := PMapFloat32Int32PtrErr(plusOneFloat32Int32PtrErr, []*float32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapFloat32Int32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapFloat32Int32PtrErr(plusOneFloat32Int32PtrErr, []*float32{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32Int32PtrErr failed")
	}

	_, err = PMapFloat32Int32PtrErr(plusOneFloat32Int32PtrErr, []*float32{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapFloat32Int32PtrErr failed")
	}

	_, err = PMapFloat32Int32PtrErr(plusOneFloat32Int32PtrErr, []*float32{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32Int32PtrErr failed")
	}
	_, err = PMapFloat32Int32PtrErr(plusOneFloat32Int32PtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32Int32PtrErr failed")
	}

	expectedList = []*int32{&vo2, &vo3}
	newList, _ = PMapFloat32Int32PtrErr(plusOneFloat32Int32PtrErr, []*float32{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapFloat32Int32Ptr failed")
	}

	_, err = PMapFloat32Int32PtrErr(plusOneFloat32Int32PtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Int32PtrErr failed")
	}
	_, err = PMapFloat32Int32PtrErr(plusOneFloat32Int32PtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Int32PtrErr failed")
	}

	_, err = PMapFloat32Int32PtrErr(plusOneFloat32Int32PtrErr, []*float32{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Int32PtrErr failed")
	}

	_, err = PMapFloat32Int32PtrErr(plusOneFloat32Int32PtrErr, []*float32{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Int32PtrErr failed")
	}

	_, err = PMapFloat32Int32PtrErr(plusOneFloat32Int32PtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Int32PtrErr failed")
	}
}

func TestPmapFloat32Int16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3

	expectedList := []*int16{&vo2, &vo3}
	newList, _ := PMapFloat32Int16PtrErr(plusOneFloat32Int16PtrErr, []*float32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapFloat32Int16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat32Int16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat32Int16PtrErr failed")
	}

	r, _ = PMapFloat32Int16PtrErr(nil, []*float32{})
	if len(r) > 0 {
		t.Errorf("PMapFloat32Int16PtrErr failed")
	}
	_, err := PMapFloat32Int16PtrErr(plusOneFloat32Int16PtrErr, []*float32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapFloat32Int16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapFloat32Int16PtrErr(plusOneFloat32Int16PtrErr, []*float32{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32Int16PtrErr failed")
	}

	_, err = PMapFloat32Int16PtrErr(plusOneFloat32Int16PtrErr, []*float32{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapFloat32Int16PtrErr failed")
	}

	_, err = PMapFloat32Int16PtrErr(plusOneFloat32Int16PtrErr, []*float32{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32Int16PtrErr failed")
	}
	_, err = PMapFloat32Int16PtrErr(plusOneFloat32Int16PtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32Int16PtrErr failed")
	}

	expectedList = []*int16{&vo2, &vo3}
	newList, _ = PMapFloat32Int16PtrErr(plusOneFloat32Int16PtrErr, []*float32{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapFloat32Int16Ptr failed")
	}

	_, err = PMapFloat32Int16PtrErr(plusOneFloat32Int16PtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Int16PtrErr failed")
	}
	_, err = PMapFloat32Int16PtrErr(plusOneFloat32Int16PtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Int16PtrErr failed")
	}

	_, err = PMapFloat32Int16PtrErr(plusOneFloat32Int16PtrErr, []*float32{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Int16PtrErr failed")
	}

	_, err = PMapFloat32Int16PtrErr(plusOneFloat32Int16PtrErr, []*float32{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Int16PtrErr failed")
	}

	_, err = PMapFloat32Int16PtrErr(plusOneFloat32Int16PtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Int16PtrErr failed")
	}
}

func TestPmapFloat32Int8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3

	expectedList := []*int8{&vo2, &vo3}
	newList, _ := PMapFloat32Int8PtrErr(plusOneFloat32Int8PtrErr, []*float32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapFloat32Int8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat32Int8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat32Int8PtrErr failed")
	}

	r, _ = PMapFloat32Int8PtrErr(nil, []*float32{})
	if len(r) > 0 {
		t.Errorf("PMapFloat32Int8PtrErr failed")
	}
	_, err := PMapFloat32Int8PtrErr(plusOneFloat32Int8PtrErr, []*float32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapFloat32Int8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapFloat32Int8PtrErr(plusOneFloat32Int8PtrErr, []*float32{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32Int8PtrErr failed")
	}

	_, err = PMapFloat32Int8PtrErr(plusOneFloat32Int8PtrErr, []*float32{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapFloat32Int8PtrErr failed")
	}

	_, err = PMapFloat32Int8PtrErr(plusOneFloat32Int8PtrErr, []*float32{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32Int8PtrErr failed")
	}
	_, err = PMapFloat32Int8PtrErr(plusOneFloat32Int8PtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32Int8PtrErr failed")
	}

	expectedList = []*int8{&vo2, &vo3}
	newList, _ = PMapFloat32Int8PtrErr(plusOneFloat32Int8PtrErr, []*float32{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapFloat32Int8Ptr failed")
	}

	_, err = PMapFloat32Int8PtrErr(plusOneFloat32Int8PtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Int8PtrErr failed")
	}
	_, err = PMapFloat32Int8PtrErr(plusOneFloat32Int8PtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Int8PtrErr failed")
	}

	_, err = PMapFloat32Int8PtrErr(plusOneFloat32Int8PtrErr, []*float32{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Int8PtrErr failed")
	}

	_, err = PMapFloat32Int8PtrErr(plusOneFloat32Int8PtrErr, []*float32{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Int8PtrErr failed")
	}

	_, err = PMapFloat32Int8PtrErr(plusOneFloat32Int8PtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Int8PtrErr failed")
	}
}

func TestPmapFloat32UintPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3

	expectedList := []*uint{&vo2, &vo3}
	newList, _ := PMapFloat32UintPtrErr(plusOneFloat32UintPtrErr, []*float32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapFloat32UintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat32UintPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat32UintPtrErr failed")
	}

	r, _ = PMapFloat32UintPtrErr(nil, []*float32{})
	if len(r) > 0 {
		t.Errorf("PMapFloat32UintPtrErr failed")
	}
	_, err := PMapFloat32UintPtrErr(plusOneFloat32UintPtrErr, []*float32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapFloat32UintPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapFloat32UintPtrErr(plusOneFloat32UintPtrErr, []*float32{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32UintPtrErr failed")
	}

	_, err = PMapFloat32UintPtrErr(plusOneFloat32UintPtrErr, []*float32{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapFloat32UintPtrErr failed")
	}

	_, err = PMapFloat32UintPtrErr(plusOneFloat32UintPtrErr, []*float32{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32UintPtrErr failed")
	}
	_, err = PMapFloat32UintPtrErr(plusOneFloat32UintPtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32UintPtrErr failed")
	}

	expectedList = []*uint{&vo2, &vo3}
	newList, _ = PMapFloat32UintPtrErr(plusOneFloat32UintPtrErr, []*float32{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapFloat32UintPtr failed")
	}

	_, err = PMapFloat32UintPtrErr(plusOneFloat32UintPtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32UintPtrErr failed")
	}
	_, err = PMapFloat32UintPtrErr(plusOneFloat32UintPtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32UintPtrErr failed")
	}

	_, err = PMapFloat32UintPtrErr(plusOneFloat32UintPtrErr, []*float32{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32UintPtrErr failed")
	}

	_, err = PMapFloat32UintPtrErr(plusOneFloat32UintPtrErr, []*float32{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32UintPtrErr failed")
	}

	_, err = PMapFloat32UintPtrErr(plusOneFloat32UintPtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32UintPtrErr failed")
	}
}

func TestPmapFloat32Uint64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3

	expectedList := []*uint64{&vo2, &vo3}
	newList, _ := PMapFloat32Uint64PtrErr(plusOneFloat32Uint64PtrErr, []*float32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapFloat32Uint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat32Uint64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat32Uint64PtrErr failed")
	}

	r, _ = PMapFloat32Uint64PtrErr(nil, []*float32{})
	if len(r) > 0 {
		t.Errorf("PMapFloat32Uint64PtrErr failed")
	}
	_, err := PMapFloat32Uint64PtrErr(plusOneFloat32Uint64PtrErr, []*float32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapFloat32Uint64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapFloat32Uint64PtrErr(plusOneFloat32Uint64PtrErr, []*float32{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32Uint64PtrErr failed")
	}

	_, err = PMapFloat32Uint64PtrErr(plusOneFloat32Uint64PtrErr, []*float32{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapFloat32Uint64PtrErr failed")
	}

	_, err = PMapFloat32Uint64PtrErr(plusOneFloat32Uint64PtrErr, []*float32{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32Uint64PtrErr failed")
	}
	_, err = PMapFloat32Uint64PtrErr(plusOneFloat32Uint64PtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32Uint64PtrErr failed")
	}

	expectedList = []*uint64{&vo2, &vo3}
	newList, _ = PMapFloat32Uint64PtrErr(plusOneFloat32Uint64PtrErr, []*float32{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapFloat32Uint64Ptr failed")
	}

	_, err = PMapFloat32Uint64PtrErr(plusOneFloat32Uint64PtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Uint64PtrErr failed")
	}
	_, err = PMapFloat32Uint64PtrErr(plusOneFloat32Uint64PtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Uint64PtrErr failed")
	}

	_, err = PMapFloat32Uint64PtrErr(plusOneFloat32Uint64PtrErr, []*float32{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Uint64PtrErr failed")
	}

	_, err = PMapFloat32Uint64PtrErr(plusOneFloat32Uint64PtrErr, []*float32{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Uint64PtrErr failed")
	}

	_, err = PMapFloat32Uint64PtrErr(plusOneFloat32Uint64PtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Uint64PtrErr failed")
	}
}

func TestPmapFloat32Uint32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3

	expectedList := []*uint32{&vo2, &vo3}
	newList, _ := PMapFloat32Uint32PtrErr(plusOneFloat32Uint32PtrErr, []*float32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapFloat32Uint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat32Uint32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat32Uint32PtrErr failed")
	}

	r, _ = PMapFloat32Uint32PtrErr(nil, []*float32{})
	if len(r) > 0 {
		t.Errorf("PMapFloat32Uint32PtrErr failed")
	}
	_, err := PMapFloat32Uint32PtrErr(plusOneFloat32Uint32PtrErr, []*float32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapFloat32Uint32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapFloat32Uint32PtrErr(plusOneFloat32Uint32PtrErr, []*float32{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32Uint32PtrErr failed")
	}

	_, err = PMapFloat32Uint32PtrErr(plusOneFloat32Uint32PtrErr, []*float32{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapFloat32Uint32PtrErr failed")
	}

	_, err = PMapFloat32Uint32PtrErr(plusOneFloat32Uint32PtrErr, []*float32{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32Uint32PtrErr failed")
	}
	_, err = PMapFloat32Uint32PtrErr(plusOneFloat32Uint32PtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32Uint32PtrErr failed")
	}

	expectedList = []*uint32{&vo2, &vo3}
	newList, _ = PMapFloat32Uint32PtrErr(plusOneFloat32Uint32PtrErr, []*float32{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapFloat32Uint32Ptr failed")
	}

	_, err = PMapFloat32Uint32PtrErr(plusOneFloat32Uint32PtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Uint32PtrErr failed")
	}
	_, err = PMapFloat32Uint32PtrErr(plusOneFloat32Uint32PtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Uint32PtrErr failed")
	}

	_, err = PMapFloat32Uint32PtrErr(plusOneFloat32Uint32PtrErr, []*float32{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Uint32PtrErr failed")
	}

	_, err = PMapFloat32Uint32PtrErr(plusOneFloat32Uint32PtrErr, []*float32{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Uint32PtrErr failed")
	}

	_, err = PMapFloat32Uint32PtrErr(plusOneFloat32Uint32PtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Uint32PtrErr failed")
	}
}

func TestPmapFloat32Uint16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3

	expectedList := []*uint16{&vo2, &vo3}
	newList, _ := PMapFloat32Uint16PtrErr(plusOneFloat32Uint16PtrErr, []*float32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapFloat32Uint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat32Uint16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat32Uint16PtrErr failed")
	}

	r, _ = PMapFloat32Uint16PtrErr(nil, []*float32{})
	if len(r) > 0 {
		t.Errorf("PMapFloat32Uint16PtrErr failed")
	}
	_, err := PMapFloat32Uint16PtrErr(plusOneFloat32Uint16PtrErr, []*float32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapFloat32Uint16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapFloat32Uint16PtrErr(plusOneFloat32Uint16PtrErr, []*float32{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32Uint16PtrErr failed")
	}

	_, err = PMapFloat32Uint16PtrErr(plusOneFloat32Uint16PtrErr, []*float32{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapFloat32Uint16PtrErr failed")
	}

	_, err = PMapFloat32Uint16PtrErr(plusOneFloat32Uint16PtrErr, []*float32{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32Uint16PtrErr failed")
	}
	_, err = PMapFloat32Uint16PtrErr(plusOneFloat32Uint16PtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32Uint16PtrErr failed")
	}

	expectedList = []*uint16{&vo2, &vo3}
	newList, _ = PMapFloat32Uint16PtrErr(plusOneFloat32Uint16PtrErr, []*float32{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapFloat32Uint16Ptr failed")
	}

	_, err = PMapFloat32Uint16PtrErr(plusOneFloat32Uint16PtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Uint16PtrErr failed")
	}
	_, err = PMapFloat32Uint16PtrErr(plusOneFloat32Uint16PtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Uint16PtrErr failed")
	}

	_, err = PMapFloat32Uint16PtrErr(plusOneFloat32Uint16PtrErr, []*float32{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Uint16PtrErr failed")
	}

	_, err = PMapFloat32Uint16PtrErr(plusOneFloat32Uint16PtrErr, []*float32{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Uint16PtrErr failed")
	}

	_, err = PMapFloat32Uint16PtrErr(plusOneFloat32Uint16PtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Uint16PtrErr failed")
	}
}

func TestPmapFloat32Uint8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3

	expectedList := []*uint8{&vo2, &vo3}
	newList, _ := PMapFloat32Uint8PtrErr(plusOneFloat32Uint8PtrErr, []*float32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapFloat32Uint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat32Uint8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat32Uint8PtrErr failed")
	}

	r, _ = PMapFloat32Uint8PtrErr(nil, []*float32{})
	if len(r) > 0 {
		t.Errorf("PMapFloat32Uint8PtrErr failed")
	}
	_, err := PMapFloat32Uint8PtrErr(plusOneFloat32Uint8PtrErr, []*float32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapFloat32Uint8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapFloat32Uint8PtrErr(plusOneFloat32Uint8PtrErr, []*float32{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32Uint8PtrErr failed")
	}

	_, err = PMapFloat32Uint8PtrErr(plusOneFloat32Uint8PtrErr, []*float32{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapFloat32Uint8PtrErr failed")
	}

	_, err = PMapFloat32Uint8PtrErr(plusOneFloat32Uint8PtrErr, []*float32{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32Uint8PtrErr failed")
	}
	_, err = PMapFloat32Uint8PtrErr(plusOneFloat32Uint8PtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32Uint8PtrErr failed")
	}

	expectedList = []*uint8{&vo2, &vo3}
	newList, _ = PMapFloat32Uint8PtrErr(plusOneFloat32Uint8PtrErr, []*float32{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapFloat32Uint8Ptr failed")
	}

	_, err = PMapFloat32Uint8PtrErr(plusOneFloat32Uint8PtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Uint8PtrErr failed")
	}
	_, err = PMapFloat32Uint8PtrErr(plusOneFloat32Uint8PtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Uint8PtrErr failed")
	}

	_, err = PMapFloat32Uint8PtrErr(plusOneFloat32Uint8PtrErr, []*float32{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Uint8PtrErr failed")
	}

	_, err = PMapFloat32Uint8PtrErr(plusOneFloat32Uint8PtrErr, []*float32{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Uint8PtrErr failed")
	}

	_, err = PMapFloat32Uint8PtrErr(plusOneFloat32Uint8PtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Uint8PtrErr failed")
	}
}

func TestPmapFloat32StrPtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 float32 = 10
	var vi0 float32

	expectedList := []*string{&vo10}
	newList, _ := PMapFloat32StrPtrErr(someLogicFloat32StrPtrErr, []*float32{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapFloat32StrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat32StrPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat32StrPtrErr failed")
	}

	r, _ = PMapFloat32StrPtrErr(nil, []*float32{})
	if len(r) > 0 {
		t.Errorf("PMapFloat32StrPtrErr failed")
	}
	_, err := PMapFloat32StrPtrErr(someLogicFloat32StrPtrErr, []*float32{&vi10, &vi0})
	if err == nil {
		t.Errorf("PMapFloat32StrPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapFloat32StrPtrErr(someLogicFloat32StrPtrErr, []*float32{&vi10, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32StrPtrErr failed")
	}

	_, err = PMapFloat32StrPtrErr(someLogicFloat32StrPtrErr, []*float32{&vi0, &vi0, &vi10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapFloat32StrPtrErr failed")
	}

	_, err = PMapFloat32StrPtrErr(someLogicFloat32StrPtrErr, []*float32{&vi0, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32StrPtrErr failed")
	}
	_, err = PMapFloat32StrPtrErr(someLogicFloat32StrPtrErr, []*float32{&vi10, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32StrPtrErr failed")
	}

	expectedList = []*string{&vo10, &vo10}
	newList, _ = PMapFloat32StrPtrErr(someLogicFloat32StrPtrErr, []*float32{&vi10, &vi10}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapFloat32StrPtr failed")
	}

	_, err = PMapFloat32StrPtrErr(someLogicFloat32StrPtrErr, []*float32{&vi10, &vi10, &vi0}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32StrPtrErr failed")
	}
	_, err = PMapFloat32StrPtrErr(someLogicFloat32StrPtrErr, []*float32{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32StrPtrErr failed")
	}

	_, err = PMapFloat32StrPtrErr(someLogicFloat32StrPtrErr, []*float32{&vi0, &vi0, &vi10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32StrPtrErr failed")
	}

	_, err = PMapFloat32StrPtrErr(someLogicFloat32StrPtrErr, []*float32{&vi0, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32StrPtrErr failed")
	}

	_, err = PMapFloat32StrPtrErr(someLogicFloat32StrPtrErr, []*float32{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32StrPtrErr failed")
	}
}

func TestPmapFloat32BoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 float32 = 10
	var v0 float32
	var v3 float32 = 3

	expectedList := []*bool{&vt, &vf}
	newList, _ := PMapFloat32BoolPtrErr(someLogicFloat32BoolPtrErr, []*float32{&v10, &v0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapFloat32BoolPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat32BoolPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat32BoolPtrErr failed")
	}

	r, _ = PMapFloat32BoolPtrErr(nil, []*float32{})
	if len(r) > 0 {
		t.Errorf("PMapFloat32BoolPtrErr failed")
	}
	_, err := PMapFloat32BoolPtrErr(someLogicFloat32BoolPtrErr, []*float32{&v10, &v3})
	if err == nil {
		t.Errorf("PMapFloat32BoolPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapFloat32BoolPtrErr(someLogicFloat32BoolPtrErr, []*float32{&v10, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32BoolPtrErr failed")
	}

	_, err = PMapFloat32BoolPtrErr(someLogicFloat32BoolPtrErr, []*float32{&v3, &v3, &v10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapFloat32BoolPtrErr failed")
	}

	_, err = PMapFloat32BoolPtrErr(someLogicFloat32BoolPtrErr, []*float32{&v3, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32BoolErr failed")
	}
	_, err = PMapFloat32BoolPtrErr(someLogicFloat32BoolPtrErr, []*float32{&v10, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32BoolPtrErr failed")
	}

	expectedList = []*bool{&vt, &vt}
	newList, _ = PMapFloat32BoolPtrErr(someLogicFloat32BoolPtrErr, []*float32{&v10, &v10}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapFloat32BoolPtr failed")
	}

	_, err = PMapFloat32BoolPtrErr(someLogicFloat32BoolPtrErr, []*float32{&v10, &v10, &v3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32BoolErr failed")
	}
	_, err = PMapFloat32BoolPtrErr(someLogicFloat32BoolPtrErr, []*float32{&v10, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32BoolPtrErr failed")
	}

	_, err = PMapFloat32BoolPtrErr(someLogicFloat32BoolPtrErr, []*float32{&v3, &v3, &v10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32BoolPtrErr failed")
	}

	_, err = PMapFloat32BoolPtrErr(someLogicFloat32BoolPtrErr, []*float32{&v3, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32BoolPtrErr failed")
	}

	_, err = PMapFloat32BoolPtrErr(someLogicFloat32BoolPtrErr, []*float32{&v10, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32BoolPtrErr failed")
	}
}

func TestPmapFloat32Float64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3

	expectedList := []*float64{&vo2, &vo3}
	newList, _ := PMapFloat32Float64PtrErr(plusOneFloat32Float64PtrErr, []*float32{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapFloat32Float64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat32Float64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat32Float64PtrErr failed")
	}

	r, _ = PMapFloat32Float64PtrErr(nil, []*float32{})
	if len(r) > 0 {
		t.Errorf("PMapFloat32Float64PtrErr failed")
	}
	_, err := PMapFloat32Float64PtrErr(plusOneFloat32Float64PtrErr, []*float32{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapFloat32Float64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapFloat32Float64PtrErr(plusOneFloat32Float64PtrErr, []*float32{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32Float64PtrErr failed")
	}

	_, err = PMapFloat32Float64PtrErr(plusOneFloat32Float64PtrErr, []*float32{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapFloat32Float64PtrErr failed")
	}

	_, err = PMapFloat32Float64PtrErr(plusOneFloat32Float64PtrErr, []*float32{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32Float64PtrErr failed")
	}
	_, err = PMapFloat32Float64PtrErr(plusOneFloat32Float64PtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32Float64PtrErr failed")
	}

	expectedList = []*float64{&vo2, &vo3}
	newList, _ = PMapFloat32Float64PtrErr(plusOneFloat32Float64PtrErr, []*float32{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapFloat32Float64Ptr failed")
	}

	_, err = PMapFloat32Float64PtrErr(plusOneFloat32Float64PtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Float64PtrErr failed")
	}
	_, err = PMapFloat32Float64PtrErr(plusOneFloat32Float64PtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Float64PtrErr failed")
	}

	_, err = PMapFloat32Float64PtrErr(plusOneFloat32Float64PtrErr, []*float32{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Float64PtrErr failed")
	}

	_, err = PMapFloat32Float64PtrErr(plusOneFloat32Float64PtrErr, []*float32{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Float64PtrErr failed")
	}

	_, err = PMapFloat32Float64PtrErr(plusOneFloat32Float64PtrErr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat32Float64PtrErr failed")
	}
}

func TestPmapFloat64IntPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3

	expectedList := []*int{&vo2, &vo3}
	newList, _ := PMapFloat64IntPtrErr(plusOneFloat64IntPtrErr, []*float64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapFloat64IntPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat64IntPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat64IntPtrErr failed")
	}

	r, _ = PMapFloat64IntPtrErr(nil, []*float64{})
	if len(r) > 0 {
		t.Errorf("PMapFloat64IntPtrErr failed")
	}
	_, err := PMapFloat64IntPtrErr(plusOneFloat64IntPtrErr, []*float64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapFloat64IntPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapFloat64IntPtrErr(plusOneFloat64IntPtrErr, []*float64{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64IntPtrErr failed")
	}

	_, err = PMapFloat64IntPtrErr(plusOneFloat64IntPtrErr, []*float64{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapFloat64IntPtrErr failed")
	}

	_, err = PMapFloat64IntPtrErr(plusOneFloat64IntPtrErr, []*float64{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64IntPtrErr failed")
	}
	_, err = PMapFloat64IntPtrErr(plusOneFloat64IntPtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64IntPtrErr failed")
	}

	expectedList = []*int{&vo2, &vo3}
	newList, _ = PMapFloat64IntPtrErr(plusOneFloat64IntPtrErr, []*float64{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapFloat64IntPtr failed")
	}

	_, err = PMapFloat64IntPtrErr(plusOneFloat64IntPtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64IntPtrErr failed")
	}
	_, err = PMapFloat64IntPtrErr(plusOneFloat64IntPtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64IntPtrErr failed")
	}

	_, err = PMapFloat64IntPtrErr(plusOneFloat64IntPtrErr, []*float64{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64IntPtrErr failed")
	}

	_, err = PMapFloat64IntPtrErr(plusOneFloat64IntPtrErr, []*float64{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64IntPtrErr failed")
	}

	_, err = PMapFloat64IntPtrErr(plusOneFloat64IntPtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64IntPtrErr failed")
	}
}

func TestPmapFloat64Int64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3

	expectedList := []*int64{&vo2, &vo3}
	newList, _ := PMapFloat64Int64PtrErr(plusOneFloat64Int64PtrErr, []*float64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapFloat64Int64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat64Int64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat64Int64PtrErr failed")
	}

	r, _ = PMapFloat64Int64PtrErr(nil, []*float64{})
	if len(r) > 0 {
		t.Errorf("PMapFloat64Int64PtrErr failed")
	}
	_, err := PMapFloat64Int64PtrErr(plusOneFloat64Int64PtrErr, []*float64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapFloat64Int64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapFloat64Int64PtrErr(plusOneFloat64Int64PtrErr, []*float64{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64Int64PtrErr failed")
	}

	_, err = PMapFloat64Int64PtrErr(plusOneFloat64Int64PtrErr, []*float64{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapFloat64Int64PtrErr failed")
	}

	_, err = PMapFloat64Int64PtrErr(plusOneFloat64Int64PtrErr, []*float64{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64Int64PtrErr failed")
	}
	_, err = PMapFloat64Int64PtrErr(plusOneFloat64Int64PtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64Int64PtrErr failed")
	}

	expectedList = []*int64{&vo2, &vo3}
	newList, _ = PMapFloat64Int64PtrErr(plusOneFloat64Int64PtrErr, []*float64{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapFloat64Int64Ptr failed")
	}

	_, err = PMapFloat64Int64PtrErr(plusOneFloat64Int64PtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Int64PtrErr failed")
	}
	_, err = PMapFloat64Int64PtrErr(plusOneFloat64Int64PtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Int64PtrErr failed")
	}

	_, err = PMapFloat64Int64PtrErr(plusOneFloat64Int64PtrErr, []*float64{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Int64PtrErr failed")
	}

	_, err = PMapFloat64Int64PtrErr(plusOneFloat64Int64PtrErr, []*float64{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Int64PtrErr failed")
	}

	_, err = PMapFloat64Int64PtrErr(plusOneFloat64Int64PtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Int64PtrErr failed")
	}
}

func TestPmapFloat64Int32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3

	expectedList := []*int32{&vo2, &vo3}
	newList, _ := PMapFloat64Int32PtrErr(plusOneFloat64Int32PtrErr, []*float64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapFloat64Int32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat64Int32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat64Int32PtrErr failed")
	}

	r, _ = PMapFloat64Int32PtrErr(nil, []*float64{})
	if len(r) > 0 {
		t.Errorf("PMapFloat64Int32PtrErr failed")
	}
	_, err := PMapFloat64Int32PtrErr(plusOneFloat64Int32PtrErr, []*float64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapFloat64Int32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapFloat64Int32PtrErr(plusOneFloat64Int32PtrErr, []*float64{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64Int32PtrErr failed")
	}

	_, err = PMapFloat64Int32PtrErr(plusOneFloat64Int32PtrErr, []*float64{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapFloat64Int32PtrErr failed")
	}

	_, err = PMapFloat64Int32PtrErr(plusOneFloat64Int32PtrErr, []*float64{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64Int32PtrErr failed")
	}
	_, err = PMapFloat64Int32PtrErr(plusOneFloat64Int32PtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64Int32PtrErr failed")
	}

	expectedList = []*int32{&vo2, &vo3}
	newList, _ = PMapFloat64Int32PtrErr(plusOneFloat64Int32PtrErr, []*float64{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapFloat64Int32Ptr failed")
	}

	_, err = PMapFloat64Int32PtrErr(plusOneFloat64Int32PtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Int32PtrErr failed")
	}
	_, err = PMapFloat64Int32PtrErr(plusOneFloat64Int32PtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Int32PtrErr failed")
	}

	_, err = PMapFloat64Int32PtrErr(plusOneFloat64Int32PtrErr, []*float64{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Int32PtrErr failed")
	}

	_, err = PMapFloat64Int32PtrErr(plusOneFloat64Int32PtrErr, []*float64{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Int32PtrErr failed")
	}

	_, err = PMapFloat64Int32PtrErr(plusOneFloat64Int32PtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Int32PtrErr failed")
	}
}

func TestPmapFloat64Int16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3

	expectedList := []*int16{&vo2, &vo3}
	newList, _ := PMapFloat64Int16PtrErr(plusOneFloat64Int16PtrErr, []*float64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapFloat64Int16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat64Int16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat64Int16PtrErr failed")
	}

	r, _ = PMapFloat64Int16PtrErr(nil, []*float64{})
	if len(r) > 0 {
		t.Errorf("PMapFloat64Int16PtrErr failed")
	}
	_, err := PMapFloat64Int16PtrErr(plusOneFloat64Int16PtrErr, []*float64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapFloat64Int16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapFloat64Int16PtrErr(plusOneFloat64Int16PtrErr, []*float64{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64Int16PtrErr failed")
	}

	_, err = PMapFloat64Int16PtrErr(plusOneFloat64Int16PtrErr, []*float64{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapFloat64Int16PtrErr failed")
	}

	_, err = PMapFloat64Int16PtrErr(plusOneFloat64Int16PtrErr, []*float64{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64Int16PtrErr failed")
	}
	_, err = PMapFloat64Int16PtrErr(plusOneFloat64Int16PtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64Int16PtrErr failed")
	}

	expectedList = []*int16{&vo2, &vo3}
	newList, _ = PMapFloat64Int16PtrErr(plusOneFloat64Int16PtrErr, []*float64{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapFloat64Int16Ptr failed")
	}

	_, err = PMapFloat64Int16PtrErr(plusOneFloat64Int16PtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Int16PtrErr failed")
	}
	_, err = PMapFloat64Int16PtrErr(plusOneFloat64Int16PtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Int16PtrErr failed")
	}

	_, err = PMapFloat64Int16PtrErr(plusOneFloat64Int16PtrErr, []*float64{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Int16PtrErr failed")
	}

	_, err = PMapFloat64Int16PtrErr(plusOneFloat64Int16PtrErr, []*float64{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Int16PtrErr failed")
	}

	_, err = PMapFloat64Int16PtrErr(plusOneFloat64Int16PtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Int16PtrErr failed")
	}
}

func TestPmapFloat64Int8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3

	expectedList := []*int8{&vo2, &vo3}
	newList, _ := PMapFloat64Int8PtrErr(plusOneFloat64Int8PtrErr, []*float64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapFloat64Int8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat64Int8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat64Int8PtrErr failed")
	}

	r, _ = PMapFloat64Int8PtrErr(nil, []*float64{})
	if len(r) > 0 {
		t.Errorf("PMapFloat64Int8PtrErr failed")
	}
	_, err := PMapFloat64Int8PtrErr(plusOneFloat64Int8PtrErr, []*float64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapFloat64Int8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapFloat64Int8PtrErr(plusOneFloat64Int8PtrErr, []*float64{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64Int8PtrErr failed")
	}

	_, err = PMapFloat64Int8PtrErr(plusOneFloat64Int8PtrErr, []*float64{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapFloat64Int8PtrErr failed")
	}

	_, err = PMapFloat64Int8PtrErr(plusOneFloat64Int8PtrErr, []*float64{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64Int8PtrErr failed")
	}
	_, err = PMapFloat64Int8PtrErr(plusOneFloat64Int8PtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64Int8PtrErr failed")
	}

	expectedList = []*int8{&vo2, &vo3}
	newList, _ = PMapFloat64Int8PtrErr(plusOneFloat64Int8PtrErr, []*float64{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapFloat64Int8Ptr failed")
	}

	_, err = PMapFloat64Int8PtrErr(plusOneFloat64Int8PtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Int8PtrErr failed")
	}
	_, err = PMapFloat64Int8PtrErr(plusOneFloat64Int8PtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Int8PtrErr failed")
	}

	_, err = PMapFloat64Int8PtrErr(plusOneFloat64Int8PtrErr, []*float64{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Int8PtrErr failed")
	}

	_, err = PMapFloat64Int8PtrErr(plusOneFloat64Int8PtrErr, []*float64{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Int8PtrErr failed")
	}

	_, err = PMapFloat64Int8PtrErr(plusOneFloat64Int8PtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Int8PtrErr failed")
	}
}

func TestPmapFloat64UintPtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3

	expectedList := []*uint{&vo2, &vo3}
	newList, _ := PMapFloat64UintPtrErr(plusOneFloat64UintPtrErr, []*float64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapFloat64UintPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat64UintPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat64UintPtrErr failed")
	}

	r, _ = PMapFloat64UintPtrErr(nil, []*float64{})
	if len(r) > 0 {
		t.Errorf("PMapFloat64UintPtrErr failed")
	}
	_, err := PMapFloat64UintPtrErr(plusOneFloat64UintPtrErr, []*float64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapFloat64UintPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapFloat64UintPtrErr(plusOneFloat64UintPtrErr, []*float64{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64UintPtrErr failed")
	}

	_, err = PMapFloat64UintPtrErr(plusOneFloat64UintPtrErr, []*float64{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapFloat64UintPtrErr failed")
	}

	_, err = PMapFloat64UintPtrErr(plusOneFloat64UintPtrErr, []*float64{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64UintPtrErr failed")
	}
	_, err = PMapFloat64UintPtrErr(plusOneFloat64UintPtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64UintPtrErr failed")
	}

	expectedList = []*uint{&vo2, &vo3}
	newList, _ = PMapFloat64UintPtrErr(plusOneFloat64UintPtrErr, []*float64{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapFloat64UintPtr failed")
	}

	_, err = PMapFloat64UintPtrErr(plusOneFloat64UintPtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64UintPtrErr failed")
	}
	_, err = PMapFloat64UintPtrErr(plusOneFloat64UintPtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64UintPtrErr failed")
	}

	_, err = PMapFloat64UintPtrErr(plusOneFloat64UintPtrErr, []*float64{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64UintPtrErr failed")
	}

	_, err = PMapFloat64UintPtrErr(plusOneFloat64UintPtrErr, []*float64{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64UintPtrErr failed")
	}

	_, err = PMapFloat64UintPtrErr(plusOneFloat64UintPtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64UintPtrErr failed")
	}
}

func TestPmapFloat64Uint64PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3

	expectedList := []*uint64{&vo2, &vo3}
	newList, _ := PMapFloat64Uint64PtrErr(plusOneFloat64Uint64PtrErr, []*float64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapFloat64Uint64PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat64Uint64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat64Uint64PtrErr failed")
	}

	r, _ = PMapFloat64Uint64PtrErr(nil, []*float64{})
	if len(r) > 0 {
		t.Errorf("PMapFloat64Uint64PtrErr failed")
	}
	_, err := PMapFloat64Uint64PtrErr(plusOneFloat64Uint64PtrErr, []*float64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapFloat64Uint64PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapFloat64Uint64PtrErr(plusOneFloat64Uint64PtrErr, []*float64{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64Uint64PtrErr failed")
	}

	_, err = PMapFloat64Uint64PtrErr(plusOneFloat64Uint64PtrErr, []*float64{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapFloat64Uint64PtrErr failed")
	}

	_, err = PMapFloat64Uint64PtrErr(plusOneFloat64Uint64PtrErr, []*float64{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64Uint64PtrErr failed")
	}
	_, err = PMapFloat64Uint64PtrErr(plusOneFloat64Uint64PtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64Uint64PtrErr failed")
	}

	expectedList = []*uint64{&vo2, &vo3}
	newList, _ = PMapFloat64Uint64PtrErr(plusOneFloat64Uint64PtrErr, []*float64{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapFloat64Uint64Ptr failed")
	}

	_, err = PMapFloat64Uint64PtrErr(plusOneFloat64Uint64PtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Uint64PtrErr failed")
	}
	_, err = PMapFloat64Uint64PtrErr(plusOneFloat64Uint64PtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Uint64PtrErr failed")
	}

	_, err = PMapFloat64Uint64PtrErr(plusOneFloat64Uint64PtrErr, []*float64{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Uint64PtrErr failed")
	}

	_, err = PMapFloat64Uint64PtrErr(plusOneFloat64Uint64PtrErr, []*float64{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Uint64PtrErr failed")
	}

	_, err = PMapFloat64Uint64PtrErr(plusOneFloat64Uint64PtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Uint64PtrErr failed")
	}
}

func TestPmapFloat64Uint32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3

	expectedList := []*uint32{&vo2, &vo3}
	newList, _ := PMapFloat64Uint32PtrErr(plusOneFloat64Uint32PtrErr, []*float64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapFloat64Uint32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat64Uint32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat64Uint32PtrErr failed")
	}

	r, _ = PMapFloat64Uint32PtrErr(nil, []*float64{})
	if len(r) > 0 {
		t.Errorf("PMapFloat64Uint32PtrErr failed")
	}
	_, err := PMapFloat64Uint32PtrErr(plusOneFloat64Uint32PtrErr, []*float64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapFloat64Uint32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapFloat64Uint32PtrErr(plusOneFloat64Uint32PtrErr, []*float64{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64Uint32PtrErr failed")
	}

	_, err = PMapFloat64Uint32PtrErr(plusOneFloat64Uint32PtrErr, []*float64{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapFloat64Uint32PtrErr failed")
	}

	_, err = PMapFloat64Uint32PtrErr(plusOneFloat64Uint32PtrErr, []*float64{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64Uint32PtrErr failed")
	}
	_, err = PMapFloat64Uint32PtrErr(plusOneFloat64Uint32PtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64Uint32PtrErr failed")
	}

	expectedList = []*uint32{&vo2, &vo3}
	newList, _ = PMapFloat64Uint32PtrErr(plusOneFloat64Uint32PtrErr, []*float64{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapFloat64Uint32Ptr failed")
	}

	_, err = PMapFloat64Uint32PtrErr(plusOneFloat64Uint32PtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Uint32PtrErr failed")
	}
	_, err = PMapFloat64Uint32PtrErr(plusOneFloat64Uint32PtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Uint32PtrErr failed")
	}

	_, err = PMapFloat64Uint32PtrErr(plusOneFloat64Uint32PtrErr, []*float64{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Uint32PtrErr failed")
	}

	_, err = PMapFloat64Uint32PtrErr(plusOneFloat64Uint32PtrErr, []*float64{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Uint32PtrErr failed")
	}

	_, err = PMapFloat64Uint32PtrErr(plusOneFloat64Uint32PtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Uint32PtrErr failed")
	}
}

func TestPmapFloat64Uint16PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3

	expectedList := []*uint16{&vo2, &vo3}
	newList, _ := PMapFloat64Uint16PtrErr(plusOneFloat64Uint16PtrErr, []*float64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapFloat64Uint16PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat64Uint16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat64Uint16PtrErr failed")
	}

	r, _ = PMapFloat64Uint16PtrErr(nil, []*float64{})
	if len(r) > 0 {
		t.Errorf("PMapFloat64Uint16PtrErr failed")
	}
	_, err := PMapFloat64Uint16PtrErr(plusOneFloat64Uint16PtrErr, []*float64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapFloat64Uint16PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapFloat64Uint16PtrErr(plusOneFloat64Uint16PtrErr, []*float64{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64Uint16PtrErr failed")
	}

	_, err = PMapFloat64Uint16PtrErr(plusOneFloat64Uint16PtrErr, []*float64{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapFloat64Uint16PtrErr failed")
	}

	_, err = PMapFloat64Uint16PtrErr(plusOneFloat64Uint16PtrErr, []*float64{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64Uint16PtrErr failed")
	}
	_, err = PMapFloat64Uint16PtrErr(plusOneFloat64Uint16PtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64Uint16PtrErr failed")
	}

	expectedList = []*uint16{&vo2, &vo3}
	newList, _ = PMapFloat64Uint16PtrErr(plusOneFloat64Uint16PtrErr, []*float64{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapFloat64Uint16Ptr failed")
	}

	_, err = PMapFloat64Uint16PtrErr(plusOneFloat64Uint16PtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Uint16PtrErr failed")
	}
	_, err = PMapFloat64Uint16PtrErr(plusOneFloat64Uint16PtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Uint16PtrErr failed")
	}

	_, err = PMapFloat64Uint16PtrErr(plusOneFloat64Uint16PtrErr, []*float64{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Uint16PtrErr failed")
	}

	_, err = PMapFloat64Uint16PtrErr(plusOneFloat64Uint16PtrErr, []*float64{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Uint16PtrErr failed")
	}

	_, err = PMapFloat64Uint16PtrErr(plusOneFloat64Uint16PtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Uint16PtrErr failed")
	}
}

func TestPmapFloat64Uint8PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3

	expectedList := []*uint8{&vo2, &vo3}
	newList, _ := PMapFloat64Uint8PtrErr(plusOneFloat64Uint8PtrErr, []*float64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapFloat64Uint8PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat64Uint8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat64Uint8PtrErr failed")
	}

	r, _ = PMapFloat64Uint8PtrErr(nil, []*float64{})
	if len(r) > 0 {
		t.Errorf("PMapFloat64Uint8PtrErr failed")
	}
	_, err := PMapFloat64Uint8PtrErr(plusOneFloat64Uint8PtrErr, []*float64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapFloat64Uint8PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapFloat64Uint8PtrErr(plusOneFloat64Uint8PtrErr, []*float64{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64Uint8PtrErr failed")
	}

	_, err = PMapFloat64Uint8PtrErr(plusOneFloat64Uint8PtrErr, []*float64{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapFloat64Uint8PtrErr failed")
	}

	_, err = PMapFloat64Uint8PtrErr(plusOneFloat64Uint8PtrErr, []*float64{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64Uint8PtrErr failed")
	}
	_, err = PMapFloat64Uint8PtrErr(plusOneFloat64Uint8PtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64Uint8PtrErr failed")
	}

	expectedList = []*uint8{&vo2, &vo3}
	newList, _ = PMapFloat64Uint8PtrErr(plusOneFloat64Uint8PtrErr, []*float64{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapFloat64Uint8Ptr failed")
	}

	_, err = PMapFloat64Uint8PtrErr(plusOneFloat64Uint8PtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Uint8PtrErr failed")
	}
	_, err = PMapFloat64Uint8PtrErr(plusOneFloat64Uint8PtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Uint8PtrErr failed")
	}

	_, err = PMapFloat64Uint8PtrErr(plusOneFloat64Uint8PtrErr, []*float64{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Uint8PtrErr failed")
	}

	_, err = PMapFloat64Uint8PtrErr(plusOneFloat64Uint8PtrErr, []*float64{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Uint8PtrErr failed")
	}

	_, err = PMapFloat64Uint8PtrErr(plusOneFloat64Uint8PtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Uint8PtrErr failed")
	}
}

func TestPmapFloat64StrPtrErr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 float64 = 10
	var vi0 float64

	expectedList := []*string{&vo10}
	newList, _ := PMapFloat64StrPtrErr(someLogicFloat64StrPtrErr, []*float64{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapFloat64StrPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat64StrPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat64StrPtrErr failed")
	}

	r, _ = PMapFloat64StrPtrErr(nil, []*float64{})
	if len(r) > 0 {
		t.Errorf("PMapFloat64StrPtrErr failed")
	}
	_, err := PMapFloat64StrPtrErr(someLogicFloat64StrPtrErr, []*float64{&vi10, &vi0})
	if err == nil {
		t.Errorf("PMapFloat64StrPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapFloat64StrPtrErr(someLogicFloat64StrPtrErr, []*float64{&vi10, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64StrPtrErr failed")
	}

	_, err = PMapFloat64StrPtrErr(someLogicFloat64StrPtrErr, []*float64{&vi0, &vi0, &vi10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapFloat64StrPtrErr failed")
	}

	_, err = PMapFloat64StrPtrErr(someLogicFloat64StrPtrErr, []*float64{&vi0, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64StrPtrErr failed")
	}
	_, err = PMapFloat64StrPtrErr(someLogicFloat64StrPtrErr, []*float64{&vi10, &vi10, &vi0}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64StrPtrErr failed")
	}

	expectedList = []*string{&vo10, &vo10}
	newList, _ = PMapFloat64StrPtrErr(someLogicFloat64StrPtrErr, []*float64{&vi10, &vi10}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapFloat64StrPtr failed")
	}

	_, err = PMapFloat64StrPtrErr(someLogicFloat64StrPtrErr, []*float64{&vi10, &vi10, &vi0}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64StrPtrErr failed")
	}
	_, err = PMapFloat64StrPtrErr(someLogicFloat64StrPtrErr, []*float64{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64StrPtrErr failed")
	}

	_, err = PMapFloat64StrPtrErr(someLogicFloat64StrPtrErr, []*float64{&vi0, &vi0, &vi10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64StrPtrErr failed")
	}

	_, err = PMapFloat64StrPtrErr(someLogicFloat64StrPtrErr, []*float64{&vi0, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64StrPtrErr failed")
	}

	_, err = PMapFloat64StrPtrErr(someLogicFloat64StrPtrErr, []*float64{&vi10, &vi10, &vi0}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64StrPtrErr failed")
	}
}

func TestPmapFloat64BoolPtrErr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 float64 = 10
	var v0 float64
	var v3 float64 = 3

	expectedList := []*bool{&vt, &vf}
	newList, _ := PMapFloat64BoolPtrErr(someLogicFloat64BoolPtrErr, []*float64{&v10, &v0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapFloat64BoolPtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat64BoolPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat64BoolPtrErr failed")
	}

	r, _ = PMapFloat64BoolPtrErr(nil, []*float64{})
	if len(r) > 0 {
		t.Errorf("PMapFloat64BoolPtrErr failed")
	}
	_, err := PMapFloat64BoolPtrErr(someLogicFloat64BoolPtrErr, []*float64{&v10, &v3})
	if err == nil {
		t.Errorf("PMapFloat64BoolPtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapFloat64BoolPtrErr(someLogicFloat64BoolPtrErr, []*float64{&v10, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64BoolPtrErr failed")
	}

	_, err = PMapFloat64BoolPtrErr(someLogicFloat64BoolPtrErr, []*float64{&v3, &v3, &v10}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapFloat64BoolPtrErr failed")
	}

	_, err = PMapFloat64BoolPtrErr(someLogicFloat64BoolPtrErr, []*float64{&v3, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64BoolErr failed")
	}
	_, err = PMapFloat64BoolPtrErr(someLogicFloat64BoolPtrErr, []*float64{&v10, &v10, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64BoolPtrErr failed")
	}

	expectedList = []*bool{&vt, &vt}
	newList, _ = PMapFloat64BoolPtrErr(someLogicFloat64BoolPtrErr, []*float64{&v10, &v10}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapFloat64BoolPtr failed")
	}

	_, err = PMapFloat64BoolPtrErr(someLogicFloat64BoolPtrErr, []*float64{&v10, &v10, &v3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64BoolErr failed")
	}
	_, err = PMapFloat64BoolPtrErr(someLogicFloat64BoolPtrErr, []*float64{&v10, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64BoolPtrErr failed")
	}

	_, err = PMapFloat64BoolPtrErr(someLogicFloat64BoolPtrErr, []*float64{&v3, &v3, &v10}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64BoolPtrErr failed")
	}

	_, err = PMapFloat64BoolPtrErr(someLogicFloat64BoolPtrErr, []*float64{&v3, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64BoolPtrErr failed")
	}

	_, err = PMapFloat64BoolPtrErr(someLogicFloat64BoolPtrErr, []*float64{&v10, &v10, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64BoolPtrErr failed")
	}
}

func TestPmapFloat64Float32PtrErr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3

	expectedList := []*float32{&vo2, &vo3}
	newList, _ := PMapFloat64Float32PtrErr(plusOneFloat64Float32PtrErr, []*float64{&vi1, &vi2})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapFloat64Float32PtrErr failed. expected=%v, actual=%v", expectedList, newList)
	}

	r, _ := PMapFloat64Float32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat64Float32PtrErr failed")
	}

	r, _ = PMapFloat64Float32PtrErr(nil, []*float64{})
	if len(r) > 0 {
		t.Errorf("PMapFloat64Float32PtrErr failed")
	}
	_, err := PMapFloat64Float32PtrErr(plusOneFloat64Float32PtrErr, []*float64{&vi1, &vi2, &vi3})
	if err == nil {
		t.Errorf("PMapFloat64Float32PtrErr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	_, err = PMapFloat64Float32PtrErr(plusOneFloat64Float32PtrErr, []*float64{&vi3, &vi3, &vi2}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64Float32PtrErr failed")
	}

	_, err = PMapFloat64Float32PtrErr(plusOneFloat64Float32PtrErr, []*float64{&vi3, &vi3, &vi2}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("PMapFloat64Float32PtrErr failed")
	}

	_, err = PMapFloat64Float32PtrErr(plusOneFloat64Float32PtrErr, []*float64{&vi3, &vi1, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64Float32PtrErr failed")
	}
	_, err = PMapFloat64Float32PtrErr(plusOneFloat64Float32PtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64Float32PtrErr failed")
	}

	expectedList = []*float32{&vo2, &vo3}
	newList, _ = PMapFloat64Float32PtrErr(plusOneFloat64Float32PtrErr, []*float64{&vi1, &vi2}, Optional{RandomOrder: true})

	counter := 0
	for i := 0; i < len(expectedList); i++ {
		for j := 0; j < len(newList); j++ {
			if *expectedList[i] == *newList[j] {
				counter++
				break
			}
		}
	}
	if counter != len(expectedList) {
		t.Errorf("PMapFloat64Float32Ptr failed")
	}

	_, err = PMapFloat64Float32PtrErr(plusOneFloat64Float32PtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Float32PtrErr failed")
	}
	_, err = PMapFloat64Float32PtrErr(plusOneFloat64Float32PtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Float32PtrErr failed")
	}

	_, err = PMapFloat64Float32PtrErr(plusOneFloat64Float32PtrErr, []*float64{&vi3, &vi3, &vi1}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Float32PtrErr failed")
	}

	_, err = PMapFloat64Float32PtrErr(plusOneFloat64Float32PtrErr, []*float64{&vi3, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Float32PtrErr failed")
	}

	_, err = PMapFloat64Float32PtrErr(plusOneFloat64Float32PtrErr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapFloat64Float32PtrErr failed")
	}
}
