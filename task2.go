package main

import "fmt"

func main() {
// 1: Integers are stored in binary (base 2) using bits (0s and 1s)

// 2: Largest 8-digit binary number
// Binary: 11111111 = 2^8 - 1 = 255
fmt.Print("Largest 8-digit binary number = ")
fmt.Println(255)

// 3: Calculate 321325 * 424521
fmt.Print("321325 * 424521 = ")
fmt.Println(321325 * 424521)

// 4: String and its length
// A string is text inside double quotes
// Use len() to find length
fmt.Print("Length of 'Hello' = ")
fmt.Println(len("Hello"))

// Problem 5: Boolean expression
fmt.Print("(true && false) || (false && true) || !(false && false) = ")
fmt.Println((true && false) || (false && true) || !(false && false))
}