package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/samandar2605/medium_post_service/storage/postgres"
	"github.com/samandar2605/medium_post_service/storage/repo"
)

type StorageI interface {
	Category() repo.CategoryStorageI
	Comment() repo.CommentStorageI
	Post() repo.PostStorageI
	Like() repo.LikeStorageI
}

type storagePg struct {
	categoryRepo repo.CategoryStorageI
	commentRepo  repo.CommentStorageI
	postRepo     repo.PostStorageI
	likeRepo     repo.LikeStorageI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		categoryRepo: postgres.NewCategory(db),
		commentRepo:  postgres.NewComment(db),
		postRepo:     postgres.NewPost(db),
		likeRepo:     postgres.NewLike(db),
	}
}

func (s *storagePg) Category() repo.CategoryStorageI {
	return s.categoryRepo
}

func (s *storagePg) Comment() repo.CommentStorageI {
	return s.commentRepo
}

func (s *storagePg) Post() repo.PostStorageI {
	return s.postRepo
}

func (s *storagePg) Like() repo.LikeStorageI {
	return s.likeRepo
}
