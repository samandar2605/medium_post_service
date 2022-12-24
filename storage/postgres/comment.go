package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/samandar2605/medium_post_service/storage/repo"
)

type commentRepo struct {
	db *sqlx.DB
}

func NewComment(db *sqlx.DB) repo.CommentStorageI {
	return &commentRepo{db: db}
}

func (cr *commentRepo) Create(comment *repo.Comment) (*repo.Comment, error) {
	query := `
		INSERT INTO comments(
			post_id,
			user_id,
			description
		) values ($1,$2,$3)
		RETURNING
			id,
			created_at
	`
	result := cr.db.QueryRow(
		query,
		comment.PostId,
		comment.UserId,
		comment.Description,
	)
	if err := result.Scan(
		&comment.Id,
		&comment.CreatedAt,
	); err != nil {
		return nil, err
	}
	return comment, nil
}

func (cr *commentRepo) Get(id int) (*repo.Comment, error) {
	var Comment repo.Comment
	query := `
		SELECT
			id,
			user_id,
			post_id,
			description,
			created_at
		FROM comments
		where id=$1`

	result := cr.db.QueryRow(
		query,
		id,
	)
	if err := result.Scan(
		&Comment.Id,
		&Comment.UserId,
		&Comment.PostId,
		&Comment.Description,
		&Comment.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &Comment, nil
}

func (cr *commentRepo) GetAll(param repo.GetCommentQuery) (*repo.GetAllCommentsResult, error) {
	result := repo.GetAllCommentsResult{
		Comments: make([]*repo.Comment, 0),
	}

	offset := (param.Page - 1) * param.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", param.Limit, offset)
	filter := ""
	if param.PostId > 0 {
		filter += fmt.Sprintf(" where post_id=%d and id!=0 ", param.PostId)
	}

	if param.UserId > 0 {
		if filter == "" {
			filter += fmt.Sprintf(" where user_id=%d", param.UserId)
		} else {
			filter += fmt.Sprintf(" and user_id=%d ", param.UserId)
		}
	}

	query := `
		SELECT
			id,
			user_id,
			post_id,
			description,
			created_at
		FROM comments 
		` + filter + `
		ORDER BY created_at ` + param.SortByDate + ` ` + limit

	rows, err := cr.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var Comment repo.Comment
		if err := rows.Scan(
			&Comment.Id,
			&Comment.UserId,
			&Comment.PostId,
			&Comment.Description,
			&Comment.CreatedAt,
		); err != nil {
			return nil, err
		}
		result.Comments = append(result.Comments, &Comment)
	}
	queryCount := `
		SELECT count(1) FROM comments ` + filter
	err = cr.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return nil, err
	}
	fmt.Println(result)
	return &result, nil
}

func (cr *commentRepo) Update(comme *repo.Comment) (*repo.Comment, error) {
	result := cr.db.QueryRow(`
		update comments set 
			description=$1,
			updated_at=$2
		where id=$3 and user_id=$4
		RETURNING post_id,created_at
	`,
		comme.Description,
		time.Now(),
		comme.Id,
		comme.UserId,
	)

	comme.UpdatedAt = time.Now().Format(time.RFC3339)
	if err := result.Scan(
		&comme.PostId,
		&comme.CreatedAt,
	); err != nil {
		return nil, err
	}

	return comme, nil
}

func (cr *commentRepo) Delete(id int) error {
	res, err := cr.db.Exec("delete from comments where id=$1", id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (cr *commentRepo) GetUserInfo(id int) int {
	var userId int

	query := `
		SELECT 
			user_id
		from comments
		where id=$1
	`
	row := cr.db.QueryRow(query, id)
	if err := row.Scan(
		&userId,
	); err != nil {
		return -1
	}
	return userId
}
