package main

import (
	"context"
	"net"
	"log"
	"google.golang.org/grpc"
	pb "project/rpc"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

type server struct {
	pb.UnimplementedRegisterServer
}

func (s *server) Arrival(ctx context.Context, in *pb.Ship) (*pb.Reply, error) {
	log.Println(in.GetImo(), in.GetName())
	return &pb.Reply{Message: "Dati ricevuti."}, nil
}

func main() {
	log.Println("Port on")
	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	s := grpc.NewServer()
	pb.RegisterRegisterServer(s, &server{})
	lis.Addr()
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err.Error())
	}
}