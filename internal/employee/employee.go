package employee

//go:generate gofp -destination fp.go -pkg employee -type "Employee, Teacher"
type Employee struct {
	Id     int
	Name   string
	Salary float64
}

type Teacher struct {
	Id     int
	Name   string
	Salary float64
}
