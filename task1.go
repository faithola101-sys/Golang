
package main

import "fmt"

func main() {
// 1: Two ways to create a new variable
// 1. var x int = 10
// 2. x := 10

// 2: Value of x after x := 5; x += 1
x := 5
x += 1
fmt.Print("Problem 2: x = ")
fmt.Println(x)

// 3: Scope - where a variable can be accessed
// In Go, scope is determined by blocks { }

// 4: var vs const
// var - can change (mutable)
// const - cannot change (immutable)

// 5
f := 100 
c := (f - 32) * 5 / 9
fmt.Print("Problem 5: ", f, "F = ")
fmt.Print(c)
fmt.Println("C")

// 6
feet := 10 
meters := float64(feet) * 0.3048
fmt.Print("Problem 6: ", feet, "ft = ")
fmt.Print(meters)
fmt.Println("m")
}
