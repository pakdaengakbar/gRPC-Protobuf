package main

import (
	"context"
	"log"
	"time"

	pbranch "grpcserver/proto"

	"google.golang.org/grpc"
)

func main() {
	// Connect ke gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pbranch.NewBranchServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// --- CreateBranch
	createRes, err := client.CreateBranch(ctx, &pbranch.CreateBranchRequest{
		Name:    "Jakarta HQ",
		Address: "Jl. Sudirman No. 1",
	})
	if err != nil {
		log.Fatalf("CreateBranch error: %v", err)
	}
	log.Printf("Created Branch: ID=%s, Message=%s\n", createRes.Id, createRes.Message)

	// --- GetBranch
	getRes, err := client.GetBranch(ctx, &pbranch.GetBranchRequest{
		Id: createRes.Id,
	})
	if err != nil {
		log.Fatalf("GetBranch error: %v", err)
	}
	log.Printf("Fetched Branch: ID=%s, Name=%s, Address=%s\n", getRes.Id, getRes.Name, getRes.Address)
}
