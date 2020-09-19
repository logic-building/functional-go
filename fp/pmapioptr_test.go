package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestPmapIntInt64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3
	var vo4 int64 = 4

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3
	expectedList := []*int64{&vo2, &vo3, &vo4}
	newList := PMapIntInt64Ptr(plusOneIntInt64Ptr, []*int{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapIntInt64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapIntInt64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapIntInt64Ptr failed")
	}

	if len(PMapIntInt64Ptr(nil, []*int{})) > 0 {
		t.Errorf("PMapIntInt64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int64{&vo2, &vo3, &vo4}
	newList = PMapIntInt64Ptr(plusOneIntInt64Ptr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapIntInt64Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapIntInt32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3
	var vo4 int32 = 4

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3
	expectedList := []*int32{&vo2, &vo3, &vo4}
	newList := PMapIntInt32Ptr(plusOneIntInt32Ptr, []*int{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapIntInt32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapIntInt32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapIntInt32Ptr failed")
	}

	if len(PMapIntInt32Ptr(nil, []*int{})) > 0 {
		t.Errorf("PMapIntInt32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int32{&vo2, &vo3, &vo4}
	newList = PMapIntInt32Ptr(plusOneIntInt32Ptr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapIntInt32Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapIntInt16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3
	var vo4 int16 = 4

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3
	expectedList := []*int16{&vo2, &vo3, &vo4}
	newList := PMapIntInt16Ptr(plusOneIntInt16Ptr, []*int{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapIntInt16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapIntInt16Ptr(nil, nil)) > 0 {
		t.Errorf("PMapIntInt16Ptr failed")
	}

	if len(PMapIntInt16Ptr(nil, []*int{})) > 0 {
		t.Errorf("PMapIntInt16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int16{&vo2, &vo3, &vo4}
	newList = PMapIntInt16Ptr(plusOneIntInt16Ptr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapIntInt16Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapIntInt8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3
	var vo4 int8 = 4

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3
	expectedList := []*int8{&vo2, &vo3, &vo4}
	newList := PMapIntInt8Ptr(plusOneIntInt8Ptr, []*int{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapIntInt8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapIntInt8Ptr(nil, nil)) > 0 {
		t.Errorf("PMapIntInt8Ptr failed")
	}

	if len(PMapIntInt8Ptr(nil, []*int{})) > 0 {
		t.Errorf("PMapIntInt8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int8{&vo2, &vo3, &vo4}
	newList = PMapIntInt8Ptr(plusOneIntInt8Ptr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapIntInt8Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapIntUintPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3
	var vo4 uint = 4

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3
	expectedList := []*uint{&vo2, &vo3, &vo4}
	newList := PMapIntUintPtr(plusOneIntUintPtr, []*int{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapIntUintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapIntUintPtr(nil, nil)) > 0 {
		t.Errorf("PMapIntUintPtr failed")
	}

	if len(PMapIntUintPtr(nil, []*int{})) > 0 {
		t.Errorf("PMapIntUintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint{&vo2, &vo3, &vo4}
	newList = PMapIntUintPtr(plusOneIntUintPtr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapIntUintPtr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapIntUint64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3
	var vo4 uint64 = 4

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3
	expectedList := []*uint64{&vo2, &vo3, &vo4}
	newList := PMapIntUint64Ptr(plusOneIntUint64Ptr, []*int{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapIntUint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapIntUint64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapIntUint64Ptr failed")
	}

	if len(PMapIntUint64Ptr(nil, []*int{})) > 0 {
		t.Errorf("PMapIntUint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint64{&vo2, &vo3, &vo4}
	newList = PMapIntUint64Ptr(plusOneIntUint64Ptr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapIntUint64Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapIntUint32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3
	var vo4 uint32 = 4

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3
	expectedList := []*uint32{&vo2, &vo3, &vo4}
	newList := PMapIntUint32Ptr(plusOneIntUint32Ptr, []*int{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapIntUint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapIntUint32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapIntUint32Ptr failed")
	}

	if len(PMapIntUint32Ptr(nil, []*int{})) > 0 {
		t.Errorf("PMapIntUint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint32{&vo2, &vo3, &vo4}
	newList = PMapIntUint32Ptr(plusOneIntUint32Ptr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapIntUint32Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapIntUint16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3
	var vo4 uint16 = 4

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3
	expectedList := []*uint16{&vo2, &vo3, &vo4}
	newList := PMapIntUint16Ptr(plusOneIntUint16Ptr, []*int{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapIntUint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapIntUint16Ptr(nil, nil)) > 0 {
		t.Errorf("PMapIntUint16Ptr failed")
	}

	if len(PMapIntUint16Ptr(nil, []*int{})) > 0 {
		t.Errorf("PMapIntUint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint16{&vo2, &vo3, &vo4}
	newList = PMapIntUint16Ptr(plusOneIntUint16Ptr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapIntUint16Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapIntUint8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3
	var vo4 uint8 = 4

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3
	expectedList := []*uint8{&vo2, &vo3, &vo4}
	newList := PMapIntUint8Ptr(plusOneIntUint8Ptr, []*int{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapIntUint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapIntUint8Ptr(nil, nil)) > 0 {
		t.Errorf("PMapIntUint8Ptr failed")
	}

	if len(PMapIntUint8Ptr(nil, []*int{})) > 0 {
		t.Errorf("PMapIntUint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint8{&vo2, &vo3, &vo4}
	newList = PMapIntUint8Ptr(plusOneIntUint8Ptr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapIntUint8Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapIntStrPtr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 int = 10
	var vi20 int = 20

	expectedList := []*string{&vo10}
	newList := PMapIntStrPtr(someLogicIntStrPtr, []*int{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapIntStrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapIntStrPtr(nil, nil)) > 0 {
		t.Errorf("PMapIntStrPtr failed")
	}

	if len(PMapIntStrPtr(nil, []*int{})) > 0 {
		t.Errorf("PMapIntStrPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*string{&vo10}
	newList = PMapIntStrPtr(someLogicIntStrPtr, []*int{&vi10, &vi20}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != "0" {
		t.Errorf("PMapIntStrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapIntBoolPtr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 int = 10
	var v0 int = 0

	expectedList := []*bool{&vt, &vf}
	newList := PMapIntBoolPtr(someLogicIntBoolPtr, []*int{&v10, &v0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapIntBoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapIntBoolPtr(nil, nil)) > 0 {
		t.Errorf("PMapIntBoolPtr failed")
	}

	if len(PMapIntBoolPtr(nil, []*int{})) > 0 {
		t.Errorf("PMapIntBoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*bool{&vt, &vf}
	newList = PMapIntBoolPtr(someLogicIntBoolPtr, []*int{&v10, &v0}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != false {
		t.Errorf("PMapIntBoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapIntFloat32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3
	var vo4 float32 = 4

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3
	expectedList := []*float32{&vo2, &vo3, &vo4}
	newList := PMapIntFloat32Ptr(plusOneIntFloat32Ptr, []*int{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapIntFloat32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapIntFloat32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapIntFloat32Ptr failed")
	}

	if len(PMapIntFloat32Ptr(nil, []*int{})) > 0 {
		t.Errorf("PMapIntFloat32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*float32{&vo2, &vo3, &vo4}
	newList = PMapIntFloat32Ptr(plusOneIntFloat32Ptr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapIntFloat32Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapIntFloat64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3
	var vo4 float64 = 4

	var vi1 int = 1
	var vi2 int = 2
	var vi3 int = 3
	expectedList := []*float64{&vo2, &vo3, &vo4}
	newList := PMapIntFloat64Ptr(plusOneIntFloat64Ptr, []*int{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapIntFloat64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapIntFloat64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapIntFloat64Ptr failed")
	}

	if len(PMapIntFloat64Ptr(nil, []*int{})) > 0 {
		t.Errorf("PMapIntFloat64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*float64{&vo2, &vo3, &vo4}
	newList = PMapIntFloat64Ptr(plusOneIntFloat64Ptr, []*int{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapIntFloat64Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt64IntPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3
	var vo4 int = 4

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3
	expectedList := []*int{&vo2, &vo3, &vo4}
	newList := PMapInt64IntPtr(plusOneInt64IntPtr, []*int64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt64IntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt64IntPtr(nil, nil)) > 0 {
		t.Errorf("PMapInt64IntPtr failed")
	}

	if len(PMapInt64IntPtr(nil, []*int64{})) > 0 {
		t.Errorf("PMapInt64IntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int{&vo2, &vo3, &vo4}
	newList = PMapInt64IntPtr(plusOneInt64IntPtr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt64IntPtr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt64Int32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3
	var vo4 int32 = 4

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3
	expectedList := []*int32{&vo2, &vo3, &vo4}
	newList := PMapInt64Int32Ptr(plusOneInt64Int32Ptr, []*int64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt64Int32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt64Int32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt64Int32Ptr failed")
	}

	if len(PMapInt64Int32Ptr(nil, []*int64{})) > 0 {
		t.Errorf("PMapInt64Int32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int32{&vo2, &vo3, &vo4}
	newList = PMapInt64Int32Ptr(plusOneInt64Int32Ptr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt64Int32Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt64Int16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3
	var vo4 int16 = 4

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3
	expectedList := []*int16{&vo2, &vo3, &vo4}
	newList := PMapInt64Int16Ptr(plusOneInt64Int16Ptr, []*int64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt64Int16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt64Int16Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt64Int16Ptr failed")
	}

	if len(PMapInt64Int16Ptr(nil, []*int64{})) > 0 {
		t.Errorf("PMapInt64Int16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int16{&vo2, &vo3, &vo4}
	newList = PMapInt64Int16Ptr(plusOneInt64Int16Ptr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt64Int16Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt64Int8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3
	var vo4 int8 = 4

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3
	expectedList := []*int8{&vo2, &vo3, &vo4}
	newList := PMapInt64Int8Ptr(plusOneInt64Int8Ptr, []*int64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt64Int8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt64Int8Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt64Int8Ptr failed")
	}

	if len(PMapInt64Int8Ptr(nil, []*int64{})) > 0 {
		t.Errorf("PMapInt64Int8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int8{&vo2, &vo3, &vo4}
	newList = PMapInt64Int8Ptr(plusOneInt64Int8Ptr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt64Int8Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt64UintPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3
	var vo4 uint = 4

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3
	expectedList := []*uint{&vo2, &vo3, &vo4}
	newList := PMapInt64UintPtr(plusOneInt64UintPtr, []*int64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt64UintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt64UintPtr(nil, nil)) > 0 {
		t.Errorf("PMapInt64UintPtr failed")
	}

	if len(PMapInt64UintPtr(nil, []*int64{})) > 0 {
		t.Errorf("PMapInt64UintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint{&vo2, &vo3, &vo4}
	newList = PMapInt64UintPtr(plusOneInt64UintPtr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt64UintPtr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt64Uint64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3
	var vo4 uint64 = 4

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3
	expectedList := []*uint64{&vo2, &vo3, &vo4}
	newList := PMapInt64Uint64Ptr(plusOneInt64Uint64Ptr, []*int64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt64Uint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt64Uint64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt64Uint64Ptr failed")
	}

	if len(PMapInt64Uint64Ptr(nil, []*int64{})) > 0 {
		t.Errorf("PMapInt64Uint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint64{&vo2, &vo3, &vo4}
	newList = PMapInt64Uint64Ptr(plusOneInt64Uint64Ptr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt64Uint64Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt64Uint32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3
	var vo4 uint32 = 4

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3
	expectedList := []*uint32{&vo2, &vo3, &vo4}
	newList := PMapInt64Uint32Ptr(plusOneInt64Uint32Ptr, []*int64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt64Uint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt64Uint32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt64Uint32Ptr failed")
	}

	if len(PMapInt64Uint32Ptr(nil, []*int64{})) > 0 {
		t.Errorf("PMapInt64Uint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint32{&vo2, &vo3, &vo4}
	newList = PMapInt64Uint32Ptr(plusOneInt64Uint32Ptr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt64Uint32Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt64Uint16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3
	var vo4 uint16 = 4

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3
	expectedList := []*uint16{&vo2, &vo3, &vo4}
	newList := PMapInt64Uint16Ptr(plusOneInt64Uint16Ptr, []*int64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt64Uint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt64Uint16Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt64Uint16Ptr failed")
	}

	if len(PMapInt64Uint16Ptr(nil, []*int64{})) > 0 {
		t.Errorf("PMapInt64Uint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint16{&vo2, &vo3, &vo4}
	newList = PMapInt64Uint16Ptr(plusOneInt64Uint16Ptr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt64Uint16Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt64Uint8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3
	var vo4 uint8 = 4

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3
	expectedList := []*uint8{&vo2, &vo3, &vo4}
	newList := PMapInt64Uint8Ptr(plusOneInt64Uint8Ptr, []*int64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt64Uint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt64Uint8Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt64Uint8Ptr failed")
	}

	if len(PMapInt64Uint8Ptr(nil, []*int64{})) > 0 {
		t.Errorf("PMapInt64Uint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint8{&vo2, &vo3, &vo4}
	newList = PMapInt64Uint8Ptr(plusOneInt64Uint8Ptr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt64Uint8Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt64StrPtr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 int64 = 10
	var vi20 int64 = 20

	expectedList := []*string{&vo10}
	newList := PMapInt64StrPtr(someLogicInt64StrPtr, []*int64{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapInt64StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt64StrPtr(nil, nil)) > 0 {
		t.Errorf("PMapInt64StrPtr failed")
	}

	if len(PMapInt64StrPtr(nil, []*int64{})) > 0 {
		t.Errorf("PMapInt64StrPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*string{&vo10}
	newList = PMapInt64StrPtr(someLogicInt64StrPtr, []*int64{&vi10, &vi20}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != "0" {
		t.Errorf("PMapInt64StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapInt64BoolPtr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 int64 = 10
	var v0 int64 = 0

	expectedList := []*bool{&vt, &vf}
	newList := PMapInt64BoolPtr(someLogicInt64BoolPtr, []*int64{&v10, &v0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt64BoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt64BoolPtr(nil, nil)) > 0 {
		t.Errorf("PMapInt64BoolPtr failed")
	}

	if len(PMapInt64BoolPtr(nil, []*int64{})) > 0 {
		t.Errorf("PMapInt64BoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*bool{&vt, &vf}
	newList = PMapInt64BoolPtr(someLogicInt64BoolPtr, []*int64{&v10, &v0}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != false {
		t.Errorf("PMapInt64BoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapInt64Float32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3
	var vo4 float32 = 4

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3
	expectedList := []*float32{&vo2, &vo3, &vo4}
	newList := PMapInt64Float32Ptr(plusOneInt64Float32Ptr, []*int64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt64Float32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt64Float32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt64Float32Ptr failed")
	}

	if len(PMapInt64Float32Ptr(nil, []*int64{})) > 0 {
		t.Errorf("PMapInt64Float32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*float32{&vo2, &vo3, &vo4}
	newList = PMapInt64Float32Ptr(plusOneInt64Float32Ptr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt64Float32Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt64Float64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3
	var vo4 float64 = 4

	var vi1 int64 = 1
	var vi2 int64 = 2
	var vi3 int64 = 3
	expectedList := []*float64{&vo2, &vo3, &vo4}
	newList := PMapInt64Float64Ptr(plusOneInt64Float64Ptr, []*int64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt64Float64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt64Float64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt64Float64Ptr failed")
	}

	if len(PMapInt64Float64Ptr(nil, []*int64{})) > 0 {
		t.Errorf("PMapInt64Float64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*float64{&vo2, &vo3, &vo4}
	newList = PMapInt64Float64Ptr(plusOneInt64Float64Ptr, []*int64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt64Float64Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt32IntPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3
	var vo4 int = 4

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3
	expectedList := []*int{&vo2, &vo3, &vo4}
	newList := PMapInt32IntPtr(plusOneInt32IntPtr, []*int32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt32IntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt32IntPtr(nil, nil)) > 0 {
		t.Errorf("PMapInt32IntPtr failed")
	}

	if len(PMapInt32IntPtr(nil, []*int32{})) > 0 {
		t.Errorf("PMapInt32IntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int{&vo2, &vo3, &vo4}
	newList = PMapInt32IntPtr(plusOneInt32IntPtr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt32IntPtr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt32Int64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3
	var vo4 int64 = 4

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3
	expectedList := []*int64{&vo2, &vo3, &vo4}
	newList := PMapInt32Int64Ptr(plusOneInt32Int64Ptr, []*int32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt32Int64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt32Int64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt32Int64Ptr failed")
	}

	if len(PMapInt32Int64Ptr(nil, []*int32{})) > 0 {
		t.Errorf("PMapInt32Int64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int64{&vo2, &vo3, &vo4}
	newList = PMapInt32Int64Ptr(plusOneInt32Int64Ptr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt32Int64Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt32Int16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3
	var vo4 int16 = 4

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3
	expectedList := []*int16{&vo2, &vo3, &vo4}
	newList := PMapInt32Int16Ptr(plusOneInt32Int16Ptr, []*int32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt32Int16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt32Int16Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt32Int16Ptr failed")
	}

	if len(PMapInt32Int16Ptr(nil, []*int32{})) > 0 {
		t.Errorf("PMapInt32Int16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int16{&vo2, &vo3, &vo4}
	newList = PMapInt32Int16Ptr(plusOneInt32Int16Ptr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt32Int16Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt32Int8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3
	var vo4 int8 = 4

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3
	expectedList := []*int8{&vo2, &vo3, &vo4}
	newList := PMapInt32Int8Ptr(plusOneInt32Int8Ptr, []*int32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt32Int8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt32Int8Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt32Int8Ptr failed")
	}

	if len(PMapInt32Int8Ptr(nil, []*int32{})) > 0 {
		t.Errorf("PMapInt32Int8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int8{&vo2, &vo3, &vo4}
	newList = PMapInt32Int8Ptr(plusOneInt32Int8Ptr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt32Int8Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt32UintPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3
	var vo4 uint = 4

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3
	expectedList := []*uint{&vo2, &vo3, &vo4}
	newList := PMapInt32UintPtr(plusOneInt32UintPtr, []*int32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt32UintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt32UintPtr(nil, nil)) > 0 {
		t.Errorf("PMapInt32UintPtr failed")
	}

	if len(PMapInt32UintPtr(nil, []*int32{})) > 0 {
		t.Errorf("PMapInt32UintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint{&vo2, &vo3, &vo4}
	newList = PMapInt32UintPtr(plusOneInt32UintPtr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt32UintPtr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt32Uint64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3
	var vo4 uint64 = 4

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3
	expectedList := []*uint64{&vo2, &vo3, &vo4}
	newList := PMapInt32Uint64Ptr(plusOneInt32Uint64Ptr, []*int32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt32Uint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt32Uint64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt32Uint64Ptr failed")
	}

	if len(PMapInt32Uint64Ptr(nil, []*int32{})) > 0 {
		t.Errorf("PMapInt32Uint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint64{&vo2, &vo3, &vo4}
	newList = PMapInt32Uint64Ptr(plusOneInt32Uint64Ptr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt32Uint64Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt32Uint32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3
	var vo4 uint32 = 4

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3
	expectedList := []*uint32{&vo2, &vo3, &vo4}
	newList := PMapInt32Uint32Ptr(plusOneInt32Uint32Ptr, []*int32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt32Uint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt32Uint32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt32Uint32Ptr failed")
	}

	if len(PMapInt32Uint32Ptr(nil, []*int32{})) > 0 {
		t.Errorf("PMapInt32Uint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint32{&vo2, &vo3, &vo4}
	newList = PMapInt32Uint32Ptr(plusOneInt32Uint32Ptr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt32Uint32Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt32Uint16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3
	var vo4 uint16 = 4

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3
	expectedList := []*uint16{&vo2, &vo3, &vo4}
	newList := PMapInt32Uint16Ptr(plusOneInt32Uint16Ptr, []*int32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt32Uint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt32Uint16Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt32Uint16Ptr failed")
	}

	if len(PMapInt32Uint16Ptr(nil, []*int32{})) > 0 {
		t.Errorf("PMapInt32Uint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint16{&vo2, &vo3, &vo4}
	newList = PMapInt32Uint16Ptr(plusOneInt32Uint16Ptr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt32Uint16Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt32Uint8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3
	var vo4 uint8 = 4

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3
	expectedList := []*uint8{&vo2, &vo3, &vo4}
	newList := PMapInt32Uint8Ptr(plusOneInt32Uint8Ptr, []*int32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt32Uint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt32Uint8Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt32Uint8Ptr failed")
	}

	if len(PMapInt32Uint8Ptr(nil, []*int32{})) > 0 {
		t.Errorf("PMapInt32Uint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint8{&vo2, &vo3, &vo4}
	newList = PMapInt32Uint8Ptr(plusOneInt32Uint8Ptr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt32Uint8Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt32StrPtr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 int32 = 10
	var vi20 int32 = 20

	expectedList := []*string{&vo10}
	newList := PMapInt32StrPtr(someLogicInt32StrPtr, []*int32{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapInt32StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt32StrPtr(nil, nil)) > 0 {
		t.Errorf("PMapInt32StrPtr failed")
	}

	if len(PMapInt32StrPtr(nil, []*int32{})) > 0 {
		t.Errorf("PMapInt32StrPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*string{&vo10}
	newList = PMapInt32StrPtr(someLogicInt32StrPtr, []*int32{&vi10, &vi20}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != "0" {
		t.Errorf("PMapInt32StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapInt32BoolPtr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 int32 = 10
	var v0 int32 = 0

	expectedList := []*bool{&vt, &vf}
	newList := PMapInt32BoolPtr(someLogicInt32BoolPtr, []*int32{&v10, &v0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt32BoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt32BoolPtr(nil, nil)) > 0 {
		t.Errorf("PMapInt32BoolPtr failed")
	}

	if len(PMapInt32BoolPtr(nil, []*int32{})) > 0 {
		t.Errorf("PMapInt32BoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*bool{&vt, &vf}
	newList = PMapInt32BoolPtr(someLogicInt32BoolPtr, []*int32{&v10, &v0}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != false {
		t.Errorf("PMapInt32BoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapInt32Float32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3
	var vo4 float32 = 4

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3
	expectedList := []*float32{&vo2, &vo3, &vo4}
	newList := PMapInt32Float32Ptr(plusOneInt32Float32Ptr, []*int32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt32Float32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt32Float32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt32Float32Ptr failed")
	}

	if len(PMapInt32Float32Ptr(nil, []*int32{})) > 0 {
		t.Errorf("PMapInt32Float32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*float32{&vo2, &vo3, &vo4}
	newList = PMapInt32Float32Ptr(plusOneInt32Float32Ptr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt32Float32Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt32Float64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3
	var vo4 float64 = 4

	var vi1 int32 = 1
	var vi2 int32 = 2
	var vi3 int32 = 3
	expectedList := []*float64{&vo2, &vo3, &vo4}
	newList := PMapInt32Float64Ptr(plusOneInt32Float64Ptr, []*int32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt32Float64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt32Float64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt32Float64Ptr failed")
	}

	if len(PMapInt32Float64Ptr(nil, []*int32{})) > 0 {
		t.Errorf("PMapInt32Float64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*float64{&vo2, &vo3, &vo4}
	newList = PMapInt32Float64Ptr(plusOneInt32Float64Ptr, []*int32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt32Float64Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt16IntPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3
	var vo4 int = 4

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3
	expectedList := []*int{&vo2, &vo3, &vo4}
	newList := PMapInt16IntPtr(plusOneInt16IntPtr, []*int16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt16IntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt16IntPtr(nil, nil)) > 0 {
		t.Errorf("PMapInt16IntPtr failed")
	}

	if len(PMapInt16IntPtr(nil, []*int16{})) > 0 {
		t.Errorf("PMapInt16IntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int{&vo2, &vo3, &vo4}
	newList = PMapInt16IntPtr(plusOneInt16IntPtr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt16IntPtr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt16Int64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3
	var vo4 int64 = 4

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3
	expectedList := []*int64{&vo2, &vo3, &vo4}
	newList := PMapInt16Int64Ptr(plusOneInt16Int64Ptr, []*int16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt16Int64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt16Int64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt16Int64Ptr failed")
	}

	if len(PMapInt16Int64Ptr(nil, []*int16{})) > 0 {
		t.Errorf("PMapInt16Int64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int64{&vo2, &vo3, &vo4}
	newList = PMapInt16Int64Ptr(plusOneInt16Int64Ptr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt16Int64Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt16Int32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3
	var vo4 int32 = 4

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3
	expectedList := []*int32{&vo2, &vo3, &vo4}
	newList := PMapInt16Int32Ptr(plusOneInt16Int32Ptr, []*int16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt16Int32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt16Int32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt16Int32Ptr failed")
	}

	if len(PMapInt16Int32Ptr(nil, []*int16{})) > 0 {
		t.Errorf("PMapInt16Int32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int32{&vo2, &vo3, &vo4}
	newList = PMapInt16Int32Ptr(plusOneInt16Int32Ptr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt16Int32Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt16Int8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3
	var vo4 int8 = 4

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3
	expectedList := []*int8{&vo2, &vo3, &vo4}
	newList := PMapInt16Int8Ptr(plusOneInt16Int8Ptr, []*int16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt16Int8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt16Int8Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt16Int8Ptr failed")
	}

	if len(PMapInt16Int8Ptr(nil, []*int16{})) > 0 {
		t.Errorf("PMapInt16Int8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int8{&vo2, &vo3, &vo4}
	newList = PMapInt16Int8Ptr(plusOneInt16Int8Ptr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt16Int8Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt16UintPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3
	var vo4 uint = 4

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3
	expectedList := []*uint{&vo2, &vo3, &vo4}
	newList := PMapInt16UintPtr(plusOneInt16UintPtr, []*int16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt16UintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt16UintPtr(nil, nil)) > 0 {
		t.Errorf("PMapInt16UintPtr failed")
	}

	if len(PMapInt16UintPtr(nil, []*int16{})) > 0 {
		t.Errorf("PMapInt16UintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint{&vo2, &vo3, &vo4}
	newList = PMapInt16UintPtr(plusOneInt16UintPtr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt16UintPtr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt16Uint64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3
	var vo4 uint64 = 4

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3
	expectedList := []*uint64{&vo2, &vo3, &vo4}
	newList := PMapInt16Uint64Ptr(plusOneInt16Uint64Ptr, []*int16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt16Uint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt16Uint64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt16Uint64Ptr failed")
	}

	if len(PMapInt16Uint64Ptr(nil, []*int16{})) > 0 {
		t.Errorf("PMapInt16Uint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint64{&vo2, &vo3, &vo4}
	newList = PMapInt16Uint64Ptr(plusOneInt16Uint64Ptr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt16Uint64Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt16Uint32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3
	var vo4 uint32 = 4

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3
	expectedList := []*uint32{&vo2, &vo3, &vo4}
	newList := PMapInt16Uint32Ptr(plusOneInt16Uint32Ptr, []*int16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt16Uint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt16Uint32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt16Uint32Ptr failed")
	}

	if len(PMapInt16Uint32Ptr(nil, []*int16{})) > 0 {
		t.Errorf("PMapInt16Uint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint32{&vo2, &vo3, &vo4}
	newList = PMapInt16Uint32Ptr(plusOneInt16Uint32Ptr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt16Uint32Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt16Uint16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3
	var vo4 uint16 = 4

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3
	expectedList := []*uint16{&vo2, &vo3, &vo4}
	newList := PMapInt16Uint16Ptr(plusOneInt16Uint16Ptr, []*int16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt16Uint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt16Uint16Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt16Uint16Ptr failed")
	}

	if len(PMapInt16Uint16Ptr(nil, []*int16{})) > 0 {
		t.Errorf("PMapInt16Uint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint16{&vo2, &vo3, &vo4}
	newList = PMapInt16Uint16Ptr(plusOneInt16Uint16Ptr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt16Uint16Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt16Uint8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3
	var vo4 uint8 = 4

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3
	expectedList := []*uint8{&vo2, &vo3, &vo4}
	newList := PMapInt16Uint8Ptr(plusOneInt16Uint8Ptr, []*int16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt16Uint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt16Uint8Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt16Uint8Ptr failed")
	}

	if len(PMapInt16Uint8Ptr(nil, []*int16{})) > 0 {
		t.Errorf("PMapInt16Uint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint8{&vo2, &vo3, &vo4}
	newList = PMapInt16Uint8Ptr(plusOneInt16Uint8Ptr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt16Uint8Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt16StrPtr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 int16 = 10
	var vi20 int16 = 20

	expectedList := []*string{&vo10}
	newList := PMapInt16StrPtr(someLogicInt16StrPtr, []*int16{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapInt16StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt16StrPtr(nil, nil)) > 0 {
		t.Errorf("PMapInt16StrPtr failed")
	}

	if len(PMapInt16StrPtr(nil, []*int16{})) > 0 {
		t.Errorf("PMapInt16StrPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*string{&vo10}
	newList = PMapInt16StrPtr(someLogicInt16StrPtr, []*int16{&vi10, &vi20}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != "0" {
		t.Errorf("PMapInt16StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapInt16BoolPtr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 int16 = 10
	var v0 int16 = 0

	expectedList := []*bool{&vt, &vf}
	newList := PMapInt16BoolPtr(someLogicInt16BoolPtr, []*int16{&v10, &v0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt16BoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt16BoolPtr(nil, nil)) > 0 {
		t.Errorf("PMapInt16BoolPtr failed")
	}

	if len(PMapInt16BoolPtr(nil, []*int16{})) > 0 {
		t.Errorf("PMapInt16BoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*bool{&vt, &vf}
	newList = PMapInt16BoolPtr(someLogicInt16BoolPtr, []*int16{&v10, &v0}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != false {
		t.Errorf("PMapInt16BoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapInt16Float32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3
	var vo4 float32 = 4

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3
	expectedList := []*float32{&vo2, &vo3, &vo4}
	newList := PMapInt16Float32Ptr(plusOneInt16Float32Ptr, []*int16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt16Float32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt16Float32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt16Float32Ptr failed")
	}

	if len(PMapInt16Float32Ptr(nil, []*int16{})) > 0 {
		t.Errorf("PMapInt16Float32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*float32{&vo2, &vo3, &vo4}
	newList = PMapInt16Float32Ptr(plusOneInt16Float32Ptr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt16Float32Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt16Float64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3
	var vo4 float64 = 4

	var vi1 int16 = 1
	var vi2 int16 = 2
	var vi3 int16 = 3
	expectedList := []*float64{&vo2, &vo3, &vo4}
	newList := PMapInt16Float64Ptr(plusOneInt16Float64Ptr, []*int16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt16Float64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt16Float64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt16Float64Ptr failed")
	}

	if len(PMapInt16Float64Ptr(nil, []*int16{})) > 0 {
		t.Errorf("PMapInt16Float64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*float64{&vo2, &vo3, &vo4}
	newList = PMapInt16Float64Ptr(plusOneInt16Float64Ptr, []*int16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt16Float64Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt8IntPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3
	var vo4 int = 4

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3
	expectedList := []*int{&vo2, &vo3, &vo4}
	newList := PMapInt8IntPtr(plusOneInt8IntPtr, []*int8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt8IntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt8IntPtr(nil, nil)) > 0 {
		t.Errorf("PMapInt8IntPtr failed")
	}

	if len(PMapInt8IntPtr(nil, []*int8{})) > 0 {
		t.Errorf("PMapInt8IntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int{&vo2, &vo3, &vo4}
	newList = PMapInt8IntPtr(plusOneInt8IntPtr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt8IntPtr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt8Int64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3
	var vo4 int64 = 4

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3
	expectedList := []*int64{&vo2, &vo3, &vo4}
	newList := PMapInt8Int64Ptr(plusOneInt8Int64Ptr, []*int8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt8Int64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt8Int64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt8Int64Ptr failed")
	}

	if len(PMapInt8Int64Ptr(nil, []*int8{})) > 0 {
		t.Errorf("PMapInt8Int64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int64{&vo2, &vo3, &vo4}
	newList = PMapInt8Int64Ptr(plusOneInt8Int64Ptr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt8Int64Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt8Int32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3
	var vo4 int32 = 4

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3
	expectedList := []*int32{&vo2, &vo3, &vo4}
	newList := PMapInt8Int32Ptr(plusOneInt8Int32Ptr, []*int8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt8Int32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt8Int32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt8Int32Ptr failed")
	}

	if len(PMapInt8Int32Ptr(nil, []*int8{})) > 0 {
		t.Errorf("PMapInt8Int32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int32{&vo2, &vo3, &vo4}
	newList = PMapInt8Int32Ptr(plusOneInt8Int32Ptr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt8Int32Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt8Int16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3
	var vo4 int16 = 4

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3
	expectedList := []*int16{&vo2, &vo3, &vo4}
	newList := PMapInt8Int16Ptr(plusOneInt8Int16Ptr, []*int8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt8Int16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt8Int16Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt8Int16Ptr failed")
	}

	if len(PMapInt8Int16Ptr(nil, []*int8{})) > 0 {
		t.Errorf("PMapInt8Int16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int16{&vo2, &vo3, &vo4}
	newList = PMapInt8Int16Ptr(plusOneInt8Int16Ptr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt8Int16Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt8UintPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3
	var vo4 uint = 4

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3
	expectedList := []*uint{&vo2, &vo3, &vo4}
	newList := PMapInt8UintPtr(plusOneInt8UintPtr, []*int8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt8UintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt8UintPtr(nil, nil)) > 0 {
		t.Errorf("PMapInt8UintPtr failed")
	}

	if len(PMapInt8UintPtr(nil, []*int8{})) > 0 {
		t.Errorf("PMapInt8UintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint{&vo2, &vo3, &vo4}
	newList = PMapInt8UintPtr(plusOneInt8UintPtr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt8UintPtr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt8Uint64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3
	var vo4 uint64 = 4

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3
	expectedList := []*uint64{&vo2, &vo3, &vo4}
	newList := PMapInt8Uint64Ptr(plusOneInt8Uint64Ptr, []*int8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt8Uint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt8Uint64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt8Uint64Ptr failed")
	}

	if len(PMapInt8Uint64Ptr(nil, []*int8{})) > 0 {
		t.Errorf("PMapInt8Uint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint64{&vo2, &vo3, &vo4}
	newList = PMapInt8Uint64Ptr(plusOneInt8Uint64Ptr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt8Uint64Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt8Uint32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3
	var vo4 uint32 = 4

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3
	expectedList := []*uint32{&vo2, &vo3, &vo4}
	newList := PMapInt8Uint32Ptr(plusOneInt8Uint32Ptr, []*int8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt8Uint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt8Uint32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt8Uint32Ptr failed")
	}

	if len(PMapInt8Uint32Ptr(nil, []*int8{})) > 0 {
		t.Errorf("PMapInt8Uint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint32{&vo2, &vo3, &vo4}
	newList = PMapInt8Uint32Ptr(plusOneInt8Uint32Ptr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt8Uint32Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt8Uint16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3
	var vo4 uint16 = 4

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3
	expectedList := []*uint16{&vo2, &vo3, &vo4}
	newList := PMapInt8Uint16Ptr(plusOneInt8Uint16Ptr, []*int8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt8Uint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt8Uint16Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt8Uint16Ptr failed")
	}

	if len(PMapInt8Uint16Ptr(nil, []*int8{})) > 0 {
		t.Errorf("PMapInt8Uint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint16{&vo2, &vo3, &vo4}
	newList = PMapInt8Uint16Ptr(plusOneInt8Uint16Ptr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt8Uint16Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt8Uint8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3
	var vo4 uint8 = 4

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3
	expectedList := []*uint8{&vo2, &vo3, &vo4}
	newList := PMapInt8Uint8Ptr(plusOneInt8Uint8Ptr, []*int8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt8Uint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt8Uint8Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt8Uint8Ptr failed")
	}

	if len(PMapInt8Uint8Ptr(nil, []*int8{})) > 0 {
		t.Errorf("PMapInt8Uint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint8{&vo2, &vo3, &vo4}
	newList = PMapInt8Uint8Ptr(plusOneInt8Uint8Ptr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt8Uint8Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt8StrPtr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 int8 = 10
	var vi20 int8 = 20

	expectedList := []*string{&vo10}
	newList := PMapInt8StrPtr(someLogicInt8StrPtr, []*int8{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapInt8StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt8StrPtr(nil, nil)) > 0 {
		t.Errorf("PMapInt8StrPtr failed")
	}

	if len(PMapInt8StrPtr(nil, []*int8{})) > 0 {
		t.Errorf("PMapInt8StrPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*string{&vo10}
	newList = PMapInt8StrPtr(someLogicInt8StrPtr, []*int8{&vi10, &vi20}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != "0" {
		t.Errorf("PMapInt8StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapInt8BoolPtr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 int8 = 10
	var v0 int8 = 0

	expectedList := []*bool{&vt, &vf}
	newList := PMapInt8BoolPtr(someLogicInt8BoolPtr, []*int8{&v10, &v0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapInt8BoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt8BoolPtr(nil, nil)) > 0 {
		t.Errorf("PMapInt8BoolPtr failed")
	}

	if len(PMapInt8BoolPtr(nil, []*int8{})) > 0 {
		t.Errorf("PMapInt8BoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*bool{&vt, &vf}
	newList = PMapInt8BoolPtr(someLogicInt8BoolPtr, []*int8{&v10, &v0}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != false {
		t.Errorf("PMapInt8BoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapInt8Float32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3
	var vo4 float32 = 4

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3
	expectedList := []*float32{&vo2, &vo3, &vo4}
	newList := PMapInt8Float32Ptr(plusOneInt8Float32Ptr, []*int8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt8Float32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt8Float32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt8Float32Ptr failed")
	}

	if len(PMapInt8Float32Ptr(nil, []*int8{})) > 0 {
		t.Errorf("PMapInt8Float32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*float32{&vo2, &vo3, &vo4}
	newList = PMapInt8Float32Ptr(plusOneInt8Float32Ptr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt8Float32Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapInt8Float64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3
	var vo4 float64 = 4

	var vi1 int8 = 1
	var vi2 int8 = 2
	var vi3 int8 = 3
	expectedList := []*float64{&vo2, &vo3, &vo4}
	newList := PMapInt8Float64Ptr(plusOneInt8Float64Ptr, []*int8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapInt8Float64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapInt8Float64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt8Float64Ptr failed")
	}

	if len(PMapInt8Float64Ptr(nil, []*int8{})) > 0 {
		t.Errorf("PMapInt8Float64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*float64{&vo2, &vo3, &vo4}
	newList = PMapInt8Float64Ptr(plusOneInt8Float64Ptr, []*int8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapInt8Float64Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUintIntPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3
	var vo4 int = 4

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3
	expectedList := []*int{&vo2, &vo3, &vo4}
	newList := PMapUintIntPtr(plusOneUintIntPtr, []*uint{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUintIntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUintIntPtr(nil, nil)) > 0 {
		t.Errorf("PMapUintIntPtr failed")
	}

	if len(PMapUintIntPtr(nil, []*uint{})) > 0 {
		t.Errorf("PMapUintIntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int{&vo2, &vo3, &vo4}
	newList = PMapUintIntPtr(plusOneUintIntPtr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUintIntPtr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUintInt64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3
	var vo4 int64 = 4

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3
	expectedList := []*int64{&vo2, &vo3, &vo4}
	newList := PMapUintInt64Ptr(plusOneUintInt64Ptr, []*uint{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUintInt64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUintInt64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUintInt64Ptr failed")
	}

	if len(PMapUintInt64Ptr(nil, []*uint{})) > 0 {
		t.Errorf("PMapUintInt64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int64{&vo2, &vo3, &vo4}
	newList = PMapUintInt64Ptr(plusOneUintInt64Ptr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUintInt64Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUintInt32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3
	var vo4 int32 = 4

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3
	expectedList := []*int32{&vo2, &vo3, &vo4}
	newList := PMapUintInt32Ptr(plusOneUintInt32Ptr, []*uint{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUintInt32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUintInt32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUintInt32Ptr failed")
	}

	if len(PMapUintInt32Ptr(nil, []*uint{})) > 0 {
		t.Errorf("PMapUintInt32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int32{&vo2, &vo3, &vo4}
	newList = PMapUintInt32Ptr(plusOneUintInt32Ptr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUintInt32Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUintInt16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3
	var vo4 int16 = 4

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3
	expectedList := []*int16{&vo2, &vo3, &vo4}
	newList := PMapUintInt16Ptr(plusOneUintInt16Ptr, []*uint{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUintInt16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUintInt16Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUintInt16Ptr failed")
	}

	if len(PMapUintInt16Ptr(nil, []*uint{})) > 0 {
		t.Errorf("PMapUintInt16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int16{&vo2, &vo3, &vo4}
	newList = PMapUintInt16Ptr(plusOneUintInt16Ptr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUintInt16Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUintInt8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3
	var vo4 int8 = 4

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3
	expectedList := []*int8{&vo2, &vo3, &vo4}
	newList := PMapUintInt8Ptr(plusOneUintInt8Ptr, []*uint{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUintInt8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUintInt8Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUintInt8Ptr failed")
	}

	if len(PMapUintInt8Ptr(nil, []*uint{})) > 0 {
		t.Errorf("PMapUintInt8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int8{&vo2, &vo3, &vo4}
	newList = PMapUintInt8Ptr(plusOneUintInt8Ptr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUintInt8Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUintUint64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3
	var vo4 uint64 = 4

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3
	expectedList := []*uint64{&vo2, &vo3, &vo4}
	newList := PMapUintUint64Ptr(plusOneUintUint64Ptr, []*uint{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUintUint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUintUint64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUintUint64Ptr failed")
	}

	if len(PMapUintUint64Ptr(nil, []*uint{})) > 0 {
		t.Errorf("PMapUintUint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint64{&vo2, &vo3, &vo4}
	newList = PMapUintUint64Ptr(plusOneUintUint64Ptr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUintUint64Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUintUint32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3
	var vo4 uint32 = 4

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3
	expectedList := []*uint32{&vo2, &vo3, &vo4}
	newList := PMapUintUint32Ptr(plusOneUintUint32Ptr, []*uint{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUintUint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUintUint32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUintUint32Ptr failed")
	}

	if len(PMapUintUint32Ptr(nil, []*uint{})) > 0 {
		t.Errorf("PMapUintUint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint32{&vo2, &vo3, &vo4}
	newList = PMapUintUint32Ptr(plusOneUintUint32Ptr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUintUint32Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUintUint16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3
	var vo4 uint16 = 4

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3
	expectedList := []*uint16{&vo2, &vo3, &vo4}
	newList := PMapUintUint16Ptr(plusOneUintUint16Ptr, []*uint{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUintUint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUintUint16Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUintUint16Ptr failed")
	}

	if len(PMapUintUint16Ptr(nil, []*uint{})) > 0 {
		t.Errorf("PMapUintUint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint16{&vo2, &vo3, &vo4}
	newList = PMapUintUint16Ptr(plusOneUintUint16Ptr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUintUint16Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUintUint8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3
	var vo4 uint8 = 4

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3
	expectedList := []*uint8{&vo2, &vo3, &vo4}
	newList := PMapUintUint8Ptr(plusOneUintUint8Ptr, []*uint{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUintUint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUintUint8Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUintUint8Ptr failed")
	}

	if len(PMapUintUint8Ptr(nil, []*uint{})) > 0 {
		t.Errorf("PMapUintUint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint8{&vo2, &vo3, &vo4}
	newList = PMapUintUint8Ptr(plusOneUintUint8Ptr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUintUint8Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUintStrPtr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 uint = 10
	var vi20 uint = 20

	expectedList := []*string{&vo10}
	newList := PMapUintStrPtr(someLogicUintStrPtr, []*uint{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapUintStrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUintStrPtr(nil, nil)) > 0 {
		t.Errorf("PMapUintStrPtr failed")
	}

	if len(PMapUintStrPtr(nil, []*uint{})) > 0 {
		t.Errorf("PMapUintStrPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*string{&vo10}
	newList = PMapUintStrPtr(someLogicUintStrPtr, []*uint{&vi10, &vi20}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != "0" {
		t.Errorf("PMapUintStrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapUintBoolPtr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 uint = 10
	var v0 uint = 0

	expectedList := []*bool{&vt, &vf}
	newList := PMapUintBoolPtr(someLogicUintBoolPtr, []*uint{&v10, &v0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUintBoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUintBoolPtr(nil, nil)) > 0 {
		t.Errorf("PMapUintBoolPtr failed")
	}

	if len(PMapUintBoolPtr(nil, []*uint{})) > 0 {
		t.Errorf("PMapUintBoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*bool{&vt, &vf}
	newList = PMapUintBoolPtr(someLogicUintBoolPtr, []*uint{&v10, &v0}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != false {
		t.Errorf("PMapUintBoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapUintFloat32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3
	var vo4 float32 = 4

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3
	expectedList := []*float32{&vo2, &vo3, &vo4}
	newList := PMapUintFloat32Ptr(plusOneUintFloat32Ptr, []*uint{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUintFloat32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUintFloat32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUintFloat32Ptr failed")
	}

	if len(PMapUintFloat32Ptr(nil, []*uint{})) > 0 {
		t.Errorf("PMapUintFloat32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*float32{&vo2, &vo3, &vo4}
	newList = PMapUintFloat32Ptr(plusOneUintFloat32Ptr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUintFloat32Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUintFloat64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3
	var vo4 float64 = 4

	var vi1 uint = 1
	var vi2 uint = 2
	var vi3 uint = 3
	expectedList := []*float64{&vo2, &vo3, &vo4}
	newList := PMapUintFloat64Ptr(plusOneUintFloat64Ptr, []*uint{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUintFloat64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUintFloat64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUintFloat64Ptr failed")
	}

	if len(PMapUintFloat64Ptr(nil, []*uint{})) > 0 {
		t.Errorf("PMapUintFloat64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*float64{&vo2, &vo3, &vo4}
	newList = PMapUintFloat64Ptr(plusOneUintFloat64Ptr, []*uint{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUintFloat64Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint64IntPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3
	var vo4 int = 4

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3
	expectedList := []*int{&vo2, &vo3, &vo4}
	newList := PMapUint64IntPtr(plusOneUint64IntPtr, []*uint64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint64IntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint64IntPtr(nil, nil)) > 0 {
		t.Errorf("PMapUint64IntPtr failed")
	}

	if len(PMapUint64IntPtr(nil, []*uint64{})) > 0 {
		t.Errorf("PMapUint64IntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int{&vo2, &vo3, &vo4}
	newList = PMapUint64IntPtr(plusOneUint64IntPtr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint64IntPtr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint64Int64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3
	var vo4 int64 = 4

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3
	expectedList := []*int64{&vo2, &vo3, &vo4}
	newList := PMapUint64Int64Ptr(plusOneUint64Int64Ptr, []*uint64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint64Int64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint64Int64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint64Int64Ptr failed")
	}

	if len(PMapUint64Int64Ptr(nil, []*uint64{})) > 0 {
		t.Errorf("PMapUint64Int64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int64{&vo2, &vo3, &vo4}
	newList = PMapUint64Int64Ptr(plusOneUint64Int64Ptr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint64Int64Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint64Int32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3
	var vo4 int32 = 4

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3
	expectedList := []*int32{&vo2, &vo3, &vo4}
	newList := PMapUint64Int32Ptr(plusOneUint64Int32Ptr, []*uint64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint64Int32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint64Int32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint64Int32Ptr failed")
	}

	if len(PMapUint64Int32Ptr(nil, []*uint64{})) > 0 {
		t.Errorf("PMapUint64Int32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int32{&vo2, &vo3, &vo4}
	newList = PMapUint64Int32Ptr(plusOneUint64Int32Ptr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint64Int32Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint64Int16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3
	var vo4 int16 = 4

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3
	expectedList := []*int16{&vo2, &vo3, &vo4}
	newList := PMapUint64Int16Ptr(plusOneUint64Int16Ptr, []*uint64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint64Int16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint64Int16Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint64Int16Ptr failed")
	}

	if len(PMapUint64Int16Ptr(nil, []*uint64{})) > 0 {
		t.Errorf("PMapUint64Int16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int16{&vo2, &vo3, &vo4}
	newList = PMapUint64Int16Ptr(plusOneUint64Int16Ptr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint64Int16Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint64Int8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3
	var vo4 int8 = 4

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3
	expectedList := []*int8{&vo2, &vo3, &vo4}
	newList := PMapUint64Int8Ptr(plusOneUint64Int8Ptr, []*uint64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint64Int8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint64Int8Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint64Int8Ptr failed")
	}

	if len(PMapUint64Int8Ptr(nil, []*uint64{})) > 0 {
		t.Errorf("PMapUint64Int8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int8{&vo2, &vo3, &vo4}
	newList = PMapUint64Int8Ptr(plusOneUint64Int8Ptr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint64Int8Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint64UintPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3
	var vo4 uint = 4

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3
	expectedList := []*uint{&vo2, &vo3, &vo4}
	newList := PMapUint64UintPtr(plusOneUint64UintPtr, []*uint64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint64UintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint64UintPtr(nil, nil)) > 0 {
		t.Errorf("PMapUint64UintPtr failed")
	}

	if len(PMapUint64UintPtr(nil, []*uint64{})) > 0 {
		t.Errorf("PMapUint64UintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint{&vo2, &vo3, &vo4}
	newList = PMapUint64UintPtr(plusOneUint64UintPtr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint64UintPtr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint64Uint32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3
	var vo4 uint32 = 4

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3
	expectedList := []*uint32{&vo2, &vo3, &vo4}
	newList := PMapUint64Uint32Ptr(plusOneUint64Uint32Ptr, []*uint64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint64Uint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint64Uint32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint64Uint32Ptr failed")
	}

	if len(PMapUint64Uint32Ptr(nil, []*uint64{})) > 0 {
		t.Errorf("PMapUint64Uint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint32{&vo2, &vo3, &vo4}
	newList = PMapUint64Uint32Ptr(plusOneUint64Uint32Ptr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint64Uint32Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint64Uint16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3
	var vo4 uint16 = 4

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3
	expectedList := []*uint16{&vo2, &vo3, &vo4}
	newList := PMapUint64Uint16Ptr(plusOneUint64Uint16Ptr, []*uint64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint64Uint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint64Uint16Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint64Uint16Ptr failed")
	}

	if len(PMapUint64Uint16Ptr(nil, []*uint64{})) > 0 {
		t.Errorf("PMapUint64Uint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint16{&vo2, &vo3, &vo4}
	newList = PMapUint64Uint16Ptr(plusOneUint64Uint16Ptr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint64Uint16Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint64Uint8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3
	var vo4 uint8 = 4

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3
	expectedList := []*uint8{&vo2, &vo3, &vo4}
	newList := PMapUint64Uint8Ptr(plusOneUint64Uint8Ptr, []*uint64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint64Uint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint64Uint8Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint64Uint8Ptr failed")
	}

	if len(PMapUint64Uint8Ptr(nil, []*uint64{})) > 0 {
		t.Errorf("PMapUint64Uint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint8{&vo2, &vo3, &vo4}
	newList = PMapUint64Uint8Ptr(plusOneUint64Uint8Ptr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint64Uint8Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint64StrPtr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 uint64 = 10
	var vi20 uint64 = 20

	expectedList := []*string{&vo10}
	newList := PMapUint64StrPtr(someLogicUint64StrPtr, []*uint64{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapUint64StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint64StrPtr(nil, nil)) > 0 {
		t.Errorf("PMapUint64StrPtr failed")
	}

	if len(PMapUint64StrPtr(nil, []*uint64{})) > 0 {
		t.Errorf("PMapUint64StrPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*string{&vo10}
	newList = PMapUint64StrPtr(someLogicUint64StrPtr, []*uint64{&vi10, &vi20}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != "0" {
		t.Errorf("PMapUint64StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapUint64BoolPtr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 uint64 = 10
	var v0 uint64 = 0

	expectedList := []*bool{&vt, &vf}
	newList := PMapUint64BoolPtr(someLogicUint64BoolPtr, []*uint64{&v10, &v0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint64BoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint64BoolPtr(nil, nil)) > 0 {
		t.Errorf("PMapUint64BoolPtr failed")
	}

	if len(PMapUint64BoolPtr(nil, []*uint64{})) > 0 {
		t.Errorf("PMapUint64BoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*bool{&vt, &vf}
	newList = PMapUint64BoolPtr(someLogicUint64BoolPtr, []*uint64{&v10, &v0}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != false {
		t.Errorf("PMapUint64BoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapUint64Float32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3
	var vo4 float32 = 4

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3
	expectedList := []*float32{&vo2, &vo3, &vo4}
	newList := PMapUint64Float32Ptr(plusOneUint64Float32Ptr, []*uint64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint64Float32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint64Float32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint64Float32Ptr failed")
	}

	if len(PMapUint64Float32Ptr(nil, []*uint64{})) > 0 {
		t.Errorf("PMapUint64Float32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*float32{&vo2, &vo3, &vo4}
	newList = PMapUint64Float32Ptr(plusOneUint64Float32Ptr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint64Float32Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint64Float64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3
	var vo4 float64 = 4

	var vi1 uint64 = 1
	var vi2 uint64 = 2
	var vi3 uint64 = 3
	expectedList := []*float64{&vo2, &vo3, &vo4}
	newList := PMapUint64Float64Ptr(plusOneUint64Float64Ptr, []*uint64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint64Float64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint64Float64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint64Float64Ptr failed")
	}

	if len(PMapUint64Float64Ptr(nil, []*uint64{})) > 0 {
		t.Errorf("PMapUint64Float64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*float64{&vo2, &vo3, &vo4}
	newList = PMapUint64Float64Ptr(plusOneUint64Float64Ptr, []*uint64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint64Float64Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint32IntPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3
	var vo4 int = 4

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3
	expectedList := []*int{&vo2, &vo3, &vo4}
	newList := PMapUint32IntPtr(plusOneUint32IntPtr, []*uint32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint32IntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint32IntPtr(nil, nil)) > 0 {
		t.Errorf("PMapUint32IntPtr failed")
	}

	if len(PMapUint32IntPtr(nil, []*uint32{})) > 0 {
		t.Errorf("PMapUint32IntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int{&vo2, &vo3, &vo4}
	newList = PMapUint32IntPtr(plusOneUint32IntPtr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint32IntPtr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint32Int64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3
	var vo4 int64 = 4

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3
	expectedList := []*int64{&vo2, &vo3, &vo4}
	newList := PMapUint32Int64Ptr(plusOneUint32Int64Ptr, []*uint32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint32Int64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint32Int64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint32Int64Ptr failed")
	}

	if len(PMapUint32Int64Ptr(nil, []*uint32{})) > 0 {
		t.Errorf("PMapUint32Int64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int64{&vo2, &vo3, &vo4}
	newList = PMapUint32Int64Ptr(plusOneUint32Int64Ptr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint32Int64Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint32Int32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3
	var vo4 int32 = 4

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3
	expectedList := []*int32{&vo2, &vo3, &vo4}
	newList := PMapUint32Int32Ptr(plusOneUint32Int32Ptr, []*uint32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint32Int32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint32Int32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint32Int32Ptr failed")
	}

	if len(PMapUint32Int32Ptr(nil, []*uint32{})) > 0 {
		t.Errorf("PMapUint32Int32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int32{&vo2, &vo3, &vo4}
	newList = PMapUint32Int32Ptr(plusOneUint32Int32Ptr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint32Int32Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint32Int16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3
	var vo4 int16 = 4

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3
	expectedList := []*int16{&vo2, &vo3, &vo4}
	newList := PMapUint32Int16Ptr(plusOneUint32Int16Ptr, []*uint32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint32Int16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint32Int16Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint32Int16Ptr failed")
	}

	if len(PMapUint32Int16Ptr(nil, []*uint32{})) > 0 {
		t.Errorf("PMapUint32Int16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int16{&vo2, &vo3, &vo4}
	newList = PMapUint32Int16Ptr(plusOneUint32Int16Ptr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint32Int16Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint32Int8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3
	var vo4 int8 = 4

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3
	expectedList := []*int8{&vo2, &vo3, &vo4}
	newList := PMapUint32Int8Ptr(plusOneUint32Int8Ptr, []*uint32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint32Int8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint32Int8Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint32Int8Ptr failed")
	}

	if len(PMapUint32Int8Ptr(nil, []*uint32{})) > 0 {
		t.Errorf("PMapUint32Int8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int8{&vo2, &vo3, &vo4}
	newList = PMapUint32Int8Ptr(plusOneUint32Int8Ptr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint32Int8Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint32UintPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3
	var vo4 uint = 4

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3
	expectedList := []*uint{&vo2, &vo3, &vo4}
	newList := PMapUint32UintPtr(plusOneUint32UintPtr, []*uint32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint32UintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint32UintPtr(nil, nil)) > 0 {
		t.Errorf("PMapUint32UintPtr failed")
	}

	if len(PMapUint32UintPtr(nil, []*uint32{})) > 0 {
		t.Errorf("PMapUint32UintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint{&vo2, &vo3, &vo4}
	newList = PMapUint32UintPtr(plusOneUint32UintPtr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint32UintPtr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint32Uint64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3
	var vo4 uint64 = 4

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3
	expectedList := []*uint64{&vo2, &vo3, &vo4}
	newList := PMapUint32Uint64Ptr(plusOneUint32Uint64Ptr, []*uint32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint32Uint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint32Uint64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint32Uint64Ptr failed")
	}

	if len(PMapUint32Uint64Ptr(nil, []*uint32{})) > 0 {
		t.Errorf("PMapUint32Uint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint64{&vo2, &vo3, &vo4}
	newList = PMapUint32Uint64Ptr(plusOneUint32Uint64Ptr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint32Uint64Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint32Uint16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3
	var vo4 uint16 = 4

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3
	expectedList := []*uint16{&vo2, &vo3, &vo4}
	newList := PMapUint32Uint16Ptr(plusOneUint32Uint16Ptr, []*uint32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint32Uint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint32Uint16Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint32Uint16Ptr failed")
	}

	if len(PMapUint32Uint16Ptr(nil, []*uint32{})) > 0 {
		t.Errorf("PMapUint32Uint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint16{&vo2, &vo3, &vo4}
	newList = PMapUint32Uint16Ptr(plusOneUint32Uint16Ptr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint32Uint16Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint32Uint8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3
	var vo4 uint8 = 4

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3
	expectedList := []*uint8{&vo2, &vo3, &vo4}
	newList := PMapUint32Uint8Ptr(plusOneUint32Uint8Ptr, []*uint32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint32Uint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint32Uint8Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint32Uint8Ptr failed")
	}

	if len(PMapUint32Uint8Ptr(nil, []*uint32{})) > 0 {
		t.Errorf("PMapUint32Uint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint8{&vo2, &vo3, &vo4}
	newList = PMapUint32Uint8Ptr(plusOneUint32Uint8Ptr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint32Uint8Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint32StrPtr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 uint32 = 10
	var vi20 uint32 = 20

	expectedList := []*string{&vo10}
	newList := PMapUint32StrPtr(someLogicUint32StrPtr, []*uint32{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapUint32StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint32StrPtr(nil, nil)) > 0 {
		t.Errorf("PMapUint32StrPtr failed")
	}

	if len(PMapUint32StrPtr(nil, []*uint32{})) > 0 {
		t.Errorf("PMapUint32StrPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*string{&vo10}
	newList = PMapUint32StrPtr(someLogicUint32StrPtr, []*uint32{&vi10, &vi20}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != "0" {
		t.Errorf("PMapUint32StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapUint32BoolPtr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 uint32 = 10
	var v0 uint32 = 0

	expectedList := []*bool{&vt, &vf}
	newList := PMapUint32BoolPtr(someLogicUint32BoolPtr, []*uint32{&v10, &v0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint32BoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint32BoolPtr(nil, nil)) > 0 {
		t.Errorf("PMapUint32BoolPtr failed")
	}

	if len(PMapUint32BoolPtr(nil, []*uint32{})) > 0 {
		t.Errorf("PMapUint32BoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*bool{&vt, &vf}
	newList = PMapUint32BoolPtr(someLogicUint32BoolPtr, []*uint32{&v10, &v0}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != false {
		t.Errorf("PMapUint32BoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapUint32Float32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3
	var vo4 float32 = 4

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3
	expectedList := []*float32{&vo2, &vo3, &vo4}
	newList := PMapUint32Float32Ptr(plusOneUint32Float32Ptr, []*uint32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint32Float32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint32Float32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint32Float32Ptr failed")
	}

	if len(PMapUint32Float32Ptr(nil, []*uint32{})) > 0 {
		t.Errorf("PMapUint32Float32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*float32{&vo2, &vo3, &vo4}
	newList = PMapUint32Float32Ptr(plusOneUint32Float32Ptr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint32Float32Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint32Float64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3
	var vo4 float64 = 4

	var vi1 uint32 = 1
	var vi2 uint32 = 2
	var vi3 uint32 = 3
	expectedList := []*float64{&vo2, &vo3, &vo4}
	newList := PMapUint32Float64Ptr(plusOneUint32Float64Ptr, []*uint32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint32Float64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint32Float64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint32Float64Ptr failed")
	}

	if len(PMapUint32Float64Ptr(nil, []*uint32{})) > 0 {
		t.Errorf("PMapUint32Float64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*float64{&vo2, &vo3, &vo4}
	newList = PMapUint32Float64Ptr(plusOneUint32Float64Ptr, []*uint32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint32Float64Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint16IntPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3
	var vo4 int = 4

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3
	expectedList := []*int{&vo2, &vo3, &vo4}
	newList := PMapUint16IntPtr(plusOneUint16IntPtr, []*uint16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint16IntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint16IntPtr(nil, nil)) > 0 {
		t.Errorf("PMapUint16IntPtr failed")
	}

	if len(PMapUint16IntPtr(nil, []*uint16{})) > 0 {
		t.Errorf("PMapUint16IntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int{&vo2, &vo3, &vo4}
	newList = PMapUint16IntPtr(plusOneUint16IntPtr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint16IntPtr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint16Int64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3
	var vo4 int64 = 4

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3
	expectedList := []*int64{&vo2, &vo3, &vo4}
	newList := PMapUint16Int64Ptr(plusOneUint16Int64Ptr, []*uint16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint16Int64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint16Int64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint16Int64Ptr failed")
	}

	if len(PMapUint16Int64Ptr(nil, []*uint16{})) > 0 {
		t.Errorf("PMapUint16Int64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int64{&vo2, &vo3, &vo4}
	newList = PMapUint16Int64Ptr(plusOneUint16Int64Ptr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint16Int64Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint16Int32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3
	var vo4 int32 = 4

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3
	expectedList := []*int32{&vo2, &vo3, &vo4}
	newList := PMapUint16Int32Ptr(plusOneUint16Int32Ptr, []*uint16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint16Int32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint16Int32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint16Int32Ptr failed")
	}

	if len(PMapUint16Int32Ptr(nil, []*uint16{})) > 0 {
		t.Errorf("PMapUint16Int32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int32{&vo2, &vo3, &vo4}
	newList = PMapUint16Int32Ptr(plusOneUint16Int32Ptr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint16Int32Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint16Int16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3
	var vo4 int16 = 4

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3
	expectedList := []*int16{&vo2, &vo3, &vo4}
	newList := PMapUint16Int16Ptr(plusOneUint16Int16Ptr, []*uint16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint16Int16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint16Int16Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint16Int16Ptr failed")
	}

	if len(PMapUint16Int16Ptr(nil, []*uint16{})) > 0 {
		t.Errorf("PMapUint16Int16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int16{&vo2, &vo3, &vo4}
	newList = PMapUint16Int16Ptr(plusOneUint16Int16Ptr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint16Int16Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint16Int8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3
	var vo4 int8 = 4

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3
	expectedList := []*int8{&vo2, &vo3, &vo4}
	newList := PMapUint16Int8Ptr(plusOneUint16Int8Ptr, []*uint16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint16Int8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint16Int8Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint16Int8Ptr failed")
	}

	if len(PMapUint16Int8Ptr(nil, []*uint16{})) > 0 {
		t.Errorf("PMapUint16Int8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int8{&vo2, &vo3, &vo4}
	newList = PMapUint16Int8Ptr(plusOneUint16Int8Ptr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint16Int8Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint16UintPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3
	var vo4 uint = 4

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3
	expectedList := []*uint{&vo2, &vo3, &vo4}
	newList := PMapUint16UintPtr(plusOneUint16UintPtr, []*uint16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint16UintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint16UintPtr(nil, nil)) > 0 {
		t.Errorf("PMapUint16UintPtr failed")
	}

	if len(PMapUint16UintPtr(nil, []*uint16{})) > 0 {
		t.Errorf("PMapUint16UintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint{&vo2, &vo3, &vo4}
	newList = PMapUint16UintPtr(plusOneUint16UintPtr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint16UintPtr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint16Uint64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3
	var vo4 uint64 = 4

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3
	expectedList := []*uint64{&vo2, &vo3, &vo4}
	newList := PMapUint16Uint64Ptr(plusOneUint16Uint64Ptr, []*uint16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint16Uint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint16Uint64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint16Uint64Ptr failed")
	}

	if len(PMapUint16Uint64Ptr(nil, []*uint16{})) > 0 {
		t.Errorf("PMapUint16Uint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint64{&vo2, &vo3, &vo4}
	newList = PMapUint16Uint64Ptr(plusOneUint16Uint64Ptr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint16Uint64Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint16Uint32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3
	var vo4 uint32 = 4

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3
	expectedList := []*uint32{&vo2, &vo3, &vo4}
	newList := PMapUint16Uint32Ptr(plusOneUint16Uint32Ptr, []*uint16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint16Uint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint16Uint32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint16Uint32Ptr failed")
	}

	if len(PMapUint16Uint32Ptr(nil, []*uint16{})) > 0 {
		t.Errorf("PMapUint16Uint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint32{&vo2, &vo3, &vo4}
	newList = PMapUint16Uint32Ptr(plusOneUint16Uint32Ptr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint16Uint32Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint16Uint8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3
	var vo4 uint8 = 4

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3
	expectedList := []*uint8{&vo2, &vo3, &vo4}
	newList := PMapUint16Uint8Ptr(plusOneUint16Uint8Ptr, []*uint16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint16Uint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint16Uint8Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint16Uint8Ptr failed")
	}

	if len(PMapUint16Uint8Ptr(nil, []*uint16{})) > 0 {
		t.Errorf("PMapUint16Uint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint8{&vo2, &vo3, &vo4}
	newList = PMapUint16Uint8Ptr(plusOneUint16Uint8Ptr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint16Uint8Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint16StrPtr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 uint16 = 10
	var vi20 uint16 = 20

	expectedList := []*string{&vo10}
	newList := PMapUint16StrPtr(someLogicUint16StrPtr, []*uint16{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapUint16StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint16StrPtr(nil, nil)) > 0 {
		t.Errorf("PMapUint16StrPtr failed")
	}

	if len(PMapUint16StrPtr(nil, []*uint16{})) > 0 {
		t.Errorf("PMapUint16StrPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*string{&vo10}
	newList = PMapUint16StrPtr(someLogicUint16StrPtr, []*uint16{&vi10, &vi20}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != "0" {
		t.Errorf("PMapUint16StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapUint16BoolPtr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 uint16 = 10
	var v0 uint16 = 0

	expectedList := []*bool{&vt, &vf}
	newList := PMapUint16BoolPtr(someLogicUint16BoolPtr, []*uint16{&v10, &v0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint16BoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint16BoolPtr(nil, nil)) > 0 {
		t.Errorf("PMapUint16BoolPtr failed")
	}

	if len(PMapUint16BoolPtr(nil, []*uint16{})) > 0 {
		t.Errorf("PMapUint16BoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*bool{&vt, &vf}
	newList = PMapUint16BoolPtr(someLogicUint16BoolPtr, []*uint16{&v10, &v0}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != false {
		t.Errorf("PMapUint16BoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapUint16Float32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3
	var vo4 float32 = 4

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3
	expectedList := []*float32{&vo2, &vo3, &vo4}
	newList := PMapUint16Float32Ptr(plusOneUint16Float32Ptr, []*uint16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint16Float32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint16Float32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint16Float32Ptr failed")
	}

	if len(PMapUint16Float32Ptr(nil, []*uint16{})) > 0 {
		t.Errorf("PMapUint16Float32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*float32{&vo2, &vo3, &vo4}
	newList = PMapUint16Float32Ptr(plusOneUint16Float32Ptr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint16Float32Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint16Float64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3
	var vo4 float64 = 4

	var vi1 uint16 = 1
	var vi2 uint16 = 2
	var vi3 uint16 = 3
	expectedList := []*float64{&vo2, &vo3, &vo4}
	newList := PMapUint16Float64Ptr(plusOneUint16Float64Ptr, []*uint16{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint16Float64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint16Float64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint16Float64Ptr failed")
	}

	if len(PMapUint16Float64Ptr(nil, []*uint16{})) > 0 {
		t.Errorf("PMapUint16Float64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*float64{&vo2, &vo3, &vo4}
	newList = PMapUint16Float64Ptr(plusOneUint16Float64Ptr, []*uint16{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint16Float64Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint8IntPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3
	var vo4 int = 4

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3
	expectedList := []*int{&vo2, &vo3, &vo4}
	newList := PMapUint8IntPtr(plusOneUint8IntPtr, []*uint8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint8IntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint8IntPtr(nil, nil)) > 0 {
		t.Errorf("PMapUint8IntPtr failed")
	}

	if len(PMapUint8IntPtr(nil, []*uint8{})) > 0 {
		t.Errorf("PMapUint8IntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int{&vo2, &vo3, &vo4}
	newList = PMapUint8IntPtr(plusOneUint8IntPtr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint8IntPtr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint8Int64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3
	var vo4 int64 = 4

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3
	expectedList := []*int64{&vo2, &vo3, &vo4}
	newList := PMapUint8Int64Ptr(plusOneUint8Int64Ptr, []*uint8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint8Int64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint8Int64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint8Int64Ptr failed")
	}

	if len(PMapUint8Int64Ptr(nil, []*uint8{})) > 0 {
		t.Errorf("PMapUint8Int64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int64{&vo2, &vo3, &vo4}
	newList = PMapUint8Int64Ptr(plusOneUint8Int64Ptr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint8Int64Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint8Int32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3
	var vo4 int32 = 4

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3
	expectedList := []*int32{&vo2, &vo3, &vo4}
	newList := PMapUint8Int32Ptr(plusOneUint8Int32Ptr, []*uint8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint8Int32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint8Int32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint8Int32Ptr failed")
	}

	if len(PMapUint8Int32Ptr(nil, []*uint8{})) > 0 {
		t.Errorf("PMapUint8Int32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int32{&vo2, &vo3, &vo4}
	newList = PMapUint8Int32Ptr(plusOneUint8Int32Ptr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint8Int32Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint8Int16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3
	var vo4 int16 = 4

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3
	expectedList := []*int16{&vo2, &vo3, &vo4}
	newList := PMapUint8Int16Ptr(plusOneUint8Int16Ptr, []*uint8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint8Int16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint8Int16Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint8Int16Ptr failed")
	}

	if len(PMapUint8Int16Ptr(nil, []*uint8{})) > 0 {
		t.Errorf("PMapUint8Int16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int16{&vo2, &vo3, &vo4}
	newList = PMapUint8Int16Ptr(plusOneUint8Int16Ptr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint8Int16Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint8Int8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3
	var vo4 int8 = 4

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3
	expectedList := []*int8{&vo2, &vo3, &vo4}
	newList := PMapUint8Int8Ptr(plusOneUint8Int8Ptr, []*uint8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint8Int8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint8Int8Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint8Int8Ptr failed")
	}

	if len(PMapUint8Int8Ptr(nil, []*uint8{})) > 0 {
		t.Errorf("PMapUint8Int8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int8{&vo2, &vo3, &vo4}
	newList = PMapUint8Int8Ptr(plusOneUint8Int8Ptr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint8Int8Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint8UintPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3
	var vo4 uint = 4

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3
	expectedList := []*uint{&vo2, &vo3, &vo4}
	newList := PMapUint8UintPtr(plusOneUint8UintPtr, []*uint8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint8UintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint8UintPtr(nil, nil)) > 0 {
		t.Errorf("PMapUint8UintPtr failed")
	}

	if len(PMapUint8UintPtr(nil, []*uint8{})) > 0 {
		t.Errorf("PMapUint8UintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint{&vo2, &vo3, &vo4}
	newList = PMapUint8UintPtr(plusOneUint8UintPtr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint8UintPtr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint8Uint64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3
	var vo4 uint64 = 4

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3
	expectedList := []*uint64{&vo2, &vo3, &vo4}
	newList := PMapUint8Uint64Ptr(plusOneUint8Uint64Ptr, []*uint8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint8Uint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint8Uint64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint8Uint64Ptr failed")
	}

	if len(PMapUint8Uint64Ptr(nil, []*uint8{})) > 0 {
		t.Errorf("PMapUint8Uint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint64{&vo2, &vo3, &vo4}
	newList = PMapUint8Uint64Ptr(plusOneUint8Uint64Ptr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint8Uint64Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint8Uint32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3
	var vo4 uint32 = 4

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3
	expectedList := []*uint32{&vo2, &vo3, &vo4}
	newList := PMapUint8Uint32Ptr(plusOneUint8Uint32Ptr, []*uint8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint8Uint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint8Uint32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint8Uint32Ptr failed")
	}

	if len(PMapUint8Uint32Ptr(nil, []*uint8{})) > 0 {
		t.Errorf("PMapUint8Uint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint32{&vo2, &vo3, &vo4}
	newList = PMapUint8Uint32Ptr(plusOneUint8Uint32Ptr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint8Uint32Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint8Uint16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3
	var vo4 uint16 = 4

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3
	expectedList := []*uint16{&vo2, &vo3, &vo4}
	newList := PMapUint8Uint16Ptr(plusOneUint8Uint16Ptr, []*uint8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint8Uint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint8Uint16Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint8Uint16Ptr failed")
	}

	if len(PMapUint8Uint16Ptr(nil, []*uint8{})) > 0 {
		t.Errorf("PMapUint8Uint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint16{&vo2, &vo3, &vo4}
	newList = PMapUint8Uint16Ptr(plusOneUint8Uint16Ptr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint8Uint16Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint8StrPtr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 uint8 = 10
	var vi20 uint8 = 20

	expectedList := []*string{&vo10}
	newList := PMapUint8StrPtr(someLogicUint8StrPtr, []*uint8{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapUint8StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint8StrPtr(nil, nil)) > 0 {
		t.Errorf("PMapUint8StrPtr failed")
	}

	if len(PMapUint8StrPtr(nil, []*uint8{})) > 0 {
		t.Errorf("PMapUint8StrPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*string{&vo10}
	newList = PMapUint8StrPtr(someLogicUint8StrPtr, []*uint8{&vi10, &vi20}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != "0" {
		t.Errorf("PMapUint8StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapUint8BoolPtr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 uint8 = 10
	var v0 uint8 = 0

	expectedList := []*bool{&vt, &vf}
	newList := PMapUint8BoolPtr(someLogicUint8BoolPtr, []*uint8{&v10, &v0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapUint8BoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint8BoolPtr(nil, nil)) > 0 {
		t.Errorf("PMapUint8BoolPtr failed")
	}

	if len(PMapUint8BoolPtr(nil, []*uint8{})) > 0 {
		t.Errorf("PMapUint8BoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*bool{&vt, &vf}
	newList = PMapUint8BoolPtr(someLogicUint8BoolPtr, []*uint8{&v10, &v0}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != false {
		t.Errorf("PMapUint8BoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapUint8Float32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3
	var vo4 float32 = 4

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3
	expectedList := []*float32{&vo2, &vo3, &vo4}
	newList := PMapUint8Float32Ptr(plusOneUint8Float32Ptr, []*uint8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint8Float32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint8Float32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint8Float32Ptr failed")
	}

	if len(PMapUint8Float32Ptr(nil, []*uint8{})) > 0 {
		t.Errorf("PMapUint8Float32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*float32{&vo2, &vo3, &vo4}
	newList = PMapUint8Float32Ptr(plusOneUint8Float32Ptr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint8Float32Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapUint8Float64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3
	var vo4 float64 = 4

	var vi1 uint8 = 1
	var vi2 uint8 = 2
	var vi3 uint8 = 3
	expectedList := []*float64{&vo2, &vo3, &vo4}
	newList := PMapUint8Float64Ptr(plusOneUint8Float64Ptr, []*uint8{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapUint8Float64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapUint8Float64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint8Float64Ptr failed")
	}

	if len(PMapUint8Float64Ptr(nil, []*uint8{})) > 0 {
		t.Errorf("PMapUint8Float64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*float64{&vo2, &vo3, &vo4}
	newList = PMapUint8Float64Ptr(plusOneUint8Float64Ptr, []*uint8{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapUint8Float64Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapStrIntPtr(t *testing.T) {
	// Test : someLogic
	var vo10 int = 10

	var vi10 string = "10"
	var vi1 string = "1"

	expectedList := []*int{&vo10}
	newList := PMapStrIntPtr(someLogicStrIntPtr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapStrIntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapStrIntPtr(nil, nil)) > 0 {
		t.Errorf("PMapStrIntPtr failed")
	}

	if len(PMapStrIntPtr(nil, []*string{})) > 0 {
		t.Errorf("PMapStrIntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
	expectedList = []*int{&vo10}
	newList = PMapStrIntPtr(someLogicStrIntPtr, []*string{&vi10, &vi1}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != 0 {
		t.Errorf("PMapStrInt failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapStrInt64Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 int64 = 10

	var vi10 string = "10"
	var vi1 string = "1"

	expectedList := []*int64{&vo10}
	newList := PMapStrInt64Ptr(someLogicStrInt64Ptr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapStrInt64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapStrInt64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapStrInt64Ptr failed")
	}

	if len(PMapStrInt64Ptr(nil, []*string{})) > 0 {
		t.Errorf("PMapStrInt64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
	expectedList = []*int64{&vo10}
	newList = PMapStrInt64Ptr(someLogicStrInt64Ptr, []*string{&vi10, &vi1}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != 0 {
		t.Errorf("PMapStrInt64 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapStrInt32Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 int32 = 10

	var vi10 string = "10"
	var vi1 string = "1"

	expectedList := []*int32{&vo10}
	newList := PMapStrInt32Ptr(someLogicStrInt32Ptr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapStrInt32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapStrInt32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapStrInt32Ptr failed")
	}

	if len(PMapStrInt32Ptr(nil, []*string{})) > 0 {
		t.Errorf("PMapStrInt32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
	expectedList = []*int32{&vo10}
	newList = PMapStrInt32Ptr(someLogicStrInt32Ptr, []*string{&vi10, &vi1}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != 0 {
		t.Errorf("PMapStrInt32 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapStrInt16Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 int16 = 10

	var vi10 string = "10"
	var vi1 string = "1"

	expectedList := []*int16{&vo10}
	newList := PMapStrInt16Ptr(someLogicStrInt16Ptr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapStrInt16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapStrInt16Ptr(nil, nil)) > 0 {
		t.Errorf("PMapStrInt16Ptr failed")
	}

	if len(PMapStrInt16Ptr(nil, []*string{})) > 0 {
		t.Errorf("PMapStrInt16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
	expectedList = []*int16{&vo10}
	newList = PMapStrInt16Ptr(someLogicStrInt16Ptr, []*string{&vi10, &vi1}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != 0 {
		t.Errorf("PMapStrInt16 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapStrInt8Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 int8 = 10

	var vi10 string = "10"
	var vi1 string = "1"

	expectedList := []*int8{&vo10}
	newList := PMapStrInt8Ptr(someLogicStrInt8Ptr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapStrInt8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapStrInt8Ptr(nil, nil)) > 0 {
		t.Errorf("PMapStrInt8Ptr failed")
	}

	if len(PMapStrInt8Ptr(nil, []*string{})) > 0 {
		t.Errorf("PMapStrInt8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
	expectedList = []*int8{&vo10}
	newList = PMapStrInt8Ptr(someLogicStrInt8Ptr, []*string{&vi10, &vi1}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != 0 {
		t.Errorf("PMapStrInt8 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapStrUintPtr(t *testing.T) {
	// Test : someLogic
	var vo10 uint = 10

	var vi10 string = "10"
	var vi1 string = "1"

	expectedList := []*uint{&vo10}
	newList := PMapStrUintPtr(someLogicStrUintPtr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapStrUintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapStrUintPtr(nil, nil)) > 0 {
		t.Errorf("PMapStrUintPtr failed")
	}

	if len(PMapStrUintPtr(nil, []*string{})) > 0 {
		t.Errorf("PMapStrUintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
	expectedList = []*uint{&vo10}
	newList = PMapStrUintPtr(someLogicStrUintPtr, []*string{&vi10, &vi1}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != 0 {
		t.Errorf("PMapStrUint failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapStrUint64Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 uint64 = 10

	var vi10 string = "10"
	var vi1 string = "1"

	expectedList := []*uint64{&vo10}
	newList := PMapStrUint64Ptr(someLogicStrUint64Ptr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapStrUint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapStrUint64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapStrUint64Ptr failed")
	}

	if len(PMapStrUint64Ptr(nil, []*string{})) > 0 {
		t.Errorf("PMapStrUint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
	expectedList = []*uint64{&vo10}
	newList = PMapStrUint64Ptr(someLogicStrUint64Ptr, []*string{&vi10, &vi1}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != 0 {
		t.Errorf("PMapStrUint64 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapStrUint32Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 uint32 = 10

	var vi10 string = "10"
	var vi1 string = "1"

	expectedList := []*uint32{&vo10}
	newList := PMapStrUint32Ptr(someLogicStrUint32Ptr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapStrUint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapStrUint32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapStrUint32Ptr failed")
	}

	if len(PMapStrUint32Ptr(nil, []*string{})) > 0 {
		t.Errorf("PMapStrUint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
	expectedList = []*uint32{&vo10}
	newList = PMapStrUint32Ptr(someLogicStrUint32Ptr, []*string{&vi10, &vi1}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != 0 {
		t.Errorf("PMapStrUint32 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapStrUint16Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 uint16 = 10

	var vi10 string = "10"
	var vi1 string = "1"

	expectedList := []*uint16{&vo10}
	newList := PMapStrUint16Ptr(someLogicStrUint16Ptr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapStrUint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapStrUint16Ptr(nil, nil)) > 0 {
		t.Errorf("PMapStrUint16Ptr failed")
	}

	if len(PMapStrUint16Ptr(nil, []*string{})) > 0 {
		t.Errorf("PMapStrUint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
	expectedList = []*uint16{&vo10}
	newList = PMapStrUint16Ptr(someLogicStrUint16Ptr, []*string{&vi10, &vi1}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != 0 {
		t.Errorf("PMapStrUint16 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapStrUint8Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 uint8 = 10

	var vi10 string = "10"
	var vi1 string = "1"

	expectedList := []*uint8{&vo10}
	newList := PMapStrUint8Ptr(someLogicStrUint8Ptr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapStrUint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapStrUint8Ptr(nil, nil)) > 0 {
		t.Errorf("PMapStrUint8Ptr failed")
	}

	if len(PMapStrUint8Ptr(nil, []*string{})) > 0 {
		t.Errorf("PMapStrUint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
	expectedList = []*uint8{&vo10}
	newList = PMapStrUint8Ptr(someLogicStrUint8Ptr, []*string{&vi10, &vi1}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != 0 {
		t.Errorf("PMapStrUint8 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapStrBoolPtr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 string = "10"
	var v0 string = "0"

	expectedList := []*bool{&vt, &vf}
	newList := PMapStrBoolPtr(someLogicStrBoolPtr, []*string{&v10, &v0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapStrBoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapStrBoolPtr(nil, nil)) > 0 {
		t.Errorf("PMapStrBoolPtr failed")
	}

	if len(PMapStrBoolPtr(nil, []*string{})) > 0 {
		t.Errorf("PMapStrBoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*bool{&vt, &vf}
	newList = PMapStrBoolPtr(someLogicStrBoolPtr, []*string{&v10, &v0}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != false {
		t.Errorf("PMapStrBoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapStrFloat32Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 float32 = 10

	var vi10 string = "10"
	var vi1 string = "1"

	expectedList := []*float32{&vo10}
	newList := PMapStrFloat32Ptr(someLogicStrFloat32Ptr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapStrFloat32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapStrFloat32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapStrFloat32Ptr failed")
	}

	if len(PMapStrFloat32Ptr(nil, []*string{})) > 0 {
		t.Errorf("PMapStrFloat32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
	expectedList = []*float32{&vo10}
	newList = PMapStrFloat32Ptr(someLogicStrFloat32Ptr, []*string{&vi10, &vi1}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != 0 {
		t.Errorf("PMapStrFloat32 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapStrFloat64Ptr(t *testing.T) {
	// Test : someLogic
	var vo10 float64 = 10

	var vi10 string = "10"
	var vi1 string = "1"

	expectedList := []*float64{&vo10}
	newList := PMapStrFloat64Ptr(someLogicStrFloat64Ptr, []*string{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapStrFloat64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapStrFloat64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapStrFloat64Ptr failed")
	}

	if len(PMapStrFloat64Ptr(nil, []*string{})) > 0 {
		t.Errorf("PMapStrFloat64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect
	expectedList = []*float64{&vo10}
	newList = PMapStrFloat64Ptr(someLogicStrFloat64Ptr, []*string{&vi10, &vi1}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != 0 {
		t.Errorf("PMapStrFloat64 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapBoolIntPtr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 int = 10
	var v0 int = 0

	expectedList := []*int{&v10, &v0}
	newList := PMapBoolIntPtr(someLogicBoolIntPtr, []*bool{&vt, &vf})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolIntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapBoolIntPtr(nil, nil)) > 0 {
		t.Errorf("PMapBoolIntPtr failed")
	}

	if len(PMapBoolIntPtr(nil, []*bool{})) > 0 {
		t.Errorf("PMapBoolIntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int{&v10, &v0}
	newList = PMapBoolIntPtr(someLogicBoolIntPtr, []*bool{&vt, &vf}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolIntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapBoolInt64Ptr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 int64 = 10
	var v0 int64 = 0

	expectedList := []*int64{&v10, &v0}
	newList := PMapBoolInt64Ptr(someLogicBoolInt64Ptr, []*bool{&vt, &vf})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolInt64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapBoolInt64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapBoolInt64Ptr failed")
	}

	if len(PMapBoolInt64Ptr(nil, []*bool{})) > 0 {
		t.Errorf("PMapBoolInt64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int64{&v10, &v0}
	newList = PMapBoolInt64Ptr(someLogicBoolInt64Ptr, []*bool{&vt, &vf}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolInt64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapBoolInt32Ptr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 int32 = 10
	var v0 int32 = 0

	expectedList := []*int32{&v10, &v0}
	newList := PMapBoolInt32Ptr(someLogicBoolInt32Ptr, []*bool{&vt, &vf})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolInt32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapBoolInt32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapBoolInt32Ptr failed")
	}

	if len(PMapBoolInt32Ptr(nil, []*bool{})) > 0 {
		t.Errorf("PMapBoolInt32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int32{&v10, &v0}
	newList = PMapBoolInt32Ptr(someLogicBoolInt32Ptr, []*bool{&vt, &vf}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolInt32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapBoolInt16Ptr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 int16 = 10
	var v0 int16 = 0

	expectedList := []*int16{&v10, &v0}
	newList := PMapBoolInt16Ptr(someLogicBoolInt16Ptr, []*bool{&vt, &vf})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolInt16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapBoolInt16Ptr(nil, nil)) > 0 {
		t.Errorf("PMapBoolInt16Ptr failed")
	}

	if len(PMapBoolInt16Ptr(nil, []*bool{})) > 0 {
		t.Errorf("PMapBoolInt16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int16{&v10, &v0}
	newList = PMapBoolInt16Ptr(someLogicBoolInt16Ptr, []*bool{&vt, &vf}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolInt16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapBoolInt8Ptr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 int8 = 10
	var v0 int8 = 0

	expectedList := []*int8{&v10, &v0}
	newList := PMapBoolInt8Ptr(someLogicBoolInt8Ptr, []*bool{&vt, &vf})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolInt8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapBoolInt8Ptr(nil, nil)) > 0 {
		t.Errorf("PMapBoolInt8Ptr failed")
	}

	if len(PMapBoolInt8Ptr(nil, []*bool{})) > 0 {
		t.Errorf("PMapBoolInt8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int8{&v10, &v0}
	newList = PMapBoolInt8Ptr(someLogicBoolInt8Ptr, []*bool{&vt, &vf}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolInt8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapBoolUintPtr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 uint = 10
	var v0 uint = 0

	expectedList := []*uint{&v10, &v0}
	newList := PMapBoolUintPtr(someLogicBoolUintPtr, []*bool{&vt, &vf})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolUintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapBoolUintPtr(nil, nil)) > 0 {
		t.Errorf("PMapBoolUintPtr failed")
	}

	if len(PMapBoolUintPtr(nil, []*bool{})) > 0 {
		t.Errorf("PMapBoolUintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint{&v10, &v0}
	newList = PMapBoolUintPtr(someLogicBoolUintPtr, []*bool{&vt, &vf}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolUintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapBoolUint64Ptr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 uint64 = 10
	var v0 uint64 = 0

	expectedList := []*uint64{&v10, &v0}
	newList := PMapBoolUint64Ptr(someLogicBoolUint64Ptr, []*bool{&vt, &vf})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolUint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapBoolUint64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapBoolUint64Ptr failed")
	}

	if len(PMapBoolUint64Ptr(nil, []*bool{})) > 0 {
		t.Errorf("PMapBoolUint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint64{&v10, &v0}
	newList = PMapBoolUint64Ptr(someLogicBoolUint64Ptr, []*bool{&vt, &vf}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolUint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapBoolUint32Ptr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 uint32 = 10
	var v0 uint32 = 0

	expectedList := []*uint32{&v10, &v0}
	newList := PMapBoolUint32Ptr(someLogicBoolUint32Ptr, []*bool{&vt, &vf})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolUint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapBoolUint32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapBoolUint32Ptr failed")
	}

	if len(PMapBoolUint32Ptr(nil, []*bool{})) > 0 {
		t.Errorf("PMapBoolUint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint32{&v10, &v0}
	newList = PMapBoolUint32Ptr(someLogicBoolUint32Ptr, []*bool{&vt, &vf}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolUint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapBoolUint16Ptr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 uint16 = 10
	var v0 uint16 = 0

	expectedList := []*uint16{&v10, &v0}
	newList := PMapBoolUint16Ptr(someLogicBoolUint16Ptr, []*bool{&vt, &vf})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolUint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapBoolUint16Ptr(nil, nil)) > 0 {
		t.Errorf("PMapBoolUint16Ptr failed")
	}

	if len(PMapBoolUint16Ptr(nil, []*bool{})) > 0 {
		t.Errorf("PMapBoolUint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint16{&v10, &v0}
	newList = PMapBoolUint16Ptr(someLogicBoolUint16Ptr, []*bool{&vt, &vf}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolUint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapBoolUint8Ptr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 uint8 = 10
	var v0 uint8 = 0

	expectedList := []*uint8{&v10, &v0}
	newList := PMapBoolUint8Ptr(someLogicBoolUint8Ptr, []*bool{&vt, &vf})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolUint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapBoolUint8Ptr(nil, nil)) > 0 {
		t.Errorf("PMapBoolUint8Ptr failed")
	}

	if len(PMapBoolUint8Ptr(nil, []*bool{})) > 0 {
		t.Errorf("PMapBoolUint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint8{&v10, &v0}
	newList = PMapBoolUint8Ptr(someLogicBoolUint8Ptr, []*bool{&vt, &vf}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolUint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapBoolStrPtr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 string = "10"
	var v0 string = "0"

	expectedList := []*string{&v10, &v0}
	newList := PMapBoolStrPtr(someLogicBoolStrPtr, []*bool{&vt, &vf})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolStrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapBoolStrPtr(nil, nil)) > 0 {
		t.Errorf("PMapBoolStrPtr failed")
	}

	if len(PMapBoolStrPtr(nil, []*bool{})) > 0 {
		t.Errorf("PMapBoolStrPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*string{&v10, &v0}
	newList = PMapBoolStrPtr(someLogicBoolStrPtr, []*bool{&vt, &vf}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolStrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapBoolFloat32Ptr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 float32 = 10
	var v0 float32 = 0

	expectedList := []*float32{&v10, &v0}
	newList := PMapBoolFloat32Ptr(someLogicBoolFloat32Ptr, []*bool{&vt, &vf})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolFloat32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapBoolFloat32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapBoolFloat32Ptr failed")
	}

	if len(PMapBoolFloat32Ptr(nil, []*bool{})) > 0 {
		t.Errorf("PMapBoolFloat32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*float32{&v10, &v0}
	newList = PMapBoolFloat32Ptr(someLogicBoolFloat32Ptr, []*bool{&vt, &vf}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolFloat32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapBoolFloat64Ptr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 float64 = 10
	var v0 float64 = 0

	expectedList := []*float64{&v10, &v0}
	newList := PMapBoolFloat64Ptr(someLogicBoolFloat64Ptr, []*bool{&vt, &vf})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolFloat64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapBoolFloat64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapBoolFloat64Ptr failed")
	}

	if len(PMapBoolFloat64Ptr(nil, []*bool{})) > 0 {
		t.Errorf("PMapBoolFloat64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*float64{&v10, &v0}
	newList = PMapBoolFloat64Ptr(someLogicBoolFloat64Ptr, []*bool{&vt, &vf}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapBoolFloat64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapFloat32IntPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3
	var vo4 int = 4

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3
	expectedList := []*int{&vo2, &vo3, &vo4}
	newList := PMapFloat32IntPtr(plusOneFloat32IntPtr, []*float32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapFloat32IntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat32IntPtr(nil, nil)) > 0 {
		t.Errorf("PMapFloat32IntPtr failed")
	}

	if len(PMapFloat32IntPtr(nil, []*float32{})) > 0 {
		t.Errorf("PMapFloat32IntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int{&vo2, &vo3, &vo4}
	newList = PMapFloat32IntPtr(plusOneFloat32IntPtr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat32IntPtr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapFloat32Int64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3
	var vo4 int64 = 4

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3
	expectedList := []*int64{&vo2, &vo3, &vo4}
	newList := PMapFloat32Int64Ptr(plusOneFloat32Int64Ptr, []*float32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapFloat32Int64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat32Int64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapFloat32Int64Ptr failed")
	}

	if len(PMapFloat32Int64Ptr(nil, []*float32{})) > 0 {
		t.Errorf("PMapFloat32Int64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int64{&vo2, &vo3, &vo4}
	newList = PMapFloat32Int64Ptr(plusOneFloat32Int64Ptr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat32Int64Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapFloat32Int32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3
	var vo4 int32 = 4

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3
	expectedList := []*int32{&vo2, &vo3, &vo4}
	newList := PMapFloat32Int32Ptr(plusOneFloat32Int32Ptr, []*float32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapFloat32Int32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat32Int32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapFloat32Int32Ptr failed")
	}

	if len(PMapFloat32Int32Ptr(nil, []*float32{})) > 0 {
		t.Errorf("PMapFloat32Int32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int32{&vo2, &vo3, &vo4}
	newList = PMapFloat32Int32Ptr(plusOneFloat32Int32Ptr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat32Int32Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapFloat32Int16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3
	var vo4 int16 = 4

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3
	expectedList := []*int16{&vo2, &vo3, &vo4}
	newList := PMapFloat32Int16Ptr(plusOneFloat32Int16Ptr, []*float32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapFloat32Int16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat32Int16Ptr(nil, nil)) > 0 {
		t.Errorf("PMapFloat32Int16Ptr failed")
	}

	if len(PMapFloat32Int16Ptr(nil, []*float32{})) > 0 {
		t.Errorf("PMapFloat32Int16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int16{&vo2, &vo3, &vo4}
	newList = PMapFloat32Int16Ptr(plusOneFloat32Int16Ptr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat32Int16Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapFloat32Int8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3
	var vo4 int8 = 4

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3
	expectedList := []*int8{&vo2, &vo3, &vo4}
	newList := PMapFloat32Int8Ptr(plusOneFloat32Int8Ptr, []*float32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapFloat32Int8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat32Int8Ptr(nil, nil)) > 0 {
		t.Errorf("PMapFloat32Int8Ptr failed")
	}

	if len(PMapFloat32Int8Ptr(nil, []*float32{})) > 0 {
		t.Errorf("PMapFloat32Int8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int8{&vo2, &vo3, &vo4}
	newList = PMapFloat32Int8Ptr(plusOneFloat32Int8Ptr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat32Int8Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapFloat32UintPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3
	var vo4 uint = 4

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3
	expectedList := []*uint{&vo2, &vo3, &vo4}
	newList := PMapFloat32UintPtr(plusOneFloat32UintPtr, []*float32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapFloat32UintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat32UintPtr(nil, nil)) > 0 {
		t.Errorf("PMapFloat32UintPtr failed")
	}

	if len(PMapFloat32UintPtr(nil, []*float32{})) > 0 {
		t.Errorf("PMapFloat32UintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint{&vo2, &vo3, &vo4}
	newList = PMapFloat32UintPtr(plusOneFloat32UintPtr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat32UintPtr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapFloat32Uint64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3
	var vo4 uint64 = 4

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3
	expectedList := []*uint64{&vo2, &vo3, &vo4}
	newList := PMapFloat32Uint64Ptr(plusOneFloat32Uint64Ptr, []*float32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapFloat32Uint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat32Uint64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapFloat32Uint64Ptr failed")
	}

	if len(PMapFloat32Uint64Ptr(nil, []*float32{})) > 0 {
		t.Errorf("PMapFloat32Uint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint64{&vo2, &vo3, &vo4}
	newList = PMapFloat32Uint64Ptr(plusOneFloat32Uint64Ptr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat32Uint64Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapFloat32Uint32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3
	var vo4 uint32 = 4

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3
	expectedList := []*uint32{&vo2, &vo3, &vo4}
	newList := PMapFloat32Uint32Ptr(plusOneFloat32Uint32Ptr, []*float32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapFloat32Uint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat32Uint32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapFloat32Uint32Ptr failed")
	}

	if len(PMapFloat32Uint32Ptr(nil, []*float32{})) > 0 {
		t.Errorf("PMapFloat32Uint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint32{&vo2, &vo3, &vo4}
	newList = PMapFloat32Uint32Ptr(plusOneFloat32Uint32Ptr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat32Uint32Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapFloat32Uint16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3
	var vo4 uint16 = 4

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3
	expectedList := []*uint16{&vo2, &vo3, &vo4}
	newList := PMapFloat32Uint16Ptr(plusOneFloat32Uint16Ptr, []*float32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapFloat32Uint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat32Uint16Ptr(nil, nil)) > 0 {
		t.Errorf("PMapFloat32Uint16Ptr failed")
	}

	if len(PMapFloat32Uint16Ptr(nil, []*float32{})) > 0 {
		t.Errorf("PMapFloat32Uint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint16{&vo2, &vo3, &vo4}
	newList = PMapFloat32Uint16Ptr(plusOneFloat32Uint16Ptr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat32Uint16Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapFloat32Uint8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3
	var vo4 uint8 = 4

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3
	expectedList := []*uint8{&vo2, &vo3, &vo4}
	newList := PMapFloat32Uint8Ptr(plusOneFloat32Uint8Ptr, []*float32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapFloat32Uint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat32Uint8Ptr(nil, nil)) > 0 {
		t.Errorf("PMapFloat32Uint8Ptr failed")
	}

	if len(PMapFloat32Uint8Ptr(nil, []*float32{})) > 0 {
		t.Errorf("PMapFloat32Uint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint8{&vo2, &vo3, &vo4}
	newList = PMapFloat32Uint8Ptr(plusOneFloat32Uint8Ptr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat32Uint8Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapFloat32StrPtr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 float32 = 10
	var vi20 float32 = 20

	expectedList := []*string{&vo10}
	newList := PMapFloat32StrPtr(someLogicFloat32StrPtr, []*float32{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapFloat32StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat32StrPtr(nil, nil)) > 0 {
		t.Errorf("PMapFloat32StrPtr failed")
	}

	if len(PMapFloat32StrPtr(nil, []*float32{})) > 0 {
		t.Errorf("PMapFloat32StrPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*string{&vo10}
	newList = PMapFloat32StrPtr(someLogicFloat32StrPtr, []*float32{&vi10, &vi20}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != "0" {
		t.Errorf("PMapFloat32StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapFloat32BoolPtr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 float32 = 10
	var v0 float32 = 0

	expectedList := []*bool{&vt, &vf}
	newList := PMapFloat32BoolPtr(someLogicFloat32BoolPtr, []*float32{&v10, &v0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapFloat32BoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat32BoolPtr(nil, nil)) > 0 {
		t.Errorf("PMapFloat32BoolPtr failed")
	}

	if len(PMapFloat32BoolPtr(nil, []*float32{})) > 0 {
		t.Errorf("PMapFloat32BoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*bool{&vt, &vf}
	newList = PMapFloat32BoolPtr(someLogicFloat32BoolPtr, []*float32{&v10, &v0}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != false {
		t.Errorf("PMapFloat32BoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapFloat32Float64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float64 = 2
	var vo3 float64 = 3
	var vo4 float64 = 4

	var vi1 float32 = 1
	var vi2 float32 = 2
	var vi3 float32 = 3
	expectedList := []*float64{&vo2, &vo3, &vo4}
	newList := PMapFloat32Float64Ptr(plusOneFloat32Float64Ptr, []*float32{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapFloat32Float64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat32Float64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapFloat32Float64Ptr failed")
	}

	if len(PMapFloat32Float64Ptr(nil, []*float32{})) > 0 {
		t.Errorf("PMapFloat32Float64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*float64{&vo2, &vo3, &vo4}
	newList = PMapFloat32Float64Ptr(plusOneFloat32Float64Ptr, []*float32{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat32Float64Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapFloat64IntPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int = 2
	var vo3 int = 3
	var vo4 int = 4

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3
	expectedList := []*int{&vo2, &vo3, &vo4}
	newList := PMapFloat64IntPtr(plusOneFloat64IntPtr, []*float64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapFloat64IntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat64IntPtr(nil, nil)) > 0 {
		t.Errorf("PMapFloat64IntPtr failed")
	}

	if len(PMapFloat64IntPtr(nil, []*float64{})) > 0 {
		t.Errorf("PMapFloat64IntPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int{&vo2, &vo3, &vo4}
	newList = PMapFloat64IntPtr(plusOneFloat64IntPtr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat64IntPtr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapFloat64Int64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int64 = 2
	var vo3 int64 = 3
	var vo4 int64 = 4

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3
	expectedList := []*int64{&vo2, &vo3, &vo4}
	newList := PMapFloat64Int64Ptr(plusOneFloat64Int64Ptr, []*float64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapFloat64Int64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat64Int64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapFloat64Int64Ptr failed")
	}

	if len(PMapFloat64Int64Ptr(nil, []*float64{})) > 0 {
		t.Errorf("PMapFloat64Int64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int64{&vo2, &vo3, &vo4}
	newList = PMapFloat64Int64Ptr(plusOneFloat64Int64Ptr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat64Int64Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapFloat64Int32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int32 = 2
	var vo3 int32 = 3
	var vo4 int32 = 4

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3
	expectedList := []*int32{&vo2, &vo3, &vo4}
	newList := PMapFloat64Int32Ptr(plusOneFloat64Int32Ptr, []*float64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapFloat64Int32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat64Int32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapFloat64Int32Ptr failed")
	}

	if len(PMapFloat64Int32Ptr(nil, []*float64{})) > 0 {
		t.Errorf("PMapFloat64Int32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int32{&vo2, &vo3, &vo4}
	newList = PMapFloat64Int32Ptr(plusOneFloat64Int32Ptr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat64Int32Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapFloat64Int16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int16 = 2
	var vo3 int16 = 3
	var vo4 int16 = 4

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3
	expectedList := []*int16{&vo2, &vo3, &vo4}
	newList := PMapFloat64Int16Ptr(plusOneFloat64Int16Ptr, []*float64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapFloat64Int16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat64Int16Ptr(nil, nil)) > 0 {
		t.Errorf("PMapFloat64Int16Ptr failed")
	}

	if len(PMapFloat64Int16Ptr(nil, []*float64{})) > 0 {
		t.Errorf("PMapFloat64Int16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int16{&vo2, &vo3, &vo4}
	newList = PMapFloat64Int16Ptr(plusOneFloat64Int16Ptr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat64Int16Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapFloat64Int8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 int8 = 2
	var vo3 int8 = 3
	var vo4 int8 = 4

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3
	expectedList := []*int8{&vo2, &vo3, &vo4}
	newList := PMapFloat64Int8Ptr(plusOneFloat64Int8Ptr, []*float64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapFloat64Int8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat64Int8Ptr(nil, nil)) > 0 {
		t.Errorf("PMapFloat64Int8Ptr failed")
	}

	if len(PMapFloat64Int8Ptr(nil, []*float64{})) > 0 {
		t.Errorf("PMapFloat64Int8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*int8{&vo2, &vo3, &vo4}
	newList = PMapFloat64Int8Ptr(plusOneFloat64Int8Ptr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat64Int8Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapFloat64UintPtr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint = 2
	var vo3 uint = 3
	var vo4 uint = 4

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3
	expectedList := []*uint{&vo2, &vo3, &vo4}
	newList := PMapFloat64UintPtr(plusOneFloat64UintPtr, []*float64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapFloat64UintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat64UintPtr(nil, nil)) > 0 {
		t.Errorf("PMapFloat64UintPtr failed")
	}

	if len(PMapFloat64UintPtr(nil, []*float64{})) > 0 {
		t.Errorf("PMapFloat64UintPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint{&vo2, &vo3, &vo4}
	newList = PMapFloat64UintPtr(plusOneFloat64UintPtr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat64UintPtr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapFloat64Uint64Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint64 = 2
	var vo3 uint64 = 3
	var vo4 uint64 = 4

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3
	expectedList := []*uint64{&vo2, &vo3, &vo4}
	newList := PMapFloat64Uint64Ptr(plusOneFloat64Uint64Ptr, []*float64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapFloat64Uint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat64Uint64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapFloat64Uint64Ptr failed")
	}

	if len(PMapFloat64Uint64Ptr(nil, []*float64{})) > 0 {
		t.Errorf("PMapFloat64Uint64Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint64{&vo2, &vo3, &vo4}
	newList = PMapFloat64Uint64Ptr(plusOneFloat64Uint64Ptr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat64Uint64Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapFloat64Uint32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint32 = 2
	var vo3 uint32 = 3
	var vo4 uint32 = 4

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3
	expectedList := []*uint32{&vo2, &vo3, &vo4}
	newList := PMapFloat64Uint32Ptr(plusOneFloat64Uint32Ptr, []*float64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapFloat64Uint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat64Uint32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapFloat64Uint32Ptr failed")
	}

	if len(PMapFloat64Uint32Ptr(nil, []*float64{})) > 0 {
		t.Errorf("PMapFloat64Uint32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint32{&vo2, &vo3, &vo4}
	newList = PMapFloat64Uint32Ptr(plusOneFloat64Uint32Ptr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat64Uint32Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapFloat64Uint16Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint16 = 2
	var vo3 uint16 = 3
	var vo4 uint16 = 4

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3
	expectedList := []*uint16{&vo2, &vo3, &vo4}
	newList := PMapFloat64Uint16Ptr(plusOneFloat64Uint16Ptr, []*float64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapFloat64Uint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat64Uint16Ptr(nil, nil)) > 0 {
		t.Errorf("PMapFloat64Uint16Ptr failed")
	}

	if len(PMapFloat64Uint16Ptr(nil, []*float64{})) > 0 {
		t.Errorf("PMapFloat64Uint16Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint16{&vo2, &vo3, &vo4}
	newList = PMapFloat64Uint16Ptr(plusOneFloat64Uint16Ptr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat64Uint16Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapFloat64Uint8Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 uint8 = 2
	var vo3 uint8 = 3
	var vo4 uint8 = 4

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3
	expectedList := []*uint8{&vo2, &vo3, &vo4}
	newList := PMapFloat64Uint8Ptr(plusOneFloat64Uint8Ptr, []*float64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapFloat64Uint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat64Uint8Ptr(nil, nil)) > 0 {
		t.Errorf("PMapFloat64Uint8Ptr failed")
	}

	if len(PMapFloat64Uint8Ptr(nil, []*float64{})) > 0 {
		t.Errorf("PMapFloat64Uint8Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*uint8{&vo2, &vo3, &vo4}
	newList = PMapFloat64Uint8Ptr(plusOneFloat64Uint8Ptr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat64Uint8Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapFloat64StrPtr(t *testing.T) {
	// Test : someLogic
	var vo10 string = "10"
	var vi10 float64 = 10
	var vi20 float64 = 20

	expectedList := []*string{&vo10}
	newList := PMapFloat64StrPtr(someLogicFloat64StrPtr, []*float64{&vi10})

	if *newList[0] != *expectedList[0] {
		t.Errorf("PMapFloat64StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat64StrPtr(nil, nil)) > 0 {
		t.Errorf("PMapFloat64StrPtr failed")
	}

	if len(PMapFloat64StrPtr(nil, []*float64{})) > 0 {
		t.Errorf("PMapFloat64StrPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*string{&vo10}
	newList = PMapFloat64StrPtr(someLogicFloat64StrPtr, []*float64{&vi10, &vi20}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != "0" {
		t.Errorf("PMapFloat64StrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapFloat64BoolPtr(t *testing.T) {
	// Test : someLogic
	var vt bool = true
	var vf bool = false

	var v10 float64 = 10
	var v0 float64 = 0

	expectedList := []*bool{&vt, &vf}
	newList := PMapFloat64BoolPtr(someLogicFloat64BoolPtr, []*float64{&v10, &v0})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] {
		t.Errorf("PMapFloat64BoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat64BoolPtr(nil, nil)) > 0 {
		t.Errorf("PMapFloat64BoolPtr failed")
	}

	if len(PMapFloat64BoolPtr(nil, []*float64{})) > 0 {
		t.Errorf("PMapFloat64BoolPtr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*bool{&vt, &vf}
	newList = PMapFloat64BoolPtr(someLogicFloat64BoolPtr, []*float64{&v10, &v0}, Optional{FixedPool: 1, RandomOrder: true})

	if *newList[0] != *expectedList[0] || *newList[1] != false {
		t.Errorf("PMapFloat64BoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestPmapFloat64Float32Ptr(t *testing.T) {
	// Test : add 1 to the list
	var vo2 float32 = 2
	var vo3 float32 = 3
	var vo4 float32 = 4

	var vi1 float64 = 1
	var vi2 float64 = 2
	var vi3 float64 = 3
	expectedList := []*float32{&vo2, &vo3, &vo4}
	newList := PMapFloat64Float32Ptr(plusOneFloat64Float32Ptr, []*float64{&vi1, &vi2, &vi3})

	if *newList[0] != *expectedList[0] || *newList[1] != *expectedList[1] || *newList[2] != *expectedList[2] {
		t.Errorf("PMapFloat64Float32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	if len(PMapFloat64Float32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapFloat64Float32Ptr failed")
	}

	if len(PMapFloat64Float32Ptr(nil, []*float64{})) > 0 {
		t.Errorf("PMapFloat64Float32Ptr failed")
	}
	reflect.TypeOf("Nandeshwar") // Leaving it here to make use of import reflect

	expectedList = []*float32{&vo2, &vo3, &vo4}
	newList = PMapFloat64Float32Ptr(plusOneFloat64Float32Ptr, []*float64{&vi1, &vi2, &vi3}, Optional{FixedPool: 2, RandomOrder: true})
	count := 0
	for _, v := range expectedList {
		for _, x := range newList {
			if *v == *x {
				count++
			}
		}
	}
	if count != len(expectedList) {
		t.Errorf("PMapFloat64Float32Ptr failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}
