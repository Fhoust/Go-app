package routes

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/Fhoust/Go-app/goApp"
	"github.com/Fhoust/Go-app/server/common"
	"github.com/Fhoust/Go-app/server/controllers"
)

type server struct {
	pb.UnimplementedAddUserServer
	pb.UnimplementedGetUserServer
	pb.UnimplementedDeleteUserServer
	pb.UnimplementedUpdateUserServer
}

var (
	user *controllers.User
)

func (s *server) AddNewUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	log.Printf("Received: %v", in.GetName())
	id, err := controllers.Insert(in.GetName())

	if err != nil {
		return nil, err
	}

	return &pb.User{Name: in.GetName(), Id: int64(id)}, nil
}

func (s *server) GetUserInfo(ctx context.Context, in *pb.User) (*pb.User, error) {
	log.Printf("Received: %v", in.GetId())
	name := controllers.UserPerID(int(in.GetId()))
	return &pb.User{Name: name, Id: in.GetId()}, nil
}

func (s *server) UpdateOneUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	log.Printf("Received: %v", in.GetId())
	controllers.UpdateUser(int(in.GetId()), in.GetName())
	return &pb.User{Name: in.GetName(), Id: in.GetId()}, nil
}

func (s *server) DeleteOldUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	log.Printf("Received: %v", in.GetId())
	name := controllers.DeletePerId(int(in.GetId()))
	return &pb.User{Name: name, Id: in.GetId()}, nil
}

// Routes -> define endpoints
func Routes() {

	lis, err := net.Listen("tcp", common.GetPort())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	//findXByY
	pb.RegisterAddUserServer(s, &server{})
	pb.RegisterGetUserServer(s, &server{})
	pb.RegisterDeleteUserServer(s, &server{})
	pb.RegisterUpdateUserServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}