package main

import (
	"context"
	"testing"

	pb "github.com/trstruth/sheep/pkg/proto"
)

func TestGetBoolFlag(t *testing.T) {
	testFlagName := "test-flag"
	contents := "{\"key\":\"test-flag\",\"probability\":0}"

	kv := NewMemoryKV()
	kv.store[testFlagName] = contents

	s := newSheepServer(kv)

	resp, err := s.GetBoolFlag(context.Background(), &pb.GetBoolFlagRequest{
		Key: testFlagName,
	})
	if err != nil {
		t.Fatalf("Failed to get flag: %s", err)
	}

	if resp.Key != testFlagName {
		t.Fatalf("Expected key to be %s, got %s", testFlagName, resp.Key)
	}

	if resp.Value != false {
		t.Fatalf("Expected value to be %t, got %t", false, resp.Value)
	}
}

func TestBoolFlagProbability(t *testing.T) {
	testFlagName := "test-flag"
	contents := "{\"key\":\"test-flag\",\"probability\":25}"

	kv := NewMemoryKV()
	kv.store[testFlagName] = contents

	s := newSheepServer(kv)

	trueCount := 0
	falseCount := 0
	for i := 0; i < 10000; i++ {
		resp, err := s.GetBoolFlag(context.Background(), &pb.GetBoolFlagRequest{
			Key: testFlagName,
		})
		if err != nil {
			t.Fatalf("Failed to get flag: %s", err)
		}

		if resp.Value {
			trueCount++
		} else {
			falseCount++
		}
	}

	t.Logf("true/false ratio: %f", float64(trueCount)/float64(falseCount+trueCount))

}
