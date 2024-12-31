package service

import (
	"A1_BMS_PROJECT/model"
	"A1_BMS_PROJECT/repository"
)

type BlogService struct {
	Repo *repository.BlogRepository
}

func NewBlogService(repo *repository.BlogRepository) *BlogService {
	return &BlogService{Repo: repo}
}

func (s *BlogService) CreateBlog(blog *model.Blog) (*model.Blog, error) {
	return s.Repo.CreateBlog(blog)
}

func (s *BlogService) GetBlogByID(id int) (*model.Blog, error) {
	return s.Repo.GetBlogByID(id)
}

func (s *BlogService) GetAllBlogs() ([]model.Blog, error) {
	return s.Repo.GetAllBlogs()
}

func (s *BlogService) UpdateBlog(blog *model.Blog) (*model.Blog, error) {
	return s.Repo.UpdateBlog(blog)
}

func (s *BlogService) DeleteBlog(id int) error {
	return s.Repo.DeleteBlog(id)
}
