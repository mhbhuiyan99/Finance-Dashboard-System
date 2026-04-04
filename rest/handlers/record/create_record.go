package record

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/mhbhuiyan99/Finance-Dashboard-System/domain"
	"github.com/mhbhuiyan99/Finance-Dashboard-System/rest/middlewares"
	"github.com/mhbhuiyan99/Finance-Dashboard-System/util"
)

type ReqCreateRecord struct {
	Amount   float64           `json:"amount"`
	Type     domain.RecordType `json:"type"`
	Category string            `json:"category"`
	Date     time.Time         `json:"date"`
	Notes    string            `json:"notes"`
}

func (h *Handler) CreateRecord(w http.ResponseWriter, r *http.Request) {

	claims, ok := middlewares.ClaimsFrom(r.Context())
	if !ok {
		util.SendError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var req ReqCreateRecord
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		util.SendError(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if req.Amount <= 0 {
		util.SendError(w, "Amount must be greater than zero", http.StatusBadRequest)
		return
	}
	if req.Type != "income" && req.Type != "expense" {
		util.SendError(w, "Type must be income or expense", http.StatusBadRequest)
		return
	}
	if req.Category == "" {
		util.SendError(w, "Category is required", http.StatusBadRequest)
		return
	}
	if req.Date.IsZero() {
		util.SendError(w, "Date is required", http.StatusBadRequest)
		return
	}

	createdRecord, err := h.svc.Create(domain.FinancialRecord{
		UserID:   claims.UserID,
		Amount:   req.Amount,
		Type:     domain.RecordType(req.Type),
		Category: req.Category,
		Date:     req.Date,
		Notes:    req.Notes,
	})

	if err != nil {
		fmt.Println("Failed to create record:", err)
		util.SendError(w, "Failed to create record", http.StatusInternalServerError)
		return
	}

	util.SendData(w, createdRecord, http.StatusCreated)
}
