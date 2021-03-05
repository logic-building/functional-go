package employee

import (
	"fmt"
	"testing"
	"time"
)

func TestDistinct(t *testing.T) {
	now := time.Now()
	employees := []Employee{{Id: 1, Name: "Ram", Salary: 700, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 3, Name: "Radha", Salary: 900, CreationDate: now}}

	distinctEmployees := Distinct(employees)
	if len(distinctEmployees) != 3 {
		t.Errorf("Error in Distinct for Employee. actualLength=%d. expectedLength=3", len(distinctEmployees))
	}

	address := "Usa"
	teachers := []Teacher{{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 2, Name: "Shaym", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 3, Name: "Damodar", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}}}

	distinctTeachers := DistinctTeacher(teachers)
	if len(distinctTeachers) != 3 {
		t.Errorf("Error in Distinct for Employee. actualLength=%d. expectedLength=3", len(distinctTeachers))
	}

	// for pointer
	employeesPtr := []*Employee{{Id: 1, Name: "Ram", Salary: 700, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 3, Name: "Radha", Salary: 900, CreationDate: now}}

	distinctEmployeesPtr := DistinctPtr(employeesPtr)
	if len(distinctEmployeesPtr) != 3 {
		t.Errorf("Error in Distinct for Employee. actualLength=%d. expectedLength=3", len(distinctEmployeesPtr))
	}

	teachersPtr := []*Teacher{{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 2, Name: "Shaym", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 3, Name: "Damodar", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}}}

	distinctTeachersPtr := DistinctTeacherPtr(teachersPtr)
	if len(distinctTeachersPtr) != 3 {
		t.Errorf("Error in Distinct for Employee. actualLength=%d. expectedLength=3", len(distinctTeachersPtr))
	}
}

func TestDistinctP(t *testing.T) {
	now := time.Now()
	employees := []Employee{{Id: 1, Name: "Ram", Salary: 700, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 3, Name: "Radha", Salary: 900, CreationDate: now}}

	actual := DistinctP(employees)
	if actual {
		t.Errorf("Error in DistinctP for Employee. actual=%v. expected=false", actual)
	}

	address := "Usa"
	teachers := []Teacher{{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 2, Name: "Shaym", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 3, Name: "Damodar", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}}}

	actual = DistinctTeacherP(teachers)
	if actual {
		t.Errorf("Error in DistinctP for Employee. actual=%v. expected=false", actual)
	}

	// pointer
	employeesPtr := []*Employee{{Id: 1, Name: "Ram", Salary: 700, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 3, Name: "Radha", Salary: 900, CreationDate: now}}

	actualPtr := DistinctPPtr(employeesPtr)
	if actualPtr {
		t.Errorf("Error in DistinctPPtr for Employee. actual=%v. expected=false", actualPtr)
	}

	teachersPtr := []*Teacher{{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 2, Name: "Shaym", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 3, Name: "Damodar", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}}}

	actualPtr = DistinctTeacherPPtr(teachersPtr)
	if actualPtr {
		t.Errorf("Error in DistinctPPtr for Employee. actual=%v. expected=false", actual)
	}
}

func TestSet(t *testing.T) {
	now := time.Now()
	employees := []Employee{{Id: 1, Name: "Ram", Salary: 700, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 3, Name: "Radha", Salary: 900, CreationDate: now}}

	distinctEmployees := Set(employees)
	if len(distinctEmployees) != 3 {
		t.Errorf("Error in Distinct for Employee. actualLength=%d. expectedLength=3", len(distinctEmployees))
	}

	address := "Usa"
	teachers := []Teacher{{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 2, Name: "Shaym", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 3, Name: "Damodar", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}}}

	distinctTeachers := SetTeacher(teachers)
	if len(distinctTeachers) != 3 {
		t.Errorf("Error in Set for Employee. actualLength=%d. expectedLength=3", len(distinctTeachers))
	}

	// for pointer
	employeesPtr := []*Employee{{Id: 1, Name: "Ram", Salary: 700, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 3, Name: "Radha", Salary: 900, CreationDate: now}}

	distinctEmployeesPtr := SetPtr(employeesPtr)
	if len(distinctEmployeesPtr) != 3 {
		t.Errorf("Error in Set for Employee. actualLength=%d. expectedLength=3", len(distinctEmployeesPtr))
	}

	teachersPtr := []*Teacher{{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 2, Name: "Shaym", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 3, Name: "Damodar", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}}}

	distinctTeachersPtr := SetTeacherPtr(teachersPtr)
	if len(distinctTeachersPtr) != 3 {
		t.Errorf("Error in Set for Employee. actualLength=%d. expectedLength=3", len(distinctTeachersPtr))
	}
}

