package main

import (
	"fmt"
	"hash/fnv"
	"math/rand"
)

type Flag struct {
	Key         string `json:"key"`
	Probability int    `json:"probability"`
}

func NewFlag(key string) *Flag {
	return &Flag{
		Key:         key,
		Probability: 0,
	}
}

func (f *Flag) ToBool(consistencyKey string) bool {
	seed := consistencyKeyToSeed(consistencyKey)
	r := rand.New(rand.NewSource(seed))

	return (r.Intn(100) + 1) <= f.Probability
}

func (f *Flag) SetProbability(probability int) error {
	if probability < 0 || probability > 100 {
		return fmt.Errorf("Probability must be between 0 and 100, got %d", probability)
	}
	f.Probability = probability

	return nil
}

func consistencyKeyToSeed(consistencyKey string) int64 {
	h := fnv.New32a()
	h.Write([]byte(consistencyKey))
	return int64(h.Sum32())
}
