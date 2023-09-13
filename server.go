package main

import (
	pb "github.com/trstruth/sheep/pkg/proto"
)

type Server struct {
	pb.UnimplementedSheepServer
	store KeyValuer
}

func newSheepServer(store KeyValuer) *Server {
	return &Server{
		store: store,
	}
}
