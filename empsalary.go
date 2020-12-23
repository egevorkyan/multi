package main

import(
	"fmt"
	"github.com/egevorkyan/multi/employee"
	"github.com/egevorkyan/multi/sources"
	"strings"
	"strconv"
)

func main(){
	var path string
	fmt.Print("Enter Employees source file: ")
	fmt.Scanln(&path)
	result := sources.Source{Path: path}
	f, err := result.OpenFile()
	if err != nil{
		panic(err)
	}
	for i := 0; i < len(f); i++{
		sliceStr := strings.Split(f[i], ",")
		basic,_ := strconv.Atoi(sliceStr[1])
		tax,_ := strconv.Atoi(sliceStr[2])
		var emp employee.Employee
		emp = employee.Emp(i)
		emp.PrintName(sliceStr[0])
		fmt.Print("Employee Salary: ", emp.PrintSalary(basic, tax))
	}

}