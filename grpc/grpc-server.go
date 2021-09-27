package main

import (
	"apiproto"
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type DateServer struct {
	apiproto.UnimplementedDateServer
}

func main() {
	l, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatal("Unable to listen on port 9999", err)
	}
	grpcServer := grpc.NewServer()
	var dateServer DateServer
	apiproto.RegisterDateServer(grpcServer, dateServer)
	fmt.Println("serving grpc requests...")
	if err := grpcServer.Serve(l); err != nil {
		log.Fatal("Failed to serve grpc", err)
	}

}

func (DateServer) GetDateTime(ctx context.Context, r *apiproto.RequestDateTime) (*apiproto.DateTime, error) {
	currentTime := time.Now()
	response := &apiproto.DateTime{
		Value: currentTime.String(),
	}
	return response, nil
}