func TestUnion(t *testing.T) {
	now := time.Now()
	employees1 := []Employee{{Id: 1, Name: "Ram", Salary: 700, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 3, Name: "Radha", Salary: 900, CreationDate: now}}

	employees2 := []Employee{{Id: 1, Name: "Ram", Salary: 700, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 4, Name: "Radha", Salary: 900, CreationDate: now}}

	unionList := Union(employees1, employees2)
	fmt.Println(unionList)
	if len(unionList) != 4 {
		t.Errorf("Error in UnionList")
	}

	address := "Usa"
	teachers1 := []Teacher{{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 2, Name: "Shaym", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 3, Name: "Damodar", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}}}

	teachers2 := []Teacher{{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 2, Name: "Shaym", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 4, Name: "Damodar", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}}}

	teacherUnion := UnionTeacher(teachers1, teachers2)
	if len(teacherUnion) != 4 {
		t.Errorf("Error in UnionTeacher")
	}
	// ptr
	employees1Ptr := []*Employee{{Id: 1, Name: "Ram", Salary: 700, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 3, Name: "Radha", Salary: 900, CreationDate: now}}

	employees2Ptr := []*Employee{{Id: 1, Name: "Ram", Salary: 700, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 4, Name: "Radha", Salary: 900, CreationDate: now}}

	unionListPtr := UnionPtr(employees1Ptr, employees2Ptr)
	fmt.Println(unionListPtr)
	if len(unionListPtr) != 4 {
		t.Errorf("Error in UnionListPtr")
	}

	teachersPtr1 := []*Teacher{{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 2, Name: "Shaym", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 3, Name: "Damodar", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}}}

	teachersPtr2 := []*Teacher{{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 2, Name: "Shaym", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 4, Name: "Damodar", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}}}

	teacherUnionPtr := UnionTeacherPtr(teachersPtr1, teachersPtr2)
	if len(teacherUnionPtr) != 4 {
		t.Errorf("Error in UnionTeacherPtr")
	}

}

func TestIntersection(t *testing.T) {
	now := time.Now()
	employees1 := []Employee{{Id: 1, Name: "Ram", Salary: 700, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 3, Name: "Radha", Salary: 900, CreationDate: now}}

	employees2 := []Employee{{Id: 1, Name: "Ram", Salary: 700, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 4, Name: "Radha", Salary: 900, CreationDate: now}}

	intersectionList := Intersection(employees1, employees2)
	fmt.Println(intersectionList)
	if len(intersectionList) != 2 {
		t.Errorf("Error in IntersectionList")
	}

	address := "Usa"
	teachers1 := []Teacher{{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 2, Name: "Shaym", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 3, Name: "Damodar", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}}}

	teachers2 := []Teacher{{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 2, Name: "Shaym", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 4, Name: "Damodar", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}}}

	teacherIntersection := IntersectionTeacher(teachers1, teachers2)
	if len(teacherIntersection) != 2 {
		t.Errorf("Error in teacherIntersection")
	}
	// ptr
	employees1Ptr := []*Employee{{Id: 1, Name: "Ram", Salary: 700, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 3, Name: "Radha", Salary: 900, CreationDate: now}}

	employees2Ptr := []*Employee{{Id: 1, Name: "Ram", Salary: 700, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 4, Name: "Radha", Salary: 900, CreationDate: now}}

	intersectionListPtr := IntersectionPtr(employees1Ptr, employees2Ptr)
	fmt.Println(intersectionListPtr)
	if len(intersectionListPtr) != 2 {
		t.Errorf("Error in intersectionListPtr")
	}

	teachersPtr1 := []*Teacher{{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 2, Name: "Shaym", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 3, Name: "Damodar", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}}}

	teachersPtr2 := []*Teacher{{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 2, Name: "Shaym", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 4, Name: "Damodar", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}}}

	teacherIntersectionPtr := IntersectionTeacherPtr(teachersPtr1, teachersPtr2)
	if len(teacherIntersectionPtr) != 2 {
		t.Errorf("Error in IntersectionTeacherPtr")
	}
}

