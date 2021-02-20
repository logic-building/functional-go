package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestSortInt64(t *testing.T) {
	expectedList := []int64{1, 2, 3, 4, 5}
	sortedList := SortInts64([]int64{5, 1, 4, 2, 3})
	matched := reflect.DeepEqual(sortedList, expectedList)
	if !matched {
		t.Errorf("Sotredint64s failed")
	}

	sortedList = SortInts64([]int64{})
	if len(sortedList) > 0 {
		t.Errorf("Sotredint64s failed")
	}
}

func TestSortInt64Ptr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	expectedList := []*int64{&v1, &v2, &v3, &v4, &v5}
	sortedList := SortInts64Ptr([]*int64{&v5, &v1, &v4, &v2, &v3})

	for i, val := range sortedList {
		if *val != *expectedList[i] {
			for _, val := range expectedList {
				t.Errorf("expected item: %v", *val)
			}

			for _, val := range sortedList {
				t.Errorf("sorted item: %v", *val)
			}
			t.Errorf("Sotredint64s failed")
		}
	}

	sortedList = SortInts64Ptr([]*int64{})
	if len(sortedList) > 0 {
		t.Errorf("Sotredint64sPtr failed")
	}
}

func TestSortInt64Desc(t *testing.T) {
	expectedList := []int64{5, 4, 3, 2, 1}
	sortedList := SortInts64Desc([]int64{5, 1, 4, 2, 3})
	matched := reflect.DeepEqual(sortedList, expectedList)
	if !matched {
		t.Errorf("Sotredint64s failed")
	}

	sortedList = SortInts64Desc([]int64{})
	if len(sortedList) > 0 {
		t.Errorf("Sotredint64sDesc failed")
	}
}

func TestSortInt64DescPtr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5

	expectedList := []*int64{&v5, &v4, &v3, &v2, &v1}
	sortedList := SortInts64DescPtr([]*int64{&v5, &v1, &v4, &v2, &v3})

	for i, val := range sortedList {
		if *val != *expectedList[i] {
			for _, val := range expectedList {
				t.Errorf("expected item: %v", *val)
			}

			for _, val := range sortedList {
				t.Errorf("sorted item: %v", *val)
			}
			t.Errorf("Sotredint64s failed")
		}
	}

	sortedList = SortInts64DescPtr([]*int64{})
	if len(sortedList) > 0 {
		t.Errorf("Sotredint64sDescPtr failed")
	}
}

func TestSortInt32(t *testing.T) {
	expectedList := []int32{1, 2, 3, 4, 5}
	sortedList := SortInts32([]int32{5, 1, 4, 2, 3})
	matched := reflect.DeepEqual(sortedList, expectedList)
	if !matched {
		t.Errorf("Sotredint32s failed")
	}

	sortedList = SortInts32([]int32{})
	if len(sortedList) > 0 {
		t.Errorf("Sotredint32s failed")
	}
}

func TestSortInt32Ptr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	expectedList := []*int32{&v1, &v2, &v3, &v4, &v5}
	sortedList := SortInts32Ptr([]*int32{&v5, &v1, &v4, &v2, &v3})

	for i, val := range sortedList {
		if *val != *expectedList[i] {
			for _, val := range expectedList {
				t.Errorf("expected item: %v", *val)
			}

			for _, val := range sortedList {
				t.Errorf("sorted item: %v", *val)
			}
			t.Errorf("Sotredint32s failed")
		}
	}

	sortedList = SortInts32Ptr([]*int32{})
	if len(sortedList) > 0 {
		t.Errorf("Sotredint32sPtr failed")
	}
}

func TestSortInt32Desc(t *testing.T) {
	expectedList := []int32{5, 4, 3, 2, 1}
	sortedList := SortInts32Desc([]int32{5, 1, 4, 2, 3})
	matched := reflect.DeepEqual(sortedList, expectedList)
	if !matched {
		t.Errorf("Sotredint32s failed")
	}

	sortedList = SortInts32Desc([]int32{})
	if len(sortedList) > 0 {
		t.Errorf("Sotredint32sDesc failed")
	}
}

