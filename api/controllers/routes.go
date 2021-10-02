package controllers

import (
	"log"
	"net/http"
)

func (server *Server) InitializeRoutes() {
	server.router.HandleFunc("/uploadFile", server.UploadFile)
	server.router.HandleFunc("uploadPost", server.UploadPost)
	log.Fatal(http.ListenAndServe(":8081", server.router))
}
