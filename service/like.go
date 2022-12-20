package service

import (
	"context"

	pb "github.com/samandar2605/medium_post_service/genproto/post_service"
	"github.com/samandar2605/medium_post_service/storage"
	"github.com/samandar2605/medium_post_service/storage/repo"
)

type LikeService struct {
	pb.UnimplementedLikeServiceServer
	storage storage.StorageI
}

func NewLikeService(strg storage.StorageI) *LikeService {
	return &LikeService{
		UnimplementedLikeServiceServer: pb.UnimplementedLikeServiceServer{},
		storage:                        strg,
	}
}

func (l *LikeService) CreateOrUpdateLike(ctx context.Context, req *pb.CreateOrUpdateLikeRequest) error {
	err := l.storage.Like().CreateOrUpdate(&repo.Like{
		UserID: req.UserId,
		PostID: req.PostId,
		Status: req.Status,
	})
	if err != nil {
		return err
	}
	return nil
}

func (l *LikeService) GetLike(ctx context.Context, req *pb.CreateOrUpdateLikeRequest) (*pb.CreateOrUpdateLikeRequest, error) {
	resp, err := l.storage.Like().Get(req.UserId, req.PostId)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrUpdateLikeRequest{
		Id:     resp.ID,
		PostId: resp.PostID,
		UserId: resp.UserID,
		Status: resp.Status,
	}, nil
}
