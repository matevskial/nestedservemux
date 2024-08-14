package main

import (
	"log"
	"net/http"
)

type userHandler struct{}

func (u *userHandler) Handler() http.Handler {
	httpServeMux := http.NewServeMux()
	/* when below handler is commented, requests to POST /api/users result with 405 Method Not Allowed, why? */
	/* when below handlefunc is uncommented, requests to POST /api/users is handled by the handler for GET /api/users, why? */
	//httpServeMux.HandleFunc("GET /", func(w http.ResponseWriter, req *http.Request) {
	//	RespondWithText(w, http.StatusOK, "GET /api/users works")
	//})
	httpServeMux.HandleFunc("POST /", func(w http.ResponseWriter, req *http.Request) {
		RespondWithText(w, http.StatusOK, "POST /api/users works")
	})
	return httpServeMux
}

type postHandler struct{}

func (p *postHandler) Handler() http.Handler {
	httpServeMux := http.NewServeMux()
	/* when below handler is commented, requests to POST /api/posts result with 405 Method Not Allowed, why? */
	/* when below handlefunc is uncommented, requests to POST /api/posts is handled by the handler for GET /api/posts, why? */
	//httpServeMux.HandleFunc("GET /", func(w http.ResponseWriter, req *http.Request) {
	//	RespondWithText(w, http.StatusOK, "GET /api/posts works")
	//})
	httpServeMux.HandleFunc("POST /", func(w http.ResponseWriter, req *http.Request) {
		RespondWithText(w, http.StatusOK, "POST /api/posts works")
	})
	return httpServeMux
}

func main() {
	userHandler := userHandler{}
	postHandler := postHandler{}

	httpServeMux := http.NewServeMux()
	httpServeMux.Handle("/api/posts/", http.StripPrefix("/api/posts", postHandler.Handler()))
	httpServeMux.Handle("/api/users/", http.StripPrefix("/api/users", userHandler.Handler()))

	httpServer := http.Server{
		Handler: httpServeMux,
		Addr:    ":8080",
	}

	httpServerError := httpServer.ListenAndServe()
	log.Fatal(httpServerError)
}
