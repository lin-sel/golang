package main

import (
	"fmt"
)

func main() {

	/*

		All Datatype Present in Golang.
		(1) Integer
			Represent Numeric Value.

			Signed Integer
				int8	8 bit	-128 to 127
				int16	16 bit	-2^16 to (2^15)-1
				int32	32 bit	-2^32 to (2^31)-1
				int64	64 bit	-2^64 to (2^63)-1
				int 	Plateform dependent 32 bit for 32 bit System and 64 bit for 64 bit System.

			unsigned Integer
				uint8	8 bit	0 to 127
				uint16	16 bit	0 to (2^16)-1
				uint32	32 bit	0 to (2^32)-1
				uint64	64 bit	0 to (2^64)-1
				uint 	Plateform dependent 32 bit for 32 bit System and 64 bit for 64 bit System.
	*/

	// Signed Integer
	var Age int8 = 40
	fmt.Printf("%v, %T\n", Age, Age)

	// Unsigned Integer
	var Ages uint8 = 40
	fmt.Printf("%v, %T\n", Ages, Ages)

	/*

		(2) Boolean
			Containing Only two Value true and false.
	*/

	// Float
	var booles bool = true
	fmt.Printf("%v, %T\n", booles, booles)

	/*

		(3) Float
			Represent decimal value.

			  type	 Size
			float32  32 bit
			float64  64 bit
	*/

	// Float
	var pointer float64 = 40
	fmt.Printf("%v, %T\n", pointer, pointer)

	/*
		(4) String
			Sequence of bytes.

	*/

	var name string = "Wordpress"
	fmt.Println(name)

	/*

		(5) Complex
			Both Imaginary and Real Number will be same datatype.

			complex64	64 bit (32 bit real and imaginary)
			complex128	128 bit (64 bit real and imaginary)
	*/

	var complexnum complex64 = 23 + 98i
	fmt.Println(complexnum)

	// Build-in function used for declare complex number declaration.
	fmt.Println(complex(23, 87))

	// Three way of Variable Declaration

	// (1) Complete Declaration ( as we saw above)
	var max int = 98
	fmt.Println(max)

	// (2) Just Declaration (But Not Initialize)
	var status int8
	fmt.Println(status)

	// (3) Shorthand Declaration
	profession := "Developer"
	fmt.Println(profession)

	// var block
	// this is shorthand declaration of varible in bulk

	var (
		data    int64 = 98
		status2 bool  = false
	)

	fmt.Println(data, status2)

	// Another Style to declare multiple variable

	var a, b, c, d int16 = 12, 13, 14, 15
	fmt.Println(a, b, c, d)

}
