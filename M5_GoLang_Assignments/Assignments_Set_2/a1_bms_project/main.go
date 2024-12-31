package main

import (
	db "A1_BMS_PROJECT/config"
	"A1_BMS_PROJECT/controller"
	"A1_BMS_PROJECT/middlewares"
	"A1_BMS_PROJECT/repository"
	"A1_BMS_PROJECT/service"
	"fmt"
	"net/http"
)

func main() {
	db := db.InitializeDatabase()
	defer db.Close()

	authRepo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo)
	authController := controller.NewAuthController(authService)

	blogRepo := repository.NewBlogRepository(db)
	blogService := service.NewBlogService(blogRepo)
	blogController := controller.NewBlogController(blogService)

	http.HandleFunc("/signup", authController.Signup)

	protectedMux := http.NewServeMux()
	protectedMux.HandleFunc("/blogs", blogController.GetAllBlogs)
	protectedMux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			blogController.CreateBlog(w, r)
		case http.MethodGet:
			blogController.GetBlogByID(w, r)
		case http.MethodPut:
			blogController.UpdateBlog(w, r)
		case http.MethodDelete:
			blogController.DeleteBlog(w, r)
		default:
			http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)
		}
	})
	protectedRoutes := middlewares.AuthMiddleware(db, protectedMux)
	http.Handle("/blogs", protectedRoutes)
	http.Handle("/blog", protectedRoutes)
	loggedRoutes := middlewares.LoggingMiddleware(http.DefaultServeMux)

	if err := http.ListenAndServe(":8085", loggedRoutes); err != nil {
		fmt.Println("Error Starting Server:", err)
	}
}
