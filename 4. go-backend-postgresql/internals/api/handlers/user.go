package handlers

import (
	"encoding/json"
	"go-backend-postgresql/internals/models"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserHandler struct {
	DB *pgxpool.Pool
}

func NewUserHandler(db *pgxpool.Pool) *UserHandler {
	return &UserHandler{DB: db}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {

	var req models.CreateUserPayload

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid Payload", http.StatusBadRequest)
		return
	}

	query := `
		INSERT INTO users (fname, lname, email, age) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id, created_at;
	`

	newUser := &models.User{
		FName: req.FName,
		LName: req.LName,
		Email: req.Email,
		Age: req.Age,
	}

	err := h.DB.QueryRow(r.Context(),query, req.FName, req.LName, req.Email, req.Age).Scan(&newUser.Id, &newUser.CreatedAt)
	if err != nil {
		// In production, you would check for unique constraint violations (duplicate email) here
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 Created
	json.NewEncoder(w).Encode(newUser)
}