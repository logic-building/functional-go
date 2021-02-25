// Generates functional code locally for user defined data type.
/*
```
1. Install "gofp" to generate code
   go get github.com/logic-building/functional-go/gofp
   go get -u github.com/logic-building/functional-go/gofp
   go install github.com/logic-building/functional-go/gofp

2. Add this line in a file where user defined data type exists
   //go:generate gofp -destination <file> -pkg <pkg> -type <Types separated by comma>

example:
    package employee

   //go:generate gofp -destination fp.go -pkg employee -type "Employee, Teacher"
   type Employee struct {
   	id     int
   	name   string
   	salary float64
   }

   type Teacher struct {
   	id     int
   	name   string
   	salary float64
   }

Note:
   A. fp.go               :  generated code
   B. employee            :  package name
   C. "Employee, Teacher" :  User defined data types

3. Generate functional code
   go generate ./...

4. Now write your code

    emp1 := employee.Employee{1, "A", 1000}
   	emp2 := employee.Employee{1, "A", 1000}
   	emp3 := employee.Employee{1, "A", 1000}

   	empList := []employee.Employee{emp1, emp2, emp3}

   	newEmpList := employee.Map(incrementSalary, empList) //  Returns: [{1 A 1500} {1 A 1500} {1 A 1500}]

   func incrementSalary(emp employee.Employee) employee.Employee {
        emp.Salary = emp.Salary + 500
        return emp
   }

```
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"math/rand"
	"runtime"
	"time"

	"github.com/logic-building/functional-go/fp"
	template2 "github.com/logic-building/functional-go/internal/template"
	"github.com/logic-building/functional-go/internal/template/basic"
	template3 "github.com/logic-building/functional-go/internal/template/methodchain"
)

var (
	destination = flag.String("destination", "", "functional code for user defined data type")
	pkgName     = flag.String("pkg", "", "package name for generated files")
	types       = flag.String("type", "", "user defined type")
	imports     = flag.String("imports", "", "import statements for user defined types when structs are in different package")
	mapFunction = flag.String("mapfun", "", "this allows to create map function such as zip, merge.")
	only        = flag.String("only", "", "includes list of function to be auto-generated")
	sortStr     = flag.String("sort", "", "generate sorting functions for user defined type. format- <StructName>:<FieldName> eg. Employee:Name, Employee:Salary")
	setStr      = flag.String("set", "", "generate 'set' functions(union, intersection, difference, subset, superset) for user defined type. format- <StructName>:<FieldName> eg. Employee:Name, Employee:Salary")

	onlyList []string

	switchTemplateDistinct = false
)

func main() {
	// when option is provided with operator(=), double quote is read too if provided. So trim it.
	*destination = strings.Trim(*destination, "\"")
	*pkgName = strings.Trim(*pkgName, "\"")
	*types = strings.Trim(*types, "\"")
	*imports = strings.Trim(*imports, "\"")
	*mapFunction = strings.Trim(*mapFunction, "\"")
	*only = strings.Trim(*only, "\"")
	*sortStr = strings.Trim(*sortStr, "\"")
	*setStr = strings.Trim(*setStr, "\"")

	isAlreadyRun := runWithin(time.Second * 15)
	defer func() {
		if !isAlreadyRun {
			fmt.Println("\n\t\t\"" + quoteForTheDay() + "\"\n")
		}
	}()

	flag.Parse()

	if *destination == "" || *pkgName == "" || *types == "" {
		fmt.Println("either of these fields : (destination, package, types) - is not provided")
		usage()
		return
	}

	if len(*destination) > 0 {
		os.Chmod(*destination, 0777)
		if err := os.MkdirAll(filepath.Dir(*destination), os.ModePerm); err != nil {
			log.Fatalf("Unable to create destination directory: %v", err)
			usage()
		}
		f, err := os.Create(*destination)
		if err != nil {
			log.Fatalf("Failed opening destination file: %v", err)
		}

		if *only != "" {
			// when -sort option is provided with operator(=), double quote is read too if provided
			onlyList = strings.Split(*only, ",")
			onlyList = fp.MapStr(strings.TrimSpace, onlyList)
		}

		structToFieldsMap, structToFieldsMapUnExpected := findStructNamesAndFieldsGivenInGoGenerate()

		generatedCode, err := generateFPCode(*pkgName, *types, *imports, structToFieldsMapUnExpected)
		if err != nil {
			usage()
			log.Fatalf("Failed code generation: %v", err)
		}

		generatedCodeIO, err := generateFPCodeIO(*pkgName, *types)
		if err != nil {
			usage()
			log.Fatalf("Failed code generation for different IO combination: %v", err)
		}

		generatedCodeII, err := generateFPCodeII(*pkgName, *types)
		if err != nil {
			usage()
			log.Fatalf("Failed code generation for different IO combination: %v", err)
		}

		// Generate zip or merge function if this is true
		if *mapFunction != "true" {
			generatedCodeII = ""
		}

		var setCode string

		if *setStr != "" {
			if strings.Contains(strings.ToLower(*setStr), "time") {
				generatedCode = strings.Replace(generatedCode,
					`import "sync" `,
					`import "sync" 
import "time"`, -1)
			}
			setCode = generateSetMethods(*setStr, nil)
		}

		var sortingCode string
		// Generate sort functions
		if *sortStr != "" {

			generatedCode = strings.Replace(generatedCode,
				`import "sync" `,
				`import "sort" 
import "sync"`, -1)
			sortingCode = generateSortMethods(*sortStr, nil)
		}

		// Generate set and sort functions for all the types of struct
		if *sortStr == "" || *setStr == "" {

			// 2nd condition is ignoring auto generation of sort functions for `all fp in 1 place strategy"
			// Can be done in future : change the logic to find struct name(employee.Teacher) current logic is not checking .
			if *sortStr == "" && strings.Index(*types, ".") < 0 {
				generatedCode = strings.Replace(generatedCode,
					`import "sync" `,
					`import "sort" 
import "sync"`, -1)
				sortingCode = generateSortMethods(*sortStr, structToFieldsMap)
			}

			if *setStr == "" {
				if strings.Contains(strings.ToLower(fmt.Sprintf("%s", structToFieldsMap)), "time") {
					generatedCode = strings.Replace(generatedCode,
						`import "sync" `,
						`import "sync" 
import "time"`, -1)

					generatedCode = strings.Replace(generatedCode,
						`import "sync"`,
						`import "sync" 
import "time"`, -1)
				}
				setCode = generateSetMethods(*setStr, structToFieldsMap)
			}
		}

		if switchTemplateDistinct {
			generatedCode = strings.Replace(generatedCode, `import _ "reflect"`, `import "reflect"`, -1)
		}
		generatedCode = strings.Replace(generatedCode, "...Optional", "...fp.Optional", -1)
		generatedCodeIO = strings.Replace(generatedCodeIO, "...Optional", "...fp.Optional", -1)
		generatedCodeII = strings.Replace(generatedCodeII, "...Optional", "...fp.Optional", -1)
		sortingCode = strings.Replace(sortingCode, "...Optional", "...fp.Optional", -1)
		setCode = strings.Replace(setCode, "...Optional", "...fp.Optional", -1)

		f.Write([]byte(generatedCode + "\n" + generatedCodeIO + "\n" + generatedCodeII + sortingCode + setCode))
		defer f.Close()

		os.Chmod(*destination, 0444)
	}

	if !isAlreadyRun {
		fmt.Println("Functional code generation is successful.")
	}
}

// When imports are passed, Remove 1st part of "." in <FOUTPUT_TYPE> and <FINPUT_TYPE>
func removeFirstPartOfDot(str string) string {
	if strings.Contains(str, ".") {
		return strings.Split(str, ".")[1]
	}
	return str
}

func generateFPCode(pkg, dataTypes, imports string, structToFieldsMapUnexpected map[string][]string) (string, error) {
	basicTypes := "int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32, string, bool"
	conditionalType := ""
	types := strings.Split(dataTypes, ",")
	types = fp.DistinctStrIgnoreCase(types)

	template := "// Code generated by 'gofp'. DO NOT EDIT.\n"
	template += "package <PACKAGE>\n"
	template += "import _ \"reflect\" \n"
	template += "import \"sync\" \n"
	template += "import \"github.com/logic-building/functional-go/fp\" \n"

	if imports != "" {
		importList := strings.Split(imports, ",")
		importList = fp.DistinctStr(importList)
		for _, v := range importList {
			template += fmt.Sprintf("import \"%s\" \n", strings.TrimSpace(v))
		}
	}

	for _, t := range types {

		if strings.TrimSpace(strings.ToLower(t)) != strings.ToLower(pkg) {
			conditionalType = strings.TrimSpace(t)
		}
		t = strings.TrimSpace(t)

		if strings.Contains(basicTypes, t) {
			continue
		}
		r := strings.NewReplacer("<PACKAGE>", pkg, "<TYPE>", t, "<CONDITIONAL_TYPE>", removeFirstPartOfDot(conditionalType))

		// This template is used to use same function templated used for fp code for basic types
		r2 := strings.NewReplacer("<PACKAGE>", pkg, "<TYPE>", t, "<FTYPE>", removeFirstPartOfDot(conditionalType))

		firstLetterLowerCase := func(s string) string {
			if len(s) > 0 {
				return strings.ToLower(s[0:1]) + s[1:]
			}
			return s
		}

		// template for method chain
		r3 := strings.NewReplacer(
			"<PACKAGE>", pkg,
			"<TYPE>", t,
			"<FTYPE>", removeFirstPartOfDot(t),
			"<CONDITIONAL_TYPE>", removeFirstPartOfDot(conditionalType),
			"<NEWTYPE>", firstLetterLowerCase(removeFirstPartOfDot(t)))

		template = r.Replace(template)

		// it collects info of members of struct other than basic types
		complextFieldsInStructSuchAsSlice := structToFieldsMapUnexpected[t]

		if len(onlyList) > 0 {
			// Always include these functions

			template += template2.Map()
			template = r.Replace(template)

			template += template2.MapPtr()
			template = r.Replace(template)

			template += basic.MapPtrErr()
			template = r2.Replace(template)

			template += basic.MapErr()
			template = r2.Replace(template)

			template += basic.PMap()
			template = r2.Replace(template)

			template += basic.PMapPtr()
			template = r2.Replace(template)

			template += basic.PMapPtrErr()
			template = r2.Replace(template)

			template += basic.PMapErr()
			template = r2.Replace(template)

			template += template2.Filter()
			template = r.Replace(template)

			template += template2.FilterPtr()
			template = r.Replace(template)

			template += basic.FilterPtrErr()
			template = r2.Replace(template)

			template += basic.FilterErr()
			template = r2.Replace(template)

			template += template3.MethodChainStruct()
			template = r3.Replace(template)

			if fp.ExistsStrIgnoreCase("Remove", onlyList) {
				template += template2.Remove()
				template = r.Replace(template)

				template += template3.MethodChainStructForRemove()
				template = r3.Replace(template)
			}

			if fp.ExistsStrIgnoreCase("RemovePtr", onlyList) {
				template += template2.RemovePtr()
				template = r.Replace(template)

				template += template3.MethodChainStructForRemovePtr()
				template = r3.Replace(template)
			}

			if fp.ExistsStrIgnoreCase("RemovePtrErr", onlyList) {
				template += basic.RemovePtrErr()
				template = r2.Replace(template)
			}

			if fp.ExistsStrIgnoreCase("RemoveErr", onlyList) {
				template += basic.RemoveErr()
				template = r2.Replace(template)
			}

			if fp.ExistsStrIgnoreCase("Some", onlyList) {
				template += template2.Some()
				template = r.Replace(template)
			}

			if fp.ExistsStrIgnoreCase("SomePtr", onlyList) {
				template += template2.SomePtr()
				template = r.Replace(template)
			}

			if fp.ExistsStrIgnoreCase("SomePtrErr", onlyList) {
				template += basic.SomePtrErr()
				template = r2.Replace(template)
			}

			if fp.ExistsStrIgnoreCase("SomeErr", onlyList) {
				template += basic.SomeErr()
				template = r2.Replace(template)
			}

			if fp.ExistsStrIgnoreCase("Every", onlyList) {
				template += template2.Every()
				template = r.Replace(template)
			}

			if fp.ExistsStrIgnoreCase("EveryPtr", onlyList) {
				template += template2.EveryPtr()
				template = r.Replace(template)
			}

			if fp.ExistsStrIgnoreCase("EveryPtrErr", onlyList) {
				template += basic.EveryPtrErr()
				template = r2.Replace(template)
			}

			if fp.ExistsStrIgnoreCase("EveryErr", onlyList) {
				template += basic.EveryErr()
				template = r2.Replace(template)
			}

			if fp.ExistsStrIgnoreCase("DropWhile", onlyList) {
				template += template2.DropWhile()
				template = r.Replace(template)
			}

			if fp.ExistsStrIgnoreCase("DropWhilePtr", onlyList) {
				template += template2.DropWhilePtr()
				template = r.Replace(template)
			}

			if fp.ExistsStrIgnoreCase("DropWhilePtrErr", onlyList) {
				template += basic.DropWhilePtrErr()
				template = r2.Replace(template)
			}

			if fp.ExistsStrIgnoreCase("DropWhileErr", onlyList) {
				template += basic.DropWhileErr()
				template = r2.Replace(template)
			}

			if fp.ExistsStrIgnoreCase("TakeWhile", onlyList) {
				template += template2.TakeWhile()
				template = r.Replace(template)
			}

			if fp.ExistsStrIgnoreCase("TakeWhilePtr", onlyList) {
				template += template2.TakeWhilePtr()
				template = r.Replace(template)
			}

			if fp.ExistsStrIgnoreCase("TakeWhilePtrErr", onlyList) {
				template += basic.TakeWhilePtrErr()
				template = r2.Replace(template)
			}

			if fp.ExistsStrIgnoreCase("TakeWhileErr", onlyList) {
				template += basic.TakeWhileErr()
				template = r2.Replace(template)
			}

			if fp.ExistsStrIgnoreCase("FilterMap", onlyList) {
				template += template2.FilterMap()
				template = r.Replace(template)
			}

			if fp.ExistsStrIgnoreCase("FilterMapPtr", onlyList) {
				template += template2.FilterMapPtr()
				template = r.Replace(template)
			}

			if fp.ExistsStrIgnoreCase("FilterMapPtrErr", onlyList) {
				template += basic.FilterMapPtrErr()
				template = r2.Replace(template)
			}

			if fp.ExistsStrIgnoreCase("FilterMapErr", onlyList) {
				template += basic.FilterMapErr()
				template = r2.Replace(template)
			}

			if fp.ExistsStrIgnoreCase("Rest", onlyList) {
				template += template2.RestPtr()
				template = r.Replace(template)
			}

			if fp.ExistsStrIgnoreCase("RestPtr", onlyList) {
				template += template2.RestPtr()
				template = r.Replace(template)
			}

			if fp.ExistsStrIgnoreCase("Reduce", onlyList) {
				template += template2.Reduce()
				template = r.Replace(template)
			}

			if fp.ExistsStrIgnoreCase("ReducePtr", onlyList) {
				template += template2.ReducePtr()
				template = r.Replace(template)
			}

			if fp.ExistsStrIgnoreCase("ReducePtrErr", onlyList) {
				template += basic.ReducePtrErr()
				template = r2.Replace(template)
			}

			if fp.ExistsStrIgnoreCase("ReduceErr", onlyList) {
				template += basic.ReduceErr()
				template = r2.Replace(template)
			}

			if fp.ExistsStrIgnoreCase("DropLast", onlyList) {
				template += template2.DropLast()
				template = r.Replace(template)
			}

			if fp.ExistsStrIgnoreCase("DropLastPtr", onlyList) {
				template += template2.DropLastPtr()
				template = r.Replace(template)
			}

			if fp.ExistsStrIgnoreCase("Reverse", onlyList) {
				template += basic.Reverse()
				template = r2.Replace(template)
			}

			if fp.ExistsStrIgnoreCase("ReversePtr", onlyList) {
				template += basic.ReversePtr()
				template = r2.Replace(template)
			}

			if fp.ExistsStrIgnoreCase("Take", onlyList) {
				template += basic.Take()
				template = r2.Replace(template)
			}

			if fp.ExistsStrIgnoreCase("TakePtr", onlyList) {
				template += basic.TakePtr()
				template = r2.Replace(template)
			}

			if len(complextFieldsInStructSuchAsSlice) > 1 {
				switchTemplateDistinct = true

				if fp.ExistsStrIgnoreCase("DistinctP", onlyList) {
					template += basic.DistinctP2()
					template = r2.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("DistinctPPtr", onlyList) {
					template += basic.DistinctPPtr2()
					template = r2.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("Distinct", onlyList) {
					template += basic.Distinct2()
					template = r2.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("DistinctPtr", onlyList) {
					template += basic.DistinctPtr2()
					template = r2.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("Union", onlyList) {
					template += basic.Union2()
					template = r2.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("UnionPtr", onlyList) {
					template += basic.Union2Ptr()
					template = r2.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("Intersection", onlyList) {
					template += basic.Intersection2()
					template = r2.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("IntersectionPtr", onlyList) {
					template += basic.IntersectionPtr2()
					template = r2.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("Difference", onlyList) {
					template += basic.Difference2()
					template = r2.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("DifferencePtr", onlyList) {
					template += basic.DifferencePtr2()
					template = r2.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("Subset", onlyList) {
					template += basic.Subset2()
					template = r2.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("SubsetPtr", onlyList) {
					template += basic.SubsetPtr2()
					template = r2.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("Superset", onlyList) {
					template += basic.Superset2()
					template = r2.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("SupersetPtr", onlyList) {
					template += basic.SupersetPtr2()
					template = r2.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("Set", onlyList) {
					template += basic.Set2()
					template = r2.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("SetPtr", onlyList) {
					template += basic.SetPtr2()
					template = r2.Replace(template)
				}

			} else {
				if fp.ExistsStrIgnoreCase("DistinctP", onlyList) {
					template += basic.DistinctP()
					template = r2.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("DistinctPPtr", onlyList) {
					template += basic.DistinctPPtr()
					template = r2.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("Distinct", onlyList) {
					template += basic.Distinct()
					template = r2.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("DistinctPtr", onlyList) {
					template += basic.DistinctPtr()
					template = r2.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("Union", onlyList) {
					template += basic.Union()
					template = r2.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("UnionPtr", onlyList) {
					template += basic.UnionPtr()
					template = r2.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("Intersection", onlyList) {
					template += basic.Intersection()
					template = r2.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("IntersectionPtr", onlyList) {
					template += basic.IntersectionPtr()
					template = r2.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("Difference", onlyList) {
					template += basic.Difference()
					template = r2.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("DifferencePtr", onlyList) {
					template += basic.DifferencePtr()
					template = r2.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("Subset", onlyList) {
					template += basic.Subset()
					template = r2.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("SubsetPtr", onlyList) {
					template += basic.SubsetPtr()
					template = r2.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("Superset", onlyList) {
					template += basic.Superset()
					template = r2.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("SupersetPtr", onlyList) {
					template += basic.SupersetPtr()
					template = r2.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("Set", onlyList) {
					template += basic.Set()
					template = r2.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("SetPtr", onlyList) {
					template += basic.SetPtr()
					template = r2.Replace(template)
				}
			}

		} else {
			template += template2.Map()
			template = r.Replace(template)

			template += template2.MapPtr()
			template = r.Replace(template)

			template += basic.MapPtrErr()
			template = r2.Replace(template)

			template += basic.MapErr()
			template = r2.Replace(template)

			template += template2.Filter()
			template = r.Replace(template)

			template += template2.FilterPtr()
			template = r.Replace(template)

			template += basic.FilterPtrErr()
			template = r2.Replace(template)

			template += basic.FilterErr()
			template = r2.Replace(template)

			template += template2.Remove()
			template = r.Replace(template)

			template += template2.RemovePtr()
			template = r.Replace(template)

			template += basic.RemovePtrErr()
			template = r2.Replace(template)

			template += basic.RemoveErr()
			template = r2.Replace(template)

			template += template2.Some()
			template = r.Replace(template)

			template += template2.SomePtr()
			template = r.Replace(template)

			template += basic.SomePtrErr()
			template = r2.Replace(template)

			template += basic.SomeErr()
			template = r2.Replace(template)

			template += template2.Every()
			template = r.Replace(template)

			template += template2.EveryPtr()
			template = r.Replace(template)

			template += basic.EveryPtrErr()
			template = r2.Replace(template)

			template += basic.EveryErr()
			template = r2.Replace(template)

			template += template2.DropWhile()
			template = r.Replace(template)

			template += template2.DropWhilePtr()
			template = r.Replace(template)

			template += basic.DropWhilePtrErr()
			template = r2.Replace(template)

			template += basic.DropWhileErr()
			template = r2.Replace(template)

			template += template2.TakeWhile()
			template = r.Replace(template)

			template += template2.TakeWhilePtr()
			template = r.Replace(template)

			template += basic.TakeWhilePtrErr()
			template = r2.Replace(template)

			template += basic.TakeWhileErr()
			template = r2.Replace(template)

			template += basic.PMap()
			template = r2.Replace(template)

			template += basic.PMapPtr()
			template = r2.Replace(template)

			template += basic.PMapPtrErr()
			template = r2.Replace(template)

			template += basic.PMapErr()
			template = r2.Replace(template)

			template += template2.FilterMap()
			template = r.Replace(template)

			template += template2.FilterMapPtr()
			template = r.Replace(template)

			template += basic.FilterMapPtrErr()
			template = r2.Replace(template)

			template += basic.FilterMapErr()
			template = r2.Replace(template)

			template += template2.Rest()
			template = r.Replace(template)

			template += template2.RestPtr()
			template = r.Replace(template)

			template += template2.Reduce()
			template = r.Replace(template)

			template += template2.ReducePtr()
			template = r.Replace(template)

			template += basic.ReducePtrErr()
			template = r2.Replace(template)

			template += basic.ReduceErr()
			template = r2.Replace(template)

			template += template2.DropLast()
			template = r.Replace(template)

			template += template2.DropLastPtr()
			template = r.Replace(template)

			template += basic.Reverse()
			template = r2.Replace(template)

			template += basic.ReversePtr()
			template = r2.Replace(template)

			template += basic.Take()
			template = r2.Replace(template)

			template += basic.TakePtr()
			template = r2.Replace(template)

			template += template3.MethodChainStruct()
			template = r3.Replace(template)

			template += template3.MethodChainStructForRemove()
			template = r3.Replace(template)

			template += template3.MethodChainStructForRemovePtr()
			template = r3.Replace(template)

			// if struct's has member of type other than basic types such as list then use template which uses reflect
			if len(complextFieldsInStructSuchAsSlice) > 1 {
				switchTemplateDistinct = true

				template += basic.DistinctP2()
				template = r2.Replace(template)

				template += basic.DistinctPPtr2()
				template = r2.Replace(template)

				template += basic.Distinct2()
				template = r2.Replace(template)

				template += basic.DistinctPtr2()
				template = r2.Replace(template)

				template += basic.Union2()
				template = r2.Replace(template)

				template += basic.Union2Ptr()
				template = r2.Replace(template)

				template += basic.Intersection2()
				template = r2.Replace(template)

				template += basic.IntersectionPtr2()
				template = r2.Replace(template)

				template += basic.Difference2()
				template = r2.Replace(template)

				template += basic.DifferencePtr2()
				template = r2.Replace(template)

				template += basic.Subset2()
				template = r2.Replace(template)

				template += basic.SubsetPtr2()
				template = r2.Replace(template)

				template += basic.Superset2()
				template = r2.Replace(template)

				template += basic.SupersetPtr2()
				template = r2.Replace(template)

				template += basic.Set2()
				template = r2.Replace(template)

				template += basic.SetPtr2()
				template = r2.Replace(template)

			} else {
				template += basic.DistinctP()
				template = r2.Replace(template)

				template += basic.DistinctPPtr()
				template = r2.Replace(template)

				template += basic.Distinct()
				template = r2.Replace(template)

				template += basic.DistinctPtr()
				template = r2.Replace(template)

				template += basic.Union()
				template = r2.Replace(template)

				template += basic.UnionPtr()
				template = r2.Replace(template)

				template += basic.Intersection()
				template = r2.Replace(template)

				template += basic.IntersectionPtr()
				template = r2.Replace(template)

				template += basic.Difference()
				template = r2.Replace(template)

				template += basic.DifferencePtr()
				template = r2.Replace(template)

				template += basic.Subset()
				template = r2.Replace(template)

				template += basic.SubsetPtr()
				template = r2.Replace(template)

				template += basic.Superset()
				template = r2.Replace(template)

				template += basic.SupersetPtr()
				template = r2.Replace(template)

				template += basic.Set()
				template = r2.Replace(template)

				template += basic.SetPtr()
				template = r2.Replace(template)
			}
		}
	}
	return template, nil
}

func generateFPCodeIO(pkg, dataTypes string) (string, error) {
	basicTypes := "int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32, string, bool"
	template := ""
	types := strings.Split(dataTypes, ",")

	types = fp.DistinctStrIgnoreCase(types)

	// For different input output combination
	for _, inputType := range types {
		for _, outputType := range types {

			inputType = strings.TrimSpace(inputType)
			outputType = strings.TrimSpace(outputType)

			if inputType == outputType || (strings.Contains(basicTypes, inputType) && strings.Contains(basicTypes, outputType)) {
				continue
			}

			if strings.Contains(basicTypes, strings.ToLower(inputType)) {
				inputType = strings.ToLower(inputType)
			}

			if strings.Contains(basicTypes, strings.ToLower(outputType)) {
				outputType = strings.ToLower(outputType)
			}

			fInputType := strings.Title(inputType)
			fOutputType := strings.Title(outputType)

			if fInputType == "String" {
				fInputType = "Str"
			}

			if fOutputType == "String" {
				fOutputType = "Str"
			}

			r := strings.NewReplacer("<FINPUT_TYPE>", removeFirstPartOfDot(fInputType), "<FOUTPUT_TYPE>", removeFirstPartOfDot(fOutputType), "<INPUT_TYPE>", inputType, "<OUTPUT_TYPE>", outputType)

			if len(onlyList) > 0 {
				if fp.ExistsStrIgnoreCase("MapIO", onlyList) {
					template += basic.MapIO()
					template = r.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("MapIOErr", onlyList) {
					template += basic.MapIOErr()
					template = r.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("MapIOPtr", onlyList) {
					template += basic.MapIOPtr()
					template = r.Replace(template)
				}
				if fp.ExistsStrIgnoreCase("MapIOPtrErr", onlyList) {
					template += basic.MapIOPtrErr()
					template = r.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("PMapIO", onlyList) {
					template += basic.PMapIO()
					template = r.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("PMapIOErr", onlyList) {
					template += basic.PMapIOErr()
					template = r.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("PMapIOPtr", onlyList) {
					template += basic.PMapIOPtr()
					template = r.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("PMapIOPtrErr", onlyList) {
					template += basic.PMapIOPtrErr()
					template = r.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("FilterMapIO", onlyList) {
					template += basic.FilterMapIO()
					template = r.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("FilterMapIOPtrErr", onlyList) {
					template += basic.FilterMapIOPtrErr()
					template = r.Replace(template)
				}

				if fp.ExistsStrIgnoreCase("FilterMapIOErr", onlyList) {
					template += basic.FilterMapIOErr()
					template = r.Replace(template)
				}
			} else {
				template += basic.MapIO()
				template = r.Replace(template)

				template += basic.MapIOErr()
				template = r.Replace(template)

				template += basic.MapIOPtr()
				template = r.Replace(template)

				template += basic.MapIOPtrErr()
				template = r.Replace(template)

				template += basic.PMapIO()
				template = r.Replace(template)

				template += basic.PMapIOErr()
				template = r.Replace(template)

				template += basic.PMapIOPtr()
				template = r.Replace(template)

				template += basic.PMapIOPtrErr()
				template = r.Replace(template)

				template += basic.FilterMapIO()
				template = r.Replace(template)

				template += basic.FilterMapIOPtrErr()
				template = r.Replace(template)

				template += basic.FilterMapIOErr()
				template = r.Replace(template)
			}
		}
	}

	return template, nil
}

func generateFPCodeII(pkg, dataTypes string) (string, error) {
	basicTypes := "int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32, string, bool"
	template := ""
	types := strings.Split(dataTypes, ",")

	types = fp.DistinctStrIgnoreCase(types)

	// For different input output combination
	for _, inputType1 := range types {
		for _, inputType2 := range types {

			inputType1 = strings.TrimSpace(inputType1)
			inputType2 = strings.TrimSpace(inputType2)

			// Skip same basic data type
			if strings.Contains(basicTypes, inputType1) && strings.Contains(basicTypes, inputType2) {
				continue
			}

			if strings.Contains(basicTypes, strings.ToLower(inputType1)) {
				inputType1 = strings.ToLower(inputType1)
			}

			if strings.Contains(basicTypes, strings.ToLower(inputType2)) {
				inputType2 = strings.ToLower(inputType2)
			}

			fInputType1 := strings.Title(inputType1)
			fInputType2 := strings.Title(inputType2)

			if fInputType1 == "String" {
				fInputType1 = "Str"
			}

			if fInputType2 == "String" {
				fInputType2 = "Str"
			}

			// Standard function name for same user defined type
			if inputType1 == inputType2 && strings.ToLower(inputType1) == strings.ToLower(pkg) {
				fInputType1 = ""
				fInputType2 = ""
			}

			// Standard function name for same user defined type
			if inputType1 == inputType2 {
				fInputType2 = ""
			}

			r := strings.NewReplacer("<FINPUT_TYPE1>", removeFirstPartOfDot(fInputType1), "<FINPUT_TYPE2>", removeFirstPartOfDot(fInputType2), "<INPUT_TYPE1>", inputType1, "<INPUT_TYPE2>", inputType2)

			template += basic.Merge()
			template = r.Replace(template)

			template += basic.Zip()
			template = r.Replace(template)
		}
	}

	return template, nil
}

func usage() {
	fmt.Println("\nUsage:")
	fmt.Println("go:generate -destination fp.go -source employee.go -pkg Employee")
}
func quoteForTheDay() string {
	quotes := []string{
		"Time spent in love is never waste",
		"Enjoy every moment",
		"Wherever there is love, there is God",
		"The real way to loving is giving not demanding",
		"No one is greater or smaller than other. Everyone in this world is unique. Love everyone",
		"The person who has heart full of love, has always something to give",
		"Hell has three gates lust, anger, greed",
		"Be happy with nothing and you will be happy with everything",
		"Detachment is not that you should own nothing, but that nothing should own you",
		"Devotion has the power to burn down any karma",
		"Love is the greatest power on earth",
		"When you wish good for others, good things come back to you. This is the Law of Nature",
		"If you can win over your mind, you can win over the whole world",
		"Darkness cannot drive out darkness, only light can do that. Hate cannot drive out hate. only love can do that",
		"Silence says so much. Just listen",
		"The greatest gift human can give to himself and others are tolerance and forgiveness",
		"The practice of devotion involves replacing desires for the world with the desires for God",
		"The wealth of divine love is the only true wealth. Every other form of wealth simply enhances one's pride",
		"Speak only when you feel your words are better than the silence",
		"For our spiritual growth, negative people are often placed in our path, so we may learn selfless love, forgiveness & surrender to the will of God",
		"The happiest people are givers not takers",
		"Why do we close our eyes when we pray, cry, kiss or dream? Because the most beautiful things in life are not seen, but felt by the heart",
		"If you have to choose between being kind and being right choose being kind and you will always be right",
		"Silence & Smile are two powerful tools.Smile is the way to solve many problems & Silence is the way to avoid many problems",
		"Don't get upset with people and situations, because both are powerless without your reaction",
		"Most of the people are in lack of knowledge.Don't hate them.Love them and understand that they are under influence of ignorance. Always do righteously.",
		"Every way and means that leads our mind to God is Devotion",
		"The Only Purpose of Our Human Life is to Attain God",
		"Meditation. Because some questions can't be answered by Google!",
		"This is the nature of existence - if you do the right things, the right things will happen to you",
		"Devotion is the only way to be liberated from material attachment. It is only then that we become free from lust, anger and greed",
		"I belong to no religion. My religion is love. Every heart is my temple",
		"Don't focus too much on the defects, be aware of them, but our endeavor should be towards positive",
		"To purify the mind, we must engage in devotion to the lord, When our mind is purified, out attitude and our experience will change towards the outer world",
		"The reason that we are in a state of suffering and we are enveloped in the darkness of material existence, is our forgetfulness of God",
		"If you can establish your relationship with God, that ultimate satisfaction that you have been searching for since innumerable lifetimes, will eventually be attained",
		"The Joy of the mind is the measure of its strength",
		"When you come to a point where you have no need to impress anybody, your freedom will begin",
		"Ritualistic worship, chanting and meditation are done with the body, voice and the mind: they excel each other in the ascending order - Ramana Maharshi",
		"Uttering the sacred word, either in a loud or low tone is preferable to chants in praise of the Supreme. Mental contemplation is superior to both",
		"When one learns to turn the mind away from material allurements and renounces the desires of the senses, such a person comes in touch with the inner bliss of the soul",
		"When we decide that God is ours and the whole world is His, then our consciousness transforms from seeking self-enjoyment to serving the Lord with everything that we have",
		"If you want to change the world, go home and love your family - Mother Teresa",
		"Every time you smile at someone, it is an action of love, a gift to that person, a beautiful thing. - Mother Teresa",
		"No color, no religion, no nationality should come between us, we are all children of God. - Mother Teresa",
		"In this life, we cannot always do great things. But we can do small things with great love. - Mother Teresa",
		`Lord, make an instrument of the peace,
    Where there is hatred, let me show love;
    Where there is injury, pardon;
    Where there is doubt, faith;
    Where there is despair, hope;
    Where there is darkness, light;
    Where there is sadness, Joy.
`,
		"A generous heart, kind speech, and a life of service and compassion are the things which renew humanity, Buddha",
	}

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s) // initialize local pseudorandom generator
	ind := r.Intn(len(quotes))

	return quotes[ind]
}

func runWithin(duration time.Duration) bool {

	runWithin := func(file string, duration time.Duration) bool {
		writeToFile := func(file string) {
			f, _ := os.Create(file)
			defer f.Close()
			f.WriteString("Functional go")
		}
		runWithin := true
		if f, err := os.Stat(file); err == nil {
			modificationTime := f.ModTime()

			currentTime := time.Now()
			diffSeconds := currentTime.Sub(modificationTime).Seconds()
			if diffSeconds > duration.Seconds() {
				runWithin = false
				writeToFile(file)
			}

		} else {
			writeToFile(file)
		}
		return runWithin
	}

	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}

		return runWithin(home+"\functional-go.txt", duration)
	}
	return runWithin("/tmp/functional-go.txt", duration)
}

func generateSortMethods(sortStr string, allFields map[string][]string) string {
	if allFields != nil {
		template := ""
		for structName, fieldsAndTypes := range allFields {
			if len(fieldsAndTypes) > 0 {
				for _, fieldType := range fieldsAndTypes {

					fieldName := strings.Split(fieldType, " ")[0]
					dataType := strings.Split(fieldType, " ")[1]

					fStructName := strings.Title(structName)
					fFieldName := strings.Title(fieldName)

					if strings.Contains(strings.ToLower(dataType), "*time") {
						r := strings.NewReplacer("<STRUCT_NAME>", structName, "<FIELD_NAME>", fieldName, "<FSTRUCT_NAME>", fStructName, "<FFIELD_NAME>", fFieldName)
						template += template2.SortStructForTimeFieldPtr()
						template = r.Replace(template)
						continue
					}

					if strings.Contains(strings.ToLower(dataType), "time") {
						r := strings.NewReplacer("<STRUCT_NAME>", structName, "<FIELD_NAME>", fieldName, "<FSTRUCT_NAME>", fStructName, "<FFIELD_NAME>", fFieldName)
						template += template2.SortStructForTimeField()
						template = r.Replace(template)
						continue
					}

					// "*int", "*int64", "*int32", "*int16", "*int8", "*uint", "*uint64", "*uint32", "*uint16", "*uint8", "*float64", "*float32", "*string"
					if strings.Contains(strings.ToLower(dataType), "*int") ||
						strings.Contains(strings.ToLower(dataType), "*int64") ||
						strings.Contains(strings.ToLower(dataType), "*int32") ||
						strings.Contains(strings.ToLower(dataType), "*int16") ||
						strings.Contains(strings.ToLower(dataType), "*int8") ||
						strings.Contains(strings.ToLower(dataType), "*uint") ||
						strings.Contains(strings.ToLower(dataType), "*uint64") ||
						strings.Contains(strings.ToLower(dataType), "*uint32") ||
						strings.Contains(strings.ToLower(dataType), "*uint16") ||
						strings.Contains(strings.ToLower(dataType), "*uint8") ||
						strings.Contains(strings.ToLower(dataType), "*float64") ||
						strings.Contains(strings.ToLower(dataType), "*float32") ||
						strings.Contains(strings.ToLower(dataType), "*string") {
						r := strings.NewReplacer("<STRUCT_NAME>", structName, "<FIELD_NAME>", fieldName, "<FSTRUCT_NAME>", fStructName, "<FFIELD_NAME>", fFieldName)
						template += template2.SortStructPtr()
						template = r.Replace(template)
						continue
					}

					r := strings.NewReplacer("<STRUCT_NAME>", structName, "<FIELD_NAME>", fieldName, "<FSTRUCT_NAME>", fStructName, "<FFIELD_NAME>", fFieldName)
					template += template2.SortStruct()
					template = r.Replace(template)

				}
			}
		}
		return template
	}
	if len(strings.TrimSpace(sortStr)) == 0 {
		fmt.Println("-sort: value is empty. ignoring sort methods")
		return ""
	}

	template := ""
	sortedStrList := strings.Split(sortStr, ",")

	for _, sortedStrItem := range sortedStrList {
		sortedStrItem = strings.TrimSpace(sortedStrItem)
		sortStrItemElements := strings.Split(sortedStrItem, ":")
		if len(sortStrItemElements) < 2 || len(sortStrItemElements) > 3 {
			fmt.Println("-sort: format is not valid. expected format for option -sort: <struct_name>:<field_name>. eg. -sort=\"Employee:Salary\"")
			fmt.Println("ignoring sort methods")
			return ""
		}

		structName := strings.TrimSpace(sortStrItemElements[0])
		fieldName := strings.TrimSpace(sortStrItemElements[1])

		fStructName := strings.Title(structName)
		fFieldName := strings.Title(fieldName)

		if len(sortStrItemElements) == 3 {
			dataType := strings.TrimSpace(sortStrItemElements[2])
			if strings.Contains(strings.ToLower(dataType), "time") {
				r := strings.NewReplacer("<STRUCT_NAME>", structName, "<FIELD_NAME>", fieldName, "<FSTRUCT_NAME>", fStructName, "<FFIELD_NAME>", fFieldName)
				template += template2.SortStructForTimeField()
				template = r.Replace(template)
				continue
			}
		}

		r := strings.NewReplacer("<STRUCT_NAME>", structName, "<FIELD_NAME>", fieldName, "<FSTRUCT_NAME>", fStructName, "<FFIELD_NAME>", fFieldName)
		template += template2.SortStruct()
		template = r.Replace(template)
	}

	return template
}

func generateSetMethods(setStr string, allFields map[string][]string) string {

	if allFields != nil {
		template := ""
		for structName, fieldsAndTypes := range allFields {
			if len(fieldsAndTypes) > 0 {
				for _, fieldType := range fieldsAndTypes {

					fieldName := strings.Split(fieldType, " ")[0]
					fieldType := strings.Split(fieldType, " ")[1]

					fStructName := strings.Title(structName)
					fFieldName := strings.Title(fieldName)

					r := strings.NewReplacer("<STRUCT_NAME>", structName, "<FIELD_NAME>", fieldName, "<FSTRUCT_NAME>", fStructName, "<FFIELD_NAME>", fFieldName, "<FIELD_TYPE>", fieldType)
					template += template2.SetStruct()
					template = r.Replace(template)

				}
			}
		}
		return template
	}
	if len(strings.TrimSpace(setStr)) == 0 {
		fmt.Println("-set: value is empty. ignoring set methods")
		return ""
	}

	template := ""
	sortedStrList := strings.Split(setStr, ",")

	for _, sortedStrItem := range sortedStrList {
		sortedStrItem = strings.TrimSpace(sortedStrItem)
		sortStrItemElements := strings.Split(sortedStrItem, ":")
		if len(sortStrItemElements) != 3 {
			fmt.Println("-set: format is not valid. expected format for option -sort: <struct_name>:<field_name>:<field_type>. eg. -set=\"Employee:Salary:float64\"")
			fmt.Println("ignoring set methods")
			return ""
		}

		structName := strings.TrimSpace(sortStrItemElements[0])
		fieldName := strings.TrimSpace(sortStrItemElements[1])
		fieldType := strings.TrimSpace(sortStrItemElements[2])

		fStructName := strings.Title(structName)
		fFieldName := strings.Title(fieldName)

		r := strings.NewReplacer("<STRUCT_NAME>", structName, "<FIELD_NAME>", fieldName, "<FSTRUCT_NAME>", fStructName, "<FFIELD_NAME>", fFieldName, "<FIELD_TYPE>", fieldType)
		template += template2.SetStruct()
		template = r.Replace(template)
	}

	return template
}

func listDir(dirName string) ([]string, error) {
	var files []string

	root := dirName
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func findStructNamesAndFieldsGivenInGoGenerate() (map[string][]string, map[string][]string) {
	allTypesInGoGenerate := strings.Split(*types, ",")

	isUserDefinedType := func(dataType string) bool {
		switch strings.ToLower(strings.TrimSpace(dataType)) {
		case "int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "float64", "float32", "string", "bool":
			return false
		}
		return true
	}
	userDefinedTypesInGoGenerate := fp.MapStr(strings.TrimSpace, fp.FilterStr(isUserDefinedType, allTypesInGoGenerate))

	structToFieldsMap := make(map[string][]string, len(userDefinedTypesInGoGenerate))
	structToFieldsMapUnExpected := make(map[string][]string, len(userDefinedTypesInGoGenerate))
	structToFieldsMapIndex := 0

	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	files, err := listDir(path)
	if err != nil {
		fmt.Printf("\n error scanning current folder=%v, error=%v. sort and set methods will be skipped", path, err)
		return nil, nil
	}

	onlyGoFiles := fp.FilterStr(func(str string) bool {
		return strings.Contains(str, ".go")
	}, files)

	totalStructCount := 0
	for _, fileStr := range onlyGoFiles {

		if totalStructCount == len(userDefinedTypesInGoGenerate) {
			break
		}

		file, err := os.Open(fileStr)
		if err != nil {
			fmt.Printf("\n error reading file=%s to generate sort and set methods. skipping set and sort functions", fileStr)
			return nil, nil
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		packageFound := false
		startCollectingStructInfo := false
		var structFields []string
		var structFieldsUnExpected []string

		for scanner.Scan() {
			txtLine := scanner.Text()
			if len(txtLine) > 0 && strings.Contains(txtLine, *pkgName) {
				packageFound = true
			}
			// reading lines of file of package mentioned in go:generate
			if packageFound {

				words := strings.Fields(txtLine)

				// Found struct
				if len(words) == 4 && strings.Contains(words[0], "type") && fp.ExistsStr(words[1], userDefinedTypesInGoGenerate) && strings.Contains(words[2], "struct") && strings.Contains(words[3], "{") {
					startCollectingStructInfo = true
					totalStructCount++
				}

				if startCollectingStructInfo && strings.TrimSpace(txtLine) == "}" {
					startCollectingStructInfo = false

					newStructFieldsArr := make([]string, len(structFields))
					for i, v := range structFields {
						newStructFieldsArr[i] = v
					}

					newStructFieldsArrUnExpected := make([]string, len(structFieldsUnExpected))
					copy(newStructFieldsArrUnExpected, structFieldsUnExpected)

					structToFieldsMap[userDefinedTypesInGoGenerate[structToFieldsMapIndex]] = newStructFieldsArr
					structToFieldsMapUnExpected[userDefinedTypesInGoGenerate[structToFieldsMapIndex]] = newStructFieldsArrUnExpected
					structToFieldsMapIndex++
					structFields = make([]string, 0)
					structFieldsUnExpected = make([]string, 0)
				}

				if startCollectingStructInfo {
					if len(words) >= 2 {
						field := strings.TrimSpace(words[0])
						dataType := strings.TrimSpace(words[1])

						switch dataType {
						case "int", "int64", "int32", "int16", "int8", "uint", "uint64", "uint32", "uint16", "uint8", "float64", "float32", "string", "time.Time", "*time.Time", "*int", "*int64", "*int32", "*int16", "*int8", "*uint", "*uint64", "*uint32", "*uint16", "*uint8", "*float64", "*float32", "*string":
							structFields = append(structFields, field+" "+dataType)
						default:
							structFieldsUnExpected = append(structFieldsUnExpected, field+" "+dataType)
						}
					}
				}
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("\n error scanning data from file %s. skipping generation of sort and set functions", fileStr)
			return nil, nil
		}
	}

	return structToFieldsMap, structToFieldsMapUnExpected
}
