package main

import (
	"fmt"
	"strconv"
)

func main() {
	strconv.Itoa(123)
	// "123"
	i, err := strconv.Atoi("abc")
	if err != nil {
		fmt.Println(err)
		//strconv.Atoi: parsing "abc": invalid syntax
	}
	fmt.Println(i)
	//0

	strconv.FormatInt(123, 2)
	// "1111011"
	i64, err := strconv.ParseInt("abc", 16, 64)
	if err != nil {
		fmt.Println(err)
		//не будет ошибки
	}
	fmt.Println(i64)
	//2748

	strconv.FormatFloat(12.34, 'f', -1, 32)
	// "12.34"
	strconv.FormatFloat(12.34, 'e', 10, 32)
	// "1.2340000153e+01"
	strconv.ParseFloat("12.34", 64)
	//12.34

	strconv.FormatBool(true)
	// "true"
	strconv.ParseBool("true")
	// true
}
