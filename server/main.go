
package main

import (
	pb "github.com/jangozw/grpc_service/source"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"//注意跟client 调用的一致
)

type Server struct{}

func (s *Server) Invoke(ctx context.Context, req *pb.Request) (res *pb.Response, err error) {
	log.Println("请求ctx", ctx)
	log.Print(req.ServiceName, req.MethodName, req.Params)
	res = &pb.Response{Code:0, Msg:"ok", Success:true, Data:""}
	err = nil
	return res, err
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Server start")

	s := grpc.NewServer()
	pb.RegisterCallServiceMethodServer(s, &Server{})
	e := s.Serve(lis)
	if e != nil {
		log.Fatal(e)
	}
}
