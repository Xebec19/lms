package handler

import (
	"net/http"

	"github.com/Xebec19/lms/common/pkg/bcrypt"
	"github.com/Xebec19/lms/common/utils"
	"github.com/Xebec19/lms/provision/internal/db"
	"github.com/Xebec19/lms/provision/internal/db/models"
	"github.com/Xebec19/lms/provision/internal/db/repository"
	"github.com/Xebec19/lms/provision/internal/requests"
)

func HandleSignup(w http.ResponseWriter, r *http.Request) {

	var req requests.SignupRequest

	err := utils.GetValidatedStruct(r, &req)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	hash, err := bcrypt.HashPassword(req.Password)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	db, err := db.GetDB()
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	userRepo := repository.NewUserRepo(db)
	user := &models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  hash,
		Phone:     req.Phone,
		Role:      req.Role,
	}

	if err := userRepo.CreateUser(user); err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteResponse(w, http.StatusOK, "Signup successful", nil)
}