func TestSortInt32DescPtr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5

	expectedList := []*int32{&v5, &v4, &v3, &v2, &v1}
	sortedList := SortInts32DescPtr([]*int32{&v5, &v1, &v4, &v2, &v3})

	for i, val := range sortedList {
		if *val != *expectedList[i] {
			for _, val := range expectedList {
				t.Errorf("expected item: %v", *val)
			}

			for _, val := range sortedList {
				t.Errorf("sorted item: %v", *val)
			}
			t.Errorf("Sotredint32s failed")
		}
	}

	sortedList = SortInts32DescPtr([]*int32{})
	if len(sortedList) > 0 {
		t.Errorf("Sotredint32sDescPtr failed")
	}
}

func TestSortInt16(t *testing.T) {
	expectedList := []int16{1, 2, 3, 4, 5}
	sortedList := SortInts16([]int16{5, 1, 4, 2, 3})
	matched := reflect.DeepEqual(sortedList, expectedList)
	if !matched {
		t.Errorf("Sotredint16s failed")
	}

	sortedList = SortInts16([]int16{})
	if len(sortedList) > 0 {
		t.Errorf("Sotredint16s failed")
	}
}

func TestSortInt16Ptr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	expectedList := []*int16{&v1, &v2, &v3, &v4, &v5}
	sortedList := SortInts16Ptr([]*int16{&v5, &v1, &v4, &v2, &v3})

	for i, val := range sortedList {
		if *val != *expectedList[i] {
			for _, val := range expectedList {
				t.Errorf("expected item: %v", *val)
			}

			for _, val := range sortedList {
				t.Errorf("sorted item: %v", *val)
			}
			t.Errorf("Sotredint16s failed")
		}
	}

	sortedList = SortInts16Ptr([]*int16{})
	if len(sortedList) > 0 {
		t.Errorf("Sotredint16sPtr failed")
	}
}

func TestSortInt16Desc(t *testing.T) {
	expectedList := []int16{5, 4, 3, 2, 1}
	sortedList := SortInts16Desc([]int16{5, 1, 4, 2, 3})
	matched := reflect.DeepEqual(sortedList, expectedList)
	if !matched {
		t.Errorf("Sotredint16s failed")
	}

	sortedList = SortInts16Desc([]int16{})
	if len(sortedList) > 0 {
		t.Errorf("Sotredint16sDesc failed")
	}
}

func TestSortInt16DescPtr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5

	expectedList := []*int16{&v5, &v4, &v3, &v2, &v1}
	sortedList := SortInts16DescPtr([]*int16{&v5, &v1, &v4, &v2, &v3})

	for i, val := range sortedList {
		if *val != *expectedList[i] {
			for _, val := range expectedList {
				t.Errorf("expected item: %v", *val)
			}

			for _, val := range sortedList {
				t.Errorf("sorted item: %v", *val)
			}
			t.Errorf("Sotredint16s failed")
		}
	}

	sortedList = SortInts16DescPtr([]*int16{})
	if len(sortedList) > 0 {
		t.Errorf("Sotredint16sDescPtr failed")
	}
}

func TestSortInt8(t *testing.T) {
	expectedList := []int8{1, 2, 3, 4, 5}
	sortedList := SortInts8([]int8{5, 1, 4, 2, 3})
	matched := reflect.DeepEqual(sortedList, expectedList)
	if !matched {
		t.Errorf("Sotredint8s failed")
	}

	sortedList = SortInts8([]int8{})
	if len(sortedList) > 0 {
		t.Errorf("Sotredint8s failed")
	}
}

func TestSortInt8Ptr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	expectedList := []*int8{&v1, &v2, &v3, &v4, &v5}
	sortedList := SortInts8Ptr([]*int8{&v5, &v1, &v4, &v2, &v3})

	for i, val := range sortedList {
		if *val != *expectedList[i] {
			for _, val := range expectedList {
				t.Errorf("expected item: %v", *val)
			}

			for _, val := range sortedList {
				t.Errorf("sorted item: %v", *val)
			}
			t.Errorf("Sotredint8s failed")
		}
	}

	sortedList = SortInts8Ptr([]*int8{})
	if len(sortedList) > 0 {
		t.Errorf("Sotredint8sPtr failed")
	}
}

