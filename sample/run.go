package main

import "fmt"

import (
	js "sample/jsontostruct"
	es "sample/typeandfunc"
	e "sample/typeandinterface"
)

func main() {

	fmt.Println(`Employee Interface implementation`)
	var e1 e.Employee
	e1 = e.Emp(1)
	e1.PrintName("John Doe")
	fmt.Println("Employee Salary:", e1.PrintSalary(25000, 5))

	// json str to struct and back
	fmt.Println(`Json str to struct and back.`)
	js.Run()

	// add function to type - user defined type, struct etc...
	fmt.Println(`Add function to type.`)
	structToFunc := es.EmpSample(999)
	structToFunc.PrintName("John Doe")
	fmt.Println("Employee Salary:", structToFunc.PrintSalary(250000, 5))
}
