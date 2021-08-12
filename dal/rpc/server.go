package rpc

import (
	user_center "Blog/proto/user-center"
	"log"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", "3333")
	if err != nil {
		logrus.Fatalf("failed to listen : %v", err)
	}
	ser := grpc.NewServer()
	user_center.RegisterUserCenterServer(ser, &UserCenterServer{})
	reflection.Register(ser)
	if err := ser.Serve(lis); err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
