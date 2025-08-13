// This program demonstrates the use of different variable types in Go.

// Defining and printing different types of variables
// Format:
// 1. String
// 2. Integer
// 3. Float
// 4. Boolean
// 5. Rune (for Unicode characters)

// Defining format: var variableName variableType = value
// Example: var age int = 30

package main
import "fmt"
func main() {
	// String variable
	fmt.Println("======String======")
	fmt.Println("Hello, World!")

	// "\n" is used to print a new line
	fmt.Println("\n")

	// Integer variable
	fmt.Println("======Integer======")
	var age int = 30
	fmt.Println("Age:", age)

	fmt.Println("\n")

	// Float variable
	fmt.Println("======Float======")
	var height float64 = 5.9
	fmt.Println("Height:", height)

	fmt.Println("\n")

	// Boolean variable
	fmt.Println("======Boolean======")
	var isStudent bool = true
	fmt.Println("Is Student:", isStudent)

	fmt.Println("\n")

	// Rune variable (for Unicode characters)
	fmt.Println("======Rune======")
	var letter rune = 'A'
	fmt.Println("Letter:", string(letter))
	fmt.Println("Unicode of Letter:", letter)
}