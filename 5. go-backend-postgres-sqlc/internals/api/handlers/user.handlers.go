package handlers

import (
	"go-backend-postgres-sqlc/internals/api/utils/bcrypt"
	"go-backend-postgres-sqlc/internals/api/utils/request"
	"go-backend-postgres-sqlc/internals/api/utils/response"
	"go-backend-postgres-sqlc/internals/api/utils/validator"
	"go-backend-postgres-sqlc/internals/database/repository"
	"go-backend-postgres-sqlc/internals/dtos"
	"net/http"
)

type UserHandler struct {
	Repo *repository.Queries
}

func NewUserHandler(repo *repository.Queries) *UserHandler {
 	return &UserHandler{Repo: repo}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	
	var req dtos.CreateUserRequest

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	if err := request.Decode(r, &req); err != nil {
		response.Error(w, "Invalid request payload", http.StatusBadRequest, err) 
		return
	}

	if err := validator.ValidateStruct(req); err != nil {
		response.Error(w, err.Error(), http.StatusBadRequest, err)
		return
	}

	hashedPassowrd, err := bcrypt.HashPassword(req.Password)
	if err != nil {
		response.Error(w, "Failed to secure password", http.StatusInternalServerError, err)
		return
	}
	user, err := h.Repo.CreateUserInDB(r.Context(), repository.CreateUserInDBParams{
		Fname: req.Fname,
		Lname: req.Lname,
		Email: req.Email,
		Age: req.Age,
		Password: hashedPassowrd,
	})

	if err != nil {
		response.Error(w, "Failed to create user", http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, "User registered successfuly", http.StatusCreated, user)
}

// func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {}