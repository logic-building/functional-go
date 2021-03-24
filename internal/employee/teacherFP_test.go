package employee

import (
	"testing"
	"time"
)

func TestMethodChainWithSortTeacher(t *testing.T) {
	now := time.Now()
	teachers := []Teacher{
		{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: &now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: &now},
		{Id: 3, Name: "Radha", Salary: 900, CreationDate: &now}}

	salaryIncrement := func(emp Teacher) Teacher {
		emp.Salary = emp.Salary + 1000
		return emp
	}

	salary700 := func(emp Teacher) bool {
		return emp.Salary == 700
	}

	teachersWithIncrementedSalary := MakeTeacherSlice(teachers...).
		DropWhile(salary700).
		Map(salaryIncrement).
		SortBySalary()

	if 1800 != teachersWithIncrementedSalary[0].Salary {
		t.Error("error in method chain with Map", teachers)
	}

	employeeDescendingOrder := MakeTeacherSlice(teachers...).
		DropWhile(salary700).
		Map(salaryIncrement).
		SortBySalaryDesc()

	if 1900 != employeeDescendingOrder[0].Salary {
		t.Error("error in method chain with Map", teachers)
	}
}

func TestMethodChainWithSortTeacherByCreationDate(t *testing.T) {
	now1 := time.Now()
	now2 := time.Now()
	now3 := time.Now()
	now4 := time.Now()
	teachers := []Teacher{
		{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now1},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: &now2},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: &now3},
		{Id: 3, Name: "Radha", Salary: 900, CreationDate: &now4}}

	salaryIncrement := func(emp Teacher) Teacher {
		emp.Salary = emp.Salary + 1000
		return emp
	}

	salary700 := func(emp Teacher) bool {
		return emp.Salary == 700
	}

	teachersWithIncrementedSalary := MakeTeacherSlice(teachers...).
		DropWhile(salary700).
		Map(salaryIncrement).
		SortBySalary().SortByCreationDateDesc()

	if 1900 != teachersWithIncrementedSalary[0].Salary {
		t.Error("error in method chain with Map", teachers)
	}

	employeeDescendingOrder := MakeTeacherSlice(teachers...).
		DropWhile(salary700).
		Map(salaryIncrement).
		SortBySalaryDesc()

	if 1900 != employeeDescendingOrder[0].Salary {
		t.Error("error in method chain with Map", teachers)
	}
}

func TestMethodChainWithSortTeacherByAddress(t *testing.T) {
	now1 := time.Now()
	now2 := time.Now()
	now3 := time.Now()
	now4 := time.Now()

	ramAddress := "Ayodhya"
	shyamAddress := "Gokul"
	radhaAddress := "Barsana"
	teachers := []Teacher{
		{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now1, Address: &ramAddress},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: &now2, Address: &shyamAddress},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: &now3, Address: &shyamAddress},
		{Id: 3, Name: "Radha", Salary: 900, CreationDate: &now4, Address: &radhaAddress},
	}

	salaryIncrement := func(emp Teacher) Teacher {
		emp.Salary = emp.Salary + 1000
		return emp
	}

	salary700 := func(emp Teacher) bool {
		return emp.Salary == 700
	}

	teachersWithIncrementedSalary := MakeTeacherSlice(teachers...).
		DropWhile(salary700).
		Map(salaryIncrement).
		SortBySalary().SortByCreationDateDesc()

	if 1900 != teachersWithIncrementedSalary[0].Salary {
		t.Error("error in method chain with Map", teachers)
	}

	employeeDescendingOrder := MakeTeacherSlice(teachers...).
		Map(salaryIncrement).
		SortByAddress()

	if 1700 != employeeDescendingOrder[0].Salary {
		t.Error("error in method chain with Map", teachers)
	}

	employeeDescendingOrder2 := MakeTeacherSlice(teachers...).
		Map(salaryIncrement).
		SortByAddressDesc()

	if 1800 != employeeDescendingOrder2[0].Salary {
		t.Error("error in method chain with Map", employeeDescendingOrder2[0].Salary)
	}
}

func TestMethodChainWithSortTeacherByCreationDatePtr(t *testing.T) {
	now1 := time.Now()
	now2 := time.Now()
	now3 := time.Now()
	now4 := time.Now()
	teachers := []*Teacher{
		{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now1},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: &now2},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: &now3},
		{Id: 3, Name: "Radha", Salary: 900, CreationDate: &now4}}

	salaryIncrement := func(emp *Teacher) *Teacher {
		emp.Salary = emp.Salary + 1000
		return emp
	}

	salary700 := func(emp *Teacher) bool {
		return emp.Salary == 700
	}

	teachersWithIncrementedSalary := MakeTeacherSlicePtr(teachers...).
		DropWhilePtr(salary700).
		MapPtr(salaryIncrement).
		SortByCreationDateDescPtr()

	if 1900 != teachersWithIncrementedSalary[0].Salary {
		t.Error("error in method chain with Map", teachers)
	}

}

func TestMethodChainWithSortTeacherBySalaryPtr(t *testing.T) {
	now1 := time.Now()
	now2 := time.Now()
	now3 := time.Now()
	now4 := time.Now()
	teachers := []*Teacher{
		{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now1},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: &now2},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: &now3},
		{Id: 3, Name: "Radha", Salary: 900, CreationDate: &now4}}

	salaryIncrement := func(emp *Teacher) *Teacher {
		emp.Salary = emp.Salary + 1000
		return emp
	}

	salary700 := func(emp *Teacher) bool {
		return emp.Salary == 700
	}

	teachersWithIncrementedSalary := MakeTeacherSlicePtr(teachers...).
		DropWhilePtr(salary700).
		MapPtr(salaryIncrement).
		SortBySalaryDescPtr()

	if 1900 != teachersWithIncrementedSalary[0].Salary {
		t.Error("error in method chain with Map", teachers)
	}

}
