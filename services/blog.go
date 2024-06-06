package services

import (
	"database/sql"
	"fmt"

	"github.com/vikram761/backend/models"
)

type blogService struct {
	Db *sql.DB
}

type BlogService interface {
	Save(models.Blog) error
	Delete(string) error
	GetAll() ([]models.Blog, error)
    GetOne(string) (models.Blog, error)
}

func NewBlogService(db *sql.DB) BlogService {
	return &blogService{Db: db}
}

func (b *blogService) Save(blog models.Blog) error {
	var query string = "INSERT INTO BLOG( IMGURL, HEADING, TAG, DESCRIPTION, CONTENT, AUTHOR) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := b.Db.Exec(query, blog.ImgURL, blog.Heading, blog.Tag, blog.Description, blog.Content, blog.Author)
	if err != nil {
		return err
	}
	return nil
}

func (b *blogService) Delete(id string) error {
	var query string = "DELETE FROM BLOG WHERE ID = $1"
	_, err := b.Db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (b *blogService) GetAll() ([]models.Blog, error) {
	query, err := b.Db.Query("SELECT ID, IMGURL, HEADING, TAG, DESCRIPTION, AUTHOR, CREATED_AT, UPDATED_AT FROM BLOG")
	if err != nil {
		return nil, err
	}

	var result []models.Blog
	defer query.Close()
	for query.Next() {
		var blog models.Blog
		err := query.Scan(&blog.ID, &blog.ImgURL, &blog.Heading, &blog.Tag, &blog.Description, &blog.Author, &blog.CreatedAt, &blog.UpdatedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, blog)
	}
	if len(result) == 0 {
		return []models.Blog{}, nil
	}
	return result, nil
}

func (b *blogService) GetOne(id string) (models.Blog, error) {
	query := b.Db.QueryRow("SELECT * FROM BLOG WHERE ID = $1", id)

	var blog models.Blog
	err := query.Scan(&blog.ID, &blog.ImgURL, &blog.Heading, &blog.Tag, &blog.Description, &blog.Content, &blog.Author, &blog.CreatedAt, &blog.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Blog{}, fmt.Errorf("user with ID %s not found", id)
		}
		return models.Blog{}, err
	}
	return blog, nil
}
