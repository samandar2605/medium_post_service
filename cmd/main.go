package main

import (
	"fmt"
	"log"
	"net"

	_"github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	pb "github.com/samandar2605/medium_post_service/genproto/post_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/samandar2605/medium_post_service/config"
	"github.com/samandar2605/medium_post_service/service"
	"github.com/samandar2605/medium_post_service/storage"
)

func main() {
	cfg := config.Load(".")

	psqlUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)

	psqlConn, err := sqlx.Connect("postgres", psqlUrl)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	strg := storage.NewStoragePg(psqlConn)

	postService := service.NewPostService(strg)
	categoryService := service.NewCategoryService(strg)
	commentService := service.NewCommentService(strg)

	lis, err := net.Listen("tcp", cfg.GrpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterPostServiceServer(s, postService)
	pb.RegisterCategoryServiceServer(s, categoryService)
	pb.RegisterCommentServiceServer(s,commentService)

	log.Println("Grpc server started in port ", cfg.GrpcPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error while listening: %v", err)
	}

}
