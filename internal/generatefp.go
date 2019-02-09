package main

import (
	"fmt"
	"github.com/logic-building/functional-go/internal/template/basic"
	"log"
	"os"
	"strings"
)

type fpCode struct {
	function              string
	codeTemplate          string
	testTemplate          string
	dataTypes             []string
	generatedFileName     string
	generatedTestFileName string
}

var fpCodeList = []fpCode{fpCode{
	function:          "DropLast",
	codeTemplate:      basic.DropLast(),
	dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string"},
	generatedFileName: "droplast.go",

	testTemplate:          basic.DropLastTest(),
	generatedTestFileName: "droplast_test.go",
}, fpCode{
	function:          "PMapIO",
	codeTemplate:      basic.PMapIO(),
	dataTypes:         []string{"int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "string"},
	generatedFileName: "pmapio.go",
}}

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
			strings.Contains(fpCode.codeTemplate, "<OUTPUT_TYPE") {

			for _, inputType := range fpCode.dataTypes {
				for _, outputType := range fpCode.dataTypes {

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
					testTemplate += fpCode.testTemplate + "\n"
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
