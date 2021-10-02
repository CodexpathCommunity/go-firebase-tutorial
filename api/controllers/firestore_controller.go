package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/victorkabata/go-firebase-tutorial/api/models"
)

func (server *Server) UploadPost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	post := models.Post{}
	err = json.Unmarshal(body, &post)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	post.Prepare()
	err = post.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}

	responses.SUCCESS(w, http.StatusOK, nil)
}
