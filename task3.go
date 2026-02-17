
package main

import "fmt"

func main() {
// 1: What does this print?
// i := 10
// if i > 10 {
//     fmt.Println("Big")
// } else {
//     fmt.Println("Small")
// }
// Answer: "Small" (since i = 10, not > 10)
i := 10
if i > 10 {
fmt.Println("Big")
} else {
fmt.Println("Small")
}

// 2: Numbers divisible by 3 between 1 and 100
fmt.Print("Problem 2: ")
for n := 1; n <= 100; n++ {
if n%3 == 0 {
fmt.Print(n, " ")
}
}
fmt.Println()

// 3: FizzBuzz
fmt.Print("Problem 3: ")
for n := 1; n <= 100; n++ {
if n%3 == 0 && n%5 == 0 {
fmt.Print("FizzBuzz ")
} else if n%3 == 0 {
fmt.Print("Fizz ")
} else if n%5 == 0 {
fmt.Print("Buzz ")
} else {
fmt.Print(n, " ")
}
}
fmt.Println()
}
