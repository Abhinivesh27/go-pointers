package main // Package Name

import "fmt" /* importing fmt package
Package fmt implements formatted Input and Output */

func main() {
	var count = int(27) // declearing int value

	point := &count /* "&" operator used to retrieve
	the address in memory of some variable */

	fmt.Println(point) // printing memory address of count variable

	fmt.Println(*point) // "*" operator to dereference the address

	*point = 100 /* assigning a new value
	to the memory location pointed to by point */

	fmt.Println(count)
}
