/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

package main

import (
	"context"
	"log"
	"net"

	pb "github.com/rajeshpachar/grpc-go-sample/greeter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":5555"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Println("receiver reque from client " + in.Name)
	return &pb.HelloReply{Response: "Hello my name is: " + in.Name}, nil
}

func (s *server) Check(ctx context.Context, in *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	log.Println("health check received for  " + in.Service)
	return nil, &pb.HealthCheckResponse{Status : pb.HealthCheckResponse_SERVING}
}

// SayHello implements helloworld.GreeterServer
// func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
// 	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
// }

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
