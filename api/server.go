package api

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	cloud "cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

type Server struct {
	Ctx             context.Context   //Context
	Router          *mux.Router       //Gorilla mux router
	FireStoreClient *firestore.Client //Firestore client
	Storage         *cloud.Client     //Firebase storage client
}

func (server *Server) InitServer() {
	server.Ctx = context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error fetching enviromental variables: %v\n", err)
	} else {
		fmt.Println("Successfully fetched enviromental variables")
	}

	opt := option.WithServiceAccountFile("C:/Users/Vicki/Downloads/serviceAccountKey.json") //ToDo: Replace with your path to service account key json file

	firebaseApp, err := firebase.NewApp(server.Ctx, nil, opt)
	if err != nil {
		log.Fatalf("Error creating firebase app: %v\n", err)
	}

	server.FireStoreClient, err = firebaseApp.Firestore(server.Ctx)
	if err != nil {
		log.Fatalf("Error creating firestore instance: %v\n", err)
	}

	server.Storage, err = cloud.NewClient(server.Ctx, opt)
	if err != nil {
		log.Fatalf("Error creating GCS client: %v\n", err)
	}

	server.Router = mux.NewRouter()
	server.InitializeRoutes()
}
