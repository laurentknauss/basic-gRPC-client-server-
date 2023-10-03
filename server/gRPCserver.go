


package main

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "time_service/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/durationpb"
	
)

type TimeServer struct { 
	pb.UnimplementedTimeServiceServer
	startTime time.Time
}


func (s *TimeServer) GetCurrentTime (ctx context.Context, req *pb.TimeRequest) (*pb.TimeResponse, error) {
	// Set a timeout for the following operations 
	ctx, cancel := context.WithTimeout(ctx, time.Second*5) 
	defer cancel() 
  
	currentTime := time.Now() 
	formattedTime := currentTime.Format("15:04:05") 
	return  &pb.TimeResponse{
		Value: formattedTime, 
	}, nil

} 
	

func (s *TimeServer) GetServerUptime (ctx context.Context, req *pb.ServerUptimeRequest) (*pb.ServerUptimeResponse, error) {
	// Set a timeout for the following operations 
	ctx, cancel := context.WithTimeout(ctx, time.Second*5) 
	defer cancel() 


	serverUptime := time.Since(s.startTime) 
	return &pb.ServerUptimeResponse{
		ServerUptime: durationpb.New(serverUptime), 


	}, nil 
} 


func main() {
	var port = ":8082" 

server  := grpc.NewServer() 
timeServer := TimeServer{
	startTime: time.Now() , 
}

pb.RegisterTimeServiceServer(server, &timeServer) 

reflection.Register(server) 

listener, err := net.Listen("tcp", port) 
if err != nil { 
	log.Fatalf("Failed to listen %v\n", err) 
}
fmt.Println("The gRPC server listens and serve on port ", port)
if err := server.Serve(listener) ;
err != nil {  log.Fatalf("Failed to serve: %v\n", err)  
}




}
