package employer

//go:generate gofp -destination fp.go -pkg employer -type "Employer, employee.Employee, int" -imports "github.com/logic-building/functional-go/internal/employee" -sort "Employer:Id"

type Employer struct {
	Id   int
	Name string
}
