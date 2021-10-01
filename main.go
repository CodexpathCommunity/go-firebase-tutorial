package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"cloud.google.com/go/firestore"
	cloud "cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
)

type App struct {
	ctx             context.Context   //Context
	router          *mux.Router       //Gorilla mux router
	fireStoreClient *firestore.Client //Firestore client
	storage         *cloud.Client     //Firebase storage client
}

func main() {
	app := App{}
	app.Init()
}

func (app *App) Init() {
	app.ctx = context.Background()

	opt := option.WithServiceAccountFile("C:/Users/Vicki/Downloads/serviceAccountKey.json") //ToDo: Replace with your path to service account key json file

	firebaseApp, err := firebase.NewApp(app.ctx, nil, opt)
	if err != nil {
		log.Fatalf("Error creating firebase app: %v\n", err)
	}

	app.fireStoreClient, err = firebaseApp.Firestore(app.ctx)
	if err != nil {
		log.Fatalf("Error creating firestore instance: %v\n", err)
	}

	app.storage, err = cloud.NewClient(app.ctx, opt)
	if err != nil {
		log.Fatalf("Error creating GCS client: %v\n", err)
	}

	app.router = mux.NewRouter()
	app.InitializeRoutes()
}

func (app *App) InitializeRoutes() {
	app.router.HandleFunc("/upload", app.UploadFile)
	log.Fatal(http.ListenAndServe(":8081", app.router))
}

func (app *App) UploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Fatalf("Error parsing file: %v\n", err)
		return
	}
	defer file.Close()

	fileName := strings.Join(strings.Fields(handler.Filename), "")

	bucketName := "your-bucket-url-goes-here" //ToDo: Replace with your bucket url

	writer := app.storage.Bucket(bucketName).Object(fileName).NewWriter(app.ctx)
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
