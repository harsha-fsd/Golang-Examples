package main

import (
	"apiproto"
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":9999", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Dial:", err)
		return
	}
	client := apiproto.NewDateClient(conn)
	for{
	r, err := AskingDateTime(context.Background(), client)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Server Date and Time:", r.Value)
	time.Sleep(1*time.Second)
}
}

func AskingDateTime(ctx context.Context, dateClient apiproto.DateClient) (*apiproto.DateTime, error) {
	request := &apiproto.RequestDateTime{
		Value: "Please send me current server date and time",
	}
	return dateClient.GetDateTime(ctx, request)
}
