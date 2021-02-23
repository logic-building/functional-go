package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/logic-building/functional-go/internal/template/basic"
	"github.com/logic-building/functional-go/internal/template/methodchain"
)

type fpCode struct {
	function         string
	codeTemplate     string
	testTemplate     string
	testTemplateBool string

	// Test template for different combination of data type - map, filter,
	testTemplateIONumber     string
	testTemplateIOStrNumber  string
	testTemplateIONumberStr  string
	testTemplateIONumberBool string
	testTemplateIOStrBool    string
	testTemplateIOBoolNumber string
	testTemplateIOBoolStr    string

	// Test template for merge
	testTemplateNumberStr  string
	testTemplateStrNumber  string
	testTemplateStrBool    string
	testTemplateBoolStr    string
	testTemplateNumberBool string
	testTemplateBoolNumber string
	testTemplateBoolBool   string
	testTemplateStrStr     string

	dataTypes             []string
	generatedFileName     string
	generatedTestFileName string
}

var fpCodeList = []fpCode{
	fpCode{
		function:          "methodchain",
		codeTemplate:      methodchain.MethodChainBasic(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "methodchain.go",

		testTemplate:          methodchain.MethodChainMapTest(),
		testTemplateBool:      methodchain.MethodChainMapBoolTest(),
		generatedTestFileName: "methodchain_test.go",
	},

	fpCode{
		function:          "DropLast",
		codeTemplate:      basic.DropLast(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "droplast.go",

		testTemplate:          basic.DropLastTest(),
		testTemplateBool:      basic.DropLastBoolTest(),
		generatedTestFileName: "droplast_test.go",
	},

	fpCode{
		function:          "Dedupe",
		codeTemplate:      basic.Dedupe(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "dedupe.go",

		testTemplate:          basic.DedupeTest(),
		testTemplateBool:      basic.DedupeBoolTest(),
		generatedTestFileName: "dedupe_test.go",
	},

	fpCode{
		function:          "Set",
		codeTemplate:      basic.Set() + basic.SetPtr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "set.go",

		testTemplate:          basic.SetTest(),
		testTemplateBool:      basic.SetBoolTest(),
		generatedTestFileName: "set_test.go",
	},

	fpCode{
		function:          "Union",
		codeTemplate:      basic.Union() + basic.UnionPtr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "union.go",

		testTemplate:          basic.UnionTest(),
		testTemplateBool:      basic.UnionBoolTest(),
		generatedTestFileName: "union_test.go",
	},

	fpCode{
		function:          "Intersection",
		codeTemplate:      basic.Intersection() + basic.IntersectionPtr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "intersection.go",

		testTemplate:          basic.IntersectionTest(),
		testTemplateBool:      basic.IntersectionBoolTest(),
		generatedTestFileName: "intersection_test.go",
	},

	fpCode{
		function:          "Difference",
		codeTemplate:      basic.Difference() + basic.DifferencePtr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "difference.go",

		testTemplate:          basic.DifferenceTest(),
		testTemplateBool:      basic.DifferenceBoolTest(),
		generatedTestFileName: "difference_test.go",
	},

	fpCode{
		function:          "Subset",
		codeTemplate:      basic.Subset() + basic.SubsetPtr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "subset.go",

		testTemplate:          basic.SubsetTest(),
		testTemplateBool:      basic.SubsetBoolTest(),
		generatedTestFileName: "subset_test.go",
	},

	fpCode{
		function:          "Superset",
		codeTemplate:      basic.Superset() + basic.SupersetPtr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "superset.go",

		testTemplate:          basic.SupersetTest(),
		testTemplateBool:      basic.SupersetBoolTest(),
		generatedTestFileName: "superset_test.go",
	},

	fpCode{
		function:          "Equal",
		codeTemplate:      basic.Equal(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "equal.go",

		testTemplate:          basic.EqualTest(),
		testTemplateBool:      basic.EqualBoolTest(),
		generatedTestFileName: "equal_test.go",
	},

	fpCode{
		function:          "EqualMapType1Type2",
		codeTemplate:      basic.EqualMapType1Type2(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "equalmap.go",

		testTemplate:           basic.EqualMapsTest(),
		testTemplateNumberStr:  basic.EqualMapsNumberToStringTest(),
		testTemplateStrNumber:  basic.EqualMapsStringToNumberTest(),
		testTemplateStrBool:    basic.EqualMapsStringToBoolTest(),
		testTemplateBoolStr:    basic.EqualMapsBoolToStringTest(),
		testTemplateNumberBool: basic.EqualMapsNumberToBoolTest(),
		testTemplateBoolNumber: basic.EqualMapsBoolToNumberTest(),
		testTemplateBoolBool:   basic.EqualMapsBoolTest(),
		testTemplateStrStr:     basic.EqualMapsStringToStringTest(),
		generatedTestFileName:  "equalmap_test.go",
	},

	fpCode{
		function:          "Zero",
		codeTemplate:      basic.ZeroP(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "float32", "float64"},
		generatedFileName: "zero.go",

		testTemplate: basic.ZeroPTest(),
		//testTemplateBool:      basic.SupersetBoolTest(),
		generatedTestFileName: "zero_test.go",
	},

	fpCode{
		function:          "PosP",
		codeTemplate:      basic.PosP(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "float32", "float64"},
		generatedFileName: "pos.go",

		testTemplate: basic.PosPTest(),
		//testTemplateBool:      basic.SupersetBoolTest(),
		generatedTestFileName: "pos_test.go",
	},

	fpCode{
		function:          "NegP",
		codeTemplate:      basic.NegP(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "float32", "float64"},
		generatedFileName: "neg.go",

		testTemplate: basic.NegPTest(),
		//testTemplateBool:      basic.SupersetBoolTest(),
		generatedTestFileName: "neg_test.go",
	},

	fpCode{
		function:          "EvenP",
		codeTemplate:      basic.EvenP(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8"},
		generatedFileName: "even.go",

		testTemplate: basic.EvenPTest(),
		//testTemplateBool:      basic.SupersetBoolTest(),
		generatedTestFileName: "even_test.go",
	},

	fpCode{
		function:          "OddP",
		codeTemplate:      basic.OddP(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8"},
		generatedFileName: "odd.go",

		testTemplate: basic.OddPTest(),
		//testTemplateBool:      basic.SupersetBoolTest(),
		generatedTestFileName: "odd_test.go",
	},

	fpCode{
		function:          "Take",
		codeTemplate:      basic.Take() + basic.TakePtr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "take.go",

		testTemplate:          basic.TakeTest(),
		testTemplateBool:      basic.TakeBoolTest(),
		generatedTestFileName: "take_test.go",
	},

	fpCode{
		function:          "Reverse",
		codeTemplate:      basic.Reverse() + basic.ReversePtr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "reverse.go",

		testTemplate:          basic.ReverseTest(),
		testTemplateBool:      basic.ReverseBoolTest(),
		generatedTestFileName: "reverse_test.go",
	},

	fpCode{
		function:          "SortInts",
		codeTemplate:      basic.SortInts(),
		dataTypes:         []string{"int"},
		generatedFileName: "sortints.go",

		testTemplate: basic.SortIntsTest(),
		//testTemplateBool:      basic.DropLastBoolTest(),
		generatedTestFileName: "sortints_test.go",
	},

	fpCode{
		function:          "SortFloats",
		codeTemplate:      basic.SortFloats64(),
		dataTypes:         []string{"float64"},
		generatedFileName: "sortfloats.go",

		testTemplate: basic.SortFloats64Test(),
		//testTemplateBool:      basic.DropLastBoolTest(),
		generatedTestFileName: "sortfloats_test.go",
	},

	fpCode{
		function:          "SortStrings",
		codeTemplate:      basic.SortStrs(),
		dataTypes:         []string{"string"},
		generatedFileName: "sortstrs.go",

		testTemplate: basic.SortStrsTest(),
		//testTemplateBool:      basic.DropLastBoolTest(),
		generatedTestFileName: "sortstrs_test.go",
	},

	fpCode{
		function:          "Sort",
		codeTemplate:      basic.Sort(),
		dataTypes:         []string{"int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "float32"},
		generatedFileName: "sort.go",

		testTemplate: basic.SortTest(),
		//testTemplateBool:      basic.DropLastBoolTest(),
		generatedTestFileName: "sort_test.go",
	},

	fpCode{
		function:          "ReducePtr",
		codeTemplate:      basic.ReducePtr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "float32", "float64"},
		generatedFileName: "reduceptr.go",

		testTemplate: basic.ReducePtrTest(),
		//testTemplateBool:      basic.DropLastBoolTest(), // Not required here
		generatedTestFileName: "reduceptr_test.go",
	},

	fpCode{
		function:          "ReducePtrErr",
		codeTemplate:      basic.ReducePtrErr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "float32", "float64"},
		generatedFileName: "reduceptrerr.go",

		testTemplate: basic.ReducePtrErrTest(),
		//testTemplateBool:      basic.DropLastBoolTest(), // Not required here
		generatedTestFileName: "reduceptrerr_test.go",
	},

	fpCode{
		function:          "ReduceErr",
		codeTemplate:      basic.ReduceErr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "float32", "float64"},
		generatedFileName: "reduceerr.go",

		testTemplate: basic.ReduceErrTest(),
		//testTemplateBool:      basic.DropLastBoolTest(), // Not required here
		generatedTestFileName: "reduceerr_test.go",
	},

	fpCode{
		function:          "SomePtr",
		codeTemplate:      basic.SomePtr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "someptr.go",

		testTemplate:          basic.SomePtrTest(),
		testTemplateBool:      basic.SomePtrTestBool(),
		generatedTestFileName: "someptr_test.go",
	},

	fpCode{
		function:          "SomePtrErr",
		codeTemplate:      basic.SomePtrErr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "someptrerr.go",

		testTemplate:          basic.SomePtrErrTest(),
		testTemplateBool:      basic.SomePtrErrTestBool(),
		generatedTestFileName: "someptrerr_test.go",
	},

	fpCode{
		function:          "SomeErr",
		codeTemplate:      basic.SomeErr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "someerr.go",

		testTemplate:          basic.SomeErrTest(),
		testTemplateBool:      basic.SomeErrTestBool(),
		generatedTestFileName: "someerr_test.go",
	},

	fpCode{
		function:          "TakeWhilePtr",
		codeTemplate:      basic.TakeWhilePtr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "takewhileptr.go",

		testTemplate:          basic.TakeWhilePtrTest(),
		testTemplateBool:      basic.TakeWhileTestBool(),
		generatedTestFileName: "takewhileptr_test.go",
	},

	fpCode{
		function:          "TakeWhilePtrErr",
		codeTemplate:      basic.TakeWhilePtrErr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "takewhileptrerr.go",

		testTemplate:          basic.TakeWhilePtrErrTest(),
		testTemplateBool:      basic.TakeWhilePtrErrTestBool(),
		generatedTestFileName: "takewhileptrerr_test.go",
	},

	fpCode{
		function:          "TakeWhileErr",
		codeTemplate:      basic.TakeWhileErr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "takewhileerr.go",

		testTemplate:          basic.TakeWhileErrTest(),
		testTemplateBool:      basic.TakeWhileErrTestBool(),
		generatedTestFileName: "takewhileerr_test.go",
	},

	fpCode{
		function:          "RemovePtr",
		codeTemplate:      basic.RemovePtr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "removeptr.go",

		testTemplate:          basic.RemovePtrTest(),
		testTemplateBool:      basic.RemovePtrTestBool(),
		generatedTestFileName: "removeptr_test.go",
	},

	fpCode{
		function:          "RemovePtrErr",
		codeTemplate:      basic.RemovePtrErr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "removeptrerr.go",

		testTemplate:          basic.RemovePtrErrTest(),
		testTemplateBool:      basic.RemovePtrErrTestBool(),
		generatedTestFileName: "removeptrerr_test.go",
	},

	fpCode{
		function:          "RemoveErr",
		codeTemplate:      basic.RemoveErr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "removeerr.go",

		testTemplate:          basic.RemoveErrTest(),
		testTemplateBool:      basic.RemoveErrTestBool(),
		generatedTestFileName: "removeerr_test.go",
	},

	fpCode{
		function:          "RestPtr",
		codeTemplate:      basic.RestPtr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "restptr.go",

		testTemplate:          basic.RestPtrTest(),
		testTemplateBool:      basic.RestPtrTestBool(),
		generatedTestFileName: "restptr_test.go",
	},

	fpCode{
		function:          "DropLastPtr",
		codeTemplate:      basic.DropLastPtr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "droplastPtr.go",

		testTemplate:          basic.DropLastPtrTest(),
		testTemplateBool:      basic.DropLastPtrBoolTest(),
		generatedTestFileName: "droplastPtr_test.go",
	},

	fpCode{
		function:          "MapPtr",
		codeTemplate:      basic.MapPtr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "mapPtr.go",

		testTemplate:          basic.MapPtrTest(),
		testTemplateBool:      basic.MapPtrBoolTest(),
		generatedTestFileName: "mapPtr_test.go",
	},

	fpCode{
		function:          "MapPtrErr",
		codeTemplate:      basic.MapPtrErr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "mapPtrErr.go",

		testTemplate:          basic.MapPtrErrTest(),
		testTemplateBool:      basic.MapPtrErrBoolTest(),
		generatedTestFileName: "mapPtrErr_test.go",
	},

	fpCode{
		function:          "MapErr",
		codeTemplate:      basic.MapErr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "mapErr.go",

		testTemplate:          basic.MapErrTest(),
		testTemplateBool:      basic.MapErrBoolTest(),
		generatedTestFileName: "mapErr_test.go",
	},

	fpCode{
		function:          "PMap",
		codeTemplate:      basic.PMap(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "pmap.go",

		// Not required because there is a manually generated file
		testTemplate:          basic.PMap2Test(),
		testTemplateBool:      basic.PMap2BoolTest(),
		generatedTestFileName: "pmap2_test.go",
	},

	fpCode{
		function:          "PMapPtr",
		codeTemplate:      basic.PMapPtr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "pmapPtr.go",

		testTemplate:          basic.PMapPtrTest() + basic.PMapPtr2Test(),
		testTemplateBool:      basic.PMapPtrBoolTest() + basic.PMap2PtrBoolTest(),
		generatedTestFileName: "pmapPtr_test.go",
	},

	fpCode{
		function:          "PMapPtrErr",
		codeTemplate:      basic.PMapPtrErr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "pmapPtrErr.go",

		testTemplate:          basic.PMapPtrErrTest() + basic.PMapPtr2ErrTest(),
		testTemplateBool:      basic.PMapPtrErrBoolTest() + basic.PMapPtr2ErrBoolTest(),
		generatedTestFileName: "pmapPtrErr_test.go",
	},

	fpCode{
		function:          "PMapErr",
		codeTemplate:      basic.PMapErr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "pmapErr.go",

		testTemplate:          basic.PMapErrTest() + basic.PMap2ErrTest(),
		testTemplateBool:      basic.PMapErrBoolTest() + basic.PMap2ErrBoolTest(),
		generatedTestFileName: "pmapErr_test.go",
	},

	fpCode{
		function:          "FilterPtr",
		codeTemplate:      basic.FilterPtr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "filterPtr.go",

		testTemplate:          basic.FilterPtrTest(),
		testTemplateBool:      basic.FilterPtrBoolTest(),
		generatedTestFileName: "filterPtr_test.go",
	},

	fpCode{
		function:          "FilterPtrErr",
		codeTemplate:      basic.FilterPtrErr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "filterPtrErr.go",

		testTemplate:          basic.FilterPtrErrTest(),
		testTemplateBool:      basic.FilterPtrErrBoolTest(),
		generatedTestFileName: "filterPtrErr_test.go",
	},

	fpCode{
		function:          "FilterErr",
		codeTemplate:      basic.FilterErr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "filterErr.go",

		testTemplate:          basic.FilterErrTest(),
		testTemplateBool:      basic.FilterErrBoolTest(),
		generatedTestFileName: "filterErr_test.go",
	},

	fpCode{
		function:          "DistinctPtr",
		codeTemplate:      basic.DistinctPtr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "distinctPtr.go",

		testTemplate:          basic.DistinctPtrTest(),
		testTemplateBool:      basic.DistinctPtrBoolTest(),
		generatedTestFileName: "distinctPtr_test.go",
	},

	fpCode{
		function:          "DistinctP",
		codeTemplate:      basic.DistinctP() + basic.DistinctPPtr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "distinctp.go",

		testTemplate:          basic.DistinctPTest(),
		testTemplateBool:      basic.DistinctPBoolTest(),
		generatedTestFileName: "distinctp_test.go",
	},

	fpCode{
		function:          "FilterMapPtr",
		codeTemplate:      basic.FilterMapPtr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "filtermapptr.go",

		testTemplate:          basic.FilterMapPtrTest(),
		testTemplateBool:      basic.FilterMapPtrBoolTest(),
		generatedTestFileName: "filtermapptr_test.go",
	},

	fpCode{
		function:          "FilterMapPtrErr",
		codeTemplate:      basic.FilterMapPtrErr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "float32", "float64"},
		generatedFileName: "filtermapptrerr.go",

		testTemplate: basic.FilterMapPtrErrTest(),
		//testTemplateBool:     bool is removed above so, bool function won't be generated.
		generatedTestFileName: "filtermapptrerr_test.go",
	},

	fpCode{
		function:          "FilterMapErr",
		codeTemplate:      basic.FilterMapErr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "float32", "float64"},
		generatedFileName: "filtermaperr.go",

		testTemplate: basic.FilterMapErrTest(),
		//testTemplateBool:     bool is removed above so, bool function won't be generated.
		generatedTestFileName: "filtermaperr_test.go",
	},

	fpCode{
		function:          "ExistsPtr",
		codeTemplate:      basic.ExistsPtr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "existsptr.go",

		testTemplate: basic.ExistsPtrTest(),
		//testTemplateBool:      basic.FilterMapPtrBoolTest(),
		generatedTestFileName: "existsptr_test.go",
	},

	fpCode{
		function:          "DropPtr",
		codeTemplate:      basic.DropPtr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "dropptr.go",

		testTemplate:          basic.DropPtrTest(),
		testTemplateBool:      basic.DropPtrBoolTest(),
		generatedTestFileName: "dropptr_test.go",
	},

	fpCode{
		function:          "DropWhilePtr",
		codeTemplate:      basic.DropWhilePtr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "dropwhileptr.go",

		testTemplate:          basic.DropWhilePtrTest(),
		testTemplateBool:      basic.DropWhilePtrBoolTest(),
		generatedTestFileName: "dropwhileptr_test.go",
	},

	fpCode{
		function:          "DropWhilePtrErr",
		codeTemplate:      basic.DropWhilePtrErr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "dropwhileptrerr.go",

		testTemplate:          basic.DropWhilePtrErrTest(),
		testTemplateBool:      basic.DropWhilePtrErrBoolTest(),
		generatedTestFileName: "dropwhileptrerr_test.go",
	},

	fpCode{
		function:          "DropWhileErr",
		codeTemplate:      basic.DropWhileErr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "dropwhileerr.go",

		testTemplate:          basic.DropWhileErrTest(),
		testTemplateBool:      basic.DropWhileErrBoolTest(),
		generatedTestFileName: "dropwhileerr_test.go",
	},

	fpCode{
		function:          "EveryPtr",
		codeTemplate:      basic.EveryPtr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "everyptr.go",

		testTemplate:          basic.EveryPtrTest(),
		testTemplateBool:      basic.EveryPtrBoolTest(),
		generatedTestFileName: "everyptr_test.go",
	},

	fpCode{
		function:          "EveryPtrErr",
		codeTemplate:      basic.EveryPtrErr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "everyptrerr.go",

		testTemplate:          basic.EveryPtrErrTest(),
		testTemplateBool:      basic.EveryPtrErrBoolTest(),
		generatedTestFileName: "everyptrerr_test.go",
	},

	fpCode{
		function:          "EveryErr",
		codeTemplate:      basic.EveryErr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "everyerr.go",

		testTemplate:          basic.EveryErrTest(),
		testTemplateBool:      basic.EveryErrBoolTest(),
		generatedTestFileName: "everyerr_test.go",
	},

	fpCode{
		function:          "MaxPtr",
		codeTemplate:      basic.MaxPtr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "float32", "float64"},
		generatedFileName: "maxptr.go",

		testTemplate: basic.MaxPtrTest(),
		//testTemplateBool:      basic.EveryPtrBoolTest(), // bool test is not required here
		generatedTestFileName: "maxptr_test.go",
	},

	fpCode{
		function:          "MinPtr",
		codeTemplate:      basic.MinPtr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "float32", "float64"},
		generatedFileName: "minptr.go",

		testTemplate: basic.MinPtrTest(),
		//testTemplateBool:      basic.EveryPtrBoolTest(), // bool test is not required here
		generatedTestFileName: "minptr_test.go",
	},

	fpCode{
		function:          "MinMaxPtr",
		codeTemplate:      basic.MinMaxPtr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "float32", "float64"},
		generatedFileName: "minmaxptr.go",

		testTemplate: basic.MinMaxPtrTest(),
		//testTemplateBool:      basic.EveryPtrBoolTest(), // bool test is not required here
		generatedTestFileName: "minmaxptr_test.go",
	},

	fpCode{
		function:          "Merge",
		codeTemplate:      basic.Merge(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "merge.go",

		testTemplate:           basic.MergeTest(),
		testTemplateNumberStr:  basic.MergeTestNumbersToString(),
		testTemplateStrNumber:  basic.MergeTestStringToNumbers(),
		testTemplateStrBool:    basic.MergeTestStringToBool(),
		testTemplateBoolStr:    basic.MergeTestBoolToString(),
		testTemplateNumberBool: basic.MergeTestNumberToBool(),
		testTemplateBoolNumber: basic.MergeTestBoolToNumber(),
		testTemplateBoolBool:   basic.MergeTestBoolToBool(),
		testTemplateStrStr:     basic.MergeTestStrToStr(),
		generatedTestFileName:  "merge_test.go",
	},

	fpCode{
		function:          "MergePtr",
		codeTemplate:      basic.MergePtr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "mergeptr.go",

		testTemplate:           basic.MergeTestPtr(),
		testTemplateNumberStr:  basic.MergeTestNumbersToStringPtr(),
		testTemplateStrNumber:  basic.MergeTestStringToNumbersPtr(),
		testTemplateStrBool:    basic.MergeTestStringToBoolPtr(),
		testTemplateBoolStr:    basic.MergeTestBoolToStringPtr(),
		testTemplateNumberBool: basic.MergeTestNumberToBoolPtr(),
		testTemplateBoolNumber: basic.MergeTestBoolToNumberPtr(),
		testTemplateBoolBool:   basic.MergeTestBoolToBoolPtr(),
		testTemplateStrStr:     basic.MergeTestStrToStrPtr(),
		generatedTestFileName:  "mergeptr_test.go",
	},

	fpCode{
		function:          "Zip",
		codeTemplate:      basic.Zip(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "zip.go",

		testTemplate:           basic.ZIPNumberToNumber(),
		testTemplateNumberStr:  basic.ZIPNumberToStr(),
		testTemplateStrNumber:  basic.ZIPStrToNumber(),
		testTemplateStrBool:    basic.ZIPStrToBool(),
		testTemplateBoolStr:    basic.ZIPBoolToStr(),
		testTemplateNumberBool: basic.ZIPNumberToBool(),
		testTemplateBoolNumber: basic.ZIPBoolToNumber(),
		testTemplateBoolBool:   basic.ZIPBoolToBool(),
		testTemplateStrStr:     basic.ZIPStrToStr(),
		generatedTestFileName:  "zip_test.go",
	},

	fpCode{
		function:          "ZipPtr",
		codeTemplate:      basic.ZipPtr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "zipptr.go",

		testTemplate:           basic.ZIPNumberToNumberPtr(),
		testTemplateNumberStr:  basic.ZIPNumberToStrPtr(),
		testTemplateStrNumber:  basic.ZIPStrToNumberPtr(),
		testTemplateStrBool:    basic.ZIPStrToBoolPtr(),
		testTemplateBoolStr:    basic.ZIPBoolToStrPtr(),
		testTemplateNumberBool: basic.ZIPNumberToBoolPtr(),
		testTemplateBoolNumber: basic.ZIPBoolToNumberPtr(),
		testTemplateBoolBool:   basic.ZIPBoolToBoolPtr(),
		testTemplateStrStr:     basic.ZIPStrToStrPtr(),
		generatedTestFileName:  "zipptr_test.go",
	},

	fpCode{
		function:          "Keys",
		codeTemplate:      basic.Keys(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "keys.go",

		testTemplate:           basic.KeysNumberToNumberTest(),
		testTemplateNumberStr:  basic.KeysNumberToStrTest(),
		testTemplateStrNumber:  basic.KeysStrToNumberTest(),
		testTemplateStrBool:    basic.KeysStrToBoolTest(),
		testTemplateBoolStr:    basic.KeysBoolToStrTest(),
		testTemplateNumberBool: basic.KeysNumberToBoolTest(),
		testTemplateBoolNumber: basic.KeysBoolToNumberTest(),
		testTemplateBoolBool:   basic.KeysBoolToBoolTest(),
		testTemplateStrStr:     basic.KeysStrToStrTest(),
		generatedTestFileName:  "keys_test.go",
	},

	fpCode{
		function:                 "PMapIO",
		codeTemplate:             basic.PMapIO(),
		testTemplateIONumber:     basic.PMapIONumber(),
		testTemplateIOStrNumber:  basic.PMapIOStrNumber(),
		testTemplateIONumberStr:  basic.PMapIONumberStr(),
		testTemplateIONumberBool: basic.PMapIONumberBool(),
		testTemplateIOStrBool:    basic.PMapIOStrBool(),
		testTemplateIOBoolNumber: basic.PMapIOBoolNumber(),
		testTemplateIOBoolStr:    basic.PMapIOBoolStr(),
		dataTypes:                []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName:        "pmapio.go",
		generatedTestFileName:    "pmapio_test.go",
	},

	fpCode{
		function:                 "PMapIOErr",
		codeTemplate:             basic.PMapIOErr(),
		testTemplateIONumber:     basic.PMapIONumberErr(),
		testTemplateIOStrNumber:  basic.PMapIOStrNumberErr(),
		testTemplateIONumberStr:  basic.PMapIONumberStrErr(),
		testTemplateIONumberBool: basic.PMapIONumberBoolErr(),
		testTemplateIOStrBool:    basic.PMapIOStrBoolErr(),
		testTemplateIOBoolNumber: basic.PMapIOBoolNumberErr(),
		testTemplateIOBoolStr:    basic.PMapIOBoolStrErr(),
		dataTypes:                []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName:        "pmapioerr.go",
		generatedTestFileName:    "pmapioerr_test.go",
	},

	fpCode{
		function:                 "PMapIOPtr",
		codeTemplate:             basic.PMapIOPtr(),
		testTemplateIONumber:     basic.PMapIONumberPtr(),
		testTemplateIOStrNumber:  basic.PMapIOStrNumberPtr(),
		testTemplateIONumberStr:  basic.PMapIONumberStrPtr(),
		testTemplateIONumberBool: basic.PMapIONumberBoolPtr(),
		testTemplateIOStrBool:    basic.PMapIOStrBoolPtr(),
		testTemplateIOBoolNumber: basic.PMapIOBoolNumberPtr(),
		testTemplateIOBoolStr:    basic.PMapIOBoolStrPtr(),
		dataTypes:                []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName:        "pmapioptr.go",
		generatedTestFileName:    "pmapioptr_test.go",
	},

	fpCode{
		function:                 "PMapIOPtrErr",
		codeTemplate:             basic.PMapIOPtrErr(),
		testTemplateIONumber:     basic.PMapIONumberPtrErr(),
		testTemplateIOStrNumber:  basic.PMapIOStrNumberPtrErr(),
		testTemplateIONumberStr:  basic.PMapIONumberStrPtrErr(),
		testTemplateIONumberBool: basic.PMapIONumberBoolPtrErr(),
		testTemplateIOStrBool:    basic.PMapIOStrBoolPtrErr(),
		testTemplateIOBoolNumber: basic.PMapIOBoolNumberPtrErr(),
		testTemplateIOBoolStr:    basic.PMapIOBoolStrPtrErr(),
		dataTypes:                []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName:        "pmapioptrerr.go",
		generatedTestFileName:    "pmapioptrerr_test.go",
	},

	fpCode{
		function:                 "MapIO",
		codeTemplate:             basic.MapIO(),
		testTemplateIONumber:     basic.MapIONumber(),
		testTemplateIOStrNumber:  basic.MapIOStrNumber(),
		testTemplateIONumberStr:  basic.MapIONumberStr(),
		testTemplateIONumberBool: basic.MapIONumberBool(),
		testTemplateIOStrBool:    basic.MapIOStrBool(),
		testTemplateIOBoolNumber: basic.MapIOBoolNumber(),
		testTemplateIOBoolStr:    basic.MapIOBoolStr(),
		dataTypes:                []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName:        "mapio.go",
		generatedTestFileName:    "mapio_test.go",
	},

	fpCode{
		function:                 "MapIOErr",
		codeTemplate:             basic.MapIOErr(),
		testTemplateIONumber:     basic.MapIONumberErr(),
		testTemplateIOStrNumber:  basic.MapIOStrNumberErr(),
		testTemplateIONumberStr:  basic.MapIONumberStrErr(),
		testTemplateIONumberBool: basic.MapIONumberBoolErr(),
		testTemplateIOStrBool:    basic.MapIOStrBoolErr(),
		testTemplateIOBoolNumber: basic.MapIOBoolNumberErr(),
		testTemplateIOBoolStr:    basic.MapIOBoolStrErr(),
		dataTypes:                []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName:        "mapioerr.go",
		generatedTestFileName:    "mapioerr_test.go",
	},

	fpCode{
		function:                 "MapIOPtr",
		codeTemplate:             basic.MapIOPtr(),
		testTemplateIONumber:     basic.MapIONumberPtrTest(),
		testTemplateIOStrNumber:  basic.MapIOStrNumberPtrTest(),
		testTemplateIONumberStr:  basic.MapIONumberStrPtrTest(),
		testTemplateIONumberBool: basic.MapIONumberBoolPtrTest(),
		testTemplateIOStrBool:    basic.MapIOStrBoolPtrTestPtr(),
		testTemplateIOBoolNumber: basic.MapIOBoolNumberPtrTest(),
		testTemplateIOBoolStr:    basic.MapIOBoolStrPtrTest(),
		dataTypes:                []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName:        "mapioptr.go",
		generatedTestFileName:    "mapioptr_test.go",
	},

	fpCode{
		function:                 "MapIOPtrErr",
		codeTemplate:             basic.MapIOPtrErr(),
		testTemplateIONumber:     basic.MapIONumberPtrErrTest(),
		testTemplateIOStrNumber:  basic.MapIOStrNumberPtrErrTest(),
		testTemplateIONumberStr:  basic.MapIONumberStrPtrErrTest(),
		testTemplateIONumberBool: basic.MapIONumberBoolPtrErrTest(),
		testTemplateIOStrBool:    basic.MapIOStrBoolPtrErrTest(),
		testTemplateIOBoolNumber: basic.MapIOBoolNumberPtrErrTest(),
		testTemplateIOBoolStr:    basic.MapIOBoolStrPtrErrTest(),
		dataTypes:                []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName:        "mapioptrerr.go",
		generatedTestFileName:    "mapioptrerr_test.go",
	},

	fpCode{
		function:                 "FilterMapIO",
		codeTemplate:             basic.FilterMapIO(),
		testTemplateIONumber:     basic.FilterMapIONumber(),
		testTemplateIOStrNumber:  basic.FilterMapIOStrNumber(),
		testTemplateIONumberStr:  basic.FilterMapIONumberStr(),
		testTemplateIONumberBool: basic.FilterMapIONumberBool(),
		testTemplateIOStrBool:    basic.FilterMapIOStrBool(),
		testTemplateIOBoolNumber: basic.FilterMapIOBoolNumber(),
		testTemplateIOBoolStr:    basic.FilterMapIOBoolStr(),
		dataTypes:                []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName:        "filtermapio.go",
		generatedTestFileName:    "filtermapio_test.go",
	},

	fpCode{
		function:                 "FilterMapIOPtr",
		codeTemplate:             basic.FilterMapIOPtr(),
		testTemplateIONumber:     basic.FilterMapIONumberPtrTest(),
		testTemplateIOStrNumber:  basic.FilterMapIOStrNumberPtrTest(),
		testTemplateIONumberStr:  basic.FilterMapIONumberStrPtrTest(),
		testTemplateIONumberBool: basic.FilterMapIONumberBoolPtrTest(),
		testTemplateIOStrBool:    basic.FilterMapIOStrBoolPtrTest(),
		testTemplateIOBoolNumber: basic.FilterMapIOBoolNumberPtrTest(),
		testTemplateIOBoolStr:    basic.FilterMapIOBoolStrPtrTest(),
		dataTypes:                []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName:        "filtermapioptr.go",
		generatedTestFileName:    "filtermapioptr_test.go",
	},

	fpCode{
		function:                 "FilterMapIOPtrErr",
		codeTemplate:             basic.FilterMapIOPtrErr(),
		testTemplateIONumber:     basic.FilterMapIONumberPtrErrTest(),
		testTemplateIOStrNumber:  basic.FilterMapIOStrNumberPtrErrTest(),
		testTemplateIONumberStr:  basic.FilterMapIONumberStrPtrErrTest(),
		testTemplateIONumberBool: basic.FilterMapIONumberBoolPtrErrTest(),
		testTemplateIOStrBool:    basic.FilterMapIOStrBoolPtrErrTest(),
		testTemplateIOBoolNumber: basic.FilterMapIOBoolNumberPtrErrTest(),
		testTemplateIOBoolStr:    basic.FilterMapIOBoolStrPtrErrTest(),
		dataTypes:                []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName:        "filtermapioptrerr.go",
		generatedTestFileName:    "filtermapioptrerr_test.go",
	},

	fpCode{
		function:                 "FilterMapIOErr",
		codeTemplate:             basic.FilterMapIOErr(),
		testTemplateIONumber:     basic.FilterMapIONumberErrTest(),
		testTemplateIOStrNumber:  basic.FilterMapIOStrNumberErrTest(),
		testTemplateIONumberStr:  basic.FilterMapIONumberStrErrTest(),
		testTemplateIONumberBool: basic.FilterMapIONumberBoolErrTest(),
		testTemplateIOStrBool:    basic.FilterMapIOStrBoolErrTest(),
		testTemplateIOBoolNumber: basic.FilterMapIOBoolNumberErrTest(),
		testTemplateIOBoolStr:    basic.FilterMapIOBoolStrErrTest(),
		dataTypes:                []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName:        "filtermapioerr.go",
		generatedTestFileName:    "filtermapioerr_test.go",
	},
}

var importTestTemplate = `

import (
	_ "errors"
	"reflect"
	"testing"
)
`

func main() {
	fmt.Println("Generating fp code")
	generateFpCode(fpCodeList)
}

func generateFpCode(fpCodeList []fpCode) {

	for _, fpCode := range fpCodeList {
		codeTemplate := "package fp"

		if strings.Contains(strings.ToLower(fpCode.function), "pmap") {
			codeTemplate += "\n\n" + `import "sync"`
		}

		codeTemplate += "\n"

		testTemplate := "package fp"
		testTemplate += importTestTemplate

		if strings.Contains(fpCode.codeTemplate, "<INPUT_TYPE>") &&
			strings.Contains(fpCode.codeTemplate, "<OUTPUT_TYPE>") {

			for _, inputType := range fpCode.dataTypes {
				for _, outputType := range fpCode.dataTypes {

					if inputType == outputType {
						continue
					}

					fInputType := strings.Title(inputType)
					fOutputType := strings.Title(outputType)

					if fInputType == "String" {
						fInputType = "Str"
					}

					if fOutputType == "String" {
						fOutputType = "Str"
					}

					codeTemplate += fpCode.codeTemplate
					r := strings.NewReplacer("<FINPUT_TYPE>", fInputType, "<FOUTPUT_TYPE>", fOutputType, "<INPUT_TYPE>", inputType, "<OUTPUT_TYPE>", outputType)
					codeTemplate = r.Replace(codeTemplate)

					// Generate tests for number types combination
					inputTypeNumbers := "int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32"
					outputTypeNumbers := "int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32"
					if strings.Contains(inputTypeNumbers, strings.ToLower(fInputType)) && strings.Contains(outputTypeNumbers, strings.ToLower(outputType)) {
						testTemplate += r.Replace(fpCode.testTemplateIONumber)
					}

					// Generate tests for input type string - output type numbers
					if strings.ToLower(inputType) == "string" && strings.Contains(outputTypeNumbers, strings.ToLower(outputType)) {
						testTemplate += r.Replace(fpCode.testTemplateIOStrNumber)
					}

					// Generate tests for input type Number - output type string
					if strings.Contains(inputTypeNumbers, strings.ToLower(inputType)) && strings.ToLower(outputType) == "string" {
						testTemplate += r.Replace(fpCode.testTemplateIONumberStr)
					}

					// Generate tests for input type Number - output type bool
					if strings.Contains(inputTypeNumbers, strings.ToLower(inputType)) && strings.ToLower(outputType) == "bool" {
						testTemplate += r.Replace(fpCode.testTemplateIONumberBool)
					}

					// Generate tests for input type str - output type bool
					if strings.ToLower(inputType) == "string" && strings.ToLower(outputType) == "bool" {
						testTemplate += r.Replace(fpCode.testTemplateIOStrBool)
					}

					// Generate tests for input type Bool - output type Number
					if strings.ToLower(inputType) == "bool" && strings.Contains(outputTypeNumbers, strings.ToLower(outputType)) {
						testTemplate += r.Replace(fpCode.testTemplateIOBoolNumber)
					}

					// Generate tests for input type Bool - output type Str
					if strings.ToLower(inputType) == "bool" && strings.ToLower(outputType) == "string" {
						testTemplate += r.Replace(fpCode.testTemplateIOBoolStr)
					}
				}
			}

			// For Merge, Zip -
		} else if strings.Contains(fpCode.codeTemplate, "<INPUT_TYPE1>") &&
			strings.Contains(fpCode.codeTemplate, "<INPUT_TYPE2>") {

			for _, inputType1 := range fpCode.dataTypes {
				for _, inputType2 := range fpCode.dataTypes {

					fInputType1 := strings.Title(inputType1)
					fInputType2 := strings.Title(inputType2)

					if fInputType1 == "String" {
						fInputType1 = "Str"
					}

					if fInputType2 == "String" {
						fInputType2 = "Str"
					}

					if fInputType1 == fInputType2 {
						fInputType2 = ""
					}

					codeTemplate += fpCode.codeTemplate
					r := strings.NewReplacer("<FINPUT_TYPE1>", fInputType1, "<FINPUT_TYPE2>", fInputType2, "<INPUT_TYPE1>", inputType1, "<INPUT_TYPE2>", inputType2)
					codeTemplate = r.Replace(codeTemplate)

					// Generate tests for number types combination
					inputTypeNumbers := "int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32"
					if strings.Contains(inputTypeNumbers, strings.ToLower(fInputType1)) && strings.Contains(inputTypeNumbers, strings.ToLower(inputType2)) {
						testTemplate += r.Replace(fpCode.testTemplate)
					}

					// Generate tests for input type1 Number - input type2 string
					if strings.Contains(inputTypeNumbers, strings.ToLower(inputType1)) && strings.ToLower(inputType2) == "string" {
						testTemplate += r.Replace(fpCode.testTemplateNumberStr)
					}

					// Generate tests for input type1 string - input type2 Numbers
					if strings.ToLower(inputType1) == "string" && strings.Contains(inputTypeNumbers, strings.ToLower(inputType2)) {
						testTemplate += r.Replace(fpCode.testTemplateStrNumber)
					}

					// Generate tests for input type1 string - input type2 Bool
					if strings.ToLower(inputType1) == "string" && strings.ToLower(inputType2) == "bool" {
						testTemplate += r.Replace(fpCode.testTemplateStrBool)
					}

					// Generate tests for input type1 bool - input type2 string
					if strings.ToLower(inputType1) == "bool" && strings.ToLower(inputType2) == "string" {
						testTemplate += r.Replace(fpCode.testTemplateBoolStr)
					}

					// Generate tests for input type1 Number - input type2 bool
					if strings.Contains(inputTypeNumbers, inputType1) && strings.ToLower(inputType2) == "bool" {
						testTemplate += r.Replace(fpCode.testTemplateNumberBool)
					}

					// Generate tests for input type1 bool - input type2 Number
					if strings.ToLower(inputType1) == "bool" && strings.Contains(inputTypeNumbers, inputType2) {
						testTemplate += r.Replace(fpCode.testTemplateBoolNumber)
					}

					// Generate tests for input type1 bool - input type2 bool
					if strings.ToLower(inputType1) == "bool" && strings.ToLower(inputType2) == "bool" {
						testTemplate += r.Replace(fpCode.testTemplateBoolBool)
					}

					// Generate tests for input type1 string - input type2 string
					if strings.ToLower(inputType1) == "string" && strings.ToLower(inputType2) == "string" {
						testTemplate += r.Replace(fpCode.testTemplateStrStr)
					}
				}
			}
		} else {

			for _, t := range fpCode.dataTypes {

				codeTemplate += fpCode.codeTemplate

				ftype := strings.Title(t)
				switch t {
				case "string":
					ftype = "Str"
					testTemplate += modifyTestDataToStr(fpCode.testTemplate)

				default:
					if strings.ToLower(t) == "bool" {
						testTemplate += fpCode.testTemplateBool
					} else {
						testTemplate += fpCode.testTemplate
					}
				}

				r := strings.NewReplacer("<TYPE>", t, "<FTYPE>", ftype)

				// for methodchain
				r2 := strings.NewReplacer("<TYPE>", t, "<FTYPE>", ftype, "<NEWTYPE>", t)
				codeTemplate = r2.Replace(codeTemplate)
				codeTemplate = r.Replace(codeTemplate)

				testTemplate = r.Replace(testTemplate)
			}
		}

		codeTemplate = modifyCodeData(codeTemplate)
		switch {
		case strings.Contains(codeTemplate, "Sort"):
			codeTemplate = basic.ReplaceActivitySortingCode(codeTemplate)
		}
		writeToFile(codeTemplate, fmt.Sprintf("fp/%s", fpCode.generatedFileName))
		if fpCode.generatedTestFileName != "" {
			testTemplate = modifyTestDataToStr2(testTemplate)
			testTemplate = modifyCodeData(testTemplate)
			writeToFile(testTemplate, fmt.Sprintf("fp/%s", fpCode.generatedTestFileName))
		}
	}

	fmt.Println("Functional code generated successfully")
}

func writeToFile(text, file string) {
	f, err := os.Create(file)
	if err != nil {
		log.Fatalf("Failed opening destination file=%v,error=%v", file, err)
	}

	f.Write([]byte(text))
	defer f.Close()
}

func modifyCodeData(code string) string {
	code = strings.Replace(code, "Int64sPtr", "Ints64Ptr", -1)
	code = strings.Replace(code, "Int32sPtr", "Ints32Ptr", -1)
	code = strings.Replace(code, "Int16sPtr", "Ints16Ptr", -1)
	code = strings.Replace(code, "Int8sPtr", "Ints8Ptr", -1)
	code = strings.Replace(code, "Unt64sPtr", "Uints64Ptr", -1)
	code = strings.Replace(code, "Uint32sPtr", "Uints32Ptr", -1)
	code = strings.Replace(code, "Uint16sPtr", "Uints16Ptr", -1)
	code = strings.Replace(code, "Uint8sPtr", "Uints8Ptr", -1)

	code = strings.Replace(code, "Int64s", "Ints64", -1)
	code = strings.Replace(code, "Int32s", "Ints32", -1)
	code = strings.Replace(code, "Int16s", "Ints16", -1)
	code = strings.Replace(code, "Int8s", "Ints8", -1)
	code = strings.Replace(code, "Unt64s", "Uints", -1)
	code = strings.Replace(code, "Uint32s", "Uints32", -1)
	code = strings.Replace(code, "Uint16s", "Uints16", -1)
	code = strings.Replace(code, "Uint8s", "Uints8", -1)
	return code
}

func modifyTestDataToStr(code string) string {
	code = strings.Replace(code, "{1, 2, 3, 4, 5}", "{\"1\", \"2\", \"3\", \"4\", \"5\"}", -1)
	code = strings.Replace(code, "{1, 2, 3, 4}", "{\"1\", \"2\", \"3\", \"4\"}", -1)
	code = strings.Replace(code, "{1, 2}", "{\"1\", \"2\"}", -1)
	code = strings.Replace(code, "{1}", "{\"1\"}", -1)

	code = strings.Replace(code, "var v1 string = 1", "var v1 string = \"1\"", -1)
	code = strings.Replace(code, "var v2 string = 2", "var v2 string = \"2\"", -1)
	code = strings.Replace(code, "var v3 string = 3", "var v3 string = \"3\"", -1)
	code = strings.Replace(code, "var v4 string = 4", "var v4 string = \"4\"", -1)
	code = strings.Replace(code, "var v5 string = 5", "var v5 string = \"5\"", -1)

	code = strings.Replace(code, "string(0)", "0", -1)
	code = strings.Replace(code, "string = 0", "string = \"0\"", -1)
	code = strings.Replace(code, `partialAddStr := func(num string) (string, error) {
		r, err := addStrErr(v5, num)
		if err != nil {
			return 0, err
		}
		return r, nil
	}`, `partialAddStr := func(num string) (string, error) {
		r, err := addStrErr(v5, num)
		if err != nil {
			return "0", err
		}
		return r, nil
	}`, -1)
	return code
}

func modifyTestDataToStr2(code string) string {
	code = strings.Replace(code, "var v0 string = 0", "var v0 string = \"0\"", -1)
	code = strings.Replace(code, "var v0 string = 0", "var v0 string = \"0\"", -1)
	code = strings.Replace(code, "var v1 string = 1", "var v1 string = \"1\"", -1)
	code = strings.Replace(code, "var v2 string = 2", "var v2 string = \"2\"", -1)
	code = strings.Replace(code, "var v3 string = 3", "var v3 string = \"3\"", -1)
	code = strings.Replace(code, "var v4 string = 4", "var v4 string = \"4\"", -1)
	code = strings.Replace(code, "var v5 string = 5", "var v5 string = \"5\"", -1)
	code = strings.Replace(code, "var v6 string = 6", "var v6 string = \"6\"", -1)
	code = strings.Replace(code, "var v7 string = 7", "var v7 string = \"7\"", -1)
	code = strings.Replace(code, "var v8 string = 8", "var v8 string = \"8\"", -1)
	code = strings.Replace(code, "var v9 string = 9", "var v9 string = \"9\"", -1)
	code = strings.Replace(code, "var v10 string = 10", "var v10 string = \"10\"", -1)
	code = strings.Replace(code, "var v16 string = 16", "var v16 string = \"16\"", -1)
	code = strings.Replace(code, "var v20 string = 20", "var v20 string = \"20\"", -1)
	code = strings.Replace(code, "var v30 string = 30", "var v30 string = \"30\"", -1)
	code = strings.Replace(code, "var v40 string = 40", "var v40 string = \"40\"", -1)
	code = strings.Replace(code, "var v80 string = 80", "var v80 string = \"80\"", -1)
	code = strings.Replace(code, "string(0)", "0", -1)
	code = strings.Replace(code, "string = 0", "string = \"0\"", -1)
	code = strings.Replace(code, `partialAddStr := func(num string) (string, error) {
		r, err := addStrErr(v5, num)
		if err != nil {
			return 0, err
		}
		return r, nil
	}`, `partialAddStr := func(num string) (string, error) {
		r, err := addStrErr(v5, num)
		if err != nil {
			return "0", err
		}
		return r, nil
	}`, -1)

	// Change one of the test for MapPtrStr
	s1 := `func TestMapStrPtr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v5 string = "5"
	var v6 string = "6"
	var v7 string = "7"
	var v8 string = "8"


	// Test: add 5 to each item in the list
	expectedSumList := []*string{&v6, &v7, &v8}
	partialAddStrPtr := func(num *string) *string {
		return addStrPtr(&v5, num)
	}
	sumList := MapStrPtr(partialAddStrPtr, []*string{&v1, &v2, &v3})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("MapStrPtr failed")
	}

	if len(MapStrPtr(nil, nil)) > 0 {
		t.Errorf("MapStrPtr failed.")
		t.Errorf(reflect.String.String())
	}
}`

	s2 := `func TestMapStrPtr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v5 string = "5"
	var v6 string = "51"
	var v7 string = "52"
	var v8 string = "53"


	// Test: add 5 to each item in the list
	expectedSumList := []*string{&v6, &v7, &v8}
	partialAddStrPtr := func(num *string) *string {
		return addStrPtr(&v5, num)
	}
	sumList := MapStrPtr(partialAddStrPtr, []*string{&v1, &v2, &v3})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("MapStrPtr failed")
	}

	if len(MapStrPtr(nil, nil)) > 0 {
		t.Errorf("MapStrPtr failed.")
		t.Errorf(reflect.String.String())
	}
}`

	code = strings.Replace(code, s1, s2, -1)

	s1 = `func TestFilterStrPtr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v10 string = "10"
	var v20 string = "20"
	var v40 string = "40"


	// Test : even number in the list
	expectedFilteredList := []*string{&v2, &v4}
	filteredList := FilterStrPtr(isEvenStrPtr, []*string{&v1, &v2, &v3, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] {
		t.Errorf("MapFilter failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	// Test: filter all even numbers divisible by 10 in the list
	expectedFilteredList = []*string{&v20, &v40}
	partialIsEven := func(num *string) bool { return isEvenDivisibleByStrPtr(num, &v10) }
	filteredList = FilterStrPtr(partialIsEven, []*string{&v20, &v1, &v3, &v40})

	if filteredList[0] != expectedFilteredList[0] || filteredList[1] != expectedFilteredList[1] {
		t.Errorf("MapFilter failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	if len(FilterStrPtr(nil, nil)) > 0 {
		t.Errorf("FilterInt failed.")
	}
}

func isEvenStrPtr(num *string) bool {
	return *num%2 == 0
}

func isEvenDivisibleByStrPtr(num, divisibleBy *string) bool {
	return *num%2 == 0 && *num % *divisibleBy == 0
}`
	s2 = `func TestFilterStrPtr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"

	// Test : even number in the list
	expectedFilteredList := []*string{&v2, &v4}
	filteredList := FilterStrPtr(isEvenStrPtr, []*string{&v1, &v2, &v3, &v4})

	if *filteredList[0] != *expectedFilteredList[0] {
		t.Errorf("MapFilter failed. Expected filtered list=%v, actual list=%v", *expectedFilteredList[0], *filteredList[0])
	}

	if len(FilterStrPtr(nil, nil)) > 0 {
		t.Errorf("FilterInt failed.")
        t.Errorf(reflect.String.String())
	}
}

func isEvenStrPtr(num *string) bool {
	if *num == "2" || *num == "4" || *num == "6" || *num == "8" || *num == "10" {
		return true
	}
	return false
}`
	code = strings.Replace(code, s1, s2, -1)

	s1 = `newList = DropStrsPtr(nil, []*string{&v1, &v4})
	if *newList[0] != 1 || *newList[1] != 4 {
		t.Errorf("DropStrs failed. Expected list=%v, actual list=%v", expectedList, newList)
	}`
	s2 = `newList = DropStrsPtr(nil, []*string{&v1, &v4})
	if *newList[0] != "1" || *newList[1] != "4" {
		t.Errorf("DropStrs failed. Expected list=%v, actual list=%v", expectedList, newList)
	}`
	code = strings.Replace(code, s1, s2, -1)

	s1 = `func isEvenFloat32Ptr(num *float32) bool {
	return *num%2 == 0
}

func isEvenDivisibleByFloat32Ptr(num, divisibleBy *float32) bool {
	return *num%2 == 0 && *num % *divisibleBy == 0
}`
	s2 = `func isEvenFloat32Ptr(num *float32) bool {
		return int(*num)%2 == 0
	}
	
	func isEvenDivisibleByFloat32Ptr(num, divisibleBy *float32) bool {
		return int(*num)%2 == 0 && int(*num) % int(*divisibleBy) == 0
	}`

	code = strings.Replace(code, s1, s2, -1)

	s1 = `func isEvenFloat64Ptr(num *float64) bool {
	return *num%2 == 0
}

func isEvenDivisibleByFloat64Ptr(num, divisibleBy *float64) bool {
	return *num%2 == 0 && *num % *divisibleBy == 0
}`
	s2 = `func isEvenFloat64Ptr(num *float64) bool {
	return int(*num)%2 == 0
}

func isEvenDivisibleByFloat64Ptr(num, divisibleBy *float64) bool {
	return int(*num)%2 == 0 && int(*num) % int(*divisibleBy) == 0
}
`
	code = strings.Replace(code, s1, s2, -1)

	code = basic.ReplaceActivityDropWhilePtr(code)

	code = basic.ReplaceActivityEveryPtrTest(code)

	code = basic.ReplaceActivityMapPtr(code)

	s1 = `func TestReduceStrPtr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v5 string = "5"

	list := []*string{&v1, &v2, &v3, &v4, &v5}
	var expected string = 15
	actual := ReduceStrPtr(plusStrPtr, list)
	if *actual != expected {
		t.Errorf("ReduceStrPtr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*string{&v1, &v2, &v3, &v4, &v5}
	expected = 18
	actual = ReduceStrPtr(plusStrPtr, list, 3)
	if *actual != expected {
		t.Errorf("ReduceStrPtr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*string{&v1, &v2}
	expected = 3
	actual = ReduceStrPtr(plusStrPtr, list)
	if *actual != expected {
		t.Errorf("ReduceInt failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*string{&v1}
	expected = 1
	actual = ReduceStrPtr(plusStrPtr, list)
	if *actual != expected {
		t.Errorf("ReduceStrPtr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*string{}
	expected = 0
	actual = ReduceStrPtr(plusStrPtr, list)
	if *actual != expected {
		t.Errorf("ReduceStrPtr failed. actual=%v, expected=%v", *actual, expected)
		t.Errorf(reflect.String.String())
	}
}`

	s2 = `func TestReduceStrPtr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v5 string = "5"

	list := []*string{&v1, &v2, &v3, &v4, &v5}
	var expected string = "12345"
	actual := ReduceStrPtr(plusStrPtr, list)
	if *actual != expected {
		t.Errorf("ReduceStrPtr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*string{&v1, &v2, &v3, &v4, &v5}
	expected = "312345"
	actual = ReduceStrPtr(plusStrPtr, list, "3")
	if *actual != expected {
		t.Errorf("ReduceStrPtr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*string{&v1, &v2}
	expected = "12"
	actual = ReduceStrPtr(plusStrPtr, list)
	if *actual != expected {
		t.Errorf("ReduceInt failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*string{&v1}
	expected = "1"
	actual = ReduceStrPtr(plusStrPtr, list)
	if *actual != expected {
		t.Errorf("ReduceStrPtr failed. actual=%v, expected=%v", *actual, expected)
	}

	list = []*string{}
	expected = ""
	actual = ReduceStrPtr(plusStrPtr, list)
	if *actual != expected {
		t.Errorf("ReduceStrPtr failed. actual=%v, expected=%v", *actual, expected)
		t.Errorf(reflect.String.String())
	}
}`
	code = strings.Replace(code, s1, s2, -1)

	code = basic.ReplaceActivityRemoveStrPtr(code)

	code = basic.ReplaceActivityMapPtrErrTest(code)
	code = basic.ReplaceActivityMapErrTest(code)

	code = basic.ReplaceActivityFilterPtrErrTest(code)
	code = basic.ReplaceActivityFilterErrTest(code)

	code = basic.ReplaceActivityFilterMapPtrErr(code)
	code = basic.ReplaceActivityFilterMapErr(code)

	code = basic.ReplaceActivityFilterMapIOPtrErr(code)
	code = basic.ReplaceActivityFilterMapIOErr(code)

	code = basic.ReplaceActivitySomePtrErr(code)
	code = basic.ReplaceActivitySomeErr(code)

	code = basic.ReplaceActivityEveryPtrErr(code)
	code = basic.ReplaceActivityEveryErr(code)

	code = basic.ReplaceActivityDropWhilePtrErr(code)
	code = basic.ReplaceActivityDropWhileErr(code)

	code = basic.ReplaceActivityTakeWhilePtrErr(code)
	code = basic.ReplaceActivityTakeWhileErr(code)

	code = basic.ReplaceActivityReducePtrErr(code)
	code = basic.ReplaceActivityReduceErr(code)

	code = basic.ReplaceActivityRemovePtrErr(code)
	code = basic.ReplaceActivityRemoveErr(code)

	code = basic.ReplaceActivityPMapPtrErr(code)
	code = basic.ReplaceActivityPMapErr(code)

	code = basic.ReplaceActivityMapIOErr(code)
	code = basic.ReplaceActivityMapIOPtrErr(code)

	code = basic.ReplaceActivityTakeWhilePtr(code)

	code = basic.ReplaceActivityPMapPtr2Test(code)
	code = basic.ReplaceActivityPMap2Test(code)

	code = basic.ReplaceActivityPMapPtr2ErrTest(code)

	code = basic.ReplaceActivityPMap2ErrTest(code)

	code = methodchain.ReplaceActivityMethodChainMap(code)

	return code
}
