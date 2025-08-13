package main
import "fmt"
func main() {
	// Output a simple message
	fmt.Println("Hello, World!")

	// Input from the user
	// Loading input and saving it to a variable (Pointer)
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln( &name )
	fmt.Println("Hello,", name)

	// Basic math
	var a, b int
	fmt.Print("Enter two numbers (a b): ")
	fmt.Scanf( "%d %d", &a, &b )

	var sum, diff, prod int
	sum = a + b
	fmt.Println("Sum:", a+b)
	diff = a - b
	fmt.Println("Difference:", a-b)
	prod = a * b
	fmt.Println("Product:", a*b)

	fmt.Println("Quotient:", float64(a)/float64(b))

	// Another way to take input
	var c ,d int
	fmt.Print("Enter the first number: ")
	fmt.Scan( &c )
	fmt.Print("Enter the second number: ")
	fmt.Scan( &d )
	fmt.Println("You entered:", c, "and", d)
}