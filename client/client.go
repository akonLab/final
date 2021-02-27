package clientpackage

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	//"com.grpc.tleu/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)

	doPrime(c)
	doAvg(c)
}
func doPrime() {

}
func doAvg() {

}
func doGreetingEveryone(c greetpb.GreetServiceClient) {

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("error while open stream: %v", err)
	}

	requests := []*greetpb.GreetEveryoneRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Tleu",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Bob",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Alice",
			},
		},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range requests {
			log.Printf("Sending message: %v", req)
			err := stream.Send(req)
			if err != nil {
				log.Fatalf("error while sending req to server: %v", err.Error())
			}
			time.Sleep(time.Second)
		}
		err := stream.CloseSend()
		if err != nil {
			log.Fatalf("errow while closing client's stream")
		}
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while getting message from stream: %v", err.Error())
			}
			log.Printf("Received: %v", res.GetResult())
		}
		close(waitc)
	}()

	<-waitc
}
