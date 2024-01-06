package main

import (
	"log"
	"net"

	"github.com/prabhjout/Train-Ticket-App/api"
	"github.com/prabhjout/Train-Ticket-App/api/pb"
	"google.golang.org/grpc"
)

func main() {
	port := ":50051"

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	tt := api.NewTrainTicketAppServer()
	pb.RegisterTrainTicketAppServer(s, tt)

	log.Printf("server listening on port %v", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
