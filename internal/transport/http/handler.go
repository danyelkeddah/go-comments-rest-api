package http

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Handler struct {
	Router  *mux.Router
	Service CommentService
	Server  *http.Server
}

func NewHandler(service CommentService) *Handler {
	h := &Handler{
		Service: service,
	}

	h.Router = mux.NewRouter()
	h.mapRoutes()
	h.Router.Use(JSONMiddleware)
	h.Router.Use(LoggingMiddleware)
	h.Router.Use(timeoutMiddleware)
	h.Server = &http.Server{
		Addr:    ":8080",
		Handler: h.Router,
	}
	return h
}

func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world")
	})

	h.Router.HandleFunc("/api/v1/comments", h.CreateComment).Methods("POST")
	h.Router.HandleFunc("/api/v1/comments/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/v1/comments/{id}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("/api/v1/comments/{id}", h.DeleteComment).Methods("DELETE")
}

func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	h.Server.Shutdown(ctx)

	log.Println("shutdown gracefully")
	return nil
}
