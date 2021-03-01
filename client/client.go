package clientpackage

import (
	"com.grpc.tleu/calcpb"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := calcpb.NewCalcServClient(conn)

	doPrime(c)
	doAvg(c)
}

func doPrime(c calcpb.CalcServClient) {
	ctx := context.Background()
	req := &calcpb.CalcPrimeRequest{}
	&prime
	response, err := c.C(ctx, req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC %v", err)
	}
	log.Printf("response from Greet:%v", response.Result)

	num := 120
	arr := []int{}
	for i := 2; i < num; i++ {
		for num%i == 0 {
			arr = append(arr, i)
			num /= i
		}
	}
	if num > 2 {
		arr = append(arr, num)
	}
	fmt.Print(arr)

}
func doAvg(c calcpb.CalcServClient) {
	xs := []float64{2, 4, 3, 1}
	total := 0.0
	for _, v := range xs {
		total += v
	}
	fmt.Print(total / float64(len(xs)))
}
