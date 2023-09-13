package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/redis/go-redis/v9"
	pb "github.com/trstruth/sheep/pkg/proto"
)

func main() {
	cfg, err := NewConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening on %s:%d\n", cfg.Server.Host, cfg.Server.Port)

	var kv KeyValuer
	switch cfg.Storage.StorageType {
	case "memory":
		log.Println("starting with memory backend")
		kv = NewMemoryKV()
		break
	case "redis":
		log.Println("starting with redis backend")
		opts, err := redis.ParseURL(cfg.Storage.ConnectionString)
		if err != nil {
			log.Fatalf("failed to parse redis connection string: %v", err)
		}

		kv = NewRedisKV(opts)

	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	reflection.Register(grpcServer)
	pb.RegisterSheepServer(grpcServer, newSheepServer(kv))

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
