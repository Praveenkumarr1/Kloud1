package main 
  
import "fmt"
  

type Address struct { 
    Name    string 
    city    string 
    Pincode int
} 
  
func main() { 
  
    
    var a Address  
    fmt.Println(a) 
  
    
    a1 := Address{"akash", "chennai", 600028} 
  
    fmt.Println("Address1: ", a1) 
  
    
    a2 := Address{Name: "peter", city: "tirutani", 
                                 Pincode: 600567} 
  
    fmt.Println("Address2: ", a2) 
  
    
    a3 := Address{Name: "jack"} 
    fmt.Println("Address3: ", a3) 
} 