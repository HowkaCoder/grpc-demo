package main

import (
	"context"
	pb "grpc/proto"
	"log"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
)


const (
  address = "localhost:50051"
  defaultX = 2
  defaultY = 4
)

func main(){
  conn , err := grpc.Dial(address, grpc.WithInsecure() , grpc.WithBlock())
  if err != nil {
    log.Fatalf("did not connect: %v" , err)
  }
  defer conn.Close()

  c := pb.NewAdderServiceClient(conn)
  x := defaultX
  y := defaultY
  if len(os.Args) > 1 {
    x  , _ = strconv.Atoi(os.Args[1])
    y  , _ = strconv.Atoi(os.Args[2])
  }
  ctx , cancel := context.WithTimeout(context.Background(), time.Second)
  defer cancel()

  r , err := c.Add(ctx, &pb.AddRequest{X : int32(x) , Y : int32(y)})
  if err != nil {
    log.Fatalf("Error with addRequest : %v" , err)
  }
  log.Printf("Greeting : %d" , r.GetResult())
}
