package main

import (
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/woodwo/unknown/grpc/proto"
	mock_proto "github.com/woodwo/unknown/grpc/proto/mock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Test_nextValueFromCalculus(t *testing.T) {
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("No calculus")
	}

	client := proto.NewComplicatedClient(conn)

	t.Run("Happy Path - at least one act", func(t *testing.T) {
		n1, _ := catsFromRandom(client)
		if !strings.Contains(n1, "🐈") {
			t.Error("Happy path Random should be positive")
		}
	})
}

func Test_nextValueMocked(t *testing.T) {
	// mocks
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	client := mock_proto.NewMockComplicatedClient(mockCtl)

	gomock.InOrder(
		client.EXPECT().Random(gomock.Any(), gomock.Any()).Return(&proto.Value{Value: 1}, nil).Times(1),
	)

	t.Run("Happy Path - exactly one cat", func(t *testing.T) {
		n1, _ := catsFromRandom(client)
		if n1 != "🐈\n" {
			t.Error("Expect just one cat")
		}
	})
}
