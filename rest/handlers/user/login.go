package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mhbhuiyan99/Finance-Dashboard-System/util"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) LogIn(w http.ResponseWriter, r *http.Request) {
	var req ReqLogin
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		util.SendError(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if req.Email == "" {
		util.SendError(w, "Email is required", http.StatusBadRequest)
		return
	}

	if req.Password == "" {
		util.SendError(w, "Password is required", http.StatusBadRequest)
		return
	}

	usr, err := h.svc.Find(req.Email)
	if err != nil || usr == nil {
		util.SendError(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	if !usr.IsActive {
		util.SendError(w, "Account is inactive", http.StatusUnauthorized)
		return
	}

	if err = util.CheckPasswordHash(req.Password, usr.Password); err != nil {
		util.SendError(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	token, err := util.CreateJWT(h.cnf.JwtSecretKey, util.Claims{
		UserID:         usr.ID,
		Email:          usr.Email,
		Role:           string(usr.Role),
	})

	if err != nil {
		fmt.Println("Failed to login: ", err)
		util.SendError(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	util.SendData(w, map[string]any{"token": token}, http.StatusOK)
}