package handler

import (
	"net/http"

	"github.com/Xebec19/lms/common/pkg/bcrypt"
	"github.com/Xebec19/lms/common/utils"
	"github.com/Xebec19/lms/provision/internal/db"
	"github.com/Xebec19/lms/provision/internal/db/models"
	"github.com/Xebec19/lms/provision/internal/db/repository"
	"github.com/Xebec19/lms/provision/internal/requests"
	"gorm.io/gorm"
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

	roleRepo := repository.NewRoleRepository(db)
	role, err := roleRepo.GetRole("user")
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	user := &models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  hash,
		Phone:     req.Phone,
		Roles:     []models.Role{*role},
	}

	if err := userRepo.CreateUser(user); err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteResponse(w, http.StatusOK, "Signup successful", nil)
}

func HandleSignin(w http.ResponseWriter, r *http.Request) {
	var req requests.SigninRequest

	err := utils.GetValidatedStruct(r, &req)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.GetDB()
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	userRepo := repository.NewUserRepo(db)
	user, err := userRepo.GetUserByEmail(req.Email)
	if err == gorm.ErrRecordNotFound {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	if err := bcrypt.CheckPasswordHash(req.Password, user.Password); err != nil {
		utils.WriteErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	db.Preload("Roles.Permissions").First(&user, user.ID)

	utils.WriteResponse(w, http.StatusOK, "Signin successful", nil)
}
