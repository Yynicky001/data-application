package cmd

import (
	"data_fetch/api/github_api_strategy"
	"data_fetch/config"
	pb "data_fetch/service/pb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedDataFetchServiceServer
}

func main() {
	lis, err := net.Listen("tcp", config.Conf.Server.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDataFetchServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) FetchDataStream(stream pb.DataFetchService_FetchDataStreamServer) error {
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			// 客户端完成发送
			return nil
		}
		if err != nil {
			log.Fatalf("Failed to receive a request: %v", err)
			return err
		}
	}
}

func fetchData() {
	c := &github_api_strategy.GitHubAPIContext{}
	switch config.Conf.GitHub.Strategy {
	case "v4":
		c.SetGitHubAPIContext(&github_api_strategy.GitHubAPIV4Strategy{})
	default:
		c.SetGitHubAPIContext(&github_api_strategy.GitHubAPIDefaultStrategy{})
	}
	c.Fetch()
}
