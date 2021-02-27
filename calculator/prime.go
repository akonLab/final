package main

import (
	"fmt"
)

func main() {
	num := 6

	for i := 2; i < num; i++ {
		//n:=triangular(i)
		for num % i == 0 {
			fmt.Println(i )
			num /= i

		}
	}
	if num>2{
		fmt.Println(num )
	}

}