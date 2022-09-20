package main

import "fmt"

type Employee struct{
    EmployeeId int
    EmployeeName string
}

func main() {
    e1:=Employee{10,"sravan reddy"}
    ptrTo1:=&e1

    fmt.Println(" : ", (*ptrTo1).EmployeeId )
    fmt.Println(" : ", (*ptrTo1).EmployeeName )

    (*ptrTo1).EmployeeId=102
    fmt.Println("2nd :", (*ptrTo1).EmployeeId)

    e2:=e1
    e2.EmployeeId = 103
    fmt.Println("3rd : ", (*ptrTo1).EmployeeId)


}
