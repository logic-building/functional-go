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
