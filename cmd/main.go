package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/woodwo/unknown/grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func catsFromRandom(c proto.ComplicatedClient) (string, error) {
	ctx := context.TODO()
	num, err := c.Random(ctx, &proto.Empty{})
	if err != nil {
		return "", err
	}
	return strings.Repeat("🐈", int(num.Value)) + "\n", nil
}

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("complicated server is not available")
	}
	defer conn.Close()

	client := proto.NewComplicatedClient(conn)
	cats, err := catsFromRandom(client)
	if err != nil {
		panic("complicated server returned an error")
	}

	fmt.Print(cats)
}
