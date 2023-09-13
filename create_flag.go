package main

import (
	"context"
	"encoding/json"
	"fmt"

	pb "github.com/trstruth/sheep/pkg/proto"
)

func (s *Server) CreateFlag(ctx context.Context, in *pb.CreateFlagRequest) (*pb.CreateFlagResponse, error) {
	flagExists, err := s.store.Exists(ctx, in.Key)
	if err != nil {
		return nil, fmt.Errorf("Failed to check if flag exists: %s", err)
	}

	if flagExists {
		return nil, fmt.Errorf("Flag with key %s already exists", in.Key)
	}

	newFlag := NewFlag(in.Key)

	flagBytes, err := json.Marshal(newFlag)
	if err != nil {
		return nil, fmt.Errorf("Failed to Marshal flag contents: %s", err)
	}
	flagStr := string(flagBytes)

	if err := s.store.Set(ctx, in.Key, flagStr); err != nil {
		return nil, fmt.Errorf("Failed to Set flag contents: %s", err)
	}

	return &pb.CreateFlagResponse{
		Key: in.Key,
	}, nil
}
