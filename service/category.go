package service

import (
	"context"
	"time"

	pb "github.com/samandar2605/medium_post_service/genproto/post_service"
	"github.com/samandar2605/medium_post_service/storage"
	"github.com/samandar2605/medium_post_service/storage/repo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func parseCategoryModel(c *repo.Category) *pb.Category {
	return &pb.Category{
		Id:        int64(c.Id),
		Title:     c.Title,
		CreatedAt: c.CreatedAt.Format(time.RFC3339),
	}
}

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	storage storage.StorageI
}

func NewCategoryService(strg storage.StorageI) *CategoryService {
	return &CategoryService{
		UnimplementedCategoryServiceServer: pb.UnimplementedCategoryServiceServer{},
		storage:                            strg,
	}
}

func (s *CategoryService) Create(ctx context.Context, req *pb.Category) (*pb.Category, error) {
	category, err := s.storage.Category().Create(&repo.Category{
		Title: req.Title,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}
	return parseCategoryModel(category), nil
}
func (s *CategoryService) Get(ctx context.Context, req *pb.IdRequest) (*pb.Category, error) {
	resp, err := s.storage.Category().Get(int(req.Id))
	if err != nil {
		return nil, err
	}
	return parseCategoryModel(resp), nil
}

func (s *CategoryService) GetAll(ctx context.Context, req *pb.GetCategoryRequest) (*pb.GetCategoryResponse, error) {
	resp, err := s.storage.Category().GetAll(repo.GetCategoryQuery{
		Page:   int(req.Page),
		Limit:  int(req.Limit),
		Search: req.Search,
	})
	if err != nil {
		return nil, err
	}

	result := pb.GetCategoryResponse{
		Count:      int32(resp.Count),
		Categories: make([]*pb.Category, 0),
	}
	for _, category := range resp.Categories {
		result.Categories = append(result.Categories, parseCategoryModel(category))
	}
	return &result, nil
}

func (s *CategoryService) Update(ctx context.Context, req *pb.Category) (*pb.Category, error) {
	category, err := s.storage.Category().Update(&repo.Category{
		Id:    int(req.Id),
		Title: req.Title,
	})
	if err != nil {
		return nil, err
	}
	return parseCategoryModel(category), nil
}

func (s *CategoryService) Delete(ctx context.Context, req *pb.IdRequest) (*pb.Empty, error) {
	err := s.storage.Category().Delete(int(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
