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

package main

import (
	"context"
	"log"
	"os"
	"time"

	ptypes "github.com/golang/protobuf/ptypes"
	any "github.com/golang/protobuf/ptypes/any"
	pb "github.com/rajeshpachar/grpc-go-sample/greeter"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:5555"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	person1 := &pb.Person{FirstName: "rajesh", LastName: "pachar"}
	any1, err := ptypes.MarshalAny(person1)
	person2 := &pb.Person{FirstName: "raj", LastName: "kumar"}
	any2, err := ptypes.MarshalAny(person2)

	persons := []*any.Any{any1, any2}

	// any, err := ptypes.MarshalAny(persons)

	log.Println("here is person data", persons)

	r, err := c.SayHello(ctx, &pb.HelloRequest{
		Name:    name,
		Details: persons,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Response)
}
