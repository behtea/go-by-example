package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroiptr(iptr *int) {
	fmt.Println("callptr: ", *iptr)
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial: ", i)

	zeroval(i)
	fmt.Println("zeroval: ", i)

	fmt.Println("pointer:", &i)

	zeroiptr(&i)
	fmt.Println("zeroiptr: ", i)

	fmt.Println("pointer:", &i)
}