func TestSortInt8Desc(t *testing.T) {
	expectedList := []int8{5, 4, 3, 2, 1}
	sortedList := SortInts8Desc([]int8{5, 1, 4, 2, 3})
	matched := reflect.DeepEqual(sortedList, expectedList)
	if !matched {
		t.Errorf("Sotredint8s failed")
	}

	sortedList = SortInts8Desc([]int8{})
	if len(sortedList) > 0 {
		t.Errorf("Sotredint8sDesc failed")
	}
}

func TestSortInt8DescPtr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5

	expectedList := []*int8{&v5, &v4, &v3, &v2, &v1}
	sortedList := SortInts8DescPtr([]*int8{&v5, &v1, &v4, &v2, &v3})

	for i, val := range sortedList {
		if *val != *expectedList[i] {
			for _, val := range expectedList {
				t.Errorf("expected item: %v", *val)
			}

			for _, val := range sortedList {
				t.Errorf("sorted item: %v", *val)
			}
			t.Errorf("Sotredint8s failed")
		}
	}

	sortedList = SortInts8DescPtr([]*int8{})
	if len(sortedList) > 0 {
		t.Errorf("Sotredint8sDescPtr failed")
	}
}

func TestSortUint(t *testing.T) {
	expectedList := []uint{1, 2, 3, 4, 5}
	sortedList := SortUints([]uint{5, 1, 4, 2, 3})
	matched := reflect.DeepEqual(sortedList, expectedList)
	if !matched {
		t.Errorf("Sotreduints failed")
	}

	sortedList = SortUints([]uint{})
	if len(sortedList) > 0 {
		t.Errorf("Sotreduints failed")
	}
}

func TestSortUintPtr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	expectedList := []*uint{&v1, &v2, &v3, &v4, &v5}
	sortedList := SortUintsPtr([]*uint{&v5, &v1, &v4, &v2, &v3})

	for i, val := range sortedList {
		if *val != *expectedList[i] {
			for _, val := range expectedList {
				t.Errorf("expected item: %v", *val)
			}

			for _, val := range sortedList {
				t.Errorf("sorted item: %v", *val)
			}
			t.Errorf("Sotreduints failed")
		}
	}

	sortedList = SortUintsPtr([]*uint{})
	if len(sortedList) > 0 {
		t.Errorf("SotreduintsPtr failed")
	}
}

func TestSortUintDesc(t *testing.T) {
	expectedList := []uint{5, 4, 3, 2, 1}
	sortedList := SortUintsDesc([]uint{5, 1, 4, 2, 3})
	matched := reflect.DeepEqual(sortedList, expectedList)
	if !matched {
		t.Errorf("Sotreduints failed")
	}

	sortedList = SortUintsDesc([]uint{})
	if len(sortedList) > 0 {
		t.Errorf("SotreduintsDesc failed")
	}
}

func TestSortUintDescPtr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5

	expectedList := []*uint{&v5, &v4, &v3, &v2, &v1}
	sortedList := SortUintsDescPtr([]*uint{&v5, &v1, &v4, &v2, &v3})

	for i, val := range sortedList {
		if *val != *expectedList[i] {
			for _, val := range expectedList {
				t.Errorf("expected item: %v", *val)
			}

			for _, val := range sortedList {
				t.Errorf("sorted item: %v", *val)
			}
			t.Errorf("Sotreduints failed")
		}
	}

	sortedList = SortUintsDescPtr([]*uint{})
	if len(sortedList) > 0 {
		t.Errorf("SotreduintsDescPtr failed")
	}
}

