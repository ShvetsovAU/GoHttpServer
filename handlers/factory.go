package handlers

import (
	"net/http"
	"github.com/shvetsovau/gohttpserver/bl"
)

func getLimsFolderBl(r *http.Request) *bl.LimsFolderBL {
	//return bl.NewLimsFolder(getUserId(r))
	return bl.NewLimsFolder("testUser")
}