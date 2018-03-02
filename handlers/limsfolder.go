package handlers

import (
	"net/http"
	"encoding/json"
	"github.com/shvetsovau/gohttpserver/model"
)

func GetAllLimsFoldersHandler(w http.ResponseWriter, r *http.Request) {

	limsFolderCollection := getLimsFolderBl(r).GetAll()
	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(limsFolderCollection)
}

func GetLimsFolderHandler(w http.ResponseWriter, r *http.Request) {
	limsFolder := getLimsFolderBl(r).GetLimsFolder(getParameter(r, "id"))
	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(limsFolder)
}

func CreateLimsFolderHandler(w http.ResponseWriter, r *http.Request) {
	limsFolder := &model.LimsFolder{}
	//if err := json.NewDecoder(r.Body).Decode(&limsFolder); err != nil {
	//	panic(err)
	//}

	limsFolder = getLimsFolderBl(r).Create(limsFolder)
	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(limsFolder)
}

func UpdateLimsFolderHandler(w http.ResponseWriter, r *http.Request) {
	limsFolder := model.LimsFolder{}

	if err := json.NewDecoder(r.Body).Decode(&limsFolder); err != nil {
		panic(err)
	}

	getLimsFolderBl(r).Update(&limsFolder)

	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(&limsFolder)
}

func DeleteLimsFolderHandler(w http.ResponseWriter, r *http.Request) {
	getLimsFolderBl(r).Delete(getParameter(r, "id"))
	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(true)
}