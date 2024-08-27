package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	pb "github.com/shanth1/golang-microservices/grpc/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	port = flag.String("port", ":50051", "the address to connect to")
)

func main() {
	flag.Parse()

	if flag.NArg() < 2 {
		log.Fatal("not enough args")
	}

	x, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	y, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("test", *port)
	conn, err := grpc.NewClient(*port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAdderClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Add(ctx, &pb.AddRequest{X: int32(x), Y: int32(y)})
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	log.Printf("Result: %d", r.GetResult())
}
