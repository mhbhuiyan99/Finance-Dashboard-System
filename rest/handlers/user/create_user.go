package user

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/mhbhuiyan99/Finance-Dashboard-System/domain"

	"github.com/mhbhuiyan99/Finance-Dashboard-System/util"
)

type ReqCreateUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req ReqCreateUser
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		util.SendError(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if req.Name == "" {
		util.SendError(w, "Name is required", http.StatusBadRequest)
		return
	}
	if req.Email == "" {
		util.SendError(w, "Email is required", http.StatusBadRequest)
		return
	}
	if len(req.Password) < 6 {
		util.SendError(w, "Password must be at least 6 characters", http.StatusBadRequest)
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		util.SendError(w, "Failed to process password", http.StatusInternalServerError)
		return
	}

	created, err := h.svc.Create(domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	})

	if err != nil {
		fmt.Println("Failed to register:", err)
		util.SendError(w, "Failed to register user", http.StatusInternalServerError)
		return
	}
	
	util.SendData(w, created, http.StatusCreated)
}
