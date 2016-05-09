package main

import "fmt"

func main() {
	i := 1
	fmt.Println("Initial value is:", i)

	//calling zeroval will not change the value of i
	zeroval(i)
	fmt.Println("After calling zero val:", i)

	//Now we will set the value of i to zero using its pointer
	zeroptr(&i)
	fmt.Println("After calling zero pointer", i)
}

func zeroval(val int) {
	val = 0
}

func zeroptr(val_pointer *int) {
	*val_pointer = 0
}
