package main

import (
	"context"
	"fmt"
)

// KeyValuer is an interface for a key-value store.
// It is used to abstract the underlying storage mechanism.
type KeyValuer interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key, val string) error
	Exists(ctx context.Context, key string) (bool, error)
}

type MemoryKV struct {
	store map[string]string
}

func NewMemoryKV() *MemoryKV {
	return &MemoryKV{
		store: make(map[string]string),
	}
}

func NewMemoryKVFromMap(store map[string]string) *MemoryKV {
	return &MemoryKV{
		store: store,
	}
}

func (kv *MemoryKV) Get(ctx context.Context, key string) (string, error) {
	val, ok := kv.store[key]
	if !ok {
		return "", fmt.Errorf("key %s does not exist in the store", key)
	}

	return val, nil
}

func (kv *MemoryKV) Set(ctx context.Context, key string, val string) error {
	kv.store[key] = val
	return nil
}

func (kv *MemoryKV) Exists(ctx context.Context, key string) (bool, error) {
	_, ok := kv.store[key]
	return ok, nil
}

var _ KeyValuer = (*MemoryKV)(nil)
