package main

import "fmt"

func average(xs[]float64)float64 {
	total:=0.0
	for _,v:=range xs {
		total +=v
	}
	return total/float64(len(xs))
}

func main() {
	xs:=[]float64{2,4,3,1}
	fmt.Println(average(xs))
}