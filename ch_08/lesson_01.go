package main

import (
	"fmt"
	// "go/constant"
)

// Learning Variables
func firstTry()  {
	var roomNumber, floorNumber int
	fmt.Println(roomNumber, floorNumber)

	var password string
	fmt.Println(password)
}


func secondTry()  {
	var roomNumber, floorNumber int = 154, 3
	fmt.Println(roomNumber,floorNumber)

	var password = "notSecured"
	fmt.Println(password)
}

func thirdTry()  {
	 roomNumber, floorNumber := 120, 55
	 fmt.Println(roomNumber, floorNumber)

	 password := "Secured"
	 fmt.Println(password)

}

// Learning Constant

const version string = "1.3.2"

func fourthTry()  {
	const occupancyLimit = 12

	var occupancyLimit1 uint8
	var occupancyLimit2 int64
	var occupancyLimit3 float32

	occupancyLimit1 = occupancyLimit
	occupancyLimit2 = occupancyLimit
	occupancyLimit3 = occupancyLimit

	fmt.Println(occupancyLimit1, occupancyLimit2, occupancyLimit3)


	// var occupancyLimit4 string

	// occupancyLimit4 = occupancyLimit

	// fmt.Println(occupancyLimit4)
}

func fifthTry()  {
	// defualt type is bool
	const isOpen = true

	// default type is rune (alias for int32)
	const MyRune = 'r'

	// default type is int
	const occupancyLimit = 12

	// default type is float64
	const vatRate = 29.87

	// default type is comples128
	const complexNumber = 1 + 2i

	// default type is string
	const hotelName = "Gopher Hotel"
}

func sixthFourth()  {
	// maximum value of an int is 9223372036854775807
	// 9223372036854775807 (max + 1) overflows int

	const profit = 9223372036854775808
	// the program compiles

	
}


func main()  {
	firstTry()
	secondTry()
	thirdTry()
	fourthTry()
	sixthFourth()
}