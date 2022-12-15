package service

import (
	"context"
	"time"

	pb "github.com/samandar2605/medium_post_service/genproto/post_service"
	"github.com/samandar2605/medium_post_service/storage/repo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/samandar2605/medium_post_service/storage"
)

func parsePostModel(Post *repo.Post) *pb.Post {
	return &pb.Post{
		Id:         Post.Id,
		Title:       Post.Title,
		Description: Post.Description,
		ImageUrl:    Post.ImageUrl,
		UserId:      Post.UserId,
		CategoryId:  Post.CategoryId,
		UpdatedAt:   Post.UpdatedAt.Format(time.RFC3339),
		ViewsCount:  int32(Post.ViewsCount),
		CreatedAt: Post.CreatedAt.Format(time.RFC3339),
	}
}

type PostService struct {
	pb.UnimplementedPostServiceServer
	storage storage.StorageI
}

func NewPostService(strg storage.StorageI) *PostService {
	return &PostService{
		UnimplementedPostServiceServer: pb.UnimplementedPostServiceServer{},
		storage: strg,
	}
}

func (s *PostService) CreatePost(ctx context.Context, req *pb.Post) (*pb.Post, error) {
	post, err := s.storage.Post().Create(&repo.Post{
		Title:       req.Title,
		Description: req.Description,
		ImageUrl:    req.ImageUrl,
		UserId:      req.UserId,
		CategoryId:  req.CategoryId,
		ViewsCount:  int64(req.ViewsCount),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}

	return parsePostModel(post), nil
}

func (s *PostService) Get(ctx context.Context, req *pb.GetPostRequest) (*pb.Post, error) {
	Post, err := s.storage.Post().Get(int(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}

	return parsePostModel(Post), nil
}

func (s *PostService) GetAll(ctx context.Context, req *pb.GetAllPostsRequest) (*pb.GetAllPostsResponse, error) {
	result, err := s.storage.Post().GetAll(repo.GetPostQuery{
			Limit:  req.Limit,
			Page:   req.Page,
			UserID: req.UserId,
			CategoryID: int64(req.CategoryId),
			SortByDate: req.SortByDate,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}

	response := pb.GetAllPostsResponse{
		Count: int64(result.Count),
		Posts: make([]*pb.Post,10),
	}

	for _, Post := range result.Post {
		response.Posts = append(response.Posts, parsePostModel(Post))
	}

	return &response, nil
}

func (s *PostService) Update(ctx context.Context, req *pb.Post) (*pb.Post, error) {
	post, err := s.storage.Post().Update(&repo.Post{
		Title:       req.Title,
		Description: req.Description,
		ImageUrl:    req.ImageUrl,
		UserId:      req.UserId,
		CategoryId:  req.CategoryId,
		ViewsCount:  int64(req.ViewsCount),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}

	return parsePostModel(post), nil
}

func (s *PostService) Delete(ctx context.Context, req *pb.GetPostRequest) (*pb.Blank, error) {
	err := s.storage.Post().Delete(int(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}
	return &pb.Blank{}, nil
}
