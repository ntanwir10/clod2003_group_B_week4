package main

import "fmt"

func main() {
	fmt.Println("Welcome to Group B's Week 4 Project!")
  
  // Function - 1: Palindrome function by Srinidhi Sivakumar
	fmt.Println("Function 1: Palindrome function submitted by Srinidhi Sivakumar ")
   
   // declare and initialize the string
   str := "madam"
   fmt.Println("Check the given word is palindrome or not,\n Given Word =",str)
 
   // calling the function
   if Palindrome(str) {
      fmt.Printf("'%s' is a palindrome\n", str)
   } else {
      fmt.Printf("'%s' is not a palindrome\n", str)
   }  
  
  // Function - 2: Maximum of 3 Numbers by Anjani
	fmt.Printf("Function 2: Maximum of 3 Numbers by Anjani")	
  	num1 := 56
	num2 := 35
	num3 := 20
	max := findMax(num1, num2, num3)
	fmt.Printf("The maximum of %d, %d, and %d is %d\n", num1, num2, num3, max)  

    //Function 3: Calculating the Area of a Circle
    fmt.Println("Function 3: Calculating the Area of a Circla by Oviya Krishnamoorthy ")

    radius := 7.0
	area := CircleArea(radius)
	fmt.Printf("The area of the circle with radius %.2f is %.2f\n", radius, area)
    
    // Function 4: Bubble Sort by Neelotpal
	  fmt.Println("Function 4: Bubble Sort by Neelotpal")
         ExampleBubblesort()

}
