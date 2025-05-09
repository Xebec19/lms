package handler

import (
	"net/http"

	"github.com/Xebec19/lms/common/utils"
)

func HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	utils.WriteResponse(w, http.StatusOK, "Provision is OK", "ok")
}
