package handler

import (
	"net/http"

	"github.com/Xebec19/lms/users-svc/internal/utils"
)

func HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	utils.WriteResponse(w, http.StatusOK, map[string]string{"status": "ok"})
}
