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

type CategoryService struct{
	pb.UnimplementedCategoryServiceServer
	storage storage.StorageI
}

func NewCategoryService(strg storage.StorageI)*CategoryService{
	return &CategoryService{
		storage: strg,
	}
}


func (s *CategoryService) Create(ctx context.Context,req *pb.Category)(*pb.Category,error){
	category,err:=s.storage.Category().Create(&repo.Category{
		Title: req.Title,
	})
	if err!=nil{
		return nil,status.Errorf(codes.Internal,"Internal server error: %v",err)
	}
	return parseCategoryModel(category),nil
}

func parseCategoryModel(c *repo.Category)*pb.Category{
	return &pb.Category{
		Id: int64(c.Id),
		Title: c.Title,
		CreatedAt: c.CreatedAt.Format(time.RFC3339),
	}
}