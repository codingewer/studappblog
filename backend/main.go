package main

import (
	"log"
	"net/http"
	"studapp-blog/api/controllers"
	"studapp-blog/middlewares"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//User routers
	router.HandleFunc("/api/users/new", middlewares.SetMiddlewareJSON(controllers.CreateUser)).Methods("POST")
	router.HandleFunc("/api/users/confirm/{token}", middlewares.SetMiddlewareJSON(controllers.ConfirmAcoount)).Methods("POST")
	router.HandleFunc("/api/users/getAll", middlewares.SetMiddlewareJSON(controllers.GetUsers)).Methods("GET")
	router.HandleFunc("/api/users/getById/{id}", middlewares.SetMiddlewareJSON(controllers.GetUser)).Methods("GET")
	router.HandleFunc("/api/users/delete/{id}", middlewares.SetMiddlewareAuthentication(controllers.DeleteUser)).Methods("DELETE")
	router.HandleFunc("/api/users/update/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(controllers.UpdateUser))).Methods("POST")
	router.HandleFunc("/login", middlewares.SetMiddlewareJSON(controllers.Login)).Methods("POST", http.MethodPost)

	router.HandleFunc("/api/posts/new", middlewares.SetMiddlewareJSON(controllers.CreatePost)).Methods("POST")
	router.HandleFunc("/api/posts/getAll", middlewares.SetMiddlewareJSON(controllers.GetPosts)).Methods("GET")
	router.HandleFunc("/api/posts/getById/{id}", middlewares.SetMiddlewareJSON(controllers.GetPost)).Methods("GET")
	router.HandleFunc("/api/posts/getByUserId/{userId}", middlewares.SetMiddlewareJSON(controllers.GetPostsByUserID)).Methods("GET")
	router.HandleFunc("/api/posts/delete/{id}", middlewares.SetMiddlewareAuthentication(controllers.DeletePost)).Methods("DELETE")
	router.HandleFunc("/api/posts/update/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(controllers.UpdatePost))).Methods("POST")
	router.HandleFunc("/api/images/upload", middlewares.SetMiddlewareJSON(controllers.ImgUpload)).Methods("POST")
	router.HandleFunc("/api/images/update/{imageId}", middlewares.SetMiddlewareJSON(controllers.UpdateImage)).Methods("POST")
	router.HandleFunc("/api/images/getAll", middlewares.SetMiddlewareJSON(controllers.GetAllImages)).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))

}
