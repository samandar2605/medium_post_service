package repo

import "time"

type GetPostQuery struct {
	Page       int32    `json:"page" db:"page" binding:"required" default:"1"`
	Limit      int32    `json:"limit" db:"limit" binding:"required" default:"10"`
	UserID     int64    `json:"user_id"`
	CategoryID int64    `json:"post_id"`
	SortByDate string `json:"sort_by_date" enums:"asc,desc" default:"desc"`
}

const (
	UserTypeSuperadmin = "superadmin"
	UserTypeUser       = "user"
)

type GetAllPostResult struct {
	Post  []*Post
	Count int32
}

type Post struct {
	Id          int64     `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	ImageUrl    string    `json:"image_url" db:"image_url"`
	UserId      int64     `json:"user_id" db:"user_id"`
	CategoryId  int64     `json:"category_id" db:"category_id"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	ViewsCount  int64     `json:"views_count" db:"views_count"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type PostStorageI interface {
	Create(p *Post) (*Post, error)
	Get(id int) (*Post, error)
	GetAll(param GetPostQuery) (*GetAllPostResult, error)
	Update(usr *Post) (*Post, error)
	Delete(id int) error
	GetUserInfo(id int) int
	ViewsInc(id int) error
}
