
package main

import (
	
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	protoapi "time_service/proto" 
)


func AskingTime(ctx context.Context, client protoapi.TimeServiceClient) (*protoapi.TimeResponse, error) { 
	request := &protoapi.TimeRequest{Value: "Please send me the date and time"} 
	response, err := client.GetCurrentTime(ctx, request) 
	if err != nil { 
		log.Fatalf("could not invoke GetCurrentTime RPC %v\n", err) 
	}
	return response , nil 
}


func AskingServerUptime(ctx context.Context, client protoapi.TimeServiceClient) (*protoapi.ServerUptimeResponse, error) { 
	request := &protoapi.ServerUptimeRequest{ClientId: "client-1"} 
	response , err :=  client.GetServerUptime(ctx, request) 
	if err != nil { 
		log.Fatalf("could not invoke GetServerUptime RPC %v\n", err) 
	}
	return response, nil 
}




func main() {
	var port = ":8082" 
	if len (os.Args) == 1 { 
		fmt.Println("Using Default port", port)
	} else { 
		port = os.Args[1] 
	}
	opts := []grpc.DialOption{ 
		grpc.WithTransportCredentials(insecure.NewCredentials()), 
	}

	conn , err := grpc.Dial(port, opts...)
	if err != nil { 
		log.Fatalf("Did not connect: %v\n", err )
	}

	defer func(con *grpc.ClientConn) { 
		if err := conn.Close(); err != nil { 
			log.Fatalf("Unexpected error: %v\n", err) 
		}
	}(conn) 

	client := protoapi.NewTimeServiceClient(conn) 

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5) 
	defer cancel() 

	timeResponse, err := AskingTime(ctx, client) 
	if err != nil { 
		log.Fatalf("Could not invoke RPC: %v\n", err) 
	}

	fmt.Println("The server current time is", timeResponse.Value) 



	serverUptimeResponse, err := AskingServerUptime(ctx, client)
	if err != nil { 
		log.Fatalf("Could not invoke RPC: %v\n", err)
	}


	fmt.Println("The server uptime is " , serverUptimeResponse.GetServerUptime().AsDuration())
	
	}



