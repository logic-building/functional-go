// Generates functional code locally for user defined data type.
/*
```
1. Install "gofp" to generate code
   go get github.com/logic-building/functional-go/gofp
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
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	template2 "github.com/logic-building/functional-go/internal/template"
	"math/rand"
	"time"
)

var (
	destination = flag.String("destination", "", "functional code for user defined data type")
	pkgName     = flag.String("pkg", "", "package name for generated files")
	types       = flag.String("type", "", "user defined type")
)

func main() {
	defer func() {
		fmt.Println("\n\t\t\"" + quoteForTheDay() + "\"")
	}()

	fmt.Println("Welcome to functional-go")

	flag.Parse()

	if *destination == "" || *pkgName == "" || *types == "" {
		fmt.Println("either of these fields : (destination, package, types) - is not provided")
		usage()
		return
	}

	if len(*destination) > 0 {
		if err := os.MkdirAll(filepath.Dir(*destination), os.ModePerm); err != nil {
			log.Fatalf("Unable to create destination directory: %v", err)
			usage()
		}
		f, err := os.Create(*destination)
		if err != nil {
			log.Fatalf("Failed opening destination file: %v", err)
		}
		generatedCode, err := generateFPCode(*pkgName, *types)
		if err != nil {
			log.Fatalf("Failed code generation: %v", err)
			usage()
		}
		f.Write([]byte(generatedCode))
		defer f.Close()
	}

	fmt.Println("Code is generated successfully")

}

func generateFPCode(pkg, dataTypes string) (string, error) {
	conditionalType := ""
	types := strings.Split(dataTypes, ",")

	template := "// Code generated by 'gofp'. DO NOT EDIT.\n"
	template += "package <PACKAGE> \n"
	template += "import \"sync\" \n"

	for _, t := range types {

		if strings.TrimSpace(strings.ToLower(t)) != strings.ToLower(pkg) {
			conditionalType = strings.TrimSpace(t)
		}
		t = strings.TrimSpace(t)
		r := strings.NewReplacer("<PACKAGE>", pkg, "<TYPE>", t, "<CONDITIONAL_TYPE>", conditionalType)

		template = r.Replace(template)

		template += template2.Map()
		template = r.Replace(template)

		template += template2.Filter()
		template = r.Replace(template)

		template += template2.Remove()
		template = r.Replace(template)

		template += template2.Some()
		template = r.Replace(template)

		template += template2.Every()
		template = r.Replace(template)

		template += template2.DropWhile()
		template = r.Replace(template)

		template += template2.TakeWhile()
		template = r.Replace(template)

		template += template2.Pmap()
		template = r.Replace(template)

		template += template2.FilterMap()
		template = r.Replace(template)

		template += template2.Rest()
		template = r.Replace(template)

		template += template2.Reduce()
		template = r.Replace(template)
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
		"Darkness cannot drive out darkness, only light can do that. Hate cannot drive out hate. only love cna do that",
		"Silence says so mcuh. Just listen",
		"The greatest gift human can give to himself and others are tolerance and forgiveness",
		"The practice of devotion involves replacing desires for the world with the desires for God",
	}

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s) // initialize local pseudorandom generator
	ind := r.Intn(len(quotes))

	return quotes[ind]
}
