package main
 
import (
         "fmt"
          "reflect"
)
 
func main() {
	var employee = make(map[string]int)
	employee["xyz"] = 10
	employee["zyx"] = 20
	fmt.Println(employee)
	employeeList := make(map[string]int)
	employeeList["xyz"] = 10
	employeeList["zyx"] = 20
	fmt.Println(employeeList)
	var intSlice = make([]int, 10)       
	var strSlice = make([]string, 10, 20) 
	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
	fmt.Println(reflect.ValueOf(intSlice).Kind())
	fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(strSlice), cap(strSlice))
	fmt.Println(reflect.ValueOf(strSlice).Kind())
}