package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func (server *Server) UploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Fatalf("Error parsing file: %v\n", err)
		return
	}
	defer file.Close()

	fileName := strings.Join(strings.Fields(handler.Filename), "")

	bucketName := os.Getenv("FIREBASE_BUCKET_NAME") //ToDo: Replace with your bucket url

	writer := server.storage.Bucket(bucketName).Object(fileName).NewWriter(server.ctx)
	defer writer.Close()

	byteSize, err := io.Copy(writer, file)
	if err != nil {
		log.Fatalf("Error coping file: %v\n", err)
		return
	}
	fmt.Printf("File size uploaded: %v\n", byteSize)

	uploadedImageUrl := fmt.Sprintf("https://storage.cloud.google.com/%s/%s", bucketName, fileName)

	fmt.Printf("Image uploaded successfully: %v", uploadedImageUrl)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"imageUrl": uploadedImageUrl})
}
