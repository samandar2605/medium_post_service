package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/samandar2605/medium_post_service/storage/repo"
)

type postRepo struct {
	db *sqlx.DB
}

func NewPost(db *sqlx.DB) repo.PostStorageI {
	return &postRepo{db: db}
}

func (pr *postRepo) Create(p *repo.Post) (*repo.Post, error) {
	query := `
		INSERT INTO posts(
			title,
			description,
			image_url,
			user_id,
			category_id
		)values($1,$2,$3,$4,$5)
		RETURNING id,created_at,views_count
	`
	row := pr.db.QueryRow(
		query,
		p.Title,
		p.Description,
		p.ImageUrl,
		p.UserId,
		p.CategoryId,
	)

	if err := row.Scan(
		&p.Id,
		&p.CreatedAt,
		&p.ViewsCount,
	); err != nil {
		return nil, err
	}
	return p, nil
}

func (pr *postRepo) Get(id int) (*repo.Post, error) {
	var Post repo.Post

	query := `
		SELECT 
			id,
			title,
			description,
			image_url,
			user_id,
			category_id,
			views_count,
			created_at
		from posts
		where id=$1
	`
	row := pr.db.QueryRow(query, id)
	if err := row.Scan(
		&Post.Id,
		&Post.Title,
		&Post.Description,
		&Post.ImageUrl,
		&Post.UserId,
		&Post.CategoryId,
		&Post.ViewsCount,
		&Post.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &Post, nil
}

func (pr *postRepo) GetAll(param repo.GetPostQuery) (*repo.GetAllPostResult, error) {
	result := repo.GetAllPostResult{
		Post: make([]*repo.Post, 0),
	}

	offset := (param.Page - 1) * param.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", param.Limit, offset)
	filter := ""

	if param.CategoryID > 0 {
		filter += fmt.Sprintf("where category_id=%d", param.CategoryID)
	}
	if param.UserID > 0 {
		if filter == "" {
			filter += fmt.Sprintf("where user_id=%d", param.UserID)
		} else {
			filter += fmt.Sprintf("and user_id=%d", param.UserID)
		}
	}
	if param.SortByDate == "" {
		param.SortByDate = "desc"
	}
	query := `
		SELECT 
			id,
			title,
			description,
			image_url,
			user_id,
			category_id,
			views_count,
			created_at
		FROM posts
		` + filter + `
		ORDER BY created_at ` + param.SortByDate + ` ` + limit

	rows, err := pr.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var Post repo.Post
		if err := rows.Scan(
			&Post.Id,
			&Post.Title,
			&Post.Description,
			&Post.ImageUrl,
			&Post.UserId,
			&Post.CategoryId,
			&Post.ViewsCount,
			&Post.CreatedAt,
		); err != nil {
			return nil, err
		}
		result.Post = append(result.Post, &Post)
	}
	queryCount := `SELECT count(1) FROM posts ` + filter
	err = pr.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (pr *postRepo) Update(post *repo.ChangePost) (*repo.Post, error) {
	var ResponsePost repo.Post
	fmt.Println("Id: ",post.Id)
	fmt.Println("Title: ",post.Title)
	fmt.Println("UserId: ",post.UserId)
	fmt.Println("Description: ",post.Description)
	fmt.Println("ImageUrl: ",post.ImageUrl)
	query := `
		update posts set 
			title=$1,
			description=$2,
			image_url=$3,
			updated_at=$4
		where id=$5 and user_id=$6
		RETURNING
			category_id,
			views_count,
			created_at
	`
	row := pr.db.QueryRow(
		query,
		post.Title,
		post.Description,
		post.ImageUrl,
		time.Now(),
		post.Id,
		post.UserId,
	)
	ResponsePost.Id = post.Id
	ResponsePost.Title = post.Title
	ResponsePost.Description = post.Description
	ResponsePost.ImageUrl = post.ImageUrl
	ResponsePost.UserId = post.UserId
	ResponsePost.UpdatedAt = time.Now()

	if err := row.Scan(
		&ResponsePost.CategoryId,
		&ResponsePost.ViewsCount,
		&ResponsePost.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &ResponsePost, nil
}

func (ur *postRepo) Delete(id int) error {
	res, err := ur.db.Exec("delete from posts where id=$1", id)
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

func (ur *postRepo) ViewsInc(id int) error {
	_, err := ur.db.Exec(`UPDATE posts SET views_count=views_count+1 WHERE id=$1`, id)
	if err != nil {
		return err
	}
	return nil
}