func TestSuperset(t *testing.T) {
	now := time.Now()
	employees1 := []Employee{{Id: 1, Name: "Ram", Salary: 700, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 3, Name: "Radha", Salary: 900, CreationDate: now}}

	employees2 := []Employee{{Id: 1, Name: "Ram", Salary: 700, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now}}

	ok := Superset(employees1, employees2)
	if !ok {
		t.Errorf("Error in SuperSEt")
	}

	address := "Usa"
	teachers1 := []Teacher{{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 2, Name: "Shaym", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 3, Name: "Damodar", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}}}

	teachers2 := []Teacher{{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 3, Name: "Damodar", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}}}

	ok = SupersetTeacher(teachers1, teachers2)
	if !ok {
		t.Errorf("Error in SupersetTeacher")
	}
	// ptr
	employees1Ptr := []*Employee{{Id: 1, Name: "Ram", Salary: 700, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 3, Name: "Radha", Salary: 900, CreationDate: now}}

	employees2Ptr := []*Employee{{Id: 1, Name: "Ram", Salary: 700, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 3, Name: "Radha", Salary: 900, CreationDate: now}}

	okPtr := SupersetPtr(employees1Ptr, employees2Ptr)
	if okPtr {
		t.Errorf("Error in SupersetPtr")
	}

	teachersPtr1 := []*Teacher{{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 2, Name: "Shaym", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 3, Name: "Damodar", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}}}

	teachersPtr2 := []*Teacher{{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 3, Name: "Damodar", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}}}

	okPtr = SupersetTeacherPtr(teachersPtr1, teachersPtr2)
	if !okPtr {
		t.Errorf("Error in SupersetTeacherPtr")
	}
}

func TestSubset(t *testing.T) {
	now := time.Now()
	employees1 := []Employee{{Id: 1, Name: "Ram", Salary: 700, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 3, Name: "Radha", Salary: 900, CreationDate: now}}

	employees2 := []Employee{{Id: 1, Name: "Ram", Salary: 700, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now}}

	ok := Subset(employees2, employees1)
	if !ok {
		t.Errorf("Error in SubsetList")
	}

	address := "Usa"
	teachers1 := []Teacher{{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 2, Name: "Shaym", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 3, Name: "Damodar", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}}}

	teachers2 := []Teacher{{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 3, Name: "Damodar", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}}}

	ok = SubsetTeacher(teachers2, teachers1)
	if !ok {
		t.Errorf("Error in SubsetTeacher")
	}
	// ptr
	employees1Ptr := []*Employee{{Id: 1, Name: "Ram", Salary: 700, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 3, Name: "Radha", Salary: 900, CreationDate: now}}

	employees2Ptr := []*Employee{{Id: 1, Name: "Ram", Salary: 700, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 3, Name: "Radha", Salary: 900, CreationDate: now}}

	okPtr := SubsetPtr(employees2Ptr, employees1Ptr)
	if okPtr {
		t.Errorf("Error in SubsetPtr")
	}

	teachersPtr1 := []*Teacher{{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 2, Name: "Shaym", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 3, Name: "Damodar", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}}}

	teachersPtr2 := []*Teacher{{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 3, Name: "Damodar", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}}}

	okPtr = SubsetTeacherPtr(teachersPtr2, teachersPtr1)
	if !okPtr {
		t.Errorf("Error in SubsetTeacherPtr")
	}
}

func TestDifference(t *testing.T) {
	now := time.Now()
	employees1 := []Employee{{Id: 1, Name: "Ram", Salary: 700, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 3, Name: "Radha", Salary: 900, CreationDate: now}}

	employees2 := []Employee{{Id: 1, Name: "Ram", Salary: 700, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 4, Name: "Radha", Salary: 900, CreationDate: now}}

	differenceList := Difference(employees1, employees2)
	if len(differenceList) != 1 {
		t.Errorf("Error in differenceList")
	}

	address := "Usa"
	teachers1 := []Teacher{{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 2, Name: "Shaym", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 3, Name: "Damodar", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}}}

	teachers2 := []Teacher{{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 2, Name: "Shaym", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 4, Name: "Damodar", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}}}

	teacherDifferenceList := DifferenceTeacher(teachers1, teachers2)
	if len(teacherDifferenceList) != 1 {
		t.Errorf("Error in DifferenceTeacher")
	}
	// ptr
	employees1Ptr := []*Employee{{Id: 1, Name: "Ram", Salary: 700, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 3, Name: "Radha", Salary: 900, CreationDate: now}}

	employees2Ptr := []*Employee{{Id: 1, Name: "Ram", Salary: 700, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 4, Name: "Radha", Salary: 900, CreationDate: now}}

	differenceListPtr := DifferencePtr(employees1Ptr, employees2Ptr)
	if len(differenceListPtr) != 1 {
		t.Errorf("Error in differenceListPtr")
	}

	teachersPtr1 := []*Teacher{{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 2, Name: "Shaym", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 3, Name: "Damodar", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}}}

	teachersPtr2 := []*Teacher{{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 1, Name: "Ram", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 2, Name: "Shaym", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}},
		{Id: 4, Name: "Damodar", Salary: 700, CreationDate: &now, Address: &address, Students: []string{"Radha", "Krishna"}}}

	teacherDifferencePtr := DifferenceTeacherPtr(teachersPtr1, teachersPtr2)
	if len(teacherDifferencePtr) != 1 {
		t.Errorf("Error in teacherDifferencePtr")
	}
}

