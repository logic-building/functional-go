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

	// Test template for different combination of data type
	testTemplateIONumber     string
	testTemplateIOStrNumber  string
	testTemplateIONumberStr  string
	testTemplateIONumberBool string
	testTemplateIOStrBool    string
	testTemplateIOBoolNumber string
	testTemplateIOBoolStr    string

	dataTypes             []string
	generatedFileName     string
	generatedTestFileName string
}

var fpCodeList = []fpCode{
	fpCode{
		function:          "DropLast",
		codeTemplate:      basic.DropLast(),
		dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool"},
		generatedFileName: "droplast.go",

		testTemplate:          basic.DropLastTest(),
		testTemplateBool:      basic.DropLastBoolTest(),
		generatedTestFileName: "droplast_test.go",
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
		dataTypes:                []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool"},
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
		dataTypes:                []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool"},
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
		dataTypes:                []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string", "bool"},
		generatedFileName:        "filtermapio.go",
		generatedTestFileName:    "filtermapio_test.go",
	},
}

var importTestTemplate = `
import (
	"reflect"
	"testing"
)`

func main() {
	fmt.Println("Generating fp code")
	generateFpCode(fpCodeList)
}

func generateFpCode(fpCodeList []fpCode) {

	for _, fpCode := range fpCodeList {
		codeTemplate := "package fp \n"

		if strings.Contains(strings.ToLower(fpCode.function), "pmap") {
			codeTemplate += "\n" + `import "sync"` + "\n"
		}

		testTemplate := "package fp \n"
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

					codeTemplate += fpCode.codeTemplate + "\n"
					r := strings.NewReplacer("<FINPUT_TYPE>", fInputType, "<FOUTPUT_TYPE>", fOutputType, "<INPUT_TYPE>", inputType, "<OUTPUT_TYPE>", outputType)
					codeTemplate = r.Replace(codeTemplate)

					// Generate tests for number types combination
					inputTypeNumbers := "int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32"
					outputTypeNumbers := "int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32"
					if strings.Contains(inputTypeNumbers, strings.ToLower(fInputType)) && strings.Contains(outputTypeNumbers, strings.ToLower(outputType)) {
						testTemplate += "\n" + r.Replace(fpCode.testTemplateIONumber)
					}

					// Generate tests for input type string - output type numbers
					if strings.ToLower(inputType) == "string" && strings.Contains(outputTypeNumbers, strings.ToLower(outputType)) {
						testTemplate += "\n" + r.Replace(fpCode.testTemplateIOStrNumber)
					}

					// Generate tests for input type Number - output type string
					if strings.Contains(inputTypeNumbers, strings.ToLower(inputType)) && strings.ToLower(outputType) == "string" {
						testTemplate += "\n" + r.Replace(fpCode.testTemplateIONumberStr)
					}

					// Generate tests for input type Number - output type bool
					if strings.Contains(inputTypeNumbers, strings.ToLower(inputType)) && strings.ToLower(outputType) == "bool" {
						testTemplate += "\n" + r.Replace(fpCode.testTemplateIONumberBool)
					}

					// Generate tests for input type str - output type bool
					if strings.ToLower(inputType) == "string" && strings.ToLower(outputType) == "bool" {
						testTemplate += "\n" + r.Replace(fpCode.testTemplateIOStrBool)
					}

					// Generate tests for input type Bool - output type Number
					if strings.ToLower(inputType) == "bool" && strings.Contains(outputTypeNumbers, strings.ToLower(outputType)) {
						testTemplate += "\n" + r.Replace(fpCode.testTemplateIOBoolNumber)
					}

					// Generate tests for input type Bool - output type Str
					if strings.ToLower(inputType) == "bool" && strings.ToLower(outputType) == "string" {
						testTemplate += "\n" + r.Replace(fpCode.testTemplateIOBoolStr)
					}
				}
			}

		} else {
			for _, t := range fpCode.dataTypes {

				codeTemplate += fpCode.codeTemplate + "\n"

				ftype := strings.Title(t)
				switch t {
				case "string":
					ftype = "Str"
					testTemplate += modifyTestDataToStr(fpCode.testTemplate) + "\n"

				default:
					if strings.ToLower(t) == "bool" {
						testTemplate += fpCode.testTemplateBool + "\n"
					} else {
						testTemplate += fpCode.testTemplate + "\n"
					}
				}

				r := strings.NewReplacer("<TYPE>", t, "<FTYPE>", ftype)
				codeTemplate = r.Replace(codeTemplate)

				testTemplate = r.Replace(testTemplate)
			}
		}

		writeToFile(codeTemplate, fmt.Sprintf("fp/%s", fpCode.generatedFileName))
		if fpCode.generatedTestFileName != "" {
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

func modifyTestDataToStr(code string) string {
	code = strings.Replace(code, "{1, 2, 3, 4, 5}", "{\"1\", \"2\", \"3\", \"4\", \"5\"}", -1)
	code = strings.Replace(code, "{1, 2, 3, 4}", "{\"1\", \"2\", \"3\", \"4\"}", -1)
	code = strings.Replace(code, "{1, 2}", "{\"1\", \"2\"}", -1)
	code = strings.Replace(code, "{1}", "{\"1\"}", -1)
	return code
}
