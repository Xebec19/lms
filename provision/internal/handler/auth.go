package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Xebec19/lms/common/utils"
	"github.com/Xebec19/lms/provision/internal/models"
)

func HandleSignup(w http.ResponseWriter, r *http.Request) {

	var req models.SignupRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	utils.WriteResponse(w, http.StatusOK, "Signup successful")
}
