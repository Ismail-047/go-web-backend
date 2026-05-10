-- name: CreateUserInDB :one
INSERT INTO users (fname, lname, email, age, password)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, fname, lname, email, age, created_at;; 