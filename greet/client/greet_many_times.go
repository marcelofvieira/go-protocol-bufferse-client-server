package main

import (
	"context"
	"io"
	"log"

	pb "github.com/marcelofvieira/protocol-buffer3/greet/proto"
)

func goGreetManyTimes(c pb.GreetServiceClient) {

	log.Printf("doGreetManyTimes was invoked")

	req := &pb.GreetRequest{
		FirstName: "Marcelo",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error reading streaam: %v\n", err)
		}

		log.Printf("Greet Many times: %s\n", msg.Result)

	}

}
