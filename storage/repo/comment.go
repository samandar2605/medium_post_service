package repo

type GetCommentQuery struct {
	Page       int    `json:"page" db:"page" binding:"required" default:"1"`
	Limit      int    `json:"limit" db:"limit" binding:"required" default:"10"`
	PostId     int    `json:"post_id" db:"post_id"`
	UserId     int    `json:"user_id" db:"user_id"`
	SortByDate string `json:"sort_by_date" enums:"asc,desc" default:"desc"`
}

type GetAllCommentsResult struct {
	Comments []*Comment
	Count    int
}

type Comment struct {
	Id          int    `json:"id" db:"id"`
	PostId      int    `json:"post_id" db:"post_id"`
	UserId      int    `json:"user_id" db:"user_id"`
	Description string `json:"description" db:"description"`
	CreatedAt   string `json:"created_at" db:"created_at"`
	UpdatedAt   string `json:"updated_at" db:"updated_at"`
}

type CommentStorageI interface {
	Create(comment *Comment) (*Comment, error)
	Get(id int) (*Comment, error)
	GetAll(param GetCommentQuery) (*GetAllCommentsResult, error)
	Update(cr *Comment) (*Comment, error)
	GetUserInfo(id int) int
	Delete(id int) error
}
