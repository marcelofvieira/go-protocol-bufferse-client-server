package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/marcelofvieira/protocol-buffer3/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {

	log.Println("LongGreet was invoked!")

	res := ""

	i := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{Result: res})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		res += fmt.Sprintf("Request %d: %s\n", i, req.FirstName)
	}

}
