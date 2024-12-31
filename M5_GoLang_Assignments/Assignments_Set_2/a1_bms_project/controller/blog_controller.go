package controller

import (
	"A1_BMS_PROJECT/model"
	"A1_BMS_PROJECT/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type BlogController struct {
	Service *service.BlogService
}

func NewBlogController(service *service.BlogService) *BlogController {
	return &BlogController{Service: service}
}

func (ctrl *BlogController) CreateBlog(w http.ResponseWriter, r *http.Request) {
	var blog model.Blog
	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	blogr, err := ctrl.Service.CreateBlog(&blog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(blogr)
}

func (ctrl *BlogController) GetBlogByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Employee ID", http.StatusBadRequest)
		return
	}
	blog, err := ctrl.Service.GetBlogByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if blog == nil {
		http.Error(w, "Blog not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(blog)
}

func (ctrl *BlogController) GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	blogs, err := ctrl.Service.GetAllBlogs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(blogs)
}

func (ctrl *BlogController) UpdateBlog(w http.ResponseWriter, r *http.Request) {
	var blog model.Blog
	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedBlog, err := ctrl.Service.UpdateBlog(&blog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedBlog)
}

func (ctrl *BlogController) DeleteBlog(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	if err := ctrl.Service.DeleteBlog(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
