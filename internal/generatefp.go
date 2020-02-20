package main

import (
	"fmt"
	"github.com/logic-building/functional-go/internal/template/basic"
	"log"
	"os"
	"strings"
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
		function:          "DropLast",
		codeTemplate:      basic.DropLast(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "droplast.go",

		testTemplate:          basic.DropLastTest(),
		testTemplateBool:      basic.DropLastBoolTest(),
		generatedTestFileName: "droplast_test.go",
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
		function:          "FilterPtr",
		codeTemplate:      basic.FilterPtr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "filterPtr.go",

		testTemplate:          basic.FilterPtrTest(),
		testTemplateBool:      basic.FilterPtrBoolTest(),
		generatedTestFileName: "filterPtr_test.go",
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
		function:          "FilterMapPtr",
		codeTemplate:      basic.FilterMapPtr(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool", "float32", "float64"},
		generatedFileName: "filtermapptr.go",

		testTemplate:          basic.FilterMapPtrTest(),
		testTemplateBool:      basic.FilterMapPtrBoolTest(),
		generatedTestFileName: "filtermapptr_test.go",
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
}

var importTestTemplate = `

import (
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
				codeTemplate = r.Replace(codeTemplate)

				testTemplate = r.Replace(testTemplate)
			}
		}

		codeTemplate = modifyCodeData(codeTemplate)
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
	return code
}

func modifyTestDataToStr2(code string) string {
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
	code = strings.Replace(code, "var v20 string = 20", "var v20 string = \"20\"", -1)
	code = strings.Replace(code, "var v40 string = 40", "var v40 string = \"40\"", -1)
	code = strings.Replace(code, "var v80 string = 80", "var v80 string = \"80\"", -1)

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
	expectedFilteredList := []*string{&v1, &v2, &v3, &v4}
	filteredList := FilterStrPtr(isEvenStrPtr, []*string{&v1, &v2, &v3, &v4})

	if *filteredList[0] != *expectedFilteredList[0] {
		t.Errorf("MapFilter failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}

	if len(FilterStrPtr(nil, nil)) > 0 {
		t.Errorf("FilterInt failed.")
        t.Errorf(reflect.String.String())
	}
}

func isEvenStrPtr(num *string) bool {
	return true
}`
	code = strings.Replace(code, s1, s2, -1)

	s1 = `func TestFilterMapStrPtr(t *testing.T) {
	// Test : Multiply all positive numbers in the list by 2

	var v1 string = "1"
	var v4 string = "4"
	var v8 string = "8"
	var v0 string = "0"
	var v2 string = "2"

	expectedFilteredList := []*string{&v2, &v4, &v8}
	filteredList := FilterMapStrPtr(isPositiveStrPtr, multiplyBy2StrPtr, []*string{&v1, &v0, &v2, &v4})

	if *filteredList[0] != *expectedFilteredList[0] || *filteredList[1] != *expectedFilteredList[1] || *filteredList[2] != *expectedFilteredList[2]{
		t.Errorf("FilterMapStrPtr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterMapStrPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapStrPtr failed.")
		t.Errorf(reflect.String.String())
	}
}

func isPositiveStrPtr(num *string) bool {
	return *num > 0
}
func multiplyBy2StrPtr(num *string) *string {
	result := *num * 2
	return &result
}`
	s2 = `func TestFilterMapStrPtr(t *testing.T) {
	var v1 string = "1"
	var v4 string = "4"
	var v0 string = "0"
	var v2 string = "2"

	expectedFilteredList := []*string{&v1, &v0, &v2, &v4}
	filteredList := FilterMapStrPtr(noFilter, concatA, []*string{&v1, &v0, &v2, &v4})

	if *filteredList[0] != *expectedFilteredList[0]+"A" || *filteredList[1] != *expectedFilteredList[1]+"A" {
		t.Errorf("FilterMapStrPtr failed. Expected filtered list=%v, actual list=%v", expectedFilteredList, filteredList)
	}
	if len(FilterMapStrPtr(nil, nil, nil)) > 0 {
		t.Errorf("FilterMapStrPtr failed.")
	}
}

func noFilter(num *string) bool {
	return true
}
func concatA(num *string) *string {
	result := *num + "A"
	return &result
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

	s1 = `func TestDropWhileStrPtr(t *testing.T) {
	// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails

	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v5 string = "5"

	expectedNewList := []*string{&v3, &v4, &v5}
	NewList := DropWhileStrPtr(isEvenStrPtr, []*string{&v4, &v2, &v3, &v4, &v5})
	if *NewList[0] != *expectedNewList[0] || *NewList[1] != *expectedNewList[1] || *NewList[2] != *expectedNewList[2] {
		t.Errorf("DropWhileStrPtr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
	}

	if len(DropWhileStrPtr(nil, nil)) > 0 {
		t.Errorf("DropWhileStrPtr failed.")
	}

	if len(DropWhileStrPtr(nil, []*string{})) > 0 {
		t.Errorf("DropWhileStrPtr failed.")
		t.Errorf(reflect.String.String())
	}
}`
	s2 = `
	func TestDropWhileStrPtr(t *testing.T) {
		// Test : drop the numbers as long as condition match and returns remaining number in the list once condition fails
	
		var v2 string = "2"
		var v3 string = "3"
		var v4 string = "4"
		var v5 string = "5"
	
		expectedNewList := []*string{&v3, &v4, &v5}
		NewList := DropWhileStrPtr(isEvenStrPtr, []*string{&v4, &v2, &v3, &v4, &v5})
		if len(NewList) > 0 {
			t.Errorf("DropWhileStrPtr failed. Expected New list=%v, actual list=%v", expectedNewList, NewList)
		}
	
		if len(DropWhileStrPtr(nil, nil)) > 0 {
			t.Errorf("DropWhileStrPtr failed.")
		}
	
		if len(DropWhileStrPtr(nil, []*string{})) > 0 {
			t.Errorf("DropWhileStrPtr failed.")
		}
	}`

	code = strings.Replace(code, s1, s2, -1)
	return code
}
