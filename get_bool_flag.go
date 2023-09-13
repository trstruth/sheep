package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	pb "github.com/trstruth/sheep/pkg/proto"
)

func (s *Server) GetBoolFlag(ctx context.Context, in *pb.GetBoolFlagRequest) (*pb.GetBoolFlagResponse, error) {
	flagExists, err := s.store.Exists(ctx, in.Key)
	if err != nil {
		return nil, fmt.Errorf("Failed to check if flag exists: %s", err)
	}

	if !flagExists {
		return nil, fmt.Errorf("No such flag exists: %s", in.Key)
	}

	flagContents, err := s.store.Get(ctx, in.Key)
	if err != nil {
		return nil, fmt.Errorf("Failed to Get flag contents: %s", err)
	}

	var flag Flag
	if err := json.Unmarshal([]byte(flagContents), &flag); err != nil {
		return nil, fmt.Errorf("Failed to Unmarshal flag contents: %s", err)
	}

	consistencyKey := in.GetConsistencyKey()
	if consistencyKey == "" {
		consistencyKey = fmt.Sprint(time.Now().UnixNano())
		fmt.Println(consistencyKey)
	}

	return &pb.GetBoolFlagResponse{
		Key:   in.Key,
		Value: flag.ToBool(consistencyKey),
	}, nil
}
