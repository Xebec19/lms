package handler

import (
	"net/http"

	"github.com/Xebec19/lms/common/utils"
	"github.com/Xebec19/lms/provision/internal/requests"
)

func HandleSignup(w http.ResponseWriter, r *http.Request) {

	var req requests.SignupRequest

	err := utils.GetValidatedStruct(r, &req)
	if err != nil {
		utils.WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.WriteResponse(w, http.StatusOK, "Signup successful")
}
