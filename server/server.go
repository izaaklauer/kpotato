package server

import (
	"context"
	"log"

        "github.com/izaaklauer/kpotato/config"
	kpotatov1 "github.com/izaaklauer/kpotato/gen/proto/go/kpotato/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type KpotatoServer struct {
	kpotatov1.UnimplementedKpotatoServiceServer

	config config.Kpotato
}

// NewKpotatoServer initializes a new server from config
func NewKpotatoServer(config config.Kpotato) (*KpotatoServer, error) {
	// Server-specific initialization, like DB clients, goes here.

	server := KpotatoServer{
		config: config,
	}

	return &server, nil
}

func (s * KpotatoServer) HelloWorld(
	ctx context.Context,
	req *kpotatov1.HelloWorldRequest,
) (*kpotatov1.HelloWorldResponse, error) {
	log.Printf("HelloWorld request with message %q", req.Message)

	resp := &kpotatov1.HelloWorldResponse{
		RequestMessage: req.Message,
		ConfigMessage:  s.config.HelloWorldMessage,
		Now:            timestamppb.Now(),
	}

	return resp, nil
}
