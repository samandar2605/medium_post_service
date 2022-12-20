package service

import (
	"context"

	pb "github.com/samandar2605/medium_post_service/genproto/post_service"
	"github.com/samandar2605/medium_post_service/storage/repo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/samandar2605/medium_post_service/storage"
)

func parseComment(comment *repo.Comment) *pb.Comment {
	return &pb.Comment{
		Id:          int64(comment.Id),
		PostId:      int64(comment.PostId),
		UserId:      int64(comment.UserId),
		Description: comment.Description,
		CreatedAt:   comment.CreatedAt,
	}
}

type CommentService struct {
	pb.UnimplementedPostServiceServer
	storage storage.StorageI
}

func NewCommentService(strg storage.StorageI) *CommentService {
	return &CommentService{
		UnimplementedPostServiceServer: pb.UnimplementedPostServiceServer{},
		storage:                        strg,
	}
}

func (s *CommentService) CreateComment(ctx context.Context, req *pb.CreateComment) (*pb.Comment, error) {
	comment, err := s.storage.Comment().Create(&repo.Comment{
		PostId:      int(req.PostId),
		UserId:      int(req.UserId),
		Description: req.Description,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}

	return parseComment(comment), nil
}

func (s *CommentService) Get(ctx context.Context, req *pb.IdWithRequest) (*pb.Comment, error) {
	Comment, err := s.storage.Comment().Get(int(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}

	return parseComment(Comment), nil
}

func (s *CommentService) GetAll(ctx context.Context, req *pb.GetCommentQuery) (*pb.GetAllCommentsResult, error) {
	result, err := s.storage.Comment().GetAll(repo.GetCommentQuery{
		Limit:      int(req.Limit),
		Page:       int(req.Page),
		UserId:     int(req.UserId),
		PostId:     int(req.PostId),
		SortByDate: req.SortByDate,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}

	response := pb.GetAllCommentsResult{
		Count:    int64(result.Count),
		Comments: make([]*pb.Comment, 10),
	}

	for _, Comment := range result.Comments {
		response.Comments = append(response.Comments, parseComment(Comment))
	}

	return &response, nil
}

func (s *CommentService) Update(ctx context.Context, req *pb.Comment) (*pb.Comment, error) {
	Comment, err := s.storage.Comment().Update(&repo.Comment{
		Id:          int(req.Id),
		PostId:      int(req.PostId),
		UserId:      int(req.UserId),
		Description: req.Description,
		CreatedAt:   req.CreatedAt,
		UpdatedAt:   req.UpdatedAt,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}

	return parseComment(Comment), nil
}

func (s *CommentService) Delete(ctx context.Context, req *pb.IdWithRequest) (*pb.Boosh, error) {
	err := s.storage.Comment().Delete(int(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}
	return &pb.Boosh{}, nil
}
