package main 

import "fmt"

func main(){

    fmt.Println("Enter temperature in Fahrenheit")
    var fahrenheit float64
    fmt.Scanf("%f", &fahrenheit)
    
    celsius := (fahrenheit - 32) * 5/9
    fmt.Println(celsius)
    
}