package record

import "github.com/mhbhuiyan99/Finance-Dashboard-System/domain"


type Service interface {
	Create(r domain.FinancialRecord) (*domain.FinancialRecord, error)
	GetByID(id string) (*domain.FinancialRecord, error)
	List() ([]domain.FinancialRecord, error)
	Delete(id string) error
	Update(r domain.FinancialRecord) (*domain.FinancialRecord, error)

	Transactions(userID string, filter domain.RecordFilter) ([]domain.FinancialRecord, error)
}