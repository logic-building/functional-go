package employee

import "time"

//go:generate gofp -destination fp.go -pkg employee -type "Employee, int, string" -sort "Employee:Name, Employee:Salary, Employee:CreationDate:time.Time" -set "Employee:Name:string, Employee:Salary:float64, Employee:CreationDate:time.Time"

type Employee struct {
	Id           int
	Name         string
	Salary       float64
	CreationDate time.Time
}

//go:generate gofp -destination teacherFP.go -pkg employee -type "Teacher"
type Teacher struct {
	Id           int
	Name         string
	Salary       float64
	CreationDate *time.Time
	Address      *string
	Students     []string
}
