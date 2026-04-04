package report

import (
	"github.com/mhbhuiyan99/Finance-Dashboard-System/domain"
	rcdHndlr "github.com/mhbhuiyan99/Finance-Dashboard-System/rest/handlers/record"
)

type Service interface {
	rcdHndlr.Service
}

type RecordRepo interface {
	Create(r domain.FinancialRecord) (*domain.FinancialRecord, error)
	GetByID(id string) (*domain.FinancialRecord, error)
	List() ([]domain.FinancialRecord, error)
	Delete(id string) error
	Update(r domain.FinancialRecord) (*domain.FinancialRecord, error)

	Transactions(TID string, filter domain.RecordFilter) ([]domain.FinancialRecord, error)
}