func TestSortUint64(t *testing.T) {
	expectedList := []uint64{1, 2, 3, 4, 5}
	sortedList := SortUint64s([]uint64{5, 1, 4, 2, 3})
	matched := reflect.DeepEqual(sortedList, expectedList)
	if !matched {
		t.Errorf("Sotreduint64s failed")
	}

	sortedList = SortUint64s([]uint64{})
	if len(sortedList) > 0 {
		t.Errorf("Sotreduint64s failed")
	}
}

func TestSortUint64Ptr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	expectedList := []*uint64{&v1, &v2, &v3, &v4, &v5}
	sortedList := SortUint64sPtr([]*uint64{&v5, &v1, &v4, &v2, &v3})

	for i, val := range sortedList {
		if *val != *expectedList[i] {
			for _, val := range expectedList {
				t.Errorf("expected item: %v", *val)
			}

			for _, val := range sortedList {
				t.Errorf("sorted item: %v", *val)
			}
			t.Errorf("Sotreduint64s failed")
		}
	}

	sortedList = SortUint64sPtr([]*uint64{})
	if len(sortedList) > 0 {
		t.Errorf("Sotreduint64sPtr failed")
	}
}

func TestSortUint64Desc(t *testing.T) {
	expectedList := []uint64{5, 4, 3, 2, 1}
	sortedList := SortUint64sDesc([]uint64{5, 1, 4, 2, 3})
	matched := reflect.DeepEqual(sortedList, expectedList)
	if !matched {
		t.Errorf("Sotreduint64s failed")
	}

	sortedList = SortUint64sDesc([]uint64{})
	if len(sortedList) > 0 {
		t.Errorf("Sotreduint64sDesc failed")
	}
}

func TestSortUint64DescPtr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5

	expectedList := []*uint64{&v5, &v4, &v3, &v2, &v1}
	sortedList := SortUint64sDescPtr([]*uint64{&v5, &v1, &v4, &v2, &v3})

	for i, val := range sortedList {
		if *val != *expectedList[i] {
			for _, val := range expectedList {
				t.Errorf("expected item: %v", *val)
			}

			for _, val := range sortedList {
				t.Errorf("sorted item: %v", *val)
			}
			t.Errorf("Sotreduint64s failed")
		}
	}

	sortedList = SortUint64sDescPtr([]*uint64{})
	if len(sortedList) > 0 {
		t.Errorf("Sotreduint64sDescPtr failed")
	}
}

func TestSortUint32(t *testing.T) {
	expectedList := []uint32{1, 2, 3, 4, 5}
	sortedList := SortUints32([]uint32{5, 1, 4, 2, 3})
	matched := reflect.DeepEqual(sortedList, expectedList)
	if !matched {
		t.Errorf("Sotreduint32s failed")
	}

	sortedList = SortUints32([]uint32{})
	if len(sortedList) > 0 {
		t.Errorf("Sotreduint32s failed")
	}
}

func TestSortUint32Ptr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	expectedList := []*uint32{&v1, &v2, &v3, &v4, &v5}
	sortedList := SortUints32Ptr([]*uint32{&v5, &v1, &v4, &v2, &v3})

	for i, val := range sortedList {
		if *val != *expectedList[i] {
			for _, val := range expectedList {
				t.Errorf("expected item: %v", *val)
			}

			for _, val := range sortedList {
				t.Errorf("sorted item: %v", *val)
			}
			t.Errorf("Sotreduint32s failed")
		}
	}

	sortedList = SortUints32Ptr([]*uint32{})
	if len(sortedList) > 0 {
		t.Errorf("Sotreduint32sPtr failed")
	}
}

func TestSortUint32Desc(t *testing.T) {
	expectedList := []uint32{5, 4, 3, 2, 1}
	sortedList := SortUints32Desc([]uint32{5, 1, 4, 2, 3})
	matched := reflect.DeepEqual(sortedList, expectedList)
	if !matched {
		t.Errorf("Sotreduint32s failed")
	}

	sortedList = SortUints32Desc([]uint32{})
	if len(sortedList) > 0 {
		t.Errorf("Sotreduint32sDesc failed")
	}
}

