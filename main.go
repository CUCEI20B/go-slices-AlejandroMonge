package main

import "fmt"

func main() {

	var n int
	var s []int
	var app int
	suma := 0

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&app)
		s = append(s, app)
	}

	for _, v := range s {
		suma += v
	}

	fmt.Println(suma)

}
