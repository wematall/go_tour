package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	fmt.Println(p)  // read and print address in memory

	*p = 21        // set i new value through the pointer
	fmt.Println(i) // see the new value of i
	fmt.Println(p) // see if the memory address the same

	p = &j // point to j
	fmt.Println(p)
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

}
