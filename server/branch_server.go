package main

import (
	"context"
	"fmt"

	pbranch "grpcserver/proto"
)

type branchServer struct {
	pbranch.UnimplementedBranchServiceServer
	data map[string]*pbranch.GetBranchResponse
}

func newBranchServer() *branchServer {
	return &branchServer{
		data: make(map[string]*pbranch.GetBranchResponse),
	}
}

func (s *branchServer) CreateBranch(ctx context.Context, req *pbranch.CreateBranchRequest) (*pbranch.CreateBranchResponse, error) {
	id := fmt.Sprintf("branch_%d", len(s.data)+1)

	branch := &pbranch.GetBranchResponse{
		Id:      id,
		Name:    req.Name,
		Address: req.Address,
	}

	s.data[id] = branch

	return &pbranch.CreateBranchResponse{
		Id:      id,
		Message: "Branch created successfully",
	}, nil
}

func (s *branchServer) GetBranch(ctx context.Context, req *pbranch.GetBranchRequest) (*pbranch.GetBranchResponse, error) {
	branch, exists := s.data[req.Id]
	if !exists {
		return nil, fmt.Errorf("branch not found")
	}

	return branch, nil
}
