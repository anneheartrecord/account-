package main

import (
	"account/biz"
	"account/internal"
	pb "account/proto"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func init() {
	internal.InitDB()
}

func main() {
	ip := flag.String("ip", "127.0.0.1", "请输入ip")
	port := flag.Int64("port", 9999, "请输入端port")
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", *ip, *port)
	server := grpc.NewServer()
	pb.RegisterAccountServiceServer(server, &biz.AccountServer{})
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	err = server.Serve(listen)

}
