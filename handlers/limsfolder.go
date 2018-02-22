package handlers

import (
	"net/http"
	"encoding/json"
	"github.com/shvetsovau/gohttpserver/model"
)

func CreateLimsFolderHandler(w http.ResponseWriter, r *http.Request) {
	limsFolder := &model.LimsFolder{}
	if err := json.NewDecoder(r.Body).Decode(&limsFolder); err != nil {
		panic(err)
	}
	//limsFolder = getBillBL(r).Create(limsFolder, getParameter(r, "id"))
	//limsFolder = getBillBL(r).Create(limsFolder)
	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(limsFolder)
}

func UpdateLimsFolderHandler(w http.ResponseWriter, r *http.Request) {
	limsFolder := model.LimsFolder{}

	if err := json.NewDecoder(r.Body).Decode(&limsFolder); err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/vnd.api+json")
	//json.NewEncoder(w).Encode(getBillBL(r).Update(&limsFolder))
}

func DeleteLimsFolderHandler(w http.ResponseWriter, r *http.Request) {
	//getBillBL(r).Delete(getParameter(r, "id"))
	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(true)
}