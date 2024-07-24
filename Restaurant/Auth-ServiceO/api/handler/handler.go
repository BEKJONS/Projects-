package handler

import (
	"auth-service/storage/postgres"
	"database/sql"
)

type Handler struct {
	Auth postgres.UserRepo
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{*postgres.NewUserRepo(db)}
}
