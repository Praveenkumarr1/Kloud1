package main 
   
import ( 
    "fmt"
    "encoding/json"
) 

type Human struct{ 
       
    // defining struct variables 
    Name string 
    Age int
    Address string 
} 

func main() { 

    human1 := Human{"Rohit", 33, "Tamilnadu"} 
    human_enc, err := json.Marshal(human1) 
       
    if err != nil { 
        fmt.Println(err) 
    } 
        
    fmt.Println(string(human_enc)) 
    human2 := []Human{ 
        {Name: "Rohit", Age: 33, Address: "Tamilnadu"}, 
        {Name: "Bhavi", Age: 20, Address: "kolkata"}, 
        {Name: "Dube", Age: 22, Address: "Manglore"}, 
    } 

    human2_enc, err := json.Marshal(human2) 
        
        if err != nil { 
            fmt.Println(err) 
        } 
           
    fmt.Println() 
        fmt.Println(string(human2_enc)) 
} 