package grpc

import (
	"blockchain_services/config"
	"blockchain_services/grpc/impl"
	"blockchain_services/grpc/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func StartGrpcServer() error {
	lis, err := net.Listen("tcp", ":"+config.GrpcPort)
	if err != nil {
		return err
	}

	log.Println("start grpc server at :" + config.GrpcPort)

	grpcServer := grpc.NewServer()
	pb.RegisterBlockChainTicketServerServer(grpcServer, impl.NewTicketsServer())

	err = grpcServer.Serve(lis)
	if err != nil {
		return err
	}

	return nil
}
