package handlers

import (
	"net/http"
	"github.com/smbody/gohhtpserver/bl/client"
)

func getLimsFolderBl(r *http.Request) *bl.LimsFolderBL {
	return bl.NewLimsFolder(getUserId(r))
}