func TestSortUint32DescPtr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5

	expectedList := []*uint32{&v5, &v4, &v3, &v2, &v1}
	sortedList := SortUints32DescPtr([]*uint32{&v5, &v1, &v4, &v2, &v3})

	for i, val := range sortedList {
		if *val != *expectedList[i] {
			for _, val := range expectedList {
				t.Errorf("expected item: %v", *val)
			}

			for _, val := range sortedList {
				t.Errorf("sorted item: %v", *val)
			}
			t.Errorf("Sotreduint32s failed")
		}
	}

	sortedList = SortUints32DescPtr([]*uint32{})
	if len(sortedList) > 0 {
		t.Errorf("Sotreduint32sDescPtr failed")
	}
}

func TestSortUint16(t *testing.T) {
	expectedList := []uint16{1, 2, 3, 4, 5}
	sortedList := SortUints16([]uint16{5, 1, 4, 2, 3})
	matched := reflect.DeepEqual(sortedList, expectedList)
	if !matched {
		t.Errorf("Sotreduint16s failed")
	}

	sortedList = SortUints16([]uint16{})
	if len(sortedList) > 0 {
		t.Errorf("Sotreduint16s failed")
	}
}

func TestSortUint16Ptr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	expectedList := []*uint16{&v1, &v2, &v3, &v4, &v5}
	sortedList := SortUints16Ptr([]*uint16{&v5, &v1, &v4, &v2, &v3})

	for i, val := range sortedList {
		if *val != *expectedList[i] {
			for _, val := range expectedList {
				t.Errorf("expected item: %v", *val)
			}

			for _, val := range sortedList {
				t.Errorf("sorted item: %v", *val)
			}
			t.Errorf("Sotreduint16s failed")
		}
	}

	sortedList = SortUints16Ptr([]*uint16{})
	if len(sortedList) > 0 {
		t.Errorf("Sotreduint16sPtr failed")
	}
}

func TestSortUint16Desc(t *testing.T) {
	expectedList := []uint16{5, 4, 3, 2, 1}
	sortedList := SortUints16Desc([]uint16{5, 1, 4, 2, 3})
	matched := reflect.DeepEqual(sortedList, expectedList)
	if !matched {
		t.Errorf("Sotreduint16s failed")
	}

	sortedList = SortUints16Desc([]uint16{})
	if len(sortedList) > 0 {
		t.Errorf("Sotreduint16sDesc failed")
	}
}

func TestSortUint16DescPtr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5

	expectedList := []*uint16{&v5, &v4, &v3, &v2, &v1}
	sortedList := SortUints16DescPtr([]*uint16{&v5, &v1, &v4, &v2, &v3})

	for i, val := range sortedList {
		if *val != *expectedList[i] {
			for _, val := range expectedList {
				t.Errorf("expected item: %v", *val)
			}

			for _, val := range sortedList {
				t.Errorf("sorted item: %v", *val)
			}
			t.Errorf("Sotreduint16s failed")
		}
	}

	sortedList = SortUints16DescPtr([]*uint16{})
	if len(sortedList) > 0 {
		t.Errorf("Sotreduint16sDescPtr failed")
	}
}

func TestSortUint8(t *testing.T) {
	expectedList := []uint8{1, 2, 3, 4, 5}
	sortedList := SortUints8([]uint8{5, 1, 4, 2, 3})
	matched := reflect.DeepEqual(sortedList, expectedList)
	if !matched {
		t.Errorf("Sotreduint8s failed")
	}

	sortedList = SortUints8([]uint8{})
	if len(sortedList) > 0 {
		t.Errorf("Sotreduint8s failed")
	}
}

func TestSortUint8Ptr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	expectedList := []*uint8{&v1, &v2, &v3, &v4, &v5}
	sortedList := SortUints8Ptr([]*uint8{&v5, &v1, &v4, &v2, &v3})

	for i, val := range sortedList {
		if *val != *expectedList[i] {
			for _, val := range expectedList {
				t.Errorf("expected item: %v", *val)
			}

			for _, val := range sortedList {
				t.Errorf("sorted item: %v", *val)
			}
			t.Errorf("Sotreduint8s failed")
		}
	}

	sortedList = SortUints8Ptr([]*uint8{})
	if len(sortedList) > 0 {
		t.Errorf("Sotreduint8sPtr failed")
	}
}

