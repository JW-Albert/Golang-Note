package main
import "fmt"
func main() {
	// Arithmetic Operations + - * / %
	fmt.Println("======Arithmetic Operations======")
	a := 10
	b := 5
	sum := a + b
	diff := a - b
	prod := a * b
	quot := a / b
	mod := a % b
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", diff)
	fmt.Println("Product:", prod)
	fmt.Println("Quotient:", quot)
	fmt.Println("Modulus:", mod)
	fmt.Println("\n")

	// Comparison Operations == != > < >= <=
	fmt.Println("======Comparison Operations======")
	isEqual := a == b
	isNotEqual := a != b
	isGreater := a > b
	isLess := a < b
	isGreaterOrEqual := a >= b
	isLessOrEqual := a <= b
	fmt.Println("Is Equal:", isEqual)
	fmt.Println("Is Not Equal:", isNotEqual)
	fmt.Println("Is Greater:", isGreater)
	fmt.Println("Is Less:", isLess)
	fmt.Println("Is Greater or Equal:", isGreaterOrEqual)
	fmt.Println("Is Less or Equal:", isLessOrEqual)

	// Logical Operations && || !
	fmt.Println("\n======Logical Operations======")
	isTrue := true
	isFalse := false
	andOp := isTrue && isFalse
	orOp := isTrue || isFalse
	notOp := !isTrue
	fmt.Println("AND Operation:", andOp)
	fmt.Println("OR Operation:", orOp)
	fmt.Println("NOT Operation:", notOp)

	// Bitwise Operations & | ^ << >>
	fmt.Println("\n======Bitwise Operations======")
	x := 5  // 0101 in binary
	y := 3  // 0011 in binary
	andBitwise := x & y   // 0001 in binary
	orBitwise := x | y    // 0111 in binary
	xorBitwise := x ^ y   // 0110 in binary
	// Left shift and right shift
	// Left shift moves bits to the left, right shift moves bits to the right
	// Each shift operation effectively multiplies or divides the number by 2
	leftShift := x << 1   // 1010 in binary
	rightShift := x >> 1  // 0010 in binary
	fmt.Println("AND Bitwise:", andBitwise)
	fmt.Println("OR Bitwise:", orBitwise)
	fmt.Println("XOR Bitwise:", xorBitwise)
	fmt.Println("Left Shift:", leftShift)
	fmt.Println("Right Shift:", rightShift)

	// Assignment Operations = += -= *= /= %=
	fmt.Println("\n======Assignment Operations======")
	c := 10
	c += 5  // c = c + 5
	d := 20
	d -= 5  // d = d - 5
	e := 30
	e *= 2  // e = e * 2
	f := 40
	f /= 2  // f = f / 2
	g := 50
	g %= 3  // g = g % 3
	fmt.Println("c after += 5:", c)
	fmt.Println("d after -= 5:", d)
	fmt.Println("e after *= 2:", e)
	fmt.Println("f after /= 2:", f)
	fmt.Println("g after %= 3:", g)

	// Unary Operations + - ! ++ --
	fmt.Println("\n======Unary Operations======")
	h := 10
	i := 20
	h++  // Increment h by 1
	i--  // Decrement i by 1
	fmt.Println("h after ++:", h)
	fmt.Println("i after --:", i)
}