func TestMethodChainWithMap(t *testing.T) {
	now := time.Now()
	employees := []Employee{
		{Id: 1, Name: "Ram", Salary: 700, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 3, Name: "Radha", Salary: 900, CreationDate: now}}

	salaryIncrement := func(emp Employee) Employee {
		emp.Salary = emp.Salary + 1000
		return emp
	}

	salaryGreaterThan800 := func(emp Employee) bool {
		return emp.Salary > 800
	}
	employeesWithIncrementedSalary := MakeEmployeeSlice(employees...).
		Filter(salaryGreaterThan800).
		Map(salaryIncrement)
	if employees[3].Salary+1000 != employeesWithIncrementedSalary[0].Salary {
		t.Error("error in method chain with Map", employees)
	}
}

func TestMethodChainWithMapPtr(t *testing.T) {
	now := time.Now()
	employees := []*Employee{
		{Id: 1, Name: "Ram", Salary: 700, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 3, Name: "Radha", Salary: 900, CreationDate: now}}

	salaryIncrement := func(emp *Employee) *Employee {
		emp.Salary = emp.Salary + 1000
		return emp
	}

	salaryGreaterThan800 := func(emp *Employee) bool {
		return emp.Salary > 800
	}
	employeesWithIncrementedSalary := MakeEmployeeSlicePtr(employees...).
		FilterPtr(salaryGreaterThan800).
		MapPtr(salaryIncrement)
	if 1900 != employeesWithIncrementedSalary[0].Salary {
		t.Errorf("error in method chain with Map %v %v", employees[0].Salary, employeesWithIncrementedSalary[0].Salary)
	}
}

func TestMethodChainWithRemove(t *testing.T) {
	now := time.Now()
	employees := []Employee{
		{Id: 1, Name: "Ram", Salary: 700, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 3, Name: "Radha", Salary: 900, CreationDate: now}}

	salaryIncrement := func(emp Employee) Employee {
		emp.Salary = emp.Salary + 1000
		return emp
	}

	salaryGreaterThan700 := func(emp Employee) bool {
		return emp.Salary > 700
	}

	salaryEqualTo800 := func(emp Employee) bool {
		return emp.Salary == 800
	}
	employeesWithIncrementedSalary := MakeEmployeeSlice(employees...).
		Filter(salaryGreaterThan700).
		Remove(salaryEqualTo800).
		Map(salaryIncrement)
	if employees[3].Salary+1000 != employeesWithIncrementedSalary[0].Salary {
		t.Error("error in method chain with Map", employees)
	}
}

func TestMethodChainWithRemovePtr(t *testing.T) {
	now := time.Now()
	employees := []*Employee{
		{Id: 1, Name: "Ram", Salary: 700, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 3, Name: "Radha", Salary: 900, CreationDate: now}}

	salaryIncrement := func(emp *Employee) *Employee {
		emp.Salary = emp.Salary + 1000
		return emp
	}

	salaryGreaterThan700 := func(emp *Employee) bool {
		return emp.Salary > 700
	}

	salaryEqualTo800 := func(emp *Employee) bool {
		return emp.Salary == 800
	}
	employeesWithIncrementedSalary := MakeEmployeeSlicePtr(employees...).
		FilterPtr(salaryGreaterThan700).
		RemovePtr(salaryEqualTo800).
		MapPtr(salaryIncrement)
	if 1900 != employeesWithIncrementedSalary[0].Salary {
		t.Error("error in method chain with Map", employees)
	}
}

func TestMethodChainWithDropWhilePtr(t *testing.T) {
	now := time.Now()
	employees := []*Employee{
		{Id: 1, Name: "Ram", Salary: 700, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 2, Name: "Shyam", Salary: 800, CreationDate: now},
		{Id: 3, Name: "Radha", Salary: 900, CreationDate: now}}

	salaryIncrement := func(emp *Employee) *Employee {
		emp.Salary = emp.Salary + 1000
		return emp
	}

	salary700 := func(emp *Employee) bool {
		return emp.Salary == 700
	}

	salaryEqualTo800 := func(emp *Employee) bool {
		return emp.Salary == 800
	}
	employeesWithIncrementedSalary := MakeEmployeeSlicePtr(employees...).
		DropWhilePtr(salary700).
		RemovePtr(salaryEqualTo800).
		MapPtr(salaryIncrement)
	if 1900 != employeesWithIncrementedSalary[0].Salary {
		t.Error("error in method chain with Map", employees)
	}
}
