package internal

import (
	"net/http"

	"github.com/Xebec19/lms/common/utils"
)

func CreateServer() *http.Server {

	r := GetRoutes()

	return &http.Server{
		Addr:    ":" + utils.GetConfig().Port,
		Handler: r,
	}
}