func TestSortUint8Desc(t *testing.T) {
	expectedList := []uint8{5, 4, 3, 2, 1}
	sortedList := SortUints8Desc([]uint8{5, 1, 4, 2, 3})
	matched := reflect.DeepEqual(sortedList, expectedList)
	if !matched {
		t.Errorf("Sotreduint8s failed")
	}

	sortedList = SortUints8Desc([]uint8{})
	if len(sortedList) > 0 {
		t.Errorf("Sotreduint8sDesc failed")
	}
}

func TestSortUint8DescPtr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5

	expectedList := []*uint8{&v5, &v4, &v3, &v2, &v1}
	sortedList := SortUints8DescPtr([]*uint8{&v5, &v1, &v4, &v2, &v3})

	for i, val := range sortedList {
		if *val != *expectedList[i] {
			for _, val := range expectedList {
				t.Errorf("expected item: %v", *val)
			}

			for _, val := range sortedList {
				t.Errorf("sorted item: %v", *val)
			}
			t.Errorf("Sotreduint8s failed")
		}
	}

	sortedList = SortUints8DescPtr([]*uint8{})
	if len(sortedList) > 0 {
		t.Errorf("Sotreduint8sDescPtr failed")
	}
}

func TestSortFloat32(t *testing.T) {
	expectedList := []float32{1, 2, 3, 4, 5}
	sortedList := SortFloat32s([]float32{5, 1, 4, 2, 3})
	matched := reflect.DeepEqual(sortedList, expectedList)
	if !matched {
		t.Errorf("Sotredfloat32s failed")
	}

	sortedList = SortFloat32s([]float32{})
	if len(sortedList) > 0 {
		t.Errorf("Sotredfloat32s failed")
	}
}

func TestSortFloat32Ptr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	expectedList := []*float32{&v1, &v2, &v3, &v4, &v5}
	sortedList := SortFloat32sPtr([]*float32{&v5, &v1, &v4, &v2, &v3})

	for i, val := range sortedList {
		if *val != *expectedList[i] {
			for _, val := range expectedList {
				t.Errorf("expected item: %v", *val)
			}

			for _, val := range sortedList {
				t.Errorf("sorted item: %v", *val)
			}
			t.Errorf("Sotredfloat32s failed")
		}
	}

	sortedList = SortFloat32sPtr([]*float32{})
	if len(sortedList) > 0 {
		t.Errorf("Sotredfloat32sPtr failed")
	}
}

func TestSortFloat32Desc(t *testing.T) {
	expectedList := []float32{5, 4, 3, 2, 1}
	sortedList := SortFloat32sDesc([]float32{5, 1, 4, 2, 3})
	matched := reflect.DeepEqual(sortedList, expectedList)
	if !matched {
		t.Errorf("Sotredfloat32s failed")
	}

	sortedList = SortFloat32sDesc([]float32{})
	if len(sortedList) > 0 {
		t.Errorf("Sotredfloat32sDesc failed")
	}
}

func TestSortFloat32DescPtr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5

	expectedList := []*float32{&v5, &v4, &v3, &v2, &v1}
	sortedList := SortFloat32sDescPtr([]*float32{&v5, &v1, &v4, &v2, &v3})

	for i, val := range sortedList {
		if *val != *expectedList[i] {
			for _, val := range expectedList {
				t.Errorf("expected item: %v", *val)
			}

			for _, val := range sortedList {
				t.Errorf("sorted item: %v", *val)
			}
			t.Errorf("Sotredfloat32s failed")
		}
	}

	sortedList = SortFloat32sDescPtr([]*float32{})
	if len(sortedList) > 0 {
		t.Errorf("Sotredfloat32sDescPtr failed")
	}
}
