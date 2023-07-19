package main

import (
	"context"
	"log"
	"net"
  pb "grpc/proto"
	"google.golang.org/grpc"
)

const (
  port = ":50051"
)
type server struct {
  pb.UnimplementedAdderServiceServer
}

func (s *server) Add(ctx context.Context , in *pb.AddRequest) (*pb.AddResponse , error){
  x := in.GetX()
  y := in.GetY()
  log.Printf("Resieved :%d  and %d " , x , y)
  return &pb.AddResponse{Result: x + y} , nil
}

func main(){
  lis , err := net.Listen("tcp" , port)
  if err != nil {
    log.Fatal(err)
  }
  s := grpc.NewServer()
  pb.RegisterAdderServiceServer(s , &server{})
  log.Printf("Listening on port :%v", port)
  if err := s.Serve(lis);err != nil {
    log.Fatal(err)
  }

}
