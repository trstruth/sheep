package main

import (
	"context"
	"testing"

	pb "github.com/trstruth/sheep/pkg/proto"
)

var ctx context.Context = context.Background()

func TestCreateFlag(t *testing.T) {
	kv := NewMemoryKV()
	s := newSheepServer(kv)

	testFlagName := "test-flag"

	resp, err := s.CreateFlag(ctx, &pb.CreateFlagRequest{
		Key: testFlagName,
	})
	if err != nil {
		t.Fatalf("Failed to create flag: %s", err)
	}

	if resp.Key != "test-flag" {
		t.Fatalf("Expected key to be %s, got %s", testFlagName, resp.Key)
	}

	val, exists := kv.store[testFlagName]
	if !exists {
		t.Fatalf("Expected flag to exist in store")
	}

	if val != "{\"key\":\"test-flag\",\"probability\":0}" {
		t.Fatalf("Expected flag value to be %s, got %s", "{\"key\":\"test-flag\",\"probability\":0}", val)
	}
}
